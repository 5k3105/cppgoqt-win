#include <QtCore/qobjectdefs.h>
#include "gen_qmetaobject__data.h"

extern "C" {

void* QMetaObject__Data_new() {
	return new QMetaObject::Data();
}
void* QMetaObject__Data_new2(void* arg1) {
	return new QMetaObject::Data(*(QMetaObject::Data*)arg1);
}
void* QMetaObject__Data_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QMetaObject::Data*)self)->operator=(*(QMetaObject::Data*)arg1);
	return (void*)&(__r);
}
void* QMetaObject__Data_superdata_get(void* self) {
	auto __r = ((QMetaObject::Data*)self)->superdata;
	return new QMetaObject::SuperData(__r);
}
void QMetaObject__Data_superdata_set(void* self, void* superdata) {
	((QMetaObject::Data*)self)->superdata = *(QMetaObject::SuperData*)superdata;
}
void* QMetaObject__Data_relatedMetaObjects_get(void* self) {
	auto __r = ((QMetaObject::Data*)self)->relatedMetaObjects;
	return (void*)(__r);
}
void QMetaObject__Data_relatedMetaObjects_set(void* self, void* relatedMetaObjects) {
	((QMetaObject::Data*)self)->relatedMetaObjects = (QMetaObject::SuperData*)relatedMetaObjects;
}
void QMetaObject__Data_delete(void* self) { delete (QMetaObject::Data*)self; }

}

#include <QtCore/qmetacontainer.h>
#include <QtCore/qmetatype.h>
#include "gen_qmetasequence.h"

extern "C" {

void* QMetaSequence_base(void* self) { return (void*)static_cast<QMetaContainer*>((QMetaSequence*)self); }
void* QMetaSequence_new() {
	return new QMetaSequence();
}
void* QMetaSequence_new2(void* arg1) {
	return new QMetaSequence(*(QMetaSequence*)arg1);
}
void* QMetaSequence_valueMetaType(void* self) {
	auto&& __r = ((QMetaSequence*)self)->valueMetaType();
	return new QMetaType(__r);
}
int QMetaSequence_isSortable(void* self) {
	auto&& __r = ((QMetaSequence*)self)->isSortable();
	return __r;
}
int QMetaSequence_canAddValueAtBegin(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canAddValueAtBegin();
	return __r;
}
int QMetaSequence_canAddValueAtEnd(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canAddValueAtEnd();
	return __r;
}
int QMetaSequence_canRemoveValueAtBegin(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canRemoveValueAtBegin();
	return __r;
}
int QMetaSequence_canRemoveValueAtEnd(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canRemoveValueAtEnd();
	return __r;
}
int QMetaSequence_canGetValueAtIndex(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canGetValueAtIndex();
	return __r;
}
int QMetaSequence_canSetValueAtIndex(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canSetValueAtIndex();
	return __r;
}
int QMetaSequence_canAddValue(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canAddValue();
	return __r;
}
int QMetaSequence_canRemoveValue(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canRemoveValue();
	return __r;
}
int QMetaSequence_canGetValueAtIterator(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canGetValueAtIterator();
	return __r;
}
int QMetaSequence_canSetValueAtIterator(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canSetValueAtIterator();
	return __r;
}
int QMetaSequence_canInsertValueAtIterator(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canInsertValueAtIterator();
	return __r;
}
int QMetaSequence_canEraseValueAtIterator(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canEraseValueAtIterator();
	return __r;
}
int QMetaSequence_canEraseRangeAtIterator(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canEraseRangeAtIterator();
	return __r;
}
int QMetaSequence_canGetValueAtConstIterator(void* self) {
	auto&& __r = ((QMetaSequence*)self)->canGetValueAtConstIterator();
	return __r;
}
int QMetaSequence_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaSequence*)self) == (*(QMetaSequence*)rhs);
	return __r;
}
int QMetaSequence_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMetaSequence*)self) != (*(QMetaSequence*)rhs);
	return __r;
}
void QMetaSequence_delete(void* self) { delete (QMetaSequence*)self; }

}

#include <QtCore/qdebug.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qlogging.h>
#include <QtCore/qchar.h>
#include <QtCore/qfloat16.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qtextstream.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qobject.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qabstractitemmodel.h>
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
#include <QtCore/qfileinfo.h>
#include <QtCore/qdir.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qjsondocument.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qtyperevision.h>
#include <QtCore/qversionnumber.h>
#include <QtCore/qpoint.h>
#include <QtCore/qline.h>
#include <QtCore/qmargins.h>
#include <QtCore/qmimetype.h>
#include <QtCore/qoperatingsystemversion.h>
#include <QtCore/qpermissions.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qstorageinfo.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpixelformat.h>
#include <QtGui/qpolygon.h>
#include <QtGui/qregion.h>
#include <QtGui/qtransform.h>
#include <QtGui/qimage.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qicon.h>
#include <QtGui/qvectornd.h>
#include <QtGui/qscreen.h>
#include <QtGui/qinputdevice.h>
#include <QtGui/qpointingdevice.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qkeysequence.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qfont.h>
#include <QtGui/qbrush.h>
#include <QtGui/qpen.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qpalette.h>
#include <QtGui/qaccessible.h>
#include <QtGui/qaction.h>
#include <QtGui/qsurfaceformat.h>
#include <QtGui/qcursor.h>
#include <QtGui/qwindow.h>
#include <QtGui/qcolorspace.h>
#include <QtGui/qfontvariableaxis.h>
#include <QtGui/qmatrix4x4.h>
#include <QtGui/qopenglcontext.h>
#include <QtGui/qpagesize.h>
#include <QtGui/qpagelayout.h>
#include <QtGui/qpageranges.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qquaternion.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qdockwidget.h>
#include <QtWidgets/qgesture.h>
#include <QtWidgets/qgraphicsitem.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include "gen_qdebug.h"

extern "C" {

void* QDebug_base(void* self) { return (void*)static_cast<QIODeviceBase*>((QDebug*)self); }
void* QDebug_new(void* device) {
	return new QDebug((QIODevice*)device);
}
void* QDebug_new2(void* string_) {
	return new QDebug((QString*)string_);
}
void* QDebug_new3(void* bytes) {
	return new QDebug((QByteArray*)bytes);
}
void* QDebug_new4(int t) {
	return new QDebug((QtMsgType)(t));
}
void* QDebug_new5(void* o) {
	return new QDebug(*(QDebug*)o);
}
void* QDebug_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QDebug*)self)->operator=(*(QDebug*)other);
	return (void*)&(__r);
}
void QDebug_swap(void* self, void* other) {
	((QDebug*)self)->swap(*(QDebug*)other);
}
void* QDebug_resetFormat(void* self) {
	auto&& __r = ((QDebug*)self)->resetFormat();
	return (void*)&(__r);
}
void* QDebug_space(void* self) {
	auto&& __r = ((QDebug*)self)->space();
	return (void*)&(__r);
}
void* QDebug_nospace(void* self) {
	auto&& __r = ((QDebug*)self)->nospace();
	return (void*)&(__r);
}
void* QDebug_maybeSpace(void* self) {
	auto&& __r = ((QDebug*)self)->maybeSpace();
	return (void*)&(__r);
}
void* QDebug_verbosity(void* self, int verbosityLevel) {
	auto&& __r = ((QDebug*)self)->verbosity(verbosityLevel);
	return (void*)&(__r);
}
int QDebug_verbosity2(void* self) {
	auto&& __r = ((QDebug*)self)->verbosity();
	return __r;
}
void QDebug_setVerbosity(void* self, int verbosityLevel) {
	((QDebug*)self)->setVerbosity(verbosityLevel);
}
int QDebug_autoInsertSpaces(void* self) {
	auto&& __r = ((QDebug*)self)->autoInsertSpaces();
	return __r;
}
void QDebug_setAutoInsertSpaces(void* self, int b) {
	((QDebug*)self)->setAutoInsertSpaces(b);
}
int QDebug_quoteStrings(void* self) {
	auto&& __r = ((QDebug*)self)->quoteStrings();
	return __r;
}
void QDebug_setQuoteStrings(void* self, int b) {
	((QDebug*)self)->setQuoteStrings(b);
}
void* QDebug_quote(void* self) {
	auto&& __r = ((QDebug*)self)->quote();
	return (void*)&(__r);
}
void* QDebug_noquote(void* self) {
	auto&& __r = ((QDebug*)self)->noquote();
	return (void*)&(__r);
}
void* QDebug_maybeQuote(void* self) {
	auto&& __r = ((QDebug*)self)->maybeQuote();
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft(void* self, void* t) {
	auto&& __r = ((QDebug*)self)->operator<<(*(QChar*)t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft2(void* self, int t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft3(void* self, char t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft4(void* self, short t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft5(void* self, unsigned short t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft6(void* self, unsigned short t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft7(void* self, unsigned int t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft8(void* self, int t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft9(void* self, unsigned int t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft10(void* self, long t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft11(void* self, unsigned long t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft12(void* self, long long t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft13(void* self, unsigned long long t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft14(void* self, void* t) {
	auto&& __r = ((QDebug*)self)->operator<<(*(qfloat16*)t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft15(void* self, float t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft16(void* self, double t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft17(void* self, const char* t) {
	auto&& __r = ((QDebug*)self)->operator<<(t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft18(void* self, const char* t, long long t_len) {
	QString t_qstr = QString::fromUtf8(t, t_len);
	auto&& __r = ((QDebug*)self)->operator<<(t_qstr);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft19(void* self, void* s) {
	auto&& __r = ((QDebug*)self)->operator<<(*(QStringView*)s);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft20(void* self, void* t) {
	auto&& __r = ((QDebug*)self)->operator<<(*(QLatin1String*)t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft21(void* self, void* t) {
	auto&& __r = ((QDebug*)self)->operator<<(*(QByteArray*)t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft22(void* self, void* t) {
	auto&& __r = ((QDebug*)self)->operator<<(*(QByteArrayView*)t);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft23(void* self, void* m) {
	auto&& __r = ((QDebug*)self)->operator<<(*(QTextStreamManipulator*)m);
	return (void*)&(__r);
}
void* QDebug_maybeQuote2(void* self, char c) {
	auto&& __r = ((QDebug*)self)->maybeQuote(c);
	return (void*)&(__r);
}
void* QDebug_OperatorShiftLeft24(void* self, const char* s, long long s_len) {
	QAnyStringView s_qstr = QAnyStringView(s, s_len);
	auto&& __r = (*(QDebug*)self) << (s_qstr);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft25(void* self, void* m) {
	auto&& __r = (*(QDebug*)self) << (*(QMetaType*)m);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft26(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QObject*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft27(void* self, void* combination) {
	auto&& __r = (*(QDebug*)self) << (*(QKeyCombination*)combination);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft28(void* self, int arg1) {
	auto&& __r = (*(QDebug*)self) << ((QVariant::Type)(arg1));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft29(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QModelIndex*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft30(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPersistentModelIndex*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft31(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QBitArray*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft32(void* self, int modes) {
	auto&& __r = (*(QDebug*)self) << (QFlags<QIODeviceBase::OpenModeFlag>(modes));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft33(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QLocale*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft34(void* self, int st) {
	auto&& __r = (*(QDebug*)self) << ((QCborSimpleType)(st));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft35(void* self, int tg) {
	auto&& __r = (*(QDebug*)self) << ((QCborKnownTags)(tg));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft36(void* self, int tg) {
	auto&& __r = (*(QDebug*)self) << ((QCborTag)(tg));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft37(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QDateTime*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft38(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QDate*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft39(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QTime*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft40(void* self, void* re) {
	auto&& __r = (*(QDebug*)self) << (*(QRegularExpression*)re);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft41(void* self, int patternOptions) {
	auto&& __r = (*(QDebug*)self) << (QFlags<QRegularExpression::PatternOption>(patternOptions));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft42(void* self, void* match) {
	auto&& __r = (*(QDebug*)self) << (*(QRegularExpressionMatch*)match);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft43(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QUrl*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft44(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QUuid*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft45(void* self, void* v) {
	auto&& __r = (*(QDebug*)self) << (*(QCborValue*)v);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft46(void* self, void* a) {
	auto&& __r = (*(QDebug*)self) << (*(QCborArray*)a);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft47(void* self, void* m) {
	auto&& __r = (*(QDebug*)self) << (*(QCborMap*)m);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft48(void* self, void* tz) {
	auto&& __r = (*(QDebug*)self) << (*(QTimeZone*)tz);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft49(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QFileInfo*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft50(void* self, int filters) {
	auto&& __r = (*(QDebug*)self) << (QFlags<QDir::Filter>(filters));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft51(void* self, void* dir) {
	auto&& __r = (*(QDebug*)self) << (*(QDir*)dir);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft52(void* self, void* item) {
	auto&& __r = (*(QDebug*)self) << (*(QEasingCurve*)item);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft53(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QItemSelectionRange*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft54(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QJsonDocument*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft55(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QJsonValue*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft56(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QJsonArray*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft57(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QJsonObject*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft58(void* self, void* revision) {
	auto&& __r = (*(QDebug*)self) << (*(QTypeRevision*)revision);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft59(void* self, void* version) {
	auto&& __r = (*(QDebug*)self) << (*(QVersionNumber*)version);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft60(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPoint*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft61(void* self, void* p) {
	auto&& __r = (*(QDebug*)self) << (*(QPointF*)p);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft62(void* self, void* p) {
	auto&& __r = (*(QDebug*)self) << (*(QLine*)p);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft63(void* self, void* p) {
	auto&& __r = (*(QDebug*)self) << (*(QLineF*)p);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft64(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QMargins*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft65(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QMarginsF*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft66(void* self, void* mime) {
	auto&& __r = (*(QDebug*)self) << (*(QMimeType*)mime);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft67(void* self, void* ov) {
	auto&& __r = (*(QDebug*)self) << (*(QOperatingSystemVersion*)ov);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft68(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPermission*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft69(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QSize*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft70(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QSizeF*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft71(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QRect*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft72(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QRectF*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft73(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QStorageInfo*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft74(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QColor*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft75(void* self, void* format) {
	auto&& __r = (*(QDebug*)self) << (*(QPixelFormat*)format);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft76(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPolygon*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft77(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPolygonF*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft78(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QRegion*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft79(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QTransform*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft80(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QImage*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft81(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPixmap*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft82(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QIcon*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft83(void* self, void* vector) {
	auto&& __r = (*(QDebug*)self) << (*(QVector2D*)vector);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft84(void* self, void* vector) {
	auto&& __r = (*(QDebug*)self) << (*(QVector3D*)vector);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft85(void* self, void* vector) {
	auto&& __r = (*(QDebug*)self) << (*(QVector4D*)vector);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft86(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QScreen*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft87(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QInputDevice*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft88(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QPointingDevice*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft89(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QEventPoint*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft90(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QEventPoint*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft91(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QKeySequence*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft92(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QEvent*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft93(void* self, void* tag) {
	auto&& __r = (*(QDebug*)self) << (*(QFont::Tag*)tag);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft94(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QFont*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft95(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QBrush*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft96(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPen*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft97(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QTextLength*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft98(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QTextFormat*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft99(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPalette*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft100(void* self, void* iface) {
	auto&& __r = (*(QDebug*)self) << ((QAccessibleInterface*)iface);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft101(void* self, void* ev) {
	auto&& __r = (*(QDebug*)self) << (*(QAccessibleEvent*)ev);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft102(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QAction*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft103(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QSurfaceFormat*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft104(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QCursor*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft105(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QWindow*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft106(void* self, void* colorSpace) {
	auto&& __r = (*(QDebug*)self) << (*(QColorSpace*)colorSpace);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft107(void* self, void* axis) {
	auto&& __r = (*(QDebug*)self) << (*(QFontVariableAxis*)axis);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft108(void* self, void* m) {
	auto&& __r = (*(QDebug*)self) << (*(QMatrix4x4*)m);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft109(void* self, void* ctx) {
	auto&& __r = (*(QDebug*)self) << ((QOpenGLContext*)ctx);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft110(void* self, void* cg) {
	auto&& __r = (*(QDebug*)self) << ((QOpenGLContextGroup*)cg);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft111(void* self, void* pageSize) {
	auto&& __r = (*(QDebug*)self) << (*(QPageSize*)pageSize);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft112(void* self, void* pageLayout) {
	auto&& __r = (*(QDebug*)self) << (*(QPageLayout*)pageLayout);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft113(void* self, void* pageRanges) {
	auto&& __r = (*(QDebug*)self) << (*(QPageRanges*)pageRanges);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft114(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QPainterPath*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft115(void* self, void* q) {
	auto&& __r = (*(QDebug*)self) << (*(QQuaternion*)q);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft116(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << (*(QSizePolicy*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft117(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QWidget*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft118(void* self, void* option) {
	auto&& __r = (*(QDebug*)self) << (*(QStyleOption*)option);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft119(void* self, void* dockWidget) {
	auto&& __r = (*(QDebug*)self) << ((QDockWidget*)dockWidget);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft120(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QGesture*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft121(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QGestureEvent*)arg1);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft122(void* self, void* item) {
	auto&& __r = (*(QDebug*)self) << ((QGraphicsItem*)item);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft123(void* self, void* item) {
	auto&& __r = (*(QDebug*)self) << ((QGraphicsObject*)item);
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft124(void* self, int change) {
	auto&& __r = (*(QDebug*)self) << ((QGraphicsItem::GraphicsItemChange)(change));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft125(void* self, int flag) {
	auto&& __r = (*(QDebug*)self) << ((QGraphicsItem::GraphicsItemFlag)(flag));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft126(void* self, int flags) {
	auto&& __r = (*(QDebug*)self) << (QFlags<QGraphicsItem::GraphicsItemFlag>(flags));
	return new QDebug(__r);
}
void* QDebug_OperatorShiftLeft127(void* self, void* arg1) {
	auto&& __r = (*(QDebug*)self) << ((QGraphicsSceneEvent*)arg1);
	return new QDebug(__r);
}
void QDebug_delete(void* self) { delete (QDebug*)self; }

}

#include <QtCore/qassociativeiterable.h>
#include <QtCore/qvariant.h>
#include <QtCore/qmetacontainer.h>
#include "gen_qassociativeiterator.h"

extern "C" {

void* QAssociativeIterator_new(void* arg1) {
	return new QAssociativeIterator(*(QAssociativeIterator*)arg1);
}
void* QAssociativeIterator_key(void* self) {
	auto&& __r = ((QAssociativeIterator*)self)->key();
	return new QVariant(__r);
}
void* QAssociativeIterator_metaContainer(void* self) {
	auto&& __r = ((QAssociativeIterator*)self)->metaContainer();
	return new QMetaAssociation(__r);
}
void QAssociativeIterator_delete(void* self) { delete (QAssociativeIterator*)self; }

}

#include <QtCore/qcborvalue.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qcborarray.h>
#include <QtCore/qcbormap.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qurl.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/quuid.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qvariant.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qcborstreamreader.h>
#include <QtCore/qcborstreamwriter.h>
#include "gen_qcborvalue.h"

extern "C" {

void* QCborValue_new() {
	return new QCborValue();
}
void* QCborValue_new2(int t_) {
	return new QCborValue((QCborValue::Type)(t_));
}
void* QCborValue_new3(int b_) {
	return new QCborValue(b_);
}
void* QCborValue_new4(int i) {
	return new QCborValue(i);
}
void* QCborValue_new5(unsigned int u) {
	return new QCborValue(u);
}
void* QCborValue_new6(long long i) {
	return new QCborValue(i);
}
void* QCborValue_new7(double v) {
	return new QCborValue(v);
}
void* QCborValue_new8(int st) {
	return new QCborValue((QCborSimpleType)(st));
}
void* QCborValue_new9(void* ba) {
	return new QCborValue(*(QByteArray*)ba);
}
void* QCborValue_new10(const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	return new QCborValue(s_qstr);
}
void* QCborValue_new11(void* s) {
	return new QCborValue(*(QStringView*)s);
}
void* QCborValue_new12(void* s) {
	return new QCborValue(*(QLatin1String*)s);
}
void* QCborValue_new13(const char* s) {
	return new QCborValue(s);
}
void* QCborValue_new14(void* a) {
	return new QCborValue(*(QCborArray*)a);
}
void* QCborValue_new15(void* m) {
	return new QCborValue(*(QCborMap*)m);
}
void* QCborValue_new16(int tag) {
	return new QCborValue((QCborTag)(tag));
}
void* QCborValue_new17(int t_) {
	return new QCborValue((QCborKnownTags)(t_));
}
void* QCborValue_new18(void* dt) {
	return new QCborValue(*(QDateTime*)dt);
}
void* QCborValue_new19(void* url) {
	return new QCborValue(*(QUrl*)url);
}
void* QCborValue_new20(void* rx) {
	return new QCborValue(*(QRegularExpression*)rx);
}
void* QCborValue_new21(void* uuid) {
	return new QCborValue(*(QUuid*)uuid);
}
void* QCborValue_new22(void* other) {
	return new QCborValue(*(QCborValue*)other);
}
void* QCborValue_new23(int tag, void* taggedValue) {
	return new QCborValue((QCborTag)(tag), *(QCborValue*)taggedValue);
}
void* QCborValue_new24(int t_, void* tv) {
	return new QCborValue((QCborKnownTags)(t_), *(QCborValue*)tv);
}
void* QCborValue_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCborValue*)self)->operator=(*(QCborValue*)other);
	return (void*)&(__r);
}
void QCborValue_swap(void* self, void* other) {
	((QCborValue*)self)->swap(*(QCborValue*)other);
}
int QCborValue_type(void* self) {
	auto&& __r = ((QCborValue*)self)->type();
	return (int)(__r);
}
int QCborValue_isInteger(void* self) {
	auto&& __r = ((QCborValue*)self)->isInteger();
	return __r;
}
int QCborValue_isByteArray(void* self) {
	auto&& __r = ((QCborValue*)self)->isByteArray();
	return __r;
}
int QCborValue_isString(void* self) {
	auto&& __r = ((QCborValue*)self)->isString();
	return __r;
}
int QCborValue_isArray(void* self) {
	auto&& __r = ((QCborValue*)self)->isArray();
	return __r;
}
int QCborValue_isMap(void* self) {
	auto&& __r = ((QCborValue*)self)->isMap();
	return __r;
}
int QCborValue_isTag(void* self) {
	auto&& __r = ((QCborValue*)self)->isTag();
	return __r;
}
int QCborValue_isFalse(void* self) {
	auto&& __r = ((QCborValue*)self)->isFalse();
	return __r;
}
int QCborValue_isTrue(void* self) {
	auto&& __r = ((QCborValue*)self)->isTrue();
	return __r;
}
int QCborValue_isBool(void* self) {
	auto&& __r = ((QCborValue*)self)->isBool();
	return __r;
}
int QCborValue_isNull(void* self) {
	auto&& __r = ((QCborValue*)self)->isNull();
	return __r;
}
int QCborValue_isUndefined(void* self) {
	auto&& __r = ((QCborValue*)self)->isUndefined();
	return __r;
}
int QCborValue_isDouble(void* self) {
	auto&& __r = ((QCborValue*)self)->isDouble();
	return __r;
}
int QCborValue_isDateTime(void* self) {
	auto&& __r = ((QCborValue*)self)->isDateTime();
	return __r;
}
int QCborValue_isUrl(void* self) {
	auto&& __r = ((QCborValue*)self)->isUrl();
	return __r;
}
int QCborValue_isRegularExpression(void* self) {
	auto&& __r = ((QCborValue*)self)->isRegularExpression();
	return __r;
}
int QCborValue_isUuid(void* self) {
	auto&& __r = ((QCborValue*)self)->isUuid();
	return __r;
}
int QCborValue_isInvalid(void* self) {
	auto&& __r = ((QCborValue*)self)->isInvalid();
	return __r;
}
int QCborValue_isContainer(void* self) {
	auto&& __r = ((QCborValue*)self)->isContainer();
	return __r;
}
int QCborValue_isSimpleType(void* self) {
	auto&& __r = ((QCborValue*)self)->isSimpleType();
	return __r;
}
int QCborValue_isSimpleType2(void* self, int st) {
	auto&& __r = ((QCborValue*)self)->isSimpleType((QCborSimpleType)(st));
	return __r;
}
int QCborValue_toSimpleType(void* self) {
	auto&& __r = ((QCborValue*)self)->toSimpleType();
	return (int)(__r);
}
long long QCborValue_toInteger(void* self) {
	auto&& __r = ((QCborValue*)self)->toInteger();
	return __r;
}
int QCborValue_toBool(void* self) {
	auto&& __r = ((QCborValue*)self)->toBool();
	return __r;
}
double QCborValue_toDouble(void* self) {
	auto&& __r = ((QCborValue*)self)->toDouble();
	return __r;
}
int QCborValue_tag(void* self) {
	auto&& __r = ((QCborValue*)self)->tag();
	return (int)(__r);
}
void* QCborValue_taggedValue(void* self) {
	auto&& __r = ((QCborValue*)self)->taggedValue();
	return new QCborValue(__r);
}
void* QCborValue_toByteArray(void* self) {
	auto&& __r = ((QCborValue*)self)->toByteArray();
	return new QByteArray(__r);
}
char* QCborValue_toString(void* self, long long* __out_len) {
	auto&& __r = ((QCborValue*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborValue_toStringView(void* self) {
	auto&& __r = ((QCborValue*)self)->toStringView();
	return new QAnyStringView(__r);
}
void* QCborValue_toDateTime(void* self) {
	auto&& __r = ((QCborValue*)self)->toDateTime();
	return new QDateTime(__r);
}
void* QCborValue_toUrl(void* self) {
	auto&& __r = ((QCborValue*)self)->toUrl();
	return new QUrl(__r);
}
void* QCborValue_toRegularExpression(void* self) {
	auto&& __r = ((QCborValue*)self)->toRegularExpression();
	return new QRegularExpression(__r);
}
void* QCborValue_toUuid(void* self) {
	auto&& __r = ((QCborValue*)self)->toUuid();
	return new QUuid(__r);
}
void* QCborValue_toArray(void* self) {
	auto&& __r = ((QCborValue*)self)->toArray();
	return new QCborArray(__r);
}
void* QCborValue_toArray2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toArray(*(QCborArray*)defaultValue);
	return new QCborArray(__r);
}
void* QCborValue_toMap(void* self) {
	auto&& __r = ((QCborValue*)self)->toMap();
	return new QCborMap(__r);
}
void* QCborValue_toMap2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toMap(*(QCborMap*)defaultValue);
	return new QCborMap(__r);
}
void* QCborValue_OperatorSubscript(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborValue*)self)->operator[](key_qstr);
	return new QCborValue(__r);
}
void* QCborValue_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QCborValue*)self)->operator[](*(QLatin1String*)key);
	return new QCborValue(__r);
}
void* QCborValue_OperatorSubscript3(void* self, long long key) {
	auto&& __r = ((QCborValue*)self)->operator[](key);
	return new QCborValue(__r);
}
void* QCborValue_OperatorSubscript4(void* self, long long key) {
	auto&& __r = ((QCborValue*)self)->operator[](key);
	return new QCborValueRef(__r);
}
void* QCborValue_OperatorSubscript5(void* self, void* key) {
	auto&& __r = ((QCborValue*)self)->operator[](*(QLatin1String*)key);
	return new QCborValueRef(__r);
}
void* QCborValue_OperatorSubscript6(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QCborValue*)self)->operator[](key_qstr);
	return new QCborValueRef(__r);
}
int QCborValue_compare(void* self, void* other) {
	auto&& __r = ((QCborValue*)self)->compare(*(QCborValue*)other);
	return __r;
}
void* QCborValue_fromVariant(void* variant) {
	auto&& __r = QCborValue::fromVariant(*(QVariant*)variant);
	return new QCborValue(__r);
}
void* QCborValue_toVariant(void* self) {
	auto&& __r = ((QCborValue*)self)->toVariant();
	return new QVariant(__r);
}
void* QCborValue_fromJsonValue(void* v) {
	auto&& __r = QCborValue::fromJsonValue(*(QJsonValue*)v);
	return new QCborValue(__r);
}
void* QCborValue_toJsonValue(void* self) {
	auto&& __r = ((QCborValue*)self)->toJsonValue();
	return new QJsonValue(__r);
}
void* QCborValue_fromCbor(void* reader) {
	auto&& __r = QCborValue::fromCbor(*(QCborStreamReader*)reader);
	return new QCborValue(__r);
}
void* QCborValue_fromCbor2(void* ba) {
	auto&& __r = QCborValue::fromCbor(*(QByteArray*)ba);
	return new QCborValue(__r);
}
void* QCborValue_fromCbor3(const char* data, long long len_) {
	auto&& __r = QCborValue::fromCbor(data, len_);
	return new QCborValue(__r);
}
void* QCborValue_fromCbor4(const unsigned char* data, long long len_) {
	auto&& __r = QCborValue::fromCbor(data, len_);
	return new QCborValue(__r);
}
void* QCborValue_toCbor(void* self) {
	auto&& __r = ((QCborValue*)self)->toCbor();
	return new QByteArray(__r);
}
void QCborValue_toCbor2(void* self, void* writer) {
	((QCborValue*)self)->toCbor(*(QCborStreamWriter*)writer);
}
char* QCborValue_toDiagnosticNotation(void* self, long long* __out_len) {
	auto&& __r = ((QCborValue*)self)->toDiagnosticNotation();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCborValue_toSimpleType2(void* self, int defaultValue) {
	auto&& __r = ((QCborValue*)self)->toSimpleType((QCborSimpleType)(defaultValue));
	return (int)(__r);
}
long long QCborValue_toInteger2(void* self, long long defaultValue) {
	auto&& __r = ((QCborValue*)self)->toInteger(defaultValue);
	return __r;
}
int QCborValue_toBool2(void* self, int defaultValue) {
	auto&& __r = ((QCborValue*)self)->toBool(defaultValue);
	return __r;
}
double QCborValue_toDouble2(void* self, double defaultValue) {
	auto&& __r = ((QCborValue*)self)->toDouble(defaultValue);
	return __r;
}
int QCborValue_tag2(void* self, int defaultValue) {
	auto&& __r = ((QCborValue*)self)->tag((QCborTag)(defaultValue));
	return (int)(__r);
}
void* QCborValue_taggedValue2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->taggedValue(*(QCborValue*)defaultValue);
	return new QCborValue(__r);
}
void* QCborValue_toByteArray2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toByteArray(*(QByteArray*)defaultValue);
	return new QByteArray(__r);
}
char* QCborValue_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QCborValue*)self)->toString(defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborValue_toStringView2(void* self, const char* defaultValue, long long defaultValue_len) {
	QAnyStringView defaultValue_qstr = QAnyStringView(defaultValue, defaultValue_len);
	auto&& __r = ((QCborValue*)self)->toStringView(defaultValue_qstr);
	return new QAnyStringView(__r);
}
void* QCborValue_toDateTime2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toDateTime(*(QDateTime*)defaultValue);
	return new QDateTime(__r);
}
void* QCborValue_toUrl2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toUrl(*(QUrl*)defaultValue);
	return new QUrl(__r);
}
void* QCborValue_toRegularExpression2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toRegularExpression(*(QRegularExpression*)defaultValue);
	return new QRegularExpression(__r);
}
void* QCborValue_toUuid2(void* self, void* defaultValue) {
	auto&& __r = ((QCborValue*)self)->toUuid(*(QUuid*)defaultValue);
	return new QUuid(__r);
}
void* QCborValue_fromCbor5(void* ba, void* error_) {
	auto&& __r = QCborValue::fromCbor(*(QByteArray*)ba, (QCborParserError*)error_);
	return new QCborValue(__r);
}
void* QCborValue_fromCbor6(const char* data, long long len_, void* error_) {
	auto&& __r = QCborValue::fromCbor(data, len_, (QCborParserError*)error_);
	return new QCborValue(__r);
}
void* QCborValue_fromCbor7(const unsigned char* data, long long len_, void* error_) {
	auto&& __r = QCborValue::fromCbor(data, len_, (QCborParserError*)error_);
	return new QCborValue(__r);
}
void* QCborValue_toCbor3(void* self, int opt) {
	auto&& __r = ((QCborValue*)self)->toCbor(QFlags<QCborValue::EncodingOption>(opt));
	return new QByteArray(__r);
}
void QCborValue_toCbor4(void* self, void* writer, int opt) {
	((QCborValue*)self)->toCbor(*(QCborStreamWriter*)writer, QFlags<QCborValue::EncodingOption>(opt));
}
char* QCborValue_toDiagnosticNotation2(void* self, int opts, long long* __out_len) {
	auto&& __r = ((QCborValue*)self)->toDiagnosticNotation(QFlags<QCborValue::DiagnosticNotationOption>(opts));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCborValue_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) == (*(QCborValue*)rhs);
	return __r;
}
int QCborValue_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) != (*(QCborValue*)rhs);
	return __r;
}
int QCborValue_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) < (*(QCborValue*)rhs);
	return __r;
}
int QCborValue_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) > (*(QCborValue*)rhs);
	return __r;
}
int QCborValue_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) <= (*(QCborValue*)rhs);
	return __r;
}
int QCborValue_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) >= (*(QCborValue*)rhs);
	return __r;
}
int QCborValue_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) == (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValue_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) != (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValue_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) < (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValue_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) > (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValue_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) <= (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValue_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) >= (*(QCborValueConstRef*)rhs);
	return __r;
}
int QCborValue_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) == (*(QCborArray*)rhs);
	return __r;
}
int QCborValue_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) != (*(QCborArray*)rhs);
	return __r;
}
int QCborValue_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) < (*(QCborArray*)rhs);
	return __r;
}
int QCborValue_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) > (*(QCborArray*)rhs);
	return __r;
}
int QCborValue_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) <= (*(QCborArray*)rhs);
	return __r;
}
int QCborValue_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) >= (*(QCborArray*)rhs);
	return __r;
}
int QCborValue_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) == (*(QCborMap*)rhs);
	return __r;
}
int QCborValue_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) != (*(QCborMap*)rhs);
	return __r;
}
int QCborValue_OperatorLesser4(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) < (*(QCborMap*)rhs);
	return __r;
}
int QCborValue_OperatorGreater4(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) > (*(QCborMap*)rhs);
	return __r;
}
int QCborValue_OperatorLesserOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) <= (*(QCborMap*)rhs);
	return __r;
}
int QCborValue_OperatorGreaterOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QCborValue*)self) >= (*(QCborMap*)rhs);
	return __r;
}
void QCborValue_delete(void* self) { delete (QCborValue*)self; }

}

#include <QtCore/qcryptographichash.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include "gen_qcryptographichash.h"

extern "C" {

void* QCryptographicHash_new(int method) {
	return new QCryptographicHash((QCryptographicHash::Algorithm)(method));
}
void QCryptographicHash_swap(void* self, void* other) {
	((QCryptographicHash*)self)->swap(*(QCryptographicHash*)other);
}
void QCryptographicHash_reset(void* self) {
	((QCryptographicHash*)self)->reset();
}
int QCryptographicHash_algorithm(void* self) {
	auto&& __r = ((QCryptographicHash*)self)->algorithm();
	return (int)(__r);
}
void QCryptographicHash_addData(void* self, const char* data, long long length) {
	((QCryptographicHash*)self)->addData(data, length);
}
void QCryptographicHash_addData2(void* self, void* data) {
	((QCryptographicHash*)self)->addData(*(QByteArrayView*)data);
}
int QCryptographicHash_addData3(void* self, void* device) {
	auto&& __r = ((QCryptographicHash*)self)->addData((QIODevice*)device);
	return __r;
}
void* QCryptographicHash_result(void* self) {
	auto&& __r = ((QCryptographicHash*)self)->result();
	return new QByteArray(__r);
}
void* QCryptographicHash_resultView(void* self) {
	auto&& __r = ((QCryptographicHash*)self)->resultView();
	return new QByteArrayView(__r);
}
void* QCryptographicHash_hash(void* data, int method) {
	auto&& __r = QCryptographicHash::hash(*(QByteArrayView*)data, (QCryptographicHash::Algorithm)(method));
	return new QByteArray(__r);
}
int QCryptographicHash_hashLength(int method) {
	auto&& __r = QCryptographicHash::hashLength((QCryptographicHash::Algorithm)(method));
	return __r;
}
int QCryptographicHash_supportsAlgorithm(int method) {
	auto&& __r = QCryptographicHash::supportsAlgorithm((QCryptographicHash::Algorithm)(method));
	return __r;
}
void QCryptographicHash_delete(void* self) { delete (QCryptographicHash*)self; }

}

#include <QtCore/qidentityproxymodel.h>
#include <QtCore/qabstractproxymodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qsize.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qidentityproxymodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QIdentityProxyModel__Go : public QIdentityProxyModel {
public:
	long long __goId;
	unsigned char __goOn[47] = {};
	template<typename... A> QIdentityProxyModel__Go(long long id, A&&... a) : QIdentityProxyModel(std::forward<A>(a)...), __goId(id) {}
	using QIdentityProxyModel::setHandleSourceLayoutChanges;
	using QIdentityProxyModel::setHandleSourceDataChanges;
	using QIdentityProxyModel::createIndex;
	using QIdentityProxyModel::encodeData;
	using QIdentityProxyModel::decodeData;
	using QIdentityProxyModel::beginInsertRows;
	using QIdentityProxyModel::endInsertRows;
	using QIdentityProxyModel::beginRemoveRows;
	using QIdentityProxyModel::endRemoveRows;
	using QIdentityProxyModel::beginMoveRows;
	using QIdentityProxyModel::endMoveRows;
	using QIdentityProxyModel::beginInsertColumns;
	using QIdentityProxyModel::endInsertColumns;
	using QIdentityProxyModel::beginRemoveColumns;
	using QIdentityProxyModel::endRemoveColumns;
	using QIdentityProxyModel::beginMoveColumns;
	using QIdentityProxyModel::endMoveColumns;
	using QIdentityProxyModel::beginResetModel;
	using QIdentityProxyModel::endResetModel;
	using QIdentityProxyModel::changePersistentIndex;
	using QIdentityProxyModel::changePersistentIndexList;
	using QIdentityProxyModel::persistentIndexList;
	using QIdentityProxyModel::sender;
	using QIdentityProxyModel::senderSignalIndex;
	using QIdentityProxyModel::receivers;
	using QIdentityProxyModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QIdentityProxyModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QIdentityProxyModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QIdentityProxyModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QIdentityProxyModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QIdentityProxyModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QIdentityProxyModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QIdentityProxyModel::metaObject();
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::columnCount(parent);
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QIdentityProxyModel::index(row, column, parent);
	}
	QModelIndex mapFromSource(const QModelIndex& sourceIndex) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1717009778236344012LL, (void*)&sourceIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QIdentityProxyModel::mapFromSource(sourceIndex);
	}
	QModelIndex mapToSource(const QModelIndex& proxyIndex) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -2333321292016234891LL, (void*)&proxyIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QIdentityProxyModel::mapToSource(proxyIndex);
	}
	QModelIndex parent(const QModelIndex& child) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&child, &__h); if (__h) return *(QModelIndex*)__r; }
		return QIdentityProxyModel::parent(child);
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::rowCount(parent);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QIdentityProxyModel::headerData(section, orientation, role);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::dropMimeData(data, action, row, column, parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QIdentityProxyModel::sibling(row, column, idx);
	}
	QItemSelection mapSelectionFromSource(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 9067522407071207004LL, (void*)&selection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QIdentityProxyModel::mapSelectionFromSource(selection);
	}
	QItemSelection mapSelectionToSource(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2788628866280405837LL, (void*)&selection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QIdentityProxyModel::mapSelectionToSource(selection);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
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
		return QIdentityProxyModel::match(start, role, value, hits, flags);
	}
	void setSourceModel(QAbstractItemModel* sourceModel) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7818871925547456643LL, (void*)sourceModel, &__h); if (__h) return; }
		QIdentityProxyModel::setSourceModel(sourceModel);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::insertColumns(column, count, parent);
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::insertRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::removeColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::removeRows(row, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QIdentityProxyModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QIdentityProxyModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QIdentityProxyModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QIdentityProxyModel::revert();
	}
	QVariant data(const QModelIndex& proxyIndex, int role) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&proxyIndex, role, &__h); if (__h) return *(QVariant*)__r; }
		return QIdentityProxyModel::data(proxyIndex, role);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QIdentityProxyModel::flags(index);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QIdentityProxyModel::setData(index, value, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QIdentityProxyModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QIdentityProxyModel::clearItemData(index);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QIdentityProxyModel::buddy(index);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::canFetchMore(parent);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QIdentityProxyModel::fetchMore(parent);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QIdentityProxyModel::sort(column, order);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QIdentityProxyModel::span(index);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::hasChildren(parent);
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[33] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QIdentityProxyModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QIdentityProxyModel::canDropMimeData(data, action, row, column, parent);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virtp_voidp) {
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
		return QIdentityProxyModel::mimeTypes();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QIdentityProxyModel::supportedDragActions();
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QIdentityProxyModel::supportedDropActions();
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QIdentityProxyModel::multiData(index, roleDataSpan);
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QIdentityProxyModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QIdentityProxyModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QIdentityProxyModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QIdentityProxyModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QIdentityProxyModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QIdentityProxyModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QIdentityProxyModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QIdentityProxyModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QIdentityProxyModel;
static std::mutex __connsMu_QIdentityProxyModel;

extern "C" {

void* QIdentityProxyModel_base(void* self) { return (void*)static_cast<QAbstractProxyModel*>((QIdentityProxyModel*)self); }
void QIdentityProxyModel_chain(void* self, void** out) {
	QIdentityProxyModel* __p = (QIdentityProxyModel*)self;
	QAbstractProxyModel* __b0 = static_cast<QAbstractProxyModel*>(__p); out[0] = (void*)__b0;
	QAbstractItemModel* __b1 = static_cast<QAbstractItemModel*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QIdentityProxyModel__set_override(void* self, int idx) { ((QIdentityProxyModel__Go*)self)->__goOn[idx] = 1; }
void* QIdentityProxyModel_fromQObject(void* o) { return (void*)qobject_cast<QIdentityProxyModel*>((QObject*)o); }
void* QIdentityProxyModel_new(long long __goid) {
	return new QIdentityProxyModel__Go(__goid);
}
void* QIdentityProxyModel_new2(long long __goid, void* parent) {
	return new QIdentityProxyModel__Go(__goid, (QObject*)parent);
}
char* QIdentityProxyModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QIdentityProxyModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QIdentityProxyModel_handleSourceLayoutChanges(void* self) {
	auto&& __r = ((QIdentityProxyModel*)self)->handleSourceLayoutChanges();
	return __r;
}
int QIdentityProxyModel_handleSourceDataChanges(void* self) {
	auto&& __r = ((QIdentityProxyModel*)self)->handleSourceDataChanges();
	return __r;
}
void QIdentityProxyModel_setHandleSourceLayoutChanges(void* self, int arg1) {
	((QIdentityProxyModel__Go*)self)->setHandleSourceLayoutChanges(arg1);
}
void QIdentityProxyModel_setHandleSourceDataChanges(void* self, int arg1) {
	((QIdentityProxyModel__Go*)self)->setHandleSourceDataChanges(arg1);
}
void* QIdentityProxyModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QIdentityProxyModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QIdentityProxyModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QIdentityProxyModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QIdentityProxyModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QIdentityProxyModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QIdentityProxyModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QIdentityProxyModel_endInsertRows(void* self) {
	((QIdentityProxyModel__Go*)self)->endInsertRows();
}
void QIdentityProxyModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QIdentityProxyModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QIdentityProxyModel_endRemoveRows(void* self) {
	((QIdentityProxyModel__Go*)self)->endRemoveRows();
}
int QIdentityProxyModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QIdentityProxyModel_endMoveRows(void* self) {
	((QIdentityProxyModel__Go*)self)->endMoveRows();
}
void QIdentityProxyModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QIdentityProxyModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QIdentityProxyModel_endInsertColumns(void* self) {
	((QIdentityProxyModel__Go*)self)->endInsertColumns();
}
void QIdentityProxyModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QIdentityProxyModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QIdentityProxyModel_endRemoveColumns(void* self) {
	((QIdentityProxyModel__Go*)self)->endRemoveColumns();
}
int QIdentityProxyModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QIdentityProxyModel_endMoveColumns(void* self) {
	((QIdentityProxyModel__Go*)self)->endMoveColumns();
}
void QIdentityProxyModel_beginResetModel(void* self) {
	((QIdentityProxyModel__Go*)self)->beginResetModel();
}
void QIdentityProxyModel_endResetModel(void* self) {
	((QIdentityProxyModel__Go*)self)->endResetModel();
}
void QIdentityProxyModel_changePersistentIndex(void* self, void* from, void* to) {
	((QIdentityProxyModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QIdentityProxyModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QIdentityProxyModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QIdentityProxyModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QIdentityProxyModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QIdentityProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QIdentityProxyModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QIdentityProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QIdentityProxyModel_sender(void* self) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->sender();
	return (void*)(__r);
}
int QIdentityProxyModel_senderSignalIndex(void* self) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QIdentityProxyModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->receivers(signal);
	return __r;
}
int QIdentityProxyModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QIdentityProxyModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QIdentityProxyModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QIdentityProxyModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QIdentityProxyModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QIdentityProxyModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QIdentityProxyModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QIdentityProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QIdentityProxyModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QIdentityProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QIdentityProxyModel_metaObject_default(void* self) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::metaObject();
	return (void*)(__r);
}
int QIdentityProxyModel_columnCount_default(void* self, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::columnCount(*(QModelIndex*)parent);
	return __r;
}
void* QIdentityProxyModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QIdentityProxyModel_mapFromSource_default(void* self, void* sourceIndex) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::mapFromSource(*(QModelIndex*)sourceIndex);
	return new QModelIndex(__r);
}
void* QIdentityProxyModel_mapToSource_default(void* self, void* proxyIndex) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::mapToSource(*(QModelIndex*)proxyIndex);
	return new QModelIndex(__r);
}
void* QIdentityProxyModel_parent_default(void* self, void* child) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::parent(*(QModelIndex*)child);
	return new QModelIndex(__r);
}
int QIdentityProxyModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
void* QIdentityProxyModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QIdentityProxyModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
void* QIdentityProxyModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
void* QIdentityProxyModel_mapSelectionFromSource_default(void* self, void* selection) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::mapSelectionFromSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void* QIdentityProxyModel_mapSelectionToSource_default(void* self, void* selection) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::mapSelectionToSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void** QIdentityProxyModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QIdentityProxyModel_setSourceModel_default(void* self, void* sourceModel) {
	((QIdentityProxyModel*)self)->QIdentityProxyModel::setSourceModel((QAbstractItemModel*)sourceModel);
}
int QIdentityProxyModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QIdentityProxyModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QIdentityProxyModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QIdentityProxyModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QIdentityProxyModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QIdentityProxyModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QIdentityProxyModel_submit_default(void* self) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::submit();
	return __r;
}
void QIdentityProxyModel_revert_default(void* self) {
	((QIdentityProxyModel*)self)->QIdentityProxyModel::revert();
}
void* QIdentityProxyModel_data_default(void* self, void* proxyIndex, int role) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::data(*(QModelIndex*)proxyIndex, role);
	return new QVariant(__r);
}
int QIdentityProxyModel_flags_default(void* self, void* index) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QIdentityProxyModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
int QIdentityProxyModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QIdentityProxyModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
void* QIdentityProxyModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
int QIdentityProxyModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QIdentityProxyModel_fetchMore_default(void* self, void* parent) {
	((QIdentityProxyModel*)self)->QIdentityProxyModel::fetchMore(*(QModelIndex*)parent);
}
void QIdentityProxyModel_sort_default(void* self, int column, int order) {
	((QIdentityProxyModel*)self)->QIdentityProxyModel::sort(column, (Qt::SortOrder)(order));
}
void* QIdentityProxyModel_span_default(void* self, void* index) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
int QIdentityProxyModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QIdentityProxyModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QIdentityProxyModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
char** QIdentityProxyModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QIdentityProxyModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::supportedDragActions();
	return (int)(__r);
}
int QIdentityProxyModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::supportedDropActions();
	return (int)(__r);
}
void QIdentityProxyModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QIdentityProxyModel*)self)->QIdentityProxyModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
void QIdentityProxyModel_resetInternalData_default(void* self) {
	((QIdentityProxyModel__Go*)self)->cppgoBase_resetInternalData();
}
int QIdentityProxyModel_event_default(void* self, void* event) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::event((QEvent*)event);
	return __r;
}
int QIdentityProxyModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QIdentityProxyModel*)self)->QIdentityProxyModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QIdentityProxyModel_timerEvent_default(void* self, void* event) {
	((QIdentityProxyModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QIdentityProxyModel_childEvent_default(void* self, void* event) {
	((QIdentityProxyModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QIdentityProxyModel_customEvent_default(void* self, void* event) {
	((QIdentityProxyModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QIdentityProxyModel_connectNotify_default(void* self, void* signal) {
	((QIdentityProxyModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QIdentityProxyModel_disconnectNotify_default(void* self, void* signal) {
	((QIdentityProxyModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QIdentityProxyModel_connect_sourceModelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::sourceModelChanged, (QIdentityProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::headerDataChanged, (QIdentityProxyModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::layoutChanged, (QIdentityProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::layoutChanged, (QIdentityProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::layoutAboutToBeChanged, (QIdentityProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::layoutAboutToBeChanged, (QIdentityProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::rowsAboutToBeInserted, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::rowsInserted, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::rowsAboutToBeRemoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::rowsRemoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::columnsAboutToBeInserted, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::columnsInserted, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::columnsAboutToBeRemoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::columnsRemoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::modelAboutToBeReset, (QIdentityProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::modelReset, (QIdentityProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::rowsAboutToBeMoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::rowsMoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::columnsAboutToBeMoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::columnsMoved, (QIdentityProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::destroyed, (QIdentityProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::destroyed, (QIdentityProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIdentityProxyModel*)self, &QIdentityProxyModel::objectNameChanged, (QIdentityProxyModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
	__conns_QIdentityProxyModel[__id] = __c;
}
void QIdentityProxyModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QIdentityProxyModel);
		auto __it = __conns_QIdentityProxyModel.find(id);
		if (__it == __conns_QIdentityProxyModel.end()) return;
		__c = __it->second;
		__conns_QIdentityProxyModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QIdentityProxyModel__Go_new(long long id, void* parent) { return new QIdentityProxyModel__Go(id, (QObject*)parent); }
void QIdentityProxyModel_delete(void* self) { delete (QIdentityProxyModel*)self; }

}

#include <QtCore/qmessageauthenticationcode.h>
#include <QtCore/qcryptographichash.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include "gen_qmessageauthenticationcode.h"

extern "C" {

void* QMessageAuthenticationCode_new(int method) {
	return new QMessageAuthenticationCode((QCryptographicHash::Algorithm)(method));
}
void* QMessageAuthenticationCode_new2(int method, void* key) {
	return new QMessageAuthenticationCode((QCryptographicHash::Algorithm)(method), *(QByteArrayView*)key);
}
void QMessageAuthenticationCode_swap(void* self, void* other) {
	((QMessageAuthenticationCode*)self)->swap(*(QMessageAuthenticationCode*)other);
}
void QMessageAuthenticationCode_reset(void* self) {
	((QMessageAuthenticationCode*)self)->reset();
}
void QMessageAuthenticationCode_setKey(void* self, void* key) {
	((QMessageAuthenticationCode*)self)->setKey(*(QByteArrayView*)key);
}
void QMessageAuthenticationCode_addData(void* self, const char* data, long long length) {
	((QMessageAuthenticationCode*)self)->addData(data, length);
}
void QMessageAuthenticationCode_addData2(void* self, void* data) {
	((QMessageAuthenticationCode*)self)->addData(*(QByteArrayView*)data);
}
int QMessageAuthenticationCode_addData3(void* self, void* device) {
	auto&& __r = ((QMessageAuthenticationCode*)self)->addData((QIODevice*)device);
	return __r;
}
void* QMessageAuthenticationCode_resultView(void* self) {
	auto&& __r = ((QMessageAuthenticationCode*)self)->resultView();
	return new QByteArrayView(__r);
}
void* QMessageAuthenticationCode_result(void* self) {
	auto&& __r = ((QMessageAuthenticationCode*)self)->result();
	return new QByteArray(__r);
}
void* QMessageAuthenticationCode_hash(void* message, void* key, int method) {
	auto&& __r = QMessageAuthenticationCode::hash(*(QByteArrayView*)message, *(QByteArrayView*)key, (QCryptographicHash::Algorithm)(method));
	return new QByteArray(__r);
}
void QMessageAuthenticationCode_delete(void* self) { delete (QMessageAuthenticationCode*)self; }

}

#include <QtCore/qplugin.h>
#include "gen_qpluginmetadata__elfnoteheader.h"

extern "C" {

void* QPluginMetaData__ElfNoteHeader_new(unsigned int payloadSize) {
	return new QPluginMetaData::ElfNoteHeader(payloadSize);
}
void* QPluginMetaData__ElfNoteHeader_new2(void* arg1) {
	return new QPluginMetaData::ElfNoteHeader(*(QPluginMetaData::ElfNoteHeader*)arg1);
}
unsigned int QPluginMetaData__ElfNoteHeader_n_namesz_get(void* self) {
	auto __r = ((QPluginMetaData::ElfNoteHeader*)self)->n_namesz;
	return __r;
}
void QPluginMetaData__ElfNoteHeader_n_namesz_set(void* self, unsigned int n_namesz) {
	((QPluginMetaData::ElfNoteHeader*)self)->n_namesz = n_namesz;
}
unsigned int QPluginMetaData__ElfNoteHeader_n_descsz_get(void* self) {
	auto __r = ((QPluginMetaData::ElfNoteHeader*)self)->n_descsz;
	return __r;
}
void QPluginMetaData__ElfNoteHeader_n_descsz_set(void* self, unsigned int n_descsz) {
	((QPluginMetaData::ElfNoteHeader*)self)->n_descsz = n_descsz;
}
unsigned int QPluginMetaData__ElfNoteHeader_n_type_get(void* self) {
	auto __r = ((QPluginMetaData::ElfNoteHeader*)self)->n_type;
	return __r;
}
void QPluginMetaData__ElfNoteHeader_n_type_set(void* self, unsigned int n_type) {
	((QPluginMetaData::ElfNoteHeader*)self)->n_type = n_type;
}
void* QPluginMetaData__ElfNoteHeader_header_get(void* self) {
	auto __r = ((QPluginMetaData::ElfNoteHeader*)self)->header;
	return new QPluginMetaData::Header(__r);
}
void QPluginMetaData__ElfNoteHeader_header_set(void* self, void* header) {
	((QPluginMetaData::ElfNoteHeader*)self)->header = *(QPluginMetaData::Header*)header;
}
void QPluginMetaData__ElfNoteHeader_delete(void* self) { delete (QPluginMetaData::ElfNoteHeader*)self; }

}

#include <QtCore/qsequentialiterable.h>
#include <QtCore/qmetacontainer.h>
#include "gen_qsequentialiterator.h"

extern "C" {

void* QSequentialIterator_new(void* arg1) {
	return new QSequentialIterator(*(QSequentialIterator*)arg1);
}
void* QSequentialIterator_metaContainer(void* self) {
	auto&& __r = ((QSequentialIterator*)self)->metaContainer();
	return new QMetaSequence(__r);
}
void QSequentialIterator_delete(void* self) { delete (QSequentialIterator*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include "gen_qxmlstreamattribute.h"

extern "C" {

void* QXmlStreamAttribute_new() {
	return new QXmlStreamAttribute();
}
void* QXmlStreamAttribute_new2(const char* qualifiedName, long long qualifiedName_len, const char* value, long long value_len) {
	QString qualifiedName_qstr = QString::fromUtf8(qualifiedName, qualifiedName_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	return new QXmlStreamAttribute(qualifiedName_qstr, value_qstr);
}
void* QXmlStreamAttribute_new3(const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* value, long long value_len) {
	QString namespaceUri_qstr = QString::fromUtf8(namespaceUri, namespaceUri_len);
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	return new QXmlStreamAttribute(namespaceUri_qstr, name_qstr, value_qstr);
}
void* QXmlStreamAttribute_new4(void* arg1) {
	return new QXmlStreamAttribute(*(QXmlStreamAttribute*)arg1);
}
void* QXmlStreamAttribute_namespaceUri(void* self) {
	auto&& __r = ((QXmlStreamAttribute*)self)->namespaceUri();
	return new QStringView(__r);
}
void* QXmlStreamAttribute_name(void* self) {
	auto&& __r = ((QXmlStreamAttribute*)self)->name();
	return new QStringView(__r);
}
void* QXmlStreamAttribute_qualifiedName(void* self) {
	auto&& __r = ((QXmlStreamAttribute*)self)->qualifiedName();
	return new QStringView(__r);
}
void* QXmlStreamAttribute_prefix(void* self) {
	auto&& __r = ((QXmlStreamAttribute*)self)->prefix();
	return new QStringView(__r);
}
void* QXmlStreamAttribute_value(void* self) {
	auto&& __r = ((QXmlStreamAttribute*)self)->value();
	return new QStringView(__r);
}
int QXmlStreamAttribute_isDefault(void* self) {
	auto&& __r = ((QXmlStreamAttribute*)self)->isDefault();
	return __r;
}
int QXmlStreamAttribute_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamAttribute*)self) == (*(QXmlStreamAttribute*)rhs);
	return __r;
}
int QXmlStreamAttribute_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QXmlStreamAttribute*)self) != (*(QXmlStreamAttribute*)rhs);
	return __r;
}
void QXmlStreamAttribute_delete(void* self) { delete (QXmlStreamAttribute*)self; }

}

#include <QtGui/qinputmethod.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtransform.h>
#include <QtCore/qrect.h>
#include <QtCore/qlocale.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qinputmethod.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QInputMethod;
static std::mutex __connsMu_QInputMethod;

extern "C" {

void* QInputMethod_base(void* self) { return (void*)static_cast<QObject*>((QInputMethod*)self); }
void* QInputMethod_fromQObject(void* o) { return (void*)qobject_cast<QInputMethod*>((QObject*)o); }
char* QInputMethod_tr(const char* s, long long* __out_len) {
	auto&& __r = QInputMethod::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QInputMethod_inputItemTransform(void* self) {
	auto&& __r = ((QInputMethod*)self)->inputItemTransform();
	return new QTransform(__r);
}
void QInputMethod_setInputItemTransform(void* self, void* transform) {
	((QInputMethod*)self)->setInputItemTransform(*(QTransform*)transform);
}
void* QInputMethod_inputItemRectangle(void* self) {
	auto&& __r = ((QInputMethod*)self)->inputItemRectangle();
	return new QRectF(__r);
}
void QInputMethod_setInputItemRectangle(void* self, void* rect) {
	((QInputMethod*)self)->setInputItemRectangle(*(QRectF*)rect);
}
void* QInputMethod_cursorRectangle(void* self) {
	auto&& __r = ((QInputMethod*)self)->cursorRectangle();
	return new QRectF(__r);
}
void* QInputMethod_anchorRectangle(void* self) {
	auto&& __r = ((QInputMethod*)self)->anchorRectangle();
	return new QRectF(__r);
}
void* QInputMethod_keyboardRectangle(void* self) {
	auto&& __r = ((QInputMethod*)self)->keyboardRectangle();
	return new QRectF(__r);
}
void* QInputMethod_inputItemClipRectangle(void* self) {
	auto&& __r = ((QInputMethod*)self)->inputItemClipRectangle();
	return new QRectF(__r);
}
int QInputMethod_isVisible(void* self) {
	auto&& __r = ((QInputMethod*)self)->isVisible();
	return __r;
}
void QInputMethod_setVisible(void* self, int visible) {
	((QInputMethod*)self)->setVisible(visible);
}
int QInputMethod_isAnimating(void* self) {
	auto&& __r = ((QInputMethod*)self)->isAnimating();
	return __r;
}
void* QInputMethod_locale(void* self) {
	auto&& __r = ((QInputMethod*)self)->locale();
	return new QLocale(__r);
}
int QInputMethod_inputDirection(void* self) {
	auto&& __r = ((QInputMethod*)self)->inputDirection();
	return (int)(__r);
}
void* QInputMethod_queryFocusObject(int query, void* argument) {
	auto&& __r = QInputMethod::queryFocusObject((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void QInputMethod_show(void* self) {
	((QInputMethod*)self)->show();
}
void QInputMethod_hide(void* self) {
	((QInputMethod*)self)->hide();
}
void QInputMethod_update(void* self, int queries) {
	((QInputMethod*)self)->update(QFlags<Qt::InputMethodQuery>(queries));
}
void QInputMethod_reset(void* self) {
	((QInputMethod*)self)->reset();
}
void QInputMethod_commit(void* self) {
	((QInputMethod*)self)->commit();
}
void QInputMethod_invokeAction(void* self, int a, int cursorPosition) {
	((QInputMethod*)self)->invokeAction((QInputMethod::Action)(a), cursorPosition);
}
void QInputMethod_cursorRectangleChanged(void* self) {
	((QInputMethod*)self)->cursorRectangleChanged();
}
void QInputMethod_anchorRectangleChanged(void* self) {
	((QInputMethod*)self)->anchorRectangleChanged();
}
void QInputMethod_keyboardRectangleChanged(void* self) {
	((QInputMethod*)self)->keyboardRectangleChanged();
}
void QInputMethod_inputItemClipRectangleChanged(void* self) {
	((QInputMethod*)self)->inputItemClipRectangleChanged();
}
void QInputMethod_visibleChanged(void* self) {
	((QInputMethod*)self)->visibleChanged();
}
void QInputMethod_animatingChanged(void* self) {
	((QInputMethod*)self)->animatingChanged();
}
void QInputMethod_localeChanged(void* self) {
	((QInputMethod*)self)->localeChanged();
}
void QInputMethod_inputDirectionChanged(void* self, int newDirection) {
	((QInputMethod*)self)->inputDirectionChanged((Qt::LayoutDirection)(newDirection));
}
void* QInputMethod_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QInputMethod::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QInputMethod_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QInputMethod::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QInputMethod_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QInputMethod::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QInputMethod_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QInputMethod::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QInputMethod_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QInputMethod::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QInputMethod_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QInputMethod::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QInputMethod_connect_cursorRectangleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::cursorRectangleChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_anchorRectangleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::anchorRectangleChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_keyboardRectangleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::keyboardRectangleChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_inputItemClipRectangleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::inputItemClipRectangleChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::visibleChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_animatingChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::animatingChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_localeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::localeChanged, (QInputMethod*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_inputDirectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::inputDirectionChanged, (QInputMethod*)self, [__id](Qt::LayoutDirection newDirection){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(newDirection));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::destroyed, (QInputMethod*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::destroyed, (QInputMethod*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QInputMethod*)self, &QInputMethod::objectNameChanged, (QInputMethod*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
	__conns_QInputMethod[__id] = __c;
}
void QInputMethod__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QInputMethod);
		auto __it = __conns_QInputMethod.find(id);
		if (__it == __conns_QInputMethod.end()) return;
		__c = __it->second;
		__conns_QInputMethod.erase(__it);
	}
	QObject::disconnect(__c);
}

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qicondragevent.h"
#include "gen_purego_cb.h"

class QIconDragEvent__Go : public QIconDragEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QIconDragEvent__Go(long long id, A&&... a) : QIconDragEvent(std::forward<A>(a)...), __goId(id) {}
	QIconDragEvent* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QIconDragEvent*)__r; }
		return QIconDragEvent::clone();
	}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QIconDragEvent::setAccepted(accepted);
	}
};

extern "C" {

void* QIconDragEvent_base(void* self) { return (void*)static_cast<QEvent*>((QIconDragEvent*)self); }
void QIconDragEvent__set_override(void* self, int idx) { ((QIconDragEvent__Go*)self)->__goOn[idx] = 1; }
void* QIconDragEvent_new(long long __goid) {
	return new QIconDragEvent__Go(__goid);
}
int QIconDragEvent_registerEventType() {
	auto&& __r = QIconDragEvent::registerEventType();
	return __r;
}
int QIconDragEvent_registerEventType2(int hint) {
	auto&& __r = QIconDragEvent::registerEventType(hint);
	return __r;
}
void* QIconDragEvent_clone_default(void* self) {
	auto&& __r = ((QIconDragEvent*)self)->QIconDragEvent::clone();
	return (void*)(__r);
}
void QIconDragEvent_setAccepted_default(void* self, int accepted) {
	((QIconDragEvent*)self)->QIconDragEvent::setAccepted(accepted);
}
void* QIconDragEvent__Go_new(long long id) { return new QIconDragEvent__Go(id); }
void QIconDragEvent_delete(void* self) { delete (QIconDragEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qwindow.h>
#include "gen_qchildwindowevent.h"

extern "C" {

void* QChildWindowEvent_base(void* self) { return (void*)static_cast<QEvent*>((QChildWindowEvent*)self); }
void* QChildWindowEvent_new(int type_, void* childWindow) {
	return new QChildWindowEvent((QEvent::Type)(type_), (QWindow*)childWindow);
}
void* QChildWindowEvent_child(void* self) {
	auto&& __r = ((QChildWindowEvent*)self)->child();
	return (void*)(__r);
}
int QChildWindowEvent_registerEventType() {
	auto&& __r = QChildWindowEvent::registerEventType();
	return __r;
}
int QChildWindowEvent_registerEventType2(int hint) {
	auto&& __r = QChildWindowEvent::registerEventType(hint);
	return __r;
}
void QChildWindowEvent_delete(void* self) { delete (QChildWindowEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtGui/qbrush.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstring.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtCore/qnamespace.h>
#include "gen_qtextframeformat.h"

extern "C" {

void* QTextFrameFormat_base(void* self) { return (void*)static_cast<QTextFormat*>((QTextFrameFormat*)self); }
void* QTextFrameFormat_new() {
	return new QTextFrameFormat();
}
void* QTextFrameFormat_new2(void* arg1) {
	return new QTextFrameFormat(*(QTextFrameFormat*)arg1);
}
int QTextFrameFormat_isValid(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->isValid();
	return __r;
}
void QTextFrameFormat_setPosition(void* self, int f) {
	((QTextFrameFormat*)self)->setPosition((QTextFrameFormat::Position)(f));
}
int QTextFrameFormat_position(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->position();
	return (int)(__r);
}
void QTextFrameFormat_setBorder(void* self, double border) {
	((QTextFrameFormat*)self)->setBorder(border);
}
double QTextFrameFormat_border(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->border();
	return __r;
}
void QTextFrameFormat_setBorderBrush(void* self, void* brush) {
	((QTextFrameFormat*)self)->setBorderBrush(*(QBrush*)brush);
}
void* QTextFrameFormat_borderBrush(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->borderBrush();
	return new QBrush(__r);
}
void QTextFrameFormat_setBorderStyle(void* self, int style) {
	((QTextFrameFormat*)self)->setBorderStyle((QTextFrameFormat::BorderStyle)(style));
}
int QTextFrameFormat_borderStyle(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->borderStyle();
	return (int)(__r);
}
void QTextFrameFormat_setMargin(void* self, double margin) {
	((QTextFrameFormat*)self)->setMargin(margin);
}
double QTextFrameFormat_margin(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->margin();
	return __r;
}
void QTextFrameFormat_setTopMargin(void* self, double margin) {
	((QTextFrameFormat*)self)->setTopMargin(margin);
}
double QTextFrameFormat_topMargin(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->topMargin();
	return __r;
}
void QTextFrameFormat_setBottomMargin(void* self, double margin) {
	((QTextFrameFormat*)self)->setBottomMargin(margin);
}
double QTextFrameFormat_bottomMargin(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->bottomMargin();
	return __r;
}
void QTextFrameFormat_setLeftMargin(void* self, double margin) {
	((QTextFrameFormat*)self)->setLeftMargin(margin);
}
double QTextFrameFormat_leftMargin(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->leftMargin();
	return __r;
}
void QTextFrameFormat_setRightMargin(void* self, double margin) {
	((QTextFrameFormat*)self)->setRightMargin(margin);
}
double QTextFrameFormat_rightMargin(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->rightMargin();
	return __r;
}
void QTextFrameFormat_setPadding(void* self, double padding) {
	((QTextFrameFormat*)self)->setPadding(padding);
}
double QTextFrameFormat_padding(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->padding();
	return __r;
}
void QTextFrameFormat_setWidth(void* self, double width) {
	((QTextFrameFormat*)self)->setWidth(width);
}
void QTextFrameFormat_setWidth2(void* self, void* length) {
	((QTextFrameFormat*)self)->setWidth(*(QTextLength*)length);
}
void* QTextFrameFormat_width(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->width();
	return new QTextLength(__r);
}
void QTextFrameFormat_setHeight(void* self, double height) {
	((QTextFrameFormat*)self)->setHeight(height);
}
void QTextFrameFormat_setHeight2(void* self, void* height) {
	((QTextFrameFormat*)self)->setHeight(*(QTextLength*)height);
}
void* QTextFrameFormat_height(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->height();
	return new QTextLength(__r);
}
void QTextFrameFormat_setPageBreakPolicy(void* self, int flags) {
	((QTextFrameFormat*)self)->setPageBreakPolicy(QFlags<QTextFormat::PageBreakFlag>(flags));
}
int QTextFrameFormat_pageBreakPolicy(void* self) {
	auto&& __r = ((QTextFrameFormat*)self)->pageBreakPolicy();
	return (int)(__r);
}
void* QTextFrameFormat_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextFrameFormat*)self)->operator=(*(QTextFrameFormat*)arg1);
	return (void*)&(__r);
}
void QTextFrameFormat_delete(void* self) { delete (QTextFrameFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qpoint.h>
#include "gen_qaccessibleimageinterface.h"
#include "gen_purego_cb.h"

class QAccessibleImageInterface__Go : public QAccessibleImageInterface {
public:
	long long __goId;
	unsigned char __goOn[3] = {};
	template<typename... A> QAccessibleImageInterface__Go(long long id, A&&... a) : QAccessibleImageInterface(std::forward<A>(a)...), __goId(id) {}
	QString imageDescription() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, 503321968110705180LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	QSize imageSize() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -5595411163605219293LL, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	QPoint imagePosition() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7694988315546962491LL, &__h); if (__h) return *(QPoint*)__r; }
		return QPoint{};
	}
};

extern "C" {

void QAccessibleImageInterface__set_override(void* self, int idx) { ((QAccessibleImageInterface__Go*)self)->__goOn[idx] = 1; }
char* QAccessibleImageInterface_imageDescription(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleImageInterface*)self)->imageDescription();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleImageInterface_imageSize(void* self) {
	auto&& __r = ((QAccessibleImageInterface*)self)->imageSize();
	return new QSize(__r);
}
void* QAccessibleImageInterface_imagePosition(void* self) {
	auto&& __r = ((QAccessibleImageInterface*)self)->imagePosition();
	return new QPoint(__r);
}
void* QAccessibleImageInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleImageInterface*)self)->operator=(*(QAccessibleImageInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleImageInterface__Go_new(long long id) { return new QAccessibleImageInterface__Go(id); }
void QAccessibleImageInterface_delete(void* self) { delete (QAccessibleImageInterface*)self; }

}

#include <QtGui/qcursor.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qvariant.h>
#include <QtCore/qpoint.h>
#include <QtGui/qscreen.h>
#include "gen_qcursor.h"

extern "C" {

void* QCursor_new() {
	return new QCursor();
}
void* QCursor_new2(int shape) {
	return new QCursor((Qt::CursorShape)(shape));
}
void* QCursor_new3(void* bitmap, void* mask) {
	return new QCursor(*(QBitmap*)bitmap, *(QBitmap*)mask);
}
void* QCursor_new4(void* pixmap) {
	return new QCursor(*(QPixmap*)pixmap);
}
void* QCursor_new5(void* cursor) {
	return new QCursor(*(QCursor*)cursor);
}
void* QCursor_new6(void* bitmap, void* mask, int hotX) {
	return new QCursor(*(QBitmap*)bitmap, *(QBitmap*)mask, hotX);
}
void* QCursor_new7(void* bitmap, void* mask, int hotX, int hotY) {
	return new QCursor(*(QBitmap*)bitmap, *(QBitmap*)mask, hotX, hotY);
}
void* QCursor_new8(void* pixmap, int hotX) {
	return new QCursor(*(QPixmap*)pixmap, hotX);
}
void* QCursor_new9(void* pixmap, int hotX, int hotY) {
	return new QCursor(*(QPixmap*)pixmap, hotX, hotY);
}
void* QCursor_OperatorAssign(void* self, void* cursor) {
	auto&& __r = ((QCursor*)self)->operator=(*(QCursor*)cursor);
	return (void*)&(__r);
}
void QCursor_swap(void* self, void* other) {
	((QCursor*)self)->swap(*(QCursor*)other);
}
void* QCursor_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QCursor*)self);
	return new QVariant(__r);
}
int QCursor_shape(void* self) {
	auto&& __r = ((QCursor*)self)->shape();
	return (int)(__r);
}
void QCursor_setShape(void* self, int newShape) {
	((QCursor*)self)->setShape((Qt::CursorShape)(newShape));
}
void* QCursor_bitmap(void* self, int arg1) {
	auto&& __r = ((QCursor*)self)->bitmap((Qt::ReturnByValueConstant)(arg1));
	return new QBitmap(__r);
}
void* QCursor_mask(void* self, int arg1) {
	auto&& __r = ((QCursor*)self)->mask((Qt::ReturnByValueConstant)(arg1));
	return new QBitmap(__r);
}
void* QCursor_bitmap2(void* self) {
	auto&& __r = ((QCursor*)self)->bitmap();
	return new QBitmap(__r);
}
void* QCursor_mask2(void* self) {
	auto&& __r = ((QCursor*)self)->mask();
	return new QBitmap(__r);
}
void* QCursor_pixmap(void* self) {
	auto&& __r = ((QCursor*)self)->pixmap();
	return new QPixmap(__r);
}
void* QCursor_hotSpot(void* self) {
	auto&& __r = ((QCursor*)self)->hotSpot();
	return new QPoint(__r);
}
void* QCursor_pos() {
	auto&& __r = QCursor::pos();
	return new QPoint(__r);
}
void* QCursor_pos2(void* screen) {
	auto&& __r = QCursor::pos((QScreen*)screen);
	return new QPoint(__r);
}
void QCursor_setPos(int x, int y) {
	QCursor::setPos(x, y);
}
void QCursor_setPos2(void* screen, int x, int y) {
	QCursor::setPos((QScreen*)screen, x, y);
}
void QCursor_setPos3(void* p) {
	QCursor::setPos(*(QPoint*)p);
}
void QCursor_setPos4(void* screen, void* p) {
	QCursor::setPos((QScreen*)screen, *(QPoint*)p);
}
int QCursor_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QCursor*)self) == (*(QCursor*)rhs);
	return __r;
}
int QCursor_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QCursor*)self) != (*(QCursor*)rhs);
	return __r;
}
void* QCursor_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QCursor*)self);
	return new QVariant(__r);
}
void QCursor_delete(void* self) { delete (QCursor*)self; }

}

#include <QtGui/qmovie.h>
#include <QtCore/qobject.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qcolor.h>
#include <QtCore/qrect.h>
#include <QtGui/qimage.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qimagereader.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qmovie.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMovie__Go : public QMovie {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QMovie__Go(long long id, A&&... a) : QMovie(std::forward<A>(a)...), __goId(id) {}
	using QMovie::sender;
	using QMovie::senderSignalIndex;
	using QMovie::receivers;
	using QMovie::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMovie::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMovie::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMovie::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMovie::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMovie::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMovie::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QMovie::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QMovie::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QMovie::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QMovie::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMovie::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMovie::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMovie::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMovie;
static std::mutex __connsMu_QMovie;

extern "C" {

void* QMovie_base(void* self) { return (void*)static_cast<QObject*>((QMovie*)self); }
void QMovie__set_override(void* self, int idx) { ((QMovie__Go*)self)->__goOn[idx] = 1; }
void* QMovie_fromQObject(void* o) { return (void*)qobject_cast<QMovie*>((QObject*)o); }
void* QMovie_new(long long __goid) {
	return new QMovie__Go(__goid);
}
void* QMovie_new2(long long __goid, void* device) {
	return new QMovie__Go(__goid, (QIODevice*)device);
}
void* QMovie_new3(long long __goid, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QMovie__Go(__goid, fileName_qstr);
}
void* QMovie_new4(long long __goid, void* parent) {
	return new QMovie__Go(__goid, (QObject*)parent);
}
void* QMovie_new5(long long __goid, void* device, void* format) {
	return new QMovie__Go(__goid, (QIODevice*)device, *(QByteArray*)format);
}
void* QMovie_new6(long long __goid, void* device, void* format, void* parent) {
	return new QMovie__Go(__goid, (QIODevice*)device, *(QByteArray*)format, (QObject*)parent);
}
void* QMovie_new7(long long __goid, const char* fileName, long long fileName_len, void* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QMovie__Go(__goid, fileName_qstr, *(QByteArray*)format);
}
void* QMovie_new8(long long __goid, const char* fileName, long long fileName_len, void* format, void* parent) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QMovie__Go(__goid, fileName_qstr, *(QByteArray*)format, (QObject*)parent);
}
char* QMovie_tr(const char* s, long long* __out_len) {
	auto&& __r = QMovie::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QMovie_supportedFormats(long long* __out_len) {
	auto&& __r = QMovie::supportedFormats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void QMovie_setDevice(void* self, void* device) {
	((QMovie*)self)->setDevice((QIODevice*)device);
}
void* QMovie_device(void* self) {
	auto&& __r = ((QMovie*)self)->device();
	return (void*)(__r);
}
void QMovie_setFileName(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QMovie*)self)->setFileName(fileName_qstr);
}
char* QMovie_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QMovie*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMovie_setFormat(void* self, void* format) {
	((QMovie*)self)->setFormat(*(QByteArray*)format);
}
void* QMovie_format(void* self) {
	auto&& __r = ((QMovie*)self)->format();
	return new QByteArray(__r);
}
void QMovie_setBackgroundColor(void* self, void* color) {
	((QMovie*)self)->setBackgroundColor(*(QColor*)color);
}
void* QMovie_backgroundColor(void* self) {
	auto&& __r = ((QMovie*)self)->backgroundColor();
	return new QColor(__r);
}
int QMovie_state(void* self) {
	auto&& __r = ((QMovie*)self)->state();
	return (int)(__r);
}
void* QMovie_frameRect(void* self) {
	auto&& __r = ((QMovie*)self)->frameRect();
	return new QRect(__r);
}
void* QMovie_currentImage(void* self) {
	auto&& __r = ((QMovie*)self)->currentImage();
	return new QImage(__r);
}
void* QMovie_currentPixmap(void* self) {
	auto&& __r = ((QMovie*)self)->currentPixmap();
	return new QPixmap(__r);
}
int QMovie_isValid(void* self) {
	auto&& __r = ((QMovie*)self)->isValid();
	return __r;
}
int QMovie_lastError(void* self) {
	auto&& __r = ((QMovie*)self)->lastError();
	return (int)(__r);
}
char* QMovie_lastErrorString(void* self, long long* __out_len) {
	auto&& __r = ((QMovie*)self)->lastErrorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QMovie_jumpToFrame(void* self, int frameNumber) {
	auto&& __r = ((QMovie*)self)->jumpToFrame(frameNumber);
	return __r;
}
int QMovie_loopCount(void* self) {
	auto&& __r = ((QMovie*)self)->loopCount();
	return __r;
}
int QMovie_frameCount(void* self) {
	auto&& __r = ((QMovie*)self)->frameCount();
	return __r;
}
int QMovie_nextFrameDelay(void* self) {
	auto&& __r = ((QMovie*)self)->nextFrameDelay();
	return __r;
}
int QMovie_currentFrameNumber(void* self) {
	auto&& __r = ((QMovie*)self)->currentFrameNumber();
	return __r;
}
int QMovie_speed(void* self) {
	auto&& __r = ((QMovie*)self)->speed();
	return __r;
}
void* QMovie_scaledSize(void* self) {
	auto&& __r = ((QMovie*)self)->scaledSize();
	return new QSize(__r);
}
void QMovie_setScaledSize(void* self, void* size) {
	((QMovie*)self)->setScaledSize(*(QSize*)size);
}
int QMovie_cacheMode(void* self) {
	auto&& __r = ((QMovie*)self)->cacheMode();
	return (int)(__r);
}
void QMovie_setCacheMode(void* self, int mode) {
	((QMovie*)self)->setCacheMode((QMovie::CacheMode)(mode));
}
void QMovie_started(void* self) {
	((QMovie*)self)->started();
}
void QMovie_resized(void* self, void* size) {
	((QMovie*)self)->resized(*(QSize*)size);
}
void QMovie_updated(void* self, void* rect) {
	((QMovie*)self)->updated(*(QRect*)rect);
}
void QMovie_stateChanged(void* self, int state) {
	((QMovie*)self)->stateChanged((QMovie::MovieState)(state));
}
void QMovie_error(void* self, int error_) {
	((QMovie*)self)->error((QImageReader::ImageReaderError)(error_));
}
void QMovie_finished(void* self) {
	((QMovie*)self)->finished();
}
void QMovie_frameChanged(void* self, int frameNumber) {
	((QMovie*)self)->frameChanged(frameNumber);
}
void QMovie_start(void* self) {
	((QMovie*)self)->start();
}
int QMovie_jumpToNextFrame(void* self) {
	auto&& __r = ((QMovie*)self)->jumpToNextFrame();
	return __r;
}
void QMovie_setPaused(void* self, int paused) {
	((QMovie*)self)->setPaused(paused);
}
void QMovie_stop(void* self) {
	((QMovie*)self)->stop();
}
void QMovie_setSpeed(void* self, int percentSpeed) {
	((QMovie*)self)->setSpeed(percentSpeed);
}
void* QMovie_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMovie::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMovie_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMovie::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMovie_sender(void* self) {
	auto&& __r = ((QMovie__Go*)self)->sender();
	return (void*)(__r);
}
int QMovie_senderSignalIndex(void* self) {
	auto&& __r = ((QMovie__Go*)self)->senderSignalIndex();
	return __r;
}
int QMovie_receivers(void* self, const char* signal) {
	auto&& __r = ((QMovie__Go*)self)->receivers(signal);
	return __r;
}
int QMovie_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMovie__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QMovie_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMovie::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMovie_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMovie::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMovie_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMovie::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMovie_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMovie::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMovie_metaObject_default(void* self) {
	auto&& __r = ((QMovie*)self)->QMovie::metaObject();
	return (void*)(__r);
}
int QMovie_event_default(void* self, void* event) {
	auto&& __r = ((QMovie*)self)->QMovie::event((QEvent*)event);
	return __r;
}
int QMovie_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QMovie*)self)->QMovie::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QMovie_timerEvent_default(void* self, void* event) {
	((QMovie__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QMovie_childEvent_default(void* self, void* event) {
	((QMovie__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QMovie_customEvent_default(void* self, void* event) {
	((QMovie__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMovie_connectNotify_default(void* self, void* signal) {
	((QMovie__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMovie_disconnectNotify_default(void* self, void* signal) {
	((QMovie__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMovie_connect_started(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::started, (QMovie*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_resized(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::resized, (QMovie*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_updated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::updated, (QMovie*)self, [__id](const QRect& rect){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&rect);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::stateChanged, (QMovie*)self, [__id](QMovie::MovieState state){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(state));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_error(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::error, (QMovie*)self, [__id](QImageReader::ImageReaderError error_){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(error_));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::finished, (QMovie*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_frameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::frameChanged, (QMovie*)self, [__id](int frameNumber){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, frameNumber);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::destroyed, (QMovie*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::destroyed, (QMovie*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMovie*)self, &QMovie::objectNameChanged, (QMovie*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
	__conns_QMovie[__id] = __c;
}
void QMovie__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMovie);
		auto __it = __conns_QMovie.find(id);
		if (__it == __conns_QMovie.end()) return;
		__c = __it->second;
		__conns_QMovie.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMovie__Go_new(long long id, void* parent) { return new QMovie__Go(id, (QObject*)parent); }
void QMovie_delete(void* self) { delete (QMovie*)self; }

}

#include <QtGui/qpixmapcache.h>
#include <QtCore/qstring.h>
#include <QtGui/qpixmap.h>
#include "gen_qpixmapcache.h"

extern "C" {

int QPixmapCache_cacheLimit() {
	auto&& __r = QPixmapCache::cacheLimit();
	return __r;
}
void QPixmapCache_setCacheLimit(int arg1) {
	QPixmapCache::setCacheLimit(arg1);
}
int QPixmapCache_find(const char* key, long long key_len, void* pixmap) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QPixmapCache::find(key_qstr, (QPixmap*)pixmap);
	return __r;
}
int QPixmapCache_find2(void* key, void* pixmap) {
	auto&& __r = QPixmapCache::find(*(QPixmapCache::Key*)key, (QPixmap*)pixmap);
	return __r;
}
int QPixmapCache_insert(const char* key, long long key_len, void* pixmap) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QPixmapCache::insert(key_qstr, *(QPixmap*)pixmap);
	return __r;
}
void* QPixmapCache_insert2(void* pixmap) {
	auto&& __r = QPixmapCache::insert(*(QPixmap*)pixmap);
	return new QPixmapCache::Key(__r);
}
int QPixmapCache_replace(void* key, void* pixmap) {
	auto&& __r = QPixmapCache::replace(*(QPixmapCache::Key*)key, *(QPixmap*)pixmap);
	return __r;
}
void QPixmapCache_remove(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	QPixmapCache::remove(key_qstr);
}
void QPixmapCache_remove2(void* key) {
	QPixmapCache::remove(*(QPixmapCache::Key*)key);
}
void QPixmapCache_clear() {
	QPixmapCache::clear();
}
void QPixmapCache_delete(void* self) { delete (QPixmapCache*)self; }

}

#include <QtGui/qtexttable.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextcursor.h>
#include <QtGui/qtextformat.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtexttable.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextTable;
static std::mutex __connsMu_QTextTable;

extern "C" {

void* QTextTable_base(void* self) { return (void*)static_cast<QTextFrame*>((QTextTable*)self); }
void QTextTable_chain(void* self, void** out) {
	QTextTable* __p = (QTextTable*)self;
	QTextFrame* __b0 = static_cast<QTextFrame*>(__p); out[0] = (void*)__b0;
	QTextObject* __b1 = static_cast<QTextObject*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void* QTextTable_fromQObject(void* o) { return (void*)qobject_cast<QTextTable*>((QObject*)o); }
void* QTextTable_new(void* doc) {
	return new QTextTable((QTextDocument*)doc);
}
char* QTextTable_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextTable::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextTable_resize(void* self, int rows, int cols) {
	((QTextTable*)self)->resize(rows, cols);
}
void QTextTable_insertRows(void* self, int pos, int num) {
	((QTextTable*)self)->insertRows(pos, num);
}
void QTextTable_insertColumns(void* self, int pos, int num) {
	((QTextTable*)self)->insertColumns(pos, num);
}
void QTextTable_appendRows(void* self, int count) {
	((QTextTable*)self)->appendRows(count);
}
void QTextTable_appendColumns(void* self, int count) {
	((QTextTable*)self)->appendColumns(count);
}
void QTextTable_removeRows(void* self, int pos, int num) {
	((QTextTable*)self)->removeRows(pos, num);
}
void QTextTable_removeColumns(void* self, int pos, int num) {
	((QTextTable*)self)->removeColumns(pos, num);
}
void QTextTable_mergeCells(void* self, int row, int col, int numRows, int numCols) {
	((QTextTable*)self)->mergeCells(row, col, numRows, numCols);
}
void QTextTable_mergeCells2(void* self, void* cursor) {
	((QTextTable*)self)->mergeCells(*(QTextCursor*)cursor);
}
void QTextTable_splitCell(void* self, int row, int col, int numRows, int numCols) {
	((QTextTable*)self)->splitCell(row, col, numRows, numCols);
}
int QTextTable_rows(void* self) {
	auto&& __r = ((QTextTable*)self)->rows();
	return __r;
}
int QTextTable_columns(void* self) {
	auto&& __r = ((QTextTable*)self)->columns();
	return __r;
}
void* QTextTable_cellAt(void* self, int row, int col) {
	auto&& __r = ((QTextTable*)self)->cellAt(row, col);
	return new QTextTableCell(__r);
}
void* QTextTable_cellAt2(void* self, int position) {
	auto&& __r = ((QTextTable*)self)->cellAt(position);
	return new QTextTableCell(__r);
}
void* QTextTable_cellAt3(void* self, void* c) {
	auto&& __r = ((QTextTable*)self)->cellAt(*(QTextCursor*)c);
	return new QTextTableCell(__r);
}
void* QTextTable_rowStart(void* self, void* c) {
	auto&& __r = ((QTextTable*)self)->rowStart(*(QTextCursor*)c);
	return new QTextCursor(__r);
}
void* QTextTable_rowEnd(void* self, void* c) {
	auto&& __r = ((QTextTable*)self)->rowEnd(*(QTextCursor*)c);
	return new QTextCursor(__r);
}
void QTextTable_setFormat(void* self, void* format) {
	((QTextTable*)self)->setFormat(*(QTextTableFormat*)format);
}
void* QTextTable_format(void* self) {
	auto&& __r = ((QTextTable*)self)->format();
	return new QTextTableFormat(__r);
}
void* QTextTable_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextTable::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextTable_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextTable::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QTextTable_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextTable::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextTable_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextTable::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextTable_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextTable::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextTable_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextTable::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextTable_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextTable*)self, &QTextTable::destroyed, (QTextTable*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextTable);
	__conns_QTextTable[__id] = __c;
}
void QTextTable_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextTable*)self, &QTextTable::destroyed, (QTextTable*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextTable);
	__conns_QTextTable[__id] = __c;
}
void QTextTable_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextTable*)self, &QTextTable::objectNameChanged, (QTextTable*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextTable);
	__conns_QTextTable[__id] = __c;
}
void QTextTable__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextTable);
		auto __it = __conns_QTextTable.find(id);
		if (__it == __conns_QTextTable.end()) return;
		__c = __it->second;
		__conns_QTextTable.erase(__it);
	}
	QObject::disconnect(__c);
}
void QTextTable_delete(void* self) { delete (QTextTable*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qtabbar.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include "gen_qstyleoptiontabwidgetframe.h"

extern "C" {

void* QStyleOptionTabWidgetFrame_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionTabWidgetFrame*)self); }
void* QStyleOptionTabWidgetFrame_new() {
	return new QStyleOptionTabWidgetFrame();
}
void* QStyleOptionTabWidgetFrame_new2(void* other) {
	return new QStyleOptionTabWidgetFrame(*(QStyleOptionTabWidgetFrame*)other);
}
void* QStyleOptionTabWidgetFrame_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionTabWidgetFrame*)self)->operator=(*(QStyleOptionTabWidgetFrame*)arg1);
	return (void*)&(__r);
}
int QStyleOptionTabWidgetFrame_lineWidth_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->lineWidth;
	return __r;
}
void QStyleOptionTabWidgetFrame_lineWidth_set(void* self, int lineWidth) {
	((QStyleOptionTabWidgetFrame*)self)->lineWidth = lineWidth;
}
int QStyleOptionTabWidgetFrame_midLineWidth_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->midLineWidth;
	return __r;
}
void QStyleOptionTabWidgetFrame_midLineWidth_set(void* self, int midLineWidth) {
	((QStyleOptionTabWidgetFrame*)self)->midLineWidth = midLineWidth;
}
int QStyleOptionTabWidgetFrame_shape_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->shape;
	return (int)(__r);
}
void QStyleOptionTabWidgetFrame_shape_set(void* self, int shape) {
	((QStyleOptionTabWidgetFrame*)self)->shape = (QTabBar::Shape)(shape);
}
void* QStyleOptionTabWidgetFrame_tabBarSize_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->tabBarSize;
	return new QSize(__r);
}
void QStyleOptionTabWidgetFrame_tabBarSize_set(void* self, void* tabBarSize) {
	((QStyleOptionTabWidgetFrame*)self)->tabBarSize = *(QSize*)tabBarSize;
}
void* QStyleOptionTabWidgetFrame_rightCornerWidgetSize_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->rightCornerWidgetSize;
	return new QSize(__r);
}
void QStyleOptionTabWidgetFrame_rightCornerWidgetSize_set(void* self, void* rightCornerWidgetSize) {
	((QStyleOptionTabWidgetFrame*)self)->rightCornerWidgetSize = *(QSize*)rightCornerWidgetSize;
}
void* QStyleOptionTabWidgetFrame_leftCornerWidgetSize_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->leftCornerWidgetSize;
	return new QSize(__r);
}
void QStyleOptionTabWidgetFrame_leftCornerWidgetSize_set(void* self, void* leftCornerWidgetSize) {
	((QStyleOptionTabWidgetFrame*)self)->leftCornerWidgetSize = *(QSize*)leftCornerWidgetSize;
}
void* QStyleOptionTabWidgetFrame_tabBarRect_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->tabBarRect;
	return new QRect(__r);
}
void QStyleOptionTabWidgetFrame_tabBarRect_set(void* self, void* tabBarRect) {
	((QStyleOptionTabWidgetFrame*)self)->tabBarRect = *(QRect*)tabBarRect;
}
void* QStyleOptionTabWidgetFrame_selectedTabRect_get(void* self) {
	auto __r = ((QStyleOptionTabWidgetFrame*)self)->selectedTabRect;
	return new QRect(__r);
}
void QStyleOptionTabWidgetFrame_selectedTabRect_set(void* self, void* selectedTabRect) {
	((QStyleOptionTabWidgetFrame*)self)->selectedTabRect = *(QRect*)selectedTabRect;
}
void QStyleOptionTabWidgetFrame_delete(void* self) { delete (QStyleOptionTabWidgetFrame*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtGui/qregion.h>
#include "gen_qstylehintreturnmask.h"

extern "C" {

void* QStyleHintReturnMask_base(void* self) { return (void*)static_cast<QStyleHintReturn*>((QStyleHintReturnMask*)self); }
void* QStyleHintReturnMask_new() {
	return new QStyleHintReturnMask();
}
void* QStyleHintReturnMask_new2(void* arg1) {
	return new QStyleHintReturnMask(*(QStyleHintReturnMask*)arg1);
}
void* QStyleHintReturnMask_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleHintReturnMask*)self)->operator=(*(QStyleHintReturnMask*)arg1);
	return (void*)&(__r);
}
void* QStyleHintReturnMask_region_get(void* self) {
	auto __r = ((QStyleHintReturnMask*)self)->region;
	return new QRegion(__r);
}
void QStyleHintReturnMask_region_set(void* self, void* region) {
	((QStyleHintReturnMask*)self)->region = *(QRegion*)region;
}
void QStyleHintReturnMask_delete(void* self) { delete (QStyleHintReturnMask*)self; }

}

#include <QtWidgets/qcombobox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qlineedit.h>
#include <QtGui/qvalidator.h>
#include <QtWidgets/qcompleter.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtGui/qicon.h>
#include <QtWidgets/qabstractitemview.h>
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
#include "gen_qcombobox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QComboBox__Go : public QComboBox {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QComboBox__Go(long long id, A&&... a) : QComboBox(std::forward<A>(a)...), __goId(id) {}
	using QComboBox::updateMicroFocus;
	using QComboBox::create;
	using QComboBox::destroy;
	using QComboBox::focusNextChild;
	using QComboBox::focusPreviousChild;
	using QComboBox::sender;
	using QComboBox::senderSignalIndex;
	using QComboBox::receivers;
	using QComboBox::isSignalConnected;
	using QComboBox::getDecodedMetricF;
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QComboBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QComboBox::focusOutEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QComboBox::changeEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QComboBox::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QComboBox::paintEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QComboBox::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QComboBox::hideEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QComboBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QComboBox::mouseReleaseEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QComboBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QComboBox::keyReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QComboBox::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QComboBox::contextMenuEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QComboBox::inputMethodEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionComboBox* __a0) { QComboBox::initStyleOption(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QComboBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QComboBox::mouseMoveEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QComboBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QComboBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QComboBox::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QComboBox::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QComboBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QComboBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QComboBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QComboBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QComboBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QComboBox::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QComboBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QComboBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QComboBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QComboBox::sharedPainter(); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QComboBox::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QComboBox::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QComboBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QComboBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QComboBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QComboBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QComboBox::metaObject();
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QComboBox::setModel(model);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QComboBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QComboBox::minimumSizeHint();
	}
	void showPopup() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -977583612921719482LL, &__h); if (__h) return; }
		QComboBox::showPopup();
	}
	void hidePopup() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 3308350117881191909LL, &__h); if (__h) return; }
		QComboBox::hidePopup();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QComboBox::event(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QComboBox::inputMethodQuery(arg1);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QComboBox::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QComboBox::focusOutEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QComboBox::changeEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QComboBox::resizeEvent(e);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QComboBox::paintEvent(e);
	}
	void showEvent(QShowEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)e, &__h); if (__h) return; }
		QComboBox::showEvent(e);
	}
	void hideEvent(QHideEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)e, &__h); if (__h) return; }
		QComboBox::hideEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QComboBox::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QComboBox::mouseReleaseEvent(e);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QComboBox::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QComboBox::keyReleaseEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QComboBox::wheelEvent(e);
	}
	void contextMenuEvent(QContextMenuEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)e, &__h); if (__h) return; }
		QComboBox::contextMenuEvent(e);
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QComboBox::inputMethodEvent(arg1);
	}
	void initStyleOption(QStyleOptionComboBox* option) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QComboBox::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QComboBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QComboBox::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QComboBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QComboBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QComboBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QComboBox::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QComboBox::mouseMoveEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QComboBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QComboBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QComboBox::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QComboBox::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QComboBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QComboBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QComboBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QComboBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QComboBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QComboBox::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QComboBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QComboBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QComboBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QComboBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QComboBox::sharedPainter();
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QComboBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QComboBox::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QComboBox::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QComboBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QComboBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QComboBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QComboBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QComboBox;
static std::mutex __connsMu_QComboBox;

extern "C" {

void* QComboBox_base(void* self) { return (void*)static_cast<QWidget*>((QComboBox*)self); }
void QComboBox_chain(void* self, void** out) {
	QComboBox* __p = (QComboBox*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QComboBox__set_override(void* self, int idx) { ((QComboBox__Go*)self)->__goOn[idx] = 1; }
void* QComboBox_fromQObject(void* o) { return (void*)qobject_cast<QComboBox*>((QObject*)o); }
void* QComboBox_new(long long __goid, void* parent) {
	return new QComboBox__Go(__goid, (QWidget*)parent);
}
void* QComboBox_new2(long long __goid) {
	return new QComboBox__Go(__goid);
}
char* QComboBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QComboBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QComboBox_maxVisibleItems(void* self) {
	auto&& __r = ((QComboBox*)self)->maxVisibleItems();
	return __r;
}
void QComboBox_setMaxVisibleItems(void* self, int maxItems) {
	((QComboBox*)self)->setMaxVisibleItems(maxItems);
}
int QComboBox_count(void* self) {
	auto&& __r = ((QComboBox*)self)->count();
	return __r;
}
void QComboBox_setMaxCount(void* self, int max) {
	((QComboBox*)self)->setMaxCount(max);
}
int QComboBox_maxCount(void* self) {
	auto&& __r = ((QComboBox*)self)->maxCount();
	return __r;
}
int QComboBox_duplicatesEnabled(void* self) {
	auto&& __r = ((QComboBox*)self)->duplicatesEnabled();
	return __r;
}
void QComboBox_setDuplicatesEnabled(void* self, int enable) {
	((QComboBox*)self)->setDuplicatesEnabled(enable);
}
void QComboBox_setFrame(void* self, int arg1) {
	((QComboBox*)self)->setFrame(arg1);
}
int QComboBox_hasFrame(void* self) {
	auto&& __r = ((QComboBox*)self)->hasFrame();
	return __r;
}
int QComboBox_findText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QComboBox*)self)->findText(text_qstr);
	return __r;
}
int QComboBox_findData(void* self, void* data) {
	auto&& __r = ((QComboBox*)self)->findData(*(QVariant*)data);
	return __r;
}
int QComboBox_insertPolicy(void* self) {
	auto&& __r = ((QComboBox*)self)->insertPolicy();
	return (int)(__r);
}
void QComboBox_setInsertPolicy(void* self, int policy) {
	((QComboBox*)self)->setInsertPolicy((QComboBox::InsertPolicy)(policy));
}
int QComboBox_sizeAdjustPolicy(void* self) {
	auto&& __r = ((QComboBox*)self)->sizeAdjustPolicy();
	return (int)(__r);
}
void QComboBox_setSizeAdjustPolicy(void* self, int policy) {
	((QComboBox*)self)->setSizeAdjustPolicy((QComboBox::SizeAdjustPolicy)(policy));
}
int QComboBox_minimumContentsLength(void* self) {
	auto&& __r = ((QComboBox*)self)->minimumContentsLength();
	return __r;
}
void QComboBox_setMinimumContentsLength(void* self, int characters) {
	((QComboBox*)self)->setMinimumContentsLength(characters);
}
void* QComboBox_iconSize(void* self) {
	auto&& __r = ((QComboBox*)self)->iconSize();
	return new QSize(__r);
}
void QComboBox_setIconSize(void* self, void* size) {
	((QComboBox*)self)->setIconSize(*(QSize*)size);
}
void QComboBox_setPlaceholderText(void* self, const char* placeholderText, long long placeholderText_len) {
	QString placeholderText_qstr = QString::fromUtf8(placeholderText, placeholderText_len);
	((QComboBox*)self)->setPlaceholderText(placeholderText_qstr);
}
char* QComboBox_placeholderText(void* self, long long* __out_len) {
	auto&& __r = ((QComboBox*)self)->placeholderText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QComboBox_isEditable(void* self) {
	auto&& __r = ((QComboBox*)self)->isEditable();
	return __r;
}
void QComboBox_setEditable(void* self, int editable) {
	((QComboBox*)self)->setEditable(editable);
}
void QComboBox_setLineEdit(void* self, void* edit) {
	((QComboBox*)self)->setLineEdit((QLineEdit*)edit);
}
void* QComboBox_lineEdit(void* self) {
	auto&& __r = ((QComboBox*)self)->lineEdit();
	return (void*)(__r);
}
void QComboBox_setValidator(void* self, void* v) {
	((QComboBox*)self)->setValidator((QValidator*)v);
}
void* QComboBox_validator(void* self) {
	auto&& __r = ((QComboBox*)self)->validator();
	return (void*)(__r);
}
void QComboBox_setCompleter(void* self, void* c) {
	((QComboBox*)self)->setCompleter((QCompleter*)c);
}
void* QComboBox_completer(void* self) {
	auto&& __r = ((QComboBox*)self)->completer();
	return (void*)(__r);
}
void* QComboBox_itemDelegate(void* self) {
	auto&& __r = ((QComboBox*)self)->itemDelegate();
	return (void*)(__r);
}
void QComboBox_setItemDelegate(void* self, void* delegate) {
	((QComboBox*)self)->setItemDelegate((QAbstractItemDelegate*)delegate);
}
void* QComboBox_model(void* self) {
	auto&& __r = ((QComboBox*)self)->model();
	return (void*)(__r);
}
void QComboBox_setModel(void* self, void* model) {
	((QComboBox*)self)->setModel((QAbstractItemModel*)model);
}
void* QComboBox_rootModelIndex(void* self) {
	auto&& __r = ((QComboBox*)self)->rootModelIndex();
	return new QModelIndex(__r);
}
void QComboBox_setRootModelIndex(void* self, void* index) {
	((QComboBox*)self)->setRootModelIndex(*(QModelIndex*)index);
}
int QComboBox_modelColumn(void* self) {
	auto&& __r = ((QComboBox*)self)->modelColumn();
	return __r;
}
void QComboBox_setModelColumn(void* self, int visibleColumn) {
	((QComboBox*)self)->setModelColumn(visibleColumn);
}
int QComboBox_labelDrawingMode(void* self) {
	auto&& __r = ((QComboBox*)self)->labelDrawingMode();
	return (int)(__r);
}
void QComboBox_setLabelDrawingMode(void* self, int labelDrawing) {
	((QComboBox*)self)->setLabelDrawingMode((QComboBox::LabelDrawingMode)(labelDrawing));
}
int QComboBox_currentIndex(void* self) {
	auto&& __r = ((QComboBox*)self)->currentIndex();
	return __r;
}
char* QComboBox_currentText(void* self, long long* __out_len) {
	auto&& __r = ((QComboBox*)self)->currentText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QComboBox_currentData(void* self) {
	auto&& __r = ((QComboBox*)self)->currentData();
	return new QVariant(__r);
}
char* QComboBox_itemText(void* self, int index, long long* __out_len) {
	auto&& __r = ((QComboBox*)self)->itemText(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QComboBox_itemIcon(void* self, int index) {
	auto&& __r = ((QComboBox*)self)->itemIcon(index);
	return new QIcon(__r);
}
void* QComboBox_itemData(void* self, int index) {
	auto&& __r = ((QComboBox*)self)->itemData(index);
	return new QVariant(__r);
}
void QComboBox_addItem(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->addItem(text_qstr);
}
void QComboBox_addItem2(void* self, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->addItem(*(QIcon*)icon, text_qstr);
}
void QComboBox_addItems(void* self, char** texts, long long texts_len) {
	QList<QString> texts_list; texts_list.reserve(texts_len);
	for (long long __i = 0; __i < texts_len; __i++) texts_list.append(QString::fromUtf8(texts[__i]));
	((QComboBox*)self)->addItems(texts_list);
}
void QComboBox_insertItem(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->insertItem(index, text_qstr);
}
void QComboBox_insertItem2(void* self, int index, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->insertItem(index, *(QIcon*)icon, text_qstr);
}
void QComboBox_insertItems(void* self, int index, char** texts, long long texts_len) {
	QList<QString> texts_list; texts_list.reserve(texts_len);
	for (long long __i = 0; __i < texts_len; __i++) texts_list.append(QString::fromUtf8(texts[__i]));
	((QComboBox*)self)->insertItems(index, texts_list);
}
void QComboBox_insertSeparator(void* self, int index) {
	((QComboBox*)self)->insertSeparator(index);
}
void QComboBox_removeItem(void* self, int index) {
	((QComboBox*)self)->removeItem(index);
}
void QComboBox_setItemText(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->setItemText(index, text_qstr);
}
void QComboBox_setItemIcon(void* self, int index, void* icon) {
	((QComboBox*)self)->setItemIcon(index, *(QIcon*)icon);
}
void QComboBox_setItemData(void* self, int index, void* value) {
	((QComboBox*)self)->setItemData(index, *(QVariant*)value);
}
void* QComboBox_view(void* self) {
	auto&& __r = ((QComboBox*)self)->view();
	return (void*)(__r);
}
void QComboBox_setView(void* self, void* itemView) {
	((QComboBox*)self)->setView((QAbstractItemView*)itemView);
}
void QComboBox_showPopup(void* self) {
	((QComboBox*)self)->showPopup();
}
void QComboBox_hidePopup(void* self) {
	((QComboBox*)self)->hidePopup();
}
void* QComboBox_inputMethodQuery(void* self, int arg1) {
	auto&& __r = ((QComboBox*)self)->inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void* QComboBox_inputMethodQuery2(void* self, int query, void* argument) {
	auto&& __r = ((QComboBox*)self)->inputMethodQuery((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void QComboBox_clear(void* self) {
	((QComboBox*)self)->clear();
}
void QComboBox_clearEditText(void* self) {
	((QComboBox*)self)->clearEditText();
}
void QComboBox_setEditText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->setEditText(text_qstr);
}
void QComboBox_setCurrentIndex(void* self, int index) {
	((QComboBox*)self)->setCurrentIndex(index);
}
void QComboBox_setCurrentText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->setCurrentText(text_qstr);
}
void QComboBox_editTextChanged(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QComboBox*)self)->editTextChanged(arg1_qstr);
}
void QComboBox_activated(void* self, int index) {
	((QComboBox*)self)->activated(index);
}
void QComboBox_textActivated(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QComboBox*)self)->textActivated(arg1_qstr);
}
void QComboBox_highlighted(void* self, int index) {
	((QComboBox*)self)->highlighted(index);
}
void QComboBox_textHighlighted(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QComboBox*)self)->textHighlighted(arg1_qstr);
}
void QComboBox_currentIndexChanged(void* self, int index) {
	((QComboBox*)self)->currentIndexChanged(index);
}
void QComboBox_currentTextChanged(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QComboBox*)self)->currentTextChanged(arg1_qstr);
}
void QComboBox_setTabOrder(void* arg1, void* arg2) {
	QComboBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QComboBox_mouseGrabber() {
	auto&& __r = QComboBox::mouseGrabber();
	return (void*)(__r);
}
void* QComboBox_keyboardGrabber() {
	auto&& __r = QComboBox::keyboardGrabber();
	return (void*)(__r);
}
void* QComboBox_find(unsigned long long arg1) {
	auto&& __r = QComboBox::find(arg1);
	return (void*)(__r);
}
void* QComboBox_createWindowContainer(void* window) {
	auto&& __r = QComboBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QComboBox_updateMicroFocus(void* self) {
	((QComboBox__Go*)self)->updateMicroFocus();
}
void QComboBox_create(void* self) {
	((QComboBox__Go*)self)->create();
}
void QComboBox_destroy(void* self) {
	((QComboBox__Go*)self)->destroy();
}
int QComboBox_focusNextChild(void* self) {
	auto&& __r = ((QComboBox__Go*)self)->focusNextChild();
	return __r;
}
int QComboBox_focusPreviousChild(void* self) {
	auto&& __r = ((QComboBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QComboBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QComboBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QComboBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QComboBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QComboBox_sender(void* self) {
	auto&& __r = ((QComboBox__Go*)self)->sender();
	return (void*)(__r);
}
int QComboBox_senderSignalIndex(void* self) {
	auto&& __r = ((QComboBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QComboBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QComboBox__Go*)self)->receivers(signal);
	return __r;
}
int QComboBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QComboBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QComboBox_devicePixelRatioFScale() {
	auto&& __r = QComboBox::devicePixelRatioFScale();
	return __r;
}
int QComboBox_encodeMetricF(int metric, double value) {
	auto&& __r = QComboBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QComboBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QComboBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QComboBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QComboBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QComboBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QComboBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QComboBox_findText2(void* self, const char* text, long long text_len, int flags) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QComboBox*)self)->findText(text_qstr, QFlags<Qt::MatchFlag>(flags));
	return __r;
}
int QComboBox_findData2(void* self, void* data, int role) {
	auto&& __r = ((QComboBox*)self)->findData(*(QVariant*)data, role);
	return __r;
}
int QComboBox_findData3(void* self, void* data, int role, int flags) {
	auto&& __r = ((QComboBox*)self)->findData(*(QVariant*)data, role, QFlags<Qt::MatchFlag>(flags));
	return __r;
}
void* QComboBox_currentData2(void* self, int role) {
	auto&& __r = ((QComboBox*)self)->currentData(role);
	return new QVariant(__r);
}
void* QComboBox_itemData2(void* self, int index, int role) {
	auto&& __r = ((QComboBox*)self)->itemData(index, role);
	return new QVariant(__r);
}
void QComboBox_addItem3(void* self, const char* text, long long text_len, void* userData) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->addItem(text_qstr, *(QVariant*)userData);
}
void QComboBox_addItem4(void* self, void* icon, const char* text, long long text_len, void* userData) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->addItem(*(QIcon*)icon, text_qstr, *(QVariant*)userData);
}
void QComboBox_insertItem3(void* self, int index, const char* text, long long text_len, void* userData) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->insertItem(index, text_qstr, *(QVariant*)userData);
}
void QComboBox_insertItem4(void* self, int index, void* icon, const char* text, long long text_len, void* userData) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QComboBox*)self)->insertItem(index, *(QIcon*)icon, text_qstr, *(QVariant*)userData);
}
void QComboBox_setItemData2(void* self, int index, void* value, int role) {
	((QComboBox*)self)->setItemData(index, *(QVariant*)value, role);
}
void* QComboBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QComboBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QComboBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QComboBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QComboBox_updateMicroFocus2(void* self, int query) {
	((QComboBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QComboBox_create2(void* self, unsigned long long arg1) {
	((QComboBox__Go*)self)->create(arg1);
}
void QComboBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QComboBox__Go*)self)->create(arg1, initializeWindow);
}
void QComboBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QComboBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QComboBox_destroy2(void* self, int destroyWindow) {
	((QComboBox__Go*)self)->destroy(destroyWindow);
}
void QComboBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QComboBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QComboBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QComboBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QComboBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QComboBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QComboBox_setModel_default(void* self, void* model) {
	((QComboBox*)self)->QComboBox::setModel((QAbstractItemModel*)model);
}
void QComboBox_showPopup_default(void* self) {
	((QComboBox*)self)->QComboBox::showPopup();
}
void QComboBox_hidePopup_default(void* self) {
	((QComboBox*)self)->QComboBox::hidePopup();
}
void* QComboBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QComboBox*)self)->QComboBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void* QComboBox_inputMethodQuery_default2(void* self, int query, void* argument) {
	auto&& __r = ((QComboBox*)self)->QComboBox::inputMethodQuery((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void* QComboBox_metaObject_default(void* self) {
	auto&& __r = ((QComboBox*)self)->QComboBox::metaObject();
	return (void*)(__r);
}
void* QComboBox_sizeHint_default(void* self) {
	auto&& __r = ((QComboBox*)self)->QComboBox::sizeHint();
	return new QSize(__r);
}
void* QComboBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QComboBox*)self)->QComboBox::minimumSizeHint();
	return new QSize(__r);
}
int QComboBox_event_default(void* self, void* event) {
	auto&& __r = ((QComboBox*)self)->QComboBox::event((QEvent*)event);
	return __r;
}
void QComboBox_focusInEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QComboBox_focusOutEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QComboBox_changeEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QComboBox_resizeEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QComboBox_paintEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QComboBox_showEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)e);
}
void QComboBox_hideEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)e);
}
void QComboBox_mousePressEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QComboBox_mouseReleaseEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QComboBox_keyPressEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QComboBox_keyReleaseEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QComboBox_wheelEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QComboBox_contextMenuEvent_default(void* self, void* e) {
	((QComboBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)e);
}
void QComboBox_inputMethodEvent_default(void* self, void* arg1) {
	((QComboBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void QComboBox_initStyleOption_default(void* self, void* option) {
	((QComboBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionComboBox*)option);
}
int QComboBox_devType_default(void* self) {
	auto&& __r = ((QComboBox*)self)->QComboBox::devType();
	return __r;
}
void QComboBox_setVisible_default(void* self, int visible) {
	((QComboBox*)self)->QComboBox::setVisible(visible);
}
int QComboBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QComboBox*)self)->QComboBox::heightForWidth(arg1);
	return __r;
}
int QComboBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QComboBox*)self)->QComboBox::hasHeightForWidth();
	return __r;
}
void* QComboBox_paintEngine_default(void* self) {
	auto&& __r = ((QComboBox*)self)->QComboBox::paintEngine();
	return (void*)(__r);
}
void QComboBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QComboBox_mouseMoveEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QComboBox_enterEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QComboBox_leaveEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QComboBox_moveEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QComboBox_closeEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QComboBox_tabletEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QComboBox_actionEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QComboBox_dragEnterEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QComboBox_dragMoveEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QComboBox_dragLeaveEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QComboBox_dropEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QComboBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QComboBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QComboBox_initPainter_default(void* self, void* painter) {
	((QComboBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QComboBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QComboBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QComboBox_sharedPainter_default(void* self) {
	auto&& __r = ((QComboBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
int QComboBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QComboBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QComboBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QComboBox*)self)->QComboBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QComboBox_timerEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QComboBox_childEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QComboBox_customEvent_default(void* self, void* event) {
	((QComboBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QComboBox_connectNotify_default(void* self, void* signal) {
	((QComboBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QComboBox_disconnectNotify_default(void* self, void* signal) {
	((QComboBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QComboBox_connect_editTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::editTextChanged, (QComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::activated, (QComboBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_textActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::textActivated, (QComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_highlighted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::highlighted, (QComboBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_textHighlighted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::textHighlighted, (QComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_currentIndexChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::currentIndexChanged, (QComboBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_currentTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::currentTextChanged, (QComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::windowTitleChanged, (QComboBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::windowIconChanged, (QComboBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::windowIconTextChanged, (QComboBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::customContextMenuRequested, (QComboBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::destroyed, (QComboBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::destroyed, (QComboBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QComboBox*)self, &QComboBox::objectNameChanged, (QComboBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
	__conns_QComboBox[__id] = __c;
}
void QComboBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QComboBox);
		auto __it = __conns_QComboBox.find(id);
		if (__it == __conns_QComboBox.end()) return;
		__c = __it->second;
		__conns_QComboBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QComboBox__Go_new(long long id) { return new QComboBox__Go(id); }
void QComboBox_delete(void* self) { delete (QComboBox*)self; }

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
#include "gen_qswipegesture.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSwipeGesture__Go : public QSwipeGesture {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QSwipeGesture__Go(long long id, A&&... a) : QSwipeGesture(std::forward<A>(a)...), __goId(id) {}
	using QSwipeGesture::sender;
	using QSwipeGesture::senderSignalIndex;
	using QSwipeGesture::receivers;
	using QSwipeGesture::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSwipeGesture::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSwipeGesture::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSwipeGesture::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSwipeGesture::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSwipeGesture::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSwipeGesture::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSwipeGesture::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSwipeGesture::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSwipeGesture::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSwipeGesture::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSwipeGesture::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSwipeGesture::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSwipeGesture::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSwipeGesture;
static std::mutex __connsMu_QSwipeGesture;

extern "C" {

void* QSwipeGesture_base(void* self) { return (void*)static_cast<QGesture*>((QSwipeGesture*)self); }
void QSwipeGesture_chain(void* self, void** out) {
	QSwipeGesture* __p = (QSwipeGesture*)self;
	QGesture* __b0 = static_cast<QGesture*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QSwipeGesture__set_override(void* self, int idx) { ((QSwipeGesture__Go*)self)->__goOn[idx] = 1; }
void* QSwipeGesture_fromQObject(void* o) { return (void*)qobject_cast<QSwipeGesture*>((QObject*)o); }
void* QSwipeGesture_new(long long __goid) {
	return new QSwipeGesture__Go(__goid);
}
void* QSwipeGesture_new2(long long __goid, void* parent) {
	return new QSwipeGesture__Go(__goid, (QObject*)parent);
}
char* QSwipeGesture_tr(const char* s, long long* __out_len) {
	auto&& __r = QSwipeGesture::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSwipeGesture_horizontalDirection(void* self) {
	auto&& __r = ((QSwipeGesture*)self)->horizontalDirection();
	return (int)(__r);
}
int QSwipeGesture_verticalDirection(void* self) {
	auto&& __r = ((QSwipeGesture*)self)->verticalDirection();
	return (int)(__r);
}
double QSwipeGesture_swipeAngle(void* self) {
	auto&& __r = ((QSwipeGesture*)self)->swipeAngle();
	return __r;
}
void QSwipeGesture_setSwipeAngle(void* self, double value) {
	((QSwipeGesture*)self)->setSwipeAngle(value);
}
void* QSwipeGesture_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSwipeGesture::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSwipeGesture_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSwipeGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSwipeGesture_sender(void* self) {
	auto&& __r = ((QSwipeGesture__Go*)self)->sender();
	return (void*)(__r);
}
int QSwipeGesture_senderSignalIndex(void* self) {
	auto&& __r = ((QSwipeGesture__Go*)self)->senderSignalIndex();
	return __r;
}
int QSwipeGesture_receivers(void* self, const char* signal) {
	auto&& __r = ((QSwipeGesture__Go*)self)->receivers(signal);
	return __r;
}
int QSwipeGesture_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSwipeGesture__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSwipeGesture_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSwipeGesture::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSwipeGesture_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSwipeGesture::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSwipeGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSwipeGesture::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSwipeGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSwipeGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSwipeGesture_metaObject_default(void* self) {
	auto&& __r = ((QSwipeGesture*)self)->QSwipeGesture::metaObject();
	return (void*)(__r);
}
int QSwipeGesture_event_default(void* self, void* event) {
	auto&& __r = ((QSwipeGesture*)self)->QSwipeGesture::event((QEvent*)event);
	return __r;
}
int QSwipeGesture_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSwipeGesture*)self)->QSwipeGesture::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSwipeGesture_timerEvent_default(void* self, void* event) {
	((QSwipeGesture__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSwipeGesture_childEvent_default(void* self, void* event) {
	((QSwipeGesture__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSwipeGesture_customEvent_default(void* self, void* event) {
	((QSwipeGesture__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSwipeGesture_connectNotify_default(void* self, void* signal) {
	((QSwipeGesture__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSwipeGesture_disconnectNotify_default(void* self, void* signal) {
	((QSwipeGesture__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSwipeGesture_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSwipeGesture*)self, &QSwipeGesture::destroyed, (QSwipeGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSwipeGesture);
	__conns_QSwipeGesture[__id] = __c;
}
void QSwipeGesture_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSwipeGesture*)self, &QSwipeGesture::destroyed, (QSwipeGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSwipeGesture);
	__conns_QSwipeGesture[__id] = __c;
}
void QSwipeGesture_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSwipeGesture*)self, &QSwipeGesture::objectNameChanged, (QSwipeGesture*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSwipeGesture);
	__conns_QSwipeGesture[__id] = __c;
}
void QSwipeGesture__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSwipeGesture);
		auto __it = __conns_QSwipeGesture.find(id);
		if (__it == __conns_QSwipeGesture.end()) return;
		__c = __it->second;
		__conns_QSwipeGesture.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSwipeGesture__Go_new(long long id, void* parent) { return new QSwipeGesture__Go(id, (QObject*)parent); }
void QSwipeGesture_delete(void* self) { delete (QSwipeGesture*)self; }

}

#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpainter.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsdropshadoweffect.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsDropShadowEffect__Go : public QGraphicsDropShadowEffect {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QGraphicsDropShadowEffect__Go(long long id, A&&... a) : QGraphicsDropShadowEffect(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsDropShadowEffect::updateBoundingRect;
	using QGraphicsDropShadowEffect::sourceIsPixmap;
	using QGraphicsDropShadowEffect::sourceBoundingRect;
	using QGraphicsDropShadowEffect::drawSource;
	using QGraphicsDropShadowEffect::sourcePixmap;
	using QGraphicsDropShadowEffect::sender;
	using QGraphicsDropShadowEffect::senderSignalIndex;
	using QGraphicsDropShadowEffect::receivers;
	using QGraphicsDropShadowEffect::isSignalConnected;
	void cppgoBase_draw(QPainter* __a0) { QGraphicsDropShadowEffect::draw(__a0); }
	void cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> __a0) { QGraphicsDropShadowEffect::sourceChanged(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsDropShadowEffect::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsDropShadowEffect::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsDropShadowEffect::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsDropShadowEffect::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsDropShadowEffect::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsDropShadowEffect::metaObject();
	}
	QRectF boundingRectFor(const QRectF& rect) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 5831744872130794138LL, (void*)&rect, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsDropShadowEffect::boundingRectFor(rect);
	}
	void draw(QPainter* painter) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1044652152830004333LL, (void*)painter, &__h); if (__h) return; }
		QGraphicsDropShadowEffect::draw(painter);
	}
	void sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> flags) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4520368087989481956LL, (int)(flags), &__h); if (__h) return; }
		QGraphicsDropShadowEffect::sourceChanged(flags);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsDropShadowEffect::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsDropShadowEffect::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsDropShadowEffect::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsDropShadowEffect::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsDropShadowEffect::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsDropShadowEffect::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsDropShadowEffect::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsDropShadowEffect;
static std::mutex __connsMu_QGraphicsDropShadowEffect;

extern "C" {

void* QGraphicsDropShadowEffect_base(void* self) { return (void*)static_cast<QGraphicsEffect*>((QGraphicsDropShadowEffect*)self); }
void QGraphicsDropShadowEffect_chain(void* self, void** out) {
	QGraphicsDropShadowEffect* __p = (QGraphicsDropShadowEffect*)self;
	QGraphicsEffect* __b0 = static_cast<QGraphicsEffect*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsDropShadowEffect__set_override(void* self, int idx) { ((QGraphicsDropShadowEffect__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsDropShadowEffect_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsDropShadowEffect*>((QObject*)o); }
void* QGraphicsDropShadowEffect_new(long long __goid) {
	return new QGraphicsDropShadowEffect__Go(__goid);
}
void* QGraphicsDropShadowEffect_new2(long long __goid, void* parent) {
	return new QGraphicsDropShadowEffect__Go(__goid, (QObject*)parent);
}
char* QGraphicsDropShadowEffect_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsDropShadowEffect::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsDropShadowEffect_offset(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->offset();
	return new QPointF(__r);
}
double QGraphicsDropShadowEffect_xOffset(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->xOffset();
	return __r;
}
double QGraphicsDropShadowEffect_yOffset(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->yOffset();
	return __r;
}
double QGraphicsDropShadowEffect_blurRadius(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->blurRadius();
	return __r;
}
void* QGraphicsDropShadowEffect_color(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->color();
	return new QColor(__r);
}
void QGraphicsDropShadowEffect_setOffset(void* self, void* ofs) {
	((QGraphicsDropShadowEffect*)self)->setOffset(*(QPointF*)ofs);
}
void QGraphicsDropShadowEffect_setOffset2(void* self, double dx, double dy) {
	((QGraphicsDropShadowEffect*)self)->setOffset(dx, dy);
}
void QGraphicsDropShadowEffect_setOffset3(void* self, double d) {
	((QGraphicsDropShadowEffect*)self)->setOffset(d);
}
void QGraphicsDropShadowEffect_setXOffset(void* self, double dx) {
	((QGraphicsDropShadowEffect*)self)->setXOffset(dx);
}
void QGraphicsDropShadowEffect_setYOffset(void* self, double dy) {
	((QGraphicsDropShadowEffect*)self)->setYOffset(dy);
}
void QGraphicsDropShadowEffect_setBlurRadius(void* self, double blurRadius) {
	((QGraphicsDropShadowEffect*)self)->setBlurRadius(blurRadius);
}
void QGraphicsDropShadowEffect_setColor(void* self, void* color) {
	((QGraphicsDropShadowEffect*)self)->setColor(*(QColor*)color);
}
void QGraphicsDropShadowEffect_offsetChanged(void* self, void* offset) {
	((QGraphicsDropShadowEffect*)self)->offsetChanged(*(QPointF*)offset);
}
void QGraphicsDropShadowEffect_blurRadiusChanged(void* self, double blurRadius) {
	((QGraphicsDropShadowEffect*)self)->blurRadiusChanged(blurRadius);
}
void QGraphicsDropShadowEffect_colorChanged(void* self, void* color) {
	((QGraphicsDropShadowEffect*)self)->colorChanged(*(QColor*)color);
}
void QGraphicsDropShadowEffect_updateBoundingRect(void* self) {
	((QGraphicsDropShadowEffect__Go*)self)->updateBoundingRect();
}
int QGraphicsDropShadowEffect_sourceIsPixmap(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourceIsPixmap();
	return __r;
}
void* QGraphicsDropShadowEffect_sourceBoundingRect(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourceBoundingRect();
	return new QRectF(__r);
}
void QGraphicsDropShadowEffect_drawSource(void* self, void* painter) {
	((QGraphicsDropShadowEffect__Go*)self)->drawSource((QPainter*)painter);
}
void* QGraphicsDropShadowEffect_sourcePixmap(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourcePixmap();
	return new QPixmap(__r);
}
void* QGraphicsDropShadowEffect_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsDropShadowEffect::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsDropShadowEffect_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsDropShadowEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsDropShadowEffect_sender(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsDropShadowEffect_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsDropShadowEffect_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsDropShadowEffect_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsDropShadowEffect_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsDropShadowEffect::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsDropShadowEffect_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsDropShadowEffect::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsDropShadowEffect_sourceBoundingRect2(void* self, int system) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourceBoundingRect((Qt::CoordinateSystem)(system));
	return new QRectF(__r);
}
void* QGraphicsDropShadowEffect_sourcePixmap2(void* self, int system) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system));
	return new QPixmap(__r);
}
void* QGraphicsDropShadowEffect_sourcePixmap3(void* self, int system, void* offset) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset);
	return new QPixmap(__r);
}
void* QGraphicsDropShadowEffect_sourcePixmap4(void* self, int system, void* offset, int mode) {
	auto&& __r = ((QGraphicsDropShadowEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset, (QGraphicsEffect::PixmapPadMode)(mode));
	return new QPixmap(__r);
}
void* QGraphicsDropShadowEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsDropShadowEffect::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsDropShadowEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsDropShadowEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsDropShadowEffect_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->QGraphicsDropShadowEffect::metaObject();
	return (void*)(__r);
}
void* QGraphicsDropShadowEffect_boundingRectFor_default(void* self, void* rect) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->QGraphicsDropShadowEffect::boundingRectFor(*(QRectF*)rect);
	return new QRectF(__r);
}
void QGraphicsDropShadowEffect_draw_default(void* self, void* painter) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_draw((QPainter*)painter);
}
void QGraphicsDropShadowEffect_sourceChanged_default(void* self, int flags) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag>(flags));
}
int QGraphicsDropShadowEffect_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->QGraphicsDropShadowEffect::event((QEvent*)event);
	return __r;
}
int QGraphicsDropShadowEffect_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsDropShadowEffect*)self)->QGraphicsDropShadowEffect::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsDropShadowEffect_timerEvent_default(void* self, void* event) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsDropShadowEffect_childEvent_default(void* self, void* event) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsDropShadowEffect_customEvent_default(void* self, void* event) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsDropShadowEffect_connectNotify_default(void* self, void* signal) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsDropShadowEffect_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsDropShadowEffect__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsDropShadowEffect_connect_offsetChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::offsetChanged, (QGraphicsDropShadowEffect*)self, [__id](const QPointF& offset){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&offset);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect_connect_blurRadiusChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::blurRadiusChanged, (QGraphicsDropShadowEffect*)self, [__id](qreal blurRadius){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(blurRadius)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect_connect_colorChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::colorChanged, (QGraphicsDropShadowEffect*)self, [__id](const QColor& color){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&color);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::enabledChanged, (QGraphicsDropShadowEffect*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::destroyed, (QGraphicsDropShadowEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::destroyed, (QGraphicsDropShadowEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsDropShadowEffect*)self, &QGraphicsDropShadowEffect::objectNameChanged, (QGraphicsDropShadowEffect*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
	__conns_QGraphicsDropShadowEffect[__id] = __c;
}
void QGraphicsDropShadowEffect__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsDropShadowEffect);
		auto __it = __conns_QGraphicsDropShadowEffect.find(id);
		if (__it == __conns_QGraphicsDropShadowEffect.end()) return;
		__c = __it->second;
		__conns_QGraphicsDropShadowEffect.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsDropShadowEffect__Go_new(long long id, void* parent) { return new QGraphicsDropShadowEffect__Go(id, (QObject*)parent); }
void QGraphicsDropShadowEffect_delete(void* self) { delete (QGraphicsDropShadowEffect*)self; }

}

#include <QtWidgets/qlineedit.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qvalidator.h>
#include <QtWidgets/qcompleter.h>
#include <QtCore/qsize.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmargins.h>
#include <QtGui/qaction.h>
#include <QtGui/qicon.h>
#include <QtWidgets/qmenu.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qvariant.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qstyle.h>
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
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qlineedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QLineEdit__Go : public QLineEdit {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QLineEdit__Go(long long id, A&&... a) : QLineEdit(std::forward<A>(a)...), __goId(id) {}
	using QLineEdit::cursorRect;
	using QLineEdit::updateMicroFocus;
	using QLineEdit::create;
	using QLineEdit::destroy;
	using QLineEdit::focusNextChild;
	using QLineEdit::focusPreviousChild;
	using QLineEdit::sender;
	using QLineEdit::senderSignalIndex;
	using QLineEdit::receivers;
	using QLineEdit::isSignalConnected;
	using QLineEdit::getDecodedMetricF;
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QLineEdit::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QLineEdit::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QLineEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QLineEdit::mouseDoubleClickEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QLineEdit::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QLineEdit::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QLineEdit::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QLineEdit::focusOutEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QLineEdit::paintEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QLineEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QLineEdit::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QLineEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QLineEdit::dropEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QLineEdit::changeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QLineEdit::contextMenuEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QLineEdit::inputMethodEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QLineEdit::initStyleOption(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QLineEdit::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QLineEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QLineEdit::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QLineEdit::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QLineEdit::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QLineEdit::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QLineEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QLineEdit::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QLineEdit::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QLineEdit::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QLineEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QLineEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QLineEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QLineEdit::sharedPainter(); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QLineEdit::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QLineEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QLineEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QLineEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QLineEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QLineEdit::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QLineEdit::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QLineEdit::minimumSizeHint();
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::mousePressEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::mouseMoveEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::mouseDoubleClickEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::keyPressEvent(arg1);
	}
	void keyReleaseEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::keyReleaseEvent(arg1);
	}
	void focusInEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::focusInEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::focusOutEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::paintEvent(arg1);
	}
	void dragEnterEvent(QDragEnterEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::dragEnterEvent(arg1);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QLineEdit::dragMoveEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QLineEdit::dragLeaveEvent(e);
	}
	void dropEvent(QDropEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::dropEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::changeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::contextMenuEvent(arg1);
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::inputMethodEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QLineEdit::initStyleOption(option);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QLineEdit::inputMethodQuery(arg1);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QLineEdit::timerEvent(arg1);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QLineEdit::event(arg1);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QLineEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QLineEdit::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QLineEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QLineEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QLineEdit::paintEngine();
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::wheelEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QLineEdit::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QLineEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QLineEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QLineEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QLineEdit::sharedPainter();
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QLineEdit::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QLineEdit::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QLineEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QLineEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QLineEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QLineEdit;
static std::mutex __connsMu_QLineEdit;

extern "C" {

void* QLineEdit_base(void* self) { return (void*)static_cast<QWidget*>((QLineEdit*)self); }
void QLineEdit_chain(void* self, void** out) {
	QLineEdit* __p = (QLineEdit*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QLineEdit__set_override(void* self, int idx) { ((QLineEdit__Go*)self)->__goOn[idx] = 1; }
void* QLineEdit_fromQObject(void* o) { return (void*)qobject_cast<QLineEdit*>((QObject*)o); }
void* QLineEdit_new(long long __goid, void* parent) {
	return new QLineEdit__Go(__goid, (QWidget*)parent);
}
void* QLineEdit_new2(long long __goid) {
	return new QLineEdit__Go(__goid);
}
void* QLineEdit_new3(long long __goid, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	return new QLineEdit__Go(__goid, arg1_qstr);
}
void* QLineEdit_new4(long long __goid, const char* arg1, long long arg1_len, void* parent) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	return new QLineEdit__Go(__goid, arg1_qstr, (QWidget*)parent);
}
char* QLineEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QLineEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLineEdit_text(void* self, long long* __out_len) {
	auto&& __r = ((QLineEdit*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLineEdit_displayText(void* self, long long* __out_len) {
	auto&& __r = ((QLineEdit*)self)->displayText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLineEdit_placeholderText(void* self, long long* __out_len) {
	auto&& __r = ((QLineEdit*)self)->placeholderText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QLineEdit_setPlaceholderText(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QLineEdit*)self)->setPlaceholderText(arg1_qstr);
}
int QLineEdit_maxLength(void* self) {
	auto&& __r = ((QLineEdit*)self)->maxLength();
	return __r;
}
void QLineEdit_setMaxLength(void* self, int arg1) {
	((QLineEdit*)self)->setMaxLength(arg1);
}
void QLineEdit_setFrame(void* self, int arg1) {
	((QLineEdit*)self)->setFrame(arg1);
}
int QLineEdit_hasFrame(void* self) {
	auto&& __r = ((QLineEdit*)self)->hasFrame();
	return __r;
}
void QLineEdit_setClearButtonEnabled(void* self, int enable) {
	((QLineEdit*)self)->setClearButtonEnabled(enable);
}
int QLineEdit_isClearButtonEnabled(void* self) {
	auto&& __r = ((QLineEdit*)self)->isClearButtonEnabled();
	return __r;
}
int QLineEdit_echoMode(void* self) {
	auto&& __r = ((QLineEdit*)self)->echoMode();
	return (int)(__r);
}
void QLineEdit_setEchoMode(void* self, int arg1) {
	((QLineEdit*)self)->setEchoMode((QLineEdit::EchoMode)(arg1));
}
int QLineEdit_isReadOnly(void* self) {
	auto&& __r = ((QLineEdit*)self)->isReadOnly();
	return __r;
}
void QLineEdit_setReadOnly(void* self, int arg1) {
	((QLineEdit*)self)->setReadOnly(arg1);
}
void QLineEdit_setValidator(void* self, void* arg1) {
	((QLineEdit*)self)->setValidator((QValidator*)arg1);
}
void* QLineEdit_validator(void* self) {
	auto&& __r = ((QLineEdit*)self)->validator();
	return (void*)(__r);
}
void QLineEdit_setCompleter(void* self, void* completer) {
	((QLineEdit*)self)->setCompleter((QCompleter*)completer);
}
void* QLineEdit_completer(void* self) {
	auto&& __r = ((QLineEdit*)self)->completer();
	return (void*)(__r);
}
int QLineEdit_cursorPosition(void* self) {
	auto&& __r = ((QLineEdit*)self)->cursorPosition();
	return __r;
}
void QLineEdit_setCursorPosition(void* self, int arg1) {
	((QLineEdit*)self)->setCursorPosition(arg1);
}
int QLineEdit_cursorPositionAt(void* self, void* pos) {
	auto&& __r = ((QLineEdit*)self)->cursorPositionAt(*(QPoint*)pos);
	return __r;
}
void QLineEdit_setAlignment(void* self, int flag) {
	((QLineEdit*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(flag));
}
int QLineEdit_alignment(void* self) {
	auto&& __r = ((QLineEdit*)self)->alignment();
	return (int)(__r);
}
void QLineEdit_cursorForward(void* self, int mark) {
	((QLineEdit*)self)->cursorForward(mark);
}
void QLineEdit_cursorBackward(void* self, int mark) {
	((QLineEdit*)self)->cursorBackward(mark);
}
void QLineEdit_cursorWordForward(void* self, int mark) {
	((QLineEdit*)self)->cursorWordForward(mark);
}
void QLineEdit_cursorWordBackward(void* self, int mark) {
	((QLineEdit*)self)->cursorWordBackward(mark);
}
void QLineEdit_backspace(void* self) {
	((QLineEdit*)self)->backspace();
}
void QLineEdit_del(void* self) {
	((QLineEdit*)self)->del();
}
void QLineEdit_home(void* self, int mark) {
	((QLineEdit*)self)->home(mark);
}
void QLineEdit_end(void* self, int mark) {
	((QLineEdit*)self)->end(mark);
}
int QLineEdit_isModified(void* self) {
	auto&& __r = ((QLineEdit*)self)->isModified();
	return __r;
}
void QLineEdit_setModified(void* self, int arg1) {
	((QLineEdit*)self)->setModified(arg1);
}
void QLineEdit_setSelection(void* self, int arg1, int arg2) {
	((QLineEdit*)self)->setSelection(arg1, arg2);
}
int QLineEdit_hasSelectedText(void* self) {
	auto&& __r = ((QLineEdit*)self)->hasSelectedText();
	return __r;
}
char* QLineEdit_selectedText(void* self, long long* __out_len) {
	auto&& __r = ((QLineEdit*)self)->selectedText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLineEdit_selectionStart(void* self) {
	auto&& __r = ((QLineEdit*)self)->selectionStart();
	return __r;
}
int QLineEdit_selectionEnd(void* self) {
	auto&& __r = ((QLineEdit*)self)->selectionEnd();
	return __r;
}
int QLineEdit_selectionLength(void* self) {
	auto&& __r = ((QLineEdit*)self)->selectionLength();
	return __r;
}
int QLineEdit_isUndoAvailable(void* self) {
	auto&& __r = ((QLineEdit*)self)->isUndoAvailable();
	return __r;
}
int QLineEdit_isRedoAvailable(void* self) {
	auto&& __r = ((QLineEdit*)self)->isRedoAvailable();
	return __r;
}
void QLineEdit_setDragEnabled(void* self, int b) {
	((QLineEdit*)self)->setDragEnabled(b);
}
int QLineEdit_dragEnabled(void* self) {
	auto&& __r = ((QLineEdit*)self)->dragEnabled();
	return __r;
}
void QLineEdit_setCursorMoveStyle(void* self, int style) {
	((QLineEdit*)self)->setCursorMoveStyle((Qt::CursorMoveStyle)(style));
}
int QLineEdit_cursorMoveStyle(void* self) {
	auto&& __r = ((QLineEdit*)self)->cursorMoveStyle();
	return (int)(__r);
}
char* QLineEdit_inputMask(void* self, long long* __out_len) {
	auto&& __r = ((QLineEdit*)self)->inputMask();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QLineEdit_setInputMask(void* self, const char* inputMask, long long inputMask_len) {
	QString inputMask_qstr = QString::fromUtf8(inputMask, inputMask_len);
	((QLineEdit*)self)->setInputMask(inputMask_qstr);
}
int QLineEdit_hasAcceptableInput(void* self) {
	auto&& __r = ((QLineEdit*)self)->hasAcceptableInput();
	return __r;
}
void QLineEdit_setTextMargins(void* self, int left, int top, int right, int bottom) {
	((QLineEdit*)self)->setTextMargins(left, top, right, bottom);
}
void QLineEdit_setTextMargins2(void* self, void* margins) {
	((QLineEdit*)self)->setTextMargins(*(QMargins*)margins);
}
void* QLineEdit_textMargins(void* self) {
	auto&& __r = ((QLineEdit*)self)->textMargins();
	return new QMargins(__r);
}
void QLineEdit_addAction(void* self, void* action, int position) {
	((QLineEdit*)self)->addAction((QAction*)action, (QLineEdit::ActionPosition)(position));
}
void* QLineEdit_addAction2(void* self, void* icon, int position) {
	auto&& __r = ((QLineEdit*)self)->addAction(*(QIcon*)icon, (QLineEdit::ActionPosition)(position));
	return (void*)(__r);
}
void QLineEdit_setText(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QLineEdit*)self)->setText(arg1_qstr);
}
void QLineEdit_clear(void* self) {
	((QLineEdit*)self)->clear();
}
void QLineEdit_selectAll(void* self) {
	((QLineEdit*)self)->selectAll();
}
void QLineEdit_undo(void* self) {
	((QLineEdit*)self)->undo();
}
void QLineEdit_redo(void* self) {
	((QLineEdit*)self)->redo();
}
void QLineEdit_cut(void* self) {
	((QLineEdit*)self)->cut();
}
void QLineEdit_copy(void* self) {
	((QLineEdit*)self)->copy();
}
void QLineEdit_paste(void* self) {
	((QLineEdit*)self)->paste();
}
void QLineEdit_deselect(void* self) {
	((QLineEdit*)self)->deselect();
}
void QLineEdit_insert(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QLineEdit*)self)->insert(arg1_qstr);
}
void* QLineEdit_createStandardContextMenu(void* self) {
	auto&& __r = ((QLineEdit*)self)->createStandardContextMenu();
	return (void*)(__r);
}
void QLineEdit_textChanged(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QLineEdit*)self)->textChanged(arg1_qstr);
}
void QLineEdit_textEdited(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QLineEdit*)self)->textEdited(arg1_qstr);
}
void QLineEdit_cursorPositionChanged(void* self, int arg1, int arg2) {
	((QLineEdit*)self)->cursorPositionChanged(arg1, arg2);
}
void QLineEdit_returnPressed(void* self) {
	((QLineEdit*)self)->returnPressed();
}
void QLineEdit_editingFinished(void* self) {
	((QLineEdit*)self)->editingFinished();
}
void QLineEdit_selectionChanged(void* self) {
	((QLineEdit*)self)->selectionChanged();
}
void QLineEdit_inputRejected(void* self) {
	((QLineEdit*)self)->inputRejected();
}
void* QLineEdit_inputMethodQuery(void* self, int arg1) {
	auto&& __r = ((QLineEdit*)self)->inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void* QLineEdit_inputMethodQuery2(void* self, int property, void* argument) {
	auto&& __r = ((QLineEdit*)self)->inputMethodQuery((Qt::InputMethodQuery)(property), *(QVariant*)argument);
	return new QVariant(__r);
}
void* QLineEdit_cursorRect(void* self) {
	auto&& __r = ((QLineEdit__Go*)self)->cursorRect();
	return new QRect(__r);
}
void QLineEdit_setTabOrder(void* arg1, void* arg2) {
	QLineEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QLineEdit_mouseGrabber() {
	auto&& __r = QLineEdit::mouseGrabber();
	return (void*)(__r);
}
void* QLineEdit_keyboardGrabber() {
	auto&& __r = QLineEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QLineEdit_find(unsigned long long arg1) {
	auto&& __r = QLineEdit::find(arg1);
	return (void*)(__r);
}
void* QLineEdit_createWindowContainer(void* window) {
	auto&& __r = QLineEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QLineEdit_updateMicroFocus(void* self) {
	((QLineEdit__Go*)self)->updateMicroFocus();
}
void QLineEdit_create(void* self) {
	((QLineEdit__Go*)self)->create();
}
void QLineEdit_destroy(void* self) {
	((QLineEdit__Go*)self)->destroy();
}
int QLineEdit_focusNextChild(void* self) {
	auto&& __r = ((QLineEdit__Go*)self)->focusNextChild();
	return __r;
}
int QLineEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QLineEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QLineEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QLineEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QLineEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QLineEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QLineEdit_sender(void* self) {
	auto&& __r = ((QLineEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QLineEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QLineEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QLineEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QLineEdit__Go*)self)->receivers(signal);
	return __r;
}
int QLineEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QLineEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QLineEdit_devicePixelRatioFScale() {
	auto&& __r = QLineEdit::devicePixelRatioFScale();
	return __r;
}
int QLineEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QLineEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QLineEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QLineEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QLineEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QLineEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLineEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QLineEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QLineEdit_cursorForward2(void* self, int mark, int steps) {
	((QLineEdit*)self)->cursorForward(mark, steps);
}
void QLineEdit_cursorBackward2(void* self, int mark, int steps) {
	((QLineEdit*)self)->cursorBackward(mark, steps);
}
void* QLineEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QLineEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QLineEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QLineEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QLineEdit_updateMicroFocus2(void* self, int query) {
	((QLineEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QLineEdit_create2(void* self, unsigned long long arg1) {
	((QLineEdit__Go*)self)->create(arg1);
}
void QLineEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QLineEdit__Go*)self)->create(arg1, initializeWindow);
}
void QLineEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QLineEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QLineEdit_destroy2(void* self, int destroyWindow) {
	((QLineEdit__Go*)self)->destroy(destroyWindow);
}
void QLineEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QLineEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QLineEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QLineEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QLineEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QLineEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QLineEdit_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void* QLineEdit_inputMethodQuery_default2(void* self, int property, void* argument) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::inputMethodQuery((Qt::InputMethodQuery)(property), *(QVariant*)argument);
	return new QVariant(__r);
}
void* QLineEdit_metaObject_default(void* self) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::metaObject();
	return (void*)(__r);
}
void* QLineEdit_sizeHint_default(void* self) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::sizeHint();
	return new QSize(__r);
}
void* QLineEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::minimumSizeHint();
	return new QSize(__r);
}
void QLineEdit_mousePressEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QLineEdit_mouseMoveEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QLineEdit_mouseReleaseEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QLineEdit_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QLineEdit_keyPressEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QLineEdit_keyReleaseEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)arg1);
}
void QLineEdit_focusInEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)arg1);
}
void QLineEdit_focusOutEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
void QLineEdit_paintEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QLineEdit_dragEnterEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)arg1);
}
void QLineEdit_dragMoveEvent_default(void* self, void* e) {
	((QLineEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QLineEdit_dragLeaveEvent_default(void* self, void* e) {
	((QLineEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QLineEdit_dropEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)arg1);
}
void QLineEdit_changeEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QLineEdit_contextMenuEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QLineEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QLineEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void QLineEdit_initStyleOption_default(void* self, void* option) {
	((QLineEdit__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
void QLineEdit_timerEvent_default(void* self, void* arg1) {
	((QLineEdit*)self)->QLineEdit::timerEvent((QTimerEvent*)arg1);
}
int QLineEdit_event_default(void* self, void* arg1) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::event((QEvent*)arg1);
	return __r;
}
int QLineEdit_devType_default(void* self) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::devType();
	return __r;
}
void QLineEdit_setVisible_default(void* self, int visible) {
	((QLineEdit*)self)->QLineEdit::setVisible(visible);
}
int QLineEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::heightForWidth(arg1);
	return __r;
}
int QLineEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::hasHeightForWidth();
	return __r;
}
void* QLineEdit_paintEngine_default(void* self) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::paintEngine();
	return (void*)(__r);
}
void QLineEdit_wheelEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QLineEdit_enterEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QLineEdit_leaveEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QLineEdit_moveEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QLineEdit_resizeEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QLineEdit_closeEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QLineEdit_tabletEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QLineEdit_actionEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QLineEdit_showEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QLineEdit_hideEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QLineEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QLineEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QLineEdit_initPainter_default(void* self, void* painter) {
	((QLineEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QLineEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QLineEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QLineEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QLineEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
int QLineEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QLineEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QLineEdit_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QLineEdit*)self)->QLineEdit::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QLineEdit_childEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QLineEdit_customEvent_default(void* self, void* event) {
	((QLineEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QLineEdit_connectNotify_default(void* self, void* signal) {
	((QLineEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QLineEdit_disconnectNotify_default(void* self, void* signal) {
	((QLineEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QLineEdit_connect_textChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::textChanged, (QLineEdit*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_textEdited(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::textEdited, (QLineEdit*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_cursorPositionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::cursorPositionChanged, (QLineEdit*)self, [__id](int arg1, int arg2){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, arg1, arg2);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::returnPressed, (QLineEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::editingFinished, (QLineEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::selectionChanged, (QLineEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_inputRejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::inputRejected, (QLineEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::windowTitleChanged, (QLineEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::windowIconChanged, (QLineEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::windowIconTextChanged, (QLineEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::customContextMenuRequested, (QLineEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::destroyed, (QLineEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::destroyed, (QLineEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLineEdit*)self, &QLineEdit::objectNameChanged, (QLineEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
	__conns_QLineEdit[__id] = __c;
}
void QLineEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QLineEdit);
		auto __it = __conns_QLineEdit.find(id);
		if (__it == __conns_QLineEdit.end()) return;
		__c = __it->second;
		__conns_QLineEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QLineEdit__Go_new(long long id) { return new QLineEdit__Go(id); }
void QLineEdit_delete(void* self) { delete (QLineEdit*)self; }

}

#include <QtWidgets/qradiobutton.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
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
#include "gen_qradiobutton.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QRadioButton__Go : public QRadioButton {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QRadioButton__Go(long long id, A&&... a) : QRadioButton(std::forward<A>(a)...), __goId(id) {}
	using QRadioButton::updateMicroFocus;
	using QRadioButton::create;
	using QRadioButton::destroy;
	using QRadioButton::focusNextChild;
	using QRadioButton::focusPreviousChild;
	using QRadioButton::sender;
	using QRadioButton::senderSignalIndex;
	using QRadioButton::receivers;
	using QRadioButton::isSignalConnected;
	using QRadioButton::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QRadioButton::event(__a0); }
	bool cppgoBase_hitButton(const QPoint& __a0) { return QRadioButton::hitButton(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QRadioButton::paintEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QRadioButton::mouseMoveEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionButton* __a0) { QRadioButton::initStyleOption(__a0); }
	void cppgoBase_checkStateSet() { QRadioButton::checkStateSet(); }
	void cppgoBase_nextCheckState() { QRadioButton::nextCheckState(); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QRadioButton::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QRadioButton::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QRadioButton::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QRadioButton::mouseReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QRadioButton::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QRadioButton::focusOutEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QRadioButton::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QRadioButton::timerEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QRadioButton::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QRadioButton::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QRadioButton::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QRadioButton::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QRadioButton::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QRadioButton::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QRadioButton::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QRadioButton::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QRadioButton::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QRadioButton::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QRadioButton::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QRadioButton::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QRadioButton::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QRadioButton::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QRadioButton::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QRadioButton::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QRadioButton::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QRadioButton::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QRadioButton::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QRadioButton::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QRadioButton::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QRadioButton::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QRadioButton::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QRadioButton::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QRadioButton::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QRadioButton::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QRadioButton::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QRadioButton::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QRadioButton::minimumSizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QRadioButton::event(e);
	}
	bool hitButton(const QPoint& arg1) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5736040353629178930LL, (void*)&arg1, &__h); if (__h) return __r; }
		return QRadioButton::hitButton(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QRadioButton::paintEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QRadioButton::mouseMoveEvent(arg1);
	}
	void initStyleOption(QStyleOptionButton* button) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)button, &__h); if (__h) return; }
		QRadioButton::initStyleOption(button);
	}
	void checkStateSet() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5747806663224476200LL, &__h); if (__h) return; }
		QRadioButton::checkStateSet();
	}
	void nextCheckState() override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1474647914901213797LL, &__h); if (__h) return; }
		QRadioButton::nextCheckState();
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::keyReleaseEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::mouseReleaseEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::focusOutEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::changeEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QRadioButton::timerEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QRadioButton::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QRadioButton::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QRadioButton::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QRadioButton::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QRadioButton::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::wheelEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QRadioButton::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QRadioButton::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QRadioButton::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QRadioButton::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QRadioButton::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QRadioButton::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QRadioButton::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QRadioButton::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QRadioButton::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QRadioButton::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QRadioButton::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QRadioButton::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QRadioButton;
static std::mutex __connsMu_QRadioButton;

extern "C" {

void* QRadioButton_base(void* self) { return (void*)static_cast<QAbstractButton*>((QRadioButton*)self); }
void QRadioButton_chain(void* self, void** out) {
	QRadioButton* __p = (QRadioButton*)self;
	QAbstractButton* __b0 = static_cast<QAbstractButton*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QRadioButton__set_override(void* self, int idx) { ((QRadioButton__Go*)self)->__goOn[idx] = 1; }
void* QRadioButton_fromQObject(void* o) { return (void*)qobject_cast<QRadioButton*>((QObject*)o); }
void* QRadioButton_new(long long __goid, void* parent) {
	return new QRadioButton__Go(__goid, (QWidget*)parent);
}
void* QRadioButton_new2(long long __goid) {
	return new QRadioButton__Go(__goid);
}
void* QRadioButton_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QRadioButton__Go(__goid, text_qstr);
}
void* QRadioButton_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QRadioButton__Go(__goid, text_qstr, (QWidget*)parent);
}
char* QRadioButton_tr(const char* s, long long* __out_len) {
	auto&& __r = QRadioButton::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QRadioButton_setTabOrder(void* arg1, void* arg2) {
	QRadioButton::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QRadioButton_mouseGrabber() {
	auto&& __r = QRadioButton::mouseGrabber();
	return (void*)(__r);
}
void* QRadioButton_keyboardGrabber() {
	auto&& __r = QRadioButton::keyboardGrabber();
	return (void*)(__r);
}
void* QRadioButton_find(unsigned long long arg1) {
	auto&& __r = QRadioButton::find(arg1);
	return (void*)(__r);
}
void* QRadioButton_createWindowContainer(void* window) {
	auto&& __r = QRadioButton::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QRadioButton_updateMicroFocus(void* self) {
	((QRadioButton__Go*)self)->updateMicroFocus();
}
void QRadioButton_create(void* self) {
	((QRadioButton__Go*)self)->create();
}
void QRadioButton_destroy(void* self) {
	((QRadioButton__Go*)self)->destroy();
}
int QRadioButton_focusNextChild(void* self) {
	auto&& __r = ((QRadioButton__Go*)self)->focusNextChild();
	return __r;
}
int QRadioButton_focusPreviousChild(void* self) {
	auto&& __r = ((QRadioButton__Go*)self)->focusPreviousChild();
	return __r;
}
void* QRadioButton_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QRadioButton::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QRadioButton_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QRadioButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QRadioButton_sender(void* self) {
	auto&& __r = ((QRadioButton__Go*)self)->sender();
	return (void*)(__r);
}
int QRadioButton_senderSignalIndex(void* self) {
	auto&& __r = ((QRadioButton__Go*)self)->senderSignalIndex();
	return __r;
}
int QRadioButton_receivers(void* self, const char* signal) {
	auto&& __r = ((QRadioButton__Go*)self)->receivers(signal);
	return __r;
}
int QRadioButton_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QRadioButton__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QRadioButton_devicePixelRatioFScale() {
	auto&& __r = QRadioButton::devicePixelRatioFScale();
	return __r;
}
int QRadioButton_encodeMetricF(int metric, double value) {
	auto&& __r = QRadioButton::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QRadioButton_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QRadioButton__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QRadioButton_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QRadioButton::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRadioButton_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QRadioButton::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRadioButton_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QRadioButton::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QRadioButton_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QRadioButton::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QRadioButton_updateMicroFocus2(void* self, int query) {
	((QRadioButton__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QRadioButton_create2(void* self, unsigned long long arg1) {
	((QRadioButton__Go*)self)->create(arg1);
}
void QRadioButton_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QRadioButton__Go*)self)->create(arg1, initializeWindow);
}
void QRadioButton_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QRadioButton__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QRadioButton_destroy2(void* self, int destroyWindow) {
	((QRadioButton__Go*)self)->destroy(destroyWindow);
}
void QRadioButton_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QRadioButton__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QRadioButton_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QRadioButton::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QRadioButton_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QRadioButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QRadioButton_metaObject_default(void* self) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::metaObject();
	return (void*)(__r);
}
void* QRadioButton_sizeHint_default(void* self) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::sizeHint();
	return new QSize(__r);
}
void* QRadioButton_minimumSizeHint_default(void* self) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::minimumSizeHint();
	return new QSize(__r);
}
int QRadioButton_event_default(void* self, void* e) {
	auto&& __r = ((QRadioButton__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QRadioButton_hitButton_default(void* self, void* arg1) {
	auto&& __r = ((QRadioButton__Go*)self)->cppgoBase_hitButton(*(QPoint*)arg1);
	return __r;
}
void QRadioButton_paintEvent_default(void* self, void* arg1) {
	((QRadioButton__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QRadioButton_mouseMoveEvent_default(void* self, void* arg1) {
	((QRadioButton__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QRadioButton_initStyleOption_default(void* self, void* button) {
	((QRadioButton__Go*)self)->cppgoBase_initStyleOption((QStyleOptionButton*)button);
}
void QRadioButton_checkStateSet_default(void* self) {
	((QRadioButton__Go*)self)->cppgoBase_checkStateSet();
}
void QRadioButton_nextCheckState_default(void* self) {
	((QRadioButton__Go*)self)->cppgoBase_nextCheckState();
}
void QRadioButton_keyPressEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QRadioButton_keyReleaseEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QRadioButton_mousePressEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QRadioButton_mouseReleaseEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QRadioButton_focusInEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QRadioButton_focusOutEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QRadioButton_changeEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QRadioButton_timerEvent_default(void* self, void* e) {
	((QRadioButton__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
int QRadioButton_devType_default(void* self) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::devType();
	return __r;
}
void QRadioButton_setVisible_default(void* self, int visible) {
	((QRadioButton*)self)->QRadioButton::setVisible(visible);
}
int QRadioButton_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::heightForWidth(arg1);
	return __r;
}
int QRadioButton_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::hasHeightForWidth();
	return __r;
}
void* QRadioButton_paintEngine_default(void* self) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::paintEngine();
	return (void*)(__r);
}
void QRadioButton_mouseDoubleClickEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QRadioButton_wheelEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QRadioButton_enterEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QRadioButton_leaveEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QRadioButton_moveEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QRadioButton_resizeEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QRadioButton_closeEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QRadioButton_contextMenuEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QRadioButton_tabletEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QRadioButton_actionEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QRadioButton_dragEnterEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QRadioButton_dragMoveEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QRadioButton_dragLeaveEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QRadioButton_dropEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QRadioButton_showEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QRadioButton_hideEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QRadioButton_metric_default(void* self, int arg1) {
	auto&& __r = ((QRadioButton__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QRadioButton_initPainter_default(void* self, void* painter) {
	((QRadioButton__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QRadioButton_redirected_default(void* self, void* offset) {
	auto&& __r = ((QRadioButton__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QRadioButton_sharedPainter_default(void* self) {
	auto&& __r = ((QRadioButton__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QRadioButton_inputMethodEvent_default(void* self, void* arg1) {
	((QRadioButton__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QRadioButton_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QRadioButton_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QRadioButton__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QRadioButton_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QRadioButton*)self)->QRadioButton::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QRadioButton_childEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QRadioButton_customEvent_default(void* self, void* event) {
	((QRadioButton__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QRadioButton_connectNotify_default(void* self, void* signal) {
	((QRadioButton__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QRadioButton_disconnectNotify_default(void* self, void* signal) {
	((QRadioButton__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QRadioButton_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::pressed, (QRadioButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_released(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::released, (QRadioButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::clicked, (QRadioButton*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::clicked, (QRadioButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::toggled, (QRadioButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::windowTitleChanged, (QRadioButton*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::windowIconChanged, (QRadioButton*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::windowIconTextChanged, (QRadioButton*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::customContextMenuRequested, (QRadioButton*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::destroyed, (QRadioButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::destroyed, (QRadioButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRadioButton*)self, &QRadioButton::objectNameChanged, (QRadioButton*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
	__conns_QRadioButton[__id] = __c;
}
void QRadioButton__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QRadioButton);
		auto __it = __conns_QRadioButton.find(id);
		if (__it == __conns_QRadioButton.end()) return;
		__c = __it->second;
		__conns_QRadioButton.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QRadioButton__Go_new(long long id) { return new QRadioButton__Go(id); }
void QRadioButton_delete(void* self) { delete (QRadioButton*)self; }

}

#include <QtWidgets/qtoolbar.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qaction.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
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
#include "gen_qtoolbar.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QToolBar__Go : public QToolBar {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QToolBar__Go(long long id, A&&... a) : QToolBar(std::forward<A>(a)...), __goId(id) {}
	using QToolBar::updateMicroFocus;
	using QToolBar::create;
	using QToolBar::destroy;
	using QToolBar::focusNextChild;
	using QToolBar::focusPreviousChild;
	using QToolBar::sender;
	using QToolBar::senderSignalIndex;
	using QToolBar::receivers;
	using QToolBar::isSignalConnected;
	using QToolBar::getDecodedMetricF;
	void cppgoBase_actionEvent(QActionEvent* __a0) { QToolBar::actionEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QToolBar::changeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QToolBar::paintEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QToolBar::event(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionToolBar* __a0) { QToolBar::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QToolBar::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QToolBar::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QToolBar::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QToolBar::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QToolBar::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QToolBar::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QToolBar::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QToolBar::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QToolBar::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QToolBar::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QToolBar::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QToolBar::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QToolBar::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QToolBar::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QToolBar::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QToolBar::tabletEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QToolBar::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QToolBar::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QToolBar::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QToolBar::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QToolBar::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QToolBar::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QToolBar::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QToolBar::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QToolBar::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QToolBar::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QToolBar::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QToolBar::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QToolBar::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QToolBar::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QToolBar::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QToolBar::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QToolBar::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QToolBar::metaObject();
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QToolBar::actionEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QToolBar::changeEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QToolBar::paintEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QToolBar::event(event);
	}
	void initStyleOption(QStyleOptionToolBar* option) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QToolBar::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QToolBar::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QToolBar::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QToolBar::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QToolBar::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QToolBar::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QToolBar::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QToolBar::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QToolBar::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QToolBar::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QToolBar::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QToolBar::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QToolBar::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QToolBar::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QToolBar::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QToolBar::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QToolBar::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QToolBar::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QToolBar::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QToolBar::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QToolBar::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QToolBar::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QToolBar::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QToolBar::tabletEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QToolBar::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QToolBar::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QToolBar::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QToolBar::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QToolBar::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QToolBar::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QToolBar::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QToolBar::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QToolBar::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QToolBar::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QToolBar::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QToolBar::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QToolBar::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QToolBar::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QToolBar::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QToolBar::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QToolBar::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QToolBar::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QToolBar::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QToolBar::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QToolBar;
static std::mutex __connsMu_QToolBar;

extern "C" {

void* QToolBar_base(void* self) { return (void*)static_cast<QWidget*>((QToolBar*)self); }
void QToolBar_chain(void* self, void** out) {
	QToolBar* __p = (QToolBar*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QToolBar__set_override(void* self, int idx) { ((QToolBar__Go*)self)->__goOn[idx] = 1; }
void* QToolBar_fromQObject(void* o) { return (void*)qobject_cast<QToolBar*>((QObject*)o); }
void* QToolBar_new(long long __goid, void* parent) {
	return new QToolBar__Go(__goid, (QWidget*)parent);
}
void* QToolBar_new2(long long __goid, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QToolBar__Go(__goid, title_qstr);
}
void* QToolBar_new3(long long __goid) {
	return new QToolBar__Go(__goid);
}
void* QToolBar_new4(long long __goid, const char* title, long long title_len, void* parent) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QToolBar__Go(__goid, title_qstr, (QWidget*)parent);
}
char* QToolBar_tr(const char* s, long long* __out_len) {
	auto&& __r = QToolBar::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QToolBar_setMovable(void* self, int movable) {
	((QToolBar*)self)->setMovable(movable);
}
int QToolBar_isMovable(void* self) {
	auto&& __r = ((QToolBar*)self)->isMovable();
	return __r;
}
void QToolBar_setAllowedAreas(void* self, int areas) {
	((QToolBar*)self)->setAllowedAreas(QFlags<Qt::ToolBarArea>(areas));
}
int QToolBar_allowedAreas(void* self) {
	auto&& __r = ((QToolBar*)self)->allowedAreas();
	return (int)(__r);
}
int QToolBar_isAreaAllowed(void* self, int area) {
	auto&& __r = ((QToolBar*)self)->isAreaAllowed((Qt::ToolBarArea)(area));
	return __r;
}
void QToolBar_setOrientation(void* self, int orientation) {
	((QToolBar*)self)->setOrientation((Qt::Orientation)(orientation));
}
int QToolBar_orientation(void* self) {
	auto&& __r = ((QToolBar*)self)->orientation();
	return (int)(__r);
}
void QToolBar_clear(void* self) {
	((QToolBar*)self)->clear();
}
void* QToolBar_addSeparator(void* self) {
	auto&& __r = ((QToolBar*)self)->addSeparator();
	return (void*)(__r);
}
void* QToolBar_insertSeparator(void* self, void* before) {
	auto&& __r = ((QToolBar*)self)->insertSeparator((QAction*)before);
	return (void*)(__r);
}
void* QToolBar_addWidget(void* self, void* widget) {
	auto&& __r = ((QToolBar*)self)->addWidget((QWidget*)widget);
	return (void*)(__r);
}
void* QToolBar_insertWidget(void* self, void* before, void* widget) {
	auto&& __r = ((QToolBar*)self)->insertWidget((QAction*)before, (QWidget*)widget);
	return (void*)(__r);
}
void* QToolBar_actionGeometry(void* self, void* action) {
	auto&& __r = ((QToolBar*)self)->actionGeometry((QAction*)action);
	return new QRect(__r);
}
void* QToolBar_actionAt(void* self, void* p) {
	auto&& __r = ((QToolBar*)self)->actionAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QToolBar_actionAt2(void* self, int x, int y) {
	auto&& __r = ((QToolBar*)self)->actionAt(x, y);
	return (void*)(__r);
}
void* QToolBar_toggleViewAction(void* self) {
	auto&& __r = ((QToolBar*)self)->toggleViewAction();
	return (void*)(__r);
}
void* QToolBar_iconSize(void* self) {
	auto&& __r = ((QToolBar*)self)->iconSize();
	return new QSize(__r);
}
int QToolBar_toolButtonStyle(void* self) {
	auto&& __r = ((QToolBar*)self)->toolButtonStyle();
	return (int)(__r);
}
void* QToolBar_widgetForAction(void* self, void* action) {
	auto&& __r = ((QToolBar*)self)->widgetForAction((QAction*)action);
	return (void*)(__r);
}
int QToolBar_isFloatable(void* self) {
	auto&& __r = ((QToolBar*)self)->isFloatable();
	return __r;
}
void QToolBar_setFloatable(void* self, int floatable) {
	((QToolBar*)self)->setFloatable(floatable);
}
int QToolBar_isFloating(void* self) {
	auto&& __r = ((QToolBar*)self)->isFloating();
	return __r;
}
void QToolBar_setIconSize(void* self, void* iconSize) {
	((QToolBar*)self)->setIconSize(*(QSize*)iconSize);
}
void QToolBar_setToolButtonStyle(void* self, int toolButtonStyle) {
	((QToolBar*)self)->setToolButtonStyle((Qt::ToolButtonStyle)(toolButtonStyle));
}
void QToolBar_actionTriggered(void* self, void* action) {
	((QToolBar*)self)->actionTriggered((QAction*)action);
}
void QToolBar_movableChanged(void* self, int movable) {
	((QToolBar*)self)->movableChanged(movable);
}
void QToolBar_allowedAreasChanged(void* self, int allowedAreas) {
	((QToolBar*)self)->allowedAreasChanged(QFlags<Qt::ToolBarArea>(allowedAreas));
}
void QToolBar_orientationChanged(void* self, int orientation) {
	((QToolBar*)self)->orientationChanged((Qt::Orientation)(orientation));
}
void QToolBar_iconSizeChanged(void* self, void* iconSize) {
	((QToolBar*)self)->iconSizeChanged(*(QSize*)iconSize);
}
void QToolBar_toolButtonStyleChanged(void* self, int toolButtonStyle) {
	((QToolBar*)self)->toolButtonStyleChanged((Qt::ToolButtonStyle)(toolButtonStyle));
}
void QToolBar_topLevelChanged(void* self, int topLevel) {
	((QToolBar*)self)->topLevelChanged(topLevel);
}
void QToolBar_visibilityChanged(void* self, int visible) {
	((QToolBar*)self)->visibilityChanged(visible);
}
void QToolBar_setTabOrder(void* arg1, void* arg2) {
	QToolBar::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QToolBar_mouseGrabber() {
	auto&& __r = QToolBar::mouseGrabber();
	return (void*)(__r);
}
void* QToolBar_keyboardGrabber() {
	auto&& __r = QToolBar::keyboardGrabber();
	return (void*)(__r);
}
void* QToolBar_find(unsigned long long arg1) {
	auto&& __r = QToolBar::find(arg1);
	return (void*)(__r);
}
void* QToolBar_createWindowContainer(void* window) {
	auto&& __r = QToolBar::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QToolBar_updateMicroFocus(void* self) {
	((QToolBar__Go*)self)->updateMicroFocus();
}
void QToolBar_create(void* self) {
	((QToolBar__Go*)self)->create();
}
void QToolBar_destroy(void* self) {
	((QToolBar__Go*)self)->destroy();
}
int QToolBar_focusNextChild(void* self) {
	auto&& __r = ((QToolBar__Go*)self)->focusNextChild();
	return __r;
}
int QToolBar_focusPreviousChild(void* self) {
	auto&& __r = ((QToolBar__Go*)self)->focusPreviousChild();
	return __r;
}
void* QToolBar_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QToolBar::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QToolBar_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QToolBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QToolBar_sender(void* self) {
	auto&& __r = ((QToolBar__Go*)self)->sender();
	return (void*)(__r);
}
int QToolBar_senderSignalIndex(void* self) {
	auto&& __r = ((QToolBar__Go*)self)->senderSignalIndex();
	return __r;
}
int QToolBar_receivers(void* self, const char* signal) {
	auto&& __r = ((QToolBar__Go*)self)->receivers(signal);
	return __r;
}
int QToolBar_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QToolBar__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QToolBar_devicePixelRatioFScale() {
	auto&& __r = QToolBar::devicePixelRatioFScale();
	return __r;
}
int QToolBar_encodeMetricF(int metric, double value) {
	auto&& __r = QToolBar::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QToolBar_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QToolBar__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QToolBar_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QToolBar::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QToolBar_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QToolBar::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QToolBar_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QToolBar::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QToolBar_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QToolBar::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QToolBar_updateMicroFocus2(void* self, int query) {
	((QToolBar__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QToolBar_create2(void* self, unsigned long long arg1) {
	((QToolBar__Go*)self)->create(arg1);
}
void QToolBar_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QToolBar__Go*)self)->create(arg1, initializeWindow);
}
void QToolBar_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QToolBar__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QToolBar_destroy2(void* self, int destroyWindow) {
	((QToolBar__Go*)self)->destroy(destroyWindow);
}
void QToolBar_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QToolBar__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QToolBar_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QToolBar::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QToolBar_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QToolBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QToolBar_metaObject_default(void* self) {
	auto&& __r = ((QToolBar*)self)->QToolBar::metaObject();
	return (void*)(__r);
}
void QToolBar_actionEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QToolBar_changeEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QToolBar_paintEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
int QToolBar_event_default(void* self, void* event) {
	auto&& __r = ((QToolBar__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QToolBar_initStyleOption_default(void* self, void* option) {
	((QToolBar__Go*)self)->cppgoBase_initStyleOption((QStyleOptionToolBar*)option);
}
int QToolBar_devType_default(void* self) {
	auto&& __r = ((QToolBar*)self)->QToolBar::devType();
	return __r;
}
void QToolBar_setVisible_default(void* self, int visible) {
	((QToolBar*)self)->QToolBar::setVisible(visible);
}
void* QToolBar_sizeHint_default(void* self) {
	auto&& __r = ((QToolBar*)self)->QToolBar::sizeHint();
	return new QSize(__r);
}
void* QToolBar_minimumSizeHint_default(void* self) {
	auto&& __r = ((QToolBar*)self)->QToolBar::minimumSizeHint();
	return new QSize(__r);
}
int QToolBar_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QToolBar*)self)->QToolBar::heightForWidth(arg1);
	return __r;
}
int QToolBar_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QToolBar*)self)->QToolBar::hasHeightForWidth();
	return __r;
}
void* QToolBar_paintEngine_default(void* self) {
	auto&& __r = ((QToolBar*)self)->QToolBar::paintEngine();
	return (void*)(__r);
}
void QToolBar_mousePressEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QToolBar_mouseReleaseEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QToolBar_mouseDoubleClickEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QToolBar_mouseMoveEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QToolBar_wheelEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QToolBar_keyPressEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QToolBar_keyReleaseEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QToolBar_focusInEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QToolBar_focusOutEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QToolBar_enterEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QToolBar_leaveEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QToolBar_moveEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QToolBar_resizeEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QToolBar_closeEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QToolBar_contextMenuEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QToolBar_tabletEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QToolBar_dragEnterEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QToolBar_dragMoveEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QToolBar_dragLeaveEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QToolBar_dropEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QToolBar_showEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QToolBar_hideEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QToolBar_metric_default(void* self, int arg1) {
	auto&& __r = ((QToolBar__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QToolBar_initPainter_default(void* self, void* painter) {
	((QToolBar__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QToolBar_redirected_default(void* self, void* offset) {
	auto&& __r = ((QToolBar__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QToolBar_sharedPainter_default(void* self) {
	auto&& __r = ((QToolBar__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QToolBar_inputMethodEvent_default(void* self, void* arg1) {
	((QToolBar__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QToolBar_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QToolBar*)self)->QToolBar::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QToolBar_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QToolBar__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QToolBar_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QToolBar*)self)->QToolBar::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QToolBar_timerEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QToolBar_childEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QToolBar_customEvent_default(void* self, void* event) {
	((QToolBar__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QToolBar_connectNotify_default(void* self, void* signal) {
	((QToolBar__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QToolBar_disconnectNotify_default(void* self, void* signal) {
	((QToolBar__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QToolBar_connect_actionTriggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::actionTriggered, (QToolBar*)self, [__id](QAction* action){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_movableChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::movableChanged, (QToolBar*)self, [__id](bool movable){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, movable);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_allowedAreasChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::allowedAreasChanged, (QToolBar*)self, [__id](QFlags<Qt::ToolBarArea> allowedAreas){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(allowedAreas));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_orientationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::orientationChanged, (QToolBar*)self, [__id](Qt::Orientation orientation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(orientation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::iconSizeChanged, (QToolBar*)self, [__id](const QSize& iconSize){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&iconSize);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_toolButtonStyleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::toolButtonStyleChanged, (QToolBar*)self, [__id](Qt::ToolButtonStyle toolButtonStyle){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(toolButtonStyle));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_topLevelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::topLevelChanged, (QToolBar*)self, [__id](bool topLevel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, topLevel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_visibilityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::visibilityChanged, (QToolBar*)self, [__id](bool visible){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, visible);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::windowTitleChanged, (QToolBar*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::windowIconChanged, (QToolBar*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::windowIconTextChanged, (QToolBar*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::customContextMenuRequested, (QToolBar*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::destroyed, (QToolBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::destroyed, (QToolBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QToolBar*)self, &QToolBar::objectNameChanged, (QToolBar*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
	__conns_QToolBar[__id] = __c;
}
void QToolBar__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QToolBar);
		auto __it = __conns_QToolBar.find(id);
		if (__it == __conns_QToolBar.end()) return;
		__c = __it->second;
		__conns_QToolBar.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QToolBar__Go_new(long long id) { return new QToolBar__Go(id); }
void QToolBar_delete(void* self) { delete (QToolBar*)self; }

}

