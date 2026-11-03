#include <QtCore/qchar.h>
#include "gen_qlatin1char.h"

extern "C" {

void* QLatin1Char_new(char c) {
	return new QLatin1Char(c);
}
void* QLatin1Char_new2(void* arg1) {
	return new QLatin1Char(*(QLatin1Char*)arg1);
}
char QLatin1Char_toLatin1(void* self) {
	auto&& __r = ((QLatin1Char*)self)->toLatin1();
	return __r;
}
unsigned short QLatin1Char_unicode(void* self) {
	auto&& __r = ((QLatin1Char*)self)->unicode();
	return __r;
}
int QLatin1Char_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QLatin1Char*)self) == (*(QLatin1Char*)rhs);
	return __r;
}
int QLatin1Char_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QLatin1Char*)self) != (*(QLatin1Char*)rhs);
	return __r;
}
int QLatin1Char_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QLatin1Char*)self) < (*(QLatin1Char*)rhs);
	return __r;
}
int QLatin1Char_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QLatin1Char*)self) > (*(QLatin1Char*)rhs);
	return __r;
}
int QLatin1Char_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QLatin1Char*)self) <= (*(QLatin1Char*)rhs);
	return __r;
}
int QLatin1Char_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QLatin1Char*)self) >= (*(QLatin1Char*)rhs);
	return __r;
}
void QLatin1Char_delete(void* self) { delete (QLatin1Char*)self; }

}

#include <QtCore/qmetatype.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qcompare.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qdebug.h>
#include "gen_qmetatype.h"

extern "C" {

void* QMetaType_new(int type_) {
	return new QMetaType(type_);
}
void* QMetaType_new2() {
	return new QMetaType();
}
void* QMetaType_new3(void* arg1) {
	return new QMetaType(*(QMetaType*)arg1);
}
void QMetaType_registerNormalizedTypedef(void* normalizedTypeName, void* type_) {
	QMetaType::registerNormalizedTypedef(*(QByteArray*)normalizedTypeName, *(QMetaType*)type_);
}
int QMetaType_type(const char* typeName) {
	auto&& __r = QMetaType::type(typeName);
	return __r;
}
int QMetaType_type2(void* typeName) {
	auto&& __r = QMetaType::type(*(QByteArray*)typeName);
	return __r;
}
const char* QMetaType_typeName(int type_) {
	auto&& __r = QMetaType::typeName(type_);
	return __r;
}
int QMetaType_sizeOf(int type_) {
	auto&& __r = QMetaType::sizeOf(type_);
	return __r;
}
int QMetaType_typeFlags(int type_) {
	auto&& __r = QMetaType::typeFlags(type_);
	return (int)(__r);
}
void* QMetaType_metaObjectForType(int type_) {
	auto&& __r = QMetaType::metaObjectForType(type_);
	return (void*)(__r);
}
int QMetaType_isRegistered(int type_) {
	auto&& __r = QMetaType::isRegistered(type_);
	return __r;
}
int QMetaType_isValid(void* self) {
	auto&& __r = ((QMetaType*)self)->isValid();
	return __r;
}
int QMetaType_isRegistered2(void* self) {
	auto&& __r = ((QMetaType*)self)->isRegistered();
	return __r;
}
void QMetaType_registerType(void* self) {
	((QMetaType*)self)->registerType();
}
int QMetaType_id(void* self) {
	auto&& __r = ((QMetaType*)self)->id();
	return __r;
}
long long QMetaType_sizeOf2(void* self) {
	auto&& __r = ((QMetaType*)self)->sizeOf();
	return __r;
}
long long QMetaType_alignOf(void* self) {
	auto&& __r = ((QMetaType*)self)->alignOf();
	return __r;
}
int QMetaType_flags(void* self) {
	auto&& __r = ((QMetaType*)self)->flags();
	return (int)(__r);
}
void* QMetaType_metaObject(void* self) {
	auto&& __r = ((QMetaType*)self)->metaObject();
	return (void*)(__r);
}
const char* QMetaType_name(void* self) {
	auto&& __r = ((QMetaType*)self)->name();
	return __r;
}
int QMetaType_isDefaultConstructible(void* self) {
	auto&& __r = ((QMetaType*)self)->isDefaultConstructible();
	return __r;
}
int QMetaType_isCopyConstructible(void* self) {
	auto&& __r = ((QMetaType*)self)->isCopyConstructible();
	return __r;
}
int QMetaType_isMoveConstructible(void* self) {
	auto&& __r = ((QMetaType*)self)->isMoveConstructible();
	return __r;
}
int QMetaType_isDestructible(void* self) {
	auto&& __r = ((QMetaType*)self)->isDestructible();
	return __r;
}
int QMetaType_isEqualityComparable(void* self) {
	auto&& __r = ((QMetaType*)self)->isEqualityComparable();
	return __r;
}
int QMetaType_isOrdered(void* self) {
	auto&& __r = ((QMetaType*)self)->isOrdered();
	return __r;
}
int QMetaType_hasRegisteredDataStreamOperators(void* self) {
	auto&& __r = ((QMetaType*)self)->hasRegisteredDataStreamOperators();
	return __r;
}
void* QMetaType_underlyingType(void* self) {
	auto&& __r = ((QMetaType*)self)->underlyingType();
	return new QMetaType(__r);
}
void* QMetaType_fromName(void* name) {
	auto&& __r = QMetaType::fromName(*(QByteArrayView*)name);
	return new QMetaType(__r);
}
int QMetaType_hasRegisteredDebugStreamOperator(void* self) {
	auto&& __r = ((QMetaType*)self)->hasRegisteredDebugStreamOperator();
	return __r;
}
int QMetaType_hasRegisteredDebugStreamOperator2(int typeId) {
	auto&& __r = QMetaType::hasRegisteredDebugStreamOperator(typeId);
	return __r;
}
int QMetaType_canConvert(void* fromType, void* toType) {
	auto&& __r = QMetaType::canConvert(*(QMetaType*)fromType, *(QMetaType*)toType);
	return __r;
}
int QMetaType_canView(void* fromType, void* toType) {
	auto&& __r = QMetaType::canView(*(QMetaType*)fromType, *(QMetaType*)toType);
	return __r;
}
int QMetaType_hasRegisteredConverterFunction(void* fromType, void* toType) {
	auto&& __r = QMetaType::hasRegisteredConverterFunction(*(QMetaType*)fromType, *(QMetaType*)toType);
	return __r;
}
int QMetaType_hasRegisteredMutableViewFunction(void* fromType, void* toType) {
	auto&& __r = QMetaType::hasRegisteredMutableViewFunction(*(QMetaType*)fromType, *(QMetaType*)toType);
	return __r;
}
void QMetaType_unregisterConverterFunction(void* from, void* to) {
	QMetaType::unregisterConverterFunction(*(QMetaType*)from, *(QMetaType*)to);
}
void QMetaType_unregisterMutableViewFunction(void* from, void* to) {
	QMetaType::unregisterMutableViewFunction(*(QMetaType*)from, *(QMetaType*)to);
}
void QMetaType_unregisterMetaType(void* type_) {
	QMetaType::unregisterMetaType(*(QMetaType*)type_);
}
int QMetaType_id2(void* self, int arg1) {
	auto&& __r = ((QMetaType*)self)->id(arg1);
	return __r;
}
int QMetaType_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaType*)self) == (*(QMetaType*)rhs);
	return __r;
}
int QMetaType_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaType*)self) != (*(QMetaType*)rhs);
	return __r;
}
void QMetaType_delete(void* self) { delete (QMetaType*)self; }

}

#include <QtCore/qdebug.h>
#include <QtCore/qtextstream.h>
#include "gen_qnodebug.h"

extern "C" {

void* QNoDebug_OperatorShiftLeft(void* self, void* arg1) {
	auto&& __r = ((QNoDebug*)self)->operator<<(*(QTextStreamManipulator*)arg1);
	return (void*)&(__r);
}
void* QNoDebug_space(void* self) {
	auto&& __r = ((QNoDebug*)self)->space();
	return (void*)&(__r);
}
void* QNoDebug_nospace(void* self) {
	auto&& __r = ((QNoDebug*)self)->nospace();
	return (void*)&(__r);
}
void* QNoDebug_maybeSpace(void* self) {
	auto&& __r = ((QNoDebug*)self)->maybeSpace();
	return (void*)&(__r);
}
void* QNoDebug_quote(void* self) {
	auto&& __r = ((QNoDebug*)self)->quote();
	return (void*)&(__r);
}
void* QNoDebug_noquote(void* self) {
	auto&& __r = ((QNoDebug*)self)->noquote();
	return (void*)&(__r);
}
void* QNoDebug_maybeQuote(void* self) {
	auto&& __r = ((QNoDebug*)self)->maybeQuote();
	return (void*)&(__r);
}
void* QNoDebug_verbosity(void* self, int arg1) {
	auto&& __r = ((QNoDebug*)self)->verbosity(arg1);
	return (void*)&(__r);
}
void* QNoDebug_maybeQuote2(void* self, char arg1) {
	auto&& __r = ((QNoDebug*)self)->maybeQuote(arg1);
	return (void*)&(__r);
}
void QNoDebug_delete(void* self) { delete (QNoDebug*)self; }

}

#include <QtCore/qassociativeiterable.h>
#include <QtCore/qmetacontainer.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qvariant.h>
#include "gen_qassociativeiterable.h"

extern "C" {

void* QAssociativeIterable_new() {
	return new QAssociativeIterable();
}
void* QAssociativeIterable_new2(void* arg1) {
	return new QAssociativeIterable(*(QAssociativeIterable*)arg1);
}
int QAssociativeIterable_containsKey(void* self, void* key) {
	auto&& __r = ((QAssociativeIterable*)self)->containsKey(*(QVariant*)key);
	return __r;
}
void QAssociativeIterable_insertKey(void* self, void* key) {
	((QAssociativeIterable*)self)->insertKey(*(QVariant*)key);
}
void QAssociativeIterable_removeKey(void* self, void* key) {
	((QAssociativeIterable*)self)->removeKey(*(QVariant*)key);
}
void* QAssociativeIterable_value(void* self, void* key) {
	auto&& __r = ((QAssociativeIterable*)self)->value(*(QVariant*)key);
	return new QVariant(__r);
}
void QAssociativeIterable_setValue(void* self, void* key, void* mapped) {
	((QAssociativeIterable*)self)->setValue(*(QVariant*)key, *(QVariant*)mapped);
}
int QAssociativeIterable_canRandomAccessIterate(void* self) {
	auto&& __r = ((QAssociativeIterable*)self)->canRandomAccessIterate();
	return __r;
}
int QAssociativeIterable_canInputIterate(void* self) {
	auto&& __r = ((QAssociativeIterable*)self)->canInputIterate();
	return __r;
}
int QAssociativeIterable_canForwardIterate(void* self) {
	auto&& __r = ((QAssociativeIterable*)self)->canForwardIterate();
	return __r;
}
int QAssociativeIterable_canReverseIterate(void* self) {
	auto&& __r = ((QAssociativeIterable*)self)->canReverseIterate();
	return __r;
}
long long QAssociativeIterable_size(void* self) {
	auto&& __r = ((QAssociativeIterable*)self)->size();
	return __r;
}
void QAssociativeIterable_clear(void* self) {
	((QAssociativeIterable*)self)->clear();
}
void* QAssociativeIterable_metaContainer(void* self) {
	auto&& __r = ((QAssociativeIterable*)self)->metaContainer();
	return new QMetaAssociation(__r);
}
void QAssociativeIterable_delete(void* self) { delete (QAssociativeIterable*)self; }

}

#include <QtCore/qcborvalue.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qurl.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/quuid.h>
#include <QtCore/qcborarray.h>
#include <QtCore/qcbormap.h>
#include <QtCore/qvariant.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qcborstreamwriter.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qcompare.h>
#include "gen_qcborvalueref.h"

extern "C" {

void* QCborValueRef_base(void* self) { return (void*)static_cast<QCborValueConstRef*>((QCborValueRef*)self); }
void* QCborValueRef_new(void* arg1) {
	return new QCborValueRef(*(QCborValueRef*)arg1);
}
void* QCborValueRef_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCborValueRef*)self)->operator=(*(QCborValue*)other);
	return (void*)&(__r);
}
void* QCborValueRef_OperatorAssign2(void* self, void* other) {
	auto&& __r = ((QCborValueRef*)self)->operator=(*(QCborValueRef*)other);
	return (void*)&(__r);
}
void* QCborValueRef_OperatorSubscript(void* self, long long key) {
	auto&& __r = ((QCborValueRef*)self)->operator[](key);
	return new QCborValueRef(__r);
}
void* QCborValueRef_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QCborValueRef*)self)->operator[](*(QLatin1String*)key);
	return new QCborValueRef(__r);
}
void* QCborValueRef_OperatorSubscript3(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborValueRef*)self)->operator[](key_qstr);
	return new QCborValueRef(__r);
}
void* QCborValueRef_ToQCborValue(void* self) {
	auto&& __r = static_cast<QCborValue>(*(QCborValueRef*)self);
	return new QCborValue(__r);
}
int QCborValueRef_type(void* self) {
	auto&& __r = ((QCborValueRef*)self)->type();
	return (int)(__r);
}
int QCborValueRef_isInteger(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isInteger();
	return __r;
}
int QCborValueRef_isByteArray(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isByteArray();
	return __r;
}
int QCborValueRef_isString(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isString();
	return __r;
}
int QCborValueRef_isArray(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isArray();
	return __r;
}
int QCborValueRef_isMap(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isMap();
	return __r;
}
int QCborValueRef_isTag(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isTag();
	return __r;
}
int QCborValueRef_isFalse(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isFalse();
	return __r;
}
int QCborValueRef_isTrue(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isTrue();
	return __r;
}
int QCborValueRef_isBool(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isBool();
	return __r;
}
int QCborValueRef_isNull(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isNull();
	return __r;
}
int QCborValueRef_isUndefined(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isUndefined();
	return __r;
}
int QCborValueRef_isDouble(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isDouble();
	return __r;
}
int QCborValueRef_isDateTime(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isDateTime();
	return __r;
}
int QCborValueRef_isUrl(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isUrl();
	return __r;
}
int QCborValueRef_isRegularExpression(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isRegularExpression();
	return __r;
}
int QCborValueRef_isUuid(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isUuid();
	return __r;
}
int QCborValueRef_isInvalid(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isInvalid();
	return __r;
}
int QCborValueRef_isContainer(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isContainer();
	return __r;
}
int QCborValueRef_isSimpleType(void* self) {
	auto&& __r = ((QCborValueRef*)self)->isSimpleType();
	return __r;
}
int QCborValueRef_isSimpleType2(void* self, int st) {
	auto&& __r = ((QCborValueRef*)self)->isSimpleType((QCborSimpleType)(st));
	return __r;
}
int QCborValueRef_toSimpleType(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toSimpleType();
	return (int)(__r);
}
int QCborValueRef_tag(void* self) {
	auto&& __r = ((QCborValueRef*)self)->tag();
	return (int)(__r);
}
void* QCborValueRef_taggedValue(void* self) {
	auto&& __r = ((QCborValueRef*)self)->taggedValue();
	return new QCborValue(__r);
}
long long QCborValueRef_toInteger(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toInteger();
	return __r;
}
int QCborValueRef_toBool(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toBool();
	return __r;
}
double QCborValueRef_toDouble(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toDouble();
	return __r;
}
void* QCborValueRef_toByteArray(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toByteArray();
	return new QByteArray(__r);
}
char* QCborValueRef_toString(void* self, long long* __out_len) {
	auto&& __r = ((QCborValueRef*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborValueRef_toDateTime(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toDateTime();
	return new QDateTime(__r);
}
void* QCborValueRef_toUrl(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toUrl();
	return new QUrl(__r);
}
void* QCborValueRef_toRegularExpression(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toRegularExpression();
	return new QRegularExpression(__r);
}
void* QCborValueRef_toUuid(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toUuid();
	return new QUuid(__r);
}
void* QCborValueRef_toArray(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toArray();
	return new QCborArray(__r);
}
void* QCborValueRef_toArray2(void* self, void* a) {
	auto&& __r = ((QCborValueRef*)self)->toArray(*(QCborArray*)a);
	return new QCborArray(__r);
}
void* QCborValueRef_toMap(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toMap();
	return new QCborMap(__r);
}
void* QCborValueRef_toMap2(void* self, void* m) {
	auto&& __r = ((QCborValueRef*)self)->toMap(*(QCborMap*)m);
	return new QCborMap(__r);
}
void* QCborValueRef_OperatorSubscript4(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborValueRef*)self)->operator[](key_qstr);
	return new QCborValue(__r);
}
void* QCborValueRef_OperatorSubscript5(void* self, void* key) {
	auto&& __r = ((QCborValueRef*)self)->operator[](*(QLatin1String*)key);
	return new QCborValue(__r);
}
void* QCborValueRef_OperatorSubscript6(void* self, long long key) {
	auto&& __r = ((QCborValueRef*)self)->operator[](key);
	return new QCborValue(__r);
}
int QCborValueRef_compare(void* self, void* other) {
	auto&& __r = ((QCborValueRef*)self)->compare(*(QCborValue*)other);
	return __r;
}
void* QCborValueRef_toVariant(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toVariant();
	return new QVariant(__r);
}
void* QCborValueRef_toJsonValue(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toJsonValue();
	return new QJsonValue(__r);
}
void* QCborValueRef_toCbor(void* self) {
	auto&& __r = ((QCborValueRef*)self)->toCbor();
	return new QByteArray(__r);
}
void QCborValueRef_toCbor2(void* self, void* writer) {
	((QCborValueRef*)self)->toCbor(*(QCborStreamWriter*)writer);
}
char* QCborValueRef_toDiagnosticNotation(void* self, long long* __out_len) {
	auto&& __r = ((QCborValueRef*)self)->toDiagnosticNotation();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCborValueRef_toSimpleType2(void* self, int defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toSimpleType((QCborSimpleType)(defaultValue));
	return (int)(__r);
}
int QCborValueRef_tag2(void* self, int defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->tag((QCborTag)(defaultValue));
	return (int)(__r);
}
void* QCborValueRef_taggedValue2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->taggedValue(*(QCborValue*)defaultValue);
	return new QCborValue(__r);
}
long long QCborValueRef_toInteger2(void* self, long long defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toInteger(defaultValue);
	return __r;
}
int QCborValueRef_toBool2(void* self, int defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toBool(defaultValue);
	return __r;
}
double QCborValueRef_toDouble2(void* self, double defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toDouble(defaultValue);
	return __r;
}
void* QCborValueRef_toByteArray2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toByteArray(*(QByteArray*)defaultValue);
	return new QByteArray(__r);
}
char* QCborValueRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QCborValueRef*)self)->toString(defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborValueRef_toDateTime2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toDateTime(*(QDateTime*)defaultValue);
	return new QDateTime(__r);
}
void* QCborValueRef_toUrl2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toUrl(*(QUrl*)defaultValue);
	return new QUrl(__r);
}
void* QCborValueRef_toRegularExpression2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toRegularExpression(*(QRegularExpression*)defaultValue);
	return new QRegularExpression(__r);
}
void* QCborValueRef_toUuid2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValueRef*)self)->toUuid(*(QUuid*)defaultValue);
	return new QUuid(__r);
}
void* QCborValueRef_toCbor3(void* self, int opt) {
	auto&& __r = ((QCborValueRef*)self)->toCbor(QFlags<QCborValue::EncodingOption>(opt));
	return new QByteArray(__r);
}
void QCborValueRef_toCbor4(void* self, void* writer, int opt) {
	((QCborValueRef*)self)->toCbor(*(QCborStreamWriter*)writer, QFlags<QCborValue::EncodingOption>(opt));
}
char* QCborValueRef_toDiagnosticNotation2(void* self, int opt, long long* __out_len) {
	auto&& __r = ((QCborValueRef*)self)->toDiagnosticNotation(QFlags<QCborValue::DiagnosticNotationOption>(opt));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCborValueRef_delete(void* self) { delete (QCborValueRef*)self; }

}

#include <QtCore/qdirlisting.h>
#include <QtCore/qstring.h>
#include "gen_qdirlisting.h"

extern "C" {

void* QDirListing_new(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDirListing(path_qstr);
}
void* QDirListing_new2(const char* path, long long path_len, char** nameFilters, long long nameFilters_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	return new QDirListing(path_qstr, nameFilters_list);
}
void* QDirListing_new3(const char* path, long long path_len, int flags) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDirListing(path_qstr, QFlags<QDirListing::IteratorFlag>(flags));
}
void* QDirListing_new4(const char* path, long long path_len, char** nameFilters, long long nameFilters_len, int flags) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	return new QDirListing(path_qstr, nameFilters_list, QFlags<QDirListing::IteratorFlag>(flags));
}
void QDirListing_swap(void* self, void* other) {
	((QDirListing*)self)->swap(*(QDirListing*)other);
}
char* QDirListing_iteratorPath(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing*)self)->iteratorPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDirListing_iteratorFlags(void* self) {
	auto&& __r = ((QDirListing*)self)->iteratorFlags();
	return (int)(__r);
}
char** QDirListing_nameFilters(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing*)self)->nameFilters();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QDirListing_end(void* self) {
	auto&& __r = ((QDirListing*)self)->end();
	return new QDirListing::sentinel(__r);
}
void* QDirListing_cend(void* self) {
	auto&& __r = ((QDirListing*)self)->cend();
	return new QDirListing::sentinel(__r);
}
void* QDirListing_constEnd(void* self) {
	auto&& __r = ((QDirListing*)self)->constEnd();
	return new QDirListing::sentinel(__r);
}
void QDirListing_delete(void* self) { delete (QDirListing*)self; }

}

#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qabstractitemmodel.h>
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
#include "gen_qitemselectionmodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QItemSelectionModel__Go : public QItemSelectionModel {
public:
	long long __goId;
	unsigned char __goOn[14] = {};
	template<typename... A> QItemSelectionModel__Go(long long id, A&&... a) : QItemSelectionModel(std::forward<A>(a)...), __goId(id) {}
	using QItemSelectionModel::emitSelectionChanged;
	using QItemSelectionModel::sender;
	using QItemSelectionModel::senderSignalIndex;
	using QItemSelectionModel::receivers;
	using QItemSelectionModel::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QItemSelectionModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QItemSelectionModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QItemSelectionModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QItemSelectionModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QItemSelectionModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QItemSelectionModel::metaObject();
	}
	void setCurrentIndex(const QModelIndex& index, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -6491574503686613496LL, (void*)&index, (int)(command), &__h); if (__h) return; }
		QItemSelectionModel::setCurrentIndex(index, command);
	}
	void select(const QModelIndex& index, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4794627110275867635LL, (void*)&index, (int)(command), &__h); if (__h) return; }
		QItemSelectionModel::select(index, command);
	}
	void select(const QItemSelection& selection, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8951329890086256397LL, (void*)&selection, (int)(command), &__h); if (__h) return; }
		QItemSelectionModel::select(selection, command);
	}
	void clear() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QItemSelectionModel::clear();
	}
	void reset() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QItemSelectionModel::reset();
	}
	void clearCurrentIndex() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9143260233085619705LL, &__h); if (__h) return; }
		QItemSelectionModel::clearCurrentIndex();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QItemSelectionModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QItemSelectionModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QItemSelectionModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QItemSelectionModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QItemSelectionModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QItemSelectionModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QItemSelectionModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QItemSelectionModel;
static std::mutex __connsMu_QItemSelectionModel;

extern "C" {

void* QItemSelectionModel_base(void* self) { return (void*)static_cast<QObject*>((QItemSelectionModel*)self); }
void QItemSelectionModel__set_override(void* self, int idx) { ((QItemSelectionModel__Go*)self)->__goOn[idx] = 1; }
void* QItemSelectionModel_fromQObject(void* o) { return (void*)qobject_cast<QItemSelectionModel*>((QObject*)o); }
void* QItemSelectionModel_new(long long __goid) {
	return new QItemSelectionModel__Go(__goid);
}
void* QItemSelectionModel_new2(long long __goid, void* model, void* parent) {
	return new QItemSelectionModel__Go(__goid, (QAbstractItemModel*)model, (QObject*)parent);
}
void* QItemSelectionModel_new3(long long __goid, void* model) {
	return new QItemSelectionModel__Go(__goid, (QAbstractItemModel*)model);
}
char* QItemSelectionModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QItemSelectionModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QItemSelectionModel_currentIndex(void* self) {
	auto&& __r = ((QItemSelectionModel*)self)->currentIndex();
	return new QModelIndex(__r);
}
int QItemSelectionModel_isSelected(void* self, void* index) {
	auto&& __r = ((QItemSelectionModel*)self)->isSelected(*(QModelIndex*)index);
	return __r;
}
int QItemSelectionModel_isRowSelected(void* self, int row) {
	auto&& __r = ((QItemSelectionModel*)self)->isRowSelected(row);
	return __r;
}
int QItemSelectionModel_isColumnSelected(void* self, int column) {
	auto&& __r = ((QItemSelectionModel*)self)->isColumnSelected(column);
	return __r;
}
int QItemSelectionModel_rowIntersectsSelection(void* self, int row) {
	auto&& __r = ((QItemSelectionModel*)self)->rowIntersectsSelection(row);
	return __r;
}
int QItemSelectionModel_columnIntersectsSelection(void* self, int column) {
	auto&& __r = ((QItemSelectionModel*)self)->columnIntersectsSelection(column);
	return __r;
}
int QItemSelectionModel_hasSelection(void* self) {
	auto&& __r = ((QItemSelectionModel*)self)->hasSelection();
	return __r;
}
void** QItemSelectionModel_selectedIndexes(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelectionModel*)self)->selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void** QItemSelectionModel_selectedRows(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelectionModel*)self)->selectedRows();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void** QItemSelectionModel_selectedColumns(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelectionModel*)self)->selectedColumns();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QItemSelectionModel_selection(void* self) {
	auto&& __r = ((QItemSelectionModel*)self)->selection();
	return new QItemSelection(__r);
}
void* QItemSelectionModel_model(void* self) {
	auto&& __r = ((QItemSelectionModel*)self)->model();
	return (void*)(__r);
}
void* QItemSelectionModel_model2(void* self) {
	auto&& __r = ((QItemSelectionModel*)self)->model();
	return (void*)(__r);
}
void QItemSelectionModel_setModel(void* self, void* model) {
	((QItemSelectionModel*)self)->setModel((QAbstractItemModel*)model);
}
void QItemSelectionModel_setCurrentIndex(void* self, void* index, int command) {
	((QItemSelectionModel*)self)->setCurrentIndex(*(QModelIndex*)index, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelectionModel_select(void* self, void* index, int command) {
	((QItemSelectionModel*)self)->select(*(QModelIndex*)index, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelectionModel_select2(void* self, void* selection, int command) {
	((QItemSelectionModel*)self)->select(*(QItemSelection*)selection, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelectionModel_clear(void* self) {
	((QItemSelectionModel*)self)->clear();
}
void QItemSelectionModel_reset(void* self) {
	((QItemSelectionModel*)self)->reset();
}
void QItemSelectionModel_clearSelection(void* self) {
	((QItemSelectionModel*)self)->clearSelection();
}
void QItemSelectionModel_clearCurrentIndex(void* self) {
	((QItemSelectionModel*)self)->clearCurrentIndex();
}
void QItemSelectionModel_selectionChanged(void* self, void* selected, void* deselected) {
	((QItemSelectionModel*)self)->selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QItemSelectionModel_currentChanged(void* self, void* current, void* previous) {
	((QItemSelectionModel*)self)->currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QItemSelectionModel_currentRowChanged(void* self, void* current, void* previous) {
	((QItemSelectionModel*)self)->currentRowChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QItemSelectionModel_currentColumnChanged(void* self, void* current, void* previous) {
	((QItemSelectionModel*)self)->currentColumnChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QItemSelectionModel_modelChanged(void* self, void* model) {
	((QItemSelectionModel*)self)->modelChanged((QAbstractItemModel*)model);
}
void QItemSelectionModel_emitSelectionChanged(void* self, void* newSelection, void* oldSelection) {
	((QItemSelectionModel__Go*)self)->emitSelectionChanged(*(QItemSelection*)newSelection, *(QItemSelection*)oldSelection);
}
void* QItemSelectionModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QItemSelectionModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QItemSelectionModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QItemSelectionModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QItemSelectionModel_sender(void* self) {
	auto&& __r = ((QItemSelectionModel__Go*)self)->sender();
	return (void*)(__r);
}
int QItemSelectionModel_senderSignalIndex(void* self) {
	auto&& __r = ((QItemSelectionModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QItemSelectionModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QItemSelectionModel__Go*)self)->receivers(signal);
	return __r;
}
int QItemSelectionModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QItemSelectionModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QItemSelectionModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QItemSelectionModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QItemSelectionModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QItemSelectionModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QItemSelectionModel_isRowSelected2(void* self, int row, void* parent) {
	auto&& __r = ((QItemSelectionModel*)self)->isRowSelected(row, *(QModelIndex*)parent);
	return __r;
}
int QItemSelectionModel_isColumnSelected2(void* self, int column, void* parent) {
	auto&& __r = ((QItemSelectionModel*)self)->isColumnSelected(column, *(QModelIndex*)parent);
	return __r;
}
int QItemSelectionModel_rowIntersectsSelection2(void* self, int row, void* parent) {
	auto&& __r = ((QItemSelectionModel*)self)->rowIntersectsSelection(row, *(QModelIndex*)parent);
	return __r;
}
int QItemSelectionModel_columnIntersectsSelection2(void* self, int column, void* parent) {
	auto&& __r = ((QItemSelectionModel*)self)->columnIntersectsSelection(column, *(QModelIndex*)parent);
	return __r;
}
void** QItemSelectionModel_selectedRows2(void* self, int column, long long* __out_len) {
	auto&& __r = ((QItemSelectionModel*)self)->selectedRows(column);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void** QItemSelectionModel_selectedColumns2(void* self, int row, long long* __out_len) {
	auto&& __r = ((QItemSelectionModel*)self)->selectedColumns(row);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QItemSelectionModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QItemSelectionModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QItemSelectionModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QItemSelectionModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QItemSelectionModel_setCurrentIndex_default(void* self, void* index, int command) {
	((QItemSelectionModel*)self)->QItemSelectionModel::setCurrentIndex(*(QModelIndex*)index, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelectionModel_select_default(void* self, void* index, int command) {
	((QItemSelectionModel*)self)->QItemSelectionModel::select(*(QModelIndex*)index, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelectionModel_select_default2(void* self, void* selection, int command) {
	((QItemSelectionModel*)self)->QItemSelectionModel::select(*(QItemSelection*)selection, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelectionModel_clear_default(void* self) {
	((QItemSelectionModel*)self)->QItemSelectionModel::clear();
}
void QItemSelectionModel_reset_default(void* self) {
	((QItemSelectionModel*)self)->QItemSelectionModel::reset();
}
void QItemSelectionModel_clearCurrentIndex_default(void* self) {
	((QItemSelectionModel*)self)->QItemSelectionModel::clearCurrentIndex();
}
void* QItemSelectionModel_metaObject_default(void* self) {
	auto&& __r = ((QItemSelectionModel*)self)->QItemSelectionModel::metaObject();
	return (void*)(__r);
}
int QItemSelectionModel_event_default(void* self, void* event) {
	auto&& __r = ((QItemSelectionModel*)self)->QItemSelectionModel::event((QEvent*)event);
	return __r;
}
int QItemSelectionModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QItemSelectionModel*)self)->QItemSelectionModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QItemSelectionModel_timerEvent_default(void* self, void* event) {
	((QItemSelectionModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QItemSelectionModel_childEvent_default(void* self, void* event) {
	((QItemSelectionModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QItemSelectionModel_customEvent_default(void* self, void* event) {
	((QItemSelectionModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QItemSelectionModel_connectNotify_default(void* self, void* signal) {
	((QItemSelectionModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QItemSelectionModel_disconnectNotify_default(void* self, void* signal) {
	((QItemSelectionModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QItemSelectionModel_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::selectionChanged, (QItemSelectionModel*)self, [__id](const QItemSelection& selected, const QItemSelection& deselected){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)&selected, (void*)&deselected);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::currentChanged, (QItemSelectionModel*)self, [__id](const QModelIndex& current, const QModelIndex& previous){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)&current, (void*)&previous);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_currentRowChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::currentRowChanged, (QItemSelectionModel*)self, [__id](const QModelIndex& current, const QModelIndex& previous){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)&current, (void*)&previous);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_currentColumnChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::currentColumnChanged, (QItemSelectionModel*)self, [__id](const QModelIndex& current, const QModelIndex& previous){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)&current, (void*)&previous);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_modelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::modelChanged, (QItemSelectionModel*)self, [__id](QAbstractItemModel* model){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)model);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::destroyed, (QItemSelectionModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::destroyed, (QItemSelectionModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemSelectionModel*)self, &QItemSelectionModel::objectNameChanged, (QItemSelectionModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
	__conns_QItemSelectionModel[__id] = __c;
}
void QItemSelectionModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QItemSelectionModel);
		auto __it = __conns_QItemSelectionModel.find(id);
		if (__it == __conns_QItemSelectionModel.end()) return;
		__c = __it->second;
		__conns_QItemSelectionModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QItemSelectionModel__Go_new(long long id) { return new QItemSelectionModel__Go(id); }
void QItemSelectionModel_delete(void* self) { delete (QItemSelectionModel*)self; }

}

#include <QtCore/qmetaobject.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qobjectdefs.h>
#include "gen_qmetaenum.h"

extern "C" {

void* QMetaEnum_new() {
	return new QMetaEnum();
}
void* QMetaEnum_new2(void* arg1) {
	return new QMetaEnum(*(QMetaEnum*)arg1);
}
const char* QMetaEnum_name(void* self) {
	auto&& __r = ((QMetaEnum*)self)->name();
	return __r;
}
const char* QMetaEnum_enumName(void* self) {
	auto&& __r = ((QMetaEnum*)self)->enumName();
	return __r;
}
void* QMetaEnum_metaType(void* self) {
	auto&& __r = ((QMetaEnum*)self)->metaType();
	return new QMetaType(__r);
}
int QMetaEnum_isFlag(void* self) {
	auto&& __r = ((QMetaEnum*)self)->isFlag();
	return __r;
}
int QMetaEnum_isScoped(void* self) {
	auto&& __r = ((QMetaEnum*)self)->isScoped();
	return __r;
}
int QMetaEnum_is64Bit(void* self) {
	auto&& __r = ((QMetaEnum*)self)->is64Bit();
	return __r;
}
int QMetaEnum_keyCount(void* self) {
	auto&& __r = ((QMetaEnum*)self)->keyCount();
	return __r;
}
const char* QMetaEnum_key(void* self, int index) {
	auto&& __r = ((QMetaEnum*)self)->key(index);
	return __r;
}
int QMetaEnum_value(void* self, int index) {
	auto&& __r = ((QMetaEnum*)self)->value(index);
	return __r;
}
unsigned long long QMetaEnum_value64(void* self, int index, int* __ok) {
	auto&& __r = ((QMetaEnum*)self)->value64(index);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? *__r : (unsigned long long)0;
}
const char* QMetaEnum_scope(void* self) {
	auto&& __r = ((QMetaEnum*)self)->scope();
	return __r;
}
int QMetaEnum_keyToValue(void* self, const char* key) {
	auto&& __r = ((QMetaEnum*)self)->keyToValue(key);
	return __r;
}
int QMetaEnum_keysToValue(void* self, const char* keys) {
	auto&& __r = ((QMetaEnum*)self)->keysToValue(keys);
	return __r;
}
unsigned long long QMetaEnum_keyToValue64(void* self, const char* key, int* __ok) {
	auto&& __r = ((QMetaEnum*)self)->keyToValue64(key);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? *__r : (unsigned long long)0;
}
unsigned long long QMetaEnum_keysToValue64(void* self, const char* keys, int* __ok) {
	auto&& __r = ((QMetaEnum*)self)->keysToValue64(keys);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? *__r : (unsigned long long)0;
}
const char* QMetaEnum_valueToKey(void* self, unsigned long long value) {
	auto&& __r = ((QMetaEnum*)self)->valueToKey(value);
	return __r;
}
void* QMetaEnum_valueToKeys(void* self, unsigned long long value) {
	auto&& __r = ((QMetaEnum*)self)->valueToKeys(value);
	return new QByteArray(__r);
}
void* QMetaEnum_enclosingMetaObject(void* self) {
	auto&& __r = ((QMetaEnum*)self)->enclosingMetaObject();
	return (void*)(__r);
}
int QMetaEnum_isValid(void* self) {
	auto&& __r = ((QMetaEnum*)self)->isValid();
	return __r;
}
int QMetaEnum_keyToValue2(void* self, const char* key, int* ok) {
	bool ok_local{};
	auto&& __r = ((QMetaEnum*)self)->keyToValue(key, &ok_local);
	*ok = ok_local;
	return __r;
}
int QMetaEnum_keysToValue2(void* self, const char* keys, int* ok) {
	bool ok_local{};
	auto&& __r = ((QMetaEnum*)self)->keysToValue(keys, &ok_local);
	*ok = ok_local;
	return __r;
}
void QMetaEnum_delete(void* self) { delete (QMetaEnum*)self; }

}

#include <QtCore/qpluginloader.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qplugin.h>
#include <QtCore/qlibrary.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qpluginloader.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPluginLoader__Go : public QPluginLoader {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QPluginLoader__Go(long long id, A&&... a) : QPluginLoader(std::forward<A>(a)...), __goId(id) {}
	using QPluginLoader::sender;
	using QPluginLoader::senderSignalIndex;
	using QPluginLoader::receivers;
	using QPluginLoader::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPluginLoader::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPluginLoader::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPluginLoader::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPluginLoader::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPluginLoader::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPluginLoader::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QPluginLoader::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPluginLoader::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QPluginLoader::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPluginLoader::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPluginLoader::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPluginLoader::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPluginLoader::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPluginLoader;
static std::mutex __connsMu_QPluginLoader;

extern "C" {

void* QPluginLoader_base(void* self) { return (void*)static_cast<QObject*>((QPluginLoader*)self); }
void QPluginLoader__set_override(void* self, int idx) { ((QPluginLoader__Go*)self)->__goOn[idx] = 1; }
void* QPluginLoader_fromQObject(void* o) { return (void*)qobject_cast<QPluginLoader*>((QObject*)o); }
void* QPluginLoader_new(long long __goid) {
	return new QPluginLoader__Go(__goid);
}
void* QPluginLoader_new2(long long __goid, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QPluginLoader__Go(__goid, fileName_qstr);
}
void* QPluginLoader_new3(long long __goid, void* parent) {
	return new QPluginLoader__Go(__goid, (QObject*)parent);
}
void* QPluginLoader_new4(long long __goid, const char* fileName, long long fileName_len, void* parent) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QPluginLoader__Go(__goid, fileName_qstr, (QObject*)parent);
}
char* QPluginLoader_tr(const char* s, long long* __out_len) {
	auto&& __r = QPluginLoader::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPluginLoader_instance(void* self) {
	auto&& __r = ((QPluginLoader*)self)->instance();
	return (void*)(__r);
}
void* QPluginLoader_metaData(void* self) {
	auto&& __r = ((QPluginLoader*)self)->metaData();
	return new QJsonObject(__r);
}
void** QPluginLoader_staticInstances(long long* __out_len) {
	auto&& __r = QPluginLoader::staticInstances();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QPluginLoader_staticPlugins(long long* __out_len) {
	auto&& __r = QPluginLoader::staticPlugins();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStaticPlugin(__e);
	return __arr;
}
int QPluginLoader_load(void* self) {
	auto&& __r = ((QPluginLoader*)self)->load();
	return __r;
}
int QPluginLoader_unload(void* self) {
	auto&& __r = ((QPluginLoader*)self)->unload();
	return __r;
}
int QPluginLoader_isLoaded(void* self) {
	auto&& __r = ((QPluginLoader*)self)->isLoaded();
	return __r;
}
void QPluginLoader_setFileName(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QPluginLoader*)self)->setFileName(fileName_qstr);
}
char* QPluginLoader_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QPluginLoader*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPluginLoader_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QPluginLoader*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPluginLoader_setLoadHints(void* self, int loadHints) {
	((QPluginLoader*)self)->setLoadHints(QFlags<QLibrary::LoadHint>(loadHints));
}
int QPluginLoader_loadHints(void* self) {
	auto&& __r = ((QPluginLoader*)self)->loadHints();
	return (int)(__r);
}
void* QPluginLoader_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPluginLoader::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPluginLoader_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPluginLoader::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPluginLoader_sender(void* self) {
	auto&& __r = ((QPluginLoader__Go*)self)->sender();
	return (void*)(__r);
}
int QPluginLoader_senderSignalIndex(void* self) {
	auto&& __r = ((QPluginLoader__Go*)self)->senderSignalIndex();
	return __r;
}
int QPluginLoader_receivers(void* self, const char* signal) {
	auto&& __r = ((QPluginLoader__Go*)self)->receivers(signal);
	return __r;
}
int QPluginLoader_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPluginLoader__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QPluginLoader_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPluginLoader::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPluginLoader_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPluginLoader::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPluginLoader_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPluginLoader::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPluginLoader_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPluginLoader::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPluginLoader_metaObject_default(void* self) {
	auto&& __r = ((QPluginLoader*)self)->QPluginLoader::metaObject();
	return (void*)(__r);
}
int QPluginLoader_event_default(void* self, void* event) {
	auto&& __r = ((QPluginLoader*)self)->QPluginLoader::event((QEvent*)event);
	return __r;
}
int QPluginLoader_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPluginLoader*)self)->QPluginLoader::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPluginLoader_timerEvent_default(void* self, void* event) {
	((QPluginLoader__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QPluginLoader_childEvent_default(void* self, void* event) {
	((QPluginLoader__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPluginLoader_customEvent_default(void* self, void* event) {
	((QPluginLoader__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPluginLoader_connectNotify_default(void* self, void* signal) {
	((QPluginLoader__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPluginLoader_disconnectNotify_default(void* self, void* signal) {
	((QPluginLoader__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPluginLoader_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPluginLoader*)self, &QPluginLoader::destroyed, (QPluginLoader*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPluginLoader);
	__conns_QPluginLoader[__id] = __c;
}
void QPluginLoader_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPluginLoader*)self, &QPluginLoader::destroyed, (QPluginLoader*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPluginLoader);
	__conns_QPluginLoader[__id] = __c;
}
void QPluginLoader_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPluginLoader*)self, &QPluginLoader::objectNameChanged, (QPluginLoader*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPluginLoader);
	__conns_QPluginLoader[__id] = __c;
}
void QPluginLoader__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPluginLoader);
		auto __it = __conns_QPluginLoader.find(id);
		if (__it == __conns_QPluginLoader.end()) return;
		__c = __it->second;
		__conns_QPluginLoader.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPluginLoader__Go_new(long long id, void* parent) { return new QPluginLoader__Go(id, (QObject*)parent); }
void QPluginLoader_delete(void* self) { delete (QPluginLoader*)self; }

}

#include <QtCore/qsequentialiterable.h>
#include <QtCore/qmetacontainer.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qvariant.h>
#include "gen_qsequentialiterable.h"

extern "C" {

void* QSequentialIterable_new() {
	return new QSequentialIterable();
}
void* QSequentialIterable_new2(void* arg1) {
	return new QSequentialIterable(*(QSequentialIterable*)arg1);
}
void* QSequentialIterable_at(void* self, long long idx) {
	auto&& __r = ((QSequentialIterable*)self)->at(idx);
	return new QVariant(__r);
}
void QSequentialIterable_set(void* self, long long idx, void* value) {
	((QSequentialIterable*)self)->set(idx, *(QVariant*)value);
}
void QSequentialIterable_addValue(void* self, void* value) {
	((QSequentialIterable*)self)->addValue(*(QVariant*)value);
}
void QSequentialIterable_removeValue(void* self) {
	((QSequentialIterable*)self)->removeValue();
}
void* QSequentialIterable_valueMetaType(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->valueMetaType();
	return new QMetaType(__r);
}
int QSequentialIterable_canRandomAccessIterate(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->canRandomAccessIterate();
	return __r;
}
int QSequentialIterable_canInputIterate(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->canInputIterate();
	return __r;
}
int QSequentialIterable_canForwardIterate(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->canForwardIterate();
	return __r;
}
int QSequentialIterable_canReverseIterate(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->canReverseIterate();
	return __r;
}
long long QSequentialIterable_size(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->size();
	return __r;
}
void QSequentialIterable_clear(void* self) {
	((QSequentialIterable*)self)->clear();
}
void* QSequentialIterable_metaContainer(void* self) {
	auto&& __r = ((QSequentialIterable*)self)->metaContainer();
	return new QMetaSequence(__r);
}
void QSequentialIterable_addValue2(void* self, void* value, int position) {
	((QSequentialIterable*)self)->addValue(*(QVariant*)value, (QSequentialIterable::Position)(position));
}
void QSequentialIterable_removeValue2(void* self, int position) {
	((QSequentialIterable*)self)->removeValue((QSequentialIterable::Position)(position));
}
void QSequentialIterable_delete(void* self) { delete (QSequentialIterable*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include "gen_qxmlstreamnamespacedeclaration.h"

extern "C" {

void* QXmlStreamNamespaceDeclaration_new() {
	return new QXmlStreamNamespaceDeclaration();
}
void* QXmlStreamNamespaceDeclaration_new2(const char* prefix, long long prefix_len, const char* namespaceUri, long long namespaceUri_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	QString namespaceUri_qstr = QString::fromUtf8(namespaceUri, namespaceUri_len);
	return new QXmlStreamNamespaceDeclaration(prefix_qstr, namespaceUri_qstr);
}
void* QXmlStreamNamespaceDeclaration_new3(void* arg1) {
	return new QXmlStreamNamespaceDeclaration(*(QXmlStreamNamespaceDeclaration*)arg1);
}
void* QXmlStreamNamespaceDeclaration_prefix(void* self) {
	auto&& __r = ((QXmlStreamNamespaceDeclaration*)self)->prefix();
	return new QStringView(__r);
}
void* QXmlStreamNamespaceDeclaration_namespaceUri(void* self) {
	auto&& __r = ((QXmlStreamNamespaceDeclaration*)self)->namespaceUri();
	return new QStringView(__r);
}
int QXmlStreamNamespaceDeclaration_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamNamespaceDeclaration*)self) == (*(QXmlStreamNamespaceDeclaration*)rhs);
	return __r;
}
int QXmlStreamNamespaceDeclaration_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamNamespaceDeclaration*)self) != (*(QXmlStreamNamespaceDeclaration*)rhs);
	return __r;
}
void QXmlStreamNamespaceDeclaration_delete(void* self) { delete (QXmlStreamNamespaceDeclaration*)self; }

}

#include <QtGui/qinputdevice.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qinputdevice.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QInputDevice__Go : public QInputDevice {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QInputDevice__Go(long long id, A&&... a) : QInputDevice(std::forward<A>(a)...), __goId(id) {}
	using QInputDevice::sender;
	using QInputDevice::senderSignalIndex;
	using QInputDevice::receivers;
	using QInputDevice::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QInputDevice::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QInputDevice::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QInputDevice::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QInputDevice::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QInputDevice::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QInputDevice::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QInputDevice::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QInputDevice::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QInputDevice::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QInputDevice::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QInputDevice::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QInputDevice::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QInputDevice::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QInputDevice;
static std::mutex __connsMu_QInputDevice;

extern "C" {

void* QInputDevice_base(void* self) { return (void*)static_cast<QObject*>((QInputDevice*)self); }
void QInputDevice__set_override(void* self, int idx) { ((QInputDevice__Go*)self)->__goOn[idx] = 1; }
void* QInputDevice_fromQObject(void* o) { return (void*)qobject_cast<QInputDevice*>((QObject*)o); }
void* QInputDevice_new(long long __goid) {
	return new QInputDevice__Go(__goid);
}
void* QInputDevice_new2(long long __goid, const char* name, long long name_len, long long systemId, int type_) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QInputDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(type_));
}
void* QInputDevice_new3(long long __goid, void* parent) {
	return new QInputDevice__Go(__goid, (QObject*)parent);
}
void* QInputDevice_new4(long long __goid, const char* name, long long name_len, long long systemId, int type_, const char* seatName, long long seatName_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	return new QInputDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(type_), seatName_qstr);
}
void* QInputDevice_new5(long long __goid, const char* name, long long name_len, long long systemId, int type_, const char* seatName, long long seatName_len, void* parent) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	return new QInputDevice__Go(__goid, name_qstr, systemId, (QInputDevice::DeviceType)(type_), seatName_qstr, (QObject*)parent);
}
char* QInputDevice_tr(const char* s, long long* __out_len) {
	auto&& __r = QInputDevice::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDevice_name(void* self, long long* __out_len) {
	auto&& __r = ((QInputDevice*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QInputDevice_type(void* self) {
	auto&& __r = ((QInputDevice*)self)->type();
	return (int)(__r);
}
int QInputDevice_capabilities(void* self) {
	auto&& __r = ((QInputDevice*)self)->capabilities();
	return (int)(__r);
}
int QInputDevice_hasCapability(void* self, int cap_) {
	auto&& __r = ((QInputDevice*)self)->hasCapability((QInputDevice::Capability)(cap_));
	return __r;
}
long long QInputDevice_systemId(void* self) {
	auto&& __r = ((QInputDevice*)self)->systemId();
	return __r;
}
char* QInputDevice_seatName(void* self, long long* __out_len) {
	auto&& __r = ((QInputDevice*)self)->seatName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QInputDevice_availableVirtualGeometry(void* self) {
	auto&& __r = ((QInputDevice*)self)->availableVirtualGeometry();
	return new QRect(__r);
}
char** QInputDevice_seatNames(long long* __out_len) {
	auto&& __r = QInputDevice::seatNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QInputDevice_devices(long long* __out_len) {
	auto&& __r = QInputDevice::devices();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QInputDevice_primaryKeyboard() {
	auto&& __r = QInputDevice::primaryKeyboard();
	return (void*)(__r);
}
int QInputDevice_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QInputDevice*)self)->operator==(*(QInputDevice*)other);
	return __r;
}
void QInputDevice_availableVirtualGeometryChanged(void* self, void* area) {
	((QInputDevice*)self)->availableVirtualGeometryChanged(*(QRect*)area);
}
void QInputDevice_capabilitiesChanged(void* self, int capabilities) {
	((QInputDevice*)self)->capabilitiesChanged(QFlags<QInputDevice::Capability>(capabilities));
}
void* QInputDevice_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QInputDevice::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QInputDevice_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QInputDevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QInputDevice_sender(void* self) {
	auto&& __r = ((QInputDevice__Go*)self)->sender();
	return (void*)(__r);
}
int QInputDevice_senderSignalIndex(void* self) {
	auto&& __r = ((QInputDevice__Go*)self)->senderSignalIndex();
	return __r;
}
int QInputDevice_receivers(void* self, const char* signal) {
	auto&& __r = ((QInputDevice__Go*)self)->receivers(signal);
	return __r;
}
int QInputDevice_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QInputDevice__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QInputDevice_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QInputDevice::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputDevice_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QInputDevice::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QInputDevice_primaryKeyboard2(const char* seatName, long long seatName_len) {
	QString seatName_qstr = QString::fromUtf8(seatName, seatName_len);
	auto&& __r = QInputDevice::primaryKeyboard(seatName_qstr);
	return (void*)(__r);
}
void* QInputDevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QInputDevice::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QInputDevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QInputDevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QInputDevice_metaObject_default(void* self) {
	auto&& __r = ((QInputDevice*)self)->QInputDevice::metaObject();
	return (void*)(__r);
}
int QInputDevice_event_default(void* self, void* event) {
	auto&& __r = ((QInputDevice*)self)->QInputDevice::event((QEvent*)event);
	return __r;
}
int QInputDevice_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QInputDevice*)self)->QInputDevice::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QInputDevice_timerEvent_default(void* self, void* event) {
	((QInputDevice__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QInputDevice_childEvent_default(void* self, void* event) {
	((QInputDevice__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QInputDevice_customEvent_default(void* self, void* event) {
	((QInputDevice__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QInputDevice_connectNotify_default(void* self, void* signal) {
	((QInputDevice__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QInputDevice_disconnectNotify_default(void* self, void* signal) {
	((QInputDevice__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QInputDevice_connect_availableVirtualGeometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDevice*)self, &QInputDevice::availableVirtualGeometryChanged, (QInputDevice*)self, [__id](QRect area){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&area);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDevice);
	__conns_QInputDevice[__id] = __c;
}
void QInputDevice_connect_capabilitiesChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDevice*)self, &QInputDevice::capabilitiesChanged, (QInputDevice*)self, [__id](QFlags<QInputDevice::Capability> capabilities){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(capabilities));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDevice);
	__conns_QInputDevice[__id] = __c;
}
void QInputDevice_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDevice*)self, &QInputDevice::destroyed, (QInputDevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDevice);
	__conns_QInputDevice[__id] = __c;
}
void QInputDevice_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDevice*)self, &QInputDevice::destroyed, (QInputDevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDevice);
	__conns_QInputDevice[__id] = __c;
}
void QInputDevice_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputDevice*)self, &QInputDevice::objectNameChanged, (QInputDevice*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputDevice);
	__conns_QInputDevice[__id] = __c;
}
void QInputDevice__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QInputDevice);
		auto __it = __conns_QInputDevice.find(id);
		if (__it == __conns_QInputDevice.end()) return;
		__c = __it->second;
		__conns_QInputDevice.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QInputDevice__Go_new(long long id, void* parent) { return new QInputDevice__Go(id, (QObject*)parent); }
void QInputDevice_delete(void* self) { delete (QInputDevice*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qhideevent.h"
#include "gen_purego_cb.h"

class QHideEvent__Go : public QHideEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QHideEvent__Go(long long id, A&&... a) : QHideEvent(std::forward<A>(a)...), __goId(id) {}
	QHideEvent* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QHideEvent*)__r; }
		return QHideEvent::clone();
	}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QHideEvent::setAccepted(accepted);
	}
};

extern "C" {

void* QHideEvent_base(void* self) { return (void*)static_cast<QEvent*>((QHideEvent*)self); }
void QHideEvent__set_override(void* self, int idx) { ((QHideEvent__Go*)self)->__goOn[idx] = 1; }
void* QHideEvent_new(long long __goid) {
	return new QHideEvent__Go(__goid);
}
int QHideEvent_registerEventType() {
	auto&& __r = QHideEvent::registerEventType();
	return __r;
}
int QHideEvent_registerEventType2(int hint) {
	auto&& __r = QHideEvent::registerEventType(hint);
	return __r;
}
void* QHideEvent_clone_default(void* self) {
	auto&& __r = ((QHideEvent*)self)->QHideEvent::clone();
	return (void*)(__r);
}
void QHideEvent_setAccepted_default(void* self, int accepted) {
	((QHideEvent*)self)->QHideEvent::setAccepted(accepted);
}
void* QHideEvent__Go_new(long long id) { return new QHideEvent__Go(id); }
void QHideEvent_delete(void* self) { delete (QHideEvent*)self; }

}

#include <QtGui/qfont.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qanystringview.h>
#include "gen_qfont__tag.h"

extern "C" {

void* QFont__Tag_new() {
	return new QFont::Tag();
}
void* QFont__Tag_new2(void* arg1) {
	return new QFont::Tag(*(QFont::Tag*)arg1);
}
int QFont__Tag_isValid(void* self) {
	auto&& __r = ((QFont::Tag*)self)->isValid();
	return __r;
}
unsigned int QFont__Tag_value(void* self) {
	auto&& __r = ((QFont::Tag*)self)->value();
	return __r;
}
void* QFont__Tag_toString(void* self) {
	auto&& __r = ((QFont::Tag*)self)->toString();
	return new QByteArray(__r);
}
void* QFont__Tag_fromValue(unsigned int value, int* __ok) {
	auto&& __r = QFont::Tag::fromValue(value);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? new QFont::Tag(*__r) : (void*)0;
}
void* QFont__Tag_fromString(const char* view, long long view_len, int* __ok) {
	QAnyStringView view_qstr = QAnyStringView(view, view_len);
	auto&& __r = QFont::Tag::fromString(view_qstr);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? new QFont::Tag(*__r) : (void*)0;
}
void* QFont__Tag_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QFont::Tag*)self)->operator=(*(QFont::Tag*)arg1);
	return (void*)&(__r);
}
int QFont__Tag_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QFont::Tag*)self) == (*(QFont::Tag*)rhs);
	return __r;
}
int QFont__Tag_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QFont::Tag*)self) != (*(QFont::Tag*)rhs);
	return __r;
}
int QFont__Tag_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QFont::Tag*)self) < (*(QFont::Tag*)rhs);
	return __r;
}
int QFont__Tag_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QFont::Tag*)self) > (*(QFont::Tag*)rhs);
	return __r;
}
int QFont__Tag_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QFont::Tag*)self) <= (*(QFont::Tag*)rhs);
	return __r;
}
int QFont__Tag_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QFont::Tag*)self) >= (*(QFont::Tag*)rhs);
	return __r;
}
void QFont__Tag_delete(void* self) { delete (QFont::Tag*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtGui/qbrush.h>
#include <QtGui/qfont.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtCore/qnamespace.h>
#include "gen_qtexttablecellformat.h"

extern "C" {

void* QTextTableCellFormat_base(void* self) { return (void*)static_cast<QTextCharFormat*>((QTextTableCellFormat*)self); }
void QTextTableCellFormat_chain(void* self, void** out) {
	QTextTableCellFormat* __p = (QTextTableCellFormat*)self;
	QTextCharFormat* __b0 = static_cast<QTextCharFormat*>(__p); out[0] = (void*)__b0;
	QTextFormat* __b1 = static_cast<QTextFormat*>(__b0); out[1] = (void*)__b1;
}
void* QTextTableCellFormat_new() {
	return new QTextTableCellFormat();
}
void* QTextTableCellFormat_new2(void* arg1) {
	return new QTextTableCellFormat(*(QTextTableCellFormat*)arg1);
}
int QTextTableCellFormat_isValid(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->isValid();
	return __r;
}
void QTextTableCellFormat_setTopPadding(void* self, double padding) {
	((QTextTableCellFormat*)self)->setTopPadding(padding);
}
double QTextTableCellFormat_topPadding(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->topPadding();
	return __r;
}
void QTextTableCellFormat_setBottomPadding(void* self, double padding) {
	((QTextTableCellFormat*)self)->setBottomPadding(padding);
}
double QTextTableCellFormat_bottomPadding(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->bottomPadding();
	return __r;
}
void QTextTableCellFormat_setLeftPadding(void* self, double padding) {
	((QTextTableCellFormat*)self)->setLeftPadding(padding);
}
double QTextTableCellFormat_leftPadding(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->leftPadding();
	return __r;
}
void QTextTableCellFormat_setRightPadding(void* self, double padding) {
	((QTextTableCellFormat*)self)->setRightPadding(padding);
}
double QTextTableCellFormat_rightPadding(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->rightPadding();
	return __r;
}
void QTextTableCellFormat_setPadding(void* self, double padding) {
	((QTextTableCellFormat*)self)->setPadding(padding);
}
void QTextTableCellFormat_setTopBorder(void* self, double width) {
	((QTextTableCellFormat*)self)->setTopBorder(width);
}
double QTextTableCellFormat_topBorder(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->topBorder();
	return __r;
}
void QTextTableCellFormat_setBottomBorder(void* self, double width) {
	((QTextTableCellFormat*)self)->setBottomBorder(width);
}
double QTextTableCellFormat_bottomBorder(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->bottomBorder();
	return __r;
}
void QTextTableCellFormat_setLeftBorder(void* self, double width) {
	((QTextTableCellFormat*)self)->setLeftBorder(width);
}
double QTextTableCellFormat_leftBorder(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->leftBorder();
	return __r;
}
void QTextTableCellFormat_setRightBorder(void* self, double width) {
	((QTextTableCellFormat*)self)->setRightBorder(width);
}
double QTextTableCellFormat_rightBorder(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->rightBorder();
	return __r;
}
void QTextTableCellFormat_setBorder(void* self, double width) {
	((QTextTableCellFormat*)self)->setBorder(width);
}
void QTextTableCellFormat_setTopBorderStyle(void* self, int style) {
	((QTextTableCellFormat*)self)->setTopBorderStyle((QTextFrameFormat::BorderStyle)(style));
}
int QTextTableCellFormat_topBorderStyle(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->topBorderStyle();
	return (int)(__r);
}
void QTextTableCellFormat_setBottomBorderStyle(void* self, int style) {
	((QTextTableCellFormat*)self)->setBottomBorderStyle((QTextFrameFormat::BorderStyle)(style));
}
int QTextTableCellFormat_bottomBorderStyle(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->bottomBorderStyle();
	return (int)(__r);
}
void QTextTableCellFormat_setLeftBorderStyle(void* self, int style) {
	((QTextTableCellFormat*)self)->setLeftBorderStyle((QTextFrameFormat::BorderStyle)(style));
}
int QTextTableCellFormat_leftBorderStyle(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->leftBorderStyle();
	return (int)(__r);
}
void QTextTableCellFormat_setRightBorderStyle(void* self, int style) {
	((QTextTableCellFormat*)self)->setRightBorderStyle((QTextFrameFormat::BorderStyle)(style));
}
int QTextTableCellFormat_rightBorderStyle(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->rightBorderStyle();
	return (int)(__r);
}
void QTextTableCellFormat_setBorderStyle(void* self, int style) {
	((QTextTableCellFormat*)self)->setBorderStyle((QTextFrameFormat::BorderStyle)(style));
}
void QTextTableCellFormat_setTopBorderBrush(void* self, void* brush) {
	((QTextTableCellFormat*)self)->setTopBorderBrush(*(QBrush*)brush);
}
void* QTextTableCellFormat_topBorderBrush(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->topBorderBrush();
	return new QBrush(__r);
}
void QTextTableCellFormat_setBottomBorderBrush(void* self, void* brush) {
	((QTextTableCellFormat*)self)->setBottomBorderBrush(*(QBrush*)brush);
}
void* QTextTableCellFormat_bottomBorderBrush(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->bottomBorderBrush();
	return new QBrush(__r);
}
void QTextTableCellFormat_setLeftBorderBrush(void* self, void* brush) {
	((QTextTableCellFormat*)self)->setLeftBorderBrush(*(QBrush*)brush);
}
void* QTextTableCellFormat_leftBorderBrush(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->leftBorderBrush();
	return new QBrush(__r);
}
void QTextTableCellFormat_setRightBorderBrush(void* self, void* brush) {
	((QTextTableCellFormat*)self)->setRightBorderBrush(*(QBrush*)brush);
}
void* QTextTableCellFormat_rightBorderBrush(void* self) {
	auto&& __r = ((QTextTableCellFormat*)self)->rightBorderBrush();
	return new QBrush(__r);
}
void QTextTableCellFormat_setBorderBrush(void* self, void* brush) {
	((QTextTableCellFormat*)self)->setBorderBrush(*(QBrush*)brush);
}
void QTextTableCellFormat_delete(void* self) { delete (QTextTableCellFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include "gen_qaccessibleselectioninterface.h"
#include "gen_purego_cb.h"

class QAccessibleSelectionInterface__Go : public QAccessibleSelectionInterface {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QAccessibleSelectionInterface__Go(long long id, A&&... a) : QAccessibleSelectionInterface(std::forward<A>(a)...), __goId(id) {}
	int selectedItemCount() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1835560489624661370LL, &__h); if (__h) return __r; }
		return int{};
	}
	QList<QAccessibleInterface*> selectedItems() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, -5807941332178582040LL, &__len, &__h);
			if (__h) {
				QList<QAccessibleInterface*> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append((QAccessibleInterface*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QAccessibleInterface*>{};
	}
	bool select(QAccessibleInterface* childItem) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -4794627110275867635LL, (void*)childItem, &__h); if (__h) return __r; }
		return bool{};
	}
	bool unselect(QAccessibleInterface* childItem) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5098092392905213826LL, (void*)childItem, &__h); if (__h) return __r; }
		return bool{};
	}
	bool selectAll() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 8758113809152574300LL, &__h); if (__h) return __r; }
		return bool{};
	}
	bool clear() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -778567940572639070LL, &__h); if (__h) return __r; }
		return bool{};
	}
	QAccessibleInterface* selectedItem(int selectionIndex) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -864810461788103157LL, selectionIndex, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QAccessibleSelectionInterface::selectedItem(selectionIndex);
	}
	bool isSelected(QAccessibleInterface* childItem) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3350427747592150876LL, (void*)childItem, &__h); if (__h) return __r; }
		return QAccessibleSelectionInterface::isSelected(childItem);
	}
};

extern "C" {

void QAccessibleSelectionInterface__set_override(void* self, int idx) { ((QAccessibleSelectionInterface__Go*)self)->__goOn[idx] = 1; }
int QAccessibleSelectionInterface_selectedItemCount(void* self) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->selectedItemCount();
	return __r;
}
void** QAccessibleSelectionInterface_selectedItems(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->selectedItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QAccessibleSelectionInterface_selectedItem(void* self, int selectionIndex) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->selectedItem(selectionIndex);
	return (void*)(__r);
}
int QAccessibleSelectionInterface_isSelected(void* self, void* childItem) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->isSelected((QAccessibleInterface*)childItem);
	return __r;
}
int QAccessibleSelectionInterface_select(void* self, void* childItem) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->select((QAccessibleInterface*)childItem);
	return __r;
}
int QAccessibleSelectionInterface_unselect(void* self, void* childItem) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->unselect((QAccessibleInterface*)childItem);
	return __r;
}
int QAccessibleSelectionInterface_selectAll(void* self) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->selectAll();
	return __r;
}
int QAccessibleSelectionInterface_clear(void* self) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->clear();
	return __r;
}
void* QAccessibleSelectionInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->operator=(*(QAccessibleSelectionInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleSelectionInterface_selectedItem_default(void* self, int selectionIndex) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->QAccessibleSelectionInterface::selectedItem(selectionIndex);
	return (void*)(__r);
}
int QAccessibleSelectionInterface_isSelected_default(void* self, void* childItem) {
	auto&& __r = ((QAccessibleSelectionInterface*)self)->QAccessibleSelectionInterface::isSelected((QAccessibleInterface*)childItem);
	return __r;
}
void* QAccessibleSelectionInterface__Go_new(long long id) { return new QAccessibleSelectionInterface__Go(id); }
void QAccessibleSelectionInterface_delete(void* self) { delete (QAccessibleSelectionInterface*)self; }

}

#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qregion.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include "gen_qbackingstore.h"

extern "C" {

void* QBackingStore_new(void* window) {
	return new QBackingStore((QWindow*)window);
}
void* QBackingStore_window(void* self) {
	auto&& __r = ((QBackingStore*)self)->window();
	return (void*)(__r);
}
void* QBackingStore_paintDevice(void* self) {
	auto&& __r = ((QBackingStore*)self)->paintDevice();
	return (void*)(__r);
}
void QBackingStore_flush(void* self, void* region) {
	((QBackingStore*)self)->flush(*(QRegion*)region);
}
void QBackingStore_resize(void* self, void* size) {
	((QBackingStore*)self)->resize(*(QSize*)size);
}
void* QBackingStore_size(void* self) {
	auto&& __r = ((QBackingStore*)self)->size();
	return new QSize(__r);
}
int QBackingStore_scroll(void* self, void* area, int dx, int dy) {
	auto&& __r = ((QBackingStore*)self)->scroll(*(QRegion*)area, dx, dy);
	return __r;
}
void QBackingStore_beginPaint(void* self, void* arg1) {
	((QBackingStore*)self)->beginPaint(*(QRegion*)arg1);
}
void QBackingStore_endPaint(void* self) {
	((QBackingStore*)self)->endPaint();
}
void QBackingStore_setStaticContents(void* self, void* region) {
	((QBackingStore*)self)->setStaticContents(*(QRegion*)region);
}
void* QBackingStore_staticContents(void* self) {
	auto&& __r = ((QBackingStore*)self)->staticContents();
	return new QRegion(__r);
}
int QBackingStore_hasStaticContents(void* self) {
	auto&& __r = ((QBackingStore*)self)->hasStaticContents();
	return __r;
}
void QBackingStore_flush2(void* self, void* region, void* window) {
	((QBackingStore*)self)->flush(*(QRegion*)region, (QWindow*)window);
}
void QBackingStore_flush3(void* self, void* region, void* window, void* offset) {
	((QBackingStore*)self)->flush(*(QRegion*)region, (QWindow*)window, *(QPoint*)offset);
}
void QBackingStore_delete(void* self) { delete (QBackingStore*)self; }

}

#include <QtGui/qopenglcontext.h>
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
#include "gen_qopenglcontextgroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QOpenGLContextGroup;
static std::mutex __connsMu_QOpenGLContextGroup;

extern "C" {

void* QOpenGLContextGroup_base(void* self) { return (void*)static_cast<QObject*>((QOpenGLContextGroup*)self); }
void* QOpenGLContextGroup_fromQObject(void* o) { return (void*)qobject_cast<QOpenGLContextGroup*>((QObject*)o); }
char* QOpenGLContextGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QOpenGLContextGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QOpenGLContextGroup_shares(void* self, long long* __out_len) {
	auto&& __r = ((QOpenGLContextGroup*)self)->shares();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QOpenGLContextGroup_currentContextGroup() {
	auto&& __r = QOpenGLContextGroup::currentContextGroup();
	return (void*)(__r);
}
void* QOpenGLContextGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QOpenGLContextGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QOpenGLContextGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QOpenGLContextGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QOpenGLContextGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QOpenGLContextGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QOpenGLContextGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QOpenGLContextGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QOpenGLContextGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QOpenGLContextGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QOpenGLContextGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QOpenGLContextGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QOpenGLContextGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContextGroup*)self, &QOpenGLContextGroup::destroyed, (QOpenGLContextGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContextGroup);
	__conns_QOpenGLContextGroup[__id] = __c;
}
void QOpenGLContextGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContextGroup*)self, &QOpenGLContextGroup::destroyed, (QOpenGLContextGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContextGroup);
	__conns_QOpenGLContextGroup[__id] = __c;
}
void QOpenGLContextGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContextGroup*)self, &QOpenGLContextGroup::objectNameChanged, (QOpenGLContextGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContextGroup);
	__conns_QOpenGLContextGroup[__id] = __c;
}
void QOpenGLContextGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContextGroup);
		auto __it = __conns_QOpenGLContextGroup.find(id);
		if (__it == __conns_QOpenGLContextGroup.end()) return;
		__c = __it->second;
		__conns_QOpenGLContextGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void QOpenGLContextGroup_delete(void* self) { delete (QOpenGLContextGroup*)self; }

}

#include <QtGui/qquaternion.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qvectornd.h>
#include <QtCore/qvariant.h>
#include "gen_qquaternion.h"

extern "C" {

void* QQuaternion_new() {
	return new QQuaternion();
}
void* QQuaternion_new2(int arg1) {
	return new QQuaternion((Qt::Initialization)(arg1));
}
void* QQuaternion_new3(float scalar, float xpos, float ypos, float zpos) {
	return new QQuaternion(scalar, xpos, ypos, zpos);
}
void* QQuaternion_new4(float scalar, void* vector) {
	return new QQuaternion(scalar, *(QVector3D*)vector);
}
void* QQuaternion_new5(void* vector) {
	return new QQuaternion(*(QVector4D*)vector);
}
void* QQuaternion_new6(void* arg1) {
	return new QQuaternion(*(QQuaternion*)arg1);
}
int QQuaternion_isNull(void* self) {
	auto&& __r = ((QQuaternion*)self)->isNull();
	return __r;
}
int QQuaternion_isIdentity(void* self) {
	auto&& __r = ((QQuaternion*)self)->isIdentity();
	return __r;
}
void* QQuaternion_vector(void* self) {
	auto&& __r = ((QQuaternion*)self)->vector();
	return new QVector3D(__r);
}
void QQuaternion_setVector(void* self, void* vector) {
	((QQuaternion*)self)->setVector(*(QVector3D*)vector);
}
void QQuaternion_setVector2(void* self, float x, float y, float z) {
	((QQuaternion*)self)->setVector(x, y, z);
}
float QQuaternion_x(void* self) {
	auto&& __r = ((QQuaternion*)self)->x();
	return __r;
}
float QQuaternion_y(void* self) {
	auto&& __r = ((QQuaternion*)self)->y();
	return __r;
}
float QQuaternion_z(void* self) {
	auto&& __r = ((QQuaternion*)self)->z();
	return __r;
}
float QQuaternion_scalar(void* self) {
	auto&& __r = ((QQuaternion*)self)->scalar();
	return __r;
}
void QQuaternion_setX(void* self, float x) {
	((QQuaternion*)self)->setX(x);
}
void QQuaternion_setY(void* self, float y) {
	((QQuaternion*)self)->setY(y);
}
void QQuaternion_setZ(void* self, float z) {
	((QQuaternion*)self)->setZ(z);
}
void QQuaternion_setScalar(void* self, float scalar) {
	((QQuaternion*)self)->setScalar(scalar);
}
float QQuaternion_dotProduct(void* q1, void* q2) {
	auto&& __r = QQuaternion::dotProduct(*(QQuaternion*)q1, *(QQuaternion*)q2);
	return __r;
}
float QQuaternion_length(void* self) {
	auto&& __r = ((QQuaternion*)self)->length();
	return __r;
}
float QQuaternion_lengthSquared(void* self) {
	auto&& __r = ((QQuaternion*)self)->lengthSquared();
	return __r;
}
void* QQuaternion_normalized(void* self) {
	auto&& __r = ((QQuaternion*)self)->normalized();
	return new QQuaternion(__r);
}
void QQuaternion_normalize(void* self) {
	((QQuaternion*)self)->normalize();
}
void* QQuaternion_inverted(void* self) {
	auto&& __r = ((QQuaternion*)self)->inverted();
	return new QQuaternion(__r);
}
void* QQuaternion_conjugated(void* self) {
	auto&& __r = ((QQuaternion*)self)->conjugated();
	return new QQuaternion(__r);
}
void* QQuaternion_rotatedVector(void* self, void* vector) {
	auto&& __r = ((QQuaternion*)self)->rotatedVector(*(QVector3D*)vector);
	return new QVector3D(__r);
}
void* QQuaternion_OperatorPlusAssign(void* self, void* quaternion) {
	auto&& __r = ((QQuaternion*)self)->operator+=(*(QQuaternion*)quaternion);
	return (void*)&(__r);
}
void* QQuaternion_OperatorMinusAssign(void* self, void* quaternion) {
	auto&& __r = ((QQuaternion*)self)->operator-=(*(QQuaternion*)quaternion);
	return (void*)&(__r);
}
void* QQuaternion_OperatorMultiplyAssign(void* self, float factor) {
	auto&& __r = ((QQuaternion*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QQuaternion_OperatorMultiplyAssign2(void* self, void* quaternion) {
	auto&& __r = ((QQuaternion*)self)->operator*=(*(QQuaternion*)quaternion);
	return (void*)&(__r);
}
void* QQuaternion_OperatorDivideAssign(void* self, float divisor) {
	auto&& __r = ((QQuaternion*)self)->operator/=(divisor);
	return (void*)&(__r);
}
void* QQuaternion_toVector4D(void* self) {
	auto&& __r = ((QQuaternion*)self)->toVector4D();
	return new QVector4D(__r);
}
void* QQuaternion_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QQuaternion*)self);
	return new QVariant(__r);
}
void QQuaternion_getAxisAndAngle(void* self, void* axis, float* angle) {
	float angle_local{};
	((QQuaternion*)self)->getAxisAndAngle((QVector3D*)axis, &angle_local);
	*angle = angle_local;
}
void* QQuaternion_fromAxisAndAngle(void* axis, float angle) {
	auto&& __r = QQuaternion::fromAxisAndAngle(*(QVector3D*)axis, angle);
	return new QQuaternion(__r);
}
void QQuaternion_getAxisAndAngle2(void* self, float* x, float* y, float* z, float* angle) {
	float x_local{};
	float y_local{};
	float z_local{};
	float angle_local{};
	((QQuaternion*)self)->getAxisAndAngle(&x_local, &y_local, &z_local, &angle_local);
	*x = x_local;
	*y = y_local;
	*z = z_local;
	*angle = angle_local;
}
void* QQuaternion_fromAxisAndAngle2(float x, float y, float z, float angle) {
	auto&& __r = QQuaternion::fromAxisAndAngle(x, y, z, angle);
	return new QQuaternion(__r);
}
void* QQuaternion_toEulerAngles(void* self) {
	auto&& __r = ((QQuaternion*)self)->toEulerAngles();
	return new QVector3D(__r);
}
void QQuaternion_getEulerAngles(void* self, float* pitch, float* yaw, float* roll) {
	float pitch_local{};
	float yaw_local{};
	float roll_local{};
	((QQuaternion*)self)->getEulerAngles(&pitch_local, &yaw_local, &roll_local);
	*pitch = pitch_local;
	*yaw = yaw_local;
	*roll = roll_local;
}
void* QQuaternion_fromEulerAngles(float pitch, float yaw, float roll) {
	auto&& __r = QQuaternion::fromEulerAngles(pitch, yaw, roll);
	return new QQuaternion(__r);
}
void* QQuaternion_toAxes(void* self) {
	auto&& __r = ((QQuaternion*)self)->toAxes();
	return new QQuaternion::Axes(__r);
}
void* QQuaternion_fromAxes(void* axes) {
	auto&& __r = QQuaternion::fromAxes(*(QQuaternion::Axes*)axes);
	return new QQuaternion(__r);
}
void QQuaternion_getAxes(void* self, void* xAxis, void* yAxis, void* zAxis) {
	((QQuaternion*)self)->getAxes((QVector3D*)xAxis, (QVector3D*)yAxis, (QVector3D*)zAxis);
}
void* QQuaternion_fromAxes2(void* xAxis, void* yAxis, void* zAxis) {
	auto&& __r = QQuaternion::fromAxes(*(QVector3D*)xAxis, *(QVector3D*)yAxis, *(QVector3D*)zAxis);
	return new QQuaternion(__r);
}
void* QQuaternion_fromDirection(void* direction, void* up) {
	auto&& __r = QQuaternion::fromDirection(*(QVector3D*)direction, *(QVector3D*)up);
	return new QQuaternion(__r);
}
void* QQuaternion_rotationTo(void* from, void* to) {
	auto&& __r = QQuaternion::rotationTo(*(QVector3D*)from, *(QVector3D*)to);
	return new QQuaternion(__r);
}
void* QQuaternion_slerp(void* q1, void* q2, float t) {
	auto&& __r = QQuaternion::slerp(*(QQuaternion*)q1, *(QQuaternion*)q2, t);
	return new QQuaternion(__r);
}
void* QQuaternion_nlerp(void* q1, void* q2, float t) {
	auto&& __r = QQuaternion::nlerp(*(QQuaternion*)q1, *(QQuaternion*)q2, t);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QQuaternion*)self)->operator=(*(QQuaternion*)arg1);
	return (void*)&(__r);
}
int QQuaternion_OperatorEqual(void* self, void* q2) {
	auto&& __r = (*(QQuaternion*)self) == (*(QQuaternion*)q2);
	return __r;
}
int QQuaternion_OperatorNotEqual(void* self, void* q2) {
	auto&& __r = (*(QQuaternion*)self) != (*(QQuaternion*)q2);
	return __r;
}
void* QQuaternion_OperatorPlus(void* self, void* q2) {
	auto&& __r = (*(QQuaternion*)self) + (*(QQuaternion*)q2);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorMinus(void* self, void* q2) {
	auto&& __r = (*(QQuaternion*)self) - (*(QQuaternion*)q2);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorMultiply(void* self, float factor) {
	auto&& __r = (*(QQuaternion*)self) * (factor);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorMultiply2(void* self, void* q2) {
	auto&& __r = (*(QQuaternion*)self) * (*(QQuaternion*)q2);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorMinus2(void* self) {
	auto&& __r = -(*(QQuaternion*)self);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorDivide(void* self, float divisor) {
	auto&& __r = (*(QQuaternion*)self) / (divisor);
	return new QQuaternion(__r);
}
void* QQuaternion_OperatorMultiply3(void* self, void* vec) {
	auto&& __r = (*(QQuaternion*)self) * (*(QVector3D*)vec);
	return new QVector3D(__r);
}
void QQuaternion_delete(void* self) { delete (QQuaternion*)self; }

}

#include <QtGui/qundostack.h>
#include <QtCore/qstring.h>
#include "gen_qundocommand.h"
#include "gen_purego_cb.h"

class QUndoCommand__Go : public QUndoCommand {
public:
	long long __goId;
	unsigned char __goOn[4] = {};
	template<typename... A> QUndoCommand__Go(long long id, A&&... a) : QUndoCommand(std::forward<A>(a)...), __goId(id) {}
	void undo() override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 7259214242580717359LL, &__h); if (__h) return; }
		QUndoCommand::undo();
	}
	void redo() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5544063219051948393LL, &__h); if (__h) return; }
		QUndoCommand::redo();
	}
	int id() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 628021283683842752LL, &__h); if (__h) return __r; }
		return QUndoCommand::id();
	}
	bool mergeWith(const QUndoCommand* other) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -5894931084898981809LL, (void*)other, &__h); if (__h) return __r; }
		return QUndoCommand::mergeWith(other);
	}
};

extern "C" {

void QUndoCommand__set_override(void* self, int idx) { ((QUndoCommand__Go*)self)->__goOn[idx] = 1; }
void* QUndoCommand_new(long long __goid) {
	return new QUndoCommand__Go(__goid);
}
void* QUndoCommand_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QUndoCommand__Go(__goid, text_qstr);
}
void* QUndoCommand_new3(long long __goid, void* parent) {
	return new QUndoCommand__Go(__goid, (QUndoCommand*)parent);
}
void* QUndoCommand_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QUndoCommand__Go(__goid, text_qstr, (QUndoCommand*)parent);
}
void QUndoCommand_undo(void* self) {
	((QUndoCommand*)self)->undo();
}
void QUndoCommand_redo(void* self) {
	((QUndoCommand*)self)->redo();
}
char* QUndoCommand_text(void* self, long long* __out_len) {
	auto&& __r = ((QUndoCommand*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUndoCommand_actionText(void* self, long long* __out_len) {
	auto&& __r = ((QUndoCommand*)self)->actionText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUndoCommand_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QUndoCommand*)self)->setText(text_qstr);
}
int QUndoCommand_isObsolete(void* self) {
	auto&& __r = ((QUndoCommand*)self)->isObsolete();
	return __r;
}
void QUndoCommand_setObsolete(void* self, int obsolete) {
	((QUndoCommand*)self)->setObsolete(obsolete);
}
int QUndoCommand_id(void* self) {
	auto&& __r = ((QUndoCommand*)self)->id();
	return __r;
}
int QUndoCommand_mergeWith(void* self, void* other) {
	auto&& __r = ((QUndoCommand*)self)->mergeWith((QUndoCommand*)other);
	return __r;
}
int QUndoCommand_childCount(void* self) {
	auto&& __r = ((QUndoCommand*)self)->childCount();
	return __r;
}
void* QUndoCommand_child(void* self, int index) {
	auto&& __r = ((QUndoCommand*)self)->child(index);
	return (void*)(__r);
}
void QUndoCommand_undo_default(void* self) {
	((QUndoCommand*)self)->QUndoCommand::undo();
}
void QUndoCommand_redo_default(void* self) {
	((QUndoCommand*)self)->QUndoCommand::redo();
}
int QUndoCommand_id_default(void* self) {
	auto&& __r = ((QUndoCommand*)self)->QUndoCommand::id();
	return __r;
}
int QUndoCommand_mergeWith_default(void* self, void* other) {
	auto&& __r = ((QUndoCommand*)self)->QUndoCommand::mergeWith((QUndoCommand*)other);
	return __r;
}
void* QUndoCommand__Go_new(long long id) { return new QUndoCommand__Go(id); }
void QUndoCommand_delete(void* self) { delete (QUndoCommand*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qicon.h>
#include "gen_qstyleoptionheader.h"

extern "C" {

void* QStyleOptionHeader_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionHeader*)self); }
void* QStyleOptionHeader_new() {
	return new QStyleOptionHeader();
}
void* QStyleOptionHeader_new2(void* other) {
	return new QStyleOptionHeader(*(QStyleOptionHeader*)other);
}
void* QStyleOptionHeader_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionHeader*)self)->operator=(*(QStyleOptionHeader*)arg1);
	return (void*)&(__r);
}
int QStyleOptionHeader_section_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->section;
	return __r;
}
void QStyleOptionHeader_section_set(void* self, int section) {
	((QStyleOptionHeader*)self)->section = section;
}
char* QStyleOptionHeader_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionHeader*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionHeader_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionHeader*)self)->text = text_qstr;
}
int QStyleOptionHeader_textAlignment_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->textAlignment;
	return (int)(__r);
}
void QStyleOptionHeader_textAlignment_set(void* self, int textAlignment) {
	((QStyleOptionHeader*)self)->textAlignment = QFlags<Qt::AlignmentFlag>(textAlignment);
}
void* QStyleOptionHeader_icon_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionHeader_icon_set(void* self, void* icon) {
	((QStyleOptionHeader*)self)->icon = *(QIcon*)icon;
}
int QStyleOptionHeader_iconAlignment_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->iconAlignment;
	return (int)(__r);
}
void QStyleOptionHeader_iconAlignment_set(void* self, int iconAlignment) {
	((QStyleOptionHeader*)self)->iconAlignment = QFlags<Qt::AlignmentFlag>(iconAlignment);
}
int QStyleOptionHeader_position_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->position;
	return (int)(__r);
}
void QStyleOptionHeader_position_set(void* self, int position) {
	((QStyleOptionHeader*)self)->position = (QStyleOptionHeader::SectionPosition)(position);
}
int QStyleOptionHeader_selectedPosition_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->selectedPosition;
	return (int)(__r);
}
void QStyleOptionHeader_selectedPosition_set(void* self, int selectedPosition) {
	((QStyleOptionHeader*)self)->selectedPosition = (QStyleOptionHeader::SelectedPosition)(selectedPosition);
}
int QStyleOptionHeader_sortIndicator_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->sortIndicator;
	return (int)(__r);
}
void QStyleOptionHeader_sortIndicator_set(void* self, int sortIndicator) {
	((QStyleOptionHeader*)self)->sortIndicator = (QStyleOptionHeader::SortIndicator)(sortIndicator);
}
int QStyleOptionHeader_orientation_get(void* self) {
	auto __r = ((QStyleOptionHeader*)self)->orientation;
	return (int)(__r);
}
void QStyleOptionHeader_orientation_set(void* self, int orientation) {
	((QStyleOptionHeader*)self)->orientation = (Qt::Orientation)(orientation);
}
void QStyleOptionHeader_delete(void* self) { delete (QStyleOptionHeader*)self; }

}

#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstractitemdelegate.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractItemDelegate__Go : public QAbstractItemDelegate {
public:
	long long __goId;
	unsigned char __goOn[18] = {};
	template<typename... A> QAbstractItemDelegate__Go(long long id, A&&... a) : QAbstractItemDelegate(std::forward<A>(a)...), __goId(id) {}
	using QAbstractItemDelegate::sender;
	using QAbstractItemDelegate::senderSignalIndex;
	using QAbstractItemDelegate::receivers;
	using QAbstractItemDelegate::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractItemDelegate::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractItemDelegate::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractItemDelegate::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractItemDelegate::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractItemDelegate::disconnectNotify(__a0); }
	void paint(QPainter* painter, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)&option, (void*)&index, &__h); if (__h) return; }
		return;
	}
	QSize sizeHint(const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, int*))cppgo_virt_voidp_voidp_voidp)(__goId, 729690261457684101LL, (void*)&option, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractItemDelegate::metaObject();
	}
	QWidget* createEditor(QWidget* parent, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_voidp_voidp_voidp_voidp)(__goId, -2118980437705941456LL, (void*)parent, (void*)&option, (void*)&index, &__h); if (__h) return (QWidget*)__r; }
		return QAbstractItemDelegate::createEditor(parent, option, index);
	}
	void destroyEditor(QWidget* editor, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -7699250917160078568LL, (void*)editor, (void*)&index, &__h); if (__h) return; }
		QAbstractItemDelegate::destroyEditor(editor, index);
	}
	void setEditorData(QWidget* editor, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -858237687273351194LL, (void*)editor, (void*)&index, &__h); if (__h) return; }
		QAbstractItemDelegate::setEditorData(editor, index);
	}
	void setModelData(QWidget* editor, QAbstractItemModel* model, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -4740900545199884600LL, (void*)editor, (void*)model, (void*)&index, &__h); if (__h) return; }
		QAbstractItemDelegate::setModelData(editor, model, index);
	}
	void updateEditorGeometry(QWidget* editor, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -1005009386014896055LL, (void*)editor, (void*)&option, (void*)&index, &__h); if (__h) return; }
		QAbstractItemDelegate::updateEditorGeometry(editor, option, index);
	}
	bool editorEvent(QEvent* event, QAbstractItemModel* model, const QStyleOptionViewItem& option, const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp_voidp)(__goId, -2261279149879782144LL, (void*)event, (void*)model, (void*)&option, (void*)&index, &__h); if (__h) return __r; }
		return QAbstractItemDelegate::editorEvent(event, model, option, index);
	}
	bool helpEvent(QHelpEvent* event, QAbstractItemView* view, const QStyleOptionViewItem& option, const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp_voidp)(__goId, 1696901855134526294LL, (void*)event, (void*)view, (void*)&option, (void*)&index, &__h); if (__h) return __r; }
		return QAbstractItemDelegate::helpEvent(event, view, option, index);
	}
	QList<int> paintingRoles() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virtp_voidp) {
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
		return QAbstractItemDelegate::paintingRoles();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemDelegate::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemDelegate::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemDelegate::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemDelegate::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemDelegate::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractItemDelegate::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractItemDelegate::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractItemDelegate;
static std::mutex __connsMu_QAbstractItemDelegate;

extern "C" {

void* QAbstractItemDelegate_base(void* self) { return (void*)static_cast<QObject*>((QAbstractItemDelegate*)self); }
void QAbstractItemDelegate__set_override(void* self, int idx) { ((QAbstractItemDelegate__Go*)self)->__goOn[idx] = 1; }
void* QAbstractItemDelegate_fromQObject(void* o) { return (void*)qobject_cast<QAbstractItemDelegate*>((QObject*)o); }
char* QAbstractItemDelegate_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractItemDelegate::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractItemDelegate_paint(void* self, void* painter, void* option, void* index) {
	((QAbstractItemDelegate*)self)->paint((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
void* QAbstractItemDelegate_sizeHint(void* self, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->sizeHint(*(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return new QSize(__r);
}
void* QAbstractItemDelegate_createEditor(void* self, void* parent, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->createEditor((QWidget*)parent, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return (void*)(__r);
}
void QAbstractItemDelegate_destroyEditor(void* self, void* editor, void* index) {
	((QAbstractItemDelegate*)self)->destroyEditor((QWidget*)editor, *(QModelIndex*)index);
}
void QAbstractItemDelegate_setEditorData(void* self, void* editor, void* index) {
	((QAbstractItemDelegate*)self)->setEditorData((QWidget*)editor, *(QModelIndex*)index);
}
void QAbstractItemDelegate_setModelData(void* self, void* editor, void* model, void* index) {
	((QAbstractItemDelegate*)self)->setModelData((QWidget*)editor, (QAbstractItemModel*)model, *(QModelIndex*)index);
}
void QAbstractItemDelegate_updateEditorGeometry(void* self, void* editor, void* option, void* index) {
	((QAbstractItemDelegate*)self)->updateEditorGeometry((QWidget*)editor, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
int QAbstractItemDelegate_editorEvent(void* self, void* event, void* model, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->editorEvent((QEvent*)event, (QAbstractItemModel*)model, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
int QAbstractItemDelegate_handleEditorEvent(void* self, void* object, void* event) {
	auto&& __r = ((QAbstractItemDelegate*)self)->handleEditorEvent((QObject*)object, (QEvent*)event);
	return __r;
}
int QAbstractItemDelegate_helpEvent(void* self, void* event, void* view, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->helpEvent((QHelpEvent*)event, (QAbstractItemView*)view, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
int* QAbstractItemDelegate_paintingRoles(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractItemDelegate*)self)->paintingRoles();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QAbstractItemDelegate_commitData(void* self, void* editor) {
	((QAbstractItemDelegate*)self)->commitData((QWidget*)editor);
}
void QAbstractItemDelegate_closeEditor(void* self, void* editor) {
	((QAbstractItemDelegate*)self)->closeEditor((QWidget*)editor);
}
void QAbstractItemDelegate_sizeHintChanged(void* self, void* arg1) {
	((QAbstractItemDelegate*)self)->sizeHintChanged(*(QModelIndex*)arg1);
}
void* QAbstractItemDelegate_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractItemDelegate::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemDelegate_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractItemDelegate::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemDelegate_sender(void* self) {
	auto&& __r = ((QAbstractItemDelegate__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractItemDelegate_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractItemDelegate__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractItemDelegate_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractItemDelegate__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractItemDelegate_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractItemDelegate__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractItemDelegate_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractItemDelegate::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractItemDelegate_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractItemDelegate::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractItemDelegate_closeEditor2(void* self, void* editor, int hint) {
	((QAbstractItemDelegate*)self)->closeEditor((QWidget*)editor, (QAbstractItemDelegate::EndEditHint)(hint));
}
void* QAbstractItemDelegate_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractItemDelegate::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemDelegate_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractItemDelegate::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemDelegate_createEditor_default(void* self, void* parent, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::createEditor((QWidget*)parent, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return (void*)(__r);
}
void QAbstractItemDelegate_destroyEditor_default(void* self, void* editor, void* index) {
	((QAbstractItemDelegate*)self)->QAbstractItemDelegate::destroyEditor((QWidget*)editor, *(QModelIndex*)index);
}
void QAbstractItemDelegate_setEditorData_default(void* self, void* editor, void* index) {
	((QAbstractItemDelegate*)self)->QAbstractItemDelegate::setEditorData((QWidget*)editor, *(QModelIndex*)index);
}
void QAbstractItemDelegate_setModelData_default(void* self, void* editor, void* model, void* index) {
	((QAbstractItemDelegate*)self)->QAbstractItemDelegate::setModelData((QWidget*)editor, (QAbstractItemModel*)model, *(QModelIndex*)index);
}
void QAbstractItemDelegate_updateEditorGeometry_default(void* self, void* editor, void* option, void* index) {
	((QAbstractItemDelegate*)self)->QAbstractItemDelegate::updateEditorGeometry((QWidget*)editor, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
int QAbstractItemDelegate_editorEvent_default(void* self, void* event, void* model, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::editorEvent((QEvent*)event, (QAbstractItemModel*)model, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
int QAbstractItemDelegate_helpEvent_default(void* self, void* event, void* view, void* option, void* index) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::helpEvent((QHelpEvent*)event, (QAbstractItemView*)view, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
int* QAbstractItemDelegate_paintingRoles_default(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::paintingRoles();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void* QAbstractItemDelegate_metaObject_default(void* self) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::metaObject();
	return (void*)(__r);
}
int QAbstractItemDelegate_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::event((QEvent*)event);
	return __r;
}
int QAbstractItemDelegate_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractItemDelegate*)self)->QAbstractItemDelegate::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractItemDelegate_timerEvent_default(void* self, void* event) {
	((QAbstractItemDelegate__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractItemDelegate_childEvent_default(void* self, void* event) {
	((QAbstractItemDelegate__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractItemDelegate_customEvent_default(void* self, void* event) {
	((QAbstractItemDelegate__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractItemDelegate_connectNotify_default(void* self, void* signal) {
	((QAbstractItemDelegate__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractItemDelegate_disconnectNotify_default(void* self, void* signal) {
	((QAbstractItemDelegate__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractItemDelegate_connect_commitData(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::commitData, (QAbstractItemDelegate*)self, [__id](QWidget* editor){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)editor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate_connect_closeEditor(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::closeEditor, (QAbstractItemDelegate*)self, [__id](QWidget* editor, QAbstractItemDelegate::EndEditHint hint){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)editor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate_connect_closeEditor2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::closeEditor, (QAbstractItemDelegate*)self, [__id](QWidget* editor, QAbstractItemDelegate::EndEditHint hint){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)editor, (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate_connect_sizeHintChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::sizeHintChanged, (QAbstractItemDelegate*)self, [__id](const QModelIndex& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::destroyed, (QAbstractItemDelegate*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::destroyed, (QAbstractItemDelegate*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemDelegate*)self, &QAbstractItemDelegate::objectNameChanged, (QAbstractItemDelegate*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
	__conns_QAbstractItemDelegate[__id] = __c;
}
void QAbstractItemDelegate__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemDelegate);
		auto __it = __conns_QAbstractItemDelegate.find(id);
		if (__it == __conns_QAbstractItemDelegate.end()) return;
		__c = __it->second;
		__conns_QAbstractItemDelegate.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractItemDelegate__Go_new(long long id, void* parent) { return new QAbstractItemDelegate__Go(id, (QObject*)parent); }
void QAbstractItemDelegate_delete(void* self) { delete (QAbstractItemDelegate*)self; }

}

#include <QtWidgets/qcommandlinkbutton.h>
#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qpoint.h>
#include <QtGui/qicon.h>
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
#include "gen_qcommandlinkbutton.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QCommandLinkButton__Go : public QCommandLinkButton {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QCommandLinkButton__Go(long long id, A&&... a) : QCommandLinkButton(std::forward<A>(a)...), __goId(id) {}
	using QCommandLinkButton::updateMicroFocus;
	using QCommandLinkButton::create;
	using QCommandLinkButton::destroy;
	using QCommandLinkButton::focusNextChild;
	using QCommandLinkButton::focusPreviousChild;
	using QCommandLinkButton::sender;
	using QCommandLinkButton::senderSignalIndex;
	using QCommandLinkButton::receivers;
	using QCommandLinkButton::isSignalConnected;
	using QCommandLinkButton::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QCommandLinkButton::event(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QCommandLinkButton::paintEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QCommandLinkButton::keyPressEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QCommandLinkButton::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QCommandLinkButton::focusOutEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QCommandLinkButton::mouseMoveEvent(__a0); }
	bool cppgoBase_hitButton(const QPoint& __a0) { return QCommandLinkButton::hitButton(__a0); }
	void cppgoBase_checkStateSet() { QCommandLinkButton::checkStateSet(); }
	void cppgoBase_nextCheckState() { QCommandLinkButton::nextCheckState(); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QCommandLinkButton::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QCommandLinkButton::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QCommandLinkButton::mouseReleaseEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QCommandLinkButton::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QCommandLinkButton::timerEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QCommandLinkButton::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QCommandLinkButton::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QCommandLinkButton::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QCommandLinkButton::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QCommandLinkButton::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QCommandLinkButton::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QCommandLinkButton::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QCommandLinkButton::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QCommandLinkButton::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QCommandLinkButton::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QCommandLinkButton::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QCommandLinkButton::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QCommandLinkButton::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QCommandLinkButton::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QCommandLinkButton::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QCommandLinkButton::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QCommandLinkButton::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QCommandLinkButton::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QCommandLinkButton::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QCommandLinkButton::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QCommandLinkButton::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QCommandLinkButton::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QCommandLinkButton::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QCommandLinkButton::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QCommandLinkButton::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QCommandLinkButton::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QCommandLinkButton::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QCommandLinkButton::sizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QCommandLinkButton::heightForWidth(arg1);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QCommandLinkButton::minimumSizeHint();
	}
	void initStyleOption(QStyleOptionButton* option) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QCommandLinkButton::initStyleOption(option);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QCommandLinkButton::event(e);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QCommandLinkButton::paintEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QCommandLinkButton::keyPressEvent(arg1);
	}
	void focusInEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)arg1, &__h); if (__h) return; }
		QCommandLinkButton::focusInEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QCommandLinkButton::focusOutEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QCommandLinkButton::mouseMoveEvent(arg1);
	}
	bool hitButton(const QPoint& pos) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5736040353629178930LL, (void*)&pos, &__h); if (__h) return __r; }
		return QCommandLinkButton::hitButton(pos);
	}
	void checkStateSet() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5747806663224476200LL, &__h); if (__h) return; }
		QCommandLinkButton::checkStateSet();
	}
	void nextCheckState() override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1474647914901213797LL, &__h); if (__h) return; }
		QCommandLinkButton::nextCheckState();
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QCommandLinkButton::keyReleaseEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QCommandLinkButton::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QCommandLinkButton::mouseReleaseEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QCommandLinkButton::changeEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QCommandLinkButton::timerEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QCommandLinkButton::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QCommandLinkButton::setVisible(visible);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QCommandLinkButton::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QCommandLinkButton::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::wheelEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QCommandLinkButton::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QCommandLinkButton::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QCommandLinkButton::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QCommandLinkButton::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QCommandLinkButton::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QCommandLinkButton::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QCommandLinkButton::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QCommandLinkButton::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QCommandLinkButton::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QCommandLinkButton::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QCommandLinkButton::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QCommandLinkButton::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QCommandLinkButton;
static std::mutex __connsMu_QCommandLinkButton;

extern "C" {

void* QCommandLinkButton_base(void* self) { return (void*)static_cast<QPushButton*>((QCommandLinkButton*)self); }
void QCommandLinkButton_chain(void* self, void** out) {
	QCommandLinkButton* __p = (QCommandLinkButton*)self;
	QPushButton* __b0 = static_cast<QPushButton*>(__p); out[0] = (void*)__b0;
	QAbstractButton* __b1 = static_cast<QAbstractButton*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QCommandLinkButton__set_override(void* self, int idx) { ((QCommandLinkButton__Go*)self)->__goOn[idx] = 1; }
void* QCommandLinkButton_fromQObject(void* o) { return (void*)qobject_cast<QCommandLinkButton*>((QObject*)o); }
void* QCommandLinkButton_new(long long __goid, void* parent) {
	return new QCommandLinkButton__Go(__goid, (QWidget*)parent);
}
void* QCommandLinkButton_new2(long long __goid) {
	return new QCommandLinkButton__Go(__goid);
}
void* QCommandLinkButton_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QCommandLinkButton__Go(__goid, text_qstr);
}
void* QCommandLinkButton_new4(long long __goid, const char* text, long long text_len, const char* description, long long description_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	return new QCommandLinkButton__Go(__goid, text_qstr, description_qstr);
}
void* QCommandLinkButton_new5(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QCommandLinkButton__Go(__goid, text_qstr, (QWidget*)parent);
}
void* QCommandLinkButton_new6(long long __goid, const char* text, long long text_len, const char* description, long long description_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	return new QCommandLinkButton__Go(__goid, text_qstr, description_qstr, (QWidget*)parent);
}
char* QCommandLinkButton_tr(const char* s, long long* __out_len) {
	auto&& __r = QCommandLinkButton::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCommandLinkButton_description(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLinkButton*)self)->description();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLinkButton_setDescription(void* self, const char* description, long long description_len) {
	QString description_qstr = QString::fromUtf8(description, description_len);
	((QCommandLinkButton*)self)->setDescription(description_qstr);
}
void QCommandLinkButton_setTabOrder(void* arg1, void* arg2) {
	QCommandLinkButton::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QCommandLinkButton_mouseGrabber() {
	auto&& __r = QCommandLinkButton::mouseGrabber();
	return (void*)(__r);
}
void* QCommandLinkButton_keyboardGrabber() {
	auto&& __r = QCommandLinkButton::keyboardGrabber();
	return (void*)(__r);
}
void* QCommandLinkButton_find(unsigned long long arg1) {
	auto&& __r = QCommandLinkButton::find(arg1);
	return (void*)(__r);
}
void* QCommandLinkButton_createWindowContainer(void* window) {
	auto&& __r = QCommandLinkButton::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QCommandLinkButton_updateMicroFocus(void* self) {
	((QCommandLinkButton__Go*)self)->updateMicroFocus();
}
void QCommandLinkButton_create(void* self) {
	((QCommandLinkButton__Go*)self)->create();
}
void QCommandLinkButton_destroy(void* self) {
	((QCommandLinkButton__Go*)self)->destroy();
}
int QCommandLinkButton_focusNextChild(void* self) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->focusNextChild();
	return __r;
}
int QCommandLinkButton_focusPreviousChild(void* self) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->focusPreviousChild();
	return __r;
}
void* QCommandLinkButton_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QCommandLinkButton::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QCommandLinkButton_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QCommandLinkButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QCommandLinkButton_sender(void* self) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->sender();
	return (void*)(__r);
}
int QCommandLinkButton_senderSignalIndex(void* self) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->senderSignalIndex();
	return __r;
}
int QCommandLinkButton_receivers(void* self, const char* signal) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->receivers(signal);
	return __r;
}
int QCommandLinkButton_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QCommandLinkButton_devicePixelRatioFScale() {
	auto&& __r = QCommandLinkButton::devicePixelRatioFScale();
	return __r;
}
int QCommandLinkButton_encodeMetricF(int metric, double value) {
	auto&& __r = QCommandLinkButton::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QCommandLinkButton_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QCommandLinkButton_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QCommandLinkButton::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCommandLinkButton_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QCommandLinkButton::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCommandLinkButton_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QCommandLinkButton::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QCommandLinkButton_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QCommandLinkButton::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QCommandLinkButton_updateMicroFocus2(void* self, int query) {
	((QCommandLinkButton__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QCommandLinkButton_create2(void* self, unsigned long long arg1) {
	((QCommandLinkButton__Go*)self)->create(arg1);
}
void QCommandLinkButton_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QCommandLinkButton__Go*)self)->create(arg1, initializeWindow);
}
void QCommandLinkButton_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QCommandLinkButton__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QCommandLinkButton_destroy2(void* self, int destroyWindow) {
	((QCommandLinkButton__Go*)self)->destroy(destroyWindow);
}
void QCommandLinkButton_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QCommandLinkButton__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QCommandLinkButton_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QCommandLinkButton::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QCommandLinkButton_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QCommandLinkButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QCommandLinkButton_metaObject_default(void* self) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::metaObject();
	return (void*)(__r);
}
void* QCommandLinkButton_sizeHint_default(void* self) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::sizeHint();
	return new QSize(__r);
}
int QCommandLinkButton_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::heightForWidth(arg1);
	return __r;
}
void* QCommandLinkButton_minimumSizeHint_default(void* self) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::minimumSizeHint();
	return new QSize(__r);
}
void QCommandLinkButton_initStyleOption_default(void* self, void* option) {
	((QCommandLinkButton*)self)->QCommandLinkButton::initStyleOption((QStyleOptionButton*)option);
}
int QCommandLinkButton_event_default(void* self, void* e) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QCommandLinkButton_paintEvent_default(void* self, void* arg1) {
	((QCommandLinkButton__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QCommandLinkButton_keyPressEvent_default(void* self, void* arg1) {
	((QCommandLinkButton__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QCommandLinkButton_focusInEvent_default(void* self, void* arg1) {
	((QCommandLinkButton__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)arg1);
}
void QCommandLinkButton_focusOutEvent_default(void* self, void* arg1) {
	((QCommandLinkButton__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
void QCommandLinkButton_mouseMoveEvent_default(void* self, void* arg1) {
	((QCommandLinkButton__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
int QCommandLinkButton_hitButton_default(void* self, void* pos) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->cppgoBase_hitButton(*(QPoint*)pos);
	return __r;
}
void QCommandLinkButton_checkStateSet_default(void* self) {
	((QCommandLinkButton__Go*)self)->cppgoBase_checkStateSet();
}
void QCommandLinkButton_nextCheckState_default(void* self) {
	((QCommandLinkButton__Go*)self)->cppgoBase_nextCheckState();
}
void QCommandLinkButton_keyReleaseEvent_default(void* self, void* e) {
	((QCommandLinkButton__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QCommandLinkButton_mousePressEvent_default(void* self, void* e) {
	((QCommandLinkButton__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QCommandLinkButton_mouseReleaseEvent_default(void* self, void* e) {
	((QCommandLinkButton__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QCommandLinkButton_changeEvent_default(void* self, void* e) {
	((QCommandLinkButton__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QCommandLinkButton_timerEvent_default(void* self, void* e) {
	((QCommandLinkButton__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
int QCommandLinkButton_devType_default(void* self) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::devType();
	return __r;
}
void QCommandLinkButton_setVisible_default(void* self, int visible) {
	((QCommandLinkButton*)self)->QCommandLinkButton::setVisible(visible);
}
int QCommandLinkButton_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::hasHeightForWidth();
	return __r;
}
void* QCommandLinkButton_paintEngine_default(void* self) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::paintEngine();
	return (void*)(__r);
}
void QCommandLinkButton_mouseDoubleClickEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QCommandLinkButton_wheelEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QCommandLinkButton_enterEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QCommandLinkButton_leaveEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QCommandLinkButton_moveEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QCommandLinkButton_resizeEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QCommandLinkButton_closeEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QCommandLinkButton_contextMenuEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QCommandLinkButton_tabletEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QCommandLinkButton_actionEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QCommandLinkButton_dragEnterEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QCommandLinkButton_dragMoveEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QCommandLinkButton_dragLeaveEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QCommandLinkButton_dropEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QCommandLinkButton_showEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QCommandLinkButton_hideEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QCommandLinkButton_metric_default(void* self, int arg1) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QCommandLinkButton_initPainter_default(void* self, void* painter) {
	((QCommandLinkButton__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QCommandLinkButton_redirected_default(void* self, void* offset) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QCommandLinkButton_sharedPainter_default(void* self) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QCommandLinkButton_inputMethodEvent_default(void* self, void* arg1) {
	((QCommandLinkButton__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QCommandLinkButton_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QCommandLinkButton_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QCommandLinkButton__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QCommandLinkButton_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QCommandLinkButton*)self)->QCommandLinkButton::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QCommandLinkButton_childEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QCommandLinkButton_customEvent_default(void* self, void* event) {
	((QCommandLinkButton__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QCommandLinkButton_connectNotify_default(void* self, void* signal) {
	((QCommandLinkButton__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QCommandLinkButton_disconnectNotify_default(void* self, void* signal) {
	((QCommandLinkButton__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QCommandLinkButton_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::pressed, (QCommandLinkButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_released(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::released, (QCommandLinkButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::clicked, (QCommandLinkButton*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::clicked, (QCommandLinkButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::toggled, (QCommandLinkButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::windowTitleChanged, (QCommandLinkButton*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::windowIconChanged, (QCommandLinkButton*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::windowIconTextChanged, (QCommandLinkButton*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::customContextMenuRequested, (QCommandLinkButton*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::destroyed, (QCommandLinkButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::destroyed, (QCommandLinkButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommandLinkButton*)self, &QCommandLinkButton::objectNameChanged, (QCommandLinkButton*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
	__conns_QCommandLinkButton[__id] = __c;
}
void QCommandLinkButton__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QCommandLinkButton);
		auto __it = __conns_QCommandLinkButton.find(id);
		if (__it == __conns_QCommandLinkButton.end()) return;
		__c = __it->second;
		__conns_QCommandLinkButton.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QCommandLinkButton__Go_new(long long id) { return new QCommandLinkButton__Go(id); }
void QCommandLinkButton_delete(void* self) { delete (QCommandLinkButton*)self; }

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
#include "gen_qtapandholdgesture.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTapAndHoldGesture__Go : public QTapAndHoldGesture {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QTapAndHoldGesture__Go(long long id, A&&... a) : QTapAndHoldGesture(std::forward<A>(a)...), __goId(id) {}
	using QTapAndHoldGesture::sender;
	using QTapAndHoldGesture::senderSignalIndex;
	using QTapAndHoldGesture::receivers;
	using QTapAndHoldGesture::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTapAndHoldGesture::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTapAndHoldGesture::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTapAndHoldGesture::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTapAndHoldGesture::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTapAndHoldGesture::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTapAndHoldGesture::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTapAndHoldGesture::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTapAndHoldGesture::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTapAndHoldGesture::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTapAndHoldGesture::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTapAndHoldGesture::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTapAndHoldGesture::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTapAndHoldGesture::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTapAndHoldGesture;
static std::mutex __connsMu_QTapAndHoldGesture;

extern "C" {

void* QTapAndHoldGesture_base(void* self) { return (void*)static_cast<QGesture*>((QTapAndHoldGesture*)self); }
void QTapAndHoldGesture_chain(void* self, void** out) {
	QTapAndHoldGesture* __p = (QTapAndHoldGesture*)self;
	QGesture* __b0 = static_cast<QGesture*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QTapAndHoldGesture__set_override(void* self, int idx) { ((QTapAndHoldGesture__Go*)self)->__goOn[idx] = 1; }
void* QTapAndHoldGesture_fromQObject(void* o) { return (void*)qobject_cast<QTapAndHoldGesture*>((QObject*)o); }
void* QTapAndHoldGesture_new(long long __goid) {
	return new QTapAndHoldGesture__Go(__goid);
}
void* QTapAndHoldGesture_new2(long long __goid, void* parent) {
	return new QTapAndHoldGesture__Go(__goid, (QObject*)parent);
}
char* QTapAndHoldGesture_tr(const char* s, long long* __out_len) {
	auto&& __r = QTapAndHoldGesture::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTapAndHoldGesture_position(void* self) {
	auto&& __r = ((QTapAndHoldGesture*)self)->position();
	return new QPointF(__r);
}
void QTapAndHoldGesture_setPosition(void* self, void* pos) {
	((QTapAndHoldGesture*)self)->setPosition(*(QPointF*)pos);
}
void QTapAndHoldGesture_setTimeout(int msecs) {
	QTapAndHoldGesture::setTimeout(msecs);
}
int QTapAndHoldGesture_timeout() {
	auto&& __r = QTapAndHoldGesture::timeout();
	return __r;
}
void* QTapAndHoldGesture_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTapAndHoldGesture::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTapAndHoldGesture_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTapAndHoldGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTapAndHoldGesture_sender(void* self) {
	auto&& __r = ((QTapAndHoldGesture__Go*)self)->sender();
	return (void*)(__r);
}
int QTapAndHoldGesture_senderSignalIndex(void* self) {
	auto&& __r = ((QTapAndHoldGesture__Go*)self)->senderSignalIndex();
	return __r;
}
int QTapAndHoldGesture_receivers(void* self, const char* signal) {
	auto&& __r = ((QTapAndHoldGesture__Go*)self)->receivers(signal);
	return __r;
}
int QTapAndHoldGesture_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTapAndHoldGesture__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTapAndHoldGesture_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTapAndHoldGesture::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTapAndHoldGesture_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTapAndHoldGesture::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTapAndHoldGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTapAndHoldGesture::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTapAndHoldGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTapAndHoldGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTapAndHoldGesture_metaObject_default(void* self) {
	auto&& __r = ((QTapAndHoldGesture*)self)->QTapAndHoldGesture::metaObject();
	return (void*)(__r);
}
int QTapAndHoldGesture_event_default(void* self, void* event) {
	auto&& __r = ((QTapAndHoldGesture*)self)->QTapAndHoldGesture::event((QEvent*)event);
	return __r;
}
int QTapAndHoldGesture_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTapAndHoldGesture*)self)->QTapAndHoldGesture::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTapAndHoldGesture_timerEvent_default(void* self, void* event) {
	((QTapAndHoldGesture__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTapAndHoldGesture_childEvent_default(void* self, void* event) {
	((QTapAndHoldGesture__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTapAndHoldGesture_customEvent_default(void* self, void* event) {
	((QTapAndHoldGesture__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTapAndHoldGesture_connectNotify_default(void* self, void* signal) {
	((QTapAndHoldGesture__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTapAndHoldGesture_disconnectNotify_default(void* self, void* signal) {
	((QTapAndHoldGesture__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTapAndHoldGesture_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTapAndHoldGesture*)self, &QTapAndHoldGesture::destroyed, (QTapAndHoldGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTapAndHoldGesture);
	__conns_QTapAndHoldGesture[__id] = __c;
}
void QTapAndHoldGesture_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTapAndHoldGesture*)self, &QTapAndHoldGesture::destroyed, (QTapAndHoldGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTapAndHoldGesture);
	__conns_QTapAndHoldGesture[__id] = __c;
}
void QTapAndHoldGesture_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTapAndHoldGesture*)self, &QTapAndHoldGesture::objectNameChanged, (QTapAndHoldGesture*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTapAndHoldGesture);
	__conns_QTapAndHoldGesture[__id] = __c;
}
void QTapAndHoldGesture__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTapAndHoldGesture);
		auto __it = __conns_QTapAndHoldGesture.find(id);
		if (__it == __conns_QTapAndHoldGesture.end()) return;
		__c = __it->second;
		__conns_QTapAndHoldGesture.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTapAndHoldGesture__Go_new(long long id, void* parent) { return new QTapAndHoldGesture__Go(id, (QObject*)parent); }
void QTapAndHoldGesture_delete(void* self) { delete (QTapAndHoldGesture*)self; }

}

#include <QtWidgets/qgraphicsgridlayout.h>
#include <QtWidgets/qgraphicslayout.h>
#include <QtWidgets/qgraphicslayoutitem.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qgraphicsitem.h>
#include "gen_qgraphicsgridlayout.h"
#include "gen_purego_cb.h"

class QGraphicsGridLayout__Go : public QGraphicsGridLayout {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QGraphicsGridLayout__Go(long long id, A&&... a) : QGraphicsGridLayout(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsGridLayout::addChildLayoutItem;
	using QGraphicsGridLayout::setGraphicsItem;
	using QGraphicsGridLayout::setOwnedByLayout;
	int count() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QGraphicsGridLayout::count();
	}
	QGraphicsLayoutItem* itemAt(int index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, index, &__h); if (__h) return (QGraphicsLayoutItem*)__r; }
		return QGraphicsGridLayout::itemAt(index);
	}
	void removeAt(int index) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1230346564330554752LL, index, &__h); if (__h) return; }
		QGraphicsGridLayout::removeAt(index);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QGraphicsGridLayout::invalidate();
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsGridLayout::setGeometry(rect);
	}
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QGraphicsGridLayout::sizeHint(which, constraint);
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsGridLayout::getContentsMargins(left, top, right, bottom);
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsGridLayout::updateGeometry();
	}
	void widgetEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 8535746785413991679LL, (void*)e, &__h); if (__h) return; }
		QGraphicsGridLayout::widgetEvent(e);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsGridLayout::isEmpty();
	}
};

extern "C" {

void* QGraphicsGridLayout_base(void* self) { return (void*)static_cast<QGraphicsLayout*>((QGraphicsGridLayout*)self); }
void QGraphicsGridLayout_chain(void* self, void** out) {
	QGraphicsGridLayout* __p = (QGraphicsGridLayout*)self;
	QGraphicsLayout* __b0 = static_cast<QGraphicsLayout*>(__p); out[0] = (void*)__b0;
	QGraphicsLayoutItem* __b1 = static_cast<QGraphicsLayoutItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsGridLayout__set_override(void* self, int idx) { ((QGraphicsGridLayout__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsGridLayout_new(long long __goid) {
	return new QGraphicsGridLayout__Go(__goid);
}
void* QGraphicsGridLayout_new2(long long __goid, void* parent) {
	return new QGraphicsGridLayout__Go(__goid, (QGraphicsLayoutItem*)parent);
}
void QGraphicsGridLayout_addItem(void* self, void* item, int row, int column, int rowSpan, int columnSpan) {
	((QGraphicsGridLayout*)self)->addItem((QGraphicsLayoutItem*)item, row, column, rowSpan, columnSpan);
}
void QGraphicsGridLayout_addItem2(void* self, void* item, int row, int column) {
	((QGraphicsGridLayout*)self)->addItem((QGraphicsLayoutItem*)item, row, column);
}
void QGraphicsGridLayout_setHorizontalSpacing(void* self, double spacing) {
	((QGraphicsGridLayout*)self)->setHorizontalSpacing(spacing);
}
double QGraphicsGridLayout_horizontalSpacing(void* self) {
	auto&& __r = ((QGraphicsGridLayout*)self)->horizontalSpacing();
	return __r;
}
void QGraphicsGridLayout_setVerticalSpacing(void* self, double spacing) {
	((QGraphicsGridLayout*)self)->setVerticalSpacing(spacing);
}
double QGraphicsGridLayout_verticalSpacing(void* self) {
	auto&& __r = ((QGraphicsGridLayout*)self)->verticalSpacing();
	return __r;
}
void QGraphicsGridLayout_setSpacing(void* self, double spacing) {
	((QGraphicsGridLayout*)self)->setSpacing(spacing);
}
void QGraphicsGridLayout_setRowSpacing(void* self, int row, double spacing) {
	((QGraphicsGridLayout*)self)->setRowSpacing(row, spacing);
}
double QGraphicsGridLayout_rowSpacing(void* self, int row) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowSpacing(row);
	return __r;
}
void QGraphicsGridLayout_setColumnSpacing(void* self, int column, double spacing) {
	((QGraphicsGridLayout*)self)->setColumnSpacing(column, spacing);
}
double QGraphicsGridLayout_columnSpacing(void* self, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnSpacing(column);
	return __r;
}
void QGraphicsGridLayout_setRowStretchFactor(void* self, int row, int stretch) {
	((QGraphicsGridLayout*)self)->setRowStretchFactor(row, stretch);
}
int QGraphicsGridLayout_rowStretchFactor(void* self, int row) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowStretchFactor(row);
	return __r;
}
void QGraphicsGridLayout_setColumnStretchFactor(void* self, int column, int stretch) {
	((QGraphicsGridLayout*)self)->setColumnStretchFactor(column, stretch);
}
int QGraphicsGridLayout_columnStretchFactor(void* self, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnStretchFactor(column);
	return __r;
}
void QGraphicsGridLayout_setRowMinimumHeight(void* self, int row, double height) {
	((QGraphicsGridLayout*)self)->setRowMinimumHeight(row, height);
}
double QGraphicsGridLayout_rowMinimumHeight(void* self, int row) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowMinimumHeight(row);
	return __r;
}
void QGraphicsGridLayout_setRowPreferredHeight(void* self, int row, double height) {
	((QGraphicsGridLayout*)self)->setRowPreferredHeight(row, height);
}
double QGraphicsGridLayout_rowPreferredHeight(void* self, int row) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowPreferredHeight(row);
	return __r;
}
void QGraphicsGridLayout_setRowMaximumHeight(void* self, int row, double height) {
	((QGraphicsGridLayout*)self)->setRowMaximumHeight(row, height);
}
double QGraphicsGridLayout_rowMaximumHeight(void* self, int row) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowMaximumHeight(row);
	return __r;
}
void QGraphicsGridLayout_setRowFixedHeight(void* self, int row, double height) {
	((QGraphicsGridLayout*)self)->setRowFixedHeight(row, height);
}
void QGraphicsGridLayout_setColumnMinimumWidth(void* self, int column, double width) {
	((QGraphicsGridLayout*)self)->setColumnMinimumWidth(column, width);
}
double QGraphicsGridLayout_columnMinimumWidth(void* self, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnMinimumWidth(column);
	return __r;
}
void QGraphicsGridLayout_setColumnPreferredWidth(void* self, int column, double width) {
	((QGraphicsGridLayout*)self)->setColumnPreferredWidth(column, width);
}
double QGraphicsGridLayout_columnPreferredWidth(void* self, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnPreferredWidth(column);
	return __r;
}
void QGraphicsGridLayout_setColumnMaximumWidth(void* self, int column, double width) {
	((QGraphicsGridLayout*)self)->setColumnMaximumWidth(column, width);
}
double QGraphicsGridLayout_columnMaximumWidth(void* self, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnMaximumWidth(column);
	return __r;
}
void QGraphicsGridLayout_setColumnFixedWidth(void* self, int column, double width) {
	((QGraphicsGridLayout*)self)->setColumnFixedWidth(column, width);
}
void QGraphicsGridLayout_setRowAlignment(void* self, int row, int alignment) {
	((QGraphicsGridLayout*)self)->setRowAlignment(row, QFlags<Qt::AlignmentFlag>(alignment));
}
int QGraphicsGridLayout_rowAlignment(void* self, int row) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowAlignment(row);
	return (int)(__r);
}
void QGraphicsGridLayout_setColumnAlignment(void* self, int column, int alignment) {
	((QGraphicsGridLayout*)self)->setColumnAlignment(column, QFlags<Qt::AlignmentFlag>(alignment));
}
int QGraphicsGridLayout_columnAlignment(void* self, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnAlignment(column);
	return (int)(__r);
}
void QGraphicsGridLayout_setAlignment(void* self, void* item, int alignment) {
	((QGraphicsGridLayout*)self)->setAlignment((QGraphicsLayoutItem*)item, QFlags<Qt::AlignmentFlag>(alignment));
}
int QGraphicsGridLayout_alignment(void* self, void* item) {
	auto&& __r = ((QGraphicsGridLayout*)self)->alignment((QGraphicsLayoutItem*)item);
	return (int)(__r);
}
int QGraphicsGridLayout_rowCount(void* self) {
	auto&& __r = ((QGraphicsGridLayout*)self)->rowCount();
	return __r;
}
int QGraphicsGridLayout_columnCount(void* self) {
	auto&& __r = ((QGraphicsGridLayout*)self)->columnCount();
	return __r;
}
void* QGraphicsGridLayout_itemAt(void* self, int row, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->itemAt(row, column);
	return (void*)(__r);
}
void* QGraphicsGridLayout_itemAt2(void* self, int index) {
	auto&& __r = ((QGraphicsGridLayout*)self)->itemAt(index);
	return (void*)(__r);
}
void QGraphicsGridLayout_removeItem(void* self, void* item) {
	((QGraphicsGridLayout*)self)->removeItem((QGraphicsLayoutItem*)item);
}
void QGraphicsGridLayout_setInstantInvalidatePropagation(int enable) {
	QGraphicsGridLayout::setInstantInvalidatePropagation(enable);
}
int QGraphicsGridLayout_instantInvalidatePropagation() {
	auto&& __r = QGraphicsGridLayout::instantInvalidatePropagation();
	return __r;
}
void QGraphicsGridLayout_addChildLayoutItem(void* self, void* layoutItem) {
	((QGraphicsGridLayout__Go*)self)->addChildLayoutItem((QGraphicsLayoutItem*)layoutItem);
}
void QGraphicsGridLayout_setGraphicsItem(void* self, void* item) {
	((QGraphicsGridLayout__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsGridLayout_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsGridLayout__Go*)self)->setOwnedByLayout(ownedByLayout);
}
void QGraphicsGridLayout_addItem3(void* self, void* item, int row, int column, int rowSpan, int columnSpan, int alignment) {
	((QGraphicsGridLayout*)self)->addItem((QGraphicsLayoutItem*)item, row, column, rowSpan, columnSpan, QFlags<Qt::AlignmentFlag>(alignment));
}
void QGraphicsGridLayout_addItem4(void* self, void* item, int row, int column, int alignment) {
	((QGraphicsGridLayout*)self)->addItem((QGraphicsLayoutItem*)item, row, column, QFlags<Qt::AlignmentFlag>(alignment));
}
void* QGraphicsGridLayout_itemAt_default(void* self, int row, int column) {
	auto&& __r = ((QGraphicsGridLayout*)self)->QGraphicsGridLayout::itemAt(row, column);
	return (void*)(__r);
}
void* QGraphicsGridLayout_itemAt_default2(void* self, int index) {
	auto&& __r = ((QGraphicsGridLayout*)self)->QGraphicsGridLayout::itemAt(index);
	return (void*)(__r);
}
int QGraphicsGridLayout_count_default(void* self) {
	auto&& __r = ((QGraphicsGridLayout*)self)->QGraphicsGridLayout::count();
	return __r;
}
void QGraphicsGridLayout_removeAt_default(void* self, int index) {
	((QGraphicsGridLayout*)self)->QGraphicsGridLayout::removeAt(index);
}
void QGraphicsGridLayout_invalidate_default(void* self) {
	((QGraphicsGridLayout*)self)->QGraphicsGridLayout::invalidate();
}
void QGraphicsGridLayout_setGeometry_default(void* self, void* rect) {
	((QGraphicsGridLayout*)self)->QGraphicsGridLayout::setGeometry(*(QRectF*)rect);
}
void* QGraphicsGridLayout_sizeHint_default(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsGridLayout*)self)->QGraphicsGridLayout::sizeHint((Qt::SizeHint)(which), *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void QGraphicsGridLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsGridLayout*)self)->QGraphicsGridLayout::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsGridLayout_updateGeometry_default(void* self) {
	((QGraphicsGridLayout*)self)->QGraphicsGridLayout::updateGeometry();
}
void QGraphicsGridLayout_widgetEvent_default(void* self, void* e) {
	((QGraphicsGridLayout*)self)->QGraphicsGridLayout::widgetEvent((QEvent*)e);
}
int QGraphicsGridLayout_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsGridLayout*)self)->QGraphicsGridLayout::isEmpty();
	return __r;
}
void* QGraphicsGridLayout__Go_new(long long id) { return new QGraphicsGridLayout__Go(id); }
void QGraphicsGridLayout_delete(void* self) { delete (QGraphicsGridLayout*)self; }

}

#include <QtWidgets/qitemdelegate.h>
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
#include <QtCore/qrect.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qvariant.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qitemdelegate.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QItemDelegate__Go : public QItemDelegate {
public:
	long long __goId;
	unsigned char __goOn[22] = {};
	template<typename... A> QItemDelegate__Go(long long id, A&&... a) : QItemDelegate(std::forward<A>(a)...), __goId(id) {}
	using QItemDelegate::drawBackground;
	using QItemDelegate::doLayout;
	using QItemDelegate::rect;
	using QItemDelegate::setOptions;
	using QItemDelegate::decoration;
	using QItemDelegate::doCheck;
	using QItemDelegate::textRectangle;
	using QItemDelegate::sender;
	using QItemDelegate::senderSignalIndex;
	using QItemDelegate::receivers;
	using QItemDelegate::isSignalConnected;
	void cppgoBase_drawDisplay(QPainter* __a0, const QStyleOptionViewItem& __a1, const QRect& __a2, const QString& __a3) { QItemDelegate::drawDisplay(__a0, __a1, __a2, __a3); }
	void cppgoBase_drawDecoration(QPainter* __a0, const QStyleOptionViewItem& __a1, const QRect& __a2, const QPixmap& __a3) { QItemDelegate::drawDecoration(__a0, __a1, __a2, __a3); }
	void cppgoBase_drawFocus(QPainter* __a0, const QStyleOptionViewItem& __a1, const QRect& __a2) { QItemDelegate::drawFocus(__a0, __a1, __a2); }
	void cppgoBase_drawCheck(QPainter* __a0, const QStyleOptionViewItem& __a1, const QRect& __a2, int __a3) { QItemDelegate::drawCheck(__a0, __a1, __a2, (Qt::CheckState)__a3); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QItemDelegate::eventFilter(__a0, __a1); }
	bool cppgoBase_editorEvent(QEvent* __a0, QAbstractItemModel* __a1, const QStyleOptionViewItem& __a2, const QModelIndex& __a3) { return QItemDelegate::editorEvent(__a0, __a1, __a2, __a3); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QItemDelegate::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QItemDelegate::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QItemDelegate::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QItemDelegate::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QItemDelegate::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QItemDelegate::metaObject();
	}
	void paint(QPainter* painter, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)&option, (void*)&index, &__h); if (__h) return; }
		QItemDelegate::paint(painter, option, index);
	}
	QSize sizeHint(const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, int*))cppgo_virt_voidp_voidp_voidp)(__goId, 729690261457684101LL, (void*)&option, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QItemDelegate::sizeHint(option, index);
	}
	QWidget* createEditor(QWidget* parent, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_voidp_voidp_voidp_voidp)(__goId, -2118980437705941456LL, (void*)parent, (void*)&option, (void*)&index, &__h); if (__h) return (QWidget*)__r; }
		return QItemDelegate::createEditor(parent, option, index);
	}
	void setEditorData(QWidget* editor, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -858237687273351194LL, (void*)editor, (void*)&index, &__h); if (__h) return; }
		QItemDelegate::setEditorData(editor, index);
	}
	void setModelData(QWidget* editor, QAbstractItemModel* model, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -4740900545199884600LL, (void*)editor, (void*)model, (void*)&index, &__h); if (__h) return; }
		QItemDelegate::setModelData(editor, model, index);
	}
	void updateEditorGeometry(QWidget* editor, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -1005009386014896055LL, (void*)editor, (void*)&option, (void*)&index, &__h); if (__h) return; }
		QItemDelegate::updateEditorGeometry(editor, option, index);
	}
	void drawDisplay(QPainter* painter, const QStyleOptionViewItem& option, const QRect& rect, const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[7] && cppgo_virt_void_voidp_voidp_voidp_constcharp_longlong) { ((void(*)(long long, long long, void*, void*, void*, const char*, long long, int*))cppgo_virt_void_voidp_voidp_voidp_constcharp_longlong)(__goId, -5011630111077710685LL, (void*)painter, (void*)&option, (void*)&rect, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return; }
		QItemDelegate::drawDisplay(painter, option, rect, text);
	}
	void drawDecoration(QPainter* painter, const QStyleOptionViewItem& option, const QRect& rect, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp_voidp)(__goId, -6476941580654430195LL, (void*)painter, (void*)&option, (void*)&rect, (void*)&pixmap, &__h); if (__h) return; }
		QItemDelegate::drawDecoration(painter, option, rect, pixmap);
	}
	void drawFocus(QPainter* painter, const QStyleOptionViewItem& option, const QRect& rect) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -5402985858130777771LL, (void*)painter, (void*)&option, (void*)&rect, &__h); if (__h) return; }
		QItemDelegate::drawFocus(painter, option, rect);
	}
	void drawCheck(QPainter* painter, const QStyleOptionViewItem& option, const QRect& rect, Qt::CheckState state) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp_voidp_voidp_int) { ((void(*)(long long, long long, void*, void*, void*, int, int*))cppgo_virt_void_voidp_voidp_voidp_int)(__goId, -8060427496448215275LL, (void*)painter, (void*)&option, (void*)&rect, (int)(state), &__h); if (__h) return; }
		QItemDelegate::drawCheck(painter, option, rect, state);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QItemDelegate::eventFilter(object, event);
	}
	bool editorEvent(QEvent* event, QAbstractItemModel* model, const QStyleOptionViewItem& option, const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp_voidp)(__goId, -2261279149879782144LL, (void*)event, (void*)model, (void*)&option, (void*)&index, &__h); if (__h) return __r; }
		return QItemDelegate::editorEvent(event, model, option, index);
	}
	void destroyEditor(QWidget* editor, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -7699250917160078568LL, (void*)editor, (void*)&index, &__h); if (__h) return; }
		QItemDelegate::destroyEditor(editor, index);
	}
	bool helpEvent(QHelpEvent* event, QAbstractItemView* view, const QStyleOptionViewItem& option, const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp_voidp)(__goId, 1696901855134526294LL, (void*)event, (void*)view, (void*)&option, (void*)&index, &__h); if (__h) return __r; }
		return QItemDelegate::helpEvent(event, view, option, index);
	}
	QList<int> paintingRoles() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virtp_voidp) {
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
		return QItemDelegate::paintingRoles();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QItemDelegate::event(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QItemDelegate::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QItemDelegate::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QItemDelegate::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QItemDelegate::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QItemDelegate::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QItemDelegate;
static std::mutex __connsMu_QItemDelegate;

extern "C" {

void* QItemDelegate_base(void* self) { return (void*)static_cast<QAbstractItemDelegate*>((QItemDelegate*)self); }
void QItemDelegate_chain(void* self, void** out) {
	QItemDelegate* __p = (QItemDelegate*)self;
	QAbstractItemDelegate* __b0 = static_cast<QAbstractItemDelegate*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QItemDelegate__set_override(void* self, int idx) { ((QItemDelegate__Go*)self)->__goOn[idx] = 1; }
void* QItemDelegate_fromQObject(void* o) { return (void*)qobject_cast<QItemDelegate*>((QObject*)o); }
void* QItemDelegate_new(long long __goid) {
	return new QItemDelegate__Go(__goid);
}
void* QItemDelegate_new2(long long __goid, void* parent) {
	return new QItemDelegate__Go(__goid, (QObject*)parent);
}
char* QItemDelegate_tr(const char* s, long long* __out_len) {
	auto&& __r = QItemDelegate::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QItemDelegate_hasClipping(void* self) {
	auto&& __r = ((QItemDelegate*)self)->hasClipping();
	return __r;
}
void QItemDelegate_setClipping(void* self, int clip) {
	((QItemDelegate*)self)->setClipping(clip);
}
void* QItemDelegate_itemEditorFactory(void* self) {
	auto&& __r = ((QItemDelegate*)self)->itemEditorFactory();
	return (void*)(__r);
}
void QItemDelegate_setItemEditorFactory(void* self, void* factory) {
	((QItemDelegate*)self)->setItemEditorFactory((QItemEditorFactory*)factory);
}
void QItemDelegate_drawBackground(void* self, void* painter, void* option, void* index) {
	((QItemDelegate__Go*)self)->drawBackground((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
void QItemDelegate_doLayout(void* self, void* option, void* checkRect, void* iconRect, void* textRect, int hint) {
	((QItemDelegate__Go*)self)->doLayout(*(QStyleOptionViewItem*)option, (QRect*)checkRect, (QRect*)iconRect, (QRect*)textRect, hint);
}
void* QItemDelegate_rect(void* self, void* option, void* index, int role) {
	auto&& __r = ((QItemDelegate__Go*)self)->rect(*(QStyleOptionViewItem*)option, *(QModelIndex*)index, role);
	return new QRect(__r);
}
void* QItemDelegate_setOptions(void* self, void* index, void* option) {
	auto&& __r = ((QItemDelegate__Go*)self)->setOptions(*(QModelIndex*)index, *(QStyleOptionViewItem*)option);
	return new QStyleOptionViewItem(__r);
}
void* QItemDelegate_decoration(void* self, void* option, void* variant) {
	auto&& __r = ((QItemDelegate__Go*)self)->decoration(*(QStyleOptionViewItem*)option, *(QVariant*)variant);
	return new QPixmap(__r);
}
void* QItemDelegate_doCheck(void* self, void* option, void* bounding, void* variant) {
	auto&& __r = ((QItemDelegate__Go*)self)->doCheck(*(QStyleOptionViewItem*)option, *(QRect*)bounding, *(QVariant*)variant);
	return new QRect(__r);
}
void* QItemDelegate_textRectangle(void* self, void* painter, void* rect, void* font, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QItemDelegate__Go*)self)->textRectangle((QPainter*)painter, *(QRect*)rect, *(QFont*)font, text_qstr);
	return new QRect(__r);
}
void* QItemDelegate_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QItemDelegate::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QItemDelegate_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QItemDelegate::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QItemDelegate_sender(void* self) {
	auto&& __r = ((QItemDelegate__Go*)self)->sender();
	return (void*)(__r);
}
int QItemDelegate_senderSignalIndex(void* self) {
	auto&& __r = ((QItemDelegate__Go*)self)->senderSignalIndex();
	return __r;
}
int QItemDelegate_receivers(void* self, const char* signal) {
	auto&& __r = ((QItemDelegate__Go*)self)->receivers(signal);
	return __r;
}
int QItemDelegate_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QItemDelegate__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QItemDelegate_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QItemDelegate::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QItemDelegate_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QItemDelegate::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QItemDelegate_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QItemDelegate::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QItemDelegate_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QItemDelegate::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QItemDelegate_metaObject_default(void* self) {
	auto&& __r = ((QItemDelegate*)self)->QItemDelegate::metaObject();
	return (void*)(__r);
}
void QItemDelegate_paint_default(void* self, void* painter, void* option, void* index) {
	((QItemDelegate*)self)->QItemDelegate::paint((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
void* QItemDelegate_sizeHint_default(void* self, void* option, void* index) {
	auto&& __r = ((QItemDelegate*)self)->QItemDelegate::sizeHint(*(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return new QSize(__r);
}
void* QItemDelegate_createEditor_default(void* self, void* parent, void* option, void* index) {
	auto&& __r = ((QItemDelegate*)self)->QItemDelegate::createEditor((QWidget*)parent, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return (void*)(__r);
}
void QItemDelegate_setEditorData_default(void* self, void* editor, void* index) {
	((QItemDelegate*)self)->QItemDelegate::setEditorData((QWidget*)editor, *(QModelIndex*)index);
}
void QItemDelegate_setModelData_default(void* self, void* editor, void* model, void* index) {
	((QItemDelegate*)self)->QItemDelegate::setModelData((QWidget*)editor, (QAbstractItemModel*)model, *(QModelIndex*)index);
}
void QItemDelegate_updateEditorGeometry_default(void* self, void* editor, void* option, void* index) {
	((QItemDelegate*)self)->QItemDelegate::updateEditorGeometry((QWidget*)editor, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
void QItemDelegate_drawDisplay_default(void* self, void* painter, void* option, void* rect, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QItemDelegate__Go*)self)->cppgoBase_drawDisplay((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QRect*)rect, text_qstr);
}
void QItemDelegate_drawDecoration_default(void* self, void* painter, void* option, void* rect, void* pixmap) {
	((QItemDelegate__Go*)self)->cppgoBase_drawDecoration((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QRect*)rect, *(QPixmap*)pixmap);
}
void QItemDelegate_drawFocus_default(void* self, void* painter, void* option, void* rect) {
	((QItemDelegate__Go*)self)->cppgoBase_drawFocus((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QRect*)rect);
}
void QItemDelegate_drawCheck_default(void* self, void* painter, void* option, void* rect, int state) {
	((QItemDelegate__Go*)self)->cppgoBase_drawCheck((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QRect*)rect, state);
}
int QItemDelegate_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QItemDelegate__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
int QItemDelegate_editorEvent_default(void* self, void* event, void* model, void* option, void* index) {
	auto&& __r = ((QItemDelegate__Go*)self)->cppgoBase_editorEvent((QEvent*)event, (QAbstractItemModel*)model, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
void QItemDelegate_destroyEditor_default(void* self, void* editor, void* index) {
	((QItemDelegate*)self)->QItemDelegate::destroyEditor((QWidget*)editor, *(QModelIndex*)index);
}
int QItemDelegate_helpEvent_default(void* self, void* event, void* view, void* option, void* index) {
	auto&& __r = ((QItemDelegate*)self)->QItemDelegate::helpEvent((QHelpEvent*)event, (QAbstractItemView*)view, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
int* QItemDelegate_paintingRoles_default(void* self, long long* __out_len) {
	auto&& __r = ((QItemDelegate*)self)->QItemDelegate::paintingRoles();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int QItemDelegate_event_default(void* self, void* event) {
	auto&& __r = ((QItemDelegate*)self)->QItemDelegate::event((QEvent*)event);
	return __r;
}
void QItemDelegate_timerEvent_default(void* self, void* event) {
	((QItemDelegate__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QItemDelegate_childEvent_default(void* self, void* event) {
	((QItemDelegate__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QItemDelegate_customEvent_default(void* self, void* event) {
	((QItemDelegate__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QItemDelegate_connectNotify_default(void* self, void* signal) {
	((QItemDelegate__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QItemDelegate_disconnectNotify_default(void* self, void* signal) {
	((QItemDelegate__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QItemDelegate_connect_commitData(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::commitData, (QItemDelegate*)self, [__id](QWidget* editor){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)editor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate_connect_closeEditor(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::closeEditor, (QItemDelegate*)self, [__id](QWidget* editor, QAbstractItemDelegate::EndEditHint hint){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)editor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate_connect_closeEditor2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::closeEditor, (QItemDelegate*)self, [__id](QWidget* editor, QAbstractItemDelegate::EndEditHint hint){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)editor, (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate_connect_sizeHintChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::sizeHintChanged, (QItemDelegate*)self, [__id](const QModelIndex& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::destroyed, (QItemDelegate*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::destroyed, (QItemDelegate*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QItemDelegate*)self, &QItemDelegate::objectNameChanged, (QItemDelegate*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
	__conns_QItemDelegate[__id] = __c;
}
void QItemDelegate__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QItemDelegate);
		auto __it = __conns_QItemDelegate.find(id);
		if (__it == __conns_QItemDelegate.end()) return;
		__c = __it->second;
		__conns_QItemDelegate.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QItemDelegate__Go_new(long long id, void* parent) { return new QItemDelegate__Go(id, (QObject*)parent); }
void QItemDelegate_delete(void* self) { delete (QItemDelegate*)self; }

}

#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractslider.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qpoint.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
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
#include "gen_qscrollbar.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QScrollBar__Go : public QScrollBar {
public:
	long long __goId;
	unsigned char __goOn[50] = {};
	template<typename... A> QScrollBar__Go(long long id, A&&... a) : QScrollBar(std::forward<A>(a)...), __goId(id) {}
	using QScrollBar::setRepeatAction;
	using QScrollBar::repeatAction;
	using QScrollBar::updateMicroFocus;
	using QScrollBar::create;
	using QScrollBar::destroy;
	using QScrollBar::focusNextChild;
	using QScrollBar::focusPreviousChild;
	using QScrollBar::sender;
	using QScrollBar::senderSignalIndex;
	using QScrollBar::receivers;
	using QScrollBar::isSignalConnected;
	using QScrollBar::getDecodedMetricF;
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QScrollBar::wheelEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QScrollBar::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QScrollBar::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QScrollBar::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QScrollBar::mouseMoveEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QScrollBar::hideEvent(__a0); }
	void cppgoBase_sliderChange(int __a0) { QScrollBar::sliderChange((QAbstractSlider::SliderChange)__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QScrollBar::contextMenuEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSlider* __a0) { QScrollBar::initStyleOption(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QScrollBar::keyPressEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QScrollBar::timerEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QScrollBar::changeEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QScrollBar::mouseDoubleClickEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QScrollBar::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QScrollBar::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QScrollBar::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QScrollBar::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QScrollBar::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QScrollBar::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QScrollBar::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QScrollBar::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QScrollBar::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QScrollBar::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QScrollBar::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QScrollBar::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QScrollBar::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QScrollBar::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QScrollBar::showEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QScrollBar::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QScrollBar::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QScrollBar::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QScrollBar::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QScrollBar::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QScrollBar::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QScrollBar::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QScrollBar::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QScrollBar::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QScrollBar::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QScrollBar::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QScrollBar::sizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QScrollBar::event(event);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::wheelEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::paintEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::mouseReleaseEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::mouseMoveEvent(arg1);
	}
	void hideEvent(QHideEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::hideEvent(arg1);
	}
	void sliderChange(QAbstractSlider::SliderChange change) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4658628703300230988LL, (int)(change), &__h); if (__h) return; }
		QScrollBar::sliderChange(change);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::contextMenuEvent(arg1);
	}
	void initStyleOption(QStyleOptionSlider* option) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QScrollBar::initStyleOption(option);
	}
	void keyPressEvent(QKeyEvent* ev) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)ev, &__h); if (__h) return; }
		QScrollBar::keyPressEvent(ev);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::timerEvent(arg1);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QScrollBar::changeEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QScrollBar::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QScrollBar::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QScrollBar::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QScrollBar::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QScrollBar::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QScrollBar::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::mouseDoubleClickEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::showEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QScrollBar::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QScrollBar::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QScrollBar::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QScrollBar::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QScrollBar::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QScrollBar::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QScrollBar::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QScrollBar::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QScrollBar::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QScrollBar::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QScrollBar::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QScrollBar::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QScrollBar;
static std::mutex __connsMu_QScrollBar;

extern "C" {

void* QScrollBar_base(void* self) { return (void*)static_cast<QAbstractSlider*>((QScrollBar*)self); }
void QScrollBar_chain(void* self, void** out) {
	QScrollBar* __p = (QScrollBar*)self;
	QAbstractSlider* __b0 = static_cast<QAbstractSlider*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QScrollBar__set_override(void* self, int idx) { ((QScrollBar__Go*)self)->__goOn[idx] = 1; }
void* QScrollBar_fromQObject(void* o) { return (void*)qobject_cast<QScrollBar*>((QObject*)o); }
void* QScrollBar_new(long long __goid, void* parent) {
	return new QScrollBar__Go(__goid, (QWidget*)parent);
}
void* QScrollBar_new2(long long __goid) {
	return new QScrollBar__Go(__goid);
}
void* QScrollBar_new3(long long __goid, int arg1) {
	return new QScrollBar__Go(__goid, (Qt::Orientation)(arg1));
}
void* QScrollBar_new4(long long __goid, int arg1, void* parent) {
	return new QScrollBar__Go(__goid, (Qt::Orientation)(arg1), (QWidget*)parent);
}
char* QScrollBar_tr(const char* s, long long* __out_len) {
	auto&& __r = QScrollBar::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QScrollBar_createStandardContextMenu(void* self, void* position) {
	auto&& __r = ((QScrollBar*)self)->createStandardContextMenu(*(QPoint*)position);
	return (void*)(__r);
}
void QScrollBar_setRepeatAction(void* self, int action) {
	((QScrollBar__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action));
}
int QScrollBar_repeatAction(void* self) {
	auto&& __r = ((QScrollBar__Go*)self)->repeatAction();
	return (int)(__r);
}
void QScrollBar_setTabOrder(void* arg1, void* arg2) {
	QScrollBar::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QScrollBar_mouseGrabber() {
	auto&& __r = QScrollBar::mouseGrabber();
	return (void*)(__r);
}
void* QScrollBar_keyboardGrabber() {
	auto&& __r = QScrollBar::keyboardGrabber();
	return (void*)(__r);
}
void* QScrollBar_find(unsigned long long arg1) {
	auto&& __r = QScrollBar::find(arg1);
	return (void*)(__r);
}
void* QScrollBar_createWindowContainer(void* window) {
	auto&& __r = QScrollBar::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QScrollBar_updateMicroFocus(void* self) {
	((QScrollBar__Go*)self)->updateMicroFocus();
}
void QScrollBar_create(void* self) {
	((QScrollBar__Go*)self)->create();
}
void QScrollBar_destroy(void* self) {
	((QScrollBar__Go*)self)->destroy();
}
int QScrollBar_focusNextChild(void* self) {
	auto&& __r = ((QScrollBar__Go*)self)->focusNextChild();
	return __r;
}
int QScrollBar_focusPreviousChild(void* self) {
	auto&& __r = ((QScrollBar__Go*)self)->focusPreviousChild();
	return __r;
}
void* QScrollBar_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QScrollBar::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QScrollBar_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QScrollBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QScrollBar_sender(void* self) {
	auto&& __r = ((QScrollBar__Go*)self)->sender();
	return (void*)(__r);
}
int QScrollBar_senderSignalIndex(void* self) {
	auto&& __r = ((QScrollBar__Go*)self)->senderSignalIndex();
	return __r;
}
int QScrollBar_receivers(void* self, const char* signal) {
	auto&& __r = ((QScrollBar__Go*)self)->receivers(signal);
	return __r;
}
int QScrollBar_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QScrollBar__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QScrollBar_devicePixelRatioFScale() {
	auto&& __r = QScrollBar::devicePixelRatioFScale();
	return __r;
}
int QScrollBar_encodeMetricF(int metric, double value) {
	auto&& __r = QScrollBar::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QScrollBar_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QScrollBar__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QScrollBar_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QScrollBar::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScrollBar_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QScrollBar::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QScrollBar_setRepeatAction2(void* self, int action, int thresholdTime) {
	((QScrollBar__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime);
}
void QScrollBar_setRepeatAction3(void* self, int action, int thresholdTime, int repeatTime) {
	((QScrollBar__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime, repeatTime);
}
void* QScrollBar_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QScrollBar::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QScrollBar_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QScrollBar::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QScrollBar_updateMicroFocus2(void* self, int query) {
	((QScrollBar__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QScrollBar_create2(void* self, unsigned long long arg1) {
	((QScrollBar__Go*)self)->create(arg1);
}
void QScrollBar_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QScrollBar__Go*)self)->create(arg1, initializeWindow);
}
void QScrollBar_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QScrollBar__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QScrollBar_destroy2(void* self, int destroyWindow) {
	((QScrollBar__Go*)self)->destroy(destroyWindow);
}
void QScrollBar_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QScrollBar__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QScrollBar_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QScrollBar::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QScrollBar_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QScrollBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QScrollBar_metaObject_default(void* self) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::metaObject();
	return (void*)(__r);
}
void* QScrollBar_sizeHint_default(void* self) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::sizeHint();
	return new QSize(__r);
}
int QScrollBar_event_default(void* self, void* event) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::event((QEvent*)event);
	return __r;
}
void QScrollBar_wheelEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QScrollBar_paintEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QScrollBar_mousePressEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QScrollBar_mouseReleaseEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QScrollBar_mouseMoveEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QScrollBar_hideEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_hideEvent((QHideEvent*)arg1);
}
void QScrollBar_sliderChange_default(void* self, int change) {
	((QScrollBar__Go*)self)->cppgoBase_sliderChange(change);
}
void QScrollBar_contextMenuEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QScrollBar_initStyleOption_default(void* self, void* option) {
	((QScrollBar__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSlider*)option);
}
void QScrollBar_keyPressEvent_default(void* self, void* ev) {
	((QScrollBar__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)ev);
}
void QScrollBar_timerEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
void QScrollBar_changeEvent_default(void* self, void* e) {
	((QScrollBar__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
int QScrollBar_devType_default(void* self) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::devType();
	return __r;
}
void QScrollBar_setVisible_default(void* self, int visible) {
	((QScrollBar*)self)->QScrollBar::setVisible(visible);
}
void* QScrollBar_minimumSizeHint_default(void* self) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::minimumSizeHint();
	return new QSize(__r);
}
int QScrollBar_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::heightForWidth(arg1);
	return __r;
}
int QScrollBar_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::hasHeightForWidth();
	return __r;
}
void* QScrollBar_paintEngine_default(void* self) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::paintEngine();
	return (void*)(__r);
}
void QScrollBar_mouseDoubleClickEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QScrollBar_keyReleaseEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QScrollBar_focusInEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QScrollBar_focusOutEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QScrollBar_enterEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QScrollBar_leaveEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QScrollBar_moveEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QScrollBar_resizeEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QScrollBar_closeEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QScrollBar_tabletEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QScrollBar_actionEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QScrollBar_dragEnterEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QScrollBar_dragMoveEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QScrollBar_dragLeaveEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QScrollBar_dropEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QScrollBar_showEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
int QScrollBar_metric_default(void* self, int arg1) {
	auto&& __r = ((QScrollBar__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QScrollBar_initPainter_default(void* self, void* painter) {
	((QScrollBar__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QScrollBar_redirected_default(void* self, void* offset) {
	auto&& __r = ((QScrollBar__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QScrollBar_sharedPainter_default(void* self) {
	auto&& __r = ((QScrollBar__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QScrollBar_inputMethodEvent_default(void* self, void* arg1) {
	((QScrollBar__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QScrollBar_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QScrollBar_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QScrollBar__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QScrollBar_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QScrollBar*)self)->QScrollBar::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QScrollBar_childEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QScrollBar_customEvent_default(void* self, void* event) {
	((QScrollBar__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QScrollBar_connectNotify_default(void* self, void* signal) {
	((QScrollBar__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QScrollBar_disconnectNotify_default(void* self, void* signal) {
	((QScrollBar__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QScrollBar_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::valueChanged, (QScrollBar*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_sliderPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::sliderPressed, (QScrollBar*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_sliderMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::sliderMoved, (QScrollBar*)self, [__id](int position){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, position);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_sliderReleased(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::sliderReleased, (QScrollBar*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_rangeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::rangeChanged, (QScrollBar*)self, [__id](int min, int max){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, min, max);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_actionTriggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::actionTriggered, (QScrollBar*)self, [__id](int action){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::windowTitleChanged, (QScrollBar*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::windowIconChanged, (QScrollBar*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::windowIconTextChanged, (QScrollBar*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::customContextMenuRequested, (QScrollBar*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::destroyed, (QScrollBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::destroyed, (QScrollBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollBar*)self, &QScrollBar::objectNameChanged, (QScrollBar*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
	__conns_QScrollBar[__id] = __c;
}
void QScrollBar__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QScrollBar);
		auto __it = __conns_QScrollBar.find(id);
		if (__it == __conns_QScrollBar.end()) return;
		__c = __it->second;
		__conns_QScrollBar.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QScrollBar__Go_new(long long id) { return new QScrollBar__Go(id); }
void QScrollBar_delete(void* self) { delete (QScrollBar*)self; }

}

#include <QtWidgets/qtoolbutton.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qmenu.h>
#include <QtGui/qaction.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qicon.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qbuttongroup.h>
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
#include <QtCore/qlocale.h>
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
#include "gen_qtoolbutton.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QToolButton__Go : public QToolButton {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QToolButton__Go(long long id, A&&... a) : QToolButton(std::forward<A>(a)...), __goId(id) {}
	using QToolButton::updateMicroFocus;
	using QToolButton::create;
	using QToolButton::destroy;
	using QToolButton::focusNextChild;
	using QToolButton::focusPreviousChild;
	using QToolButton::sender;
	using QToolButton::senderSignalIndex;
	using QToolButton::receivers;
	using QToolButton::isSignalConnected;
	using QToolButton::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QToolButton::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QToolButton::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QToolButton::mouseReleaseEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QToolButton::paintEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QToolButton::actionEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QToolButton::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QToolButton::leaveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QToolButton::timerEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QToolButton::changeEvent(__a0); }
	bool cppgoBase_hitButton(const QPoint& __a0) { return QToolButton::hitButton(__a0); }
	void cppgoBase_checkStateSet() { QToolButton::checkStateSet(); }
	void cppgoBase_nextCheckState() { QToolButton::nextCheckState(); }
	void cppgoBase_initStyleOption(QStyleOptionToolButton* __a0) { QToolButton::initStyleOption(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QToolButton::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QToolButton::keyReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QToolButton::mouseMoveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QToolButton::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QToolButton::focusOutEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QToolButton::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QToolButton::wheelEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QToolButton::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QToolButton::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QToolButton::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QToolButton::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QToolButton::tabletEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QToolButton::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QToolButton::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QToolButton::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QToolButton::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QToolButton::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QToolButton::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QToolButton::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QToolButton::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QToolButton::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QToolButton::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QToolButton::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QToolButton::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QToolButton::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QToolButton::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QToolButton::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QToolButton::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QToolButton::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QToolButton::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QToolButton::minimumSizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QToolButton::event(e);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::mouseReleaseEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::paintEvent(arg1);
	}
	void actionEvent(QActionEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::actionEvent(arg1);
	}
	void enterEvent(QEnterEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::enterEvent(arg1);
	}
	void leaveEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::leaveEvent(arg1);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::timerEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::changeEvent(arg1);
	}
	bool hitButton(const QPoint& pos) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5736040353629178930LL, (void*)&pos, &__h); if (__h) return __r; }
		return QToolButton::hitButton(pos);
	}
	void checkStateSet() override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5747806663224476200LL, &__h); if (__h) return; }
		QToolButton::checkStateSet();
	}
	void nextCheckState() override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1474647914901213797LL, &__h); if (__h) return; }
		QToolButton::nextCheckState();
	}
	void initStyleOption(QStyleOptionToolButton* option) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QToolButton::initStyleOption(option);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QToolButton::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QToolButton::keyReleaseEvent(e);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QToolButton::mouseMoveEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QToolButton::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QToolButton::focusOutEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QToolButton::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QToolButton::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QToolButton::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QToolButton::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QToolButton::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QToolButton::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QToolButton::wheelEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QToolButton::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QToolButton::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QToolButton::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QToolButton::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QToolButton::tabletEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QToolButton::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QToolButton::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QToolButton::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QToolButton::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QToolButton::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QToolButton::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QToolButton::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QToolButton::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QToolButton::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QToolButton::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QToolButton::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QToolButton::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QToolButton::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QToolButton::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QToolButton::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QToolButton::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QToolButton::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QToolButton::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QToolButton::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QToolButton;
static std::mutex __connsMu_QToolButton;

extern "C" {

void* QToolButton_base(void* self) { return (void*)static_cast<QAbstractButton*>((QToolButton*)self); }
void QToolButton_chain(void* self, void** out) {
	QToolButton* __p = (QToolButton*)self;
	QAbstractButton* __b0 = static_cast<QAbstractButton*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QToolButton__set_override(void* self, int idx) { ((QToolButton__Go*)self)->__goOn[idx] = 1; }
void* QToolButton_fromQObject(void* o) { return (void*)qobject_cast<QToolButton*>((QObject*)o); }
void* QToolButton_new(long long __goid, void* parent) {
	return new QToolButton__Go(__goid, (QWidget*)parent);
}
void* QToolButton_new2(long long __goid) {
	return new QToolButton__Go(__goid);
}
char* QToolButton_tr(const char* s, long long* __out_len) {
	auto&& __r = QToolButton::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QToolButton_toolButtonStyle(void* self) {
	auto&& __r = ((QToolButton*)self)->toolButtonStyle();
	return (int)(__r);
}
int QToolButton_arrowType(void* self) {
	auto&& __r = ((QToolButton*)self)->arrowType();
	return (int)(__r);
}
void QToolButton_setArrowType(void* self, int type_) {
	((QToolButton*)self)->setArrowType((Qt::ArrowType)(type_));
}
void QToolButton_setMenu(void* self, void* menu) {
	((QToolButton*)self)->setMenu((QMenu*)menu);
}
void* QToolButton_menu(void* self) {
	auto&& __r = ((QToolButton*)self)->menu();
	return (void*)(__r);
}
void QToolButton_setPopupMode(void* self, int mode) {
	((QToolButton*)self)->setPopupMode((QToolButton::ToolButtonPopupMode)(mode));
}
int QToolButton_popupMode(void* self) {
	auto&& __r = ((QToolButton*)self)->popupMode();
	return (int)(__r);
}
void* QToolButton_defaultAction(void* self) {
	auto&& __r = ((QToolButton*)self)->defaultAction();
	return (void*)(__r);
}
void QToolButton_setAutoRaise(void* self, int enable) {
	((QToolButton*)self)->setAutoRaise(enable);
}
int QToolButton_autoRaise(void* self) {
	auto&& __r = ((QToolButton*)self)->autoRaise();
	return __r;
}
void QToolButton_showMenu(void* self) {
	((QToolButton*)self)->showMenu();
}
void QToolButton_setToolButtonStyle(void* self, int style) {
	((QToolButton*)self)->setToolButtonStyle((Qt::ToolButtonStyle)(style));
}
void QToolButton_setDefaultAction(void* self, void* arg1) {
	((QToolButton*)self)->setDefaultAction((QAction*)arg1);
}
void QToolButton_triggered(void* self, void* arg1) {
	((QToolButton*)self)->triggered((QAction*)arg1);
}
void QToolButton_setTabOrder(void* arg1, void* arg2) {
	QToolButton::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QToolButton_mouseGrabber() {
	auto&& __r = QToolButton::mouseGrabber();
	return (void*)(__r);
}
void* QToolButton_keyboardGrabber() {
	auto&& __r = QToolButton::keyboardGrabber();
	return (void*)(__r);
}
void* QToolButton_find(unsigned long long arg1) {
	auto&& __r = QToolButton::find(arg1);
	return (void*)(__r);
}
void* QToolButton_createWindowContainer(void* window) {
	auto&& __r = QToolButton::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QToolButton_updateMicroFocus(void* self) {
	((QToolButton__Go*)self)->updateMicroFocus();
}
void QToolButton_create(void* self) {
	((QToolButton__Go*)self)->create();
}
void QToolButton_destroy(void* self) {
	((QToolButton__Go*)self)->destroy();
}
int QToolButton_focusNextChild(void* self) {
	auto&& __r = ((QToolButton__Go*)self)->focusNextChild();
	return __r;
}
int QToolButton_focusPreviousChild(void* self) {
	auto&& __r = ((QToolButton__Go*)self)->focusPreviousChild();
	return __r;
}
void* QToolButton_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QToolButton::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QToolButton_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QToolButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QToolButton_sender(void* self) {
	auto&& __r = ((QToolButton__Go*)self)->sender();
	return (void*)(__r);
}
int QToolButton_senderSignalIndex(void* self) {
	auto&& __r = ((QToolButton__Go*)self)->senderSignalIndex();
	return __r;
}
int QToolButton_receivers(void* self, const char* signal) {
	auto&& __r = ((QToolButton__Go*)self)->receivers(signal);
	return __r;
}
int QToolButton_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QToolButton__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QToolButton_devicePixelRatioFScale() {
	auto&& __r = QToolButton::devicePixelRatioFScale();
	return __r;
}
int QToolButton_encodeMetricF(int metric, double value) {
	auto&& __r = QToolButton::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QToolButton_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QToolButton__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QToolButton_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QToolButton::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QToolButton_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QToolButton::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QToolButton_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QToolButton::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QToolButton_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QToolButton::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QToolButton_updateMicroFocus2(void* self, int query) {
	((QToolButton__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QToolButton_create2(void* self, unsigned long long arg1) {
	((QToolButton__Go*)self)->create(arg1);
}
void QToolButton_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QToolButton__Go*)self)->create(arg1, initializeWindow);
}
void QToolButton_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QToolButton__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QToolButton_destroy2(void* self, int destroyWindow) {
	((QToolButton__Go*)self)->destroy(destroyWindow);
}
void QToolButton_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QToolButton__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QToolButton_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QToolButton::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QToolButton_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QToolButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QToolButton_metaObject_default(void* self) {
	auto&& __r = ((QToolButton*)self)->QToolButton::metaObject();
	return (void*)(__r);
}
void* QToolButton_sizeHint_default(void* self) {
	auto&& __r = ((QToolButton*)self)->QToolButton::sizeHint();
	return new QSize(__r);
}
void* QToolButton_minimumSizeHint_default(void* self) {
	auto&& __r = ((QToolButton*)self)->QToolButton::minimumSizeHint();
	return new QSize(__r);
}
int QToolButton_event_default(void* self, void* e) {
	auto&& __r = ((QToolButton__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QToolButton_mousePressEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QToolButton_mouseReleaseEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QToolButton_paintEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QToolButton_actionEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_actionEvent((QActionEvent*)arg1);
}
void QToolButton_enterEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)arg1);
}
void QToolButton_leaveEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_leaveEvent((QEvent*)arg1);
}
void QToolButton_timerEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
void QToolButton_changeEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QToolButton_hitButton_default(void* self, void* pos) {
	auto&& __r = ((QToolButton__Go*)self)->cppgoBase_hitButton(*(QPoint*)pos);
	return __r;
}
void QToolButton_checkStateSet_default(void* self) {
	((QToolButton__Go*)self)->cppgoBase_checkStateSet();
}
void QToolButton_nextCheckState_default(void* self) {
	((QToolButton__Go*)self)->cppgoBase_nextCheckState();
}
void QToolButton_initStyleOption_default(void* self, void* option) {
	((QToolButton__Go*)self)->cppgoBase_initStyleOption((QStyleOptionToolButton*)option);
}
void QToolButton_keyPressEvent_default(void* self, void* e) {
	((QToolButton__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QToolButton_keyReleaseEvent_default(void* self, void* e) {
	((QToolButton__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QToolButton_mouseMoveEvent_default(void* self, void* e) {
	((QToolButton__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QToolButton_focusInEvent_default(void* self, void* e) {
	((QToolButton__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QToolButton_focusOutEvent_default(void* self, void* e) {
	((QToolButton__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
int QToolButton_devType_default(void* self) {
	auto&& __r = ((QToolButton*)self)->QToolButton::devType();
	return __r;
}
void QToolButton_setVisible_default(void* self, int visible) {
	((QToolButton*)self)->QToolButton::setVisible(visible);
}
int QToolButton_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QToolButton*)self)->QToolButton::heightForWidth(arg1);
	return __r;
}
int QToolButton_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QToolButton*)self)->QToolButton::hasHeightForWidth();
	return __r;
}
void* QToolButton_paintEngine_default(void* self) {
	auto&& __r = ((QToolButton*)self)->QToolButton::paintEngine();
	return (void*)(__r);
}
void QToolButton_mouseDoubleClickEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QToolButton_wheelEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QToolButton_moveEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QToolButton_resizeEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QToolButton_closeEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QToolButton_contextMenuEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QToolButton_tabletEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QToolButton_dragEnterEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QToolButton_dragMoveEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QToolButton_dragLeaveEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QToolButton_dropEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QToolButton_showEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QToolButton_hideEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QToolButton_metric_default(void* self, int arg1) {
	auto&& __r = ((QToolButton__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QToolButton_initPainter_default(void* self, void* painter) {
	((QToolButton__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QToolButton_redirected_default(void* self, void* offset) {
	auto&& __r = ((QToolButton__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QToolButton_sharedPainter_default(void* self) {
	auto&& __r = ((QToolButton__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QToolButton_inputMethodEvent_default(void* self, void* arg1) {
	((QToolButton__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QToolButton_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QToolButton*)self)->QToolButton::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QToolButton_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QToolButton__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QToolButton_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QToolButton*)self)->QToolButton::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QToolButton_childEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QToolButton_customEvent_default(void* self, void* event) {
	((QToolButton__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QToolButton_connectNotify_default(void* self, void* signal) {
	((QToolButton__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QToolButton_disconnectNotify_default(void* self, void* signal) {
	((QToolButton__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QToolButton_connect_triggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::triggered, (QToolButton*)self, [__id](QAction* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::pressed, (QToolButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_released(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::released, (QToolButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::clicked, (QToolButton*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::clicked, (QToolButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::toggled, (QToolButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::windowTitleChanged, (QToolButton*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::windowIconChanged, (QToolButton*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::windowIconTextChanged, (QToolButton*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::customContextMenuRequested, (QToolButton*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::destroyed, (QToolButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::destroyed, (QToolButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolButton*)self, &QToolButton::objectNameChanged, (QToolButton*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
	__conns_QToolButton[__id] = __c;
}
void QToolButton__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QToolButton);
		auto __it = __conns_QToolButton.find(id);
		if (__it == __conns_QToolButton.end()) return;
		__c = __it->second;
		__conns_QToolButton.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QToolButton__Go_new(long long id) { return new QToolButton__Go(id); }
void QToolButton_delete(void* self) { delete (QToolButton*)self; }

}

