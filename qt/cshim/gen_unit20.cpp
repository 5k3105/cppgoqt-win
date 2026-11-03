#include <QtCore/qchar.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include "gen_qchar.h"

extern "C" {

void* QChar_new() {
	return new QChar();
}
void* QChar_new2(unsigned char c, unsigned char r) {
	return new QChar(c, r);
}
void* QChar_new3(void* arg1) {
	return new QChar(*(QChar*)arg1);
}
void* QChar_fromUcs2(unsigned short c) {
	auto&& __r = QChar::fromUcs2(c);
	return new QChar(__r);
}
int QChar_category(void* self) {
	auto&& __r = ((QChar*)self)->category();
	return (int)(__r);
}
int QChar_direction(void* self) {
	auto&& __r = ((QChar*)self)->direction();
	return (int)(__r);
}
int QChar_joiningType(void* self) {
	auto&& __r = ((QChar*)self)->joiningType();
	return (int)(__r);
}
unsigned char QChar_combiningClass(void* self) {
	auto&& __r = ((QChar*)self)->combiningClass();
	return __r;
}
void* QChar_mirroredChar(void* self) {
	auto&& __r = ((QChar*)self)->mirroredChar();
	return new QChar(__r);
}
int QChar_hasMirrored(void* self) {
	auto&& __r = ((QChar*)self)->hasMirrored();
	return __r;
}
char* QChar_decomposition(void* self, long long* __out_len) {
	auto&& __r = ((QChar*)self)->decomposition();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QChar_decompositionTag(void* self) {
	auto&& __r = ((QChar*)self)->decompositionTag();
	return (int)(__r);
}
int QChar_digitValue(void* self) {
	auto&& __r = ((QChar*)self)->digitValue();
	return __r;
}
void* QChar_toLower(void* self) {
	auto&& __r = ((QChar*)self)->toLower();
	return new QChar(__r);
}
void* QChar_toUpper(void* self) {
	auto&& __r = ((QChar*)self)->toUpper();
	return new QChar(__r);
}
void* QChar_toTitleCase(void* self) {
	auto&& __r = ((QChar*)self)->toTitleCase();
	return new QChar(__r);
}
void* QChar_toCaseFolded(void* self) {
	auto&& __r = ((QChar*)self)->toCaseFolded();
	return new QChar(__r);
}
int QChar_script(void* self) {
	auto&& __r = ((QChar*)self)->script();
	return (int)(__r);
}
int QChar_unicodeVersion(void* self) {
	auto&& __r = ((QChar*)self)->unicodeVersion();
	return (int)(__r);
}
char QChar_toLatin1(void* self) {
	auto&& __r = ((QChar*)self)->toLatin1();
	return __r;
}
unsigned short QChar_unicode(void* self) {
	auto&& __r = ((QChar*)self)->unicode();
	return __r;
}
void* QChar_fromLatin1(char c) {
	auto&& __r = QChar::fromLatin1(c);
	return new QChar(__r);
}
int QChar_isNull(void* self) {
	auto&& __r = ((QChar*)self)->isNull();
	return __r;
}
int QChar_isPrint(void* self) {
	auto&& __r = ((QChar*)self)->isPrint();
	return __r;
}
int QChar_isSpace(void* self) {
	auto&& __r = ((QChar*)self)->isSpace();
	return __r;
}
int QChar_isMark(void* self) {
	auto&& __r = ((QChar*)self)->isMark();
	return __r;
}
int QChar_isPunct(void* self) {
	auto&& __r = ((QChar*)self)->isPunct();
	return __r;
}
int QChar_isSymbol(void* self) {
	auto&& __r = ((QChar*)self)->isSymbol();
	return __r;
}
int QChar_isLetter(void* self) {
	auto&& __r = ((QChar*)self)->isLetter();
	return __r;
}
int QChar_isNumber(void* self) {
	auto&& __r = ((QChar*)self)->isNumber();
	return __r;
}
int QChar_isLetterOrNumber(void* self) {
	auto&& __r = ((QChar*)self)->isLetterOrNumber();
	return __r;
}
int QChar_isDigit(void* self) {
	auto&& __r = ((QChar*)self)->isDigit();
	return __r;
}
int QChar_isLower(void* self) {
	auto&& __r = ((QChar*)self)->isLower();
	return __r;
}
int QChar_isUpper(void* self) {
	auto&& __r = ((QChar*)self)->isUpper();
	return __r;
}
int QChar_isTitleCase(void* self) {
	auto&& __r = ((QChar*)self)->isTitleCase();
	return __r;
}
int QChar_isNonCharacter(void* self) {
	auto&& __r = ((QChar*)self)->isNonCharacter();
	return __r;
}
int QChar_isHighSurrogate(void* self) {
	auto&& __r = ((QChar*)self)->isHighSurrogate();
	return __r;
}
int QChar_isLowSurrogate(void* self) {
	auto&& __r = ((QChar*)self)->isLowSurrogate();
	return __r;
}
int QChar_isSurrogate(void* self) {
	auto&& __r = ((QChar*)self)->isSurrogate();
	return __r;
}
unsigned char QChar_cell(void* self) {
	auto&& __r = ((QChar*)self)->cell();
	return __r;
}
unsigned char QChar_row(void* self) {
	auto&& __r = ((QChar*)self)->row();
	return __r;
}
void QChar_setCell(void* self, unsigned char acell) {
	((QChar*)self)->setCell(acell);
}
void QChar_setRow(void* self, unsigned char arow) {
	((QChar*)self)->setRow(arow);
}
int QChar_isNonCharacter2(unsigned int ucs4) {
	auto&& __r = QChar::isNonCharacter(ucs4);
	return __r;
}
int QChar_isHighSurrogate2(unsigned int ucs4) {
	auto&& __r = QChar::isHighSurrogate(ucs4);
	return __r;
}
int QChar_isLowSurrogate2(unsigned int ucs4) {
	auto&& __r = QChar::isLowSurrogate(ucs4);
	return __r;
}
int QChar_isSurrogate2(unsigned int ucs4) {
	auto&& __r = QChar::isSurrogate(ucs4);
	return __r;
}
int QChar_requiresSurrogates(unsigned int ucs4) {
	auto&& __r = QChar::requiresSurrogates(ucs4);
	return __r;
}
unsigned int QChar_surrogateToUcs4(unsigned short high, unsigned short low) {
	auto&& __r = QChar::surrogateToUcs4(high, low);
	return __r;
}
unsigned int QChar_surrogateToUcs42(void* high, void* low) {
	auto&& __r = QChar::surrogateToUcs4(*(QChar*)high, *(QChar*)low);
	return __r;
}
unsigned short QChar_highSurrogate(unsigned int ucs4) {
	auto&& __r = QChar::highSurrogate(ucs4);
	return __r;
}
unsigned short QChar_lowSurrogate(unsigned int ucs4) {
	auto&& __r = QChar::lowSurrogate(ucs4);
	return __r;
}
int QChar_category2(unsigned int ucs4) {
	auto&& __r = QChar::category(ucs4);
	return (int)(__r);
}
int QChar_direction2(unsigned int ucs4) {
	auto&& __r = QChar::direction(ucs4);
	return (int)(__r);
}
int QChar_joiningType2(unsigned int ucs4) {
	auto&& __r = QChar::joiningType(ucs4);
	return (int)(__r);
}
unsigned char QChar_combiningClass2(unsigned int ucs4) {
	auto&& __r = QChar::combiningClass(ucs4);
	return __r;
}
unsigned int QChar_mirroredChar2(unsigned int ucs4) {
	auto&& __r = QChar::mirroredChar(ucs4);
	return __r;
}
int QChar_hasMirrored2(unsigned int ucs4) {
	auto&& __r = QChar::hasMirrored(ucs4);
	return __r;
}
char* QChar_decomposition2(unsigned int ucs4, long long* __out_len) {
	auto&& __r = QChar::decomposition(ucs4);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QChar_decompositionTag2(unsigned int ucs4) {
	auto&& __r = QChar::decompositionTag(ucs4);
	return (int)(__r);
}
int QChar_digitValue2(unsigned int ucs4) {
	auto&& __r = QChar::digitValue(ucs4);
	return __r;
}
unsigned int QChar_toLower2(unsigned int ucs4) {
	auto&& __r = QChar::toLower(ucs4);
	return __r;
}
unsigned int QChar_toUpper2(unsigned int ucs4) {
	auto&& __r = QChar::toUpper(ucs4);
	return __r;
}
unsigned int QChar_toTitleCase2(unsigned int ucs4) {
	auto&& __r = QChar::toTitleCase(ucs4);
	return __r;
}
unsigned int QChar_toCaseFolded2(unsigned int ucs4) {
	auto&& __r = QChar::toCaseFolded(ucs4);
	return __r;
}
int QChar_script2(unsigned int ucs4) {
	auto&& __r = QChar::script(ucs4);
	return (int)(__r);
}
int QChar_unicodeVersion2(unsigned int ucs4) {
	auto&& __r = QChar::unicodeVersion(ucs4);
	return (int)(__r);
}
int QChar_currentUnicodeVersion() {
	auto&& __r = QChar::currentUnicodeVersion();
	return (int)(__r);
}
int QChar_isPrint2(unsigned int ucs4) {
	auto&& __r = QChar::isPrint(ucs4);
	return __r;
}
int QChar_isSpace2(unsigned int ucs4) {
	auto&& __r = QChar::isSpace(ucs4);
	return __r;
}
int QChar_isMark2(unsigned int ucs4) {
	auto&& __r = QChar::isMark(ucs4);
	return __r;
}
int QChar_isPunct2(unsigned int ucs4) {
	auto&& __r = QChar::isPunct(ucs4);
	return __r;
}
int QChar_isSymbol2(unsigned int ucs4) {
	auto&& __r = QChar::isSymbol(ucs4);
	return __r;
}
int QChar_isLetter2(unsigned int ucs4) {
	auto&& __r = QChar::isLetter(ucs4);
	return __r;
}
int QChar_isNumber2(unsigned int ucs4) {
	auto&& __r = QChar::isNumber(ucs4);
	return __r;
}
int QChar_isLetterOrNumber2(unsigned int ucs4) {
	auto&& __r = QChar::isLetterOrNumber(ucs4);
	return __r;
}
int QChar_isDigit2(unsigned int ucs4) {
	auto&& __r = QChar::isDigit(ucs4);
	return __r;
}
int QChar_isLower2(unsigned int ucs4) {
	auto&& __r = QChar::isLower(ucs4);
	return __r;
}
int QChar_isUpper2(unsigned int ucs4) {
	auto&& __r = QChar::isUpper(ucs4);
	return __r;
}
int QChar_isTitleCase2(unsigned int ucs4) {
	auto&& __r = QChar::isTitleCase(ucs4);
	return __r;
}
int QChar_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) == (*(QChar*)rhs);
	return __r;
}
int QChar_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) != (*(QChar*)rhs);
	return __r;
}
int QChar_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) < (*(QChar*)rhs);
	return __r;
}
int QChar_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) > (*(QChar*)rhs);
	return __r;
}
int QChar_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) <= (*(QChar*)rhs);
	return __r;
}
int QChar_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) >= (*(QChar*)rhs);
	return __r;
}
int QChar_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) == (*(QByteArrayView*)rhs);
	return __r;
}
int QChar_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) != (*(QByteArrayView*)rhs);
	return __r;
}
int QChar_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) < (*(QByteArrayView*)rhs);
	return __r;
}
int QChar_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) > (*(QByteArrayView*)rhs);
	return __r;
}
int QChar_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) <= (*(QByteArrayView*)rhs);
	return __r;
}
int QChar_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) >= (*(QByteArrayView*)rhs);
	return __r;
}
int QChar_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) == (*(QByteArray*)rhs);
	return __r;
}
int QChar_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) != (*(QByteArray*)rhs);
	return __r;
}
int QChar_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) < (*(QByteArray*)rhs);
	return __r;
}
int QChar_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) > (*(QByteArray*)rhs);
	return __r;
}
int QChar_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) <= (*(QByteArray*)rhs);
	return __r;
}
int QChar_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) >= (*(QByteArray*)rhs);
	return __r;
}
int QChar_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) == (*(QStringView*)rhs);
	return __r;
}
int QChar_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) != (*(QStringView*)rhs);
	return __r;
}
int QChar_OperatorLesser4(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) < (*(QStringView*)rhs);
	return __r;
}
int QChar_OperatorGreater4(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) > (*(QStringView*)rhs);
	return __r;
}
int QChar_OperatorLesserOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) <= (*(QStringView*)rhs);
	return __r;
}
int QChar_OperatorGreaterOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QChar*)self) >= (*(QStringView*)rhs);
	return __r;
}
int QChar_OperatorEqual5(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QChar*)self) == (rhs_qstr);
	return __r;
}
int QChar_OperatorNotEqual5(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QChar*)self) != (rhs_qstr);
	return __r;
}
int QChar_OperatorLesser5(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QChar*)self) < (rhs_qstr);
	return __r;
}
int QChar_OperatorGreater5(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QChar*)self) > (rhs_qstr);
	return __r;
}
int QChar_OperatorLesserOrEqual5(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QChar*)self) <= (rhs_qstr);
	return __r;
}
int QChar_OperatorGreaterOrEqual5(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QChar*)self) >= (rhs_qstr);
	return __r;
}
char* QChar_OperatorPlus(void* self, const char* s2, long long s2_len, long long* __out_len) {
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = (*(QChar*)self) + (s2_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QChar_delete(void* self) { delete (QChar*)self; }

}

#include <QtCore/qbindingstorage.h>
#include "gen_qbindingstatus.h"

extern "C" {

void QBindingStatus_delete(void* self) { delete (QBindingStatus*)self; }

}

#include <QtCore/qvariant.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qchar.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qbitarray.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qlocale.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qurl.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/quuid.h>
#include <QtCore/qsize.h>
#include <QtCore/qpoint.h>
#include <QtCore/qline.h>
#include <QtCore/qrect.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qjsondocument.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcompare.h>
#include "gen_qvariant.h"

extern "C" {

void* QVariant_new() {
	return new QVariant();
}
void* QVariant_new2(void* type_) {
	return new QVariant(*(QMetaType*)type_);
}
void* QVariant_new3(void* other) {
	return new QVariant(*(QVariant*)other);
}
void* QVariant_new4(int i) {
	return new QVariant(i);
}
void* QVariant_new5(unsigned int ui) {
	return new QVariant(ui);
}
void* QVariant_new6(long long ll) {
	return new QVariant(ll);
}
void* QVariant_new7(unsigned long long ull) {
	return new QVariant(ull);
}
void* QVariant_new8(int b) {
	return new QVariant(b);
}
void* QVariant_new9(double d) {
	return new QVariant(d);
}
void* QVariant_new10(float f) {
	return new QVariant(f);
}
void* QVariant_new11(void* qchar) {
	return new QVariant(*(QChar*)qchar);
}
void* QVariant_new12(void* date) {
	return new QVariant(*(QDate*)date);
}
void* QVariant_new13(void* time) {
	return new QVariant(*(QTime*)time);
}
void* QVariant_new14(void* bitarray) {
	return new QVariant(*(QBitArray*)bitarray);
}
void* QVariant_new15(void* bytearray) {
	return new QVariant(*(QByteArray*)bytearray);
}
void* QVariant_new16(void* datetime) {
	return new QVariant(*(QDateTime*)datetime);
}
void* QVariant_new17(char** hash_keys, void* hash_vals, long long hash_len) {
	QHash<QString, QVariant> hash_map;
	void** hash_vv = (void**)hash_vals;
	for (long long __i = 0; __i < hash_len; __i++) hash_map.insert(QString::fromUtf8(hash_keys[__i]), *(QVariant*)hash_vv[__i]);
	return new QVariant(hash_map);
}
void* QVariant_new18(void* jsonArray) {
	return new QVariant(*(QJsonArray*)jsonArray);
}
void* QVariant_new19(void* jsonObject) {
	return new QVariant(*(QJsonObject*)jsonObject);
}
void* QVariant_new20(void** list, long long list_len) {
	QList<QVariant> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QVariant*)list[__i]);
	return new QVariant(list_list);
}
void* QVariant_new21(void* locale) {
	return new QVariant(*(QLocale*)locale);
}
void* QVariant_new22(char** map__keys, void* map__vals, long long map__len) {
	QMap<QString, QVariant> map__map;
	void** map__vv = (void**)map__vals;
	for (long long __i = 0; __i < map__len; __i++) map__map.insert(QString::fromUtf8(map__keys[__i]), *(QVariant*)map__vv[__i]);
	return new QVariant(map__map);
}
void* QVariant_new23(void* re) {
	return new QVariant(*(QRegularExpression*)re);
}
void* QVariant_new24(const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	return new QVariant(string__qstr);
}
void* QVariant_new25(char** stringlist, long long stringlist_len) {
	QList<QString> stringlist_list; stringlist_list.reserve(stringlist_len);
	for (long long __i = 0; __i < stringlist_len; __i++) stringlist_list.append(QString::fromUtf8(stringlist[__i]));
	return new QVariant(stringlist_list);
}
void* QVariant_new26(void* url) {
	return new QVariant(*(QUrl*)url);
}
void* QVariant_new27(void* jsonValue) {
	return new QVariant(*(QJsonValue*)jsonValue);
}
void* QVariant_new28(void* modelIndex) {
	return new QVariant(*(QModelIndex*)modelIndex);
}
void* QVariant_new29(void* uuid) {
	return new QVariant(*(QUuid*)uuid);
}
void* QVariant_new30(void* size) {
	return new QVariant(*(QSize*)size);
}
void* QVariant_new31(void* size) {
	return new QVariant(*(QSizeF*)size);
}
void* QVariant_new32(void* pt) {
	return new QVariant(*(QPoint*)pt);
}
void* QVariant_new33(void* pt) {
	return new QVariant(*(QPointF*)pt);
}
void* QVariant_new34(void* line) {
	return new QVariant(*(QLine*)line);
}
void* QVariant_new35(void* line) {
	return new QVariant(*(QLineF*)line);
}
void* QVariant_new36(void* rect) {
	return new QVariant(*(QRect*)rect);
}
void* QVariant_new37(void* rect) {
	return new QVariant(*(QRectF*)rect);
}
void* QVariant_new38(void* easing) {
	return new QVariant(*(QEasingCurve*)easing);
}
void* QVariant_new39(void* jsonDocument) {
	return new QVariant(*(QJsonDocument*)jsonDocument);
}
void* QVariant_new40(void* modelIndex) {
	return new QVariant(*(QPersistentModelIndex*)modelIndex);
}
void* QVariant_new41(const char* str) {
	return new QVariant(str);
}
void* QVariant_new42(void* string_) {
	return new QVariant(*(QLatin1String*)string_);
}
void* QVariant_new43(int type_) {
	return new QVariant((QVariant::Type)(type_));
}
void* QVariant_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QVariant*)self)->operator=(*(QVariant*)other);
	return (void*)&(__r);
}
void QVariant_swap(void* self, void* other) {
	((QVariant*)self)->swap(*(QVariant*)other);
}
int QVariant_userType(void* self) {
	auto&& __r = ((QVariant*)self)->userType();
	return __r;
}
int QVariant_typeId(void* self) {
	auto&& __r = ((QVariant*)self)->typeId();
	return __r;
}
const char* QVariant_typeName(void* self) {
	auto&& __r = ((QVariant*)self)->typeName();
	return __r;
}
void* QVariant_metaType(void* self) {
	auto&& __r = ((QVariant*)self)->metaType();
	return new QMetaType(__r);
}
int QVariant_canConvert(void* self, void* targetType) {
	auto&& __r = ((QVariant*)self)->canConvert(*(QMetaType*)targetType);
	return __r;
}
int QVariant_convert(void* self, void* type_) {
	auto&& __r = ((QVariant*)self)->convert(*(QMetaType*)type_);
	return __r;
}
int QVariant_canView(void* self, void* targetType) {
	auto&& __r = ((QVariant*)self)->canView(*(QMetaType*)targetType);
	return __r;
}
int QVariant_canConvert2(void* self, int targetTypeId) {
	auto&& __r = ((QVariant*)self)->canConvert(targetTypeId);
	return __r;
}
int QVariant_convert2(void* self, int targetTypeId) {
	auto&& __r = ((QVariant*)self)->convert(targetTypeId);
	return __r;
}
int QVariant_isValid(void* self) {
	auto&& __r = ((QVariant*)self)->isValid();
	return __r;
}
int QVariant_isNull(void* self) {
	auto&& __r = ((QVariant*)self)->isNull();
	return __r;
}
void QVariant_clear(void* self) {
	((QVariant*)self)->clear();
}
void QVariant_detach(void* self) {
	((QVariant*)self)->detach();
}
int QVariant_isDetached(void* self) {
	auto&& __r = ((QVariant*)self)->isDetached();
	return __r;
}
int QVariant_toInt(void* self) {
	auto&& __r = ((QVariant*)self)->toInt();
	return __r;
}
unsigned int QVariant_toUInt(void* self) {
	auto&& __r = ((QVariant*)self)->toUInt();
	return __r;
}
long long QVariant_toLongLong(void* self) {
	auto&& __r = ((QVariant*)self)->toLongLong();
	return __r;
}
unsigned long long QVariant_toULongLong(void* self) {
	auto&& __r = ((QVariant*)self)->toULongLong();
	return __r;
}
int QVariant_toBool(void* self) {
	auto&& __r = ((QVariant*)self)->toBool();
	return __r;
}
double QVariant_toDouble(void* self) {
	auto&& __r = ((QVariant*)self)->toDouble();
	return __r;
}
float QVariant_toFloat(void* self) {
	auto&& __r = ((QVariant*)self)->toFloat();
	return __r;
}
double QVariant_toReal(void* self) {
	auto&& __r = ((QVariant*)self)->toReal();
	return __r;
}
void* QVariant_toByteArray(void* self) {
	auto&& __r = ((QVariant*)self)->toByteArray();
	return new QByteArray(__r);
}
void* QVariant_toBitArray(void* self) {
	auto&& __r = ((QVariant*)self)->toBitArray();
	return new QBitArray(__r);
}
char* QVariant_toString(void* self, long long* __out_len) {
	auto&& __r = ((QVariant*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QVariant_toStringList(void* self, long long* __out_len) {
	auto&& __r = ((QVariant*)self)->toStringList();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QVariant_toChar(void* self) {
	auto&& __r = ((QVariant*)self)->toChar();
	return new QChar(__r);
}
void* QVariant_toDate(void* self) {
	auto&& __r = ((QVariant*)self)->toDate();
	return new QDate(__r);
}
void* QVariant_toTime(void* self) {
	auto&& __r = ((QVariant*)self)->toTime();
	return new QTime(__r);
}
void* QVariant_toDateTime(void* self) {
	auto&& __r = ((QVariant*)self)->toDateTime();
	return new QDateTime(__r);
}
void** QVariant_toList(void* self, long long* __out_len) {
	auto&& __r = ((QVariant*)self)->toList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QVariant(__e);
	return __arr;
}
void* QVariant_toPoint(void* self) {
	auto&& __r = ((QVariant*)self)->toPoint();
	return new QPoint(__r);
}
void* QVariant_toPointF(void* self) {
	auto&& __r = ((QVariant*)self)->toPointF();
	return new QPointF(__r);
}
void* QVariant_toRect(void* self) {
	auto&& __r = ((QVariant*)self)->toRect();
	return new QRect(__r);
}
void* QVariant_toSize(void* self) {
	auto&& __r = ((QVariant*)self)->toSize();
	return new QSize(__r);
}
void* QVariant_toSizeF(void* self) {
	auto&& __r = ((QVariant*)self)->toSizeF();
	return new QSizeF(__r);
}
void* QVariant_toLine(void* self) {
	auto&& __r = ((QVariant*)self)->toLine();
	return new QLine(__r);
}
void* QVariant_toLineF(void* self) {
	auto&& __r = ((QVariant*)self)->toLineF();
	return new QLineF(__r);
}
void* QVariant_toRectF(void* self) {
	auto&& __r = ((QVariant*)self)->toRectF();
	return new QRectF(__r);
}
void* QVariant_toLocale(void* self) {
	auto&& __r = ((QVariant*)self)->toLocale();
	return new QLocale(__r);
}
void* QVariant_toRegularExpression(void* self) {
	auto&& __r = ((QVariant*)self)->toRegularExpression();
	return new QRegularExpression(__r);
}
void* QVariant_toEasingCurve(void* self) {
	auto&& __r = ((QVariant*)self)->toEasingCurve();
	return new QEasingCurve(__r);
}
void* QVariant_toUuid(void* self) {
	auto&& __r = ((QVariant*)self)->toUuid();
	return new QUuid(__r);
}
void* QVariant_toUrl(void* self) {
	auto&& __r = ((QVariant*)self)->toUrl();
	return new QUrl(__r);
}
void* QVariant_toJsonValue(void* self) {
	auto&& __r = ((QVariant*)self)->toJsonValue();
	return new QJsonValue(__r);
}
void* QVariant_toJsonObject(void* self) {
	auto&& __r = ((QVariant*)self)->toJsonObject();
	return new QJsonObject(__r);
}
void* QVariant_toJsonArray(void* self) {
	auto&& __r = ((QVariant*)self)->toJsonArray();
	return new QJsonArray(__r);
}
void* QVariant_toJsonDocument(void* self) {
	auto&& __r = ((QVariant*)self)->toJsonDocument();
	return new QJsonDocument(__r);
}
void* QVariant_toModelIndex(void* self) {
	auto&& __r = ((QVariant*)self)->toModelIndex();
	return new QModelIndex(__r);
}
void* QVariant_toPersistentModelIndex(void* self) {
	auto&& __r = ((QVariant*)self)->toPersistentModelIndex();
	return new QPersistentModelIndex(__r);
}
void QVariant_load(void* self, void* ds) {
	((QVariant*)self)->load(*(QDataStream*)ds);
}
void QVariant_save(void* self, void* ds) {
	((QVariant*)self)->save(*(QDataStream*)ds);
}
int QVariant_type(void* self) {
	auto&& __r = ((QVariant*)self)->type();
	return (int)(__r);
}
const char* QVariant_typeToName(int typeId) {
	auto&& __r = QVariant::typeToName(typeId);
	return __r;
}
int QVariant_nameToType(const char* name) {
	auto&& __r = QVariant::nameToType(name);
	return (int)(__r);
}
void QVariant_setValue(void* self, void* avalue) {
	((QVariant*)self)->setValue(*(QVariant*)avalue);
}
void* QVariant_fromMetaType(void* type_) {
	auto&& __r = QVariant::fromMetaType(*(QMetaType*)type_);
	return new QVariant(__r);
}
void* QVariant_compare(void* lhs, void* rhs) {
	auto&& __r = QVariant::compare(*(QVariant*)lhs, *(QVariant*)rhs);
	return new QPartialOrdering(__r);
}
int QVariant_toInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toInt(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QVariant_toUInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toUInt(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QVariant_toLongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toLongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QVariant_toULongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toULongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
double QVariant_toDouble2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toDouble(&ok_local);
	*ok = ok_local;
	return __r;
}
float QVariant_toFloat2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toFloat(&ok_local);
	*ok = ok_local;
	return __r;
}
double QVariant_toReal2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QVariant*)self)->toReal(&ok_local);
	*ok = ok_local;
	return __r;
}
int QVariant_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QVariant*)self) == (*(QVariant*)rhs);
	return __r;
}
int QVariant_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QVariant*)self) != (*(QVariant*)rhs);
	return __r;
}
void QVariant_delete(void* self) { delete (QVariant*)self; }

}

#include <QtCore/qbitarray.h>
#include <QtCore/qsysinfo.h>
#include "gen_qbitarray.h"

extern "C" {

void* QBitArray_new() {
	return new QBitArray();
}
void* QBitArray_new2(long long size) {
	return new QBitArray(size);
}
void* QBitArray_new3(void* other) {
	return new QBitArray(*(QBitArray*)other);
}
void* QBitArray_new4(long long size, int val) {
	return new QBitArray(size, val);
}
void* QBitArray_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QBitArray*)self)->operator=(*(QBitArray*)other);
	return (void*)&(__r);
}
void QBitArray_swap(void* self, void* other) {
	((QBitArray*)self)->swap(*(QBitArray*)other);
}
long long QBitArray_size(void* self) {
	auto&& __r = ((QBitArray*)self)->size();
	return __r;
}
long long QBitArray_count(void* self) {
	auto&& __r = ((QBitArray*)self)->count();
	return __r;
}
long long QBitArray_count2(void* self, int on) {
	auto&& __r = ((QBitArray*)self)->count(on);
	return __r;
}
int QBitArray_isEmpty(void* self) {
	auto&& __r = ((QBitArray*)self)->isEmpty();
	return __r;
}
int QBitArray_isNull(void* self) {
	auto&& __r = ((QBitArray*)self)->isNull();
	return __r;
}
void QBitArray_resize(void* self, long long size) {
	((QBitArray*)self)->resize(size);
}
void QBitArray_detach(void* self) {
	((QBitArray*)self)->detach();
}
int QBitArray_isDetached(void* self) {
	auto&& __r = ((QBitArray*)self)->isDetached();
	return __r;
}
void QBitArray_clear(void* self) {
	((QBitArray*)self)->clear();
}
int QBitArray_testBit(void* self, long long i) {
	auto&& __r = ((QBitArray*)self)->testBit(i);
	return __r;
}
void QBitArray_setBit(void* self, long long i) {
	((QBitArray*)self)->setBit(i);
}
void QBitArray_setBit2(void* self, long long i, int val) {
	((QBitArray*)self)->setBit(i, val);
}
void QBitArray_clearBit(void* self, long long i) {
	((QBitArray*)self)->clearBit(i);
}
int QBitArray_toggleBit(void* self, long long i) {
	auto&& __r = ((QBitArray*)self)->toggleBit(i);
	return __r;
}
int QBitArray_at(void* self, long long i) {
	auto&& __r = ((QBitArray*)self)->at(i);
	return __r;
}
void* QBitArray_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QBitArray*)self)->operator[](i);
	return new QBitRef(__r);
}
int QBitArray_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QBitArray*)self)->operator[](i);
	return __r;
}
void* QBitArray_OperatorBitwiseAndAssign(void* self, void* arg1) {
	auto&& __r = ((QBitArray*)self)->operator&=(*(QBitArray*)arg1);
	return (void*)&(__r);
}
void* QBitArray_OperatorBitwiseOrAssign(void* self, void* arg1) {
	auto&& __r = ((QBitArray*)self)->operator|=(*(QBitArray*)arg1);
	return (void*)&(__r);
}
void* QBitArray_OperatorBitwiseXorAssign(void* self, void* arg1) {
	auto&& __r = ((QBitArray*)self)->operator^=(*(QBitArray*)arg1);
	return (void*)&(__r);
}
int QBitArray_fill(void* self, int aval) {
	auto&& __r = ((QBitArray*)self)->fill(aval);
	return __r;
}
void QBitArray_fill2(void* self, int val, long long first, long long last) {
	((QBitArray*)self)->fill(val, first, last);
}
void QBitArray_truncate(void* self, long long pos) {
	((QBitArray*)self)->truncate(pos);
}
const char* QBitArray_bits(void* self) {
	auto&& __r = ((QBitArray*)self)->bits();
	return __r;
}
void* QBitArray_fromBits(const char* data, long long len_) {
	auto&& __r = QBitArray::fromBits(data, len_);
	return new QBitArray(__r);
}
unsigned int QBitArray_toUInt32(void* self, int endianness) {
	auto&& __r = ((QBitArray*)self)->toUInt32((QSysInfo::Endian)(endianness));
	return __r;
}
int QBitArray_fill3(void* self, int aval, long long asize) {
	auto&& __r = ((QBitArray*)self)->fill(aval, asize);
	return __r;
}
unsigned int QBitArray_toUInt322(void* self, int endianness, int* ok) {
	bool ok_local{};
	auto&& __r = ((QBitArray*)self)->toUInt32((QSysInfo::Endian)(endianness), &ok_local);
	*ok = ok_local;
	return __r;
}
void* QBitArray_OperatorBitwiseAnd(void* self, void* a2) {
	auto&& __r = (*(QBitArray*)self) & (*(QBitArray*)a2);
	return new QBitArray(__r);
}
void* QBitArray_OperatorBitwiseOr(void* self, void* a2) {
	auto&& __r = (*(QBitArray*)self) | (*(QBitArray*)a2);
	return new QBitArray(__r);
}
void* QBitArray_OperatorBitwiseXor(void* self, void* a2) {
	auto&& __r = (*(QBitArray*)self) ^ (*(QBitArray*)a2);
	return new QBitArray(__r);
}
void* QBitArray_OperatorBitwiseNot(void* self) {
	auto&& __r = ~(*(QBitArray*)self);
	return new QBitArray(__r);
}
int QBitArray_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QBitArray*)self) == (*(QBitArray*)rhs);
	return __r;
}
int QBitArray_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QBitArray*)self) != (*(QBitArray*)rhs);
	return __r;
}
void QBitArray_delete(void* self) { delete (QBitArray*)self; }

}

#include <QtCore/qcborarray.h>
#include <QtCore/qcborvalue.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qjsonarray.h>
#include "gen_qcborarray.h"

extern "C" {

void* QCborArray_new() {
	return new QCborArray();
}
void* QCborArray_new2(void* other) {
	return new QCborArray(*(QCborArray*)other);
}
void* QCborArray_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCborArray*)self)->operator=(*(QCborArray*)other);
	return (void*)&(__r);
}
void QCborArray_swap(void* self, void* other) {
	((QCborArray*)self)->swap(*(QCborArray*)other);
}
void* QCborArray_toCborValue(void* self) {
	auto&& __r = ((QCborArray*)self)->toCborValue();
	return new QCborValue(__r);
}
long long QCborArray_size(void* self) {
	auto&& __r = ((QCborArray*)self)->size();
	return __r;
}
int QCborArray_isEmpty(void* self) {
	auto&& __r = ((QCborArray*)self)->isEmpty();
	return __r;
}
void QCborArray_clear(void* self) {
	((QCborArray*)self)->clear();
}
void* QCborArray_at(void* self, long long i) {
	auto&& __r = ((QCborArray*)self)->at(i);
	return new QCborValue(__r);
}
void* QCborArray_first(void* self) {
	auto&& __r = ((QCborArray*)self)->first();
	return new QCborValue(__r);
}
void* QCborArray_last(void* self) {
	auto&& __r = ((QCborArray*)self)->last();
	return new QCborValue(__r);
}
void* QCborArray_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QCborArray*)self)->operator[](i);
	return new QCborValue(__r);
}
void* QCborArray_first2(void* self) {
	auto&& __r = ((QCborArray*)self)->first();
	return new QCborValueRef(__r);
}
void* QCborArray_last2(void* self) {
	auto&& __r = ((QCborArray*)self)->last();
	return new QCborValueRef(__r);
}
void* QCborArray_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QCborArray*)self)->operator[](i);
	return new QCborValueRef(__r);
}
void QCborArray_insert(void* self, long long i, void* value) {
	((QCborArray*)self)->insert(i, *(QCborValue*)value);
}
void QCborArray_prepend(void* self, void* value) {
	((QCborArray*)self)->prepend(*(QCborValue*)value);
}
void QCborArray_append(void* self, void* value) {
	((QCborArray*)self)->append(*(QCborValue*)value);
}
void QCborArray_removeAt(void* self, long long i) {
	((QCborArray*)self)->removeAt(i);
}
void* QCborArray_takeAt(void* self, long long i) {
	auto&& __r = ((QCborArray*)self)->takeAt(i);
	return new QCborValue(__r);
}
void QCborArray_removeFirst(void* self) {
	((QCborArray*)self)->removeFirst();
}
void QCborArray_removeLast(void* self) {
	((QCborArray*)self)->removeLast();
}
void* QCborArray_takeFirst(void* self) {
	auto&& __r = ((QCborArray*)self)->takeFirst();
	return new QCborValue(__r);
}
void* QCborArray_takeLast(void* self) {
	auto&& __r = ((QCborArray*)self)->takeLast();
	return new QCborValue(__r);
}
int QCborArray_contains(void* self, void* value) {
	auto&& __r = ((QCborArray*)self)->contains(*(QCborValue*)value);
	return __r;
}
int QCborArray_compare(void* self, void* other) {
	auto&& __r = ((QCborArray*)self)->compare(*(QCborArray*)other);
	return __r;
}
void QCborArray_push_back(void* self, void* t) {
	((QCborArray*)self)->push_back(*(QCborValue*)t);
}
void QCborArray_push_front(void* self, void* t) {
	((QCborArray*)self)->push_front(*(QCborValue*)t);
}
void QCborArray_pop_front(void* self) {
	((QCborArray*)self)->pop_front();
}
void QCborArray_pop_back(void* self) {
	((QCborArray*)self)->pop_back();
}
int QCborArray_empty(void* self) {
	auto&& __r = ((QCborArray*)self)->empty();
	return __r;
}
void* QCborArray_OperatorPlus(void* self, void* v) {
	auto&& __r = ((QCborArray*)self)->operator+(*(QCborValue*)v);
	return new QCborArray(__r);
}
void* QCborArray_OperatorPlusAssign(void* self, void* v) {
	auto&& __r = ((QCborArray*)self)->operator+=(*(QCborValue*)v);
	return (void*)&(__r);
}
void* QCborArray_OperatorShiftLeft(void* self, void* v) {
	auto&& __r = ((QCborArray*)self)->operator<<(*(QCborValue*)v);
	return (void*)&(__r);
}
void* QCborArray_fromStringList(char** list, long long list_len) {
	QList<QString> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(QString::fromUtf8(list[__i]));
	auto&& __r = QCborArray::fromStringList(list_list);
	return new QCborArray(__r);
}
void* QCborArray_fromVariantList(void** list, long long list_len) {
	QList<QVariant> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QVariant*)list[__i]);
	auto&& __r = QCborArray::fromVariantList(list_list);
	return new QCborArray(__r);
}
void* QCborArray_fromJsonArray(void* array) {
	auto&& __r = QCborArray::fromJsonArray(*(QJsonArray*)array);
	return new QCborArray(__r);
}
void** QCborArray_toVariantList(void* self, long long* __out_len) {
	auto&& __r = ((QCborArray*)self)->toVariantList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QVariant(__e);
	return __arr;
}
void* QCborArray_toJsonArray(void* self) {
	auto&& __r = ((QCborArray*)self)->toJsonArray();
	return new QJsonArray(__r);
}
int QCborArray_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) == (*(QCborArray*)rhs);
	return __r;
}
int QCborArray_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) != (*(QCborArray*)rhs);
	return __r;
}
int QCborArray_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) < (*(QCborArray*)rhs);
	return __r;
}
int QCborArray_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) > (*(QCborArray*)rhs);
	return __r;
}
int QCborArray_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) <= (*(QCborArray*)rhs);
	return __r;
}
int QCborArray_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) >= (*(QCborArray*)rhs);
	return __r;
}
int QCborArray_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) == (*(QCborValue*)rhs);
	return __r;
}
int QCborArray_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) != (*(QCborValue*)rhs);
	return __r;
}
int QCborArray_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) < (*(QCborValue*)rhs);
	return __r;
}
int QCborArray_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) > (*(QCborValue*)rhs);
	return __r;
}
int QCborArray_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) <= (*(QCborValue*)rhs);
	return __r;
}
int QCborArray_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) >= (*(QCborValue*)rhs);
	return __r;
}
int QCborArray_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) == (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborArray_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) != (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborArray_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) < (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborArray_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) > (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborArray_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) <= (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborArray_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborArray*)self) >= (*(QCborValueConstRef*)rhs);
	return __r;
}
void QCborArray_delete(void* self) { delete (QCborArray*)self; }

}

#include <QtCore/qdirlisting.h>
#include <QtCore/qstring.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qtimezone.h>
#include <QtCore/qfiledevice.h>
#include "gen_qdirlisting__direntry.h"

extern "C" {

void* QDirListing__DirEntry_new(void* arg1) {
	return new QDirListing::DirEntry(*(QDirListing::DirEntry*)arg1);
}
void* QDirListing__DirEntry_new2() {
	return new QDirListing::DirEntry();
}
char* QDirListing__DirEntry_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_baseName(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->baseName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_completeBaseName(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->completeBaseName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_suffix(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->suffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_bundleName(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->bundleName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_completeSuffix(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->completeSuffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_filePath(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->filePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDirListing__DirEntry_isDir(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isDir();
	return __r;
}
int QDirListing__DirEntry_isFile(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isFile();
	return __r;
}
int QDirListing__DirEntry_isSymLink(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isSymLink();
	return __r;
}
int QDirListing__DirEntry_exists(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->exists();
	return __r;
}
int QDirListing__DirEntry_isHidden(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isHidden();
	return __r;
}
int QDirListing__DirEntry_isReadable(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isReadable();
	return __r;
}
int QDirListing__DirEntry_isWritable(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isWritable();
	return __r;
}
int QDirListing__DirEntry_isExecutable(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->isExecutable();
	return __r;
}
void* QDirListing__DirEntry_fileInfo(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->fileInfo();
	return new QFileInfo(__r);
}
char* QDirListing__DirEntry_canonicalFilePath(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->canonicalFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_absoluteFilePath(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->absoluteFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirListing__DirEntry_absolutePath(void* self, long long* __out_len) {
	auto&& __r = ((QDirListing::DirEntry*)self)->absolutePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QDirListing__DirEntry_size(void* self) {
	auto&& __r = ((QDirListing::DirEntry*)self)->size();
	return __r;
}
void* QDirListing__DirEntry_birthTime(void* self, void* tz) {
	auto&& __r = ((QDirListing::DirEntry*)self)->birthTime(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QDirListing__DirEntry_metadataChangeTime(void* self, void* tz) {
	auto&& __r = ((QDirListing::DirEntry*)self)->metadataChangeTime(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QDirListing__DirEntry_lastModified(void* self, void* tz) {
	auto&& __r = ((QDirListing::DirEntry*)self)->lastModified(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QDirListing__DirEntry_lastRead(void* self, void* tz) {
	auto&& __r = ((QDirListing::DirEntry*)self)->lastRead(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QDirListing__DirEntry_fileTime(void* self, int type_, void* tz) {
	auto&& __r = ((QDirListing::DirEntry*)self)->fileTime((QFileDevice::FileTime)(type_), *(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QDirListing__DirEntry_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QDirListing::DirEntry*)self)->operator=(*(QDirListing::DirEntry*)arg1);
	return (void*)&(__r);
}
void QDirListing__DirEntry_delete(void* self) { delete (QDirListing::DirEntry*)self; }

}

#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qabstractitemmodel.h>
#include "gen_qitemselection.h"

extern "C" {

void* QItemSelection_new(void* topLeft, void* bottomRight) {
	return new QItemSelection(*(QModelIndex*)topLeft, *(QModelIndex*)bottomRight);
}
void* QItemSelection_new2() {
	return new QItemSelection();
}
void* QItemSelection_new3(void* arg1) {
	return new QItemSelection(*(QItemSelection*)arg1);
}
void QItemSelection_select(void* self, void* topLeft, void* bottomRight) {
	((QItemSelection*)self)->select(*(QModelIndex*)topLeft, *(QModelIndex*)bottomRight);
}
int QItemSelection_contains(void* self, void* index) {
	auto&& __r = ((QItemSelection*)self)->contains(*(QModelIndex*)index);
	return __r;
}
void** QItemSelection_indexes(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->indexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QItemSelection_merge(void* self, void* other, int command) {
	((QItemSelection*)self)->merge(*(QItemSelection*)other, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QItemSelection_split(void* range_, void* other, void* result) {
	QItemSelection::split(*(QItemSelectionRange*)range_, *(QItemSelectionRange*)other, (QItemSelection*)result);
}
long long QItemSelection_maxSize() {
	auto&& __r = QItemSelection::maxSize();
	return __r;
}
long long QItemSelection_size(void* self) {
	auto&& __r = ((QItemSelection*)self)->size();
	return __r;
}
long long QItemSelection_count(void* self) {
	auto&& __r = ((QItemSelection*)self)->count();
	return __r;
}
long long QItemSelection_length(void* self) {
	auto&& __r = ((QItemSelection*)self)->length();
	return __r;
}
int QItemSelection_isEmpty(void* self) {
	auto&& __r = ((QItemSelection*)self)->isEmpty();
	return __r;
}
void QItemSelection_resize(void* self, long long size) {
	((QItemSelection*)self)->resize(size);
}
void QItemSelection_resize2(void* self, long long size, void* c) {
	((QItemSelection*)self)->resize(size, *(QItemSelectionRange*)c);
}
void QItemSelection_resizeForOverwrite(void* self, long long size) {
	((QItemSelection*)self)->resizeForOverwrite(size);
}
long long QItemSelection_capacity(void* self) {
	auto&& __r = ((QItemSelection*)self)->capacity();
	return __r;
}
void QItemSelection_reserve(void* self, long long size) {
	((QItemSelection*)self)->reserve(size);
}
void QItemSelection_squeeze(void* self) {
	((QItemSelection*)self)->squeeze();
}
void QItemSelection_detach(void* self) {
	((QItemSelection*)self)->detach();
}
int QItemSelection_isDetached(void* self) {
	auto&& __r = ((QItemSelection*)self)->isDetached();
	return __r;
}
int QItemSelection_isSharedWith(void* self, void** other, long long other_len) {
	QList<QItemSelectionRange> other_list; other_list.reserve(other_len);
	for (long long __i = 0; __i < other_len; __i++) other_list.append(*(QItemSelectionRange*)other[__i]);
	auto&& __r = ((QItemSelection*)self)->isSharedWith(other_list);
	return __r;
}
void* QItemSelection_data(void* self) {
	auto&& __r = ((QItemSelection*)self)->data();
	return (void*)(__r);
}
void* QItemSelection_data2(void* self) {
	auto&& __r = ((QItemSelection*)self)->data();
	return (void*)(__r);
}
void* QItemSelection_constData(void* self) {
	auto&& __r = ((QItemSelection*)self)->constData();
	return (void*)(__r);
}
void QItemSelection_clear(void* self) {
	((QItemSelection*)self)->clear();
}
void* QItemSelection_at(void* self, long long i) {
	auto&& __r = ((QItemSelection*)self)->at(i);
	return (void*)&(__r);
}
void* QItemSelection_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QItemSelection*)self)->operator[](i);
	return (void*)&(__r);
}
void* QItemSelection_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QItemSelection*)self)->operator[](i);
	return (void*)&(__r);
}
void QItemSelection_append(void* self, void* t) {
	((QItemSelection*)self)->append(*(QItemSelectionRange*)t);
}
void QItemSelection_append2(void* self, void** l, long long l_len) {
	QList<QItemSelectionRange> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QItemSelectionRange*)l[__i]);
	((QItemSelection*)self)->append(l_list);
}
void QItemSelection_prepend(void* self, void* t) {
	((QItemSelection*)self)->prepend(*(QItemSelectionRange*)t);
}
void QItemSelection_replace(void* self, long long i, void* t) {
	((QItemSelection*)self)->replace(i, *(QItemSelectionRange*)t);
}
void QItemSelection_remove(void* self, long long i) {
	((QItemSelection*)self)->remove(i);
}
void QItemSelection_removeFirst(void* self) {
	((QItemSelection*)self)->removeFirst();
}
void QItemSelection_removeLast(void* self) {
	((QItemSelection*)self)->removeLast();
}
void* QItemSelection_takeFirst(void* self) {
	auto&& __r = ((QItemSelection*)self)->takeFirst();
	return new QItemSelectionRange(__r);
}
void* QItemSelection_takeLast(void* self) {
	auto&& __r = ((QItemSelection*)self)->takeLast();
	return new QItemSelectionRange(__r);
}
void QItemSelection_removeAt(void* self, long long i) {
	((QItemSelection*)self)->removeAt(i);
}
void* QItemSelection_takeAt(void* self, long long i) {
	auto&& __r = ((QItemSelection*)self)->takeAt(i);
	return new QItemSelectionRange(__r);
}
void QItemSelection_move(void* self, long long from, long long to) {
	((QItemSelection*)self)->move(from, to);
}
void* QItemSelection_first(void* self) {
	auto&& __r = ((QItemSelection*)self)->first();
	return (void*)&(__r);
}
void* QItemSelection_first2(void* self) {
	auto&& __r = ((QItemSelection*)self)->first();
	return (void*)&(__r);
}
void* QItemSelection_constFirst(void* self) {
	auto&& __r = ((QItemSelection*)self)->constFirst();
	return (void*)&(__r);
}
void* QItemSelection_last(void* self) {
	auto&& __r = ((QItemSelection*)self)->last();
	return (void*)&(__r);
}
void* QItemSelection_last2(void* self) {
	auto&& __r = ((QItemSelection*)self)->last();
	return (void*)&(__r);
}
void* QItemSelection_constLast(void* self) {
	auto&& __r = ((QItemSelection*)self)->constLast();
	return (void*)&(__r);
}
int QItemSelection_startsWith(void* self, void* t) {
	auto&& __r = ((QItemSelection*)self)->startsWith(*(QItemSelectionRange*)t);
	return __r;
}
int QItemSelection_endsWith(void* self, void* t) {
	auto&& __r = ((QItemSelection*)self)->endsWith(*(QItemSelectionRange*)t);
	return __r;
}
void** QItemSelection_mid(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->mid(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_first3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->first(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_last3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->last(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_sliced(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->sliced(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_sliced2(void* self, long long pos, long long n, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->sliced(pos, n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void* QItemSelection_value(void* self, long long i) {
	auto&& __r = ((QItemSelection*)self)->value(i);
	return new QItemSelectionRange(__r);
}
void* QItemSelection_value2(void* self, long long i, void* defaultValue) {
	auto&& __r = ((QItemSelection*)self)->value(i, *(QItemSelectionRange*)defaultValue);
	return new QItemSelectionRange(__r);
}
void QItemSelection_swapItemsAt(void* self, long long i, long long j) {
	((QItemSelection*)self)->swapItemsAt(i, j);
}
void QItemSelection_push_back(void* self, void* t) {
	((QItemSelection*)self)->push_back(*(QItemSelectionRange*)t);
}
void QItemSelection_push_front(void* self, void* t) {
	((QItemSelection*)self)->push_front(*(QItemSelectionRange*)t);
}
void QItemSelection_pop_back(void* self) {
	((QItemSelection*)self)->pop_back();
}
void QItemSelection_pop_front(void* self) {
	((QItemSelection*)self)->pop_front();
}
int QItemSelection_empty(void* self) {
	auto&& __r = ((QItemSelection*)self)->empty();
	return __r;
}
void* QItemSelection_front(void* self) {
	auto&& __r = ((QItemSelection*)self)->front();
	return (void*)&(__r);
}
void* QItemSelection_front2(void* self) {
	auto&& __r = ((QItemSelection*)self)->front();
	return (void*)&(__r);
}
void* QItemSelection_back(void* self) {
	auto&& __r = ((QItemSelection*)self)->back();
	return (void*)&(__r);
}
void* QItemSelection_back2(void* self) {
	auto&& __r = ((QItemSelection*)self)->back();
	return (void*)&(__r);
}
void QItemSelection_shrink_to_fit(void* self) {
	((QItemSelection*)self)->shrink_to_fit();
}
long long QItemSelection_max_size(void* self) {
	auto&& __r = ((QItemSelection*)self)->max_size();
	return __r;
}
void** QItemSelection_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QItemSelectionRange> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QItemSelectionRange*)l[__i]);
	auto&& __r = ((QItemSelection*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QItemSelectionRange> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QItemSelectionRange*)l[__i]);
	auto&& __r = ((QItemSelection*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_fromList(void** list, long long list_len, long long* __out_len) {
	QList<QItemSelectionRange> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QItemSelectionRange*)list[__i]);
	auto&& __r = QItemSelection::fromList(list_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_toList(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->toList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_fromVector(void** vector, long long vector_len, long long* __out_len) {
	QList<QItemSelectionRange> vector_list; vector_list.reserve(vector_len);
	for (long long __i = 0; __i < vector_len; __i++) vector_list.append(*(QItemSelectionRange*)vector[__i]);
	auto&& __r = QItemSelection::fromVector(vector_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void** QItemSelection_toVector(void* self, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->toVector();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void QItemSelection_remove2(void* self, long long i, long long n) {
	((QItemSelection*)self)->remove(i, n);
}
void** QItemSelection_mid2(void* self, long long pos, long long len_, long long* __out_len) {
	auto&& __r = ((QItemSelection*)self)->mid(pos, len_);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QItemSelectionRange(__e);
	return __arr;
}
void QItemSelection_delete(void* self) { delete (QItemSelection*)self; }

}

#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qobject.h>
#include <QtCore/qproperty.h>
#include <QtCore/qobjectdefs.h>
#include "gen_qmetaproperty.h"

extern "C" {

void* QMetaProperty_new() {
	return new QMetaProperty();
}
const char* QMetaProperty_name(void* self) {
	auto&& __r = ((QMetaProperty*)self)->name();
	return __r;
}
const char* QMetaProperty_typeName(void* self) {
	auto&& __r = ((QMetaProperty*)self)->typeName();
	return __r;
}
int QMetaProperty_type(void* self) {
	auto&& __r = ((QMetaProperty*)self)->type();
	return (int)(__r);
}
int QMetaProperty_userType(void* self) {
	auto&& __r = ((QMetaProperty*)self)->userType();
	return __r;
}
int QMetaProperty_typeId(void* self) {
	auto&& __r = ((QMetaProperty*)self)->typeId();
	return __r;
}
void* QMetaProperty_metaType(void* self) {
	auto&& __r = ((QMetaProperty*)self)->metaType();
	return new QMetaType(__r);
}
int QMetaProperty_propertyIndex(void* self) {
	auto&& __r = ((QMetaProperty*)self)->propertyIndex();
	return __r;
}
int QMetaProperty_relativePropertyIndex(void* self) {
	auto&& __r = ((QMetaProperty*)self)->relativePropertyIndex();
	return __r;
}
int QMetaProperty_isReadable(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isReadable();
	return __r;
}
int QMetaProperty_isWritable(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isWritable();
	return __r;
}
int QMetaProperty_isResettable(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isResettable();
	return __r;
}
int QMetaProperty_isDesignable(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isDesignable();
	return __r;
}
int QMetaProperty_isScriptable(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isScriptable();
	return __r;
}
int QMetaProperty_isStored(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isStored();
	return __r;
}
int QMetaProperty_isUser(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isUser();
	return __r;
}
int QMetaProperty_isConstant(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isConstant();
	return __r;
}
int QMetaProperty_isFinal(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isFinal();
	return __r;
}
int QMetaProperty_isVirtual(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isVirtual();
	return __r;
}
int QMetaProperty_isOverride(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isOverride();
	return __r;
}
int QMetaProperty_isRequired(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isRequired();
	return __r;
}
int QMetaProperty_isBindable(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isBindable();
	return __r;
}
int QMetaProperty_isFlagType(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isFlagType();
	return __r;
}
int QMetaProperty_isEnumType(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isEnumType();
	return __r;
}
void* QMetaProperty_enumerator(void* self) {
	auto&& __r = ((QMetaProperty*)self)->enumerator();
	return new QMetaEnum(__r);
}
int QMetaProperty_hasNotifySignal(void* self) {
	auto&& __r = ((QMetaProperty*)self)->hasNotifySignal();
	return __r;
}
void* QMetaProperty_notifySignal(void* self) {
	auto&& __r = ((QMetaProperty*)self)->notifySignal();
	return new QMetaMethod(__r);
}
int QMetaProperty_notifySignalIndex(void* self) {
	auto&& __r = ((QMetaProperty*)self)->notifySignalIndex();
	return __r;
}
int QMetaProperty_revision(void* self) {
	auto&& __r = ((QMetaProperty*)self)->revision();
	return __r;
}
void* QMetaProperty_read(void* self, void* obj) {
	auto&& __r = ((QMetaProperty*)self)->read((QObject*)obj);
	return new QVariant(__r);
}
int QMetaProperty_write(void* self, void* obj, void* value) {
	auto&& __r = ((QMetaProperty*)self)->write((QObject*)obj, *(QVariant*)value);
	return __r;
}
int QMetaProperty_reset(void* self, void* obj) {
	auto&& __r = ((QMetaProperty*)self)->reset((QObject*)obj);
	return __r;
}
void* QMetaProperty_bindable(void* self, void* object) {
	auto&& __r = ((QMetaProperty*)self)->bindable((QObject*)object);
	return new QUntypedBindable(__r);
}
int QMetaProperty_hasStdCppSet(void* self) {
	auto&& __r = ((QMetaProperty*)self)->hasStdCppSet();
	return __r;
}
int QMetaProperty_isAlias(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isAlias();
	return __r;
}
int QMetaProperty_isValid(void* self) {
	auto&& __r = ((QMetaProperty*)self)->isValid();
	return __r;
}
void* QMetaProperty_enclosingMetaObject(void* self) {
	auto&& __r = ((QMetaProperty*)self)->enclosingMetaObject();
	return (void*)(__r);
}
void QMetaProperty_delete(void* self) { delete (QMetaProperty*)self; }

}

#include <QtCore/qprocess.h>
#include <QtCore/qstring.h>
#include "gen_qprocessenvironment.h"

extern "C" {

void* QProcessEnvironment_new() {
	return new QProcessEnvironment();
}
void* QProcessEnvironment_new2(int arg1) {
	return new QProcessEnvironment((QProcessEnvironment::Initialization)(arg1));
}
void* QProcessEnvironment_new3(void* other) {
	return new QProcessEnvironment(*(QProcessEnvironment*)other);
}
void* QProcessEnvironment_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QProcessEnvironment*)self)->operator=(*(QProcessEnvironment*)other);
	return (void*)&(__r);
}
void QProcessEnvironment_swap(void* self, void* other) {
	((QProcessEnvironment*)self)->swap(*(QProcessEnvironment*)other);
}
int QProcessEnvironment_isEmpty(void* self) {
	auto&& __r = ((QProcessEnvironment*)self)->isEmpty();
	return __r;
}
int QProcessEnvironment_inheritsFromParent(void* self) {
	auto&& __r = ((QProcessEnvironment*)self)->inheritsFromParent();
	return __r;
}
void QProcessEnvironment_clear(void* self) {
	((QProcessEnvironment*)self)->clear();
}
int QProcessEnvironment_contains(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QProcessEnvironment*)self)->contains(name_qstr);
	return __r;
}
void QProcessEnvironment_insert(void* self, const char* name, long long name_len, const char* value, long long value_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QProcessEnvironment*)self)->insert(name_qstr, value_qstr);
}
void QProcessEnvironment_remove(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QProcessEnvironment*)self)->remove(name_qstr);
}
char* QProcessEnvironment_value(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QProcessEnvironment*)self)->value(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QProcessEnvironment_toStringList(void* self, long long* __out_len) {
	auto&& __r = ((QProcessEnvironment*)self)->toStringList();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QProcessEnvironment_keys(void* self, long long* __out_len) {
	auto&& __r = ((QProcessEnvironment*)self)->keys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QProcessEnvironment_insert2(void* self, void* e) {
	((QProcessEnvironment*)self)->insert(*(QProcessEnvironment*)e);
}
void* QProcessEnvironment_systemEnvironment() {
	auto&& __r = QProcessEnvironment::systemEnvironment();
	return new QProcessEnvironment(__r);
}
char* QProcessEnvironment_value2(void* self, const char* name, long long name_len, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QProcessEnvironment*)self)->value(name_qstr, defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QProcessEnvironment_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QProcessEnvironment*)self) == (*(QProcessEnvironment*)rhs);
	return __r;
}
int QProcessEnvironment_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QProcessEnvironment*)self) != (*(QProcessEnvironment*)rhs);
	return __r;
}
void QProcessEnvironment_delete(void* self) { delete (QProcessEnvironment*)self; }

}

#include <QtCore/qsettings.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsettings.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSettings__Go : public QSettings {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QSettings__Go(long long id, A&&... a) : QSettings(std::forward<A>(a)...), __goId(id) {}
	using QSettings::sender;
	using QSettings::senderSignalIndex;
	using QSettings::receivers;
	using QSettings::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QSettings::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSettings::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSettings::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSettings::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSettings::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSettings::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSettings::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSettings::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSettings::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSettings::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSettings::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSettings::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSettings::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSettings::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSettings;
static std::mutex __connsMu_QSettings;

extern "C" {

void* QSettings_base(void* self) { return (void*)static_cast<QObject*>((QSettings*)self); }
void QSettings__set_override(void* self, int idx) { ((QSettings__Go*)self)->__goOn[idx] = 1; }
void* QSettings_fromQObject(void* o) { return (void*)qobject_cast<QSettings*>((QObject*)o); }
void* QSettings_new(long long __goid, const char* organization, long long organization_len) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	return new QSettings__Go(__goid, organization_qstr);
}
void* QSettings_new2(long long __goid, int scope, const char* organization, long long organization_len) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	return new QSettings__Go(__goid, (QSettings::Scope)(scope), organization_qstr);
}
void* QSettings_new3(long long __goid, int format, int scope, const char* organization, long long organization_len) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	return new QSettings__Go(__goid, (QSettings::Format)(format), (QSettings::Scope)(scope), organization_qstr);
}
void* QSettings_new4(long long __goid, const char* fileName, long long fileName_len, int format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QSettings__Go(__goid, fileName_qstr, (QSettings::Format)(format));
}
void* QSettings_new5(long long __goid) {
	return new QSettings__Go(__goid);
}
void* QSettings_new6(long long __goid, int scope) {
	return new QSettings__Go(__goid, (QSettings::Scope)(scope));
}
void* QSettings_new7(long long __goid, const char* organization, long long organization_len, const char* application, long long application_len) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	QString application_qstr = QString::fromUtf8(application, application_len);
	return new QSettings__Go(__goid, organization_qstr, application_qstr);
}
void* QSettings_new8(long long __goid, const char* organization, long long organization_len, const char* application, long long application_len, void* parent) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	QString application_qstr = QString::fromUtf8(application, application_len);
	return new QSettings__Go(__goid, organization_qstr, application_qstr, (QObject*)parent);
}
void* QSettings_new9(long long __goid, int scope, const char* organization, long long organization_len, const char* application, long long application_len) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	QString application_qstr = QString::fromUtf8(application, application_len);
	return new QSettings__Go(__goid, (QSettings::Scope)(scope), organization_qstr, application_qstr);
}
void* QSettings_new10(long long __goid, int scope, const char* organization, long long organization_len, const char* application, long long application_len, void* parent) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	QString application_qstr = QString::fromUtf8(application, application_len);
	return new QSettings__Go(__goid, (QSettings::Scope)(scope), organization_qstr, application_qstr, (QObject*)parent);
}
void* QSettings_new11(long long __goid, int format, int scope, const char* organization, long long organization_len, const char* application, long long application_len) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	QString application_qstr = QString::fromUtf8(application, application_len);
	return new QSettings__Go(__goid, (QSettings::Format)(format), (QSettings::Scope)(scope), organization_qstr, application_qstr);
}
void* QSettings_new12(long long __goid, int format, int scope, const char* organization, long long organization_len, const char* application, long long application_len, void* parent) {
	QString organization_qstr = QString::fromUtf8(organization, organization_len);
	QString application_qstr = QString::fromUtf8(application, application_len);
	return new QSettings__Go(__goid, (QSettings::Format)(format), (QSettings::Scope)(scope), organization_qstr, application_qstr, (QObject*)parent);
}
void* QSettings_new13(long long __goid, const char* fileName, long long fileName_len, int format, void* parent) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QSettings__Go(__goid, fileName_qstr, (QSettings::Format)(format), (QObject*)parent);
}
void* QSettings_new14(long long __goid, void* parent) {
	return new QSettings__Go(__goid, (QObject*)parent);
}
void* QSettings_new15(long long __goid, int scope, void* parent) {
	return new QSettings__Go(__goid, (QSettings::Scope)(scope), (QObject*)parent);
}
char* QSettings_tr(const char* s, long long* __out_len) {
	auto&& __r = QSettings::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSettings_clear(void* self) {
	((QSettings*)self)->clear();
}
void QSettings_sync(void* self) {
	((QSettings*)self)->sync();
}
int QSettings_status(void* self) {
	auto&& __r = ((QSettings*)self)->status();
	return (int)(__r);
}
int QSettings_isAtomicSyncRequired(void* self) {
	auto&& __r = ((QSettings*)self)->isAtomicSyncRequired();
	return __r;
}
void QSettings_setAtomicSyncRequired(void* self, int enable) {
	((QSettings*)self)->setAtomicSyncRequired(enable);
}
void QSettings_beginGroup(void* self, const char* prefix, long long prefix_len) {
	QAnyStringView prefix_qstr = QAnyStringView(prefix, prefix_len);
	((QSettings*)self)->beginGroup(prefix_qstr);
}
void QSettings_endGroup(void* self) {
	((QSettings*)self)->endGroup();
}
char* QSettings_group(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->group();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSettings_beginReadArray(void* self, const char* prefix, long long prefix_len) {
	QAnyStringView prefix_qstr = QAnyStringView(prefix, prefix_len);
	auto&& __r = ((QSettings*)self)->beginReadArray(prefix_qstr);
	return __r;
}
void QSettings_beginWriteArray(void* self, const char* prefix, long long prefix_len) {
	QAnyStringView prefix_qstr = QAnyStringView(prefix, prefix_len);
	((QSettings*)self)->beginWriteArray(prefix_qstr);
}
void QSettings_endArray(void* self) {
	((QSettings*)self)->endArray();
}
void QSettings_setArrayIndex(void* self, int i) {
	((QSettings*)self)->setArrayIndex(i);
}
char** QSettings_allKeys(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->allKeys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QSettings_childKeys(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->childKeys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QSettings_childGroups(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->childGroups();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QSettings_isWritable(void* self) {
	auto&& __r = ((QSettings*)self)->isWritable();
	return __r;
}
void QSettings_setValue(void* self, const char* key, long long key_len, void* value) {
	QAnyStringView key_qstr = QAnyStringView(key, key_len);
	((QSettings*)self)->setValue(key_qstr, *(QVariant*)value);
}
void* QSettings_value(void* self, const char* key, long long key_len, void* defaultValue) {
	QAnyStringView key_qstr = QAnyStringView(key, key_len);
	auto&& __r = ((QSettings*)self)->value(key_qstr, *(QVariant*)defaultValue);
	return new QVariant(__r);
}
void* QSettings_value2(void* self, const char* key, long long key_len) {
	QAnyStringView key_qstr = QAnyStringView(key, key_len);
	auto&& __r = ((QSettings*)self)->value(key_qstr);
	return new QVariant(__r);
}
void QSettings_remove(void* self, const char* key, long long key_len) {
	QAnyStringView key_qstr = QAnyStringView(key, key_len);
	((QSettings*)self)->remove(key_qstr);
}
int QSettings_contains(void* self, const char* key, long long key_len) {
	QAnyStringView key_qstr = QAnyStringView(key, key_len);
	auto&& __r = ((QSettings*)self)->contains(key_qstr);
	return __r;
}
void QSettings_setFallbacksEnabled(void* self, int b) {
	((QSettings*)self)->setFallbacksEnabled(b);
}
int QSettings_fallbacksEnabled(void* self) {
	auto&& __r = ((QSettings*)self)->fallbacksEnabled();
	return __r;
}
char* QSettings_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSettings_format(void* self) {
	auto&& __r = ((QSettings*)self)->format();
	return (int)(__r);
}
int QSettings_scope(void* self) {
	auto&& __r = ((QSettings*)self)->scope();
	return (int)(__r);
}
char* QSettings_organizationName(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->organizationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSettings_applicationName(void* self, long long* __out_len) {
	auto&& __r = ((QSettings*)self)->applicationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSettings_setDefaultFormat(int format) {
	QSettings::setDefaultFormat((QSettings::Format)(format));
}
int QSettings_defaultFormat() {
	auto&& __r = QSettings::defaultFormat();
	return (int)(__r);
}
void QSettings_setPath(int format, int scope, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QSettings::setPath((QSettings::Format)(format), (QSettings::Scope)(scope), path_qstr);
}
void* QSettings_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSettings::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSettings_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSettings::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSettings_sender(void* self) {
	auto&& __r = ((QSettings__Go*)self)->sender();
	return (void*)(__r);
}
int QSettings_senderSignalIndex(void* self) {
	auto&& __r = ((QSettings__Go*)self)->senderSignalIndex();
	return __r;
}
int QSettings_receivers(void* self, const char* signal) {
	auto&& __r = ((QSettings__Go*)self)->receivers(signal);
	return __r;
}
int QSettings_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSettings__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSettings_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSettings::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSettings_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSettings::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSettings_beginWriteArray2(void* self, const char* prefix, long long prefix_len, int size) {
	QAnyStringView prefix_qstr = QAnyStringView(prefix, prefix_len);
	((QSettings*)self)->beginWriteArray(prefix_qstr, size);
}
void* QSettings_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSettings::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSettings_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSettings::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSettings_metaObject_default(void* self) {
	auto&& __r = ((QSettings*)self)->QSettings::metaObject();
	return (void*)(__r);
}
int QSettings_event_default(void* self, void* event) {
	auto&& __r = ((QSettings__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QSettings_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSettings*)self)->QSettings::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSettings_timerEvent_default(void* self, void* event) {
	((QSettings__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSettings_childEvent_default(void* self, void* event) {
	((QSettings__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSettings_customEvent_default(void* self, void* event) {
	((QSettings__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSettings_connectNotify_default(void* self, void* signal) {
	((QSettings__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSettings_disconnectNotify_default(void* self, void* signal) {
	((QSettings__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSettings_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSettings*)self, &QSettings::destroyed, (QSettings*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSettings);
	__conns_QSettings[__id] = __c;
}
void QSettings_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSettings*)self, &QSettings::destroyed, (QSettings*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSettings);
	__conns_QSettings[__id] = __c;
}
void QSettings_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSettings*)self, &QSettings::objectNameChanged, (QSettings*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSettings);
	__conns_QSettings[__id] = __c;
}
void QSettings__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSettings);
		auto __it = __conns_QSettings.find(id);
		if (__it == __conns_QSettings.end()) return;
		__c = __it->second;
		__conns_QSettings.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSettings__Go_new(long long id, void* parent) { return new QSettings__Go(id, (QObject*)parent); }
void QSettings_delete(void* self) { delete (QSettings*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qstringview.h>
#include "gen_qxmlstreamnotationdeclaration.h"

extern "C" {

void* QXmlStreamNotationDeclaration_new() {
	return new QXmlStreamNotationDeclaration();
}
void* QXmlStreamNotationDeclaration_new2(void* arg1) {
	return new QXmlStreamNotationDeclaration(*(QXmlStreamNotationDeclaration*)arg1);
}
void* QXmlStreamNotationDeclaration_name(void* self) {
	auto&& __r = ((QXmlStreamNotationDeclaration*)self)->name();
	return new QStringView(__r);
}
void* QXmlStreamNotationDeclaration_systemId(void* self) {
	auto&& __r = ((QXmlStreamNotationDeclaration*)self)->systemId();
	return new QStringView(__r);
}
void* QXmlStreamNotationDeclaration_publicId(void* self) {
	auto&& __r = ((QXmlStreamNotationDeclaration*)self)->publicId();
	return new QStringView(__r);
}
int QXmlStreamNotationDeclaration_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamNotationDeclaration*)self) == (*(QXmlStreamNotationDeclaration*)rhs);
	return __r;
}
int QXmlStreamNotationDeclaration_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamNotationDeclaration*)self) != (*(QXmlStreamNotationDeclaration*)rhs);
	return __r;
}
void QXmlStreamNotationDeclaration_delete(void* self) { delete (QXmlStreamNotationDeclaration*)self; }

}

#include <QtGui/qpointingdevice.h>
#include "gen_qpointingdeviceuniqueid.h"

extern "C" {

void* QPointingDeviceUniqueId_new() {
	return new QPointingDeviceUniqueId();
}
void* QPointingDeviceUniqueId_new2(void* arg1) {
	return new QPointingDeviceUniqueId(*(QPointingDeviceUniqueId*)arg1);
}
void* QPointingDeviceUniqueId_fromNumericId(long long id) {
	auto&& __r = QPointingDeviceUniqueId::fromNumericId(id);
	return new QPointingDeviceUniqueId(__r);
}
int QPointingDeviceUniqueId_isValid(void* self) {
	auto&& __r = ((QPointingDeviceUniqueId*)self)->isValid();
	return __r;
}
long long QPointingDeviceUniqueId_numericId(void* self) {
	auto&& __r = ((QPointingDeviceUniqueId*)self)->numericId();
	return __r;
}
int QPointingDeviceUniqueId_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPointingDeviceUniqueId*)self) == (*(QPointingDeviceUniqueId*)rhs);
	return __r;
}
int QPointingDeviceUniqueId_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPointingDeviceUniqueId*)self) != (*(QPointingDeviceUniqueId*)rhs);
	return __r;
}
void QPointingDeviceUniqueId_delete(void* self) { delete (QPointingDeviceUniqueId*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qcoreevent.h>
#include "gen_qcontextmenuevent.h"

extern "C" {

void* QContextMenuEvent_base(void* self) { return (void*)static_cast<QInputEvent*>((QContextMenuEvent*)self); }
void QContextMenuEvent_chain(void* self, void** out) {
	QContextMenuEvent* __p = (QContextMenuEvent*)self;
	QInputEvent* __b0 = static_cast<QInputEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void* QContextMenuEvent_new(int reason, void* pos, void* globalPos) {
	return new QContextMenuEvent((QContextMenuEvent::Reason)(reason), *(QPoint*)pos, *(QPoint*)globalPos);
}
void* QContextMenuEvent_new2(int reason, void* pos) {
	return new QContextMenuEvent((QContextMenuEvent::Reason)(reason), *(QPoint*)pos);
}
void* QContextMenuEvent_new3(int reason, void* pos, void* globalPos, int modifiers) {
	return new QContextMenuEvent((QContextMenuEvent::Reason)(reason), *(QPoint*)pos, *(QPoint*)globalPos, QFlags<Qt::KeyboardModifier>(modifiers));
}
int QContextMenuEvent_x(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->x();
	return __r;
}
int QContextMenuEvent_y(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->y();
	return __r;
}
int QContextMenuEvent_globalX(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->globalX();
	return __r;
}
int QContextMenuEvent_globalY(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->globalY();
	return __r;
}
void* QContextMenuEvent_pos(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->pos();
	return (void*)&(__r);
}
void* QContextMenuEvent_globalPos(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->globalPos();
	return (void*)&(__r);
}
int QContextMenuEvent_reason(void* self) {
	auto&& __r = ((QContextMenuEvent*)self)->reason();
	return (int)(__r);
}
int QContextMenuEvent_registerEventType() {
	auto&& __r = QContextMenuEvent::registerEventType();
	return __r;
}
int QContextMenuEvent_registerEventType2(int hint) {
	auto&& __r = QContextMenuEvent::registerEventType(hint);
	return __r;
}
void QContextMenuEvent_delete(void* self) { delete (QContextMenuEvent*)self; }

}

#include <QtGui/qfontdatabase.h>
#include <QtCore/qstring.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontinfo.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qchar.h>
#include "gen_qfontdatabase.h"

extern "C" {

void* QFontDatabase_new() {
	return new QFontDatabase();
}
int* QFontDatabase_standardSizes(long long* __out_len) {
	auto&& __r = QFontDatabase::standardSizes();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int* QFontDatabase_writingSystems(long long* __out_len) {
	auto&& __r = QFontDatabase::writingSystems();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
int* QFontDatabase_writingSystems2(const char* family, long long family_len, long long* __out_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::writingSystems(family_qstr);
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
char** QFontDatabase_families(long long* __out_len) {
	auto&& __r = QFontDatabase::families();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFontDatabase_styles(const char* family, long long family_len, long long* __out_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::styles(family_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int* QFontDatabase_pointSizes(const char* family, long long family_len, long long* __out_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::pointSizes(family_qstr);
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int* QFontDatabase_smoothSizes(const char* family, long long family_len, const char* style, long long style_len, long long* __out_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::smoothSizes(family_qstr, style_qstr);
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
char* QFontDatabase_styleString(void* font, long long* __out_len) {
	auto&& __r = QFontDatabase::styleString(*(QFont*)font);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFontDatabase_styleString2(void* fontInfo, long long* __out_len) {
	auto&& __r = QFontDatabase::styleString(*(QFontInfo*)fontInfo);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFontDatabase_font(const char* family, long long family_len, const char* style, long long style_len, int pointSize) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::font(family_qstr, style_qstr, pointSize);
	return new QFont(__r);
}
int QFontDatabase_isBitmapScalable(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::isBitmapScalable(family_qstr);
	return __r;
}
int QFontDatabase_isSmoothlyScalable(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::isSmoothlyScalable(family_qstr);
	return __r;
}
int QFontDatabase_isScalable(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::isScalable(family_qstr);
	return __r;
}
int QFontDatabase_isFixedPitch(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::isFixedPitch(family_qstr);
	return __r;
}
int QFontDatabase_italic(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::italic(family_qstr, style_qstr);
	return __r;
}
int QFontDatabase_bold(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::bold(family_qstr, style_qstr);
	return __r;
}
int QFontDatabase_weight(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::weight(family_qstr, style_qstr);
	return __r;
}
int QFontDatabase_hasFamily(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::hasFamily(family_qstr);
	return __r;
}
int QFontDatabase_isPrivateFamily(const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	auto&& __r = QFontDatabase::isPrivateFamily(family_qstr);
	return __r;
}
char* QFontDatabase_writingSystemName(int writingSystem, long long* __out_len) {
	auto&& __r = QFontDatabase::writingSystemName((QFontDatabase::WritingSystem)(writingSystem));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFontDatabase_writingSystemSample(int writingSystem, long long* __out_len) {
	auto&& __r = QFontDatabase::writingSystemSample((QFontDatabase::WritingSystem)(writingSystem));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFontDatabase_addApplicationFont(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFontDatabase::addApplicationFont(fileName_qstr);
	return __r;
}
int QFontDatabase_addApplicationFontFromData(void* fontData) {
	auto&& __r = QFontDatabase::addApplicationFontFromData(*(QByteArray*)fontData);
	return __r;
}
char** QFontDatabase_applicationFontFamilies(int id, long long* __out_len) {
	auto&& __r = QFontDatabase::applicationFontFamilies(id);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QFontDatabase_removeApplicationFont(int id) {
	auto&& __r = QFontDatabase::removeApplicationFont(id);
	return __r;
}
int QFontDatabase_removeAllApplicationFonts() {
	auto&& __r = QFontDatabase::removeAllApplicationFonts();
	return __r;
}
void QFontDatabase_addApplicationFallbackFontFamily(int script, const char* familyName, long long familyName_len) {
	QString familyName_qstr = QString::fromUtf8(familyName, familyName_len);
	QFontDatabase::addApplicationFallbackFontFamily((QChar::Script)(script), familyName_qstr);
}
int QFontDatabase_removeApplicationFallbackFontFamily(int script, const char* familyName, long long familyName_len) {
	QString familyName_qstr = QString::fromUtf8(familyName, familyName_len);
	auto&& __r = QFontDatabase::removeApplicationFallbackFontFamily((QChar::Script)(script), familyName_qstr);
	return __r;
}
void QFontDatabase_setApplicationFallbackFontFamilies(int arg1, char** familyNames, long long familyNames_len) {
	QList<QString> familyNames_list; familyNames_list.reserve(familyNames_len);
	for (long long __i = 0; __i < familyNames_len; __i++) familyNames_list.append(QString::fromUtf8(familyNames[__i]));
	QFontDatabase::setApplicationFallbackFontFamilies((QChar::Script)(arg1), familyNames_list);
}
char** QFontDatabase_applicationFallbackFontFamilies(int script, long long* __out_len) {
	auto&& __r = QFontDatabase::applicationFallbackFontFamilies((QChar::Script)(script));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFontDatabase_addApplicationEmojiFontFamily(const char* familyName, long long familyName_len) {
	QString familyName_qstr = QString::fromUtf8(familyName, familyName_len);
	QFontDatabase::addApplicationEmojiFontFamily(familyName_qstr);
}
int QFontDatabase_removeApplicationEmojiFontFamily(const char* familyName, long long familyName_len) {
	QString familyName_qstr = QString::fromUtf8(familyName, familyName_len);
	auto&& __r = QFontDatabase::removeApplicationEmojiFontFamily(familyName_qstr);
	return __r;
}
void QFontDatabase_setApplicationEmojiFontFamilies(char** familyNames, long long familyNames_len) {
	QList<QString> familyNames_list; familyNames_list.reserve(familyNames_len);
	for (long long __i = 0; __i < familyNames_len; __i++) familyNames_list.append(QString::fromUtf8(familyNames[__i]));
	QFontDatabase::setApplicationEmojiFontFamilies(familyNames_list);
}
char** QFontDatabase_applicationEmojiFontFamilies(long long* __out_len) {
	auto&& __r = QFontDatabase::applicationEmojiFontFamilies();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QFontDatabase_systemFont(int type_) {
	auto&& __r = QFontDatabase::systemFont((QFontDatabase::SystemFont)(type_));
	return new QFont(__r);
}
char** QFontDatabase_families2(int writingSystem, long long* __out_len) {
	auto&& __r = QFontDatabase::families((QFontDatabase::WritingSystem)(writingSystem));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int* QFontDatabase_pointSizes2(const char* family, long long family_len, const char* style, long long style_len, long long* __out_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::pointSizes(family_qstr, style_qstr);
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int QFontDatabase_isBitmapScalable2(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::isBitmapScalable(family_qstr, style_qstr);
	return __r;
}
int QFontDatabase_isSmoothlyScalable2(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::isSmoothlyScalable(family_qstr, style_qstr);
	return __r;
}
int QFontDatabase_isScalable2(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::isScalable(family_qstr, style_qstr);
	return __r;
}
int QFontDatabase_isFixedPitch2(const char* family, long long family_len, const char* style, long long style_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	QString style_qstr = QString::fromUtf8(style, style_len);
	auto&& __r = QFontDatabase::isFixedPitch(family_qstr, style_qstr);
	return __r;
}
void QFontDatabase_delete(void* self) { delete (QFontDatabase*)self; }

}

#include <QtGui/qtextcursor.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextobject.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextdocumentfragment.h>
#include <QtGui/qtextlist.h>
#include <QtGui/qtexttable.h>
#include <QtGui/qimage.h>
#include "gen_qtextcursor.h"

extern "C" {

void* QTextCursor_new() {
	return new QTextCursor();
}
void* QTextCursor_new2(void* document) {
	return new QTextCursor((QTextDocument*)document);
}
void* QTextCursor_new3(void* frame) {
	return new QTextCursor((QTextFrame*)frame);
}
void* QTextCursor_new4(void* block) {
	return new QTextCursor(*(QTextBlock*)block);
}
void* QTextCursor_new5(void* cursor) {
	return new QTextCursor(*(QTextCursor*)cursor);
}
void* QTextCursor_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QTextCursor*)self)->operator=(*(QTextCursor*)other);
	return (void*)&(__r);
}
void QTextCursor_swap(void* self, void* other) {
	((QTextCursor*)self)->swap(*(QTextCursor*)other);
}
int QTextCursor_isNull(void* self) {
	auto&& __r = ((QTextCursor*)self)->isNull();
	return __r;
}
void QTextCursor_setPosition(void* self, int pos) {
	((QTextCursor*)self)->setPosition(pos);
}
int QTextCursor_position(void* self) {
	auto&& __r = ((QTextCursor*)self)->position();
	return __r;
}
int QTextCursor_positionInBlock(void* self) {
	auto&& __r = ((QTextCursor*)self)->positionInBlock();
	return __r;
}
int QTextCursor_anchor(void* self) {
	auto&& __r = ((QTextCursor*)self)->anchor();
	return __r;
}
void QTextCursor_insertText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextCursor*)self)->insertText(text_qstr);
}
void QTextCursor_insertText2(void* self, const char* text, long long text_len, void* format) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextCursor*)self)->insertText(text_qstr, *(QTextCharFormat*)format);
}
int QTextCursor_movePosition(void* self, int op) {
	auto&& __r = ((QTextCursor*)self)->movePosition((QTextCursor::MoveOperation)(op));
	return __r;
}
int QTextCursor_visualNavigation(void* self) {
	auto&& __r = ((QTextCursor*)self)->visualNavigation();
	return __r;
}
void QTextCursor_setVisualNavigation(void* self, int b) {
	((QTextCursor*)self)->setVisualNavigation(b);
}
void QTextCursor_setVerticalMovementX(void* self, int x) {
	((QTextCursor*)self)->setVerticalMovementX(x);
}
int QTextCursor_verticalMovementX(void* self) {
	auto&& __r = ((QTextCursor*)self)->verticalMovementX();
	return __r;
}
void QTextCursor_setKeepPositionOnInsert(void* self, int b) {
	((QTextCursor*)self)->setKeepPositionOnInsert(b);
}
int QTextCursor_keepPositionOnInsert(void* self) {
	auto&& __r = ((QTextCursor*)self)->keepPositionOnInsert();
	return __r;
}
void QTextCursor_deleteChar(void* self) {
	((QTextCursor*)self)->deleteChar();
}
void QTextCursor_deletePreviousChar(void* self) {
	((QTextCursor*)self)->deletePreviousChar();
}
void QTextCursor_select(void* self, int selection) {
	((QTextCursor*)self)->select((QTextCursor::SelectionType)(selection));
}
int QTextCursor_hasSelection(void* self) {
	auto&& __r = ((QTextCursor*)self)->hasSelection();
	return __r;
}
int QTextCursor_hasComplexSelection(void* self) {
	auto&& __r = ((QTextCursor*)self)->hasComplexSelection();
	return __r;
}
void QTextCursor_removeSelectedText(void* self) {
	((QTextCursor*)self)->removeSelectedText();
}
void QTextCursor_clearSelection(void* self) {
	((QTextCursor*)self)->clearSelection();
}
int QTextCursor_selectionStart(void* self) {
	auto&& __r = ((QTextCursor*)self)->selectionStart();
	return __r;
}
int QTextCursor_selectionEnd(void* self) {
	auto&& __r = ((QTextCursor*)self)->selectionEnd();
	return __r;
}
char* QTextCursor_selectedText(void* self, long long* __out_len) {
	auto&& __r = ((QTextCursor*)self)->selectedText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextCursor_selection(void* self) {
	auto&& __r = ((QTextCursor*)self)->selection();
	return new QTextDocumentFragment(__r);
}
void QTextCursor_selectedTableCells(void* self, int* firstRow, int* numRows, int* firstColumn, int* numColumns) {
	int firstRow_local{};
	int numRows_local{};
	int firstColumn_local{};
	int numColumns_local{};
	((QTextCursor*)self)->selectedTableCells(&firstRow_local, &numRows_local, &firstColumn_local, &numColumns_local);
	*firstRow = firstRow_local;
	*numRows = numRows_local;
	*firstColumn = firstColumn_local;
	*numColumns = numColumns_local;
}
void* QTextCursor_block(void* self) {
	auto&& __r = ((QTextCursor*)self)->block();
	return new QTextBlock(__r);
}
void* QTextCursor_charFormat(void* self) {
	auto&& __r = ((QTextCursor*)self)->charFormat();
	return new QTextCharFormat(__r);
}
void QTextCursor_setCharFormat(void* self, void* format) {
	((QTextCursor*)self)->setCharFormat(*(QTextCharFormat*)format);
}
void QTextCursor_mergeCharFormat(void* self, void* modifier) {
	((QTextCursor*)self)->mergeCharFormat(*(QTextCharFormat*)modifier);
}
void* QTextCursor_blockFormat(void* self) {
	auto&& __r = ((QTextCursor*)self)->blockFormat();
	return new QTextBlockFormat(__r);
}
void QTextCursor_setBlockFormat(void* self, void* format) {
	((QTextCursor*)self)->setBlockFormat(*(QTextBlockFormat*)format);
}
void QTextCursor_mergeBlockFormat(void* self, void* modifier) {
	((QTextCursor*)self)->mergeBlockFormat(*(QTextBlockFormat*)modifier);
}
void* QTextCursor_blockCharFormat(void* self) {
	auto&& __r = ((QTextCursor*)self)->blockCharFormat();
	return new QTextCharFormat(__r);
}
void QTextCursor_setBlockCharFormat(void* self, void* format) {
	((QTextCursor*)self)->setBlockCharFormat(*(QTextCharFormat*)format);
}
void QTextCursor_mergeBlockCharFormat(void* self, void* modifier) {
	((QTextCursor*)self)->mergeBlockCharFormat(*(QTextCharFormat*)modifier);
}
int QTextCursor_atBlockStart(void* self) {
	auto&& __r = ((QTextCursor*)self)->atBlockStart();
	return __r;
}
int QTextCursor_atBlockEnd(void* self) {
	auto&& __r = ((QTextCursor*)self)->atBlockEnd();
	return __r;
}
int QTextCursor_atStart(void* self) {
	auto&& __r = ((QTextCursor*)self)->atStart();
	return __r;
}
int QTextCursor_atEnd(void* self) {
	auto&& __r = ((QTextCursor*)self)->atEnd();
	return __r;
}
void QTextCursor_insertBlock(void* self) {
	((QTextCursor*)self)->insertBlock();
}
void QTextCursor_insertBlock2(void* self, void* format) {
	((QTextCursor*)self)->insertBlock(*(QTextBlockFormat*)format);
}
void QTextCursor_insertBlock3(void* self, void* format, void* charFormat) {
	((QTextCursor*)self)->insertBlock(*(QTextBlockFormat*)format, *(QTextCharFormat*)charFormat);
}
void* QTextCursor_insertList(void* self, void* format) {
	auto&& __r = ((QTextCursor*)self)->insertList(*(QTextListFormat*)format);
	return (void*)(__r);
}
void* QTextCursor_insertList2(void* self, int style) {
	auto&& __r = ((QTextCursor*)self)->insertList((QTextListFormat::Style)(style));
	return (void*)(__r);
}
void* QTextCursor_createList(void* self, void* format) {
	auto&& __r = ((QTextCursor*)self)->createList(*(QTextListFormat*)format);
	return (void*)(__r);
}
void* QTextCursor_createList2(void* self, int style) {
	auto&& __r = ((QTextCursor*)self)->createList((QTextListFormat::Style)(style));
	return (void*)(__r);
}
void* QTextCursor_currentList(void* self) {
	auto&& __r = ((QTextCursor*)self)->currentList();
	return (void*)(__r);
}
void* QTextCursor_insertTable(void* self, int rows, int cols, void* format) {
	auto&& __r = ((QTextCursor*)self)->insertTable(rows, cols, *(QTextTableFormat*)format);
	return (void*)(__r);
}
void* QTextCursor_insertTable2(void* self, int rows, int cols) {
	auto&& __r = ((QTextCursor*)self)->insertTable(rows, cols);
	return (void*)(__r);
}
void* QTextCursor_currentTable(void* self) {
	auto&& __r = ((QTextCursor*)self)->currentTable();
	return (void*)(__r);
}
void* QTextCursor_insertFrame(void* self, void* format) {
	auto&& __r = ((QTextCursor*)self)->insertFrame(*(QTextFrameFormat*)format);
	return (void*)(__r);
}
void* QTextCursor_currentFrame(void* self) {
	auto&& __r = ((QTextCursor*)self)->currentFrame();
	return (void*)(__r);
}
void QTextCursor_insertFragment(void* self, void* fragment) {
	((QTextCursor*)self)->insertFragment(*(QTextDocumentFragment*)fragment);
}
void QTextCursor_insertHtml(void* self, const char* html, long long html_len) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	((QTextCursor*)self)->insertHtml(html_qstr);
}
void QTextCursor_insertMarkdown(void* self, const char* markdown, long long markdown_len) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	((QTextCursor*)self)->insertMarkdown(markdown_qstr);
}
void QTextCursor_insertImage(void* self, void* format, int alignment) {
	((QTextCursor*)self)->insertImage(*(QTextImageFormat*)format, (QTextFrameFormat::Position)(alignment));
}
void QTextCursor_insertImage2(void* self, void* format) {
	((QTextCursor*)self)->insertImage(*(QTextImageFormat*)format);
}
void QTextCursor_insertImage3(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QTextCursor*)self)->insertImage(name_qstr);
}
void QTextCursor_insertImage4(void* self, void* image) {
	((QTextCursor*)self)->insertImage(*(QImage*)image);
}
void QTextCursor_beginEditBlock(void* self) {
	((QTextCursor*)self)->beginEditBlock();
}
void QTextCursor_joinPreviousEditBlock(void* self) {
	((QTextCursor*)self)->joinPreviousEditBlock();
}
void QTextCursor_endEditBlock(void* self) {
	((QTextCursor*)self)->endEditBlock();
}
int QTextCursor_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = ((QTextCursor*)self)->operator!=(*(QTextCursor*)rhs);
	return __r;
}
int QTextCursor_OperatorLesser(void* self, void* rhs) {
	auto&& __r = ((QTextCursor*)self)->operator<(*(QTextCursor*)rhs);
	return __r;
}
int QTextCursor_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = ((QTextCursor*)self)->operator<=(*(QTextCursor*)rhs);
	return __r;
}
int QTextCursor_OperatorEqual(void* self, void* rhs) {
	auto&& __r = ((QTextCursor*)self)->operator==(*(QTextCursor*)rhs);
	return __r;
}
int QTextCursor_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = ((QTextCursor*)self)->operator>=(*(QTextCursor*)rhs);
	return __r;
}
int QTextCursor_OperatorGreater(void* self, void* rhs) {
	auto&& __r = ((QTextCursor*)self)->operator>(*(QTextCursor*)rhs);
	return __r;
}
int QTextCursor_isCopyOf(void* self, void* other) {
	auto&& __r = ((QTextCursor*)self)->isCopyOf(*(QTextCursor*)other);
	return __r;
}
int QTextCursor_blockNumber(void* self) {
	auto&& __r = ((QTextCursor*)self)->blockNumber();
	return __r;
}
int QTextCursor_columnNumber(void* self) {
	auto&& __r = ((QTextCursor*)self)->columnNumber();
	return __r;
}
void* QTextCursor_document(void* self) {
	auto&& __r = ((QTextCursor*)self)->document();
	return (void*)(__r);
}
void QTextCursor_setPosition2(void* self, int pos, int mode) {
	((QTextCursor*)self)->setPosition(pos, (QTextCursor::MoveMode)(mode));
}
int QTextCursor_movePosition2(void* self, int op, int arg2) {
	auto&& __r = ((QTextCursor*)self)->movePosition((QTextCursor::MoveOperation)(op), (QTextCursor::MoveMode)(arg2));
	return __r;
}
int QTextCursor_movePosition3(void* self, int op, int arg2, int n) {
	auto&& __r = ((QTextCursor*)self)->movePosition((QTextCursor::MoveOperation)(op), (QTextCursor::MoveMode)(arg2), n);
	return __r;
}
void QTextCursor_insertMarkdown2(void* self, const char* markdown, long long markdown_len, int features) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	((QTextCursor*)self)->insertMarkdown(markdown_qstr, QFlags<QTextDocument::MarkdownFeature>(features));
}
void QTextCursor_insertImage5(void* self, void* image, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QTextCursor*)self)->insertImage(*(QImage*)image, name_qstr);
}
void QTextCursor_delete(void* self) { delete (QTextCursor*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qvariant.h>
#include "gen_qaccessibleattributesinterface.h"
#include "gen_purego_cb.h"

class QAccessibleAttributesInterface__Go : public QAccessibleAttributesInterface {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QAccessibleAttributesInterface__Go(long long id, A&&... a) : QAccessibleAttributesInterface(std::forward<A>(a)...), __goId(id) {}
	QList<QAccessible::Attribute> attributeKeys() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 9162027173711402047LL, &__len, &__h);
			if (__h) {
				QList<QAccessible::Attribute> __r; __r.reserve(__len);
				int* __aa = (int*)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append((QAccessible::Attribute)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QAccessible::Attribute>{};
	}
	QVariant attributeValue(QAccessible::Attribute key) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 4705686874656021092LL, (int)(key), &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
};

extern "C" {

void QAccessibleAttributesInterface__set_override(void* self, int idx) { ((QAccessibleAttributesInterface__Go*)self)->__goOn[idx] = 1; }
int* QAccessibleAttributesInterface_attributeKeys(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleAttributesInterface*)self)->attributeKeys();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
void* QAccessibleAttributesInterface_attributeValue(void* self, int key) {
	auto&& __r = ((QAccessibleAttributesInterface*)self)->attributeValue((QAccessible::Attribute)(key));
	return new QVariant(__r);
}
void* QAccessibleAttributesInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleAttributesInterface*)self)->operator=(*(QAccessibleAttributesInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleAttributesInterface__Go_new(long long id) { return new QAccessibleAttributesInterface__Go(id); }
void QAccessibleAttributesInterface_delete(void* self) { delete (QAccessibleAttributesInterface*)self; }

}

#include <QtGui/qclipboard.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qmimedata.h>
#include <QtGui/qimage.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qclipboard.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QClipboard;
static std::mutex __connsMu_QClipboard;

extern "C" {

void* QClipboard_base(void* self) { return (void*)static_cast<QObject*>((QClipboard*)self); }
void* QClipboard_fromQObject(void* o) { return (void*)qobject_cast<QClipboard*>((QObject*)o); }
char* QClipboard_tr(const char* s, long long* __out_len) {
	auto&& __r = QClipboard::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QClipboard_clear(void* self) {
	((QClipboard*)self)->clear();
}
int QClipboard_supportsSelection(void* self) {
	auto&& __r = ((QClipboard*)self)->supportsSelection();
	return __r;
}
int QClipboard_supportsFindBuffer(void* self) {
	auto&& __r = ((QClipboard*)self)->supportsFindBuffer();
	return __r;
}
int QClipboard_ownsSelection(void* self) {
	auto&& __r = ((QClipboard*)self)->ownsSelection();
	return __r;
}
int QClipboard_ownsClipboard(void* self) {
	auto&& __r = ((QClipboard*)self)->ownsClipboard();
	return __r;
}
int QClipboard_ownsFindBuffer(void* self) {
	auto&& __r = ((QClipboard*)self)->ownsFindBuffer();
	return __r;
}
char* QClipboard_text(void* self, long long* __out_len) {
	auto&& __r = ((QClipboard*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QClipboard_text2(void* self, void* subtype, long long* __out_len) {
	auto&& __r = ((QClipboard*)self)->text(*(QString*)subtype);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QClipboard_setText(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QClipboard*)self)->setText(arg1_qstr);
}
void* QClipboard_mimeData(void* self) {
	auto&& __r = ((QClipboard*)self)->mimeData();
	return (void*)(__r);
}
void QClipboard_setMimeData(void* self, void* data) {
	((QClipboard*)self)->setMimeData((QMimeData*)data);
}
void* QClipboard_image(void* self) {
	auto&& __r = ((QClipboard*)self)->image();
	return new QImage(__r);
}
void* QClipboard_pixmap(void* self) {
	auto&& __r = ((QClipboard*)self)->pixmap();
	return new QPixmap(__r);
}
void QClipboard_setImage(void* self, void* arg1) {
	((QClipboard*)self)->setImage(*(QImage*)arg1);
}
void QClipboard_setPixmap(void* self, void* arg1) {
	((QClipboard*)self)->setPixmap(*(QPixmap*)arg1);
}
void QClipboard_changed(void* self, int mode) {
	((QClipboard*)self)->changed((QClipboard::Mode)(mode));
}
void QClipboard_selectionChanged(void* self) {
	((QClipboard*)self)->selectionChanged();
}
void QClipboard_findBufferChanged(void* self) {
	((QClipboard*)self)->findBufferChanged();
}
void QClipboard_dataChanged(void* self) {
	((QClipboard*)self)->dataChanged();
}
void* QClipboard_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QClipboard::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QClipboard_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QClipboard::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QClipboard_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QClipboard::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QClipboard_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QClipboard::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QClipboard_clear2(void* self, int mode) {
	((QClipboard*)self)->clear((QClipboard::Mode)(mode));
}
char* QClipboard_text3(void* self, int mode, long long* __out_len) {
	auto&& __r = ((QClipboard*)self)->text((QClipboard::Mode)(mode));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QClipboard_text4(void* self, void* subtype, int mode, long long* __out_len) {
	auto&& __r = ((QClipboard*)self)->text(*(QString*)subtype, (QClipboard::Mode)(mode));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QClipboard_setText2(void* self, const char* arg1, long long arg1_len, int mode) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QClipboard*)self)->setText(arg1_qstr, (QClipboard::Mode)(mode));
}
void* QClipboard_mimeData2(void* self, int mode) {
	auto&& __r = ((QClipboard*)self)->mimeData((QClipboard::Mode)(mode));
	return (void*)(__r);
}
void QClipboard_setMimeData2(void* self, void* data, int mode) {
	((QClipboard*)self)->setMimeData((QMimeData*)data, (QClipboard::Mode)(mode));
}
void* QClipboard_image2(void* self, int mode) {
	auto&& __r = ((QClipboard*)self)->image((QClipboard::Mode)(mode));
	return new QImage(__r);
}
void* QClipboard_pixmap2(void* self, int mode) {
	auto&& __r = ((QClipboard*)self)->pixmap((QClipboard::Mode)(mode));
	return new QPixmap(__r);
}
void QClipboard_setImage2(void* self, void* arg1, int mode) {
	((QClipboard*)self)->setImage(*(QImage*)arg1, (QClipboard::Mode)(mode));
}
void QClipboard_setPixmap2(void* self, void* arg1, int mode) {
	((QClipboard*)self)->setPixmap(*(QPixmap*)arg1, (QClipboard::Mode)(mode));
}
void* QClipboard_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QClipboard::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QClipboard_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QClipboard::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QClipboard_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::changed, (QClipboard*)self, [__id](QClipboard::Mode mode){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(mode));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::selectionChanged, (QClipboard*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard_connect_findBufferChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::findBufferChanged, (QClipboard*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard_connect_dataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::dataChanged, (QClipboard*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::destroyed, (QClipboard*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::destroyed, (QClipboard*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QClipboard*)self, &QClipboard::objectNameChanged, (QClipboard*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
	__conns_QClipboard[__id] = __c;
}
void QClipboard__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QClipboard);
		auto __it = __conns_QClipboard.find(id);
		if (__it == __conns_QClipboard.end()) return;
		__c = __it->second;
		__conns_QClipboard.erase(__it);
	}
	QObject::disconnect(__c);
}

}

#include <QtGui/qopenglcontext.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qsurfaceformat.h>
#include <QtGui/qscreen.h>
#include <QtGui/qsurface.h>
#include <QtCore/qbytearray.h>
#include <QtGui/qopenglfunctions.h>
#include <QtGui/qopenglextrafunctions.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qopenglcontext.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QOpenGLContext__Go : public QOpenGLContext {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QOpenGLContext__Go(long long id, A&&... a) : QOpenGLContext(std::forward<A>(a)...), __goId(id) {}
	using QOpenGLContext::sender;
	using QOpenGLContext::senderSignalIndex;
	using QOpenGLContext::receivers;
	using QOpenGLContext::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QOpenGLContext::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QOpenGLContext::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QOpenGLContext::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QOpenGLContext::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QOpenGLContext::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QOpenGLContext::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QOpenGLContext::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QOpenGLContext::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QOpenGLContext::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QOpenGLContext::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QOpenGLContext::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QOpenGLContext::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QOpenGLContext::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QOpenGLContext;
static std::mutex __connsMu_QOpenGLContext;

extern "C" {

void* QOpenGLContext_base(void* self) { return (void*)static_cast<QObject*>((QOpenGLContext*)self); }
void QOpenGLContext__set_override(void* self, int idx) { ((QOpenGLContext__Go*)self)->__goOn[idx] = 1; }
void* QOpenGLContext_fromQObject(void* o) { return (void*)qobject_cast<QOpenGLContext*>((QObject*)o); }
void* QOpenGLContext_new(long long __goid) {
	return new QOpenGLContext__Go(__goid);
}
void* QOpenGLContext_new2(long long __goid, void* parent) {
	return new QOpenGLContext__Go(__goid, (QObject*)parent);
}
char* QOpenGLContext_tr(const char* s, long long* __out_len) {
	auto&& __r = QOpenGLContext::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QOpenGLContext_setFormat(void* self, void* format) {
	((QOpenGLContext*)self)->setFormat(*(QSurfaceFormat*)format);
}
void QOpenGLContext_setShareContext(void* self, void* shareContext) {
	((QOpenGLContext*)self)->setShareContext((QOpenGLContext*)shareContext);
}
void QOpenGLContext_setScreen(void* self, void* screen) {
	((QOpenGLContext*)self)->setScreen((QScreen*)screen);
}
int QOpenGLContext_create(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->create();
	return __r;
}
int QOpenGLContext_isValid(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->isValid();
	return __r;
}
void* QOpenGLContext_format(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->format();
	return new QSurfaceFormat(__r);
}
void* QOpenGLContext_shareContext(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->shareContext();
	return (void*)(__r);
}
void* QOpenGLContext_shareGroup(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->shareGroup();
	return (void*)(__r);
}
void* QOpenGLContext_screen(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->screen();
	return (void*)(__r);
}
unsigned int QOpenGLContext_defaultFramebufferObject(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->defaultFramebufferObject();
	return __r;
}
int QOpenGLContext_makeCurrent(void* self, void* surface) {
	auto&& __r = ((QOpenGLContext*)self)->makeCurrent((QSurface*)surface);
	return __r;
}
void QOpenGLContext_doneCurrent(void* self) {
	((QOpenGLContext*)self)->doneCurrent();
}
void QOpenGLContext_swapBuffers(void* self, void* surface) {
	((QOpenGLContext*)self)->swapBuffers((QSurface*)surface);
}
void* QOpenGLContext_surface(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->surface();
	return (void*)(__r);
}
void* QOpenGLContext_currentContext() {
	auto&& __r = QOpenGLContext::currentContext();
	return (void*)(__r);
}
int QOpenGLContext_areSharing(void* first, void* second) {
	auto&& __r = QOpenGLContext::areSharing((QOpenGLContext*)first, (QOpenGLContext*)second);
	return __r;
}
void* QOpenGLContext_functions(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->functions();
	return (void*)(__r);
}
void* QOpenGLContext_extraFunctions(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->extraFunctions();
	return (void*)(__r);
}
void** QOpenGLContext_extensions(void* self, long long* __out_len) {
	auto&& __r = ((QOpenGLContext*)self)->extensions();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
int QOpenGLContext_hasExtension(void* self, void* extension) {
	auto&& __r = ((QOpenGLContext*)self)->hasExtension(*(QByteArray*)extension);
	return __r;
}
int QOpenGLContext_openGLModuleType() {
	auto&& __r = QOpenGLContext::openGLModuleType();
	return (int)(__r);
}
int QOpenGLContext_isOpenGLES(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->isOpenGLES();
	return __r;
}
int QOpenGLContext_supportsThreadedOpenGL() {
	auto&& __r = QOpenGLContext::supportsThreadedOpenGL();
	return __r;
}
void* QOpenGLContext_globalShareContext() {
	auto&& __r = QOpenGLContext::globalShareContext();
	return (void*)(__r);
}
void QOpenGLContext_aboutToBeDestroyed(void* self) {
	((QOpenGLContext*)self)->aboutToBeDestroyed();
}
void* QOpenGLContext_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QOpenGLContext::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QOpenGLContext_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QOpenGLContext::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QOpenGLContext_sender(void* self) {
	auto&& __r = ((QOpenGLContext__Go*)self)->sender();
	return (void*)(__r);
}
int QOpenGLContext_senderSignalIndex(void* self) {
	auto&& __r = ((QOpenGLContext__Go*)self)->senderSignalIndex();
	return __r;
}
int QOpenGLContext_receivers(void* self, const char* signal) {
	auto&& __r = ((QOpenGLContext__Go*)self)->receivers(signal);
	return __r;
}
int QOpenGLContext_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QOpenGLContext__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QOpenGLContext_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QOpenGLContext::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QOpenGLContext_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QOpenGLContext::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QOpenGLContext_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QOpenGLContext::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QOpenGLContext_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QOpenGLContext::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QOpenGLContext_metaObject_default(void* self) {
	auto&& __r = ((QOpenGLContext*)self)->QOpenGLContext::metaObject();
	return (void*)(__r);
}
int QOpenGLContext_event_default(void* self, void* event) {
	auto&& __r = ((QOpenGLContext*)self)->QOpenGLContext::event((QEvent*)event);
	return __r;
}
int QOpenGLContext_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QOpenGLContext*)self)->QOpenGLContext::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QOpenGLContext_timerEvent_default(void* self, void* event) {
	((QOpenGLContext__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QOpenGLContext_childEvent_default(void* self, void* event) {
	((QOpenGLContext__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QOpenGLContext_customEvent_default(void* self, void* event) {
	((QOpenGLContext__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QOpenGLContext_connectNotify_default(void* self, void* signal) {
	((QOpenGLContext__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QOpenGLContext_disconnectNotify_default(void* self, void* signal) {
	((QOpenGLContext__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QOpenGLContext_connect_aboutToBeDestroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContext*)self, &QOpenGLContext::aboutToBeDestroyed, (QOpenGLContext*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContext);
	__conns_QOpenGLContext[__id] = __c;
}
void QOpenGLContext_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContext*)self, &QOpenGLContext::destroyed, (QOpenGLContext*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContext);
	__conns_QOpenGLContext[__id] = __c;
}
void QOpenGLContext_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContext*)self, &QOpenGLContext::destroyed, (QOpenGLContext*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContext);
	__conns_QOpenGLContext[__id] = __c;
}
void QOpenGLContext_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QOpenGLContext*)self, &QOpenGLContext::objectNameChanged, (QOpenGLContext*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContext);
	__conns_QOpenGLContext[__id] = __c;
}
void QOpenGLContext__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QOpenGLContext);
		auto __it = __conns_QOpenGLContext.find(id);
		if (__it == __conns_QOpenGLContext.end()) return;
		__c = __it->second;
		__conns_QOpenGLContext.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QOpenGLContext__Go_new(long long id, void* parent) { return new QOpenGLContext__Go(id, (QObject*)parent); }
void QOpenGLContext_delete(void* self) { delete (QOpenGLContext*)self; }

}

#include <QtGui/qquaternion.h>
#include <QtGui/qvectornd.h>
#include "gen_qquaternion__axis.h"

extern "C" {

void* QQuaternion__Axis_fromVector3D(void* v) {
	auto&& __r = QQuaternion::Axis::fromVector3D(*(QVector3D*)v);
	return new QQuaternion::Axis(__r);
}
void* QQuaternion__Axis_toVector3D(void* self) {
	auto&& __r = ((QQuaternion::Axis*)self)->toVector3D();
	return new QVector3D(__r);
}
float QQuaternion__Axis_x_get(void* self) {
	auto __r = ((QQuaternion::Axis*)self)->x;
	return __r;
}
void QQuaternion__Axis_x_set(void* self, float x) {
	((QQuaternion::Axis*)self)->x = x;
}
float QQuaternion__Axis_y_get(void* self) {
	auto __r = ((QQuaternion::Axis*)self)->y;
	return __r;
}
void QQuaternion__Axis_y_set(void* self, float y) {
	((QQuaternion::Axis*)self)->y = y;
}
float QQuaternion__Axis_z_get(void* self) {
	auto __r = ((QQuaternion::Axis*)self)->z;
	return __r;
}
void QQuaternion__Axis_z_set(void* self, float z) {
	((QQuaternion::Axis*)self)->z = z;
}
void QQuaternion__Axis_delete(void* self) { delete (QQuaternion::Axis*)self; }

}

#include <QtGui/qundostack.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qaction.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qundostack.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QUndoStack__Go : public QUndoStack {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QUndoStack__Go(long long id, A&&... a) : QUndoStack(std::forward<A>(a)...), __goId(id) {}
	using QUndoStack::sender;
	using QUndoStack::senderSignalIndex;
	using QUndoStack::receivers;
	using QUndoStack::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QUndoStack::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QUndoStack::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QUndoStack::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QUndoStack::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QUndoStack::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QUndoStack::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QUndoStack::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QUndoStack::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QUndoStack::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QUndoStack::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QUndoStack::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QUndoStack::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QUndoStack::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QUndoStack;
static std::mutex __connsMu_QUndoStack;

extern "C" {

void* QUndoStack_base(void* self) { return (void*)static_cast<QObject*>((QUndoStack*)self); }
void QUndoStack__set_override(void* self, int idx) { ((QUndoStack__Go*)self)->__goOn[idx] = 1; }
void* QUndoStack_fromQObject(void* o) { return (void*)qobject_cast<QUndoStack*>((QObject*)o); }
void* QUndoStack_new(long long __goid) {
	return new QUndoStack__Go(__goid);
}
void* QUndoStack_new2(long long __goid, void* parent) {
	return new QUndoStack__Go(__goid, (QObject*)parent);
}
char* QUndoStack_tr(const char* s, long long* __out_len) {
	auto&& __r = QUndoStack::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUndoStack_clear(void* self) {
	((QUndoStack*)self)->clear();
}
void QUndoStack_push(void* self, void* cmd) {
	((QUndoStack*)self)->push((QUndoCommand*)cmd);
}
int QUndoStack_canUndo(void* self) {
	auto&& __r = ((QUndoStack*)self)->canUndo();
	return __r;
}
int QUndoStack_canRedo(void* self) {
	auto&& __r = ((QUndoStack*)self)->canRedo();
	return __r;
}
char* QUndoStack_undoText(void* self, long long* __out_len) {
	auto&& __r = ((QUndoStack*)self)->undoText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUndoStack_redoText(void* self, long long* __out_len) {
	auto&& __r = ((QUndoStack*)self)->redoText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QUndoStack_count(void* self) {
	auto&& __r = ((QUndoStack*)self)->count();
	return __r;
}
int QUndoStack_index(void* self) {
	auto&& __r = ((QUndoStack*)self)->index();
	return __r;
}
char* QUndoStack_text(void* self, int idx, long long* __out_len) {
	auto&& __r = ((QUndoStack*)self)->text(idx);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUndoStack_createUndoAction(void* self, void* parent) {
	auto&& __r = ((QUndoStack*)self)->createUndoAction((QObject*)parent);
	return (void*)(__r);
}
void* QUndoStack_createRedoAction(void* self, void* parent) {
	auto&& __r = ((QUndoStack*)self)->createRedoAction((QObject*)parent);
	return (void*)(__r);
}
int QUndoStack_isActive(void* self) {
	auto&& __r = ((QUndoStack*)self)->isActive();
	return __r;
}
int QUndoStack_isClean(void* self) {
	auto&& __r = ((QUndoStack*)self)->isClean();
	return __r;
}
int QUndoStack_cleanIndex(void* self) {
	auto&& __r = ((QUndoStack*)self)->cleanIndex();
	return __r;
}
void QUndoStack_beginMacro(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QUndoStack*)self)->beginMacro(text_qstr);
}
void QUndoStack_endMacro(void* self) {
	((QUndoStack*)self)->endMacro();
}
void QUndoStack_setUndoLimit(void* self, int limit) {
	((QUndoStack*)self)->setUndoLimit(limit);
}
int QUndoStack_undoLimit(void* self) {
	auto&& __r = ((QUndoStack*)self)->undoLimit();
	return __r;
}
void* QUndoStack_command(void* self, int index) {
	auto&& __r = ((QUndoStack*)self)->command(index);
	return (void*)(__r);
}
void QUndoStack_setClean(void* self) {
	((QUndoStack*)self)->setClean();
}
void QUndoStack_resetClean(void* self) {
	((QUndoStack*)self)->resetClean();
}
void QUndoStack_setIndex(void* self, int idx) {
	((QUndoStack*)self)->setIndex(idx);
}
void QUndoStack_undo(void* self) {
	((QUndoStack*)self)->undo();
}
void QUndoStack_redo(void* self) {
	((QUndoStack*)self)->redo();
}
void QUndoStack_setActive(void* self) {
	((QUndoStack*)self)->setActive();
}
void QUndoStack_indexChanged(void* self, int idx) {
	((QUndoStack*)self)->indexChanged(idx);
}
void QUndoStack_cleanChanged(void* self, int clean) {
	((QUndoStack*)self)->cleanChanged(clean);
}
void QUndoStack_canUndoChanged(void* self, int canUndo) {
	((QUndoStack*)self)->canUndoChanged(canUndo);
}
void QUndoStack_canRedoChanged(void* self, int canRedo) {
	((QUndoStack*)self)->canRedoChanged(canRedo);
}
void QUndoStack_undoTextChanged(void* self, const char* undoText, long long undoText_len) {
	QString undoText_qstr = QString::fromUtf8(undoText, undoText_len);
	((QUndoStack*)self)->undoTextChanged(undoText_qstr);
}
void QUndoStack_redoTextChanged(void* self, const char* redoText, long long redoText_len) {
	QString redoText_qstr = QString::fromUtf8(redoText, redoText_len);
	((QUndoStack*)self)->redoTextChanged(redoText_qstr);
}
void* QUndoStack_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QUndoStack::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QUndoStack_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QUndoStack::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QUndoStack_sender(void* self) {
	auto&& __r = ((QUndoStack__Go*)self)->sender();
	return (void*)(__r);
}
int QUndoStack_senderSignalIndex(void* self) {
	auto&& __r = ((QUndoStack__Go*)self)->senderSignalIndex();
	return __r;
}
int QUndoStack_receivers(void* self, const char* signal) {
	auto&& __r = ((QUndoStack__Go*)self)->receivers(signal);
	return __r;
}
int QUndoStack_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QUndoStack__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QUndoStack_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QUndoStack::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUndoStack_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QUndoStack::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUndoStack_createUndoAction2(void* self, void* parent, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	auto&& __r = ((QUndoStack*)self)->createUndoAction((QObject*)parent, prefix_qstr);
	return (void*)(__r);
}
void* QUndoStack_createRedoAction2(void* self, void* parent, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	auto&& __r = ((QUndoStack*)self)->createRedoAction((QObject*)parent, prefix_qstr);
	return (void*)(__r);
}
void QUndoStack_setActive2(void* self, int active) {
	((QUndoStack*)self)->setActive(active);
}
void* QUndoStack_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QUndoStack::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QUndoStack_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QUndoStack::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QUndoStack_metaObject_default(void* self) {
	auto&& __r = ((QUndoStack*)self)->QUndoStack::metaObject();
	return (void*)(__r);
}
int QUndoStack_event_default(void* self, void* event) {
	auto&& __r = ((QUndoStack*)self)->QUndoStack::event((QEvent*)event);
	return __r;
}
int QUndoStack_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QUndoStack*)self)->QUndoStack::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QUndoStack_timerEvent_default(void* self, void* event) {
	((QUndoStack__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QUndoStack_childEvent_default(void* self, void* event) {
	((QUndoStack__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QUndoStack_customEvent_default(void* self, void* event) {
	((QUndoStack__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QUndoStack_connectNotify_default(void* self, void* signal) {
	((QUndoStack__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QUndoStack_disconnectNotify_default(void* self, void* signal) {
	((QUndoStack__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QUndoStack_connect_indexChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::indexChanged, (QUndoStack*)self, [__id](int idx){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, idx);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_cleanChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::cleanChanged, (QUndoStack*)self, [__id](bool clean){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, clean);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_canUndoChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::canUndoChanged, (QUndoStack*)self, [__id](bool canUndo){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, canUndo);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_canRedoChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::canRedoChanged, (QUndoStack*)self, [__id](bool canRedo){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, canRedo);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_undoTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::undoTextChanged, (QUndoStack*)self, [__id](const QString& undoText){
		QByteArray undoText_ba = undoText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)undoText_ba.constData(), undoText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_redoTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::redoTextChanged, (QUndoStack*)self, [__id](const QString& redoText){
		QByteArray redoText_ba = redoText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)redoText_ba.constData(), redoText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::destroyed, (QUndoStack*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::destroyed, (QUndoStack*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoStack*)self, &QUndoStack::objectNameChanged, (QUndoStack*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
	__conns_QUndoStack[__id] = __c;
}
void QUndoStack__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QUndoStack);
		auto __it = __conns_QUndoStack.find(id);
		if (__it == __conns_QUndoStack.end()) return;
		__c = __it->second;
		__conns_QUndoStack.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QUndoStack__Go_new(long long id, void* parent) { return new QUndoStack__Go(id, (QObject*)parent); }
void QUndoStack_delete(void* self) { delete (QUndoStack*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include "gen_qstyleoptionheaderv2.h"

extern "C" {

void* QStyleOptionHeaderV2_base(void* self) { return (void*)static_cast<QStyleOptionHeader*>((QStyleOptionHeaderV2*)self); }
void QStyleOptionHeaderV2_chain(void* self, void** out) {
	QStyleOptionHeaderV2* __p = (QStyleOptionHeaderV2*)self;
	QStyleOptionHeader* __b0 = static_cast<QStyleOptionHeader*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionHeaderV2_new() {
	return new QStyleOptionHeaderV2();
}
void* QStyleOptionHeaderV2_new2(void* other) {
	return new QStyleOptionHeaderV2(*(QStyleOptionHeaderV2*)other);
}
void* QStyleOptionHeaderV2_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionHeaderV2*)self)->operator=(*(QStyleOptionHeaderV2*)arg1);
	return (void*)&(__r);
}
int QStyleOptionHeaderV2_textElideMode_get(void* self) {
	auto __r = ((QStyleOptionHeaderV2*)self)->textElideMode;
	return (int)(__r);
}
void QStyleOptionHeaderV2_textElideMode_set(void* self, int textElideMode) {
	((QStyleOptionHeaderV2*)self)->textElideMode = (Qt::TextElideMode)(textElideMode);
}
int QStyleOptionHeaderV2_isSectionDragTarget_get(void* self) {
	auto __r = ((QStyleOptionHeaderV2*)self)->isSectionDragTarget;
	return __r;
}
void QStyleOptionHeaderV2_isSectionDragTarget_set(void* self, int isSectionDragTarget) {
	((QStyleOptionHeaderV2*)self)->isSectionDragTarget = isSectionDragTarget;
}
int QStyleOptionHeaderV2_unused_get(void* self) {
	auto __r = ((QStyleOptionHeaderV2*)self)->unused;
	return __r;
}
void QStyleOptionHeaderV2_unused_set(void* self, int unused) {
	((QStyleOptionHeaderV2*)self)->unused = unused;
}
void QStyleOptionHeaderV2_delete(void* self) { delete (QStyleOptionHeaderV2*)self; }

}

#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qscrollbar.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
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
#include "gen_qabstractscrollarea.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractScrollArea__Go : public QAbstractScrollArea {
public:
	long long __goId;
	unsigned char __goOn[53] = {};
	template<typename... A> QAbstractScrollArea__Go(long long id, A&&... a) : QAbstractScrollArea(std::forward<A>(a)...), __goId(id) {}
	using QAbstractScrollArea::setViewportMargins;
	using QAbstractScrollArea::viewportMargins;
	using QAbstractScrollArea::drawFrame;
	using QAbstractScrollArea::updateMicroFocus;
	using QAbstractScrollArea::create;
	using QAbstractScrollArea::destroy;
	using QAbstractScrollArea::focusNextChild;
	using QAbstractScrollArea::focusPreviousChild;
	using QAbstractScrollArea::sender;
	using QAbstractScrollArea::senderSignalIndex;
	using QAbstractScrollArea::receivers;
	using QAbstractScrollArea::isSignalConnected;
	using QAbstractScrollArea::getDecodedMetricF;
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QAbstractScrollArea::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QAbstractScrollArea::event(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QAbstractScrollArea::viewportEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QAbstractScrollArea::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QAbstractScrollArea::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QAbstractScrollArea::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QAbstractScrollArea::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QAbstractScrollArea::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QAbstractScrollArea::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QAbstractScrollArea::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QAbstractScrollArea::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QAbstractScrollArea::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QAbstractScrollArea::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QAbstractScrollArea::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QAbstractScrollArea::dropEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QAbstractScrollArea::keyPressEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QAbstractScrollArea::scrollContentsBy(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QAbstractScrollArea::viewportSizeHint(); }
	void cppgoBase_changeEvent(QEvent* __a0) { QAbstractScrollArea::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QAbstractScrollArea::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QAbstractScrollArea::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QAbstractScrollArea::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QAbstractScrollArea::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QAbstractScrollArea::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QAbstractScrollArea::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QAbstractScrollArea::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QAbstractScrollArea::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QAbstractScrollArea::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QAbstractScrollArea::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QAbstractScrollArea::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QAbstractScrollArea::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QAbstractScrollArea::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QAbstractScrollArea::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QAbstractScrollArea::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QAbstractScrollArea::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QAbstractScrollArea::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QAbstractScrollArea::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractScrollArea::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractScrollArea::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractScrollArea::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractScrollArea::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractScrollArea::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractScrollArea::metaObject();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractScrollArea::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractScrollArea::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QAbstractScrollArea::setupViewport(viewport);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QAbstractScrollArea::eventFilter(arg1, arg2);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QAbstractScrollArea::event(arg1);
	}
	bool viewportEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)arg1, &__h); if (__h) return __r; }
		return QAbstractScrollArea::viewportEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::resizeEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::paintEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::mouseDoubleClickEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::mouseMoveEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::contextMenuEvent(arg1);
	}
	void dragEnterEvent(QDragEnterEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::dragEnterEvent(arg1);
	}
	void dragMoveEvent(QDragMoveEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::dragMoveEvent(arg1);
	}
	void dragLeaveEvent(QDragLeaveEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::dragLeaveEvent(arg1);
	}
	void dropEvent(QDropEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::dropEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::keyPressEvent(arg1);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QAbstractScrollArea::scrollContentsBy(dx, dy);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractScrollArea::viewportSizeHint();
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QAbstractScrollArea::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QAbstractScrollArea::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QAbstractScrollArea::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QAbstractScrollArea::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QAbstractScrollArea::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QAbstractScrollArea::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QAbstractScrollArea::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QAbstractScrollArea::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QAbstractScrollArea::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QAbstractScrollArea::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QAbstractScrollArea::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractScrollArea::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractScrollArea::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QAbstractScrollArea::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractScrollArea::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractScrollArea::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractScrollArea::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractScrollArea;
static std::mutex __connsMu_QAbstractScrollArea;

extern "C" {

void* QAbstractScrollArea_base(void* self) { return (void*)static_cast<QFrame*>((QAbstractScrollArea*)self); }
void QAbstractScrollArea_chain(void* self, void** out) {
	QAbstractScrollArea* __p = (QAbstractScrollArea*)self;
	QFrame* __b0 = static_cast<QFrame*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QAbstractScrollArea__set_override(void* self, int idx) { ((QAbstractScrollArea__Go*)self)->__goOn[idx] = 1; }
void* QAbstractScrollArea_fromQObject(void* o) { return (void*)qobject_cast<QAbstractScrollArea*>((QObject*)o); }
void* QAbstractScrollArea_new(long long __goid, void* parent) {
	return new QAbstractScrollArea__Go(__goid, (QWidget*)parent);
}
void* QAbstractScrollArea_new2(long long __goid) {
	return new QAbstractScrollArea__Go(__goid);
}
char* QAbstractScrollArea_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractScrollArea::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractScrollArea_verticalScrollBarPolicy(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->verticalScrollBarPolicy();
	return (int)(__r);
}
void QAbstractScrollArea_setVerticalScrollBarPolicy(void* self, int arg1) {
	((QAbstractScrollArea*)self)->setVerticalScrollBarPolicy((Qt::ScrollBarPolicy)(arg1));
}
void* QAbstractScrollArea_verticalScrollBar(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->verticalScrollBar();
	return (void*)(__r);
}
void QAbstractScrollArea_setVerticalScrollBar(void* self, void* scrollbar) {
	((QAbstractScrollArea*)self)->setVerticalScrollBar((QScrollBar*)scrollbar);
}
int QAbstractScrollArea_horizontalScrollBarPolicy(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->horizontalScrollBarPolicy();
	return (int)(__r);
}
void QAbstractScrollArea_setHorizontalScrollBarPolicy(void* self, int arg1) {
	((QAbstractScrollArea*)self)->setHorizontalScrollBarPolicy((Qt::ScrollBarPolicy)(arg1));
}
void* QAbstractScrollArea_horizontalScrollBar(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->horizontalScrollBar();
	return (void*)(__r);
}
void QAbstractScrollArea_setHorizontalScrollBar(void* self, void* scrollbar) {
	((QAbstractScrollArea*)self)->setHorizontalScrollBar((QScrollBar*)scrollbar);
}
void* QAbstractScrollArea_cornerWidget(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->cornerWidget();
	return (void*)(__r);
}
void QAbstractScrollArea_setCornerWidget(void* self, void* widget) {
	((QAbstractScrollArea*)self)->setCornerWidget((QWidget*)widget);
}
void QAbstractScrollArea_addScrollBarWidget(void* self, void* widget, int alignment) {
	((QAbstractScrollArea*)self)->addScrollBarWidget((QWidget*)widget, QFlags<Qt::AlignmentFlag>(alignment));
}
void** QAbstractScrollArea_scrollBarWidgets(void* self, int alignment, long long* __out_len) {
	auto&& __r = ((QAbstractScrollArea*)self)->scrollBarWidgets(QFlags<Qt::AlignmentFlag>(alignment));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QAbstractScrollArea_viewport(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->viewport();
	return (void*)(__r);
}
void QAbstractScrollArea_setViewport(void* self, void* widget) {
	((QAbstractScrollArea*)self)->setViewport((QWidget*)widget);
}
void* QAbstractScrollArea_maximumViewportSize(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->maximumViewportSize();
	return new QSize(__r);
}
void QAbstractScrollArea_setupViewport(void* self, void* viewport) {
	((QAbstractScrollArea*)self)->setupViewport((QWidget*)viewport);
}
int QAbstractScrollArea_sizeAdjustPolicy(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->sizeAdjustPolicy();
	return (int)(__r);
}
void QAbstractScrollArea_setSizeAdjustPolicy(void* self, int policy) {
	((QAbstractScrollArea*)self)->setSizeAdjustPolicy((QAbstractScrollArea::SizeAdjustPolicy)(policy));
}
void QAbstractScrollArea_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QAbstractScrollArea__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QAbstractScrollArea_setViewportMargins2(void* self, void* margins) {
	((QAbstractScrollArea__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QAbstractScrollArea_viewportMargins(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QAbstractScrollArea_drawFrame(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->drawFrame((QPainter*)arg1);
}
void QAbstractScrollArea_setTabOrder(void* arg1, void* arg2) {
	QAbstractScrollArea::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QAbstractScrollArea_mouseGrabber() {
	auto&& __r = QAbstractScrollArea::mouseGrabber();
	return (void*)(__r);
}
void* QAbstractScrollArea_keyboardGrabber() {
	auto&& __r = QAbstractScrollArea::keyboardGrabber();
	return (void*)(__r);
}
void* QAbstractScrollArea_find(unsigned long long arg1) {
	auto&& __r = QAbstractScrollArea::find(arg1);
	return (void*)(__r);
}
void* QAbstractScrollArea_createWindowContainer(void* window) {
	auto&& __r = QAbstractScrollArea::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QAbstractScrollArea_updateMicroFocus(void* self) {
	((QAbstractScrollArea__Go*)self)->updateMicroFocus();
}
void QAbstractScrollArea_create(void* self) {
	((QAbstractScrollArea__Go*)self)->create();
}
void QAbstractScrollArea_destroy(void* self) {
	((QAbstractScrollArea__Go*)self)->destroy();
}
int QAbstractScrollArea_focusNextChild(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->focusNextChild();
	return __r;
}
int QAbstractScrollArea_focusPreviousChild(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->focusPreviousChild();
	return __r;
}
void* QAbstractScrollArea_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractScrollArea::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractScrollArea_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractScrollArea::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractScrollArea_sender(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractScrollArea_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractScrollArea_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractScrollArea_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QAbstractScrollArea_devicePixelRatioFScale() {
	auto&& __r = QAbstractScrollArea::devicePixelRatioFScale();
	return __r;
}
int QAbstractScrollArea_encodeMetricF(int metric, double value) {
	auto&& __r = QAbstractScrollArea::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QAbstractScrollArea_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QAbstractScrollArea_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractScrollArea::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractScrollArea_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractScrollArea::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractScrollArea_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QAbstractScrollArea::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QAbstractScrollArea_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QAbstractScrollArea::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QAbstractScrollArea_updateMicroFocus2(void* self, int query) {
	((QAbstractScrollArea__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QAbstractScrollArea_create2(void* self, unsigned long long arg1) {
	((QAbstractScrollArea__Go*)self)->create(arg1);
}
void QAbstractScrollArea_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QAbstractScrollArea__Go*)self)->create(arg1, initializeWindow);
}
void QAbstractScrollArea_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QAbstractScrollArea__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QAbstractScrollArea_destroy2(void* self, int destroyWindow) {
	((QAbstractScrollArea__Go*)self)->destroy(destroyWindow);
}
void QAbstractScrollArea_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QAbstractScrollArea__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QAbstractScrollArea_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractScrollArea::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractScrollArea_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractScrollArea::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAbstractScrollArea_setupViewport_default(void* self, void* viewport) {
	((QAbstractScrollArea*)self)->QAbstractScrollArea::setupViewport((QWidget*)viewport);
}
void* QAbstractScrollArea_metaObject_default(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::metaObject();
	return (void*)(__r);
}
void* QAbstractScrollArea_minimumSizeHint_default(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::minimumSizeHint();
	return new QSize(__r);
}
void* QAbstractScrollArea_sizeHint_default(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::sizeHint();
	return new QSize(__r);
}
int QAbstractScrollArea_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QAbstractScrollArea_event_default(void* self, void* arg1) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
int QAbstractScrollArea_viewportEvent_default(void* self, void* arg1) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_viewportEvent((QEvent*)arg1);
	return __r;
}
void QAbstractScrollArea_resizeEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QAbstractScrollArea_paintEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QAbstractScrollArea_mousePressEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QAbstractScrollArea_mouseReleaseEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QAbstractScrollArea_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QAbstractScrollArea_mouseMoveEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QAbstractScrollArea_wheelEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QAbstractScrollArea_contextMenuEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QAbstractScrollArea_dragEnterEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)arg1);
}
void QAbstractScrollArea_dragMoveEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)arg1);
}
void QAbstractScrollArea_dragLeaveEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)arg1);
}
void QAbstractScrollArea_dropEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_dropEvent((QDropEvent*)arg1);
}
void QAbstractScrollArea_keyPressEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QAbstractScrollArea_scrollContentsBy_default(void* self, int dx, int dy) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void* QAbstractScrollArea_viewportSizeHint_default(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QAbstractScrollArea_changeEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QAbstractScrollArea_initStyleOption_default(void* self, void* option) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QAbstractScrollArea_devType_default(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::devType();
	return __r;
}
void QAbstractScrollArea_setVisible_default(void* self, int visible) {
	((QAbstractScrollArea*)self)->QAbstractScrollArea::setVisible(visible);
}
int QAbstractScrollArea_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::heightForWidth(arg1);
	return __r;
}
int QAbstractScrollArea_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::hasHeightForWidth();
	return __r;
}
void* QAbstractScrollArea_paintEngine_default(void* self) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::paintEngine();
	return (void*)(__r);
}
void QAbstractScrollArea_keyReleaseEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QAbstractScrollArea_focusInEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QAbstractScrollArea_focusOutEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QAbstractScrollArea_enterEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QAbstractScrollArea_leaveEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QAbstractScrollArea_moveEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QAbstractScrollArea_closeEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QAbstractScrollArea_tabletEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QAbstractScrollArea_actionEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QAbstractScrollArea_showEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QAbstractScrollArea_hideEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QAbstractScrollArea_metric_default(void* self, int arg1) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QAbstractScrollArea_initPainter_default(void* self, void* painter) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QAbstractScrollArea_redirected_default(void* self, void* offset) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QAbstractScrollArea_sharedPainter_default(void* self) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QAbstractScrollArea_inputMethodEvent_default(void* self, void* arg1) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QAbstractScrollArea_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QAbstractScrollArea*)self)->QAbstractScrollArea::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QAbstractScrollArea_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QAbstractScrollArea__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QAbstractScrollArea_timerEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractScrollArea_childEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractScrollArea_customEvent_default(void* self, void* event) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractScrollArea_connectNotify_default(void* self, void* signal) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractScrollArea_disconnectNotify_default(void* self, void* signal) {
	((QAbstractScrollArea__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractScrollArea_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::windowTitleChanged, (QAbstractScrollArea*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::windowIconChanged, (QAbstractScrollArea*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::windowIconTextChanged, (QAbstractScrollArea*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::customContextMenuRequested, (QAbstractScrollArea*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::destroyed, (QAbstractScrollArea*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::destroyed, (QAbstractScrollArea*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractScrollArea*)self, &QAbstractScrollArea::objectNameChanged, (QAbstractScrollArea*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
	__conns_QAbstractScrollArea[__id] = __c;
}
void QAbstractScrollArea__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractScrollArea);
		auto __it = __conns_QAbstractScrollArea.find(id);
		if (__it == __conns_QAbstractScrollArea.end()) return;
		__c = __it->second;
		__conns_QAbstractScrollArea.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractScrollArea__Go_new(long long id) { return new QAbstractScrollArea__Go(id); }
void QAbstractScrollArea_delete(void* self) { delete (QAbstractScrollArea*)self; }

}

#include <QtWidgets/qcommonstyle.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtGui/qicon.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpalette.h>
#include <QtWidgets/qapplication.h>
#include <QtGui/qfontmetrics.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qcommonstyle.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QCommonStyle__Go : public QCommonStyle {
public:
	long long __goId;
	unsigned char __goOn[31] = {};
	template<typename... A> QCommonStyle__Go(long long id, A&&... a) : QCommonStyle(std::forward<A>(a)...), __goId(id) {}
	using QCommonStyle::sender;
	using QCommonStyle::senderSignalIndex;
	using QCommonStyle::receivers;
	using QCommonStyle::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QCommonStyle::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QCommonStyle::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QCommonStyle::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QCommonStyle::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QCommonStyle::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QCommonStyle::metaObject();
	}
	void drawPrimitive(QStyle::PrimitiveElement pe, const QStyleOption* opt, QPainter* p, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, 6098275085281943360LL, (int)(pe), (void*)opt, (void*)p, (void*)w, &__h); if (__h) return; }
		QCommonStyle::drawPrimitive(pe, opt, p, w);
	}
	void drawControl(QStyle::ControlElement element, const QStyleOption* opt, QPainter* p, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, -8697913287309583788LL, (int)(element), (void*)opt, (void*)p, (void*)w, &__h); if (__h) return; }
		QCommonStyle::drawControl(element, opt, p, w);
	}
	QRect subElementRect(QStyle::SubElement r, const QStyleOption* opt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -1580195784239416471LL, (int)(r), (void*)opt, (void*)widget, &__h); if (__h) return *(QRect*)__r; }
		return QCommonStyle::subElementRect(r, opt, widget);
	}
	void drawComplexControl(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, QPainter* p, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, -778167327566211030LL, (int)(cc), (void*)opt, (void*)p, (void*)w, &__h); if (__h) return; }
		QCommonStyle::drawComplexControl(cc, opt, p, w);
	}
	QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, const QPoint& pt, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_int_int_voidp_voidp_voidp)(__goId, 7058885155526130727LL, (int)(cc), (void*)opt, (void*)&pt, (void*)w, &__h); if (__h) return (QStyle::SubControl)(__r); }
		return QCommonStyle::hitTestComplexControl(cc, opt, pt, w);
	}
	QRect subControlRect(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, QStyle::SubControl sc, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_int_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int, void*, int*))cppgo_virt_voidp_int_voidp_int_voidp)(__goId, -8244607819117461060LL, (int)(cc), (void*)opt, (int)(sc), (void*)w, &__h); if (__h) return *(QRect*)__r; }
		return QCommonStyle::subControlRect(cc, opt, sc, w);
	}
	QSize sizeFromContents(QStyle::ContentsType ct, const QStyleOption* opt, const QSize& contentsSize, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp_voidp)(__goId, -3316981301817172730LL, (int)(ct), (void*)opt, (void*)&contentsSize, (void*)widget, &__h); if (__h) return *(QSize*)__r; }
		return QCommonStyle::sizeFromContents(ct, opt, contentsSize, widget);
	}
	int pixelMetric(QStyle::PixelMetric m, const QStyleOption* opt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, int*))cppgo_virt_int_int_voidp_voidp)(__goId, -7621686327863150821LL, (int)(m), (void*)opt, (void*)widget, &__h); if (__h) return __r; }
		return QCommonStyle::pixelMetric(m, opt, widget);
	}
	int styleHint(QStyle::StyleHint sh, const QStyleOption* opt, const QWidget* w, QStyleHintReturn* shret) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_int_int_voidp_voidp_voidp)(__goId, 2017029407462380579LL, (int)(sh), (void*)opt, (void*)w, (void*)shret, &__h); if (__h) return __r; }
		return QCommonStyle::styleHint(sh, opt, w, shret);
	}
	QIcon standardIcon(QStyle::StandardPixmap standardIcon, const QStyleOption* opt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -1870022100809264807LL, (int)(standardIcon), (void*)opt, (void*)widget, &__h); if (__h) return *(QIcon*)__r; }
		return QCommonStyle::standardIcon(standardIcon, opt, widget);
	}
	QPixmap standardPixmap(QStyle::StandardPixmap sp, const QStyleOption* opt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -5288619860977043091LL, (int)(sp), (void*)opt, (void*)widget, &__h); if (__h) return *(QPixmap*)__r; }
		return QCommonStyle::standardPixmap(sp, opt, widget);
	}
	QPixmap generatedIconPixmap(QIcon::Mode iconMode, const QPixmap& pixmap, const QStyleOption* opt) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -4387834948099211678LL, (int)(iconMode), (void*)&pixmap, (void*)opt, &__h); if (__h) return *(QPixmap*)__r; }
		return QCommonStyle::generatedIconPixmap(iconMode, pixmap, opt);
	}
	int layoutSpacing(QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int_int_int_voidp_voidp) { int __r = ((int(*)(long long, long long, int, int, int, void*, void*, int*))cppgo_virt_int_int_int_int_voidp_voidp)(__goId, -4874522738816254324LL, (int)(control1), (int)(control2), (int)(orientation), (void*)option, (void*)widget, &__h); if (__h) return __r; }
		return QCommonStyle::layoutSpacing(control1, control2, orientation, option, widget);
	}
	void polish(QPalette& arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4633126903937721620LL, (void*)&arg1, &__h); if (__h) return; }
		QCommonStyle::polish(arg1);
	}
	void polish(QApplication* app) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4263633720130086714LL, (void*)app, &__h); if (__h) return; }
		QCommonStyle::polish(app);
	}
	void polish(QWidget* widget) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4263634819641714925LL, (void*)widget, &__h); if (__h) return; }
		QCommonStyle::polish(widget);
	}
	void unpolish(QWidget* widget) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5547720835653541585LL, (void*)widget, &__h); if (__h) return; }
		QCommonStyle::unpolish(widget);
	}
	void unpolish(QApplication* application) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 265826155645830215LL, (void*)application, &__h); if (__h) return; }
		QCommonStyle::unpolish(application);
	}
	QRect itemTextRect(const QFontMetrics& fm, const QRect& r, int flags, bool enabled, const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[19] && cppgo_virt_voidp_voidp_voidp_int_int_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int, const char*, long long, int*))cppgo_virt_voidp_voidp_voidp_int_int_constcharp_longlong)(__goId, 8974830244642632823LL, (void*)&fm, (void*)&r, flags, enabled, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return *(QRect*)__r; }
		return QCommonStyle::itemTextRect(fm, r, flags, enabled, text);
	}
	QRect itemPixmapRect(const QRect& r, int flags, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, void*, int, void*, int*))cppgo_virt_voidp_voidp_int_voidp)(__goId, 3847657773274941773LL, (void*)&r, flags, (void*)&pixmap, &__h); if (__h) return *(QRect*)__r; }
		return QCommonStyle::itemPixmapRect(r, flags, pixmap);
	}
	void drawItemText(QPainter* painter, const QRect& rect, int flags, const QPalette& pal, bool enabled, const QString& text, QPalette::ColorRole textRole) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[21] && cppgo_virt_void_voidp_voidp_int_voidp_int_constcharp_longlong_int) { ((void(*)(long long, long long, void*, void*, int, void*, int, const char*, long long, int, int*))cppgo_virt_void_voidp_voidp_int_voidp_int_constcharp_longlong_int)(__goId, 2315857538243927747LL, (void*)painter, (void*)&rect, flags, (void*)&pal, enabled, (char*)text_ba.constData(), text_ba.size(), (int)(textRole), &__h); if (__h) return; }
		QCommonStyle::drawItemText(painter, rect, flags, pal, enabled, text, textRole);
	}
	void drawItemPixmap(QPainter* painter, const QRect& rect, int alignment, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp_voidp_int_voidp) { ((void(*)(long long, long long, void*, void*, int, void*, int*))cppgo_virt_void_voidp_voidp_int_voidp)(__goId, -3740415202673381007LL, (void*)painter, (void*)&rect, alignment, (void*)&pixmap, &__h); if (__h) return; }
		QCommonStyle::drawItemPixmap(painter, rect, alignment, pixmap);
	}
	QPalette standardPalette() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7278964765692882119LL, &__h); if (__h) return *(QPalette*)__r; }
		return QCommonStyle::standardPalette();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QCommonStyle::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QCommonStyle::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QCommonStyle::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QCommonStyle::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QCommonStyle::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QCommonStyle::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QCommonStyle::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QCommonStyle;
static std::mutex __connsMu_QCommonStyle;

extern "C" {

void* QCommonStyle_base(void* self) { return (void*)static_cast<QStyle*>((QCommonStyle*)self); }
void QCommonStyle_chain(void* self, void** out) {
	QCommonStyle* __p = (QCommonStyle*)self;
	QStyle* __b0 = static_cast<QStyle*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QCommonStyle__set_override(void* self, int idx) { ((QCommonStyle__Go*)self)->__goOn[idx] = 1; }
void* QCommonStyle_fromQObject(void* o) { return (void*)qobject_cast<QCommonStyle*>((QObject*)o); }
void* QCommonStyle_new(long long __goid) {
	return new QCommonStyle__Go(__goid);
}
char* QCommonStyle_tr(const char* s, long long* __out_len) {
	auto&& __r = QCommonStyle::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCommonStyle_visualRect(int direction, void* boundingRect, void* logicalRect) {
	auto&& __r = QCommonStyle::visualRect((Qt::LayoutDirection)(direction), *(QRect*)boundingRect, *(QRect*)logicalRect);
	return new QRect(__r);
}
void* QCommonStyle_visualPos(int direction, void* boundingRect, void* logicalPos) {
	auto&& __r = QCommonStyle::visualPos((Qt::LayoutDirection)(direction), *(QRect*)boundingRect, *(QPoint*)logicalPos);
	return new QPoint(__r);
}
int QCommonStyle_sliderPositionFromValue(int min, int max, int val, int space) {
	auto&& __r = QCommonStyle::sliderPositionFromValue(min, max, val, space);
	return __r;
}
int QCommonStyle_sliderValueFromPosition(int min, int max, int pos, int space) {
	auto&& __r = QCommonStyle::sliderValueFromPosition(min, max, pos, space);
	return __r;
}
int QCommonStyle_visualAlignment(int direction, int alignment) {
	auto&& __r = QCommonStyle::visualAlignment((Qt::LayoutDirection)(direction), QFlags<Qt::AlignmentFlag>(alignment));
	return (int)(__r);
}
void* QCommonStyle_alignedRect(int direction, int alignment, void* size, void* rectangle) {
	auto&& __r = QCommonStyle::alignedRect((Qt::LayoutDirection)(direction), QFlags<Qt::AlignmentFlag>(alignment), *(QSize*)size, *(QRect*)rectangle);
	return new QRect(__r);
}
void* QCommonStyle_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QCommonStyle::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QCommonStyle_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QCommonStyle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QCommonStyle_sender(void* self) {
	auto&& __r = ((QCommonStyle__Go*)self)->sender();
	return (void*)(__r);
}
int QCommonStyle_senderSignalIndex(void* self) {
	auto&& __r = ((QCommonStyle__Go*)self)->senderSignalIndex();
	return __r;
}
int QCommonStyle_receivers(void* self, const char* signal) {
	auto&& __r = ((QCommonStyle__Go*)self)->receivers(signal);
	return __r;
}
int QCommonStyle_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QCommonStyle__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QCommonStyle_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QCommonStyle::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCommonStyle_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QCommonStyle::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCommonStyle_sliderPositionFromValue2(int min, int max, int val, int space, int upsideDown) {
	auto&& __r = QCommonStyle::sliderPositionFromValue(min, max, val, space, upsideDown);
	return __r;
}
int QCommonStyle_sliderValueFromPosition2(int min, int max, int pos, int space, int upsideDown) {
	auto&& __r = QCommonStyle::sliderValueFromPosition(min, max, pos, space, upsideDown);
	return __r;
}
void* QCommonStyle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QCommonStyle::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QCommonStyle_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QCommonStyle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QCommonStyle_metaObject_default(void* self) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::metaObject();
	return (void*)(__r);
}
void QCommonStyle_drawPrimitive_default(void* self, int pe, void* opt, void* p, void* w) {
	((QCommonStyle*)self)->QCommonStyle::drawPrimitive((QStyle::PrimitiveElement)(pe), (QStyleOption*)opt, (QPainter*)p, (QWidget*)w);
}
void QCommonStyle_drawControl_default(void* self, int element, void* opt, void* p, void* w) {
	((QCommonStyle*)self)->QCommonStyle::drawControl((QStyle::ControlElement)(element), (QStyleOption*)opt, (QPainter*)p, (QWidget*)w);
}
void* QCommonStyle_subElementRect_default(void* self, int r, void* opt, void* widget) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::subElementRect((QStyle::SubElement)(r), (QStyleOption*)opt, (QWidget*)widget);
	return new QRect(__r);
}
void QCommonStyle_drawComplexControl_default(void* self, int cc, void* opt, void* p, void* w) {
	((QCommonStyle*)self)->QCommonStyle::drawComplexControl((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, (QPainter*)p, (QWidget*)w);
}
int QCommonStyle_hitTestComplexControl_default(void* self, int cc, void* opt, void* pt, void* w) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::hitTestComplexControl((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, *(QPoint*)pt, (QWidget*)w);
	return (int)(__r);
}
void* QCommonStyle_subControlRect_default(void* self, int cc, void* opt, int sc, void* w) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::subControlRect((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, (QStyle::SubControl)(sc), (QWidget*)w);
	return new QRect(__r);
}
void* QCommonStyle_sizeFromContents_default(void* self, int ct, void* opt, void* contentsSize, void* widget) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::sizeFromContents((QStyle::ContentsType)(ct), (QStyleOption*)opt, *(QSize*)contentsSize, (QWidget*)widget);
	return new QSize(__r);
}
int QCommonStyle_pixelMetric_default(void* self, int m, void* opt, void* widget) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::pixelMetric((QStyle::PixelMetric)(m), (QStyleOption*)opt, (QWidget*)widget);
	return __r;
}
int QCommonStyle_styleHint_default(void* self, int sh, void* opt, void* w, void* shret) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::styleHint((QStyle::StyleHint)(sh), (QStyleOption*)opt, (QWidget*)w, (QStyleHintReturn*)shret);
	return __r;
}
void* QCommonStyle_standardIcon_default(void* self, int standardIcon, void* opt, void* widget) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::standardIcon((QStyle::StandardPixmap)(standardIcon), (QStyleOption*)opt, (QWidget*)widget);
	return new QIcon(__r);
}
void* QCommonStyle_standardPixmap_default(void* self, int sp, void* opt, void* widget) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::standardPixmap((QStyle::StandardPixmap)(sp), (QStyleOption*)opt, (QWidget*)widget);
	return new QPixmap(__r);
}
void* QCommonStyle_generatedIconPixmap_default(void* self, int iconMode, void* pixmap, void* opt) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::generatedIconPixmap((QIcon::Mode)(iconMode), *(QPixmap*)pixmap, (QStyleOption*)opt);
	return new QPixmap(__r);
}
int QCommonStyle_layoutSpacing_default(void* self, int control1, int control2, int orientation, void* option, void* widget) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::layoutSpacing((QSizePolicy::ControlType)(control1), (QSizePolicy::ControlType)(control2), (Qt::Orientation)(orientation), (QStyleOption*)option, (QWidget*)widget);
	return __r;
}
void QCommonStyle_polish_default(void* self, void* arg1) {
	((QCommonStyle*)self)->QCommonStyle::polish(*(QPalette*)arg1);
}
void QCommonStyle_unpolish_default(void* self, void* widget) {
	((QCommonStyle*)self)->QCommonStyle::unpolish((QWidget*)widget);
}
void* QCommonStyle_itemTextRect_default(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::itemTextRect(*(QFontMetrics*)fm, *(QRect*)r, flags, enabled, text_qstr);
	return new QRect(__r);
}
void* QCommonStyle_itemPixmapRect_default(void* self, void* r, int flags, void* pixmap) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::itemPixmapRect(*(QRect*)r, flags, *(QPixmap*)pixmap);
	return new QRect(__r);
}
void QCommonStyle_drawItemText_default(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QCommonStyle*)self)->QCommonStyle::drawItemText((QPainter*)painter, *(QRect*)rect, flags, *(QPalette*)pal, enabled, text_qstr, (QPalette::ColorRole)(textRole));
}
void QCommonStyle_drawItemPixmap_default(void* self, void* painter, void* rect, int alignment, void* pixmap) {
	((QCommonStyle*)self)->QCommonStyle::drawItemPixmap((QPainter*)painter, *(QRect*)rect, alignment, *(QPixmap*)pixmap);
}
void* QCommonStyle_standardPalette_default(void* self) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::standardPalette();
	return new QPalette(__r);
}
int QCommonStyle_event_default(void* self, void* event) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::event((QEvent*)event);
	return __r;
}
int QCommonStyle_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QCommonStyle*)self)->QCommonStyle::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QCommonStyle_timerEvent_default(void* self, void* event) {
	((QCommonStyle__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QCommonStyle_childEvent_default(void* self, void* event) {
	((QCommonStyle__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QCommonStyle_customEvent_default(void* self, void* event) {
	((QCommonStyle__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QCommonStyle_connectNotify_default(void* self, void* signal) {
	((QCommonStyle__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QCommonStyle_disconnectNotify_default(void* self, void* signal) {
	((QCommonStyle__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QCommonStyle_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommonStyle*)self, &QCommonStyle::destroyed, (QCommonStyle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommonStyle);
	__conns_QCommonStyle[__id] = __c;
}
void QCommonStyle_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommonStyle*)self, &QCommonStyle::destroyed, (QCommonStyle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommonStyle);
	__conns_QCommonStyle[__id] = __c;
}
void QCommonStyle_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCommonStyle*)self, &QCommonStyle::objectNameChanged, (QCommonStyle*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCommonStyle);
	__conns_QCommonStyle[__id] = __c;
}
void QCommonStyle__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QCommonStyle);
		auto __it = __conns_QCommonStyle.find(id);
		if (__it == __conns_QCommonStyle.end()) return;
		__c = __it->second;
		__conns_QCommonStyle.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QCommonStyle__Go_new(long long id) { return new QCommonStyle__Go(id); }
void QCommonStyle_delete(void* self) { delete (QCommonStyle*)self; }

}

#include <QtWidgets/qgesture.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qpoint.h>
#include "gen_qgestureevent.h"

extern "C" {

void* QGestureEvent_base(void* self) { return (void*)static_cast<QEvent*>((QGestureEvent*)self); }
void* QGestureEvent_new(void** gestures, long long gestures_len) {
	QList<QGesture*> gestures_list; gestures_list.reserve(gestures_len);
	for (long long __i = 0; __i < gestures_len; __i++) gestures_list.append((QGesture*)gestures[__i]);
	return new QGestureEvent(gestures_list);
}
void* QGestureEvent_new2(void* arg1) {
	return new QGestureEvent(*(QGestureEvent*)arg1);
}
void** QGestureEvent_gestures(void* self, long long* __out_len) {
	auto&& __r = ((QGestureEvent*)self)->gestures();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QGestureEvent_gesture(void* self, int type_) {
	auto&& __r = ((QGestureEvent*)self)->gesture((Qt::GestureType)(type_));
	return (void*)(__r);
}
void** QGestureEvent_activeGestures(void* self, long long* __out_len) {
	auto&& __r = ((QGestureEvent*)self)->activeGestures();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGestureEvent_canceledGestures(void* self, long long* __out_len) {
	auto&& __r = ((QGestureEvent*)self)->canceledGestures();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGestureEvent_setAccepted(void* self, void* arg1, int arg2) {
	((QGestureEvent*)self)->setAccepted((QGesture*)arg1, arg2);
}
void QGestureEvent_accept(void* self, void* arg1) {
	((QGestureEvent*)self)->accept((QGesture*)arg1);
}
void QGestureEvent_ignore(void* self, void* arg1) {
	((QGestureEvent*)self)->ignore((QGesture*)arg1);
}
int QGestureEvent_isAccepted(void* self, void* arg1) {
	auto&& __r = ((QGestureEvent*)self)->isAccepted((QGesture*)arg1);
	return __r;
}
void QGestureEvent_setAccepted2(void* self, int arg1, int arg2) {
	((QGestureEvent*)self)->setAccepted((Qt::GestureType)(arg1), arg2);
}
void QGestureEvent_accept2(void* self, int arg1) {
	((QGestureEvent*)self)->accept((Qt::GestureType)(arg1));
}
void QGestureEvent_ignore2(void* self, int arg1) {
	((QGestureEvent*)self)->ignore((Qt::GestureType)(arg1));
}
int QGestureEvent_isAccepted2(void* self, int arg1) {
	auto&& __r = ((QGestureEvent*)self)->isAccepted((Qt::GestureType)(arg1));
	return __r;
}
void QGestureEvent_setWidget(void* self, void* widget) {
	((QGestureEvent*)self)->setWidget((QWidget*)widget);
}
void* QGestureEvent_widget(void* self) {
	auto&& __r = ((QGestureEvent*)self)->widget();
	return (void*)(__r);
}
void* QGestureEvent_mapToGraphicsScene(void* self, void* gesturePoint) {
	auto&& __r = ((QGestureEvent*)self)->mapToGraphicsScene(*(QPointF*)gesturePoint);
	return new QPointF(__r);
}
void* QGestureEvent_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QGestureEvent*)self)->operator=(*(QGestureEvent*)arg1);
	return (void*)&(__r);
}
int QGestureEvent_registerEventType() {
	auto&& __r = QGestureEvent::registerEventType();
	return __r;
}
int QGestureEvent_registerEventType2(int hint) {
	auto&& __r = QGestureEvent::registerEventType(hint);
	return __r;
}
void QGestureEvent_delete(void* self) { delete (QGestureEvent*)self; }

}

#include <QtWidgets/qgraphicsitemanimation.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qtimeline.h>
#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsitemanimation.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsItemAnimation__Go : public QGraphicsItemAnimation {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QGraphicsItemAnimation__Go(long long id, A&&... a) : QGraphicsItemAnimation(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsItemAnimation::sender;
	using QGraphicsItemAnimation::senderSignalIndex;
	using QGraphicsItemAnimation::receivers;
	using QGraphicsItemAnimation::isSignalConnected;
	void cppgoBase_beforeAnimationStep(qreal __a0) { QGraphicsItemAnimation::beforeAnimationStep(__a0); }
	void cppgoBase_afterAnimationStep(qreal __a0) { QGraphicsItemAnimation::afterAnimationStep(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsItemAnimation::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsItemAnimation::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsItemAnimation::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsItemAnimation::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsItemAnimation::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsItemAnimation::metaObject();
	}
	void beforeAnimationStep(qreal step) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_unsignedlonglong) { ((void(*)(long long, long long, unsigned long long, int*))cppgo_virt_void_unsignedlonglong)(__goId, -1838142790170279106LL, __builtin_bit_cast(unsigned long long, (double)(step)), &__h); if (__h) return; }
		QGraphicsItemAnimation::beforeAnimationStep(step);
	}
	void afterAnimationStep(qreal step) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_unsignedlonglong) { ((void(*)(long long, long long, unsigned long long, int*))cppgo_virt_void_unsignedlonglong)(__goId, 7260939011549207179LL, __builtin_bit_cast(unsigned long long, (double)(step)), &__h); if (__h) return; }
		QGraphicsItemAnimation::afterAnimationStep(step);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsItemAnimation::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsItemAnimation::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemAnimation::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemAnimation::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemAnimation::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsItemAnimation::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsItemAnimation::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsItemAnimation;
static std::mutex __connsMu_QGraphicsItemAnimation;

extern "C" {

void* QGraphicsItemAnimation_base(void* self) { return (void*)static_cast<QObject*>((QGraphicsItemAnimation*)self); }
void QGraphicsItemAnimation__set_override(void* self, int idx) { ((QGraphicsItemAnimation__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsItemAnimation_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsItemAnimation*>((QObject*)o); }
void* QGraphicsItemAnimation_new(long long __goid) {
	return new QGraphicsItemAnimation__Go(__goid);
}
void* QGraphicsItemAnimation_new2(long long __goid, void* parent) {
	return new QGraphicsItemAnimation__Go(__goid, (QObject*)parent);
}
char* QGraphicsItemAnimation_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsItemAnimation::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsItemAnimation_item(void* self) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->item();
	return (void*)(__r);
}
void QGraphicsItemAnimation_setItem(void* self, void* item) {
	((QGraphicsItemAnimation*)self)->setItem((QGraphicsItem*)item);
}
void* QGraphicsItemAnimation_timeLine(void* self) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->timeLine();
	return (void*)(__r);
}
void QGraphicsItemAnimation_setTimeLine(void* self, void* timeLine) {
	((QGraphicsItemAnimation*)self)->setTimeLine((QTimeLine*)timeLine);
}
void* QGraphicsItemAnimation_posAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->posAt(step);
	return new QPointF(__r);
}
void QGraphicsItemAnimation_setPosAt(void* self, double step, void* pos) {
	((QGraphicsItemAnimation*)self)->setPosAt(step, *(QPointF*)pos);
}
void* QGraphicsItemAnimation_transformAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->transformAt(step);
	return new QTransform(__r);
}
double QGraphicsItemAnimation_rotationAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->rotationAt(step);
	return __r;
}
void QGraphicsItemAnimation_setRotationAt(void* self, double step, double angle) {
	((QGraphicsItemAnimation*)self)->setRotationAt(step, angle);
}
double QGraphicsItemAnimation_xTranslationAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->xTranslationAt(step);
	return __r;
}
double QGraphicsItemAnimation_yTranslationAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->yTranslationAt(step);
	return __r;
}
void QGraphicsItemAnimation_setTranslationAt(void* self, double step, double dx, double dy) {
	((QGraphicsItemAnimation*)self)->setTranslationAt(step, dx, dy);
}
double QGraphicsItemAnimation_verticalScaleAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->verticalScaleAt(step);
	return __r;
}
double QGraphicsItemAnimation_horizontalScaleAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->horizontalScaleAt(step);
	return __r;
}
void QGraphicsItemAnimation_setScaleAt(void* self, double step, double sx, double sy) {
	((QGraphicsItemAnimation*)self)->setScaleAt(step, sx, sy);
}
double QGraphicsItemAnimation_verticalShearAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->verticalShearAt(step);
	return __r;
}
double QGraphicsItemAnimation_horizontalShearAt(void* self, double step) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->horizontalShearAt(step);
	return __r;
}
void QGraphicsItemAnimation_setShearAt(void* self, double step, double sh, double sv) {
	((QGraphicsItemAnimation*)self)->setShearAt(step, sh, sv);
}
void QGraphicsItemAnimation_clear(void* self) {
	((QGraphicsItemAnimation*)self)->clear();
}
void QGraphicsItemAnimation_setStep(void* self, double x) {
	((QGraphicsItemAnimation*)self)->setStep(x);
}
void* QGraphicsItemAnimation_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsItemAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsItemAnimation_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsItemAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsItemAnimation_sender(void* self) {
	auto&& __r = ((QGraphicsItemAnimation__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsItemAnimation_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsItemAnimation__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsItemAnimation_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsItemAnimation__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsItemAnimation_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsItemAnimation__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsItemAnimation_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsItemAnimation::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsItemAnimation_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsItemAnimation::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsItemAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsItemAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsItemAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsItemAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsItemAnimation_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->QGraphicsItemAnimation::metaObject();
	return (void*)(__r);
}
void QGraphicsItemAnimation_beforeAnimationStep_default(void* self, double step) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_beforeAnimationStep(step);
}
void QGraphicsItemAnimation_afterAnimationStep_default(void* self, double step) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_afterAnimationStep(step);
}
int QGraphicsItemAnimation_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->QGraphicsItemAnimation::event((QEvent*)event);
	return __r;
}
int QGraphicsItemAnimation_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsItemAnimation*)self)->QGraphicsItemAnimation::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsItemAnimation_timerEvent_default(void* self, void* event) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsItemAnimation_childEvent_default(void* self, void* event) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsItemAnimation_customEvent_default(void* self, void* event) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsItemAnimation_connectNotify_default(void* self, void* signal) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsItemAnimation_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsItemAnimation__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsItemAnimation_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsItemAnimation*)self, &QGraphicsItemAnimation::destroyed, (QGraphicsItemAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsItemAnimation);
	__conns_QGraphicsItemAnimation[__id] = __c;
}
void QGraphicsItemAnimation_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsItemAnimation*)self, &QGraphicsItemAnimation::destroyed, (QGraphicsItemAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsItemAnimation);
	__conns_QGraphicsItemAnimation[__id] = __c;
}
void QGraphicsItemAnimation_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsItemAnimation*)self, &QGraphicsItemAnimation::objectNameChanged, (QGraphicsItemAnimation*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsItemAnimation);
	__conns_QGraphicsItemAnimation[__id] = __c;
}
void QGraphicsItemAnimation__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsItemAnimation);
		auto __it = __conns_QGraphicsItemAnimation.find(id);
		if (__it == __conns_QGraphicsItemAnimation.end()) return;
		__c = __it->second;
		__conns_QGraphicsItemAnimation.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsItemAnimation__Go_new(long long id, void* parent) { return new QGraphicsItemAnimation__Go(id, (QObject*)parent); }
void QGraphicsItemAnimation_delete(void* self) { delete (QGraphicsItemAnimation*)self; }

}

#include <QtWidgets/qitemeditorfactory.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qbytearray.h>
#include "gen_qitemeditorcreatorbase.h"
#include "gen_purego_cb.h"

class QItemEditorCreatorBase__Go : public QItemEditorCreatorBase {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QItemEditorCreatorBase__Go(long long id, A&&... a) : QItemEditorCreatorBase(std::forward<A>(a)...), __goId(id) {}
	QWidget* createWidget(QWidget* parent) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8470300370605365395LL, (void*)parent, &__h); if (__h) return (QWidget*)__r; }
		return nullptr;
	}
	QByteArray valuePropertyName() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -3548500303342436100LL, &__h); if (__h) return *(QByteArray*)__r; }
		return QByteArray{};
	}
};

extern "C" {

void QItemEditorCreatorBase__set_override(void* self, int idx) { ((QItemEditorCreatorBase__Go*)self)->__goOn[idx] = 1; }
void* QItemEditorCreatorBase_createWidget(void* self, void* parent) {
	auto&& __r = ((QItemEditorCreatorBase*)self)->createWidget((QWidget*)parent);
	return (void*)(__r);
}
void* QItemEditorCreatorBase_valuePropertyName(void* self) {
	auto&& __r = ((QItemEditorCreatorBase*)self)->valuePropertyName();
	return new QByteArray(__r);
}
void* QItemEditorCreatorBase_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QItemEditorCreatorBase*)self)->operator=(*(QItemEditorCreatorBase*)arg1);
	return (void*)&(__r);
}
void* QItemEditorCreatorBase__Go_new(long long id) { return new QItemEditorCreatorBase__Go(id); }
void QItemEditorCreatorBase_delete(void* self) { delete (QItemEditorCreatorBase*)self; }

}

#include <QtWidgets/qscrollerproperties.h>
#include <QtCore/qvariant.h>
#include "gen_qscrollerproperties.h"

extern "C" {

void* QScrollerProperties_new() {
	return new QScrollerProperties();
}
void* QScrollerProperties_new2(void* sp) {
	return new QScrollerProperties(*(QScrollerProperties*)sp);
}
void* QScrollerProperties_OperatorAssign(void* self, void* sp) {
	auto&& __r = ((QScrollerProperties*)self)->operator=(*(QScrollerProperties*)sp);
	return (void*)&(__r);
}
int QScrollerProperties_OperatorEqual(void* self, void* sp) {
	auto&& __r = ((QScrollerProperties*)self)->operator==(*(QScrollerProperties*)sp);
	return __r;
}
int QScrollerProperties_OperatorNotEqual(void* self, void* sp) {
	auto&& __r = ((QScrollerProperties*)self)->operator!=(*(QScrollerProperties*)sp);
	return __r;
}
void QScrollerProperties_setDefaultScrollerProperties(void* sp) {
	QScrollerProperties::setDefaultScrollerProperties(*(QScrollerProperties*)sp);
}
void QScrollerProperties_unsetDefaultScrollerProperties() {
	QScrollerProperties::unsetDefaultScrollerProperties();
}
void* QScrollerProperties_scrollMetric(void* self, int metric) {
	auto&& __r = ((QScrollerProperties*)self)->scrollMetric((QScrollerProperties::ScrollMetric)(metric));
	return new QVariant(__r);
}
void QScrollerProperties_setScrollMetric(void* self, int metric, void* value) {
	((QScrollerProperties*)self)->setScrollMetric((QScrollerProperties::ScrollMetric)(metric), *(QVariant*)value);
}
void QScrollerProperties_delete(void* self) { delete (QScrollerProperties*)self; }

}

#include <QtWidgets/qtooltip.h>
#include <QtCore/qpoint.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include "gen_qtooltip.h"

extern "C" {

void QToolTip_showText(void* pos, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QToolTip::showText(*(QPoint*)pos, text_qstr);
}
void QToolTip_hideText() {
	QToolTip::hideText();
}
int QToolTip_isVisible() {
	auto&& __r = QToolTip::isVisible();
	return __r;
}
char* QToolTip_text(long long* __out_len) {
	auto&& __r = QToolTip::text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QToolTip_palette() {
	auto&& __r = QToolTip::palette();
	return new QPalette(__r);
}
void QToolTip_setPalette(void* arg1) {
	QToolTip::setPalette(*(QPalette*)arg1);
}
void* QToolTip_font() {
	auto&& __r = QToolTip::font();
	return new QFont(__r);
}
void QToolTip_setFont(void* arg1) {
	QToolTip::setFont(*(QFont*)arg1);
}
void QToolTip_showText2(void* pos, const char* text, long long text_len, void* w) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QToolTip::showText(*(QPoint*)pos, text_qstr, (QWidget*)w);
}
void QToolTip_showText3(void* pos, const char* text, long long text_len, void* w, void* rect) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QToolTip::showText(*(QPoint*)pos, text_qstr, (QWidget*)w, *(QRect*)rect);
}
void QToolTip_showText4(void* pos, const char* text, long long text_len, void* w, void* rect, int msecShowTime) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QToolTip::showText(*(QPoint*)pos, text_qstr, (QWidget*)w, *(QRect*)rect, msecShowTime);
}
void QToolTip_delete(void* self) { delete (QToolTip*)self; }

}

