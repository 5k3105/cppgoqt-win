#include <QtCore/qarraydata.h>
#include "gen_qarraydata.h"

extern "C" {

long long QArrayData_allocatedCapacity(void* self) {
	auto&& __r = ((QArrayData*)self)->allocatedCapacity();
	return __r;
}
long long QArrayData_constAllocatedCapacity(void* self) {
	auto&& __r = ((QArrayData*)self)->constAllocatedCapacity();
	return __r;
}
int QArrayData_ref(void* self) {
	auto&& __r = ((QArrayData*)self)->ref();
	return __r;
}
int QArrayData_deref(void* self) {
	auto&& __r = ((QArrayData*)self)->deref();
	return __r;
}
int QArrayData_isShared(void* self) {
	auto&& __r = ((QArrayData*)self)->isShared();
	return __r;
}
int QArrayData_needsDetach(void* self) {
	auto&& __r = ((QArrayData*)self)->needsDetach();
	return __r;
}
long long QArrayData_detachCapacity(void* self, long long newSize) {
	auto&& __r = ((QArrayData*)self)->detachCapacity(newSize);
	return __r;
}
void QArrayData_deallocate(void* data, long long objectSize, long long alignment) {
	QArrayData::deallocate((QArrayData*)data, objectSize, alignment);
}
int QArrayData_flags_get(void* self) {
	auto __r = ((QArrayData*)self)->flags;
	return (int)(__r);
}
void QArrayData_flags_set(void* self, int flags) {
	((QArrayData*)self)->flags = QFlags<QArrayData::ArrayOption>(flags);
}
long long QArrayData_alloc_get(void* self) {
	auto __r = ((QArrayData*)self)->alloc;
	return __r;
}
void QArrayData_alloc_set(void* self, long long alloc) {
	((QArrayData*)self)->alloc = alloc;
}
void QArrayData_delete(void* self) { delete (QArrayData*)self; }

}

#include <QtCore/qbindingstorage.h>
#include <QtCore/qpropertyprivate.h>
#include "gen_qbindingstorage.h"

extern "C" {

void* QBindingStorage_new() {
	return new QBindingStorage();
}
int QBindingStorage_isEmpty(void* self) {
	auto&& __r = ((QBindingStorage*)self)->isEmpty();
	return __r;
}
int QBindingStorage_isValid(void* self) {
	auto&& __r = ((QBindingStorage*)self)->isValid();
	return __r;
}
void QBindingStorage_registerDependency(void* self, void* data) {
	((QBindingStorage*)self)->registerDependency((QUntypedPropertyData*)data);
}
void QBindingStorage_delete(void* self) { delete (QBindingStorage*)self; }

}

#include <QtCore/qvariant.h>
#include "gen_qvariantconstpointer.h"

extern "C" {

void* QVariantConstPointer_new(void* variant) {
	return new QVariantConstPointer(*(QVariant*)variant);
}
void* QVariantConstPointer_new2(void* arg1) {
	return new QVariantConstPointer(*(QVariantConstPointer*)arg1);
}
void* QVariantConstPointer_OperatorMultiply(void* self) {
	auto&& __r = ((QVariantConstPointer*)self)->operator*();
	return new QVariant(__r);
}
void* QVariantConstPointer_OperatorMinusGreater(void* self) {
	auto&& __r = ((QVariantConstPointer*)self)->operator->();
	return (void*)(__r);
}
void* QVariantConstPointer_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QVariantConstPointer*)self)->operator=(*(QVariantConstPointer*)arg1);
	return (void*)&(__r);
}
void QVariantConstPointer_delete(void* self) { delete (QVariantConstPointer*)self; }

}

#include <QtCore/qbitarray.h>
#include "gen_qbitref.h"

extern "C" {

int QBitRef_ToBool(void* self) {
	auto&& __r = static_cast<bool>(*(QBitRef*)self);
	return __r;
}
int QBitRef_OperatorNot(void* self) {
	auto&& __r = ((QBitRef*)self)->operator!();
	return __r;
}
void* QBitRef_OperatorAssign(void* self, void* val) {
	auto&& __r = ((QBitRef*)self)->operator=(*(QBitRef*)val);
	return (void*)&(__r);
}
void* QBitRef_OperatorAssign2(void* self, int val) {
	auto&& __r = ((QBitRef*)self)->operator=(val);
	return (void*)&(__r);
}
void QBitRef_delete(void* self) { delete (QBitRef*)self; }

}

#include <QtCore/qcbormap.h>
#include <QtCore/qcborvalue.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qjsonobject.h>
#include "gen_qcbormap.h"

extern "C" {

void* QCborMap_new() {
	return new QCborMap();
}
void* QCborMap_new2(void* other) {
	return new QCborMap(*(QCborMap*)other);
}
void* QCborMap_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCborMap*)self)->operator=(*(QCborMap*)other);
	return (void*)&(__r);
}
void QCborMap_swap(void* self, void* other) {
	((QCborMap*)self)->swap(*(QCborMap*)other);
}
void* QCborMap_toCborValue(void* self) {
	auto&& __r = ((QCborMap*)self)->toCborValue();
	return new QCborValue(__r);
}
long long QCborMap_size(void* self) {
	auto&& __r = ((QCborMap*)self)->size();
	return __r;
}
int QCborMap_isEmpty(void* self) {
	auto&& __r = ((QCborMap*)self)->isEmpty();
	return __r;
}
void QCborMap_clear(void* self) {
	((QCborMap*)self)->clear();
}
void** QCborMap_keys(void* self, long long* __out_len) {
	auto&& __r = ((QCborMap*)self)->keys();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QCborValue(__e);
	return __arr;
}
void* QCborMap_value(void* self, long long key) {
	auto&& __r = ((QCborMap*)self)->value(key);
	return new QCborValue(__r);
}
void* QCborMap_value2(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->value(*(QLatin1String*)key);
	return new QCborValue(__r);
}
void* QCborMap_value3(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborMap*)self)->value(key_qstr);
	return new QCborValue(__r);
}
void* QCborMap_value4(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->value(*(QCborValue*)key);
	return new QCborValue(__r);
}
void* QCborMap_OperatorSubscript(void* self, long long key) {
	auto&& __r = ((QCborMap*)self)->operator[](key);
	return new QCborValue(__r);
}
void* QCborMap_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->operator[](*(QLatin1String*)key);
	return new QCborValue(__r);
}
void* QCborMap_OperatorSubscript3(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborMap*)self)->operator[](key_qstr);
	return new QCborValue(__r);
}
void* QCborMap_OperatorSubscript4(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->operator[](*(QCborValue*)key);
	return new QCborValue(__r);
}
void* QCborMap_OperatorSubscript5(void* self, long long key) {
	auto&& __r = ((QCborMap*)self)->operator[](key);
	return new QCborValueRef(__r);
}
void* QCborMap_OperatorSubscript6(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->operator[](*(QLatin1String*)key);
	return new QCborValueRef(__r);
}
void* QCborMap_OperatorSubscript7(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborMap*)self)->operator[](key_qstr);
	return new QCborValueRef(__r);
}
void* QCborMap_OperatorSubscript8(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->operator[](*(QCborValue*)key);
	return new QCborValueRef(__r);
}
void* QCborMap_take(void* self, long long key) {
	auto&& __r = ((QCborMap*)self)->take(key);
	return new QCborValue(__r);
}
void* QCborMap_take2(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->take(*(QLatin1String*)key);
	return new QCborValue(__r);
}
void* QCborMap_take3(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborMap*)self)->take(key_qstr);
	return new QCborValue(__r);
}
void* QCborMap_take4(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->take(*(QCborValue*)key);
	return new QCborValue(__r);
}
void QCborMap_remove(void* self, long long key) {
	((QCborMap*)self)->remove(key);
}
void QCborMap_remove2(void* self, void* key) {
	((QCborMap*)self)->remove(*(QLatin1String*)key);
}
void QCborMap_remove3(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QCborMap*)self)->remove(key_qstr);
}
void QCborMap_remove4(void* self, void* key) {
	((QCborMap*)self)->remove(*(QCborValue*)key);
}
int QCborMap_contains(void* self, long long key) {
	auto&& __r = ((QCborMap*)self)->contains(key);
	return __r;
}
int QCborMap_contains2(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->contains(*(QLatin1String*)key);
	return __r;
}
int QCborMap_contains3(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborMap*)self)->contains(key_qstr);
	return __r;
}
int QCborMap_contains4(void* self, void* key) {
	auto&& __r = ((QCborMap*)self)->contains(*(QCborValue*)key);
	return __r;
}
int QCborMap_compare(void* self, void* other) {
	auto&& __r = ((QCborMap*)self)->compare(*(QCborMap*)other);
	return __r;
}
int QCborMap_empty(void* self) {
	auto&& __r = ((QCborMap*)self)->empty();
	return __r;
}
void* QCborMap_fromVariantMap(char** map__keys, void* map__vals, long long map__len) {
	QMap<QString, QVariant> map__map;
	void** map__vv = (void**)map__vals;
	for (long long __i = 0; __i < map__len; __i++) map__map.insert(QString::fromUtf8(map__keys[__i]), *(QVariant*)map__vv[__i]);
	auto&& __r = QCborMap::fromVariantMap(map__map);
	return new QCborMap(__r);
}
void* QCborMap_fromVariantHash(char** hash_keys, void* hash_vals, long long hash_len) {
	QHash<QString, QVariant> hash_map;
	void** hash_vv = (void**)hash_vals;
	for (long long __i = 0; __i < hash_len; __i++) hash_map.insert(QString::fromUtf8(hash_keys[__i]), *(QVariant*)hash_vv[__i]);
	auto&& __r = QCborMap::fromVariantHash(hash_map);
	return new QCborMap(__r);
}
void* QCborMap_fromJsonObject(void* o) {
	auto&& __r = QCborMap::fromJsonObject(*(QJsonObject*)o);
	return new QCborMap(__r);
}
void* QCborMap_toJsonObject(void* self) {
	auto&& __r = ((QCborMap*)self)->toJsonObject();
	return new QJsonObject(__r);
}
int QCborMap_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) == (*(QCborMap*)rhs);
	return __r;
}
int QCborMap_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) != (*(QCborMap*)rhs);
	return __r;
}
int QCborMap_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) < (*(QCborMap*)rhs);
	return __r;
}
int QCborMap_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) > (*(QCborMap*)rhs);
	return __r;
}
int QCborMap_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) <= (*(QCborMap*)rhs);
	return __r;
}
int QCborMap_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) >= (*(QCborMap*)rhs);
	return __r;
}
int QCborMap_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) == (*(QCborValue*)rhs);
	return __r;
}
int QCborMap_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) != (*(QCborValue*)rhs);
	return __r;
}
int QCborMap_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) < (*(QCborValue*)rhs);
	return __r;
}
int QCborMap_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) > (*(QCborValue*)rhs);
	return __r;
}
int QCborMap_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) <= (*(QCborValue*)rhs);
	return __r;
}
int QCborMap_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) >= (*(QCborValue*)rhs);
	return __r;
}
int QCborMap_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) == (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborMap_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) != (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborMap_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) < (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborMap_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) > (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborMap_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) <= (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborMap_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborMap*)self) >= (*(QCborValueConstRef*)rhs);
	return __r;
}
void QCborMap_delete(void* self) { delete (QCborMap*)self; }

}

#include <QtCore/qdirlisting.h>
#include "gen_qdirlisting__sentinel.h"

extern "C" {

int QDirListing__sentinel_OperatorEqual(void* self, void* arg1) {
	auto&& __r = (*(QDirListing::sentinel*)self) == (*(QDirListing::sentinel*)arg1);
	return __r;
}
int QDirListing__sentinel_OperatorNotEqual(void* self, void* arg1) {
	auto&& __r = (*(QDirListing::sentinel*)self) != (*(QDirListing::sentinel*)arg1);
	return __r;
}
void QDirListing__sentinel_delete(void* self) { delete (QDirListing::sentinel*)self; }

}

#include <QtCore/qjsonparseerror.h>
#include <QtCore/qstring.h>
#include "gen_qjsonparseerror.h"

extern "C" {

char* QJsonParseError_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QJsonParseError*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QJsonParseError_offset_get(void* self) {
	auto __r = ((QJsonParseError*)self)->offset;
	return __r;
}
void QJsonParseError_offset_set(void* self, int offset) {
	((QJsonParseError*)self)->offset = offset;
}
int QJsonParseError_error_get(void* self) {
	auto __r = ((QJsonParseError*)self)->error;
	return (int)(__r);
}
void QJsonParseError_error_set(void* self, int error_) {
	((QJsonParseError*)self)->error = (QJsonParseError::ParseError)(error_);
}
void QJsonParseError_delete(void* self) { delete (QJsonParseError*)self; }

}

#include <QtCore/qmetaobject.h>
#include <QtCore/qobjectdefs.h>
#include "gen_qmetaclassinfo.h"

extern "C" {

void* QMetaClassInfo_new() {
	return new QMetaClassInfo();
}
void* QMetaClassInfo_new2(void* arg1) {
	return new QMetaClassInfo(*(QMetaClassInfo*)arg1);
}
const char* QMetaClassInfo_name(void* self) {
	auto&& __r = ((QMetaClassInfo*)self)->name();
	return __r;
}
const char* QMetaClassInfo_value(void* self) {
	auto&& __r = ((QMetaClassInfo*)self)->value();
	return __r;
}
void* QMetaClassInfo_enclosingMetaObject(void* self) {
	auto&& __r = ((QMetaClassInfo*)self)->enclosingMetaObject();
	return (void*)(__r);
}
void QMetaClassInfo_delete(void* self) { delete (QMetaClassInfo*)self; }

}

#include <QtCore/qprocess.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstringview.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qprocess.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QProcess__Go : public QProcess {
public:
	long long __goId;
	unsigned char __goOn[22] = {};
	template<typename... A> QProcess__Go(long long id, A&&... a) : QProcess(std::forward<A>(a)...), __goId(id) {}
	using QProcess::setProcessState;
	using QProcess::setOpenMode;
	using QProcess::setErrorString;
	using QProcess::sender;
	using QProcess::senderSignalIndex;
	using QProcess::receivers;
	using QProcess::isSignalConnected;
	qint64 cppgoBase_skipData(qint64 __a0) { return QProcess::skipData(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QProcess::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QProcess::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QProcess::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QProcess::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QProcess::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QProcess::metaObject();
	}
	bool open(QFlags<QIODeviceBase::OpenModeFlag> mode) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -554057428128861975LL, (int)(mode), &__h); if (__h) return __r; }
		return QProcess::open(mode);
	}
	bool waitForReadyRead(int msecs) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 8356976193102722774LL, msecs, &__h); if (__h) return __r; }
		return QProcess::waitForReadyRead(msecs);
	}
	bool waitForBytesWritten(int msecs) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 7726833620054767889LL, msecs, &__h); if (__h) return __r; }
		return QProcess::waitForBytesWritten(msecs);
	}
	qint64 bytesToWrite() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 1024747515116322948LL, &__h); if (__h) return __r; }
		return QProcess::bytesToWrite();
	}
	bool isSequential() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 8888098400087344862LL, &__h); if (__h) return __r; }
		return QProcess::isSequential();
	}
	void close() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -7006921981232612925LL, &__h); if (__h) return; }
		QProcess::close();
	}
	qint64 pos() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 8618582476283917417LL, &__h); if (__h) return __r; }
		return QProcess::pos();
	}
	qint64 size() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 5614464919154503228LL, &__h); if (__h) return __r; }
		return QProcess::size();
	}
	bool seek(qint64 pos) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, 3241140482992819217LL, pos, &__h); if (__h) return __r; }
		return QProcess::seek(pos);
	}
	bool atEnd() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1487528724786208339LL, &__h); if (__h) return __r; }
		return QProcess::atEnd();
	}
	bool reset() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1086335023529244512LL, &__h); if (__h) return __r; }
		return QProcess::reset();
	}
	qint64 bytesAvailable() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 4814449630054007943LL, &__h); if (__h) return __r; }
		return QProcess::bytesAvailable();
	}
	bool canReadLine() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1250682508954504761LL, &__h); if (__h) return __r; }
		return QProcess::canReadLine();
	}
	qint64 skipData(qint64 maxSize) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_longlong_longlong) { long long __r = ((long long(*)(long long, long long, long long, int*))cppgo_virt_longlong_longlong)(__goId, 8178738247667652082LL, maxSize, &__h); if (__h) return __r; }
		return QProcess::skipData(maxSize);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QProcess::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QProcess::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QProcess::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QProcess::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QProcess::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QProcess::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QProcess::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QProcess;
static std::mutex __connsMu_QProcess;

extern "C" {

void* QProcess_base(void* self) { return (void*)static_cast<QIODevice*>((QProcess*)self); }
void QProcess_chain(void* self, void** out) {
	QProcess* __p = (QProcess*)self;
	QIODevice* __b0 = static_cast<QIODevice*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QProcess__set_override(void* self, int idx) { ((QProcess__Go*)self)->__goOn[idx] = 1; }
void* QProcess_fromQObject(void* o) { return (void*)qobject_cast<QProcess*>((QObject*)o); }
void* QProcess_new(long long __goid) {
	return new QProcess__Go(__goid);
}
void* QProcess_new2(long long __goid, void* parent) {
	return new QProcess__Go(__goid, (QObject*)parent);
}
char* QProcess_tr(const char* s, long long* __out_len) {
	auto&& __r = QProcess::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProcess_start(void* self, const char* program, long long program_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	((QProcess*)self)->start(program_qstr);
}
void QProcess_start2(void* self) {
	((QProcess*)self)->start();
}
void QProcess_startCommand(void* self, const char* command, long long command_len) {
	QString command_qstr = QString::fromUtf8(command, command_len);
	((QProcess*)self)->startCommand(command_qstr);
}
int QProcess_startDetached(void* self) {
	auto&& __r = ((QProcess*)self)->startDetached();
	return __r;
}
char* QProcess_program(void* self, long long* __out_len) {
	auto&& __r = ((QProcess*)self)->program();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProcess_setProgram(void* self, const char* program, long long program_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	((QProcess*)self)->setProgram(program_qstr);
}
char** QProcess_arguments(void* self, long long* __out_len) {
	auto&& __r = ((QProcess*)self)->arguments();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QProcess_setArguments(void* self, char** arguments, long long arguments_len) {
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	((QProcess*)self)->setArguments(arguments_list);
}
int QProcess_processChannelMode(void* self) {
	auto&& __r = ((QProcess*)self)->processChannelMode();
	return (int)(__r);
}
void QProcess_setProcessChannelMode(void* self, int mode) {
	((QProcess*)self)->setProcessChannelMode((QProcess::ProcessChannelMode)(mode));
}
int QProcess_inputChannelMode(void* self) {
	auto&& __r = ((QProcess*)self)->inputChannelMode();
	return (int)(__r);
}
void QProcess_setInputChannelMode(void* self, int mode) {
	((QProcess*)self)->setInputChannelMode((QProcess::InputChannelMode)(mode));
}
int QProcess_readChannel(void* self) {
	auto&& __r = ((QProcess*)self)->readChannel();
	return (int)(__r);
}
void QProcess_setReadChannel(void* self, int channel) {
	((QProcess*)self)->setReadChannel((QProcess::ProcessChannel)(channel));
}
void QProcess_closeReadChannel(void* self, int channel) {
	((QProcess*)self)->closeReadChannel((QProcess::ProcessChannel)(channel));
}
void QProcess_closeWriteChannel(void* self) {
	((QProcess*)self)->closeWriteChannel();
}
void QProcess_setStandardInputFile(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QProcess*)self)->setStandardInputFile(fileName_qstr);
}
void QProcess_setStandardOutputFile(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QProcess*)self)->setStandardOutputFile(fileName_qstr);
}
void QProcess_setStandardErrorFile(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QProcess*)self)->setStandardErrorFile(fileName_qstr);
}
void QProcess_setStandardOutputProcess(void* self, void* destination) {
	((QProcess*)self)->setStandardOutputProcess((QProcess*)destination);
}
char* QProcess_nativeArguments(void* self, long long* __out_len) {
	auto&& __r = ((QProcess*)self)->nativeArguments();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProcess_setNativeArguments(void* self, const char* arguments, long long arguments_len) {
	QString arguments_qstr = QString::fromUtf8(arguments, arguments_len);
	((QProcess*)self)->setNativeArguments(arguments_qstr);
}
char* QProcess_workingDirectory(void* self, long long* __out_len) {
	auto&& __r = ((QProcess*)self)->workingDirectory();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProcess_setWorkingDirectory(void* self, const char* dir, long long dir_len) {
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	((QProcess*)self)->setWorkingDirectory(dir_qstr);
}
void QProcess_setEnvironment(void* self, char** environment, long long environment_len) {
	QList<QString> environment_list; environment_list.reserve(environment_len);
	for (long long __i = 0; __i < environment_len; __i++) environment_list.append(QString::fromUtf8(environment[__i]));
	((QProcess*)self)->setEnvironment(environment_list);
}
char** QProcess_environment(void* self, long long* __out_len) {
	auto&& __r = ((QProcess*)self)->environment();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QProcess_setProcessEnvironment(void* self, void* environment) {
	((QProcess*)self)->setProcessEnvironment(*(QProcessEnvironment*)environment);
}
void* QProcess_processEnvironment(void* self) {
	auto&& __r = ((QProcess*)self)->processEnvironment();
	return new QProcessEnvironment(__r);
}
int QProcess_error(void* self) {
	auto&& __r = ((QProcess*)self)->error();
	return (int)(__r);
}
int QProcess_state(void* self) {
	auto&& __r = ((QProcess*)self)->state();
	return (int)(__r);
}
long long QProcess_processId(void* self) {
	auto&& __r = ((QProcess*)self)->processId();
	return __r;
}
int QProcess_waitForStarted(void* self) {
	auto&& __r = ((QProcess*)self)->waitForStarted();
	return __r;
}
int QProcess_waitForFinished(void* self) {
	auto&& __r = ((QProcess*)self)->waitForFinished();
	return __r;
}
void* QProcess_readAllStandardOutput(void* self) {
	auto&& __r = ((QProcess*)self)->readAllStandardOutput();
	return new QByteArray(__r);
}
void* QProcess_readAllStandardError(void* self) {
	auto&& __r = ((QProcess*)self)->readAllStandardError();
	return new QByteArray(__r);
}
int QProcess_exitCode(void* self) {
	auto&& __r = ((QProcess*)self)->exitCode();
	return __r;
}
int QProcess_exitStatus(void* self) {
	auto&& __r = ((QProcess*)self)->exitStatus();
	return (int)(__r);
}
int QProcess_execute(const char* program, long long program_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	auto&& __r = QProcess::execute(program_qstr);
	return __r;
}
int QProcess_startDetached2(const char* program, long long program_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	auto&& __r = QProcess::startDetached(program_qstr);
	return __r;
}
char** QProcess_systemEnvironment(long long* __out_len) {
	auto&& __r = QProcess::systemEnvironment();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QProcess_nullDevice(long long* __out_len) {
	auto&& __r = QProcess::nullDevice();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QProcess_splitCommand(void* command, long long* __out_len) {
	auto&& __r = QProcess::splitCommand(*(QStringView*)command);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QProcess_terminate(void* self) {
	((QProcess*)self)->terminate();
}
void QProcess_kill(void* self) {
	((QProcess*)self)->kill();
}
void QProcess_finished(void* self, int exitCode) {
	((QProcess*)self)->finished(exitCode);
}
void QProcess_errorOccurred(void* self, int error_) {
	((QProcess*)self)->errorOccurred((QProcess::ProcessError)(error_));
}
void QProcess_setProcessState(void* self, int state) {
	((QProcess__Go*)self)->setProcessState((QProcess::ProcessState)(state));
}
void QProcess_setOpenMode(void* self, int openMode) {
	((QProcess__Go*)self)->setOpenMode(QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void QProcess_setErrorString(void* self, const char* errorString, long long errorString_len) {
	QString errorString_qstr = QString::fromUtf8(errorString, errorString_len);
	((QProcess__Go*)self)->setErrorString(errorString_qstr);
}
void* QProcess_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QProcess::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QProcess_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QProcess::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QProcess_sender(void* self) {
	auto&& __r = ((QProcess__Go*)self)->sender();
	return (void*)(__r);
}
int QProcess_senderSignalIndex(void* self) {
	auto&& __r = ((QProcess__Go*)self)->senderSignalIndex();
	return __r;
}
int QProcess_receivers(void* self, const char* signal) {
	auto&& __r = ((QProcess__Go*)self)->receivers(signal);
	return __r;
}
int QProcess_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QProcess__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QProcess_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QProcess::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QProcess_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QProcess::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProcess_start3(void* self, const char* program, long long program_len, char** arguments, long long arguments_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	((QProcess*)self)->start(program_qstr, arguments_list);
}
void QProcess_start4(void* self, const char* program, long long program_len, char** arguments, long long arguments_len, int mode) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	((QProcess*)self)->start(program_qstr, arguments_list, QFlags<QIODeviceBase::OpenModeFlag>(mode));
}
void QProcess_start5(void* self, int mode) {
	((QProcess*)self)->start(QFlags<QIODeviceBase::OpenModeFlag>(mode));
}
void QProcess_startCommand2(void* self, const char* command, long long command_len, int mode) {
	QString command_qstr = QString::fromUtf8(command, command_len);
	((QProcess*)self)->startCommand(command_qstr, QFlags<QIODeviceBase::OpenModeFlag>(mode));
}
int QProcess_startDetached3(void* self, long long* pid) {
	qint64 pid_local{};
	auto&& __r = ((QProcess*)self)->startDetached(&pid_local);
	*pid = pid_local;
	return __r;
}
void QProcess_setStandardOutputFile2(void* self, const char* fileName, long long fileName_len, int mode) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QProcess*)self)->setStandardOutputFile(fileName_qstr, QFlags<QIODeviceBase::OpenModeFlag>(mode));
}
void QProcess_setStandardErrorFile2(void* self, const char* fileName, long long fileName_len, int mode) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QProcess*)self)->setStandardErrorFile(fileName_qstr, QFlags<QIODeviceBase::OpenModeFlag>(mode));
}
int QProcess_waitForStarted2(void* self, int msecs) {
	auto&& __r = ((QProcess*)self)->waitForStarted(msecs);
	return __r;
}
int QProcess_waitForFinished2(void* self, int msecs) {
	auto&& __r = ((QProcess*)self)->waitForFinished(msecs);
	return __r;
}
int QProcess_execute2(const char* program, long long program_len, char** arguments, long long arguments_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	auto&& __r = QProcess::execute(program_qstr, arguments_list);
	return __r;
}
int QProcess_startDetached4(const char* program, long long program_len, char** arguments, long long arguments_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	auto&& __r = QProcess::startDetached(program_qstr, arguments_list);
	return __r;
}
int QProcess_startDetached5(const char* program, long long program_len, char** arguments, long long arguments_len, const char* workingDirectory, long long workingDirectory_len) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	QString workingDirectory_qstr = QString::fromUtf8(workingDirectory, workingDirectory_len);
	auto&& __r = QProcess::startDetached(program_qstr, arguments_list, workingDirectory_qstr);
	return __r;
}
int QProcess_startDetached6(const char* program, long long program_len, char** arguments, long long arguments_len, const char* workingDirectory, long long workingDirectory_len, long long* pid) {
	QString program_qstr = QString::fromUtf8(program, program_len);
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	QString workingDirectory_qstr = QString::fromUtf8(workingDirectory, workingDirectory_len);
	qint64 pid_local{};
	auto&& __r = QProcess::startDetached(program_qstr, arguments_list, workingDirectory_qstr, &pid_local);
	*pid = pid_local;
	return __r;
}
void QProcess_finished2(void* self, int exitCode, int exitStatus) {
	((QProcess*)self)->finished(exitCode, (QProcess::ExitStatus)(exitStatus));
}
void* QProcess_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QProcess::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QProcess_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QProcess::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QProcess_metaObject_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::metaObject();
	return (void*)(__r);
}
int QProcess_open_default(void* self, int mode) {
	auto&& __r = ((QProcess*)self)->QProcess::open(QFlags<QIODeviceBase::OpenModeFlag>(mode));
	return __r;
}
int QProcess_waitForReadyRead_default(void* self, int msecs) {
	auto&& __r = ((QProcess*)self)->QProcess::waitForReadyRead(msecs);
	return __r;
}
int QProcess_waitForBytesWritten_default(void* self, int msecs) {
	auto&& __r = ((QProcess*)self)->QProcess::waitForBytesWritten(msecs);
	return __r;
}
long long QProcess_bytesToWrite_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::bytesToWrite();
	return __r;
}
int QProcess_isSequential_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::isSequential();
	return __r;
}
void QProcess_close_default(void* self) {
	((QProcess*)self)->QProcess::close();
}
long long QProcess_pos_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::pos();
	return __r;
}
long long QProcess_size_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::size();
	return __r;
}
int QProcess_seek_default(void* self, long long pos) {
	auto&& __r = ((QProcess*)self)->QProcess::seek(pos);
	return __r;
}
int QProcess_atEnd_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::atEnd();
	return __r;
}
int QProcess_reset_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::reset();
	return __r;
}
long long QProcess_bytesAvailable_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::bytesAvailable();
	return __r;
}
int QProcess_canReadLine_default(void* self) {
	auto&& __r = ((QProcess*)self)->QProcess::canReadLine();
	return __r;
}
long long QProcess_skipData_default(void* self, long long maxSize) {
	auto&& __r = ((QProcess__Go*)self)->cppgoBase_skipData(maxSize);
	return __r;
}
int QProcess_event_default(void* self, void* event) {
	auto&& __r = ((QProcess*)self)->QProcess::event((QEvent*)event);
	return __r;
}
int QProcess_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QProcess*)self)->QProcess::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QProcess_timerEvent_default(void* self, void* event) {
	((QProcess__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QProcess_childEvent_default(void* self, void* event) {
	((QProcess__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QProcess_customEvent_default(void* self, void* event) {
	((QProcess__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QProcess_connectNotify_default(void* self, void* signal) {
	((QProcess__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QProcess_disconnectNotify_default(void* self, void* signal) {
	((QProcess__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QProcess_connect_started(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::started, (QProcess*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::finished, (QProcess*)self, [__id](int exitCode, QProcess::ExitStatus exitStatus){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, exitCode);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_finished2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::finished, (QProcess*)self, [__id](int exitCode, QProcess::ExitStatus exitStatus){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, exitCode, (int)(exitStatus));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_errorOccurred(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::errorOccurred, (QProcess*)self, [__id](QProcess::ProcessError error_){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(error_));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::stateChanged, (QProcess*)self, [__id](QProcess::ProcessState state){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(state));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_readyReadStandardOutput(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::readyReadStandardOutput, (QProcess*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_readyReadStandardError(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::readyReadStandardError, (QProcess*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::readyRead, (QProcess*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::channelReadyRead, (QProcess*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::bytesWritten, (QProcess*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::channelBytesWritten, (QProcess*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::aboutToClose, (QProcess*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::readChannelFinished, (QProcess*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::destroyed, (QProcess*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::destroyed, (QProcess*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProcess*)self, &QProcess::objectNameChanged, (QProcess*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
	__conns_QProcess[__id] = __c;
}
void QProcess__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QProcess);
		auto __it = __conns_QProcess.find(id);
		if (__it == __conns_QProcess.end()) return;
		__c = __it->second;
		__conns_QProcess.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QProcess__Go_new(long long id, void* parent) { return new QProcess__Go(id, (QObject*)parent); }
void QProcess_delete(void* self) { delete (QProcess*)self; }

}

#include <QtCore/qtipccommon.h>
#include <QtCore/qstring.h>
#include "gen_qnativeipckey.h"

extern "C" {

void* QNativeIpcKey_new() {
	return new QNativeIpcKey();
}
void* QNativeIpcKey_new2(int type_) {
	return new QNativeIpcKey((QNativeIpcKey::Type)(type_));
}
void* QNativeIpcKey_new3(const char* k, long long k_len) {
	QString k_qstr = QString::fromUtf8(k, k_len);
	return new QNativeIpcKey(k_qstr);
}
void* QNativeIpcKey_new4(void* other) {
	return new QNativeIpcKey(*(QNativeIpcKey*)other);
}
void* QNativeIpcKey_new5(const char* k, long long k_len, int type_) {
	QString k_qstr = QString::fromUtf8(k, k_len);
	return new QNativeIpcKey(k_qstr, (QNativeIpcKey::Type)(type_));
}
int QNativeIpcKey_legacyDefaultTypeForOs() {
	auto&& __r = QNativeIpcKey::legacyDefaultTypeForOs();
	return (int)(__r);
}
void* QNativeIpcKey_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QNativeIpcKey*)self)->operator=(*(QNativeIpcKey*)other);
	return (void*)&(__r);
}
void QNativeIpcKey_swap(void* self, void* other) {
	((QNativeIpcKey*)self)->swap(*(QNativeIpcKey*)other);
}
int QNativeIpcKey_isEmpty(void* self) {
	auto&& __r = ((QNativeIpcKey*)self)->isEmpty();
	return __r;
}
int QNativeIpcKey_isValid(void* self) {
	auto&& __r = ((QNativeIpcKey*)self)->isValid();
	return __r;
}
int QNativeIpcKey_type(void* self) {
	auto&& __r = ((QNativeIpcKey*)self)->type();
	return (int)(__r);
}
void QNativeIpcKey_setType(void* self, int type_) {
	((QNativeIpcKey*)self)->setType((QNativeIpcKey::Type)(type_));
}
char* QNativeIpcKey_nativeKey(void* self, long long* __out_len) {
	auto&& __r = ((QNativeIpcKey*)self)->nativeKey();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QNativeIpcKey_setNativeKey(void* self, const char* newKey, long long newKey_len) {
	QString newKey_qstr = QString::fromUtf8(newKey, newKey_len);
	((QNativeIpcKey*)self)->setNativeKey(newKey_qstr);
}
char* QNativeIpcKey_toString(void* self, long long* __out_len) {
	auto&& __r = ((QNativeIpcKey*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QNativeIpcKey_fromString(const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QNativeIpcKey::fromString(string__qstr);
	return new QNativeIpcKey(__r);
}
int QNativeIpcKey_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QNativeIpcKey*)self) == (*(QNativeIpcKey*)rhs);
	return __r;
}
int QNativeIpcKey_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QNativeIpcKey*)self) != (*(QNativeIpcKey*)rhs);
	return __r;
}
void QNativeIpcKey_delete(void* self) { delete (QNativeIpcKey*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qstringview.h>
#include "gen_qxmlstreamentitydeclaration.h"

extern "C" {

void* QXmlStreamEntityDeclaration_new() {
	return new QXmlStreamEntityDeclaration();
}
void* QXmlStreamEntityDeclaration_new2(void* arg1) {
	return new QXmlStreamEntityDeclaration(*(QXmlStreamEntityDeclaration*)arg1);
}
void* QXmlStreamEntityDeclaration_name(void* self) {
	auto&& __r = ((QXmlStreamEntityDeclaration*)self)->name();
	return new QStringView(__r);
}
void* QXmlStreamEntityDeclaration_notationName(void* self) {
	auto&& __r = ((QXmlStreamEntityDeclaration*)self)->notationName();
	return new QStringView(__r);
}
void* QXmlStreamEntityDeclaration_systemId(void* self) {
	auto&& __r = ((QXmlStreamEntityDeclaration*)self)->systemId();
	return new QStringView(__r);
}
void* QXmlStreamEntityDeclaration_publicId(void* self) {
	auto&& __r = ((QXmlStreamEntityDeclaration*)self)->publicId();
	return new QStringView(__r);
}
void* QXmlStreamEntityDeclaration_value(void* self) {
	auto&& __r = ((QXmlStreamEntityDeclaration*)self)->value();
	return new QStringView(__r);
}
int QXmlStreamEntityDeclaration_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamEntityDeclaration*)self) == (*(QXmlStreamEntityDeclaration*)rhs);
	return __r;
}
int QXmlStreamEntityDeclaration_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamEntityDeclaration*)self) != (*(QXmlStreamEntityDeclaration*)rhs);
	return __r;
}
void QXmlStreamEntityDeclaration_delete(void* self) { delete (QXmlStreamEntityDeclaration*)self; }

}

#include <QtGui/qpointingdevice.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qevent.h>
#include <QtGui/qeventpoint.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qpointingdevice.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPointingDevice__Go : public QPointingDevice {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QPointingDevice__Go(long long id, A&&... a) : QPointingDevice(std::forward<A>(a)...), __goId(id) {}
	using QPointingDevice::sender;
	using QPointingDevice::senderSignalIndex;
	using QPointingDevice::receivers;
	using QPointingDevice::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPointingDevice::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPointingDevice::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPointingDevice::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPointingDevice::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPointingDevice::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPointingDevice::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QPointingDevice::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPointingDevice::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QPointingDevice::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPointingDevice::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPointingDevice::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPointingDevice::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPointingDevice::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPointingDevice;
static std::mutex __connsMu_QPointingDevice;

extern "C" {

void* QPointingDevice_base(void* self) { return (void*)static_cast<QInputDevice*>((QPointingDevice*)self); }
void QPointingDevice_chain(void* self, void** out) {
	QPointingDevice* __p = (QPointingDevice*)self;
	QInputDevice* __b0 = static_cast<QInputDevice*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QPointingDevice__set_override(void* self, int idx) { ((QPointingDevice__Go*)self)->__goOn[idx] = 1; }
void* QPointingDevice_fromQObject(void* o) { return (void*)qobject_cast<QPointingDevice*>((QObject*)o); }
void* QPointingDevice_new(long long __goid) {
	return new QPointingDevice__Go(__goid);
}
void* QPointingDevice_new2(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QPointingDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(devType), (QPointingDevice::PointerType)(pType), QFlags<QInputDevice::Capability>(caps), maxPoints, buttonCount);
}
void* QPointingDevice_new3(long long __goid, void* parent) {
	return new QPointingDevice__Go(__goid, (QObject*)parent);
}
void* QPointingDevice_new4(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount, const char* seatName, long long seatName_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	return new QPointingDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(devType), (QPointingDevice::PointerType)(pType), QFlags<QInputDevice::Capability>(caps), maxPoints, buttonCount, seatName_qstr);
}
void* QPointingDevice_new5(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount, const char* seatName, long long seatName_len, void* uniqueId) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	return new QPointingDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(devType), (QPointingDevice::PointerType)(pType), QFlags<QInputDevice::Capability>(caps), maxPoints, buttonCount, seatName_qstr, *(QPointingDeviceUniqueId*)uniqueId);
}
void* QPointingDevice_new6(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount, const char* seatName, long long seatName_len, void* uniqueId, void* parent) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	return new QPointingDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(devType), (QPointingDevice::PointerType)(pType), QFlags<QInputDevice::Capability>(caps), maxPoints, buttonCount, seatName_qstr, *(QPointingDeviceUniqueId*)uniqueId, (QObject*)parent);
}
char* QPointingDevice_tr(const char* s, long long* __out_len) {
	auto&& __r = QPointingDevice::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPointingDevice_setType(void* self, int devType) {
	((QPointingDevice*)self)->setType((QInputDevice::DeviceType)(devType));
}
void QPointingDevice_setCapabilities(void* self, int caps) {
	((QPointingDevice*)self)->setCapabilities(QFlags<QInputDevice::Capability>(caps));
}
void QPointingDevice_setMaximumTouchPoints(void* self, int c) {
	((QPointingDevice*)self)->setMaximumTouchPoints(c);
}
int QPointingDevice_pointerType(void* self) {
	auto&& __r = ((QPointingDevice*)self)->pointerType();
	return (int)(__r);
}
int QPointingDevice_maximumPoints(void* self) {
	auto&& __r = ((QPointingDevice*)self)->maximumPoints();
	return __r;
}
int QPointingDevice_buttonCount(void* self) {
	auto&& __r = ((QPointingDevice*)self)->buttonCount();
	return __r;
}
void* QPointingDevice_uniqueId(void* self) {
	auto&& __r = ((QPointingDevice*)self)->uniqueId();
	return new QPointingDeviceUniqueId(__r);
}
void* QPointingDevice_primaryPointingDevice() {
	auto&& __r = QPointingDevice::primaryPointingDevice();
	return (void*)(__r);
}
int QPointingDevice_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QPointingDevice*)self)->operator==(*(QPointingDevice*)other);
	return __r;
}
void QPointingDevice_grabChanged(void* self, void* grabber, int transition, void* event, void* point) {
	((QPointingDevice*)self)->grabChanged((QObject*)grabber, (QPointingDevice::GrabTransition)(transition), (QPointerEvent*)event, *(QEventPoint*)point);
}
char** QPointingDevice_seatNames(long long* __out_len) {
	auto&& __r = QPointingDevice::seatNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QPointingDevice_devices(long long* __out_len) {
	auto&& __r = QPointingDevice::devices();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QPointingDevice_primaryKeyboard() {
	auto&& __r = QPointingDevice::primaryKeyboard();
	return (void*)(__r);
}
void* QPointingDevice_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPointingDevice::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPointingDevice_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPointingDevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPointingDevice_sender(void* self) {
	auto&& __r = ((QPointingDevice__Go*)self)->sender();
	return (void*)(__r);
}
int QPointingDevice_senderSignalIndex(void* self) {
	auto&& __r = ((QPointingDevice__Go*)self)->senderSignalIndex();
	return __r;
}
int QPointingDevice_receivers(void* self, const char* signal) {
	auto&& __r = ((QPointingDevice__Go*)self)->receivers(signal);
	return __r;
}
int QPointingDevice_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPointingDevice__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QPointingDevice_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPointingDevice::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPointingDevice_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPointingDevice::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPointingDevice_primaryPointingDevice2(const char* seatName, long long seatName_len) {
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	auto&& __r = QPointingDevice::primaryPointingDevice(seatName_qstr);
	return (void*)(__r);
}
void* QPointingDevice_primaryKeyboard2(const char* seatName, long long seatName_len) {
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	auto&& __r = QPointingDevice::primaryKeyboard(seatName_qstr);
	return (void*)(__r);
}
void* QPointingDevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPointingDevice::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPointingDevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPointingDevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPointingDevice_metaObject_default(void* self) {
	auto&& __r = ((QPointingDevice*)self)->QPointingDevice::metaObject();
	return (void*)(__r);
}
int QPointingDevice_event_default(void* self, void* event) {
	auto&& __r = ((QPointingDevice*)self)->QPointingDevice::event((QEvent*)event);
	return __r;
}
int QPointingDevice_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPointingDevice*)self)->QPointingDevice::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPointingDevice_timerEvent_default(void* self, void* event) {
	((QPointingDevice__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QPointingDevice_childEvent_default(void* self, void* event) {
	((QPointingDevice__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPointingDevice_customEvent_default(void* self, void* event) {
	((QPointingDevice__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPointingDevice_connectNotify_default(void* self, void* signal) {
	((QPointingDevice__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPointingDevice_disconnectNotify_default(void* self, void* signal) {
	((QPointingDevice__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPointingDevice_connect_grabChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPointingDevice*)self, &QPointingDevice::grabChanged, (QPointingDevice*)self, [__id](QObject* grabber, QPointingDevice::GrabTransition transition, const QPointerEvent* event, const QEventPoint& point){
		if (cppgo_sig_void_voidp_int_voidp_voidp) ((void(*)(long long, void*, int, void*, void*))cppgo_sig_void_voidp_int_voidp_voidp)(__id, (void*)grabber, (int)(transition), (void*)event, (void*)&point);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
	__conns_QPointingDevice[__id] = __c;
}
void QPointingDevice_connect_availableVirtualGeometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPointingDevice*)self, &QPointingDevice::availableVirtualGeometryChanged, (QPointingDevice*)self, [__id](QRect area){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&area);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
	__conns_QPointingDevice[__id] = __c;
}
void QPointingDevice_connect_capabilitiesChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPointingDevice*)self, &QPointingDevice::capabilitiesChanged, (QPointingDevice*)self, [__id](QFlags<QInputDevice::Capability> capabilities){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(capabilities));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
	__conns_QPointingDevice[__id] = __c;
}
void QPointingDevice_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPointingDevice*)self, &QPointingDevice::destroyed, (QPointingDevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
	__conns_QPointingDevice[__id] = __c;
}
void QPointingDevice_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPointingDevice*)self, &QPointingDevice::destroyed, (QPointingDevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
	__conns_QPointingDevice[__id] = __c;
}
void QPointingDevice_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPointingDevice*)self, &QPointingDevice::objectNameChanged, (QPointingDevice*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
	__conns_QPointingDevice[__id] = __c;
}
void QPointingDevice__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPointingDevice);
		auto __it = __conns_QPointingDevice.find(id);
		if (__it == __conns_QPointingDevice.end()) return;
		__c = __it->second;
		__conns_QPointingDevice.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPointingDevice__Go_new(long long id, void* parent) { return new QPointingDevice__Go(id, (QObject*)parent); }
void QPointingDevice_delete(void* self) { delete (QPointingDevice*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qstring.h>
#include "gen_qinputmethodevent.h"
#include "gen_purego_cb.h"

class QInputMethodEvent__Go : public QInputMethodEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QInputMethodEvent__Go(long long id, A&&... a) : QInputMethodEvent(std::forward<A>(a)...), __goId(id) {}
	QInputMethodEvent* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QInputMethodEvent*)__r; }
		return QInputMethodEvent::clone();
	}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QInputMethodEvent::setAccepted(accepted);
	}
};

extern "C" {

void* QInputMethodEvent_base(void* self) { return (void*)static_cast<QEvent*>((QInputMethodEvent*)self); }
void QInputMethodEvent__set_override(void* self, int idx) { ((QInputMethodEvent__Go*)self)->__goOn[idx] = 1; }
void* QInputMethodEvent_new(long long __goid) {
	return new QInputMethodEvent__Go(__goid);
}
void* QInputMethodEvent_new2(long long __goid, const char* preeditText, long long preeditText_len, void** attributes, long long attributes_len) {
	QString preeditText_qstr = QString::fromUtf8(preeditText, preeditText_len);
	QList<QInputMethodEvent::Attribute> attributes_list; attributes_list.reserve(attributes_len);
	for (long long __i = 0; __i < attributes_len; __i++) attributes_list.append(*(QInputMethodEvent::Attribute*)attributes[__i]);
	return new QInputMethodEvent__Go(__goid, preeditText_qstr, attributes_list);
}
void QInputMethodEvent_setCommitString(void* self, const char* commitString, long long commitString_len) {
	QString commitString_qstr = QString::fromUtf8(commitString, commitString_len);
	((QInputMethodEvent*)self)->setCommitString(commitString_qstr);
}
void** QInputMethodEvent_attributes(void* self, long long* __out_len) {
	auto&& __r = ((QInputMethodEvent*)self)->attributes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QInputMethodEvent::Attribute(__e);
	return __arr;
}
char* QInputMethodEvent_preeditString(void* self, long long* __out_len) {
	auto&& __r = ((QInputMethodEvent*)self)->preeditString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputMethodEvent_commitString(void* self, long long* __out_len) {
	auto&& __r = ((QInputMethodEvent*)self)->commitString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QInputMethodEvent_replacementStart(void* self) {
	auto&& __r = ((QInputMethodEvent*)self)->replacementStart();
	return __r;
}
int QInputMethodEvent_replacementLength(void* self) {
	auto&& __r = ((QInputMethodEvent*)self)->replacementLength();
	return __r;
}
int QInputMethodEvent_registerEventType() {
	auto&& __r = QInputMethodEvent::registerEventType();
	return __r;
}
void QInputMethodEvent_setCommitString2(void* self, const char* commitString, long long commitString_len, int replaceFrom) {
	QString commitString_qstr = QString::fromUtf8(commitString, commitString_len);
	((QInputMethodEvent*)self)->setCommitString(commitString_qstr, replaceFrom);
}
void QInputMethodEvent_setCommitString3(void* self, const char* commitString, long long commitString_len, int replaceFrom, int replaceLength) {
	QString commitString_qstr = QString::fromUtf8(commitString, commitString_len);
	((QInputMethodEvent*)self)->setCommitString(commitString_qstr, replaceFrom, replaceLength);
}
int QInputMethodEvent_registerEventType2(int hint) {
	auto&& __r = QInputMethodEvent::registerEventType(hint);
	return __r;
}
void* QInputMethodEvent_clone_default(void* self) {
	auto&& __r = ((QInputMethodEvent*)self)->QInputMethodEvent::clone();
	return (void*)(__r);
}
void QInputMethodEvent_setAccepted_default(void* self, int accepted) {
	((QInputMethodEvent*)self)->QInputMethodEvent::setAccepted(accepted);
}
void* QInputMethodEvent__Go_new(long long id) { return new QInputMethodEvent__Go(id); }
void QInputMethodEvent_delete(void* self) { delete (QInputMethodEvent*)self; }

}

#include <QtGui/qrawfont.h>
#include <QtCore/qstring.h>
#include <QtGui/qfont.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qpoint.h>
#include <QtCore/qchar.h>
#include <QtGui/qimage.h>
#include <QtGui/qtransform.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qrect.h>
#include <QtGui/qfontdatabase.h>
#include "gen_qrawfont.h"

extern "C" {

void* QRawFont_new() {
	return new QRawFont();
}
void* QRawFont_new2(const char* fileName, long long fileName_len, double pixelSize) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QRawFont(fileName_qstr, pixelSize);
}
void* QRawFont_new3(void* fontData, double pixelSize) {
	return new QRawFont(*(QByteArray*)fontData, pixelSize);
}
void* QRawFont_new4(void* other) {
	return new QRawFont(*(QRawFont*)other);
}
void* QRawFont_new5(const char* fileName, long long fileName_len, double pixelSize, int hintingPreference) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QRawFont(fileName_qstr, pixelSize, (QFont::HintingPreference)(hintingPreference));
}
void* QRawFont_new6(void* fontData, double pixelSize, int hintingPreference) {
	return new QRawFont(*(QByteArray*)fontData, pixelSize, (QFont::HintingPreference)(hintingPreference));
}
void* QRawFont_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QRawFont*)self)->operator=(*(QRawFont*)other);
	return (void*)&(__r);
}
void QRawFont_swap(void* self, void* other) {
	((QRawFont*)self)->swap(*(QRawFont*)other);
}
int QRawFont_isValid(void* self) {
	auto&& __r = ((QRawFont*)self)->isValid();
	return __r;
}
int QRawFont_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QRawFont*)self)->operator==(*(QRawFont*)other);
	return __r;
}
int QRawFont_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QRawFont*)self)->operator!=(*(QRawFont*)other);
	return __r;
}
unsigned int QRawFont_glyphCount(void* self) {
	auto&& __r = ((QRawFont*)self)->glyphCount();
	return __r;
}
char* QRawFont_familyName(void* self, long long* __out_len) {
	auto&& __r = ((QRawFont*)self)->familyName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRawFont_styleName(void* self, long long* __out_len) {
	auto&& __r = ((QRawFont*)self)->styleName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QRawFont_style(void* self) {
	auto&& __r = ((QRawFont*)self)->style();
	return (int)(__r);
}
int QRawFont_weight(void* self) {
	auto&& __r = ((QRawFont*)self)->weight();
	return __r;
}
unsigned int* QRawFont_glyphIndexesForString(void* self, const char* text, long long text_len, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QRawFont*)self)->glyphIndexesForString(text_qstr);
	*__out_len = __r.size();
	unsigned int* __arr = (unsigned int*)malloc(sizeof(unsigned int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void** QRawFont_advancesForGlyphIndexes(void* self, unsigned int* glyphIndexes, long long glyphIndexes_len, long long* __out_len) {
	QList<unsigned int> glyphIndexes_list; glyphIndexes_list.reserve(glyphIndexes_len);
	for (long long __i = 0; __i < glyphIndexes_len; __i++) glyphIndexes_list.append(glyphIndexes[__i]);
	auto&& __r = ((QRawFont*)self)->advancesForGlyphIndexes(glyphIndexes_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QRawFont_advancesForGlyphIndexes2(void* self, unsigned int* glyphIndexes, long long glyphIndexes_len, int layoutFlags, long long* __out_len) {
	QList<unsigned int> glyphIndexes_list; glyphIndexes_list.reserve(glyphIndexes_len);
	for (long long __i = 0; __i < glyphIndexes_len; __i++) glyphIndexes_list.append(glyphIndexes[__i]);
	auto&& __r = ((QRawFont*)self)->advancesForGlyphIndexes(glyphIndexes_list, QFlags<QRawFont::LayoutFlag>(layoutFlags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
int QRawFont_glyphIndexesForChars(void* self, void* chars, int numChars, unsigned int* glyphIndexes, int* numGlyphs) {
	quint32 glyphIndexes_local{};
	int numGlyphs_local{};
	auto&& __r = ((QRawFont*)self)->glyphIndexesForChars((QChar*)chars, numChars, &glyphIndexes_local, &numGlyphs_local);
	*glyphIndexes = glyphIndexes_local;
	*numGlyphs = numGlyphs_local;
	return __r;
}
void* QRawFont_alphaMapForGlyph(void* self, unsigned int glyphIndex) {
	auto&& __r = ((QRawFont*)self)->alphaMapForGlyph(glyphIndex);
	return new QImage(__r);
}
void* QRawFont_pathForGlyph(void* self, unsigned int glyphIndex) {
	auto&& __r = ((QRawFont*)self)->pathForGlyph(glyphIndex);
	return new QPainterPath(__r);
}
void* QRawFont_boundingRect(void* self, unsigned int glyphIndex) {
	auto&& __r = ((QRawFont*)self)->boundingRect(glyphIndex);
	return new QRectF(__r);
}
void QRawFont_setPixelSize(void* self, double pixelSize) {
	((QRawFont*)self)->setPixelSize(pixelSize);
}
double QRawFont_pixelSize(void* self) {
	auto&& __r = ((QRawFont*)self)->pixelSize();
	return __r;
}
int QRawFont_hintingPreference(void* self) {
	auto&& __r = ((QRawFont*)self)->hintingPreference();
	return (int)(__r);
}
double QRawFont_ascent(void* self) {
	auto&& __r = ((QRawFont*)self)->ascent();
	return __r;
}
double QRawFont_capHeight(void* self) {
	auto&& __r = ((QRawFont*)self)->capHeight();
	return __r;
}
double QRawFont_descent(void* self) {
	auto&& __r = ((QRawFont*)self)->descent();
	return __r;
}
double QRawFont_leading(void* self) {
	auto&& __r = ((QRawFont*)self)->leading();
	return __r;
}
double QRawFont_xHeight(void* self) {
	auto&& __r = ((QRawFont*)self)->xHeight();
	return __r;
}
double QRawFont_averageCharWidth(void* self) {
	auto&& __r = ((QRawFont*)self)->averageCharWidth();
	return __r;
}
double QRawFont_maxCharWidth(void* self) {
	auto&& __r = ((QRawFont*)self)->maxCharWidth();
	return __r;
}
double QRawFont_lineThickness(void* self) {
	auto&& __r = ((QRawFont*)self)->lineThickness();
	return __r;
}
double QRawFont_underlinePosition(void* self) {
	auto&& __r = ((QRawFont*)self)->underlinePosition();
	return __r;
}
double QRawFont_unitsPerEm(void* self) {
	auto&& __r = ((QRawFont*)self)->unitsPerEm();
	return __r;
}
char* QRawFont_glyphName(void* self, unsigned int glyphIndex, long long* __out_len) {
	auto&& __r = ((QRawFont*)self)->glyphName(glyphIndex);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QRawFont_loadFromFile(void* self, const char* fileName, long long fileName_len, double pixelSize, int hintingPreference) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QRawFont*)self)->loadFromFile(fileName_qstr, pixelSize, (QFont::HintingPreference)(hintingPreference));
}
void QRawFont_loadFromData(void* self, void* fontData, double pixelSize, int hintingPreference) {
	((QRawFont*)self)->loadFromData(*(QByteArray*)fontData, pixelSize, (QFont::HintingPreference)(hintingPreference));
}
int QRawFont_supportsCharacter(void* self, unsigned int ucs4) {
	auto&& __r = ((QRawFont*)self)->supportsCharacter(ucs4);
	return __r;
}
int QRawFont_supportsCharacter2(void* self, void* character) {
	auto&& __r = ((QRawFont*)self)->supportsCharacter(*(QChar*)character);
	return __r;
}
int* QRawFont_supportedWritingSystems(void* self, long long* __out_len) {
	auto&& __r = ((QRawFont*)self)->supportedWritingSystems();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
void* QRawFont_fontTable(void* self, const char* tagName) {
	auto&& __r = ((QRawFont*)self)->fontTable(tagName);
	return new QByteArray(__r);
}
void* QRawFont_fontTable2(void* self, void* tag) {
	auto&& __r = ((QRawFont*)self)->fontTable(*(QFont::Tag*)tag);
	return new QByteArray(__r);
}
void* QRawFont_fromFont(void* font) {
	auto&& __r = QRawFont::fromFont(*(QFont*)font);
	return new QRawFont(__r);
}
void* QRawFont_alphaMapForGlyph2(void* self, unsigned int glyphIndex, int antialiasingType) {
	auto&& __r = ((QRawFont*)self)->alphaMapForGlyph(glyphIndex, (QRawFont::AntialiasingType)(antialiasingType));
	return new QImage(__r);
}
void* QRawFont_alphaMapForGlyph3(void* self, unsigned int glyphIndex, int antialiasingType, void* transform) {
	auto&& __r = ((QRawFont*)self)->alphaMapForGlyph(glyphIndex, (QRawFont::AntialiasingType)(antialiasingType), *(QTransform*)transform);
	return new QImage(__r);
}
void* QRawFont_fromFont2(void* font, int writingSystem) {
	auto&& __r = QRawFont::fromFont(*(QFont*)font, (QFontDatabase::WritingSystem)(writingSystem));
	return new QRawFont(__r);
}
void QRawFont_delete(void* self) { delete (QRawFont*)self; }

}

#include <QtGui/qtextlayout.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtextformat.h>
#include "gen_qtextinlineobject.h"

extern "C" {

void* QTextInlineObject_new() {
	return new QTextInlineObject();
}
int QTextInlineObject_isValid(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->isValid();
	return __r;
}
void* QTextInlineObject_rect(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->rect();
	return new QRectF(__r);
}
double QTextInlineObject_width(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->width();
	return __r;
}
double QTextInlineObject_ascent(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->ascent();
	return __r;
}
double QTextInlineObject_descent(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->descent();
	return __r;
}
double QTextInlineObject_height(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->height();
	return __r;
}
int QTextInlineObject_textDirection(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->textDirection();
	return (int)(__r);
}
void QTextInlineObject_setWidth(void* self, double w) {
	((QTextInlineObject*)self)->setWidth(w);
}
void QTextInlineObject_setAscent(void* self, double a) {
	((QTextInlineObject*)self)->setAscent(a);
}
void QTextInlineObject_setDescent(void* self, double d) {
	((QTextInlineObject*)self)->setDescent(d);
}
int QTextInlineObject_textPosition(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->textPosition();
	return __r;
}
int QTextInlineObject_formatIndex(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->formatIndex();
	return __r;
}
void* QTextInlineObject_format(void* self) {
	auto&& __r = ((QTextInlineObject*)self)->format();
	return new QTextFormat(__r);
}
void QTextInlineObject_delete(void* self) { delete (QTextInlineObject*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibleevent.h"

extern "C" {

void* QAccessibleEvent_new(void* obj, int typ) {
	return new QAccessibleEvent((QObject*)obj, (QAccessible::Event)(typ));
}
void* QAccessibleEvent_new2(void* iface, int typ) {
	return new QAccessibleEvent((QAccessibleInterface*)iface, (QAccessible::Event)(typ));
}
int QAccessibleEvent_type(void* self) {
	auto&& __r = ((QAccessibleEvent*)self)->type();
	return (int)(__r);
}
void* QAccessibleEvent_object(void* self) {
	auto&& __r = ((QAccessibleEvent*)self)->object();
	return (void*)(__r);
}
unsigned int QAccessibleEvent_uniqueId(void* self) {
	auto&& __r = ((QAccessibleEvent*)self)->uniqueId();
	return __r;
}
void QAccessibleEvent_setChild(void* self, int chld) {
	((QAccessibleEvent*)self)->setChild(chld);
}
int QAccessibleEvent_child(void* self) {
	auto&& __r = ((QAccessibleEvent*)self)->child();
	return __r;
}
void* QAccessibleEvent_accessibleInterface(void* self) {
	auto&& __r = ((QAccessibleEvent*)self)->accessibleInterface();
	return (void*)(__r);
}
void QAccessibleEvent_delete(void* self) { delete (QAccessibleEvent*)self; }

}

#include <QtGui/qcolortransform.h>
#include <QtGui/qrgba64.h>
#include <QtCore/qfloat16.h>
#include <QtGui/qcolor.h>
#include "gen_qcolortransform.h"

extern "C" {

void* QColorTransform_new() {
	return new QColorTransform();
}
void* QColorTransform_new2(void* colorTransform) {
	return new QColorTransform(*(QColorTransform*)colorTransform);
}
void* QColorTransform_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QColorTransform*)self)->operator=(*(QColorTransform*)other);
	return (void*)&(__r);
}
void QColorTransform_swap(void* self, void* other) {
	((QColorTransform*)self)->swap(*(QColorTransform*)other);
}
int QColorTransform_isIdentity(void* self) {
	auto&& __r = ((QColorTransform*)self)->isIdentity();
	return __r;
}
unsigned int QColorTransform_map(void* self, unsigned int argb) {
	auto&& __r = ((QColorTransform*)self)->map(argb);
	return __r;
}
void* QColorTransform_map2(void* self, void* rgba64) {
	auto&& __r = ((QColorTransform*)self)->map(*(QRgba64*)rgba64);
	return new QRgba64(__r);
}
void* QColorTransform_map3(void* self, void* color) {
	auto&& __r = ((QColorTransform*)self)->map(*(QColor*)color);
	return new QColor(__r);
}
int QColorTransform_OperatorEqual(void* self, void* ct2) {
	auto&& __r = (*(QColorTransform*)self) == (*(QColorTransform*)ct2);
	return __r;
}
int QColorTransform_OperatorNotEqual(void* self, void* ct2) {
	auto&& __r = (*(QColorTransform*)self) != (*(QColorTransform*)ct2);
	return __r;
}
void QColorTransform_delete(void* self) { delete (QColorTransform*)self; }

}

#include <QtGui/qopenglfunctions.h>
#include <QtGui/qopenglcontext.h>
#include "gen_qopenglfunctions.h"

extern "C" {

void* QOpenGLFunctions_new() {
	return new QOpenGLFunctions();
}
void* QOpenGLFunctions_new2(void* context) {
	return new QOpenGLFunctions((QOpenGLContext*)context);
}
void* QOpenGLFunctions_new3(void* arg1) {
	return new QOpenGLFunctions(*(QOpenGLFunctions*)arg1);
}
int QOpenGLFunctions_openGLFeatures(void* self) {
	auto&& __r = ((QOpenGLFunctions*)self)->openGLFeatures();
	return (int)(__r);
}
int QOpenGLFunctions_hasOpenGLFeature(void* self, int feature) {
	auto&& __r = ((QOpenGLFunctions*)self)->hasOpenGLFeature((QOpenGLFunctions::OpenGLFeature)(feature));
	return __r;
}
void QOpenGLFunctions_initializeOpenGLFunctions(void* self) {
	((QOpenGLFunctions*)self)->initializeOpenGLFunctions();
}
void QOpenGLFunctions_glBindTexture(void* self, unsigned int target, unsigned int texture) {
	((QOpenGLFunctions*)self)->glBindTexture(target, texture);
}
void QOpenGLFunctions_glBlendFunc(void* self, unsigned int sfactor, unsigned int dfactor) {
	((QOpenGLFunctions*)self)->glBlendFunc(sfactor, dfactor);
}
void QOpenGLFunctions_glClear(void* self, unsigned int mask) {
	((QOpenGLFunctions*)self)->glClear(mask);
}
void QOpenGLFunctions_glClearColor(void* self, float red, float green, float blue, float alpha) {
	((QOpenGLFunctions*)self)->glClearColor(red, green, blue, alpha);
}
void QOpenGLFunctions_glClearStencil(void* self, int s) {
	((QOpenGLFunctions*)self)->glClearStencil(s);
}
void QOpenGLFunctions_glColorMask(void* self, unsigned char red, unsigned char green, unsigned char blue, unsigned char alpha) {
	((QOpenGLFunctions*)self)->glColorMask(red, green, blue, alpha);
}
void QOpenGLFunctions_glCopyTexImage2D(void* self, unsigned int target, int level, unsigned int internalformat, int x, int y, int width, int height, int border) {
	((QOpenGLFunctions*)self)->glCopyTexImage2D(target, level, internalformat, x, y, width, height, border);
}
void QOpenGLFunctions_glCopyTexSubImage2D(void* self, unsigned int target, int level, int xoffset, int yoffset, int x, int y, int width, int height) {
	((QOpenGLFunctions*)self)->glCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}
void QOpenGLFunctions_glCullFace(void* self, unsigned int mode) {
	((QOpenGLFunctions*)self)->glCullFace(mode);
}
void QOpenGLFunctions_glDepthFunc(void* self, unsigned int func_) {
	((QOpenGLFunctions*)self)->glDepthFunc(func_);
}
void QOpenGLFunctions_glDepthMask(void* self, unsigned char flag) {
	((QOpenGLFunctions*)self)->glDepthMask(flag);
}
void QOpenGLFunctions_glDisable(void* self, unsigned int cap_) {
	((QOpenGLFunctions*)self)->glDisable(cap_);
}
void QOpenGLFunctions_glDrawArrays(void* self, unsigned int mode, int first, int count) {
	((QOpenGLFunctions*)self)->glDrawArrays(mode, first, count);
}
void QOpenGLFunctions_glEnable(void* self, unsigned int cap_) {
	((QOpenGLFunctions*)self)->glEnable(cap_);
}
void QOpenGLFunctions_glFinish(void* self) {
	((QOpenGLFunctions*)self)->glFinish();
}
void QOpenGLFunctions_glFlush(void* self) {
	((QOpenGLFunctions*)self)->glFlush();
}
void QOpenGLFunctions_glFrontFace(void* self, unsigned int mode) {
	((QOpenGLFunctions*)self)->glFrontFace(mode);
}
void QOpenGLFunctions_glGenTextures(void* self, int n, unsigned int* textures) {
	GLuint textures_local{};
	((QOpenGLFunctions*)self)->glGenTextures(n, &textures_local);
	*textures = textures_local;
}
void QOpenGLFunctions_glGetBooleanv(void* self, unsigned int pname, unsigned char* params) {
	GLboolean params_local{};
	((QOpenGLFunctions*)self)->glGetBooleanv(pname, &params_local);
	*params = params_local;
}
unsigned int QOpenGLFunctions_glGetError(void* self) {
	auto&& __r = ((QOpenGLFunctions*)self)->glGetError();
	return __r;
}
void QOpenGLFunctions_glGetFloatv(void* self, unsigned int pname, float* params) {
	GLfloat params_local{};
	((QOpenGLFunctions*)self)->glGetFloatv(pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetIntegerv(void* self, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetIntegerv(pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetTexParameterfv(void* self, unsigned int target, unsigned int pname, float* params) {
	GLfloat params_local{};
	((QOpenGLFunctions*)self)->glGetTexParameterfv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetTexParameteriv(void* self, unsigned int target, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetTexParameteriv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glHint(void* self, unsigned int target, unsigned int mode) {
	((QOpenGLFunctions*)self)->glHint(target, mode);
}
unsigned char QOpenGLFunctions_glIsEnabled(void* self, unsigned int cap_) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsEnabled(cap_);
	return __r;
}
unsigned char QOpenGLFunctions_glIsTexture(void* self, unsigned int texture) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsTexture(texture);
	return __r;
}
void QOpenGLFunctions_glLineWidth(void* self, float width) {
	((QOpenGLFunctions*)self)->glLineWidth(width);
}
void QOpenGLFunctions_glPixelStorei(void* self, unsigned int pname, int param) {
	((QOpenGLFunctions*)self)->glPixelStorei(pname, param);
}
void QOpenGLFunctions_glPolygonOffset(void* self, float factor, float units) {
	((QOpenGLFunctions*)self)->glPolygonOffset(factor, units);
}
void QOpenGLFunctions_glScissor(void* self, int x, int y, int width, int height) {
	((QOpenGLFunctions*)self)->glScissor(x, y, width, height);
}
void QOpenGLFunctions_glStencilFunc(void* self, unsigned int func_, int ref, unsigned int mask) {
	((QOpenGLFunctions*)self)->glStencilFunc(func_, ref, mask);
}
void QOpenGLFunctions_glStencilMask(void* self, unsigned int mask) {
	((QOpenGLFunctions*)self)->glStencilMask(mask);
}
void QOpenGLFunctions_glStencilOp(void* self, unsigned int fail, unsigned int zfail, unsigned int zpass) {
	((QOpenGLFunctions*)self)->glStencilOp(fail, zfail, zpass);
}
void QOpenGLFunctions_glTexParameterf(void* self, unsigned int target, unsigned int pname, float param) {
	((QOpenGLFunctions*)self)->glTexParameterf(target, pname, param);
}
void QOpenGLFunctions_glTexParameteri(void* self, unsigned int target, unsigned int pname, int param) {
	((QOpenGLFunctions*)self)->glTexParameteri(target, pname, param);
}
void QOpenGLFunctions_glViewport(void* self, int x, int y, int width, int height) {
	((QOpenGLFunctions*)self)->glViewport(x, y, width, height);
}
void QOpenGLFunctions_glActiveTexture(void* self, unsigned int texture) {
	((QOpenGLFunctions*)self)->glActiveTexture(texture);
}
void QOpenGLFunctions_glAttachShader(void* self, unsigned int program, unsigned int shader) {
	((QOpenGLFunctions*)self)->glAttachShader(program, shader);
}
void QOpenGLFunctions_glBindAttribLocation(void* self, unsigned int program, unsigned int index, const char* name) {
	((QOpenGLFunctions*)self)->glBindAttribLocation(program, index, name);
}
void QOpenGLFunctions_glBindBuffer(void* self, unsigned int target, unsigned int buffer) {
	((QOpenGLFunctions*)self)->glBindBuffer(target, buffer);
}
void QOpenGLFunctions_glBindFramebuffer(void* self, unsigned int target, unsigned int framebuffer) {
	((QOpenGLFunctions*)self)->glBindFramebuffer(target, framebuffer);
}
void QOpenGLFunctions_glBindRenderbuffer(void* self, unsigned int target, unsigned int renderbuffer) {
	((QOpenGLFunctions*)self)->glBindRenderbuffer(target, renderbuffer);
}
void QOpenGLFunctions_glBlendColor(void* self, float red, float green, float blue, float alpha) {
	((QOpenGLFunctions*)self)->glBlendColor(red, green, blue, alpha);
}
void QOpenGLFunctions_glBlendEquation(void* self, unsigned int mode) {
	((QOpenGLFunctions*)self)->glBlendEquation(mode);
}
void QOpenGLFunctions_glBlendEquationSeparate(void* self, unsigned int modeRGB, unsigned int modeAlpha) {
	((QOpenGLFunctions*)self)->glBlendEquationSeparate(modeRGB, modeAlpha);
}
void QOpenGLFunctions_glBlendFuncSeparate(void* self, unsigned int srcRGB, unsigned int dstRGB, unsigned int srcAlpha, unsigned int dstAlpha) {
	((QOpenGLFunctions*)self)->glBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}
unsigned int QOpenGLFunctions_glCheckFramebufferStatus(void* self, unsigned int target) {
	auto&& __r = ((QOpenGLFunctions*)self)->glCheckFramebufferStatus(target);
	return __r;
}
void QOpenGLFunctions_glClearDepthf(void* self, float depth) {
	((QOpenGLFunctions*)self)->glClearDepthf(depth);
}
void QOpenGLFunctions_glCompileShader(void* self, unsigned int shader) {
	((QOpenGLFunctions*)self)->glCompileShader(shader);
}
unsigned int QOpenGLFunctions_glCreateProgram(void* self) {
	auto&& __r = ((QOpenGLFunctions*)self)->glCreateProgram();
	return __r;
}
unsigned int QOpenGLFunctions_glCreateShader(void* self, unsigned int type_) {
	auto&& __r = ((QOpenGLFunctions*)self)->glCreateShader(type_);
	return __r;
}
void QOpenGLFunctions_glDeleteProgram(void* self, unsigned int program) {
	((QOpenGLFunctions*)self)->glDeleteProgram(program);
}
void QOpenGLFunctions_glDeleteShader(void* self, unsigned int shader) {
	((QOpenGLFunctions*)self)->glDeleteShader(shader);
}
void QOpenGLFunctions_glDepthRangef(void* self, float zNear, float zFar) {
	((QOpenGLFunctions*)self)->glDepthRangef(zNear, zFar);
}
void QOpenGLFunctions_glDetachShader(void* self, unsigned int program, unsigned int shader) {
	((QOpenGLFunctions*)self)->glDetachShader(program, shader);
}
void QOpenGLFunctions_glDisableVertexAttribArray(void* self, unsigned int index) {
	((QOpenGLFunctions*)self)->glDisableVertexAttribArray(index);
}
void QOpenGLFunctions_glEnableVertexAttribArray(void* self, unsigned int index) {
	((QOpenGLFunctions*)self)->glEnableVertexAttribArray(index);
}
void QOpenGLFunctions_glFramebufferRenderbuffer(void* self, unsigned int target, unsigned int attachment, unsigned int renderbuffertarget, unsigned int renderbuffer) {
	((QOpenGLFunctions*)self)->glFramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer);
}
void QOpenGLFunctions_glFramebufferTexture2D(void* self, unsigned int target, unsigned int attachment, unsigned int textarget, unsigned int texture, int level) {
	((QOpenGLFunctions*)self)->glFramebufferTexture2D(target, attachment, textarget, texture, level);
}
void QOpenGLFunctions_glGenBuffers(void* self, int n, unsigned int* buffers) {
	GLuint buffers_local{};
	((QOpenGLFunctions*)self)->glGenBuffers(n, &buffers_local);
	*buffers = buffers_local;
}
void QOpenGLFunctions_glGenerateMipmap(void* self, unsigned int target) {
	((QOpenGLFunctions*)self)->glGenerateMipmap(target);
}
void QOpenGLFunctions_glGenFramebuffers(void* self, int n, unsigned int* framebuffers) {
	GLuint framebuffers_local{};
	((QOpenGLFunctions*)self)->glGenFramebuffers(n, &framebuffers_local);
	*framebuffers = framebuffers_local;
}
void QOpenGLFunctions_glGenRenderbuffers(void* self, int n, unsigned int* renderbuffers) {
	GLuint renderbuffers_local{};
	((QOpenGLFunctions*)self)->glGenRenderbuffers(n, &renderbuffers_local);
	*renderbuffers = renderbuffers_local;
}
void QOpenGLFunctions_glGetActiveAttrib(void* self, unsigned int program, unsigned int index, int bufsize, int* length, int* size, unsigned int* type_, char* name) {
	GLsizei length_local{};
	GLint size_local{};
	GLenum type__local{};
	((QOpenGLFunctions*)self)->glGetActiveAttrib(program, index, bufsize, &length_local, &size_local, &type__local, (char*)name);
	*length = length_local;
	*size = size_local;
	*type_ = type__local;
}
void QOpenGLFunctions_glGetActiveUniform(void* self, unsigned int program, unsigned int index, int bufsize, int* length, int* size, unsigned int* type_, char* name) {
	GLsizei length_local{};
	GLint size_local{};
	GLenum type__local{};
	((QOpenGLFunctions*)self)->glGetActiveUniform(program, index, bufsize, &length_local, &size_local, &type__local, (char*)name);
	*length = length_local;
	*size = size_local;
	*type_ = type__local;
}
void QOpenGLFunctions_glGetAttachedShaders(void* self, unsigned int program, int maxcount, int* count, unsigned int* shaders) {
	GLsizei count_local{};
	GLuint shaders_local{};
	((QOpenGLFunctions*)self)->glGetAttachedShaders(program, maxcount, &count_local, &shaders_local);
	*count = count_local;
	*shaders = shaders_local;
}
int QOpenGLFunctions_glGetAttribLocation(void* self, unsigned int program, const char* name) {
	auto&& __r = ((QOpenGLFunctions*)self)->glGetAttribLocation(program, name);
	return __r;
}
void QOpenGLFunctions_glGetBufferParameteriv(void* self, unsigned int target, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetBufferParameteriv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetFramebufferAttachmentParameteriv(void* self, unsigned int target, unsigned int attachment, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetFramebufferAttachmentParameteriv(target, attachment, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetProgramiv(void* self, unsigned int program, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetProgramiv(program, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetProgramInfoLog(void* self, unsigned int program, int bufsize, int* length, char* infolog) {
	GLsizei length_local{};
	((QOpenGLFunctions*)self)->glGetProgramInfoLog(program, bufsize, &length_local, (char*)infolog);
	*length = length_local;
}
void QOpenGLFunctions_glGetRenderbufferParameteriv(void* self, unsigned int target, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetRenderbufferParameteriv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetShaderiv(void* self, unsigned int shader, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetShaderiv(shader, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetShaderInfoLog(void* self, unsigned int shader, int bufsize, int* length, char* infolog) {
	GLsizei length_local{};
	((QOpenGLFunctions*)self)->glGetShaderInfoLog(shader, bufsize, &length_local, (char*)infolog);
	*length = length_local;
}
void QOpenGLFunctions_glGetShaderPrecisionFormat(void* self, unsigned int shadertype, unsigned int precisiontype, int* range_, int* precision) {
	GLint range__local{};
	GLint precision_local{};
	((QOpenGLFunctions*)self)->glGetShaderPrecisionFormat(shadertype, precisiontype, &range__local, &precision_local);
	*range_ = range__local;
	*precision = precision_local;
}
void QOpenGLFunctions_glGetShaderSource(void* self, unsigned int shader, int bufsize, int* length, char* source) {
	GLsizei length_local{};
	((QOpenGLFunctions*)self)->glGetShaderSource(shader, bufsize, &length_local, (char*)source);
	*length = length_local;
}
void QOpenGLFunctions_glGetUniformfv(void* self, unsigned int program, int location, float* params) {
	GLfloat params_local{};
	((QOpenGLFunctions*)self)->glGetUniformfv(program, location, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetUniformiv(void* self, unsigned int program, int location, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetUniformiv(program, location, &params_local);
	*params = params_local;
}
int QOpenGLFunctions_glGetUniformLocation(void* self, unsigned int program, const char* name) {
	auto&& __r = ((QOpenGLFunctions*)self)->glGetUniformLocation(program, name);
	return __r;
}
void QOpenGLFunctions_glGetVertexAttribfv(void* self, unsigned int index, unsigned int pname, float* params) {
	GLfloat params_local{};
	((QOpenGLFunctions*)self)->glGetVertexAttribfv(index, pname, &params_local);
	*params = params_local;
}
void QOpenGLFunctions_glGetVertexAttribiv(void* self, unsigned int index, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLFunctions*)self)->glGetVertexAttribiv(index, pname, &params_local);
	*params = params_local;
}
unsigned char QOpenGLFunctions_glIsBuffer(void* self, unsigned int buffer) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsBuffer(buffer);
	return __r;
}
unsigned char QOpenGLFunctions_glIsFramebuffer(void* self, unsigned int framebuffer) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsFramebuffer(framebuffer);
	return __r;
}
unsigned char QOpenGLFunctions_glIsProgram(void* self, unsigned int program) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsProgram(program);
	return __r;
}
unsigned char QOpenGLFunctions_glIsRenderbuffer(void* self, unsigned int renderbuffer) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsRenderbuffer(renderbuffer);
	return __r;
}
unsigned char QOpenGLFunctions_glIsShader(void* self, unsigned int shader) {
	auto&& __r = ((QOpenGLFunctions*)self)->glIsShader(shader);
	return __r;
}
void QOpenGLFunctions_glLinkProgram(void* self, unsigned int program) {
	((QOpenGLFunctions*)self)->glLinkProgram(program);
}
void QOpenGLFunctions_glReleaseShaderCompiler(void* self) {
	((QOpenGLFunctions*)self)->glReleaseShaderCompiler();
}
void QOpenGLFunctions_glRenderbufferStorage(void* self, unsigned int target, unsigned int internalformat, int width, int height) {
	((QOpenGLFunctions*)self)->glRenderbufferStorage(target, internalformat, width, height);
}
void QOpenGLFunctions_glSampleCoverage(void* self, float value, unsigned char invert) {
	((QOpenGLFunctions*)self)->glSampleCoverage(value, invert);
}
void QOpenGLFunctions_glStencilFuncSeparate(void* self, unsigned int face, unsigned int func_, int ref, unsigned int mask) {
	((QOpenGLFunctions*)self)->glStencilFuncSeparate(face, func_, ref, mask);
}
void QOpenGLFunctions_glStencilMaskSeparate(void* self, unsigned int face, unsigned int mask) {
	((QOpenGLFunctions*)self)->glStencilMaskSeparate(face, mask);
}
void QOpenGLFunctions_glStencilOpSeparate(void* self, unsigned int face, unsigned int fail, unsigned int zfail, unsigned int zpass) {
	((QOpenGLFunctions*)self)->glStencilOpSeparate(face, fail, zfail, zpass);
}
void QOpenGLFunctions_glUniform1f(void* self, int location, float x) {
	((QOpenGLFunctions*)self)->glUniform1f(location, x);
}
void QOpenGLFunctions_glUniform1i(void* self, int location, int x) {
	((QOpenGLFunctions*)self)->glUniform1i(location, x);
}
void QOpenGLFunctions_glUniform2f(void* self, int location, float x, float y) {
	((QOpenGLFunctions*)self)->glUniform2f(location, x, y);
}
void QOpenGLFunctions_glUniform2i(void* self, int location, int x, int y) {
	((QOpenGLFunctions*)self)->glUniform2i(location, x, y);
}
void QOpenGLFunctions_glUniform3f(void* self, int location, float x, float y, float z) {
	((QOpenGLFunctions*)self)->glUniform3f(location, x, y, z);
}
void QOpenGLFunctions_glUniform3i(void* self, int location, int x, int y, int z) {
	((QOpenGLFunctions*)self)->glUniform3i(location, x, y, z);
}
void QOpenGLFunctions_glUniform4f(void* self, int location, float x, float y, float z, float w) {
	((QOpenGLFunctions*)self)->glUniform4f(location, x, y, z, w);
}
void QOpenGLFunctions_glUniform4i(void* self, int location, int x, int y, int z, int w) {
	((QOpenGLFunctions*)self)->glUniform4i(location, x, y, z, w);
}
void QOpenGLFunctions_glUseProgram(void* self, unsigned int program) {
	((QOpenGLFunctions*)self)->glUseProgram(program);
}
void QOpenGLFunctions_glValidateProgram(void* self, unsigned int program) {
	((QOpenGLFunctions*)self)->glValidateProgram(program);
}
void QOpenGLFunctions_glVertexAttrib1f(void* self, unsigned int indx, float x) {
	((QOpenGLFunctions*)self)->glVertexAttrib1f(indx, x);
}
void QOpenGLFunctions_glVertexAttrib2f(void* self, unsigned int indx, float x, float y) {
	((QOpenGLFunctions*)self)->glVertexAttrib2f(indx, x, y);
}
void QOpenGLFunctions_glVertexAttrib3f(void* self, unsigned int indx, float x, float y, float z) {
	((QOpenGLFunctions*)self)->glVertexAttrib3f(indx, x, y, z);
}
void QOpenGLFunctions_glVertexAttrib4f(void* self, unsigned int indx, float x, float y, float z, float w) {
	((QOpenGLFunctions*)self)->glVertexAttrib4f(indx, x, y, z, w);
}
void QOpenGLFunctions_delete(void* self) { delete (QOpenGLFunctions*)self; }

}

#include <QtGui/qquaternion.h>
#include "gen_qquaternion__axes.h"

extern "C" {

void* QQuaternion__Axes_x_get(void* self) {
	auto __r = ((QQuaternion::Axes*)self)->x;
	return new QQuaternion::Axis(__r);
}
void QQuaternion__Axes_x_set(void* self, void* x) {
	((QQuaternion::Axes*)self)->x = *(QQuaternion::Axis*)x;
}
void* QQuaternion__Axes_y_get(void* self) {
	auto __r = ((QQuaternion::Axes*)self)->y;
	return new QQuaternion::Axis(__r);
}
void QQuaternion__Axes_y_set(void* self, void* y) {
	((QQuaternion::Axes*)self)->y = *(QQuaternion::Axis*)y;
}
void* QQuaternion__Axes_z_get(void* self) {
	auto __r = ((QQuaternion::Axes*)self)->z;
	return new QQuaternion::Axis(__r);
}
void QQuaternion__Axes_z_set(void* self, void* z) {
	((QQuaternion::Axes*)self)->z = *(QQuaternion::Axis*)z;
}
void QQuaternion__Axes_delete(void* self) { delete (QQuaternion::Axes*)self; }

}

#include <QtGui/qvalidator.h>
#include <QtCore/qobject.h>
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
#include "gen_qvalidator.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QValidator__Go : public QValidator {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QValidator__Go(long long id, A&&... a) : QValidator(std::forward<A>(a)...), __goId(id) {}
	using QValidator::sender;
	using QValidator::senderSignalIndex;
	using QValidator::receivers;
	using QValidator::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QValidator::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QValidator::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QValidator::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QValidator::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QValidator::disconnectNotify(__a0); }
	QValidator::State validate(QString& arg1, int& arg2) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&arg1, (int*)&arg2, &__h); if (__h) return (QValidator::State)(__r); }
		return QValidator::State{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QValidator::metaObject();
	}
	void fixup(QString& arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&arg1, &__h); if (__h) return; }
		QValidator::fixup(arg1);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QValidator::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QValidator::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QValidator::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QValidator::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QValidator::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QValidator::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QValidator::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QValidator;
static std::mutex __connsMu_QValidator;

extern "C" {

void* QValidator_base(void* self) { return (void*)static_cast<QObject*>((QValidator*)self); }
void QValidator__set_override(void* self, int idx) { ((QValidator__Go*)self)->__goOn[idx] = 1; }
void* QValidator_fromQObject(void* o) { return (void*)qobject_cast<QValidator*>((QObject*)o); }
char* QValidator_tr(const char* s, long long* __out_len) {
	auto&& __r = QValidator::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QValidator_setLocale(void* self, void* locale) {
	((QValidator*)self)->setLocale(*(QLocale*)locale);
}
void* QValidator_locale(void* self) {
	auto&& __r = ((QValidator*)self)->locale();
	return new QLocale(__r);
}
int QValidator_validate(void* self, void* arg1, int* arg2) {
	int arg2_ref = *arg2;
	auto&& __r = ((QValidator*)self)->validate(*(QString*)arg1, arg2_ref);
	*arg2 = arg2_ref;
	return (int)(__r);
}
void QValidator_fixup(void* self, void* arg1) {
	((QValidator*)self)->fixup(*(QString*)arg1);
}
void QValidator_changed(void* self) {
	((QValidator*)self)->changed();
}
void* QValidator_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QValidator::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QValidator_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QValidator_sender(void* self) {
	auto&& __r = ((QValidator__Go*)self)->sender();
	return (void*)(__r);
}
int QValidator_senderSignalIndex(void* self) {
	auto&& __r = ((QValidator__Go*)self)->senderSignalIndex();
	return __r;
}
int QValidator_receivers(void* self, const char* signal) {
	auto&& __r = ((QValidator__Go*)self)->receivers(signal);
	return __r;
}
int QValidator_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QValidator__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QValidator_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QValidator::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QValidator_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QValidator::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QValidator::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QValidator_fixup_default(void* self, void* arg1) {
	((QValidator*)self)->QValidator::fixup(*(QString*)arg1);
}
void* QValidator_metaObject_default(void* self) {
	auto&& __r = ((QValidator*)self)->QValidator::metaObject();
	return (void*)(__r);
}
int QValidator_event_default(void* self, void* event) {
	auto&& __r = ((QValidator*)self)->QValidator::event((QEvent*)event);
	return __r;
}
int QValidator_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QValidator*)self)->QValidator::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QValidator_timerEvent_default(void* self, void* event) {
	((QValidator__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QValidator_childEvent_default(void* self, void* event) {
	((QValidator__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QValidator_customEvent_default(void* self, void* event) {
	((QValidator__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QValidator_connectNotify_default(void* self, void* signal) {
	((QValidator__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QValidator_disconnectNotify_default(void* self, void* signal) {
	((QValidator__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QValidator_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QValidator*)self, &QValidator::changed, (QValidator*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QValidator);
	__conns_QValidator[__id] = __c;
}
void QValidator_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QValidator*)self, &QValidator::destroyed, (QValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QValidator);
	__conns_QValidator[__id] = __c;
}
void QValidator_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QValidator*)self, &QValidator::destroyed, (QValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QValidator);
	__conns_QValidator[__id] = __c;
}
void QValidator_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QValidator*)self, &QValidator::objectNameChanged, (QValidator*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QValidator);
	__conns_QValidator[__id] = __c;
}
void QValidator__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QValidator);
		auto __it = __conns_QValidator.find(id);
		if (__it == __conns_QValidator.end()) return;
		__c = __it->second;
		__conns_QValidator.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QValidator__Go_new(long long id, void* parent) { return new QValidator__Go(id, (QObject*)parent); }
void QValidator_delete(void* self) { delete (QValidator*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include "gen_qstyleoptionbutton.h"

extern "C" {

void* QStyleOptionButton_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionButton*)self); }
void* QStyleOptionButton_new() {
	return new QStyleOptionButton();
}
void* QStyleOptionButton_new2(void* other) {
	return new QStyleOptionButton(*(QStyleOptionButton*)other);
}
void* QStyleOptionButton_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionButton*)self)->operator=(*(QStyleOptionButton*)arg1);
	return (void*)&(__r);
}
int QStyleOptionButton_features_get(void* self) {
	auto __r = ((QStyleOptionButton*)self)->features;
	return (int)(__r);
}
void QStyleOptionButton_features_set(void* self, int features) {
	((QStyleOptionButton*)self)->features = QFlags<QStyleOptionButton::ButtonFeature>(features);
}
char* QStyleOptionButton_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionButton*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionButton_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionButton*)self)->text = text_qstr;
}
void* QStyleOptionButton_icon_get(void* self) {
	auto __r = ((QStyleOptionButton*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionButton_icon_set(void* self, void* icon) {
	((QStyleOptionButton*)self)->icon = *(QIcon*)icon;
}
void* QStyleOptionButton_iconSize_get(void* self) {
	auto __r = ((QStyleOptionButton*)self)->iconSize;
	return new QSize(__r);
}
void QStyleOptionButton_iconSize_set(void* self, void* iconSize) {
	((QStyleOptionButton*)self)->iconSize = *(QSize*)iconSize;
}
void QStyleOptionButton_delete(void* self) { delete (QStyleOptionButton*)self; }

}

#include <QtWidgets/qabstractitemview.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtGui/qregion.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qscrollbar.h>
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
#include "gen_qabstractitemview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractItemView__Go : public QAbstractItemView {
public:
	long long __goId;
	unsigned char __goOn[91] = {};
	template<typename... A> QAbstractItemView__Go(long long id, A&&... a) : QAbstractItemView(std::forward<A>(a)...), __goId(id) {}
	using QAbstractItemView::state;
	using QAbstractItemView::scheduleDelayedItemsLayout;
	using QAbstractItemView::executeDelayedItemsLayout;
	using QAbstractItemView::setDirtyRegion;
	using QAbstractItemView::scrollDirtyRegion;
	using QAbstractItemView::dirtyRegionOffset;
	using QAbstractItemView::startAutoScroll;
	using QAbstractItemView::stopAutoScroll;
	using QAbstractItemView::doAutoScroll;
	using QAbstractItemView::dropIndicatorPosition;
	using QAbstractItemView::setViewportMargins;
	using QAbstractItemView::viewportMargins;
	using QAbstractItemView::drawFrame;
	using QAbstractItemView::updateMicroFocus;
	using QAbstractItemView::create;
	using QAbstractItemView::destroy;
	using QAbstractItemView::focusNextChild;
	using QAbstractItemView::focusPreviousChild;
	using QAbstractItemView::sender;
	using QAbstractItemView::senderSignalIndex;
	using QAbstractItemView::receivers;
	using QAbstractItemView::isSignalConnected;
	using QAbstractItemView::getDecodedMetricF;
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QAbstractItemView::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QAbstractItemView::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QAbstractItemView::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QAbstractItemView::currentChanged(__a0, __a1); }
	void cppgoBase_updateEditorData() { QAbstractItemView::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QAbstractItemView::updateEditorGeometries(); }
	void cppgoBase_updateGeometries() { QAbstractItemView::updateGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QAbstractItemView::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QAbstractItemView::horizontalScrollbarAction(__a0); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QAbstractItemView::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QAbstractItemView::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QAbstractItemView::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QAbstractItemView::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QAbstractItemView::editorDestroyed(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QAbstractItemView::selectedIndexes(); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QAbstractItemView::selectionCommand(__a0, __a1); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QAbstractItemView::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QAbstractItemView::initViewItemOption(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QAbstractItemView::focusNextPrevChild(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QAbstractItemView::event(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QAbstractItemView::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QAbstractItemView::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QAbstractItemView::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QAbstractItemView::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QAbstractItemView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QAbstractItemView::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QAbstractItemView::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QAbstractItemView::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QAbstractItemView::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QAbstractItemView::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QAbstractItemView::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QAbstractItemView::keyPressEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QAbstractItemView::resizeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractItemView::timerEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QAbstractItemView::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QAbstractItemView::eventFilter(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QAbstractItemView::viewportSizeHint(); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QAbstractItemView::paintEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QAbstractItemView::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QAbstractItemView::contextMenuEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QAbstractItemView::scrollContentsBy(__a0, __a1); }
	void cppgoBase_changeEvent(QEvent* __a0) { QAbstractItemView::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QAbstractItemView::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QAbstractItemView::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QAbstractItemView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QAbstractItemView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QAbstractItemView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QAbstractItemView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QAbstractItemView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QAbstractItemView::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QAbstractItemView::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QAbstractItemView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QAbstractItemView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QAbstractItemView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QAbstractItemView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QAbstractItemView::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractItemView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractItemView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractItemView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractItemView::disconnectNotify(__a0); }
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QRect{};
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		return;
	}
	QModelIndex indexAt(const QPoint& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&point, &__h); if (__h) return *(QModelIndex*)__r; }
		return QModelIndex{};
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QModelIndex{};
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return int{};
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return int{};
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return bool{};
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		return;
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QRegion{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractItemView::metaObject();
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QAbstractItemView::setModel(model);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QAbstractItemView::setSelectionModel(selectionModel);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[12] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QAbstractItemView::keyboardSearch(search);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QAbstractItemView::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QAbstractItemView::sizeHintForColumn(column);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QAbstractItemView::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractItemView::inputMethodQuery(query);
	}
	void reset() override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QAbstractItemView::reset();
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QAbstractItemView::setRootIndex(index);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QAbstractItemView::doItemsLayout();
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QAbstractItemView::selectAll();
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QAbstractItemView::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QAbstractItemView::rowsAboutToBeRemoved(parent, start, end);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QAbstractItemView::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QAbstractItemView::currentChanged(current, previous);
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QAbstractItemView::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QAbstractItemView::updateEditorGeometries();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QAbstractItemView::updateGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QAbstractItemView::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QAbstractItemView::horizontalScrollbarAction(action);
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QAbstractItemView::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QAbstractItemView::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QAbstractItemView::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QAbstractItemView::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QAbstractItemView::editorDestroyed(editor);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virtp_voidp) {
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
		return QAbstractItemView::selectedIndexes();
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemView::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QAbstractItemView::selectionCommand(index, event);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QAbstractItemView::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QAbstractItemView::initViewItemOption(option);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QAbstractItemView::focusNextPrevChild(next);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemView::event(event);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemView::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::mousePressEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::keyPressEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::resizeEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::timerEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemView::eventFilter(object, event);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractItemView::viewportSizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractItemView::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractItemView::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QAbstractItemView::setupViewport(viewport);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractItemView::paintEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractItemView::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractItemView::contextMenuEvent(arg1);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QAbstractItemView::scrollContentsBy(dx, dy);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractItemView::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QAbstractItemView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QAbstractItemView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QAbstractItemView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QAbstractItemView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QAbstractItemView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QAbstractItemView::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QAbstractItemView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QAbstractItemView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QAbstractItemView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QAbstractItemView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QAbstractItemView::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractItemView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractItemView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractItemView;
static std::mutex __connsMu_QAbstractItemView;

extern "C" {

void* QAbstractItemView_base(void* self) { return (void*)static_cast<QAbstractScrollArea*>((QAbstractItemView*)self); }
void QAbstractItemView_chain(void* self, void** out) {
	QAbstractItemView* __p = (QAbstractItemView*)self;
	QAbstractScrollArea* __b0 = static_cast<QAbstractScrollArea*>(__p); out[0] = (void*)__b0;
	QFrame* __b1 = static_cast<QFrame*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QAbstractItemView__set_override(void* self, int idx) { ((QAbstractItemView__Go*)self)->__goOn[idx] = 1; }
void* QAbstractItemView_fromQObject(void* o) { return (void*)qobject_cast<QAbstractItemView*>((QObject*)o); }
char* QAbstractItemView_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractItemView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractItemView_setModel(void* self, void* model) {
	((QAbstractItemView*)self)->setModel((QAbstractItemModel*)model);
}
void* QAbstractItemView_model(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->model();
	return (void*)(__r);
}
void QAbstractItemView_setSelectionModel(void* self, void* selectionModel) {
	((QAbstractItemView*)self)->setSelectionModel((QItemSelectionModel*)selectionModel);
}
void* QAbstractItemView_selectionModel(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->selectionModel();
	return (void*)(__r);
}
void QAbstractItemView_setItemDelegate(void* self, void* delegate) {
	((QAbstractItemView*)self)->setItemDelegate((QAbstractItemDelegate*)delegate);
}
void* QAbstractItemView_itemDelegate(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->itemDelegate();
	return (void*)(__r);
}
void QAbstractItemView_setSelectionMode(void* self, int mode) {
	((QAbstractItemView*)self)->setSelectionMode((QAbstractItemView::SelectionMode)(mode));
}
int QAbstractItemView_selectionMode(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->selectionMode();
	return (int)(__r);
}
void QAbstractItemView_setSelectionBehavior(void* self, int behavior) {
	((QAbstractItemView*)self)->setSelectionBehavior((QAbstractItemView::SelectionBehavior)(behavior));
}
int QAbstractItemView_selectionBehavior(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->selectionBehavior();
	return (int)(__r);
}
void* QAbstractItemView_currentIndex(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->currentIndex();
	return new QModelIndex(__r);
}
void* QAbstractItemView_rootIndex(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->rootIndex();
	return new QModelIndex(__r);
}
void QAbstractItemView_setEditTriggers(void* self, int triggers) {
	((QAbstractItemView*)self)->setEditTriggers(QFlags<QAbstractItemView::EditTrigger>(triggers));
}
int QAbstractItemView_editTriggers(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->editTriggers();
	return (int)(__r);
}
void QAbstractItemView_setVerticalScrollMode(void* self, int mode) {
	((QAbstractItemView*)self)->setVerticalScrollMode((QAbstractItemView::ScrollMode)(mode));
}
int QAbstractItemView_verticalScrollMode(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->verticalScrollMode();
	return (int)(__r);
}
void QAbstractItemView_resetVerticalScrollMode(void* self) {
	((QAbstractItemView*)self)->resetVerticalScrollMode();
}
void QAbstractItemView_setHorizontalScrollMode(void* self, int mode) {
	((QAbstractItemView*)self)->setHorizontalScrollMode((QAbstractItemView::ScrollMode)(mode));
}
int QAbstractItemView_horizontalScrollMode(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->horizontalScrollMode();
	return (int)(__r);
}
void QAbstractItemView_resetHorizontalScrollMode(void* self) {
	((QAbstractItemView*)self)->resetHorizontalScrollMode();
}
void QAbstractItemView_setAutoScroll(void* self, int enable) {
	((QAbstractItemView*)self)->setAutoScroll(enable);
}
int QAbstractItemView_hasAutoScroll(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->hasAutoScroll();
	return __r;
}
void QAbstractItemView_setAutoScrollMargin(void* self, int margin) {
	((QAbstractItemView*)self)->setAutoScrollMargin(margin);
}
int QAbstractItemView_autoScrollMargin(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->autoScrollMargin();
	return __r;
}
void QAbstractItemView_setTabKeyNavigation(void* self, int enable) {
	((QAbstractItemView*)self)->setTabKeyNavigation(enable);
}
int QAbstractItemView_tabKeyNavigation(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->tabKeyNavigation();
	return __r;
}
void QAbstractItemView_setDropIndicatorShown(void* self, int enable) {
	((QAbstractItemView*)self)->setDropIndicatorShown(enable);
}
int QAbstractItemView_showDropIndicator(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->showDropIndicator();
	return __r;
}
void QAbstractItemView_setDragEnabled(void* self, int enable) {
	((QAbstractItemView*)self)->setDragEnabled(enable);
}
int QAbstractItemView_dragEnabled(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->dragEnabled();
	return __r;
}
void QAbstractItemView_setDragDropOverwriteMode(void* self, int overwrite) {
	((QAbstractItemView*)self)->setDragDropOverwriteMode(overwrite);
}
int QAbstractItemView_dragDropOverwriteMode(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->dragDropOverwriteMode();
	return __r;
}
void QAbstractItemView_setDragDropMode(void* self, int behavior) {
	((QAbstractItemView*)self)->setDragDropMode((QAbstractItemView::DragDropMode)(behavior));
}
int QAbstractItemView_dragDropMode(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->dragDropMode();
	return (int)(__r);
}
void QAbstractItemView_setDefaultDropAction(void* self, int dropAction) {
	((QAbstractItemView*)self)->setDefaultDropAction((Qt::DropAction)(dropAction));
}
int QAbstractItemView_defaultDropAction(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->defaultDropAction();
	return (int)(__r);
}
void QAbstractItemView_setAlternatingRowColors(void* self, int enable) {
	((QAbstractItemView*)self)->setAlternatingRowColors(enable);
}
int QAbstractItemView_alternatingRowColors(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->alternatingRowColors();
	return __r;
}
void QAbstractItemView_setIconSize(void* self, void* size) {
	((QAbstractItemView*)self)->setIconSize(*(QSize*)size);
}
void* QAbstractItemView_iconSize(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->iconSize();
	return new QSize(__r);
}
void QAbstractItemView_setTextElideMode(void* self, int mode) {
	((QAbstractItemView*)self)->setTextElideMode((Qt::TextElideMode)(mode));
}
int QAbstractItemView_textElideMode(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->textElideMode();
	return (int)(__r);
}
void QAbstractItemView_keyboardSearch(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QAbstractItemView*)self)->keyboardSearch(search_qstr);
}
void* QAbstractItemView_visualRect(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QAbstractItemView_scrollTo(void* self, void* index, int hint) {
	((QAbstractItemView*)self)->scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QAbstractItemView_indexAt(void* self, void* point) {
	auto&& __r = ((QAbstractItemView*)self)->indexAt(*(QPoint*)point);
	return new QModelIndex(__r);
}
void* QAbstractItemView_sizeHintForIndex(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->sizeHintForIndex(*(QModelIndex*)index);
	return new QSize(__r);
}
int QAbstractItemView_sizeHintForRow(void* self, int row) {
	auto&& __r = ((QAbstractItemView*)self)->sizeHintForRow(row);
	return __r;
}
int QAbstractItemView_sizeHintForColumn(void* self, int column) {
	auto&& __r = ((QAbstractItemView*)self)->sizeHintForColumn(column);
	return __r;
}
int QAbstractItemView_updateThreshold(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->updateThreshold();
	return __r;
}
void QAbstractItemView_setUpdateThreshold(void* self, int threshold) {
	((QAbstractItemView*)self)->setUpdateThreshold(threshold);
}
int QAbstractItemView_keyboardSearchFlags(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->keyboardSearchFlags();
	return (int)(__r);
}
void QAbstractItemView_setKeyboardSearchFlags(void* self, int searchFlags) {
	((QAbstractItemView*)self)->setKeyboardSearchFlags(QFlags<Qt::MatchFlag>(searchFlags));
}
void QAbstractItemView_openPersistentEditor(void* self, void* index) {
	((QAbstractItemView*)self)->openPersistentEditor(*(QModelIndex*)index);
}
void QAbstractItemView_closePersistentEditor(void* self, void* index) {
	((QAbstractItemView*)self)->closePersistentEditor(*(QModelIndex*)index);
}
int QAbstractItemView_isPersistentEditorOpen(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->isPersistentEditorOpen(*(QModelIndex*)index);
	return __r;
}
void QAbstractItemView_setIndexWidget(void* self, void* index, void* widget) {
	((QAbstractItemView*)self)->setIndexWidget(*(QModelIndex*)index, (QWidget*)widget);
}
void* QAbstractItemView_indexWidget(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->indexWidget(*(QModelIndex*)index);
	return (void*)(__r);
}
void QAbstractItemView_setItemDelegateForRow(void* self, int row, void* delegate) {
	((QAbstractItemView*)self)->setItemDelegateForRow(row, (QAbstractItemDelegate*)delegate);
}
void* QAbstractItemView_itemDelegateForRow(void* self, int row) {
	auto&& __r = ((QAbstractItemView*)self)->itemDelegateForRow(row);
	return (void*)(__r);
}
void QAbstractItemView_setItemDelegateForColumn(void* self, int column, void* delegate) {
	((QAbstractItemView*)self)->setItemDelegateForColumn(column, (QAbstractItemDelegate*)delegate);
}
void* QAbstractItemView_itemDelegateForColumn(void* self, int column) {
	auto&& __r = ((QAbstractItemView*)self)->itemDelegateForColumn(column);
	return (void*)(__r);
}
void* QAbstractItemView_itemDelegate2(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->itemDelegate(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QAbstractItemView_itemDelegateForIndex(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void QAbstractItemView_reset(void* self) {
	((QAbstractItemView*)self)->reset();
}
void QAbstractItemView_setRootIndex(void* self, void* index) {
	((QAbstractItemView*)self)->setRootIndex(*(QModelIndex*)index);
}
void QAbstractItemView_doItemsLayout(void* self) {
	((QAbstractItemView*)self)->doItemsLayout();
}
void QAbstractItemView_selectAll(void* self) {
	((QAbstractItemView*)self)->selectAll();
}
void QAbstractItemView_edit(void* self, void* index) {
	((QAbstractItemView*)self)->edit(*(QModelIndex*)index);
}
void QAbstractItemView_clearSelection(void* self) {
	((QAbstractItemView*)self)->clearSelection();
}
void QAbstractItemView_setCurrentIndex(void* self, void* index) {
	((QAbstractItemView*)self)->setCurrentIndex(*(QModelIndex*)index);
}
void QAbstractItemView_scrollToTop(void* self) {
	((QAbstractItemView*)self)->scrollToTop();
}
void QAbstractItemView_scrollToBottom(void* self) {
	((QAbstractItemView*)self)->scrollToBottom();
}
void QAbstractItemView_update(void* self, void* index) {
	((QAbstractItemView*)self)->update(*(QModelIndex*)index);
}
void QAbstractItemView_pressed(void* self, void* index) {
	((QAbstractItemView*)self)->pressed(*(QModelIndex*)index);
}
void QAbstractItemView_clicked(void* self, void* index) {
	((QAbstractItemView*)self)->clicked(*(QModelIndex*)index);
}
void QAbstractItemView_doubleClicked(void* self, void* index) {
	((QAbstractItemView*)self)->doubleClicked(*(QModelIndex*)index);
}
void QAbstractItemView_activated(void* self, void* index) {
	((QAbstractItemView*)self)->activated(*(QModelIndex*)index);
}
void QAbstractItemView_entered(void* self, void* index) {
	((QAbstractItemView*)self)->entered(*(QModelIndex*)index);
}
void QAbstractItemView_viewportEntered(void* self) {
	((QAbstractItemView*)self)->viewportEntered();
}
void QAbstractItemView_iconSizeChanged(void* self, void* size) {
	((QAbstractItemView*)self)->iconSizeChanged(*(QSize*)size);
}
int QAbstractItemView_state(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->state();
	return (int)(__r);
}
void QAbstractItemView_scheduleDelayedItemsLayout(void* self) {
	((QAbstractItemView__Go*)self)->scheduleDelayedItemsLayout();
}
void QAbstractItemView_executeDelayedItemsLayout(void* self) {
	((QAbstractItemView__Go*)self)->executeDelayedItemsLayout();
}
void QAbstractItemView_setDirtyRegion(void* self, void* region) {
	((QAbstractItemView__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QAbstractItemView_scrollDirtyRegion(void* self, int dx, int dy) {
	((QAbstractItemView__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QAbstractItemView_dirtyRegionOffset(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QAbstractItemView_startAutoScroll(void* self) {
	((QAbstractItemView__Go*)self)->startAutoScroll();
}
void QAbstractItemView_stopAutoScroll(void* self) {
	((QAbstractItemView__Go*)self)->stopAutoScroll();
}
void QAbstractItemView_doAutoScroll(void* self) {
	((QAbstractItemView__Go*)self)->doAutoScroll();
}
int QAbstractItemView_dropIndicatorPosition(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QAbstractItemView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QAbstractItemView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QAbstractItemView_setViewportMargins2(void* self, void* margins) {
	((QAbstractItemView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QAbstractItemView_viewportMargins(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QAbstractItemView_drawFrame(void* self, void* arg1) {
	((QAbstractItemView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QAbstractItemView_setTabOrder(void* arg1, void* arg2) {
	QAbstractItemView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QAbstractItemView_mouseGrabber() {
	auto&& __r = QAbstractItemView::mouseGrabber();
	return (void*)(__r);
}
void* QAbstractItemView_keyboardGrabber() {
	auto&& __r = QAbstractItemView::keyboardGrabber();
	return (void*)(__r);
}
void* QAbstractItemView_find(unsigned long long arg1) {
	auto&& __r = QAbstractItemView::find(arg1);
	return (void*)(__r);
}
void* QAbstractItemView_createWindowContainer(void* window) {
	auto&& __r = QAbstractItemView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QAbstractItemView_updateMicroFocus(void* self) {
	((QAbstractItemView__Go*)self)->updateMicroFocus();
}
void QAbstractItemView_create(void* self) {
	((QAbstractItemView__Go*)self)->create();
}
void QAbstractItemView_destroy(void* self) {
	((QAbstractItemView__Go*)self)->destroy();
}
int QAbstractItemView_focusNextChild(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->focusNextChild();
	return __r;
}
int QAbstractItemView_focusPreviousChild(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QAbstractItemView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractItemView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractItemView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemView_sender(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractItemView_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractItemView_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractItemView__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractItemView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractItemView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QAbstractItemView_devicePixelRatioFScale() {
	auto&& __r = QAbstractItemView::devicePixelRatioFScale();
	return __r;
}
int QAbstractItemView_encodeMetricF(int metric, double value) {
	auto&& __r = QAbstractItemView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QAbstractItemView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QAbstractItemView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QAbstractItemView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractItemView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractItemView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractItemView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractItemView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QAbstractItemView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QAbstractItemView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QAbstractItemView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QAbstractItemView_updateMicroFocus2(void* self, int query) {
	((QAbstractItemView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QAbstractItemView_create2(void* self, unsigned long long arg1) {
	((QAbstractItemView__Go*)self)->create(arg1);
}
void QAbstractItemView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QAbstractItemView__Go*)self)->create(arg1, initializeWindow);
}
void QAbstractItemView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QAbstractItemView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QAbstractItemView_destroy2(void* self, int destroyWindow) {
	((QAbstractItemView__Go*)self)->destroy(destroyWindow);
}
void QAbstractItemView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QAbstractItemView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QAbstractItemView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractItemView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractItemView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAbstractItemView_setModel_default(void* self, void* model) {
	((QAbstractItemView*)self)->QAbstractItemView::setModel((QAbstractItemModel*)model);
}
void QAbstractItemView_setSelectionModel_default(void* self, void* selectionModel) {
	((QAbstractItemView*)self)->QAbstractItemView::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QAbstractItemView_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QAbstractItemView*)self)->QAbstractItemView::keyboardSearch(search_qstr);
}
int QAbstractItemView_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::sizeHintForRow(row);
	return __r;
}
int QAbstractItemView_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::sizeHintForColumn(column);
	return __r;
}
void* QAbstractItemView_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void QAbstractItemView_reset_default(void* self) {
	((QAbstractItemView*)self)->QAbstractItemView::reset();
}
void QAbstractItemView_setRootIndex_default(void* self, void* index) {
	((QAbstractItemView*)self)->QAbstractItemView::setRootIndex(*(QModelIndex*)index);
}
void QAbstractItemView_doItemsLayout_default(void* self) {
	((QAbstractItemView*)self)->QAbstractItemView::doItemsLayout();
}
void QAbstractItemView_selectAll_default(void* self) {
	((QAbstractItemView*)self)->QAbstractItemView::selectAll();
}
void QAbstractItemView_edit_default(void* self, void* index) {
	((QAbstractItemView*)self)->QAbstractItemView::edit(*(QModelIndex*)index);
}
void* QAbstractItemView_metaObject_default(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::metaObject();
	return (void*)(__r);
}
void* QAbstractItemView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QAbstractItemView_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QAbstractItemView__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QAbstractItemView_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QAbstractItemView__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QAbstractItemView_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QAbstractItemView__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QAbstractItemView_currentChanged_default(void* self, void* current, void* previous) {
	((QAbstractItemView__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QAbstractItemView_updateEditorData_default(void* self) {
	((QAbstractItemView__Go*)self)->cppgoBase_updateEditorData();
}
void QAbstractItemView_updateEditorGeometries_default(void* self) {
	((QAbstractItemView__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QAbstractItemView_updateGeometries_default(void* self) {
	((QAbstractItemView__Go*)self)->cppgoBase_updateGeometries();
}
void QAbstractItemView_verticalScrollbarAction_default(void* self, int action) {
	((QAbstractItemView__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QAbstractItemView_horizontalScrollbarAction_default(void* self, int action) {
	((QAbstractItemView__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
void QAbstractItemView_verticalScrollbarValueChanged_default(void* self, int value) {
	((QAbstractItemView__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QAbstractItemView_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QAbstractItemView__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QAbstractItemView_closeEditor_default(void* self, void* editor, int hint) {
	((QAbstractItemView__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QAbstractItemView_commitData_default(void* self, void* editor) {
	((QAbstractItemView__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QAbstractItemView_editorDestroyed_default(void* self, void* editor) {
	((QAbstractItemView__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
void** QAbstractItemView_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
int QAbstractItemView_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
void QAbstractItemView_startDrag_default(void* self, int supportedActions) {
	((QAbstractItemView__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QAbstractItemView_initViewItemOption_default(void* self, void* option) {
	((QAbstractItemView__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
int QAbstractItemView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QAbstractItemView_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QAbstractItemView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QAbstractItemView_mousePressEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QAbstractItemView_mouseMoveEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QAbstractItemView_mouseReleaseEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QAbstractItemView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QAbstractItemView_dragEnterEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QAbstractItemView_dragMoveEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QAbstractItemView_dragLeaveEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QAbstractItemView_dropEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QAbstractItemView_focusInEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QAbstractItemView_focusOutEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QAbstractItemView_keyPressEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QAbstractItemView_resizeEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QAbstractItemView_timerEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractItemView_inputMethodEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QAbstractItemView_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QAbstractItemView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void* QAbstractItemView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::minimumSizeHint();
	return new QSize(__r);
}
void* QAbstractItemView_sizeHint_default(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::sizeHint();
	return new QSize(__r);
}
void QAbstractItemView_setupViewport_default(void* self, void* viewport) {
	((QAbstractItemView*)self)->QAbstractItemView::setupViewport((QWidget*)viewport);
}
void QAbstractItemView_paintEvent_default(void* self, void* arg1) {
	((QAbstractItemView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QAbstractItemView_wheelEvent_default(void* self, void* arg1) {
	((QAbstractItemView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QAbstractItemView_contextMenuEvent_default(void* self, void* arg1) {
	((QAbstractItemView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QAbstractItemView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QAbstractItemView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QAbstractItemView_changeEvent_default(void* self, void* arg1) {
	((QAbstractItemView__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QAbstractItemView_initStyleOption_default(void* self, void* option) {
	((QAbstractItemView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QAbstractItemView_devType_default(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::devType();
	return __r;
}
void QAbstractItemView_setVisible_default(void* self, int visible) {
	((QAbstractItemView*)self)->QAbstractItemView::setVisible(visible);
}
int QAbstractItemView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::heightForWidth(arg1);
	return __r;
}
int QAbstractItemView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::hasHeightForWidth();
	return __r;
}
void* QAbstractItemView_paintEngine_default(void* self) {
	auto&& __r = ((QAbstractItemView*)self)->QAbstractItemView::paintEngine();
	return (void*)(__r);
}
void QAbstractItemView_keyReleaseEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QAbstractItemView_enterEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QAbstractItemView_leaveEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QAbstractItemView_moveEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QAbstractItemView_closeEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QAbstractItemView_tabletEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QAbstractItemView_actionEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QAbstractItemView_showEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QAbstractItemView_hideEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QAbstractItemView_metric_default(void* self, int arg1) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QAbstractItemView_initPainter_default(void* self, void* painter) {
	((QAbstractItemView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QAbstractItemView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QAbstractItemView_sharedPainter_default(void* self) {
	auto&& __r = ((QAbstractItemView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QAbstractItemView_childEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractItemView_customEvent_default(void* self, void* event) {
	((QAbstractItemView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractItemView_connectNotify_default(void* self, void* signal) {
	((QAbstractItemView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractItemView_disconnectNotify_default(void* self, void* signal) {
	((QAbstractItemView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractItemView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::pressed, (QAbstractItemView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::clicked, (QAbstractItemView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::doubleClicked, (QAbstractItemView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::activated, (QAbstractItemView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::entered, (QAbstractItemView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::viewportEntered, (QAbstractItemView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::iconSizeChanged, (QAbstractItemView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::windowTitleChanged, (QAbstractItemView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::windowIconChanged, (QAbstractItemView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::windowIconTextChanged, (QAbstractItemView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::customContextMenuRequested, (QAbstractItemView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::destroyed, (QAbstractItemView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::destroyed, (QAbstractItemView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemView*)self, &QAbstractItemView::objectNameChanged, (QAbstractItemView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
	__conns_QAbstractItemView[__id] = __c;
}
void QAbstractItemView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemView);
		auto __it = __conns_QAbstractItemView.find(id);
		if (__it == __conns_QAbstractItemView.end()) return;
		__c = __it->second;
		__conns_QAbstractItemView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractItemView__Go_new(long long id) { return new QAbstractItemView__Go(id); }
void QAbstractItemView_delete(void* self) { delete (QAbstractItemView*)self; }

}

#include <QtWidgets/qcompleter.h>
#include <QtCore/qobject.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qcompleter.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QCompleter__Go : public QCompleter {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QCompleter__Go(long long id, A&&... a) : QCompleter(std::forward<A>(a)...), __goId(id) {}
	using QCompleter::sender;
	using QCompleter::senderSignalIndex;
	using QCompleter::receivers;
	using QCompleter::isSignalConnected;
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QCompleter::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QCompleter::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QCompleter::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QCompleter::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QCompleter::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QCompleter::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QCompleter::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QCompleter::metaObject();
	}
	QString pathFromIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_charp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_charp_voidp)(__goId, -4079497074265104506LL, (void*)&index, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QCompleter::pathFromIndex(index);
	}
	QList<QString> splitPath(const QString& path) const override {
		int __h = 0;
		QByteArray path_ba = path.toUtf8();
		if (__goOn[2] && cppgo_virtp_voidp_constcharp_longlong) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, const char*, long long, long long*, int*))cppgo_virtp_voidp_constcharp_longlong)(__goId, -7571154661167114460LL, (char*)path_ba.constData(), path_ba.size(), &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QCompleter::splitPath(path);
	}
	bool eventFilter(QObject* o, QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)o, (void*)e, &__h); if (__h) return __r; }
		return QCompleter::eventFilter(o, e);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QCompleter::event(arg1);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QCompleter::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QCompleter::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QCompleter::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QCompleter::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QCompleter::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QCompleter;
static std::mutex __connsMu_QCompleter;

extern "C" {

void* QCompleter_base(void* self) { return (void*)static_cast<QObject*>((QCompleter*)self); }
void QCompleter__set_override(void* self, int idx) { ((QCompleter__Go*)self)->__goOn[idx] = 1; }
void* QCompleter_fromQObject(void* o) { return (void*)qobject_cast<QCompleter*>((QObject*)o); }
void* QCompleter_new(long long __goid) {
	return new QCompleter__Go(__goid);
}
void* QCompleter_new2(long long __goid, void* model) {
	return new QCompleter__Go(__goid, (QAbstractItemModel*)model);
}
void* QCompleter_new3(long long __goid, char** completions, long long completions_len) {
	QList<QString> completions_list; completions_list.reserve(completions_len);
	for (long long __i = 0; __i < completions_len; __i++) completions_list.append(QString::fromUtf8(completions[__i]));
	return new QCompleter__Go(__goid, completions_list);
}
void* QCompleter_new4(long long __goid, void* parent) {
	return new QCompleter__Go(__goid, (QObject*)parent);
}
void* QCompleter_new5(long long __goid, void* model, void* parent) {
	return new QCompleter__Go(__goid, (QAbstractItemModel*)model, (QObject*)parent);
}
void* QCompleter_new6(long long __goid, char** completions, long long completions_len, void* parent) {
	QList<QString> completions_list; completions_list.reserve(completions_len);
	for (long long __i = 0; __i < completions_len; __i++) completions_list.append(QString::fromUtf8(completions[__i]));
	return new QCompleter__Go(__goid, completions_list, (QObject*)parent);
}
char* QCompleter_tr(const char* s, long long* __out_len) {
	auto&& __r = QCompleter::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCompleter_setWidget(void* self, void* widget) {
	((QCompleter*)self)->setWidget((QWidget*)widget);
}
void* QCompleter_widget(void* self) {
	auto&& __r = ((QCompleter*)self)->widget();
	return (void*)(__r);
}
void QCompleter_setModel(void* self, void* c) {
	((QCompleter*)self)->setModel((QAbstractItemModel*)c);
}
void* QCompleter_model(void* self) {
	auto&& __r = ((QCompleter*)self)->model();
	return (void*)(__r);
}
void QCompleter_setCompletionMode(void* self, int mode) {
	((QCompleter*)self)->setCompletionMode((QCompleter::CompletionMode)(mode));
}
int QCompleter_completionMode(void* self) {
	auto&& __r = ((QCompleter*)self)->completionMode();
	return (int)(__r);
}
void QCompleter_setFilterMode(void* self, int filterMode) {
	((QCompleter*)self)->setFilterMode(QFlags<Qt::MatchFlag>(filterMode));
}
int QCompleter_filterMode(void* self) {
	auto&& __r = ((QCompleter*)self)->filterMode();
	return (int)(__r);
}
void* QCompleter_popup(void* self) {
	auto&& __r = ((QCompleter*)self)->popup();
	return (void*)(__r);
}
void QCompleter_setPopup(void* self, void* popup) {
	((QCompleter*)self)->setPopup((QAbstractItemView*)popup);
}
void QCompleter_setCaseSensitivity(void* self, int caseSensitivity) {
	((QCompleter*)self)->setCaseSensitivity((Qt::CaseSensitivity)(caseSensitivity));
}
int QCompleter_caseSensitivity(void* self) {
	auto&& __r = ((QCompleter*)self)->caseSensitivity();
	return (int)(__r);
}
void QCompleter_setModelSorting(void* self, int sorting) {
	((QCompleter*)self)->setModelSorting((QCompleter::ModelSorting)(sorting));
}
int QCompleter_modelSorting(void* self) {
	auto&& __r = ((QCompleter*)self)->modelSorting();
	return (int)(__r);
}
void QCompleter_setCompletionColumn(void* self, int column) {
	((QCompleter*)self)->setCompletionColumn(column);
}
int QCompleter_completionColumn(void* self) {
	auto&& __r = ((QCompleter*)self)->completionColumn();
	return __r;
}
void QCompleter_setCompletionRole(void* self, int role) {
	((QCompleter*)self)->setCompletionRole(role);
}
int QCompleter_completionRole(void* self) {
	auto&& __r = ((QCompleter*)self)->completionRole();
	return __r;
}
int QCompleter_wrapAround(void* self) {
	auto&& __r = ((QCompleter*)self)->wrapAround();
	return __r;
}
int QCompleter_maxVisibleItems(void* self) {
	auto&& __r = ((QCompleter*)self)->maxVisibleItems();
	return __r;
}
void QCompleter_setMaxVisibleItems(void* self, int maxItems) {
	((QCompleter*)self)->setMaxVisibleItems(maxItems);
}
int QCompleter_completionCount(void* self) {
	auto&& __r = ((QCompleter*)self)->completionCount();
	return __r;
}
int QCompleter_setCurrentRow(void* self, int row) {
	auto&& __r = ((QCompleter*)self)->setCurrentRow(row);
	return __r;
}
int QCompleter_currentRow(void* self) {
	auto&& __r = ((QCompleter*)self)->currentRow();
	return __r;
}
void* QCompleter_currentIndex(void* self) {
	auto&& __r = ((QCompleter*)self)->currentIndex();
	return new QModelIndex(__r);
}
char* QCompleter_currentCompletion(void* self, long long* __out_len) {
	auto&& __r = ((QCompleter*)self)->currentCompletion();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCompleter_completionModel(void* self) {
	auto&& __r = ((QCompleter*)self)->completionModel();
	return (void*)(__r);
}
char* QCompleter_completionPrefix(void* self, long long* __out_len) {
	auto&& __r = ((QCompleter*)self)->completionPrefix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCompleter_setCompletionPrefix(void* self, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	((QCompleter*)self)->setCompletionPrefix(prefix_qstr);
}
void QCompleter_complete(void* self) {
	((QCompleter*)self)->complete();
}
void QCompleter_setWrapAround(void* self, int wrap) {
	((QCompleter*)self)->setWrapAround(wrap);
}
char* QCompleter_pathFromIndex(void* self, void* index, long long* __out_len) {
	auto&& __r = ((QCompleter*)self)->pathFromIndex(*(QModelIndex*)index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QCompleter_splitPath(void* self, const char* path, long long path_len, long long* __out_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QCompleter*)self)->splitPath(path_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QCompleter_activated(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QCompleter*)self)->activated(text_qstr);
}
void QCompleter_activated2(void* self, void* index) {
	((QCompleter*)self)->activated(*(QModelIndex*)index);
}
void QCompleter_highlighted(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QCompleter*)self)->highlighted(text_qstr);
}
void QCompleter_highlighted2(void* self, void* index) {
	((QCompleter*)self)->highlighted(*(QModelIndex*)index);
}
void* QCompleter_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QCompleter::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QCompleter_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QCompleter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QCompleter_sender(void* self) {
	auto&& __r = ((QCompleter__Go*)self)->sender();
	return (void*)(__r);
}
int QCompleter_senderSignalIndex(void* self) {
	auto&& __r = ((QCompleter__Go*)self)->senderSignalIndex();
	return __r;
}
int QCompleter_receivers(void* self, const char* signal) {
	auto&& __r = ((QCompleter__Go*)self)->receivers(signal);
	return __r;
}
int QCompleter_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QCompleter__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QCompleter_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QCompleter::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCompleter_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QCompleter::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCompleter_complete2(void* self, void* rect) {
	((QCompleter*)self)->complete(*(QRect*)rect);
}
void* QCompleter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QCompleter::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QCompleter_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QCompleter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
char* QCompleter_pathFromIndex_default(void* self, void* index, long long* __out_len) {
	auto&& __r = ((QCompleter*)self)->QCompleter::pathFromIndex(*(QModelIndex*)index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QCompleter_splitPath_default(void* self, const char* path, long long path_len, long long* __out_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QCompleter*)self)->QCompleter::splitPath(path_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QCompleter_metaObject_default(void* self) {
	auto&& __r = ((QCompleter*)self)->QCompleter::metaObject();
	return (void*)(__r);
}
int QCompleter_eventFilter_default(void* self, void* o, void* e) {
	auto&& __r = ((QCompleter__Go*)self)->cppgoBase_eventFilter((QObject*)o, (QEvent*)e);
	return __r;
}
int QCompleter_event_default(void* self, void* arg1) {
	auto&& __r = ((QCompleter__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QCompleter_timerEvent_default(void* self, void* event) {
	((QCompleter__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QCompleter_childEvent_default(void* self, void* event) {
	((QCompleter__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QCompleter_customEvent_default(void* self, void* event) {
	((QCompleter__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QCompleter_connectNotify_default(void* self, void* signal) {
	((QCompleter__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QCompleter_disconnectNotify_default(void* self, void* signal) {
	((QCompleter__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QCompleter_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, static_cast<void(QCompleter::*)(const QString&)>(&QCompleter::activated), (QCompleter*)self, [__id](const QString& text){
		QByteArray text_ba = text.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)text_ba.constData(), text_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter_connect_activated2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, static_cast<void(QCompleter::*)(const QModelIndex&)>(&QCompleter::activated), (QCompleter*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter_connect_highlighted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, static_cast<void(QCompleter::*)(const QString&)>(&QCompleter::highlighted), (QCompleter*)self, [__id](const QString& text){
		QByteArray text_ba = text.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)text_ba.constData(), text_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter_connect_highlighted2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, static_cast<void(QCompleter::*)(const QModelIndex&)>(&QCompleter::highlighted), (QCompleter*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, &QCompleter::destroyed, (QCompleter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, &QCompleter::destroyed, (QCompleter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCompleter*)self, &QCompleter::objectNameChanged, (QCompleter*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
	__conns_QCompleter[__id] = __c;
}
void QCompleter__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QCompleter);
		auto __it = __conns_QCompleter.find(id);
		if (__it == __conns_QCompleter.end()) return;
		__c = __it->second;
		__conns_QCompleter.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QCompleter__Go_new(long long id, void* parent) { return new QCompleter__Go(id, (QObject*)parent); }
void QCompleter_delete(void* self) { delete (QCompleter*)self; }

}

#include <QtWidgets/qgesturerecognizer.h>
#include <QtWidgets/qgesture.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include "gen_qgesturerecognizer.h"
#include "gen_purego_cb.h"

class QGestureRecognizer__Go : public QGestureRecognizer {
public:
	long long __goId;
	unsigned char __goOn[3] = {};
	template<typename... A> QGestureRecognizer__Go(long long id, A&&... a) : QGestureRecognizer(std::forward<A>(a)...), __goId(id) {}
	QFlags<QGestureRecognizer::ResultFlag> recognize(QGesture* state, QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp)(__goId, -8783688087013084677LL, (void*)state, (void*)watched, (void*)event, &__h); if (__h) return QFlags<QGestureRecognizer::ResultFlag>(__r); }
		return QFlags<QGestureRecognizer::ResultFlag>{};
	}
	QGesture* create(QObject* target) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -626457835795418115LL, (void*)target, &__h); if (__h) return (QGesture*)__r; }
		return QGestureRecognizer::create(target);
	}
	void reset(QGesture* state) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1086335023529244512LL, (void*)state, &__h); if (__h) return; }
		QGestureRecognizer::reset(state);
	}
};

extern "C" {

void QGestureRecognizer__set_override(void* self, int idx) { ((QGestureRecognizer__Go*)self)->__goOn[idx] = 1; }
void* QGestureRecognizer_create(void* self, void* target) {
	auto&& __r = ((QGestureRecognizer*)self)->create((QObject*)target);
	return (void*)(__r);
}
int QGestureRecognizer_recognize(void* self, void* state, void* watched, void* event) {
	auto&& __r = ((QGestureRecognizer*)self)->recognize((QGesture*)state, (QObject*)watched, (QEvent*)event);
	return (int)(__r);
}
void QGestureRecognizer_reset(void* self, void* state) {
	((QGestureRecognizer*)self)->reset((QGesture*)state);
}
int QGestureRecognizer_registerRecognizer(void* recognizer) {
	auto&& __r = QGestureRecognizer::registerRecognizer((QGestureRecognizer*)recognizer);
	return (int)(__r);
}
void QGestureRecognizer_unregisterRecognizer(int type_) {
	QGestureRecognizer::unregisterRecognizer((Qt::GestureType)(type_));
}
void* QGestureRecognizer_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QGestureRecognizer*)self)->operator=(*(QGestureRecognizer*)arg1);
	return (void*)&(__r);
}
void* QGestureRecognizer_create_default(void* self, void* target) {
	auto&& __r = ((QGestureRecognizer*)self)->QGestureRecognizer::create((QObject*)target);
	return (void*)(__r);
}
void QGestureRecognizer_reset_default(void* self, void* state) {
	((QGestureRecognizer*)self)->QGestureRecognizer::reset((QGesture*)state);
}
void* QGestureRecognizer__Go_new(long long id) { return new QGestureRecognizer__Go(id); }
void QGestureRecognizer_delete(void* self) { delete (QGestureRecognizer*)self; }

}

#include <QtWidgets/qgraphicslinearlayout.h>
#include <QtWidgets/qgraphicslayout.h>
#include <QtWidgets/qgraphicslayoutitem.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qgraphicsitem.h>
#include "gen_qgraphicslinearlayout.h"
#include "gen_purego_cb.h"

class QGraphicsLinearLayout__Go : public QGraphicsLinearLayout {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QGraphicsLinearLayout__Go(long long id, A&&... a) : QGraphicsLinearLayout(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsLinearLayout::addChildLayoutItem;
	using QGraphicsLinearLayout::setGraphicsItem;
	using QGraphicsLinearLayout::setOwnedByLayout;
	void removeAt(int index) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1230346564330554752LL, index, &__h); if (__h) return; }
		QGraphicsLinearLayout::removeAt(index);
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsLinearLayout::setGeometry(rect);
	}
	int count() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QGraphicsLinearLayout::count();
	}
	QGraphicsLayoutItem* itemAt(int index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, index, &__h); if (__h) return (QGraphicsLayoutItem*)__r; }
		return QGraphicsLinearLayout::itemAt(index);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QGraphicsLinearLayout::invalidate();
	}
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QGraphicsLinearLayout::sizeHint(which, constraint);
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsLinearLayout::getContentsMargins(left, top, right, bottom);
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsLinearLayout::updateGeometry();
	}
	void widgetEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 8535746785413991679LL, (void*)e, &__h); if (__h) return; }
		QGraphicsLinearLayout::widgetEvent(e);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsLinearLayout::isEmpty();
	}
};

extern "C" {

void* QGraphicsLinearLayout_base(void* self) { return (void*)static_cast<QGraphicsLayout*>((QGraphicsLinearLayout*)self); }
void QGraphicsLinearLayout_chain(void* self, void** out) {
	QGraphicsLinearLayout* __p = (QGraphicsLinearLayout*)self;
	QGraphicsLayout* __b0 = static_cast<QGraphicsLayout*>(__p); out[0] = (void*)__b0;
	QGraphicsLayoutItem* __b1 = static_cast<QGraphicsLayoutItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsLinearLayout__set_override(void* self, int idx) { ((QGraphicsLinearLayout__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsLinearLayout_new(long long __goid) {
	return new QGraphicsLinearLayout__Go(__goid);
}
void* QGraphicsLinearLayout_new2(long long __goid, int orientation) {
	return new QGraphicsLinearLayout__Go(__goid, (Qt::Orientation)(orientation));
}
void* QGraphicsLinearLayout_new3(long long __goid, void* parent) {
	return new QGraphicsLinearLayout__Go(__goid, (QGraphicsLayoutItem*)parent);
}
void* QGraphicsLinearLayout_new4(long long __goid, int orientation, void* parent) {
	return new QGraphicsLinearLayout__Go(__goid, (Qt::Orientation)(orientation), (QGraphicsLayoutItem*)parent);
}
void QGraphicsLinearLayout_setOrientation(void* self, int orientation) {
	((QGraphicsLinearLayout*)self)->setOrientation((Qt::Orientation)(orientation));
}
int QGraphicsLinearLayout_orientation(void* self) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->orientation();
	return (int)(__r);
}
void QGraphicsLinearLayout_addItem(void* self, void* item) {
	((QGraphicsLinearLayout*)self)->addItem((QGraphicsLayoutItem*)item);
}
void QGraphicsLinearLayout_addStretch(void* self) {
	((QGraphicsLinearLayout*)self)->addStretch();
}
void QGraphicsLinearLayout_insertItem(void* self, int index, void* item) {
	((QGraphicsLinearLayout*)self)->insertItem(index, (QGraphicsLayoutItem*)item);
}
void QGraphicsLinearLayout_insertStretch(void* self, int index) {
	((QGraphicsLinearLayout*)self)->insertStretch(index);
}
void QGraphicsLinearLayout_removeItem(void* self, void* item) {
	((QGraphicsLinearLayout*)self)->removeItem((QGraphicsLayoutItem*)item);
}
void QGraphicsLinearLayout_setSpacing(void* self, double spacing) {
	((QGraphicsLinearLayout*)self)->setSpacing(spacing);
}
double QGraphicsLinearLayout_spacing(void* self) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->spacing();
	return __r;
}
void QGraphicsLinearLayout_setItemSpacing(void* self, int index, double spacing) {
	((QGraphicsLinearLayout*)self)->setItemSpacing(index, spacing);
}
double QGraphicsLinearLayout_itemSpacing(void* self, int index) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->itemSpacing(index);
	return __r;
}
void QGraphicsLinearLayout_setStretchFactor(void* self, void* item, int stretch) {
	((QGraphicsLinearLayout*)self)->setStretchFactor((QGraphicsLayoutItem*)item, stretch);
}
int QGraphicsLinearLayout_stretchFactor(void* self, void* item) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->stretchFactor((QGraphicsLayoutItem*)item);
	return __r;
}
void QGraphicsLinearLayout_setAlignment(void* self, void* item, int alignment) {
	((QGraphicsLinearLayout*)self)->setAlignment((QGraphicsLayoutItem*)item, QFlags<Qt::AlignmentFlag>(alignment));
}
int QGraphicsLinearLayout_alignment(void* self, void* item) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->alignment((QGraphicsLayoutItem*)item);
	return (int)(__r);
}
void QGraphicsLinearLayout_dump(void* self) {
	((QGraphicsLinearLayout*)self)->dump();
}
void QGraphicsLinearLayout_setInstantInvalidatePropagation(int enable) {
	QGraphicsLinearLayout::setInstantInvalidatePropagation(enable);
}
int QGraphicsLinearLayout_instantInvalidatePropagation() {
	auto&& __r = QGraphicsLinearLayout::instantInvalidatePropagation();
	return __r;
}
void QGraphicsLinearLayout_addChildLayoutItem(void* self, void* layoutItem) {
	((QGraphicsLinearLayout__Go*)self)->addChildLayoutItem((QGraphicsLayoutItem*)layoutItem);
}
void QGraphicsLinearLayout_setGraphicsItem(void* self, void* item) {
	((QGraphicsLinearLayout__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsLinearLayout_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsLinearLayout__Go*)self)->setOwnedByLayout(ownedByLayout);
}
void QGraphicsLinearLayout_addStretch2(void* self, int stretch) {
	((QGraphicsLinearLayout*)self)->addStretch(stretch);
}
void QGraphicsLinearLayout_insertStretch2(void* self, int index, int stretch) {
	((QGraphicsLinearLayout*)self)->insertStretch(index, stretch);
}
void QGraphicsLinearLayout_dump2(void* self, int indent) {
	((QGraphicsLinearLayout*)self)->dump(indent);
}
void QGraphicsLinearLayout_removeAt_default(void* self, int index) {
	((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::removeAt(index);
}
void QGraphicsLinearLayout_setGeometry_default(void* self, void* rect) {
	((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::setGeometry(*(QRectF*)rect);
}
int QGraphicsLinearLayout_count_default(void* self) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::count();
	return __r;
}
void* QGraphicsLinearLayout_itemAt_default(void* self, int index) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::itemAt(index);
	return (void*)(__r);
}
void QGraphicsLinearLayout_invalidate_default(void* self) {
	((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::invalidate();
}
void* QGraphicsLinearLayout_sizeHint_default(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::sizeHint((Qt::SizeHint)(which), *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void QGraphicsLinearLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsLinearLayout_updateGeometry_default(void* self) {
	((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::updateGeometry();
}
void QGraphicsLinearLayout_widgetEvent_default(void* self, void* e) {
	((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::widgetEvent((QEvent*)e);
}
int QGraphicsLinearLayout_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsLinearLayout*)self)->QGraphicsLinearLayout::isEmpty();
	return __r;
}
void* QGraphicsLinearLayout__Go_new(long long id) { return new QGraphicsLinearLayout__Go(id); }
void QGraphicsLinearLayout_delete(void* self) { delete (QGraphicsLinearLayout*)self; }

}

#include <QtWidgets/qitemeditorfactory.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qbytearray.h>
#include "gen_qitemeditorfactory.h"
#include "gen_purego_cb.h"

class QItemEditorFactory__Go : public QItemEditorFactory {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QItemEditorFactory__Go(long long id, A&&... a) : QItemEditorFactory(std::forward<A>(a)...), __goId(id) {}
	QWidget* createEditor(int userType, QWidget* parent) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -2118980437705941456LL, userType, (void*)parent, &__h); if (__h) return (QWidget*)__r; }
		return QItemEditorFactory::createEditor(userType, parent);
	}
	QByteArray valuePropertyName(int userType) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -3548500303342436100LL, userType, &__h); if (__h) return *(QByteArray*)__r; }
		return QItemEditorFactory::valuePropertyName(userType);
	}
};

extern "C" {

void QItemEditorFactory__set_override(void* self, int idx) { ((QItemEditorFactory__Go*)self)->__goOn[idx] = 1; }
void* QItemEditorFactory_new(long long __goid) {
	return new QItemEditorFactory__Go(__goid);
}
void* QItemEditorFactory_new2(long long __goid, void* arg1) {
	return new QItemEditorFactory__Go(__goid, *(QItemEditorFactory*)arg1);
}
void* QItemEditorFactory_createEditor(void* self, int userType, void* parent) {
	auto&& __r = ((QItemEditorFactory*)self)->createEditor(userType, (QWidget*)parent);
	return (void*)(__r);
}
void* QItemEditorFactory_valuePropertyName(void* self, int userType) {
	auto&& __r = ((QItemEditorFactory*)self)->valuePropertyName(userType);
	return new QByteArray(__r);
}
void QItemEditorFactory_registerEditor(void* self, int userType, void* creator) {
	((QItemEditorFactory*)self)->registerEditor(userType, (QItemEditorCreatorBase*)creator);
}
void* QItemEditorFactory_defaultFactory() {
	auto&& __r = QItemEditorFactory::defaultFactory();
	return (void*)(__r);
}
void QItemEditorFactory_setDefaultFactory(void* factory) {
	QItemEditorFactory::setDefaultFactory((QItemEditorFactory*)factory);
}
void* QItemEditorFactory_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QItemEditorFactory*)self)->operator=(*(QItemEditorFactory*)arg1);
	return (void*)&(__r);
}
void* QItemEditorFactory_createEditor_default(void* self, int userType, void* parent) {
	auto&& __r = ((QItemEditorFactory*)self)->QItemEditorFactory::createEditor(userType, (QWidget*)parent);
	return (void*)(__r);
}
void* QItemEditorFactory_valuePropertyName_default(void* self, int userType) {
	auto&& __r = ((QItemEditorFactory*)self)->QItemEditorFactory::valuePropertyName(userType);
	return new QByteArray(__r);
}
void* QItemEditorFactory__Go_new(long long id) { return new QItemEditorFactory__Go(id); }
void QItemEditorFactory_delete(void* self) { delete (QItemEditorFactory*)self; }

}

#include <QtWidgets/qscroller.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qscrollerproperties.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qscroller.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QScroller;
static std::mutex __connsMu_QScroller;

extern "C" {

void* QScroller_base(void* self) { return (void*)static_cast<QObject*>((QScroller*)self); }
void* QScroller_fromQObject(void* o) { return (void*)qobject_cast<QScroller*>((QObject*)o); }
char* QScroller_tr(const char* s, long long* __out_len) {
	auto&& __r = QScroller::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QScroller_hasScroller(void* target) {
	auto&& __r = QScroller::hasScroller((QObject*)target);
	return __r;
}
void* QScroller_scroller(void* target) {
	auto&& __r = QScroller::scroller((QObject*)target);
	return (void*)(__r);
}
void* QScroller_scroller2(void* target) {
	auto&& __r = QScroller::scroller((QObject*)target);
	return (void*)(__r);
}
int QScroller_grabGesture(void* target) {
	auto&& __r = QScroller::grabGesture((QObject*)target);
	return (int)(__r);
}
int QScroller_grabbedGesture(void* target) {
	auto&& __r = QScroller::grabbedGesture((QObject*)target);
	return (int)(__r);
}
void QScroller_ungrabGesture(void* target) {
	QScroller::ungrabGesture((QObject*)target);
}
void** QScroller_activeScrollers(long long* __out_len) {
	auto&& __r = QScroller::activeScrollers();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QScroller_target(void* self) {
	auto&& __r = ((QScroller*)self)->target();
	return (void*)(__r);
}
int QScroller_state(void* self) {
	auto&& __r = ((QScroller*)self)->state();
	return (int)(__r);
}
int QScroller_handleInput(void* self, int input, void* position) {
	auto&& __r = ((QScroller*)self)->handleInput((QScroller::Input)(input), *(QPointF*)position);
	return __r;
}
void QScroller_stop(void* self) {
	((QScroller*)self)->stop();
}
void* QScroller_velocity(void* self) {
	auto&& __r = ((QScroller*)self)->velocity();
	return new QPointF(__r);
}
void* QScroller_finalPosition(void* self) {
	auto&& __r = ((QScroller*)self)->finalPosition();
	return new QPointF(__r);
}
void* QScroller_pixelPerMeter(void* self) {
	auto&& __r = ((QScroller*)self)->pixelPerMeter();
	return new QPointF(__r);
}
void* QScroller_scrollerProperties(void* self) {
	auto&& __r = ((QScroller*)self)->scrollerProperties();
	return new QScrollerProperties(__r);
}
void QScroller_setSnapPositionsX(void* self, double* positions, long long positions_len) {
	QList<double> positions_list; positions_list.reserve(positions_len);
	for (long long __i = 0; __i < positions_len; __i++) positions_list.append(positions[__i]);
	((QScroller*)self)->setSnapPositionsX(positions_list);
}
void QScroller_setSnapPositionsX2(void* self, double first, double interval) {
	((QScroller*)self)->setSnapPositionsX(first, interval);
}
void QScroller_setSnapPositionsY(void* self, double* positions, long long positions_len) {
	QList<double> positions_list; positions_list.reserve(positions_len);
	for (long long __i = 0; __i < positions_len; __i++) positions_list.append(positions[__i]);
	((QScroller*)self)->setSnapPositionsY(positions_list);
}
void QScroller_setSnapPositionsY2(void* self, double first, double interval) {
	((QScroller*)self)->setSnapPositionsY(first, interval);
}
void QScroller_setScrollerProperties(void* self, void* prop) {
	((QScroller*)self)->setScrollerProperties(*(QScrollerProperties*)prop);
}
void QScroller_scrollTo(void* self, void* pos) {
	((QScroller*)self)->scrollTo(*(QPointF*)pos);
}
void QScroller_scrollTo2(void* self, void* pos, int scrollTime) {
	((QScroller*)self)->scrollTo(*(QPointF*)pos, scrollTime);
}
void QScroller_ensureVisible(void* self, void* rect, double xmargin, double ymargin) {
	((QScroller*)self)->ensureVisible(*(QRectF*)rect, xmargin, ymargin);
}
void QScroller_ensureVisible2(void* self, void* rect, double xmargin, double ymargin, int scrollTime) {
	((QScroller*)self)->ensureVisible(*(QRectF*)rect, xmargin, ymargin, scrollTime);
}
void QScroller_resendPrepareEvent(void* self) {
	((QScroller*)self)->resendPrepareEvent();
}
void QScroller_stateChanged(void* self, int newstate) {
	((QScroller*)self)->stateChanged((QScroller::State)(newstate));
}
void QScroller_scrollerPropertiesChanged(void* self, void* arg1) {
	((QScroller*)self)->scrollerPropertiesChanged(*(QScrollerProperties*)arg1);
}
void* QScroller_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QScroller::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QScroller_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QScroller::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QScroller_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QScroller::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScroller_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QScroller::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QScroller_grabGesture2(void* target, int gestureType) {
	auto&& __r = QScroller::grabGesture((QObject*)target, (QScroller::ScrollerGestureType)(gestureType));
	return (int)(__r);
}
int QScroller_handleInput2(void* self, int input, void* position, long long timestamp) {
	auto&& __r = ((QScroller*)self)->handleInput((QScroller::Input)(input), *(QPointF*)position, timestamp);
	return __r;
}
void* QScroller_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QScroller::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QScroller_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QScroller::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QScroller_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScroller*)self, &QScroller::stateChanged, (QScroller*)self, [__id](QScroller::State newstate){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(newstate));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScroller);
	__conns_QScroller[__id] = __c;
}
void QScroller_connect_scrollerPropertiesChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScroller*)self, &QScroller::scrollerPropertiesChanged, (QScroller*)self, [__id](const QScrollerProperties& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScroller);
	__conns_QScroller[__id] = __c;
}
void QScroller_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScroller*)self, &QScroller::destroyed, (QScroller*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScroller);
	__conns_QScroller[__id] = __c;
}
void QScroller_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScroller*)self, &QScroller::destroyed, (QScroller*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScroller);
	__conns_QScroller[__id] = __c;
}
void QScroller_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScroller*)self, &QScroller::objectNameChanged, (QScroller*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScroller);
	__conns_QScroller[__id] = __c;
}
void QScroller__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QScroller);
		auto __it = __conns_QScroller.find(id);
		if (__it == __conns_QScroller.end()) return;
		__c = __it->second;
		__conns_QScroller.erase(__it);
	}
	QObject::disconnect(__c);
}

}

#include <QtWidgets/qtreeview.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtWidgets/qheaderview.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qregion.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
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
#include "gen_qtreeview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTreeView__Go : public QTreeView {
public:
	long long __goId;
	unsigned char __goOn[93] = {};
	template<typename... A> QTreeView__Go(long long id, A&&... a) : QTreeView(std::forward<A>(a)...), __goId(id) {}
	using QTreeView::columnResized;
	using QTreeView::columnCountChanged;
	using QTreeView::columnMoved;
	using QTreeView::reexpand;
	using QTreeView::rowsRemoved;
	using QTreeView::drawTree;
	using QTreeView::indexRowSizeHint;
	using QTreeView::rowHeight;
	using QTreeView::state;
	using QTreeView::scheduleDelayedItemsLayout;
	using QTreeView::executeDelayedItemsLayout;
	using QTreeView::setDirtyRegion;
	using QTreeView::scrollDirtyRegion;
	using QTreeView::dirtyRegionOffset;
	using QTreeView::startAutoScroll;
	using QTreeView::stopAutoScroll;
	using QTreeView::doAutoScroll;
	using QTreeView::dropIndicatorPosition;
	using QTreeView::setViewportMargins;
	using QTreeView::viewportMargins;
	using QTreeView::drawFrame;
	using QTreeView::updateMicroFocus;
	using QTreeView::create;
	using QTreeView::destroy;
	using QTreeView::focusNextChild;
	using QTreeView::focusPreviousChild;
	using QTreeView::sender;
	using QTreeView::senderSignalIndex;
	using QTreeView::receivers;
	using QTreeView::isSignalConnected;
	using QTreeView::getDecodedMetricF;
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QTreeView::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QTreeView::scrollContentsBy(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QTreeView::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QTreeView::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QTreeView::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	int cppgoBase_horizontalOffset() { return QTreeView::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QTreeView::verticalOffset(); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QTreeView::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QTreeView::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QTreeView::selectedIndexes(); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTreeView::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTreeView::timerEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTreeView::paintEvent(__a0); }
	void cppgoBase_drawRow(QPainter* __a0, const QStyleOptionViewItem& __a1, const QModelIndex& __a2) { QTreeView::drawRow(__a0, __a1, __a2); }
	void cppgoBase_drawBranches(QPainter* __a0, const QRect& __a1, const QModelIndex& __a2) { QTreeView::drawBranches(__a0, __a1, __a2); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTreeView::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTreeView::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTreeView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTreeView::mouseMoveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTreeView::keyPressEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTreeView::dragMoveEvent(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QTreeView::viewportEvent(__a0); }
	void cppgoBase_updateGeometries() { QTreeView::updateGeometries(); }
	QSize cppgoBase_viewportSizeHint() { return QTreeView::viewportSizeHint(); }
	int cppgoBase_sizeHintForColumn(int __a0) { return QTreeView::sizeHintForColumn(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QTreeView::horizontalScrollbarAction(__a0); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QTreeView::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QTreeView::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QTreeView::currentChanged(__a0, __a1); }
	void cppgoBase_updateEditorData() { QTreeView::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QTreeView::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QTreeView::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QTreeView::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QTreeView::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QTreeView::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QTreeView::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QTreeView::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QTreeView::selectionCommand(__a0, __a1); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QTreeView::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QTreeView::initViewItemOption(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTreeView::focusNextPrevChild(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QTreeView::event(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTreeView::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTreeView::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTreeView::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTreeView::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTreeView::focusOutEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTreeView::resizeEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTreeView::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QTreeView::eventFilter(__a0, __a1); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTreeView::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTreeView::contextMenuEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QTreeView::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTreeView::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTreeView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTreeView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTreeView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTreeView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTreeView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTreeView::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTreeView::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTreeView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTreeView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTreeView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTreeView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTreeView::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTreeView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTreeView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTreeView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTreeView::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTreeView::metaObject();
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QTreeView::setModel(model);
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QTreeView::setRootIndex(index);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QTreeView::setSelectionModel(selectionModel);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[4] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QTreeView::keyboardSearch(search);
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QTreeView::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QTreeView::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTreeView::indexAt(p);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QTreeView::doItemsLayout();
	}
	void reset() override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QTreeView::reset();
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QTreeView::selectAll();
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QTreeView::verticalScrollbarValueChanged(value);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QTreeView::scrollContentsBy(dx, dy);
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTreeView::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTreeView::rowsAboutToBeRemoved(parent, start, end);
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QTreeView::moveCursor(cursorAction, modifiers);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QTreeView::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QTreeView::verticalOffset();
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QTreeView::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QTreeView::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virtp_voidp) {
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
		return QTreeView::selectedIndexes();
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QTreeView::changeEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTreeView::timerEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QTreeView::paintEvent(event);
	}
	void drawRow(QPainter* painter, const QStyleOptionViewItem& options, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -3950604704165452999LL, (void*)painter, (void*)&options, (void*)&index, &__h); if (__h) return; }
		QTreeView::drawRow(painter, options, index);
	}
	void drawBranches(QPainter* painter, const QRect& rect, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6369455307610057755LL, (void*)painter, (void*)&rect, (void*)&index, &__h); if (__h) return; }
		QTreeView::drawBranches(painter, rect, index);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QTreeView::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QTreeView::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QTreeView::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QTreeView::mouseMoveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QTreeView::keyPressEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTreeView::dragMoveEvent(event);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QTreeView::viewportEvent(event);
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QTreeView::updateGeometries();
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QTreeView::viewportSizeHint();
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QTreeView::sizeHintForColumn(column);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QTreeView::horizontalScrollbarAction(action);
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QTreeView::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QTreeView::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QTreeView::currentChanged(current, previous);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QTreeView::sizeHintForRow(row);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QTreeView::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QTreeView::inputMethodQuery(query);
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QTreeView::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QTreeView::updateEditorGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QTreeView::verticalScrollbarAction(action);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QTreeView::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QTreeView::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QTreeView::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QTreeView::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QTreeView::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QTreeView::selectionCommand(index, event);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QTreeView::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QTreeView::initViewItemOption(option);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTreeView::focusNextPrevChild(next);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTreeView::event(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTreeView::dragEnterEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTreeView::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTreeView::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTreeView::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTreeView::focusOutEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QTreeView::resizeEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QTreeView::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QTreeView::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTreeView::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTreeView::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QTreeView::setupViewport(viewport);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QTreeView::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QTreeView::contextMenuEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTreeView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTreeView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTreeView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTreeView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTreeView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTreeView::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTreeView::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTreeView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTreeView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTreeView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTreeView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTreeView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTreeView::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QTreeView::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTreeView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTreeView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTreeView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTreeView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTreeView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTreeView::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTreeView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTreeView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[91] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTreeView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[92] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTreeView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTreeView;
static std::mutex __connsMu_QTreeView;

extern "C" {

void* QTreeView_base(void* self) { return (void*)static_cast<QAbstractItemView*>((QTreeView*)self); }
void QTreeView_chain(void* self, void** out) {
	QTreeView* __p = (QTreeView*)self;
	QAbstractItemView* __b0 = static_cast<QAbstractItemView*>(__p); out[0] = (void*)__b0;
	QAbstractScrollArea* __b1 = static_cast<QAbstractScrollArea*>(__b0); out[1] = (void*)__b1;
	QFrame* __b2 = static_cast<QFrame*>(__b1); out[2] = (void*)__b2;
	QWidget* __b3 = static_cast<QWidget*>(__b2); out[3] = (void*)__b3;
	QObject* __b4 = static_cast<QObject*>(__b3); out[4] = (void*)__b4;
}
void QTreeView__set_override(void* self, int idx) { ((QTreeView__Go*)self)->__goOn[idx] = 1; }
void* QTreeView_fromQObject(void* o) { return (void*)qobject_cast<QTreeView*>((QObject*)o); }
void* QTreeView_new(long long __goid, void* parent) {
	return new QTreeView__Go(__goid, (QWidget*)parent);
}
void* QTreeView_new2(long long __goid) {
	return new QTreeView__Go(__goid);
}
char* QTreeView_tr(const char* s, long long* __out_len) {
	auto&& __r = QTreeView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTreeView_header(void* self) {
	auto&& __r = ((QTreeView*)self)->header();
	return (void*)(__r);
}
void QTreeView_setHeader(void* self, void* header) {
	((QTreeView*)self)->setHeader((QHeaderView*)header);
}
int QTreeView_autoExpandDelay(void* self) {
	auto&& __r = ((QTreeView*)self)->autoExpandDelay();
	return __r;
}
void QTreeView_setAutoExpandDelay(void* self, int delay) {
	((QTreeView*)self)->setAutoExpandDelay(delay);
}
int QTreeView_indentation(void* self) {
	auto&& __r = ((QTreeView*)self)->indentation();
	return __r;
}
void QTreeView_setIndentation(void* self, int i) {
	((QTreeView*)self)->setIndentation(i);
}
void QTreeView_resetIndentation(void* self) {
	((QTreeView*)self)->resetIndentation();
}
int QTreeView_rootIsDecorated(void* self) {
	auto&& __r = ((QTreeView*)self)->rootIsDecorated();
	return __r;
}
void QTreeView_setRootIsDecorated(void* self, int show) {
	((QTreeView*)self)->setRootIsDecorated(show);
}
int QTreeView_uniformRowHeights(void* self) {
	auto&& __r = ((QTreeView*)self)->uniformRowHeights();
	return __r;
}
void QTreeView_setUniformRowHeights(void* self, int uniform) {
	((QTreeView*)self)->setUniformRowHeights(uniform);
}
int QTreeView_itemsExpandable(void* self) {
	auto&& __r = ((QTreeView*)self)->itemsExpandable();
	return __r;
}
void QTreeView_setItemsExpandable(void* self, int enable) {
	((QTreeView*)self)->setItemsExpandable(enable);
}
int QTreeView_expandsOnDoubleClick(void* self) {
	auto&& __r = ((QTreeView*)self)->expandsOnDoubleClick();
	return __r;
}
void QTreeView_setExpandsOnDoubleClick(void* self, int enable) {
	((QTreeView*)self)->setExpandsOnDoubleClick(enable);
}
int QTreeView_columnViewportPosition(void* self, int column) {
	auto&& __r = ((QTreeView*)self)->columnViewportPosition(column);
	return __r;
}
int QTreeView_columnWidth(void* self, int column) {
	auto&& __r = ((QTreeView*)self)->columnWidth(column);
	return __r;
}
void QTreeView_setColumnWidth(void* self, int column, int width) {
	((QTreeView*)self)->setColumnWidth(column, width);
}
int QTreeView_columnAt(void* self, int x) {
	auto&& __r = ((QTreeView*)self)->columnAt(x);
	return __r;
}
int QTreeView_isColumnHidden(void* self, int column) {
	auto&& __r = ((QTreeView*)self)->isColumnHidden(column);
	return __r;
}
void QTreeView_setColumnHidden(void* self, int column, int hide) {
	((QTreeView*)self)->setColumnHidden(column, hide);
}
int QTreeView_isHeaderHidden(void* self) {
	auto&& __r = ((QTreeView*)self)->isHeaderHidden();
	return __r;
}
void QTreeView_setHeaderHidden(void* self, int hide) {
	((QTreeView*)self)->setHeaderHidden(hide);
}
int QTreeView_isRowHidden(void* self, int row, void* parent) {
	auto&& __r = ((QTreeView*)self)->isRowHidden(row, *(QModelIndex*)parent);
	return __r;
}
void QTreeView_setRowHidden(void* self, int row, void* parent, int hide) {
	((QTreeView*)self)->setRowHidden(row, *(QModelIndex*)parent, hide);
}
int QTreeView_isFirstColumnSpanned(void* self, int row, void* parent) {
	auto&& __r = ((QTreeView*)self)->isFirstColumnSpanned(row, *(QModelIndex*)parent);
	return __r;
}
void QTreeView_setFirstColumnSpanned(void* self, int row, void* parent, int span) {
	((QTreeView*)self)->setFirstColumnSpanned(row, *(QModelIndex*)parent, span);
}
int QTreeView_isExpanded(void* self, void* index) {
	auto&& __r = ((QTreeView*)self)->isExpanded(*(QModelIndex*)index);
	return __r;
}
void QTreeView_setExpanded(void* self, void* index, int expand) {
	((QTreeView*)self)->setExpanded(*(QModelIndex*)index, expand);
}
void QTreeView_setSortingEnabled(void* self, int enable) {
	((QTreeView*)self)->setSortingEnabled(enable);
}
int QTreeView_isSortingEnabled(void* self) {
	auto&& __r = ((QTreeView*)self)->isSortingEnabled();
	return __r;
}
void QTreeView_setAnimated(void* self, int enable) {
	((QTreeView*)self)->setAnimated(enable);
}
int QTreeView_isAnimated(void* self) {
	auto&& __r = ((QTreeView*)self)->isAnimated();
	return __r;
}
void QTreeView_setAllColumnsShowFocus(void* self, int enable) {
	((QTreeView*)self)->setAllColumnsShowFocus(enable);
}
int QTreeView_allColumnsShowFocus(void* self) {
	auto&& __r = ((QTreeView*)self)->allColumnsShowFocus();
	return __r;
}
void QTreeView_setWordWrap(void* self, int on) {
	((QTreeView*)self)->setWordWrap(on);
}
int QTreeView_wordWrap(void* self) {
	auto&& __r = ((QTreeView*)self)->wordWrap();
	return __r;
}
void QTreeView_setTreePosition(void* self, int logicalIndex) {
	((QTreeView*)self)->setTreePosition(logicalIndex);
}
int QTreeView_treePosition(void* self) {
	auto&& __r = ((QTreeView*)self)->treePosition();
	return __r;
}
void* QTreeView_indexAbove(void* self, void* index) {
	auto&& __r = ((QTreeView*)self)->indexAbove(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void* QTreeView_indexBelow(void* self, void* index) {
	auto&& __r = ((QTreeView*)self)->indexBelow(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void QTreeView_expanded(void* self, void* index) {
	((QTreeView*)self)->expanded(*(QModelIndex*)index);
}
void QTreeView_collapsed(void* self, void* index) {
	((QTreeView*)self)->collapsed(*(QModelIndex*)index);
}
void QTreeView_hideColumn(void* self, int column) {
	((QTreeView*)self)->hideColumn(column);
}
void QTreeView_showColumn(void* self, int column) {
	((QTreeView*)self)->showColumn(column);
}
void QTreeView_expand(void* self, void* index) {
	((QTreeView*)self)->expand(*(QModelIndex*)index);
}
void QTreeView_collapse(void* self, void* index) {
	((QTreeView*)self)->collapse(*(QModelIndex*)index);
}
void QTreeView_resizeColumnToContents(void* self, int column) {
	((QTreeView*)self)->resizeColumnToContents(column);
}
void QTreeView_sortByColumn(void* self, int column, int order) {
	((QTreeView*)self)->sortByColumn(column, (Qt::SortOrder)(order));
}
void QTreeView_expandAll(void* self) {
	((QTreeView*)self)->expandAll();
}
void QTreeView_expandRecursively(void* self, void* index) {
	((QTreeView*)self)->expandRecursively(*(QModelIndex*)index);
}
void QTreeView_collapseAll(void* self) {
	((QTreeView*)self)->collapseAll();
}
void QTreeView_expandToDepth(void* self, int depth) {
	((QTreeView*)self)->expandToDepth(depth);
}
void QTreeView_columnResized(void* self, int column, int oldSize, int newSize) {
	((QTreeView__Go*)self)->columnResized(column, oldSize, newSize);
}
void QTreeView_columnCountChanged(void* self, int oldCount, int newCount) {
	((QTreeView__Go*)self)->columnCountChanged(oldCount, newCount);
}
void QTreeView_columnMoved(void* self) {
	((QTreeView__Go*)self)->columnMoved();
}
void QTreeView_reexpand(void* self) {
	((QTreeView__Go*)self)->reexpand();
}
void QTreeView_rowsRemoved(void* self, void* parent, int first, int last) {
	((QTreeView__Go*)self)->rowsRemoved(*(QModelIndex*)parent, first, last);
}
void QTreeView_drawTree(void* self, void* painter, void* region) {
	((QTreeView__Go*)self)->drawTree((QPainter*)painter, *(QRegion*)region);
}
int QTreeView_indexRowSizeHint(void* self, void* index) {
	auto&& __r = ((QTreeView__Go*)self)->indexRowSizeHint(*(QModelIndex*)index);
	return __r;
}
int QTreeView_rowHeight(void* self, void* index) {
	auto&& __r = ((QTreeView__Go*)self)->rowHeight(*(QModelIndex*)index);
	return __r;
}
int QTreeView_state(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->state();
	return (int)(__r);
}
void QTreeView_scheduleDelayedItemsLayout(void* self) {
	((QTreeView__Go*)self)->scheduleDelayedItemsLayout();
}
void QTreeView_executeDelayedItemsLayout(void* self) {
	((QTreeView__Go*)self)->executeDelayedItemsLayout();
}
void QTreeView_setDirtyRegion(void* self, void* region) {
	((QTreeView__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QTreeView_scrollDirtyRegion(void* self, int dx, int dy) {
	((QTreeView__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QTreeView_dirtyRegionOffset(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QTreeView_startAutoScroll(void* self) {
	((QTreeView__Go*)self)->startAutoScroll();
}
void QTreeView_stopAutoScroll(void* self) {
	((QTreeView__Go*)self)->stopAutoScroll();
}
void QTreeView_doAutoScroll(void* self) {
	((QTreeView__Go*)self)->doAutoScroll();
}
int QTreeView_dropIndicatorPosition(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QTreeView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QTreeView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QTreeView_setViewportMargins2(void* self, void* margins) {
	((QTreeView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QTreeView_viewportMargins(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QTreeView_drawFrame(void* self, void* arg1) {
	((QTreeView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QTreeView_setTabOrder(void* arg1, void* arg2) {
	QTreeView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTreeView_mouseGrabber() {
	auto&& __r = QTreeView::mouseGrabber();
	return (void*)(__r);
}
void* QTreeView_keyboardGrabber() {
	auto&& __r = QTreeView::keyboardGrabber();
	return (void*)(__r);
}
void* QTreeView_find(unsigned long long arg1) {
	auto&& __r = QTreeView::find(arg1);
	return (void*)(__r);
}
void* QTreeView_createWindowContainer(void* window) {
	auto&& __r = QTreeView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTreeView_updateMicroFocus(void* self) {
	((QTreeView__Go*)self)->updateMicroFocus();
}
void QTreeView_create(void* self) {
	((QTreeView__Go*)self)->create();
}
void QTreeView_destroy(void* self) {
	((QTreeView__Go*)self)->destroy();
}
int QTreeView_focusNextChild(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->focusNextChild();
	return __r;
}
int QTreeView_focusPreviousChild(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTreeView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTreeView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTreeView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTreeView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTreeView_sender(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->sender();
	return (void*)(__r);
}
int QTreeView_senderSignalIndex(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->senderSignalIndex();
	return __r;
}
int QTreeView_receivers(void* self, const char* signal) {
	auto&& __r = ((QTreeView__Go*)self)->receivers(signal);
	return __r;
}
int QTreeView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTreeView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTreeView_devicePixelRatioFScale() {
	auto&& __r = QTreeView::devicePixelRatioFScale();
	return __r;
}
int QTreeView_encodeMetricF(int metric, double value) {
	auto&& __r = QTreeView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTreeView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTreeView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTreeView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTreeView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTreeView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTreeView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTreeView_expandRecursively2(void* self, void* index, int depth) {
	((QTreeView*)self)->expandRecursively(*(QModelIndex*)index, depth);
}
void* QTreeView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTreeView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTreeView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTreeView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTreeView_updateMicroFocus2(void* self, int query) {
	((QTreeView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTreeView_create2(void* self, unsigned long long arg1) {
	((QTreeView__Go*)self)->create(arg1);
}
void QTreeView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTreeView__Go*)self)->create(arg1, initializeWindow);
}
void QTreeView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTreeView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTreeView_destroy2(void* self, int destroyWindow) {
	((QTreeView__Go*)self)->destroy(destroyWindow);
}
void QTreeView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTreeView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTreeView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTreeView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTreeView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTreeView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTreeView_metaObject_default(void* self) {
	auto&& __r = ((QTreeView*)self)->QTreeView::metaObject();
	return (void*)(__r);
}
void QTreeView_setModel_default(void* self, void* model) {
	((QTreeView*)self)->QTreeView::setModel((QAbstractItemModel*)model);
}
void QTreeView_setRootIndex_default(void* self, void* index) {
	((QTreeView*)self)->QTreeView::setRootIndex(*(QModelIndex*)index);
}
void QTreeView_setSelectionModel_default(void* self, void* selectionModel) {
	((QTreeView*)self)->QTreeView::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QTreeView_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QTreeView*)self)->QTreeView::keyboardSearch(search_qstr);
}
void* QTreeView_visualRect_default(void* self, void* index) {
	auto&& __r = ((QTreeView*)self)->QTreeView::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QTreeView_scrollTo_default(void* self, void* index, int hint) {
	((QTreeView*)self)->QTreeView::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QTreeView_indexAt_default(void* self, void* p) {
	auto&& __r = ((QTreeView*)self)->QTreeView::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QTreeView_doItemsLayout_default(void* self) {
	((QTreeView*)self)->QTreeView::doItemsLayout();
}
void QTreeView_reset_default(void* self) {
	((QTreeView*)self)->QTreeView::reset();
}
void QTreeView_selectAll_default(void* self) {
	((QTreeView*)self)->QTreeView::selectAll();
}
void QTreeView_verticalScrollbarValueChanged_default(void* self, int value) {
	((QTreeView__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QTreeView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QTreeView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QTreeView_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QTreeView__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QTreeView_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QTreeView__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void* QTreeView_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
int QTreeView_horizontalOffset_default(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QTreeView_verticalOffset_default(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void QTreeView_setSelection_default(void* self, void* rect, int command) {
	((QTreeView__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QTreeView_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QTreeView_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QTreeView_changeEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QTreeView_timerEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTreeView_paintEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QTreeView_drawRow_default(void* self, void* painter, void* options, void* index) {
	((QTreeView__Go*)self)->cppgoBase_drawRow((QPainter*)painter, *(QStyleOptionViewItem*)options, *(QModelIndex*)index);
}
void QTreeView_drawBranches_default(void* self, void* painter, void* rect, void* index) {
	((QTreeView__Go*)self)->cppgoBase_drawBranches((QPainter*)painter, *(QRect*)rect, *(QModelIndex*)index);
}
void QTreeView_mousePressEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QTreeView_mouseReleaseEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QTreeView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QTreeView_mouseMoveEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QTreeView_keyPressEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QTreeView_dragMoveEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
int QTreeView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QTreeView_updateGeometries_default(void* self) {
	((QTreeView__Go*)self)->cppgoBase_updateGeometries();
}
void* QTreeView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
int QTreeView_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_sizeHintForColumn(column);
	return __r;
}
void QTreeView_horizontalScrollbarAction_default(void* self, int action) {
	((QTreeView__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
int QTreeView_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QTreeView_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QTreeView__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QTreeView_currentChanged_default(void* self, void* current, void* previous) {
	((QTreeView__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
int QTreeView_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QTreeView*)self)->QTreeView::sizeHintForRow(row);
	return __r;
}
void* QTreeView_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QTreeView*)self)->QTreeView::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QTreeView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QTreeView*)self)->QTreeView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QTreeView_updateEditorData_default(void* self) {
	((QTreeView__Go*)self)->cppgoBase_updateEditorData();
}
void QTreeView_updateEditorGeometries_default(void* self) {
	((QTreeView__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QTreeView_verticalScrollbarAction_default(void* self, int action) {
	((QTreeView__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QTreeView_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QTreeView__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QTreeView_closeEditor_default(void* self, void* editor, int hint) {
	((QTreeView__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QTreeView_commitData_default(void* self, void* editor) {
	((QTreeView__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QTreeView_editorDestroyed_default(void* self, void* editor) {
	((QTreeView__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QTreeView_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QTreeView_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
void QTreeView_startDrag_default(void* self, int supportedActions) {
	((QTreeView__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QTreeView_initViewItemOption_default(void* self, void* option) {
	((QTreeView__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
int QTreeView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QTreeView_event_default(void* self, void* event) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QTreeView_dragEnterEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTreeView_dragLeaveEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTreeView_dropEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QTreeView_focusInEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QTreeView_focusOutEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTreeView_resizeEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QTreeView_inputMethodEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QTreeView_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QTreeView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTreeView*)self)->QTreeView::minimumSizeHint();
	return new QSize(__r);
}
void* QTreeView_sizeHint_default(void* self) {
	auto&& __r = ((QTreeView*)self)->QTreeView::sizeHint();
	return new QSize(__r);
}
void QTreeView_setupViewport_default(void* self, void* viewport) {
	((QTreeView*)self)->QTreeView::setupViewport((QWidget*)viewport);
}
void QTreeView_wheelEvent_default(void* self, void* arg1) {
	((QTreeView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QTreeView_contextMenuEvent_default(void* self, void* arg1) {
	((QTreeView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QTreeView_initStyleOption_default(void* self, void* option) {
	((QTreeView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QTreeView_devType_default(void* self) {
	auto&& __r = ((QTreeView*)self)->QTreeView::devType();
	return __r;
}
void QTreeView_setVisible_default(void* self, int visible) {
	((QTreeView*)self)->QTreeView::setVisible(visible);
}
int QTreeView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTreeView*)self)->QTreeView::heightForWidth(arg1);
	return __r;
}
int QTreeView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTreeView*)self)->QTreeView::hasHeightForWidth();
	return __r;
}
void* QTreeView_paintEngine_default(void* self) {
	auto&& __r = ((QTreeView*)self)->QTreeView::paintEngine();
	return (void*)(__r);
}
void QTreeView_keyReleaseEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTreeView_enterEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTreeView_leaveEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTreeView_moveEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTreeView_closeEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTreeView_tabletEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTreeView_actionEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTreeView_showEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QTreeView_hideEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTreeView_metric_default(void* self, int arg1) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTreeView_initPainter_default(void* self, void* painter) {
	((QTreeView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTreeView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTreeView_sharedPainter_default(void* self) {
	auto&& __r = ((QTreeView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTreeView_childEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTreeView_customEvent_default(void* self, void* event) {
	((QTreeView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTreeView_connectNotify_default(void* self, void* signal) {
	((QTreeView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTreeView_disconnectNotify_default(void* self, void* signal) {
	((QTreeView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTreeView_connect_expanded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::expanded, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_collapsed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::collapsed, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::pressed, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::clicked, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::doubleClicked, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::activated, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::entered, (QTreeView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::viewportEntered, (QTreeView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::iconSizeChanged, (QTreeView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::windowTitleChanged, (QTreeView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::windowIconChanged, (QTreeView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::windowIconTextChanged, (QTreeView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::customContextMenuRequested, (QTreeView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::destroyed, (QTreeView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::destroyed, (QTreeView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeView*)self, &QTreeView::objectNameChanged, (QTreeView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
	__conns_QTreeView[__id] = __c;
}
void QTreeView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTreeView);
		auto __it = __conns_QTreeView.find(id);
		if (__it == __conns_QTreeView.end()) return;
		__c = __it->second;
		__conns_QTreeView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTreeView__Go_new(long long id) { return new QTreeView__Go(id); }
void QTreeView_delete(void* self) { delete (QTreeView*)self; }

}

