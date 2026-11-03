#include <QtCore/qobjectdefs.h>
#include "gen_qmetamethodreturnargument.h"

extern "C" {

const char* QMetaMethodReturnArgument_name_get(void* self) {
	auto __r = ((QMetaMethodReturnArgument*)self)->name;
	return __r;
}
void QMetaMethodReturnArgument_name_set(void* self, const char* name) {
	((QMetaMethodReturnArgument*)self)->name = name;
}
void QMetaMethodReturnArgument_delete(void* self) { delete (QMetaMethodReturnArgument*)self; }

}

#include <QtCore/qdatastream.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qfloat16.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbitarray.h>
#include <QtCore/qlocale.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qurl.h>
#include <QtCore/quuid.h>
#include <QtCore/qcborvalue.h>
#include <QtCore/qcborarray.h>
#include <QtCore/qcbormap.h>
#include <QtCore/qtimezone.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qjsondocument.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qtyperevision.h>
#include <QtCore/qversionnumber.h>
#include <QtCore/qpoint.h>
#include <QtCore/qline.h>
#include <QtCore/qmargins.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpolygon.h>
#include <QtGui/qregion.h>
#include <QtGui/qtransform.h>
#include <QtGui/qimage.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qicon.h>
#include <QtGui/qvectornd.h>
#include <QtGui/qkeysequence.h>
#include <QtGui/qfont.h>
#include <QtGui/qbrush.h>
#include <QtGui/qpen.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qpalette.h>
#include <QtGui/qcursor.h>
#include <QtGui/qcolorspace.h>
#include <QtGui/qmatrix4x4.h>
#include <QtGui/qpageranges.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qpicture.h>
#include <QtGui/qquaternion.h>
#include <QtGui/qstandarditemmodel.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlistwidget.h>
#include <QtWidgets/qtablewidget.h>
#include <QtWidgets/qtreewidget.h>
#include "gen_qdatastream.h"

extern "C" {

void* QDataStream_base(void* self) { return (void*)static_cast<QIODeviceBase*>((QDataStream*)self); }
void* QDataStream_new() {
	return new QDataStream();
}
void* QDataStream_new2(void* arg1) {
	return new QDataStream((QIODevice*)arg1);
}
void* QDataStream_new3(void* arg1, int flags) {
	return new QDataStream((QByteArray*)arg1, QFlags<QIODeviceBase::OpenModeFlag>(flags));
}
void* QDataStream_new4(void* arg1) {
	return new QDataStream(*(QByteArray*)arg1);
}
void* QDataStream_device(void* self) {
	auto&& __r = ((QDataStream*)self)->device();
	return (void*)(__r);
}
void QDataStream_setDevice(void* self, void* arg1) {
	((QDataStream*)self)->setDevice((QIODevice*)arg1);
}
int QDataStream_atEnd(void* self) {
	auto&& __r = ((QDataStream*)self)->atEnd();
	return __r;
}
int QDataStream_status(void* self) {
	auto&& __r = ((QDataStream*)self)->status();
	return (int)(__r);
}
void QDataStream_setStatus(void* self, int status) {
	((QDataStream*)self)->setStatus((QDataStream::Status)(status));
}
void QDataStream_resetStatus(void* self) {
	((QDataStream*)self)->resetStatus();
}
int QDataStream_floatingPointPrecision(void* self) {
	auto&& __r = ((QDataStream*)self)->floatingPointPrecision();
	return (int)(__r);
}
void QDataStream_setFloatingPointPrecision(void* self, int precision) {
	((QDataStream*)self)->setFloatingPointPrecision((QDataStream::FloatingPointPrecision)(precision));
}
int QDataStream_byteOrder(void* self) {
	auto&& __r = ((QDataStream*)self)->byteOrder();
	return (int)(__r);
}
void QDataStream_setByteOrder(void* self, int arg1) {
	((QDataStream*)self)->setByteOrder((QDataStream::ByteOrder)(arg1));
}
int QDataStream_version(void* self) {
	auto&& __r = ((QDataStream*)self)->version();
	return __r;
}
void QDataStream_setVersion(void* self, int arg1) {
	((QDataStream*)self)->setVersion(arg1);
}
void* QDataStream_OperatorShiftRight(void* self, char* i) {
	char i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight2(void* self, char* i) {
	qint8 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight3(void* self, unsigned char* i) {
	quint8 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight4(void* self, short* i) {
	qint16 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight5(void* self, unsigned short* i) {
	quint16 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight6(void* self, int* i) {
	qint32 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight7(void* self, unsigned int* i) {
	quint32 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight8(void* self, long long* i) {
	qint64 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight9(void* self, unsigned long long* i) {
	quint64 i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight10(void* self, int* i) {
	bool i_ref = *i;
	auto&& __r = ((QDataStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight11(void* self, float* f) {
	float f_ref = *f;
	auto&& __r = ((QDataStream*)self)->operator>>(f_ref);
	*f = f_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight12(void* self, double* f) {
	double f_ref = *f;
	auto&& __r = ((QDataStream*)self)->operator>>(f_ref);
	*f = f_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight13(void* self, unsigned short* c) {
	char16_t c_ref = *c;
	auto&& __r = ((QDataStream*)self)->operator>>(c_ref);
	*c = c_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight14(void* self, unsigned int* c) {
	char32_t c_ref = *c;
	auto&& __r = ((QDataStream*)self)->operator>>(c_ref);
	*c = c_ref;
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft(void* self, char i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft2(void* self, char i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft3(void* self, unsigned char i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft4(void* self, short i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft5(void* self, unsigned short i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft6(void* self, int i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft7(void* self, unsigned int i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft8(void* self, long long i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft9(void* self, unsigned long long i) {
	auto&& __r = ((QDataStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft10(void* self, float f) {
	auto&& __r = ((QDataStream*)self)->operator<<(f);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft11(void* self, double f) {
	auto&& __r = ((QDataStream*)self)->operator<<(f);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft12(void* self, const char* str) {
	auto&& __r = ((QDataStream*)self)->operator<<(str);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft13(void* self, unsigned short c) {
	auto&& __r = ((QDataStream*)self)->operator<<(c);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft14(void* self, unsigned int c) {
	auto&& __r = ((QDataStream*)self)->operator<<(c);
	return (void*)&(__r);
}
int QDataStream_ToBool(void* self) {
	auto&& __r = static_cast<bool>(*(QDataStream*)self);
	return __r;
}
long long QDataStream_readRawData(void* self, char* arg1, long long len_) {
	auto&& __r = ((QDataStream*)self)->readRawData((char*)arg1, len_);
	return __r;
}
void* QDataStream_writeBytes(void* self, const char* arg1, long long len_) {
	auto&& __r = ((QDataStream*)self)->writeBytes(arg1, len_);
	return (void*)&(__r);
}
long long QDataStream_writeRawData(void* self, const char* arg1, long long len_) {
	auto&& __r = ((QDataStream*)self)->writeRawData(arg1, len_);
	return __r;
}
long long QDataStream_skipRawData(void* self, long long len_) {
	auto&& __r = ((QDataStream*)self)->skipRawData(len_);
	return __r;
}
void QDataStream_startTransaction(void* self) {
	((QDataStream*)self)->startTransaction();
}
int QDataStream_commitTransaction(void* self) {
	auto&& __r = ((QDataStream*)self)->commitTransaction();
	return __r;
}
void QDataStream_rollbackTransaction(void* self) {
	((QDataStream*)self)->rollbackTransaction();
}
void QDataStream_abortTransaction(void* self) {
	((QDataStream*)self)->abortTransaction();
}
int QDataStream_isDeviceTransactionStarted(void* self) {
	auto&& __r = ((QDataStream*)self)->isDeviceTransactionStarted();
	return __r;
}
void* QDataStream_OperatorShiftLeft15(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QByteArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight15(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QByteArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft16(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = (*(QDataStream*)self) << (arg1_qstr);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight16(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QString*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft17(void* self, void* chr) {
	auto&& __r = (*(QDataStream*)self) << (*(QChar*)chr);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight17(void* self, void* chr) {
	auto&& __r = (*(QDataStream*)self) >> (*(QChar*)chr);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight18(void* self, void* combination) {
	auto&& __r = (*(QDataStream*)self) >> (*(QKeyCombination*)combination);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft18(void* self, void* combination) {
	auto&& __r = (*(QDataStream*)self) << (*(QKeyCombination*)combination);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft19(void* self, void* f) {
	auto&& __r = (*(QDataStream*)self) << (*(qfloat16*)f);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight19(void* self, void* f) {
	auto&& __r = (*(QDataStream*)self) >> (*(qfloat16*)f);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight20(void* self, void* p) {
	auto&& __r = (*(QDataStream*)self) >> (*(QVariant*)p);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft20(void* self, void* p) {
	auto&& __r = (*(QDataStream*)self) << (*(QVariant*)p);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight21(void* self, int* p) {
	QVariant::Type p_ref = (QVariant::Type)(*p);
	auto&& __r = (*(QDataStream*)self) >> (p_ref);
	*p = (int)(p_ref);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft21(void* self, int p) {
	auto&& __r = (*(QDataStream*)self) << ((QVariant::Type)(p));
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft22(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QBitArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight22(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QBitArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft23(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QLocale*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight23(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QLocale*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft24(void* self, int st) {
	auto&& __r = (*(QDataStream*)self) << ((QCborSimpleType)(st));
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight24(void* self, int* st) {
	QCborSimpleType st_ref = (QCborSimpleType)(*st);
	auto&& __r = (*(QDataStream*)self) >> (st_ref);
	*st = (int)(st_ref);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft25(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QDate*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight25(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QDate*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft26(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTime*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight26(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTime*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft27(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QDateTime*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight27(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QDateTime*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft28(void* self, void* re) {
	auto&& __r = (*(QDataStream*)self) << (*(QRegularExpression*)re);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight28(void* self, void* re) {
	auto&& __r = (*(QDataStream*)self) >> (*(QRegularExpression*)re);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft29(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QUrl*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight29(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QUrl*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft30(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QUuid*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight30(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QUuid*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft31(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QCborValue*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight31(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QCborValue*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft32(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QCborArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight32(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QCborArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft33(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QCborMap*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight33(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QCborMap*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft34(void* self, void* tz) {
	auto&& __r = (*(QDataStream*)self) << (*(QTimeZone*)tz);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight34(void* self, void* tz) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTimeZone*)tz);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft35(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QEasingCurve*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight35(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QEasingCurve*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft36(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QJsonDocument*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight36(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QJsonDocument*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft37(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QJsonValue*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight37(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QJsonValue*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft38(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QJsonArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight38(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QJsonArray*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft39(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QJsonObject*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight39(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QJsonObject*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft40(void* self, void* revision) {
	auto&& __r = (*(QDataStream*)self) << (*(QTypeRevision*)revision);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight40(void* self, void* revision) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTypeRevision*)revision);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft41(void* self, void* version) {
	auto&& __r = (*(QDataStream*)self) << (*(QVersionNumber*)version);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight41(void* self, void* version) {
	auto&& __r = (*(QDataStream*)self) >> (*(QVersionNumber*)version);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft42(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QPoint*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight42(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPoint*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft43(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QPointF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight43(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPointF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft44(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QLine*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight44(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QLine*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft45(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QLineF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight45(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QLineF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft46(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QMargins*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight46(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QMargins*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft47(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QMarginsF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight47(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QMarginsF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft48(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QSize*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight48(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QSize*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft49(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QSizeF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight49(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QSizeF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft50(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QRect*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight50(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QRect*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft51(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QRectF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight51(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QRectF*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft52(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QColor*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight52(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QColor*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft53(void* self, void* polygon) {
	auto&& __r = (*(QDataStream*)self) << (*(QPolygon*)polygon);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight53(void* self, void* polygon) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPolygon*)polygon);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft54(void* self, void* array) {
	auto&& __r = (*(QDataStream*)self) << (*(QPolygonF*)array);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight54(void* self, void* array) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPolygonF*)array);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft55(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QRegion*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight55(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QRegion*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft56(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTransform*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight56(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTransform*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft57(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QImage*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight57(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QImage*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight58(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPixmap*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft58(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QPixmap*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft59(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QIcon*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight59(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QIcon*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft60(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QVector2D*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight60(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QVector2D*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft61(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QVector3D*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight61(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QVector3D*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft62(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QVector4D*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight62(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QVector4D*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft63(void* self, void* ks) {
	auto&& __r = (*(QDataStream*)self) << (*(QKeySequence*)ks);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight63(void* self, void* ks) {
	auto&& __r = (*(QDataStream*)self) >> (*(QKeySequence*)ks);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft64(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QFont::Tag*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight64(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QFont::Tag*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft65(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QFont*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight65(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QFont*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft66(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QBrush*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight66(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QBrush*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft67(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QPen*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight67(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPen*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft68(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextLength*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight68(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextLength*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft69(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight69(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft70(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextCharFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight70(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextCharFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft71(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextBlockFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight71(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextBlockFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft72(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextListFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight72(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextListFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft73(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextFrameFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight73(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextFrameFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft74(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QTextTableCellFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight74(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTextTableCellFormat*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft75(void* self, void* p) {
	auto&& __r = (*(QDataStream*)self) << (*(QPalette*)p);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight75(void* self, void* p) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPalette*)p);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft76(void* self, void* cursor) {
	auto&& __r = (*(QDataStream*)self) << (*(QCursor*)cursor);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight76(void* self, void* cursor) {
	auto&& __r = (*(QDataStream*)self) >> (*(QCursor*)cursor);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft77(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QColorSpace*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight77(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QColorSpace*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft78(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QMatrix4x4*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight78(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QMatrix4x4*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft79(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QPageRanges*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight79(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPageRanges*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft80(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QPainterPath*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight80(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPainterPath*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft81(void* self, void* p) {
	auto&& __r = (*(QDataStream*)self) << (*(QPicture*)p);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight81(void* self, void* p) {
	auto&& __r = (*(QDataStream*)self) >> (*(QPicture*)p);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft82(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QQuaternion*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight82(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QQuaternion*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight83(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) >> (*(QStandardItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft83(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) << (*(QStandardItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft84(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) << (*(QSizePolicy*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight84(void* self, void* arg1) {
	auto&& __r = (*(QDataStream*)self) >> (*(QSizePolicy*)arg1);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft85(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) << (*(QListWidgetItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight85(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) >> (*(QListWidgetItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight86(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTableWidgetItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft86(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) << (*(QTableWidgetItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftLeft87(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) << (*(QTreeWidgetItem*)item);
	return (void*)&(__r);
}
void* QDataStream_OperatorShiftRight87(void* self, void* item) {
	auto&& __r = (*(QDataStream*)self) >> (*(QTreeWidgetItem*)item);
	return (void*)&(__r);
}
void QDataStream_delete(void* self) { delete (QDataStream*)self; }

}

#include <QtCore/qshareddata.h>
#include <QtCore/qatomic.h>
#include "gen_qshareddata.h"

extern "C" {

void* QSharedData_new() {
	return new QSharedData();
}
void* QSharedData_new2(void* arg1) {
	return new QSharedData(*(QSharedData*)arg1);
}
void* QSharedData_ref_get(void* self) {
	auto __r = ((QSharedData*)self)->ref;
	return new QAtomicInt(__r);
}
void QSharedData_ref_set(void* self, void* ref) {
	((QSharedData*)self)->ref = *(QAtomicInt*)ref;
}
void QSharedData_delete(void* self) { delete (QSharedData*)self; }

}

#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include "gen_qchildevent.h"

extern "C" {

void* QChildEvent_base(void* self) { return (void*)static_cast<QEvent*>((QChildEvent*)self); }
void* QChildEvent_new(int type_, void* child) {
	return new QChildEvent((QEvent::Type)(type_), (QObject*)child);
}
void* QChildEvent_child(void* self) {
	auto&& __r = ((QChildEvent*)self)->child();
	return (void*)(__r);
}
int QChildEvent_added(void* self) {
	auto&& __r = ((QChildEvent*)self)->added();
	return __r;
}
int QChildEvent_polished(void* self) {
	auto&& __r = ((QChildEvent*)self)->polished();
	return __r;
}
int QChildEvent_removed(void* self) {
	auto&& __r = ((QChildEvent*)self)->removed();
	return __r;
}
int QChildEvent_registerEventType() {
	auto&& __r = QChildEvent::registerEventType();
	return __r;
}
int QChildEvent_registerEventType2(int hint) {
	auto&& __r = QChildEvent::registerEventType(hint);
	return __r;
}
void QChildEvent_delete(void* self) { delete (QChildEvent*)self; }

}

#include <QtCore/quuid.h>
#include <QtCore/qsysinfo.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include "gen_quuid.h"

extern "C" {

void* QUuid_new() {
	return new QUuid();
}
void* QUuid_new2(unsigned int l, unsigned short w1, unsigned short w2, unsigned char b1, unsigned char b2, unsigned char b3, unsigned char b4, unsigned char b5, unsigned char b6, unsigned char b7, unsigned char b8) {
	return new QUuid(l, w1, w2, b1, b2, b3, b4, b5, b6, b7, b8);
}
void* QUuid_new3(void* id128) {
	return new QUuid(*(QUuid::Id128Bytes*)id128);
}
void* QUuid_new4(const char* string_, long long string__len) {
	QAnyStringView string__qstr = QAnyStringView(string_, string__len);
	return new QUuid(string__qstr);
}
void* QUuid_new5(void* arg1) {
	return new QUuid(*(QUuid*)arg1);
}
void* QUuid_new6(void* id128, int order) {
	return new QUuid(*(QUuid::Id128Bytes*)id128, (QSysInfo::Endian)(order));
}
void* QUuid_fromString(const char* string_, long long string__len) {
	QAnyStringView string__qstr = QAnyStringView(string_, string__len);
	auto&& __r = QUuid::fromString(string__qstr);
	return new QUuid(__r);
}
char* QUuid_toString(void* self, long long* __out_len) {
	auto&& __r = ((QUuid*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUuid_toByteArray(void* self) {
	auto&& __r = ((QUuid*)self)->toByteArray();
	return new QByteArray(__r);
}
void* QUuid_toBytes(void* self) {
	auto&& __r = ((QUuid*)self)->toBytes();
	return new QUuid::Id128Bytes(__r);
}
void* QUuid_toRfc4122(void* self) {
	auto&& __r = ((QUuid*)self)->toRfc4122();
	return new QByteArray(__r);
}
void* QUuid_fromRfc4122(void* arg1) {
	auto&& __r = QUuid::fromRfc4122(*(QByteArrayView*)arg1);
	return new QUuid(__r);
}
int QUuid_isNull(void* self) {
	auto&& __r = ((QUuid*)self)->isNull();
	return __r;
}
void* QUuid_ToGUID(void* self) {
	auto&& __r = static_cast<GUID>(*(QUuid*)self);
	return new _GUID(__r);
}
void* QUuid_createUuid() {
	auto&& __r = QUuid::createUuid();
	return new QUuid(__r);
}
void* QUuid_createUuidV5(void* ns, void* baseData) {
	auto&& __r = QUuid::createUuidV5(*(QUuid*)ns, *(QByteArrayView*)baseData);
	return new QUuid(__r);
}
void* QUuid_createUuidV3(void* ns, void* baseData) {
	auto&& __r = QUuid::createUuidV3(*(QUuid*)ns, *(QByteArrayView*)baseData);
	return new QUuid(__r);
}
void* QUuid_createUuidV7() {
	auto&& __r = QUuid::createUuidV7();
	return new QUuid(__r);
}
int QUuid_variant(void* self) {
	auto&& __r = ((QUuid*)self)->variant();
	return (int)(__r);
}
int QUuid_version(void* self) {
	auto&& __r = ((QUuid*)self)->version();
	return (int)(__r);
}
void* QUuid_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QUuid*)self)->operator=(*(QUuid*)arg1);
	return (void*)&(__r);
}
char* QUuid_toString2(void* self, int mode, long long* __out_len) {
	auto&& __r = ((QUuid*)self)->toString((QUuid::StringFormat)(mode));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUuid_toByteArray2(void* self, int mode) {
	auto&& __r = ((QUuid*)self)->toByteArray((QUuid::StringFormat)(mode));
	return new QByteArray(__r);
}
void* QUuid_toBytes2(void* self, int order) {
	auto&& __r = ((QUuid*)self)->toBytes((QSysInfo::Endian)(order));
	return new QUuid::Id128Bytes(__r);
}
int QUuid_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QUuid*)self) == (*(QUuid*)rhs);
	return __r;
}
int QUuid_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QUuid*)self) != (*(QUuid*)rhs);
	return __r;
}
int QUuid_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QUuid*)self) < (*(QUuid*)rhs);
	return __r;
}
int QUuid_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QUuid*)self) > (*(QUuid*)rhs);
	return __r;
}
int QUuid_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QUuid*)self) <= (*(QUuid*)rhs);
	return __r;
}
int QUuid_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QUuid*)self) >= (*(QUuid*)rhs);
	return __r;
}
unsigned int QUuid_data1_get(void* self) {
	auto __r = ((QUuid*)self)->data1;
	return __r;
}
void QUuid_data1_set(void* self, unsigned int data1) {
	((QUuid*)self)->data1 = data1;
}
unsigned short QUuid_data2_get(void* self) {
	auto __r = ((QUuid*)self)->data2;
	return __r;
}
void QUuid_data2_set(void* self, unsigned short data2) {
	((QUuid*)self)->data2 = data2;
}
unsigned short QUuid_data3_get(void* self) {
	auto __r = ((QUuid*)self)->data3;
	return __r;
}
void QUuid_data3_set(void* self, unsigned short data3) {
	((QUuid*)self)->data3 = data3;
}
void QUuid_delete(void* self) { delete (QUuid*)self; }

}

#include <QtCore/qcommandlineoption.h>
#include <QtCore/qstring.h>
#include "gen_qcommandlineoption.h"

extern "C" {

void* QCommandLineOption_new(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QCommandLineOption(name_qstr);
}
void* QCommandLineOption_new2(char** names, long long names_len) {
	QList<QString> names_list; names_list.reserve(names_len);
	for (long long __i = 0; __i < names_len; __i++) names_list.append(QString::fromUtf8(names[__i]));
	return new QCommandLineOption(names_list);
}
void* QCommandLineOption_new3(const char* name, long long name_len, const char* description, long long description_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	return new QCommandLineOption(name_qstr, description_qstr);
}
void* QCommandLineOption_new4(char** names, long long names_len, const char* description, long long description_len) {
	QList<QString> names_list; names_list.reserve(names_len);
	for (long long __i = 0; __i < names_len; __i++) names_list.append(QString::fromUtf8(names[__i]));
	QString description_qstr = QString::fromUtf8(description, description_len);
	return new QCommandLineOption(names_list, description_qstr);
}
void* QCommandLineOption_new5(void* other) {
	return new QCommandLineOption(*(QCommandLineOption*)other);
}
void* QCommandLineOption_new6(const char* name, long long name_len, const char* description, long long description_len, const char* valueName, long long valueName_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	QString valueName_qstr = QString::fromUtf8(valueName, valueName_len);
	return new QCommandLineOption(name_qstr, description_qstr, valueName_qstr);
}
void* QCommandLineOption_new7(const char* name, long long name_len, const char* description, long long description_len, const char* valueName, long long valueName_len, const char* defaultValue, long long defaultValue_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	QString valueName_qstr = QString::fromUtf8(valueName, valueName_len);
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	return new QCommandLineOption(name_qstr, description_qstr, valueName_qstr, defaultValue_qstr);
}
void* QCommandLineOption_new8(char** names, long long names_len, const char* description, long long description_len, const char* valueName, long long valueName_len) {
	QList<QString> names_list; names_list.reserve(names_len);
	for (long long __i = 0; __i < names_len; __i++) names_list.append(QString::fromUtf8(names[__i]));
	QString description_qstr = QString::fromUtf8(description, description_len);
	QString valueName_qstr = QString::fromUtf8(valueName, valueName_len);
	return new QCommandLineOption(names_list, description_qstr, valueName_qstr);
}
void* QCommandLineOption_new9(char** names, long long names_len, const char* description, long long description_len, const char* valueName, long long valueName_len, const char* defaultValue, long long defaultValue_len) {
	QList<QString> names_list; names_list.reserve(names_len);
	for (long long __i = 0; __i < names_len; __i++) names_list.append(QString::fromUtf8(names[__i]));
	QString description_qstr = QString::fromUtf8(description, description_len);
	QString valueName_qstr = QString::fromUtf8(valueName, valueName_len);
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	return new QCommandLineOption(names_list, description_qstr, valueName_qstr, defaultValue_qstr);
}
void* QCommandLineOption_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCommandLineOption*)self)->operator=(*(QCommandLineOption*)other);
	return (void*)&(__r);
}
void QCommandLineOption_swap(void* self, void* other) {
	((QCommandLineOption*)self)->swap(*(QCommandLineOption*)other);
}
char** QCommandLineOption_names(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineOption*)self)->names();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QCommandLineOption_setValueName(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QCommandLineOption*)self)->setValueName(name_qstr);
}
char* QCommandLineOption_valueName(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineOption*)self)->valueName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLineOption_setDescription(void* self, const char* description, long long description_len) {
	QString description_qstr = QString::fromUtf8(description, description_len);
	((QCommandLineOption*)self)->setDescription(description_qstr);
}
char* QCommandLineOption_description(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineOption*)self)->description();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLineOption_setDefaultValue(void* self, const char* defaultValue, long long defaultValue_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	((QCommandLineOption*)self)->setDefaultValue(defaultValue_qstr);
}
void QCommandLineOption_setDefaultValues(void* self, char** defaultValues, long long defaultValues_len) {
	QList<QString> defaultValues_list; defaultValues_list.reserve(defaultValues_len);
	for (long long __i = 0; __i < defaultValues_len; __i++) defaultValues_list.append(QString::fromUtf8(defaultValues[__i]));
	((QCommandLineOption*)self)->setDefaultValues(defaultValues_list);
}
char** QCommandLineOption_defaultValues(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineOption*)self)->defaultValues();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QCommandLineOption_flags(void* self) {
	auto&& __r = ((QCommandLineOption*)self)->flags();
	return (int)(__r);
}
void QCommandLineOption_setFlags(void* self, int aflags) {
	((QCommandLineOption*)self)->setFlags(QFlags<QCommandLineOption::Flag>(aflags));
}
void QCommandLineOption_delete(void* self) { delete (QCommandLineOption*)self; }

}

#include <QtCore/qrunnable.h>
#include "gen_qrunnable.h"
#include "gen_purego_cb.h"

class QRunnable__Go : public QRunnable {
public:
	long long __goId;
	unsigned char __goOn[1] = {};
	template<typename... A> QRunnable__Go(long long id, A&&... a) : QRunnable(std::forward<A>(a)...), __goId(id) {}
	void run() override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8524566317707732406LL, &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QRunnable__set_override(void* self, int idx) { ((QRunnable__Go*)self)->__goOn[idx] = 1; }
void QRunnable_run(void* self) {
	((QRunnable*)self)->run();
}
int QRunnable_autoDelete(void* self) {
	auto&& __r = ((QRunnable*)self)->autoDelete();
	return __r;
}
void QRunnable_setAutoDelete(void* self, int autoDelete) {
	((QRunnable*)self)->setAutoDelete(autoDelete);
}
void* QRunnable__Go_new(long long id) { return new QRunnable__Go(id); }
void QRunnable_delete(void* self) { delete (QRunnable*)self; }

}

#include <QtCore/qloggingcategory.h>
#include <QtCore/qlogging.h>
#include <QtCore/qstring.h>
#include "gen_qloggingcategory.h"

extern "C" {

void* QLoggingCategory_new(const char* category) {
	return new QLoggingCategory(category);
}
void* QLoggingCategory_new2(const char* category, int severityLevel) {
	return new QLoggingCategory(category, (QtMsgType)(severityLevel));
}
int QLoggingCategory_isEnabled(void* self, int type_) {
	auto&& __r = ((QLoggingCategory*)self)->isEnabled((QtMsgType)(type_));
	return __r;
}
void QLoggingCategory_setEnabled(void* self, int type_, int enable) {
	((QLoggingCategory*)self)->setEnabled((QtMsgType)(type_), enable);
}
int QLoggingCategory_isDebugEnabled(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->isDebugEnabled();
	return __r;
}
int QLoggingCategory_isInfoEnabled(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->isInfoEnabled();
	return __r;
}
int QLoggingCategory_isWarningEnabled(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->isWarningEnabled();
	return __r;
}
int QLoggingCategory_isCriticalEnabled(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->isCriticalEnabled();
	return __r;
}
const char* QLoggingCategory_categoryName(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->categoryName();
	return __r;
}
void* QLoggingCategory_OperatorCall(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->operator()();
	return (void*)&(__r);
}
void* QLoggingCategory_OperatorCall2(void* self) {
	auto&& __r = ((QLoggingCategory*)self)->operator()();
	return (void*)&(__r);
}
void* QLoggingCategory_defaultCategory() {
	auto&& __r = QLoggingCategory::defaultCategory();
	return (void*)(__r);
}
void QLoggingCategory_setFilterRules(const char* rules, long long rules_len) {
	QString rules_qstr = QString::fromUtf8(rules, rules_len);
	QLoggingCategory::setFilterRules(rules_qstr);
}
void QLoggingCategory_delete(void* self) { delete (QLoggingCategory*)self; }

}

#include <QtCore/qplugin.h>
#include "gen_qpluginmetadata.h"

extern "C" {

unsigned char QPluginMetaData_archRequirements() {
	auto&& __r = QPluginMetaData::archRequirements();
	return __r;
}
unsigned long long QPluginMetaData_size_get(void* self) {
	auto __r = ((QPluginMetaData*)self)->size;
	return __r;
}
void QPluginMetaData_size_set(void* self, unsigned long long size) {
	((QPluginMetaData*)self)->size = size;
}
void QPluginMetaData_delete(void* self) { delete (QPluginMetaData*)self; }

}

#include <QtCore/qsemaphore.h>
#include <QtCore/qdeadlinetimer.h>
#include "gen_qsemaphore.h"

extern "C" {

void* QSemaphore_new() {
	return new QSemaphore();
}
void* QSemaphore_new2(int n) {
	return new QSemaphore(n);
}
void QSemaphore_acquire(void* self) {
	((QSemaphore*)self)->acquire();
}
int QSemaphore_tryAcquire(void* self) {
	auto&& __r = ((QSemaphore*)self)->tryAcquire();
	return __r;
}
int QSemaphore_tryAcquire2(void* self, int n, int timeout) {
	auto&& __r = ((QSemaphore*)self)->tryAcquire(n, timeout);
	return __r;
}
int QSemaphore_tryAcquire3(void* self, int n, void* timeout) {
	auto&& __r = ((QSemaphore*)self)->tryAcquire(n, *(QDeadlineTimer*)timeout);
	return __r;
}
void QSemaphore_release(void* self) {
	((QSemaphore*)self)->release();
}
int QSemaphore_available(void* self) {
	auto&& __r = ((QSemaphore*)self)->available();
	return __r;
}
int QSemaphore_try_acquire(void* self) {
	auto&& __r = ((QSemaphore*)self)->try_acquire();
	return __r;
}
void QSemaphore_acquire2(void* self, int n) {
	((QSemaphore*)self)->acquire(n);
}
int QSemaphore_tryAcquire4(void* self, int n) {
	auto&& __r = ((QSemaphore*)self)->tryAcquire(n);
	return __r;
}
void QSemaphore_release2(void* self, int n) {
	((QSemaphore*)self)->release(n);
}
void QSemaphore_delete(void* self) { delete (QSemaphore*)self; }

}

#include <QtCore/qurlquery.h>
#include <QtCore/qurl.h>
#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include "gen_qurlquery.h"

extern "C" {

void* QUrlQuery_new() {
	return new QUrlQuery();
}
void* QUrlQuery_new2(void* url) {
	return new QUrlQuery(*(QUrl*)url);
}
void* QUrlQuery_new3(const char* queryString, long long queryString_len) {
	QString queryString_qstr = QString::fromUtf8(queryString, queryString_len);
	return new QUrlQuery(queryString_qstr);
}
void* QUrlQuery_new4(void* other) {
	return new QUrlQuery(*(QUrlQuery*)other);
}
void* QUrlQuery_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QUrlQuery*)self)->operator=(*(QUrlQuery*)other);
	return (void*)&(__r);
}
void QUrlQuery_swap(void* self, void* other) {
	((QUrlQuery*)self)->swap(*(QUrlQuery*)other);
}
int QUrlQuery_isEmpty(void* self) {
	auto&& __r = ((QUrlQuery*)self)->isEmpty();
	return __r;
}
int QUrlQuery_isDetached(void* self) {
	auto&& __r = ((QUrlQuery*)self)->isDetached();
	return __r;
}
void QUrlQuery_clear(void* self) {
	((QUrlQuery*)self)->clear();
}
char* QUrlQuery_query(void* self, long long* __out_len) {
	auto&& __r = ((QUrlQuery*)self)->query();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrlQuery_setQuery(void* self, const char* queryString, long long queryString_len) {
	QString queryString_qstr = QString::fromUtf8(queryString, queryString_len);
	((QUrlQuery*)self)->setQuery(queryString_qstr);
}
char* QUrlQuery_toString(void* self, long long* __out_len) {
	auto&& __r = ((QUrlQuery*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrlQuery_setQueryDelimiters(void* self, void* valueDelimiter, void* pairDelimiter) {
	((QUrlQuery*)self)->setQueryDelimiters(*(QChar*)valueDelimiter, *(QChar*)pairDelimiter);
}
void* QUrlQuery_queryValueDelimiter(void* self) {
	auto&& __r = ((QUrlQuery*)self)->queryValueDelimiter();
	return new QChar(__r);
}
void* QUrlQuery_queryPairDelimiter(void* self) {
	auto&& __r = ((QUrlQuery*)self)->queryPairDelimiter();
	return new QChar(__r);
}
int QUrlQuery_hasQueryItem(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QUrlQuery*)self)->hasQueryItem(key_qstr);
	return __r;
}
void QUrlQuery_addQueryItem(void* self, const char* key, long long key_len, const char* value, long long value_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QUrlQuery*)self)->addQueryItem(key_qstr, value_qstr);
}
void QUrlQuery_removeQueryItem(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QUrlQuery*)self)->removeQueryItem(key_qstr);
}
char* QUrlQuery_queryItemValue(void* self, const char* key, long long key_len, long long* __out_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QUrlQuery*)self)->queryItemValue(key_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QUrlQuery_allQueryItemValues(void* self, const char* key, long long key_len, long long* __out_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QUrlQuery*)self)->allQueryItemValues(key_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QUrlQuery_removeAllQueryItems(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QUrlQuery*)self)->removeAllQueryItems(key_qstr);
}
unsigned short QUrlQuery_defaultQueryValueDelimiter() {
	auto&& __r = QUrlQuery::defaultQueryValueDelimiter();
	return __r;
}
unsigned short QUrlQuery_defaultQueryPairDelimiter() {
	auto&& __r = QUrlQuery::defaultQueryPairDelimiter();
	return __r;
}
char* QUrlQuery_query2(void* self, int encoding, long long* __out_len) {
	auto&& __r = ((QUrlQuery*)self)->query(QFlags<QUrl::ComponentFormattingOption>(encoding));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrlQuery_toString2(void* self, int encoding, long long* __out_len) {
	auto&& __r = ((QUrlQuery*)self)->toString(QFlags<QUrl::ComponentFormattingOption>(encoding));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrlQuery_queryItemValue2(void* self, const char* key, long long key_len, int encoding, long long* __out_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QUrlQuery*)self)->queryItemValue(key_qstr, QFlags<QUrl::ComponentFormattingOption>(encoding));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QUrlQuery_allQueryItemValues2(void* self, const char* key, long long key_len, int encoding, long long* __out_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QUrlQuery*)self)->allQueryItemValues(key_qstr, QFlags<QUrl::ComponentFormattingOption>(encoding));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QUrlQuery_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QUrlQuery*)self) == (*(QUrlQuery*)rhs);
	return __r;
}
int QUrlQuery_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QUrlQuery*)self) != (*(QUrlQuery*)rhs);
	return __r;
}
void QUrlQuery_delete(void* self) { delete (QUrlQuery*)self; }

}

#include <QtGui/qvectornd.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtGui/qmatrix4x4.h>
#include <QtCore/qrect.h>
#include <QtCore/qvariant.h>
#include "gen_qvector3d.h"

extern "C" {

void* QVector3D_new() {
	return new QVector3D();
}
void* QVector3D_new2(int arg1) {
	return new QVector3D((Qt::Initialization)(arg1));
}
void* QVector3D_new3(float xpos, float ypos, float zpos) {
	return new QVector3D(xpos, ypos, zpos);
}
void* QVector3D_new4(void* point) {
	return new QVector3D(*(QPoint*)point);
}
void* QVector3D_new5(void* point) {
	return new QVector3D(*(QPointF*)point);
}
void* QVector3D_new6(void* vector) {
	return new QVector3D(*(QVector2D*)vector);
}
void* QVector3D_new7(void* vector, float zpos) {
	return new QVector3D(*(QVector2D*)vector, zpos);
}
void* QVector3D_new8(void* vector) {
	return new QVector3D(*(QVector4D*)vector);
}
void* QVector3D_new9(void* arg1) {
	return new QVector3D(*(QVector3D*)arg1);
}
int QVector3D_isNull(void* self) {
	auto&& __r = ((QVector3D*)self)->isNull();
	return __r;
}
float QVector3D_x(void* self) {
	auto&& __r = ((QVector3D*)self)->x();
	return __r;
}
float QVector3D_y(void* self) {
	auto&& __r = ((QVector3D*)self)->y();
	return __r;
}
float QVector3D_z(void* self) {
	auto&& __r = ((QVector3D*)self)->z();
	return __r;
}
void QVector3D_setX(void* self, float x) {
	((QVector3D*)self)->setX(x);
}
void QVector3D_setY(void* self, float y) {
	((QVector3D*)self)->setY(y);
}
void QVector3D_setZ(void* self, float z) {
	((QVector3D*)self)->setZ(z);
}
float QVector3D_OperatorSubscript(void* self, int i) {
	auto&& __r = ((QVector3D*)self)->operator[](i);
	return __r;
}
float QVector3D_length(void* self) {
	auto&& __r = ((QVector3D*)self)->length();
	return __r;
}
float QVector3D_lengthSquared(void* self) {
	auto&& __r = ((QVector3D*)self)->lengthSquared();
	return __r;
}
void* QVector3D_normalized(void* self) {
	auto&& __r = ((QVector3D*)self)->normalized();
	return new QVector3D(__r);
}
void QVector3D_normalize(void* self) {
	((QVector3D*)self)->normalize();
}
void* QVector3D_OperatorPlusAssign(void* self, void* vector) {
	auto&& __r = ((QVector3D*)self)->operator+=(*(QVector3D*)vector);
	return (void*)&(__r);
}
void* QVector3D_OperatorMinusAssign(void* self, void* vector) {
	auto&& __r = ((QVector3D*)self)->operator-=(*(QVector3D*)vector);
	return (void*)&(__r);
}
void* QVector3D_OperatorMultiplyAssign(void* self, float factor) {
	auto&& __r = ((QVector3D*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QVector3D_OperatorMultiplyAssign2(void* self, void* vector) {
	auto&& __r = ((QVector3D*)self)->operator*=(*(QVector3D*)vector);
	return (void*)&(__r);
}
void* QVector3D_OperatorDivideAssign(void* self, float divisor) {
	auto&& __r = ((QVector3D*)self)->operator/=(divisor);
	return (void*)&(__r);
}
void* QVector3D_OperatorDivideAssign2(void* self, void* vector) {
	auto&& __r = ((QVector3D*)self)->operator/=(*(QVector3D*)vector);
	return (void*)&(__r);
}
float QVector3D_dotProduct(void* v1, void* v2) {
	auto&& __r = QVector3D::dotProduct(*(QVector3D*)v1, *(QVector3D*)v2);
	return __r;
}
void* QVector3D_crossProduct(void* v1, void* v2) {
	auto&& __r = QVector3D::crossProduct(*(QVector3D*)v1, *(QVector3D*)v2);
	return new QVector3D(__r);
}
void* QVector3D_normal(void* v1, void* v2) {
	auto&& __r = QVector3D::normal(*(QVector3D*)v1, *(QVector3D*)v2);
	return new QVector3D(__r);
}
void* QVector3D_normal2(void* v1, void* v2, void* v3) {
	auto&& __r = QVector3D::normal(*(QVector3D*)v1, *(QVector3D*)v2, *(QVector3D*)v3);
	return new QVector3D(__r);
}
void* QVector3D_project(void* self, void* modelView, void* projection, void* viewport) {
	auto&& __r = ((QVector3D*)self)->project(*(QMatrix4x4*)modelView, *(QMatrix4x4*)projection, *(QRect*)viewport);
	return new QVector3D(__r);
}
void* QVector3D_unproject(void* self, void* modelView, void* projection, void* viewport) {
	auto&& __r = ((QVector3D*)self)->unproject(*(QMatrix4x4*)modelView, *(QMatrix4x4*)projection, *(QRect*)viewport);
	return new QVector3D(__r);
}
float QVector3D_distanceToPoint(void* self, void* point) {
	auto&& __r = ((QVector3D*)self)->distanceToPoint(*(QVector3D*)point);
	return __r;
}
float QVector3D_distanceToPlane(void* self, void* plane, void* normal) {
	auto&& __r = ((QVector3D*)self)->distanceToPlane(*(QVector3D*)plane, *(QVector3D*)normal);
	return __r;
}
float QVector3D_distanceToPlane2(void* self, void* plane1, void* plane2, void* plane3) {
	auto&& __r = ((QVector3D*)self)->distanceToPlane(*(QVector3D*)plane1, *(QVector3D*)plane2, *(QVector3D*)plane3);
	return __r;
}
float QVector3D_distanceToLine(void* self, void* point, void* direction) {
	auto&& __r = ((QVector3D*)self)->distanceToLine(*(QVector3D*)point, *(QVector3D*)direction);
	return __r;
}
void* QVector3D_toVector2D(void* self) {
	auto&& __r = ((QVector3D*)self)->toVector2D();
	return new QVector2D(__r);
}
void* QVector3D_toVector4D(void* self) {
	auto&& __r = ((QVector3D*)self)->toVector4D();
	return new QVector4D(__r);
}
void* QVector3D_toPoint(void* self) {
	auto&& __r = ((QVector3D*)self)->toPoint();
	return new QPoint(__r);
}
void* QVector3D_toPointF(void* self) {
	auto&& __r = ((QVector3D*)self)->toPointF();
	return new QPointF(__r);
}
void* QVector3D_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QVector3D*)self);
	return new QVariant(__r);
}
int QVector3D_OperatorEqual(void* self, void* v2) {
	auto&& __r = (*(QVector3D*)self) == (*(QVector3D*)v2);
	return __r;
}
int QVector3D_OperatorNotEqual(void* self, void* v2) {
	auto&& __r = (*(QVector3D*)self) != (*(QVector3D*)v2);
	return __r;
}
void* QVector3D_OperatorPlus(void* self, void* v2) {
	auto&& __r = (*(QVector3D*)self) + (*(QVector3D*)v2);
	return new QVector3D(__r);
}
void* QVector3D_OperatorMinus(void* self, void* v2) {
	auto&& __r = (*(QVector3D*)self) - (*(QVector3D*)v2);
	return new QVector3D(__r);
}
void* QVector3D_OperatorMultiply(void* self, float factor) {
	auto&& __r = (*(QVector3D*)self) * (factor);
	return new QVector3D(__r);
}
void* QVector3D_OperatorMultiply2(void* self, void* v2) {
	auto&& __r = (*(QVector3D*)self) * (*(QVector3D*)v2);
	return new QVector3D(__r);
}
void* QVector3D_OperatorMinus2(void* self) {
	auto&& __r = -(*(QVector3D*)self);
	return new QVector3D(__r);
}
void* QVector3D_OperatorDivide(void* self, float divisor) {
	auto&& __r = (*(QVector3D*)self) / (divisor);
	return new QVector3D(__r);
}
void* QVector3D_OperatorDivide2(void* self, void* divisor) {
	auto&& __r = (*(QVector3D*)self) / (*(QVector3D*)divisor);
	return new QVector3D(__r);
}
void* QVector3D_OperatorMultiply3(void* self, void* matrix) {
	auto&& __r = (*(QVector3D*)self) * (*(QMatrix4x4*)matrix);
	return new QVector3D(__r);
}
void QVector3D_delete(void* self) { delete (QVector3D*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qplatformsurfaceevent.h"

extern "C" {

void* QPlatformSurfaceEvent_base(void* self) { return (void*)static_cast<QEvent*>((QPlatformSurfaceEvent*)self); }
void* QPlatformSurfaceEvent_new(int surfaceEventType) {
	return new QPlatformSurfaceEvent((QPlatformSurfaceEvent::SurfaceEventType)(surfaceEventType));
}
int QPlatformSurfaceEvent_surfaceEventType(void* self) {
	auto&& __r = ((QPlatformSurfaceEvent*)self)->surfaceEventType();
	return (int)(__r);
}
int QPlatformSurfaceEvent_registerEventType() {
	auto&& __r = QPlatformSurfaceEvent::registerEventType();
	return __r;
}
int QPlatformSurfaceEvent_registerEventType2(int hint) {
	auto&& __r = QPlatformSurfaceEvent::registerEventType(hint);
	return __r;
}
void QPlatformSurfaceEvent_delete(void* self) { delete (QPlatformSurfaceEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include "gen_qscrollevent.h"

extern "C" {

void* QScrollEvent_base(void* self) { return (void*)static_cast<QEvent*>((QScrollEvent*)self); }
void* QScrollEvent_new(void* contentPos, void* overshoot, int scrollState) {
	return new QScrollEvent(*(QPointF*)contentPos, *(QPointF*)overshoot, (QScrollEvent::ScrollState)(scrollState));
}
void* QScrollEvent_contentPos(void* self) {
	auto&& __r = ((QScrollEvent*)self)->contentPos();
	return new QPointF(__r);
}
void* QScrollEvent_overshootDistance(void* self) {
	auto&& __r = ((QScrollEvent*)self)->overshootDistance();
	return new QPointF(__r);
}
int QScrollEvent_scrollState(void* self) {
	auto&& __r = ((QScrollEvent*)self)->scrollState();
	return (int)(__r);
}
int QScrollEvent_registerEventType() {
	auto&& __r = QScrollEvent::registerEventType();
	return __r;
}
int QScrollEvent_registerEventType2(int hint) {
	auto&& __r = QScrollEvent::registerEventType(hint);
	return __r;
}
void QScrollEvent_delete(void* self) { delete (QScrollEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstring.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include "gen_qtextblockformat.h"

extern "C" {

void* QTextBlockFormat_base(void* self) { return (void*)static_cast<QTextFormat*>((QTextBlockFormat*)self); }
void* QTextBlockFormat_new() {
	return new QTextBlockFormat();
}
void* QTextBlockFormat_new2(void* arg1) {
	return new QTextBlockFormat(*(QTextBlockFormat*)arg1);
}
int QTextBlockFormat_isValid(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->isValid();
	return __r;
}
void QTextBlockFormat_setAlignment(void* self, int alignment) {
	((QTextBlockFormat*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QTextBlockFormat_alignment(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->alignment();
	return (int)(__r);
}
void QTextBlockFormat_setTopMargin(void* self, double margin) {
	((QTextBlockFormat*)self)->setTopMargin(margin);
}
double QTextBlockFormat_topMargin(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->topMargin();
	return __r;
}
void QTextBlockFormat_setBottomMargin(void* self, double margin) {
	((QTextBlockFormat*)self)->setBottomMargin(margin);
}
double QTextBlockFormat_bottomMargin(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->bottomMargin();
	return __r;
}
void QTextBlockFormat_setLeftMargin(void* self, double margin) {
	((QTextBlockFormat*)self)->setLeftMargin(margin);
}
double QTextBlockFormat_leftMargin(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->leftMargin();
	return __r;
}
void QTextBlockFormat_setRightMargin(void* self, double margin) {
	((QTextBlockFormat*)self)->setRightMargin(margin);
}
double QTextBlockFormat_rightMargin(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->rightMargin();
	return __r;
}
void QTextBlockFormat_setTextIndent(void* self, double aindent) {
	((QTextBlockFormat*)self)->setTextIndent(aindent);
}
double QTextBlockFormat_textIndent(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->textIndent();
	return __r;
}
void QTextBlockFormat_setIndent(void* self, int indent) {
	((QTextBlockFormat*)self)->setIndent(indent);
}
int QTextBlockFormat_indent(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->indent();
	return __r;
}
void QTextBlockFormat_setHeadingLevel(void* self, int alevel) {
	((QTextBlockFormat*)self)->setHeadingLevel(alevel);
}
int QTextBlockFormat_headingLevel(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->headingLevel();
	return __r;
}
void QTextBlockFormat_setLineHeight(void* self, double height, int heightType) {
	((QTextBlockFormat*)self)->setLineHeight(height, heightType);
}
double QTextBlockFormat_lineHeight(void* self, double scriptLineHeight, double scaling) {
	auto&& __r = ((QTextBlockFormat*)self)->lineHeight(scriptLineHeight, scaling);
	return __r;
}
double QTextBlockFormat_lineHeight2(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->lineHeight();
	return __r;
}
int QTextBlockFormat_lineHeightType(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->lineHeightType();
	return __r;
}
void QTextBlockFormat_setNonBreakableLines(void* self, int b) {
	((QTextBlockFormat*)self)->setNonBreakableLines(b);
}
int QTextBlockFormat_nonBreakableLines(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->nonBreakableLines();
	return __r;
}
void QTextBlockFormat_setPageBreakPolicy(void* self, int flags) {
	((QTextBlockFormat*)self)->setPageBreakPolicy(QFlags<QTextFormat::PageBreakFlag>(flags));
}
int QTextBlockFormat_pageBreakPolicy(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->pageBreakPolicy();
	return (int)(__r);
}
void QTextBlockFormat_setTabPositions(void* self, void** tabs, long long tabs_len) {
	QList<QTextOption::Tab> tabs_list; tabs_list.reserve(tabs_len);
	for (long long __i = 0; __i < tabs_len; __i++) tabs_list.append(*(QTextOption::Tab*)tabs[__i]);
	((QTextBlockFormat*)self)->setTabPositions(tabs_list);
}
void** QTextBlockFormat_tabPositions(void* self, long long* __out_len) {
	auto&& __r = ((QTextBlockFormat*)self)->tabPositions();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextOption::Tab(__e);
	return __arr;
}
void QTextBlockFormat_setMarker(void* self, int marker) {
	((QTextBlockFormat*)self)->setMarker((QTextBlockFormat::MarkerType)(marker));
}
int QTextBlockFormat_marker(void* self) {
	auto&& __r = ((QTextBlockFormat*)self)->marker();
	return (int)(__r);
}
void* QTextBlockFormat_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextBlockFormat*)self)->operator=(*(QTextBlockFormat*)arg1);
	return (void*)&(__r);
}
void QTextBlockFormat_delete(void* self) { delete (QTextBlockFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include "gen_qaccessibletablecellinterface.h"
#include "gen_purego_cb.h"

class QAccessibleTableCellInterface__Go : public QAccessibleTableCellInterface {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QAccessibleTableCellInterface__Go(long long id, A&&... a) : QAccessibleTableCellInterface(std::forward<A>(a)...), __goId(id) {}
	bool isSelected() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3350427747592150876LL, &__h); if (__h) return __r; }
		return bool{};
	}
	QList<QAccessibleInterface*> columnHeaderCells() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 1398826710565483375LL, &__len, &__h);
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
	QList<QAccessibleInterface*> rowHeaderCells() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, -2170165325896972357LL, &__len, &__h);
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
	int columnIndex() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1856647841747687763LL, &__h); if (__h) return __r; }
		return int{};
	}
	int rowIndex() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5428465724232560177LL, &__h); if (__h) return __r; }
		return int{};
	}
	int columnExtent() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3228800890628578745LL, &__h); if (__h) return __r; }
		return int{};
	}
	int rowExtent() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 6265358599971552931LL, &__h); if (__h) return __r; }
		return int{};
	}
	QAccessibleInterface* table() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 8583921542012250175LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return nullptr;
	}
};

extern "C" {

void QAccessibleTableCellInterface__set_override(void* self, int idx) { ((QAccessibleTableCellInterface__Go*)self)->__goOn[idx] = 1; }
int QAccessibleTableCellInterface_isSelected(void* self) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->isSelected();
	return __r;
}
void** QAccessibleTableCellInterface_columnHeaderCells(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->columnHeaderCells();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QAccessibleTableCellInterface_rowHeaderCells(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->rowHeaderCells();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QAccessibleTableCellInterface_columnIndex(void* self) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->columnIndex();
	return __r;
}
int QAccessibleTableCellInterface_rowIndex(void* self) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->rowIndex();
	return __r;
}
int QAccessibleTableCellInterface_columnExtent(void* self) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->columnExtent();
	return __r;
}
int QAccessibleTableCellInterface_rowExtent(void* self) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->rowExtent();
	return __r;
}
void* QAccessibleTableCellInterface_table(void* self) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->table();
	return (void*)(__r);
}
void* QAccessibleTableCellInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleTableCellInterface*)self)->operator=(*(QAccessibleTableCellInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleTableCellInterface__Go_new(long long id) { return new QAccessibleTableCellInterface__Go(id); }
void QAccessibleTableCellInterface_delete(void* self) { delete (QAccessibleTableCellInterface*)self; }

}

#include <QtGui/qsurfaceformat.h>
#include <QtGui/qcolorspace.h>
#include "gen_qsurfaceformat.h"

extern "C" {

void* QSurfaceFormat_new() {
	return new QSurfaceFormat();
}
void* QSurfaceFormat_new2(int options) {
	return new QSurfaceFormat(QFlags<QSurfaceFormat::FormatOption>(options));
}
void* QSurfaceFormat_new3(void* other) {
	return new QSurfaceFormat(*(QSurfaceFormat*)other);
}
void* QSurfaceFormat_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QSurfaceFormat*)self)->operator=(*(QSurfaceFormat*)other);
	return (void*)&(__r);
}
void QSurfaceFormat_setDepthBufferSize(void* self, int size) {
	((QSurfaceFormat*)self)->setDepthBufferSize(size);
}
int QSurfaceFormat_depthBufferSize(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->depthBufferSize();
	return __r;
}
void QSurfaceFormat_setStencilBufferSize(void* self, int size) {
	((QSurfaceFormat*)self)->setStencilBufferSize(size);
}
int QSurfaceFormat_stencilBufferSize(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->stencilBufferSize();
	return __r;
}
void QSurfaceFormat_setRedBufferSize(void* self, int size) {
	((QSurfaceFormat*)self)->setRedBufferSize(size);
}
int QSurfaceFormat_redBufferSize(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->redBufferSize();
	return __r;
}
void QSurfaceFormat_setGreenBufferSize(void* self, int size) {
	((QSurfaceFormat*)self)->setGreenBufferSize(size);
}
int QSurfaceFormat_greenBufferSize(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->greenBufferSize();
	return __r;
}
void QSurfaceFormat_setBlueBufferSize(void* self, int size) {
	((QSurfaceFormat*)self)->setBlueBufferSize(size);
}
int QSurfaceFormat_blueBufferSize(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->blueBufferSize();
	return __r;
}
void QSurfaceFormat_setAlphaBufferSize(void* self, int size) {
	((QSurfaceFormat*)self)->setAlphaBufferSize(size);
}
int QSurfaceFormat_alphaBufferSize(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->alphaBufferSize();
	return __r;
}
void QSurfaceFormat_setSamples(void* self, int numSamples) {
	((QSurfaceFormat*)self)->setSamples(numSamples);
}
int QSurfaceFormat_samples(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->samples();
	return __r;
}
void QSurfaceFormat_setSwapBehavior(void* self, int behavior) {
	((QSurfaceFormat*)self)->setSwapBehavior((QSurfaceFormat::SwapBehavior)(behavior));
}
int QSurfaceFormat_swapBehavior(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->swapBehavior();
	return (int)(__r);
}
int QSurfaceFormat_hasAlpha(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->hasAlpha();
	return __r;
}
void QSurfaceFormat_setProfile(void* self, int profile) {
	((QSurfaceFormat*)self)->setProfile((QSurfaceFormat::OpenGLContextProfile)(profile));
}
int QSurfaceFormat_profile(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->profile();
	return (int)(__r);
}
void QSurfaceFormat_setRenderableType(void* self, int type_) {
	((QSurfaceFormat*)self)->setRenderableType((QSurfaceFormat::RenderableType)(type_));
}
int QSurfaceFormat_renderableType(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->renderableType();
	return (int)(__r);
}
void QSurfaceFormat_setMajorVersion(void* self, int majorVersion) {
	((QSurfaceFormat*)self)->setMajorVersion(majorVersion);
}
int QSurfaceFormat_majorVersion(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->majorVersion();
	return __r;
}
void QSurfaceFormat_setMinorVersion(void* self, int minorVersion) {
	((QSurfaceFormat*)self)->setMinorVersion(minorVersion);
}
int QSurfaceFormat_minorVersion(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->minorVersion();
	return __r;
}
int QSurfaceFormat_version(void* self, int* __out_b) {
	auto&& __r = ((QSurfaceFormat*)self)->version();
	*__out_b = __r.second;
	return __r.first;
}
void QSurfaceFormat_setVersion(void* self, int major, int minor) {
	((QSurfaceFormat*)self)->setVersion(major, minor);
}
int QSurfaceFormat_stereo(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->stereo();
	return __r;
}
void QSurfaceFormat_setStereo(void* self, int enable) {
	((QSurfaceFormat*)self)->setStereo(enable);
}
void QSurfaceFormat_setOptions(void* self, int options) {
	((QSurfaceFormat*)self)->setOptions(QFlags<QSurfaceFormat::FormatOption>(options));
}
void QSurfaceFormat_setOption(void* self, int option) {
	((QSurfaceFormat*)self)->setOption((QSurfaceFormat::FormatOption)(option));
}
int QSurfaceFormat_testOption(void* self, int option) {
	auto&& __r = ((QSurfaceFormat*)self)->testOption((QSurfaceFormat::FormatOption)(option));
	return __r;
}
int QSurfaceFormat_options(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->options();
	return (int)(__r);
}
int QSurfaceFormat_swapInterval(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->swapInterval();
	return __r;
}
void QSurfaceFormat_setSwapInterval(void* self, int interval) {
	((QSurfaceFormat*)self)->setSwapInterval(interval);
}
void* QSurfaceFormat_colorSpace(void* self) {
	auto&& __r = ((QSurfaceFormat*)self)->colorSpace();
	return (void*)&(__r);
}
void QSurfaceFormat_setColorSpace(void* self, void* colorSpace) {
	((QSurfaceFormat*)self)->setColorSpace(*(QColorSpace*)colorSpace);
}
void QSurfaceFormat_setColorSpace2(void* self, int colorSpace) {
	((QSurfaceFormat*)self)->setColorSpace((QSurfaceFormat::ColorSpace)(colorSpace));
}
void QSurfaceFormat_setDefaultFormat(void* format) {
	QSurfaceFormat::setDefaultFormat(*(QSurfaceFormat*)format);
}
void* QSurfaceFormat_defaultFormat() {
	auto&& __r = QSurfaceFormat::defaultFormat();
	return new QSurfaceFormat(__r);
}
void QSurfaceFormat_setOption2(void* self, int option, int on) {
	((QSurfaceFormat*)self)->setOption((QSurfaceFormat::FormatOption)(option), on);
}
int QSurfaceFormat_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QSurfaceFormat*)self) == (*(QSurfaceFormat*)rhs);
	return __r;
}
int QSurfaceFormat_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QSurfaceFormat*)self) != (*(QSurfaceFormat*)rhs);
	return __r;
}
void QSurfaceFormat_delete(void* self) { delete (QSurfaceFormat*)self; }

}

#include <QtGui/qimagereader.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtGui/qimage.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include <QtGui/qimageiohandler.h>
#include "gen_qimagereader.h"

extern "C" {

void* QImageReader_new() {
	return new QImageReader();
}
void* QImageReader_new2(void* device) {
	return new QImageReader((QIODevice*)device);
}
void* QImageReader_new3(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QImageReader(fileName_qstr);
}
void* QImageReader_new4(void* device, void* format) {
	return new QImageReader((QIODevice*)device, *(QByteArray*)format);
}
void* QImageReader_new5(const char* fileName, long long fileName_len, void* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QImageReader(fileName_qstr, *(QByteArray*)format);
}
char* QImageReader_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QImageReader::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImageReader_setFormat(void* self, void* format) {
	((QImageReader*)self)->setFormat(*(QByteArray*)format);
}
void* QImageReader_format(void* self) {
	auto&& __r = ((QImageReader*)self)->format();
	return new QByteArray(__r);
}
void QImageReader_setAutoDetectImageFormat(void* self, int enabled) {
	((QImageReader*)self)->setAutoDetectImageFormat(enabled);
}
int QImageReader_autoDetectImageFormat(void* self) {
	auto&& __r = ((QImageReader*)self)->autoDetectImageFormat();
	return __r;
}
void QImageReader_setDecideFormatFromContent(void* self, int ignored) {
	((QImageReader*)self)->setDecideFormatFromContent(ignored);
}
int QImageReader_decideFormatFromContent(void* self) {
	auto&& __r = ((QImageReader*)self)->decideFormatFromContent();
	return __r;
}
void QImageReader_setDevice(void* self, void* device) {
	((QImageReader*)self)->setDevice((QIODevice*)device);
}
void* QImageReader_device(void* self) {
	auto&& __r = ((QImageReader*)self)->device();
	return (void*)(__r);
}
void QImageReader_setFileName(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QImageReader*)self)->setFileName(fileName_qstr);
}
char* QImageReader_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QImageReader*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QImageReader_size(void* self) {
	auto&& __r = ((QImageReader*)self)->size();
	return new QSize(__r);
}
int QImageReader_imageFormat(void* self) {
	auto&& __r = ((QImageReader*)self)->imageFormat();
	return (int)(__r);
}
char** QImageReader_textKeys(void* self, long long* __out_len) {
	auto&& __r = ((QImageReader*)self)->textKeys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QImageReader_text(void* self, const char* key, long long key_len, long long* __out_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QImageReader*)self)->text(key_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImageReader_setClipRect(void* self, void* rect) {
	((QImageReader*)self)->setClipRect(*(QRect*)rect);
}
void* QImageReader_clipRect(void* self) {
	auto&& __r = ((QImageReader*)self)->clipRect();
	return new QRect(__r);
}
void QImageReader_setScaledSize(void* self, void* size) {
	((QImageReader*)self)->setScaledSize(*(QSize*)size);
}
void* QImageReader_scaledSize(void* self) {
	auto&& __r = ((QImageReader*)self)->scaledSize();
	return new QSize(__r);
}
void QImageReader_setQuality(void* self, int quality) {
	((QImageReader*)self)->setQuality(quality);
}
int QImageReader_quality(void* self) {
	auto&& __r = ((QImageReader*)self)->quality();
	return __r;
}
void QImageReader_setScaledClipRect(void* self, void* rect) {
	((QImageReader*)self)->setScaledClipRect(*(QRect*)rect);
}
void* QImageReader_scaledClipRect(void* self) {
	auto&& __r = ((QImageReader*)self)->scaledClipRect();
	return new QRect(__r);
}
void QImageReader_setBackgroundColor(void* self, void* color) {
	((QImageReader*)self)->setBackgroundColor(*(QColor*)color);
}
void* QImageReader_backgroundColor(void* self) {
	auto&& __r = ((QImageReader*)self)->backgroundColor();
	return new QColor(__r);
}
int QImageReader_supportsAnimation(void* self) {
	auto&& __r = ((QImageReader*)self)->supportsAnimation();
	return __r;
}
int QImageReader_transformation(void* self) {
	auto&& __r = ((QImageReader*)self)->transformation();
	return (int)(__r);
}
void QImageReader_setAutoTransform(void* self, int enabled) {
	((QImageReader*)self)->setAutoTransform(enabled);
}
int QImageReader_autoTransform(void* self) {
	auto&& __r = ((QImageReader*)self)->autoTransform();
	return __r;
}
void* QImageReader_subType(void* self) {
	auto&& __r = ((QImageReader*)self)->subType();
	return new QByteArray(__r);
}
void** QImageReader_supportedSubTypes(void* self, long long* __out_len) {
	auto&& __r = ((QImageReader*)self)->supportedSubTypes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
int QImageReader_canRead(void* self) {
	auto&& __r = ((QImageReader*)self)->canRead();
	return __r;
}
void* QImageReader_read(void* self) {
	auto&& __r = ((QImageReader*)self)->read();
	return new QImage(__r);
}
int QImageReader_read2(void* self, void* image) {
	auto&& __r = ((QImageReader*)self)->read((QImage*)image);
	return __r;
}
int QImageReader_jumpToNextImage(void* self) {
	auto&& __r = ((QImageReader*)self)->jumpToNextImage();
	return __r;
}
int QImageReader_jumpToImage(void* self, int imageNumber) {
	auto&& __r = ((QImageReader*)self)->jumpToImage(imageNumber);
	return __r;
}
int QImageReader_loopCount(void* self) {
	auto&& __r = ((QImageReader*)self)->loopCount();
	return __r;
}
int QImageReader_imageCount(void* self) {
	auto&& __r = ((QImageReader*)self)->imageCount();
	return __r;
}
int QImageReader_nextImageDelay(void* self) {
	auto&& __r = ((QImageReader*)self)->nextImageDelay();
	return __r;
}
int QImageReader_currentImageNumber(void* self) {
	auto&& __r = ((QImageReader*)self)->currentImageNumber();
	return __r;
}
void* QImageReader_currentImageRect(void* self) {
	auto&& __r = ((QImageReader*)self)->currentImageRect();
	return new QRect(__r);
}
int QImageReader_error(void* self) {
	auto&& __r = ((QImageReader*)self)->error();
	return (int)(__r);
}
char* QImageReader_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QImageReader*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QImageReader_supportsOption(void* self, int option) {
	auto&& __r = ((QImageReader*)self)->supportsOption((QImageIOHandler::ImageOption)(option));
	return __r;
}
void* QImageReader_imageFormat2(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QImageReader::imageFormat(fileName_qstr);
	return new QByteArray(__r);
}
void* QImageReader_imageFormat3(void* device) {
	auto&& __r = QImageReader::imageFormat((QIODevice*)device);
	return new QByteArray(__r);
}
void** QImageReader_supportedImageFormats(long long* __out_len) {
	auto&& __r = QImageReader::supportedImageFormats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QImageReader_supportedMimeTypes(long long* __out_len) {
	auto&& __r = QImageReader::supportedMimeTypes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QImageReader_imageFormatsForMimeType(void* mimeType, long long* __out_len) {
	auto&& __r = QImageReader::imageFormatsForMimeType(*(QByteArray*)mimeType);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
int QImageReader_allocationLimit() {
	auto&& __r = QImageReader::allocationLimit();
	return __r;
}
void QImageReader_setAllocationLimit(int mbLimit) {
	QImageReader::setAllocationLimit(mbLimit);
}
char* QImageReader_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QImageReader::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QImageReader_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QImageReader::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImageReader_delete(void* self) { delete (QImageReader*)self; }

}

#include <QtGui/qpdfoutputintent.h>
#include <QtCore/qstring.h>
#include <QtCore/qurl.h>
#include <QtGui/qcolorspace.h>
#include "gen_qpdfoutputintent.h"

extern "C" {

void* QPdfOutputIntent_new() {
	return new QPdfOutputIntent();
}
void* QPdfOutputIntent_new2(void* other) {
	return new QPdfOutputIntent(*(QPdfOutputIntent*)other);
}
void* QPdfOutputIntent_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPdfOutputIntent*)self)->operator=(*(QPdfOutputIntent*)other);
	return (void*)&(__r);
}
void QPdfOutputIntent_swap(void* self, void* other) {
	((QPdfOutputIntent*)self)->swap(*(QPdfOutputIntent*)other);
}
char* QPdfOutputIntent_outputConditionIdentifier(void* self, long long* __out_len) {
	auto&& __r = ((QPdfOutputIntent*)self)->outputConditionIdentifier();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfOutputIntent_setOutputConditionIdentifier(void* self, const char* identifier, long long identifier_len) {
	QString identifier_qstr = QString::fromUtf8(identifier, identifier_len);
	((QPdfOutputIntent*)self)->setOutputConditionIdentifier(identifier_qstr);
}
char* QPdfOutputIntent_outputCondition(void* self, long long* __out_len) {
	auto&& __r = ((QPdfOutputIntent*)self)->outputCondition();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfOutputIntent_setOutputCondition(void* self, const char* condition, long long condition_len) {
	QString condition_qstr = QString::fromUtf8(condition, condition_len);
	((QPdfOutputIntent*)self)->setOutputCondition(condition_qstr);
}
void* QPdfOutputIntent_registryName(void* self) {
	auto&& __r = ((QPdfOutputIntent*)self)->registryName();
	return new QUrl(__r);
}
void QPdfOutputIntent_setRegistryName(void* self, void* name) {
	((QPdfOutputIntent*)self)->setRegistryName(*(QUrl*)name);
}
void* QPdfOutputIntent_outputProfile(void* self) {
	auto&& __r = ((QPdfOutputIntent*)self)->outputProfile();
	return new QColorSpace(__r);
}
void QPdfOutputIntent_setOutputProfile(void* self, void* profile) {
	((QPdfOutputIntent*)self)->setOutputProfile(*(QColorSpace*)profile);
}
void QPdfOutputIntent_delete(void* self) { delete (QPdfOutputIntent*)self; }

}

#include <QtGui/qtextdocumentwriter.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextdocumentfragment.h>
#include "gen_qtextdocumentwriter.h"

extern "C" {

void* QTextDocumentWriter_new() {
	return new QTextDocumentWriter();
}
void* QTextDocumentWriter_new2(void* device, void* format) {
	return new QTextDocumentWriter((QIODevice*)device, *(QByteArray*)format);
}
void* QTextDocumentWriter_new3(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QTextDocumentWriter(fileName_qstr);
}
void* QTextDocumentWriter_new4(const char* fileName, long long fileName_len, void* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QTextDocumentWriter(fileName_qstr, *(QByteArray*)format);
}
void QTextDocumentWriter_setFormat(void* self, void* format) {
	((QTextDocumentWriter*)self)->setFormat(*(QByteArray*)format);
}
void* QTextDocumentWriter_format(void* self) {
	auto&& __r = ((QTextDocumentWriter*)self)->format();
	return new QByteArray(__r);
}
void QTextDocumentWriter_setDevice(void* self, void* device) {
	((QTextDocumentWriter*)self)->setDevice((QIODevice*)device);
}
void* QTextDocumentWriter_device(void* self) {
	auto&& __r = ((QTextDocumentWriter*)self)->device();
	return (void*)(__r);
}
void QTextDocumentWriter_setFileName(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QTextDocumentWriter*)self)->setFileName(fileName_qstr);
}
char* QTextDocumentWriter_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocumentWriter*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextDocumentWriter_write(void* self, void* document) {
	auto&& __r = ((QTextDocumentWriter*)self)->write((QTextDocument*)document);
	return __r;
}
int QTextDocumentWriter_write2(void* self, void* fragment) {
	auto&& __r = ((QTextDocumentWriter*)self)->write(*(QTextDocumentFragment*)fragment);
	return __r;
}
void** QTextDocumentWriter_supportedDocumentFormats(long long* __out_len) {
	auto&& __r = QTextDocumentWriter::supportedDocumentFormats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void QTextDocumentWriter_delete(void* self) { delete (QTextDocumentWriter*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qpalette.h>
#include <QtCore/qobject.h>
#include "gen_qstyleoption.h"

extern "C" {

void* QStyleOption_new() {
	return new QStyleOption();
}
void* QStyleOption_new2(void* other) {
	return new QStyleOption(*(QStyleOption*)other);
}
void* QStyleOption_new3(int version) {
	return new QStyleOption(version);
}
void* QStyleOption_new4(int version, int type_) {
	return new QStyleOption(version, type_);
}
void QStyleOption_initFrom(void* self, void* w) {
	((QStyleOption*)self)->initFrom((QWidget*)w);
}
void* QStyleOption_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QStyleOption*)self)->operator=(*(QStyleOption*)other);
	return (void*)&(__r);
}
int QStyleOption_version_get(void* self) {
	auto __r = ((QStyleOption*)self)->version;
	return __r;
}
void QStyleOption_version_set(void* self, int version) {
	((QStyleOption*)self)->version = version;
}
int QStyleOption_type_get(void* self) {
	auto __r = ((QStyleOption*)self)->type;
	return __r;
}
void QStyleOption_type_set(void* self, int type_) {
	((QStyleOption*)self)->type = type_;
}
int QStyleOption_state_get(void* self) {
	auto __r = ((QStyleOption*)self)->state;
	return (int)(__r);
}
void QStyleOption_state_set(void* self, int state) {
	((QStyleOption*)self)->state = QFlags<QStyle::StateFlag>(state);
}
int QStyleOption_direction_get(void* self) {
	auto __r = ((QStyleOption*)self)->direction;
	return (int)(__r);
}
void QStyleOption_direction_set(void* self, int direction) {
	((QStyleOption*)self)->direction = (Qt::LayoutDirection)(direction);
}
void* QStyleOption_rect_get(void* self) {
	auto __r = ((QStyleOption*)self)->rect;
	return new QRect(__r);
}
void QStyleOption_rect_set(void* self, void* rect) {
	((QStyleOption*)self)->rect = *(QRect*)rect;
}
void* QStyleOption_fontMetrics_get(void* self) {
	auto __r = ((QStyleOption*)self)->fontMetrics;
	return new QFontMetrics(__r);
}
void QStyleOption_fontMetrics_set(void* self, void* fontMetrics) {
	((QStyleOption*)self)->fontMetrics = *(QFontMetrics*)fontMetrics;
}
void* QStyleOption_palette_get(void* self) {
	auto __r = ((QStyleOption*)self)->palette;
	return new QPalette(__r);
}
void QStyleOption_palette_set(void* self, void* palette) {
	((QStyleOption*)self)->palette = *(QPalette*)palette;
}
void* QStyleOption_styleObject_get(void* self) {
	auto __r = ((QStyleOption*)self)->styleObject;
	return (void*)(__r);
}
void QStyleOption_styleObject_set(void* self, void* styleObject) {
	((QStyleOption*)self)->styleObject = (QObject*)styleObject;
}
void QStyleOption_delete(void* self) { delete (QStyleOption*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include "gen_qstyleoptionsizegrip.h"

extern "C" {

void* QStyleOptionSizeGrip_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionSizeGrip*)self); }
void QStyleOptionSizeGrip_chain(void* self, void** out) {
	QStyleOptionSizeGrip* __p = (QStyleOptionSizeGrip*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionSizeGrip_new() {
	return new QStyleOptionSizeGrip();
}
void* QStyleOptionSizeGrip_new2(void* other) {
	return new QStyleOptionSizeGrip(*(QStyleOptionSizeGrip*)other);
}
void* QStyleOptionSizeGrip_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionSizeGrip*)self)->operator=(*(QStyleOptionSizeGrip*)arg1);
	return (void*)&(__r);
}
int QStyleOptionSizeGrip_corner_get(void* self) {
	auto __r = ((QStyleOptionSizeGrip*)self)->corner;
	return (int)(__r);
}
void QStyleOptionSizeGrip_corner_set(void* self, int corner) {
	((QStyleOptionSizeGrip*)self)->corner = (Qt::Corner)(corner);
}
void QStyleOptionSizeGrip_delete(void* self) { delete (QStyleOptionSizeGrip*)self; }

}

#include <QtWidgets/qcolordialog.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qcolor.h>
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
#include "gen_qcolordialog.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QColorDialog__Go : public QColorDialog {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QColorDialog__Go(long long id, A&&... a) : QColorDialog(std::forward<A>(a)...), __goId(id) {}
	using QColorDialog::adjustPosition;
	using QColorDialog::updateMicroFocus;
	using QColorDialog::create;
	using QColorDialog::destroy;
	using QColorDialog::focusNextChild;
	using QColorDialog::focusPreviousChild;
	using QColorDialog::sender;
	using QColorDialog::senderSignalIndex;
	using QColorDialog::receivers;
	using QColorDialog::isSignalConnected;
	using QColorDialog::getDecodedMetricF;
	void cppgoBase_changeEvent(QEvent* __a0) { QColorDialog::changeEvent(__a0); }
	void cppgoBase_done(int __a0) { QColorDialog::done(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QColorDialog::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QColorDialog::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QColorDialog::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QColorDialog::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QColorDialog::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QColorDialog::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QColorDialog::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QColorDialog::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QColorDialog::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QColorDialog::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QColorDialog::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QColorDialog::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QColorDialog::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QColorDialog::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QColorDialog::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QColorDialog::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QColorDialog::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QColorDialog::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QColorDialog::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QColorDialog::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QColorDialog::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QColorDialog::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QColorDialog::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QColorDialog::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QColorDialog::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QColorDialog::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QColorDialog::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QColorDialog::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QColorDialog::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QColorDialog::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QColorDialog::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QColorDialog::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QColorDialog::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QColorDialog::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QColorDialog::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QColorDialog::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QColorDialog::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QColorDialog::metaObject();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QColorDialog::setVisible(visible);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::changeEvent(event);
	}
	void done(int result) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, result, &__h); if (__h) return; }
		QColorDialog::done(result);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QColorDialog::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QColorDialog::minimumSizeHint();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QColorDialog::exec();
	}
	void accept() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QColorDialog::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QColorDialog::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QColorDialog::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QColorDialog::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QColorDialog::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QColorDialog::resizeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QColorDialog::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QColorDialog::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QColorDialog::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QColorDialog::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QColorDialog::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QColorDialog::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QColorDialog::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QColorDialog::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QColorDialog::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QColorDialog::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QColorDialog::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QColorDialog::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QColorDialog::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QColorDialog::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QColorDialog::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QColorDialog::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QColorDialog::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QColorDialog::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QColorDialog;
static std::mutex __connsMu_QColorDialog;

extern "C" {

void* QColorDialog_base(void* self) { return (void*)static_cast<QDialog*>((QColorDialog*)self); }
void QColorDialog_chain(void* self, void** out) {
	QColorDialog* __p = (QColorDialog*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QColorDialog__set_override(void* self, int idx) { ((QColorDialog__Go*)self)->__goOn[idx] = 1; }
void* QColorDialog_fromQObject(void* o) { return (void*)qobject_cast<QColorDialog*>((QObject*)o); }
void* QColorDialog_new(long long __goid, void* parent) {
	return new QColorDialog__Go(__goid, (QWidget*)parent);
}
void* QColorDialog_new2(long long __goid) {
	return new QColorDialog__Go(__goid);
}
void* QColorDialog_new3(long long __goid, void* initial) {
	return new QColorDialog__Go(__goid, *(QColor*)initial);
}
void* QColorDialog_new4(long long __goid, void* initial, void* parent) {
	return new QColorDialog__Go(__goid, *(QColor*)initial, (QWidget*)parent);
}
char* QColorDialog_tr(const char* s, long long* __out_len) {
	auto&& __r = QColorDialog::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QColorDialog_setCurrentColor(void* self, void* color) {
	((QColorDialog*)self)->setCurrentColor(*(QColor*)color);
}
void* QColorDialog_currentColor(void* self) {
	auto&& __r = ((QColorDialog*)self)->currentColor();
	return new QColor(__r);
}
void* QColorDialog_selectedColor(void* self) {
	auto&& __r = ((QColorDialog*)self)->selectedColor();
	return new QColor(__r);
}
void QColorDialog_setOption(void* self, int option) {
	((QColorDialog*)self)->setOption((QColorDialog::ColorDialogOption)(option));
}
int QColorDialog_testOption(void* self, int option) {
	auto&& __r = ((QColorDialog*)self)->testOption((QColorDialog::ColorDialogOption)(option));
	return __r;
}
void QColorDialog_setOptions(void* self, int options) {
	((QColorDialog*)self)->setOptions(QFlags<QColorDialog::ColorDialogOption>(options));
}
int QColorDialog_options(void* self) {
	auto&& __r = ((QColorDialog*)self)->options();
	return (int)(__r);
}
void QColorDialog_open(void* self, void* receiver, const char* member) {
	((QColorDialog*)self)->open((QObject*)receiver, member);
}
void* QColorDialog_getColor() {
	auto&& __r = QColorDialog::getColor();
	return new QColor(__r);
}
int QColorDialog_customCount() {
	auto&& __r = QColorDialog::customCount();
	return __r;
}
void* QColorDialog_customColor(int index) {
	auto&& __r = QColorDialog::customColor(index);
	return new QColor(__r);
}
void QColorDialog_setCustomColor(int index, void* color) {
	QColorDialog::setCustomColor(index, *(QColor*)color);
}
void* QColorDialog_standardColor(int index) {
	auto&& __r = QColorDialog::standardColor(index);
	return new QColor(__r);
}
void QColorDialog_setStandardColor(int index, void* color) {
	QColorDialog::setStandardColor(index, *(QColor*)color);
}
void QColorDialog_currentColorChanged(void* self, void* color) {
	((QColorDialog*)self)->currentColorChanged(*(QColor*)color);
}
void QColorDialog_colorSelected(void* self, void* color) {
	((QColorDialog*)self)->colorSelected(*(QColor*)color);
}
void QColorDialog_adjustPosition(void* self, void* arg1) {
	((QColorDialog__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QColorDialog_setTabOrder(void* arg1, void* arg2) {
	QColorDialog::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QColorDialog_mouseGrabber() {
	auto&& __r = QColorDialog::mouseGrabber();
	return (void*)(__r);
}
void* QColorDialog_keyboardGrabber() {
	auto&& __r = QColorDialog::keyboardGrabber();
	return (void*)(__r);
}
void* QColorDialog_find(unsigned long long arg1) {
	auto&& __r = QColorDialog::find(arg1);
	return (void*)(__r);
}
void* QColorDialog_createWindowContainer(void* window) {
	auto&& __r = QColorDialog::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QColorDialog_updateMicroFocus(void* self) {
	((QColorDialog__Go*)self)->updateMicroFocus();
}
void QColorDialog_create(void* self) {
	((QColorDialog__Go*)self)->create();
}
void QColorDialog_destroy(void* self) {
	((QColorDialog__Go*)self)->destroy();
}
int QColorDialog_focusNextChild(void* self) {
	auto&& __r = ((QColorDialog__Go*)self)->focusNextChild();
	return __r;
}
int QColorDialog_focusPreviousChild(void* self) {
	auto&& __r = ((QColorDialog__Go*)self)->focusPreviousChild();
	return __r;
}
void* QColorDialog_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QColorDialog::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QColorDialog_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QColorDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QColorDialog_sender(void* self) {
	auto&& __r = ((QColorDialog__Go*)self)->sender();
	return (void*)(__r);
}
int QColorDialog_senderSignalIndex(void* self) {
	auto&& __r = ((QColorDialog__Go*)self)->senderSignalIndex();
	return __r;
}
int QColorDialog_receivers(void* self, const char* signal) {
	auto&& __r = ((QColorDialog__Go*)self)->receivers(signal);
	return __r;
}
int QColorDialog_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QColorDialog__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QColorDialog_devicePixelRatioFScale() {
	auto&& __r = QColorDialog::devicePixelRatioFScale();
	return __r;
}
int QColorDialog_encodeMetricF(int metric, double value) {
	auto&& __r = QColorDialog::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QColorDialog_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QColorDialog__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QColorDialog_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QColorDialog::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QColorDialog_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QColorDialog::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QColorDialog_setOption2(void* self, int option, int on) {
	((QColorDialog*)self)->setOption((QColorDialog::ColorDialogOption)(option), on);
}
void* QColorDialog_getColor2(void* initial) {
	auto&& __r = QColorDialog::getColor(*(QColor*)initial);
	return new QColor(__r);
}
void* QColorDialog_getColor3(void* initial, void* parent) {
	auto&& __r = QColorDialog::getColor(*(QColor*)initial, (QWidget*)parent);
	return new QColor(__r);
}
void* QColorDialog_getColor4(void* initial, void* parent, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = QColorDialog::getColor(*(QColor*)initial, (QWidget*)parent, title_qstr);
	return new QColor(__r);
}
void* QColorDialog_getColor5(void* initial, void* parent, const char* title, long long title_len, int options) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = QColorDialog::getColor(*(QColor*)initial, (QWidget*)parent, title_qstr, QFlags<QColorDialog::ColorDialogOption>(options));
	return new QColor(__r);
}
void* QColorDialog_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QColorDialog::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QColorDialog_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QColorDialog::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QColorDialog_updateMicroFocus2(void* self, int query) {
	((QColorDialog__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QColorDialog_create2(void* self, unsigned long long arg1) {
	((QColorDialog__Go*)self)->create(arg1);
}
void QColorDialog_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QColorDialog__Go*)self)->create(arg1, initializeWindow);
}
void QColorDialog_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QColorDialog__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QColorDialog_destroy2(void* self, int destroyWindow) {
	((QColorDialog__Go*)self)->destroy(destroyWindow);
}
void QColorDialog_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QColorDialog__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QColorDialog_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QColorDialog::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QColorDialog_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QColorDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QColorDialog_metaObject_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::metaObject();
	return (void*)(__r);
}
void QColorDialog_setVisible_default(void* self, int visible) {
	((QColorDialog*)self)->QColorDialog::setVisible(visible);
}
void QColorDialog_changeEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QColorDialog_done_default(void* self, int result) {
	((QColorDialog__Go*)self)->cppgoBase_done(result);
}
void* QColorDialog_sizeHint_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::sizeHint();
	return new QSize(__r);
}
void* QColorDialog_minimumSizeHint_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::minimumSizeHint();
	return new QSize(__r);
}
int QColorDialog_exec_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::exec();
	return __r;
}
void QColorDialog_accept_default(void* self) {
	((QColorDialog*)self)->QColorDialog::accept();
}
void QColorDialog_reject_default(void* self) {
	((QColorDialog*)self)->QColorDialog::reject();
}
void QColorDialog_keyPressEvent_default(void* self, void* arg1) {
	((QColorDialog__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QColorDialog_closeEvent_default(void* self, void* arg1) {
	((QColorDialog__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QColorDialog_showEvent_default(void* self, void* arg1) {
	((QColorDialog__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QColorDialog_resizeEvent_default(void* self, void* arg1) {
	((QColorDialog__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QColorDialog_contextMenuEvent_default(void* self, void* arg1) {
	((QColorDialog__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QColorDialog_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QColorDialog__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QColorDialog_devType_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::devType();
	return __r;
}
int QColorDialog_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::heightForWidth(arg1);
	return __r;
}
int QColorDialog_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::hasHeightForWidth();
	return __r;
}
void* QColorDialog_paintEngine_default(void* self) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::paintEngine();
	return (void*)(__r);
}
int QColorDialog_event_default(void* self, void* event) {
	auto&& __r = ((QColorDialog__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QColorDialog_mousePressEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QColorDialog_mouseReleaseEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QColorDialog_mouseDoubleClickEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QColorDialog_mouseMoveEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QColorDialog_wheelEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QColorDialog_keyReleaseEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QColorDialog_focusInEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QColorDialog_focusOutEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QColorDialog_enterEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QColorDialog_leaveEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QColorDialog_paintEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QColorDialog_moveEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QColorDialog_tabletEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QColorDialog_actionEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QColorDialog_dragEnterEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QColorDialog_dragMoveEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QColorDialog_dragLeaveEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QColorDialog_dropEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QColorDialog_hideEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QColorDialog_metric_default(void* self, int arg1) {
	auto&& __r = ((QColorDialog__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QColorDialog_initPainter_default(void* self, void* painter) {
	((QColorDialog__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QColorDialog_redirected_default(void* self, void* offset) {
	auto&& __r = ((QColorDialog__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QColorDialog_sharedPainter_default(void* self) {
	auto&& __r = ((QColorDialog__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QColorDialog_inputMethodEvent_default(void* self, void* arg1) {
	((QColorDialog__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QColorDialog_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QColorDialog*)self)->QColorDialog::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QColorDialog_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QColorDialog__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QColorDialog_timerEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QColorDialog_childEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QColorDialog_customEvent_default(void* self, void* event) {
	((QColorDialog__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QColorDialog_connectNotify_default(void* self, void* signal) {
	((QColorDialog__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QColorDialog_disconnectNotify_default(void* self, void* signal) {
	((QColorDialog__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QColorDialog_connect_currentColorChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::currentColorChanged, (QColorDialog*)self, [__id](const QColor& color){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&color);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_colorSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::colorSelected, (QColorDialog*)self, [__id](const QColor& color){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&color);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::finished, (QColorDialog*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::accepted, (QColorDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::rejected, (QColorDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::windowTitleChanged, (QColorDialog*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::windowIconChanged, (QColorDialog*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::windowIconTextChanged, (QColorDialog*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::customContextMenuRequested, (QColorDialog*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::destroyed, (QColorDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::destroyed, (QColorDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColorDialog*)self, &QColorDialog::objectNameChanged, (QColorDialog*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
	__conns_QColorDialog[__id] = __c;
}
void QColorDialog__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QColorDialog);
		auto __it = __conns_QColorDialog.find(id);
		if (__it == __conns_QColorDialog.end()) return;
		__c = __it->second;
		__conns_QColorDialog.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QColorDialog__Go_new(long long id) { return new QColorDialog__Go(id); }
void QColorDialog_delete(void* self) { delete (QColorDialog*)self; }

}

#include <QtWidgets/qgesture.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgesture.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGesture__Go : public QGesture {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QGesture__Go(long long id, A&&... a) : QGesture(std::forward<A>(a)...), __goId(id) {}
	using QGesture::sender;
	using QGesture::senderSignalIndex;
	using QGesture::receivers;
	using QGesture::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGesture::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGesture::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGesture::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGesture::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGesture::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGesture::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGesture::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGesture::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGesture::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGesture::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGesture::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGesture::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGesture::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGesture;
static std::mutex __connsMu_QGesture;

extern "C" {

void* QGesture_base(void* self) { return (void*)static_cast<QObject*>((QGesture*)self); }
void QGesture__set_override(void* self, int idx) { ((QGesture__Go*)self)->__goOn[idx] = 1; }
void* QGesture_fromQObject(void* o) { return (void*)qobject_cast<QGesture*>((QObject*)o); }
void* QGesture_new(long long __goid) {
	return new QGesture__Go(__goid);
}
void* QGesture_new2(long long __goid, void* parent) {
	return new QGesture__Go(__goid, (QObject*)parent);
}
char* QGesture_tr(const char* s, long long* __out_len) {
	auto&& __r = QGesture::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGesture_gestureType(void* self) {
	auto&& __r = ((QGesture*)self)->gestureType();
	return (int)(__r);
}
int QGesture_state(void* self) {
	auto&& __r = ((QGesture*)self)->state();
	return (int)(__r);
}
void* QGesture_hotSpot(void* self) {
	auto&& __r = ((QGesture*)self)->hotSpot();
	return new QPointF(__r);
}
void QGesture_setHotSpot(void* self, void* value) {
	((QGesture*)self)->setHotSpot(*(QPointF*)value);
}
int QGesture_hasHotSpot(void* self) {
	auto&& __r = ((QGesture*)self)->hasHotSpot();
	return __r;
}
void QGesture_unsetHotSpot(void* self) {
	((QGesture*)self)->unsetHotSpot();
}
void QGesture_setGestureCancelPolicy(void* self, int policy) {
	((QGesture*)self)->setGestureCancelPolicy((QGesture::GestureCancelPolicy)(policy));
}
int QGesture_gestureCancelPolicy(void* self) {
	auto&& __r = ((QGesture*)self)->gestureCancelPolicy();
	return (int)(__r);
}
void* QGesture_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGesture::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGesture_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGesture_sender(void* self) {
	auto&& __r = ((QGesture__Go*)self)->sender();
	return (void*)(__r);
}
int QGesture_senderSignalIndex(void* self) {
	auto&& __r = ((QGesture__Go*)self)->senderSignalIndex();
	return __r;
}
int QGesture_receivers(void* self, const char* signal) {
	auto&& __r = ((QGesture__Go*)self)->receivers(signal);
	return __r;
}
int QGesture_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGesture__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGesture_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGesture::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGesture_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGesture::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGesture::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGesture_metaObject_default(void* self) {
	auto&& __r = ((QGesture*)self)->QGesture::metaObject();
	return (void*)(__r);
}
int QGesture_event_default(void* self, void* event) {
	auto&& __r = ((QGesture*)self)->QGesture::event((QEvent*)event);
	return __r;
}
int QGesture_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGesture*)self)->QGesture::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGesture_timerEvent_default(void* self, void* event) {
	((QGesture__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGesture_childEvent_default(void* self, void* event) {
	((QGesture__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGesture_customEvent_default(void* self, void* event) {
	((QGesture__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGesture_connectNotify_default(void* self, void* signal) {
	((QGesture__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGesture_disconnectNotify_default(void* self, void* signal) {
	((QGesture__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGesture_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGesture*)self, &QGesture::destroyed, (QGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGesture);
	__conns_QGesture[__id] = __c;
}
void QGesture_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGesture*)self, &QGesture::destroyed, (QGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGesture);
	__conns_QGesture[__id] = __c;
}
void QGesture_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGesture*)self, &QGesture::objectNameChanged, (QGesture*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGesture);
	__conns_QGesture[__id] = __c;
}
void QGesture__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGesture);
		auto __it = __conns_QGesture.find(id);
		if (__it == __conns_QGesture.end()) return;
		__c = __it->second;
		__conns_QGesture.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGesture__Go_new(long long id, void* parent) { return new QGesture__Go(id, (QObject*)parent); }
void QGesture_delete(void* self) { delete (QGesture*)self; }

}

#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainter.h>
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
#include "gen_qgraphicseffect.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsEffect__Go : public QGraphicsEffect {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QGraphicsEffect__Go(long long id, A&&... a) : QGraphicsEffect(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsEffect::updateBoundingRect;
	using QGraphicsEffect::sourceIsPixmap;
	using QGraphicsEffect::sourceBoundingRect;
	using QGraphicsEffect::drawSource;
	using QGraphicsEffect::sourcePixmap;
	using QGraphicsEffect::sender;
	using QGraphicsEffect::senderSignalIndex;
	using QGraphicsEffect::receivers;
	using QGraphicsEffect::isSignalConnected;
	void cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> __a0) { QGraphicsEffect::sourceChanged(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsEffect::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsEffect::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsEffect::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsEffect::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsEffect::disconnectNotify(__a0); }
	void draw(QPainter* painter) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1044652152830004333LL, (void*)painter, &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsEffect::metaObject();
	}
	QRectF boundingRectFor(const QRectF& sourceRect) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 5831744872130794138LL, (void*)&sourceRect, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsEffect::boundingRectFor(sourceRect);
	}
	void sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> flags) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4520368087989481956LL, (int)(flags), &__h); if (__h) return; }
		QGraphicsEffect::sourceChanged(flags);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsEffect::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsEffect::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEffect::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEffect::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsEffect::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsEffect::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsEffect::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsEffect;
static std::mutex __connsMu_QGraphicsEffect;

extern "C" {

void* QGraphicsEffect_base(void* self) { return (void*)static_cast<QObject*>((QGraphicsEffect*)self); }
void QGraphicsEffect__set_override(void* self, int idx) { ((QGraphicsEffect__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsEffect_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsEffect*>((QObject*)o); }
char* QGraphicsEffect_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsEffect::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsEffect_boundingRectFor(void* self, void* sourceRect) {
	auto&& __r = ((QGraphicsEffect*)self)->boundingRectFor(*(QRectF*)sourceRect);
	return new QRectF(__r);
}
void* QGraphicsEffect_boundingRect(void* self) {
	auto&& __r = ((QGraphicsEffect*)self)->boundingRect();
	return new QRectF(__r);
}
int QGraphicsEffect_isEnabled(void* self) {
	auto&& __r = ((QGraphicsEffect*)self)->isEnabled();
	return __r;
}
void QGraphicsEffect_setEnabled(void* self, int enable) {
	((QGraphicsEffect*)self)->setEnabled(enable);
}
void QGraphicsEffect_update(void* self) {
	((QGraphicsEffect*)self)->update();
}
void QGraphicsEffect_enabledChanged(void* self, int enabled) {
	((QGraphicsEffect*)self)->enabledChanged(enabled);
}
void QGraphicsEffect_updateBoundingRect(void* self) {
	((QGraphicsEffect__Go*)self)->updateBoundingRect();
}
int QGraphicsEffect_sourceIsPixmap(void* self) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourceIsPixmap();
	return __r;
}
void* QGraphicsEffect_sourceBoundingRect(void* self) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourceBoundingRect();
	return new QRectF(__r);
}
void QGraphicsEffect_drawSource(void* self, void* painter) {
	((QGraphicsEffect__Go*)self)->drawSource((QPainter*)painter);
}
void* QGraphicsEffect_sourcePixmap(void* self) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourcePixmap();
	return new QPixmap(__r);
}
void* QGraphicsEffect_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsEffect::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsEffect_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsEffect_sender(void* self) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsEffect_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsEffect_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsEffect_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsEffect_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsEffect::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsEffect_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsEffect::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsEffect_sourceBoundingRect2(void* self, int system) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourceBoundingRect((Qt::CoordinateSystem)(system));
	return new QRectF(__r);
}
void* QGraphicsEffect_sourcePixmap2(void* self, int system) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system));
	return new QPixmap(__r);
}
void* QGraphicsEffect_sourcePixmap3(void* self, int system, void* offset) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset);
	return new QPixmap(__r);
}
void* QGraphicsEffect_sourcePixmap4(void* self, int system, void* offset, int mode) {
	auto&& __r = ((QGraphicsEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset, (QGraphicsEffect::PixmapPadMode)(mode));
	return new QPixmap(__r);
}
void* QGraphicsEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsEffect::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsEffect_boundingRectFor_default(void* self, void* sourceRect) {
	auto&& __r = ((QGraphicsEffect*)self)->QGraphicsEffect::boundingRectFor(*(QRectF*)sourceRect);
	return new QRectF(__r);
}
void* QGraphicsEffect_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsEffect*)self)->QGraphicsEffect::metaObject();
	return (void*)(__r);
}
void QGraphicsEffect_sourceChanged_default(void* self, int flags) {
	((QGraphicsEffect__Go*)self)->cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag>(flags));
}
int QGraphicsEffect_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsEffect*)self)->QGraphicsEffect::event((QEvent*)event);
	return __r;
}
int QGraphicsEffect_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsEffect*)self)->QGraphicsEffect::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsEffect_timerEvent_default(void* self, void* event) {
	((QGraphicsEffect__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsEffect_childEvent_default(void* self, void* event) {
	((QGraphicsEffect__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsEffect_customEvent_default(void* self, void* event) {
	((QGraphicsEffect__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsEffect_connectNotify_default(void* self, void* signal) {
	((QGraphicsEffect__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsEffect_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsEffect__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsEffect_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsEffect*)self, &QGraphicsEffect::enabledChanged, (QGraphicsEffect*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsEffect);
	__conns_QGraphicsEffect[__id] = __c;
}
void QGraphicsEffect_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsEffect*)self, &QGraphicsEffect::destroyed, (QGraphicsEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsEffect);
	__conns_QGraphicsEffect[__id] = __c;
}
void QGraphicsEffect_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsEffect*)self, &QGraphicsEffect::destroyed, (QGraphicsEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsEffect);
	__conns_QGraphicsEffect[__id] = __c;
}
void QGraphicsEffect_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsEffect*)self, &QGraphicsEffect::objectNameChanged, (QGraphicsEffect*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsEffect);
	__conns_QGraphicsEffect[__id] = __c;
}
void QGraphicsEffect__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsEffect);
		auto __it = __conns_QGraphicsEffect.find(id);
		if (__it == __conns_QGraphicsEffect.end()) return;
		__c = __it->second;
		__conns_QGraphicsEffect.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsEffect__Go_new(long long id, void* parent) { return new QGraphicsEffect__Go(id, (QObject*)parent); }
void QGraphicsEffect_delete(void* self) { delete (QGraphicsEffect*)self; }

}

#include <QtWidgets/qgraphicsview.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtGui/qpainter.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtGui/qtransform.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qgraphicsitem.h>
#include <QtGui/qpolygon.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qvariant.h>
#include <QtGui/qbrush.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qscrollbar.h>
#include <QtCore/qmargins.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qstyle.h>
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
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsView__Go : public QGraphicsView {
public:
	long long __goId;
	unsigned char __goOn[55] = {};
	template<typename... A> QGraphicsView__Go(long long id, A&&... a) : QGraphicsView(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsView::setViewportMargins;
	using QGraphicsView::viewportMargins;
	using QGraphicsView::drawFrame;
	using QGraphicsView::updateMicroFocus;
	using QGraphicsView::create;
	using QGraphicsView::destroy;
	using QGraphicsView::focusNextChild;
	using QGraphicsView::focusPreviousChild;
	using QGraphicsView::sender;
	using QGraphicsView::senderSignalIndex;
	using QGraphicsView::receivers;
	using QGraphicsView::isSignalConnected;
	using QGraphicsView::getDecodedMetricF;
	void cppgoBase_setupViewport(QWidget* __a0) { QGraphicsView::setupViewport(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QGraphicsView::event(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QGraphicsView::viewportEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QGraphicsView::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QGraphicsView::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QGraphicsView::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QGraphicsView::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QGraphicsView::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsView::focusInEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QGraphicsView::focusNextPrevChild(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsView::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsView::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsView::keyReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QGraphicsView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QGraphicsView::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QGraphicsView::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QGraphicsView::mouseReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QGraphicsView::wheelEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QGraphicsView::paintEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QGraphicsView::resizeEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QGraphicsView::scrollContentsBy(__a0, __a1); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QGraphicsView::showEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsView::inputMethodEvent(__a0); }
	void cppgoBase_drawBackground(QPainter* __a0, const QRectF& __a1) { QGraphicsView::drawBackground(__a0, __a1); }
	void cppgoBase_drawForeground(QPainter* __a0, const QRectF& __a1) { QGraphicsView::drawForeground(__a0, __a1); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QGraphicsView::eventFilter(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QGraphicsView::viewportSizeHint(); }
	void cppgoBase_changeEvent(QEvent* __a0) { QGraphicsView::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QGraphicsView::initStyleOption(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QGraphicsView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QGraphicsView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QGraphicsView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QGraphicsView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QGraphicsView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QGraphicsView::actionEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QGraphicsView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QGraphicsView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QGraphicsView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QGraphicsView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QGraphicsView::sharedPainter(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsView::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsView::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsView::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QGraphicsView::sizeHint();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsView::inputMethodQuery(query);
	}
	void setupViewport(QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)widget, &__h); if (__h) return; }
		QGraphicsView::setupViewport(widget);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsView::event(event);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsView::viewportEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::contextMenuEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::dragEnterEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::dragLeaveEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::dragMoveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::focusInEvent(event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QGraphicsView::focusNextPrevChild(next);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::keyReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::mouseDoubleClickEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::mousePressEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::mouseReleaseEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::wheelEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::paintEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::resizeEvent(event);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QGraphicsView::scrollContentsBy(dx, dy);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::showEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::inputMethodEvent(event);
	}
	void drawBackground(QPainter* painter, const QRectF& rect) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -3076319991076358961LL, (void*)painter, (void*)&rect, &__h); if (__h) return; }
		QGraphicsView::drawBackground(painter, rect);
	}
	void drawForeground(QPainter* painter, const QRectF& rect) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 8267095530709235866LL, (void*)painter, (void*)&rect, &__h); if (__h) return; }
		QGraphicsView::drawForeground(painter, rect);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QGraphicsView::minimumSizeHint();
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QGraphicsView::eventFilter(arg1, arg2);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QGraphicsView::viewportSizeHint();
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QGraphicsView::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QGraphicsView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QGraphicsView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QGraphicsView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QGraphicsView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QGraphicsView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QGraphicsView::paintEngine();
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::actionEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QGraphicsView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QGraphicsView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QGraphicsView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QGraphicsView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QGraphicsView::sharedPainter();
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsView;
static std::mutex __connsMu_QGraphicsView;

extern "C" {

void* QGraphicsView_base(void* self) { return (void*)static_cast<QAbstractScrollArea*>((QGraphicsView*)self); }
void QGraphicsView_chain(void* self, void** out) {
	QGraphicsView* __p = (QGraphicsView*)self;
	QAbstractScrollArea* __b0 = static_cast<QAbstractScrollArea*>(__p); out[0] = (void*)__b0;
	QFrame* __b1 = static_cast<QFrame*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QGraphicsView__set_override(void* self, int idx) { ((QGraphicsView__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsView_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsView*>((QObject*)o); }
void* QGraphicsView_new(long long __goid, void* parent) {
	return new QGraphicsView__Go(__goid, (QWidget*)parent);
}
void* QGraphicsView_new2(long long __goid) {
	return new QGraphicsView__Go(__goid);
}
void* QGraphicsView_new3(long long __goid, void* scene) {
	return new QGraphicsView__Go(__goid, (QGraphicsScene*)scene);
}
void* QGraphicsView_new4(long long __goid, void* scene, void* parent) {
	return new QGraphicsView__Go(__goid, (QGraphicsScene*)scene, (QWidget*)parent);
}
char* QGraphicsView_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGraphicsView_renderHints(void* self) {
	auto&& __r = ((QGraphicsView*)self)->renderHints();
	return (int)(__r);
}
void QGraphicsView_setRenderHint(void* self, int hint) {
	((QGraphicsView*)self)->setRenderHint((QPainter::RenderHint)(hint));
}
void QGraphicsView_setRenderHints(void* self, int hints) {
	((QGraphicsView*)self)->setRenderHints(QFlags<QPainter::RenderHint>(hints));
}
int QGraphicsView_alignment(void* self) {
	auto&& __r = ((QGraphicsView*)self)->alignment();
	return (int)(__r);
}
void QGraphicsView_setAlignment(void* self, int alignment) {
	((QGraphicsView*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QGraphicsView_transformationAnchor(void* self) {
	auto&& __r = ((QGraphicsView*)self)->transformationAnchor();
	return (int)(__r);
}
void QGraphicsView_setTransformationAnchor(void* self, int anchor) {
	((QGraphicsView*)self)->setTransformationAnchor((QGraphicsView::ViewportAnchor)(anchor));
}
int QGraphicsView_resizeAnchor(void* self) {
	auto&& __r = ((QGraphicsView*)self)->resizeAnchor();
	return (int)(__r);
}
void QGraphicsView_setResizeAnchor(void* self, int anchor) {
	((QGraphicsView*)self)->setResizeAnchor((QGraphicsView::ViewportAnchor)(anchor));
}
int QGraphicsView_viewportUpdateMode(void* self) {
	auto&& __r = ((QGraphicsView*)self)->viewportUpdateMode();
	return (int)(__r);
}
void QGraphicsView_setViewportUpdateMode(void* self, int mode) {
	((QGraphicsView*)self)->setViewportUpdateMode((QGraphicsView::ViewportUpdateMode)(mode));
}
int QGraphicsView_optimizationFlags(void* self) {
	auto&& __r = ((QGraphicsView*)self)->optimizationFlags();
	return (int)(__r);
}
void QGraphicsView_setOptimizationFlag(void* self, int flag) {
	((QGraphicsView*)self)->setOptimizationFlag((QGraphicsView::OptimizationFlag)(flag));
}
void QGraphicsView_setOptimizationFlags(void* self, int flags) {
	((QGraphicsView*)self)->setOptimizationFlags(QFlags<QGraphicsView::OptimizationFlag>(flags));
}
int QGraphicsView_dragMode(void* self) {
	auto&& __r = ((QGraphicsView*)self)->dragMode();
	return (int)(__r);
}
void QGraphicsView_setDragMode(void* self, int mode) {
	((QGraphicsView*)self)->setDragMode((QGraphicsView::DragMode)(mode));
}
int QGraphicsView_rubberBandSelectionMode(void* self) {
	auto&& __r = ((QGraphicsView*)self)->rubberBandSelectionMode();
	return (int)(__r);
}
void QGraphicsView_setRubberBandSelectionMode(void* self, int mode) {
	((QGraphicsView*)self)->setRubberBandSelectionMode((Qt::ItemSelectionMode)(mode));
}
void* QGraphicsView_rubberBandRect(void* self) {
	auto&& __r = ((QGraphicsView*)self)->rubberBandRect();
	return new QRect(__r);
}
int QGraphicsView_cacheMode(void* self) {
	auto&& __r = ((QGraphicsView*)self)->cacheMode();
	return (int)(__r);
}
void QGraphicsView_setCacheMode(void* self, int mode) {
	((QGraphicsView*)self)->setCacheMode(QFlags<QGraphicsView::CacheModeFlag>(mode));
}
void QGraphicsView_resetCachedContent(void* self) {
	((QGraphicsView*)self)->resetCachedContent();
}
int QGraphicsView_isInteractive(void* self) {
	auto&& __r = ((QGraphicsView*)self)->isInteractive();
	return __r;
}
void QGraphicsView_setInteractive(void* self, int allowed) {
	((QGraphicsView*)self)->setInteractive(allowed);
}
void* QGraphicsView_scene(void* self) {
	auto&& __r = ((QGraphicsView*)self)->scene();
	return (void*)(__r);
}
void QGraphicsView_setScene(void* self, void* scene) {
	((QGraphicsView*)self)->setScene((QGraphicsScene*)scene);
}
void* QGraphicsView_sceneRect(void* self) {
	auto&& __r = ((QGraphicsView*)self)->sceneRect();
	return new QRectF(__r);
}
void QGraphicsView_setSceneRect(void* self, void* rect) {
	((QGraphicsView*)self)->setSceneRect(*(QRectF*)rect);
}
void QGraphicsView_setSceneRect2(void* self, double x, double y, double w, double h) {
	((QGraphicsView*)self)->setSceneRect(x, y, w, h);
}
void* QGraphicsView_transform(void* self) {
	auto&& __r = ((QGraphicsView*)self)->transform();
	return new QTransform(__r);
}
void* QGraphicsView_viewportTransform(void* self) {
	auto&& __r = ((QGraphicsView*)self)->viewportTransform();
	return new QTransform(__r);
}
int QGraphicsView_isTransformed(void* self) {
	auto&& __r = ((QGraphicsView*)self)->isTransformed();
	return __r;
}
void QGraphicsView_setTransform(void* self, void* matrix) {
	((QGraphicsView*)self)->setTransform(*(QTransform*)matrix);
}
void QGraphicsView_resetTransform(void* self) {
	((QGraphicsView*)self)->resetTransform();
}
void QGraphicsView_rotate(void* self, double angle) {
	((QGraphicsView*)self)->rotate(angle);
}
void QGraphicsView_scale(void* self, double sx, double sy) {
	((QGraphicsView*)self)->scale(sx, sy);
}
void QGraphicsView_shear(void* self, double sh, double sv) {
	((QGraphicsView*)self)->shear(sh, sv);
}
void QGraphicsView_translate(void* self, double dx, double dy) {
	((QGraphicsView*)self)->translate(dx, dy);
}
void QGraphicsView_centerOn(void* self, void* pos) {
	((QGraphicsView*)self)->centerOn(*(QPointF*)pos);
}
void QGraphicsView_centerOn2(void* self, double x, double y) {
	((QGraphicsView*)self)->centerOn(x, y);
}
void QGraphicsView_centerOn3(void* self, void* item) {
	((QGraphicsView*)self)->centerOn((QGraphicsItem*)item);
}
void QGraphicsView_ensureVisible(void* self, void* rect) {
	((QGraphicsView*)self)->ensureVisible(*(QRectF*)rect);
}
void QGraphicsView_ensureVisible2(void* self, double x, double y, double w, double h) {
	((QGraphicsView*)self)->ensureVisible(x, y, w, h);
}
void QGraphicsView_ensureVisible3(void* self, void* item) {
	((QGraphicsView*)self)->ensureVisible((QGraphicsItem*)item);
}
void QGraphicsView_fitInView(void* self, void* rect) {
	((QGraphicsView*)self)->fitInView(*(QRectF*)rect);
}
void QGraphicsView_fitInView2(void* self, double x, double y, double w, double h) {
	((QGraphicsView*)self)->fitInView(x, y, w, h);
}
void QGraphicsView_fitInView3(void* self, void* item) {
	((QGraphicsView*)self)->fitInView((QGraphicsItem*)item);
}
void QGraphicsView_render(void* self, void* painter) {
	((QGraphicsView*)self)->render((QPainter*)painter);
}
void** QGraphicsView_items(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items2(void* self, void* pos, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QPoint*)pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items3(void* self, int x, int y, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(x, y);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items4(void* self, void* rect, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QRect*)rect);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items5(void* self, int x, int y, int w, int h, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(x, y, w, h);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items6(void* self, void* polygon, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QPolygon*)polygon);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items7(void* self, void* path, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QPainterPath*)path);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QGraphicsView_itemAt(void* self, void* pos) {
	auto&& __r = ((QGraphicsView*)self)->itemAt(*(QPoint*)pos);
	return (void*)(__r);
}
void* QGraphicsView_itemAt2(void* self, int x, int y) {
	auto&& __r = ((QGraphicsView*)self)->itemAt(x, y);
	return (void*)(__r);
}
void* QGraphicsView_mapToScene(void* self, void* point) {
	auto&& __r = ((QGraphicsView*)self)->mapToScene(*(QPoint*)point);
	return new QPointF(__r);
}
void* QGraphicsView_mapToScene2(void* self, void* rect) {
	auto&& __r = ((QGraphicsView*)self)->mapToScene(*(QRect*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsView_mapToScene3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsView*)self)->mapToScene(*(QPolygon*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsView_mapToScene4(void* self, void* path) {
	auto&& __r = ((QGraphicsView*)self)->mapToScene(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsView_mapFromScene(void* self, void* point) {
	auto&& __r = ((QGraphicsView*)self)->mapFromScene(*(QPointF*)point);
	return new QPoint(__r);
}
void* QGraphicsView_mapFromScene2(void* self, void* rect) {
	auto&& __r = ((QGraphicsView*)self)->mapFromScene(*(QRectF*)rect);
	return new QPolygon(__r);
}
void* QGraphicsView_mapFromScene3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsView*)self)->mapFromScene(*(QPolygonF*)polygon);
	return new QPolygon(__r);
}
void* QGraphicsView_mapFromScene4(void* self, void* path) {
	auto&& __r = ((QGraphicsView*)self)->mapFromScene(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsView_mapToScene5(void* self, int x, int y) {
	auto&& __r = ((QGraphicsView*)self)->mapToScene(x, y);
	return new QPointF(__r);
}
void* QGraphicsView_mapToScene6(void* self, int x, int y, int w, int h) {
	auto&& __r = ((QGraphicsView*)self)->mapToScene(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsView_mapFromScene5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsView*)self)->mapFromScene(x, y);
	return new QPoint(__r);
}
void* QGraphicsView_mapFromScene6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsView*)self)->mapFromScene(x, y, w, h);
	return new QPolygon(__r);
}
void* QGraphicsView_backgroundBrush(void* self) {
	auto&& __r = ((QGraphicsView*)self)->backgroundBrush();
	return new QBrush(__r);
}
void QGraphicsView_setBackgroundBrush(void* self, void* brush) {
	((QGraphicsView*)self)->setBackgroundBrush(*(QBrush*)brush);
}
void* QGraphicsView_foregroundBrush(void* self) {
	auto&& __r = ((QGraphicsView*)self)->foregroundBrush();
	return new QBrush(__r);
}
void QGraphicsView_setForegroundBrush(void* self, void* brush) {
	((QGraphicsView*)self)->setForegroundBrush(*(QBrush*)brush);
}
void QGraphicsView_updateScene(void* self, void** rects, long long rects_len) {
	QList<QRectF> rects_list; rects_list.reserve(rects_len);
	for (long long __i = 0; __i < rects_len; __i++) rects_list.append(*(QRectF*)rects[__i]);
	((QGraphicsView*)self)->updateScene(rects_list);
}
void QGraphicsView_invalidateScene(void* self) {
	((QGraphicsView*)self)->invalidateScene();
}
void QGraphicsView_updateSceneRect(void* self, void* rect) {
	((QGraphicsView*)self)->updateSceneRect(*(QRectF*)rect);
}
void QGraphicsView_rubberBandChanged(void* self, void* viewportRect, void* fromScenePoint, void* toScenePoint) {
	((QGraphicsView*)self)->rubberBandChanged(*(QRect*)viewportRect, *(QPointF*)fromScenePoint, *(QPointF*)toScenePoint);
}
void QGraphicsView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QGraphicsView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QGraphicsView_setViewportMargins2(void* self, void* margins) {
	((QGraphicsView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QGraphicsView_viewportMargins(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QGraphicsView_drawFrame(void* self, void* arg1) {
	((QGraphicsView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QGraphicsView_setTabOrder(void* arg1, void* arg2) {
	QGraphicsView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QGraphicsView_mouseGrabber() {
	auto&& __r = QGraphicsView::mouseGrabber();
	return (void*)(__r);
}
void* QGraphicsView_keyboardGrabber() {
	auto&& __r = QGraphicsView::keyboardGrabber();
	return (void*)(__r);
}
void* QGraphicsView_find(unsigned long long arg1) {
	auto&& __r = QGraphicsView::find(arg1);
	return (void*)(__r);
}
void* QGraphicsView_createWindowContainer(void* window) {
	auto&& __r = QGraphicsView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QGraphicsView_updateMicroFocus(void* self) {
	((QGraphicsView__Go*)self)->updateMicroFocus();
}
void QGraphicsView_create(void* self) {
	((QGraphicsView__Go*)self)->create();
}
void QGraphicsView_destroy(void* self) {
	((QGraphicsView__Go*)self)->destroy();
}
int QGraphicsView_focusNextChild(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->focusNextChild();
	return __r;
}
int QGraphicsView_focusPreviousChild(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QGraphicsView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsView_sender(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsView_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsView_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsView__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QGraphicsView_devicePixelRatioFScale() {
	auto&& __r = QGraphicsView::devicePixelRatioFScale();
	return __r;
}
int QGraphicsView_encodeMetricF(int metric, double value) {
	auto&& __r = QGraphicsView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QGraphicsView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QGraphicsView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QGraphicsView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsView_setRenderHint2(void* self, int hint, int enabled) {
	((QGraphicsView*)self)->setRenderHint((QPainter::RenderHint)(hint), enabled);
}
void QGraphicsView_setOptimizationFlag2(void* self, int flag, int enabled) {
	((QGraphicsView*)self)->setOptimizationFlag((QGraphicsView::OptimizationFlag)(flag), enabled);
}
void QGraphicsView_setTransform2(void* self, void* matrix, int combine) {
	((QGraphicsView*)self)->setTransform(*(QTransform*)matrix, combine);
}
void QGraphicsView_ensureVisible4(void* self, void* rect, int xmargin) {
	((QGraphicsView*)self)->ensureVisible(*(QRectF*)rect, xmargin);
}
void QGraphicsView_ensureVisible5(void* self, void* rect, int xmargin, int ymargin) {
	((QGraphicsView*)self)->ensureVisible(*(QRectF*)rect, xmargin, ymargin);
}
void QGraphicsView_ensureVisible6(void* self, double x, double y, double w, double h, int xmargin) {
	((QGraphicsView*)self)->ensureVisible(x, y, w, h, xmargin);
}
void QGraphicsView_ensureVisible7(void* self, double x, double y, double w, double h, int xmargin, int ymargin) {
	((QGraphicsView*)self)->ensureVisible(x, y, w, h, xmargin, ymargin);
}
void QGraphicsView_ensureVisible8(void* self, void* item, int xmargin) {
	((QGraphicsView*)self)->ensureVisible((QGraphicsItem*)item, xmargin);
}
void QGraphicsView_ensureVisible9(void* self, void* item, int xmargin, int ymargin) {
	((QGraphicsView*)self)->ensureVisible((QGraphicsItem*)item, xmargin, ymargin);
}
void QGraphicsView_fitInView4(void* self, void* rect, int aspectRadioMode) {
	((QGraphicsView*)self)->fitInView(*(QRectF*)rect, (Qt::AspectRatioMode)(aspectRadioMode));
}
void QGraphicsView_fitInView5(void* self, double x, double y, double w, double h, int aspectRadioMode) {
	((QGraphicsView*)self)->fitInView(x, y, w, h, (Qt::AspectRatioMode)(aspectRadioMode));
}
void QGraphicsView_fitInView6(void* self, void* item, int aspectRadioMode) {
	((QGraphicsView*)self)->fitInView((QGraphicsItem*)item, (Qt::AspectRatioMode)(aspectRadioMode));
}
void QGraphicsView_render2(void* self, void* painter, void* target) {
	((QGraphicsView*)self)->render((QPainter*)painter, *(QRectF*)target);
}
void QGraphicsView_render3(void* self, void* painter, void* target, void* source) {
	((QGraphicsView*)self)->render((QPainter*)painter, *(QRectF*)target, *(QRect*)source);
}
void QGraphicsView_render4(void* self, void* painter, void* target, void* source, int aspectRatioMode) {
	((QGraphicsView*)self)->render((QPainter*)painter, *(QRectF*)target, *(QRect*)source, (Qt::AspectRatioMode)(aspectRatioMode));
}
void** QGraphicsView_items8(void* self, void* rect, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QRect*)rect, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items9(void* self, int x, int y, int w, int h, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(x, y, w, h, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items10(void* self, void* polygon, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QPolygon*)polygon, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsView_items11(void* self, void* path, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsView*)self)->items(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGraphicsView_invalidateScene2(void* self, void* rect) {
	((QGraphicsView*)self)->invalidateScene(*(QRectF*)rect);
}
void QGraphicsView_invalidateScene3(void* self, void* rect, int layers) {
	((QGraphicsView*)self)->invalidateScene(*(QRectF*)rect, QFlags<QGraphicsScene::SceneLayer>(layers));
}
void* QGraphicsView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QGraphicsView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QGraphicsView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QGraphicsView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QGraphicsView_updateMicroFocus2(void* self, int query) {
	((QGraphicsView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QGraphicsView_create2(void* self, unsigned long long arg1) {
	((QGraphicsView__Go*)self)->create(arg1);
}
void QGraphicsView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QGraphicsView__Go*)self)->create(arg1, initializeWindow);
}
void QGraphicsView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QGraphicsView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QGraphicsView_destroy2(void* self, int destroyWindow) {
	((QGraphicsView__Go*)self)->destroy(destroyWindow);
}
void QGraphicsView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QGraphicsView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QGraphicsView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsView_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::metaObject();
	return (void*)(__r);
}
void* QGraphicsView_sizeHint_default(void* self) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::sizeHint();
	return new QSize(__r);
}
void* QGraphicsView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QGraphicsView_setupViewport_default(void* self, void* widget) {
	((QGraphicsView__Go*)self)->cppgoBase_setupViewport((QWidget*)widget);
}
int QGraphicsView_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QGraphicsView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QGraphicsView_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QGraphicsView_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QGraphicsView_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QGraphicsView_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QGraphicsView_dropEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QGraphicsView_focusInEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
int QGraphicsView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QGraphicsView_focusOutEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsView_keyPressEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsView_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QGraphicsView_mousePressEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QGraphicsView_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QGraphicsView_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QGraphicsView_wheelEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QGraphicsView_paintEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QGraphicsView_resizeEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QGraphicsView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QGraphicsView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QGraphicsView_showEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QGraphicsView_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void QGraphicsView_drawBackground_default(void* self, void* painter, void* rect) {
	((QGraphicsView__Go*)self)->cppgoBase_drawBackground((QPainter*)painter, *(QRectF*)rect);
}
void QGraphicsView_drawForeground_default(void* self, void* painter, void* rect) {
	((QGraphicsView__Go*)self)->cppgoBase_drawForeground((QPainter*)painter, *(QRectF*)rect);
}
void* QGraphicsView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::minimumSizeHint();
	return new QSize(__r);
}
int QGraphicsView_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
void* QGraphicsView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QGraphicsView_changeEvent_default(void* self, void* arg1) {
	((QGraphicsView__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QGraphicsView_initStyleOption_default(void* self, void* option) {
	((QGraphicsView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QGraphicsView_devType_default(void* self) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::devType();
	return __r;
}
void QGraphicsView_setVisible_default(void* self, int visible) {
	((QGraphicsView*)self)->QGraphicsView::setVisible(visible);
}
int QGraphicsView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::heightForWidth(arg1);
	return __r;
}
int QGraphicsView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::hasHeightForWidth();
	return __r;
}
void* QGraphicsView_paintEngine_default(void* self) {
	auto&& __r = ((QGraphicsView*)self)->QGraphicsView::paintEngine();
	return (void*)(__r);
}
void QGraphicsView_enterEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QGraphicsView_leaveEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QGraphicsView_moveEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QGraphicsView_closeEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QGraphicsView_tabletEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QGraphicsView_actionEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QGraphicsView_hideEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QGraphicsView_metric_default(void* self, int arg1) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QGraphicsView_initPainter_default(void* self, void* painter) {
	((QGraphicsView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QGraphicsView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QGraphicsView_sharedPainter_default(void* self) {
	auto&& __r = ((QGraphicsView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QGraphicsView_timerEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsView_childEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsView_customEvent_default(void* self, void* event) {
	((QGraphicsView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsView_connectNotify_default(void* self, void* signal) {
	((QGraphicsView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsView_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsView_connect_rubberBandChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::rubberBandChanged, (QGraphicsView*)self, [__id](QRect viewportRect, QPointF fromScenePoint, QPointF toScenePoint){
		if (cppgo_sig_void_voidp_voidp_voidp) ((void(*)(long long, void*, void*, void*))cppgo_sig_void_voidp_voidp_voidp)(__id, (void*)&viewportRect, (void*)&fromScenePoint, (void*)&toScenePoint);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::windowTitleChanged, (QGraphicsView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::windowIconChanged, (QGraphicsView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::windowIconTextChanged, (QGraphicsView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::customContextMenuRequested, (QGraphicsView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::destroyed, (QGraphicsView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::destroyed, (QGraphicsView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsView*)self, &QGraphicsView::objectNameChanged, (QGraphicsView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
	__conns_QGraphicsView[__id] = __c;
}
void QGraphicsView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsView);
		auto __it = __conns_QGraphicsView.find(id);
		if (__it == __conns_QGraphicsView.end()) return;
		__c = __it->second;
		__conns_QGraphicsView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsView__Go_new(long long id) { return new QGraphicsView__Go(id); }
void QGraphicsView_delete(void* self) { delete (QGraphicsView*)self; }

}

#include <QtWidgets/qprogressbar.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
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
#include "gen_qprogressbar.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QProgressBar__Go : public QProgressBar {
public:
	long long __goId;
	unsigned char __goOn[50] = {};
	template<typename... A> QProgressBar__Go(long long id, A&&... a) : QProgressBar(std::forward<A>(a)...), __goId(id) {}
	using QProgressBar::updateMicroFocus;
	using QProgressBar::create;
	using QProgressBar::destroy;
	using QProgressBar::focusNextChild;
	using QProgressBar::focusPreviousChild;
	using QProgressBar::sender;
	using QProgressBar::senderSignalIndex;
	using QProgressBar::receivers;
	using QProgressBar::isSignalConnected;
	using QProgressBar::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QProgressBar::event(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QProgressBar::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionProgressBar* __a0) { QProgressBar::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QProgressBar::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QProgressBar::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QProgressBar::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QProgressBar::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QProgressBar::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QProgressBar::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QProgressBar::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QProgressBar::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QProgressBar::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QProgressBar::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QProgressBar::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QProgressBar::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QProgressBar::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QProgressBar::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QProgressBar::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QProgressBar::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QProgressBar::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QProgressBar::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QProgressBar::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QProgressBar::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QProgressBar::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QProgressBar::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QProgressBar::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QProgressBar::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QProgressBar::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QProgressBar::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QProgressBar::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QProgressBar::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QProgressBar::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QProgressBar::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QProgressBar::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QProgressBar::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QProgressBar::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QProgressBar::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QProgressBar::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QProgressBar::metaObject();
	}
	QString text() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, -430950356557152130LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QProgressBar::text();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QProgressBar::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QProgressBar::minimumSizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QProgressBar::event(e);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QProgressBar::paintEvent(arg1);
	}
	void initStyleOption(QStyleOptionProgressBar* option) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QProgressBar::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QProgressBar::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QProgressBar::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QProgressBar::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QProgressBar::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QProgressBar::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QProgressBar::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QProgressBar::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QProgressBar::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QProgressBar::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QProgressBar::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QProgressBar::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QProgressBar::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QProgressBar::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QProgressBar::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QProgressBar::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QProgressBar::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QProgressBar::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QProgressBar::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QProgressBar;
static std::mutex __connsMu_QProgressBar;

extern "C" {

void* QProgressBar_base(void* self) { return (void*)static_cast<QWidget*>((QProgressBar*)self); }
void QProgressBar_chain(void* self, void** out) {
	QProgressBar* __p = (QProgressBar*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QProgressBar__set_override(void* self, int idx) { ((QProgressBar__Go*)self)->__goOn[idx] = 1; }
void* QProgressBar_fromQObject(void* o) { return (void*)qobject_cast<QProgressBar*>((QObject*)o); }
void* QProgressBar_new(long long __goid, void* parent) {
	return new QProgressBar__Go(__goid, (QWidget*)parent);
}
void* QProgressBar_new2(long long __goid) {
	return new QProgressBar__Go(__goid);
}
char* QProgressBar_tr(const char* s, long long* __out_len) {
	auto&& __r = QProgressBar::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QProgressBar_minimum(void* self) {
	auto&& __r = ((QProgressBar*)self)->minimum();
	return __r;
}
int QProgressBar_maximum(void* self) {
	auto&& __r = ((QProgressBar*)self)->maximum();
	return __r;
}
int QProgressBar_value(void* self) {
	auto&& __r = ((QProgressBar*)self)->value();
	return __r;
}
char* QProgressBar_text(void* self, long long* __out_len) {
	auto&& __r = ((QProgressBar*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProgressBar_setTextVisible(void* self, int visible) {
	((QProgressBar*)self)->setTextVisible(visible);
}
int QProgressBar_isTextVisible(void* self) {
	auto&& __r = ((QProgressBar*)self)->isTextVisible();
	return __r;
}
int QProgressBar_alignment(void* self) {
	auto&& __r = ((QProgressBar*)self)->alignment();
	return (int)(__r);
}
void QProgressBar_setAlignment(void* self, int alignment) {
	((QProgressBar*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QProgressBar_orientation(void* self) {
	auto&& __r = ((QProgressBar*)self)->orientation();
	return (int)(__r);
}
void QProgressBar_setInvertedAppearance(void* self, int invert) {
	((QProgressBar*)self)->setInvertedAppearance(invert);
}
int QProgressBar_invertedAppearance(void* self) {
	auto&& __r = ((QProgressBar*)self)->invertedAppearance();
	return __r;
}
void QProgressBar_setTextDirection(void* self, int textDirection) {
	((QProgressBar*)self)->setTextDirection((QProgressBar::Direction)(textDirection));
}
int QProgressBar_textDirection(void* self) {
	auto&& __r = ((QProgressBar*)self)->textDirection();
	return (int)(__r);
}
void QProgressBar_setFormat(void* self, const char* format, long long format_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	((QProgressBar*)self)->setFormat(format_qstr);
}
void QProgressBar_resetFormat(void* self) {
	((QProgressBar*)self)->resetFormat();
}
char* QProgressBar_format(void* self, long long* __out_len) {
	auto&& __r = ((QProgressBar*)self)->format();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProgressBar_reset(void* self) {
	((QProgressBar*)self)->reset();
}
void QProgressBar_setRange(void* self, int minimum, int maximum) {
	((QProgressBar*)self)->setRange(minimum, maximum);
}
void QProgressBar_setMinimum(void* self, int minimum) {
	((QProgressBar*)self)->setMinimum(minimum);
}
void QProgressBar_setMaximum(void* self, int maximum) {
	((QProgressBar*)self)->setMaximum(maximum);
}
void QProgressBar_setValue(void* self, int value) {
	((QProgressBar*)self)->setValue(value);
}
void QProgressBar_setOrientation(void* self, int arg1) {
	((QProgressBar*)self)->setOrientation((Qt::Orientation)(arg1));
}
void QProgressBar_valueChanged(void* self, int value) {
	((QProgressBar*)self)->valueChanged(value);
}
void QProgressBar_setTabOrder(void* arg1, void* arg2) {
	QProgressBar::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QProgressBar_mouseGrabber() {
	auto&& __r = QProgressBar::mouseGrabber();
	return (void*)(__r);
}
void* QProgressBar_keyboardGrabber() {
	auto&& __r = QProgressBar::keyboardGrabber();
	return (void*)(__r);
}
void* QProgressBar_find(unsigned long long arg1) {
	auto&& __r = QProgressBar::find(arg1);
	return (void*)(__r);
}
void* QProgressBar_createWindowContainer(void* window) {
	auto&& __r = QProgressBar::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QProgressBar_updateMicroFocus(void* self) {
	((QProgressBar__Go*)self)->updateMicroFocus();
}
void QProgressBar_create(void* self) {
	((QProgressBar__Go*)self)->create();
}
void QProgressBar_destroy(void* self) {
	((QProgressBar__Go*)self)->destroy();
}
int QProgressBar_focusNextChild(void* self) {
	auto&& __r = ((QProgressBar__Go*)self)->focusNextChild();
	return __r;
}
int QProgressBar_focusPreviousChild(void* self) {
	auto&& __r = ((QProgressBar__Go*)self)->focusPreviousChild();
	return __r;
}
void* QProgressBar_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QProgressBar::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QProgressBar_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QProgressBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QProgressBar_sender(void* self) {
	auto&& __r = ((QProgressBar__Go*)self)->sender();
	return (void*)(__r);
}
int QProgressBar_senderSignalIndex(void* self) {
	auto&& __r = ((QProgressBar__Go*)self)->senderSignalIndex();
	return __r;
}
int QProgressBar_receivers(void* self, const char* signal) {
	auto&& __r = ((QProgressBar__Go*)self)->receivers(signal);
	return __r;
}
int QProgressBar_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QProgressBar__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QProgressBar_devicePixelRatioFScale() {
	auto&& __r = QProgressBar::devicePixelRatioFScale();
	return __r;
}
int QProgressBar_encodeMetricF(int metric, double value) {
	auto&& __r = QProgressBar::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QProgressBar_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QProgressBar__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QProgressBar_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QProgressBar::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QProgressBar_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QProgressBar::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QProgressBar_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QProgressBar::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QProgressBar_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QProgressBar::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QProgressBar_updateMicroFocus2(void* self, int query) {
	((QProgressBar__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QProgressBar_create2(void* self, unsigned long long arg1) {
	((QProgressBar__Go*)self)->create(arg1);
}
void QProgressBar_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QProgressBar__Go*)self)->create(arg1, initializeWindow);
}
void QProgressBar_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QProgressBar__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QProgressBar_destroy2(void* self, int destroyWindow) {
	((QProgressBar__Go*)self)->destroy(destroyWindow);
}
void QProgressBar_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QProgressBar__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QProgressBar_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QProgressBar::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QProgressBar_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QProgressBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
char* QProgressBar_text_default(void* self, long long* __out_len) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QProgressBar_metaObject_default(void* self) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::metaObject();
	return (void*)(__r);
}
void* QProgressBar_sizeHint_default(void* self) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::sizeHint();
	return new QSize(__r);
}
void* QProgressBar_minimumSizeHint_default(void* self) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::minimumSizeHint();
	return new QSize(__r);
}
int QProgressBar_event_default(void* self, void* e) {
	auto&& __r = ((QProgressBar__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QProgressBar_paintEvent_default(void* self, void* arg1) {
	((QProgressBar__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QProgressBar_initStyleOption_default(void* self, void* option) {
	((QProgressBar__Go*)self)->cppgoBase_initStyleOption((QStyleOptionProgressBar*)option);
}
int QProgressBar_devType_default(void* self) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::devType();
	return __r;
}
void QProgressBar_setVisible_default(void* self, int visible) {
	((QProgressBar*)self)->QProgressBar::setVisible(visible);
}
int QProgressBar_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::heightForWidth(arg1);
	return __r;
}
int QProgressBar_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::hasHeightForWidth();
	return __r;
}
void* QProgressBar_paintEngine_default(void* self) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::paintEngine();
	return (void*)(__r);
}
void QProgressBar_mousePressEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QProgressBar_mouseReleaseEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QProgressBar_mouseDoubleClickEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QProgressBar_mouseMoveEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QProgressBar_wheelEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QProgressBar_keyPressEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QProgressBar_keyReleaseEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QProgressBar_focusInEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QProgressBar_focusOutEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QProgressBar_enterEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QProgressBar_leaveEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QProgressBar_moveEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QProgressBar_resizeEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QProgressBar_closeEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QProgressBar_contextMenuEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QProgressBar_tabletEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QProgressBar_actionEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QProgressBar_dragEnterEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QProgressBar_dragMoveEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QProgressBar_dragLeaveEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QProgressBar_dropEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QProgressBar_showEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QProgressBar_hideEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QProgressBar_changeEvent_default(void* self, void* arg1) {
	((QProgressBar__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QProgressBar_metric_default(void* self, int arg1) {
	auto&& __r = ((QProgressBar__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QProgressBar_initPainter_default(void* self, void* painter) {
	((QProgressBar__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QProgressBar_redirected_default(void* self, void* offset) {
	auto&& __r = ((QProgressBar__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QProgressBar_sharedPainter_default(void* self) {
	auto&& __r = ((QProgressBar__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QProgressBar_inputMethodEvent_default(void* self, void* arg1) {
	((QProgressBar__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QProgressBar_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QProgressBar_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QProgressBar__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QProgressBar_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QProgressBar*)self)->QProgressBar::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QProgressBar_timerEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QProgressBar_childEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QProgressBar_customEvent_default(void* self, void* event) {
	((QProgressBar__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QProgressBar_connectNotify_default(void* self, void* signal) {
	((QProgressBar__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QProgressBar_disconnectNotify_default(void* self, void* signal) {
	((QProgressBar__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QProgressBar_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::valueChanged, (QProgressBar*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::windowTitleChanged, (QProgressBar*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::windowIconChanged, (QProgressBar*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::windowIconTextChanged, (QProgressBar*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::customContextMenuRequested, (QProgressBar*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::destroyed, (QProgressBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::destroyed, (QProgressBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressBar*)self, &QProgressBar::objectNameChanged, (QProgressBar*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
	__conns_QProgressBar[__id] = __c;
}
void QProgressBar__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QProgressBar);
		auto __it = __conns_QProgressBar.find(id);
		if (__it == __conns_QProgressBar.end()) return;
		__c = __it->second;
		__conns_QProgressBar.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QProgressBar__Go_new(long long id) { return new QProgressBar__Go(id); }
void QProgressBar_delete(void* self) { delete (QProgressBar*)self; }

}

#include <QtWidgets/qtablewidget.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qfont.h>
#include <QtGui/qbrush.h>
#include <QtCore/qsize.h>
#include <QtCore/qvariant.h>
#include <QtCore/qdatastream.h>
#include "gen_qtablewidgetitem.h"
#include "gen_purego_cb.h"

class QTableWidgetItem__Go : public QTableWidgetItem {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QTableWidgetItem__Go(long long id, A&&... a) : QTableWidgetItem(std::forward<A>(a)...), __goId(id) {}
	QTableWidgetItem* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QTableWidgetItem*)__r; }
		return QTableWidgetItem::clone();
	}
	QVariant data(int role) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -8837375942098466299LL, role, &__h); if (__h) return *(QVariant*)__r; }
		return QTableWidgetItem::data(role);
	}
	void setData(int role, const QVariant& value) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, 35191454575202219LL, role, (void*)&value, &__h); if (__h) return; }
		QTableWidgetItem::setData(role, value);
	}
	void read(QDataStream& in) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5541207787354052101LL, (void*)&in, &__h); if (__h) return; }
		QTableWidgetItem::read(in);
	}
	void write(QDataStream& out) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5099743077423337732LL, (void*)&out, &__h); if (__h) return; }
		QTableWidgetItem::write(out);
	}
};

extern "C" {

void QTableWidgetItem__set_override(void* self, int idx) { ((QTableWidgetItem__Go*)self)->__goOn[idx] = 1; }
void* QTableWidgetItem_new(long long __goid) {
	return new QTableWidgetItem__Go(__goid);
}
void* QTableWidgetItem_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTableWidgetItem__Go(__goid, text_qstr);
}
void* QTableWidgetItem_new3(long long __goid, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTableWidgetItem__Go(__goid, *(QIcon*)icon, text_qstr);
}
void* QTableWidgetItem_new4(long long __goid, void* other) {
	return new QTableWidgetItem__Go(__goid, *(QTableWidgetItem*)other);
}
void* QTableWidgetItem_new5(long long __goid, int type_) {
	return new QTableWidgetItem__Go(__goid, type_);
}
void* QTableWidgetItem_new6(long long __goid, const char* text, long long text_len, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTableWidgetItem__Go(__goid, text_qstr, type_);
}
void* QTableWidgetItem_new7(long long __goid, void* icon, const char* text, long long text_len, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTableWidgetItem__Go(__goid, *(QIcon*)icon, text_qstr, type_);
}
void* QTableWidgetItem_clone(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->clone();
	return (void*)(__r);
}
void* QTableWidgetItem_tableWidget(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->tableWidget();
	return (void*)(__r);
}
int QTableWidgetItem_row(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->row();
	return __r;
}
int QTableWidgetItem_column(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->column();
	return __r;
}
void QTableWidgetItem_setSelected(void* self, int select_) {
	((QTableWidgetItem*)self)->setSelected(select_);
}
int QTableWidgetItem_isSelected(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->isSelected();
	return __r;
}
int QTableWidgetItem_flags(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->flags();
	return (int)(__r);
}
void QTableWidgetItem_setFlags(void* self, int flags) {
	((QTableWidgetItem*)self)->setFlags(QFlags<Qt::ItemFlag>(flags));
}
char* QTableWidgetItem_text(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidgetItem*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTableWidgetItem_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTableWidgetItem*)self)->setText(text_qstr);
}
void* QTableWidgetItem_icon(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->icon();
	return new QIcon(__r);
}
void QTableWidgetItem_setIcon(void* self, void* icon) {
	((QTableWidgetItem*)self)->setIcon(*(QIcon*)icon);
}
char* QTableWidgetItem_statusTip(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidgetItem*)self)->statusTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTableWidgetItem_setStatusTip(void* self, const char* statusTip, long long statusTip_len) {
	QString statusTip_qstr = QString::fromUtf8(statusTip, statusTip_len);
	((QTableWidgetItem*)self)->setStatusTip(statusTip_qstr);
}
char* QTableWidgetItem_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidgetItem*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTableWidgetItem_setToolTip(void* self, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QTableWidgetItem*)self)->setToolTip(toolTip_qstr);
}
char* QTableWidgetItem_whatsThis(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidgetItem*)self)->whatsThis();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTableWidgetItem_setWhatsThis(void* self, const char* whatsThis, long long whatsThis_len) {
	QString whatsThis_qstr = QString::fromUtf8(whatsThis, whatsThis_len);
	((QTableWidgetItem*)self)->setWhatsThis(whatsThis_qstr);
}
void* QTableWidgetItem_font(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->font();
	return new QFont(__r);
}
void QTableWidgetItem_setFont(void* self, void* font) {
	((QTableWidgetItem*)self)->setFont(*(QFont*)font);
}
int QTableWidgetItem_textAlignment(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->textAlignment();
	return __r;
}
void QTableWidgetItem_setTextAlignment(void* self, int alignment) {
	((QTableWidgetItem*)self)->setTextAlignment(alignment);
}
void QTableWidgetItem_setTextAlignment2(void* self, int alignment) {
	((QTableWidgetItem*)self)->setTextAlignment((Qt::AlignmentFlag)(alignment));
}
void QTableWidgetItem_setTextAlignment3(void* self, int alignment) {
	((QTableWidgetItem*)self)->setTextAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
void* QTableWidgetItem_background(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->background();
	return new QBrush(__r);
}
void QTableWidgetItem_setBackground(void* self, void* brush) {
	((QTableWidgetItem*)self)->setBackground(*(QBrush*)brush);
}
void* QTableWidgetItem_foreground(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->foreground();
	return new QBrush(__r);
}
void QTableWidgetItem_setForeground(void* self, void* brush) {
	((QTableWidgetItem*)self)->setForeground(*(QBrush*)brush);
}
int QTableWidgetItem_checkState(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->checkState();
	return (int)(__r);
}
void QTableWidgetItem_setCheckState(void* self, int state) {
	((QTableWidgetItem*)self)->setCheckState((Qt::CheckState)(state));
}
void* QTableWidgetItem_sizeHint(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->sizeHint();
	return new QSize(__r);
}
void QTableWidgetItem_setSizeHint(void* self, void* size) {
	((QTableWidgetItem*)self)->setSizeHint(*(QSize*)size);
}
void* QTableWidgetItem_data(void* self, int role) {
	auto&& __r = ((QTableWidgetItem*)self)->data(role);
	return new QVariant(__r);
}
void QTableWidgetItem_setData(void* self, int role, void* value) {
	((QTableWidgetItem*)self)->setData(role, *(QVariant*)value);
}
int QTableWidgetItem_OperatorLesser(void* self, void* other) {
	auto&& __r = ((QTableWidgetItem*)self)->operator<(*(QTableWidgetItem*)other);
	return __r;
}
void QTableWidgetItem_read(void* self, void* in) {
	((QTableWidgetItem*)self)->read(*(QDataStream*)in);
}
void QTableWidgetItem_write(void* self, void* out) {
	((QTableWidgetItem*)self)->write(*(QDataStream*)out);
}
void* QTableWidgetItem_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QTableWidgetItem*)self)->operator=(*(QTableWidgetItem*)other);
	return (void*)&(__r);
}
int QTableWidgetItem_type(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->type();
	return __r;
}
void* QTableWidgetItem_clone_default(void* self) {
	auto&& __r = ((QTableWidgetItem*)self)->QTableWidgetItem::clone();
	return (void*)(__r);
}
void* QTableWidgetItem_data_default(void* self, int role) {
	auto&& __r = ((QTableWidgetItem*)self)->QTableWidgetItem::data(role);
	return new QVariant(__r);
}
void QTableWidgetItem_setData_default(void* self, int role, void* value) {
	((QTableWidgetItem*)self)->QTableWidgetItem::setData(role, *(QVariant*)value);
}
void QTableWidgetItem_read_default(void* self, void* in) {
	((QTableWidgetItem*)self)->QTableWidgetItem::read(*(QDataStream*)in);
}
void QTableWidgetItem_write_default(void* self, void* out) {
	((QTableWidgetItem*)self)->QTableWidgetItem::write(*(QDataStream*)out);
}
void* QTableWidgetItem__Go_new(long long id) { return new QTableWidgetItem__Go(id); }
void QTableWidgetItem_delete(void* self) { delete (QTableWidgetItem*)self; }

}

