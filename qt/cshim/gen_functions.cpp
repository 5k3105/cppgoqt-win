#include <QtCore/q20type_traits.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qalgorithms.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qassert.h>
#include <QtCore/qbasictimer.h>
#include <QtCore/qbindingstorage.h>
#include <QtCore/qbitarray.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayalgorithms.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcborarray.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qcbormap.h>
#include <QtCore/qcborvalue.h>
#include <QtCore/qchar.h>
#include <QtCore/qcollator.h>
#include <QtCore/qcommandlineoption.h>
#include <QtCore/qcoreapplication.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qdebug.h>
#include <QtCore/qdir.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qendian.h>
#include <QtCore/qfile.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qfloat16.h>
#include <QtCore/qfutureinterface.h>
#include <QtCore/qhash.h>
#include <QtCore/qhashfunctions.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsondocument.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qlibraryinfo.h>
#include <QtCore/qlocale.h>
#include <QtCore/qlogging.h>
#include <QtCore/qmargins.h>
#include <QtCore/qmath.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qmimetype.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qnumeric.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qpermissions.h>
#include <QtCore/qplugin.h>
#include <QtCore/qpoint.h>
#include <QtCore/qprocess.h>
#include <QtCore/qproperty.h>
#include <QtCore/qrect.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qsize.h>
#include <QtCore/qstorageinfo.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include <QtCore/qtclasshelpermacros.h>
#include <QtCore/qtemporarydir.h>
#include <QtCore/qtenvironmentvariables.h>
#include <QtCore/qtextboundaryfinder.h>
#include <QtCore/qtextstream.h>
#include <QtCore/qtimezone.h>
#include <QtCore/qtipccommon.h>
#include <QtCore/qtmetamacros.h>
#include <QtCore/qttranslation.h>
#include <QtCore/qtversion.h>
#include <QtCore/qtyperevision.h>
#include <QtCore/qurl.h>
#include <QtCore/qurlquery.h>
#include <QtCore/quuid.h>
#include <QtCore/qvariant.h>
#include <QtCore/qversionnumber.h>
#include <QtCore/qyieldcpu.h>
#include <QtGui/qaccessible.h>
#include <QtGui/qaccessible_base.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qbrush.h>
#include <QtGui/qcolor.h>
#include <QtGui/qcolorspace.h>
#include <QtGui/qcolortransform.h>
#include <QtGui/qcursor.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontvariableaxis.h>
#include <QtGui/qglyphrun.h>
#include <QtGui/qicon.h>
#include <QtGui/qimage.h>
#include <QtGui/qkeysequence.h>
#include <QtGui/qpagelayout.h>
#include <QtGui/qpageranges.h>
#include <QtGui/qpagesize.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qpalette.h>
#include <QtGui/qpdfoutputintent.h>
#include <QtGui/qpen.h>
#include <QtGui/qpicture.h>
#include <QtGui/qpixelformat.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qpixmapcache.h>
#include <QtGui/qpointingdevice.h>
#include <QtGui/qpolygon.h>
#include <QtGui/qquaternion.h>
#include <QtGui/qrawfont.h>
#include <QtGui/qregion.h>
#include <QtGui/qrgb.h>
#include <QtGui/qrgba64.h>
#include <QtGui/qstatictext.h>
#include <QtGui/qtextcursor.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qdrawutil.h>
#include "gen_functions.h"

extern "C" {

void cppgofn_qt_qt_assert(const char* assertion, const char* file, int line) {
	qt_assert(assertion, file, line);
}
void cppgofn_qt_qt_assert_x(const char* where, const char* what, const char* file, int line) {
	qt_assert_x(where, what, file, line);
}
int cppgofn_qt_qt_no_assert_x(int arg1, const char* arg2, const char* arg3) {
	auto&& __r = qt_no_assert_x(arg1, arg2, arg3);
	return __r;
}
void cppgofn_qt_qt_check_pointer(const char* arg1, int arg2) {
	qt_check_pointer(arg1, arg2);
}
void cppgofn_qt_qBadAlloc() {
	qBadAlloc();
}
int cppgofn_qt_qt_assume_is_deprecated(int cond) {
	auto&& __r = qt_assume_is_deprecated(cond);
	return __r;
}
const char* cppgofn_qt_qVersion() {
	auto&& __r = qVersion();
	return __r;
}
void cppgofn_qt_qErrnoWarning(int code, const char* msg) {
	qErrnoWarning(code, msg);
}
void cppgofn_qt_qErrnoWarning2(const char* msg) {
	qErrnoWarning(msg);
}
void cppgofn_qt_qt_message_output(int arg1, void* context, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	qt_message_output((QtMsgType)(arg1), *(QMessageLogContext*)context, message_qstr);
}
void cppgofn_qt_qSetMessagePattern(const char* messagePattern, long long messagePattern_len) {
	QString messagePattern_qstr = QString::fromUtf8(messagePattern, messagePattern_len);
	qSetMessagePattern(messagePattern_qstr);
}
char* cppgofn_qt_qFormatLogMessage(int type_, void* context, const char* buf, long long buf_len, long long* __out_len) {
	QString buf_qstr = QString::fromUtf8(buf, buf_len);
	auto&& __r = qFormatLogMessage((QtMsgType)(type_), *(QMessageLogContext*)context, buf_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* cppgofn_qt_qt_error_string(int errorCode, long long* __out_len) {
	auto&& __r = qt_error_string(errorCode);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void cppgofn_qt_qYieldCpu() {
	qYieldCpu();
}
int cppgofn_qt_is_constant_evaluated() {
	auto&& __r = q20::is_constant_evaluated();
	return __r;
}
int cppgofn_qt_qIsInf(double d) {
	auto&& __r = qIsInf(d);
	return __r;
}
int cppgofn_qt_qIsNaN(double d) {
	auto&& __r = qIsNaN(d);
	return __r;
}
int cppgofn_qt_qIsFinite(double d) {
	auto&& __r = qIsFinite(d);
	return __r;
}
int cppgofn_qt_qFpClassify(double val) {
	auto&& __r = qFpClassify(val);
	return __r;
}
int cppgofn_qt_qIsInf2(float f) {
	auto&& __r = qIsInf(f);
	return __r;
}
int cppgofn_qt_qIsNaN2(float f) {
	auto&& __r = qIsNaN(f);
	return __r;
}
int cppgofn_qt_qIsFinite2(float f) {
	auto&& __r = qIsFinite(f);
	return __r;
}
int cppgofn_qt_qFpClassify2(float val) {
	auto&& __r = qFpClassify(val);
	return __r;
}
double cppgofn_qt_qSNaN() {
	auto&& __r = qSNaN();
	return __r;
}
double cppgofn_qt_qQNaN() {
	auto&& __r = qQNaN();
	return __r;
}
double cppgofn_qt_qInf() {
	auto&& __r = qInf();
	return __r;
}
unsigned int cppgofn_qt_qFloatDistance(float a, float b) {
	auto&& __r = qFloatDistance(a, b);
	return __r;
}
unsigned long long cppgofn_qt_qFloatDistance2(double a, double b) {
	auto&& __r = qFloatDistance(a, b);
	return __r;
}
int cppgofn_qt_qRound(double d) {
	auto&& __r = qRound(d);
	return __r;
}
int cppgofn_qt_qRound2(float f) {
	auto&& __r = qRound(f);
	return __r;
}
long long cppgofn_qt_qRound64(double d) {
	auto&& __r = qRound64(d);
	return __r;
}
long long cppgofn_qt_qRound642(float f) {
	auto&& __r = qRound64(f);
	return __r;
}
int cppgofn_qt_qFuzzyCompare(double p1, double p2) {
	auto&& __r = qFuzzyCompare(p1, p2);
	return __r;
}
int cppgofn_qt_qFuzzyCompare2(float p1, float p2) {
	auto&& __r = qFuzzyCompare(p1, p2);
	return __r;
}
int cppgofn_qt_qFuzzyIsNull(double d) {
	auto&& __r = qFuzzyIsNull(d);
	return __r;
}
int cppgofn_qt_qFuzzyIsNull2(float f) {
	auto&& __r = qFuzzyIsNull(f);
	return __r;
}
int cppgofn_qt_qIsNull(double d) {
	auto&& __r = qIsNull(d);
	return __r;
}
int cppgofn_qt_qIsNull2(float f) {
	auto&& __r = qIsNull(f);
	return __r;
}
int cppgofn_qt_qIntCast(double f) {
	auto&& __r = qIntCast(f);
	return __r;
}
int cppgofn_qt_qIntCast2(float f) {
	auto&& __r = qIntCast(f);
	return __r;
}
void* cppgofn_qt_qgetenv(const char* varName) {
	auto&& __r = qgetenv(varName);
	return new QByteArray(__r);
}
char* cppgofn_qt_qEnvironmentVariable(const char* varName, long long* __out_len) {
	auto&& __r = qEnvironmentVariable(varName);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* cppgofn_qt_qEnvironmentVariable2(const char* varName, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = qEnvironmentVariable(varName, defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int cppgofn_qt_qputenv(const char* varName, void* value) {
	auto&& __r = qputenv(varName, *(QByteArrayView*)value);
	return __r;
}
int cppgofn_qt_qunsetenv(const char* varName) {
	auto&& __r = qunsetenv(varName);
	return __r;
}
int cppgofn_qt_qEnvironmentVariableIsEmpty(const char* varName) {
	auto&& __r = qEnvironmentVariableIsEmpty(varName);
	return __r;
}
int cppgofn_qt_qEnvironmentVariableIsSet(const char* varName) {
	auto&& __r = qEnvironmentVariableIsSet(varName);
	return __r;
}
int cppgofn_qt_qEnvironmentVariableIntValue(const char* varName, int* ok) {
	bool ok_local{};
	auto&& __r = qEnvironmentVariableIntValue(varName, &ok_local);
	*ok = ok_local;
	return __r;
}
long long cppgofn_qt_qEnvironmentVariableIntegerValue(const char* varName, int* __ok) {
	auto&& __r = qEnvironmentVariableIntegerValue(varName);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? *__r : (long long)0;
}
char* cppgofn_qt_qtTrId(const char* id, int n, long long* __out_len) {
	auto&& __r = qtTrId(id, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* cppgofn_qt_qt_getEnumMetaObject(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ScrollBarPolicy)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ScrollBarPolicy)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject2(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::FocusPolicy)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName2(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::FocusPolicy)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject3(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ContextMenuPolicy)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName3(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ContextMenuPolicy)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject4(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ContextMenuTrigger)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName4(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ContextMenuTrigger)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject5(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ArrowType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName5(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ArrowType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject6(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ToolButtonStyle)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName6(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ToolButtonStyle)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject7(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::PenStyle)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName7(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::PenStyle)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject8(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::PenCapStyle)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName8(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::PenCapStyle)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject9(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::PenJoinStyle)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName9(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::PenJoinStyle)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject10(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::BrushStyle)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName10(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::BrushStyle)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject11(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::FillRule)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName11(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::FillRule)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject12(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::MaskMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName12(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::MaskMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject13(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::BGMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName13(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::BGMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject14(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ClipOperation)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName14(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ClipOperation)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject15(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::SizeMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName15(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::SizeMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject16(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::Axis)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName16(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::Axis)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject17(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::Corner)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName17(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::Corner)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject18(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::Edge)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName18(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::Edge)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject19(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::LayoutDirection)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName19(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::LayoutDirection)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject20(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::SizeHint)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName20(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::SizeHint)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject21(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::Orientation)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName21(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::Orientation)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject22(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::DropAction)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName22(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::DropAction)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject23(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::AlignmentFlag>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName23(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::AlignmentFlag>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject24(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TextFlag)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName24(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TextFlag)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject25(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::Orientation>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName25(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::Orientation>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject26(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::SplitBehaviorFlags>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName26(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::SplitBehaviorFlags>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject27(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::DropAction>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName27(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::DropAction>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject28(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::Edge>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName28(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::Edge>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject29(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::DockWidgetArea>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName29(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::DockWidgetArea>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject30(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::ToolBarArea>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName30(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::ToolBarArea>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject31(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::DockWidgetArea)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName31(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::DockWidgetArea)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject32(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ToolBarArea)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName32(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ToolBarArea)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject33(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TextFormat)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName33(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TextFormat)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject34(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TextElideMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName34(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TextElideMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject35(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::DateFormat)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName35(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::DateFormat)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject36(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TimeSpec)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName36(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TimeSpec)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject37(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::DayOfWeek)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName37(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::DayOfWeek)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject38(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::CursorShape)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName38(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::CursorShape)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject39(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::GlobalColor)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName39(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::GlobalColor)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject40(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ColorScheme)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName40(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ColorScheme)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject41(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ContrastPreference)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName41(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ContrastPreference)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject42(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::AspectRatioMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName42(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::AspectRatioMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject43(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TransformationMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName43(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TransformationMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject44(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::ImageConversionFlag>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName44(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::ImageConversionFlag>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject45(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::Key)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName45(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::Key)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject46(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ShortcutContext)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName46(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ShortcutContext)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject47(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TextInteractionFlag)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName47(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TextInteractionFlag)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject48(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::TextInteractionFlag>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName48(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::TextInteractionFlag>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject49(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ItemSelectionMode)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName49(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ItemSelectionMode)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject50(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ItemSelectionOperation)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName50(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ItemSelectionOperation)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject51(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::ItemFlag>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName51(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::ItemFlag>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject52(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::CheckState)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName52(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::CheckState)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject53(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ItemDataRole)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName53(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ItemDataRole)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject54(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::SortOrder)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName54(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::SortOrder)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject55(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::CaseSensitivity)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName55(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::CaseSensitivity)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject56(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::MatchFlag>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName56(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::MatchFlag>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject57(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::Modifier)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName57(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::Modifier)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject58(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::Modifier>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName58(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::Modifier>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject59(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::KeyboardModifier)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName59(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::KeyboardModifier)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject60(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::KeyboardModifier>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName60(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::KeyboardModifier>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject61(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::MouseButton>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName61(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::MouseButton>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject62(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::WindowType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName62(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::WindowType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject63(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::WindowState)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName63(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::WindowState)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject64(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::WindowModality)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName64(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::WindowModality)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject65(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::WidgetAttribute)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName65(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::WidgetAttribute)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject66(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ApplicationAttribute)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName66(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ApplicationAttribute)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject67(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::WindowType>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName67(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::WindowType>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject68(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::WindowState>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName68(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::WindowState>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject69(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::FocusReason)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName69(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::FocusReason)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject70(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::InputMethodHint)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName70(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::InputMethodHint)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject71(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::InputMethodQuery)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName71(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::InputMethodQuery)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject72(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::InputMethodHint>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName72(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::InputMethodHint>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject73(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::EnterKeyType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName73(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::EnterKeyType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject74(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::InputMethodQuery>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName74(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::InputMethodQuery>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject75(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::TouchPointState>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName75(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::TouchPointState>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject76(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ScreenOrientation)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName76(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ScreenOrientation)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject77(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::ScreenOrientation>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName77(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::ScreenOrientation>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject78(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ConnectionType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName78(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ConnectionType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject79(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ApplicationState)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName79(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ApplicationState)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject80(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::GestureState)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName80(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::GestureState)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject81(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::GestureType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName81(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::GestureType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject82(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::NativeGestureType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName82(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::NativeGestureType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject83(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::CursorMoveStyle)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName83(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::CursorMoveStyle)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject84(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TimerType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName84(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TimerType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject85(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TimerId)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName85(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TimerId)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject86(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ScrollPhase)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName86(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ScrollPhase)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject87(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::MouseEventSource)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName87(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::MouseEventSource)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject88(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject(QFlags<Qt::MouseEventFlag>(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName88(int arg1) {
	auto&& __r = Qt::qt_getEnumName(QFlags<Qt::MouseEventFlag>(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject89(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::ChecksumType)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName89(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::ChecksumType)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject90(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::HighDpiScaleFactorRoundingPolicy)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName90(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::HighDpiScaleFactorRoundingPolicy)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject91(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::TabFocusBehavior)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName91(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::TabFocusBehavior)(arg1));
	return __r;
}
void* cppgofn_qt_qt_getEnumMetaObject92(int arg1) {
	auto&& __r = Qt::qt_getEnumMetaObject((Qt::PermissionStatus)(arg1));
	return (void*)(__r);
}
const char* cppgofn_qt_qt_getEnumName92(int arg1) {
	auto&& __r = Qt::qt_getEnumName((Qt::PermissionStatus)(arg1));
	return __r;
}
const char* cppgofn_qt_qFlagLocation(const char* method) {
	auto&& __r = qFlagLocation(method);
	return __r;
}
void cppgofn_qt_swap(void* lhs, void* rhs) {
	swap(*(QMetaObject::Connection*)lhs, *(QMetaObject::Connection*)rhs);
}
unsigned long long cppgofn_qt_qstrlen(const char* str) {
	auto&& __r = qstrlen(str);
	return __r;
}
unsigned long long cppgofn_qt_qstrnlen(const char* str, unsigned long long maxlen) {
	auto&& __r = qstrnlen(str, maxlen);
	return __r;
}
int cppgofn_qt_qstrcmp(const char* str1, const char* str2) {
	auto&& __r = qstrcmp(str1, str2);
	return __r;
}
int cppgofn_qt_qstrncmp(const char* str1, const char* str2, unsigned long long len_) {
	auto&& __r = qstrncmp(str1, str2, len_);
	return __r;
}
int cppgofn_qt_qstricmp(const char* arg1, const char* arg2) {
	auto&& __r = qstricmp(arg1, arg2);
	return __r;
}
int cppgofn_qt_qstrnicmp(const char* arg1, const char* arg2, unsigned long long len_) {
	auto&& __r = qstrnicmp(arg1, arg2, len_);
	return __r;
}
int cppgofn_qt_qstrnicmp2(const char* arg1, long long arg2, const char* arg3, long long arg4) {
	auto&& __r = qstrnicmp(arg1, arg2, arg3, arg4);
	return __r;
}
int cppgofn_qt_qsnprintf(char* str, unsigned long long n, const char* fmt) {
	auto&& __r = qsnprintf((char*)str, n, fmt);
	return __r;
}
unsigned short cppgofn_qt_qChecksum(void* data, int standard) {
	auto&& __r = qChecksum(*(QByteArrayView*)data, (Qt::ChecksumType)(standard));
	return __r;
}
unsigned short cppgofn_qt_qChecksum2(const char* s, long long len_, int standard) {
	auto&& __r = qChecksum(s, len_, (Qt::ChecksumType)(standard));
	return __r;
}
void* cppgofn_qt_qCompress(const unsigned char* data, long long nbytes, int compressionLevel) {
	auto&& __r = qCompress(data, nbytes, compressionLevel);
	return new QByteArray(__r);
}
void* cppgofn_qt_qUncompress(const unsigned char* data, long long nbytes) {
	auto&& __r = qUncompress(data, nbytes);
	return new QByteArray(__r);
}
void* cppgofn_qt_qCompress2(void* data, int compressionLevel) {
	auto&& __r = qCompress(*(QByteArray*)data, compressionLevel);
	return new QByteArray(__r);
}
void* cppgofn_qt_qUncompress2(void* data) {
	auto&& __r = qUncompress(*(QByteArray*)data);
	return new QByteArray(__r);
}
void cppgofn_qt_swap2(void* value1, void* value2) {
	swap(*(QByteArray*)value1, *(QByteArray*)value2);
}
void cppgofn_qt_swap3(void* value1, void* value2) {
	swap(*(QByteArray::FromBase64Result*)value1, *(QByteArray::FromBase64Result*)value2);
}
unsigned long long cppgofn_qt_qHash(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QByteArray::FromBase64Result*)key, seed);
	return __r;
}
void cppgofn_qt_swap4(void* value1, void* value2) {
	swap(*(QString*)value1, *(QString*)value2);
}
int cppgofn_qt_qGlobalQHashSeed() {
	auto&& __r = qGlobalQHashSeed();
	return __r;
}
void cppgofn_qt_qSetGlobalQHashSeed(int newSeed) {
	qSetGlobalQHashSeed(newSeed);
}
unsigned long long cppgofn_qt_hash(unsigned long long key, unsigned long long seed) {
	auto&& __r = QHashPrivate::hash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash2(char key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash3(unsigned char key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash4(char key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash5(unsigned short key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash6(short key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash7(unsigned int key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash8(int key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash9(unsigned long key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash10(long key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash11(unsigned long long key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash12(long long key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash13(float key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash14(double key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash15(unsigned short key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash16(unsigned short key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash17(unsigned int key, unsigned long long seed) {
	auto&& __r = qHash(key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash18(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QChar*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash19(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QByteArrayView*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash20(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QStringView*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash21(const char* key, long long key_len, unsigned long long seed) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = qHash(key_qstr, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash22(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QBitArray*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash23(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QLatin1String*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash24(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QKeyCombination*)key, seed);
	return __r;
}
unsigned int cppgofn_qt_qt_hash(void* key, unsigned int chained) {
	auto&& __r = qt_hash(*(QStringView*)key, chained);
	return __r;
}
unsigned int cppgofn_qt_qPopulationCount(unsigned int v) {
	auto&& __r = qPopulationCount(v);
	return __r;
}
unsigned int cppgofn_qt_qPopulationCount2(unsigned char v) {
	auto&& __r = qPopulationCount(v);
	return __r;
}
unsigned int cppgofn_qt_qPopulationCount3(unsigned short v) {
	auto&& __r = qPopulationCount(v);
	return __r;
}
unsigned int cppgofn_qt_qPopulationCount4(unsigned long long v) {
	auto&& __r = qPopulationCount(v);
	return __r;
}
unsigned int cppgofn_qt_qPopulationCount5(unsigned long v) {
	auto&& __r = qPopulationCount(v);
	return __r;
}
unsigned int cppgofn_qt_qCountTrailingZeroBits(unsigned int v) {
	auto&& __r = qCountTrailingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountTrailingZeroBits2(unsigned char v) {
	auto&& __r = qCountTrailingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountTrailingZeroBits3(unsigned short v) {
	auto&& __r = qCountTrailingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountTrailingZeroBits4(unsigned long long v) {
	auto&& __r = qCountTrailingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountTrailingZeroBits5(unsigned long v) {
	auto&& __r = qCountTrailingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountLeadingZeroBits(unsigned int v) {
	auto&& __r = qCountLeadingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountLeadingZeroBits2(unsigned char v) {
	auto&& __r = qCountLeadingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountLeadingZeroBits3(unsigned short v) {
	auto&& __r = qCountLeadingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountLeadingZeroBits4(unsigned long long v) {
	auto&& __r = qCountLeadingZeroBits(v);
	return __r;
}
unsigned int cppgofn_qt_qCountLeadingZeroBits5(unsigned long v) {
	auto&& __r = qCountLeadingZeroBits(v);
	return __r;
}
double cppgofn_qt_qFastSin(double x) {
	auto&& __r = qFastSin(x);
	return __r;
}
double cppgofn_qt_qFastCos(double x) {
	auto&& __r = qFastCos(x);
	return __r;
}
float cppgofn_qt_qDegreesToRadians(float degrees) {
	auto&& __r = qDegreesToRadians(degrees);
	return __r;
}
double cppgofn_qt_qDegreesToRadians2(double degrees) {
	auto&& __r = qDegreesToRadians(degrees);
	return __r;
}
float cppgofn_qt_qRadiansToDegrees(float radians) {
	auto&& __r = qRadiansToDegrees(radians);
	return __r;
}
double cppgofn_qt_qRadiansToDegrees2(double radians) {
	auto&& __r = qRadiansToDegrees(radians);
	return __r;
}
unsigned int cppgofn_qt_qNextPowerOfTwo(unsigned int v) {
	auto&& __r = qNextPowerOfTwo(v);
	return __r;
}
unsigned long long cppgofn_qt_qNextPowerOfTwo2(unsigned long long v) {
	auto&& __r = qNextPowerOfTwo(v);
	return __r;
}
unsigned int cppgofn_qt_qNextPowerOfTwo3(int v) {
	auto&& __r = qNextPowerOfTwo(v);
	return __r;
}
unsigned long long cppgofn_qt_qNextPowerOfTwo4(long long v) {
	auto&& __r = qNextPowerOfTwo(v);
	return __r;
}
unsigned long cppgofn_qt_qNextPowerOfTwo5(unsigned long v) {
	auto&& __r = qNextPowerOfTwo(v);
	return __r;
}
unsigned long cppgofn_qt_qNextPowerOfTwo6(long v) {
	auto&& __r = qNextPowerOfTwo(v);
	return __r;
}
void cppgofn_qt_qFloatFromFloat16(float* arg1, void* arg2, long long length) {
	float arg1_local{};
	qFloatFromFloat16(&arg1_local, (qfloat16*)arg2, length);
	*arg1 = arg1_local;
}
int cppgofn_qt_qIsInf3(void* f) {
	auto&& __r = qIsInf(*(qfloat16*)f);
	return __r;
}
int cppgofn_qt_qIsNaN3(void* f) {
	auto&& __r = qIsNaN(*(qfloat16*)f);
	return __r;
}
int cppgofn_qt_qIsFinite3(void* f) {
	auto&& __r = qIsFinite(*(qfloat16*)f);
	return __r;
}
int cppgofn_qt_qFpClassify3(void* f) {
	auto&& __r = qFpClassify(*(qfloat16*)f);
	return __r;
}
void* cppgofn_qt_qSqrt(void* f) {
	auto&& __r = qSqrt(*(qfloat16*)f);
	return new qfloat16(__r);
}
int cppgofn_qt_qRound3(void* d) {
	auto&& __r = qRound(*(qfloat16*)d);
	return __r;
}
long long cppgofn_qt_qRound643(void* d) {
	auto&& __r = qRound64(*(qfloat16*)d);
	return __r;
}
int cppgofn_qt_qFuzzyCompare3(void* p1, void* p2) {
	auto&& __r = qFuzzyCompare(*(qfloat16*)p1, *(qfloat16*)p2);
	return __r;
}
int cppgofn_qt_qFuzzyIsNull3(void* f) {
	auto&& __r = qFuzzyIsNull(*(qfloat16*)f);
	return __r;
}
int cppgofn_qt_qIsNull3(void* f) {
	auto&& __r = qIsNull(*(qfloat16*)f);
	return __r;
}
int cppgofn_qt_qIntCast3(void* f) {
	auto&& __r = qIntCast(*(qfloat16*)f);
	return __r;
}
void* cppgofn_qt_qHypot(void* x, void* y) {
	auto&& __r = qHypot(*(qfloat16*)x, *(qfloat16*)y);
	return new qfloat16(__r);
}
void* cppgofn_qt_qHypot2(void* x, void* y, void* z) {
	auto&& __r = qHypot(*(qfloat16*)x, *(qfloat16*)y, *(qfloat16*)z);
	return new qfloat16(__r);
}
int cppgofn_qt_qRegisterMetaType(void* meta) {
	auto&& __r = qRegisterMetaType(*(QMetaType*)meta);
	return __r;
}
unsigned long long cppgofn_qt_qHash25(void* type_, unsigned long long seed) {
	auto&& __r = qHash(*(QMetaType*)type_, seed);
	return __r;
}
void* cppgofn_qt_qt_qFindChild_helper(void* parent, const char* name, long long name_len, void* mo, int options) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = qt_qFindChild_helper((QObject*)parent, name_qstr, *(QMetaObject*)mo, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* cppgofn_qt_qGetBindingStorage(void* o) {
	auto&& __r = qGetBindingStorage((QObject*)o);
	return (void*)(__r);
}
void* cppgofn_qt_qGetBindingStorage2(void* o) {
	auto&& __r = qGetBindingStorage((QObject*)o);
	return (void*)(__r);
}
void cppgofn_qt_swap5(void* value1, void* value2) {
	swap(*(QDeadlineTimer*)value1, *(QDeadlineTimer*)value2);
}
unsigned long long cppgofn_qt_bucketsForCapacity(unsigned long long requestedCapacity) {
	auto&& __r = QHashPrivate::GrowthPolicy::bucketsForCapacity(requestedCapacity);
	return __r;
}
unsigned long long cppgofn_qt_bucketForHash(unsigned long long nBuckets, unsigned long long hash) {
	auto&& __r = QHashPrivate::GrowthPolicy::bucketForHash(nBuckets, hash);
	return __r;
}
void* cppgofn_qt_bin(void* s) {
	auto&& __r = Qt::bin(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_oct(void* s) {
	auto&& __r = Qt::oct(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_dec(void* s) {
	auto&& __r = Qt::dec(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_hex(void* s) {
	auto&& __r = Qt::hex(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_showbase(void* s) {
	auto&& __r = Qt::showbase(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_forcesign(void* s) {
	auto&& __r = Qt::forcesign(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_forcepoint(void* s) {
	auto&& __r = Qt::forcepoint(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_noshowbase(void* s) {
	auto&& __r = Qt::noshowbase(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_noforcesign(void* s) {
	auto&& __r = Qt::noforcesign(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_noforcepoint(void* s) {
	auto&& __r = Qt::noforcepoint(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_uppercasebase(void* s) {
	auto&& __r = Qt::uppercasebase(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_uppercasedigits(void* s) {
	auto&& __r = Qt::uppercasedigits(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_lowercasebase(void* s) {
	auto&& __r = Qt::lowercasebase(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_lowercasedigits(void* s) {
	auto&& __r = Qt::lowercasedigits(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_fixed(void* s) {
	auto&& __r = Qt::fixed(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_scientific(void* s) {
	auto&& __r = Qt::scientific(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_left(void* s) {
	auto&& __r = Qt::left(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_right(void* s) {
	auto&& __r = Qt::right(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_center(void* s) {
	auto&& __r = Qt::center(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_endl(void* s) {
	auto&& __r = Qt::endl(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_flush(void* s) {
	auto&& __r = Qt::flush(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_reset(void* s) {
	auto&& __r = Qt::reset(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_bom(void* s) {
	auto&& __r = Qt::bom(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_ws(void* s) {
	auto&& __r = Qt::ws(*(QTextStream*)s);
	return (void*)&(__r);
}
void* cppgofn_qt_qSetFieldWidth(int width) {
	auto&& __r = qSetFieldWidth(width);
	return new QTextStreamManipulator(__r);
}
void* cppgofn_qt_qSetPadChar(void* ch) {
	auto&& __r = qSetPadChar(*(QChar*)ch);
	return new QTextStreamManipulator(__r);
}
void* cppgofn_qt_qSetRealNumberPrecision(int precision) {
	auto&& __r = qSetRealNumberPrecision(precision);
	return new QTextStreamManipulator(__r);
}
void cppgofn_qt_swap6(void* value1, void* value2) {
	swap(*(QDebug*)value1, *(QDebug*)value2);
}
void* cppgofn_qt_qt_QMetaEnum_debugOperator(void* arg1, long long value, void* meta, const char* name) {
	auto&& __r = qt_QMetaEnum_debugOperator(*(QDebug*)arg1, value, (QMetaObject*)meta, name);
	return new QDebug(__r);
}
void* cppgofn_qt_qt_QMetaEnum_flagDebugOperator(void* dbg, unsigned long long value, void* meta, const char* name) {
	auto&& __r = qt_QMetaEnum_flagDebugOperator(*(QDebug*)dbg, value, (QMetaObject*)meta, name);
	return new QDebug(__r);
}
void cppgofn_qt_qt_QMetaEnum_flagDebugOperator2(void* debug, unsigned long long sizeofT, unsigned int value) {
	qt_QMetaEnum_flagDebugOperator(*(QDebug*)debug, sizeofT, value);
}
void cppgofn_qt_qt_QMetaEnum_flagDebugOperator3(void* debug, unsigned long long sizeofT, unsigned long long value) {
	qt_QMetaEnum_flagDebugOperator(*(QDebug*)debug, sizeofT, value);
}
void cppgofn_qt_swap7(void* value1, void* value2) {
	swap(*(QVariant*)value1, *(QVariant*)value2);
}
unsigned long long cppgofn_qt_qHash26(void* index, unsigned long long seed) {
	auto&& __r = qHash(*(QPersistentModelIndex*)index, seed);
	return __r;
}
void cppgofn_qt_swap8(void* value1, void* value2) {
	swap(*(QPersistentModelIndex*)value1, *(QPersistentModelIndex*)value2);
}
unsigned long long cppgofn_qt_qHash27(void* index, unsigned long long seed) {
	auto&& __r = qHash(*(QModelIndex*)index, seed);
	return __r;
}
void cppgofn_qt_swap9(void* lhs, void* rhs) {
	swap(*(QBasicTimer*)lhs, *(QBasicTimer*)rhs);
}
char* cppgofn_qt_qAppName(long long* __out_len) {
	auto&& __r = qAppName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void cppgofn_qt_swap10(void* value1, void* value2) {
	swap(*(QBitArray*)value1, *(QBitArray*)value2);
}
unsigned long long cppgofn_qt_qHash28(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QLocale*)key, seed);
	return __r;
}
void cppgofn_qt_swap11(void* value1, void* value2) {
	swap(*(QLocale*)value1, *(QLocale*)value2);
}
unsigned long long cppgofn_qt_qHash29(int tag, unsigned long long seed) {
	auto&& __r = qHash((QCborSimpleType)(tag), seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash30(int tag, unsigned long long seed) {
	auto&& __r = qHash((QCborTag)(tag), seed);
	return __r;
}
void cppgofn_qt_swap12(void* value1, void* value2) {
	swap(*(QDateTime*)value1, *(QDateTime*)value2);
}
unsigned long long cppgofn_qt_qHash31(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QDateTime*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash32(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QDate*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash33(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QTime*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash34(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QRegularExpression*)key, seed);
	return __r;
}
void cppgofn_qt_swap13(void* value1, void* value2) {
	swap(*(QRegularExpression*)value1, *(QRegularExpression*)value2);
}
void cppgofn_qt_swap14(void* value1, void* value2) {
	swap(*(QRegularExpressionMatch*)value1, *(QRegularExpressionMatch*)value2);
}
void cppgofn_qt_swap15(void* value1, void* value2) {
	swap(*(QRegularExpressionMatchIterator*)value1, *(QRegularExpressionMatchIterator*)value2);
}
unsigned long long cppgofn_qt_qHash35(void* url, unsigned long long seed) {
	auto&& __r = qHash(*(QUrl*)url, seed);
	return __r;
}
void cppgofn_qt_swap16(void* value1, void* value2) {
	swap(*(QUrl*)value1, *(QUrl*)value2);
}
unsigned long long cppgofn_qt_qbswap_helper(unsigned long long source) {
	auto&& __r = qbswap_helper(source);
	return __r;
}
unsigned int cppgofn_qt_qbswap_helper2(unsigned int source) {
	auto&& __r = qbswap_helper(source);
	return __r;
}
unsigned short cppgofn_qt_qbswap_helper3(unsigned short source) {
	auto&& __r = qbswap_helper(source);
	return __r;
}
unsigned char cppgofn_qt_qbswap_helper4(unsigned char source) {
	auto&& __r = qbswap_helper(source);
	return __r;
}
void* cppgofn_qt_qbswap(void* source) {
	auto&& __r = qbswap(*(qfloat16*)source);
	return new qfloat16(__r);
}
float cppgofn_qt_qbswap2(float source) {
	auto&& __r = qbswap(source);
	return __r;
}
double cppgofn_qt_qbswap3(double source) {
	auto&& __r = qbswap(source);
	return __r;
}
unsigned long long cppgofn_qt_qHash36(void* uuid, unsigned long long seed) {
	auto&& __r = qHash(*(QUuid*)uuid, seed);
	return __r;
}
void cppgofn_qt_swap17(void* value1, void* value2) {
	swap(*(QCborValue*)value1, *(QCborValue*)value2);
}
unsigned long long cppgofn_qt_qHash37(void* value, unsigned long long seed) {
	auto&& __r = qHash(*(QCborValue*)value, seed);
	return __r;
}
void cppgofn_qt_swap18(void* value1, void* value2) {
	swap(*(QCborArray*)value1, *(QCborArray*)value2);
}
unsigned long long cppgofn_qt_qHash38(void* array, unsigned long long seed) {
	auto&& __r = qHash(*(QCborArray*)array, seed);
	return __r;
}
void cppgofn_qt_swap19(void* value1, void* value2) {
	swap(*(QCborMap*)value1, *(QCborMap*)value2);
}
unsigned long long cppgofn_qt_qHash39(void* map_, unsigned long long seed) {
	auto&& __r = qHash(*(QCborMap*)map_, seed);
	return __r;
}
void cppgofn_qt_beginPropertyUpdateGroup() {
	Qt::beginPropertyUpdateGroup();
}
void cppgofn_qt_endPropertyUpdateGroup() {
	Qt::endPropertyUpdateGroup();
}
void cppgofn_qt_swap20(void* value1, void* value2) {
	swap(*(QCollatorSortKey*)value1, *(QCollatorSortKey*)value2);
}
void cppgofn_qt_swap21(void* value1, void* value2) {
	swap(*(QCollator*)value1, *(QCollator*)value2);
}
void cppgofn_qt_swap22(void* value1, void* value2) {
	swap(*(QCommandLineOption*)value1, *(QCommandLineOption*)value2);
}
int cppgofn_qt_qEnableNtfsPermissionChecks() {
	auto&& __r = qEnableNtfsPermissionChecks();
	return __r;
}
int cppgofn_qt_qDisableNtfsPermissionChecks() {
	auto&& __r = qDisableNtfsPermissionChecks();
	return __r;
}
int cppgofn_qt_qAreNtfsPermissionChecksEnabled() {
	auto&& __r = qAreNtfsPermissionChecksEnabled();
	return __r;
}
void cppgofn_qt_swap23(void* value1, void* value2) {
	swap(*(QTimeZone*)value1, *(QTimeZone*)value2);
}
void cppgofn_qt_swap24(void* value1, void* value2) {
	swap(*(QFileInfo*)value1, *(QFileInfo*)value2);
}
void cppgofn_qt_swap25(void* value1, void* value2) {
	swap(*(QDir*)value1, *(QDir*)value2);
}
void cppgofn_qt_swap26(void* value1, void* value2) {
	swap(*(QEasingCurve*)value1, *(QEasingCurve*)value2);
}
void cppgofn_qt_swap27(void* lhs, void* rhs) {
	swap(*(QFutureInterfaceBase*)lhs, *(QFutureInterfaceBase*)rhs);
}
void cppgofn_qt_swap28(void* value1, void* value2) {
	swap(*(QItemSelection*)value1, *(QItemSelection*)value2);
}
void cppgofn_qt_swap29(void* value1, void* value2) {
	swap(*(QJsonDocument*)value1, *(QJsonDocument*)value2);
}
void cppgofn_qt_swap30(void* value1, void* value2) {
	swap(*(QJsonValue*)value1, *(QJsonValue*)value2);
}
unsigned long long cppgofn_qt_qHash40(void* value, unsigned long long seed) {
	auto&& __r = qHash(*(QJsonValue*)value, seed);
	return __r;
}
void cppgofn_qt_swap31(void* value1, void* value2) {
	swap(*(QJsonArray*)value1, *(QJsonArray*)value2);
}
unsigned long long cppgofn_qt_qHash41(void* array, unsigned long long seed) {
	auto&& __r = qHash(*(QJsonArray*)array, seed);
	return __r;
}
void cppgofn_qt_swap32(void* value1, void* value2) {
	swap(*(QJsonObject*)value1, *(QJsonObject*)value2);
}
unsigned long long cppgofn_qt_qHash42(void* object, unsigned long long seed) {
	auto&& __r = qHash(*(QJsonObject*)object, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash43(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QTypeRevision*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash44(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QVersionNumber*)key, seed);
	return __r;
}
int cppgofn_qt_qSharedBuild() {
	auto&& __r = qSharedBuild();
	return __r;
}
unsigned long long cppgofn_qt_qHash45(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QPoint*)key, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash46(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QMimeType*)key, seed);
	return __r;
}
void cppgofn_qt_swap33(void* value1, void* value2) {
	swap(*(QMimeType*)value1, *(QMimeType*)value2);
}
void cppgofn_qt_swap34(void* value1, void* value2) {
	swap(*(QLocationPermission*)value1, *(QLocationPermission*)value2);
}
void cppgofn_qt_swap35(void* value1, void* value2) {
	swap(*(QCalendarPermission*)value1, *(QCalendarPermission*)value2);
}
void cppgofn_qt_swap36(void* value1, void* value2) {
	swap(*(QContactsPermission*)value1, *(QContactsPermission*)value2);
}
void cppgofn_qt_swap37(void* value1, void* value2) {
	swap(*(QBluetoothPermission*)value1, *(QBluetoothPermission*)value2);
}
void cppgofn_qt_swap38(void* value1, void* value2) {
	swap(*(QCameraPermission*)value1, *(QCameraPermission*)value2);
}
void cppgofn_qt_swap39(void* value1, void* value2) {
	swap(*(QMicrophonePermission*)value1, *(QMicrophonePermission*)value2);
}
unsigned char cppgofn_qt_qPluginArchRequirements() {
	auto&& __r = qPluginArchRequirements();
	return __r;
}
void cppgofn_qt_qRegisterStaticPluginFunction(void* staticPlugin) {
	qRegisterStaticPluginFunction(*(QStaticPlugin*)staticPlugin);
}
void cppgofn_qt_swap40(void* value1, void* value2) {
	swap(*(QProcessEnvironment*)value1, *(QProcessEnvironment*)value2);
}
unsigned long long cppgofn_qt_qHash47(void* s, unsigned long long seed) {
	auto&& __r = qHash(*(QSize*)s, seed);
	return __r;
}
unsigned long long cppgofn_qt_qHash48(void* r, unsigned long long seed) {
	auto&& __r = qHash(*(QRect*)r, seed);
	return __r;
}
void cppgofn_qt_swap41(void* value1, void* value2) {
	swap(*(QNativeIpcKey*)value1, *(QNativeIpcKey*)value2);
}
void cppgofn_qt_swap42(void* value1, void* value2) {
	swap(*(QStorageInfo*)value1, *(QStorageInfo*)value2);
}
void cppgofn_qt_swap43(void* lhs, void* rhs) {
	swap(*(QTemporaryDir*)lhs, *(QTemporaryDir*)rhs);
}
void cppgofn_qt_swap44(void* value1, void* value2) {
	swap(*(QTextBoundaryFinder*)value1, *(QTextBoundaryFinder*)value2);
}
unsigned long long cppgofn_qt_qHash49(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QUrlQuery*)key, seed);
	return __r;
}
void cppgofn_qt_swap45(void* value1, void* value2) {
	swap(*(QUrlQuery*)value1, *(QUrlQuery*)value2);
}
int cppgofn_qt_qRed(unsigned int rgb) {
	auto&& __r = qRed(rgb);
	return __r;
}
int cppgofn_qt_qGreen(unsigned int rgb) {
	auto&& __r = qGreen(rgb);
	return __r;
}
int cppgofn_qt_qBlue(unsigned int rgb) {
	auto&& __r = qBlue(rgb);
	return __r;
}
int cppgofn_qt_qAlpha(unsigned int rgb) {
	auto&& __r = qAlpha(rgb);
	return __r;
}
unsigned int cppgofn_qt_qRgb(int r, int g, int b) {
	auto&& __r = qRgb(r, g, b);
	return __r;
}
unsigned int cppgofn_qt_qRgba(int r, int g, int b, int a) {
	auto&& __r = qRgba(r, g, b, a);
	return __r;
}
int cppgofn_qt_qGray(int r, int g, int b) {
	auto&& __r = qGray(r, g, b);
	return __r;
}
int cppgofn_qt_qGray2(unsigned int rgb) {
	auto&& __r = qGray(rgb);
	return __r;
}
int cppgofn_qt_qIsGray(unsigned int rgb) {
	auto&& __r = qIsGray(rgb);
	return __r;
}
unsigned int cppgofn_qt_qPremultiply(unsigned int x) {
	auto&& __r = qPremultiply(x);
	return __r;
}
unsigned int cppgofn_qt_qUnpremultiply(unsigned int p) {
	auto&& __r = qUnpremultiply(p);
	return __r;
}
void* cppgofn_qt_qRgba64(unsigned short r, unsigned short g, unsigned short b, unsigned short a) {
	auto&& __r = qRgba64(r, g, b, a);
	return new QRgba64(__r);
}
void* cppgofn_qt_qRgba642(unsigned long long c) {
	auto&& __r = qRgba64(c);
	return new QRgba64(__r);
}
void* cppgofn_qt_qPremultiply2(void* c) {
	auto&& __r = qPremultiply(*(QRgba64*)c);
	return new QRgba64(__r);
}
void* cppgofn_qt_qUnpremultiply2(void* c) {
	auto&& __r = qUnpremultiply(*(QRgba64*)c);
	return new QRgba64(__r);
}
unsigned int cppgofn_qt_qRed2(void* rgb) {
	auto&& __r = qRed(*(QRgba64*)rgb);
	return __r;
}
unsigned int cppgofn_qt_qGreen2(void* rgb) {
	auto&& __r = qGreen(*(QRgba64*)rgb);
	return __r;
}
unsigned int cppgofn_qt_qBlue2(void* rgb) {
	auto&& __r = qBlue(*(QRgba64*)rgb);
	return __r;
}
unsigned int cppgofn_qt_qAlpha2(void* rgb) {
	auto&& __r = qAlpha(*(QRgba64*)rgb);
	return __r;
}
void* cppgofn_qt_qPixelFormatRgba(unsigned char red, unsigned char green, unsigned char blue, unsigned char alfa, int usage, int position, int pmul, int typeInt) {
	auto&& __r = qPixelFormatRgba(red, green, blue, alfa, (QPixelFormat::AlphaUsage)(usage), (QPixelFormat::AlphaPosition)(position), (QPixelFormat::AlphaPremultiplied)(pmul), (QPixelFormat::TypeInterpretation)(typeInt));
	return new QPixelFormat(__r);
}
void* cppgofn_qt_qPixelFormatGrayscale(unsigned char channelSize, int typeInt) {
	auto&& __r = qPixelFormatGrayscale(channelSize, (QPixelFormat::TypeInterpretation)(typeInt));
	return new QPixelFormat(__r);
}
void* cppgofn_qt_qPixelFormatAlpha(unsigned char channelSize, int typeInt) {
	auto&& __r = qPixelFormatAlpha(channelSize, (QPixelFormat::TypeInterpretation)(typeInt));
	return new QPixelFormat(__r);
}
void* cppgofn_qt_qPixelFormatCmyk(unsigned char channelSize, unsigned char alfa, int usage, int position, int typeInt) {
	auto&& __r = qPixelFormatCmyk(channelSize, alfa, (QPixelFormat::AlphaUsage)(usage), (QPixelFormat::AlphaPosition)(position), (QPixelFormat::TypeInterpretation)(typeInt));
	return new QPixelFormat(__r);
}
void* cppgofn_qt_qPixelFormatHsl(unsigned char channelSize, unsigned char alfa, int usage, int position, int typeInt) {
	auto&& __r = qPixelFormatHsl(channelSize, alfa, (QPixelFormat::AlphaUsage)(usage), (QPixelFormat::AlphaPosition)(position), (QPixelFormat::TypeInterpretation)(typeInt));
	return new QPixelFormat(__r);
}
void* cppgofn_qt_qPixelFormatHsv(unsigned char channelSize, unsigned char alfa, int usage, int position, int typeInt) {
	auto&& __r = qPixelFormatHsv(channelSize, alfa, (QPixelFormat::AlphaUsage)(usage), (QPixelFormat::AlphaPosition)(position), (QPixelFormat::TypeInterpretation)(typeInt));
	return new QPixelFormat(__r);
}
void* cppgofn_qt_qPixelFormatYuv(int layout, unsigned char alfa, int usage, int position, int p_mul, int typeInt, int b_order) {
	auto&& __r = qPixelFormatYuv((QPixelFormat::YUVLayout)(layout), alfa, (QPixelFormat::AlphaUsage)(usage), (QPixelFormat::AlphaPosition)(position), (QPixelFormat::AlphaPremultiplied)(p_mul), (QPixelFormat::TypeInterpretation)(typeInt), (QPixelFormat::ByteOrder)(b_order));
	return new QPixelFormat(__r);
}
void cppgofn_qt_swap46(void* value1, void* value2) {
	swap(*(QPolygon*)value1, *(QPolygon*)value2);
}
void cppgofn_qt_swap47(void* value1, void* value2) {
	swap(*(QPolygonF*)value1, *(QPolygonF*)value2);
}
void cppgofn_qt_swap48(void* value1, void* value2) {
	swap(*(QRegion*)value1, *(QRegion*)value2);
}
unsigned long long cppgofn_qt_qHash50(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QTransform*)key, seed);
	return __r;
}
int cppgofn_qt_qFuzzyCompare4(void* t1, void* t2) {
	auto&& __r = qFuzzyCompare(*(QTransform*)t1, *(QTransform*)t2);
	return __r;
}
void cppgofn_qt_swap49(void* value1, void* value2) {
	swap(*(QImage*)value1, *(QImage*)value2);
}
void cppgofn_qt_swap50(void* value1, void* value2) {
	swap(*(QPixmap*)value1, *(QPixmap*)value2);
}
void cppgofn_qt_swap51(void* value1, void* value2) {
	swap(*(QIcon*)value1, *(QIcon*)value2);
}
char* cppgofn_qt_qt_findAtNxFile(const char* baseFileName, long long baseFileName_len, double targetDevicePixelRatio, double* sourceDevicePixelRatio, long long* __out_len) {
	QString baseFileName_qstr = QString::fromUtf8(baseFileName, baseFileName_len);
	qreal sourceDevicePixelRatio_local{};
	auto&& __r = qt_findAtNxFile(baseFileName_qstr, targetDevicePixelRatio, &sourceDevicePixelRatio_local);
	*sourceDevicePixelRatio = sourceDevicePixelRatio_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
unsigned long long cppgofn_qt_qHash51(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QPointingDeviceUniqueId*)key, seed);
	return __r;
}
void cppgofn_qt_swap52(void* value1, void* value2) {
	swap(*(QEventPoint*)value1, *(QEventPoint*)value2);
}
unsigned long long cppgofn_qt_qHash52(void* key, unsigned long long seed) {
	auto&& __r = qHash(*(QKeySequence*)key, seed);
	return __r;
}
void cppgofn_qt_swap53(void* value1, void* value2) {
	swap(*(QKeySequence*)value1, *(QKeySequence*)value2);
}
void cppgofn_qt_swap54(void* value1, void* value2) {
	swap(*(QFont*)value1, *(QFont*)value2);
}
unsigned long long cppgofn_qt_qHash53(void* font, unsigned long long seed) {
	auto&& __r = qHash(*(QFont*)font, seed);
	return __r;
}
void cppgofn_qt_swap55(void* value1, void* value2) {
	swap(*(QRawFont*)value1, *(QRawFont*)value2);
}
unsigned long long cppgofn_qt_qHash54(void* font, unsigned long long seed) {
	auto&& __r = qHash(*(QRawFont*)font, seed);
	return __r;
}
void cppgofn_qt_swap56(void* value1, void* value2) {
	swap(*(QGlyphRun*)value1, *(QGlyphRun*)value2);
}
int cppgofn_qt_mightBeRichText(const char* arg1, long long arg1_len) {
	QAnyStringView arg1_qstr = QAnyStringView(arg1, arg1_len);
	auto&& __r = Qt::mightBeRichText(arg1_qstr);
	return __r;
}
char* cppgofn_qt_convertFromPlainText(const char* plain, long long plain_len, int mode, long long* __out_len) {
	QString plain_qstr = QString::fromUtf8(plain, plain_len);
	auto&& __r = Qt::convertFromPlainText(plain_qstr, (Qt::WhiteSpaceMode)(mode));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void cppgofn_qt_swap57(void* value1, void* value2) {
	swap(*(QBrush*)value1, *(QBrush*)value2);
}
void cppgofn_qt_swap58(void* value1, void* value2) {
	swap(*(QPen*)value1, *(QPen*)value2);
}
void cppgofn_qt_swap59(void* value1, void* value2) {
	swap(*(QTextFormat*)value1, *(QTextFormat*)value2);
}
void cppgofn_qt_swap60(void* value1, void* value2) {
	swap(*(QTextCharFormat*)value1, *(QTextCharFormat*)value2);
}
void cppgofn_qt_swap61(void* value1, void* value2) {
	swap(*(QTextBlockFormat*)value1, *(QTextBlockFormat*)value2);
}
void cppgofn_qt_swap62(void* value1, void* value2) {
	swap(*(QTextListFormat*)value1, *(QTextListFormat*)value2);
}
void cppgofn_qt_swap63(void* value1, void* value2) {
	swap(*(QTextImageFormat*)value1, *(QTextImageFormat*)value2);
}
void cppgofn_qt_swap64(void* value1, void* value2) {
	swap(*(QTextFrameFormat*)value1, *(QTextFrameFormat*)value2);
}
void cppgofn_qt_swap65(void* value1, void* value2) {
	swap(*(QTextTableFormat*)value1, *(QTextTableFormat*)value2);
}
void cppgofn_qt_swap66(void* value1, void* value2) {
	swap(*(QTextTableCellFormat*)value1, *(QTextTableCellFormat*)value2);
}
void cppgofn_qt_swap67(void* value1, void* value2) {
	swap(*(QTextCursor*)value1, *(QTextCursor*)value2);
}
void cppgofn_qt_swap68(void* value1, void* value2) {
	swap(*(QPalette*)value1, *(QPalette*)value2);
}
const char* cppgofn_qt_qAccessibleRoleString(int role) {
	auto&& __r = qAccessibleRoleString((QAccessible::Role)(role));
	return __r;
}
const char* cppgofn_qt_qAccessibleEventString(int event) {
	auto&& __r = qAccessibleEventString((QAccessible::Event)(event));
	return __r;
}
char* cppgofn_qt_qAccessibleLocalizedActionDescription(const char* actionName, long long actionName_len, long long* __out_len) {
	QString actionName_qstr = QString::fromUtf8(actionName, actionName_len);
	auto&& __r = qAccessibleLocalizedActionDescription(actionName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void cppgofn_qt_swap69(void* value1, void* value2) {
	swap(*(QBitmap*)value1, *(QBitmap*)value2);
}
void cppgofn_qt_swap70(void* value1, void* value2) {
	swap(*(QCursor*)value1, *(QCursor*)value2);
}
void cppgofn_qt_swap71(void* value1, void* value2) {
	swap(*(QColorTransform*)value1, *(QColorTransform*)value2);
}
void cppgofn_qt_swap72(void* value1, void* value2) {
	swap(*(QColorSpace*)value1, *(QColorSpace*)value2);
}
void cppgofn_qt_swap73(void* value1, void* value2) {
	swap(*(QFontVariableAxis*)value1, *(QFontVariableAxis*)value2);
}
void cppgofn_qt_swap74(void* value1, void* value2) {
	swap(*(QFontInfo*)value1, *(QFontInfo*)value2);
}
void cppgofn_qt_swap75(void* value1, void* value2) {
	swap(*(QFontMetrics*)value1, *(QFontMetrics*)value2);
}
void cppgofn_qt_swap76(void* value1, void* value2) {
	swap(*(QFontMetricsF*)value1, *(QFontMetricsF*)value2);
}
void cppgofn_qt_swap77(void* value1, void* value2) {
	swap(*(QPageSize*)value1, *(QPageSize*)value2);
}
void cppgofn_qt_swap78(void* value1, void* value2) {
	swap(*(QPageLayout*)value1, *(QPageLayout*)value2);
}
void cppgofn_qt_swap79(void* value1, void* value2) {
	swap(*(QPageRanges*)value1, *(QPageRanges*)value2);
}
void cppgofn_qt_swap80(void* value1, void* value2) {
	swap(*(QPainterPath*)value1, *(QPainterPath*)value2);
}
void cppgofn_qt_swap81(void* value1, void* value2) {
	swap(*(QPdfOutputIntent*)value1, *(QPdfOutputIntent*)value2);
}
void cppgofn_qt_swap82(void* value1, void* value2) {
	swap(*(QPicture*)value1, *(QPicture*)value2);
}
void cppgofn_qt_swap83(void* value1, void* value2) {
	swap(*(QPixmapCache::Key*)value1, *(QPixmapCache::Key*)value2);
}
int cppgofn_qt_qFuzzyCompare5(void* q1, void* q2) {
	auto&& __r = qFuzzyCompare(*(QQuaternion*)q1, *(QQuaternion*)q2);
	return __r;
}
void cppgofn_qt_swap84(void* value1, void* value2) {
	swap(*(QStaticText*)value1, *(QStaticText*)value2);
}
void cppgofn_qt_qDrawShadeLine(void* p, int x1, int y1, int x2, int y2, void* pal, int sunken, int lineWidth, int midLineWidth) {
	qDrawShadeLine((QPainter*)p, x1, y1, x2, y2, *(QPalette*)pal, sunken, lineWidth, midLineWidth);
}
void cppgofn_qt_qDrawShadeLine2(void* p, void* p1, void* p2, void* pal, int sunken, int lineWidth, int midLineWidth) {
	qDrawShadeLine((QPainter*)p, *(QPoint*)p1, *(QPoint*)p2, *(QPalette*)pal, sunken, lineWidth, midLineWidth);
}
void cppgofn_qt_qDrawShadeRect(void* p, int x, int y, int w, int h, void* pal, int sunken, int lineWidth, int midLineWidth, void* fill) {
	qDrawShadeRect((QPainter*)p, x, y, w, h, *(QPalette*)pal, sunken, lineWidth, midLineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawShadeRect2(void* p, void* r, void* pal, int sunken, int lineWidth, int midLineWidth, void* fill) {
	qDrawShadeRect((QPainter*)p, *(QRect*)r, *(QPalette*)pal, sunken, lineWidth, midLineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawShadePanel(void* p, int x, int y, int w, int h, void* pal, int sunken, int lineWidth, void* fill) {
	qDrawShadePanel((QPainter*)p, x, y, w, h, *(QPalette*)pal, sunken, lineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawShadePanel2(void* p, void* r, void* pal, int sunken, int lineWidth, void* fill) {
	qDrawShadePanel((QPainter*)p, *(QRect*)r, *(QPalette*)pal, sunken, lineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawWinButton(void* p, int x, int y, int w, int h, void* pal, int sunken, void* fill) {
	qDrawWinButton((QPainter*)p, x, y, w, h, *(QPalette*)pal, sunken, (QBrush*)fill);
}
void cppgofn_qt_qDrawWinButton2(void* p, void* r, void* pal, int sunken, void* fill) {
	qDrawWinButton((QPainter*)p, *(QRect*)r, *(QPalette*)pal, sunken, (QBrush*)fill);
}
void cppgofn_qt_qDrawWinPanel(void* p, int x, int y, int w, int h, void* pal, int sunken, void* fill) {
	qDrawWinPanel((QPainter*)p, x, y, w, h, *(QPalette*)pal, sunken, (QBrush*)fill);
}
void cppgofn_qt_qDrawWinPanel2(void* p, void* r, void* pal, int sunken, void* fill) {
	qDrawWinPanel((QPainter*)p, *(QRect*)r, *(QPalette*)pal, sunken, (QBrush*)fill);
}
void cppgofn_qt_qDrawPlainRect(void* p, int x, int y, int w, int h, void* arg6, int lineWidth, void* fill) {
	qDrawPlainRect((QPainter*)p, x, y, w, h, *(QColor*)arg6, lineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawPlainRect2(void* p, void* r, void* arg3, int lineWidth, void* fill) {
	qDrawPlainRect((QPainter*)p, *(QRect*)r, *(QColor*)arg3, lineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawPlainRoundedRect(void* p, int x, int y, int w, int h, double rx, double ry, void* arg8, int lineWidth, void* fill) {
	qDrawPlainRoundedRect((QPainter*)p, x, y, w, h, rx, ry, *(QColor*)arg8, lineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawPlainRoundedRect2(void* painter, void* rect, double rx, double ry, void* lineColor, int lineWidth, void* fill) {
	qDrawPlainRoundedRect((QPainter*)painter, *(QRect*)rect, rx, ry, *(QColor*)lineColor, lineWidth, (QBrush*)fill);
}
void cppgofn_qt_qDrawBorderPixmap(void* painter, void* targetRect, void* targetMargins, void* pixmap, void* sourceRect, void* sourceMargins, void* rules, int hints) {
	qDrawBorderPixmap((QPainter*)painter, *(QRect*)targetRect, *(QMargins*)targetMargins, *(QPixmap*)pixmap, *(QRect*)sourceRect, *(QMargins*)sourceMargins, *(QTileRules*)rules, QFlags<QDrawBorderPixmap::DrawingHint>(hints));
}
void cppgofn_qt_qDrawBorderPixmap2(void* painter, void* target, void* margins, void* pixmap) {
	qDrawBorderPixmap((QPainter*)painter, *(QRect*)target, *(QMargins*)margins, *(QPixmap*)pixmap);
}

}
