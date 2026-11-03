#include <QtCore/qflags.h>
#include "gen_qflag.h"

extern "C" {

void* QFlag_new(int value) {
	return new QFlag(value);
}
void* QFlag_new2(long value) {
	return new QFlag(value);
}
void* QFlag_new3(unsigned long value) {
	return new QFlag(value);
}
void* QFlag_new4(unsigned int value) {
	return new QFlag(value);
}
void* QFlag_new5(short value) {
	return new QFlag(value);
}
void* QFlag_new6(unsigned short value) {
	return new QFlag(value);
}
void* QFlag_new7(void* arg1) {
	return new QFlag(*(QFlag*)arg1);
}
int QFlag_ToInt(void* self) {
	auto&& __r = static_cast<int>(*(QFlag*)self);
	return __r;
}
unsigned int QFlag_ToUint(void* self) {
	auto&& __r = static_cast<unsigned int>(*(QFlag*)self);
	return __r;
}
void QFlag_delete(void* self) { delete (QFlag*)self; }

}

#include <QtCore/qanystringview.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstringview.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qchar.h>
#include "gen_qanystringview.h"

extern "C" {

void* QAnyStringView_new() {
	return new QAnyStringView();
}
void* QAnyStringView_new2(void* str) {
	return new QAnyStringView(*(QByteArray*)str);
}
void* QAnyStringView_new3(const char* str, long long str_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	return new QAnyStringView(str_qstr);
}
void* QAnyStringView_new4(void* str) {
	return new QAnyStringView(*(QLatin1String*)str);
}
void* QAnyStringView_new5(void* v) {
	return new QAnyStringView(*(QStringView*)v);
}
void* QAnyStringView_new6(const char* arg1, long long arg1_len) {
	QAnyStringView arg1_qstr = QAnyStringView(arg1, arg1_len);
	return new QAnyStringView(arg1_qstr);
}
void* QAnyStringView_mid(void* self, long long pos) {
	auto&& __r = ((QAnyStringView*)self)->mid(pos);
	return new QAnyStringView(__r);
}
void* QAnyStringView_left(void* self, long long n) {
	auto&& __r = ((QAnyStringView*)self)->left(n);
	return new QAnyStringView(__r);
}
void* QAnyStringView_right(void* self, long long n) {
	auto&& __r = ((QAnyStringView*)self)->right(n);
	return new QAnyStringView(__r);
}
void* QAnyStringView_sliced(void* self, long long pos) {
	auto&& __r = ((QAnyStringView*)self)->sliced(pos);
	return new QAnyStringView(__r);
}
void* QAnyStringView_sliced2(void* self, long long pos, long long n) {
	auto&& __r = ((QAnyStringView*)self)->sliced(pos, n);
	return new QAnyStringView(__r);
}
void* QAnyStringView_first(void* self, long long n) {
	auto&& __r = ((QAnyStringView*)self)->first(n);
	return new QAnyStringView(__r);
}
void* QAnyStringView_last(void* self, long long n) {
	auto&& __r = ((QAnyStringView*)self)->last(n);
	return new QAnyStringView(__r);
}
void* QAnyStringView_chopped(void* self, long long n) {
	auto&& __r = ((QAnyStringView*)self)->chopped(n);
	return new QAnyStringView(__r);
}
void* QAnyStringView_slice(void* self, long long pos) {
	auto&& __r = ((QAnyStringView*)self)->slice(pos);
	return (void*)&(__r);
}
void* QAnyStringView_slice2(void* self, long long pos, long long n) {
	auto&& __r = ((QAnyStringView*)self)->slice(pos, n);
	return (void*)&(__r);
}
void QAnyStringView_truncate(void* self, long long n) {
	((QAnyStringView*)self)->truncate(n);
}
void QAnyStringView_chop(void* self, long long n) {
	((QAnyStringView*)self)->chop(n);
}
char* QAnyStringView_toString(void* self, long long* __out_len) {
	auto&& __r = ((QAnyStringView*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QAnyStringView_size(void* self) {
	auto&& __r = ((QAnyStringView*)self)->size();
	return __r;
}
int QAnyStringView_compare(const char* lhs, long long lhs_len, const char* rhs, long long rhs_len) {
	QAnyStringView lhs_qstr = QAnyStringView(lhs, lhs_len);
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = QAnyStringView::compare(lhs_qstr, rhs_qstr);
	return __r;
}
int QAnyStringView_equal(const char* lhs, long long lhs_len, const char* rhs, long long rhs_len) {
	QAnyStringView lhs_qstr = QAnyStringView(lhs, lhs_len);
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = QAnyStringView::equal(lhs_qstr, rhs_qstr);
	return __r;
}
void* QAnyStringView_front(void* self) {
	auto&& __r = ((QAnyStringView*)self)->front();
	return new QChar(__r);
}
void* QAnyStringView_back(void* self) {
	auto&& __r = ((QAnyStringView*)self)->back();
	return new QChar(__r);
}
int QAnyStringView_empty(void* self) {
	auto&& __r = ((QAnyStringView*)self)->empty();
	return __r;
}
long long QAnyStringView_size_bytes(void* self) {
	auto&& __r = ((QAnyStringView*)self)->size_bytes();
	return __r;
}
long long QAnyStringView_max_size(void* self) {
	auto&& __r = ((QAnyStringView*)self)->max_size();
	return __r;
}
int QAnyStringView_isNull(void* self) {
	auto&& __r = ((QAnyStringView*)self)->isNull();
	return __r;
}
int QAnyStringView_isEmpty(void* self) {
	auto&& __r = ((QAnyStringView*)self)->isEmpty();
	return __r;
}
long long QAnyStringView_length(void* self) {
	auto&& __r = ((QAnyStringView*)self)->length();
	return __r;
}
void* QAnyStringView_OperatorAssign(void* self, const char* arg1, long long arg1_len) {
	QAnyStringView arg1_qstr = QAnyStringView(arg1, arg1_len);
	auto&& __r = ((QAnyStringView*)self)->operator=(arg1_qstr);
	return (void*)&(__r);
}
void* QAnyStringView_mid2(void* self, long long pos, long long n) {
	auto&& __r = ((QAnyStringView*)self)->mid(pos, n);
	return new QAnyStringView(__r);
}
int QAnyStringView_compare2(const char* lhs, long long lhs_len, const char* rhs, long long rhs_len, int cs) {
	QAnyStringView lhs_qstr = QAnyStringView(lhs, lhs_len);
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = QAnyStringView::compare(lhs_qstr, rhs_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QAnyStringView_OperatorEqual(void* self, const char* rhs, long long rhs_len) {
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = (*(QAnyStringView*)self) == (rhs_qstr);
	return __r;
}
int QAnyStringView_OperatorNotEqual(void* self, const char* rhs, long long rhs_len) {
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = (*(QAnyStringView*)self) != (rhs_qstr);
	return __r;
}
int QAnyStringView_OperatorLesser(void* self, const char* rhs, long long rhs_len) {
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = (*(QAnyStringView*)self) < (rhs_qstr);
	return __r;
}
int QAnyStringView_OperatorGreater(void* self, const char* rhs, long long rhs_len) {
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = (*(QAnyStringView*)self) > (rhs_qstr);
	return __r;
}
int QAnyStringView_OperatorLesserOrEqual(void* self, const char* rhs, long long rhs_len) {
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = (*(QAnyStringView*)self) <= (rhs_qstr);
	return __r;
}
int QAnyStringView_OperatorGreaterOrEqual(void* self, const char* rhs, long long rhs_len) {
	QAnyStringView rhs_qstr = QAnyStringView(rhs, rhs_len);
	auto&& __r = (*(QAnyStringView*)self) >= (rhs_qstr);
	return __r;
}
void QAnyStringView_delete(void* self) { delete (QAnyStringView*)self; }

}

#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qnamespace.h>
#include "gen_qdeadlinetimer.h"

extern "C" {

void* QDeadlineTimer_new() {
	return new QDeadlineTimer();
}
void* QDeadlineTimer_new2(int type_) {
	return new QDeadlineTimer((Qt::TimerType)(type_));
}
void* QDeadlineTimer_new3(int arg1) {
	return new QDeadlineTimer((QDeadlineTimer::ForeverConstant)(arg1));
}
void* QDeadlineTimer_new4(long long msecs) {
	return new QDeadlineTimer(msecs);
}
void* QDeadlineTimer_new5(void* arg1) {
	return new QDeadlineTimer(*(QDeadlineTimer*)arg1);
}
void* QDeadlineTimer_new6(int arg1, int type_) {
	return new QDeadlineTimer((QDeadlineTimer::ForeverConstant)(arg1), (Qt::TimerType)(type_));
}
void* QDeadlineTimer_new7(long long msecs, int type_) {
	return new QDeadlineTimer(msecs, (Qt::TimerType)(type_));
}
void QDeadlineTimer_swap(void* self, void* other) {
	((QDeadlineTimer*)self)->swap(*(QDeadlineTimer*)other);
}
int QDeadlineTimer_isForever(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->isForever();
	return __r;
}
int QDeadlineTimer_hasExpired(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->hasExpired();
	return __r;
}
int QDeadlineTimer_timerType(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->timerType();
	return (int)(__r);
}
void QDeadlineTimer_setTimerType(void* self, int type_) {
	((QDeadlineTimer*)self)->setTimerType((Qt::TimerType)(type_));
}
long long QDeadlineTimer_remainingTime(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->remainingTime();
	return __r;
}
long long QDeadlineTimer_remainingTimeNSecs(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->remainingTimeNSecs();
	return __r;
}
void QDeadlineTimer_setRemainingTime(void* self, long long msecs) {
	((QDeadlineTimer*)self)->setRemainingTime(msecs);
}
void QDeadlineTimer_setPreciseRemainingTime(void* self, long long secs) {
	((QDeadlineTimer*)self)->setPreciseRemainingTime(secs);
}
long long QDeadlineTimer_deadline(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->deadline();
	return __r;
}
long long QDeadlineTimer_deadlineNSecs(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->deadlineNSecs();
	return __r;
}
void QDeadlineTimer_setDeadline(void* self, long long msecs) {
	((QDeadlineTimer*)self)->setDeadline(msecs);
}
void QDeadlineTimer_setPreciseDeadline(void* self, long long secs) {
	((QDeadlineTimer*)self)->setPreciseDeadline(secs);
}
void* QDeadlineTimer_addNSecs(void* dt, long long nsecs) {
	auto&& __r = QDeadlineTimer::addNSecs(*(QDeadlineTimer*)dt, nsecs);
	return new QDeadlineTimer(__r);
}
void* QDeadlineTimer_current() {
	auto&& __r = QDeadlineTimer::current();
	return new QDeadlineTimer(__r);
}
void* QDeadlineTimer_OperatorPlusAssign(void* self, long long msecs) {
	auto&& __r = ((QDeadlineTimer*)self)->operator+=(msecs);
	return (void*)&(__r);
}
void* QDeadlineTimer_OperatorMinusAssign(void* self, long long msecs) {
	auto&& __r = ((QDeadlineTimer*)self)->operator-=(msecs);
	return (void*)&(__r);
}
long long QDeadlineTimer_remainingTimeAsDuration(void* self) {
	auto&& __r = ((QDeadlineTimer*)self)->remainingTimeAsDuration();
	return (long long)(__r).count();
}
void QDeadlineTimer_setRemainingTime2(void* self, long long msecs, int type_) {
	((QDeadlineTimer*)self)->setRemainingTime(msecs, (Qt::TimerType)(type_));
}
void QDeadlineTimer_setPreciseRemainingTime2(void* self, long long secs, long long nsecs) {
	((QDeadlineTimer*)self)->setPreciseRemainingTime(secs, nsecs);
}
void QDeadlineTimer_setPreciseRemainingTime3(void* self, long long secs, long long nsecs, int type_) {
	((QDeadlineTimer*)self)->setPreciseRemainingTime(secs, nsecs, (Qt::TimerType)(type_));
}
void QDeadlineTimer_setDeadline2(void* self, long long msecs, int timerType) {
	((QDeadlineTimer*)self)->setDeadline(msecs, (Qt::TimerType)(timerType));
}
void QDeadlineTimer_setPreciseDeadline2(void* self, long long secs, long long nsecs) {
	((QDeadlineTimer*)self)->setPreciseDeadline(secs, nsecs);
}
void QDeadlineTimer_setPreciseDeadline3(void* self, long long secs, long long nsecs, int type_) {
	((QDeadlineTimer*)self)->setPreciseDeadline(secs, nsecs, (Qt::TimerType)(type_));
}
void* QDeadlineTimer_current2(int timerType) {
	auto&& __r = QDeadlineTimer::current((Qt::TimerType)(timerType));
	return new QDeadlineTimer(__r);
}
void* QDeadlineTimer_OperatorPlus(void* self, long long msecs) {
	auto&& __r = (*(QDeadlineTimer*)self) + (msecs);
	return new QDeadlineTimer(__r);
}
void* QDeadlineTimer_OperatorMinus(void* self, long long msecs) {
	auto&& __r = (*(QDeadlineTimer*)self) - (msecs);
	return new QDeadlineTimer(__r);
}
long long QDeadlineTimer_OperatorMinus2(void* self, void* dt2) {
	auto&& __r = (*(QDeadlineTimer*)self) - (*(QDeadlineTimer*)dt2);
	return __r;
}
int QDeadlineTimer_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QDeadlineTimer*)self) == (*(QDeadlineTimer*)rhs);
	return __r;
}
int QDeadlineTimer_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QDeadlineTimer*)self) != (*(QDeadlineTimer*)rhs);
	return __r;
}
int QDeadlineTimer_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QDeadlineTimer*)self) < (*(QDeadlineTimer*)rhs);
	return __r;
}
int QDeadlineTimer_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QDeadlineTimer*)self) > (*(QDeadlineTimer*)rhs);
	return __r;
}
int QDeadlineTimer_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QDeadlineTimer*)self) <= (*(QDeadlineTimer*)rhs);
	return __r;
}
int QDeadlineTimer_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QDeadlineTimer*)self) >= (*(QDeadlineTimer*)rhs);
	return __r;
}
void QDeadlineTimer_delete(void* self) { delete (QDeadlineTimer*)self; }

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
#include "gen_qabstracttablemodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractTableModel__Go : public QAbstractTableModel {
public:
	long long __goId;
	unsigned char __goOn[40] = {};
	template<typename... A> QAbstractTableModel__Go(long long id, A&&... a) : QAbstractTableModel(std::forward<A>(a)...), __goId(id) {}
	using QAbstractTableModel::createIndex;
	using QAbstractTableModel::encodeData;
	using QAbstractTableModel::decodeData;
	using QAbstractTableModel::beginInsertRows;
	using QAbstractTableModel::endInsertRows;
	using QAbstractTableModel::beginRemoveRows;
	using QAbstractTableModel::endRemoveRows;
	using QAbstractTableModel::beginMoveRows;
	using QAbstractTableModel::endMoveRows;
	using QAbstractTableModel::beginInsertColumns;
	using QAbstractTableModel::endInsertColumns;
	using QAbstractTableModel::beginRemoveColumns;
	using QAbstractTableModel::endRemoveColumns;
	using QAbstractTableModel::beginMoveColumns;
	using QAbstractTableModel::endMoveColumns;
	using QAbstractTableModel::beginResetModel;
	using QAbstractTableModel::endResetModel;
	using QAbstractTableModel::changePersistentIndex;
	using QAbstractTableModel::changePersistentIndexList;
	using QAbstractTableModel::persistentIndexList;
	using QAbstractTableModel::sender;
	using QAbstractTableModel::senderSignalIndex;
	using QAbstractTableModel::receivers;
	using QAbstractTableModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QAbstractTableModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractTableModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractTableModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractTableModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractTableModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractTableModel::disconnectNotify(__a0); }
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return int{};
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return int{};
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractTableModel::metaObject();
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractTableModel::index(row, column, parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractTableModel::sibling(row, column, idx);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::dropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QAbstractTableModel::flags(index);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractTableModel::setData(index, value, role);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractTableModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractTableModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QAbstractTableModel::clearItemData(index);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_voidp) {
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
		return QAbstractTableModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[13] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QAbstractTableModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::canDropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractTableModel::supportedDropActions();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractTableModel::supportedDragActions();
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::removeColumns(column, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractTableModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractTableModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QAbstractTableModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractTableModel::canFetchMore(parent);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QAbstractTableModel::sort(column, order);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractTableModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
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
		return QAbstractTableModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractTableModel::span(index);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QAbstractTableModel::multiData(index, roleDataSpan);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QAbstractTableModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QAbstractTableModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QAbstractTableModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractTableModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractTableModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractTableModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractTableModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractTableModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractTableModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractTableModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractTableModel;
static std::mutex __connsMu_QAbstractTableModel;

extern "C" {

void* QAbstractTableModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QAbstractTableModel*)self); }
void QAbstractTableModel_chain(void* self, void** out) {
	QAbstractTableModel* __p = (QAbstractTableModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAbstractTableModel__set_override(void* self, int idx) { ((QAbstractTableModel__Go*)self)->__goOn[idx] = 1; }
void* QAbstractTableModel_fromQObject(void* o) { return (void*)qobject_cast<QAbstractTableModel*>((QObject*)o); }
char* QAbstractTableModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractTableModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractTableModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QAbstractTableModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QAbstractTableModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QAbstractTableModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QAbstractTableModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QAbstractTableModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QAbstractTableModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QAbstractTableModel_endInsertRows(void* self) {
	((QAbstractTableModel__Go*)self)->endInsertRows();
}
void QAbstractTableModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QAbstractTableModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QAbstractTableModel_endRemoveRows(void* self) {
	((QAbstractTableModel__Go*)self)->endRemoveRows();
}
int QAbstractTableModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QAbstractTableModel_endMoveRows(void* self) {
	((QAbstractTableModel__Go*)self)->endMoveRows();
}
void QAbstractTableModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QAbstractTableModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractTableModel_endInsertColumns(void* self) {
	((QAbstractTableModel__Go*)self)->endInsertColumns();
}
void QAbstractTableModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QAbstractTableModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractTableModel_endRemoveColumns(void* self) {
	((QAbstractTableModel__Go*)self)->endRemoveColumns();
}
int QAbstractTableModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QAbstractTableModel_endMoveColumns(void* self) {
	((QAbstractTableModel__Go*)self)->endMoveColumns();
}
void QAbstractTableModel_beginResetModel(void* self) {
	((QAbstractTableModel__Go*)self)->beginResetModel();
}
void QAbstractTableModel_endResetModel(void* self) {
	((QAbstractTableModel__Go*)self)->endResetModel();
}
void QAbstractTableModel_changePersistentIndex(void* self, void* from, void* to) {
	((QAbstractTableModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QAbstractTableModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QAbstractTableModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QAbstractTableModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractTableModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractTableModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractTableModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractTableModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractTableModel_sender(void* self) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractTableModel_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractTableModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractTableModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractTableModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractTableModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractTableModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractTableModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractTableModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractTableModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractTableModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractTableModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractTableModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractTableModel_metaObject_default(void* self) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::metaObject();
	return (void*)(__r);
}
void* QAbstractTableModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QAbstractTableModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QAbstractTableModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractTableModel_flags_default(void* self, void* index) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QAbstractTableModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
void* QAbstractTableModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QAbstractTableModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QAbstractTableModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
char** QAbstractTableModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QAbstractTableModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QAbstractTableModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractTableModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::supportedDropActions();
	return (int)(__r);
}
int QAbstractTableModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::supportedDragActions();
	return (int)(__r);
}
int QAbstractTableModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractTableModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractTableModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractTableModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractTableModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractTableModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QAbstractTableModel_fetchMore_default(void* self, void* parent) {
	((QAbstractTableModel*)self)->QAbstractTableModel::fetchMore(*(QModelIndex*)parent);
}
int QAbstractTableModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QAbstractTableModel_sort_default(void* self, int column, int order) {
	((QAbstractTableModel*)self)->QAbstractTableModel::sort(column, (Qt::SortOrder)(order));
}
void* QAbstractTableModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QAbstractTableModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractTableModel_span_default(void* self, void* index) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void QAbstractTableModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QAbstractTableModel*)self)->QAbstractTableModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QAbstractTableModel_submit_default(void* self) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::submit();
	return __r;
}
void QAbstractTableModel_revert_default(void* self) {
	((QAbstractTableModel*)self)->QAbstractTableModel::revert();
}
void QAbstractTableModel_resetInternalData_default(void* self) {
	((QAbstractTableModel__Go*)self)->cppgoBase_resetInternalData();
}
int QAbstractTableModel_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::event((QEvent*)event);
	return __r;
}
int QAbstractTableModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractTableModel*)self)->QAbstractTableModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractTableModel_timerEvent_default(void* self, void* event) {
	((QAbstractTableModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractTableModel_childEvent_default(void* self, void* event) {
	((QAbstractTableModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractTableModel_customEvent_default(void* self, void* event) {
	((QAbstractTableModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractTableModel_connectNotify_default(void* self, void* signal) {
	((QAbstractTableModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractTableModel_disconnectNotify_default(void* self, void* signal) {
	((QAbstractTableModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractTableModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::headerDataChanged, (QAbstractTableModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::layoutChanged, (QAbstractTableModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::layoutChanged, (QAbstractTableModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::layoutAboutToBeChanged, (QAbstractTableModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::layoutAboutToBeChanged, (QAbstractTableModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::rowsAboutToBeInserted, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::rowsInserted, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::rowsAboutToBeRemoved, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::rowsRemoved, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::columnsAboutToBeInserted, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::columnsInserted, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::columnsAboutToBeRemoved, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::columnsRemoved, (QAbstractTableModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::modelAboutToBeReset, (QAbstractTableModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::modelReset, (QAbstractTableModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::rowsAboutToBeMoved, (QAbstractTableModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::rowsMoved, (QAbstractTableModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::columnsAboutToBeMoved, (QAbstractTableModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::columnsMoved, (QAbstractTableModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::destroyed, (QAbstractTableModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::destroyed, (QAbstractTableModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTableModel*)self, &QAbstractTableModel::objectNameChanged, (QAbstractTableModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
	__conns_QAbstractTableModel[__id] = __c;
}
void QAbstractTableModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTableModel);
		auto __it = __conns_QAbstractTableModel.find(id);
		if (__it == __conns_QAbstractTableModel.end()) return;
		__c = __it->second;
		__conns_QAbstractTableModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractTableModel__Go_new(long long id, void* parent) { return new QAbstractTableModel__Go(id, (QObject*)parent); }
void QAbstractTableModel_delete(void* self) { delete (QAbstractTableModel*)self; }

}

#include <QtCore/qcalendar.h>
#include "gen_qcalendar__yearmonthday.h"

extern "C" {

void* QCalendar__YearMonthDay_new() {
	return new QCalendar::YearMonthDay();
}
void* QCalendar__YearMonthDay_new2(int y) {
	return new QCalendar::YearMonthDay(y);
}
void* QCalendar__YearMonthDay_new3(int y, int m) {
	return new QCalendar::YearMonthDay(y, m);
}
void* QCalendar__YearMonthDay_new4(int y, int m, int d) {
	return new QCalendar::YearMonthDay(y, m, d);
}
int QCalendar__YearMonthDay_isValid(void* self) {
	auto&& __r = ((QCalendar::YearMonthDay*)self)->isValid();
	return __r;
}
int QCalendar__YearMonthDay_year_get(void* self) {
	auto __r = ((QCalendar::YearMonthDay*)self)->year;
	return __r;
}
void QCalendar__YearMonthDay_year_set(void* self, int year) {
	((QCalendar::YearMonthDay*)self)->year = year;
}
int QCalendar__YearMonthDay_month_get(void* self) {
	auto __r = ((QCalendar::YearMonthDay*)self)->month;
	return __r;
}
void QCalendar__YearMonthDay_month_set(void* self, int month) {
	((QCalendar::YearMonthDay*)self)->month = month;
}
int QCalendar__YearMonthDay_day_get(void* self) {
	auto __r = ((QCalendar::YearMonthDay*)self)->day;
	return __r;
}
void QCalendar__YearMonthDay_day_set(void* self, int day) {
	((QCalendar::YearMonthDay*)self)->day = day;
}
void QCalendar__YearMonthDay_delete(void* self) { delete (QCalendar::YearMonthDay*)self; }

}

#include <QtCore/qproperty.h>
#include "gen_qpropertybindingsourcelocation.h"

extern "C" {

void* QPropertyBindingSourceLocation_new() {
	return new QPropertyBindingSourceLocation();
}
void* QPropertyBindingSourceLocation_new2(void* arg1) {
	return new QPropertyBindingSourceLocation(*(QPropertyBindingSourceLocation*)arg1);
}
const char* QPropertyBindingSourceLocation_fileName_get(void* self) {
	auto __r = ((QPropertyBindingSourceLocation*)self)->fileName;
	return __r;
}
void QPropertyBindingSourceLocation_fileName_set(void* self, const char* fileName) {
	((QPropertyBindingSourceLocation*)self)->fileName = fileName;
}
const char* QPropertyBindingSourceLocation_functionName_get(void* self) {
	auto __r = ((QPropertyBindingSourceLocation*)self)->functionName;
	return __r;
}
void QPropertyBindingSourceLocation_functionName_set(void* self, const char* functionName) {
	((QPropertyBindingSourceLocation*)self)->functionName = functionName;
}
unsigned int QPropertyBindingSourceLocation_line_get(void* self) {
	auto __r = ((QPropertyBindingSourceLocation*)self)->line;
	return __r;
}
void QPropertyBindingSourceLocation_line_set(void* self, unsigned int line) {
	((QPropertyBindingSourceLocation*)self)->line = line;
}
unsigned int QPropertyBindingSourceLocation_column_get(void* self) {
	auto __r = ((QPropertyBindingSourceLocation*)self)->column;
	return __r;
}
void QPropertyBindingSourceLocation_column_set(void* self, unsigned int column) {
	((QPropertyBindingSourceLocation*)self)->column = column;
}
void QPropertyBindingSourceLocation_delete(void* self) { delete (QPropertyBindingSourceLocation*)self; }

}

#include <QtCore/qdir.h>
#include <QtCore/qstring.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qfiledevice.h>
#include <QtCore/qchar.h>
#include "gen_qdir.h"

extern "C" {

void* QDir_new(void* arg1) {
	return new QDir(*(QDir*)arg1);
}
void* QDir_new2() {
	return new QDir();
}
void* QDir_new3(const char* path, long long path_len, const char* nameFilter, long long nameFilter_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QString nameFilter_qstr = QString::fromUtf8(nameFilter, nameFilter_len);
	return new QDir(path_qstr, nameFilter_qstr);
}
void* QDir_new4(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDir(path_qstr);
}
void* QDir_new5(const char* path, long long path_len, const char* nameFilter, long long nameFilter_len, int sort) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QString nameFilter_qstr = QString::fromUtf8(nameFilter, nameFilter_len);
	return new QDir(path_qstr, nameFilter_qstr, QFlags<QDir::SortFlag>(sort));
}
void* QDir_new6(const char* path, long long path_len, const char* nameFilter, long long nameFilter_len, int sort, int filter) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QString nameFilter_qstr = QString::fromUtf8(nameFilter, nameFilter_len);
	return new QDir(path_qstr, nameFilter_qstr, QFlags<QDir::SortFlag>(sort), QFlags<QDir::Filter>(filter));
}
void* QDir_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QDir*)self)->operator=(*(QDir*)arg1);
	return (void*)&(__r);
}
void QDir_swap(void* self, void* other) {
	((QDir*)self)->swap(*(QDir*)other);
}
void QDir_setPath(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	((QDir*)self)->setPath(path_qstr);
}
char* QDir_path(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->path();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_absolutePath(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->absolutePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_canonicalPath(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->canonicalPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDir_setSearchPaths(const char* prefix, long long prefix_len, char** searchPaths, long long searchPaths_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	QList<QString> searchPaths_list; searchPaths_list.reserve(searchPaths_len);
	for (long long __i = 0; __i < searchPaths_len; __i++) searchPaths_list.append(QString::fromUtf8(searchPaths[__i]));
	QDir::setSearchPaths(prefix_qstr, searchPaths_list);
}
void QDir_addSearchPath(const char* prefix, long long prefix_len, const char* path, long long path_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	QString path_qstr = QString::fromUtf8(path, path_len);
	QDir::addSearchPath(prefix_qstr, path_qstr);
}
char** QDir_searchPaths(const char* prefix, long long prefix_len, long long* __out_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	auto&& __r = QDir::searchPaths(prefix_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QDir_dirName(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->dirName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_filePath(void* self, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QDir*)self)->filePath(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_absoluteFilePath(void* self, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QDir*)self)->absoluteFilePath(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_relativeFilePath(void* self, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QDir*)self)->relativeFilePath(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_toNativeSeparators(const char* pathName, long long pathName_len, long long* __out_len) {
	QString pathName_qstr = QString::fromUtf8(pathName, pathName_len);
	auto&& __r = QDir::toNativeSeparators(pathName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDir_fromNativeSeparators(const char* pathName, long long pathName_len, long long* __out_len) {
	QString pathName_qstr = QString::fromUtf8(pathName, pathName_len);
	auto&& __r = QDir::fromNativeSeparators(pathName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDir_cd(void* self, const char* dirName, long long dirName_len) {
	QString dirName_qstr = QString::fromUtf8(dirName, dirName_len);
	auto&& __r = ((QDir*)self)->cd(dirName_qstr);
	return __r;
}
int QDir_cdUp(void* self) {
	auto&& __r = ((QDir*)self)->cdUp();
	return __r;
}
char** QDir_nameFilters(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->nameFilters();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QDir_setNameFilters(void* self, char** nameFilters, long long nameFilters_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	((QDir*)self)->setNameFilters(nameFilters_list);
}
int QDir_filter(void* self) {
	auto&& __r = ((QDir*)self)->filter();
	return (int)(__r);
}
void QDir_setFilter(void* self, int filter) {
	((QDir*)self)->setFilter(QFlags<QDir::Filter>(filter));
}
int QDir_sorting(void* self) {
	auto&& __r = ((QDir*)self)->sorting();
	return (int)(__r);
}
void QDir_setSorting(void* self, int sort) {
	((QDir*)self)->setSorting(QFlags<QDir::SortFlag>(sort));
}
long long QDir_count(void* self) {
	auto&& __r = ((QDir*)self)->count();
	return __r;
}
int QDir_isEmpty(void* self) {
	auto&& __r = ((QDir*)self)->isEmpty();
	return __r;
}
char* QDir_OperatorSubscript(void* self, long long arg1, long long* __out_len) {
	auto&& __r = ((QDir*)self)->operator[](arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QDir_nameFiltersFromString(const char* nameFilter, long long nameFilter_len, long long* __out_len) {
	QString nameFilter_qstr = QString::fromUtf8(nameFilter, nameFilter_len);
	auto&& __r = QDir::nameFiltersFromString(nameFilter_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QDir_entryList(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->entryList();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QDir_entryList2(void* self, char** nameFilters, long long nameFilters_len, long long* __out_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	auto&& __r = ((QDir*)self)->entryList(nameFilters_list);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QDir_entryInfoList(void* self, long long* __out_len) {
	auto&& __r = ((QDir*)self)->entryInfoList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
void** QDir_entryInfoList2(void* self, char** nameFilters, long long nameFilters_len, long long* __out_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	auto&& __r = ((QDir*)self)->entryInfoList(nameFilters_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
int QDir_mkdir(void* self, const char* dirName, long long dirName_len) {
	QString dirName_qstr = QString::fromUtf8(dirName, dirName_len);
	auto&& __r = ((QDir*)self)->mkdir(dirName_qstr);
	return __r;
}
int QDir_rmdir(void* self, const char* dirName, long long dirName_len) {
	QString dirName_qstr = QString::fromUtf8(dirName, dirName_len);
	auto&& __r = ((QDir*)self)->rmdir(dirName_qstr);
	return __r;
}
int QDir_mkpath(void* self, const char* dirPath, long long dirPath_len) {
	QString dirPath_qstr = QString::fromUtf8(dirPath, dirPath_len);
	auto&& __r = ((QDir*)self)->mkpath(dirPath_qstr);
	return __r;
}
int QDir_rmpath(void* self, const char* dirPath, long long dirPath_len) {
	QString dirPath_qstr = QString::fromUtf8(dirPath, dirPath_len);
	auto&& __r = ((QDir*)self)->rmpath(dirPath_qstr);
	return __r;
}
int QDir_removeRecursively(void* self) {
	auto&& __r = ((QDir*)self)->removeRecursively();
	return __r;
}
int QDir_isReadable(void* self) {
	auto&& __r = ((QDir*)self)->isReadable();
	return __r;
}
int QDir_exists(void* self) {
	auto&& __r = ((QDir*)self)->exists();
	return __r;
}
int QDir_isRoot(void* self) {
	auto&& __r = ((QDir*)self)->isRoot();
	return __r;
}
int QDir_isRelativePath(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = QDir::isRelativePath(path_qstr);
	return __r;
}
int QDir_isAbsolutePath(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = QDir::isAbsolutePath(path_qstr);
	return __r;
}
int QDir_isRelative(void* self) {
	auto&& __r = ((QDir*)self)->isRelative();
	return __r;
}
int QDir_isAbsolute(void* self) {
	auto&& __r = ((QDir*)self)->isAbsolute();
	return __r;
}
int QDir_makeAbsolute(void* self) {
	auto&& __r = ((QDir*)self)->makeAbsolute();
	return __r;
}
int QDir_remove(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QDir*)self)->remove(fileName_qstr);
	return __r;
}
int QDir_rename(void* self, const char* oldName, long long oldName_len, const char* newName, long long newName_len) {
	QString oldName_qstr = QString::fromUtf8(oldName, oldName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = ((QDir*)self)->rename(oldName_qstr, newName_qstr);
	return __r;
}
int QDir_exists2(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QDir*)self)->exists(name_qstr);
	return __r;
}
void** QDir_drives(long long* __out_len) {
	auto&& __r = QDir::drives();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
void* QDir_listSeparator() {
	auto&& __r = QDir::listSeparator();
	return new QChar(__r);
}
void* QDir_separator() {
	auto&& __r = QDir::separator();
	return new QChar(__r);
}
int QDir_setCurrent(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = QDir::setCurrent(path_qstr);
	return __r;
}
void* QDir_current() {
	auto&& __r = QDir::current();
	return new QDir(__r);
}
char* QDir_currentPath(long long* __out_len) {
	auto&& __r = QDir::currentPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDir_home() {
	auto&& __r = QDir::home();
	return new QDir(__r);
}
char* QDir_homePath(long long* __out_len) {
	auto&& __r = QDir::homePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDir_root() {
	auto&& __r = QDir::root();
	return new QDir(__r);
}
char* QDir_rootPath(long long* __out_len) {
	auto&& __r = QDir::rootPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDir_temp() {
	auto&& __r = QDir::temp();
	return new QDir(__r);
}
char* QDir_tempPath(long long* __out_len) {
	auto&& __r = QDir::tempPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDir_match(char** filters, long long filters_len, const char* fileName, long long fileName_len) {
	QList<QString> filters_list; filters_list.reserve(filters_len);
	for (long long __i = 0; __i < filters_len; __i++) filters_list.append(QString::fromUtf8(filters[__i]));
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QDir::match(filters_list, fileName_qstr);
	return __r;
}
int QDir_match2(const char* filter, long long filter_len, const char* fileName, long long fileName_len) {
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QDir::match(filter_qstr, fileName_qstr);
	return __r;
}
char* QDir_cleanPath(const char* path, long long path_len, long long* __out_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = QDir::cleanPath(path_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDir_refresh(void* self) {
	((QDir*)self)->refresh();
}
int QDir_isEmpty2(void* self, int filters) {
	auto&& __r = ((QDir*)self)->isEmpty(QFlags<QDir::Filter>(filters));
	return __r;
}
char** QDir_entryList3(void* self, int filters, long long* __out_len) {
	auto&& __r = ((QDir*)self)->entryList(QFlags<QDir::Filter>(filters));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QDir_entryList4(void* self, int filters, int sort, long long* __out_len) {
	auto&& __r = ((QDir*)self)->entryList(QFlags<QDir::Filter>(filters), QFlags<QDir::SortFlag>(sort));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QDir_entryList5(void* self, char** nameFilters, long long nameFilters_len, int filters, long long* __out_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	auto&& __r = ((QDir*)self)->entryList(nameFilters_list, QFlags<QDir::Filter>(filters));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QDir_entryList6(void* self, char** nameFilters, long long nameFilters_len, int filters, int sort, long long* __out_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	auto&& __r = ((QDir*)self)->entryList(nameFilters_list, QFlags<QDir::Filter>(filters), QFlags<QDir::SortFlag>(sort));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QDir_entryInfoList3(void* self, int filters, long long* __out_len) {
	auto&& __r = ((QDir*)self)->entryInfoList(QFlags<QDir::Filter>(filters));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
void** QDir_entryInfoList4(void* self, int filters, int sort, long long* __out_len) {
	auto&& __r = ((QDir*)self)->entryInfoList(QFlags<QDir::Filter>(filters), QFlags<QDir::SortFlag>(sort));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
void** QDir_entryInfoList5(void* self, char** nameFilters, long long nameFilters_len, int filters, long long* __out_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	auto&& __r = ((QDir*)self)->entryInfoList(nameFilters_list, QFlags<QDir::Filter>(filters));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
void** QDir_entryInfoList6(void* self, char** nameFilters, long long nameFilters_len, int filters, int sort, long long* __out_len) {
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	auto&& __r = ((QDir*)self)->entryInfoList(nameFilters_list, QFlags<QDir::Filter>(filters), QFlags<QDir::SortFlag>(sort));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFileInfo(__e);
	return __arr;
}
int QDir_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QDir*)self) == (*(QDir*)rhs);
	return __r;
}
int QDir_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QDir*)self) != (*(QDir*)rhs);
	return __r;
}
void QDir_delete(void* self) { delete (QDir*)self; }

}

#include <QtCore/qjsonobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qanystringview.h>
#include "gen_qjsonobject.h"

extern "C" {

void* QJsonObject_new() {
	return new QJsonObject();
}
void* QJsonObject_new2(void* other) {
	return new QJsonObject(*(QJsonObject*)other);
}
void* QJsonObject_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QJsonObject*)self)->operator=(*(QJsonObject*)other);
	return (void*)&(__r);
}
void QJsonObject_swap(void* self, void* other) {
	((QJsonObject*)self)->swap(*(QJsonObject*)other);
}
void* QJsonObject_fromVariantMap(char** map__keys, void* map__vals, long long map__len) {
	QMap<QString, QVariant> map__map;
	void** map__vv = (void**)map__vals;
	for (long long __i = 0; __i < map__len; __i++) map__map.insert(QString::fromUtf8(map__keys[__i]), *(QVariant*)map__vv[__i]);
	auto&& __r = QJsonObject::fromVariantMap(map__map);
	return new QJsonObject(__r);
}
void* QJsonObject_fromVariantHash(char** map__keys, void* map__vals, long long map__len) {
	QHash<QString, QVariant> map__map;
	void** map__vv = (void**)map__vals;
	for (long long __i = 0; __i < map__len; __i++) map__map.insert(QString::fromUtf8(map__keys[__i]), *(QVariant*)map__vv[__i]);
	auto&& __r = QJsonObject::fromVariantHash(map__map);
	return new QJsonObject(__r);
}
char** QJsonObject_keys(void* self, long long* __out_len) {
	auto&& __r = ((QJsonObject*)self)->keys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
long long QJsonObject_size(void* self) {
	auto&& __r = ((QJsonObject*)self)->size();
	return __r;
}
long long QJsonObject_count(void* self) {
	auto&& __r = ((QJsonObject*)self)->count();
	return __r;
}
long long QJsonObject_length(void* self) {
	auto&& __r = ((QJsonObject*)self)->length();
	return __r;
}
int QJsonObject_isEmpty(void* self) {
	auto&& __r = ((QJsonObject*)self)->isEmpty();
	return __r;
}
void* QJsonObject_value(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonObject*)self)->value(key_qstr);
	return new QJsonValue(__r);
}
void* QJsonObject_OperatorSubscript(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonObject*)self)->operator[](key_qstr);
	return new QJsonValue(__r);
}
void* QJsonObject_OperatorSubscript2(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonObject*)self)->operator[](key_qstr);
	return new QJsonValueRef(__r);
}
void* QJsonObject_value2(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->value(*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonObject_value3(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->value(*(QLatin1String*)key);
	return new QJsonValue(__r);
}
void* QJsonObject_OperatorSubscript3(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->operator[](*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonObject_OperatorSubscript4(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->operator[](*(QLatin1String*)key);
	return new QJsonValue(__r);
}
void* QJsonObject_OperatorSubscript5(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->operator[](*(QStringView*)key);
	return new QJsonValueRef(__r);
}
void* QJsonObject_OperatorSubscript6(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->operator[](*(QLatin1String*)key);
	return new QJsonValueRef(__r);
}
void QJsonObject_remove(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QJsonObject*)self)->remove(key_qstr);
}
void* QJsonObject_take(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonObject*)self)->take(key_qstr);
	return new QJsonValue(__r);
}
int QJsonObject_contains(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonObject*)self)->contains(key_qstr);
	return __r;
}
void QJsonObject_remove2(void* self, void* key) {
	((QJsonObject*)self)->remove(*(QStringView*)key);
}
void QJsonObject_remove3(void* self, void* key) {
	((QJsonObject*)self)->remove(*(QLatin1String*)key);
}
void* QJsonObject_take2(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->take(*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonObject_take3(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->take(*(QLatin1String*)key);
	return new QJsonValue(__r);
}
int QJsonObject_contains2(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->contains(*(QStringView*)key);
	return __r;
}
int QJsonObject_contains3(void* self, void* key) {
	auto&& __r = ((QJsonObject*)self)->contains(*(QLatin1String*)key);
	return __r;
}
int QJsonObject_empty(void* self) {
	auto&& __r = ((QJsonObject*)self)->empty();
	return __r;
}
int QJsonObject_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonObject*)self) == (*(QJsonObject*)rhs);
	return __r;
}
int QJsonObject_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonObject*)self) != (*(QJsonObject*)rhs);
	return __r;
}
int QJsonObject_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonObject*)self) == (*(QJsonValue*)rhs);
	return __r;
}
int QJsonObject_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonObject*)self) != (*(QJsonValue*)rhs);
	return __r;
}
int QJsonObject_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QJsonObject*)self) == (*(QJsonValueConstRef*)rhs);
	return __r;
}
int QJsonObject_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QJsonObject*)self) != (*(QJsonValueConstRef*)rhs);
	return __r;
}
void QJsonObject_delete(void* self) { delete (QJsonObject*)self; }

}

#include <QtCore/qoperatingsystemversion.h>
#include <QtCore/qstring.h>
#include <QtCore/qversionnumber.h>
#include "gen_qoperatingsystemversionunexported.h"

extern "C" {

void* QOperatingSystemVersionUnexported_base(void* self) { return (void*)static_cast<QOperatingSystemVersionBase*>((QOperatingSystemVersionUnexported*)self); }
void* QOperatingSystemVersionUnexported_new(void* other) {
	return new QOperatingSystemVersionUnexported(*(QOperatingSystemVersionBase*)other);
}
void* QOperatingSystemVersionUnexported_new2() {
	return new QOperatingSystemVersionUnexported();
}
void* QOperatingSystemVersionUnexported_new3(void* arg1) {
	return new QOperatingSystemVersionUnexported(*(QOperatingSystemVersionUnexported*)arg1);
}
void* QOperatingSystemVersionUnexported_new4(int arg1, int arg2, int arg3, int arg4) {
	return new QOperatingSystemVersionUnexported((QOperatingSystemVersionBase::OSType)(arg1), arg2, arg3, arg4);
}
void* QOperatingSystemVersionUnexported_current() {
	auto&& __r = QOperatingSystemVersionUnexported::current();
	return new QOperatingSystemVersionBase(__r);
}
char* QOperatingSystemVersionUnexported_name(void* osversion, long long* __out_len) {
	auto&& __r = QOperatingSystemVersionUnexported::name(*(QOperatingSystemVersionBase*)osversion);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QOperatingSystemVersionUnexported_currentType() {
	auto&& __r = QOperatingSystemVersionUnexported::currentType();
	return (int)(__r);
}
void QOperatingSystemVersionUnexported_delete(void* self) { delete (QOperatingSystemVersionUnexported*)self; }

}

#include <QtCore/qrangemodel.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qrangemodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QRangeModel;
static std::mutex __connsMu_QRangeModel;

extern "C" {

void* QRangeModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QRangeModel*)self); }
void QRangeModel_chain(void* self, void** out) {
	QRangeModel* __p = (QRangeModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QRangeModel_fromQObject(void* o) { return (void*)qobject_cast<QRangeModel*>((QObject*)o); }
char* QRangeModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QRangeModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRangeModel_index(void* self, int row, int column) {
	auto&& __r = ((QRangeModel*)self)->index(row, column);
	return new QModelIndex(__r);
}
void* QRangeModel_parent(void* self, void* child) {
	auto&& __r = ((QRangeModel*)self)->parent(*(QModelIndex*)child);
	return new QModelIndex(__r);
}
void* QRangeModel_sibling(void* self, int row, int column, void* index) {
	auto&& __r = ((QRangeModel*)self)->sibling(row, column, *(QModelIndex*)index);
	return new QModelIndex(__r);
}
int QRangeModel_rowCount(void* self) {
	auto&& __r = ((QRangeModel*)self)->rowCount();
	return __r;
}
int QRangeModel_columnCount(void* self) {
	auto&& __r = ((QRangeModel*)self)->columnCount();
	return __r;
}
int QRangeModel_insertColumns(void* self, int column, int count) {
	auto&& __r = ((QRangeModel*)self)->insertColumns(column, count);
	return __r;
}
int QRangeModel_removeColumns(void* self, int column, int count) {
	auto&& __r = ((QRangeModel*)self)->removeColumns(column, count);
	return __r;
}
int QRangeModel_moveColumns(void* self, void* sourceParent, int sourceColumn, int count, void* destParent, int destColumn) {
	auto&& __r = ((QRangeModel*)self)->moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destParent, destColumn);
	return __r;
}
int QRangeModel_insertRows(void* self, int row, int count) {
	auto&& __r = ((QRangeModel*)self)->insertRows(row, count);
	return __r;
}
int QRangeModel_removeRows(void* self, int row, int count) {
	auto&& __r = ((QRangeModel*)self)->removeRows(row, count);
	return __r;
}
int QRangeModel_moveRows(void* self, void* sourceParent, int sourceRow, int count, void* destParent, int destRow) {
	auto&& __r = ((QRangeModel*)self)->moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destParent, destRow);
	return __r;
}
void QRangeModel_setRoleNames(void* self, int* names_keys, void* names_vals, long long names_len) {
	QHash<int, QByteArray> names_map;
	void** names_vv = (void**)names_vals;
	for (long long __i = 0; __i < names_len; __i++) names_map.insert(names_keys[__i], *(QByteArray*)names_vv[__i]);
	((QRangeModel*)self)->setRoleNames(names_map);
}
void QRangeModel_resetRoleNames(void* self) {
	((QRangeModel*)self)->resetRoleNames();
}
int QRangeModel_hasChildren(void* self) {
	auto&& __r = ((QRangeModel*)self)->hasChildren();
	return __r;
}
int QRangeModel_autoConnectPolicy(void* self) {
	auto&& __r = ((QRangeModel*)self)->autoConnectPolicy();
	return (int)(__r);
}
void QRangeModel_setAutoConnectPolicy(void* self, int policy) {
	((QRangeModel*)self)->setAutoConnectPolicy((QRangeModel::AutoConnectPolicy)(policy));
}
void QRangeModel_roleNamesChanged(void* self) {
	((QRangeModel*)self)->roleNamesChanged();
}
void QRangeModel_autoConnectPolicyChanged(void* self, int policy) {
	((QRangeModel*)self)->autoConnectPolicyChanged((QRangeModel::AutoConnectPolicy)(policy));
}
void* QRangeModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QRangeModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QRangeModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QRangeModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QRangeModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QRangeModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRangeModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QRangeModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRangeModel_index2(void* self, int row, int column, void* parent) {
	auto&& __r = ((QRangeModel*)self)->index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
int QRangeModel_rowCount2(void* self, void* parent) {
	auto&& __r = ((QRangeModel*)self)->rowCount(*(QModelIndex*)parent);
	return __r;
}
int QRangeModel_columnCount2(void* self, void* parent) {
	auto&& __r = ((QRangeModel*)self)->columnCount(*(QModelIndex*)parent);
	return __r;
}
int QRangeModel_insertColumns2(void* self, int column, int count, void* parent) {
	auto&& __r = ((QRangeModel*)self)->insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QRangeModel_removeColumns2(void* self, int column, int count, void* parent) {
	auto&& __r = ((QRangeModel*)self)->removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QRangeModel_insertRows2(void* self, int row, int count, void* parent) {
	auto&& __r = ((QRangeModel*)self)->insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QRangeModel_removeRows2(void* self, int row, int count, void* parent) {
	auto&& __r = ((QRangeModel*)self)->removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QRangeModel_hasChildren2(void* self, void* parent) {
	auto&& __r = ((QRangeModel*)self)->hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QRangeModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QRangeModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QRangeModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QRangeModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QRangeModel_connect_roleNamesChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::roleNamesChanged, (QRangeModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_autoConnectPolicyChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::autoConnectPolicyChanged, (QRangeModel*)self, [__id](QRangeModel::AutoConnectPolicy policy){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(policy));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::headerDataChanged, (QRangeModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::layoutChanged, (QRangeModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::layoutChanged, (QRangeModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::layoutAboutToBeChanged, (QRangeModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::layoutAboutToBeChanged, (QRangeModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::rowsAboutToBeInserted, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::rowsInserted, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::rowsAboutToBeRemoved, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::rowsRemoved, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::columnsAboutToBeInserted, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::columnsInserted, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::columnsAboutToBeRemoved, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::columnsRemoved, (QRangeModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::modelAboutToBeReset, (QRangeModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::modelReset, (QRangeModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::rowsAboutToBeMoved, (QRangeModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::rowsMoved, (QRangeModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::columnsAboutToBeMoved, (QRangeModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::columnsMoved, (QRangeModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::destroyed, (QRangeModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::destroyed, (QRangeModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRangeModel*)self, &QRangeModel::objectNameChanged, (QRangeModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
	__conns_QRangeModel[__id] = __c;
}
void QRangeModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QRangeModel);
		auto __it = __conns_QRangeModel.find(id);
		if (__it == __conns_QRangeModel.end()) return;
		__c = __it->second;
		__conns_QRangeModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void QRangeModel_delete(void* self) { delete (QRangeModel*)self; }

}

#include <QtCore/qstandardpaths.h>
#include <QtCore/qstring.h>
#include "gen_qstandardpaths.h"

extern "C" {

char* QStandardPaths_writableLocation(int type_, long long* __out_len) {
	auto&& __r = QStandardPaths::writableLocation((QStandardPaths::StandardLocation)(type_));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QStandardPaths_standardLocations(int type_, long long* __out_len) {
	auto&& __r = QStandardPaths::standardLocations((QStandardPaths::StandardLocation)(type_));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QStandardPaths_locate(int type_, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QStandardPaths::locate((QStandardPaths::StandardLocation)(type_), fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QStandardPaths_locateAll(int type_, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QStandardPaths::locateAll((QStandardPaths::StandardLocation)(type_), fileName_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QStandardPaths_displayName(int type_, long long* __out_len) {
	auto&& __r = QStandardPaths::displayName((QStandardPaths::StandardLocation)(type_));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStandardPaths_findExecutable(const char* executableName, long long executableName_len, long long* __out_len) {
	QString executableName_qstr = QString::fromUtf8(executableName, executableName_len);
	auto&& __r = QStandardPaths::findExecutable(executableName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardPaths_setTestModeEnabled(int testMode) {
	QStandardPaths::setTestModeEnabled(testMode);
}
int QStandardPaths_isTestModeEnabled() {
	auto&& __r = QStandardPaths::isTestModeEnabled();
	return __r;
}
char* QStandardPaths_locate2(int type_, const char* fileName, long long fileName_len, int options, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QStandardPaths::locate((QStandardPaths::StandardLocation)(type_), fileName_qstr, QFlags<QStandardPaths::LocateOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QStandardPaths_locateAll2(int type_, const char* fileName, long long fileName_len, int options, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QStandardPaths::locateAll((QStandardPaths::StandardLocation)(type_), fileName_qstr, QFlags<QStandardPaths::LocateOption>(options));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QStandardPaths_findExecutable2(const char* executableName, long long executableName_len, char** paths, long long paths_len, long long* __out_len) {
	QString executableName_qstr = QString::fromUtf8(executableName, executableName_len);
	QList<QString> paths_list; paths_list.reserve(paths_len);
	for (long long __i = 0; __i < paths_len; __i++) paths_list.append(QString::fromUtf8(paths[__i]));
	auto&& __r = QStandardPaths::findExecutable(executableName_qstr, paths_list);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}

}

#include <QtGui/qcolor.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qrgba64.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qvariant.h>
#include <QtGui/qbrush.h>
#include <QtGui/qpen.h>
#include "gen_qcolor.h"

extern "C" {

void* QColor_new() {
	return new QColor();
}
void* QColor_new2(int color) {
	return new QColor((Qt::GlobalColor)(color));
}
void* QColor_new3(int r, int g, int b) {
	return new QColor(r, g, b);
}
void* QColor_new4(unsigned int rgb) {
	return new QColor(rgb);
}
void* QColor_new5(void* rgba64) {
	return new QColor(*(QRgba64*)rgba64);
}
void* QColor_new6(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QColor(name_qstr);
}
void* QColor_new7(void* name) {
	return new QColor(*(QStringView*)name);
}
void* QColor_new8(const char* aname) {
	return new QColor(aname);
}
void* QColor_new9(void* name) {
	return new QColor(*(QLatin1String*)name);
}
void* QColor_new10(int spec) {
	return new QColor((QColor::Spec)(spec));
}
void* QColor_new11(int spec, unsigned short a1, unsigned short a2, unsigned short a3, unsigned short a4) {
	return new QColor((QColor::Spec)(spec), a1, a2, a3, a4);
}
void* QColor_new12(void* arg1) {
	return new QColor(*(QColor*)arg1);
}
void* QColor_new13(int r, int g, int b, int a) {
	return new QColor(r, g, b, a);
}
void* QColor_new14(int spec, unsigned short a1, unsigned short a2, unsigned short a3, unsigned short a4, unsigned short a5) {
	return new QColor((QColor::Spec)(spec), a1, a2, a3, a4, a5);
}
void* QColor_fromString(const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = QColor::fromString(name_qstr);
	return new QColor(__r);
}
void* QColor_OperatorAssign(void* self, int color) {
	auto&& __r = ((QColor*)self)->operator=((Qt::GlobalColor)(color));
	return (void*)&(__r);
}
int QColor_isValid(void* self) {
	auto&& __r = ((QColor*)self)->isValid();
	return __r;
}
char* QColor_name(void* self, long long* __out_len) {
	auto&& __r = ((QColor*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QColor_setNamedColor(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QColor*)self)->setNamedColor(name_qstr);
}
void QColor_setNamedColor2(void* self, void* name) {
	((QColor*)self)->setNamedColor(*(QStringView*)name);
}
void QColor_setNamedColor3(void* self, void* name) {
	((QColor*)self)->setNamedColor(*(QLatin1String*)name);
}
char** QColor_colorNames(long long* __out_len) {
	auto&& __r = QColor::colorNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QColor_spec(void* self) {
	auto&& __r = ((QColor*)self)->spec();
	return (int)(__r);
}
int QColor_alpha(void* self) {
	auto&& __r = ((QColor*)self)->alpha();
	return __r;
}
void QColor_setAlpha(void* self, int alpha) {
	((QColor*)self)->setAlpha(alpha);
}
float QColor_alphaF(void* self) {
	auto&& __r = ((QColor*)self)->alphaF();
	return __r;
}
void QColor_setAlphaF(void* self, float alpha) {
	((QColor*)self)->setAlphaF(alpha);
}
int QColor_red(void* self) {
	auto&& __r = ((QColor*)self)->red();
	return __r;
}
int QColor_green(void* self) {
	auto&& __r = ((QColor*)self)->green();
	return __r;
}
int QColor_blue(void* self) {
	auto&& __r = ((QColor*)self)->blue();
	return __r;
}
void QColor_setRed(void* self, int red) {
	((QColor*)self)->setRed(red);
}
void QColor_setGreen(void* self, int green) {
	((QColor*)self)->setGreen(green);
}
void QColor_setBlue(void* self, int blue) {
	((QColor*)self)->setBlue(blue);
}
float QColor_redF(void* self) {
	auto&& __r = ((QColor*)self)->redF();
	return __r;
}
float QColor_greenF(void* self) {
	auto&& __r = ((QColor*)self)->greenF();
	return __r;
}
float QColor_blueF(void* self) {
	auto&& __r = ((QColor*)self)->blueF();
	return __r;
}
void QColor_setRedF(void* self, float red) {
	((QColor*)self)->setRedF(red);
}
void QColor_setGreenF(void* self, float green) {
	((QColor*)self)->setGreenF(green);
}
void QColor_setBlueF(void* self, float blue) {
	((QColor*)self)->setBlueF(blue);
}
void QColor_getRgb(void* self, int* r, int* g, int* b) {
	int r_local{};
	int g_local{};
	int b_local{};
	((QColor*)self)->getRgb(&r_local, &g_local, &b_local);
	*r = r_local;
	*g = g_local;
	*b = b_local;
}
void QColor_setRgb(void* self, int r, int g, int b) {
	((QColor*)self)->setRgb(r, g, b);
}
void QColor_getRgbF(void* self, float* r, float* g, float* b) {
	float r_local{};
	float g_local{};
	float b_local{};
	((QColor*)self)->getRgbF(&r_local, &g_local, &b_local);
	*r = r_local;
	*g = g_local;
	*b = b_local;
}
void QColor_setRgbF(void* self, float r, float g, float b) {
	((QColor*)self)->setRgbF(r, g, b);
}
void* QColor_rgba64(void* self) {
	auto&& __r = ((QColor*)self)->rgba64();
	return new QRgba64(__r);
}
void QColor_setRgba64(void* self, void* rgba) {
	((QColor*)self)->setRgba64(*(QRgba64*)rgba);
}
unsigned int QColor_rgba(void* self) {
	auto&& __r = ((QColor*)self)->rgba();
	return __r;
}
void QColor_setRgba(void* self, unsigned int rgba) {
	((QColor*)self)->setRgba(rgba);
}
unsigned int QColor_rgb(void* self) {
	auto&& __r = ((QColor*)self)->rgb();
	return __r;
}
void QColor_setRgb2(void* self, unsigned int rgb) {
	((QColor*)self)->setRgb(rgb);
}
int QColor_hue(void* self) {
	auto&& __r = ((QColor*)self)->hue();
	return __r;
}
int QColor_saturation(void* self) {
	auto&& __r = ((QColor*)self)->saturation();
	return __r;
}
int QColor_hsvHue(void* self) {
	auto&& __r = ((QColor*)self)->hsvHue();
	return __r;
}
int QColor_hsvSaturation(void* self) {
	auto&& __r = ((QColor*)self)->hsvSaturation();
	return __r;
}
int QColor_value(void* self) {
	auto&& __r = ((QColor*)self)->value();
	return __r;
}
float QColor_hueF(void* self) {
	auto&& __r = ((QColor*)self)->hueF();
	return __r;
}
float QColor_saturationF(void* self) {
	auto&& __r = ((QColor*)self)->saturationF();
	return __r;
}
float QColor_hsvHueF(void* self) {
	auto&& __r = ((QColor*)self)->hsvHueF();
	return __r;
}
float QColor_hsvSaturationF(void* self) {
	auto&& __r = ((QColor*)self)->hsvSaturationF();
	return __r;
}
float QColor_valueF(void* self) {
	auto&& __r = ((QColor*)self)->valueF();
	return __r;
}
void QColor_getHsv(void* self, int* h, int* s, int* v) {
	int h_local{};
	int s_local{};
	int v_local{};
	((QColor*)self)->getHsv(&h_local, &s_local, &v_local);
	*h = h_local;
	*s = s_local;
	*v = v_local;
}
void QColor_setHsv(void* self, int h, int s, int v) {
	((QColor*)self)->setHsv(h, s, v);
}
void QColor_getHsvF(void* self, float* h, float* s, float* v) {
	float h_local{};
	float s_local{};
	float v_local{};
	((QColor*)self)->getHsvF(&h_local, &s_local, &v_local);
	*h = h_local;
	*s = s_local;
	*v = v_local;
}
void QColor_setHsvF(void* self, float h, float s, float v) {
	((QColor*)self)->setHsvF(h, s, v);
}
int QColor_cyan(void* self) {
	auto&& __r = ((QColor*)self)->cyan();
	return __r;
}
int QColor_magenta(void* self) {
	auto&& __r = ((QColor*)self)->magenta();
	return __r;
}
int QColor_yellow(void* self) {
	auto&& __r = ((QColor*)self)->yellow();
	return __r;
}
int QColor_black(void* self) {
	auto&& __r = ((QColor*)self)->black();
	return __r;
}
float QColor_cyanF(void* self) {
	auto&& __r = ((QColor*)self)->cyanF();
	return __r;
}
float QColor_magentaF(void* self) {
	auto&& __r = ((QColor*)self)->magentaF();
	return __r;
}
float QColor_yellowF(void* self) {
	auto&& __r = ((QColor*)self)->yellowF();
	return __r;
}
float QColor_blackF(void* self) {
	auto&& __r = ((QColor*)self)->blackF();
	return __r;
}
void QColor_getCmyk(void* self, int* c, int* m, int* y, int* k) {
	int c_local{};
	int m_local{};
	int y_local{};
	int k_local{};
	((QColor*)self)->getCmyk(&c_local, &m_local, &y_local, &k_local);
	*c = c_local;
	*m = m_local;
	*y = y_local;
	*k = k_local;
}
void QColor_setCmyk(void* self, int c, int m, int y, int k) {
	((QColor*)self)->setCmyk(c, m, y, k);
}
void QColor_getCmykF(void* self, float* c, float* m, float* y, float* k) {
	float c_local{};
	float m_local{};
	float y_local{};
	float k_local{};
	((QColor*)self)->getCmykF(&c_local, &m_local, &y_local, &k_local);
	*c = c_local;
	*m = m_local;
	*y = y_local;
	*k = k_local;
}
void QColor_setCmykF(void* self, float c, float m, float y, float k) {
	((QColor*)self)->setCmykF(c, m, y, k);
}
int QColor_hslHue(void* self) {
	auto&& __r = ((QColor*)self)->hslHue();
	return __r;
}
int QColor_hslSaturation(void* self) {
	auto&& __r = ((QColor*)self)->hslSaturation();
	return __r;
}
int QColor_lightness(void* self) {
	auto&& __r = ((QColor*)self)->lightness();
	return __r;
}
float QColor_hslHueF(void* self) {
	auto&& __r = ((QColor*)self)->hslHueF();
	return __r;
}
float QColor_hslSaturationF(void* self) {
	auto&& __r = ((QColor*)self)->hslSaturationF();
	return __r;
}
float QColor_lightnessF(void* self) {
	auto&& __r = ((QColor*)self)->lightnessF();
	return __r;
}
void QColor_getHsl(void* self, int* h, int* s, int* l) {
	int h_local{};
	int s_local{};
	int l_local{};
	((QColor*)self)->getHsl(&h_local, &s_local, &l_local);
	*h = h_local;
	*s = s_local;
	*l = l_local;
}
void QColor_setHsl(void* self, int h, int s, int l) {
	((QColor*)self)->setHsl(h, s, l);
}
void QColor_getHslF(void* self, float* h, float* s, float* l) {
	float h_local{};
	float s_local{};
	float l_local{};
	((QColor*)self)->getHslF(&h_local, &s_local, &l_local);
	*h = h_local;
	*s = s_local;
	*l = l_local;
}
void QColor_setHslF(void* self, float h, float s, float l) {
	((QColor*)self)->setHslF(h, s, l);
}
void* QColor_toRgb(void* self) {
	auto&& __r = ((QColor*)self)->toRgb();
	return new QColor(__r);
}
void* QColor_toHsv(void* self) {
	auto&& __r = ((QColor*)self)->toHsv();
	return new QColor(__r);
}
void* QColor_toCmyk(void* self) {
	auto&& __r = ((QColor*)self)->toCmyk();
	return new QColor(__r);
}
void* QColor_toHsl(void* self) {
	auto&& __r = ((QColor*)self)->toHsl();
	return new QColor(__r);
}
void* QColor_toExtendedRgb(void* self) {
	auto&& __r = ((QColor*)self)->toExtendedRgb();
	return new QColor(__r);
}
void* QColor_convertTo(void* self, int colorSpec) {
	auto&& __r = ((QColor*)self)->convertTo((QColor::Spec)(colorSpec));
	return new QColor(__r);
}
void* QColor_fromRgb(unsigned int rgb) {
	auto&& __r = QColor::fromRgb(rgb);
	return new QColor(__r);
}
void* QColor_fromRgba(unsigned int rgba) {
	auto&& __r = QColor::fromRgba(rgba);
	return new QColor(__r);
}
void* QColor_fromRgb2(int r, int g, int b) {
	auto&& __r = QColor::fromRgb(r, g, b);
	return new QColor(__r);
}
void* QColor_fromRgbF(float r, float g, float b) {
	auto&& __r = QColor::fromRgbF(r, g, b);
	return new QColor(__r);
}
void* QColor_fromRgba64(unsigned short r, unsigned short g, unsigned short b) {
	auto&& __r = QColor::fromRgba64(r, g, b);
	return new QColor(__r);
}
void* QColor_fromRgba642(void* rgba) {
	auto&& __r = QColor::fromRgba64(*(QRgba64*)rgba);
	return new QColor(__r);
}
void* QColor_fromHsv(int h, int s, int v) {
	auto&& __r = QColor::fromHsv(h, s, v);
	return new QColor(__r);
}
void* QColor_fromHsvF(float h, float s, float v) {
	auto&& __r = QColor::fromHsvF(h, s, v);
	return new QColor(__r);
}
void* QColor_fromCmyk(int c, int m, int y, int k) {
	auto&& __r = QColor::fromCmyk(c, m, y, k);
	return new QColor(__r);
}
void* QColor_fromCmykF(float c, float m, float y, float k) {
	auto&& __r = QColor::fromCmykF(c, m, y, k);
	return new QColor(__r);
}
void* QColor_fromHsl(int h, int s, int l) {
	auto&& __r = QColor::fromHsl(h, s, l);
	return new QColor(__r);
}
void* QColor_fromHslF(float h, float s, float l) {
	auto&& __r = QColor::fromHslF(h, s, l);
	return new QColor(__r);
}
void* QColor_lighter(void* self) {
	auto&& __r = ((QColor*)self)->lighter();
	return new QColor(__r);
}
void* QColor_darker(void* self) {
	auto&& __r = ((QColor*)self)->darker();
	return new QColor(__r);
}
int QColor_OperatorEqual(void* self, void* c) {
	auto&& __r = ((QColor*)self)->operator==(*(QColor*)c);
	return __r;
}
int QColor_OperatorNotEqual(void* self, void* c) {
	auto&& __r = ((QColor*)self)->operator!=(*(QColor*)c);
	return __r;
}
void* QColor_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QColor*)self);
	return new QVariant(__r);
}
int QColor_isValidColor(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = QColor::isValidColor(name_qstr);
	return __r;
}
int QColor_isValidColor2(void* arg1) {
	auto&& __r = QColor::isValidColor(*(QStringView*)arg1);
	return __r;
}
int QColor_isValidColor3(void* arg1) {
	auto&& __r = QColor::isValidColor(*(QLatin1String*)arg1);
	return __r;
}
int QColor_isValidColorName(const char* arg1, long long arg1_len) {
	QAnyStringView arg1_qstr = QAnyStringView(arg1, arg1_len);
	auto&& __r = QColor::isValidColorName(arg1_qstr);
	return __r;
}
void* QColor_OperatorAssign2(void* self, void* arg1) {
	auto&& __r = ((QColor*)self)->operator=(*(QColor*)arg1);
	return (void*)&(__r);
}
char* QColor_name2(void* self, int format, long long* __out_len) {
	auto&& __r = ((QColor*)self)->name((QColor::NameFormat)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QColor_getRgb2(void* self, int* r, int* g, int* b, int* a) {
	int r_local{};
	int g_local{};
	int b_local{};
	int a_local{};
	((QColor*)self)->getRgb(&r_local, &g_local, &b_local, &a_local);
	*r = r_local;
	*g = g_local;
	*b = b_local;
	*a = a_local;
}
void QColor_setRgb3(void* self, int r, int g, int b, int a) {
	((QColor*)self)->setRgb(r, g, b, a);
}
void QColor_getRgbF2(void* self, float* r, float* g, float* b, float* a) {
	float r_local{};
	float g_local{};
	float b_local{};
	float a_local{};
	((QColor*)self)->getRgbF(&r_local, &g_local, &b_local, &a_local);
	*r = r_local;
	*g = g_local;
	*b = b_local;
	*a = a_local;
}
void QColor_setRgbF2(void* self, float r, float g, float b, float a) {
	((QColor*)self)->setRgbF(r, g, b, a);
}
void QColor_getHsv2(void* self, int* h, int* s, int* v, int* a) {
	int h_local{};
	int s_local{};
	int v_local{};
	int a_local{};
	((QColor*)self)->getHsv(&h_local, &s_local, &v_local, &a_local);
	*h = h_local;
	*s = s_local;
	*v = v_local;
	*a = a_local;
}
void QColor_setHsv2(void* self, int h, int s, int v, int a) {
	((QColor*)self)->setHsv(h, s, v, a);
}
void QColor_getHsvF2(void* self, float* h, float* s, float* v, float* a) {
	float h_local{};
	float s_local{};
	float v_local{};
	float a_local{};
	((QColor*)self)->getHsvF(&h_local, &s_local, &v_local, &a_local);
	*h = h_local;
	*s = s_local;
	*v = v_local;
	*a = a_local;
}
void QColor_setHsvF2(void* self, float h, float s, float v, float a) {
	((QColor*)self)->setHsvF(h, s, v, a);
}
void QColor_getCmyk2(void* self, int* c, int* m, int* y, int* k, int* a) {
	int c_local{};
	int m_local{};
	int y_local{};
	int k_local{};
	int a_local{};
	((QColor*)self)->getCmyk(&c_local, &m_local, &y_local, &k_local, &a_local);
	*c = c_local;
	*m = m_local;
	*y = y_local;
	*k = k_local;
	*a = a_local;
}
void QColor_setCmyk2(void* self, int c, int m, int y, int k, int a) {
	((QColor*)self)->setCmyk(c, m, y, k, a);
}
void QColor_getCmykF2(void* self, float* c, float* m, float* y, float* k, float* a) {
	float c_local{};
	float m_local{};
	float y_local{};
	float k_local{};
	float a_local{};
	((QColor*)self)->getCmykF(&c_local, &m_local, &y_local, &k_local, &a_local);
	*c = c_local;
	*m = m_local;
	*y = y_local;
	*k = k_local;
	*a = a_local;
}
void QColor_setCmykF2(void* self, float c, float m, float y, float k, float a) {
	((QColor*)self)->setCmykF(c, m, y, k, a);
}
void QColor_getHsl2(void* self, int* h, int* s, int* l, int* a) {
	int h_local{};
	int s_local{};
	int l_local{};
	int a_local{};
	((QColor*)self)->getHsl(&h_local, &s_local, &l_local, &a_local);
	*h = h_local;
	*s = s_local;
	*l = l_local;
	*a = a_local;
}
void QColor_setHsl2(void* self, int h, int s, int l, int a) {
	((QColor*)self)->setHsl(h, s, l, a);
}
void QColor_getHslF2(void* self, float* h, float* s, float* l, float* a) {
	float h_local{};
	float s_local{};
	float l_local{};
	float a_local{};
	((QColor*)self)->getHslF(&h_local, &s_local, &l_local, &a_local);
	*h = h_local;
	*s = s_local;
	*l = l_local;
	*a = a_local;
}
void QColor_setHslF2(void* self, float h, float s, float l, float a) {
	((QColor*)self)->setHslF(h, s, l, a);
}
void* QColor_fromRgb3(int r, int g, int b, int a) {
	auto&& __r = QColor::fromRgb(r, g, b, a);
	return new QColor(__r);
}
void* QColor_fromRgbF2(float r, float g, float b, float a) {
	auto&& __r = QColor::fromRgbF(r, g, b, a);
	return new QColor(__r);
}
void* QColor_fromRgba643(unsigned short r, unsigned short g, unsigned short b, unsigned short a) {
	auto&& __r = QColor::fromRgba64(r, g, b, a);
	return new QColor(__r);
}
void* QColor_fromHsv2(int h, int s, int v, int a) {
	auto&& __r = QColor::fromHsv(h, s, v, a);
	return new QColor(__r);
}
void* QColor_fromHsvF2(float h, float s, float v, float a) {
	auto&& __r = QColor::fromHsvF(h, s, v, a);
	return new QColor(__r);
}
void* QColor_fromCmyk2(int c, int m, int y, int k, int a) {
	auto&& __r = QColor::fromCmyk(c, m, y, k, a);
	return new QColor(__r);
}
void* QColor_fromCmykF2(float c, float m, float y, float k, float a) {
	auto&& __r = QColor::fromCmykF(c, m, y, k, a);
	return new QColor(__r);
}
void* QColor_fromHsl2(int h, int s, int l, int a) {
	auto&& __r = QColor::fromHsl(h, s, l, a);
	return new QColor(__r);
}
void* QColor_fromHslF2(float h, float s, float l, float a) {
	auto&& __r = QColor::fromHslF(h, s, l, a);
	return new QColor(__r);
}
void* QColor_lighter2(void* self, int f) {
	auto&& __r = ((QColor*)self)->lighter(f);
	return new QColor(__r);
}
void* QColor_darker2(void* self, int f) {
	auto&& __r = ((QColor*)self)->darker(f);
	return new QColor(__r);
}
int QColor_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QColor*)self) == (*(QBrush*)rhs);
	return __r;
}
int QColor_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QColor*)self) != (*(QBrush*)rhs);
	return __r;
}
int QColor_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QColor*)self) == (*(QPen*)rhs);
	return __r;
}
int QColor_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QColor*)self) != (*(QPen*)rhs);
	return __r;
}
void* QColor_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QColor*)self);
	return new QVariant(__r);
}
void QColor_delete(void* self) { delete (QColor*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobject.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qcoreevent.h>
#include "gen_qenterevent.h"

extern "C" {

void* QEnterEvent_base(void* self) { return (void*)static_cast<QSinglePointEvent*>((QEnterEvent*)self); }
void QEnterEvent_chain(void* self, void** out) {
	QEnterEvent* __p = (QEnterEvent*)self;
	QSinglePointEvent* __b0 = static_cast<QSinglePointEvent*>(__p); out[0] = (void*)__b0;
	QPointerEvent* __b1 = static_cast<QPointerEvent*>(__b0); out[1] = (void*)__b1;
	QInputEvent* __b2 = static_cast<QInputEvent*>(__b1); out[2] = (void*)__b2;
	QEvent* __b3 = static_cast<QEvent*>(__b2); out[3] = (void*)__b3;
}
void* QEnterEvent_new(void* localPos, void* scenePos, void* globalPos) {
	return new QEnterEvent(*(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos);
}
void* QEnterEvent_new2(void* localPos, void* scenePos, void* globalPos, void* device) {
	return new QEnterEvent(*(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, (QPointingDevice*)device);
}
void* QEnterEvent_pos(void* self) {
	auto&& __r = ((QEnterEvent*)self)->pos();
	return new QPoint(__r);
}
void* QEnterEvent_globalPos(void* self) {
	auto&& __r = ((QEnterEvent*)self)->globalPos();
	return new QPoint(__r);
}
int QEnterEvent_x(void* self) {
	auto&& __r = ((QEnterEvent*)self)->x();
	return __r;
}
int QEnterEvent_y(void* self) {
	auto&& __r = ((QEnterEvent*)self)->y();
	return __r;
}
int QEnterEvent_globalX(void* self) {
	auto&& __r = ((QEnterEvent*)self)->globalX();
	return __r;
}
int QEnterEvent_globalY(void* self) {
	auto&& __r = ((QEnterEvent*)self)->globalY();
	return __r;
}
void* QEnterEvent_localPos(void* self) {
	auto&& __r = ((QEnterEvent*)self)->localPos();
	return new QPointF(__r);
}
void* QEnterEvent_windowPos(void* self) {
	auto&& __r = ((QEnterEvent*)self)->windowPos();
	return new QPointF(__r);
}
void* QEnterEvent_screenPos(void* self) {
	auto&& __r = ((QEnterEvent*)self)->screenPos();
	return new QPointF(__r);
}
int QEnterEvent_registerEventType() {
	auto&& __r = QEnterEvent::registerEventType();
	return __r;
}
int QEnterEvent_registerEventType2(int hint) {
	auto&& __r = QEnterEvent::registerEventType(hint);
	return __r;
}
void QEnterEvent_delete(void* self) { delete (QEnterEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qdragleaveevent.h"
#include "gen_purego_cb.h"

class QDragLeaveEvent__Go : public QDragLeaveEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QDragLeaveEvent__Go(long long id, A&&... a) : QDragLeaveEvent(std::forward<A>(a)...), __goId(id) {}
	QDragLeaveEvent* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QDragLeaveEvent*)__r; }
		return QDragLeaveEvent::clone();
	}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QDragLeaveEvent::setAccepted(accepted);
	}
};

extern "C" {

void* QDragLeaveEvent_base(void* self) { return (void*)static_cast<QEvent*>((QDragLeaveEvent*)self); }
void QDragLeaveEvent__set_override(void* self, int idx) { ((QDragLeaveEvent__Go*)self)->__goOn[idx] = 1; }
void* QDragLeaveEvent_new(long long __goid) {
	return new QDragLeaveEvent__Go(__goid);
}
int QDragLeaveEvent_registerEventType() {
	auto&& __r = QDragLeaveEvent::registerEventType();
	return __r;
}
int QDragLeaveEvent_registerEventType2(int hint) {
	auto&& __r = QDragLeaveEvent::registerEventType(hint);
	return __r;
}
void* QDragLeaveEvent_clone_default(void* self) {
	auto&& __r = ((QDragLeaveEvent*)self)->QDragLeaveEvent::clone();
	return (void*)(__r);
}
void QDragLeaveEvent_setAccepted_default(void* self, int accepted) {
	((QDragLeaveEvent*)self)->QDragLeaveEvent::setAccepted(accepted);
}
void* QDragLeaveEvent__Go_new(long long id) { return new QDragLeaveEvent__Go(id); }
void QDragLeaveEvent_delete(void* self) { delete (QDragLeaveEvent*)self; }

}

#include <QtGui/qbrush.h>
#include <QtGui/qcolor.h>
#include "gen_qgradient.h"

extern "C" {

void* QGradient_new() {
	return new QGradient();
}
void* QGradient_new2(int arg1) {
	return new QGradient((QGradient::Preset)(arg1));
}
void* QGradient_new3(void* arg1) {
	return new QGradient(*(QGradient*)arg1);
}
int QGradient_type(void* self) {
	auto&& __r = ((QGradient*)self)->type();
	return (int)(__r);
}
void QGradient_setSpread(void* self, int spread) {
	((QGradient*)self)->setSpread((QGradient::Spread)(spread));
}
int QGradient_spread(void* self) {
	auto&& __r = ((QGradient*)self)->spread();
	return (int)(__r);
}
void QGradient_setColorAt(void* self, double pos, void* color) {
	((QGradient*)self)->setColorAt(pos, *(QColor*)color);
}
int QGradient_coordinateMode(void* self) {
	auto&& __r = ((QGradient*)self)->coordinateMode();
	return (int)(__r);
}
void QGradient_setCoordinateMode(void* self, int mode) {
	((QGradient*)self)->setCoordinateMode((QGradient::CoordinateMode)(mode));
}
int QGradient_interpolationMode(void* self) {
	auto&& __r = ((QGradient*)self)->interpolationMode();
	return (int)(__r);
}
void QGradient_setInterpolationMode(void* self, int mode) {
	((QGradient*)self)->setInterpolationMode((QGradient::InterpolationMode)(mode));
}
int QGradient_OperatorEqual(void* self, void* gradient) {
	auto&& __r = ((QGradient*)self)->operator==(*(QGradient*)gradient);
	return __r;
}
int QGradient_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QGradient*)self)->operator!=(*(QGradient*)other);
	return __r;
}
void QGradient_delete(void* self) { delete (QGradient*)self; }

}

#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtGui/qtextcursor.h>
#include <QtGui/qtextformat.h>
#include "gen_qabstracttextdocumentlayout__selection.h"

extern "C" {

void* QAbstractTextDocumentLayout__Selection_new(void* arg1) {
	return new QAbstractTextDocumentLayout::Selection(*(QAbstractTextDocumentLayout::Selection*)arg1);
}
void* QAbstractTextDocumentLayout__Selection_new2() {
	return new QAbstractTextDocumentLayout::Selection();
}
void* QAbstractTextDocumentLayout__Selection_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAbstractTextDocumentLayout::Selection*)self)->operator=(*(QAbstractTextDocumentLayout::Selection*)arg1);
	return (void*)&(__r);
}
void* QAbstractTextDocumentLayout__Selection_cursor_get(void* self) {
	auto __r = ((QAbstractTextDocumentLayout::Selection*)self)->cursor;
	return new QTextCursor(__r);
}
void QAbstractTextDocumentLayout__Selection_cursor_set(void* self, void* cursor) {
	((QAbstractTextDocumentLayout::Selection*)self)->cursor = *(QTextCursor*)cursor;
}
void* QAbstractTextDocumentLayout__Selection_format_get(void* self) {
	auto __r = ((QAbstractTextDocumentLayout::Selection*)self)->format;
	return new QTextCharFormat(__r);
}
void QAbstractTextDocumentLayout__Selection_format_set(void* self, void* format) {
	((QAbstractTextDocumentLayout::Selection*)self)->format = *(QTextCharFormat*)format;
}
void QAbstractTextDocumentLayout__Selection_delete(void* self) { delete (QAbstractTextDocumentLayout::Selection*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibletextupdateevent.h"

extern "C" {

void* QAccessibleTextUpdateEvent_base(void* self) { return (void*)static_cast<QAccessibleTextCursorEvent*>((QAccessibleTextUpdateEvent*)self); }
void QAccessibleTextUpdateEvent_chain(void* self, void** out) {
	QAccessibleTextUpdateEvent* __p = (QAccessibleTextUpdateEvent*)self;
	QAccessibleTextCursorEvent* __b0 = static_cast<QAccessibleTextCursorEvent*>(__p); out[0] = (void*)__b0;
	QAccessibleEvent* __b1 = static_cast<QAccessibleEvent*>(__b0); out[1] = (void*)__b1;
}
void* QAccessibleTextUpdateEvent_new(void* obj, int position, const char* oldText, long long oldText_len, const char* text, long long text_len) {
	QString oldText_qstr = QString::fromUtf8(oldText, oldText_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAccessibleTextUpdateEvent((QObject*)obj, position, oldText_qstr, text_qstr);
}
void* QAccessibleTextUpdateEvent_new2(void* iface, int position, const char* oldText, long long oldText_len, const char* text, long long text_len) {
	QString oldText_qstr = QString::fromUtf8(oldText, oldText_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAccessibleTextUpdateEvent((QAccessibleInterface*)iface, position, oldText_qstr, text_qstr);
}
char* QAccessibleTextUpdateEvent_textRemoved(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTextUpdateEvent*)self)->textRemoved();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTextUpdateEvent_textInserted(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTextUpdateEvent*)self)->textInserted();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleTextUpdateEvent_changePosition(void* self) {
	auto&& __r = ((QAccessibleTextUpdateEvent*)self)->changePosition();
	return __r;
}
void QAccessibleTextUpdateEvent_delete(void* self) { delete (QAccessibleTextUpdateEvent*)self; }

}

#include <QtGui/qfontvariableaxis.h>
#include <QtGui/qfont.h>
#include <QtCore/qstring.h>
#include "gen_qfontvariableaxis.h"

extern "C" {

void* QFontVariableAxis_new() {
	return new QFontVariableAxis();
}
void* QFontVariableAxis_new2(void* axis) {
	return new QFontVariableAxis(*(QFontVariableAxis*)axis);
}
void QFontVariableAxis_swap(void* self, void* other) {
	((QFontVariableAxis*)self)->swap(*(QFontVariableAxis*)other);
}
void* QFontVariableAxis_OperatorAssign(void* self, void* axis) {
	auto&& __r = ((QFontVariableAxis*)self)->operator=(*(QFontVariableAxis*)axis);
	return (void*)&(__r);
}
void* QFontVariableAxis_tag(void* self) {
	auto&& __r = ((QFontVariableAxis*)self)->tag();
	return new QFont::Tag(__r);
}
void QFontVariableAxis_setTag(void* self, void* tag) {
	((QFontVariableAxis*)self)->setTag(*(QFont::Tag*)tag);
}
char* QFontVariableAxis_name(void* self, long long* __out_len) {
	auto&& __r = ((QFontVariableAxis*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontVariableAxis_setName(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QFontVariableAxis*)self)->setName(name_qstr);
}
double QFontVariableAxis_minimumValue(void* self) {
	auto&& __r = ((QFontVariableAxis*)self)->minimumValue();
	return __r;
}
void QFontVariableAxis_setMinimumValue(void* self, double minimumValue) {
	((QFontVariableAxis*)self)->setMinimumValue(minimumValue);
}
double QFontVariableAxis_maximumValue(void* self) {
	auto&& __r = ((QFontVariableAxis*)self)->maximumValue();
	return __r;
}
void QFontVariableAxis_setMaximumValue(void* self, double maximumValue) {
	((QFontVariableAxis*)self)->setMaximumValue(maximumValue);
}
double QFontVariableAxis_defaultValue(void* self) {
	auto&& __r = ((QFontVariableAxis*)self)->defaultValue();
	return __r;
}
void QFontVariableAxis_setDefaultValue(void* self, double defaultValue) {
	((QFontVariableAxis*)self)->setDefaultValue(defaultValue);
}
void QFontVariableAxis_delete(void* self) { delete (QFontVariableAxis*)self; }

}

#include <QtGui/qpagedpaintdevice.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpagelayout.h>
#include <QtGui/qpagesize.h>
#include <QtCore/qmargins.h>
#include <QtGui/qpageranges.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include "gen_qpagedpaintdevice.h"

extern "C" {

void* QPagedPaintDevice_base(void* self) { return (void*)static_cast<QPaintDevice*>((QPagedPaintDevice*)self); }
int QPagedPaintDevice_newPage(void* self) {
	auto&& __r = ((QPagedPaintDevice*)self)->newPage();
	return __r;
}
int QPagedPaintDevice_setPageLayout(void* self, void* pageLayout) {
	auto&& __r = ((QPagedPaintDevice*)self)->setPageLayout(*(QPageLayout*)pageLayout);
	return __r;
}
int QPagedPaintDevice_setPageSize(void* self, void* pageSize) {
	auto&& __r = ((QPagedPaintDevice*)self)->setPageSize(*(QPageSize*)pageSize);
	return __r;
}
int QPagedPaintDevice_setPageOrientation(void* self, int orientation) {
	auto&& __r = ((QPagedPaintDevice*)self)->setPageOrientation((QPageLayout::Orientation)(orientation));
	return __r;
}
int QPagedPaintDevice_setPageMargins(void* self, void* margins, int units) {
	auto&& __r = ((QPagedPaintDevice*)self)->setPageMargins(*(QMarginsF*)margins, (QPageLayout::Unit)(units));
	return __r;
}
void* QPagedPaintDevice_pageLayout(void* self) {
	auto&& __r = ((QPagedPaintDevice*)self)->pageLayout();
	return new QPageLayout(__r);
}
void QPagedPaintDevice_setPageRanges(void* self, void* ranges) {
	((QPagedPaintDevice*)self)->setPageRanges(*(QPageRanges*)ranges);
}
void* QPagedPaintDevice_pageRanges(void* self) {
	auto&& __r = ((QPagedPaintDevice*)self)->pageRanges();
	return new QPageRanges(__r);
}
double QPagedPaintDevice_devicePixelRatioFScale() {
	auto&& __r = QPagedPaintDevice::devicePixelRatioFScale();
	return __r;
}
int QPagedPaintDevice_encodeMetricF(int metric, double value) {
	auto&& __r = QPagedPaintDevice::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
void QPagedPaintDevice_delete(void* self) { delete (QPagedPaintDevice*)self; }

}

#include <QtGui/qstatictext.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qsize.h>
#include <QtGui/qtransform.h>
#include <QtGui/qfont.h>
#include "gen_qstatictext.h"

extern "C" {

void* QStaticText_new() {
	return new QStaticText();
}
void* QStaticText_new2(const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QStaticText(text_qstr);
}
void* QStaticText_new3(void* other) {
	return new QStaticText(*(QStaticText*)other);
}
void* QStaticText_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStaticText*)self)->operator=(*(QStaticText*)arg1);
	return (void*)&(__r);
}
void QStaticText_swap(void* self, void* other) {
	((QStaticText*)self)->swap(*(QStaticText*)other);
}
void QStaticText_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStaticText*)self)->setText(text_qstr);
}
char* QStaticText_text(void* self, long long* __out_len) {
	auto&& __r = ((QStaticText*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStaticText_setTextFormat(void* self, int textFormat) {
	((QStaticText*)self)->setTextFormat((Qt::TextFormat)(textFormat));
}
int QStaticText_textFormat(void* self) {
	auto&& __r = ((QStaticText*)self)->textFormat();
	return (int)(__r);
}
void QStaticText_setTextWidth(void* self, double textWidth) {
	((QStaticText*)self)->setTextWidth(textWidth);
}
double QStaticText_textWidth(void* self) {
	auto&& __r = ((QStaticText*)self)->textWidth();
	return __r;
}
void QStaticText_setTextOption(void* self, void* textOption) {
	((QStaticText*)self)->setTextOption(*(QTextOption*)textOption);
}
void* QStaticText_textOption(void* self) {
	auto&& __r = ((QStaticText*)self)->textOption();
	return new QTextOption(__r);
}
void* QStaticText_size(void* self) {
	auto&& __r = ((QStaticText*)self)->size();
	return new QSizeF(__r);
}
void QStaticText_prepare(void* self) {
	((QStaticText*)self)->prepare();
}
void QStaticText_setPerformanceHint(void* self, int performanceHint) {
	((QStaticText*)self)->setPerformanceHint((QStaticText::PerformanceHint)(performanceHint));
}
int QStaticText_performanceHint(void* self) {
	auto&& __r = ((QStaticText*)self)->performanceHint();
	return (int)(__r);
}
int QStaticText_OperatorEqual(void* self, void* arg1) {
	auto&& __r = ((QStaticText*)self)->operator==(*(QStaticText*)arg1);
	return __r;
}
int QStaticText_OperatorNotEqual(void* self, void* arg1) {
	auto&& __r = ((QStaticText*)self)->operator!=(*(QStaticText*)arg1);
	return __r;
}
void QStaticText_prepare2(void* self, void* matrix) {
	((QStaticText*)self)->prepare(*(QTransform*)matrix);
}
void QStaticText_prepare3(void* self, void* matrix, void* font) {
	((QStaticText*)self)->prepare(*(QTransform*)matrix, *(QFont*)font);
}
void QStaticText_delete(void* self) { delete (QStaticText*)self; }

}

#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include "gen_qwidgetdata.h"

extern "C" {

void* QWidgetData_new(void* arg1) {
	return new QWidgetData(*(QWidgetData*)arg1);
}
void* QWidgetData_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QWidgetData*)self)->operator=(*(QWidgetData*)arg1);
	return (void*)&(__r);
}
unsigned long long QWidgetData_winid_get(void* self) {
	auto __r = ((QWidgetData*)self)->winid;
	return __r;
}
void QWidgetData_winid_set(void* self, unsigned long long winid) {
	((QWidgetData*)self)->winid = winid;
}
unsigned int QWidgetData_widget_attributes_get(void* self) {
	auto __r = ((QWidgetData*)self)->widget_attributes;
	return __r;
}
void QWidgetData_widget_attributes_set(void* self, unsigned int widget_attributes) {
	((QWidgetData*)self)->widget_attributes = widget_attributes;
}
int QWidgetData_window_flags_get(void* self) {
	auto __r = ((QWidgetData*)self)->window_flags;
	return (int)(__r);
}
void QWidgetData_window_flags_set(void* self, int window_flags) {
	((QWidgetData*)self)->window_flags = QFlags<Qt::WindowType>(window_flags);
}
unsigned int QWidgetData_window_state_get(void* self) {
	auto __r = ((QWidgetData*)self)->window_state;
	return __r;
}
void QWidgetData_window_state_set(void* self, unsigned int window_state) {
	((QWidgetData*)self)->window_state = window_state;
}
unsigned int QWidgetData_focus_policy_get(void* self) {
	auto __r = ((QWidgetData*)self)->focus_policy;
	return __r;
}
void QWidgetData_focus_policy_set(void* self, unsigned int focus_policy) {
	((QWidgetData*)self)->focus_policy = focus_policy;
}
unsigned int QWidgetData_sizehint_forced_get(void* self) {
	auto __r = ((QWidgetData*)self)->sizehint_forced;
	return __r;
}
void QWidgetData_sizehint_forced_set(void* self, unsigned int sizehint_forced) {
	((QWidgetData*)self)->sizehint_forced = sizehint_forced;
}
unsigned int QWidgetData_is_closing_get(void* self) {
	auto __r = ((QWidgetData*)self)->is_closing;
	return __r;
}
void QWidgetData_is_closing_set(void* self, unsigned int is_closing) {
	((QWidgetData*)self)->is_closing = is_closing;
}
unsigned int QWidgetData_in_show_get(void* self) {
	auto __r = ((QWidgetData*)self)->in_show;
	return __r;
}
void QWidgetData_in_show_set(void* self, unsigned int in_show) {
	((QWidgetData*)self)->in_show = in_show;
}
unsigned int QWidgetData_in_set_window_state_get(void* self) {
	auto __r = ((QWidgetData*)self)->in_set_window_state;
	return __r;
}
void QWidgetData_in_set_window_state_set(void* self, unsigned int in_set_window_state) {
	((QWidgetData*)self)->in_set_window_state = in_set_window_state;
}
unsigned int QWidgetData_fstrut_dirty_get(void* self) {
	auto __r = ((QWidgetData*)self)->fstrut_dirty;
	return __r;
}
void QWidgetData_fstrut_dirty_set(void* self, unsigned int fstrut_dirty) {
	((QWidgetData*)self)->fstrut_dirty = fstrut_dirty;
}
unsigned int QWidgetData_context_menu_policy_get(void* self) {
	auto __r = ((QWidgetData*)self)->context_menu_policy;
	return __r;
}
void QWidgetData_context_menu_policy_set(void* self, unsigned int context_menu_policy) {
	((QWidgetData*)self)->context_menu_policy = context_menu_policy;
}
unsigned int QWidgetData_window_modality_get(void* self) {
	auto __r = ((QWidgetData*)self)->window_modality;
	return __r;
}
void QWidgetData_window_modality_set(void* self, unsigned int window_modality) {
	((QWidgetData*)self)->window_modality = window_modality;
}
unsigned int QWidgetData_in_destructor_get(void* self) {
	auto __r = ((QWidgetData*)self)->in_destructor;
	return __r;
}
void QWidgetData_in_destructor_set(void* self, unsigned int in_destructor) {
	((QWidgetData*)self)->in_destructor = in_destructor;
}
unsigned int QWidgetData_unused_get(void* self) {
	auto __r = ((QWidgetData*)self)->unused;
	return __r;
}
void QWidgetData_unused_set(void* self, unsigned int unused) {
	((QWidgetData*)self)->unused = unused;
}
void* QWidgetData_crect_get(void* self) {
	auto __r = ((QWidgetData*)self)->crect;
	return new QRect(__r);
}
void QWidgetData_crect_set(void* self, void* crect) {
	((QWidgetData*)self)->crect = *(QRect*)crect;
}
void* QWidgetData_pal_get(void* self) {
	auto __r = ((QWidgetData*)self)->pal;
	return new QPalette(__r);
}
void QWidgetData_pal_set(void* self, void* pal) {
	((QWidgetData*)self)->pal = *(QPalette*)pal;
}
void* QWidgetData_fnt_get(void* self) {
	auto __r = ((QWidgetData*)self)->fnt;
	return new QFont(__r);
}
void QWidgetData_fnt_set(void* self, void* fnt) {
	((QWidgetData*)self)->fnt = *(QFont*)fnt;
}
void* QWidgetData_wrect_get(void* self) {
	auto __r = ((QWidgetData*)self)->wrect;
	return new QRect(__r);
}
void QWidgetData_wrect_set(void* self, void* wrect) {
	((QWidgetData*)self)->wrect = *(QRect*)wrect;
}
void QWidgetData_delete(void* self) { delete (QWidgetData*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include "gen_qstyleoptiondockwidget.h"

extern "C" {

void* QStyleOptionDockWidget_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionDockWidget*)self); }
void* QStyleOptionDockWidget_new() {
	return new QStyleOptionDockWidget();
}
void* QStyleOptionDockWidget_new2(void* other) {
	return new QStyleOptionDockWidget(*(QStyleOptionDockWidget*)other);
}
void* QStyleOptionDockWidget_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionDockWidget*)self)->operator=(*(QStyleOptionDockWidget*)arg1);
	return (void*)&(__r);
}
char* QStyleOptionDockWidget_title_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionDockWidget*)self)->title;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionDockWidget_title_set(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QStyleOptionDockWidget*)self)->title = title_qstr;
}
int QStyleOptionDockWidget_closable_get(void* self) {
	auto __r = ((QStyleOptionDockWidget*)self)->closable;
	return __r;
}
void QStyleOptionDockWidget_closable_set(void* self, int closable) {
	((QStyleOptionDockWidget*)self)->closable = closable;
}
int QStyleOptionDockWidget_movable_get(void* self) {
	auto __r = ((QStyleOptionDockWidget*)self)->movable;
	return __r;
}
void QStyleOptionDockWidget_movable_set(void* self, int movable) {
	((QStyleOptionDockWidget*)self)->movable = movable;
}
int QStyleOptionDockWidget_floatable_get(void* self) {
	auto __r = ((QStyleOptionDockWidget*)self)->floatable;
	return __r;
}
void QStyleOptionDockWidget_floatable_set(void* self, int floatable) {
	((QStyleOptionDockWidget*)self)->floatable = floatable;
}
int QStyleOptionDockWidget_verticalTitleBar_get(void* self) {
	auto __r = ((QStyleOptionDockWidget*)self)->verticalTitleBar;
	return __r;
}
void QStyleOptionDockWidget_verticalTitleBar_set(void* self, int verticalTitleBar) {
	((QStyleOptionDockWidget*)self)->verticalTitleBar = verticalTitleBar;
}
void QStyleOptionDockWidget_delete(void* self) { delete (QStyleOptionDockWidget*)self; }

}

#include <QtWidgets/qlayoutitem.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include "gen_qwidgetitem.h"

extern "C" {

void* QWidgetItem_base(void* self) { return (void*)static_cast<QLayoutItem*>((QWidgetItem*)self); }
void* QWidgetItem_new(void* w) {
	return new QWidgetItem((QWidget*)w);
}
void QWidgetItem_delete(void* self) { delete (QWidgetItem*)self; }

}

#include <QtWidgets/qdialogbuttonbox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qpushbutton.h>
#include <QtCore/qcoreevent.h>
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
#include <QtGui/qevent.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdialogbuttonbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDialogButtonBox__Go : public QDialogButtonBox {
public:
	long long __goId;
	unsigned char __goOn[48] = {};
	template<typename... A> QDialogButtonBox__Go(long long id, A&&... a) : QDialogButtonBox(std::forward<A>(a)...), __goId(id) {}
	using QDialogButtonBox::updateMicroFocus;
	using QDialogButtonBox::create;
	using QDialogButtonBox::destroy;
	using QDialogButtonBox::focusNextChild;
	using QDialogButtonBox::focusPreviousChild;
	using QDialogButtonBox::sender;
	using QDialogButtonBox::senderSignalIndex;
	using QDialogButtonBox::receivers;
	using QDialogButtonBox::isSignalConnected;
	using QDialogButtonBox::getDecodedMetricF;
	void cppgoBase_changeEvent(QEvent* __a0) { QDialogButtonBox::changeEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QDialogButtonBox::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDialogButtonBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDialogButtonBox::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDialogButtonBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDialogButtonBox::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDialogButtonBox::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDialogButtonBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDialogButtonBox::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDialogButtonBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDialogButtonBox::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDialogButtonBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDialogButtonBox::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDialogButtonBox::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDialogButtonBox::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDialogButtonBox::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDialogButtonBox::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDialogButtonBox::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDialogButtonBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDialogButtonBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDialogButtonBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDialogButtonBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDialogButtonBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDialogButtonBox::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDialogButtonBox::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDialogButtonBox::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDialogButtonBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDialogButtonBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDialogButtonBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDialogButtonBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDialogButtonBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDialogButtonBox::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDialogButtonBox::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDialogButtonBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDialogButtonBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDialogButtonBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDialogButtonBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDialogButtonBox::metaObject();
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::changeEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDialogButtonBox::event(event);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDialogButtonBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDialogButtonBox::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDialogButtonBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDialogButtonBox::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDialogButtonBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDialogButtonBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDialogButtonBox::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDialogButtonBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDialogButtonBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDialogButtonBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDialogButtonBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDialogButtonBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDialogButtonBox::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDialogButtonBox::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDialogButtonBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDialogButtonBox::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDialogButtonBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDialogButtonBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDialogButtonBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDialogButtonBox;
static std::mutex __connsMu_QDialogButtonBox;

extern "C" {

void* QDialogButtonBox_base(void* self) { return (void*)static_cast<QWidget*>((QDialogButtonBox*)self); }
void QDialogButtonBox_chain(void* self, void** out) {
	QDialogButtonBox* __p = (QDialogButtonBox*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QDialogButtonBox__set_override(void* self, int idx) { ((QDialogButtonBox__Go*)self)->__goOn[idx] = 1; }
void* QDialogButtonBox_fromQObject(void* o) { return (void*)qobject_cast<QDialogButtonBox*>((QObject*)o); }
void* QDialogButtonBox_new(long long __goid, void* parent) {
	return new QDialogButtonBox__Go(__goid, (QWidget*)parent);
}
void* QDialogButtonBox_new2(long long __goid) {
	return new QDialogButtonBox__Go(__goid);
}
void* QDialogButtonBox_new3(long long __goid, int orientation) {
	return new QDialogButtonBox__Go(__goid, (Qt::Orientation)(orientation));
}
void* QDialogButtonBox_new4(long long __goid, int buttons) {
	return new QDialogButtonBox__Go(__goid, QFlags<QDialogButtonBox::StandardButton>(buttons));
}
void* QDialogButtonBox_new5(long long __goid, int buttons, int orientation) {
	return new QDialogButtonBox__Go(__goid, QFlags<QDialogButtonBox::StandardButton>(buttons), (Qt::Orientation)(orientation));
}
void* QDialogButtonBox_new6(long long __goid, int orientation, void* parent) {
	return new QDialogButtonBox__Go(__goid, (Qt::Orientation)(orientation), (QWidget*)parent);
}
void* QDialogButtonBox_new7(long long __goid, int buttons, void* parent) {
	return new QDialogButtonBox__Go(__goid, QFlags<QDialogButtonBox::StandardButton>(buttons), (QWidget*)parent);
}
void* QDialogButtonBox_new8(long long __goid, int buttons, int orientation, void* parent) {
	return new QDialogButtonBox__Go(__goid, QFlags<QDialogButtonBox::StandardButton>(buttons), (Qt::Orientation)(orientation), (QWidget*)parent);
}
char* QDialogButtonBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QDialogButtonBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDialogButtonBox_setOrientation(void* self, int orientation) {
	((QDialogButtonBox*)self)->setOrientation((Qt::Orientation)(orientation));
}
int QDialogButtonBox_orientation(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->orientation();
	return (int)(__r);
}
void QDialogButtonBox_addButton(void* self, void* button, int role) {
	((QDialogButtonBox*)self)->addButton((QAbstractButton*)button, (QDialogButtonBox::ButtonRole)(role));
}
void* QDialogButtonBox_addButton2(void* self, const char* text, long long text_len, int role) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QDialogButtonBox*)self)->addButton(text_qstr, (QDialogButtonBox::ButtonRole)(role));
	return (void*)(__r);
}
void* QDialogButtonBox_addButton3(void* self, int button) {
	auto&& __r = ((QDialogButtonBox*)self)->addButton((QDialogButtonBox::StandardButton)(button));
	return (void*)(__r);
}
void QDialogButtonBox_removeButton(void* self, void* button) {
	((QDialogButtonBox*)self)->removeButton((QAbstractButton*)button);
}
void QDialogButtonBox_clear(void* self) {
	((QDialogButtonBox*)self)->clear();
}
void** QDialogButtonBox_buttons(void* self, long long* __out_len) {
	auto&& __r = ((QDialogButtonBox*)self)->buttons();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QDialogButtonBox_buttonRole(void* self, void* button) {
	auto&& __r = ((QDialogButtonBox*)self)->buttonRole((QAbstractButton*)button);
	return (int)(__r);
}
void QDialogButtonBox_setStandardButtons(void* self, int buttons) {
	((QDialogButtonBox*)self)->setStandardButtons(QFlags<QDialogButtonBox::StandardButton>(buttons));
}
int QDialogButtonBox_standardButtons(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->standardButtons();
	return (int)(__r);
}
int QDialogButtonBox_standardButton(void* self, void* button) {
	auto&& __r = ((QDialogButtonBox*)self)->standardButton((QAbstractButton*)button);
	return (int)(__r);
}
void* QDialogButtonBox_button(void* self, int which) {
	auto&& __r = ((QDialogButtonBox*)self)->button((QDialogButtonBox::StandardButton)(which));
	return (void*)(__r);
}
void QDialogButtonBox_setCenterButtons(void* self, int center) {
	((QDialogButtonBox*)self)->setCenterButtons(center);
}
int QDialogButtonBox_centerButtons(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->centerButtons();
	return __r;
}
void QDialogButtonBox_clicked(void* self, void* button) {
	((QDialogButtonBox*)self)->clicked((QAbstractButton*)button);
}
void QDialogButtonBox_accepted(void* self) {
	((QDialogButtonBox*)self)->accepted();
}
void QDialogButtonBox_helpRequested(void* self) {
	((QDialogButtonBox*)self)->helpRequested();
}
void QDialogButtonBox_rejected(void* self) {
	((QDialogButtonBox*)self)->rejected();
}
void QDialogButtonBox_setTabOrder(void* arg1, void* arg2) {
	QDialogButtonBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDialogButtonBox_mouseGrabber() {
	auto&& __r = QDialogButtonBox::mouseGrabber();
	return (void*)(__r);
}
void* QDialogButtonBox_keyboardGrabber() {
	auto&& __r = QDialogButtonBox::keyboardGrabber();
	return (void*)(__r);
}
void* QDialogButtonBox_find(unsigned long long arg1) {
	auto&& __r = QDialogButtonBox::find(arg1);
	return (void*)(__r);
}
void* QDialogButtonBox_createWindowContainer(void* window) {
	auto&& __r = QDialogButtonBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDialogButtonBox_updateMicroFocus(void* self) {
	((QDialogButtonBox__Go*)self)->updateMicroFocus();
}
void QDialogButtonBox_create(void* self) {
	((QDialogButtonBox__Go*)self)->create();
}
void QDialogButtonBox_destroy(void* self) {
	((QDialogButtonBox__Go*)self)->destroy();
}
int QDialogButtonBox_focusNextChild(void* self) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->focusNextChild();
	return __r;
}
int QDialogButtonBox_focusPreviousChild(void* self) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDialogButtonBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDialogButtonBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDialogButtonBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDialogButtonBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDialogButtonBox_sender(void* self) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->sender();
	return (void*)(__r);
}
int QDialogButtonBox_senderSignalIndex(void* self) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QDialogButtonBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->receivers(signal);
	return __r;
}
int QDialogButtonBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDialogButtonBox_devicePixelRatioFScale() {
	auto&& __r = QDialogButtonBox::devicePixelRatioFScale();
	return __r;
}
int QDialogButtonBox_encodeMetricF(int metric, double value) {
	auto&& __r = QDialogButtonBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDialogButtonBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDialogButtonBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDialogButtonBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDialogButtonBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDialogButtonBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDialogButtonBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDialogButtonBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDialogButtonBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDialogButtonBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDialogButtonBox_updateMicroFocus2(void* self, int query) {
	((QDialogButtonBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDialogButtonBox_create2(void* self, unsigned long long arg1) {
	((QDialogButtonBox__Go*)self)->create(arg1);
}
void QDialogButtonBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDialogButtonBox__Go*)self)->create(arg1, initializeWindow);
}
void QDialogButtonBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDialogButtonBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDialogButtonBox_destroy2(void* self, int destroyWindow) {
	((QDialogButtonBox__Go*)self)->destroy(destroyWindow);
}
void QDialogButtonBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDialogButtonBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDialogButtonBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDialogButtonBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDialogButtonBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDialogButtonBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDialogButtonBox_metaObject_default(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::metaObject();
	return (void*)(__r);
}
void QDialogButtonBox_changeEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
int QDialogButtonBox_event_default(void* self, void* event) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QDialogButtonBox_devType_default(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::devType();
	return __r;
}
void QDialogButtonBox_setVisible_default(void* self, int visible) {
	((QDialogButtonBox*)self)->QDialogButtonBox::setVisible(visible);
}
void* QDialogButtonBox_sizeHint_default(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::sizeHint();
	return new QSize(__r);
}
void* QDialogButtonBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::minimumSizeHint();
	return new QSize(__r);
}
int QDialogButtonBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::heightForWidth(arg1);
	return __r;
}
int QDialogButtonBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::hasHeightForWidth();
	return __r;
}
void* QDialogButtonBox_paintEngine_default(void* self) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::paintEngine();
	return (void*)(__r);
}
void QDialogButtonBox_mousePressEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QDialogButtonBox_mouseReleaseEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QDialogButtonBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDialogButtonBox_mouseMoveEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QDialogButtonBox_wheelEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QDialogButtonBox_keyPressEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QDialogButtonBox_keyReleaseEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDialogButtonBox_focusInEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QDialogButtonBox_focusOutEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDialogButtonBox_enterEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDialogButtonBox_leaveEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDialogButtonBox_paintEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QDialogButtonBox_moveEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDialogButtonBox_resizeEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QDialogButtonBox_closeEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QDialogButtonBox_contextMenuEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QDialogButtonBox_tabletEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDialogButtonBox_actionEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDialogButtonBox_dragEnterEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDialogButtonBox_dragMoveEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDialogButtonBox_dragLeaveEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDialogButtonBox_dropEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QDialogButtonBox_showEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QDialogButtonBox_hideEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QDialogButtonBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDialogButtonBox_initPainter_default(void* self, void* painter) {
	((QDialogButtonBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDialogButtonBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDialogButtonBox_sharedPainter_default(void* self) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDialogButtonBox_inputMethodEvent_default(void* self, void* arg1) {
	((QDialogButtonBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QDialogButtonBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QDialogButtonBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDialogButtonBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QDialogButtonBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDialogButtonBox*)self)->QDialogButtonBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDialogButtonBox_timerEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDialogButtonBox_childEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDialogButtonBox_customEvent_default(void* self, void* event) {
	((QDialogButtonBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDialogButtonBox_connectNotify_default(void* self, void* signal) {
	((QDialogButtonBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDialogButtonBox_disconnectNotify_default(void* self, void* signal) {
	((QDialogButtonBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDialogButtonBox_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::clicked, (QDialogButtonBox*)self, [__id](QAbstractButton* button){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)button);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::accepted, (QDialogButtonBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_helpRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::helpRequested, (QDialogButtonBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::rejected, (QDialogButtonBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::windowTitleChanged, (QDialogButtonBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::windowIconChanged, (QDialogButtonBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::windowIconTextChanged, (QDialogButtonBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::customContextMenuRequested, (QDialogButtonBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::destroyed, (QDialogButtonBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::destroyed, (QDialogButtonBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialogButtonBox*)self, &QDialogButtonBox::objectNameChanged, (QDialogButtonBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
	__conns_QDialogButtonBox[__id] = __c;
}
void QDialogButtonBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDialogButtonBox);
		auto __it = __conns_QDialogButtonBox.find(id);
		if (__it == __conns_QDialogButtonBox.end()) return;
		__c = __it->second;
		__conns_QDialogButtonBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDialogButtonBox__Go_new(long long id) { return new QDialogButtonBox__Go(id); }
void QDialogButtonBox_delete(void* self) { delete (QDialogButtonBox*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
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
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicsellipseitem.h"
#include "gen_purego_cb.h"

class QGraphicsEllipseItem__Go : public QGraphicsEllipseItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsEllipseItem__Go(long long id, A&&... a) : QGraphicsEllipseItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsEllipseItem::updateMicroFocus;
	using QGraphicsEllipseItem::addToIndex;
	using QGraphicsEllipseItem::removeFromIndex;
	using QGraphicsEllipseItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsEllipseItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsEllipseItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsEllipseItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsEllipseItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsEllipseItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsEllipseItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsEllipseItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsEllipseItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsEllipseItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsEllipseItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsEllipseItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsEllipseItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsEllipseItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsEllipseItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsEllipseItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsEllipseItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsEllipseItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsEllipseItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsEllipseItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsEllipseItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsEllipseItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsEllipseItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsEllipseItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsEllipseItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsEllipseItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsEllipseItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsEllipseItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsEllipseItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsEllipseItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsEllipseItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsEllipseItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsEllipseItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsEllipseItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEllipseItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsEllipseItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsEllipseItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsEllipseItem_base(void* self) { return (void*)static_cast<QAbstractGraphicsShapeItem*>((QGraphicsEllipseItem*)self); }
void QGraphicsEllipseItem_chain(void* self, void** out) {
	QGraphicsEllipseItem* __p = (QGraphicsEllipseItem*)self;
	QAbstractGraphicsShapeItem* __b0 = static_cast<QAbstractGraphicsShapeItem*>(__p); out[0] = (void*)__b0;
	QGraphicsItem* __b1 = static_cast<QGraphicsItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsEllipseItem__set_override(void* self, int idx) { ((QGraphicsEllipseItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsEllipseItem_new(long long __goid) {
	return new QGraphicsEllipseItem__Go(__goid);
}
void* QGraphicsEllipseItem_new2(long long __goid, void* rect) {
	return new QGraphicsEllipseItem__Go(__goid, *(QRectF*)rect);
}
void* QGraphicsEllipseItem_new3(long long __goid, double x, double y, double w, double h) {
	return new QGraphicsEllipseItem__Go(__goid, x, y, w, h);
}
void* QGraphicsEllipseItem_new4(long long __goid, void* parent) {
	return new QGraphicsEllipseItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsEllipseItem_new5(long long __goid, void* rect, void* parent) {
	return new QGraphicsEllipseItem__Go(__goid, *(QRectF*)rect, (QGraphicsItem*)parent);
}
void* QGraphicsEllipseItem_new6(long long __goid, double x, double y, double w, double h, void* parent) {
	return new QGraphicsEllipseItem__Go(__goid, x, y, w, h, (QGraphicsItem*)parent);
}
void* QGraphicsEllipseItem_rect(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->rect();
	return new QRectF(__r);
}
void QGraphicsEllipseItem_setRect(void* self, void* rect) {
	((QGraphicsEllipseItem*)self)->setRect(*(QRectF*)rect);
}
void QGraphicsEllipseItem_setRect2(void* self, double x, double y, double w, double h) {
	((QGraphicsEllipseItem*)self)->setRect(x, y, w, h);
}
int QGraphicsEllipseItem_startAngle(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->startAngle();
	return __r;
}
void QGraphicsEllipseItem_setStartAngle(void* self, int angle) {
	((QGraphicsEllipseItem*)self)->setStartAngle(angle);
}
int QGraphicsEllipseItem_spanAngle(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->spanAngle();
	return __r;
}
void QGraphicsEllipseItem_setSpanAngle(void* self, int angle) {
	((QGraphicsEllipseItem*)self)->setSpanAngle(angle);
}
void QGraphicsEllipseItem_updateMicroFocus(void* self) {
	((QGraphicsEllipseItem__Go*)self)->updateMicroFocus();
}
void QGraphicsEllipseItem_addToIndex(void* self) {
	((QGraphicsEllipseItem__Go*)self)->addToIndex();
}
void QGraphicsEllipseItem_removeFromIndex(void* self) {
	((QGraphicsEllipseItem__Go*)self)->removeFromIndex();
}
void QGraphicsEllipseItem_prepareGeometryChange(void* self) {
	((QGraphicsEllipseItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsEllipseItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsEllipseItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsEllipseItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsEllipseItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsEllipseItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsEllipseItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsEllipseItem_type_default(void* self) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::type();
	return __r;
}
int QGraphicsEllipseItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsEllipseItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsEllipseItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsEllipseItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsEllipseItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsEllipseItem_advance_default(void* self, int phase) {
	((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::advance(phase);
}
int QGraphicsEllipseItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsEllipseItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsEllipseItem*)self)->QGraphicsEllipseItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsEllipseItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsEllipseItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsEllipseItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsEllipseItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsEllipseItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsEllipseItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsEllipseItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsEllipseItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsEllipseItem_dropEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsEllipseItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsEllipseItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsEllipseItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsEllipseItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsEllipseItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsEllipseItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsEllipseItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsEllipseItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsEllipseItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsEllipseItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsEllipseItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsEllipseItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsEllipseItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsEllipseItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsEllipseItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsEllipseItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsEllipseItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsEllipseItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsEllipseItem__Go_new(long long id) { return new QGraphicsEllipseItem__Go(id); }
void QGraphicsEllipseItem_delete(void* self) { delete (QGraphicsEllipseItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include "gen_qgraphicsscenewheelevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneWheelEvent__Go : public QGraphicsSceneWheelEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneWheelEvent__Go(long long id, A&&... a) : QGraphicsSceneWheelEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneWheelEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneWheelEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneWheelEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneWheelEvent*)self); }
void QGraphicsSceneWheelEvent_chain(void* self, void** out) {
	QGraphicsSceneWheelEvent* __p = (QGraphicsSceneWheelEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneWheelEvent__set_override(void* self, int idx) { ((QGraphicsSceneWheelEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneWheelEvent_new(long long __goid) {
	return new QGraphicsSceneWheelEvent__Go(__goid);
}
void* QGraphicsSceneWheelEvent_new2(long long __goid, int type_) {
	return new QGraphicsSceneWheelEvent__Go(__goid, (QEvent::Type)(type_));
}
void* QGraphicsSceneWheelEvent_pos(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->pos();
	return new QPointF(__r);
}
void QGraphicsSceneWheelEvent_setPos(void* self, void* pos) {
	((QGraphicsSceneWheelEvent*)self)->setPos(*(QPointF*)pos);
}
void* QGraphicsSceneWheelEvent_scenePos(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsSceneWheelEvent_setScenePos(void* self, void* pos) {
	((QGraphicsSceneWheelEvent*)self)->setScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneWheelEvent_screenPos(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->screenPos();
	return new QPoint(__r);
}
void QGraphicsSceneWheelEvent_setScreenPos(void* self, void* pos) {
	((QGraphicsSceneWheelEvent*)self)->setScreenPos(*(QPoint*)pos);
}
int QGraphicsSceneWheelEvent_buttons(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->buttons();
	return (int)(__r);
}
void QGraphicsSceneWheelEvent_setButtons(void* self, int buttons) {
	((QGraphicsSceneWheelEvent*)self)->setButtons(QFlags<Qt::MouseButton>(buttons));
}
int QGraphicsSceneWheelEvent_modifiers(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->modifiers();
	return (int)(__r);
}
void QGraphicsSceneWheelEvent_setModifiers(void* self, int modifiers) {
	((QGraphicsSceneWheelEvent*)self)->setModifiers(QFlags<Qt::KeyboardModifier>(modifiers));
}
int QGraphicsSceneWheelEvent_delta(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->delta();
	return __r;
}
void QGraphicsSceneWheelEvent_setDelta(void* self, int delta) {
	((QGraphicsSceneWheelEvent*)self)->setDelta(delta);
}
int QGraphicsSceneWheelEvent_orientation(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->orientation();
	return (int)(__r);
}
void QGraphicsSceneWheelEvent_setOrientation(void* self, int orientation) {
	((QGraphicsSceneWheelEvent*)self)->setOrientation((Qt::Orientation)(orientation));
}
int QGraphicsSceneWheelEvent_phase(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->phase();
	return (int)(__r);
}
void QGraphicsSceneWheelEvent_setPhase(void* self, int scrollPhase) {
	((QGraphicsSceneWheelEvent*)self)->setPhase((Qt::ScrollPhase)(scrollPhase));
}
void* QGraphicsSceneWheelEvent_pixelDelta(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->pixelDelta();
	return new QPoint(__r);
}
void QGraphicsSceneWheelEvent_setPixelDelta(void* self, void* delta) {
	((QGraphicsSceneWheelEvent*)self)->setPixelDelta(*(QPoint*)delta);
}
int QGraphicsSceneWheelEvent_isInverted(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->isInverted();
	return __r;
}
void QGraphicsSceneWheelEvent_setInverted(void* self, int inverted) {
	((QGraphicsSceneWheelEvent*)self)->setInverted(inverted);
}
int QGraphicsSceneWheelEvent_registerEventType() {
	auto&& __r = QGraphicsSceneWheelEvent::registerEventType();
	return __r;
}
int QGraphicsSceneWheelEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneWheelEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneWheelEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneWheelEvent*)self)->QGraphicsSceneWheelEvent::setAccepted(accepted);
}
void* QGraphicsSceneWheelEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneWheelEvent*)self)->QGraphicsSceneWheelEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneWheelEvent__Go_new(long long id) { return new QGraphicsSceneWheelEvent__Go(id); }
void QGraphicsSceneWheelEvent_delete(void* self) { delete (QGraphicsSceneWheelEvent*)self; }

}

#include <QtWidgets/qlistwidget.h>
#include <QtWidgets/qlistview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qregion.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
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
#include "gen_qlistwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QListWidget__Go : public QListWidget {
public:
	long long __goId;
	unsigned char __goOn[93] = {};
	template<typename... A> QListWidget__Go(long long id, A&&... a) : QListWidget(std::forward<A>(a)...), __goId(id) {}
	using QListWidget::resizeContents;
	using QListWidget::contentsSize;
	using QListWidget::rectForIndex;
	using QListWidget::setPositionForIndex;
	using QListWidget::state;
	using QListWidget::scheduleDelayedItemsLayout;
	using QListWidget::executeDelayedItemsLayout;
	using QListWidget::setDirtyRegion;
	using QListWidget::scrollDirtyRegion;
	using QListWidget::dirtyRegionOffset;
	using QListWidget::startAutoScroll;
	using QListWidget::stopAutoScroll;
	using QListWidget::doAutoScroll;
	using QListWidget::dropIndicatorPosition;
	using QListWidget::setViewportMargins;
	using QListWidget::viewportMargins;
	using QListWidget::drawFrame;
	using QListWidget::updateMicroFocus;
	using QListWidget::create;
	using QListWidget::destroy;
	using QListWidget::focusNextChild;
	using QListWidget::focusPreviousChild;
	using QListWidget::sender;
	using QListWidget::senderSignalIndex;
	using QListWidget::receivers;
	using QListWidget::isSignalConnected;
	using QListWidget::getDecodedMetricF;
	void cppgoBase_dropEvent(QDropEvent* __a0) { QListWidget::dropEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QListWidget::event(__a0); }
	QList<QString> cppgoBase_mimeTypes() { return QListWidget::mimeTypes(); }
	bool cppgoBase_dropMimeData(int __a0, const QMimeData* __a1, int __a2) { return QListWidget::dropMimeData(__a0, __a1, (Qt::DropAction)__a2); }
	QFlags<Qt::DropAction> cppgoBase_supportedDropActions() { return QListWidget::supportedDropActions(); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QListWidget::scrollContentsBy(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QListWidget::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QListWidget::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QListWidget::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QListWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QListWidget::wheelEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QListWidget::timerEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QListWidget::resizeEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QListWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QListWidget::dragLeaveEvent(__a0); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QListWidget::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QListWidget::initViewItemOption(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QListWidget::paintEvent(__a0); }
	int cppgoBase_horizontalOffset() { return QListWidget::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QListWidget::verticalOffset(); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QListWidget::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QListWidget::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QListWidget::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QListWidget::selectedIndexes(); }
	void cppgoBase_updateGeometries() { QListWidget::updateGeometries(); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QListWidget::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QListWidget::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QListWidget::currentChanged(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QListWidget::viewportSizeHint(); }
	void cppgoBase_updateEditorData() { QListWidget::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QListWidget::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QListWidget::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QListWidget::horizontalScrollbarAction(__a0); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QListWidget::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QListWidget::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QListWidget::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QListWidget::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QListWidget::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QListWidget::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QListWidget::selectionCommand(__a0, __a1); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QListWidget::focusNextPrevChild(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QListWidget::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QListWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QListWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QListWidget::dragEnterEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QListWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QListWidget::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QListWidget::keyPressEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QListWidget::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QListWidget::eventFilter(__a0, __a1); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QListWidget::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QListWidget::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QListWidget::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QListWidget::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QListWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QListWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QListWidget::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QListWidget::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QListWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QListWidget::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QListWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QListWidget::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QListWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QListWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QListWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QListWidget::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QListWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QListWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QListWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QListWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QListWidget::metaObject();
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QListWidget::setSelectionModel(selectionModel);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QListWidget::dropEvent(event);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QListWidget::event(e);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virtp_voidp) {
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
		return QListWidget::mimeTypes();
	}
	bool dropMimeData(int index, const QMimeData* data, Qt::DropAction action) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, void*, int, int*))cppgo_virt_int_int_voidp_int)(__goId, -9049994684305682212LL, index, (void*)data, (int)(action), &__h); if (__h) return __r; }
		return QListWidget::dropMimeData(index, data, action);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QListWidget::supportedDropActions();
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QListWidget::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QListWidget::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QListWidget::indexAt(p);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QListWidget::doItemsLayout();
	}
	void reset() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QListWidget::reset();
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QListWidget::setRootIndex(index);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QListWidget::scrollContentsBy(dx, dy);
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QListWidget::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QListWidget::rowsAboutToBeRemoved(parent, start, end);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QListWidget::mouseMoveEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QListWidget::mouseReleaseEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QListWidget::wheelEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QListWidget::timerEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QListWidget::resizeEvent(e);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QListWidget::dragMoveEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QListWidget::dragLeaveEvent(e);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QListWidget::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QListWidget::initViewItemOption(option);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QListWidget::paintEvent(e);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QListWidget::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QListWidget::verticalOffset();
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QListWidget::moveCursor(cursorAction, modifiers);
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QListWidget::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QListWidget::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virtp_voidp) {
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
		return QListWidget::selectedIndexes();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QListWidget::updateGeometries();
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QListWidget::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QListWidget::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QListWidget::currentChanged(current, previous);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QListWidget::viewportSizeHint();
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[37] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QListWidget::keyboardSearch(search);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QListWidget::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QListWidget::sizeHintForColumn(column);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QListWidget::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QListWidget::inputMethodQuery(query);
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QListWidget::selectAll();
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QListWidget::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QListWidget::updateEditorGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QListWidget::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QListWidget::horizontalScrollbarAction(action);
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QListWidget::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QListWidget::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QListWidget::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QListWidget::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QListWidget::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QListWidget::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QListWidget::selectionCommand(index, event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QListWidget::focusNextPrevChild(next);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QListWidget::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QListWidget::mousePressEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QListWidget::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QListWidget::dragEnterEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QListWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QListWidget::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QListWidget::keyPressEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QListWidget::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QListWidget::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QListWidget::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QListWidget::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QListWidget::setupViewport(viewport);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QListWidget::contextMenuEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QListWidget::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QListWidget::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QListWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QListWidget::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QListWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QListWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QListWidget::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QListWidget::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QListWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QListWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QListWidget::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QListWidget::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QListWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QListWidget::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QListWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QListWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QListWidget::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QListWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QListWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QListWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QListWidget::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QListWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QListWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[91] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QListWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[92] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QListWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QListWidget;
static std::mutex __connsMu_QListWidget;

extern "C" {

void* QListWidget_base(void* self) { return (void*)static_cast<QListView*>((QListWidget*)self); }
void QListWidget_chain(void* self, void** out) {
	QListWidget* __p = (QListWidget*)self;
	QListView* __b0 = static_cast<QListView*>(__p); out[0] = (void*)__b0;
	QAbstractItemView* __b1 = static_cast<QAbstractItemView*>(__b0); out[1] = (void*)__b1;
	QAbstractScrollArea* __b2 = static_cast<QAbstractScrollArea*>(__b1); out[2] = (void*)__b2;
	QFrame* __b3 = static_cast<QFrame*>(__b2); out[3] = (void*)__b3;
	QWidget* __b4 = static_cast<QWidget*>(__b3); out[4] = (void*)__b4;
	QObject* __b5 = static_cast<QObject*>(__b4); out[5] = (void*)__b5;
}
void QListWidget__set_override(void* self, int idx) { ((QListWidget__Go*)self)->__goOn[idx] = 1; }
void* QListWidget_fromQObject(void* o) { return (void*)qobject_cast<QListWidget*>((QObject*)o); }
void* QListWidget_new(long long __goid, void* parent) {
	return new QListWidget__Go(__goid, (QWidget*)parent);
}
void* QListWidget_new2(long long __goid) {
	return new QListWidget__Go(__goid);
}
char* QListWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QListWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QListWidget_item(void* self, int row) {
	auto&& __r = ((QListWidget*)self)->item(row);
	return (void*)(__r);
}
int QListWidget_row(void* self, void* item) {
	auto&& __r = ((QListWidget*)self)->row((QListWidgetItem*)item);
	return __r;
}
void QListWidget_insertItem(void* self, int row, void* item) {
	((QListWidget*)self)->insertItem(row, (QListWidgetItem*)item);
}
void QListWidget_insertItem2(void* self, int row, const char* label, long long label_len) {
	QString label_qstr = QString::fromUtf8(label, label_len);
	((QListWidget*)self)->insertItem(row, label_qstr);
}
void QListWidget_insertItems(void* self, int row, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QListWidget*)self)->insertItems(row, labels_list);
}
void QListWidget_addItem(void* self, const char* label, long long label_len) {
	QString label_qstr = QString::fromUtf8(label, label_len);
	((QListWidget*)self)->addItem(label_qstr);
}
void QListWidget_addItem2(void* self, void* item) {
	((QListWidget*)self)->addItem((QListWidgetItem*)item);
}
void QListWidget_addItems(void* self, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QListWidget*)self)->addItems(labels_list);
}
void* QListWidget_takeItem(void* self, int row) {
	auto&& __r = ((QListWidget*)self)->takeItem(row);
	return (void*)(__r);
}
int QListWidget_count(void* self) {
	auto&& __r = ((QListWidget*)self)->count();
	return __r;
}
void* QListWidget_currentItem(void* self) {
	auto&& __r = ((QListWidget*)self)->currentItem();
	return (void*)(__r);
}
void QListWidget_setCurrentItem(void* self, void* item) {
	((QListWidget*)self)->setCurrentItem((QListWidgetItem*)item);
}
void QListWidget_setCurrentItem2(void* self, void* item, int command) {
	((QListWidget*)self)->setCurrentItem((QListWidgetItem*)item, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
int QListWidget_currentRow(void* self) {
	auto&& __r = ((QListWidget*)self)->currentRow();
	return __r;
}
void QListWidget_setCurrentRow(void* self, int row) {
	((QListWidget*)self)->setCurrentRow(row);
}
void QListWidget_setCurrentRow2(void* self, int row, int command) {
	((QListWidget*)self)->setCurrentRow(row, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QListWidget_itemAt(void* self, void* p) {
	auto&& __r = ((QListWidget*)self)->itemAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QListWidget_itemAt2(void* self, int x, int y) {
	auto&& __r = ((QListWidget*)self)->itemAt(x, y);
	return (void*)(__r);
}
void* QListWidget_visualItemRect(void* self, void* item) {
	auto&& __r = ((QListWidget*)self)->visualItemRect((QListWidgetItem*)item);
	return new QRect(__r);
}
void QListWidget_sortItems(void* self) {
	((QListWidget*)self)->sortItems();
}
void QListWidget_setSortingEnabled(void* self, int enable) {
	((QListWidget*)self)->setSortingEnabled(enable);
}
int QListWidget_isSortingEnabled(void* self) {
	auto&& __r = ((QListWidget*)self)->isSortingEnabled();
	return __r;
}
void QListWidget_editItem(void* self, void* item) {
	((QListWidget*)self)->editItem((QListWidgetItem*)item);
}
void QListWidget_openPersistentEditor(void* self, void* item) {
	((QListWidget*)self)->openPersistentEditor((QListWidgetItem*)item);
}
void QListWidget_closePersistentEditor(void* self, void* item) {
	((QListWidget*)self)->closePersistentEditor((QListWidgetItem*)item);
}
int QListWidget_isPersistentEditorOpen(void* self, void* item) {
	auto&& __r = ((QListWidget*)self)->isPersistentEditorOpen((QListWidgetItem*)item);
	return __r;
}
void* QListWidget_itemWidget(void* self, void* item) {
	auto&& __r = ((QListWidget*)self)->itemWidget((QListWidgetItem*)item);
	return (void*)(__r);
}
void QListWidget_setItemWidget(void* self, void* item, void* widget) {
	((QListWidget*)self)->setItemWidget((QListWidgetItem*)item, (QWidget*)widget);
}
void QListWidget_removeItemWidget(void* self, void* item) {
	((QListWidget*)self)->removeItemWidget((QListWidgetItem*)item);
}
void** QListWidget_selectedItems(void* self, long long* __out_len) {
	auto&& __r = ((QListWidget*)self)->selectedItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QListWidget_findItems(void* self, const char* text, long long text_len, int flags, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QListWidget*)self)->findItems(text_qstr, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QListWidget_items(void* self, void* data, long long* __out_len) {
	auto&& __r = ((QListWidget*)self)->items((QMimeData*)data);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QListWidget_indexFromItem(void* self, void* item) {
	auto&& __r = ((QListWidget*)self)->indexFromItem((QListWidgetItem*)item);
	return new QModelIndex(__r);
}
void* QListWidget_itemFromIndex(void* self, void* index) {
	auto&& __r = ((QListWidget*)self)->itemFromIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void QListWidget_setSupportedDragActions(void* self, int actions) {
	((QListWidget*)self)->setSupportedDragActions(QFlags<Qt::DropAction>(actions));
}
int QListWidget_supportedDragActions(void* self) {
	auto&& __r = ((QListWidget*)self)->supportedDragActions();
	return (int)(__r);
}
void QListWidget_scrollToItem(void* self, void* item) {
	((QListWidget*)self)->scrollToItem((QListWidgetItem*)item);
}
void QListWidget_clear(void* self) {
	((QListWidget*)self)->clear();
}
void QListWidget_itemPressed(void* self, void* item) {
	((QListWidget*)self)->itemPressed((QListWidgetItem*)item);
}
void QListWidget_itemClicked(void* self, void* item) {
	((QListWidget*)self)->itemClicked((QListWidgetItem*)item);
}
void QListWidget_itemDoubleClicked(void* self, void* item) {
	((QListWidget*)self)->itemDoubleClicked((QListWidgetItem*)item);
}
void QListWidget_itemActivated(void* self, void* item) {
	((QListWidget*)self)->itemActivated((QListWidgetItem*)item);
}
void QListWidget_itemEntered(void* self, void* item) {
	((QListWidget*)self)->itemEntered((QListWidgetItem*)item);
}
void QListWidget_itemChanged(void* self, void* item) {
	((QListWidget*)self)->itemChanged((QListWidgetItem*)item);
}
void QListWidget_currentItemChanged(void* self, void* current, void* previous) {
	((QListWidget*)self)->currentItemChanged((QListWidgetItem*)current, (QListWidgetItem*)previous);
}
void QListWidget_currentTextChanged(void* self, const char* currentText, long long currentText_len) {
	QString currentText_qstr = QString::fromUtf8(currentText, currentText_len);
	((QListWidget*)self)->currentTextChanged(currentText_qstr);
}
void QListWidget_currentRowChanged(void* self, int currentRow) {
	((QListWidget*)self)->currentRowChanged(currentRow);
}
void QListWidget_itemSelectionChanged(void* self) {
	((QListWidget*)self)->itemSelectionChanged();
}
void QListWidget_resizeContents(void* self, int width, int height) {
	((QListWidget__Go*)self)->resizeContents(width, height);
}
void* QListWidget_contentsSize(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->contentsSize();
	return new QSize(__r);
}
void* QListWidget_rectForIndex(void* self, void* index) {
	auto&& __r = ((QListWidget__Go*)self)->rectForIndex(*(QModelIndex*)index);
	return new QRect(__r);
}
void QListWidget_setPositionForIndex(void* self, void* position, void* index) {
	((QListWidget__Go*)self)->setPositionForIndex(*(QPoint*)position, *(QModelIndex*)index);
}
int QListWidget_state(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->state();
	return (int)(__r);
}
void QListWidget_scheduleDelayedItemsLayout(void* self) {
	((QListWidget__Go*)self)->scheduleDelayedItemsLayout();
}
void QListWidget_executeDelayedItemsLayout(void* self) {
	((QListWidget__Go*)self)->executeDelayedItemsLayout();
}
void QListWidget_setDirtyRegion(void* self, void* region) {
	((QListWidget__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QListWidget_scrollDirtyRegion(void* self, int dx, int dy) {
	((QListWidget__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QListWidget_dirtyRegionOffset(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QListWidget_startAutoScroll(void* self) {
	((QListWidget__Go*)self)->startAutoScroll();
}
void QListWidget_stopAutoScroll(void* self) {
	((QListWidget__Go*)self)->stopAutoScroll();
}
void QListWidget_doAutoScroll(void* self) {
	((QListWidget__Go*)self)->doAutoScroll();
}
int QListWidget_dropIndicatorPosition(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QListWidget_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QListWidget__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QListWidget_setViewportMargins2(void* self, void* margins) {
	((QListWidget__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QListWidget_viewportMargins(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QListWidget_drawFrame(void* self, void* arg1) {
	((QListWidget__Go*)self)->drawFrame((QPainter*)arg1);
}
void QListWidget_setTabOrder(void* arg1, void* arg2) {
	QListWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QListWidget_mouseGrabber() {
	auto&& __r = QListWidget::mouseGrabber();
	return (void*)(__r);
}
void* QListWidget_keyboardGrabber() {
	auto&& __r = QListWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QListWidget_find(unsigned long long arg1) {
	auto&& __r = QListWidget::find(arg1);
	return (void*)(__r);
}
void* QListWidget_createWindowContainer(void* window) {
	auto&& __r = QListWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QListWidget_updateMicroFocus(void* self) {
	((QListWidget__Go*)self)->updateMicroFocus();
}
void QListWidget_create(void* self) {
	((QListWidget__Go*)self)->create();
}
void QListWidget_destroy(void* self) {
	((QListWidget__Go*)self)->destroy();
}
int QListWidget_focusNextChild(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->focusNextChild();
	return __r;
}
int QListWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QListWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QListWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QListWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QListWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QListWidget_sender(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QListWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QListWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QListWidget__Go*)self)->receivers(signal);
	return __r;
}
int QListWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QListWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QListWidget_devicePixelRatioFScale() {
	auto&& __r = QListWidget::devicePixelRatioFScale();
	return __r;
}
int QListWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QListWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QListWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QListWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QListWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QListWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QListWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QListWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QListWidget_sortItems2(void* self, int order) {
	((QListWidget*)self)->sortItems((Qt::SortOrder)(order));
}
void QListWidget_scrollToItem2(void* self, void* item, int hint) {
	((QListWidget*)self)->scrollToItem((QListWidgetItem*)item, (QAbstractItemView::ScrollHint)(hint));
}
void* QListWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QListWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QListWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QListWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QListWidget_updateMicroFocus2(void* self, int query) {
	((QListWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QListWidget_create2(void* self, unsigned long long arg1) {
	((QListWidget__Go*)self)->create(arg1);
}
void QListWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QListWidget__Go*)self)->create(arg1, initializeWindow);
}
void QListWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QListWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QListWidget_destroy2(void* self, int destroyWindow) {
	((QListWidget__Go*)self)->destroy(destroyWindow);
}
void QListWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QListWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QListWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QListWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QListWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QListWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QListWidget_metaObject_default(void* self) {
	auto&& __r = ((QListWidget*)self)->QListWidget::metaObject();
	return (void*)(__r);
}
void QListWidget_setSelectionModel_default(void* self, void* selectionModel) {
	((QListWidget*)self)->QListWidget::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QListWidget_dropEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QListWidget_event_default(void* self, void* e) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
char** QListWidget_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QListWidget_dropMimeData_default(void* self, int index, void* data, int action) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_dropMimeData(index, (QMimeData*)data, action);
	return __r;
}
int QListWidget_supportedDropActions_default(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_supportedDropActions();
	return (int)(__r);
}
void* QListWidget_visualRect_default(void* self, void* index) {
	auto&& __r = ((QListWidget*)self)->QListWidget::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QListWidget_scrollTo_default(void* self, void* index, int hint) {
	((QListWidget*)self)->QListWidget::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QListWidget_indexAt_default(void* self, void* p) {
	auto&& __r = ((QListWidget*)self)->QListWidget::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QListWidget_doItemsLayout_default(void* self) {
	((QListWidget*)self)->QListWidget::doItemsLayout();
}
void QListWidget_reset_default(void* self) {
	((QListWidget*)self)->QListWidget::reset();
}
void QListWidget_setRootIndex_default(void* self, void* index) {
	((QListWidget*)self)->QListWidget::setRootIndex(*(QModelIndex*)index);
}
void QListWidget_scrollContentsBy_default(void* self, int dx, int dy) {
	((QListWidget__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QListWidget_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QListWidget__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QListWidget_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QListWidget__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QListWidget_mouseMoveEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QListWidget_mouseReleaseEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QListWidget_wheelEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QListWidget_timerEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
void QListWidget_resizeEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QListWidget_dragMoveEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QListWidget_dragLeaveEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QListWidget_startDrag_default(void* self, int supportedActions) {
	((QListWidget__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QListWidget_initViewItemOption_default(void* self, void* option) {
	((QListWidget__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
void QListWidget_paintEvent_default(void* self, void* e) {
	((QListWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
int QListWidget_horizontalOffset_default(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QListWidget_verticalOffset_default(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void* QListWidget_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
void QListWidget_setSelection_default(void* self, void* rect, int command) {
	((QListWidget__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QListWidget_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QListWidget_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QListWidget_updateGeometries_default(void* self) {
	((QListWidget__Go*)self)->cppgoBase_updateGeometries();
}
int QListWidget_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QListWidget_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QListWidget__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QListWidget_currentChanged_default(void* self, void* current, void* previous) {
	((QListWidget__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void* QListWidget_viewportSizeHint_default(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QListWidget_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QListWidget*)self)->QListWidget::keyboardSearch(search_qstr);
}
int QListWidget_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QListWidget*)self)->QListWidget::sizeHintForRow(row);
	return __r;
}
int QListWidget_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QListWidget*)self)->QListWidget::sizeHintForColumn(column);
	return __r;
}
void* QListWidget_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QListWidget*)self)->QListWidget::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QListWidget_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QListWidget*)self)->QListWidget::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QListWidget_selectAll_default(void* self) {
	((QListWidget*)self)->QListWidget::selectAll();
}
void QListWidget_updateEditorData_default(void* self) {
	((QListWidget__Go*)self)->cppgoBase_updateEditorData();
}
void QListWidget_updateEditorGeometries_default(void* self) {
	((QListWidget__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QListWidget_verticalScrollbarAction_default(void* self, int action) {
	((QListWidget__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QListWidget_horizontalScrollbarAction_default(void* self, int action) {
	((QListWidget__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
void QListWidget_verticalScrollbarValueChanged_default(void* self, int value) {
	((QListWidget__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QListWidget_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QListWidget__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QListWidget_closeEditor_default(void* self, void* editor, int hint) {
	((QListWidget__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QListWidget_commitData_default(void* self, void* editor) {
	((QListWidget__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QListWidget_editorDestroyed_default(void* self, void* editor) {
	((QListWidget__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QListWidget_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QListWidget_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
int QListWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QListWidget_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QListWidget_mousePressEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QListWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QListWidget_dragEnterEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QListWidget_focusInEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QListWidget_focusOutEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QListWidget_keyPressEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QListWidget_inputMethodEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QListWidget_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QListWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QListWidget*)self)->QListWidget::minimumSizeHint();
	return new QSize(__r);
}
void* QListWidget_sizeHint_default(void* self) {
	auto&& __r = ((QListWidget*)self)->QListWidget::sizeHint();
	return new QSize(__r);
}
void QListWidget_setupViewport_default(void* self, void* viewport) {
	((QListWidget*)self)->QListWidget::setupViewport((QWidget*)viewport);
}
void QListWidget_contextMenuEvent_default(void* self, void* arg1) {
	((QListWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QListWidget_changeEvent_default(void* self, void* arg1) {
	((QListWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QListWidget_initStyleOption_default(void* self, void* option) {
	((QListWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QListWidget_devType_default(void* self) {
	auto&& __r = ((QListWidget*)self)->QListWidget::devType();
	return __r;
}
void QListWidget_setVisible_default(void* self, int visible) {
	((QListWidget*)self)->QListWidget::setVisible(visible);
}
int QListWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QListWidget*)self)->QListWidget::heightForWidth(arg1);
	return __r;
}
int QListWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QListWidget*)self)->QListWidget::hasHeightForWidth();
	return __r;
}
void* QListWidget_paintEngine_default(void* self) {
	auto&& __r = ((QListWidget*)self)->QListWidget::paintEngine();
	return (void*)(__r);
}
void QListWidget_keyReleaseEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QListWidget_enterEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QListWidget_leaveEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QListWidget_moveEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QListWidget_closeEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QListWidget_tabletEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QListWidget_actionEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QListWidget_showEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QListWidget_hideEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QListWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QListWidget_initPainter_default(void* self, void* painter) {
	((QListWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QListWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QListWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QListWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QListWidget_childEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QListWidget_customEvent_default(void* self, void* event) {
	((QListWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QListWidget_connectNotify_default(void* self, void* signal) {
	((QListWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QListWidget_disconnectNotify_default(void* self, void* signal) {
	((QListWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QListWidget_connect_itemPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemPressed, (QListWidget*)self, [__id](QListWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_itemClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemClicked, (QListWidget*)self, [__id](QListWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_itemDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemDoubleClicked, (QListWidget*)self, [__id](QListWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_itemActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemActivated, (QListWidget*)self, [__id](QListWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_itemEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemEntered, (QListWidget*)self, [__id](QListWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_itemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemChanged, (QListWidget*)self, [__id](QListWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_currentItemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::currentItemChanged, (QListWidget*)self, [__id](QListWidgetItem* current, QListWidgetItem* previous){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)current, (void*)previous);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_currentTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::currentTextChanged, (QListWidget*)self, [__id](const QString& currentText){
		QByteArray currentText_ba = currentText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)currentText_ba.constData(), currentText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_currentRowChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::currentRowChanged, (QListWidget*)self, [__id](int currentRow){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_itemSelectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::itemSelectionChanged, (QListWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_indexesMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::indexesMoved, (QListWidget*)self, [__id](const QList<QModelIndex>& indexes){
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (cppgo_sig_void_voidpp_longlong) ((void(*)(long long, void**, long long))cppgo_sig_void_voidpp_longlong)(__id, indexes_v.data(), (long long)indexes.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::pressed, (QListWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::clicked, (QListWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::doubleClicked, (QListWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::activated, (QListWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::entered, (QListWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::viewportEntered, (QListWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::iconSizeChanged, (QListWidget*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::windowTitleChanged, (QListWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::windowIconChanged, (QListWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::windowIconTextChanged, (QListWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::customContextMenuRequested, (QListWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::destroyed, (QListWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::destroyed, (QListWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListWidget*)self, &QListWidget::objectNameChanged, (QListWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
	__conns_QListWidget[__id] = __c;
}
void QListWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QListWidget);
		auto __it = __conns_QListWidget.find(id);
		if (__it == __conns_QListWidget.end()) return;
		__c = __it->second;
		__conns_QListWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QListWidget__Go_new(long long id) { return new QListWidget__Go(id); }
void QListWidget_delete(void* self) { delete (QListWidget*)self; }

}

#include <QtWidgets/qsplitter.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtGui/qevent.h>
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
#include "gen_qsplitterhandle.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSplitterHandle;
static std::mutex __connsMu_QSplitterHandle;

extern "C" {

void* QSplitterHandle_base(void* self) { return (void*)static_cast<QWidget*>((QSplitterHandle*)self); }
void QSplitterHandle_chain(void* self, void** out) {
	QSplitterHandle* __p = (QSplitterHandle*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QSplitterHandle_fromQObject(void* o) { return (void*)qobject_cast<QSplitterHandle*>((QObject*)o); }
void* QSplitterHandle_new(int o, void* parent) {
	return new QSplitterHandle((Qt::Orientation)(o), (QSplitter*)parent);
}
char* QSplitterHandle_tr(const char* s, long long* __out_len) {
	auto&& __r = QSplitterHandle::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSplitterHandle_setOrientation(void* self, int o) {
	((QSplitterHandle*)self)->setOrientation((Qt::Orientation)(o));
}
int QSplitterHandle_orientation(void* self) {
	auto&& __r = ((QSplitterHandle*)self)->orientation();
	return (int)(__r);
}
int QSplitterHandle_opaqueResize(void* self) {
	auto&& __r = ((QSplitterHandle*)self)->opaqueResize();
	return __r;
}
void* QSplitterHandle_splitter(void* self) {
	auto&& __r = ((QSplitterHandle*)self)->splitter();
	return (void*)(__r);
}
void QSplitterHandle_setTabOrder(void* arg1, void* arg2) {
	QSplitterHandle::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QSplitterHandle_mouseGrabber() {
	auto&& __r = QSplitterHandle::mouseGrabber();
	return (void*)(__r);
}
void* QSplitterHandle_keyboardGrabber() {
	auto&& __r = QSplitterHandle::keyboardGrabber();
	return (void*)(__r);
}
void* QSplitterHandle_find(unsigned long long arg1) {
	auto&& __r = QSplitterHandle::find(arg1);
	return (void*)(__r);
}
void* QSplitterHandle_createWindowContainer(void* window) {
	auto&& __r = QSplitterHandle::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void* QSplitterHandle_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSplitterHandle::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSplitterHandle_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSplitterHandle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
double QSplitterHandle_devicePixelRatioFScale() {
	auto&& __r = QSplitterHandle::devicePixelRatioFScale();
	return __r;
}
int QSplitterHandle_encodeMetricF(int metric, double value) {
	auto&& __r = QSplitterHandle::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
char* QSplitterHandle_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSplitterHandle::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSplitterHandle_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSplitterHandle::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSplitterHandle_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QSplitterHandle::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QSplitterHandle_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QSplitterHandle::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void* QSplitterHandle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSplitterHandle::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSplitterHandle_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSplitterHandle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QSplitterHandle_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::windowTitleChanged, (QSplitterHandle*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::windowIconChanged, (QSplitterHandle*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::windowIconTextChanged, (QSplitterHandle*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::customContextMenuRequested, (QSplitterHandle*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::destroyed, (QSplitterHandle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::destroyed, (QSplitterHandle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitterHandle*)self, &QSplitterHandle::objectNameChanged, (QSplitterHandle*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
	__conns_QSplitterHandle[__id] = __c;
}
void QSplitterHandle__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSplitterHandle);
		auto __it = __conns_QSplitterHandle.find(id);
		if (__it == __conns_QSplitterHandle.end()) return;
		__c = __it->second;
		__conns_QSplitterHandle.erase(__it);
	}
	QObject::disconnect(__c);
}
void QSplitterHandle_delete(void* self) { delete (QSplitterHandle*)self; }

}

#include <QtWidgets/qwidgetaction.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qactiongroup.h>
#include <QtGui/qicon.h>
#include <QtGui/qkeysequence.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qfont.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qwidgetaction.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QWidgetAction__Go : public QWidgetAction {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QWidgetAction__Go(long long id, A&&... a) : QWidgetAction(std::forward<A>(a)...), __goId(id) {}
	using QWidgetAction::createdWidgets;
	using QWidgetAction::sender;
	using QWidgetAction::senderSignalIndex;
	using QWidgetAction::receivers;
	using QWidgetAction::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QWidgetAction::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QWidgetAction::eventFilter(__a0, __a1); }
	QWidget* cppgoBase_createWidget(QWidget* __a0) { return QWidgetAction::createWidget(__a0); }
	void cppgoBase_deleteWidget(QWidget* __a0) { QWidgetAction::deleteWidget(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QWidgetAction::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QWidgetAction::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QWidgetAction::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QWidgetAction::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QWidgetAction::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QWidgetAction::metaObject();
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QWidgetAction::event(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QWidgetAction::eventFilter(arg1, arg2);
	}
	QWidget* createWidget(QWidget* parent) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8470300370605365395LL, (void*)parent, &__h); if (__h) return (QWidget*)__r; }
		return QWidgetAction::createWidget(parent);
	}
	void deleteWidget(QWidget* widget) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6493542048625819130LL, (void*)widget, &__h); if (__h) return; }
		QWidgetAction::deleteWidget(widget);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QWidgetAction::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QWidgetAction::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QWidgetAction::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QWidgetAction::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QWidgetAction::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QWidgetAction;
static std::mutex __connsMu_QWidgetAction;

extern "C" {

void* QWidgetAction_base(void* self) { return (void*)static_cast<QAction*>((QWidgetAction*)self); }
void QWidgetAction_chain(void* self, void** out) {
	QWidgetAction* __p = (QWidgetAction*)self;
	QAction* __b0 = static_cast<QAction*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QWidgetAction__set_override(void* self, int idx) { ((QWidgetAction__Go*)self)->__goOn[idx] = 1; }
void* QWidgetAction_fromQObject(void* o) { return (void*)qobject_cast<QWidgetAction*>((QObject*)o); }
void* QWidgetAction_new(long long __goid, void* parent) {
	return new QWidgetAction__Go(__goid, (QObject*)parent);
}
char* QWidgetAction_tr(const char* s, long long* __out_len) {
	auto&& __r = QWidgetAction::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidgetAction_setDefaultWidget(void* self, void* w) {
	((QWidgetAction*)self)->setDefaultWidget((QWidget*)w);
}
void* QWidgetAction_defaultWidget(void* self) {
	auto&& __r = ((QWidgetAction*)self)->defaultWidget();
	return (void*)(__r);
}
void* QWidgetAction_requestWidget(void* self, void* parent) {
	auto&& __r = ((QWidgetAction*)self)->requestWidget((QWidget*)parent);
	return (void*)(__r);
}
void QWidgetAction_releaseWidget(void* self, void* widget) {
	((QWidgetAction*)self)->releaseWidget((QWidget*)widget);
}
void** QWidgetAction_createdWidgets(void* self, long long* __out_len) {
	auto&& __r = ((QWidgetAction__Go*)self)->createdWidgets();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QWidgetAction_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QWidgetAction::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QWidgetAction_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QWidgetAction::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QWidgetAction_sender(void* self) {
	auto&& __r = ((QWidgetAction__Go*)self)->sender();
	return (void*)(__r);
}
int QWidgetAction_senderSignalIndex(void* self) {
	auto&& __r = ((QWidgetAction__Go*)self)->senderSignalIndex();
	return __r;
}
int QWidgetAction_receivers(void* self, const char* signal) {
	auto&& __r = ((QWidgetAction__Go*)self)->receivers(signal);
	return __r;
}
int QWidgetAction_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QWidgetAction__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QWidgetAction_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QWidgetAction::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWidgetAction_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QWidgetAction::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QWidgetAction_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QWidgetAction::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QWidgetAction_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QWidgetAction::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QWidgetAction_metaObject_default(void* self) {
	auto&& __r = ((QWidgetAction*)self)->QWidgetAction::metaObject();
	return (void*)(__r);
}
int QWidgetAction_event_default(void* self, void* arg1) {
	auto&& __r = ((QWidgetAction__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
int QWidgetAction_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QWidgetAction__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
void* QWidgetAction_createWidget_default(void* self, void* parent) {
	auto&& __r = ((QWidgetAction__Go*)self)->cppgoBase_createWidget((QWidget*)parent);
	return (void*)(__r);
}
void QWidgetAction_deleteWidget_default(void* self, void* widget) {
	((QWidgetAction__Go*)self)->cppgoBase_deleteWidget((QWidget*)widget);
}
void QWidgetAction_timerEvent_default(void* self, void* event) {
	((QWidgetAction__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QWidgetAction_childEvent_default(void* self, void* event) {
	((QWidgetAction__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QWidgetAction_customEvent_default(void* self, void* event) {
	((QWidgetAction__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QWidgetAction_connectNotify_default(void* self, void* signal) {
	((QWidgetAction__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QWidgetAction_disconnectNotify_default(void* self, void* signal) {
	((QWidgetAction__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QWidgetAction_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::changed, (QWidgetAction*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::enabledChanged, (QWidgetAction*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_checkableChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::checkableChanged, (QWidgetAction*)self, [__id](bool checkable){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checkable);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::visibleChanged, (QWidgetAction*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_triggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::triggered, (QWidgetAction*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_triggered2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::triggered, (QWidgetAction*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_hovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::hovered, (QWidgetAction*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::toggled, (QWidgetAction*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::destroyed, (QWidgetAction*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::destroyed, (QWidgetAction*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidgetAction*)self, &QWidgetAction::objectNameChanged, (QWidgetAction*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
	__conns_QWidgetAction[__id] = __c;
}
void QWidgetAction__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QWidgetAction);
		auto __it = __conns_QWidgetAction.find(id);
		if (__it == __conns_QWidgetAction.end()) return;
		__c = __it->second;
		__conns_QWidgetAction.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QWidgetAction__Go_new(long long id, void* parent) { return new QWidgetAction__Go(id, (QObject*)parent); }
void QWidgetAction_delete(void* self) { delete (QWidgetAction*)self; }

}

