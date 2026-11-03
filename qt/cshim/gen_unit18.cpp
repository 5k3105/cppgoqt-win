#include <QtCore/qobjectdefs.h>
#include "gen_qmetaobject__connection.h"

extern "C" {

void* QMetaObject__Connection_new() {
	return new QMetaObject::Connection();
}
void* QMetaObject__Connection_new2(void* other) {
	return new QMetaObject::Connection(*(QMetaObject::Connection*)other);
}
void* QMetaObject__Connection_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QMetaObject::Connection*)self)->operator=(*(QMetaObject::Connection*)other);
	return (void*)&(__r);
}
void QMetaObject__Connection_swap(void* self, void* other) {
	((QMetaObject::Connection*)self)->swap(*(QMetaObject::Connection*)other);
}
void QMetaObject__Connection_delete(void* self) { delete (QMetaObject::Connection*)self; }

}

#include <QtCore/qmetacontainer.h>
#include <QtCore/qmetatype.h>
#include "gen_qmetaassociation.h"

extern "C" {

void* QMetaAssociation_base(void* self) { return (void*)static_cast<QMetaContainer*>((QMetaAssociation*)self); }
void* QMetaAssociation_new() {
	return new QMetaAssociation();
}
void* QMetaAssociation_new2(void* arg1) {
	return new QMetaAssociation(*(QMetaAssociation*)arg1);
}
void* QMetaAssociation_keyMetaType(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->keyMetaType();
	return new QMetaType(__r);
}
void* QMetaAssociation_mappedMetaType(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->mappedMetaType();
	return new QMetaType(__r);
}
int QMetaAssociation_canInsertKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canInsertKey();
	return __r;
}
int QMetaAssociation_canRemoveKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canRemoveKey();
	return __r;
}
int QMetaAssociation_canContainsKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canContainsKey();
	return __r;
}
int QMetaAssociation_canGetMappedAtKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canGetMappedAtKey();
	return __r;
}
int QMetaAssociation_canSetMappedAtKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canSetMappedAtKey();
	return __r;
}
int QMetaAssociation_canGetKeyAtIterator(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canGetKeyAtIterator();
	return __r;
}
int QMetaAssociation_canGetKeyAtConstIterator(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canGetKeyAtConstIterator();
	return __r;
}
int QMetaAssociation_canGetMappedAtIterator(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canGetMappedAtIterator();
	return __r;
}
int QMetaAssociation_canGetMappedAtConstIterator(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canGetMappedAtConstIterator();
	return __r;
}
int QMetaAssociation_canSetMappedAtIterator(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canSetMappedAtIterator();
	return __r;
}
int QMetaAssociation_canCreateIteratorAtKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canCreateIteratorAtKey();
	return __r;
}
int QMetaAssociation_canCreateConstIteratorAtKey(void* self) {
	auto&& __r = ((QMetaAssociation*)self)->canCreateConstIteratorAtKey();
	return __r;
}
int QMetaAssociation_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaAssociation*)self) == (*(QMetaAssociation*)rhs);
	return __r;
}
int QMetaAssociation_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaAssociation*)self) != (*(QMetaAssociation*)rhs);
	return __r;
}
void QMetaAssociation_delete(void* self) { delete (QMetaAssociation*)self; }

}

#include <QtCore/qdebug.h>
#include "gen_qdebugstatesaver.h"

extern "C" {

void* QDebugStateSaver_new(void* dbg) {
	return new QDebugStateSaver(*(QDebug*)dbg);
}
void QDebugStateSaver_delete(void* self) { delete (QDebugStateSaver*)self; }

}

#include <QtCore/qassociativeiterable.h>
#include <QtCore/qvariant.h>
#include <QtCore/qmetacontainer.h>
#include "gen_qassociativeconstiterator.h"

extern "C" {

void* QAssociativeConstIterator_new(void* arg1) {
	return new QAssociativeConstIterator(*(QAssociativeConstIterator*)arg1);
}
void* QAssociativeConstIterator_key(void* self) {
	auto&& __r = ((QAssociativeConstIterator*)self)->key();
	return new QVariant(__r);
}
void* QAssociativeConstIterator_value(void* self) {
	auto&& __r = ((QAssociativeConstIterator*)self)->value();
	return new QVariant(__r);
}
void* QAssociativeConstIterator_OperatorMultiply(void* self) {
	auto&& __r = ((QAssociativeConstIterator*)self)->operator*();
	return new QVariant(__r);
}
void* QAssociativeConstIterator_OperatorMinusGreater(void* self) {
	auto&& __r = ((QAssociativeConstIterator*)self)->operator->();
	return new QVariantConstPointer(__r);
}
void* QAssociativeConstIterator_metaContainer(void* self) {
	auto&& __r = ((QAssociativeConstIterator*)self)->metaContainer();
	return new QMetaAssociation(__r);
}
void QAssociativeConstIterator_delete(void* self) { delete (QAssociativeConstIterator*)self; }

}

#include <QtCore/qcborvalue.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qurl.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/quuid.h>
#include <QtCore/qcborarray.h>
#include <QtCore/qcbormap.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qvariant.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qcborstreamwriter.h>
#include <QtCore/qcompare.h>
#include "gen_qcborvalueconstref.h"

extern "C" {

void* QCborValueConstRef_new(void* arg1) {
	return new QCborValueConstRef(*(QCborValueConstRef*)arg1);
}
void* QCborValueConstRef_ToQCborValue(void* self) {
	auto&& __r = static_cast<QCborValue>(*(QCborValueConstRef*)self);
	return new QCborValue(__r);
}
int QCborValueConstRef_type(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->type();
	return (int)(__r);
}
int QCborValueConstRef_isInteger(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isInteger();
	return __r;
}
int QCborValueConstRef_isByteArray(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isByteArray();
	return __r;
}
int QCborValueConstRef_isString(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isString();
	return __r;
}
int QCborValueConstRef_isArray(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isArray();
	return __r;
}
int QCborValueConstRef_isMap(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isMap();
	return __r;
}
int QCborValueConstRef_isTag(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isTag();
	return __r;
}
int QCborValueConstRef_isFalse(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isFalse();
	return __r;
}
int QCborValueConstRef_isTrue(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isTrue();
	return __r;
}
int QCborValueConstRef_isBool(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isBool();
	return __r;
}
int QCborValueConstRef_isNull(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isNull();
	return __r;
}
int QCborValueConstRef_isUndefined(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isUndefined();
	return __r;
}
int QCborValueConstRef_isDouble(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isDouble();
	return __r;
}
int QCborValueConstRef_isDateTime(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isDateTime();
	return __r;
}
int QCborValueConstRef_isUrl(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isUrl();
	return __r;
}
int QCborValueConstRef_isRegularExpression(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isRegularExpression();
	return __r;
}
int QCborValueConstRef_isUuid(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isUuid();
	return __r;
}
int QCborValueConstRef_isInvalid(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isInvalid();
	return __r;
}
int QCborValueConstRef_isContainer(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isContainer();
	return __r;
}
int QCborValueConstRef_isSimpleType(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->isSimpleType();
	return __r;
}
int QCborValueConstRef_isSimpleType2(void* self, int st) {
	auto&& __r = ((QCborValueConstRef*)self)->isSimpleType((QCborSimpleType)(st));
	return __r;
}
int QCborValueConstRef_toSimpleType(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toSimpleType();
	return (int)(__r);
}
int QCborValueConstRef_tag(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->tag();
	return (int)(__r);
}
void* QCborValueConstRef_taggedValue(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->taggedValue();
	return new QCborValue(__r);
}
long long QCborValueConstRef_toInteger(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toInteger();
	return __r;
}
int QCborValueConstRef_toBool(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toBool();
	return __r;
}
double QCborValueConstRef_toDouble(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toDouble();
	return __r;
}
void* QCborValueConstRef_toByteArray(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toByteArray();
	return new QByteArray(__r);
}
char* QCborValueConstRef_toString(void* self, long long* __out_len) {
	auto&& __r = ((QCborValueConstRef*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborValueConstRef_toStringView(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toStringView();
	return new QAnyStringView(__r);
}
void* QCborValueConstRef_toDateTime(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toDateTime();
	return new QDateTime(__r);
}
void* QCborValueConstRef_toUrl(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toUrl();
	return new QUrl(__r);
}
void* QCborValueConstRef_toRegularExpression(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toRegularExpression();
	return new QRegularExpression(__r);
}
void* QCborValueConstRef_toUuid(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toUuid();
	return new QUuid(__r);
}
void* QCborValueConstRef_toArray(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toArray();
	return new QCborArray(__r);
}
void* QCborValueConstRef_toArray2(void* self, void* a) {
	auto&& __r = ((QCborValueConstRef*)self)->toArray(*(QCborArray*)a);
	return new QCborArray(__r);
}
void* QCborValueConstRef_toMap(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toMap();
	return new QCborMap(__r);
}
void* QCborValueConstRef_toMap2(void* self, void* m) {
	auto&& __r = ((QCborValueConstRef*)self)->toMap(*(QCborMap*)m);
	return new QCborMap(__r);
}
void* QCborValueConstRef_OperatorSubscript(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborValueConstRef*)self)->operator[](key_qstr);
	return new QCborValue(__r);
}
void* QCborValueConstRef_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QCborValueConstRef*)self)->operator[](*(QLatin1String*)key);
	return new QCborValue(__r);
}
void* QCborValueConstRef_OperatorSubscript3(void* self, long long key) {
	auto&& __r = ((QCborValueConstRef*)self)->operator[](key);
	return new QCborValue(__r);
}
int QCborValueConstRef_compare(void* self, void* other) {
	auto&& __r = ((QCborValueConstRef*)self)->compare(*(QCborValue*)other);
	return __r;
}
void* QCborValueConstRef_toVariant(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toVariant();
	return new QVariant(__r);
}
void* QCborValueConstRef_toJsonValue(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toJsonValue();
	return new QJsonValue(__r);
}
void* QCborValueConstRef_toCbor(void* self) {
	auto&& __r = ((QCborValueConstRef*)self)->toCbor();
	return new QByteArray(__r);
}
void QCborValueConstRef_toCbor2(void* self, void* writer) {
	((QCborValueConstRef*)self)->toCbor(*(QCborStreamWriter*)writer);
}
char* QCborValueConstRef_toDiagnosticNotation(void* self, long long* __out_len) {
	auto&& __r = ((QCborValueConstRef*)self)->toDiagnosticNotation();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCborValueConstRef_toSimpleType2(void* self, int defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toSimpleType((QCborSimpleType)(defaultValue));
	return (int)(__r);
}
int QCborValueConstRef_tag2(void* self, int defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->tag((QCborTag)(defaultValue));
	return (int)(__r);
}
void* QCborValueConstRef_taggedValue2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->taggedValue(*(QCborValue*)defaultValue);
	return new QCborValue(__r);
}
long long QCborValueConstRef_toInteger2(void* self, long long defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toInteger(defaultValue);
	return __r;
}
int QCborValueConstRef_toBool2(void* self, int defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toBool(defaultValue);
	return __r;
}
double QCborValueConstRef_toDouble2(void* self, double defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toDouble(defaultValue);
	return __r;
}
void* QCborValueConstRef_toByteArray2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toByteArray(*(QByteArray*)defaultValue);
	return new QByteArray(__r);
}
char* QCborValueConstRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QCborValueConstRef*)self)->toString(defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborValueConstRef_toStringView2(void* self, const char* defaultValue, long long defaultValue_len) {
	QAnyStringView defaultValue_qstr = QAnyStringView(defaultValue, defaultValue_len);
	auto&& __r = ((QCborValueConstRef*)self)->toStringView(defaultValue_qstr);
	return new QAnyStringView(__r);
}
void* QCborValueConstRef_toDateTime2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toDateTime(*(QDateTime*)defaultValue);
	return new QDateTime(__r);
}
void* QCborValueConstRef_toUrl2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toUrl(*(QUrl*)defaultValue);
	return new QUrl(__r);
}
void* QCborValueConstRef_toRegularExpression2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toRegularExpression(*(QRegularExpression*)defaultValue);
	return new QRegularExpression(__r);
}
void* QCborValueConstRef_toUuid2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueConstRef*)self)->toUuid(*(QUuid*)defaultValue);
	return new QUuid(__r);
}
void* QCborValueConstRef_toCbor3(void* self, int opt) {
	auto&& __r = ((QCborValueConstRef*)self)->toCbor(QFlags<QCborValue::EncodingOption>(opt));
	return new QByteArray(__r);
}
void QCborValueConstRef_toCbor4(void* self, void* writer, int opt) {
	((QCborValueConstRef*)self)->toCbor(*(QCborStreamWriter*)writer, QFlags<QCborValue::EncodingOption>(opt));
}
char* QCborValueConstRef_toDiagnosticNotation2(void* self, int opt, long long* __out_len) {
	auto&& __r = ((QCborValueConstRef*)self)->toDiagnosticNotation(QFlags<QCborValue::DiagnosticNotationOption>(opt));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCborValueConstRef_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) == (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) != (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) < (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) > (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) <= (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) >= (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) == (*(QCborValue*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) != (*(QCborValue*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) < (*(QCborValue*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) > (*(QCborValue*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) <= (*(QCborValue*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) >= (*(QCborValue*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) == (*(QCborArray*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) != (*(QCborArray*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) < (*(QCborArray*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) > (*(QCborArray*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) <= (*(QCborArray*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) >= (*(QCborArray*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) == (*(QCborMap*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) != (*(QCborMap*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesser4(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) < (*(QCborMap*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreater4(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) > (*(QCborMap*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorLesserOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) <= (*(QCborMap*)rhs);
	return __r;
}
int QCborValueConstRef_OperatorGreaterOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValueConstRef*)self) >= (*(QCborMap*)rhs);
	return __r;
}
void QCborValueConstRef_delete(void* self) { delete (QCborValueConstRef*)self; }

}

#include <QtCore/qfiledevice.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qfiledevice.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFileDevice;
static std::mutex __connsMu_QFileDevice;

extern "C" {

void* QFileDevice_base(void* self) { return (void*)static_cast<QIODevice*>((QFileDevice*)self); }
void QFileDevice_chain(void* self, void** out) {
	QFileDevice* __p = (QFileDevice*)self;
	QIODevice* __b0 = static_cast<QIODevice*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QFileDevice_fromQObject(void* o) { return (void*)qobject_cast<QFileDevice*>((QObject*)o); }
char* QFileDevice_tr(const char* s, long long* __out_len) {
	auto&& __r = QFileDevice::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFileDevice_error(void* self) {
	auto&& __r = ((QFileDevice*)self)->error();
	return (int)(__r);
}
void QFileDevice_unsetError(void* self) {
	((QFileDevice*)self)->unsetError();
}
int QFileDevice_handle(void* self) {
	auto&& __r = ((QFileDevice*)self)->handle();
	return __r;
}
char* QFileDevice_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QFileDevice*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFileDevice_flush(void* self) {
	auto&& __r = ((QFileDevice*)self)->flush();
	return __r;
}
int QFileDevice_resize(void* self, long long sz) {
	auto&& __r = ((QFileDevice*)self)->resize(sz);
	return __r;
}
int QFileDevice_permissions(void* self) {
	auto&& __r = ((QFileDevice*)self)->permissions();
	return (int)(__r);
}
int QFileDevice_setPermissions(void* self, int permissionSpec) {
	auto&& __r = ((QFileDevice*)self)->setPermissions(QFlags<QFileDevice::Permission>(permissionSpec));
	return __r;
}
int QFileDevice_unmap(void* self, unsigned char* address) {
	uchar address_local{};
	auto&& __r = ((QFileDevice*)self)->unmap(&address_local);
	*address = address_local;
	return __r;
}
void* QFileDevice_fileTime(void* self, int time) {
	auto&& __r = ((QFileDevice*)self)->fileTime((QFileDevice::FileTime)(time));
	return new QDateTime(__r);
}
int QFileDevice_setFileTime(void* self, void* newDate, int fileTime) {
	auto&& __r = ((QFileDevice*)self)->setFileTime(*(QDateTime*)newDate, (QFileDevice::FileTime)(fileTime));
	return __r;
}
void* QFileDevice_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFileDevice::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFileDevice_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFileDevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QFileDevice_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFileDevice::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDevice_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFileDevice::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFileDevice::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFileDevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFileDevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QFileDevice_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::readyRead, (QFileDevice*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::channelReadyRead, (QFileDevice*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::bytesWritten, (QFileDevice*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::channelBytesWritten, (QFileDevice*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::aboutToClose, (QFileDevice*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::readChannelFinished, (QFileDevice*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::destroyed, (QFileDevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::destroyed, (QFileDevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDevice*)self, &QFileDevice::objectNameChanged, (QFileDevice*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
	__conns_QFileDevice[__id] = __c;
}
void QFileDevice__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFileDevice);
		auto __it = __conns_QFileDevice.find(id);
		if (__it == __conns_QFileDevice.end()) return;
		__c = __it->second;
		__conns_QFileDevice.erase(__it);
	}
	QObject::disconnect(__c);
}
void QFileDevice_delete(void* self) { delete (QFileDevice*)self; }

}

#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qabstractitemmodel.h>
#include "gen_qitemselectionrange.h"

extern "C" {

void* QItemSelectionRange_new() {
	return new QItemSelectionRange();
}
void* QItemSelectionRange_new2(void* topL, void* bottomR) {
	return new QItemSelectionRange(*(QModelIndex*)topL, *(QModelIndex*)bottomR);
}
void* QItemSelectionRange_new3(void* index) {
	return new QItemSelectionRange(*(QModelIndex*)index);
}
void* QItemSelectionRange_new4(void* arg1) {
	return new QItemSelectionRange(*(QItemSelectionRange*)arg1);
}
void QItemSelectionRange_swap(void* self, void* other) {
	((QItemSelectionRange*)self)->swap(*(QItemSelectionRange*)other);
}
int QItemSelectionRange_top(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->top();
	return __r;
}
int QItemSelectionRange_left(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->left();
	return __r;
}
int QItemSelectionRange_bottom(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->bottom();
	return __r;
}
int QItemSelectionRange_right(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->right();
	return __r;
}
int QItemSelectionRange_width(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->width();
	return __r;
}
int QItemSelectionRange_height(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->height();
	return __r;
}
void* QItemSelectionRange_topLeft(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->topLeft();
	return (void*)&(__r);
}
void* QItemSelectionRange_bottomRight(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->bottomRight();
	return (void*)&(__r);
}
void* QItemSelectionRange_parent(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->parent();
	return new QModelIndex(__r);
}
void* QItemSelectionRange_model(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->model();
	return (void*)(__r);
}
int QItemSelectionRange_contains(void* self, void* index) {
	auto&& __r = ((QItemSelectionRange*)self)->contains(*(QModelIndex*)index);
	return __r;
}
int QItemSelectionRange_contains2(void* self, int row, int column, void* parentIndex) {
	auto&& __r = ((QItemSelectionRange*)self)->contains(row, column, *(QModelIndex*)parentIndex);
	return __r;
}
int QItemSelectionRange_intersects(void* self, void* other) {
	auto&& __r = ((QItemSelectionRange*)self)->intersects(*(QItemSelectionRange*)other);
	return __r;
}
void* QItemSelectionRange_intersected(void* self, void* other) {
	auto&& __r = ((QItemSelectionRange*)self)->intersected(*(QItemSelectionRange*)other);
	return new QItemSelectionRange(__r);
}
int QItemSelectionRange_isValid(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->isValid();
	return __r;
}
int QItemSelectionRange_isEmpty(void* self) {
	auto&& __r = ((QItemSelectionRange*)self)->isEmpty();
	return __r;
}
void** QItemSelectionRange_indexes(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelectionRange*)self)->indexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QItemSelectionRange_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QItemSelectionRange*)self)->operator=(*(QItemSelectionRange*)arg1);
	return (void*)&(__r);
}
int QItemSelectionRange_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QItemSelectionRange*)self) == (*(QItemSelectionRange*)rhs);
	return __r;
}
int QItemSelectionRange_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QItemSelectionRange*)self) != (*(QItemSelectionRange*)rhs);
	return __r;
}
void QItemSelectionRange_delete(void* self) { delete (QItemSelectionRange*)self; }

}

#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qobject.h>
#include <QtCore/qnamespace.h>
#include "gen_qmetamethod.h"

extern "C" {

void* QMetaMethod_new() {
	return new QMetaMethod();
}
void* QMetaMethod_new2(void* arg1) {
	return new QMetaMethod(*(QMetaMethod*)arg1);
}
void* QMetaMethod_methodSignature(void* self) {
	auto&& __r = ((QMetaMethod*)self)->methodSignature();
	return new QByteArray(__r);
}
void* QMetaMethod_name(void* self) {
	auto&& __r = ((QMetaMethod*)self)->name();
	return new QByteArray(__r);
}
void* QMetaMethod_nameView(void* self) {
	auto&& __r = ((QMetaMethod*)self)->nameView();
	return new QByteArrayView(__r);
}
const char* QMetaMethod_typeName(void* self) {
	auto&& __r = ((QMetaMethod*)self)->typeName();
	return __r;
}
int QMetaMethod_returnType(void* self) {
	auto&& __r = ((QMetaMethod*)self)->returnType();
	return __r;
}
void* QMetaMethod_returnMetaType(void* self) {
	auto&& __r = ((QMetaMethod*)self)->returnMetaType();
	return new QMetaType(__r);
}
int QMetaMethod_parameterCount(void* self) {
	auto&& __r = ((QMetaMethod*)self)->parameterCount();
	return __r;
}
int QMetaMethod_parameterType(void* self, int index) {
	auto&& __r = ((QMetaMethod*)self)->parameterType(index);
	return __r;
}
void* QMetaMethod_parameterMetaType(void* self, int index) {
	auto&& __r = ((QMetaMethod*)self)->parameterMetaType(index);
	return new QMetaType(__r);
}
void QMetaMethod_getParameterTypes(void* self, int* types) {
	int types_local{};
	((QMetaMethod*)self)->getParameterTypes(&types_local);
	*types = types_local;
}
void** QMetaMethod_parameterTypes(void* self, long long* __out_len) {
	auto&& __r = ((QMetaMethod*)self)->parameterTypes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void* QMetaMethod_parameterTypeName(void* self, int index) {
	auto&& __r = ((QMetaMethod*)self)->parameterTypeName(index);
	return new QByteArray(__r);
}
void** QMetaMethod_parameterNames(void* self, long long* __out_len) {
	auto&& __r = ((QMetaMethod*)self)->parameterNames();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
const char* QMetaMethod_tag(void* self) {
	auto&& __r = ((QMetaMethod*)self)->tag();
	return __r;
}
int QMetaMethod_access(void* self) {
	auto&& __r = ((QMetaMethod*)self)->access();
	return (int)(__r);
}
int QMetaMethod_methodType(void* self) {
	auto&& __r = ((QMetaMethod*)self)->methodType();
	return (int)(__r);
}
int QMetaMethod_attributes(void* self) {
	auto&& __r = ((QMetaMethod*)self)->attributes();
	return __r;
}
int QMetaMethod_methodIndex(void* self) {
	auto&& __r = ((QMetaMethod*)self)->methodIndex();
	return __r;
}
int QMetaMethod_relativeMethodIndex(void* self) {
	auto&& __r = ((QMetaMethod*)self)->relativeMethodIndex();
	return __r;
}
int QMetaMethod_revision(void* self) {
	auto&& __r = ((QMetaMethod*)self)->revision();
	return __r;
}
int QMetaMethod_isConst(void* self) {
	auto&& __r = ((QMetaMethod*)self)->isConst();
	return __r;
}
void* QMetaMethod_enclosingMetaObject(void* self) {
	auto&& __r = ((QMetaMethod*)self)->enclosingMetaObject();
	return (void*)(__r);
}
int QMetaMethod_invoke(void* self, void* object, int connectionType, void* returnValue) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue);
	return __r;
}
int QMetaMethod_invoke2(void* self, void* object, void* returnValue) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue);
	return __r;
}
int QMetaMethod_invoke3(void* self, void* object, int connectionType, void* val0) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0);
	return __r;
}
int QMetaMethod_invoke4(void* self, void* object, void* val0) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0);
	return __r;
}
int QMetaMethod_isValid(void* self) {
	auto&& __r = ((QMetaMethod*)self)->isValid();
	return __r;
}
int QMetaMethod_invoke5(void* self, void* object, int connectionType, void* returnValue, void* val0) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0);
	return __r;
}
int QMetaMethod_invoke6(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaMethod_invoke7(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaMethod_invoke8(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaMethod_invoke9(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaMethod_invoke10(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaMethod_invoke11(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaMethod_invoke12(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaMethod_invoke13(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaMethod_invoke14(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaMethod_invoke15(void* self, void* object, void* returnValue, void* val0) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0);
	return __r;
}
int QMetaMethod_invoke16(void* self, void* object, void* returnValue, void* val0, void* val1) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaMethod_invoke17(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaMethod_invoke18(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaMethod_invoke19(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaMethod_invoke20(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaMethod_invoke21(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaMethod_invoke22(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaMethod_invoke23(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaMethod_invoke24(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericReturnArgument*)returnValue, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaMethod_invoke25(void* self, void* object, int connectionType, void* val0, void* val1) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaMethod_invoke26(void* self, void* object, int connectionType, void* val0, void* val1, void* val2) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaMethod_invoke27(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaMethod_invoke28(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaMethod_invoke29(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaMethod_invoke30(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaMethod_invoke31(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaMethod_invoke32(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaMethod_invoke33(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, (Qt::ConnectionType)(connectionType), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaMethod_invoke34(void* self, void* object, void* val0, void* val1) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaMethod_invoke35(void* self, void* object, void* val0, void* val1, void* val2) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaMethod_invoke36(void* self, void* object, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaMethod_invoke37(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaMethod_invoke38(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaMethod_invoke39(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaMethod_invoke40(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaMethod_invoke41(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaMethod_invoke42(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = ((QMetaMethod*)self)->invoke((QObject*)object, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaMethod_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaMethod*)self) == (*(QMetaMethod*)rhs);
	return __r;
}
int QMetaMethod_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaMethod*)self) != (*(QMetaMethod*)rhs);
	return __r;
}
void QMetaMethod_delete(void* self) { delete (QMetaMethod*)self; }

}

#include <QtCore/qplugin.h>
#include <QtCore/qjsonobject.h>
#include "gen_qstaticplugin.h"

extern "C" {

void* QStaticPlugin_new(void* arg1) {
	return new QStaticPlugin(*(QStaticPlugin*)arg1);
}
void* QStaticPlugin_metaData(void* self) {
	auto&& __r = ((QStaticPlugin*)self)->metaData();
	return new QJsonObject(__r);
}
void QStaticPlugin_delete(void* self) { delete (QStaticPlugin*)self; }

}

#include <QtCore/qsequentialiterable.h>
#include <QtCore/qvariant.h>
#include <QtCore/qmetacontainer.h>
#include "gen_qsequentialconstiterator.h"

extern "C" {

void* QSequentialConstIterator_new(void* arg1) {
	return new QSequentialConstIterator(*(QSequentialConstIterator*)arg1);
}
void* QSequentialConstIterator_OperatorMultiply(void* self) {
	auto&& __r = ((QSequentialConstIterator*)self)->operator*();
	return new QVariant(__r);
}
void* QSequentialConstIterator_OperatorMinusGreater(void* self) {
	auto&& __r = ((QSequentialConstIterator*)self)->operator->();
	return new QVariantConstPointer(__r);
}
void* QSequentialConstIterator_metaContainer(void* self) {
	auto&& __r = ((QSequentialConstIterator*)self)->metaContainer();
	return new QMetaSequence(__r);
}
void QSequentialConstIterator_delete(void* self) { delete (QSequentialConstIterator*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qstringview.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstring.h>
#include "gen_qxmlstreamattributes.h"

extern "C" {

void* QXmlStreamAttributes_new() {
	return new QXmlStreamAttributes();
}
void* QXmlStreamAttributes_value(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QXmlStreamAttributes*)self)->value(namespaceUri_qstr, name_qstr);
	return new QStringView(__r);
}
void* QXmlStreamAttributes_value2(void* self, const char* qualifiedName, long long qualifiedName_len) {
	QAnyStringView qualifiedName_qstr = QAnyStringView(qualifiedName, qualifiedName_len);
	auto&& __r = ((QXmlStreamAttributes*)self)->value(qualifiedName_qstr);
	return new QStringView(__r);
}
void QXmlStreamAttributes_append(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* value, long long value_len) {
	QString namespaceUri_qstr = QString::fromUtf8(namespaceUri, namespaceUri_len);
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QXmlStreamAttributes*)self)->append(namespaceUri_qstr, name_qstr, value_qstr);
}
void QXmlStreamAttributes_append2(void* self, const char* qualifiedName, long long qualifiedName_len, const char* value, long long value_len) {
	QString qualifiedName_qstr = QString::fromUtf8(qualifiedName, qualifiedName_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QXmlStreamAttributes*)self)->append(qualifiedName_qstr, value_qstr);
}
int QXmlStreamAttributes_hasAttribute(void* self, const char* qualifiedName, long long qualifiedName_len) {
	QAnyStringView qualifiedName_qstr = QAnyStringView(qualifiedName, qualifiedName_len);
	auto&& __r = ((QXmlStreamAttributes*)self)->hasAttribute(qualifiedName_qstr);
	return __r;
}
int QXmlStreamAttributes_hasAttribute2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QXmlStreamAttributes*)self)->hasAttribute(namespaceUri_qstr, name_qstr);
	return __r;
}
long long QXmlStreamAttributes_maxSize() {
	auto&& __r = QXmlStreamAttributes::maxSize();
	return __r;
}
long long QXmlStreamAttributes_size(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->size();
	return __r;
}
long long QXmlStreamAttributes_count(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->count();
	return __r;
}
long long QXmlStreamAttributes_length(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->length();
	return __r;
}
int QXmlStreamAttributes_isEmpty(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->isEmpty();
	return __r;
}
void QXmlStreamAttributes_resize(void* self, long long size) {
	((QXmlStreamAttributes*)self)->resize(size);
}
void QXmlStreamAttributes_resize2(void* self, long long size, void* c) {
	((QXmlStreamAttributes*)self)->resize(size, *(QXmlStreamAttribute*)c);
}
void QXmlStreamAttributes_resizeForOverwrite(void* self, long long size) {
	((QXmlStreamAttributes*)self)->resizeForOverwrite(size);
}
long long QXmlStreamAttributes_capacity(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->capacity();
	return __r;
}
void QXmlStreamAttributes_reserve(void* self, long long size) {
	((QXmlStreamAttributes*)self)->reserve(size);
}
void QXmlStreamAttributes_squeeze(void* self) {
	((QXmlStreamAttributes*)self)->squeeze();
}
void QXmlStreamAttributes_detach(void* self) {
	((QXmlStreamAttributes*)self)->detach();
}
int QXmlStreamAttributes_isDetached(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->isDetached();
	return __r;
}
int QXmlStreamAttributes_isSharedWith(void* self, void** other, long long other_len) {
	QList<QXmlStreamAttribute> other_list; other_list.reserve(other_len);
	for (long long __i = 0; __i < other_len; __i++) other_list.append(*(QXmlStreamAttribute*)other[__i]);
	auto&& __r = ((QXmlStreamAttributes*)self)->isSharedWith(other_list);
	return __r;
}
void* QXmlStreamAttributes_data(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->data();
	return (void*)(__r);
}
void* QXmlStreamAttributes_data2(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->data();
	return (void*)(__r);
}
void* QXmlStreamAttributes_constData(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->constData();
	return (void*)(__r);
}
void QXmlStreamAttributes_clear(void* self) {
	((QXmlStreamAttributes*)self)->clear();
}
void* QXmlStreamAttributes_at(void* self, long long i) {
	auto&& __r = ((QXmlStreamAttributes*)self)->at(i);
	return (void*)&(__r);
}
void* QXmlStreamAttributes_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QXmlStreamAttributes*)self)->operator[](i);
	return (void*)&(__r);
}
void* QXmlStreamAttributes_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QXmlStreamAttributes*)self)->operator[](i);
	return (void*)&(__r);
}
void QXmlStreamAttributes_prepend(void* self, void* t) {
	((QXmlStreamAttributes*)self)->prepend(*(QXmlStreamAttribute*)t);
}
void QXmlStreamAttributes_replace(void* self, long long i, void* t) {
	((QXmlStreamAttributes*)self)->replace(i, *(QXmlStreamAttribute*)t);
}
void QXmlStreamAttributes_remove(void* self, long long i) {
	((QXmlStreamAttributes*)self)->remove(i);
}
void QXmlStreamAttributes_removeFirst(void* self) {
	((QXmlStreamAttributes*)self)->removeFirst();
}
void QXmlStreamAttributes_removeLast(void* self) {
	((QXmlStreamAttributes*)self)->removeLast();
}
void* QXmlStreamAttributes_takeFirst(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->takeFirst();
	return new QXmlStreamAttribute(__r);
}
void* QXmlStreamAttributes_takeLast(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->takeLast();
	return new QXmlStreamAttribute(__r);
}
void QXmlStreamAttributes_removeAt(void* self, long long i) {
	((QXmlStreamAttributes*)self)->removeAt(i);
}
void* QXmlStreamAttributes_takeAt(void* self, long long i) {
	auto&& __r = ((QXmlStreamAttributes*)self)->takeAt(i);
	return new QXmlStreamAttribute(__r);
}
void QXmlStreamAttributes_move(void* self, long long from, long long to) {
	((QXmlStreamAttributes*)self)->move(from, to);
}
void* QXmlStreamAttributes_first(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->first();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_first2(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->first();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_constFirst(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->constFirst();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_last(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->last();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_last2(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->last();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_constLast(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->constLast();
	return (void*)&(__r);
}
int QXmlStreamAttributes_startsWith(void* self, void* t) {
	auto&& __r = ((QXmlStreamAttributes*)self)->startsWith(*(QXmlStreamAttribute*)t);
	return __r;
}
int QXmlStreamAttributes_endsWith(void* self, void* t) {
	auto&& __r = ((QXmlStreamAttributes*)self)->endsWith(*(QXmlStreamAttribute*)t);
	return __r;
}
void** QXmlStreamAttributes_mid(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->mid(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_first3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->first(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_last3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->last(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_sliced(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->sliced(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_sliced2(void* self, long long pos, long long n, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->sliced(pos, n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void QXmlStreamAttributes_swapItemsAt(void* self, long long i, long long j) {
	((QXmlStreamAttributes*)self)->swapItemsAt(i, j);
}
void QXmlStreamAttributes_push_back(void* self, void* t) {
	((QXmlStreamAttributes*)self)->push_back(*(QXmlStreamAttribute*)t);
}
void QXmlStreamAttributes_push_front(void* self, void* t) {
	((QXmlStreamAttributes*)self)->push_front(*(QXmlStreamAttribute*)t);
}
void QXmlStreamAttributes_pop_back(void* self) {
	((QXmlStreamAttributes*)self)->pop_back();
}
void QXmlStreamAttributes_pop_front(void* self) {
	((QXmlStreamAttributes*)self)->pop_front();
}
int QXmlStreamAttributes_empty(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->empty();
	return __r;
}
void* QXmlStreamAttributes_front(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->front();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_front2(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->front();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_back(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->back();
	return (void*)&(__r);
}
void* QXmlStreamAttributes_back2(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->back();
	return (void*)&(__r);
}
void QXmlStreamAttributes_shrink_to_fit(void* self) {
	((QXmlStreamAttributes*)self)->shrink_to_fit();
}
long long QXmlStreamAttributes_max_size(void* self) {
	auto&& __r = ((QXmlStreamAttributes*)self)->max_size();
	return __r;
}
void** QXmlStreamAttributes_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QXmlStreamAttribute> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QXmlStreamAttribute*)l[__i]);
	auto&& __r = ((QXmlStreamAttributes*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QXmlStreamAttribute> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QXmlStreamAttribute*)l[__i]);
	auto&& __r = ((QXmlStreamAttributes*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_fromList(void** list, long long list_len, long long* __out_len) {
	QList<QXmlStreamAttribute> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QXmlStreamAttribute*)list[__i]);
	auto&& __r = QXmlStreamAttributes::fromList(list_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_toList(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->toList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_fromVector(void** vector, long long vector_len, long long* __out_len) {
	QList<QXmlStreamAttribute> vector_list; vector_list.reserve(vector_len);
	for (long long __i = 0; __i < vector_len; __i++) vector_list.append(*(QXmlStreamAttribute*)vector[__i]);
	auto&& __r = QXmlStreamAttributes::fromVector(vector_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void** QXmlStreamAttributes_toVector(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->toVector();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void QXmlStreamAttributes_remove2(void* self, long long i, long long n) {
	((QXmlStreamAttributes*)self)->remove(i, n);
}
void** QXmlStreamAttributes_mid2(void* self, long long pos, long long len_, long long* __out_len) {
	auto&& __r = ((QXmlStreamAttributes*)self)->mid(pos, len_);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamAttribute(__e);
	return __arr;
}
void QXmlStreamAttributes_delete(void* self) { delete (QXmlStreamAttributes*)self; }

}

#include <QtGui/qguiapplication.h>
#include <QtCore/qcoreapplication.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qwindow.h>
#include <QtCore/qpoint.h>
#include <QtGui/qicon.h>
#include <QtCore/qobject.h>
#include <QtGui/qscreen.h>
#include <QtGui/qcursor.h>
#include <QtGui/qfont.h>
#include <QtGui/qclipboard.h>
#include <QtGui/qpalette.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qstylehints.h>
#include <QtGui/qinputmethod.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qsessionmanager.h>
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
#include "gen_qguiapplication.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGuiApplication;
static std::mutex __connsMu_QGuiApplication;

extern "C" {

void* QGuiApplication_base(void* self) { return (void*)static_cast<QCoreApplication*>((QGuiApplication*)self); }
void QGuiApplication_chain(void* self, void** out) {
	QGuiApplication* __p = (QGuiApplication*)self;
	QCoreApplication* __b0 = static_cast<QCoreApplication*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QGuiApplication_fromQObject(void* o) { return (void*)qobject_cast<QGuiApplication*>((QObject*)o); }
char* QGuiApplication_tr(const char* s, long long* __out_len) {
	auto&& __r = QGuiApplication::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setApplicationDisplayName(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QGuiApplication::setApplicationDisplayName(name_qstr);
}
char* QGuiApplication_applicationDisplayName(long long* __out_len) {
	auto&& __r = QGuiApplication::applicationDisplayName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setBadgeNumber(void* self, long long number) {
	((QGuiApplication*)self)->setBadgeNumber(number);
}
void QGuiApplication_setDesktopFileName(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QGuiApplication::setDesktopFileName(name_qstr);
}
char* QGuiApplication_desktopFileName(long long* __out_len) {
	auto&& __r = QGuiApplication::desktopFileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QGuiApplication_allWindows(long long* __out_len) {
	auto&& __r = QGuiApplication::allWindows();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGuiApplication_topLevelWindows(long long* __out_len) {
	auto&& __r = QGuiApplication::topLevelWindows();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QGuiApplication_topLevelAt(void* pos) {
	auto&& __r = QGuiApplication::topLevelAt(*(QPoint*)pos);
	return (void*)(__r);
}
void QGuiApplication_setWindowIcon(void* icon) {
	QGuiApplication::setWindowIcon(*(QIcon*)icon);
}
void* QGuiApplication_windowIcon() {
	auto&& __r = QGuiApplication::windowIcon();
	return new QIcon(__r);
}
char* QGuiApplication_platformName(long long* __out_len) {
	auto&& __r = QGuiApplication::platformName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGuiApplication_modalWindow() {
	auto&& __r = QGuiApplication::modalWindow();
	return (void*)(__r);
}
void* QGuiApplication_focusWindow() {
	auto&& __r = QGuiApplication::focusWindow();
	return (void*)(__r);
}
void* QGuiApplication_focusObject() {
	auto&& __r = QGuiApplication::focusObject();
	return (void*)(__r);
}
void* QGuiApplication_primaryScreen() {
	auto&& __r = QGuiApplication::primaryScreen();
	return (void*)(__r);
}
void** QGuiApplication_screens(long long* __out_len) {
	auto&& __r = QGuiApplication::screens();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QGuiApplication_screenAt(void* point) {
	auto&& __r = QGuiApplication::screenAt(*(QPoint*)point);
	return (void*)(__r);
}
double QGuiApplication_devicePixelRatio(void* self) {
	auto&& __r = ((QGuiApplication*)self)->devicePixelRatio();
	return __r;
}
void* QGuiApplication_overrideCursor() {
	auto&& __r = QGuiApplication::overrideCursor();
	return (void*)(__r);
}
void QGuiApplication_setOverrideCursor(void* arg1) {
	QGuiApplication::setOverrideCursor(*(QCursor*)arg1);
}
void QGuiApplication_changeOverrideCursor(void* arg1) {
	QGuiApplication::changeOverrideCursor(*(QCursor*)arg1);
}
void QGuiApplication_restoreOverrideCursor() {
	QGuiApplication::restoreOverrideCursor();
}
void* QGuiApplication_font() {
	auto&& __r = QGuiApplication::font();
	return new QFont(__r);
}
void QGuiApplication_setFont(void* arg1) {
	QGuiApplication::setFont(*(QFont*)arg1);
}
void* QGuiApplication_clipboard() {
	auto&& __r = QGuiApplication::clipboard();
	return (void*)(__r);
}
void* QGuiApplication_palette() {
	auto&& __r = QGuiApplication::palette();
	return new QPalette(__r);
}
void QGuiApplication_setPalette(void* pal) {
	QGuiApplication::setPalette(*(QPalette*)pal);
}
int QGuiApplication_keyboardModifiers() {
	auto&& __r = QGuiApplication::keyboardModifiers();
	return (int)(__r);
}
int QGuiApplication_queryKeyboardModifiers() {
	auto&& __r = QGuiApplication::queryKeyboardModifiers();
	return (int)(__r);
}
int QGuiApplication_mouseButtons() {
	auto&& __r = QGuiApplication::mouseButtons();
	return (int)(__r);
}
void QGuiApplication_setLayoutDirection(int direction) {
	QGuiApplication::setLayoutDirection((Qt::LayoutDirection)(direction));
}
int QGuiApplication_layoutDirection() {
	auto&& __r = QGuiApplication::layoutDirection();
	return (int)(__r);
}
int QGuiApplication_isRightToLeft() {
	auto&& __r = QGuiApplication::isRightToLeft();
	return __r;
}
int QGuiApplication_isLeftToRight() {
	auto&& __r = QGuiApplication::isLeftToRight();
	return __r;
}
void* QGuiApplication_styleHints() {
	auto&& __r = QGuiApplication::styleHints();
	return (void*)(__r);
}
void QGuiApplication_setDesktopSettingsAware(int on) {
	QGuiApplication::setDesktopSettingsAware(on);
}
int QGuiApplication_desktopSettingsAware() {
	auto&& __r = QGuiApplication::desktopSettingsAware();
	return __r;
}
void* QGuiApplication_inputMethod() {
	auto&& __r = QGuiApplication::inputMethod();
	return (void*)(__r);
}
void QGuiApplication_setQuitOnLastWindowClosed(int quit) {
	QGuiApplication::setQuitOnLastWindowClosed(quit);
}
int QGuiApplication_quitOnLastWindowClosed() {
	auto&& __r = QGuiApplication::quitOnLastWindowClosed();
	return __r;
}
int QGuiApplication_applicationState() {
	auto&& __r = QGuiApplication::applicationState();
	return (int)(__r);
}
void QGuiApplication_setHighDpiScaleFactorRoundingPolicy(int policy) {
	QGuiApplication::setHighDpiScaleFactorRoundingPolicy((Qt::HighDpiScaleFactorRoundingPolicy)(policy));
}
int QGuiApplication_highDpiScaleFactorRoundingPolicy() {
	auto&& __r = QGuiApplication::highDpiScaleFactorRoundingPolicy();
	return (int)(__r);
}
int QGuiApplication_exec() {
	auto&& __r = QGuiApplication::exec();
	return __r;
}
int QGuiApplication_isSessionRestored(void* self) {
	auto&& __r = ((QGuiApplication*)self)->isSessionRestored();
	return __r;
}
char* QGuiApplication_sessionId(void* self, long long* __out_len) {
	auto&& __r = ((QGuiApplication*)self)->sessionId();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGuiApplication_sessionKey(void* self, long long* __out_len) {
	auto&& __r = ((QGuiApplication*)self)->sessionKey();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGuiApplication_isSavingSession(void* self) {
	auto&& __r = ((QGuiApplication*)self)->isSavingSession();
	return __r;
}
void QGuiApplication_sync() {
	QGuiApplication::sync();
}
void QGuiApplication_fontDatabaseChanged(void* self) {
	((QGuiApplication*)self)->fontDatabaseChanged();
}
void QGuiApplication_screenAdded(void* self, void* screen) {
	((QGuiApplication*)self)->screenAdded((QScreen*)screen);
}
void QGuiApplication_screenRemoved(void* self, void* screen) {
	((QGuiApplication*)self)->screenRemoved((QScreen*)screen);
}
void QGuiApplication_primaryScreenChanged(void* self, void* screen) {
	((QGuiApplication*)self)->primaryScreenChanged((QScreen*)screen);
}
void QGuiApplication_lastWindowClosed(void* self) {
	((QGuiApplication*)self)->lastWindowClosed();
}
void QGuiApplication_focusObjectChanged(void* self, void* focusObject) {
	((QGuiApplication*)self)->focusObjectChanged((QObject*)focusObject);
}
void QGuiApplication_focusWindowChanged(void* self, void* focusWindow) {
	((QGuiApplication*)self)->focusWindowChanged((QWindow*)focusWindow);
}
void QGuiApplication_applicationStateChanged(void* self, int state) {
	((QGuiApplication*)self)->applicationStateChanged((Qt::ApplicationState)(state));
}
void QGuiApplication_layoutDirectionChanged(void* self, int direction) {
	((QGuiApplication*)self)->layoutDirectionChanged((Qt::LayoutDirection)(direction));
}
void QGuiApplication_commitDataRequest(void* self, void* sessionManager) {
	((QGuiApplication*)self)->commitDataRequest(*(QSessionManager*)sessionManager);
}
void QGuiApplication_saveStateRequest(void* self, void* sessionManager) {
	((QGuiApplication*)self)->saveStateRequest(*(QSessionManager*)sessionManager);
}
void QGuiApplication_applicationDisplayNameChanged(void* self) {
	((QGuiApplication*)self)->applicationDisplayNameChanged();
}
void QGuiApplication_paletteChanged(void* self, void* pal) {
	((QGuiApplication*)self)->paletteChanged(*(QPalette*)pal);
}
void QGuiApplication_fontChanged(void* self, void* font) {
	((QGuiApplication*)self)->fontChanged(*(QFont*)font);
}
char** QGuiApplication_arguments(long long* __out_len) {
	auto&& __r = QGuiApplication::arguments();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QGuiApplication_setAttribute(int attribute) {
	QGuiApplication::setAttribute((Qt::ApplicationAttribute)(attribute));
}
int QGuiApplication_testAttribute(int attribute) {
	auto&& __r = QGuiApplication::testAttribute((Qt::ApplicationAttribute)(attribute));
	return __r;
}
void QGuiApplication_setOrganizationDomain(const char* orgDomain, long long orgDomain_len) {
	QString orgDomain_qstr = QString::fromUtf8(orgDomain, orgDomain_len);
	QGuiApplication::setOrganizationDomain(orgDomain_qstr);
}
char* QGuiApplication_organizationDomain(long long* __out_len) {
	auto&& __r = QGuiApplication::organizationDomain();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setOrganizationName(const char* orgName, long long orgName_len) {
	QString orgName_qstr = QString::fromUtf8(orgName, orgName_len);
	QGuiApplication::setOrganizationName(orgName_qstr);
}
char* QGuiApplication_organizationName(long long* __out_len) {
	auto&& __r = QGuiApplication::organizationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setApplicationName(const char* application, long long application_len) {
	QString application_qstr = QString::fromUtf8(application, application_len);
	QGuiApplication::setApplicationName(application_qstr);
}
char* QGuiApplication_applicationName(long long* __out_len) {
	auto&& __r = QGuiApplication::applicationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setApplicationVersion(const char* version, long long version_len) {
	QString version_qstr = QString::fromUtf8(version, version_len);
	QGuiApplication::setApplicationVersion(version_qstr);
}
char* QGuiApplication_applicationVersion(long long* __out_len) {
	auto&& __r = QGuiApplication::applicationVersion();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setSetuidAllowed(int allow) {
	QGuiApplication::setSetuidAllowed(allow);
}
int QGuiApplication_isSetuidAllowed() {
	auto&& __r = QGuiApplication::isSetuidAllowed();
	return __r;
}
void* QGuiApplication_instance() {
	auto&& __r = QGuiApplication::instance();
	return (void*)(__r);
}
int QGuiApplication_instanceExists() {
	auto&& __r = QGuiApplication::instanceExists();
	return __r;
}
void QGuiApplication_processEvents() {
	QGuiApplication::processEvents();
}
void QGuiApplication_processEvents2(int flags, int maxtime) {
	QGuiApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), maxtime);
}
void QGuiApplication_processEvents3(int flags, void* deadline) {
	QGuiApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), *(QDeadlineTimer*)deadline);
}
int QGuiApplication_sendEvent(void* receiver, void* event) {
	auto&& __r = QGuiApplication::sendEvent((QObject*)receiver, (QEvent*)event);
	return __r;
}
void QGuiApplication_postEvent(void* receiver, void* event) {
	QGuiApplication::postEvent((QObject*)receiver, (QEvent*)event);
}
void QGuiApplication_sendPostedEvents() {
	QGuiApplication::sendPostedEvents();
}
void QGuiApplication_removePostedEvents(void* receiver) {
	QGuiApplication::removePostedEvents((QObject*)receiver);
}
void* QGuiApplication_eventDispatcher() {
	auto&& __r = QGuiApplication::eventDispatcher();
	return (void*)(__r);
}
void QGuiApplication_setEventDispatcher(void* eventDispatcher) {
	QGuiApplication::setEventDispatcher((QAbstractEventDispatcher*)eventDispatcher);
}
int QGuiApplication_startingUp() {
	auto&& __r = QGuiApplication::startingUp();
	return __r;
}
int QGuiApplication_closingDown() {
	auto&& __r = QGuiApplication::closingDown();
	return __r;
}
char* QGuiApplication_applicationDirPath(long long* __out_len) {
	auto&& __r = QGuiApplication::applicationDirPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGuiApplication_applicationFilePath(long long* __out_len) {
	auto&& __r = QGuiApplication::applicationFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QGuiApplication_applicationPid() {
	auto&& __r = QGuiApplication::applicationPid();
	return __r;
}
void QGuiApplication_setLibraryPaths(char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	QGuiApplication::setLibraryPaths(arg1_list);
}
char** QGuiApplication_libraryPaths(long long* __out_len) {
	auto&& __r = QGuiApplication::libraryPaths();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QGuiApplication_addLibraryPath(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QGuiApplication::addLibraryPath(arg1_qstr);
}
void QGuiApplication_removeLibraryPath(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QGuiApplication::removeLibraryPath(arg1_qstr);
}
int QGuiApplication_installTranslator(void* messageFile) {
	auto&& __r = QGuiApplication::installTranslator((QTranslator*)messageFile);
	return __r;
}
int QGuiApplication_removeTranslator(void* messageFile) {
	auto&& __r = QGuiApplication::removeTranslator((QTranslator*)messageFile);
	return __r;
}
char* QGuiApplication_translate(const char* context, const char* key, long long* __out_len) {
	auto&& __r = QGuiApplication::translate(context, key);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGuiApplication_isQuitLockEnabled() {
	auto&& __r = QGuiApplication::isQuitLockEnabled();
	return __r;
}
void QGuiApplication_setQuitLockEnabled(int enabled) {
	QGuiApplication::setQuitLockEnabled(enabled);
}
void QGuiApplication_quit() {
	QGuiApplication::quit();
}
void QGuiApplication_exit() {
	QGuiApplication::exit();
}
void* QGuiApplication_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGuiApplication::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGuiApplication_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGuiApplication::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QGuiApplication_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGuiApplication::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGuiApplication_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGuiApplication::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_setAttribute2(int attribute, int on) {
	QGuiApplication::setAttribute((Qt::ApplicationAttribute)(attribute), on);
}
void QGuiApplication_processEvents4(int flags) {
	QGuiApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags));
}
void QGuiApplication_postEvent2(void* receiver, void* event, int priority) {
	QGuiApplication::postEvent((QObject*)receiver, (QEvent*)event, priority);
}
void QGuiApplication_sendPostedEvents2(void* receiver) {
	QGuiApplication::sendPostedEvents((QObject*)receiver);
}
void QGuiApplication_sendPostedEvents3(void* receiver, int event_type) {
	QGuiApplication::sendPostedEvents((QObject*)receiver, event_type);
}
void QGuiApplication_removePostedEvents2(void* receiver, int eventType) {
	QGuiApplication::removePostedEvents((QObject*)receiver, eventType);
}
char* QGuiApplication_translate2(const char* context, const char* key, const char* disambiguation, long long* __out_len) {
	auto&& __r = QGuiApplication::translate(context, key, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGuiApplication_translate3(const char* context, const char* key, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QGuiApplication::translate(context, key, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGuiApplication_exit2(int retcode) {
	QGuiApplication::exit(retcode);
}
void* QGuiApplication_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGuiApplication::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGuiApplication_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGuiApplication::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QGuiApplication_connect_fontDatabaseChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::fontDatabaseChanged, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_screenAdded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::screenAdded, (QGuiApplication*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_screenRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::screenRemoved, (QGuiApplication*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_primaryScreenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::primaryScreenChanged, (QGuiApplication*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_lastWindowClosed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::lastWindowClosed, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_focusObjectChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::focusObjectChanged, (QGuiApplication*)self, [__id](QObject* focusObject){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)focusObject);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_focusWindowChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::focusWindowChanged, (QGuiApplication*)self, [__id](QWindow* focusWindow){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)focusWindow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_applicationStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::applicationStateChanged, (QGuiApplication*)self, [__id](Qt::ApplicationState state){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(state));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_layoutDirectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::layoutDirectionChanged, (QGuiApplication*)self, [__id](Qt::LayoutDirection direction){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(direction));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_commitDataRequest(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::commitDataRequest, (QGuiApplication*)self, [__id](QSessionManager& sessionManager){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&sessionManager);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_saveStateRequest(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::saveStateRequest, (QGuiApplication*)self, [__id](QSessionManager& sessionManager){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&sessionManager);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_applicationDisplayNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::applicationDisplayNameChanged, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_paletteChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::paletteChanged, (QGuiApplication*)self, [__id](const QPalette& pal){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pal);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_fontChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::fontChanged, (QGuiApplication*)self, [__id](const QFont& font){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&font);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_aboutToQuit(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::aboutToQuit, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_organizationNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::organizationNameChanged, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_organizationDomainChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::organizationDomainChanged, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_applicationNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::applicationNameChanged, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_applicationVersionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::applicationVersionChanged, (QGuiApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::destroyed, (QGuiApplication*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::destroyed, (QGuiApplication*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGuiApplication*)self, &QGuiApplication::objectNameChanged, (QGuiApplication*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
	__conns_QGuiApplication[__id] = __c;
}
void QGuiApplication__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGuiApplication);
		auto __it = __conns_QGuiApplication.find(id);
		if (__it == __conns_QGuiApplication.end()) return;
		__c = __it->second;
		__conns_QGuiApplication.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGuiApplication_new_args(int* argc, char** argv) { return new QGuiApplication(*argc, argv); }
void QGuiApplication_delete(void* self) { delete (QGuiApplication*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qshowevent.h"
#include "gen_purego_cb.h"

class QShowEvent__Go : public QShowEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QShowEvent__Go(long long id, A&&... a) : QShowEvent(std::forward<A>(a)...), __goId(id) {}
	QShowEvent* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QShowEvent*)__r; }
		return QShowEvent::clone();
	}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QShowEvent::setAccepted(accepted);
	}
};

extern "C" {

void* QShowEvent_base(void* self) { return (void*)static_cast<QEvent*>((QShowEvent*)self); }
void QShowEvent__set_override(void* self, int idx) { ((QShowEvent__Go*)self)->__goOn[idx] = 1; }
void* QShowEvent_new(long long __goid) {
	return new QShowEvent__Go(__goid);
}
int QShowEvent_registerEventType() {
	auto&& __r = QShowEvent::registerEventType();
	return __r;
}
int QShowEvent_registerEventType2(int hint) {
	auto&& __r = QShowEvent::registerEventType(hint);
	return __r;
}
void* QShowEvent_clone_default(void* self) {
	auto&& __r = ((QShowEvent*)self)->QShowEvent::clone();
	return (void*)(__r);
}
void QShowEvent_setAccepted_default(void* self, int accepted) {
	((QShowEvent*)self)->QShowEvent::setAccepted(accepted);
}
void* QShowEvent__Go_new(long long id) { return new QShowEvent__Go(id); }
void QShowEvent_delete(void* self) { delete (QShowEvent*)self; }

}

#include <QtGui/qfont.h>
#include <QtCore/qstring.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qvariant.h>
#include "gen_qfont.h"

extern "C" {

void* QFont_new() {
	return new QFont();
}
void* QFont_new2(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	return new QFont(family_qstr);
}
void* QFont_new3(char** families, long long families_len) {
	QList<QString> families_list; families_list.reserve(families_len);
	for (long long __i = 0; __i < families_len; __i++) families_list.append(QString::fromUtf8(families[__i]));
	return new QFont(families_list);
}
void* QFont_new4(void* font, void* pd) {
	return new QFont(*(QFont*)font, (QPaintDevice*)pd);
}
void* QFont_new5(void* font) {
	return new QFont(*(QFont*)font);
}
void* QFont_new6(const char* family, long long family_len, int pointSize) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	return new QFont(family_qstr, pointSize);
}
void* QFont_new7(const char* family, long long family_len, int pointSize, int weight) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	return new QFont(family_qstr, pointSize, weight);
}
void* QFont_new8(const char* family, long long family_len, int pointSize, int weight, int italic) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	return new QFont(family_qstr, pointSize, weight, italic);
}
void* QFont_new9(char** families, long long families_len, int pointSize) {
	QList<QString> families_list; families_list.reserve(families_len);
	for (long long __i = 0; __i < families_len; __i++) families_list.append(QString::fromUtf8(families[__i]));
	return new QFont(families_list, pointSize);
}
void* QFont_new10(char** families, long long families_len, int pointSize, int weight) {
	QList<QString> families_list; families_list.reserve(families_len);
	for (long long __i = 0; __i < families_len; __i++) families_list.append(QString::fromUtf8(families[__i]));
	return new QFont(families_list, pointSize, weight);
}
void* QFont_new11(char** families, long long families_len, int pointSize, int weight, int italic) {
	QList<QString> families_list; families_list.reserve(families_len);
	for (long long __i = 0; __i < families_len; __i++) families_list.append(QString::fromUtf8(families[__i]));
	return new QFont(families_list, pointSize, weight, italic);
}
void QFont_swap(void* self, void* other) {
	((QFont*)self)->swap(*(QFont*)other);
}
char* QFont_family(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->family();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFont_setFamily(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QFont*)self)->setFamily(arg1_qstr);
}
char** QFont_families(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->families();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFont_setFamilies(void* self, char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	((QFont*)self)->setFamilies(arg1_list);
}
char* QFont_styleName(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->styleName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFont_setStyleName(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QFont*)self)->setStyleName(arg1_qstr);
}
int QFont_pointSize(void* self) {
	auto&& __r = ((QFont*)self)->pointSize();
	return __r;
}
void QFont_setPointSize(void* self, int arg1) {
	((QFont*)self)->setPointSize(arg1);
}
double QFont_pointSizeF(void* self) {
	auto&& __r = ((QFont*)self)->pointSizeF();
	return __r;
}
void QFont_setPointSizeF(void* self, double arg1) {
	((QFont*)self)->setPointSizeF(arg1);
}
int QFont_pixelSize(void* self) {
	auto&& __r = ((QFont*)self)->pixelSize();
	return __r;
}
void QFont_setPixelSize(void* self, int arg1) {
	((QFont*)self)->setPixelSize(arg1);
}
int QFont_weight(void* self) {
	auto&& __r = ((QFont*)self)->weight();
	return (int)(__r);
}
void QFont_setWeight(void* self, int weight) {
	((QFont*)self)->setWeight((QFont::Weight)(weight));
}
int QFont_bold(void* self) {
	auto&& __r = ((QFont*)self)->bold();
	return __r;
}
void QFont_setBold(void* self, int arg1) {
	((QFont*)self)->setBold(arg1);
}
void QFont_setStyle(void* self, int style) {
	((QFont*)self)->setStyle((QFont::Style)(style));
}
int QFont_style(void* self) {
	auto&& __r = ((QFont*)self)->style();
	return (int)(__r);
}
int QFont_italic(void* self) {
	auto&& __r = ((QFont*)self)->italic();
	return __r;
}
void QFont_setItalic(void* self, int b) {
	((QFont*)self)->setItalic(b);
}
int QFont_underline(void* self) {
	auto&& __r = ((QFont*)self)->underline();
	return __r;
}
void QFont_setUnderline(void* self, int arg1) {
	((QFont*)self)->setUnderline(arg1);
}
int QFont_overline(void* self) {
	auto&& __r = ((QFont*)self)->overline();
	return __r;
}
void QFont_setOverline(void* self, int arg1) {
	((QFont*)self)->setOverline(arg1);
}
int QFont_strikeOut(void* self) {
	auto&& __r = ((QFont*)self)->strikeOut();
	return __r;
}
void QFont_setStrikeOut(void* self, int arg1) {
	((QFont*)self)->setStrikeOut(arg1);
}
int QFont_fixedPitch(void* self) {
	auto&& __r = ((QFont*)self)->fixedPitch();
	return __r;
}
void QFont_setFixedPitch(void* self, int arg1) {
	((QFont*)self)->setFixedPitch(arg1);
}
int QFont_kerning(void* self) {
	auto&& __r = ((QFont*)self)->kerning();
	return __r;
}
void QFont_setKerning(void* self, int arg1) {
	((QFont*)self)->setKerning(arg1);
}
int QFont_styleHint(void* self) {
	auto&& __r = ((QFont*)self)->styleHint();
	return (int)(__r);
}
int QFont_styleStrategy(void* self) {
	auto&& __r = ((QFont*)self)->styleStrategy();
	return (int)(__r);
}
void QFont_setStyleHint(void* self, int arg1) {
	((QFont*)self)->setStyleHint((QFont::StyleHint)(arg1));
}
void QFont_setStyleStrategy(void* self, int s) {
	((QFont*)self)->setStyleStrategy((QFont::StyleStrategy)(s));
}
int QFont_stretch(void* self) {
	auto&& __r = ((QFont*)self)->stretch();
	return __r;
}
void QFont_setStretch(void* self, int arg1) {
	((QFont*)self)->setStretch(arg1);
}
double QFont_letterSpacing(void* self) {
	auto&& __r = ((QFont*)self)->letterSpacing();
	return __r;
}
int QFont_letterSpacingType(void* self) {
	auto&& __r = ((QFont*)self)->letterSpacingType();
	return (int)(__r);
}
void QFont_setLetterSpacing(void* self, int type_, double spacing) {
	((QFont*)self)->setLetterSpacing((QFont::SpacingType)(type_), spacing);
}
double QFont_wordSpacing(void* self) {
	auto&& __r = ((QFont*)self)->wordSpacing();
	return __r;
}
void QFont_setWordSpacing(void* self, double spacing) {
	((QFont*)self)->setWordSpacing(spacing);
}
void QFont_setCapitalization(void* self, int arg1) {
	((QFont*)self)->setCapitalization((QFont::Capitalization)(arg1));
}
int QFont_capitalization(void* self) {
	auto&& __r = ((QFont*)self)->capitalization();
	return (int)(__r);
}
void QFont_setHintingPreference(void* self, int hintingPreference) {
	((QFont*)self)->setHintingPreference((QFont::HintingPreference)(hintingPreference));
}
int QFont_hintingPreference(void* self) {
	auto&& __r = ((QFont*)self)->hintingPreference();
	return (int)(__r);
}
void QFont_setFeature(void* self, void* tag, unsigned int value) {
	((QFont*)self)->setFeature(*(QFont::Tag*)tag, value);
}
void QFont_unsetFeature(void* self, void* tag) {
	((QFont*)self)->unsetFeature(*(QFont::Tag*)tag);
}
unsigned int QFont_featureValue(void* self, void* tag) {
	auto&& __r = ((QFont*)self)->featureValue(*(QFont::Tag*)tag);
	return __r;
}
int QFont_isFeatureSet(void* self, void* tag) {
	auto&& __r = ((QFont*)self)->isFeatureSet(*(QFont::Tag*)tag);
	return __r;
}
void** QFont_featureTags(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->featureTags();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFont::Tag(__e);
	return __arr;
}
void QFont_clearFeatures(void* self) {
	((QFont*)self)->clearFeatures();
}
void QFont_setVariableAxis(void* self, void* tag, float value) {
	((QFont*)self)->setVariableAxis(*(QFont::Tag*)tag, value);
}
void QFont_unsetVariableAxis(void* self, void* tag) {
	((QFont*)self)->unsetVariableAxis(*(QFont::Tag*)tag);
}
int QFont_isVariableAxisSet(void* self, void* tag) {
	auto&& __r = ((QFont*)self)->isVariableAxisSet(*(QFont::Tag*)tag);
	return __r;
}
float QFont_variableAxisValue(void* self, void* tag) {
	auto&& __r = ((QFont*)self)->variableAxisValue(*(QFont::Tag*)tag);
	return __r;
}
void QFont_clearVariableAxes(void* self) {
	((QFont*)self)->clearVariableAxes();
}
void** QFont_variableAxisTags(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->variableAxisTags();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFont::Tag(__e);
	return __arr;
}
int QFont_exactMatch(void* self) {
	auto&& __r = ((QFont*)self)->exactMatch();
	return __r;
}
void* QFont_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QFont*)self)->operator=(*(QFont*)arg1);
	return (void*)&(__r);
}
int QFont_OperatorEqual(void* self, void* arg1) {
	auto&& __r = ((QFont*)self)->operator==(*(QFont*)arg1);
	return __r;
}
int QFont_OperatorNotEqual(void* self, void* arg1) {
	auto&& __r = ((QFont*)self)->operator!=(*(QFont*)arg1);
	return __r;
}
int QFont_OperatorLesser(void* self, void* arg1) {
	auto&& __r = ((QFont*)self)->operator<(*(QFont*)arg1);
	return __r;
}
void* QFont_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QFont*)self);
	return new QVariant(__r);
}
int QFont_isCopyOf(void* self, void* arg1) {
	auto&& __r = ((QFont*)self)->isCopyOf(*(QFont*)arg1);
	return __r;
}
char* QFont_key(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->key();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFont_toString(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFont_fromString(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = ((QFont*)self)->fromString(arg1_qstr);
	return __r;
}
char* QFont_substitute(const char* arg1, long long arg1_len, long long* __out_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QFont::substitute(arg1_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QFont_substitutes(const char* arg1, long long arg1_len, long long* __out_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QFont::substitutes(arg1_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFont_substitutions(long long* __out_len) {
	auto&& __r = QFont::substitutions();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFont_insertSubstitution(const char* arg1, long long arg1_len, const char* arg2, long long arg2_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QString arg2_qstr = QString::fromUtf8(arg2, arg2_len);
	QFont::insertSubstitution(arg1_qstr, arg2_qstr);
}
void QFont_insertSubstitutions(const char* arg1, long long arg1_len, char** arg2, long long arg2_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QList<QString> arg2_list; arg2_list.reserve(arg2_len);
	for (long long __i = 0; __i < arg2_len; __i++) arg2_list.append(QString::fromUtf8(arg2[__i]));
	QFont::insertSubstitutions(arg1_qstr, arg2_list);
}
void QFont_removeSubstitutions(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QFont::removeSubstitutions(arg1_qstr);
}
void QFont_initialize() {
	QFont::initialize();
}
void QFont_cleanup() {
	QFont::cleanup();
}
void QFont_cacheStatistics() {
	QFont::cacheStatistics();
}
char* QFont_defaultFamily(void* self, long long* __out_len) {
	auto&& __r = ((QFont*)self)->defaultFamily();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFont_resolve(void* self, void* arg1) {
	auto&& __r = ((QFont*)self)->resolve(*(QFont*)arg1);
	return new QFont(__r);
}
unsigned int QFont_resolveMask(void* self) {
	auto&& __r = ((QFont*)self)->resolveMask();
	return __r;
}
void QFont_setResolveMask(void* self, unsigned int mask) {
	((QFont*)self)->setResolveMask(mask);
}
void QFont_setLegacyWeight(void* self, int legacyWeight) {
	((QFont*)self)->setLegacyWeight(legacyWeight);
}
int QFont_legacyWeight(void* self) {
	auto&& __r = ((QFont*)self)->legacyWeight();
	return __r;
}
void QFont_setStyleHint2(void* self, int arg1, int arg2) {
	((QFont*)self)->setStyleHint((QFont::StyleHint)(arg1), (QFont::StyleStrategy)(arg2));
}
void* QFont_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QFont*)self);
	return new QVariant(__r);
}
void QFont_delete(void* self) { delete (QFont*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbrush.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstring.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include "gen_qtexttableformat.h"

extern "C" {

void* QTextTableFormat_base(void* self) { return (void*)static_cast<QTextFrameFormat*>((QTextTableFormat*)self); }
void QTextTableFormat_chain(void* self, void** out) {
	QTextTableFormat* __p = (QTextTableFormat*)self;
	QTextFrameFormat* __b0 = static_cast<QTextFrameFormat*>(__p); out[0] = (void*)__b0;
	QTextFormat* __b1 = static_cast<QTextFormat*>(__b0); out[1] = (void*)__b1;
}
void* QTextTableFormat_new() {
	return new QTextTableFormat();
}
void* QTextTableFormat_new2(void* arg1) {
	return new QTextTableFormat(*(QTextTableFormat*)arg1);
}
int QTextTableFormat_isValid(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->isValid();
	return __r;
}
int QTextTableFormat_columns(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->columns();
	return __r;
}
void QTextTableFormat_setColumns(void* self, int columns) {
	((QTextTableFormat*)self)->setColumns(columns);
}
void QTextTableFormat_setColumnWidthConstraints(void* self, void** constraints, long long constraints_len) {
	QList<QTextLength> constraints_list; constraints_list.reserve(constraints_len);
	for (long long __i = 0; __i < constraints_len; __i++) constraints_list.append(*(QTextLength*)constraints[__i]);
	((QTextTableFormat*)self)->setColumnWidthConstraints(constraints_list);
}
void** QTextTableFormat_columnWidthConstraints(void* self, long long* __out_len) {
	auto&& __r = ((QTextTableFormat*)self)->columnWidthConstraints();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextLength(__e);
	return __arr;
}
void QTextTableFormat_clearColumnWidthConstraints(void* self) {
	((QTextTableFormat*)self)->clearColumnWidthConstraints();
}
double QTextTableFormat_cellSpacing(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->cellSpacing();
	return __r;
}
void QTextTableFormat_setCellSpacing(void* self, double spacing) {
	((QTextTableFormat*)self)->setCellSpacing(spacing);
}
double QTextTableFormat_cellPadding(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->cellPadding();
	return __r;
}
void QTextTableFormat_setCellPadding(void* self, double padding) {
	((QTextTableFormat*)self)->setCellPadding(padding);
}
void QTextTableFormat_setAlignment(void* self, int alignment) {
	((QTextTableFormat*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QTextTableFormat_alignment(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->alignment();
	return (int)(__r);
}
void QTextTableFormat_setHeaderRowCount(void* self, int count) {
	((QTextTableFormat*)self)->setHeaderRowCount(count);
}
int QTextTableFormat_headerRowCount(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->headerRowCount();
	return __r;
}
void QTextTableFormat_setBorderCollapse(void* self, int borderCollapse) {
	((QTextTableFormat*)self)->setBorderCollapse(borderCollapse);
}
int QTextTableFormat_borderCollapse(void* self) {
	auto&& __r = ((QTextTableFormat*)self)->borderCollapse();
	return __r;
}
void QTextTableFormat_delete(void* self) { delete (QTextTableFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include "gen_qaccessiblehyperlinkinterface.h"
#include "gen_purego_cb.h"

class QAccessibleHyperlinkInterface__Go : public QAccessibleHyperlinkInterface {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QAccessibleHyperlinkInterface__Go(long long id, A&&... a) : QAccessibleHyperlinkInterface(std::forward<A>(a)...), __goId(id) {}
	QString anchor() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, 2999672799903939988LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	QString anchorTarget() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, -8101286066007612963LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	int startIndex() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1674626578571099579LL, &__h); if (__h) return __r; }
		return int{};
	}
	int endIndex() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 4812684365611001022LL, &__h); if (__h) return __r; }
		return int{};
	}
	bool isValid() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 7231157875507698241LL, &__h); if (__h) return __r; }
		return bool{};
	}
};

extern "C" {

void QAccessibleHyperlinkInterface__set_override(void* self, int idx) { ((QAccessibleHyperlinkInterface__Go*)self)->__goOn[idx] = 1; }
char* QAccessibleHyperlinkInterface_anchor(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleHyperlinkInterface*)self)->anchor();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleHyperlinkInterface_anchorTarget(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleHyperlinkInterface*)self)->anchorTarget();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleHyperlinkInterface_startIndex(void* self) {
	auto&& __r = ((QAccessibleHyperlinkInterface*)self)->startIndex();
	return __r;
}
int QAccessibleHyperlinkInterface_endIndex(void* self) {
	auto&& __r = ((QAccessibleHyperlinkInterface*)self)->endIndex();
	return __r;
}
int QAccessibleHyperlinkInterface_isValid(void* self) {
	auto&& __r = ((QAccessibleHyperlinkInterface*)self)->isValid();
	return __r;
}
void* QAccessibleHyperlinkInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleHyperlinkInterface*)self)->operator=(*(QAccessibleHyperlinkInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleHyperlinkInterface__Go_new(long long id) { return new QAccessibleHyperlinkInterface__Go(id); }
void QAccessibleHyperlinkInterface_delete(void* self) { delete (QAccessibleHyperlinkInterface*)self; }

}

#include <QtGui/qwindow.h>
#include <QtCore/qobject.h>
#include <QtGui/qsurface.h>
#include <QtGui/qscreen.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qsurfaceformat.h>
#include <QtGui/qregion.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qmargins.h>
#include <QtCore/qpoint.h>
#include <QtGui/qicon.h>
#include <QtGui/qaccessible.h>
#include <QtGui/qcursor.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qwindow.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QWindow__Go : public QWindow {
public:
	long long __goId;
	unsigned char __goOn[32] = {};
	template<typename... A> QWindow__Go(long long id, A&&... a) : QWindow(std::forward<A>(a)...), __goId(id) {}
	using QWindow::sender;
	using QWindow::senderSignalIndex;
	using QWindow::receivers;
	using QWindow::isSignalConnected;
	void cppgoBase_exposeEvent(QExposeEvent* __a0) { QWindow::exposeEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QWindow::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QWindow::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QWindow::moveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QWindow::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QWindow::focusOutEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QWindow::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QWindow::hideEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QWindow::closeEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QWindow::event(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QWindow::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QWindow::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QWindow::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QWindow::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QWindow::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QWindow::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QWindow::wheelEvent(__a0); }
	void cppgoBase_touchEvent(QTouchEvent* __a0) { QWindow::touchEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QWindow::tabletEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QWindow::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QWindow::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QWindow::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QWindow::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QWindow::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QWindow::metaObject();
	}
	QSurface::SurfaceType surfaceType() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 7339588379958472532LL, &__h); if (__h) return (QSurface::SurfaceType)(__r); }
		return QWindow::surfaceType();
	}
	QSurfaceFormat format() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7267508590720539762LL, &__h); if (__h) return *(QSurfaceFormat*)__r; }
		return QWindow::format();
	}
	QSize size() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5614464919154503228LL, &__h); if (__h) return *(QSize*)__r; }
		return QWindow::size();
	}
	QAccessibleInterface* accessibleRoot() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -2255663269888600317LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QWindow::accessibleRoot();
	}
	QObject* focusObject() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 8913927178009815944LL, &__h); if (__h) return (QObject*)__r; }
		return QWindow::focusObject();
	}
	void exposeEvent(QExposeEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3239118852587300383LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::exposeEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::resizeEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::paintEvent(arg1);
	}
	void moveEvent(QMoveEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::moveEvent(arg1);
	}
	void focusInEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::focusInEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::focusOutEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::showEvent(arg1);
	}
	void hideEvent(QHideEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::hideEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::closeEvent(arg1);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QWindow::event(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::keyPressEvent(arg1);
	}
	void keyReleaseEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::keyReleaseEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::mouseDoubleClickEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::mouseMoveEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::wheelEvent(arg1);
	}
	void touchEvent(QTouchEvent* arg1) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2204038732330958816LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::touchEvent(arg1);
	}
	void tabletEvent(QTabletEvent* arg1) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)arg1, &__h); if (__h) return; }
		QWindow::tabletEvent(arg1);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QWindow::nativeEvent(eventType, message, result);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QWindow::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QWindow::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QWindow::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QWindow::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QWindow::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QWindow::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QWindow;
static std::mutex __connsMu_QWindow;

extern "C" {

void* QWindow_base(void* self) { return (void*)static_cast<QObject*>((QWindow*)self); }
void QWindow__set_override(void* self, int idx) { ((QWindow__Go*)self)->__goOn[idx] = 1; }
void* QWindow_fromQObject(void* o) { return (void*)qobject_cast<QWindow*>((QObject*)o); }
void* QWindow_new(long long __goid) {
	return new QWindow__Go(__goid);
}
void* QWindow_new2(long long __goid, void* parent) {
	return new QWindow__Go(__goid, (QWindow*)parent);
}
void* QWindow_new3(long long __goid, void* screen) {
	return new QWindow__Go(__goid, (QScreen*)screen);
}
char* QWindow_tr(const char* s, long long* __out_len) {
	auto&& __r = QWindow::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWindow_setSurfaceType(void* self, int surfaceType) {
	((QWindow*)self)->setSurfaceType((QSurface::SurfaceType)(surfaceType));
}
int QWindow_surfaceType(void* self) {
	auto&& __r = ((QWindow*)self)->surfaceType();
	return (int)(__r);
}
int QWindow_isVisible(void* self) {
	auto&& __r = ((QWindow*)self)->isVisible();
	return __r;
}
int QWindow_visibility(void* self) {
	auto&& __r = ((QWindow*)self)->visibility();
	return (int)(__r);
}
void QWindow_setVisibility(void* self, int v) {
	((QWindow*)self)->setVisibility((QWindow::Visibility)(v));
}
void QWindow_create(void* self) {
	((QWindow*)self)->create();
}
unsigned long long QWindow_winId(void* self) {
	auto&& __r = ((QWindow*)self)->winId();
	return __r;
}
void* QWindow_parent(void* self) {
	auto&& __r = ((QWindow*)self)->parent();
	return (void*)(__r);
}
void QWindow_setParent(void* self, void* parent) {
	((QWindow*)self)->setParent((QWindow*)parent);
}
int QWindow_isTopLevel(void* self) {
	auto&& __r = ((QWindow*)self)->isTopLevel();
	return __r;
}
int QWindow_isModal(void* self) {
	auto&& __r = ((QWindow*)self)->isModal();
	return __r;
}
int QWindow_modality(void* self) {
	auto&& __r = ((QWindow*)self)->modality();
	return (int)(__r);
}
void QWindow_setModality(void* self, int modality) {
	((QWindow*)self)->setModality((Qt::WindowModality)(modality));
}
void QWindow_setFormat(void* self, void* format) {
	((QWindow*)self)->setFormat(*(QSurfaceFormat*)format);
}
void* QWindow_format(void* self) {
	auto&& __r = ((QWindow*)self)->format();
	return new QSurfaceFormat(__r);
}
void* QWindow_requestedFormat(void* self) {
	auto&& __r = ((QWindow*)self)->requestedFormat();
	return new QSurfaceFormat(__r);
}
void QWindow_setFlags(void* self, int flags) {
	((QWindow*)self)->setFlags(QFlags<Qt::WindowType>(flags));
}
int QWindow_flags(void* self) {
	auto&& __r = ((QWindow*)self)->flags();
	return (int)(__r);
}
void QWindow_setFlag(void* self, int arg1) {
	((QWindow*)self)->setFlag((Qt::WindowType)(arg1));
}
int QWindow_type(void* self) {
	auto&& __r = ((QWindow*)self)->type();
	return (int)(__r);
}
char* QWindow_title(void* self, long long* __out_len) {
	auto&& __r = ((QWindow*)self)->title();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWindow_setOpacity(void* self, double level) {
	((QWindow*)self)->setOpacity(level);
}
double QWindow_opacity(void* self) {
	auto&& __r = ((QWindow*)self)->opacity();
	return __r;
}
void QWindow_setMask(void* self, void* region) {
	((QWindow*)self)->setMask(*(QRegion*)region);
}
void* QWindow_mask(void* self) {
	auto&& __r = ((QWindow*)self)->mask();
	return new QRegion(__r);
}
int QWindow_isActive(void* self) {
	auto&& __r = ((QWindow*)self)->isActive();
	return __r;
}
void QWindow_reportContentOrientationChange(void* self, int orientation) {
	((QWindow*)self)->reportContentOrientationChange((Qt::ScreenOrientation)(orientation));
}
int QWindow_contentOrientation(void* self) {
	auto&& __r = ((QWindow*)self)->contentOrientation();
	return (int)(__r);
}
double QWindow_devicePixelRatio(void* self) {
	auto&& __r = ((QWindow*)self)->devicePixelRatio();
	return __r;
}
int QWindow_windowState(void* self) {
	auto&& __r = ((QWindow*)self)->windowState();
	return (int)(__r);
}
int QWindow_windowStates(void* self) {
	auto&& __r = ((QWindow*)self)->windowStates();
	return (int)(__r);
}
void QWindow_setWindowState(void* self, int state) {
	((QWindow*)self)->setWindowState((Qt::WindowState)(state));
}
void QWindow_setWindowStates(void* self, int states) {
	((QWindow*)self)->setWindowStates(QFlags<Qt::WindowState>(states));
}
void QWindow_setTransientParent(void* self, void* parent) {
	((QWindow*)self)->setTransientParent((QWindow*)parent);
}
void* QWindow_transientParent(void* self) {
	auto&& __r = ((QWindow*)self)->transientParent();
	return (void*)(__r);
}
int QWindow_isAncestorOf(void* self, void* child) {
	auto&& __r = ((QWindow*)self)->isAncestorOf((QWindow*)child);
	return __r;
}
int QWindow_isExposed(void* self) {
	auto&& __r = ((QWindow*)self)->isExposed();
	return __r;
}
int QWindow_minimumWidth(void* self) {
	auto&& __r = ((QWindow*)self)->minimumWidth();
	return __r;
}
int QWindow_minimumHeight(void* self) {
	auto&& __r = ((QWindow*)self)->minimumHeight();
	return __r;
}
int QWindow_maximumWidth(void* self) {
	auto&& __r = ((QWindow*)self)->maximumWidth();
	return __r;
}
int QWindow_maximumHeight(void* self) {
	auto&& __r = ((QWindow*)self)->maximumHeight();
	return __r;
}
void* QWindow_minimumSize(void* self) {
	auto&& __r = ((QWindow*)self)->minimumSize();
	return new QSize(__r);
}
void* QWindow_maximumSize(void* self) {
	auto&& __r = ((QWindow*)self)->maximumSize();
	return new QSize(__r);
}
void* QWindow_baseSize(void* self) {
	auto&& __r = ((QWindow*)self)->baseSize();
	return new QSize(__r);
}
void* QWindow_sizeIncrement(void* self) {
	auto&& __r = ((QWindow*)self)->sizeIncrement();
	return new QSize(__r);
}
void QWindow_setMinimumSize(void* self, void* size) {
	((QWindow*)self)->setMinimumSize(*(QSize*)size);
}
void QWindow_setMaximumSize(void* self, void* size) {
	((QWindow*)self)->setMaximumSize(*(QSize*)size);
}
void QWindow_setBaseSize(void* self, void* size) {
	((QWindow*)self)->setBaseSize(*(QSize*)size);
}
void QWindow_setSizeIncrement(void* self, void* size) {
	((QWindow*)self)->setSizeIncrement(*(QSize*)size);
}
void* QWindow_geometry(void* self) {
	auto&& __r = ((QWindow*)self)->geometry();
	return new QRect(__r);
}
void* QWindow_frameMargins(void* self) {
	auto&& __r = ((QWindow*)self)->frameMargins();
	return new QMargins(__r);
}
void* QWindow_frameGeometry(void* self) {
	auto&& __r = ((QWindow*)self)->frameGeometry();
	return new QRect(__r);
}
void* QWindow_framePosition(void* self) {
	auto&& __r = ((QWindow*)self)->framePosition();
	return new QPoint(__r);
}
void QWindow_setFramePosition(void* self, void* point) {
	((QWindow*)self)->setFramePosition(*(QPoint*)point);
}
void* QWindow_safeAreaMargins(void* self) {
	auto&& __r = ((QWindow*)self)->safeAreaMargins();
	return new QMargins(__r);
}
int QWindow_width(void* self) {
	auto&& __r = ((QWindow*)self)->width();
	return __r;
}
int QWindow_height(void* self) {
	auto&& __r = ((QWindow*)self)->height();
	return __r;
}
int QWindow_x(void* self) {
	auto&& __r = ((QWindow*)self)->x();
	return __r;
}
int QWindow_y(void* self) {
	auto&& __r = ((QWindow*)self)->y();
	return __r;
}
void* QWindow_size(void* self) {
	auto&& __r = ((QWindow*)self)->size();
	return new QSize(__r);
}
void* QWindow_position(void* self) {
	auto&& __r = ((QWindow*)self)->position();
	return new QPoint(__r);
}
void QWindow_setPosition(void* self, void* pt) {
	((QWindow*)self)->setPosition(*(QPoint*)pt);
}
void QWindow_setPosition2(void* self, int posx, int posy) {
	((QWindow*)self)->setPosition(posx, posy);
}
void QWindow_resize(void* self, void* newSize) {
	((QWindow*)self)->resize(*(QSize*)newSize);
}
void QWindow_resize2(void* self, int w, int h) {
	((QWindow*)self)->resize(w, h);
}
void QWindow_setFilePath(void* self, const char* filePath, long long filePath_len) {
	QString filePath_qstr = QString::fromUtf8(filePath, filePath_len);
	((QWindow*)self)->setFilePath(filePath_qstr);
}
char* QWindow_filePath(void* self, long long* __out_len) {
	auto&& __r = ((QWindow*)self)->filePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWindow_setIcon(void* self, void* icon) {
	((QWindow*)self)->setIcon(*(QIcon*)icon);
}
void* QWindow_icon(void* self) {
	auto&& __r = ((QWindow*)self)->icon();
	return new QIcon(__r);
}
void QWindow_destroy(void* self) {
	((QWindow*)self)->destroy();
}
int QWindow_setKeyboardGrabEnabled(void* self, int grab) {
	auto&& __r = ((QWindow*)self)->setKeyboardGrabEnabled(grab);
	return __r;
}
int QWindow_setMouseGrabEnabled(void* self, int grab) {
	auto&& __r = ((QWindow*)self)->setMouseGrabEnabled(grab);
	return __r;
}
void* QWindow_screen(void* self) {
	auto&& __r = ((QWindow*)self)->screen();
	return (void*)(__r);
}
void QWindow_setScreen(void* self, void* screen) {
	((QWindow*)self)->setScreen((QScreen*)screen);
}
void* QWindow_accessibleRoot(void* self) {
	auto&& __r = ((QWindow*)self)->accessibleRoot();
	return (void*)(__r);
}
void* QWindow_focusObject(void* self) {
	auto&& __r = ((QWindow*)self)->focusObject();
	return (void*)(__r);
}
void* QWindow_mapToGlobal(void* self, void* pos) {
	auto&& __r = ((QWindow*)self)->mapToGlobal(*(QPointF*)pos);
	return new QPointF(__r);
}
void* QWindow_mapFromGlobal(void* self, void* pos) {
	auto&& __r = ((QWindow*)self)->mapFromGlobal(*(QPointF*)pos);
	return new QPointF(__r);
}
void* QWindow_mapToGlobal2(void* self, void* pos) {
	auto&& __r = ((QWindow*)self)->mapToGlobal(*(QPoint*)pos);
	return new QPoint(__r);
}
void* QWindow_mapFromGlobal2(void* self, void* pos) {
	auto&& __r = ((QWindow*)self)->mapFromGlobal(*(QPoint*)pos);
	return new QPoint(__r);
}
void* QWindow_cursor(void* self) {
	auto&& __r = ((QWindow*)self)->cursor();
	return new QCursor(__r);
}
void QWindow_setCursor(void* self, void* arg1) {
	((QWindow*)self)->setCursor(*(QCursor*)arg1);
}
void QWindow_unsetCursor(void* self) {
	((QWindow*)self)->unsetCursor();
}
void* QWindow_fromWinId(unsigned long long id) {
	auto&& __r = QWindow::fromWinId(id);
	return (void*)(__r);
}
void QWindow_requestActivate(void* self) {
	((QWindow*)self)->requestActivate();
}
void QWindow_setVisible(void* self, int visible) {
	((QWindow*)self)->setVisible(visible);
}
void QWindow_show(void* self) {
	((QWindow*)self)->show();
}
void QWindow_hide(void* self) {
	((QWindow*)self)->hide();
}
void QWindow_showMinimized(void* self) {
	((QWindow*)self)->showMinimized();
}
void QWindow_showMaximized(void* self) {
	((QWindow*)self)->showMaximized();
}
void QWindow_showFullScreen(void* self) {
	((QWindow*)self)->showFullScreen();
}
void QWindow_showNormal(void* self) {
	((QWindow*)self)->showNormal();
}
int QWindow_close(void* self) {
	auto&& __r = ((QWindow*)self)->close();
	return __r;
}
void QWindow_raise(void* self) {
	((QWindow*)self)->raise();
}
void QWindow_lower(void* self) {
	((QWindow*)self)->lower();
}
int QWindow_startSystemResize(void* self, int edges) {
	auto&& __r = ((QWindow*)self)->startSystemResize(QFlags<Qt::Edge>(edges));
	return __r;
}
int QWindow_startSystemMove(void* self) {
	auto&& __r = ((QWindow*)self)->startSystemMove();
	return __r;
}
void QWindow_setTitle(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWindow*)self)->setTitle(arg1_qstr);
}
void QWindow_setX(void* self, int arg) {
	((QWindow*)self)->setX(arg);
}
void QWindow_setY(void* self, int arg) {
	((QWindow*)self)->setY(arg);
}
void QWindow_setWidth(void* self, int arg) {
	((QWindow*)self)->setWidth(arg);
}
void QWindow_setHeight(void* self, int arg) {
	((QWindow*)self)->setHeight(arg);
}
void QWindow_setGeometry(void* self, int posx, int posy, int w, int h) {
	((QWindow*)self)->setGeometry(posx, posy, w, h);
}
void QWindow_setGeometry2(void* self, void* rect) {
	((QWindow*)self)->setGeometry(*(QRect*)rect);
}
void QWindow_setMinimumWidth(void* self, int w) {
	((QWindow*)self)->setMinimumWidth(w);
}
void QWindow_setMinimumHeight(void* self, int h) {
	((QWindow*)self)->setMinimumHeight(h);
}
void QWindow_setMaximumWidth(void* self, int w) {
	((QWindow*)self)->setMaximumWidth(w);
}
void QWindow_setMaximumHeight(void* self, int h) {
	((QWindow*)self)->setMaximumHeight(h);
}
void QWindow_alert(void* self, int msec) {
	((QWindow*)self)->alert(msec);
}
void QWindow_requestUpdate(void* self) {
	((QWindow*)self)->requestUpdate();
}
void QWindow_screenChanged(void* self, void* screen) {
	((QWindow*)self)->screenChanged((QScreen*)screen);
}
void QWindow_modalityChanged(void* self, int modality) {
	((QWindow*)self)->modalityChanged((Qt::WindowModality)(modality));
}
void QWindow_flagsChanged(void* self, int flags) {
	((QWindow*)self)->flagsChanged(QFlags<Qt::WindowType>(flags));
}
void QWindow_windowStateChanged(void* self, int windowState) {
	((QWindow*)self)->windowStateChanged((Qt::WindowState)(windowState));
}
void QWindow_windowTitleChanged(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QWindow*)self)->windowTitleChanged(title_qstr);
}
void QWindow_xChanged(void* self, int arg) {
	((QWindow*)self)->xChanged(arg);
}
void QWindow_yChanged(void* self, int arg) {
	((QWindow*)self)->yChanged(arg);
}
void QWindow_widthChanged(void* self, int arg) {
	((QWindow*)self)->widthChanged(arg);
}
void QWindow_heightChanged(void* self, int arg) {
	((QWindow*)self)->heightChanged(arg);
}
void QWindow_minimumWidthChanged(void* self, int arg) {
	((QWindow*)self)->minimumWidthChanged(arg);
}
void QWindow_minimumHeightChanged(void* self, int arg) {
	((QWindow*)self)->minimumHeightChanged(arg);
}
void QWindow_maximumWidthChanged(void* self, int arg) {
	((QWindow*)self)->maximumWidthChanged(arg);
}
void QWindow_maximumHeightChanged(void* self, int arg) {
	((QWindow*)self)->maximumHeightChanged(arg);
}
void QWindow_safeAreaMarginsChanged(void* self, void* arg) {
	((QWindow*)self)->safeAreaMarginsChanged(*(QMargins*)arg);
}
void QWindow_visibleChanged(void* self, int arg) {
	((QWindow*)self)->visibleChanged(arg);
}
void QWindow_visibilityChanged(void* self, int visibility) {
	((QWindow*)self)->visibilityChanged((QWindow::Visibility)(visibility));
}
void QWindow_activeChanged(void* self) {
	((QWindow*)self)->activeChanged();
}
void QWindow_contentOrientationChanged(void* self, int orientation) {
	((QWindow*)self)->contentOrientationChanged((Qt::ScreenOrientation)(orientation));
}
void QWindow_focusObjectChanged(void* self, void* object) {
	((QWindow*)self)->focusObjectChanged((QObject*)object);
}
void QWindow_opacityChanged(void* self, double opacity) {
	((QWindow*)self)->opacityChanged(opacity);
}
void QWindow_transientParentChanged(void* self, void* transientParent) {
	((QWindow*)self)->transientParentChanged((QWindow*)transientParent);
}
void* QWindow_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QWindow::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QWindow_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QWindow_sender(void* self) {
	auto&& __r = ((QWindow__Go*)self)->sender();
	return (void*)(__r);
}
int QWindow_senderSignalIndex(void* self) {
	auto&& __r = ((QWindow__Go*)self)->senderSignalIndex();
	return __r;
}
int QWindow_receivers(void* self, const char* signal) {
	auto&& __r = ((QWindow__Go*)self)->receivers(signal);
	return __r;
}
int QWindow_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QWindow__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
int QWindow_surfaceClass(void* self) {
	auto&& __r = ((QWindow*)self)->surfaceClass();
	return (int)(__r);
}
int QWindow_supportsOpenGL(void* self) {
	auto&& __r = ((QWindow*)self)->supportsOpenGL();
	return __r;
}
char* QWindow_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QWindow::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWindow_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QWindow::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QWindow_parent2(void* self, int mode) {
	auto&& __r = ((QWindow*)self)->parent((QWindow::AncestorMode)(mode));
	return (void*)(__r);
}
void QWindow_setFlag2(void* self, int arg1, int on) {
	((QWindow*)self)->setFlag((Qt::WindowType)(arg1), on);
}
int QWindow_isAncestorOf2(void* self, void* child, int mode) {
	auto&& __r = ((QWindow*)self)->isAncestorOf((QWindow*)child, (QWindow::AncestorMode)(mode));
	return __r;
}
void* QWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QWindow::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QWindow_asQSurface(void* self) {
	auto&& __r = static_cast<QSurface*>((QWindow*)self);
	return (void*)(__r);
}
int QWindow_surfaceType_default(void* self) {
	auto&& __r = ((QWindow*)self)->QWindow::surfaceType();
	return (int)(__r);
}
void* QWindow_format_default(void* self) {
	auto&& __r = ((QWindow*)self)->QWindow::format();
	return new QSurfaceFormat(__r);
}
void* QWindow_size_default(void* self) {
	auto&& __r = ((QWindow*)self)->QWindow::size();
	return new QSize(__r);
}
void* QWindow_accessibleRoot_default(void* self) {
	auto&& __r = ((QWindow*)self)->QWindow::accessibleRoot();
	return (void*)(__r);
}
void* QWindow_focusObject_default(void* self) {
	auto&& __r = ((QWindow*)self)->QWindow::focusObject();
	return (void*)(__r);
}
void* QWindow_metaObject_default(void* self) {
	auto&& __r = ((QWindow*)self)->QWindow::metaObject();
	return (void*)(__r);
}
void QWindow_exposeEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_exposeEvent((QExposeEvent*)arg1);
}
void QWindow_resizeEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QWindow_paintEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QWindow_moveEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)arg1);
}
void QWindow_focusInEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)arg1);
}
void QWindow_focusOutEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
void QWindow_showEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QWindow_hideEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_hideEvent((QHideEvent*)arg1);
}
void QWindow_closeEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
int QWindow_event_default(void* self, void* arg1) {
	auto&& __r = ((QWindow__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QWindow_keyPressEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QWindow_keyReleaseEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)arg1);
}
void QWindow_mousePressEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QWindow_mouseReleaseEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QWindow_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QWindow_mouseMoveEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QWindow_wheelEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QWindow_touchEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_touchEvent((QTouchEvent*)arg1);
}
void QWindow_tabletEvent_default(void* self, void* arg1) {
	((QWindow__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)arg1);
}
int QWindow_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QWindow*)self)->QWindow::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QWindow_timerEvent_default(void* self, void* event) {
	((QWindow__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QWindow_childEvent_default(void* self, void* event) {
	((QWindow__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QWindow_customEvent_default(void* self, void* event) {
	((QWindow__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QWindow_connectNotify_default(void* self, void* signal) {
	((QWindow__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QWindow_disconnectNotify_default(void* self, void* signal) {
	((QWindow__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QWindow_connect_screenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::screenChanged, (QWindow*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_modalityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::modalityChanged, (QWindow*)self, [__id](Qt::WindowModality modality){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(modality));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_flagsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::flagsChanged, (QWindow*)self, [__id](QFlags<Qt::WindowType> flags){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(flags));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_windowStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::windowStateChanged, (QWindow*)self, [__id](Qt::WindowState windowState){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(windowState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::windowTitleChanged, (QWindow*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::xChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::yChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::widthChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::heightChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_minimumWidthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::minimumWidthChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_minimumHeightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::minimumHeightChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_maximumWidthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::maximumWidthChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_maximumHeightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::maximumHeightChanged, (QWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_safeAreaMarginsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::safeAreaMarginsChanged, (QWindow*)self, [__id](QMargins arg){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::visibleChanged, (QWindow*)self, [__id](bool arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_visibilityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::visibilityChanged, (QWindow*)self, [__id](QWindow::Visibility visibility){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(visibility));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_activeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::activeChanged, (QWindow*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_contentOrientationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::contentOrientationChanged, (QWindow*)self, [__id](Qt::ScreenOrientation orientation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(orientation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_focusObjectChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::focusObjectChanged, (QWindow*)self, [__id](QObject* object){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)object);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::opacityChanged, (QWindow*)self, [__id](qreal opacity){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(opacity)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_transientParentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::transientParentChanged, (QWindow*)self, [__id](QWindow* transientParent){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)transientParent);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::destroyed, (QWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::destroyed, (QWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWindow*)self, &QWindow::objectNameChanged, (QWindow*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
	__conns_QWindow[__id] = __c;
}
void QWindow__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QWindow);
		auto __it = __conns_QWindow.find(id);
		if (__it == __conns_QWindow.end()) return;
		__c = __it->second;
		__conns_QWindow.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QWindow__Go_new(long long id) { return new QWindow__Go(id); }
void QWindow_delete(void* self) { delete (QWindow*)self; }

}

#include <QtGui/qoffscreensurface.h>
#include <QtCore/qobject.h>
#include <QtGui/qsurface.h>
#include <QtGui/qscreen.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qsurfaceformat.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qoffscreensurface.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QOffscreenSurface__Go : public QOffscreenSurface {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QOffscreenSurface__Go(long long id, A&&... a) : QOffscreenSurface(std::forward<A>(a)...), __goId(id) {}
	using QOffscreenSurface::sender;
	using QOffscreenSurface::senderSignalIndex;
	using QOffscreenSurface::receivers;
	using QOffscreenSurface::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QOffscreenSurface::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QOffscreenSurface::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QOffscreenSurface::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QOffscreenSurface::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QOffscreenSurface::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QOffscreenSurface::metaObject();
	}
	QSurface::SurfaceType surfaceType() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 7339588379958472532LL, &__h); if (__h) return (QSurface::SurfaceType)(__r); }
		return QOffscreenSurface::surfaceType();
	}
	QSurfaceFormat format() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7267508590720539762LL, &__h); if (__h) return *(QSurfaceFormat*)__r; }
		return QOffscreenSurface::format();
	}
	QSize size() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5614464919154503228LL, &__h); if (__h) return *(QSize*)__r; }
		return QOffscreenSurface::size();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QOffscreenSurface::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QOffscreenSurface::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QOffscreenSurface::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QOffscreenSurface::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QOffscreenSurface::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QOffscreenSurface::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QOffscreenSurface::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QOffscreenSurface;
static std::mutex __connsMu_QOffscreenSurface;

extern "C" {

void* QOffscreenSurface_base(void* self) { return (void*)static_cast<QObject*>((QOffscreenSurface*)self); }
void QOffscreenSurface__set_override(void* self, int idx) { ((QOffscreenSurface__Go*)self)->__goOn[idx] = 1; }
void* QOffscreenSurface_fromQObject(void* o) { return (void*)qobject_cast<QOffscreenSurface*>((QObject*)o); }
void* QOffscreenSurface_new(long long __goid) {
	return new QOffscreenSurface__Go(__goid);
}
void* QOffscreenSurface_new2(long long __goid, void* screen) {
	return new QOffscreenSurface__Go(__goid, (QScreen*)screen);
}
void* QOffscreenSurface_new3(long long __goid, void* screen, void* parent) {
	return new QOffscreenSurface__Go(__goid, (QScreen*)screen, (QObject*)parent);
}
char* QOffscreenSurface_tr(const char* s, long long* __out_len) {
	auto&& __r = QOffscreenSurface::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QOffscreenSurface_surfaceType(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->surfaceType();
	return (int)(__r);
}
void QOffscreenSurface_create(void* self) {
	((QOffscreenSurface*)self)->create();
}
void QOffscreenSurface_destroy(void* self) {
	((QOffscreenSurface*)self)->destroy();
}
int QOffscreenSurface_isValid(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->isValid();
	return __r;
}
void QOffscreenSurface_setFormat(void* self, void* format) {
	((QOffscreenSurface*)self)->setFormat(*(QSurfaceFormat*)format);
}
void* QOffscreenSurface_format(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->format();
	return new QSurfaceFormat(__r);
}
void* QOffscreenSurface_requestedFormat(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->requestedFormat();
	return new QSurfaceFormat(__r);
}
void* QOffscreenSurface_size(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->size();
	return new QSize(__r);
}
void* QOffscreenSurface_screen(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->screen();
	return (void*)(__r);
}
void QOffscreenSurface_setScreen(void* self, void* screen) {
	((QOffscreenSurface*)self)->setScreen((QScreen*)screen);
}
void QOffscreenSurface_screenChanged(void* self, void* screen) {
	((QOffscreenSurface*)self)->screenChanged((QScreen*)screen);
}
void* QOffscreenSurface_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QOffscreenSurface::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QOffscreenSurface_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QOffscreenSurface::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QOffscreenSurface_sender(void* self) {
	auto&& __r = ((QOffscreenSurface__Go*)self)->sender();
	return (void*)(__r);
}
int QOffscreenSurface_senderSignalIndex(void* self) {
	auto&& __r = ((QOffscreenSurface__Go*)self)->senderSignalIndex();
	return __r;
}
int QOffscreenSurface_receivers(void* self, const char* signal) {
	auto&& __r = ((QOffscreenSurface__Go*)self)->receivers(signal);
	return __r;
}
int QOffscreenSurface_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QOffscreenSurface__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
int QOffscreenSurface_surfaceClass(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->surfaceClass();
	return (int)(__r);
}
int QOffscreenSurface_supportsOpenGL(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->supportsOpenGL();
	return __r;
}
char* QOffscreenSurface_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QOffscreenSurface::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QOffscreenSurface_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QOffscreenSurface::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QOffscreenSurface_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QOffscreenSurface::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QOffscreenSurface_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QOffscreenSurface::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QOffscreenSurface_asQSurface(void* self) {
	auto&& __r = static_cast<QSurface*>((QOffscreenSurface*)self);
	return (void*)(__r);
}
int QOffscreenSurface_surfaceType_default(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->QOffscreenSurface::surfaceType();
	return (int)(__r);
}
void* QOffscreenSurface_format_default(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->QOffscreenSurface::format();
	return new QSurfaceFormat(__r);
}
void* QOffscreenSurface_size_default(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->QOffscreenSurface::size();
	return new QSize(__r);
}
void* QOffscreenSurface_metaObject_default(void* self) {
	auto&& __r = ((QOffscreenSurface*)self)->QOffscreenSurface::metaObject();
	return (void*)(__r);
}
int QOffscreenSurface_event_default(void* self, void* event) {
	auto&& __r = ((QOffscreenSurface*)self)->QOffscreenSurface::event((QEvent*)event);
	return __r;
}
int QOffscreenSurface_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QOffscreenSurface*)self)->QOffscreenSurface::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QOffscreenSurface_timerEvent_default(void* self, void* event) {
	((QOffscreenSurface__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QOffscreenSurface_childEvent_default(void* self, void* event) {
	((QOffscreenSurface__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QOffscreenSurface_customEvent_default(void* self, void* event) {
	((QOffscreenSurface__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QOffscreenSurface_connectNotify_default(void* self, void* signal) {
	((QOffscreenSurface__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QOffscreenSurface_disconnectNotify_default(void* self, void* signal) {
	((QOffscreenSurface__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QOffscreenSurface_connect_screenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOffscreenSurface*)self, &QOffscreenSurface::screenChanged, (QOffscreenSurface*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOffscreenSurface);
	__conns_QOffscreenSurface[__id] = __c;
}
void QOffscreenSurface_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOffscreenSurface*)self, &QOffscreenSurface::destroyed, (QOffscreenSurface*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOffscreenSurface);
	__conns_QOffscreenSurface[__id] = __c;
}
void QOffscreenSurface_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOffscreenSurface*)self, &QOffscreenSurface::destroyed, (QOffscreenSurface*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOffscreenSurface);
	__conns_QOffscreenSurface[__id] = __c;
}
void QOffscreenSurface_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOffscreenSurface*)self, &QOffscreenSurface::objectNameChanged, (QOffscreenSurface*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOffscreenSurface);
	__conns_QOffscreenSurface[__id] = __c;
}
void QOffscreenSurface__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QOffscreenSurface);
		auto __it = __conns_QOffscreenSurface.find(id);
		if (__it == __conns_QOffscreenSurface.end()) return;
		__c = __it->second;
		__conns_QOffscreenSurface.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QOffscreenSurface__Go_new(long long id) { return new QOffscreenSurface__Go(id); }
void QOffscreenSurface_delete(void* self) { delete (QOffscreenSurface*)self; }

}

#include <QtGui/qpixmapcache.h>
#include "gen_qpixmapcache__key.h"

extern "C" {

void* QPixmapCache__Key_new() {
	return new QPixmapCache::Key();
}
void* QPixmapCache__Key_new2(void* other) {
	return new QPixmapCache::Key(*(QPixmapCache::Key*)other);
}
int QPixmapCache__Key_OperatorEqual(void* self, void* key) {
	auto&& __r = ((QPixmapCache::Key*)self)->operator==(*(QPixmapCache::Key*)key);
	return __r;
}
int QPixmapCache__Key_OperatorNotEqual(void* self, void* key) {
	auto&& __r = ((QPixmapCache::Key*)self)->operator!=(*(QPixmapCache::Key*)key);
	return __r;
}
void* QPixmapCache__Key_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPixmapCache::Key*)self)->operator=(*(QPixmapCache::Key*)other);
	return (void*)&(__r);
}
void QPixmapCache__Key_swap(void* self, void* other) {
	((QPixmapCache::Key*)self)->swap(*(QPixmapCache::Key*)other);
}
int QPixmapCache__Key_isValid(void* self) {
	auto&& __r = ((QPixmapCache::Key*)self)->isValid();
	return __r;
}
void QPixmapCache__Key_delete(void* self) { delete (QPixmapCache::Key*)self; }

}

#include <QtGui/qundogroup.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qundostack.h>
#include <QtGui/qaction.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qundogroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QUndoGroup__Go : public QUndoGroup {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QUndoGroup__Go(long long id, A&&... a) : QUndoGroup(std::forward<A>(a)...), __goId(id) {}
	using QUndoGroup::sender;
	using QUndoGroup::senderSignalIndex;
	using QUndoGroup::receivers;
	using QUndoGroup::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QUndoGroup::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QUndoGroup::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QUndoGroup::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QUndoGroup::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QUndoGroup::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QUndoGroup::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QUndoGroup::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QUndoGroup::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QUndoGroup::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QUndoGroup::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QUndoGroup::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QUndoGroup::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QUndoGroup::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QUndoGroup;
static std::mutex __connsMu_QUndoGroup;

extern "C" {

void* QUndoGroup_base(void* self) { return (void*)static_cast<QObject*>((QUndoGroup*)self); }
void QUndoGroup__set_override(void* self, int idx) { ((QUndoGroup__Go*)self)->__goOn[idx] = 1; }
void* QUndoGroup_fromQObject(void* o) { return (void*)qobject_cast<QUndoGroup*>((QObject*)o); }
void* QUndoGroup_new(long long __goid) {
	return new QUndoGroup__Go(__goid);
}
void* QUndoGroup_new2(long long __goid, void* parent) {
	return new QUndoGroup__Go(__goid, (QObject*)parent);
}
char* QUndoGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QUndoGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUndoGroup_addStack(void* self, void* stack) {
	((QUndoGroup*)self)->addStack((QUndoStack*)stack);
}
void QUndoGroup_removeStack(void* self, void* stack) {
	((QUndoGroup*)self)->removeStack((QUndoStack*)stack);
}
void** QUndoGroup_stacks(void* self, long long* __out_len) {
	auto&& __r = ((QUndoGroup*)self)->stacks();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QUndoGroup_activeStack(void* self) {
	auto&& __r = ((QUndoGroup*)self)->activeStack();
	return (void*)(__r);
}
void* QUndoGroup_createUndoAction(void* self, void* parent) {
	auto&& __r = ((QUndoGroup*)self)->createUndoAction((QObject*)parent);
	return (void*)(__r);
}
void* QUndoGroup_createRedoAction(void* self, void* parent) {
	auto&& __r = ((QUndoGroup*)self)->createRedoAction((QObject*)parent);
	return (void*)(__r);
}
int QUndoGroup_canUndo(void* self) {
	auto&& __r = ((QUndoGroup*)self)->canUndo();
	return __r;
}
int QUndoGroup_canRedo(void* self) {
	auto&& __r = ((QUndoGroup*)self)->canRedo();
	return __r;
}
char* QUndoGroup_undoText(void* self, long long* __out_len) {
	auto&& __r = ((QUndoGroup*)self)->undoText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUndoGroup_redoText(void* self, long long* __out_len) {
	auto&& __r = ((QUndoGroup*)self)->redoText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QUndoGroup_isClean(void* self) {
	auto&& __r = ((QUndoGroup*)self)->isClean();
	return __r;
}
void QUndoGroup_undo(void* self) {
	((QUndoGroup*)self)->undo();
}
void QUndoGroup_redo(void* self) {
	((QUndoGroup*)self)->redo();
}
void QUndoGroup_setActiveStack(void* self, void* stack) {
	((QUndoGroup*)self)->setActiveStack((QUndoStack*)stack);
}
void QUndoGroup_activeStackChanged(void* self, void* stack) {
	((QUndoGroup*)self)->activeStackChanged((QUndoStack*)stack);
}
void QUndoGroup_indexChanged(void* self, int idx) {
	((QUndoGroup*)self)->indexChanged(idx);
}
void QUndoGroup_cleanChanged(void* self, int clean) {
	((QUndoGroup*)self)->cleanChanged(clean);
}
void QUndoGroup_canUndoChanged(void* self, int canUndo) {
	((QUndoGroup*)self)->canUndoChanged(canUndo);
}
void QUndoGroup_canRedoChanged(void* self, int canRedo) {
	((QUndoGroup*)self)->canRedoChanged(canRedo);
}
void QUndoGroup_undoTextChanged(void* self, const char* undoText, long long undoText_len) {
	QString undoText_qstr = QString::fromUtf8(undoText, undoText_len);
	((QUndoGroup*)self)->undoTextChanged(undoText_qstr);
}
void QUndoGroup_redoTextChanged(void* self, const char* redoText, long long redoText_len) {
	QString redoText_qstr = QString::fromUtf8(redoText, redoText_len);
	((QUndoGroup*)self)->redoTextChanged(redoText_qstr);
}
void* QUndoGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QUndoGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QUndoGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QUndoGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QUndoGroup_sender(void* self) {
	auto&& __r = ((QUndoGroup__Go*)self)->sender();
	return (void*)(__r);
}
int QUndoGroup_senderSignalIndex(void* self) {
	auto&& __r = ((QUndoGroup__Go*)self)->senderSignalIndex();
	return __r;
}
int QUndoGroup_receivers(void* self, const char* signal) {
	auto&& __r = ((QUndoGroup__Go*)self)->receivers(signal);
	return __r;
}
int QUndoGroup_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QUndoGroup__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QUndoGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QUndoGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUndoGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QUndoGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUndoGroup_createUndoAction2(void* self, void* parent, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	auto&& __r = ((QUndoGroup*)self)->createUndoAction((QObject*)parent, prefix_qstr);
	return (void*)(__r);
}
void* QUndoGroup_createRedoAction2(void* self, void* parent, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	auto&& __r = ((QUndoGroup*)self)->createRedoAction((QObject*)parent, prefix_qstr);
	return (void*)(__r);
}
void* QUndoGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QUndoGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QUndoGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QUndoGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QUndoGroup_metaObject_default(void* self) {
	auto&& __r = ((QUndoGroup*)self)->QUndoGroup::metaObject();
	return (void*)(__r);
}
int QUndoGroup_event_default(void* self, void* event) {
	auto&& __r = ((QUndoGroup*)self)->QUndoGroup::event((QEvent*)event);
	return __r;
}
int QUndoGroup_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QUndoGroup*)self)->QUndoGroup::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QUndoGroup_timerEvent_default(void* self, void* event) {
	((QUndoGroup__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QUndoGroup_childEvent_default(void* self, void* event) {
	((QUndoGroup__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QUndoGroup_customEvent_default(void* self, void* event) {
	((QUndoGroup__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QUndoGroup_connectNotify_default(void* self, void* signal) {
	((QUndoGroup__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QUndoGroup_disconnectNotify_default(void* self, void* signal) {
	((QUndoGroup__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QUndoGroup_connect_activeStackChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::activeStackChanged, (QUndoGroup*)self, [__id](QUndoStack* stack){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)stack);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_indexChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::indexChanged, (QUndoGroup*)self, [__id](int idx){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, idx);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_cleanChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::cleanChanged, (QUndoGroup*)self, [__id](bool clean){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, clean);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_canUndoChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::canUndoChanged, (QUndoGroup*)self, [__id](bool canUndo){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, canUndo);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_canRedoChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::canRedoChanged, (QUndoGroup*)self, [__id](bool canRedo){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, canRedo);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_undoTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::undoTextChanged, (QUndoGroup*)self, [__id](const QString& undoText){
		QByteArray undoText_ba = undoText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)undoText_ba.constData(), undoText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_redoTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::redoTextChanged, (QUndoGroup*)self, [__id](const QString& redoText){
		QByteArray redoText_ba = redoText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)redoText_ba.constData(), redoText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::destroyed, (QUndoGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::destroyed, (QUndoGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoGroup*)self, &QUndoGroup::objectNameChanged, (QUndoGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
	__conns_QUndoGroup[__id] = __c;
}
void QUndoGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QUndoGroup);
		auto __it = __conns_QUndoGroup.find(id);
		if (__it == __conns_QUndoGroup.end()) return;
		__c = __it->second;
		__conns_QUndoGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QUndoGroup__Go_new(long long id, void* parent) { return new QUndoGroup__Go(id, (QObject*)parent); }
void QUndoGroup_delete(void* self) { delete (QUndoGroup*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qtabbar.h>
#include <QtCore/qrect.h>
#include "gen_qstyleoptiontabbarbase.h"

extern "C" {

void* QStyleOptionTabBarBase_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionTabBarBase*)self); }
void* QStyleOptionTabBarBase_new() {
	return new QStyleOptionTabBarBase();
}
void* QStyleOptionTabBarBase_new2(void* other) {
	return new QStyleOptionTabBarBase(*(QStyleOptionTabBarBase*)other);
}
void* QStyleOptionTabBarBase_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionTabBarBase*)self)->operator=(*(QStyleOptionTabBarBase*)arg1);
	return (void*)&(__r);
}
int QStyleOptionTabBarBase_shape_get(void* self) {
	auto __r = ((QStyleOptionTabBarBase*)self)->shape;
	return (int)(__r);
}
void QStyleOptionTabBarBase_shape_set(void* self, int shape) {
	((QStyleOptionTabBarBase*)self)->shape = (QTabBar::Shape)(shape);
}
void* QStyleOptionTabBarBase_tabBarRect_get(void* self) {
	auto __r = ((QStyleOptionTabBarBase*)self)->tabBarRect;
	return new QRect(__r);
}
void QStyleOptionTabBarBase_tabBarRect_set(void* self, void* tabBarRect) {
	((QStyleOptionTabBarBase*)self)->tabBarRect = *(QRect*)tabBarRect;
}
void* QStyleOptionTabBarBase_selectedTabRect_get(void* self) {
	auto __r = ((QStyleOptionTabBarBase*)self)->selectedTabRect;
	return new QRect(__r);
}
void QStyleOptionTabBarBase_selectedTabRect_set(void* self, void* selectedTabRect) {
	((QStyleOptionTabBarBase*)self)->selectedTabRect = *(QRect*)selectedTabRect;
}
int QStyleOptionTabBarBase_documentMode_get(void* self) {
	auto __r = ((QStyleOptionTabBarBase*)self)->documentMode;
	return __r;
}
void QStyleOptionTabBarBase_documentMode_set(void* self, int documentMode) {
	((QStyleOptionTabBarBase*)self)->documentMode = documentMode;
}
void QStyleOptionTabBarBase_delete(void* self) { delete (QStyleOptionTabBarBase*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtCore/qvariant.h>
#include "gen_qstylehintreturnvariant.h"

extern "C" {

void* QStyleHintReturnVariant_base(void* self) { return (void*)static_cast<QStyleHintReturn*>((QStyleHintReturnVariant*)self); }
void* QStyleHintReturnVariant_new() {
	return new QStyleHintReturnVariant();
}
void* QStyleHintReturnVariant_new2(void* arg1) {
	return new QStyleHintReturnVariant(*(QStyleHintReturnVariant*)arg1);
}
void* QStyleHintReturnVariant_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleHintReturnVariant*)self)->operator=(*(QStyleHintReturnVariant*)arg1);
	return (void*)&(__r);
}
void* QStyleHintReturnVariant_variant_get(void* self) {
	auto __r = ((QStyleHintReturnVariant*)self)->variant;
	return new QVariant(__r);
}
void QStyleHintReturnVariant_variant_set(void* self, void* variant) {
	((QStyleHintReturnVariant*)self)->variant = *(QVariant*)variant;
}
void QStyleHintReturnVariant_delete(void* self) { delete (QStyleHintReturnVariant*)self; }

}

#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qpoint.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qbuttongroup.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
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
#include <QtCore/qlocale.h>
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
#include "gen_qpushbutton.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPushButton__Go : public QPushButton {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QPushButton__Go(long long id, A&&... a) : QPushButton(std::forward<A>(a)...), __goId(id) {}
	using QPushButton::updateMicroFocus;
	using QPushButton::create;
	using QPushButton::destroy;
	using QPushButton::focusNextChild;
	using QPushButton::focusPreviousChild;
	using QPushButton::sender;
	using QPushButton::senderSignalIndex;
	using QPushButton::receivers;
	using QPushButton::isSignalConnected;
	using QPushButton::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QPushButton::event(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QPushButton::paintEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QPushButton::keyPressEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QPushButton::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QPushButton::focusOutEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QPushButton::mouseMoveEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionButton* __a0) { QPushButton::initStyleOption(__a0); }
	bool cppgoBase_hitButton(const QPoint& __a0) { return QPushButton::hitButton(__a0); }
	void cppgoBase_checkStateSet() { QPushButton::checkStateSet(); }
	void cppgoBase_nextCheckState() { QPushButton::nextCheckState(); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QPushButton::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QPushButton::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QPushButton::mouseReleaseEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QPushButton::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPushButton::timerEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QPushButton::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QPushButton::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QPushButton::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QPushButton::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QPushButton::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QPushButton::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QPushButton::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QPushButton::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QPushButton::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QPushButton::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QPushButton::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QPushButton::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QPushButton::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QPushButton::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QPushButton::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QPushButton::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QPushButton::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QPushButton::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QPushButton::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QPushButton::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QPushButton::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QPushButton::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPushButton::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPushButton::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPushButton::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPushButton::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPushButton::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QPushButton::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QPushButton::minimumSizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QPushButton::event(e);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QPushButton::paintEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QPushButton::keyPressEvent(arg1);
	}
	void focusInEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)arg1, &__h); if (__h) return; }
		QPushButton::focusInEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QPushButton::focusOutEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QPushButton::mouseMoveEvent(arg1);
	}
	void initStyleOption(QStyleOptionButton* option) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QPushButton::initStyleOption(option);
	}
	bool hitButton(const QPoint& pos) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5736040353629178930LL, (void*)&pos, &__h); if (__h) return __r; }
		return QPushButton::hitButton(pos);
	}
	void checkStateSet() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5747806663224476200LL, &__h); if (__h) return; }
		QPushButton::checkStateSet();
	}
	void nextCheckState() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1474647914901213797LL, &__h); if (__h) return; }
		QPushButton::nextCheckState();
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QPushButton::keyReleaseEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QPushButton::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QPushButton::mouseReleaseEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QPushButton::changeEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QPushButton::timerEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QPushButton::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QPushButton::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QPushButton::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QPushButton::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QPushButton::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QPushButton::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QPushButton::wheelEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QPushButton::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QPushButton::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QPushButton::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QPushButton::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QPushButton::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QPushButton::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QPushButton::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QPushButton::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QPushButton::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QPushButton::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QPushButton::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QPushButton::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QPushButton::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QPushButton::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QPushButton::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QPushButton::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QPushButton::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QPushButton::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QPushButton::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QPushButton::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QPushButton::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QPushButton::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPushButton::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPushButton::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPushButton::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPushButton::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPushButton::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPushButton;
static std::mutex __connsMu_QPushButton;

extern "C" {

void* QPushButton_base(void* self) { return (void*)static_cast<QAbstractButton*>((QPushButton*)self); }
void QPushButton_chain(void* self, void** out) {
	QPushButton* __p = (QPushButton*)self;
	QAbstractButton* __b0 = static_cast<QAbstractButton*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QPushButton__set_override(void* self, int idx) { ((QPushButton__Go*)self)->__goOn[idx] = 1; }
void* QPushButton_fromQObject(void* o) { return (void*)qobject_cast<QPushButton*>((QObject*)o); }
void* QPushButton_new(long long __goid, void* parent) {
	return new QPushButton__Go(__goid, (QWidget*)parent);
}
void* QPushButton_new2(long long __goid) {
	return new QPushButton__Go(__goid);
}
void* QPushButton_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QPushButton__Go(__goid, text_qstr);
}
void* QPushButton_new4(long long __goid, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QPushButton__Go(__goid, *(QIcon*)icon, text_qstr);
}
void* QPushButton_new5(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QPushButton__Go(__goid, text_qstr, (QWidget*)parent);
}
void* QPushButton_new6(long long __goid, void* icon, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QPushButton__Go(__goid, *(QIcon*)icon, text_qstr, (QWidget*)parent);
}
char* QPushButton_tr(const char* s, long long* __out_len) {
	auto&& __r = QPushButton::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPushButton_autoDefault(void* self) {
	auto&& __r = ((QPushButton*)self)->autoDefault();
	return __r;
}
void QPushButton_setAutoDefault(void* self, int arg1) {
	((QPushButton*)self)->setAutoDefault(arg1);
}
int QPushButton_isDefault(void* self) {
	auto&& __r = ((QPushButton*)self)->isDefault();
	return __r;
}
void QPushButton_setDefault(void* self, int arg1) {
	((QPushButton*)self)->setDefault(arg1);
}
void QPushButton_setMenu(void* self, void* menu) {
	((QPushButton*)self)->setMenu((QMenu*)menu);
}
void* QPushButton_menu(void* self) {
	auto&& __r = ((QPushButton*)self)->menu();
	return (void*)(__r);
}
void QPushButton_setFlat(void* self, int arg1) {
	((QPushButton*)self)->setFlat(arg1);
}
int QPushButton_isFlat(void* self) {
	auto&& __r = ((QPushButton*)self)->isFlat();
	return __r;
}
void QPushButton_showMenu(void* self) {
	((QPushButton*)self)->showMenu();
}
void QPushButton_setTabOrder(void* arg1, void* arg2) {
	QPushButton::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QPushButton_mouseGrabber() {
	auto&& __r = QPushButton::mouseGrabber();
	return (void*)(__r);
}
void* QPushButton_keyboardGrabber() {
	auto&& __r = QPushButton::keyboardGrabber();
	return (void*)(__r);
}
void* QPushButton_find(unsigned long long arg1) {
	auto&& __r = QPushButton::find(arg1);
	return (void*)(__r);
}
void* QPushButton_createWindowContainer(void* window) {
	auto&& __r = QPushButton::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QPushButton_updateMicroFocus(void* self) {
	((QPushButton__Go*)self)->updateMicroFocus();
}
void QPushButton_create(void* self) {
	((QPushButton__Go*)self)->create();
}
void QPushButton_destroy(void* self) {
	((QPushButton__Go*)self)->destroy();
}
int QPushButton_focusNextChild(void* self) {
	auto&& __r = ((QPushButton__Go*)self)->focusNextChild();
	return __r;
}
int QPushButton_focusPreviousChild(void* self) {
	auto&& __r = ((QPushButton__Go*)self)->focusPreviousChild();
	return __r;
}
void* QPushButton_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPushButton::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPushButton_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPushButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPushButton_sender(void* self) {
	auto&& __r = ((QPushButton__Go*)self)->sender();
	return (void*)(__r);
}
int QPushButton_senderSignalIndex(void* self) {
	auto&& __r = ((QPushButton__Go*)self)->senderSignalIndex();
	return __r;
}
int QPushButton_receivers(void* self, const char* signal) {
	auto&& __r = ((QPushButton__Go*)self)->receivers(signal);
	return __r;
}
int QPushButton_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPushButton__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QPushButton_devicePixelRatioFScale() {
	auto&& __r = QPushButton::devicePixelRatioFScale();
	return __r;
}
int QPushButton_encodeMetricF(int metric, double value) {
	auto&& __r = QPushButton::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QPushButton_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QPushButton__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QPushButton_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPushButton::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPushButton_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPushButton::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPushButton_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QPushButton::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QPushButton_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QPushButton::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QPushButton_updateMicroFocus2(void* self, int query) {
	((QPushButton__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QPushButton_create2(void* self, unsigned long long arg1) {
	((QPushButton__Go*)self)->create(arg1);
}
void QPushButton_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QPushButton__Go*)self)->create(arg1, initializeWindow);
}
void QPushButton_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QPushButton__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QPushButton_destroy2(void* self, int destroyWindow) {
	((QPushButton__Go*)self)->destroy(destroyWindow);
}
void QPushButton_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QPushButton__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QPushButton_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPushButton::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPushButton_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPushButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPushButton_metaObject_default(void* self) {
	auto&& __r = ((QPushButton*)self)->QPushButton::metaObject();
	return (void*)(__r);
}
void* QPushButton_sizeHint_default(void* self) {
	auto&& __r = ((QPushButton*)self)->QPushButton::sizeHint();
	return new QSize(__r);
}
void* QPushButton_minimumSizeHint_default(void* self) {
	auto&& __r = ((QPushButton*)self)->QPushButton::minimumSizeHint();
	return new QSize(__r);
}
int QPushButton_event_default(void* self, void* e) {
	auto&& __r = ((QPushButton__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QPushButton_paintEvent_default(void* self, void* arg1) {
	((QPushButton__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QPushButton_keyPressEvent_default(void* self, void* arg1) {
	((QPushButton__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QPushButton_focusInEvent_default(void* self, void* arg1) {
	((QPushButton__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)arg1);
}
void QPushButton_focusOutEvent_default(void* self, void* arg1) {
	((QPushButton__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
void QPushButton_mouseMoveEvent_default(void* self, void* arg1) {
	((QPushButton__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QPushButton_initStyleOption_default(void* self, void* option) {
	((QPushButton__Go*)self)->cppgoBase_initStyleOption((QStyleOptionButton*)option);
}
int QPushButton_hitButton_default(void* self, void* pos) {
	auto&& __r = ((QPushButton__Go*)self)->cppgoBase_hitButton(*(QPoint*)pos);
	return __r;
}
void QPushButton_checkStateSet_default(void* self) {
	((QPushButton__Go*)self)->cppgoBase_checkStateSet();
}
void QPushButton_nextCheckState_default(void* self) {
	((QPushButton__Go*)self)->cppgoBase_nextCheckState();
}
void QPushButton_keyReleaseEvent_default(void* self, void* e) {
	((QPushButton__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QPushButton_mousePressEvent_default(void* self, void* e) {
	((QPushButton__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QPushButton_mouseReleaseEvent_default(void* self, void* e) {
	((QPushButton__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QPushButton_changeEvent_default(void* self, void* e) {
	((QPushButton__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QPushButton_timerEvent_default(void* self, void* e) {
	((QPushButton__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
int QPushButton_devType_default(void* self) {
	auto&& __r = ((QPushButton*)self)->QPushButton::devType();
	return __r;
}
void QPushButton_setVisible_default(void* self, int visible) {
	((QPushButton*)self)->QPushButton::setVisible(visible);
}
int QPushButton_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QPushButton*)self)->QPushButton::heightForWidth(arg1);
	return __r;
}
int QPushButton_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QPushButton*)self)->QPushButton::hasHeightForWidth();
	return __r;
}
void* QPushButton_paintEngine_default(void* self) {
	auto&& __r = ((QPushButton*)self)->QPushButton::paintEngine();
	return (void*)(__r);
}
void QPushButton_mouseDoubleClickEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QPushButton_wheelEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QPushButton_enterEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QPushButton_leaveEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QPushButton_moveEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QPushButton_resizeEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QPushButton_closeEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QPushButton_contextMenuEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QPushButton_tabletEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QPushButton_actionEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QPushButton_dragEnterEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QPushButton_dragMoveEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QPushButton_dragLeaveEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QPushButton_dropEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QPushButton_showEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QPushButton_hideEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QPushButton_metric_default(void* self, int arg1) {
	auto&& __r = ((QPushButton__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QPushButton_initPainter_default(void* self, void* painter) {
	((QPushButton__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QPushButton_redirected_default(void* self, void* offset) {
	auto&& __r = ((QPushButton__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QPushButton_sharedPainter_default(void* self) {
	auto&& __r = ((QPushButton__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QPushButton_inputMethodEvent_default(void* self, void* arg1) {
	((QPushButton__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QPushButton_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QPushButton*)self)->QPushButton::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QPushButton_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QPushButton__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QPushButton_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPushButton*)self)->QPushButton::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPushButton_childEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPushButton_customEvent_default(void* self, void* event) {
	((QPushButton__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPushButton_connectNotify_default(void* self, void* signal) {
	((QPushButton__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPushButton_disconnectNotify_default(void* self, void* signal) {
	((QPushButton__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPushButton_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::pressed, (QPushButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_released(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::released, (QPushButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::clicked, (QPushButton*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::clicked, (QPushButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::toggled, (QPushButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::windowTitleChanged, (QPushButton*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::windowIconChanged, (QPushButton*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::windowIconTextChanged, (QPushButton*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::customContextMenuRequested, (QPushButton*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::destroyed, (QPushButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::destroyed, (QPushButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPushButton*)self, &QPushButton::objectNameChanged, (QPushButton*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
	__conns_QPushButton[__id] = __c;
}
void QPushButton__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPushButton);
		auto __it = __conns_QPushButton.find(id);
		if (__it == __conns_QPushButton.end()) return;
		__c = __it->second;
		__conns_QPushButton.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPushButton__Go_new(long long id) { return new QPushButton__Go(id); }
void QPushButton_delete(void* self) { delete (QPushButton*)self; }

}

#include <QtWidgets/qgesture.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtapgesture.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTapGesture__Go : public QTapGesture {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QTapGesture__Go(long long id, A&&... a) : QTapGesture(std::forward<A>(a)...), __goId(id) {}
	using QTapGesture::sender;
	using QTapGesture::senderSignalIndex;
	using QTapGesture::receivers;
	using QTapGesture::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTapGesture::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTapGesture::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTapGesture::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTapGesture::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTapGesture::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTapGesture::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTapGesture::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTapGesture::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTapGesture::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTapGesture::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTapGesture::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTapGesture::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTapGesture::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTapGesture;
static std::mutex __connsMu_QTapGesture;

extern "C" {

void* QTapGesture_base(void* self) { return (void*)static_cast<QGesture*>((QTapGesture*)self); }
void QTapGesture_chain(void* self, void** out) {
	QTapGesture* __p = (QTapGesture*)self;
	QGesture* __b0 = static_cast<QGesture*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QTapGesture__set_override(void* self, int idx) { ((QTapGesture__Go*)self)->__goOn[idx] = 1; }
void* QTapGesture_fromQObject(void* o) { return (void*)qobject_cast<QTapGesture*>((QObject*)o); }
void* QTapGesture_new(long long __goid) {
	return new QTapGesture__Go(__goid);
}
void* QTapGesture_new2(long long __goid, void* parent) {
	return new QTapGesture__Go(__goid, (QObject*)parent);
}
char* QTapGesture_tr(const char* s, long long* __out_len) {
	auto&& __r = QTapGesture::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTapGesture_position(void* self) {
	auto&& __r = ((QTapGesture*)self)->position();
	return new QPointF(__r);
}
void QTapGesture_setPosition(void* self, void* pos) {
	((QTapGesture*)self)->setPosition(*(QPointF*)pos);
}
void* QTapGesture_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTapGesture::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTapGesture_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTapGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTapGesture_sender(void* self) {
	auto&& __r = ((QTapGesture__Go*)self)->sender();
	return (void*)(__r);
}
int QTapGesture_senderSignalIndex(void* self) {
	auto&& __r = ((QTapGesture__Go*)self)->senderSignalIndex();
	return __r;
}
int QTapGesture_receivers(void* self, const char* signal) {
	auto&& __r = ((QTapGesture__Go*)self)->receivers(signal);
	return __r;
}
int QTapGesture_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTapGesture__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTapGesture_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTapGesture::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTapGesture_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTapGesture::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTapGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTapGesture::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTapGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTapGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTapGesture_metaObject_default(void* self) {
	auto&& __r = ((QTapGesture*)self)->QTapGesture::metaObject();
	return (void*)(__r);
}
int QTapGesture_event_default(void* self, void* event) {
	auto&& __r = ((QTapGesture*)self)->QTapGesture::event((QEvent*)event);
	return __r;
}
int QTapGesture_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTapGesture*)self)->QTapGesture::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTapGesture_timerEvent_default(void* self, void* event) {
	((QTapGesture__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTapGesture_childEvent_default(void* self, void* event) {
	((QTapGesture__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTapGesture_customEvent_default(void* self, void* event) {
	((QTapGesture__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTapGesture_connectNotify_default(void* self, void* signal) {
	((QTapGesture__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTapGesture_disconnectNotify_default(void* self, void* signal) {
	((QTapGesture__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTapGesture_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTapGesture*)self, &QTapGesture::destroyed, (QTapGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTapGesture);
	__conns_QTapGesture[__id] = __c;
}
void QTapGesture_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTapGesture*)self, &QTapGesture::destroyed, (QTapGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTapGesture);
	__conns_QTapGesture[__id] = __c;
}
void QTapGesture_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTapGesture*)self, &QTapGesture::objectNameChanged, (QTapGesture*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTapGesture);
	__conns_QTapGesture[__id] = __c;
}
void QTapGesture__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTapGesture);
		auto __it = __conns_QTapGesture.find(id);
		if (__it == __conns_QTapGesture.end()) return;
		__c = __it->second;
		__conns_QTapGesture.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTapGesture__Go_new(long long id, void* parent) { return new QTapGesture__Go(id, (QObject*)parent); }
void QTapGesture_delete(void* self) { delete (QTapGesture*)self; }

}

#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qbrush.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qpoint.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsopacityeffect.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsOpacityEffect__Go : public QGraphicsOpacityEffect {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QGraphicsOpacityEffect__Go(long long id, A&&... a) : QGraphicsOpacityEffect(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsOpacityEffect::updateBoundingRect;
	using QGraphicsOpacityEffect::sourceIsPixmap;
	using QGraphicsOpacityEffect::sourceBoundingRect;
	using QGraphicsOpacityEffect::drawSource;
	using QGraphicsOpacityEffect::sourcePixmap;
	using QGraphicsOpacityEffect::sender;
	using QGraphicsOpacityEffect::senderSignalIndex;
	using QGraphicsOpacityEffect::receivers;
	using QGraphicsOpacityEffect::isSignalConnected;
	void cppgoBase_draw(QPainter* __a0) { QGraphicsOpacityEffect::draw(__a0); }
	void cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> __a0) { QGraphicsOpacityEffect::sourceChanged(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsOpacityEffect::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsOpacityEffect::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsOpacityEffect::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsOpacityEffect::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsOpacityEffect::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsOpacityEffect::metaObject();
	}
	void draw(QPainter* painter) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1044652152830004333LL, (void*)painter, &__h); if (__h) return; }
		QGraphicsOpacityEffect::draw(painter);
	}
	QRectF boundingRectFor(const QRectF& sourceRect) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 5831744872130794138LL, (void*)&sourceRect, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsOpacityEffect::boundingRectFor(sourceRect);
	}
	void sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> flags) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4520368087989481956LL, (int)(flags), &__h); if (__h) return; }
		QGraphicsOpacityEffect::sourceChanged(flags);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsOpacityEffect::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsOpacityEffect::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsOpacityEffect::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsOpacityEffect::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsOpacityEffect::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsOpacityEffect::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsOpacityEffect::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsOpacityEffect;
static std::mutex __connsMu_QGraphicsOpacityEffect;

extern "C" {

void* QGraphicsOpacityEffect_base(void* self) { return (void*)static_cast<QGraphicsEffect*>((QGraphicsOpacityEffect*)self); }
void QGraphicsOpacityEffect_chain(void* self, void** out) {
	QGraphicsOpacityEffect* __p = (QGraphicsOpacityEffect*)self;
	QGraphicsEffect* __b0 = static_cast<QGraphicsEffect*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsOpacityEffect__set_override(void* self, int idx) { ((QGraphicsOpacityEffect__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsOpacityEffect_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsOpacityEffect*>((QObject*)o); }
void* QGraphicsOpacityEffect_new(long long __goid) {
	return new QGraphicsOpacityEffect__Go(__goid);
}
void* QGraphicsOpacityEffect_new2(long long __goid, void* parent) {
	return new QGraphicsOpacityEffect__Go(__goid, (QObject*)parent);
}
char* QGraphicsOpacityEffect_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsOpacityEffect::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
double QGraphicsOpacityEffect_opacity(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect*)self)->opacity();
	return __r;
}
void* QGraphicsOpacityEffect_opacityMask(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect*)self)->opacityMask();
	return new QBrush(__r);
}
void QGraphicsOpacityEffect_setOpacity(void* self, double opacity) {
	((QGraphicsOpacityEffect*)self)->setOpacity(opacity);
}
void QGraphicsOpacityEffect_setOpacityMask(void* self, void* mask) {
	((QGraphicsOpacityEffect*)self)->setOpacityMask(*(QBrush*)mask);
}
void QGraphicsOpacityEffect_opacityChanged(void* self, double opacity) {
	((QGraphicsOpacityEffect*)self)->opacityChanged(opacity);
}
void QGraphicsOpacityEffect_opacityMaskChanged(void* self, void* mask) {
	((QGraphicsOpacityEffect*)self)->opacityMaskChanged(*(QBrush*)mask);
}
void QGraphicsOpacityEffect_updateBoundingRect(void* self) {
	((QGraphicsOpacityEffect__Go*)self)->updateBoundingRect();
}
int QGraphicsOpacityEffect_sourceIsPixmap(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourceIsPixmap();
	return __r;
}
void* QGraphicsOpacityEffect_sourceBoundingRect(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourceBoundingRect();
	return new QRectF(__r);
}
void QGraphicsOpacityEffect_drawSource(void* self, void* painter) {
	((QGraphicsOpacityEffect__Go*)self)->drawSource((QPainter*)painter);
}
void* QGraphicsOpacityEffect_sourcePixmap(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourcePixmap();
	return new QPixmap(__r);
}
void* QGraphicsOpacityEffect_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsOpacityEffect::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsOpacityEffect_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsOpacityEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsOpacityEffect_sender(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsOpacityEffect_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsOpacityEffect_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsOpacityEffect_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsOpacityEffect_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsOpacityEffect::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsOpacityEffect_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsOpacityEffect::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsOpacityEffect_sourceBoundingRect2(void* self, int system) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourceBoundingRect((Qt::CoordinateSystem)(system));
	return new QRectF(__r);
}
void* QGraphicsOpacityEffect_sourcePixmap2(void* self, int system) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system));
	return new QPixmap(__r);
}
void* QGraphicsOpacityEffect_sourcePixmap3(void* self, int system, void* offset) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset);
	return new QPixmap(__r);
}
void* QGraphicsOpacityEffect_sourcePixmap4(void* self, int system, void* offset, int mode) {
	auto&& __r = ((QGraphicsOpacityEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset, (QGraphicsEffect::PixmapPadMode)(mode));
	return new QPixmap(__r);
}
void* QGraphicsOpacityEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsOpacityEffect::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsOpacityEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsOpacityEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsOpacityEffect_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsOpacityEffect*)self)->QGraphicsOpacityEffect::metaObject();
	return (void*)(__r);
}
void QGraphicsOpacityEffect_draw_default(void* self, void* painter) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_draw((QPainter*)painter);
}
void* QGraphicsOpacityEffect_boundingRectFor_default(void* self, void* sourceRect) {
	auto&& __r = ((QGraphicsOpacityEffect*)self)->QGraphicsOpacityEffect::boundingRectFor(*(QRectF*)sourceRect);
	return new QRectF(__r);
}
void QGraphicsOpacityEffect_sourceChanged_default(void* self, int flags) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag>(flags));
}
int QGraphicsOpacityEffect_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsOpacityEffect*)self)->QGraphicsOpacityEffect::event((QEvent*)event);
	return __r;
}
int QGraphicsOpacityEffect_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsOpacityEffect*)self)->QGraphicsOpacityEffect::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsOpacityEffect_timerEvent_default(void* self, void* event) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsOpacityEffect_childEvent_default(void* self, void* event) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsOpacityEffect_customEvent_default(void* self, void* event) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsOpacityEffect_connectNotify_default(void* self, void* signal) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsOpacityEffect_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsOpacityEffect__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsOpacityEffect_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsOpacityEffect*)self, &QGraphicsOpacityEffect::opacityChanged, (QGraphicsOpacityEffect*)self, [__id](qreal opacity){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(opacity)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
	__conns_QGraphicsOpacityEffect[__id] = __c;
}
void QGraphicsOpacityEffect_connect_opacityMaskChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsOpacityEffect*)self, &QGraphicsOpacityEffect::opacityMaskChanged, (QGraphicsOpacityEffect*)self, [__id](const QBrush& mask){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&mask);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
	__conns_QGraphicsOpacityEffect[__id] = __c;
}
void QGraphicsOpacityEffect_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsOpacityEffect*)self, &QGraphicsOpacityEffect::enabledChanged, (QGraphicsOpacityEffect*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
	__conns_QGraphicsOpacityEffect[__id] = __c;
}
void QGraphicsOpacityEffect_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsOpacityEffect*)self, &QGraphicsOpacityEffect::destroyed, (QGraphicsOpacityEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
	__conns_QGraphicsOpacityEffect[__id] = __c;
}
void QGraphicsOpacityEffect_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsOpacityEffect*)self, &QGraphicsOpacityEffect::destroyed, (QGraphicsOpacityEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
	__conns_QGraphicsOpacityEffect[__id] = __c;
}
void QGraphicsOpacityEffect_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsOpacityEffect*)self, &QGraphicsOpacityEffect::objectNameChanged, (QGraphicsOpacityEffect*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
	__conns_QGraphicsOpacityEffect[__id] = __c;
}
void QGraphicsOpacityEffect__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsOpacityEffect);
		auto __it = __conns_QGraphicsOpacityEffect.find(id);
		if (__it == __conns_QGraphicsOpacityEffect.end()) return;
		__c = __it->second;
		__conns_QGraphicsOpacityEffect.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsOpacityEffect__Go_new(long long id, void* parent) { return new QGraphicsOpacityEffect__Go(id, (QObject*)parent); }
void QGraphicsOpacityEffect_delete(void* self) { delete (QGraphicsOpacityEffect*)self; }

}

#include <QtWidgets/qinputdialog.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qlineedit.h>
#include <QtCore/qobject.h>
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
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qinputdialog.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QInputDialog__Go : public QInputDialog {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QInputDialog__Go(long long id, A&&... a) : QInputDialog(std::forward<A>(a)...), __goId(id) {}
	using QInputDialog::adjustPosition;
	using QInputDialog::updateMicroFocus;
	using QInputDialog::create;
	using QInputDialog::destroy;
	using QInputDialog::focusNextChild;
	using QInputDialog::focusPreviousChild;
	using QInputDialog::sender;
	using QInputDialog::senderSignalIndex;
	using QInputDialog::receivers;
	using QInputDialog::isSignalConnected;
	using QInputDialog::getDecodedMetricF;
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QInputDialog::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QInputDialog::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QInputDialog::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QInputDialog::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QInputDialog::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QInputDialog::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QInputDialog::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QInputDialog::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QInputDialog::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QInputDialog::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QInputDialog::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QInputDialog::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QInputDialog::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QInputDialog::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QInputDialog::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QInputDialog::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QInputDialog::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QInputDialog::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QInputDialog::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QInputDialog::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QInputDialog::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QInputDialog::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QInputDialog::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QInputDialog::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QInputDialog::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QInputDialog::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QInputDialog::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QInputDialog::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QInputDialog::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QInputDialog::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QInputDialog::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QInputDialog::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QInputDialog::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QInputDialog::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QInputDialog::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QInputDialog::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QInputDialog::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QInputDialog::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QInputDialog::metaObject();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QInputDialog::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QInputDialog::sizeHint();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QInputDialog::setVisible(visible);
	}
	void done(int result) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, result, &__h); if (__h) return; }
		QInputDialog::done(result);
	}
	int exec() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QInputDialog::exec();
	}
	void accept() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QInputDialog::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QInputDialog::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::resizeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QInputDialog::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QInputDialog::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QInputDialog::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QInputDialog::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QInputDialog::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QInputDialog::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QInputDialog::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QInputDialog::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QInputDialog::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QInputDialog::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QInputDialog::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QInputDialog::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QInputDialog::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QInputDialog::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QInputDialog::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QInputDialog::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QInputDialog::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QInputDialog;
static std::mutex __connsMu_QInputDialog;

extern "C" {

void* QInputDialog_base(void* self) { return (void*)static_cast<QDialog*>((QInputDialog*)self); }
void QInputDialog_chain(void* self, void** out) {
	QInputDialog* __p = (QInputDialog*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QInputDialog__set_override(void* self, int idx) { ((QInputDialog__Go*)self)->__goOn[idx] = 1; }
void* QInputDialog_fromQObject(void* o) { return (void*)qobject_cast<QInputDialog*>((QObject*)o); }
void* QInputDialog_new(long long __goid, void* parent) {
	return new QInputDialog__Go(__goid, (QWidget*)parent);
}
void* QInputDialog_new2(long long __goid) {
	return new QInputDialog__Go(__goid);
}
void* QInputDialog_new3(long long __goid, void* parent, int flags) {
	return new QInputDialog__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
char* QInputDialog_tr(const char* s, long long* __out_len) {
	auto&& __r = QInputDialog::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QInputDialog_setInputMode(void* self, int mode) {
	((QInputDialog*)self)->setInputMode((QInputDialog::InputMode)(mode));
}
int QInputDialog_inputMode(void* self) {
	auto&& __r = ((QInputDialog*)self)->inputMode();
	return (int)(__r);
}
void QInputDialog_setLabelText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QInputDialog*)self)->setLabelText(text_qstr);
}
char* QInputDialog_labelText(void* self, long long* __out_len) {
	auto&& __r = ((QInputDialog*)self)->labelText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QInputDialog_setOption(void* self, int option) {
	((QInputDialog*)self)->setOption((QInputDialog::InputDialogOption)(option));
}
int QInputDialog_testOption(void* self, int option) {
	auto&& __r = ((QInputDialog*)self)->testOption((QInputDialog::InputDialogOption)(option));
	return __r;
}
void QInputDialog_setOptions(void* self, int options) {
	((QInputDialog*)self)->setOptions(QFlags<QInputDialog::InputDialogOption>(options));
}
int QInputDialog_options(void* self) {
	auto&& __r = ((QInputDialog*)self)->options();
	return (int)(__r);
}
void QInputDialog_setTextValue(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QInputDialog*)self)->setTextValue(text_qstr);
}
char* QInputDialog_textValue(void* self, long long* __out_len) {
	auto&& __r = ((QInputDialog*)self)->textValue();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QInputDialog_setTextEchoMode(void* self, int mode) {
	((QInputDialog*)self)->setTextEchoMode((QLineEdit::EchoMode)(mode));
}
int QInputDialog_textEchoMode(void* self) {
	auto&& __r = ((QInputDialog*)self)->textEchoMode();
	return (int)(__r);
}
void QInputDialog_setComboBoxEditable(void* self, int editable) {
	((QInputDialog*)self)->setComboBoxEditable(editable);
}
int QInputDialog_isComboBoxEditable(void* self) {
	auto&& __r = ((QInputDialog*)self)->isComboBoxEditable();
	return __r;
}
void QInputDialog_setComboBoxItems(void* self, char** items, long long items_len) {
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	((QInputDialog*)self)->setComboBoxItems(items_list);
}
char** QInputDialog_comboBoxItems(void* self, long long* __out_len) {
	auto&& __r = ((QInputDialog*)self)->comboBoxItems();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QInputDialog_setIntValue(void* self, int value) {
	((QInputDialog*)self)->setIntValue(value);
}
int QInputDialog_intValue(void* self) {
	auto&& __r = ((QInputDialog*)self)->intValue();
	return __r;
}
void QInputDialog_setIntMinimum(void* self, int min) {
	((QInputDialog*)self)->setIntMinimum(min);
}
int QInputDialog_intMinimum(void* self) {
	auto&& __r = ((QInputDialog*)self)->intMinimum();
	return __r;
}
void QInputDialog_setIntMaximum(void* self, int max) {
	((QInputDialog*)self)->setIntMaximum(max);
}
int QInputDialog_intMaximum(void* self) {
	auto&& __r = ((QInputDialog*)self)->intMaximum();
	return __r;
}
void QInputDialog_setIntRange(void* self, int min, int max) {
	((QInputDialog*)self)->setIntRange(min, max);
}
void QInputDialog_setIntStep(void* self, int step) {
	((QInputDialog*)self)->setIntStep(step);
}
int QInputDialog_intStep(void* self) {
	auto&& __r = ((QInputDialog*)self)->intStep();
	return __r;
}
void QInputDialog_setDoubleValue(void* self, double value) {
	((QInputDialog*)self)->setDoubleValue(value);
}
double QInputDialog_doubleValue(void* self) {
	auto&& __r = ((QInputDialog*)self)->doubleValue();
	return __r;
}
void QInputDialog_setDoubleMinimum(void* self, double min) {
	((QInputDialog*)self)->setDoubleMinimum(min);
}
double QInputDialog_doubleMinimum(void* self) {
	auto&& __r = ((QInputDialog*)self)->doubleMinimum();
	return __r;
}
void QInputDialog_setDoubleMaximum(void* self, double max) {
	((QInputDialog*)self)->setDoubleMaximum(max);
}
double QInputDialog_doubleMaximum(void* self) {
	auto&& __r = ((QInputDialog*)self)->doubleMaximum();
	return __r;
}
void QInputDialog_setDoubleRange(void* self, double min, double max) {
	((QInputDialog*)self)->setDoubleRange(min, max);
}
void QInputDialog_setDoubleDecimals(void* self, int decimals) {
	((QInputDialog*)self)->setDoubleDecimals(decimals);
}
int QInputDialog_doubleDecimals(void* self) {
	auto&& __r = ((QInputDialog*)self)->doubleDecimals();
	return __r;
}
void QInputDialog_setOkButtonText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QInputDialog*)self)->setOkButtonText(text_qstr);
}
char* QInputDialog_okButtonText(void* self, long long* __out_len) {
	auto&& __r = ((QInputDialog*)self)->okButtonText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QInputDialog_setCancelButtonText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QInputDialog*)self)->setCancelButtonText(text_qstr);
}
char* QInputDialog_cancelButtonText(void* self, long long* __out_len) {
	auto&& __r = ((QInputDialog*)self)->cancelButtonText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QInputDialog_open(void* self, void* receiver, const char* member) {
	((QInputDialog*)self)->open((QObject*)receiver, member);
}
char* QInputDialog_getText(void* parent, const char* title, long long title_len, const char* label, long long label_len, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getText((QWidget*)parent, title_qstr, label_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getMultiLineText(void* parent, const char* title, long long title_len, const char* label, long long label_len, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getMultiLineText((QWidget*)parent, title_qstr, label_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getItem(void* parent, const char* title, long long title_len, const char* label, long long label_len, char** items, long long items_len, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	auto&& __r = QInputDialog::getItem((QWidget*)parent, title_qstr, label_qstr, items_list);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QInputDialog_getInt(void* parent, const char* title, long long title_len, const char* label, long long label_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr);
	return __r;
}
double QInputDialog_getDouble(void* parent, const char* title, long long title_len, const char* label, long long label_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr);
	return __r;
}
void QInputDialog_setDoubleStep(void* self, double step) {
	((QInputDialog*)self)->setDoubleStep(step);
}
double QInputDialog_doubleStep(void* self) {
	auto&& __r = ((QInputDialog*)self)->doubleStep();
	return __r;
}
void QInputDialog_textValueChanged(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QInputDialog*)self)->textValueChanged(text_qstr);
}
void QInputDialog_textValueSelected(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QInputDialog*)self)->textValueSelected(text_qstr);
}
void QInputDialog_intValueChanged(void* self, int value) {
	((QInputDialog*)self)->intValueChanged(value);
}
void QInputDialog_intValueSelected(void* self, int value) {
	((QInputDialog*)self)->intValueSelected(value);
}
void QInputDialog_doubleValueChanged(void* self, double value) {
	((QInputDialog*)self)->doubleValueChanged(value);
}
void QInputDialog_doubleValueSelected(void* self, double value) {
	((QInputDialog*)self)->doubleValueSelected(value);
}
void QInputDialog_adjustPosition(void* self, void* arg1) {
	((QInputDialog__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QInputDialog_setTabOrder(void* arg1, void* arg2) {
	QInputDialog::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QInputDialog_mouseGrabber() {
	auto&& __r = QInputDialog::mouseGrabber();
	return (void*)(__r);
}
void* QInputDialog_keyboardGrabber() {
	auto&& __r = QInputDialog::keyboardGrabber();
	return (void*)(__r);
}
void* QInputDialog_find(unsigned long long arg1) {
	auto&& __r = QInputDialog::find(arg1);
	return (void*)(__r);
}
void* QInputDialog_createWindowContainer(void* window) {
	auto&& __r = QInputDialog::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QInputDialog_updateMicroFocus(void* self) {
	((QInputDialog__Go*)self)->updateMicroFocus();
}
void QInputDialog_create(void* self) {
	((QInputDialog__Go*)self)->create();
}
void QInputDialog_destroy(void* self) {
	((QInputDialog__Go*)self)->destroy();
}
int QInputDialog_focusNextChild(void* self) {
	auto&& __r = ((QInputDialog__Go*)self)->focusNextChild();
	return __r;
}
int QInputDialog_focusPreviousChild(void* self) {
	auto&& __r = ((QInputDialog__Go*)self)->focusPreviousChild();
	return __r;
}
void* QInputDialog_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QInputDialog::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QInputDialog_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QInputDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QInputDialog_sender(void* self) {
	auto&& __r = ((QInputDialog__Go*)self)->sender();
	return (void*)(__r);
}
int QInputDialog_senderSignalIndex(void* self) {
	auto&& __r = ((QInputDialog__Go*)self)->senderSignalIndex();
	return __r;
}
int QInputDialog_receivers(void* self, const char* signal) {
	auto&& __r = ((QInputDialog__Go*)self)->receivers(signal);
	return __r;
}
int QInputDialog_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QInputDialog__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QInputDialog_devicePixelRatioFScale() {
	auto&& __r = QInputDialog::devicePixelRatioFScale();
	return __r;
}
int QInputDialog_encodeMetricF(int metric, double value) {
	auto&& __r = QInputDialog::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QInputDialog_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QInputDialog__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QInputDialog_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QInputDialog::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QInputDialog::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QInputDialog_setOption2(void* self, int option, int on) {
	((QInputDialog*)self)->setOption((QInputDialog::InputDialogOption)(option), on);
}
char* QInputDialog_getText2(void* parent, const char* title, long long title_len, const char* label, long long label_len, int echo, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getText((QWidget*)parent, title_qstr, label_qstr, (QLineEdit::EchoMode)(echo));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getText3(void* parent, const char* title, long long title_len, const char* label, long long label_len, int echo, const char* text, long long text_len, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QInputDialog::getText((QWidget*)parent, title_qstr, label_qstr, (QLineEdit::EchoMode)(echo), text_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getText4(void* parent, const char* title, long long title_len, const char* label, long long label_len, int echo, const char* text, long long text_len, int* ok, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getText((QWidget*)parent, title_qstr, label_qstr, (QLineEdit::EchoMode)(echo), text_qstr, &ok_local);
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getText5(void* parent, const char* title, long long title_len, const char* label, long long label_len, int echo, const char* text, long long text_len, int* ok, int flags, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getText((QWidget*)parent, title_qstr, label_qstr, (QLineEdit::EchoMode)(echo), text_qstr, &ok_local, QFlags<Qt::WindowType>(flags));
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getText6(void* parent, const char* title, long long title_len, const char* label, long long label_len, int echo, const char* text, long long text_len, int* ok, int flags, int inputMethodHints, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getText((QWidget*)parent, title_qstr, label_qstr, (QLineEdit::EchoMode)(echo), text_qstr, &ok_local, QFlags<Qt::WindowType>(flags), QFlags<Qt::InputMethodHint>(inputMethodHints));
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getMultiLineText2(void* parent, const char* title, long long title_len, const char* label, long long label_len, const char* text, long long text_len, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QInputDialog::getMultiLineText((QWidget*)parent, title_qstr, label_qstr, text_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getMultiLineText3(void* parent, const char* title, long long title_len, const char* label, long long label_len, const char* text, long long text_len, int* ok, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getMultiLineText((QWidget*)parent, title_qstr, label_qstr, text_qstr, &ok_local);
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getMultiLineText4(void* parent, const char* title, long long title_len, const char* label, long long label_len, const char* text, long long text_len, int* ok, int flags, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getMultiLineText((QWidget*)parent, title_qstr, label_qstr, text_qstr, &ok_local, QFlags<Qt::WindowType>(flags));
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getMultiLineText5(void* parent, const char* title, long long title_len, const char* label, long long label_len, const char* text, long long text_len, int* ok, int flags, int inputMethodHints, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getMultiLineText((QWidget*)parent, title_qstr, label_qstr, text_qstr, &ok_local, QFlags<Qt::WindowType>(flags), QFlags<Qt::InputMethodHint>(inputMethodHints));
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getItem2(void* parent, const char* title, long long title_len, const char* label, long long label_len, char** items, long long items_len, int current, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	auto&& __r = QInputDialog::getItem((QWidget*)parent, title_qstr, label_qstr, items_list, current);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getItem3(void* parent, const char* title, long long title_len, const char* label, long long label_len, char** items, long long items_len, int current, int editable, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	auto&& __r = QInputDialog::getItem((QWidget*)parent, title_qstr, label_qstr, items_list, current, editable);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getItem4(void* parent, const char* title, long long title_len, const char* label, long long label_len, char** items, long long items_len, int current, int editable, int* ok, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	bool ok_local{};
	auto&& __r = QInputDialog::getItem((QWidget*)parent, title_qstr, label_qstr, items_list, current, editable, &ok_local);
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getItem5(void* parent, const char* title, long long title_len, const char* label, long long label_len, char** items, long long items_len, int current, int editable, int* ok, int flags, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	bool ok_local{};
	auto&& __r = QInputDialog::getItem((QWidget*)parent, title_qstr, label_qstr, items_list, current, editable, &ok_local, QFlags<Qt::WindowType>(flags));
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDialog_getItem6(void* parent, const char* title, long long title_len, const char* label, long long label_len, char** items, long long items_len, int current, int editable, int* ok, int flags, int inputMethodHints, long long* __out_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	QList<QString> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append(QString::fromUtf8(items[__i]));
	bool ok_local{};
	auto&& __r = QInputDialog::getItem((QWidget*)parent, title_qstr, label_qstr, items_list, current, editable, &ok_local, QFlags<Qt::WindowType>(flags), QFlags<Qt::InputMethodHint>(inputMethodHints));
	*ok = ok_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QInputDialog_getInt2(void* parent, const char* title, long long title_len, const char* label, long long label_len, int value) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr, value);
	return __r;
}
int QInputDialog_getInt3(void* parent, const char* title, long long title_len, const char* label, long long label_len, int value, int minValue) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr, value, minValue);
	return __r;
}
int QInputDialog_getInt4(void* parent, const char* title, long long title_len, const char* label, long long label_len, int value, int minValue, int maxValue) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue);
	return __r;
}
int QInputDialog_getInt5(void* parent, const char* title, long long title_len, const char* label, long long label_len, int value, int minValue, int maxValue, int step) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, step);
	return __r;
}
int QInputDialog_getInt6(void* parent, const char* title, long long title_len, const char* label, long long label_len, int value, int minValue, int maxValue, int step, int* ok) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, step, &ok_local);
	*ok = ok_local;
	return __r;
}
int QInputDialog_getInt7(void* parent, const char* title, long long title_len, const char* label, long long label_len, int value, int minValue, int maxValue, int step, int* ok, int flags) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getInt((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, step, &ok_local, QFlags<Qt::WindowType>(flags));
	*ok = ok_local;
	return __r;
}
double QInputDialog_getDouble2(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value);
	return __r;
}
double QInputDialog_getDouble3(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value, double minValue) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value, minValue);
	return __r;
}
double QInputDialog_getDouble4(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value, double minValue, double maxValue) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue);
	return __r;
}
double QInputDialog_getDouble5(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value, double minValue, double maxValue, int decimals) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, decimals);
	return __r;
}
double QInputDialog_getDouble6(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value, double minValue, double maxValue, int decimals, int* ok) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, decimals, &ok_local);
	*ok = ok_local;
	return __r;
}
double QInputDialog_getDouble7(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value, double minValue, double maxValue, int decimals, int* ok, int flags) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, decimals, &ok_local, QFlags<Qt::WindowType>(flags));
	*ok = ok_local;
	return __r;
}
double QInputDialog_getDouble8(void* parent, const char* title, long long title_len, const char* label, long long label_len, double value, double minValue, double maxValue, int decimals, int* ok, int flags, double step) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString label_qstr = QString::fromUtf8(label, label_len);
	bool ok_local{};
	auto&& __r = QInputDialog::getDouble((QWidget*)parent, title_qstr, label_qstr, value, minValue, maxValue, decimals, &ok_local, QFlags<Qt::WindowType>(flags), step);
	*ok = ok_local;
	return __r;
}
void* QInputDialog_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QInputDialog::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QInputDialog_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QInputDialog::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QInputDialog_updateMicroFocus2(void* self, int query) {
	((QInputDialog__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QInputDialog_create2(void* self, unsigned long long arg1) {
	((QInputDialog__Go*)self)->create(arg1);
}
void QInputDialog_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QInputDialog__Go*)self)->create(arg1, initializeWindow);
}
void QInputDialog_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QInputDialog__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QInputDialog_destroy2(void* self, int destroyWindow) {
	((QInputDialog__Go*)self)->destroy(destroyWindow);
}
void QInputDialog_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QInputDialog__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QInputDialog_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QInputDialog::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QInputDialog_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QInputDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QInputDialog_metaObject_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::metaObject();
	return (void*)(__r);
}
void* QInputDialog_minimumSizeHint_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::minimumSizeHint();
	return new QSize(__r);
}
void* QInputDialog_sizeHint_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::sizeHint();
	return new QSize(__r);
}
void QInputDialog_setVisible_default(void* self, int visible) {
	((QInputDialog*)self)->QInputDialog::setVisible(visible);
}
void QInputDialog_done_default(void* self, int result) {
	((QInputDialog*)self)->QInputDialog::done(result);
}
int QInputDialog_exec_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::exec();
	return __r;
}
void QInputDialog_accept_default(void* self) {
	((QInputDialog*)self)->QInputDialog::accept();
}
void QInputDialog_reject_default(void* self) {
	((QInputDialog*)self)->QInputDialog::reject();
}
void QInputDialog_keyPressEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QInputDialog_closeEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QInputDialog_showEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QInputDialog_resizeEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QInputDialog_contextMenuEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QInputDialog_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QInputDialog__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QInputDialog_devType_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::devType();
	return __r;
}
int QInputDialog_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::heightForWidth(arg1);
	return __r;
}
int QInputDialog_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::hasHeightForWidth();
	return __r;
}
void* QInputDialog_paintEngine_default(void* self) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::paintEngine();
	return (void*)(__r);
}
int QInputDialog_event_default(void* self, void* event) {
	auto&& __r = ((QInputDialog__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QInputDialog_mousePressEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QInputDialog_mouseReleaseEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QInputDialog_mouseDoubleClickEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QInputDialog_mouseMoveEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QInputDialog_wheelEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QInputDialog_keyReleaseEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QInputDialog_focusInEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QInputDialog_focusOutEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QInputDialog_enterEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QInputDialog_leaveEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QInputDialog_paintEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QInputDialog_moveEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QInputDialog_tabletEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QInputDialog_actionEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QInputDialog_dragEnterEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QInputDialog_dragMoveEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QInputDialog_dragLeaveEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QInputDialog_dropEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QInputDialog_hideEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QInputDialog_changeEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QInputDialog_metric_default(void* self, int arg1) {
	auto&& __r = ((QInputDialog__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QInputDialog_initPainter_default(void* self, void* painter) {
	((QInputDialog__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QInputDialog_redirected_default(void* self, void* offset) {
	auto&& __r = ((QInputDialog__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QInputDialog_sharedPainter_default(void* self) {
	auto&& __r = ((QInputDialog__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QInputDialog_inputMethodEvent_default(void* self, void* arg1) {
	((QInputDialog__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QInputDialog_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QInputDialog*)self)->QInputDialog::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QInputDialog_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QInputDialog__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QInputDialog_timerEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QInputDialog_childEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QInputDialog_customEvent_default(void* self, void* event) {
	((QInputDialog__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QInputDialog_connectNotify_default(void* self, void* signal) {
	((QInputDialog__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QInputDialog_disconnectNotify_default(void* self, void* signal) {
	((QInputDialog__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QInputDialog_connect_textValueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::textValueChanged, (QInputDialog*)self, [__id](const QString& text){
		QByteArray text_ba = text.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)text_ba.constData(), text_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_textValueSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::textValueSelected, (QInputDialog*)self, [__id](const QString& text){
		QByteArray text_ba = text.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)text_ba.constData(), text_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_intValueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::intValueChanged, (QInputDialog*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_intValueSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::intValueSelected, (QInputDialog*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_doubleValueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::doubleValueChanged, (QInputDialog*)self, [__id](double value){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(value)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_doubleValueSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::doubleValueSelected, (QInputDialog*)self, [__id](double value){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(value)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::finished, (QInputDialog*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::accepted, (QInputDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::rejected, (QInputDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::windowTitleChanged, (QInputDialog*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::windowIconChanged, (QInputDialog*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::windowIconTextChanged, (QInputDialog*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::customContextMenuRequested, (QInputDialog*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::destroyed, (QInputDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::destroyed, (QInputDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDialog*)self, &QInputDialog::objectNameChanged, (QInputDialog*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
	__conns_QInputDialog[__id] = __c;
}
void QInputDialog__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QInputDialog);
		auto __it = __conns_QInputDialog.find(id);
		if (__it == __conns_QInputDialog.end()) return;
		__c = __it->second;
		__conns_QInputDialog.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QInputDialog__Go_new(long long id) { return new QInputDialog__Go(id); }
void QInputDialog_delete(void* self) { delete (QInputDialog*)self; }

}

#include <QtWidgets/qrhiwidget.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtGui/qimage.h>
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
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qrhiwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QRhiWidget__Go : public QRhiWidget {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QRhiWidget__Go(long long id, A&&... a) : QRhiWidget(std::forward<A>(a)...), __goId(id) {}
	using QRhiWidget::isAutoRenderTargetEnabled;
	using QRhiWidget::setAutoRenderTarget;
	using QRhiWidget::updateMicroFocus;
	using QRhiWidget::create;
	using QRhiWidget::destroy;
	using QRhiWidget::focusNextChild;
	using QRhiWidget::focusPreviousChild;
	using QRhiWidget::sender;
	using QRhiWidget::senderSignalIndex;
	using QRhiWidget::receivers;
	using QRhiWidget::isSignalConnected;
	using QRhiWidget::getDecodedMetricF;
	void cppgoBase_releaseResources() { QRhiWidget::releaseResources(); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QRhiWidget::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QRhiWidget::paintEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QRhiWidget::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QRhiWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QRhiWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QRhiWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QRhiWidget::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QRhiWidget::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QRhiWidget::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QRhiWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QRhiWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QRhiWidget::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QRhiWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QRhiWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QRhiWidget::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QRhiWidget::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QRhiWidget::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QRhiWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QRhiWidget::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QRhiWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QRhiWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QRhiWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QRhiWidget::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QRhiWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QRhiWidget::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QRhiWidget::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QRhiWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QRhiWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QRhiWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QRhiWidget::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QRhiWidget::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QRhiWidget::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QRhiWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QRhiWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QRhiWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QRhiWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QRhiWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QRhiWidget::metaObject();
	}
	void releaseResources() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -434767277154741265LL, &__h); if (__h) return; }
		QRhiWidget::releaseResources();
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QRhiWidget::resizeEvent(e);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QRhiWidget::paintEvent(e);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QRhiWidget::event(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QRhiWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QRhiWidget::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QRhiWidget::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QRhiWidget::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QRhiWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QRhiWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QRhiWidget::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QRhiWidget::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QRhiWidget::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QRhiWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QRhiWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QRhiWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QRhiWidget::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QRhiWidget::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QRhiWidget::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QRhiWidget::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QRhiWidget::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QRhiWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QRhiWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QRhiWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QRhiWidget;
static std::mutex __connsMu_QRhiWidget;

extern "C" {

void* QRhiWidget_base(void* self) { return (void*)static_cast<QWidget*>((QRhiWidget*)self); }
void QRhiWidget_chain(void* self, void** out) {
	QRhiWidget* __p = (QRhiWidget*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QRhiWidget__set_override(void* self, int idx) { ((QRhiWidget__Go*)self)->__goOn[idx] = 1; }
void* QRhiWidget_fromQObject(void* o) { return (void*)qobject_cast<QRhiWidget*>((QObject*)o); }
void* QRhiWidget_new(long long __goid, void* parent) {
	return new QRhiWidget__Go(__goid, (QWidget*)parent);
}
void* QRhiWidget_new2(long long __goid) {
	return new QRhiWidget__Go(__goid);
}
void* QRhiWidget_new3(long long __goid, void* parent, int f) {
	return new QRhiWidget__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QRhiWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QRhiWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QRhiWidget_api(void* self) {
	auto&& __r = ((QRhiWidget*)self)->api();
	return (int)(__r);
}
void QRhiWidget_setApi(void* self, int api) {
	((QRhiWidget*)self)->setApi((QRhiWidget::Api)(api));
}
int QRhiWidget_isDebugLayerEnabled(void* self) {
	auto&& __r = ((QRhiWidget*)self)->isDebugLayerEnabled();
	return __r;
}
void QRhiWidget_setDebugLayerEnabled(void* self, int enable) {
	((QRhiWidget*)self)->setDebugLayerEnabled(enable);
}
int QRhiWidget_sampleCount(void* self) {
	auto&& __r = ((QRhiWidget*)self)->sampleCount();
	return __r;
}
void QRhiWidget_setSampleCount(void* self, int samples) {
	((QRhiWidget*)self)->setSampleCount(samples);
}
int QRhiWidget_colorBufferFormat(void* self) {
	auto&& __r = ((QRhiWidget*)self)->colorBufferFormat();
	return (int)(__r);
}
void QRhiWidget_setColorBufferFormat(void* self, int format) {
	((QRhiWidget*)self)->setColorBufferFormat((QRhiWidget::TextureFormat)(format));
}
void* QRhiWidget_fixedColorBufferSize(void* self) {
	auto&& __r = ((QRhiWidget*)self)->fixedColorBufferSize();
	return new QSize(__r);
}
void QRhiWidget_setFixedColorBufferSize(void* self, void* pixelSize) {
	((QRhiWidget*)self)->setFixedColorBufferSize(*(QSize*)pixelSize);
}
void QRhiWidget_setFixedColorBufferSize2(void* self, int w, int h) {
	((QRhiWidget*)self)->setFixedColorBufferSize(w, h);
}
int QRhiWidget_isMirrorVerticallyEnabled(void* self) {
	auto&& __r = ((QRhiWidget*)self)->isMirrorVerticallyEnabled();
	return __r;
}
void QRhiWidget_setMirrorVertically(void* self, int enabled) {
	((QRhiWidget*)self)->setMirrorVertically(enabled);
}
void* QRhiWidget_grabFramebuffer(void* self) {
	auto&& __r = ((QRhiWidget*)self)->grabFramebuffer();
	return new QImage(__r);
}
int QRhiWidget_isAutoRenderTargetEnabled(void* self) {
	auto&& __r = ((QRhiWidget__Go*)self)->isAutoRenderTargetEnabled();
	return __r;
}
void QRhiWidget_setAutoRenderTarget(void* self, int enabled) {
	((QRhiWidget__Go*)self)->setAutoRenderTarget(enabled);
}
void QRhiWidget_frameSubmitted(void* self) {
	((QRhiWidget*)self)->frameSubmitted();
}
void QRhiWidget_renderFailed(void* self) {
	((QRhiWidget*)self)->renderFailed();
}
void QRhiWidget_sampleCountChanged(void* self, int samples) {
	((QRhiWidget*)self)->sampleCountChanged(samples);
}
void QRhiWidget_colorBufferFormatChanged(void* self, int format) {
	((QRhiWidget*)self)->colorBufferFormatChanged((QRhiWidget::TextureFormat)(format));
}
void QRhiWidget_fixedColorBufferSizeChanged(void* self, void* pixelSize) {
	((QRhiWidget*)self)->fixedColorBufferSizeChanged(*(QSize*)pixelSize);
}
void QRhiWidget_mirrorVerticallyChanged(void* self, int enabled) {
	((QRhiWidget*)self)->mirrorVerticallyChanged(enabled);
}
void QRhiWidget_setTabOrder(void* arg1, void* arg2) {
	QRhiWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QRhiWidget_mouseGrabber() {
	auto&& __r = QRhiWidget::mouseGrabber();
	return (void*)(__r);
}
void* QRhiWidget_keyboardGrabber() {
	auto&& __r = QRhiWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QRhiWidget_find(unsigned long long arg1) {
	auto&& __r = QRhiWidget::find(arg1);
	return (void*)(__r);
}
void* QRhiWidget_createWindowContainer(void* window) {
	auto&& __r = QRhiWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QRhiWidget_updateMicroFocus(void* self) {
	((QRhiWidget__Go*)self)->updateMicroFocus();
}
void QRhiWidget_create(void* self) {
	((QRhiWidget__Go*)self)->create();
}
void QRhiWidget_destroy(void* self) {
	((QRhiWidget__Go*)self)->destroy();
}
int QRhiWidget_focusNextChild(void* self) {
	auto&& __r = ((QRhiWidget__Go*)self)->focusNextChild();
	return __r;
}
int QRhiWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QRhiWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QRhiWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QRhiWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QRhiWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QRhiWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QRhiWidget_sender(void* self) {
	auto&& __r = ((QRhiWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QRhiWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QRhiWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QRhiWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QRhiWidget__Go*)self)->receivers(signal);
	return __r;
}
int QRhiWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QRhiWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QRhiWidget_devicePixelRatioFScale() {
	auto&& __r = QRhiWidget::devicePixelRatioFScale();
	return __r;
}
int QRhiWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QRhiWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QRhiWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QRhiWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QRhiWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QRhiWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRhiWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QRhiWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRhiWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QRhiWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QRhiWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QRhiWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QRhiWidget_updateMicroFocus2(void* self, int query) {
	((QRhiWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QRhiWidget_create2(void* self, unsigned long long arg1) {
	((QRhiWidget__Go*)self)->create(arg1);
}
void QRhiWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QRhiWidget__Go*)self)->create(arg1, initializeWindow);
}
void QRhiWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QRhiWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QRhiWidget_destroy2(void* self, int destroyWindow) {
	((QRhiWidget__Go*)self)->destroy(destroyWindow);
}
void QRhiWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QRhiWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QRhiWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QRhiWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QRhiWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QRhiWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QRhiWidget_metaObject_default(void* self) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::metaObject();
	return (void*)(__r);
}
void QRhiWidget_releaseResources_default(void* self) {
	((QRhiWidget__Go*)self)->cppgoBase_releaseResources();
}
void QRhiWidget_resizeEvent_default(void* self, void* e) {
	((QRhiWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QRhiWidget_paintEvent_default(void* self, void* e) {
	((QRhiWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
int QRhiWidget_event_default(void* self, void* e) {
	auto&& __r = ((QRhiWidget__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QRhiWidget_devType_default(void* self) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::devType();
	return __r;
}
void QRhiWidget_setVisible_default(void* self, int visible) {
	((QRhiWidget*)self)->QRhiWidget::setVisible(visible);
}
void* QRhiWidget_sizeHint_default(void* self) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::sizeHint();
	return new QSize(__r);
}
void* QRhiWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::minimumSizeHint();
	return new QSize(__r);
}
int QRhiWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::heightForWidth(arg1);
	return __r;
}
int QRhiWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::hasHeightForWidth();
	return __r;
}
void* QRhiWidget_paintEngine_default(void* self) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::paintEngine();
	return (void*)(__r);
}
void QRhiWidget_mousePressEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QRhiWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QRhiWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QRhiWidget_mouseMoveEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QRhiWidget_wheelEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QRhiWidget_keyPressEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QRhiWidget_keyReleaseEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QRhiWidget_focusInEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QRhiWidget_focusOutEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QRhiWidget_enterEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QRhiWidget_leaveEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QRhiWidget_moveEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QRhiWidget_closeEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QRhiWidget_contextMenuEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QRhiWidget_tabletEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QRhiWidget_actionEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QRhiWidget_dragEnterEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QRhiWidget_dragMoveEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QRhiWidget_dragLeaveEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QRhiWidget_dropEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QRhiWidget_showEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QRhiWidget_hideEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QRhiWidget_changeEvent_default(void* self, void* arg1) {
	((QRhiWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QRhiWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QRhiWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QRhiWidget_initPainter_default(void* self, void* painter) {
	((QRhiWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QRhiWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QRhiWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QRhiWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QRhiWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QRhiWidget_inputMethodEvent_default(void* self, void* arg1) {
	((QRhiWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QRhiWidget_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QRhiWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QRhiWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QRhiWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QRhiWidget*)self)->QRhiWidget::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QRhiWidget_timerEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QRhiWidget_childEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QRhiWidget_customEvent_default(void* self, void* event) {
	((QRhiWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QRhiWidget_connectNotify_default(void* self, void* signal) {
	((QRhiWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QRhiWidget_disconnectNotify_default(void* self, void* signal) {
	((QRhiWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QRhiWidget_connect_frameSubmitted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::frameSubmitted, (QRhiWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_renderFailed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::renderFailed, (QRhiWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_sampleCountChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::sampleCountChanged, (QRhiWidget*)self, [__id](int samples){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, samples);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_colorBufferFormatChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::colorBufferFormatChanged, (QRhiWidget*)self, [__id](QRhiWidget::TextureFormat format){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(format));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_fixedColorBufferSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::fixedColorBufferSizeChanged, (QRhiWidget*)self, [__id](const QSize& pixelSize){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pixelSize);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_mirrorVerticallyChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::mirrorVerticallyChanged, (QRhiWidget*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::windowTitleChanged, (QRhiWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::windowIconChanged, (QRhiWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::windowIconTextChanged, (QRhiWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::customContextMenuRequested, (QRhiWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::destroyed, (QRhiWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::destroyed, (QRhiWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRhiWidget*)self, &QRhiWidget::objectNameChanged, (QRhiWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
	__conns_QRhiWidget[__id] = __c;
}
void QRhiWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QRhiWidget);
		auto __it = __conns_QRhiWidget.find(id);
		if (__it == __conns_QRhiWidget.end()) return;
		__c = __it->second;
		__conns_QRhiWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QRhiWidget__Go_new(long long id) { return new QRhiWidget__Go(id); }
void QRhiWidget_delete(void* self) { delete (QRhiWidget*)self; }

}

#include <QtWidgets/qtoolbox.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qsize.h>
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
#include "gen_qtoolbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QToolBox__Go : public QToolBox {
public:
	long long __goId;
	unsigned char __goOn[51] = {};
	template<typename... A> QToolBox__Go(long long id, A&&... a) : QToolBox(std::forward<A>(a)...), __goId(id) {}
	using QToolBox::drawFrame;
	using QToolBox::updateMicroFocus;
	using QToolBox::create;
	using QToolBox::destroy;
	using QToolBox::focusNextChild;
	using QToolBox::focusPreviousChild;
	using QToolBox::sender;
	using QToolBox::senderSignalIndex;
	using QToolBox::receivers;
	using QToolBox::isSignalConnected;
	using QToolBox::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QToolBox::event(__a0); }
	void cppgoBase_itemInserted(int __a0) { QToolBox::itemInserted(__a0); }
	void cppgoBase_itemRemoved(int __a0) { QToolBox::itemRemoved(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QToolBox::showEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QToolBox::changeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QToolBox::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QToolBox::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QToolBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QToolBox::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QToolBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QToolBox::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QToolBox::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QToolBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QToolBox::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QToolBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QToolBox::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QToolBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QToolBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QToolBox::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QToolBox::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QToolBox::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QToolBox::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QToolBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QToolBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QToolBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QToolBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QToolBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QToolBox::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QToolBox::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QToolBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QToolBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QToolBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QToolBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QToolBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QToolBox::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QToolBox::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QToolBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QToolBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QToolBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QToolBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QToolBox::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QToolBox::event(e);
	}
	void itemInserted(int index) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2981295881700994698LL, index, &__h); if (__h) return; }
		QToolBox::itemInserted(index);
	}
	void itemRemoved(int index) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 5624127216289758258LL, index, &__h); if (__h) return; }
		QToolBox::itemRemoved(index);
	}
	void showEvent(QShowEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)e, &__h); if (__h) return; }
		QToolBox::showEvent(e);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QToolBox::changeEvent(arg1);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QToolBox::sizeHint();
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QToolBox::paintEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QToolBox::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QToolBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QToolBox::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QToolBox::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QToolBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QToolBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QToolBox::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QToolBox::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QToolBox::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QToolBox::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QToolBox::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QToolBox::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QToolBox::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QToolBox::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QToolBox::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QToolBox::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QToolBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QToolBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QToolBox::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QToolBox::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QToolBox::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QToolBox::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QToolBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QToolBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QToolBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QToolBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QToolBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QToolBox::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QToolBox::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QToolBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QToolBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QToolBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QToolBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QToolBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QToolBox::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QToolBox::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QToolBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QToolBox::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QToolBox::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QToolBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QToolBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QToolBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QToolBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QToolBox;
static std::mutex __connsMu_QToolBox;

extern "C" {

void* QToolBox_base(void* self) { return (void*)static_cast<QFrame*>((QToolBox*)self); }
void QToolBox_chain(void* self, void** out) {
	QToolBox* __p = (QToolBox*)self;
	QFrame* __b0 = static_cast<QFrame*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QToolBox__set_override(void* self, int idx) { ((QToolBox__Go*)self)->__goOn[idx] = 1; }
void* QToolBox_fromQObject(void* o) { return (void*)qobject_cast<QToolBox*>((QObject*)o); }
void* QToolBox_new(long long __goid, void* parent) {
	return new QToolBox__Go(__goid, (QWidget*)parent);
}
void* QToolBox_new2(long long __goid) {
	return new QToolBox__Go(__goid);
}
void* QToolBox_new3(long long __goid, void* parent, int f) {
	return new QToolBox__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QToolBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QToolBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QToolBox_addItem(void* self, void* widget, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QToolBox*)self)->addItem((QWidget*)widget, text_qstr);
	return __r;
}
int QToolBox_addItem2(void* self, void* widget, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QToolBox*)self)->addItem((QWidget*)widget, *(QIcon*)icon, text_qstr);
	return __r;
}
int QToolBox_insertItem(void* self, int index, void* widget, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QToolBox*)self)->insertItem(index, (QWidget*)widget, text_qstr);
	return __r;
}
int QToolBox_insertItem2(void* self, int index, void* widget, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QToolBox*)self)->insertItem(index, (QWidget*)widget, *(QIcon*)icon, text_qstr);
	return __r;
}
void QToolBox_removeItem(void* self, int index) {
	((QToolBox*)self)->removeItem(index);
}
void QToolBox_setItemEnabled(void* self, int index, int enabled) {
	((QToolBox*)self)->setItemEnabled(index, enabled);
}
int QToolBox_isItemEnabled(void* self, int index) {
	auto&& __r = ((QToolBox*)self)->isItemEnabled(index);
	return __r;
}
void QToolBox_setItemText(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QToolBox*)self)->setItemText(index, text_qstr);
}
char* QToolBox_itemText(void* self, int index, long long* __out_len) {
	auto&& __r = ((QToolBox*)self)->itemText(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QToolBox_setItemIcon(void* self, int index, void* icon) {
	((QToolBox*)self)->setItemIcon(index, *(QIcon*)icon);
}
void* QToolBox_itemIcon(void* self, int index) {
	auto&& __r = ((QToolBox*)self)->itemIcon(index);
	return new QIcon(__r);
}
void QToolBox_setItemToolTip(void* self, int index, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QToolBox*)self)->setItemToolTip(index, toolTip_qstr);
}
char* QToolBox_itemToolTip(void* self, int index, long long* __out_len) {
	auto&& __r = ((QToolBox*)self)->itemToolTip(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QToolBox_currentIndex(void* self) {
	auto&& __r = ((QToolBox*)self)->currentIndex();
	return __r;
}
void* QToolBox_currentWidget(void* self) {
	auto&& __r = ((QToolBox*)self)->currentWidget();
	return (void*)(__r);
}
void* QToolBox_widget(void* self, int index) {
	auto&& __r = ((QToolBox*)self)->widget(index);
	return (void*)(__r);
}
int QToolBox_indexOf(void* self, void* widget) {
	auto&& __r = ((QToolBox*)self)->indexOf((QWidget*)widget);
	return __r;
}
int QToolBox_count(void* self) {
	auto&& __r = ((QToolBox*)self)->count();
	return __r;
}
void QToolBox_setCurrentIndex(void* self, int index) {
	((QToolBox*)self)->setCurrentIndex(index);
}
void QToolBox_setCurrentWidget(void* self, void* widget) {
	((QToolBox*)self)->setCurrentWidget((QWidget*)widget);
}
void QToolBox_currentChanged(void* self, int index) {
	((QToolBox*)self)->currentChanged(index);
}
void QToolBox_drawFrame(void* self, void* arg1) {
	((QToolBox__Go*)self)->drawFrame((QPainter*)arg1);
}
void QToolBox_setTabOrder(void* arg1, void* arg2) {
	QToolBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QToolBox_mouseGrabber() {
	auto&& __r = QToolBox::mouseGrabber();
	return (void*)(__r);
}
void* QToolBox_keyboardGrabber() {
	auto&& __r = QToolBox::keyboardGrabber();
	return (void*)(__r);
}
void* QToolBox_find(unsigned long long arg1) {
	auto&& __r = QToolBox::find(arg1);
	return (void*)(__r);
}
void* QToolBox_createWindowContainer(void* window) {
	auto&& __r = QToolBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QToolBox_updateMicroFocus(void* self) {
	((QToolBox__Go*)self)->updateMicroFocus();
}
void QToolBox_create(void* self) {
	((QToolBox__Go*)self)->create();
}
void QToolBox_destroy(void* self) {
	((QToolBox__Go*)self)->destroy();
}
int QToolBox_focusNextChild(void* self) {
	auto&& __r = ((QToolBox__Go*)self)->focusNextChild();
	return __r;
}
int QToolBox_focusPreviousChild(void* self) {
	auto&& __r = ((QToolBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QToolBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QToolBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QToolBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QToolBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QToolBox_sender(void* self) {
	auto&& __r = ((QToolBox__Go*)self)->sender();
	return (void*)(__r);
}
int QToolBox_senderSignalIndex(void* self) {
	auto&& __r = ((QToolBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QToolBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QToolBox__Go*)self)->receivers(signal);
	return __r;
}
int QToolBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QToolBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QToolBox_devicePixelRatioFScale() {
	auto&& __r = QToolBox::devicePixelRatioFScale();
	return __r;
}
int QToolBox_encodeMetricF(int metric, double value) {
	auto&& __r = QToolBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QToolBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QToolBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QToolBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QToolBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QToolBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QToolBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QToolBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QToolBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QToolBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QToolBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QToolBox_updateMicroFocus2(void* self, int query) {
	((QToolBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QToolBox_create2(void* self, unsigned long long arg1) {
	((QToolBox__Go*)self)->create(arg1);
}
void QToolBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QToolBox__Go*)self)->create(arg1, initializeWindow);
}
void QToolBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QToolBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QToolBox_destroy2(void* self, int destroyWindow) {
	((QToolBox__Go*)self)->destroy(destroyWindow);
}
void QToolBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QToolBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QToolBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QToolBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QToolBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QToolBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QToolBox_metaObject_default(void* self) {
	auto&& __r = ((QToolBox*)self)->QToolBox::metaObject();
	return (void*)(__r);
}
int QToolBox_event_default(void* self, void* e) {
	auto&& __r = ((QToolBox__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QToolBox_itemInserted_default(void* self, int index) {
	((QToolBox__Go*)self)->cppgoBase_itemInserted(index);
}
void QToolBox_itemRemoved_default(void* self, int index) {
	((QToolBox__Go*)self)->cppgoBase_itemRemoved(index);
}
void QToolBox_showEvent_default(void* self, void* e) {
	((QToolBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)e);
}
void QToolBox_changeEvent_default(void* self, void* arg1) {
	((QToolBox__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void* QToolBox_sizeHint_default(void* self) {
	auto&& __r = ((QToolBox*)self)->QToolBox::sizeHint();
	return new QSize(__r);
}
void QToolBox_paintEvent_default(void* self, void* arg1) {
	((QToolBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QToolBox_initStyleOption_default(void* self, void* option) {
	((QToolBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QToolBox_devType_default(void* self) {
	auto&& __r = ((QToolBox*)self)->QToolBox::devType();
	return __r;
}
void QToolBox_setVisible_default(void* self, int visible) {
	((QToolBox*)self)->QToolBox::setVisible(visible);
}
void* QToolBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QToolBox*)self)->QToolBox::minimumSizeHint();
	return new QSize(__r);
}
int QToolBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QToolBox*)self)->QToolBox::heightForWidth(arg1);
	return __r;
}
int QToolBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QToolBox*)self)->QToolBox::hasHeightForWidth();
	return __r;
}
void* QToolBox_paintEngine_default(void* self) {
	auto&& __r = ((QToolBox*)self)->QToolBox::paintEngine();
	return (void*)(__r);
}
void QToolBox_mousePressEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QToolBox_mouseReleaseEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QToolBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QToolBox_mouseMoveEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QToolBox_wheelEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QToolBox_keyPressEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QToolBox_keyReleaseEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QToolBox_focusInEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QToolBox_focusOutEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QToolBox_enterEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QToolBox_leaveEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QToolBox_moveEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QToolBox_resizeEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QToolBox_closeEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QToolBox_contextMenuEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QToolBox_tabletEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QToolBox_actionEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QToolBox_dragEnterEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QToolBox_dragMoveEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QToolBox_dragLeaveEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QToolBox_dropEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QToolBox_hideEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QToolBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QToolBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QToolBox_initPainter_default(void* self, void* painter) {
	((QToolBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QToolBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QToolBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QToolBox_sharedPainter_default(void* self) {
	auto&& __r = ((QToolBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QToolBox_inputMethodEvent_default(void* self, void* arg1) {
	((QToolBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QToolBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QToolBox*)self)->QToolBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QToolBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QToolBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QToolBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QToolBox*)self)->QToolBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QToolBox_timerEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QToolBox_childEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QToolBox_customEvent_default(void* self, void* event) {
	((QToolBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QToolBox_connectNotify_default(void* self, void* signal) {
	((QToolBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QToolBox_disconnectNotify_default(void* self, void* signal) {
	((QToolBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QToolBox_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::currentChanged, (QToolBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::windowTitleChanged, (QToolBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::windowIconChanged, (QToolBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::windowIconTextChanged, (QToolBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::customContextMenuRequested, (QToolBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::destroyed, (QToolBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::destroyed, (QToolBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBox*)self, &QToolBox::objectNameChanged, (QToolBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
	__conns_QToolBox[__id] = __c;
}
void QToolBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QToolBox);
		auto __it = __conns_QToolBox.find(id);
		if (__it == __conns_QToolBox.end()) return;
		__c = __it->second;
		__conns_QToolBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QToolBox__Go_new(long long id) { return new QToolBox__Go(id); }
void QToolBox_delete(void* self) { delete (QToolBox*)self; }

}

