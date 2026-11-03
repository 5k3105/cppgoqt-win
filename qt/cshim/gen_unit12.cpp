#include <QtCore/qobjectdefs.h>
#include "gen_qgenericreturnargument.h"

extern "C" {

void* QGenericReturnArgument_base(void* self) { return (void*)static_cast<QGenericArgument*>((QGenericReturnArgument*)self); }
void* QGenericReturnArgument_new() {
	return new QGenericReturnArgument();
}
void* QGenericReturnArgument_new2(const char* aName) {
	return new QGenericReturnArgument(aName);
}
void QGenericReturnArgument_delete(void* self) { delete (QGenericReturnArgument*)self; }

}

#include <QtCore/qstringmatcher.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include "gen_qstringmatcher.h"

extern "C" {

void* QStringMatcher_new() {
	return new QStringMatcher();
}
void* QStringMatcher_new2(const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	return new QStringMatcher(pattern_qstr);
}
void* QStringMatcher_new3(void* uc, long long len_) {
	return new QStringMatcher((QChar*)uc, len_);
}
void* QStringMatcher_new4(void* pattern) {
	return new QStringMatcher(*(QStringView*)pattern);
}
void* QStringMatcher_new5(void* other) {
	return new QStringMatcher(*(QStringMatcher*)other);
}
void* QStringMatcher_new6(const char* pattern, long long pattern_len, int cs) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	return new QStringMatcher(pattern_qstr, (Qt::CaseSensitivity)(cs));
}
void* QStringMatcher_new7(void* uc, long long len_, int cs) {
	return new QStringMatcher((QChar*)uc, len_, (Qt::CaseSensitivity)(cs));
}
void* QStringMatcher_new8(void* pattern, int cs) {
	return new QStringMatcher(*(QStringView*)pattern, (Qt::CaseSensitivity)(cs));
}
void* QStringMatcher_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QStringMatcher*)self)->operator=(*(QStringMatcher*)other);
	return (void*)&(__r);
}
void QStringMatcher_setPattern(void* self, const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	((QStringMatcher*)self)->setPattern(pattern_qstr);
}
void QStringMatcher_setCaseSensitivity(void* self, int cs) {
	((QStringMatcher*)self)->setCaseSensitivity((Qt::CaseSensitivity)(cs));
}
long long QStringMatcher_indexIn(void* self, const char* str, long long str_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QStringMatcher*)self)->indexIn(str_qstr);
	return __r;
}
long long QStringMatcher_indexIn2(void* self, void* str, long long length) {
	auto&& __r = ((QStringMatcher*)self)->indexIn((QChar*)str, length);
	return __r;
}
long long QStringMatcher_indexIn3(void* self, void* str) {
	auto&& __r = ((QStringMatcher*)self)->indexIn(*(QStringView*)str);
	return __r;
}
char* QStringMatcher_pattern(void* self, long long* __out_len) {
	auto&& __r = ((QStringMatcher*)self)->pattern();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStringMatcher_patternView(void* self) {
	auto&& __r = ((QStringMatcher*)self)->patternView();
	return new QStringView(__r);
}
int QStringMatcher_caseSensitivity(void* self) {
	auto&& __r = ((QStringMatcher*)self)->caseSensitivity();
	return (int)(__r);
}
long long QStringMatcher_indexIn4(void* self, const char* str, long long str_len, long long from) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QStringMatcher*)self)->indexIn(str_qstr, from);
	return __r;
}
long long QStringMatcher_indexIn5(void* self, void* str, long long length, long long from) {
	auto&& __r = ((QStringMatcher*)self)->indexIn((QChar*)str, length, from);
	return __r;
}
long long QStringMatcher_indexIn6(void* self, void* str, long long from) {
	auto&& __r = ((QStringMatcher*)self)->indexIn(*(QStringView*)str, from);
	return __r;
}
void QStringMatcher_delete(void* self) { delete (QStringMatcher*)self; }

}

#include <QtCore/qtextstream.h>
#include <QtCore/qchar.h>
#include "gen_qtextstreammanipulator.h"

extern "C" {

void QTextStreamManipulator_exec(void* self, void* s) {
	((QTextStreamManipulator*)self)->exec(*(QTextStream*)s);
}
void QTextStreamManipulator_delete(void* self) { delete (QTextStreamManipulator*)self; }

}

#include <QtCore/qcoreevent.h>
#include "gen_qevent.h"

extern "C" {

void* QEvent_new(int type_) {
	return new QEvent((QEvent::Type)(type_));
}
int QEvent_type(void* self) {
	auto&& __r = ((QEvent*)self)->type();
	return (int)(__r);
}
int QEvent_spontaneous(void* self) {
	auto&& __r = ((QEvent*)self)->spontaneous();
	return __r;
}
void QEvent_setAccepted(void* self, int accepted) {
	((QEvent*)self)->setAccepted(accepted);
}
int QEvent_isAccepted(void* self) {
	auto&& __r = ((QEvent*)self)->isAccepted();
	return __r;
}
void QEvent_accept(void* self) {
	((QEvent*)self)->accept();
}
void QEvent_ignore(void* self) {
	((QEvent*)self)->ignore();
}
int QEvent_isInputEvent(void* self) {
	auto&& __r = ((QEvent*)self)->isInputEvent();
	return __r;
}
int QEvent_isPointerEvent(void* self) {
	auto&& __r = ((QEvent*)self)->isPointerEvent();
	return __r;
}
int QEvent_isSinglePointEvent(void* self) {
	auto&& __r = ((QEvent*)self)->isSinglePointEvent();
	return __r;
}
int QEvent_registerEventType() {
	auto&& __r = QEvent::registerEventType();
	return __r;
}
void* QEvent_clone(void* self) {
	auto&& __r = ((QEvent*)self)->clone();
	return (void*)(__r);
}
int QEvent_registerEventType2(int hint) {
	auto&& __r = QEvent::registerEventType(hint);
	return __r;
}
void QEvent_delete(void* self) { delete (QEvent*)self; }

}

#include <QtCore/qurl.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qurlquery.h>
#include "gen_qurl.h"

extern "C" {

void* QUrl_new() {
	return new QUrl();
}
void* QUrl_new2(void* copy_) {
	return new QUrl(*(QUrl*)copy_);
}
void* QUrl_new3(const char* url, long long url_len) {
	QString url_qstr = QString::fromUtf8(url, url_len);
	return new QUrl(url_qstr);
}
void* QUrl_new4(const char* url, long long url_len, int mode) {
	QString url_qstr = QString::fromUtf8(url, url_len);
	return new QUrl(url_qstr, (QUrl::ParsingMode)(mode));
}
void* QUrl_OperatorAssign(void* self, void* copy_) {
	auto&& __r = ((QUrl*)self)->operator=(*(QUrl*)copy_);
	return (void*)&(__r);
}
void* QUrl_OperatorAssign2(void* self, const char* url, long long url_len) {
	QString url_qstr = QString::fromUtf8(url, url_len);
	auto&& __r = ((QUrl*)self)->operator=(url_qstr);
	return (void*)&(__r);
}
void QUrl_swap(void* self, void* other) {
	((QUrl*)self)->swap(*(QUrl*)other);
}
void QUrl_setUrl(void* self, const char* url, long long url_len) {
	QString url_qstr = QString::fromUtf8(url, url_len);
	((QUrl*)self)->setUrl(url_qstr);
}
char* QUrl_url(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->url();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrl_toString(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrl_toDisplayString(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->toDisplayString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUrl_toEncoded(void* self) {
	auto&& __r = ((QUrl*)self)->toEncoded();
	return new QByteArray(__r);
}
void* QUrl_fromEncoded(void* input) {
	auto&& __r = QUrl::fromEncoded(*(QByteArrayView*)input);
	return new QUrl(__r);
}
void* QUrl_fromUserInput(const char* userInput, long long userInput_len) {
	QString userInput_qstr = QString::fromUtf8(userInput, userInput_len);
	auto&& __r = QUrl::fromUserInput(userInput_qstr);
	return new QUrl(__r);
}
int QUrl_isValid(void* self) {
	auto&& __r = ((QUrl*)self)->isValid();
	return __r;
}
char* QUrl_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QUrl_isEmpty(void* self) {
	auto&& __r = ((QUrl*)self)->isEmpty();
	return __r;
}
void QUrl_clear(void* self) {
	((QUrl*)self)->clear();
}
void QUrl_setScheme(void* self, const char* scheme, long long scheme_len) {
	QString scheme_qstr = QString::fromUtf8(scheme, scheme_len);
	((QUrl*)self)->setScheme(scheme_qstr);
}
char* QUrl_scheme(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->scheme();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setAuthority(void* self, const char* authority, long long authority_len) {
	QString authority_qstr = QString::fromUtf8(authority, authority_len);
	((QUrl*)self)->setAuthority(authority_qstr);
}
char* QUrl_authority(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->authority();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setUserInfo(void* self, const char* userInfo, long long userInfo_len) {
	QString userInfo_qstr = QString::fromUtf8(userInfo, userInfo_len);
	((QUrl*)self)->setUserInfo(userInfo_qstr);
}
char* QUrl_userInfo(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->userInfo();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setUserName(void* self, const char* userName, long long userName_len) {
	QString userName_qstr = QString::fromUtf8(userName, userName_len);
	((QUrl*)self)->setUserName(userName_qstr);
}
char* QUrl_userName(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->userName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setPassword(void* self, const char* password, long long password_len) {
	QString password_qstr = QString::fromUtf8(password, password_len);
	((QUrl*)self)->setPassword(password_qstr);
}
char* QUrl_password(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->password();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setHost(void* self, const char* host, long long host_len) {
	QString host_qstr = QString::fromUtf8(host, host_len);
	((QUrl*)self)->setHost(host_qstr);
}
char* QUrl_host(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->host();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setPort(void* self, int port) {
	((QUrl*)self)->setPort(port);
}
int QUrl_port(void* self) {
	auto&& __r = ((QUrl*)self)->port();
	return __r;
}
void QUrl_setPath(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	((QUrl*)self)->setPath(path_qstr);
}
char* QUrl_path(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->path();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrl_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QUrl_hasQuery(void* self) {
	auto&& __r = ((QUrl*)self)->hasQuery();
	return __r;
}
void QUrl_setQuery(void* self, const char* query, long long query_len) {
	QString query_qstr = QString::fromUtf8(query, query_len);
	((QUrl*)self)->setQuery(query_qstr);
}
void QUrl_setQuery2(void* self, void* query) {
	((QUrl*)self)->setQuery(*(QUrlQuery*)query);
}
char* QUrl_query(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->query();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QUrl_hasFragment(void* self) {
	auto&& __r = ((QUrl*)self)->hasFragment();
	return __r;
}
char* QUrl_fragment(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->fragment();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setFragment(void* self, const char* fragment, long long fragment_len) {
	QString fragment_qstr = QString::fromUtf8(fragment, fragment_len);
	((QUrl*)self)->setFragment(fragment_qstr);
}
void* QUrl_resolved(void* self, void* relative) {
	auto&& __r = ((QUrl*)self)->resolved(*(QUrl*)relative);
	return new QUrl(__r);
}
int QUrl_isRelative(void* self) {
	auto&& __r = ((QUrl*)self)->isRelative();
	return __r;
}
int QUrl_isParentOf(void* self, void* url) {
	auto&& __r = ((QUrl*)self)->isParentOf(*(QUrl*)url);
	return __r;
}
int QUrl_isLocalFile(void* self) {
	auto&& __r = ((QUrl*)self)->isLocalFile();
	return __r;
}
void* QUrl_fromLocalFile(const char* localfile, long long localfile_len) {
	QString localfile_qstr = QString::fromUtf8(localfile, localfile_len);
	auto&& __r = QUrl::fromLocalFile(localfile_qstr);
	return new QUrl(__r);
}
char* QUrl_toLocalFile(void* self, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->toLocalFile();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_detach(void* self) {
	((QUrl*)self)->detach();
}
int QUrl_isDetached(void* self) {
	auto&& __r = ((QUrl*)self)->isDetached();
	return __r;
}
char* QUrl_fromPercentEncoding(void* arg1, long long* __out_len) {
	auto&& __r = QUrl::fromPercentEncoding(*(QByteArray*)arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUrl_toPercentEncoding(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QUrl::toPercentEncoding(arg1_qstr);
	return new QByteArray(__r);
}
char* QUrl_fromAce(void* domain, long long* __out_len) {
	auto&& __r = QUrl::fromAce(*(QByteArray*)domain);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUrl_toAce(const char* domain, long long domain_len) {
	QString domain_qstr = QString::fromUtf8(domain, domain_len);
	auto&& __r = QUrl::toAce(domain_qstr);
	return new QByteArray(__r);
}
char** QUrl_idnWhitelist(long long* __out_len) {
	auto&& __r = QUrl::idnWhitelist();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QUrl_toStringList(void** uris, long long uris_len, long long* __out_len) {
	QList<QUrl> uris_list; uris_list.reserve(uris_len);
	for (long long __i = 0; __i < uris_len; __i++) uris_list.append(*(QUrl*)uris[__i]);
	auto&& __r = QUrl::toStringList(uris_list);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QUrl_fromStringList(char** uris, long long uris_len, long long* __out_len) {
	QList<QString> uris_list; uris_list.reserve(uris_len);
	for (long long __i = 0; __i < uris_len; __i++) uris_list.append(QString::fromUtf8(uris[__i]));
	auto&& __r = QUrl::fromStringList(uris_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void QUrl_setIdnWhitelist(char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	QUrl::setIdnWhitelist(arg1_list);
}
void QUrl_setUrl2(void* self, const char* url, long long url_len, int mode) {
	QString url_qstr = QString::fromUtf8(url, url_len);
	((QUrl*)self)->setUrl(url_qstr, (QUrl::ParsingMode)(mode));
}
void* QUrl_fromEncoded2(void* input, int mode) {
	auto&& __r = QUrl::fromEncoded(*(QByteArrayView*)input, (QUrl::ParsingMode)(mode));
	return new QUrl(__r);
}
void* QUrl_fromUserInput2(const char* userInput, long long userInput_len, const char* workingDirectory, long long workingDirectory_len) {
	QString userInput_qstr = QString::fromUtf8(userInput, userInput_len);
	QString workingDirectory_qstr = QString::fromUtf8(workingDirectory, workingDirectory_len);
	auto&& __r = QUrl::fromUserInput(userInput_qstr, workingDirectory_qstr);
	return new QUrl(__r);
}
void* QUrl_fromUserInput3(const char* userInput, long long userInput_len, const char* workingDirectory, long long workingDirectory_len, int options) {
	QString userInput_qstr = QString::fromUtf8(userInput, userInput_len);
	QString workingDirectory_qstr = QString::fromUtf8(workingDirectory, workingDirectory_len);
	auto&& __r = QUrl::fromUserInput(userInput_qstr, workingDirectory_qstr, QFlags<QUrl::UserInputResolutionOption>(options));
	return new QUrl(__r);
}
void QUrl_setAuthority2(void* self, const char* authority, long long authority_len, int mode) {
	QString authority_qstr = QString::fromUtf8(authority, authority_len);
	((QUrl*)self)->setAuthority(authority_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_authority2(void* self, int options, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->authority(QFlags<QUrl::ComponentFormattingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setUserInfo2(void* self, const char* userInfo, long long userInfo_len, int mode) {
	QString userInfo_qstr = QString::fromUtf8(userInfo, userInfo_len);
	((QUrl*)self)->setUserInfo(userInfo_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_userInfo2(void* self, int options, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->userInfo(QFlags<QUrl::ComponentFormattingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setUserName2(void* self, const char* userName, long long userName_len, int mode) {
	QString userName_qstr = QString::fromUtf8(userName, userName_len);
	((QUrl*)self)->setUserName(userName_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_userName2(void* self, int options, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->userName(QFlags<QUrl::ComponentFormattingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setPassword2(void* self, const char* password, long long password_len, int mode) {
	QString password_qstr = QString::fromUtf8(password, password_len);
	((QUrl*)self)->setPassword(password_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_password2(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->password(QFlags<QUrl::ComponentFormattingOption>(arg1));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setHost2(void* self, const char* host, long long host_len, int mode) {
	QString host_qstr = QString::fromUtf8(host, host_len);
	((QUrl*)self)->setHost(host_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_host2(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->host(QFlags<QUrl::ComponentFormattingOption>(arg1));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QUrl_port2(void* self, int defaultPort) {
	auto&& __r = ((QUrl*)self)->port(defaultPort);
	return __r;
}
void QUrl_setPath2(void* self, const char* path, long long path_len, int mode) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	((QUrl*)self)->setPath(path_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_path2(void* self, int options, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->path(QFlags<QUrl::ComponentFormattingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrl_fileName2(void* self, int options, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->fileName(QFlags<QUrl::ComponentFormattingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setQuery3(void* self, const char* query, long long query_len, int mode) {
	QString query_qstr = QString::fromUtf8(query, query_len);
	((QUrl*)self)->setQuery(query_qstr, (QUrl::ParsingMode)(mode));
}
char* QUrl_query2(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->query(QFlags<QUrl::ComponentFormattingOption>(arg1));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUrl_fragment2(void* self, int options, long long* __out_len) {
	auto&& __r = ((QUrl*)self)->fragment(QFlags<QUrl::ComponentFormattingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUrl_setFragment2(void* self, const char* fragment, long long fragment_len, int mode) {
	QString fragment_qstr = QString::fromUtf8(fragment, fragment_len);
	((QUrl*)self)->setFragment(fragment_qstr, (QUrl::ParsingMode)(mode));
}
void* QUrl_toPercentEncoding2(const char* arg1, long long arg1_len, void* exclude) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QUrl::toPercentEncoding(arg1_qstr, *(QByteArray*)exclude);
	return new QByteArray(__r);
}
void* QUrl_toPercentEncoding3(const char* arg1, long long arg1_len, void* exclude, void* include) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QUrl::toPercentEncoding(arg1_qstr, *(QByteArray*)exclude, *(QByteArray*)include);
	return new QByteArray(__r);
}
char* QUrl_fromAce2(void* domain, int options, long long* __out_len) {
	auto&& __r = QUrl::fromAce(*(QByteArray*)domain, QFlags<QUrl::AceProcessingOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUrl_toAce2(const char* domain, long long domain_len, int options) {
	QString domain_qstr = QString::fromUtf8(domain, domain_len);
	auto&& __r = QUrl::toAce(domain_qstr, QFlags<QUrl::AceProcessingOption>(options));
	return new QByteArray(__r);
}
void** QUrl_fromStringList2(char** uris, long long uris_len, int mode, long long* __out_len) {
	QList<QString> uris_list; uris_list.reserve(uris_len);
	for (long long __i = 0; __i < uris_len; __i++) uris_list.append(QString::fromUtf8(uris[__i]));
	auto&& __r = QUrl::fromStringList(uris_list, (QUrl::ParsingMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
int QUrl_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QUrl*)self) == (*(QUrl*)rhs);
	return __r;
}
int QUrl_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QUrl*)self) != (*(QUrl*)rhs);
	return __r;
}
int QUrl_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QUrl*)self) < (*(QUrl*)rhs);
	return __r;
}
int QUrl_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QUrl*)self) > (*(QUrl*)rhs);
	return __r;
}
int QUrl_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QUrl*)self) <= (*(QUrl*)rhs);
	return __r;
}
int QUrl_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QUrl*)self) >= (*(QUrl*)rhs);
	return __r;
}
void QUrl_delete(void* self) { delete (QUrl*)self; }

}

#include <QtCore/qcollator.h>
#include "gen_qcollatorsortkey.h"

extern "C" {

void* QCollatorSortKey_new(void* other) {
	return new QCollatorSortKey(*(QCollatorSortKey*)other);
}
void* QCollatorSortKey_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCollatorSortKey*)self)->operator=(*(QCollatorSortKey*)other);
	return (void*)&(__r);
}
void QCollatorSortKey_swap(void* self, void* other) {
	((QCollatorSortKey*)self)->swap(*(QCollatorSortKey*)other);
}
int QCollatorSortKey_compare(void* self, void* key) {
	auto&& __r = ((QCollatorSortKey*)self)->compare(*(QCollatorSortKey*)key);
	return __r;
}
int QCollatorSortKey_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QCollatorSortKey*)self) < (*(QCollatorSortKey*)rhs);
	return __r;
}
void QCollatorSortKey_delete(void* self) { delete (QCollatorSortKey*)self; }

}

#include <QtCore/qfutureinterface.h>
#include <QtCore/qexception.h>
#include <QtCore/qrunnable.h>
#include <QtCore/qthreadpool.h>
#include <QtCore/qstring.h>
#include <QtCore/qmutex.h>
#include <QtCore/qobject.h>
#include <QtCore/qvariant.h>
#include "gen_qfutureinterfacebase.h"

extern "C" {

void* QFutureInterfaceBase_new() {
	return new QFutureInterfaceBase();
}
void* QFutureInterfaceBase_new2(void* other) {
	return new QFutureInterfaceBase(*(QFutureInterfaceBase*)other);
}
void* QFutureInterfaceBase_new3(int initialState) {
	return new QFutureInterfaceBase((QFutureInterfaceBase::State)(initialState));
}
void* QFutureInterfaceBase_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QFutureInterfaceBase*)self)->operator=(*(QFutureInterfaceBase*)other);
	return (void*)&(__r);
}
void QFutureInterfaceBase_reportStarted(void* self) {
	((QFutureInterfaceBase*)self)->reportStarted();
}
void QFutureInterfaceBase_reportFinished(void* self) {
	((QFutureInterfaceBase*)self)->reportFinished();
}
void QFutureInterfaceBase_reportCanceled(void* self) {
	((QFutureInterfaceBase*)self)->reportCanceled();
}
void QFutureInterfaceBase_reportException(void* self, void* e) {
	((QFutureInterfaceBase*)self)->reportException(*(QException*)e);
}
void QFutureInterfaceBase_reportResultsReady(void* self, int beginIndex, int endIndex) {
	((QFutureInterfaceBase*)self)->reportResultsReady(beginIndex, endIndex);
}
void QFutureInterfaceBase_setRunnable(void* self, void* runnable) {
	((QFutureInterfaceBase*)self)->setRunnable((QRunnable*)runnable);
}
void QFutureInterfaceBase_setThreadPool(void* self, void* pool) {
	((QFutureInterfaceBase*)self)->setThreadPool((QThreadPool*)pool);
}
void* QFutureInterfaceBase_threadPool(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->threadPool();
	return (void*)(__r);
}
void QFutureInterfaceBase_setFilterMode(void* self, int enable) {
	((QFutureInterfaceBase*)self)->setFilterMode(enable);
}
void QFutureInterfaceBase_setProgressRange(void* self, int minimum, int maximum) {
	((QFutureInterfaceBase*)self)->setProgressRange(minimum, maximum);
}
int QFutureInterfaceBase_progressMinimum(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->progressMinimum();
	return __r;
}
int QFutureInterfaceBase_progressMaximum(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->progressMaximum();
	return __r;
}
int QFutureInterfaceBase_isProgressUpdateNeeded(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isProgressUpdateNeeded();
	return __r;
}
void QFutureInterfaceBase_setProgressValue(void* self, int progressValue) {
	((QFutureInterfaceBase*)self)->setProgressValue(progressValue);
}
int QFutureInterfaceBase_progressValue(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->progressValue();
	return __r;
}
void QFutureInterfaceBase_setProgressValueAndText(void* self, int progressValue, const char* progressText, long long progressText_len) {
	QString progressText_qstr = QString::fromUtf8(progressText, progressText_len);
	((QFutureInterfaceBase*)self)->setProgressValueAndText(progressValue, progressText_qstr);
}
char* QFutureInterfaceBase_progressText(void* self, long long* __out_len) {
	auto&& __r = ((QFutureInterfaceBase*)self)->progressText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFutureInterfaceBase_setExpectedResultCount(void* self, int resultCount) {
	((QFutureInterfaceBase*)self)->setExpectedResultCount(resultCount);
}
int QFutureInterfaceBase_expectedResultCount(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->expectedResultCount();
	return __r;
}
int QFutureInterfaceBase_resultCount(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->resultCount();
	return __r;
}
int QFutureInterfaceBase_queryState(void* self, int state) {
	auto&& __r = ((QFutureInterfaceBase*)self)->queryState((QFutureInterfaceBase::State)(state));
	return __r;
}
int QFutureInterfaceBase_isRunning(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isRunning();
	return __r;
}
int QFutureInterfaceBase_isStarted(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isStarted();
	return __r;
}
int QFutureInterfaceBase_isCanceled(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isCanceled();
	return __r;
}
int QFutureInterfaceBase_isFinished(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isFinished();
	return __r;
}
int QFutureInterfaceBase_isPaused(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isPaused();
	return __r;
}
void QFutureInterfaceBase_setPaused(void* self, int paused) {
	((QFutureInterfaceBase*)self)->setPaused(paused);
}
void QFutureInterfaceBase_togglePaused(void* self) {
	((QFutureInterfaceBase*)self)->togglePaused();
}
int QFutureInterfaceBase_isSuspending(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isSuspending();
	return __r;
}
int QFutureInterfaceBase_isSuspended(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isSuspended();
	return __r;
}
int QFutureInterfaceBase_isThrottled(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isThrottled();
	return __r;
}
int QFutureInterfaceBase_isResultReadyAt(void* self, int index) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isResultReadyAt(index);
	return __r;
}
int QFutureInterfaceBase_isValid(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isValid();
	return __r;
}
int QFutureInterfaceBase_loadState(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->loadState();
	return __r;
}
void QFutureInterfaceBase_cancel(void* self) {
	((QFutureInterfaceBase*)self)->cancel();
}
void QFutureInterfaceBase_cancelAndFinish(void* self) {
	((QFutureInterfaceBase*)self)->cancelAndFinish();
}
void QFutureInterfaceBase_cancelChain(void* self) {
	((QFutureInterfaceBase*)self)->cancelChain();
}
void QFutureInterfaceBase_setAddResultsIfCanceledEnabled(void* self, int enable) {
	((QFutureInterfaceBase*)self)->setAddResultsIfCanceledEnabled(enable);
}
int QFutureInterfaceBase_isAddResultsIfCanceledEnabled(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isAddResultsIfCanceledEnabled();
	return __r;
}
void QFutureInterfaceBase_setSuspended(void* self, int suspend) {
	((QFutureInterfaceBase*)self)->setSuspended(suspend);
}
void QFutureInterfaceBase_toggleSuspended(void* self) {
	((QFutureInterfaceBase*)self)->toggleSuspended();
}
void QFutureInterfaceBase_reportSuspended(void* self) {
	((QFutureInterfaceBase*)self)->reportSuspended();
}
void QFutureInterfaceBase_setThrottled(void* self, int enable) {
	((QFutureInterfaceBase*)self)->setThrottled(enable);
}
void QFutureInterfaceBase_waitForFinished(void* self) {
	((QFutureInterfaceBase*)self)->waitForFinished();
}
int QFutureInterfaceBase_waitForNextResult(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->waitForNextResult();
	return __r;
}
void QFutureInterfaceBase_waitForResult(void* self, int resultIndex) {
	((QFutureInterfaceBase*)self)->waitForResult(resultIndex);
}
void QFutureInterfaceBase_waitForResume(void* self) {
	((QFutureInterfaceBase*)self)->waitForResume();
}
void QFutureInterfaceBase_suspendIfRequested(void* self) {
	((QFutureInterfaceBase*)self)->suspendIfRequested();
}
void* QFutureInterfaceBase_mutex(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->mutex();
	return (void*)&(__r);
}
int QFutureInterfaceBase_hasException(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->hasException();
	return __r;
}
int QFutureInterfaceBase_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QFutureInterfaceBase*)self)->operator==(*(QFutureInterfaceBase*)other);
	return __r;
}
int QFutureInterfaceBase_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QFutureInterfaceBase*)self)->operator!=(*(QFutureInterfaceBase*)other);
	return __r;
}
void QFutureInterfaceBase_swap(void* self, void* other) {
	((QFutureInterfaceBase*)self)->swap(*(QFutureInterfaceBase*)other);
}
int QFutureInterfaceBase_isChainCanceled(void* self) {
	auto&& __r = ((QFutureInterfaceBase*)self)->isChainCanceled();
	return __r;
}
void QFutureInterfaceBase_delete(void* self) { delete (QFutureInterfaceBase*)self; }

}

#include <QtCore/qline.h>
#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include "gen_qlinef.h"

extern "C" {

void* QLineF_new() {
	return new QLineF();
}
void* QLineF_new2(void* pt1, void* pt2) {
	return new QLineF(*(QPointF*)pt1, *(QPointF*)pt2);
}
void* QLineF_new3(double x1, double y1, double x2, double y2) {
	return new QLineF(x1, y1, x2, y2);
}
void* QLineF_new4(void* line) {
	return new QLineF(*(QLine*)line);
}
void* QLineF_new5(void* arg1) {
	return new QLineF(*(QLineF*)arg1);
}
void* QLineF_fromPolar(double length, double angle) {
	auto&& __r = QLineF::fromPolar(length, angle);
	return new QLineF(__r);
}
int QLineF_isNull(void* self) {
	auto&& __r = ((QLineF*)self)->isNull();
	return __r;
}
void* QLineF_p1(void* self) {
	auto&& __r = ((QLineF*)self)->p1();
	return new QPointF(__r);
}
void* QLineF_p2(void* self) {
	auto&& __r = ((QLineF*)self)->p2();
	return new QPointF(__r);
}
double QLineF_x1(void* self) {
	auto&& __r = ((QLineF*)self)->x1();
	return __r;
}
double QLineF_y1(void* self) {
	auto&& __r = ((QLineF*)self)->y1();
	return __r;
}
double QLineF_x2(void* self) {
	auto&& __r = ((QLineF*)self)->x2();
	return __r;
}
double QLineF_y2(void* self) {
	auto&& __r = ((QLineF*)self)->y2();
	return __r;
}
double QLineF_dx(void* self) {
	auto&& __r = ((QLineF*)self)->dx();
	return __r;
}
double QLineF_dy(void* self) {
	auto&& __r = ((QLineF*)self)->dy();
	return __r;
}
double QLineF_length(void* self) {
	auto&& __r = ((QLineF*)self)->length();
	return __r;
}
void QLineF_setLength(void* self, double len_) {
	((QLineF*)self)->setLength(len_);
}
double QLineF_angle(void* self) {
	auto&& __r = ((QLineF*)self)->angle();
	return __r;
}
void QLineF_setAngle(void* self, double angle) {
	((QLineF*)self)->setAngle(angle);
}
double QLineF_angleTo(void* self, void* l) {
	auto&& __r = ((QLineF*)self)->angleTo(*(QLineF*)l);
	return __r;
}
void* QLineF_unitVector(void* self) {
	auto&& __r = ((QLineF*)self)->unitVector();
	return new QLineF(__r);
}
void* QLineF_normalVector(void* self) {
	auto&& __r = ((QLineF*)self)->normalVector();
	return new QLineF(__r);
}
int QLineF_intersects(void* self, void* l) {
	auto&& __r = ((QLineF*)self)->intersects(*(QLineF*)l);
	return (int)(__r);
}
void* QLineF_pointAt(void* self, double t) {
	auto&& __r = ((QLineF*)self)->pointAt(t);
	return new QPointF(__r);
}
void QLineF_translate(void* self, void* p) {
	((QLineF*)self)->translate(*(QPointF*)p);
}
void QLineF_translate2(void* self, double dx, double dy) {
	((QLineF*)self)->translate(dx, dy);
}
void* QLineF_translated(void* self, void* p) {
	auto&& __r = ((QLineF*)self)->translated(*(QPointF*)p);
	return new QLineF(__r);
}
void* QLineF_translated2(void* self, double dx, double dy) {
	auto&& __r = ((QLineF*)self)->translated(dx, dy);
	return new QLineF(__r);
}
void* QLineF_center(void* self) {
	auto&& __r = ((QLineF*)self)->center();
	return new QPointF(__r);
}
void QLineF_setP1(void* self, void* p1) {
	((QLineF*)self)->setP1(*(QPointF*)p1);
}
void QLineF_setP2(void* self, void* p2) {
	((QLineF*)self)->setP2(*(QPointF*)p2);
}
void QLineF_setPoints(void* self, void* p1, void* p2) {
	((QLineF*)self)->setPoints(*(QPointF*)p1, *(QPointF*)p2);
}
void QLineF_setLine(void* self, double x1, double y1, double x2, double y2) {
	((QLineF*)self)->setLine(x1, y1, x2, y2);
}
void* QLineF_toLine(void* self) {
	auto&& __r = ((QLineF*)self)->toLine();
	return new QLine(__r);
}
int QLineF_intersects2(void* self, void* l, void* intersectionPoint) {
	auto&& __r = ((QLineF*)self)->intersects(*(QLineF*)l, (QPointF*)intersectionPoint);
	return (int)(__r);
}
int QLineF_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QLineF*)self) == (*(QLineF*)rhs);
	return __r;
}
int QLineF_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QLineF*)self) != (*(QLineF*)rhs);
	return __r;
}
int QLineF_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QLineF*)self) == (*(QLine*)rhs);
	return __r;
}
int QLineF_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QLineF*)self) != (*(QLine*)rhs);
	return __r;
}
void* QLineF_OperatorMultiply(void* self, void* m) {
	auto&& __r = (*(QLineF*)self) * (*(QTransform*)m);
	return new QLineF(__r);
}
void QLineF_delete(void* self) { delete (QLineF*)self; }

}

#include <QtCore/qpermissions.h>
#include "gen_qcamerapermission.h"

extern "C" {

void* QCameraPermission_new() {
	return new QCameraPermission();
}
void* QCameraPermission_new2(void* other) {
	return new QCameraPermission(*(QCameraPermission*)other);
}
void* QCameraPermission_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCameraPermission*)self)->operator=(*(QCameraPermission*)other);
	return (void*)&(__r);
}
void QCameraPermission_swap(void* self, void* other) {
	((QCameraPermission*)self)->swap(*(QCameraPermission*)other);
}
void QCameraPermission_delete(void* self) { delete (QCameraPermission*)self; }

}

#include <QtCore/qresource.h>
#include <QtCore/qstring.h>
#include <QtCore/qlocale.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatetime.h>
#include "gen_qresource.h"

extern "C" {

void* QResource_new() {
	return new QResource();
}
void* QResource_new2(const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	return new QResource(file_qstr);
}
void* QResource_new3(const char* file, long long file_len, void* locale) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	return new QResource(file_qstr, *(QLocale*)locale);
}
void QResource_setFileName(void* self, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	((QResource*)self)->setFileName(file_qstr);
}
char* QResource_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QResource*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QResource_absoluteFilePath(void* self, long long* __out_len) {
	auto&& __r = ((QResource*)self)->absoluteFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QResource_setLocale(void* self, void* locale) {
	((QResource*)self)->setLocale(*(QLocale*)locale);
}
void* QResource_locale(void* self) {
	auto&& __r = ((QResource*)self)->locale();
	return new QLocale(__r);
}
int QResource_isValid(void* self) {
	auto&& __r = ((QResource*)self)->isValid();
	return __r;
}
int QResource_compressionAlgorithm(void* self) {
	auto&& __r = ((QResource*)self)->compressionAlgorithm();
	return (int)(__r);
}
long long QResource_size(void* self) {
	auto&& __r = ((QResource*)self)->size();
	return __r;
}
long long QResource_uncompressedSize(void* self) {
	auto&& __r = ((QResource*)self)->uncompressedSize();
	return __r;
}
void* QResource_uncompressedData(void* self) {
	auto&& __r = ((QResource*)self)->uncompressedData();
	return new QByteArray(__r);
}
void* QResource_lastModified(void* self) {
	auto&& __r = ((QResource*)self)->lastModified();
	return new QDateTime(__r);
}
int QResource_registerResource(const char* rccFilename, long long rccFilename_len) {
	QString rccFilename_qstr = QString::fromUtf8(rccFilename, rccFilename_len);
	auto&& __r = QResource::registerResource(rccFilename_qstr);
	return __r;
}
int QResource_unregisterResource(const char* rccFilename, long long rccFilename_len) {
	QString rccFilename_qstr = QString::fromUtf8(rccFilename, rccFilename_len);
	auto&& __r = QResource::unregisterResource(rccFilename_qstr);
	return __r;
}
int QResource_registerResource2(const unsigned char* rccData) {
	auto&& __r = QResource::registerResource(rccData);
	return __r;
}
int QResource_unregisterResource2(const unsigned char* rccData) {
	auto&& __r = QResource::unregisterResource(rccData);
	return __r;
}
int QResource_registerResource3(const char* rccFilename, long long rccFilename_len, const char* resourceRoot, long long resourceRoot_len) {
	QString rccFilename_qstr = QString::fromUtf8(rccFilename, rccFilename_len);
	QString resourceRoot_qstr = QString::fromUtf8(resourceRoot, resourceRoot_len);
	auto&& __r = QResource::registerResource(rccFilename_qstr, resourceRoot_qstr);
	return __r;
}
int QResource_unregisterResource3(const char* rccFilename, long long rccFilename_len, const char* resourceRoot, long long resourceRoot_len) {
	QString rccFilename_qstr = QString::fromUtf8(rccFilename, rccFilename_len);
	QString resourceRoot_qstr = QString::fromUtf8(resourceRoot, resourceRoot_len);
	auto&& __r = QResource::unregisterResource(rccFilename_qstr, resourceRoot_qstr);
	return __r;
}
int QResource_registerResource4(const unsigned char* rccData, const char* resourceRoot, long long resourceRoot_len) {
	QString resourceRoot_qstr = QString::fromUtf8(resourceRoot, resourceRoot_len);
	auto&& __r = QResource::registerResource(rccData, resourceRoot_qstr);
	return __r;
}
int QResource_unregisterResource4(const unsigned char* rccData, const char* resourceRoot, long long resourceRoot_len) {
	QString resourceRoot_qstr = QString::fromUtf8(resourceRoot, resourceRoot_len);
	auto&& __r = QResource::unregisterResource(rccData, resourceRoot_qstr);
	return __r;
}
void QResource_delete(void* self) { delete (QResource*)self; }

}

#include <QtCore/qtranslator.h>
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
#include "gen_qtranslator.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTranslator__Go : public QTranslator {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QTranslator__Go(long long id, A&&... a) : QTranslator(std::forward<A>(a)...), __goId(id) {}
	using QTranslator::sender;
	using QTranslator::senderSignalIndex;
	using QTranslator::receivers;
	using QTranslator::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTranslator::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTranslator::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTranslator::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTranslator::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTranslator::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTranslator::metaObject();
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QTranslator::isEmpty();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTranslator::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTranslator::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTranslator::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTranslator::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTranslator::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTranslator::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTranslator::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTranslator;
static std::mutex __connsMu_QTranslator;

extern "C" {

void* QTranslator_base(void* self) { return (void*)static_cast<QObject*>((QTranslator*)self); }
void QTranslator__set_override(void* self, int idx) { ((QTranslator__Go*)self)->__goOn[idx] = 1; }
void* QTranslator_fromQObject(void* o) { return (void*)qobject_cast<QTranslator*>((QObject*)o); }
void* QTranslator_new(long long __goid) {
	return new QTranslator__Go(__goid);
}
void* QTranslator_new2(long long __goid, void* parent) {
	return new QTranslator__Go(__goid, (QObject*)parent);
}
char* QTranslator_tr(const char* s, long long* __out_len) {
	auto&& __r = QTranslator::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTranslator_translate(void* self, const char* context, const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = ((QTranslator*)self)->translate(context, sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTranslator_isEmpty(void* self) {
	auto&& __r = ((QTranslator*)self)->isEmpty();
	return __r;
}
char* QTranslator_language(void* self, long long* __out_len) {
	auto&& __r = ((QTranslator*)self)->language();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTranslator_filePath(void* self, long long* __out_len) {
	auto&& __r = ((QTranslator*)self)->filePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTranslator_load(void* self, const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = ((QTranslator*)self)->load(filename_qstr);
	return __r;
}
int QTranslator_load2(void* self, void* locale, const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = ((QTranslator*)self)->load(*(QLocale*)locale, filename_qstr);
	return __r;
}
int QTranslator_load3(void* self, const unsigned char* data, int len_) {
	auto&& __r = ((QTranslator*)self)->load(data, len_);
	return __r;
}
void* QTranslator_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTranslator::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTranslator_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTranslator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTranslator_sender(void* self) {
	auto&& __r = ((QTranslator__Go*)self)->sender();
	return (void*)(__r);
}
int QTranslator_senderSignalIndex(void* self) {
	auto&& __r = ((QTranslator__Go*)self)->senderSignalIndex();
	return __r;
}
int QTranslator_receivers(void* self, const char* signal) {
	auto&& __r = ((QTranslator__Go*)self)->receivers(signal);
	return __r;
}
int QTranslator_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTranslator__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTranslator_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTranslator::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTranslator_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTranslator::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTranslator_load4(void* self, const char* filename, long long filename_len, const char* directory, long long directory_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	auto&& __r = ((QTranslator*)self)->load(filename_qstr, directory_qstr);
	return __r;
}
int QTranslator_load5(void* self, const char* filename, long long filename_len, const char* directory, long long directory_len, const char* search_delimiters, long long search_delimiters_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	QString search_delimiters_qstr = QString::fromUtf8(search_delimiters, search_delimiters_len);
	auto&& __r = ((QTranslator*)self)->load(filename_qstr, directory_qstr, search_delimiters_qstr);
	return __r;
}
int QTranslator_load6(void* self, const char* filename, long long filename_len, const char* directory, long long directory_len, const char* search_delimiters, long long search_delimiters_len, const char* suffix, long long suffix_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	QString search_delimiters_qstr = QString::fromUtf8(search_delimiters, search_delimiters_len);
	QString suffix_qstr = QString::fromUtf8(suffix, suffix_len);
	auto&& __r = ((QTranslator*)self)->load(filename_qstr, directory_qstr, search_delimiters_qstr, suffix_qstr);
	return __r;
}
int QTranslator_load7(void* self, void* locale, const char* filename, long long filename_len, const char* prefix, long long prefix_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	auto&& __r = ((QTranslator*)self)->load(*(QLocale*)locale, filename_qstr, prefix_qstr);
	return __r;
}
int QTranslator_load8(void* self, void* locale, const char* filename, long long filename_len, const char* prefix, long long prefix_len, const char* directory, long long directory_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	auto&& __r = ((QTranslator*)self)->load(*(QLocale*)locale, filename_qstr, prefix_qstr, directory_qstr);
	return __r;
}
int QTranslator_load9(void* self, void* locale, const char* filename, long long filename_len, const char* prefix, long long prefix_len, const char* directory, long long directory_len, const char* suffix, long long suffix_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	QString suffix_qstr = QString::fromUtf8(suffix, suffix_len);
	auto&& __r = ((QTranslator*)self)->load(*(QLocale*)locale, filename_qstr, prefix_qstr, directory_qstr, suffix_qstr);
	return __r;
}
int QTranslator_load10(void* self, const unsigned char* data, int len_, const char* directory, long long directory_len) {
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	auto&& __r = ((QTranslator*)self)->load(data, len_, directory_qstr);
	return __r;
}
void* QTranslator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTranslator::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTranslator_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTranslator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
int QTranslator_isEmpty_default(void* self) {
	auto&& __r = ((QTranslator*)self)->QTranslator::isEmpty();
	return __r;
}
void* QTranslator_metaObject_default(void* self) {
	auto&& __r = ((QTranslator*)self)->QTranslator::metaObject();
	return (void*)(__r);
}
int QTranslator_event_default(void* self, void* event) {
	auto&& __r = ((QTranslator*)self)->QTranslator::event((QEvent*)event);
	return __r;
}
int QTranslator_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTranslator*)self)->QTranslator::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTranslator_timerEvent_default(void* self, void* event) {
	((QTranslator__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTranslator_childEvent_default(void* self, void* event) {
	((QTranslator__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTranslator_customEvent_default(void* self, void* event) {
	((QTranslator__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTranslator_connectNotify_default(void* self, void* signal) {
	((QTranslator__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTranslator_disconnectNotify_default(void* self, void* signal) {
	((QTranslator__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTranslator_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTranslator*)self, &QTranslator::destroyed, (QTranslator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTranslator);
	__conns_QTranslator[__id] = __c;
}
void QTranslator_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTranslator*)self, &QTranslator::destroyed, (QTranslator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTranslator);
	__conns_QTranslator[__id] = __c;
}
void QTranslator_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTranslator*)self, &QTranslator::objectNameChanged, (QTranslator*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTranslator);
	__conns_QTranslator[__id] = __c;
}
void QTranslator__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTranslator);
		auto __it = __conns_QTranslator.find(id);
		if (__it == __conns_QTranslator.end()) return;
		__c = __it->second;
		__conns_QTranslator.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTranslator__Go_new(long long id, void* parent) { return new QTranslator__Go(id, (QObject*)parent); }
void QTranslator_delete(void* self) { delete (QTranslator*)self; }

}

#include <QtGui/qabstractfileiconprovider.h>
#include <QtGui/qicon.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qstring.h>
#include "gen_qabstractfileiconprovider.h"
#include "gen_purego_cb.h"

class QAbstractFileIconProvider__Go : public QAbstractFileIconProvider {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QAbstractFileIconProvider__Go(long long id, A&&... a) : QAbstractFileIconProvider(std::forward<A>(a)...), __goId(id) {}
	QIcon icon(QAbstractFileIconProvider::IconType arg1) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -2594270449535773968LL, (int)(arg1), &__h); if (__h) return *(QIcon*)__r; }
		return QAbstractFileIconProvider::icon(arg1);
	}
	QIcon icon(const QFileInfo& arg1) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -909105390709809242LL, (void*)&arg1, &__h); if (__h) return *(QIcon*)__r; }
		return QAbstractFileIconProvider::icon(arg1);
	}
	QString type(const QFileInfo& arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virtp_charp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_charp_voidp)(__goId, -6371403872082420691LL, (void*)&arg1, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAbstractFileIconProvider::type(arg1);
	}
	void setOptions(QFlags<QAbstractFileIconProvider::Option> arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 5783776084322042327LL, (int)(arg1), &__h); if (__h) return; }
		QAbstractFileIconProvider::setOptions(arg1);
	}
	QFlags<QAbstractFileIconProvider::Option> options() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3871419584579008699LL, &__h); if (__h) return QFlags<QAbstractFileIconProvider::Option>(__r); }
		return QAbstractFileIconProvider::options();
	}
};

extern "C" {

void QAbstractFileIconProvider__set_override(void* self, int idx) { ((QAbstractFileIconProvider__Go*)self)->__goOn[idx] = 1; }
void* QAbstractFileIconProvider_new(long long __goid) {
	return new QAbstractFileIconProvider__Go(__goid);
}
void* QAbstractFileIconProvider_icon(void* self, int arg1) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->icon((QAbstractFileIconProvider::IconType)(arg1));
	return new QIcon(__r);
}
void* QAbstractFileIconProvider_icon2(void* self, void* arg1) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->icon(*(QFileInfo*)arg1);
	return new QIcon(__r);
}
char* QAbstractFileIconProvider_type(void* self, void* arg1, long long* __out_len) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->type(*(QFileInfo*)arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractFileIconProvider_setOptions(void* self, int arg1) {
	((QAbstractFileIconProvider*)self)->setOptions(QFlags<QAbstractFileIconProvider::Option>(arg1));
}
int QAbstractFileIconProvider_options(void* self) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->options();
	return (int)(__r);
}
void* QAbstractFileIconProvider_icon_default(void* self, int arg1) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->QAbstractFileIconProvider::icon((QAbstractFileIconProvider::IconType)(arg1));
	return new QIcon(__r);
}
void* QAbstractFileIconProvider_icon_default2(void* self, void* arg1) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->QAbstractFileIconProvider::icon(*(QFileInfo*)arg1);
	return new QIcon(__r);
}
char* QAbstractFileIconProvider_type_default(void* self, void* arg1, long long* __out_len) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->QAbstractFileIconProvider::type(*(QFileInfo*)arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractFileIconProvider_setOptions_default(void* self, int arg1) {
	((QAbstractFileIconProvider*)self)->QAbstractFileIconProvider::setOptions(QFlags<QAbstractFileIconProvider::Option>(arg1));
}
int QAbstractFileIconProvider_options_default(void* self) {
	auto&& __r = ((QAbstractFileIconProvider*)self)->QAbstractFileIconProvider::options();
	return (int)(__r);
}
void* QAbstractFileIconProvider__Go_new(long long id) { return new QAbstractFileIconProvider__Go(id); }
void QAbstractFileIconProvider_delete(void* self) { delete (QAbstractFileIconProvider*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include "gen_qmoveevent.h"

extern "C" {

void* QMoveEvent_base(void* self) { return (void*)static_cast<QEvent*>((QMoveEvent*)self); }
void* QMoveEvent_new(void* pos, void* oldPos) {
	return new QMoveEvent(*(QPoint*)pos, *(QPoint*)oldPos);
}
void* QMoveEvent_pos(void* self) {
	auto&& __r = ((QMoveEvent*)self)->pos();
	return (void*)&(__r);
}
void* QMoveEvent_oldPos(void* self) {
	auto&& __r = ((QMoveEvent*)self)->oldPos();
	return (void*)&(__r);
}
int QMoveEvent_registerEventType() {
	auto&& __r = QMoveEvent::registerEventType();
	return __r;
}
int QMoveEvent_registerEventType2(int hint) {
	auto&& __r = QMoveEvent::registerEventType(hint);
	return __r;
}
void QMoveEvent_delete(void* self) { delete (QMoveEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qeventpoint.h>
#include <QtCore/qobject.h>
#include <QtGui/qinputdevice.h>
#include "gen_qtouchevent.h"

extern "C" {

void* QTouchEvent_base(void* self) { return (void*)static_cast<QPointerEvent*>((QTouchEvent*)self); }
void QTouchEvent_chain(void* self, void** out) {
	QTouchEvent* __p = (QTouchEvent*)self;
	QPointerEvent* __b0 = static_cast<QPointerEvent*>(__p); out[0] = (void*)__b0;
	QInputEvent* __b1 = static_cast<QInputEvent*>(__b0); out[1] = (void*)__b1;
	QEvent* __b2 = static_cast<QEvent*>(__b1); out[2] = (void*)__b2;
}
void* QTouchEvent_new(int eventType) {
	return new QTouchEvent((QEvent::Type)(eventType));
}
void* QTouchEvent_new2(int eventType, void* device, int modifiers, int touchPointStates) {
	return new QTouchEvent((QEvent::Type)(eventType), (QPointingDevice*)device, QFlags<Qt::KeyboardModifier>(modifiers), QFlags<QEventPoint::State>(touchPointStates));
}
void* QTouchEvent_new3(int eventType, void* device) {
	return new QTouchEvent((QEvent::Type)(eventType), (QPointingDevice*)device);
}
void* QTouchEvent_new4(int eventType, void* device, int modifiers) {
	return new QTouchEvent((QEvent::Type)(eventType), (QPointingDevice*)device, QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QTouchEvent_new5(int eventType, void* device, int modifiers, void** touchPoints, long long touchPoints_len) {
	QList<QEventPoint> touchPoints_list; touchPoints_list.reserve(touchPoints_len);
	for (long long __i = 0; __i < touchPoints_len; __i++) touchPoints_list.append(*(QEventPoint*)touchPoints[__i]);
	return new QTouchEvent((QEvent::Type)(eventType), (QPointingDevice*)device, QFlags<Qt::KeyboardModifier>(modifiers), touchPoints_list);
}
void* QTouchEvent_new6(int eventType, void* device, int modifiers, int touchPointStates, void** touchPoints, long long touchPoints_len) {
	QList<QEventPoint> touchPoints_list; touchPoints_list.reserve(touchPoints_len);
	for (long long __i = 0; __i < touchPoints_len; __i++) touchPoints_list.append(*(QEventPoint*)touchPoints[__i]);
	return new QTouchEvent((QEvent::Type)(eventType), (QPointingDevice*)device, QFlags<Qt::KeyboardModifier>(modifiers), QFlags<QEventPoint::State>(touchPointStates), touchPoints_list);
}
void* QTouchEvent_target(void* self) {
	auto&& __r = ((QTouchEvent*)self)->target();
	return (void*)(__r);
}
int QTouchEvent_touchPointStates(void* self) {
	auto&& __r = ((QTouchEvent*)self)->touchPointStates();
	return (int)(__r);
}
void** QTouchEvent_touchPoints(void* self, long long* __out_len) {
	auto&& __r = ((QTouchEvent*)self)->touchPoints();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QEventPoint(__e);
	return __arr;
}
int QTouchEvent_registerEventType() {
	auto&& __r = QTouchEvent::registerEventType();
	return __r;
}
int QTouchEvent_registerEventType2(int hint) {
	auto&& __r = QTouchEvent::registerEventType(hint);
	return __r;
}
void QTouchEvent_delete(void* self) { delete (QTouchEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstring.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtCore/qnamespace.h>
#include "gen_qtextformat.h"

extern "C" {

void* QTextFormat_new() {
	return new QTextFormat();
}
void* QTextFormat_new2(int type_) {
	return new QTextFormat(type_);
}
void* QTextFormat_new3(void* rhs) {
	return new QTextFormat(*(QTextFormat*)rhs);
}
void* QTextFormat_OperatorAssign(void* self, void* rhs) {
	auto&& __r = ((QTextFormat*)self)->operator=(*(QTextFormat*)rhs);
	return (void*)&(__r);
}
void QTextFormat_swap(void* self, void* other) {
	((QTextFormat*)self)->swap(*(QTextFormat*)other);
}
void QTextFormat_merge(void* self, void* other) {
	((QTextFormat*)self)->merge(*(QTextFormat*)other);
}
int QTextFormat_isValid(void* self) {
	auto&& __r = ((QTextFormat*)self)->isValid();
	return __r;
}
int QTextFormat_isEmpty(void* self) {
	auto&& __r = ((QTextFormat*)self)->isEmpty();
	return __r;
}
int QTextFormat_type(void* self) {
	auto&& __r = ((QTextFormat*)self)->type();
	return __r;
}
int QTextFormat_objectIndex(void* self) {
	auto&& __r = ((QTextFormat*)self)->objectIndex();
	return __r;
}
void QTextFormat_setObjectIndex(void* self, int object) {
	((QTextFormat*)self)->setObjectIndex(object);
}
void* QTextFormat_property(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->property(propertyId);
	return new QVariant(__r);
}
void QTextFormat_setProperty(void* self, int propertyId, void* value) {
	((QTextFormat*)self)->setProperty(propertyId, *(QVariant*)value);
}
void QTextFormat_clearProperty(void* self, int propertyId) {
	((QTextFormat*)self)->clearProperty(propertyId);
}
int QTextFormat_hasProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->hasProperty(propertyId);
	return __r;
}
int QTextFormat_boolProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->boolProperty(propertyId);
	return __r;
}
int QTextFormat_intProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->intProperty(propertyId);
	return __r;
}
double QTextFormat_doubleProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->doubleProperty(propertyId);
	return __r;
}
char* QTextFormat_stringProperty(void* self, int propertyId, long long* __out_len) {
	auto&& __r = ((QTextFormat*)self)->stringProperty(propertyId);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextFormat_colorProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->colorProperty(propertyId);
	return new QColor(__r);
}
void* QTextFormat_penProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->penProperty(propertyId);
	return new QPen(__r);
}
void* QTextFormat_brushProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->brushProperty(propertyId);
	return new QBrush(__r);
}
void* QTextFormat_lengthProperty(void* self, int propertyId) {
	auto&& __r = ((QTextFormat*)self)->lengthProperty(propertyId);
	return new QTextLength(__r);
}
void** QTextFormat_lengthVectorProperty(void* self, int propertyId, long long* __out_len) {
	auto&& __r = ((QTextFormat*)self)->lengthVectorProperty(propertyId);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextLength(__e);
	return __arr;
}
void QTextFormat_setProperty2(void* self, int propertyId, void** lengths, long long lengths_len) {
	QList<QTextLength> lengths_list; lengths_list.reserve(lengths_len);
	for (long long __i = 0; __i < lengths_len; __i++) lengths_list.append(*(QTextLength*)lengths[__i]);
	((QTextFormat*)self)->setProperty(propertyId, lengths_list);
}
int* QTextFormat_properties(void* self, void** __out_vals, long long* __out_len) {
	auto&& __r = ((QTextFormat*)self)->properties();
	*__out_len = __r.size();
	int* __keys = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	void** __vals = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto __it = __r.cbegin(); __it != __r.cend(); ++__it) { __keys[__i] = __it.key(); __vals[__i] = new QVariant(__it.value()); __i++; }
	*__out_vals = (void*)__vals;
	return __keys;
}
int QTextFormat_propertyCount(void* self) {
	auto&& __r = ((QTextFormat*)self)->propertyCount();
	return __r;
}
void QTextFormat_setObjectType(void* self, int type_) {
	((QTextFormat*)self)->setObjectType(type_);
}
int QTextFormat_objectType(void* self) {
	auto&& __r = ((QTextFormat*)self)->objectType();
	return __r;
}
int QTextFormat_isCharFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isCharFormat();
	return __r;
}
int QTextFormat_isBlockFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isBlockFormat();
	return __r;
}
int QTextFormat_isListFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isListFormat();
	return __r;
}
int QTextFormat_isFrameFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isFrameFormat();
	return __r;
}
int QTextFormat_isImageFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isImageFormat();
	return __r;
}
int QTextFormat_isTableFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isTableFormat();
	return __r;
}
int QTextFormat_isTableCellFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->isTableCellFormat();
	return __r;
}
void* QTextFormat_toBlockFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toBlockFormat();
	return new QTextBlockFormat(__r);
}
void* QTextFormat_toCharFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toCharFormat();
	return new QTextCharFormat(__r);
}
void* QTextFormat_toListFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toListFormat();
	return new QTextListFormat(__r);
}
void* QTextFormat_toTableFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toTableFormat();
	return new QTextTableFormat(__r);
}
void* QTextFormat_toFrameFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toFrameFormat();
	return new QTextFrameFormat(__r);
}
void* QTextFormat_toImageFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toImageFormat();
	return new QTextImageFormat(__r);
}
void* QTextFormat_toTableCellFormat(void* self) {
	auto&& __r = ((QTextFormat*)self)->toTableCellFormat();
	return new QTextTableCellFormat(__r);
}
int QTextFormat_OperatorEqual(void* self, void* rhs) {
	auto&& __r = ((QTextFormat*)self)->operator==(*(QTextFormat*)rhs);
	return __r;
}
int QTextFormat_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = ((QTextFormat*)self)->operator!=(*(QTextFormat*)rhs);
	return __r;
}
void* QTextFormat_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QTextFormat*)self);
	return new QVariant(__r);
}
void QTextFormat_setLayoutDirection(void* self, int direction) {
	((QTextFormat*)self)->setLayoutDirection((Qt::LayoutDirection)(direction));
}
int QTextFormat_layoutDirection(void* self) {
	auto&& __r = ((QTextFormat*)self)->layoutDirection();
	return (int)(__r);
}
void QTextFormat_setBackground(void* self, void* brush) {
	((QTextFormat*)self)->setBackground(*(QBrush*)brush);
}
void* QTextFormat_background(void* self) {
	auto&& __r = ((QTextFormat*)self)->background();
	return new QBrush(__r);
}
void QTextFormat_clearBackground(void* self) {
	((QTextFormat*)self)->clearBackground();
}
void QTextFormat_setForeground(void* self, void* brush) {
	((QTextFormat*)self)->setForeground(*(QBrush*)brush);
}
void* QTextFormat_foreground(void* self) {
	auto&& __r = ((QTextFormat*)self)->foreground();
	return new QBrush(__r);
}
void QTextFormat_clearForeground(void* self) {
	((QTextFormat*)self)->clearForeground();
}
void* QTextFormat_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QTextFormat*)self);
	return new QVariant(__r);
}
void QTextFormat_delete(void* self) { delete (QTextFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include "gen_qaccessibleeditabletextinterface.h"
#include "gen_purego_cb.h"

class QAccessibleEditableTextInterface__Go : public QAccessibleEditableTextInterface {
public:
	long long __goId;
	unsigned char __goOn[3] = {};
	template<typename... A> QAccessibleEditableTextInterface__Go(long long id, A&&... a) : QAccessibleEditableTextInterface(std::forward<A>(a)...), __goId(id) {}
	void deleteText(int startOffset, int endOffset) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 4554921347280463265LL, startOffset, endOffset, &__h); if (__h) return; }
		return;
	}
	void insertText(int offset, const QString& text) override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[1] && cppgo_virt_void_int_constcharp_longlong) { ((void(*)(long long, long long, int, const char*, long long, int*))cppgo_virt_void_int_constcharp_longlong)(__goId, -7342818475556881681LL, offset, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return; }
		return;
	}
	void replaceText(int startOffset, int endOffset, const QString& text) override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[2] && cppgo_virt_void_int_int_constcharp_longlong) { ((void(*)(long long, long long, int, int, const char*, long long, int*))cppgo_virt_void_int_int_constcharp_longlong)(__goId, 7391831930358964452LL, startOffset, endOffset, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QAccessibleEditableTextInterface__set_override(void* self, int idx) { ((QAccessibleEditableTextInterface__Go*)self)->__goOn[idx] = 1; }
void QAccessibleEditableTextInterface_deleteText(void* self, int startOffset, int endOffset) {
	((QAccessibleEditableTextInterface*)self)->deleteText(startOffset, endOffset);
}
void QAccessibleEditableTextInterface_insertText(void* self, int offset, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAccessibleEditableTextInterface*)self)->insertText(offset, text_qstr);
}
void QAccessibleEditableTextInterface_replaceText(void* self, int startOffset, int endOffset, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAccessibleEditableTextInterface*)self)->replaceText(startOffset, endOffset, text_qstr);
}
void* QAccessibleEditableTextInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleEditableTextInterface*)self)->operator=(*(QAccessibleEditableTextInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleEditableTextInterface__Go_new(long long id) { return new QAccessibleEditableTextInterface__Go(id); }
void QAccessibleEditableTextInterface_delete(void* self) { delete (QAccessibleEditableTextInterface*)self; }

}

#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qactiongroup.h>
#include <QtGui/qkeysequence.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qfont.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qaction.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAction__Go : public QAction {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QAction__Go(long long id, A&&... a) : QAction(std::forward<A>(a)...), __goId(id) {}
	using QAction::sender;
	using QAction::senderSignalIndex;
	using QAction::receivers;
	using QAction::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QAction::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAction::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAction::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAction::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAction::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAction::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAction::metaObject();
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QAction::event(arg1);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAction::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAction::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAction::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAction::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAction::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAction::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAction;
static std::mutex __connsMu_QAction;

extern "C" {

void* QAction_base(void* self) { return (void*)static_cast<QObject*>((QAction*)self); }
void QAction__set_override(void* self, int idx) { ((QAction__Go*)self)->__goOn[idx] = 1; }
void* QAction_fromQObject(void* o) { return (void*)qobject_cast<QAction*>((QObject*)o); }
void* QAction_new(long long __goid) {
	return new QAction__Go(__goid);
}
void* QAction_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAction__Go(__goid, text_qstr);
}
void* QAction_new3(long long __goid, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAction__Go(__goid, *(QIcon*)icon, text_qstr);
}
void* QAction_new4(long long __goid, void* parent) {
	return new QAction__Go(__goid, (QObject*)parent);
}
void* QAction_new5(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAction__Go(__goid, text_qstr, (QObject*)parent);
}
void* QAction_new6(long long __goid, void* icon, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAction__Go(__goid, *(QIcon*)icon, text_qstr, (QObject*)parent);
}
char* QAction_tr(const char* s, long long* __out_len) {
	auto&& __r = QAction::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QAction_associatedObjects(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->associatedObjects();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QAction_setActionGroup(void* self, void* group) {
	((QAction*)self)->setActionGroup((QActionGroup*)group);
}
void* QAction_actionGroup(void* self) {
	auto&& __r = ((QAction*)self)->actionGroup();
	return (void*)(__r);
}
void QAction_setIcon(void* self, void* icon) {
	((QAction*)self)->setIcon(*(QIcon*)icon);
}
void* QAction_icon(void* self) {
	auto&& __r = ((QAction*)self)->icon();
	return new QIcon(__r);
}
void QAction_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAction*)self)->setText(text_qstr);
}
char* QAction_text(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAction_setIconText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAction*)self)->setIconText(text_qstr);
}
char* QAction_iconText(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->iconText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAction_setToolTip(void* self, const char* tip, long long tip_len) {
	QString tip_qstr = QString::fromUtf8(tip, tip_len);
	((QAction*)self)->setToolTip(tip_qstr);
}
char* QAction_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAction_setStatusTip(void* self, const char* statusTip, long long statusTip_len) {
	QString statusTip_qstr = QString::fromUtf8(statusTip, statusTip_len);
	((QAction*)self)->setStatusTip(statusTip_qstr);
}
char* QAction_statusTip(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->statusTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAction_setWhatsThis(void* self, const char* what, long long what_len) {
	QString what_qstr = QString::fromUtf8(what, what_len);
	((QAction*)self)->setWhatsThis(what_qstr);
}
char* QAction_whatsThis(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->whatsThis();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAction_setPriority(void* self, int priority) {
	((QAction*)self)->setPriority((QAction::Priority)(priority));
}
int QAction_priority(void* self) {
	auto&& __r = ((QAction*)self)->priority();
	return (int)(__r);
}
void QAction_setSeparator(void* self, int b) {
	((QAction*)self)->setSeparator(b);
}
int QAction_isSeparator(void* self) {
	auto&& __r = ((QAction*)self)->isSeparator();
	return __r;
}
void QAction_setShortcut(void* self, void* shortcut) {
	((QAction*)self)->setShortcut(*(QKeySequence*)shortcut);
}
void* QAction_shortcut(void* self) {
	auto&& __r = ((QAction*)self)->shortcut();
	return new QKeySequence(__r);
}
void QAction_setShortcuts(void* self, void** shortcuts, long long shortcuts_len) {
	QList<QKeySequence> shortcuts_list; shortcuts_list.reserve(shortcuts_len);
	for (long long __i = 0; __i < shortcuts_len; __i++) shortcuts_list.append(*(QKeySequence*)shortcuts[__i]);
	((QAction*)self)->setShortcuts(shortcuts_list);
}
void QAction_setShortcuts2(void* self, int arg1) {
	((QAction*)self)->setShortcuts((QKeySequence::StandardKey)(arg1));
}
void** QAction_shortcuts(void* self, long long* __out_len) {
	auto&& __r = ((QAction*)self)->shortcuts();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QKeySequence(__e);
	return __arr;
}
void QAction_setShortcutContext(void* self, int context) {
	((QAction*)self)->setShortcutContext((Qt::ShortcutContext)(context));
}
int QAction_shortcutContext(void* self) {
	auto&& __r = ((QAction*)self)->shortcutContext();
	return (int)(__r);
}
void QAction_setAutoRepeat(void* self, int arg1) {
	((QAction*)self)->setAutoRepeat(arg1);
}
int QAction_autoRepeat(void* self) {
	auto&& __r = ((QAction*)self)->autoRepeat();
	return __r;
}
void QAction_setFont(void* self, void* font) {
	((QAction*)self)->setFont(*(QFont*)font);
}
void* QAction_font(void* self) {
	auto&& __r = ((QAction*)self)->font();
	return new QFont(__r);
}
void QAction_setCheckable(void* self, int arg1) {
	((QAction*)self)->setCheckable(arg1);
}
int QAction_isCheckable(void* self) {
	auto&& __r = ((QAction*)self)->isCheckable();
	return __r;
}
void* QAction_data(void* self) {
	auto&& __r = ((QAction*)self)->data();
	return new QVariant(__r);
}
void QAction_setData(void* self, void* var_) {
	((QAction*)self)->setData(*(QVariant*)var_);
}
int QAction_isChecked(void* self) {
	auto&& __r = ((QAction*)self)->isChecked();
	return __r;
}
int QAction_isEnabled(void* self) {
	auto&& __r = ((QAction*)self)->isEnabled();
	return __r;
}
int QAction_isVisible(void* self) {
	auto&& __r = ((QAction*)self)->isVisible();
	return __r;
}
void QAction_activate(void* self, int event) {
	((QAction*)self)->activate((QAction::ActionEvent)(event));
}
void QAction_setMenuRole(void* self, int menuRole) {
	((QAction*)self)->setMenuRole((QAction::MenuRole)(menuRole));
}
int QAction_menuRole(void* self) {
	auto&& __r = ((QAction*)self)->menuRole();
	return (int)(__r);
}
void QAction_setIconVisibleInMenu(void* self, int visible) {
	((QAction*)self)->setIconVisibleInMenu(visible);
}
int QAction_isIconVisibleInMenu(void* self) {
	auto&& __r = ((QAction*)self)->isIconVisibleInMenu();
	return __r;
}
void QAction_setShortcutVisibleInContextMenu(void* self, int show) {
	((QAction*)self)->setShortcutVisibleInContextMenu(show);
}
int QAction_isShortcutVisibleInContextMenu(void* self) {
	auto&& __r = ((QAction*)self)->isShortcutVisibleInContextMenu();
	return __r;
}
int QAction_showStatusText(void* self) {
	auto&& __r = ((QAction*)self)->showStatusText();
	return __r;
}
void QAction_trigger(void* self) {
	((QAction*)self)->trigger();
}
void QAction_hover(void* self) {
	((QAction*)self)->hover();
}
void QAction_setChecked(void* self, int arg1) {
	((QAction*)self)->setChecked(arg1);
}
void QAction_toggle(void* self) {
	((QAction*)self)->toggle();
}
void QAction_setEnabled(void* self, int arg1) {
	((QAction*)self)->setEnabled(arg1);
}
void QAction_resetEnabled(void* self) {
	((QAction*)self)->resetEnabled();
}
void QAction_setDisabled(void* self, int b) {
	((QAction*)self)->setDisabled(b);
}
void QAction_setVisible(void* self, int arg1) {
	((QAction*)self)->setVisible(arg1);
}
void QAction_changed(void* self) {
	((QAction*)self)->changed();
}
void QAction_enabledChanged(void* self, int enabled) {
	((QAction*)self)->enabledChanged(enabled);
}
void QAction_checkableChanged(void* self, int checkable) {
	((QAction*)self)->checkableChanged(checkable);
}
void QAction_visibleChanged(void* self) {
	((QAction*)self)->visibleChanged();
}
void QAction_triggered(void* self) {
	((QAction*)self)->triggered();
}
void QAction_hovered(void* self) {
	((QAction*)self)->hovered();
}
void QAction_toggled(void* self, int arg1) {
	((QAction*)self)->toggled(arg1);
}
void* QAction_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAction::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAction_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAction::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAction_sender(void* self) {
	auto&& __r = ((QAction__Go*)self)->sender();
	return (void*)(__r);
}
int QAction_senderSignalIndex(void* self) {
	auto&& __r = ((QAction__Go*)self)->senderSignalIndex();
	return __r;
}
int QAction_receivers(void* self, const char* signal) {
	auto&& __r = ((QAction__Go*)self)->receivers(signal);
	return __r;
}
int QAction_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAction__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAction_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAction::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAction_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAction::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAction_showStatusText2(void* self, void* object) {
	auto&& __r = ((QAction*)self)->showStatusText((QObject*)object);
	return __r;
}
void QAction_triggered2(void* self, int checked) {
	((QAction*)self)->triggered(checked);
}
void* QAction_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAction::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAction_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAction::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAction_metaObject_default(void* self) {
	auto&& __r = ((QAction*)self)->QAction::metaObject();
	return (void*)(__r);
}
int QAction_event_default(void* self, void* arg1) {
	auto&& __r = ((QAction__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
int QAction_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAction*)self)->QAction::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAction_timerEvent_default(void* self, void* event) {
	((QAction__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAction_childEvent_default(void* self, void* event) {
	((QAction__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAction_customEvent_default(void* self, void* event) {
	((QAction__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAction_connectNotify_default(void* self, void* signal) {
	((QAction__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAction_disconnectNotify_default(void* self, void* signal) {
	((QAction__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAction_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::changed, (QAction*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::enabledChanged, (QAction*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_checkableChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::checkableChanged, (QAction*)self, [__id](bool checkable){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checkable);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::visibleChanged, (QAction*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_triggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::triggered, (QAction*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_triggered2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::triggered, (QAction*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_hovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::hovered, (QAction*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::toggled, (QAction*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::destroyed, (QAction*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::destroyed, (QAction*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAction*)self, &QAction::objectNameChanged, (QAction*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAction);
	__conns_QAction[__id] = __c;
}
void QAction__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAction);
		auto __it = __conns_QAction.find(id);
		if (__it == __conns_QAction.end()) return;
		__c = __it->second;
		__conns_QAction.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAction__Go_new(long long id, void* parent) { return new QAction__Go(id, (QObject*)parent); }
void QAction_delete(void* self) { delete (QAction*)self; }

}

#include <QtGui/qimageiohandler.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtGui/qimage.h>
#include <QtCore/qvariant.h>
#include <QtCore/qrect.h>
#include <QtCore/qsize.h>
#include "gen_qimageiohandler.h"
#include "gen_purego_cb.h"

class QImageIOHandler__Go : public QImageIOHandler {
public:
	long long __goId;
	unsigned char __goOn[13] = {};
	template<typename... A> QImageIOHandler__Go(long long id, A&&... a) : QImageIOHandler(std::forward<A>(a)...), __goId(id) {}
	bool canRead() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5697610016496362555LL, &__h); if (__h) return __r; }
		return bool{};
	}
	bool read(QImage* image) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5541207787354052101LL, (void*)image, &__h); if (__h) return __r; }
		return bool{};
	}
	bool write(const QImage& image) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -5099743077423337732LL, (void*)&image, &__h); if (__h) return __r; }
		return QImageIOHandler::write(image);
	}
	QVariant option(QImageIOHandler::ImageOption option) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 270525795752552276LL, (int)(option), &__h); if (__h) return *(QVariant*)__r; }
		return QImageIOHandler::option(option);
	}
	void setOption(QImageIOHandler::ImageOption option, const QVariant& value) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, 5527147985269015870LL, (int)(option), (void*)&value, &__h); if (__h) return; }
		QImageIOHandler::setOption(option, value);
	}
	bool supportsOption(QImageIOHandler::ImageOption option) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 89707390584546352LL, (int)(option), &__h); if (__h) return __r; }
		return QImageIOHandler::supportsOption(option);
	}
	bool jumpToNextImage() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3906762202460275946LL, &__h); if (__h) return __r; }
		return QImageIOHandler::jumpToNextImage();
	}
	bool jumpToImage(int imageNumber) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 732703795301504057LL, imageNumber, &__h); if (__h) return __r; }
		return QImageIOHandler::jumpToImage(imageNumber);
	}
	int loopCount() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 60103087788910490LL, &__h); if (__h) return __r; }
		return QImageIOHandler::loopCount();
	}
	int imageCount() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7262609269034271663LL, &__h); if (__h) return __r; }
		return QImageIOHandler::imageCount();
	}
	int nextImageDelay() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5074503007600271320LL, &__h); if (__h) return __r; }
		return QImageIOHandler::nextImageDelay();
	}
	int currentImageNumber() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 2054658236031792398LL, &__h); if (__h) return __r; }
		return QImageIOHandler::currentImageNumber();
	}
	QRect currentImageRect() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2007698496588377553LL, &__h); if (__h) return *(QRect*)__r; }
		return QImageIOHandler::currentImageRect();
	}
};

extern "C" {

void QImageIOHandler__set_override(void* self, int idx) { ((QImageIOHandler__Go*)self)->__goOn[idx] = 1; }
void QImageIOHandler_setDevice(void* self, void* device) {
	((QImageIOHandler*)self)->setDevice((QIODevice*)device);
}
void* QImageIOHandler_device(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->device();
	return (void*)(__r);
}
void QImageIOHandler_setFormat(void* self, void* format) {
	((QImageIOHandler*)self)->setFormat(*(QByteArray*)format);
}
void QImageIOHandler_setFormat2(void* self, void* format) {
	((QImageIOHandler*)self)->setFormat(*(QByteArray*)format);
}
void* QImageIOHandler_format(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->format();
	return new QByteArray(__r);
}
int QImageIOHandler_canRead(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->canRead();
	return __r;
}
int QImageIOHandler_read(void* self, void* image) {
	auto&& __r = ((QImageIOHandler*)self)->read((QImage*)image);
	return __r;
}
int QImageIOHandler_write(void* self, void* image) {
	auto&& __r = ((QImageIOHandler*)self)->write(*(QImage*)image);
	return __r;
}
void* QImageIOHandler_option(void* self, int option) {
	auto&& __r = ((QImageIOHandler*)self)->option((QImageIOHandler::ImageOption)(option));
	return new QVariant(__r);
}
void QImageIOHandler_setOption(void* self, int option, void* value) {
	((QImageIOHandler*)self)->setOption((QImageIOHandler::ImageOption)(option), *(QVariant*)value);
}
int QImageIOHandler_supportsOption(void* self, int option) {
	auto&& __r = ((QImageIOHandler*)self)->supportsOption((QImageIOHandler::ImageOption)(option));
	return __r;
}
int QImageIOHandler_jumpToNextImage(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->jumpToNextImage();
	return __r;
}
int QImageIOHandler_jumpToImage(void* self, int imageNumber) {
	auto&& __r = ((QImageIOHandler*)self)->jumpToImage(imageNumber);
	return __r;
}
int QImageIOHandler_loopCount(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->loopCount();
	return __r;
}
int QImageIOHandler_imageCount(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->imageCount();
	return __r;
}
int QImageIOHandler_nextImageDelay(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->nextImageDelay();
	return __r;
}
int QImageIOHandler_currentImageNumber(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->currentImageNumber();
	return __r;
}
void* QImageIOHandler_currentImageRect(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->currentImageRect();
	return new QRect(__r);
}
int QImageIOHandler_allocateImage(void* size, int format, void* image) {
	auto&& __r = QImageIOHandler::allocateImage(*(QSize*)size, (QImage::Format)(format), (QImage*)image);
	return __r;
}
int QImageIOHandler_write_default(void* self, void* image) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::write(*(QImage*)image);
	return __r;
}
void* QImageIOHandler_option_default(void* self, int option) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::option((QImageIOHandler::ImageOption)(option));
	return new QVariant(__r);
}
void QImageIOHandler_setOption_default(void* self, int option, void* value) {
	((QImageIOHandler*)self)->QImageIOHandler::setOption((QImageIOHandler::ImageOption)(option), *(QVariant*)value);
}
int QImageIOHandler_supportsOption_default(void* self, int option) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::supportsOption((QImageIOHandler::ImageOption)(option));
	return __r;
}
int QImageIOHandler_jumpToNextImage_default(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::jumpToNextImage();
	return __r;
}
int QImageIOHandler_jumpToImage_default(void* self, int imageNumber) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::jumpToImage(imageNumber);
	return __r;
}
int QImageIOHandler_loopCount_default(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::loopCount();
	return __r;
}
int QImageIOHandler_imageCount_default(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::imageCount();
	return __r;
}
int QImageIOHandler_nextImageDelay_default(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::nextImageDelay();
	return __r;
}
int QImageIOHandler_currentImageNumber_default(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::currentImageNumber();
	return __r;
}
void* QImageIOHandler_currentImageRect_default(void* self) {
	auto&& __r = ((QImageIOHandler*)self)->QImageIOHandler::currentImageRect();
	return new QRect(__r);
}
void* QImageIOHandler__Go_new(long long id) { return new QImageIOHandler__Go(id); }
void QImageIOHandler_delete(void* self) { delete (QImageIOHandler*)self; }

}

#include <QtGui/qpainterpath.h>
#include <QtGui/qpen.h>
#include <QtCore/qnamespace.h>
#include "gen_qpainterpathstroker.h"

extern "C" {

void* QPainterPathStroker_new() {
	return new QPainterPathStroker();
}
void* QPainterPathStroker_new2(void* pen) {
	return new QPainterPathStroker(*(QPen*)pen);
}
void QPainterPathStroker_setWidth(void* self, double width) {
	((QPainterPathStroker*)self)->setWidth(width);
}
double QPainterPathStroker_width(void* self) {
	auto&& __r = ((QPainterPathStroker*)self)->width();
	return __r;
}
void QPainterPathStroker_setCapStyle(void* self, int style) {
	((QPainterPathStroker*)self)->setCapStyle((Qt::PenCapStyle)(style));
}
int QPainterPathStroker_capStyle(void* self) {
	auto&& __r = ((QPainterPathStroker*)self)->capStyle();
	return (int)(__r);
}
void QPainterPathStroker_setJoinStyle(void* self, int style) {
	((QPainterPathStroker*)self)->setJoinStyle((Qt::PenJoinStyle)(style));
}
int QPainterPathStroker_joinStyle(void* self) {
	auto&& __r = ((QPainterPathStroker*)self)->joinStyle();
	return (int)(__r);
}
void QPainterPathStroker_setMiterLimit(void* self, double length) {
	((QPainterPathStroker*)self)->setMiterLimit(length);
}
double QPainterPathStroker_miterLimit(void* self) {
	auto&& __r = ((QPainterPathStroker*)self)->miterLimit();
	return __r;
}
void QPainterPathStroker_setCurveThreshold(void* self, double threshold) {
	((QPainterPathStroker*)self)->setCurveThreshold(threshold);
}
double QPainterPathStroker_curveThreshold(void* self) {
	auto&& __r = ((QPainterPathStroker*)self)->curveThreshold();
	return __r;
}
void QPainterPathStroker_setDashPattern(void* self, int arg1) {
	((QPainterPathStroker*)self)->setDashPattern((Qt::PenStyle)(arg1));
}
void QPainterPathStroker_setDashPattern2(void* self, double* dashPattern, long long dashPattern_len) {
	QList<double> dashPattern_list; dashPattern_list.reserve(dashPattern_len);
	for (long long __i = 0; __i < dashPattern_len; __i++) dashPattern_list.append(dashPattern[__i]);
	((QPainterPathStroker*)self)->setDashPattern(dashPattern_list);
}
double* QPainterPathStroker_dashPattern(void* self, long long* __out_len) {
	auto&& __r = ((QPainterPathStroker*)self)->dashPattern();
	*__out_len = __r.size();
	double* __arr = (double*)malloc(sizeof(double) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QPainterPathStroker_setDashOffset(void* self, double offset) {
	((QPainterPathStroker*)self)->setDashOffset(offset);
}
double QPainterPathStroker_dashOffset(void* self) {
	auto&& __r = ((QPainterPathStroker*)self)->dashOffset();
	return __r;
}
void* QPainterPathStroker_createStroke(void* self, void* path) {
	auto&& __r = ((QPainterPathStroker*)self)->createStroke(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void QPainterPathStroker_delete(void* self) { delete (QPainterPathStroker*)self; }

}

#include <QtGui/qsyntaxhighlighter.h>
#include <QtCore/qobject.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qcolor.h>
#include <QtGui/qfont.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsyntaxhighlighter.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSyntaxHighlighter__Go : public QSyntaxHighlighter {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QSyntaxHighlighter__Go(long long id, A&&... a) : QSyntaxHighlighter(std::forward<A>(a)...), __goId(id) {}
	using QSyntaxHighlighter::setFormat;
	using QSyntaxHighlighter::format;
	using QSyntaxHighlighter::previousBlockState;
	using QSyntaxHighlighter::currentBlockState;
	using QSyntaxHighlighter::setCurrentBlockState;
	using QSyntaxHighlighter::setCurrentBlockUserData;
	using QSyntaxHighlighter::currentBlockUserData;
	using QSyntaxHighlighter::currentBlock;
	using QSyntaxHighlighter::sender;
	using QSyntaxHighlighter::senderSignalIndex;
	using QSyntaxHighlighter::receivers;
	using QSyntaxHighlighter::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSyntaxHighlighter::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSyntaxHighlighter::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSyntaxHighlighter::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSyntaxHighlighter::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSyntaxHighlighter::disconnectNotify(__a0); }
	void highlightBlock(const QString& text) override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[0] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -2478846640236730472LL, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSyntaxHighlighter::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSyntaxHighlighter::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSyntaxHighlighter::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSyntaxHighlighter::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSyntaxHighlighter::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSyntaxHighlighter::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSyntaxHighlighter::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSyntaxHighlighter::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSyntaxHighlighter;
static std::mutex __connsMu_QSyntaxHighlighter;

extern "C" {

void* QSyntaxHighlighter_base(void* self) { return (void*)static_cast<QObject*>((QSyntaxHighlighter*)self); }
void QSyntaxHighlighter__set_override(void* self, int idx) { ((QSyntaxHighlighter__Go*)self)->__goOn[idx] = 1; }
void* QSyntaxHighlighter_fromQObject(void* o) { return (void*)qobject_cast<QSyntaxHighlighter*>((QObject*)o); }
char* QSyntaxHighlighter_tr(const char* s, long long* __out_len) {
	auto&& __r = QSyntaxHighlighter::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSyntaxHighlighter_setDocument(void* self, void* doc) {
	((QSyntaxHighlighter*)self)->setDocument((QTextDocument*)doc);
}
void* QSyntaxHighlighter_document(void* self) {
	auto&& __r = ((QSyntaxHighlighter*)self)->document();
	return (void*)(__r);
}
void QSyntaxHighlighter_rehighlight(void* self) {
	((QSyntaxHighlighter*)self)->rehighlight();
}
void QSyntaxHighlighter_rehighlightBlock(void* self, void* block) {
	((QSyntaxHighlighter*)self)->rehighlightBlock(*(QTextBlock*)block);
}
void QSyntaxHighlighter_setFormat(void* self, int start, int count, void* format) {
	((QSyntaxHighlighter__Go*)self)->setFormat(start, count, *(QTextCharFormat*)format);
}
void QSyntaxHighlighter_setFormat2(void* self, int start, int count, void* color) {
	((QSyntaxHighlighter__Go*)self)->setFormat(start, count, *(QColor*)color);
}
void QSyntaxHighlighter_setFormat3(void* self, int start, int count, void* font) {
	((QSyntaxHighlighter__Go*)self)->setFormat(start, count, *(QFont*)font);
}
void* QSyntaxHighlighter_format(void* self, int pos) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->format(pos);
	return new QTextCharFormat(__r);
}
int QSyntaxHighlighter_previousBlockState(void* self) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->previousBlockState();
	return __r;
}
int QSyntaxHighlighter_currentBlockState(void* self) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->currentBlockState();
	return __r;
}
void QSyntaxHighlighter_setCurrentBlockState(void* self, int newState) {
	((QSyntaxHighlighter__Go*)self)->setCurrentBlockState(newState);
}
void QSyntaxHighlighter_setCurrentBlockUserData(void* self, void* data) {
	((QSyntaxHighlighter__Go*)self)->setCurrentBlockUserData((QTextBlockUserData*)data);
}
void* QSyntaxHighlighter_currentBlockUserData(void* self) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->currentBlockUserData();
	return (void*)(__r);
}
void* QSyntaxHighlighter_currentBlock(void* self) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->currentBlock();
	return new QTextBlock(__r);
}
void* QSyntaxHighlighter_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSyntaxHighlighter::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSyntaxHighlighter_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSyntaxHighlighter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSyntaxHighlighter_sender(void* self) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->sender();
	return (void*)(__r);
}
int QSyntaxHighlighter_senderSignalIndex(void* self) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->senderSignalIndex();
	return __r;
}
int QSyntaxHighlighter_receivers(void* self, const char* signal) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->receivers(signal);
	return __r;
}
int QSyntaxHighlighter_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSyntaxHighlighter__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSyntaxHighlighter_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSyntaxHighlighter::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSyntaxHighlighter_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSyntaxHighlighter::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSyntaxHighlighter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSyntaxHighlighter::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSyntaxHighlighter_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSyntaxHighlighter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSyntaxHighlighter_metaObject_default(void* self) {
	auto&& __r = ((QSyntaxHighlighter*)self)->QSyntaxHighlighter::metaObject();
	return (void*)(__r);
}
int QSyntaxHighlighter_event_default(void* self, void* event) {
	auto&& __r = ((QSyntaxHighlighter*)self)->QSyntaxHighlighter::event((QEvent*)event);
	return __r;
}
int QSyntaxHighlighter_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSyntaxHighlighter*)self)->QSyntaxHighlighter::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSyntaxHighlighter_timerEvent_default(void* self, void* event) {
	((QSyntaxHighlighter__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSyntaxHighlighter_childEvent_default(void* self, void* event) {
	((QSyntaxHighlighter__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSyntaxHighlighter_customEvent_default(void* self, void* event) {
	((QSyntaxHighlighter__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSyntaxHighlighter_connectNotify_default(void* self, void* signal) {
	((QSyntaxHighlighter__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSyntaxHighlighter_disconnectNotify_default(void* self, void* signal) {
	((QSyntaxHighlighter__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSyntaxHighlighter_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSyntaxHighlighter*)self, &QSyntaxHighlighter::destroyed, (QSyntaxHighlighter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSyntaxHighlighter);
	__conns_QSyntaxHighlighter[__id] = __c;
}
void QSyntaxHighlighter_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSyntaxHighlighter*)self, &QSyntaxHighlighter::destroyed, (QSyntaxHighlighter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSyntaxHighlighter);
	__conns_QSyntaxHighlighter[__id] = __c;
}
void QSyntaxHighlighter_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSyntaxHighlighter*)self, &QSyntaxHighlighter::objectNameChanged, (QSyntaxHighlighter*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSyntaxHighlighter);
	__conns_QSyntaxHighlighter[__id] = __c;
}
void QSyntaxHighlighter__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSyntaxHighlighter);
		auto __it = __conns_QSyntaxHighlighter.find(id);
		if (__it == __conns_QSyntaxHighlighter.end()) return;
		__c = __it->second;
		__conns_QSyntaxHighlighter.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSyntaxHighlighter__Go_new(long long id, void* parent) { return new QSyntaxHighlighter__Go(id, (QObject*)parent); }
void QSyntaxHighlighter_delete(void* self) { delete (QSyntaxHighlighter*)self; }

}

#include <QtWidgets/qrubberband.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
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
#include "gen_qrubberband.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QRubberBand;
static std::mutex __connsMu_QRubberBand;

extern "C" {

void* QRubberBand_base(void* self) { return (void*)static_cast<QWidget*>((QRubberBand*)self); }
void QRubberBand_chain(void* self, void** out) {
	QRubberBand* __p = (QRubberBand*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QRubberBand_fromQObject(void* o) { return (void*)qobject_cast<QRubberBand*>((QObject*)o); }
void* QRubberBand_new(int arg1) {
	return new QRubberBand((QRubberBand::Shape)(arg1));
}
void* QRubberBand_new2(int arg1, void* arg2) {
	return new QRubberBand((QRubberBand::Shape)(arg1), (QWidget*)arg2);
}
char* QRubberBand_tr(const char* s, long long* __out_len) {
	auto&& __r = QRubberBand::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QRubberBand_shape(void* self) {
	auto&& __r = ((QRubberBand*)self)->shape();
	return (int)(__r);
}
void QRubberBand_setGeometry(void* self, void* r) {
	((QRubberBand*)self)->setGeometry(*(QRect*)r);
}
void QRubberBand_setGeometry2(void* self, int x, int y, int w, int h) {
	((QRubberBand*)self)->setGeometry(x, y, w, h);
}
void QRubberBand_move(void* self, int x, int y) {
	((QRubberBand*)self)->move(x, y);
}
void QRubberBand_move2(void* self, void* p) {
	((QRubberBand*)self)->move(*(QPoint*)p);
}
void QRubberBand_resize(void* self, int w, int h) {
	((QRubberBand*)self)->resize(w, h);
}
void QRubberBand_resize2(void* self, void* s) {
	((QRubberBand*)self)->resize(*(QSize*)s);
}
void QRubberBand_setTabOrder(void* arg1, void* arg2) {
	QRubberBand::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QRubberBand_mouseGrabber() {
	auto&& __r = QRubberBand::mouseGrabber();
	return (void*)(__r);
}
void* QRubberBand_keyboardGrabber() {
	auto&& __r = QRubberBand::keyboardGrabber();
	return (void*)(__r);
}
void* QRubberBand_find(unsigned long long arg1) {
	auto&& __r = QRubberBand::find(arg1);
	return (void*)(__r);
}
void* QRubberBand_createWindowContainer(void* window) {
	auto&& __r = QRubberBand::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void* QRubberBand_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QRubberBand::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QRubberBand_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QRubberBand::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
double QRubberBand_devicePixelRatioFScale() {
	auto&& __r = QRubberBand::devicePixelRatioFScale();
	return __r;
}
int QRubberBand_encodeMetricF(int metric, double value) {
	auto&& __r = QRubberBand::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
char* QRubberBand_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QRubberBand::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRubberBand_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QRubberBand::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRubberBand_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QRubberBand::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QRubberBand_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QRubberBand::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void* QRubberBand_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QRubberBand::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QRubberBand_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QRubberBand::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QRubberBand_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::windowTitleChanged, (QRubberBand*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::windowIconChanged, (QRubberBand*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::windowIconTextChanged, (QRubberBand*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::customContextMenuRequested, (QRubberBand*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::destroyed, (QRubberBand*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::destroyed, (QRubberBand*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRubberBand*)self, &QRubberBand::objectNameChanged, (QRubberBand*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
	__conns_QRubberBand[__id] = __c;
}
void QRubberBand__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QRubberBand);
		auto __it = __conns_QRubberBand.find(id);
		if (__it == __conns_QRubberBand.end()) return;
		__c = __it->second;
		__conns_QRubberBand.erase(__it);
	}
	QObject::disconnect(__c);
}
void QRubberBand_delete(void* self) { delete (QRubberBand*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qnamespace.h>
#include "gen_qstyleoptiontitlebar.h"

extern "C" {

void* QStyleOptionTitleBar_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionTitleBar*)self); }
void QStyleOptionTitleBar_chain(void* self, void** out) {
	QStyleOptionTitleBar* __p = (QStyleOptionTitleBar*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionTitleBar_new() {
	return new QStyleOptionTitleBar();
}
void* QStyleOptionTitleBar_new2(void* other) {
	return new QStyleOptionTitleBar(*(QStyleOptionTitleBar*)other);
}
void* QStyleOptionTitleBar_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionTitleBar*)self)->operator=(*(QStyleOptionTitleBar*)arg1);
	return (void*)&(__r);
}
char* QStyleOptionTitleBar_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionTitleBar*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionTitleBar_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionTitleBar*)self)->text = text_qstr;
}
void* QStyleOptionTitleBar_icon_get(void* self) {
	auto __r = ((QStyleOptionTitleBar*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionTitleBar_icon_set(void* self, void* icon) {
	((QStyleOptionTitleBar*)self)->icon = *(QIcon*)icon;
}
int QStyleOptionTitleBar_titleBarState_get(void* self) {
	auto __r = ((QStyleOptionTitleBar*)self)->titleBarState;
	return __r;
}
void QStyleOptionTitleBar_titleBarState_set(void* self, int titleBarState) {
	((QStyleOptionTitleBar*)self)->titleBarState = titleBarState;
}
int QStyleOptionTitleBar_titleBarFlags_get(void* self) {
	auto __r = ((QStyleOptionTitleBar*)self)->titleBarFlags;
	return (int)(__r);
}
void QStyleOptionTitleBar_titleBarFlags_set(void* self, int titleBarFlags) {
	((QStyleOptionTitleBar*)self)->titleBarFlags = QFlags<Qt::WindowType>(titleBarFlags);
}
void QStyleOptionTitleBar_delete(void* self) { delete (QStyleOptionTitleBar*)self; }

}

#include <QtWidgets/qcheckbox.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtGui/qevent.h>
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
#include "gen_qcheckbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QCheckBox__Go : public QCheckBox {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QCheckBox__Go(long long id, A&&... a) : QCheckBox(std::forward<A>(a)...), __goId(id) {}
	using QCheckBox::updateMicroFocus;
	using QCheckBox::create;
	using QCheckBox::destroy;
	using QCheckBox::focusNextChild;
	using QCheckBox::focusPreviousChild;
	using QCheckBox::sender;
	using QCheckBox::senderSignalIndex;
	using QCheckBox::receivers;
	using QCheckBox::isSignalConnected;
	using QCheckBox::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QCheckBox::event(__a0); }
	bool cppgoBase_hitButton(const QPoint& __a0) { return QCheckBox::hitButton(__a0); }
	void cppgoBase_checkStateSet() { QCheckBox::checkStateSet(); }
	void cppgoBase_nextCheckState() { QCheckBox::nextCheckState(); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QCheckBox::paintEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QCheckBox::mouseMoveEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionButton* __a0) { QCheckBox::initStyleOption(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QCheckBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QCheckBox::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QCheckBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QCheckBox::mouseReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QCheckBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QCheckBox::focusOutEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QCheckBox::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QCheckBox::timerEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QCheckBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QCheckBox::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QCheckBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QCheckBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QCheckBox::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QCheckBox::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QCheckBox::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QCheckBox::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QCheckBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QCheckBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QCheckBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QCheckBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QCheckBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QCheckBox::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QCheckBox::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QCheckBox::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QCheckBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QCheckBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QCheckBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QCheckBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QCheckBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QCheckBox::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QCheckBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QCheckBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QCheckBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QCheckBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QCheckBox::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QCheckBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QCheckBox::minimumSizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QCheckBox::event(e);
	}
	bool hitButton(const QPoint& pos) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5736040353629178930LL, (void*)&pos, &__h); if (__h) return __r; }
		return QCheckBox::hitButton(pos);
	}
	void checkStateSet() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5747806663224476200LL, &__h); if (__h) return; }
		QCheckBox::checkStateSet();
	}
	void nextCheckState() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1474647914901213797LL, &__h); if (__h) return; }
		QCheckBox::nextCheckState();
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QCheckBox::paintEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QCheckBox::mouseMoveEvent(arg1);
	}
	void initStyleOption(QStyleOptionButton* option) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QCheckBox::initStyleOption(option);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::keyReleaseEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::mouseReleaseEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::focusOutEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::changeEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QCheckBox::timerEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QCheckBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QCheckBox::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QCheckBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QCheckBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QCheckBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::wheelEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QCheckBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QCheckBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QCheckBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QCheckBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QCheckBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QCheckBox::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QCheckBox::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QCheckBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QCheckBox::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QCheckBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QCheckBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QCheckBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QCheckBox;
static std::mutex __connsMu_QCheckBox;

extern "C" {

void* QCheckBox_base(void* self) { return (void*)static_cast<QAbstractButton*>((QCheckBox*)self); }
void QCheckBox_chain(void* self, void** out) {
	QCheckBox* __p = (QCheckBox*)self;
	QAbstractButton* __b0 = static_cast<QAbstractButton*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QCheckBox__set_override(void* self, int idx) { ((QCheckBox__Go*)self)->__goOn[idx] = 1; }
void* QCheckBox_fromQObject(void* o) { return (void*)qobject_cast<QCheckBox*>((QObject*)o); }
void* QCheckBox_new(long long __goid, void* parent) {
	return new QCheckBox__Go(__goid, (QWidget*)parent);
}
void* QCheckBox_new2(long long __goid) {
	return new QCheckBox__Go(__goid);
}
void* QCheckBox_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QCheckBox__Go(__goid, text_qstr);
}
void* QCheckBox_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QCheckBox__Go(__goid, text_qstr, (QWidget*)parent);
}
char* QCheckBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QCheckBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCheckBox_setTristate(void* self) {
	((QCheckBox*)self)->setTristate();
}
int QCheckBox_isTristate(void* self) {
	auto&& __r = ((QCheckBox*)self)->isTristate();
	return __r;
}
int QCheckBox_checkState(void* self) {
	auto&& __r = ((QCheckBox*)self)->checkState();
	return (int)(__r);
}
void QCheckBox_setCheckState(void* self, int state) {
	((QCheckBox*)self)->setCheckState((Qt::CheckState)(state));
}
void QCheckBox_stateChanged(void* self, int arg1) {
	((QCheckBox*)self)->stateChanged(arg1);
}
void QCheckBox_checkStateChanged(void* self, int arg1) {
	((QCheckBox*)self)->checkStateChanged((Qt::CheckState)(arg1));
}
void QCheckBox_setTabOrder(void* arg1, void* arg2) {
	QCheckBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QCheckBox_mouseGrabber() {
	auto&& __r = QCheckBox::mouseGrabber();
	return (void*)(__r);
}
void* QCheckBox_keyboardGrabber() {
	auto&& __r = QCheckBox::keyboardGrabber();
	return (void*)(__r);
}
void* QCheckBox_find(unsigned long long arg1) {
	auto&& __r = QCheckBox::find(arg1);
	return (void*)(__r);
}
void* QCheckBox_createWindowContainer(void* window) {
	auto&& __r = QCheckBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QCheckBox_updateMicroFocus(void* self) {
	((QCheckBox__Go*)self)->updateMicroFocus();
}
void QCheckBox_create(void* self) {
	((QCheckBox__Go*)self)->create();
}
void QCheckBox_destroy(void* self) {
	((QCheckBox__Go*)self)->destroy();
}
int QCheckBox_focusNextChild(void* self) {
	auto&& __r = ((QCheckBox__Go*)self)->focusNextChild();
	return __r;
}
int QCheckBox_focusPreviousChild(void* self) {
	auto&& __r = ((QCheckBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QCheckBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QCheckBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QCheckBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QCheckBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QCheckBox_sender(void* self) {
	auto&& __r = ((QCheckBox__Go*)self)->sender();
	return (void*)(__r);
}
int QCheckBox_senderSignalIndex(void* self) {
	auto&& __r = ((QCheckBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QCheckBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QCheckBox__Go*)self)->receivers(signal);
	return __r;
}
int QCheckBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QCheckBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QCheckBox_devicePixelRatioFScale() {
	auto&& __r = QCheckBox::devicePixelRatioFScale();
	return __r;
}
int QCheckBox_encodeMetricF(int metric, double value) {
	auto&& __r = QCheckBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QCheckBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QCheckBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QCheckBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QCheckBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCheckBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QCheckBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCheckBox_setTristate2(void* self, int y) {
	((QCheckBox*)self)->setTristate(y);
}
void* QCheckBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QCheckBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QCheckBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QCheckBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QCheckBox_updateMicroFocus2(void* self, int query) {
	((QCheckBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QCheckBox_create2(void* self, unsigned long long arg1) {
	((QCheckBox__Go*)self)->create(arg1);
}
void QCheckBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QCheckBox__Go*)self)->create(arg1, initializeWindow);
}
void QCheckBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QCheckBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QCheckBox_destroy2(void* self, int destroyWindow) {
	((QCheckBox__Go*)self)->destroy(destroyWindow);
}
void QCheckBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QCheckBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QCheckBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QCheckBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QCheckBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QCheckBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QCheckBox_metaObject_default(void* self) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::metaObject();
	return (void*)(__r);
}
void* QCheckBox_sizeHint_default(void* self) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::sizeHint();
	return new QSize(__r);
}
void* QCheckBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::minimumSizeHint();
	return new QSize(__r);
}
int QCheckBox_event_default(void* self, void* e) {
	auto&& __r = ((QCheckBox__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QCheckBox_hitButton_default(void* self, void* pos) {
	auto&& __r = ((QCheckBox__Go*)self)->cppgoBase_hitButton(*(QPoint*)pos);
	return __r;
}
void QCheckBox_checkStateSet_default(void* self) {
	((QCheckBox__Go*)self)->cppgoBase_checkStateSet();
}
void QCheckBox_nextCheckState_default(void* self) {
	((QCheckBox__Go*)self)->cppgoBase_nextCheckState();
}
void QCheckBox_paintEvent_default(void* self, void* arg1) {
	((QCheckBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QCheckBox_mouseMoveEvent_default(void* self, void* arg1) {
	((QCheckBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QCheckBox_initStyleOption_default(void* self, void* option) {
	((QCheckBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionButton*)option);
}
void QCheckBox_keyPressEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QCheckBox_keyReleaseEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QCheckBox_mousePressEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QCheckBox_mouseReleaseEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QCheckBox_focusInEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QCheckBox_focusOutEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QCheckBox_changeEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QCheckBox_timerEvent_default(void* self, void* e) {
	((QCheckBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
int QCheckBox_devType_default(void* self) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::devType();
	return __r;
}
void QCheckBox_setVisible_default(void* self, int visible) {
	((QCheckBox*)self)->QCheckBox::setVisible(visible);
}
int QCheckBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::heightForWidth(arg1);
	return __r;
}
int QCheckBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::hasHeightForWidth();
	return __r;
}
void* QCheckBox_paintEngine_default(void* self) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::paintEngine();
	return (void*)(__r);
}
void QCheckBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QCheckBox_wheelEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QCheckBox_enterEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QCheckBox_leaveEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QCheckBox_moveEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QCheckBox_resizeEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QCheckBox_closeEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QCheckBox_contextMenuEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QCheckBox_tabletEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QCheckBox_actionEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QCheckBox_dragEnterEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QCheckBox_dragMoveEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QCheckBox_dragLeaveEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QCheckBox_dropEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QCheckBox_showEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QCheckBox_hideEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QCheckBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QCheckBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QCheckBox_initPainter_default(void* self, void* painter) {
	((QCheckBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QCheckBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QCheckBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QCheckBox_sharedPainter_default(void* self) {
	auto&& __r = ((QCheckBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QCheckBox_inputMethodEvent_default(void* self, void* arg1) {
	((QCheckBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QCheckBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QCheckBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QCheckBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QCheckBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QCheckBox*)self)->QCheckBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QCheckBox_childEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QCheckBox_customEvent_default(void* self, void* event) {
	((QCheckBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QCheckBox_connectNotify_default(void* self, void* signal) {
	((QCheckBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QCheckBox_disconnectNotify_default(void* self, void* signal) {
	((QCheckBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QCheckBox_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::stateChanged, (QCheckBox*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_checkStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::checkStateChanged, (QCheckBox*)self, [__id](Qt::CheckState arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::pressed, (QCheckBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_released(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::released, (QCheckBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::clicked, (QCheckBox*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::clicked, (QCheckBox*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::toggled, (QCheckBox*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::windowTitleChanged, (QCheckBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::windowIconChanged, (QCheckBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::windowIconTextChanged, (QCheckBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::customContextMenuRequested, (QCheckBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::destroyed, (QCheckBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::destroyed, (QCheckBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCheckBox*)self, &QCheckBox::objectNameChanged, (QCheckBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
	__conns_QCheckBox[__id] = __c;
}
void QCheckBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QCheckBox);
		auto __it = __conns_QCheckBox.find(id);
		if (__it == __conns_QCheckBox.end()) return;
		__c = __it->second;
		__conns_QCheckBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QCheckBox__Go_new(long long id) { return new QCheckBox__Go(id); }
void QCheckBox_delete(void* self) { delete (QCheckBox*)self; }

}

#include <QtWidgets/qformlayout.h>
#include <QtWidgets/qlayout.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qlayoutitem.h>
#include <QtCore/qrect.h>
#include <QtCore/qsize.h>
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
#include "gen_qformlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFormLayout__Go : public QFormLayout {
public:
	long long __goId;
	unsigned char __goOn[32] = {};
	template<typename... A> QFormLayout__Go(long long id, A&&... a) : QFormLayout(std::forward<A>(a)...), __goId(id) {}
	using QFormLayout::widgetEvent;
	using QFormLayout::addChildLayout;
	using QFormLayout::addChildWidget;
	using QFormLayout::adoptLayout;
	using QFormLayout::alignmentRect;
	using QFormLayout::sender;
	using QFormLayout::senderSignalIndex;
	using QFormLayout::receivers;
	using QFormLayout::isSignalConnected;
	void cppgoBase_childEvent(QChildEvent* __a0) { QFormLayout::childEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFormLayout::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFormLayout::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFormLayout::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFormLayout::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFormLayout::metaObject();
	}
	int spacing() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4587105490415799682LL, &__h); if (__h) return __r; }
		return QFormLayout::spacing();
	}
	void setSpacing(int arg1) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4039479028406881496LL, arg1, &__h); if (__h) return; }
		QFormLayout::setSpacing(arg1);
	}
	void addItem(QLayoutItem* item) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6884465960736393867LL, (void*)item, &__h); if (__h) return; }
		QFormLayout::addItem(item);
	}
	QLayoutItem* itemAt(int index) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, index, &__h); if (__h) return (QLayoutItem*)__r; }
		return QFormLayout::itemAt(index);
	}
	QLayoutItem* takeAt(int index) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7549889108533486643LL, index, &__h); if (__h) return (QLayoutItem*)__r; }
		return QFormLayout::takeAt(index);
	}
	void setGeometry(const QRect& rect) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QFormLayout::setGeometry(rect);
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QFormLayout::minimumSize();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QFormLayout::sizeHint();
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QFormLayout::invalidate();
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QFormLayout::hasHeightForWidth();
	}
	int heightForWidth(int width) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, width, &__h); if (__h) return __r; }
		return QFormLayout::heightForWidth(width);
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QFormLayout::expandingDirections();
	}
	int count() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QFormLayout::count();
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QFormLayout::geometry();
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QFormLayout::maximumSize();
	}
	int indexOf(const QWidget* arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7258473138030329570LL, (void*)arg1, &__h); if (__h) return __r; }
		return QFormLayout::indexOf(arg1);
	}
	int indexOf(const QLayoutItem* arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3276435765994963516LL, (void*)arg1, &__h); if (__h) return __r; }
		return QFormLayout::indexOf(arg1);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QFormLayout::isEmpty();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QFormLayout::controlTypes();
	}
	QLayoutItem* replaceWidget(QWidget* from, QWidget* to, QFlags<Qt::FindChildOption> options) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int*))cppgo_virt_voidp_voidp_voidp_int)(__goId, -712862742159708061LL, (void*)from, (void*)to, (int)(options), &__h); if (__h) return (QLayoutItem*)__r; }
		return QFormLayout::replaceWidget(from, to, options);
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QFormLayout::layout();
	}
	void childEvent(QChildEvent* e) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)e, &__h); if (__h) return; }
		QFormLayout::childEvent(e);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFormLayout::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFormLayout::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFormLayout::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFormLayout::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFormLayout::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFormLayout::disconnectNotify(signal);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QFormLayout::minimumHeightForWidth(arg1);
	}
	QWidget* widget() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4575937485689502171LL, &__h); if (__h) return (QWidget*)__r; }
		return QFormLayout::widget();
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QFormLayout::spacerItem();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFormLayout;
static std::mutex __connsMu_QFormLayout;

extern "C" {

void* QFormLayout_base(void* self) { return (void*)static_cast<QLayout*>((QFormLayout*)self); }
void QFormLayout_chain(void* self, void** out) {
	QFormLayout* __p = (QFormLayout*)self;
	QLayout* __b0 = static_cast<QLayout*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QFormLayout__set_override(void* self, int idx) { ((QFormLayout__Go*)self)->__goOn[idx] = 1; }
void* QFormLayout_fromQObject(void* o) { return (void*)qobject_cast<QFormLayout*>((QObject*)o); }
void* QFormLayout_new(long long __goid, void* parent) {
	return new QFormLayout__Go(__goid, (QWidget*)parent);
}
void* QFormLayout_new2(long long __goid) {
	return new QFormLayout__Go(__goid);
}
char* QFormLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QFormLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFormLayout_setFieldGrowthPolicy(void* self, int policy) {
	((QFormLayout*)self)->setFieldGrowthPolicy((QFormLayout::FieldGrowthPolicy)(policy));
}
int QFormLayout_fieldGrowthPolicy(void* self) {
	auto&& __r = ((QFormLayout*)self)->fieldGrowthPolicy();
	return (int)(__r);
}
void QFormLayout_setRowWrapPolicy(void* self, int policy) {
	((QFormLayout*)self)->setRowWrapPolicy((QFormLayout::RowWrapPolicy)(policy));
}
int QFormLayout_rowWrapPolicy(void* self) {
	auto&& __r = ((QFormLayout*)self)->rowWrapPolicy();
	return (int)(__r);
}
void QFormLayout_setLabelAlignment(void* self, int alignment) {
	((QFormLayout*)self)->setLabelAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QFormLayout_labelAlignment(void* self) {
	auto&& __r = ((QFormLayout*)self)->labelAlignment();
	return (int)(__r);
}
void QFormLayout_setFormAlignment(void* self, int alignment) {
	((QFormLayout*)self)->setFormAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QFormLayout_formAlignment(void* self) {
	auto&& __r = ((QFormLayout*)self)->formAlignment();
	return (int)(__r);
}
void QFormLayout_setHorizontalSpacing(void* self, int spacing) {
	((QFormLayout*)self)->setHorizontalSpacing(spacing);
}
int QFormLayout_horizontalSpacing(void* self) {
	auto&& __r = ((QFormLayout*)self)->horizontalSpacing();
	return __r;
}
void QFormLayout_setVerticalSpacing(void* self, int spacing) {
	((QFormLayout*)self)->setVerticalSpacing(spacing);
}
int QFormLayout_verticalSpacing(void* self) {
	auto&& __r = ((QFormLayout*)self)->verticalSpacing();
	return __r;
}
void QFormLayout_addRow(void* self, void* label, void* field) {
	((QFormLayout*)self)->addRow((QWidget*)label, (QWidget*)field);
}
void QFormLayout_addRow2(void* self, void* label, void* field) {
	((QFormLayout*)self)->addRow((QWidget*)label, (QLayout*)field);
}
void QFormLayout_addRow3(void* self, const char* labelText, long long labelText_len, void* field) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	((QFormLayout*)self)->addRow(labelText_qstr, (QWidget*)field);
}
void QFormLayout_addRow4(void* self, const char* labelText, long long labelText_len, void* field) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	((QFormLayout*)self)->addRow(labelText_qstr, (QLayout*)field);
}
void QFormLayout_addRow5(void* self, void* widget) {
	((QFormLayout*)self)->addRow((QWidget*)widget);
}
void QFormLayout_addRow6(void* self, void* layout) {
	((QFormLayout*)self)->addRow((QLayout*)layout);
}
void QFormLayout_insertRow(void* self, int row, void* label, void* field) {
	((QFormLayout*)self)->insertRow(row, (QWidget*)label, (QWidget*)field);
}
void QFormLayout_insertRow2(void* self, int row, void* label, void* field) {
	((QFormLayout*)self)->insertRow(row, (QWidget*)label, (QLayout*)field);
}
void QFormLayout_insertRow3(void* self, int row, const char* labelText, long long labelText_len, void* field) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	((QFormLayout*)self)->insertRow(row, labelText_qstr, (QWidget*)field);
}
void QFormLayout_insertRow4(void* self, int row, const char* labelText, long long labelText_len, void* field) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	((QFormLayout*)self)->insertRow(row, labelText_qstr, (QLayout*)field);
}
void QFormLayout_insertRow5(void* self, int row, void* widget) {
	((QFormLayout*)self)->insertRow(row, (QWidget*)widget);
}
void QFormLayout_insertRow6(void* self, int row, void* layout) {
	((QFormLayout*)self)->insertRow(row, (QLayout*)layout);
}
void QFormLayout_removeRow(void* self, int row) {
	((QFormLayout*)self)->removeRow(row);
}
void QFormLayout_removeRow2(void* self, void* widget) {
	((QFormLayout*)self)->removeRow((QWidget*)widget);
}
void QFormLayout_removeRow3(void* self, void* layout) {
	((QFormLayout*)self)->removeRow((QLayout*)layout);
}
void* QFormLayout_takeRow(void* self, int row) {
	auto&& __r = ((QFormLayout*)self)->takeRow(row);
	return new QFormLayout::TakeRowResult(__r);
}
void* QFormLayout_takeRow2(void* self, void* widget) {
	auto&& __r = ((QFormLayout*)self)->takeRow((QWidget*)widget);
	return new QFormLayout::TakeRowResult(__r);
}
void* QFormLayout_takeRow3(void* self, void* layout) {
	auto&& __r = ((QFormLayout*)self)->takeRow((QLayout*)layout);
	return new QFormLayout::TakeRowResult(__r);
}
void QFormLayout_setItem(void* self, int row, int role, void* item) {
	((QFormLayout*)self)->setItem(row, (QFormLayout::ItemRole)(role), (QLayoutItem*)item);
}
void QFormLayout_setWidget(void* self, int row, int role, void* widget) {
	((QFormLayout*)self)->setWidget(row, (QFormLayout::ItemRole)(role), (QWidget*)widget);
}
void QFormLayout_setLayout(void* self, int row, int role, void* layout) {
	((QFormLayout*)self)->setLayout(row, (QFormLayout::ItemRole)(role), (QLayout*)layout);
}
void QFormLayout_setRowVisible(void* self, int row, int on) {
	((QFormLayout*)self)->setRowVisible(row, on);
}
void QFormLayout_setRowVisible2(void* self, void* widget, int on) {
	((QFormLayout*)self)->setRowVisible((QWidget*)widget, on);
}
void QFormLayout_setRowVisible3(void* self, void* layout, int on) {
	((QFormLayout*)self)->setRowVisible((QLayout*)layout, on);
}
int QFormLayout_isRowVisible(void* self, int row) {
	auto&& __r = ((QFormLayout*)self)->isRowVisible(row);
	return __r;
}
int QFormLayout_isRowVisible2(void* self, void* widget) {
	auto&& __r = ((QFormLayout*)self)->isRowVisible((QWidget*)widget);
	return __r;
}
int QFormLayout_isRowVisible3(void* self, void* layout) {
	auto&& __r = ((QFormLayout*)self)->isRowVisible((QLayout*)layout);
	return __r;
}
void* QFormLayout_itemAt(void* self, int row, int role) {
	auto&& __r = ((QFormLayout*)self)->itemAt(row, (QFormLayout::ItemRole)(role));
	return (void*)(__r);
}
void QFormLayout_getItemPosition(void* self, int index, int* rowPtr, int* rolePtr) {
	int rowPtr_local{};
	QFormLayout::ItemRole rolePtr_local{};
	((QFormLayout*)self)->getItemPosition(index, &rowPtr_local, &rolePtr_local);
	*rowPtr = rowPtr_local;
	*rolePtr = (int)(rolePtr_local);
}
void QFormLayout_getWidgetPosition(void* self, void* widget, int* rowPtr, int* rolePtr) {
	int rowPtr_local{};
	QFormLayout::ItemRole rolePtr_local{};
	((QFormLayout*)self)->getWidgetPosition((QWidget*)widget, &rowPtr_local, &rolePtr_local);
	*rowPtr = rowPtr_local;
	*rolePtr = (int)(rolePtr_local);
}
void QFormLayout_getLayoutPosition(void* self, void* layout, int* rowPtr, int* rolePtr) {
	int rowPtr_local{};
	QFormLayout::ItemRole rolePtr_local{};
	((QFormLayout*)self)->getLayoutPosition((QLayout*)layout, &rowPtr_local, &rolePtr_local);
	*rowPtr = rowPtr_local;
	*rolePtr = (int)(rolePtr_local);
}
void* QFormLayout_labelForField(void* self, void* field) {
	auto&& __r = ((QFormLayout*)self)->labelForField((QWidget*)field);
	return (void*)(__r);
}
void* QFormLayout_labelForField2(void* self, void* field) {
	auto&& __r = ((QFormLayout*)self)->labelForField((QLayout*)field);
	return (void*)(__r);
}
void* QFormLayout_itemAt2(void* self, int index) {
	auto&& __r = ((QFormLayout*)self)->itemAt(index);
	return (void*)(__r);
}
int QFormLayout_rowCount(void* self) {
	auto&& __r = ((QFormLayout*)self)->rowCount();
	return __r;
}
void* QFormLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QFormLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void QFormLayout_widgetEvent(void* self, void* arg1) {
	((QFormLayout__Go*)self)->widgetEvent((QEvent*)arg1);
}
void QFormLayout_addChildLayout(void* self, void* l) {
	((QFormLayout__Go*)self)->addChildLayout((QLayout*)l);
}
void QFormLayout_addChildWidget(void* self, void* w) {
	((QFormLayout__Go*)self)->addChildWidget((QWidget*)w);
}
int QFormLayout_adoptLayout(void* self, void* layout) {
	auto&& __r = ((QFormLayout__Go*)self)->adoptLayout((QLayout*)layout);
	return __r;
}
void* QFormLayout_alignmentRect(void* self, void* arg1) {
	auto&& __r = ((QFormLayout__Go*)self)->alignmentRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QFormLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFormLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFormLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFormLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFormLayout_sender(void* self) {
	auto&& __r = ((QFormLayout__Go*)self)->sender();
	return (void*)(__r);
}
int QFormLayout_senderSignalIndex(void* self) {
	auto&& __r = ((QFormLayout__Go*)self)->senderSignalIndex();
	return __r;
}
int QFormLayout_receivers(void* self, const char* signal) {
	auto&& __r = ((QFormLayout__Go*)self)->receivers(signal);
	return __r;
}
int QFormLayout_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFormLayout__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QFormLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFormLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFormLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFormLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFormLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFormLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFormLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFormLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFormLayout_itemAt_default(void* self, int row, int role) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::itemAt(row, (QFormLayout::ItemRole)(role));
	return (void*)(__r);
}
void* QFormLayout_itemAt_default2(void* self, int index) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::itemAt(index);
	return (void*)(__r);
}
void* QFormLayout_metaObject_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::metaObject();
	return (void*)(__r);
}
int QFormLayout_spacing_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::spacing();
	return __r;
}
void QFormLayout_setSpacing_default(void* self, int arg1) {
	((QFormLayout*)self)->QFormLayout::setSpacing(arg1);
}
void QFormLayout_addItem_default(void* self, void* item) {
	((QFormLayout*)self)->QFormLayout::addItem((QLayoutItem*)item);
}
void* QFormLayout_takeAt_default(void* self, int index) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::takeAt(index);
	return (void*)(__r);
}
void QFormLayout_setGeometry_default(void* self, void* rect) {
	((QFormLayout*)self)->QFormLayout::setGeometry(*(QRect*)rect);
}
void* QFormLayout_minimumSize_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::minimumSize();
	return new QSize(__r);
}
void* QFormLayout_sizeHint_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::sizeHint();
	return new QSize(__r);
}
void QFormLayout_invalidate_default(void* self) {
	((QFormLayout*)self)->QFormLayout::invalidate();
}
int QFormLayout_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::hasHeightForWidth();
	return __r;
}
int QFormLayout_heightForWidth_default(void* self, int width) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::heightForWidth(width);
	return __r;
}
int QFormLayout_expandingDirections_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::expandingDirections();
	return (int)(__r);
}
int QFormLayout_count_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::count();
	return __r;
}
void* QFormLayout_geometry_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::geometry();
	return new QRect(__r);
}
void* QFormLayout_maximumSize_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::maximumSize();
	return new QSize(__r);
}
int QFormLayout_indexOf_default(void* self, void* arg1) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::indexOf((QWidget*)arg1);
	return __r;
}
int QFormLayout_isEmpty_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::isEmpty();
	return __r;
}
int QFormLayout_controlTypes_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::controlTypes();
	return (int)(__r);
}
void* QFormLayout_replaceWidget_default(void* self, void* from, void* to, int options) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* QFormLayout_layout_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::layout();
	return (void*)(__r);
}
void QFormLayout_childEvent_default(void* self, void* e) {
	((QFormLayout__Go*)self)->cppgoBase_childEvent((QChildEvent*)e);
}
int QFormLayout_event_default(void* self, void* event) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::event((QEvent*)event);
	return __r;
}
int QFormLayout_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFormLayout_timerEvent_default(void* self, void* event) {
	((QFormLayout__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFormLayout_customEvent_default(void* self, void* event) {
	((QFormLayout__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFormLayout_connectNotify_default(void* self, void* signal) {
	((QFormLayout__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFormLayout_disconnectNotify_default(void* self, void* signal) {
	((QFormLayout__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
int QFormLayout_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::minimumHeightForWidth(arg1);
	return __r;
}
void* QFormLayout_widget_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::widget();
	return (void*)(__r);
}
void* QFormLayout_spacerItem_default(void* self) {
	auto&& __r = ((QFormLayout*)self)->QFormLayout::spacerItem();
	return (void*)(__r);
}
void QFormLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFormLayout*)self, &QFormLayout::destroyed, (QFormLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFormLayout);
	__conns_QFormLayout[__id] = __c;
}
void QFormLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFormLayout*)self, &QFormLayout::destroyed, (QFormLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFormLayout);
	__conns_QFormLayout[__id] = __c;
}
void QFormLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFormLayout*)self, &QFormLayout::objectNameChanged, (QFormLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFormLayout);
	__conns_QFormLayout[__id] = __c;
}
void QFormLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFormLayout);
		auto __it = __conns_QFormLayout.find(id);
		if (__it == __conns_QFormLayout.end()) return;
		__c = __it->second;
		__conns_QFormLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFormLayout__Go_new(long long id) { return new QFormLayout__Go(id); }
void QFormLayout_delete(void* self) { delete (QFormLayout*)self; }

}

#include <QtWidgets/qgraphicsanchorlayout.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsanchor.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsAnchor;
static std::mutex __connsMu_QGraphicsAnchor;

extern "C" {

void* QGraphicsAnchor_base(void* self) { return (void*)static_cast<QObject*>((QGraphicsAnchor*)self); }
void* QGraphicsAnchor_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsAnchor*>((QObject*)o); }
char* QGraphicsAnchor_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsAnchor::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsAnchor_setSpacing(void* self, double spacing) {
	((QGraphicsAnchor*)self)->setSpacing(spacing);
}
void QGraphicsAnchor_unsetSpacing(void* self) {
	((QGraphicsAnchor*)self)->unsetSpacing();
}
double QGraphicsAnchor_spacing(void* self) {
	auto&& __r = ((QGraphicsAnchor*)self)->spacing();
	return __r;
}
void QGraphicsAnchor_setSizePolicy(void* self, int policy) {
	((QGraphicsAnchor*)self)->setSizePolicy((QSizePolicy::Policy)(policy));
}
int QGraphicsAnchor_sizePolicy(void* self) {
	auto&& __r = ((QGraphicsAnchor*)self)->sizePolicy();
	return (int)(__r);
}
void* QGraphicsAnchor_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsAnchor::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsAnchor_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsAnchor::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QGraphicsAnchor_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsAnchor::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsAnchor_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsAnchor::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsAnchor_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsAnchor::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsAnchor_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsAnchor::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QGraphicsAnchor_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsAnchor*)self, &QGraphicsAnchor::destroyed, (QGraphicsAnchor*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsAnchor);
	__conns_QGraphicsAnchor[__id] = __c;
}
void QGraphicsAnchor_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsAnchor*)self, &QGraphicsAnchor::destroyed, (QGraphicsAnchor*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsAnchor);
	__conns_QGraphicsAnchor[__id] = __c;
}
void QGraphicsAnchor_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsAnchor*)self, &QGraphicsAnchor::objectNameChanged, (QGraphicsAnchor*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsAnchor);
	__conns_QGraphicsAnchor[__id] = __c;
}
void QGraphicsAnchor__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsAnchor);
		auto __it = __conns_QGraphicsAnchor.find(id);
		if (__it == __conns_QGraphicsAnchor.end()) return;
		__c = __it->second;
		__conns_QGraphicsAnchor.erase(__it);
	}
	QObject::disconnect(__c);
}
void QGraphicsAnchor_delete(void* self) { delete (QGraphicsAnchor*)self; }

}

#include <QtWidgets/qgraphicstransform.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qvectornd.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qmatrix4x4.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsrotation.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsRotation__Go : public QGraphicsRotation {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QGraphicsRotation__Go(long long id, A&&... a) : QGraphicsRotation(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsRotation::update;
	using QGraphicsRotation::sender;
	using QGraphicsRotation::senderSignalIndex;
	using QGraphicsRotation::receivers;
	using QGraphicsRotation::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsRotation::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsRotation::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsRotation::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsRotation::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsRotation::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsRotation::metaObject();
	}
	void applyTo(QMatrix4x4* matrix) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1580758529722784438LL, (void*)matrix, &__h); if (__h) return; }
		QGraphicsRotation::applyTo(matrix);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsRotation::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsRotation::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRotation::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRotation::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRotation::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsRotation::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsRotation::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsRotation;
static std::mutex __connsMu_QGraphicsRotation;

extern "C" {

void* QGraphicsRotation_base(void* self) { return (void*)static_cast<QGraphicsTransform*>((QGraphicsRotation*)self); }
void QGraphicsRotation_chain(void* self, void** out) {
	QGraphicsRotation* __p = (QGraphicsRotation*)self;
	QGraphicsTransform* __b0 = static_cast<QGraphicsTransform*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsRotation__set_override(void* self, int idx) { ((QGraphicsRotation__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsRotation_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsRotation*>((QObject*)o); }
void* QGraphicsRotation_new(long long __goid) {
	return new QGraphicsRotation__Go(__goid);
}
void* QGraphicsRotation_new2(long long __goid, void* parent) {
	return new QGraphicsRotation__Go(__goid, (QObject*)parent);
}
char* QGraphicsRotation_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsRotation::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsRotation_origin(void* self) {
	auto&& __r = ((QGraphicsRotation*)self)->origin();
	return new QVector3D(__r);
}
void QGraphicsRotation_setOrigin(void* self, void* point) {
	((QGraphicsRotation*)self)->setOrigin(*(QVector3D*)point);
}
double QGraphicsRotation_angle(void* self) {
	auto&& __r = ((QGraphicsRotation*)self)->angle();
	return __r;
}
void QGraphicsRotation_setAngle(void* self, double arg1) {
	((QGraphicsRotation*)self)->setAngle(arg1);
}
void* QGraphicsRotation_axis(void* self) {
	auto&& __r = ((QGraphicsRotation*)self)->axis();
	return new QVector3D(__r);
}
void QGraphicsRotation_setAxis(void* self, void* axis) {
	((QGraphicsRotation*)self)->setAxis(*(QVector3D*)axis);
}
void QGraphicsRotation_setAxis2(void* self, int axis) {
	((QGraphicsRotation*)self)->setAxis((Qt::Axis)(axis));
}
void QGraphicsRotation_originChanged(void* self) {
	((QGraphicsRotation*)self)->originChanged();
}
void QGraphicsRotation_angleChanged(void* self) {
	((QGraphicsRotation*)self)->angleChanged();
}
void QGraphicsRotation_axisChanged(void* self) {
	((QGraphicsRotation*)self)->axisChanged();
}
void QGraphicsRotation_update(void* self) {
	((QGraphicsRotation__Go*)self)->update();
}
void* QGraphicsRotation_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsRotation::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsRotation_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsRotation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsRotation_sender(void* self) {
	auto&& __r = ((QGraphicsRotation__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsRotation_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsRotation__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsRotation_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsRotation__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsRotation_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsRotation__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsRotation_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsRotation::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsRotation_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsRotation::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsRotation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsRotation::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsRotation_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsRotation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsRotation_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsRotation*)self)->QGraphicsRotation::metaObject();
	return (void*)(__r);
}
void QGraphicsRotation_applyTo_default(void* self, void* matrix) {
	((QGraphicsRotation*)self)->QGraphicsRotation::applyTo((QMatrix4x4*)matrix);
}
int QGraphicsRotation_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsRotation*)self)->QGraphicsRotation::event((QEvent*)event);
	return __r;
}
int QGraphicsRotation_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsRotation*)self)->QGraphicsRotation::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsRotation_timerEvent_default(void* self, void* event) {
	((QGraphicsRotation__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsRotation_childEvent_default(void* self, void* event) {
	((QGraphicsRotation__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsRotation_customEvent_default(void* self, void* event) {
	((QGraphicsRotation__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsRotation_connectNotify_default(void* self, void* signal) {
	((QGraphicsRotation__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsRotation_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsRotation__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsRotation_connect_originChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsRotation*)self, &QGraphicsRotation::originChanged, (QGraphicsRotation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
	__conns_QGraphicsRotation[__id] = __c;
}
void QGraphicsRotation_connect_angleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsRotation*)self, &QGraphicsRotation::angleChanged, (QGraphicsRotation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
	__conns_QGraphicsRotation[__id] = __c;
}
void QGraphicsRotation_connect_axisChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsRotation*)self, &QGraphicsRotation::axisChanged, (QGraphicsRotation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
	__conns_QGraphicsRotation[__id] = __c;
}
void QGraphicsRotation_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsRotation*)self, &QGraphicsRotation::destroyed, (QGraphicsRotation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
	__conns_QGraphicsRotation[__id] = __c;
}
void QGraphicsRotation_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsRotation*)self, &QGraphicsRotation::destroyed, (QGraphicsRotation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
	__conns_QGraphicsRotation[__id] = __c;
}
void QGraphicsRotation_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsRotation*)self, &QGraphicsRotation::objectNameChanged, (QGraphicsRotation*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
	__conns_QGraphicsRotation[__id] = __c;
}
void QGraphicsRotation__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsRotation);
		auto __it = __conns_QGraphicsRotation.find(id);
		if (__it == __conns_QGraphicsRotation.end()) return;
		__c = __it->second;
		__conns_QGraphicsRotation.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsRotation__Go_new(long long id, void* parent) { return new QGraphicsRotation__Go(id, (QObject*)parent); }
void QGraphicsRotation_delete(void* self) { delete (QGraphicsRotation*)self; }

}

#include <QtWidgets/qplaintextedit.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qvariant.h>
#include <QtCore/qurl.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qtextedit.h>
#include <QtGui/qpagedpaintdevice.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qmimedata.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtWidgets/qscrollbar.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
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
#include "gen_qplaintextedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPlainTextEdit__Go : public QPlainTextEdit {
public:
	long long __goId;
	unsigned char __goOn[58] = {};
	template<typename... A> QPlainTextEdit__Go(long long id, A&&... a) : QPlainTextEdit(std::forward<A>(a)...), __goId(id) {}
	using QPlainTextEdit::firstVisibleBlock;
	using QPlainTextEdit::contentOffset;
	using QPlainTextEdit::blockBoundingRect;
	using QPlainTextEdit::blockBoundingGeometry;
	using QPlainTextEdit::getPaintContext;
	using QPlainTextEdit::zoomInF;
	using QPlainTextEdit::setViewportMargins;
	using QPlainTextEdit::viewportMargins;
	using QPlainTextEdit::drawFrame;
	using QPlainTextEdit::updateMicroFocus;
	using QPlainTextEdit::create;
	using QPlainTextEdit::destroy;
	using QPlainTextEdit::focusNextChild;
	using QPlainTextEdit::focusPreviousChild;
	using QPlainTextEdit::sender;
	using QPlainTextEdit::senderSignalIndex;
	using QPlainTextEdit::receivers;
	using QPlainTextEdit::isSignalConnected;
	using QPlainTextEdit::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QPlainTextEdit::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPlainTextEdit::timerEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QPlainTextEdit::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QPlainTextEdit::keyReleaseEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QPlainTextEdit::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QPlainTextEdit::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QPlainTextEdit::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QPlainTextEdit::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QPlainTextEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QPlainTextEdit::mouseDoubleClickEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QPlainTextEdit::focusNextPrevChild(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QPlainTextEdit::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QPlainTextEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QPlainTextEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QPlainTextEdit::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QPlainTextEdit::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QPlainTextEdit::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QPlainTextEdit::focusOutEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QPlainTextEdit::showEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QPlainTextEdit::changeEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QPlainTextEdit::wheelEvent(__a0); }
	QMimeData* cppgoBase_createMimeDataFromSelection() { return QPlainTextEdit::createMimeDataFromSelection(); }
	bool cppgoBase_canInsertFromMimeData(const QMimeData* __a0) { return QPlainTextEdit::canInsertFromMimeData(__a0); }
	void cppgoBase_insertFromMimeData(const QMimeData* __a0) { QPlainTextEdit::insertFromMimeData(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QPlainTextEdit::inputMethodEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QPlainTextEdit::scrollContentsBy(__a0, __a1); }
	void cppgoBase_doSetTextCursor(const QTextCursor& __a0) { QPlainTextEdit::doSetTextCursor(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QPlainTextEdit::eventFilter(__a0, __a1); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QPlainTextEdit::viewportEvent(__a0); }
	QSize cppgoBase_viewportSizeHint() { return QPlainTextEdit::viewportSizeHint(); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QPlainTextEdit::initStyleOption(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QPlainTextEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QPlainTextEdit::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QPlainTextEdit::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QPlainTextEdit::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QPlainTextEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QPlainTextEdit::actionEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QPlainTextEdit::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QPlainTextEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QPlainTextEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QPlainTextEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QPlainTextEdit::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPlainTextEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPlainTextEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPlainTextEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPlainTextEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPlainTextEdit::metaObject();
	}
	QVariant loadResource(int type_, const QUrl& name) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -519616681208050567LL, type_, (void*)&name, &__h); if (__h) return *(QVariant*)__r; }
		return QPlainTextEdit::loadResource(type_, name);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery property) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(property), &__h); if (__h) return *(QVariant*)__r; }
		return QPlainTextEdit::inputMethodQuery(property);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QPlainTextEdit::event(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::timerEvent(e);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::keyReleaseEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::resizeEvent(e);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::paintEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::mousePressEvent(e);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::mouseMoveEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::mouseReleaseEvent(e);
	}
	void mouseDoubleClickEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::mouseDoubleClickEvent(e);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QPlainTextEdit::focusNextPrevChild(next);
	}
	void contextMenuEvent(QContextMenuEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::contextMenuEvent(e);
	}
	void dragEnterEvent(QDragEnterEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::dragEnterEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::dragLeaveEvent(e);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::dragMoveEvent(e);
	}
	void dropEvent(QDropEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::dropEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::focusOutEvent(e);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QPlainTextEdit::showEvent(arg1);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::changeEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QPlainTextEdit::wheelEvent(e);
	}
	QMimeData* createMimeDataFromSelection() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 3493558106800165317LL, &__h); if (__h) return (QMimeData*)__r; }
		return QPlainTextEdit::createMimeDataFromSelection();
	}
	bool canInsertFromMimeData(const QMimeData* source) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2605205276815256896LL, (void*)source, &__h); if (__h) return __r; }
		return QPlainTextEdit::canInsertFromMimeData(source);
	}
	void insertFromMimeData(const QMimeData* source) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -466553387102125056LL, (void*)source, &__h); if (__h) return; }
		QPlainTextEdit::insertFromMimeData(source);
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QPlainTextEdit::inputMethodEvent(arg1);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QPlainTextEdit::scrollContentsBy(dx, dy);
	}
	void doSetTextCursor(const QTextCursor& cursor) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4090661491944147209LL, (void*)&cursor, &__h); if (__h) return; }
		QPlainTextEdit::doSetTextCursor(cursor);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QPlainTextEdit::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QPlainTextEdit::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QPlainTextEdit::setupViewport(viewport);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QPlainTextEdit::eventFilter(arg1, arg2);
	}
	bool viewportEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)arg1, &__h); if (__h) return __r; }
		return QPlainTextEdit::viewportEvent(arg1);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QPlainTextEdit::viewportSizeHint();
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QPlainTextEdit::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QPlainTextEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QPlainTextEdit::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QPlainTextEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QPlainTextEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QPlainTextEdit::paintEngine();
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::actionEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QPlainTextEdit::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QPlainTextEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QPlainTextEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QPlainTextEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QPlainTextEdit::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPlainTextEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPlainTextEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPlainTextEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPlainTextEdit;
static std::mutex __connsMu_QPlainTextEdit;

extern "C" {

void* QPlainTextEdit_base(void* self) { return (void*)static_cast<QAbstractScrollArea*>((QPlainTextEdit*)self); }
void QPlainTextEdit_chain(void* self, void** out) {
	QPlainTextEdit* __p = (QPlainTextEdit*)self;
	QAbstractScrollArea* __b0 = static_cast<QAbstractScrollArea*>(__p); out[0] = (void*)__b0;
	QFrame* __b1 = static_cast<QFrame*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QPlainTextEdit__set_override(void* self, int idx) { ((QPlainTextEdit__Go*)self)->__goOn[idx] = 1; }
void* QPlainTextEdit_fromQObject(void* o) { return (void*)qobject_cast<QPlainTextEdit*>((QObject*)o); }
void* QPlainTextEdit_new(long long __goid, void* parent) {
	return new QPlainTextEdit__Go(__goid, (QWidget*)parent);
}
void* QPlainTextEdit_new2(long long __goid) {
	return new QPlainTextEdit__Go(__goid);
}
void* QPlainTextEdit_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QPlainTextEdit__Go(__goid, text_qstr);
}
void* QPlainTextEdit_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QPlainTextEdit__Go(__goid, text_qstr, (QWidget*)parent);
}
char* QPlainTextEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QPlainTextEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPlainTextEdit_setDocument(void* self, void* document) {
	((QPlainTextEdit*)self)->setDocument((QTextDocument*)document);
}
void* QPlainTextEdit_document(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->document();
	return (void*)(__r);
}
void QPlainTextEdit_setPlaceholderText(void* self, const char* placeholderText, long long placeholderText_len) {
	QString placeholderText_qstr = QString::fromUtf8(placeholderText, placeholderText_len);
	((QPlainTextEdit*)self)->setPlaceholderText(placeholderText_qstr);
}
char* QPlainTextEdit_placeholderText(void* self, long long* __out_len) {
	auto&& __r = ((QPlainTextEdit*)self)->placeholderText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPlainTextEdit_setTextCursor(void* self, void* cursor) {
	((QPlainTextEdit*)self)->setTextCursor(*(QTextCursor*)cursor);
}
void* QPlainTextEdit_textCursor(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->textCursor();
	return new QTextCursor(__r);
}
int QPlainTextEdit_isReadOnly(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->isReadOnly();
	return __r;
}
void QPlainTextEdit_setReadOnly(void* self, int ro) {
	((QPlainTextEdit*)self)->setReadOnly(ro);
}
void QPlainTextEdit_setTextInteractionFlags(void* self, int flags) {
	((QPlainTextEdit*)self)->setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>(flags));
}
int QPlainTextEdit_textInteractionFlags(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->textInteractionFlags();
	return (int)(__r);
}
void QPlainTextEdit_mergeCurrentCharFormat(void* self, void* modifier) {
	((QPlainTextEdit*)self)->mergeCurrentCharFormat(*(QTextCharFormat*)modifier);
}
void QPlainTextEdit_setCurrentCharFormat(void* self, void* format) {
	((QPlainTextEdit*)self)->setCurrentCharFormat(*(QTextCharFormat*)format);
}
void* QPlainTextEdit_currentCharFormat(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->currentCharFormat();
	return new QTextCharFormat(__r);
}
int QPlainTextEdit_tabChangesFocus(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->tabChangesFocus();
	return __r;
}
void QPlainTextEdit_setTabChangesFocus(void* self, int b) {
	((QPlainTextEdit*)self)->setTabChangesFocus(b);
}
void QPlainTextEdit_setDocumentTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QPlainTextEdit*)self)->setDocumentTitle(title_qstr);
}
char* QPlainTextEdit_documentTitle(void* self, long long* __out_len) {
	auto&& __r = ((QPlainTextEdit*)self)->documentTitle();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPlainTextEdit_isUndoRedoEnabled(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->isUndoRedoEnabled();
	return __r;
}
void QPlainTextEdit_setUndoRedoEnabled(void* self, int enable) {
	((QPlainTextEdit*)self)->setUndoRedoEnabled(enable);
}
void QPlainTextEdit_setMaximumBlockCount(void* self, int maximum) {
	((QPlainTextEdit*)self)->setMaximumBlockCount(maximum);
}
int QPlainTextEdit_maximumBlockCount(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->maximumBlockCount();
	return __r;
}
int QPlainTextEdit_lineWrapMode(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->lineWrapMode();
	return (int)(__r);
}
void QPlainTextEdit_setLineWrapMode(void* self, int mode) {
	((QPlainTextEdit*)self)->setLineWrapMode((QPlainTextEdit::LineWrapMode)(mode));
}
int QPlainTextEdit_wordWrapMode(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->wordWrapMode();
	return (int)(__r);
}
void QPlainTextEdit_setWordWrapMode(void* self, int policy) {
	((QPlainTextEdit*)self)->setWordWrapMode((QTextOption::WrapMode)(policy));
}
void QPlainTextEdit_setBackgroundVisible(void* self, int visible) {
	((QPlainTextEdit*)self)->setBackgroundVisible(visible);
}
int QPlainTextEdit_backgroundVisible(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->backgroundVisible();
	return __r;
}
void QPlainTextEdit_setCenterOnScroll(void* self, int enabled) {
	((QPlainTextEdit*)self)->setCenterOnScroll(enabled);
}
int QPlainTextEdit_centerOnScroll(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->centerOnScroll();
	return __r;
}
int QPlainTextEdit_find(void* self, const char* exp, long long exp_len) {
	QString exp_qstr = QString::fromUtf8(exp, exp_len);
	auto&& __r = ((QPlainTextEdit*)self)->find(exp_qstr);
	return __r;
}
int QPlainTextEdit_find2(void* self, void* exp) {
	auto&& __r = ((QPlainTextEdit*)self)->find(*(QRegularExpression*)exp);
	return __r;
}
char* QPlainTextEdit_toPlainText(void* self, long long* __out_len) {
	auto&& __r = ((QPlainTextEdit*)self)->toPlainText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPlainTextEdit_ensureCursorVisible(void* self) {
	((QPlainTextEdit*)self)->ensureCursorVisible();
}
void* QPlainTextEdit_loadResource(void* self, int type_, void* name) {
	auto&& __r = ((QPlainTextEdit*)self)->loadResource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
void* QPlainTextEdit_createStandardContextMenu(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->createStandardContextMenu();
	return (void*)(__r);
}
void* QPlainTextEdit_createStandardContextMenu2(void* self, void* position) {
	auto&& __r = ((QPlainTextEdit*)self)->createStandardContextMenu(*(QPoint*)position);
	return (void*)(__r);
}
void* QPlainTextEdit_cursorForPosition(void* self, void* pos) {
	auto&& __r = ((QPlainTextEdit*)self)->cursorForPosition(*(QPoint*)pos);
	return new QTextCursor(__r);
}
void* QPlainTextEdit_cursorRect(void* self, void* cursor) {
	auto&& __r = ((QPlainTextEdit*)self)->cursorRect(*(QTextCursor*)cursor);
	return new QRect(__r);
}
void* QPlainTextEdit_cursorRect2(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->cursorRect();
	return new QRect(__r);
}
char* QPlainTextEdit_anchorAt(void* self, void* pos, long long* __out_len) {
	auto&& __r = ((QPlainTextEdit*)self)->anchorAt(*(QPoint*)pos);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPlainTextEdit_overwriteMode(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->overwriteMode();
	return __r;
}
void QPlainTextEdit_setOverwriteMode(void* self, int overwrite) {
	((QPlainTextEdit*)self)->setOverwriteMode(overwrite);
}
double QPlainTextEdit_tabStopDistance(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->tabStopDistance();
	return __r;
}
void QPlainTextEdit_setTabStopDistance(void* self, double distance) {
	((QPlainTextEdit*)self)->setTabStopDistance(distance);
}
int QPlainTextEdit_cursorWidth(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->cursorWidth();
	return __r;
}
void QPlainTextEdit_setCursorWidth(void* self, int width) {
	((QPlainTextEdit*)self)->setCursorWidth(width);
}
void QPlainTextEdit_setExtraSelections(void* self, void** selections, long long selections_len) {
	QList<QTextEdit::ExtraSelection> selections_list; selections_list.reserve(selections_len);
	for (long long __i = 0; __i < selections_len; __i++) selections_list.append(*(QTextEdit::ExtraSelection*)selections[__i]);
	((QPlainTextEdit*)self)->setExtraSelections(selections_list);
}
void** QPlainTextEdit_extraSelections(void* self, long long* __out_len) {
	auto&& __r = ((QPlainTextEdit*)self)->extraSelections();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextEdit::ExtraSelection(__e);
	return __arr;
}
void QPlainTextEdit_moveCursor(void* self, int operation) {
	((QPlainTextEdit*)self)->moveCursor((QTextCursor::MoveOperation)(operation));
}
int QPlainTextEdit_canPaste(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->canPaste();
	return __r;
}
void QPlainTextEdit_print(void* self, void* printer) {
	((QPlainTextEdit*)self)->print((QPagedPaintDevice*)printer);
}
int QPlainTextEdit_blockCount(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->blockCount();
	return __r;
}
void* QPlainTextEdit_inputMethodQuery(void* self, int property) {
	auto&& __r = ((QPlainTextEdit*)self)->inputMethodQuery((Qt::InputMethodQuery)(property));
	return new QVariant(__r);
}
void* QPlainTextEdit_inputMethodQuery2(void* self, int query, void* argument) {
	auto&& __r = ((QPlainTextEdit*)self)->inputMethodQuery((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void QPlainTextEdit_setPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPlainTextEdit*)self)->setPlainText(text_qstr);
}
void QPlainTextEdit_cut(void* self) {
	((QPlainTextEdit*)self)->cut();
}
void QPlainTextEdit_copy(void* self) {
	((QPlainTextEdit*)self)->copy();
}
void QPlainTextEdit_paste(void* self) {
	((QPlainTextEdit*)self)->paste();
}
void QPlainTextEdit_undo(void* self) {
	((QPlainTextEdit*)self)->undo();
}
void QPlainTextEdit_redo(void* self) {
	((QPlainTextEdit*)self)->redo();
}
void QPlainTextEdit_clear(void* self) {
	((QPlainTextEdit*)self)->clear();
}
void QPlainTextEdit_selectAll(void* self) {
	((QPlainTextEdit*)self)->selectAll();
}
void QPlainTextEdit_insertPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPlainTextEdit*)self)->insertPlainText(text_qstr);
}
void QPlainTextEdit_appendPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPlainTextEdit*)self)->appendPlainText(text_qstr);
}
void QPlainTextEdit_appendHtml(void* self, const char* html, long long html_len) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	((QPlainTextEdit*)self)->appendHtml(html_qstr);
}
void QPlainTextEdit_centerCursor(void* self) {
	((QPlainTextEdit*)self)->centerCursor();
}
void QPlainTextEdit_zoomIn(void* self) {
	((QPlainTextEdit*)self)->zoomIn();
}
void QPlainTextEdit_zoomOut(void* self) {
	((QPlainTextEdit*)self)->zoomOut();
}
void QPlainTextEdit_textChanged(void* self) {
	((QPlainTextEdit*)self)->textChanged();
}
void QPlainTextEdit_undoAvailable(void* self, int b) {
	((QPlainTextEdit*)self)->undoAvailable(b);
}
void QPlainTextEdit_redoAvailable(void* self, int b) {
	((QPlainTextEdit*)self)->redoAvailable(b);
}
void QPlainTextEdit_copyAvailable(void* self, int b) {
	((QPlainTextEdit*)self)->copyAvailable(b);
}
void QPlainTextEdit_selectionChanged(void* self) {
	((QPlainTextEdit*)self)->selectionChanged();
}
void QPlainTextEdit_cursorPositionChanged(void* self) {
	((QPlainTextEdit*)self)->cursorPositionChanged();
}
void QPlainTextEdit_updateRequest(void* self, void* rect, int dy) {
	((QPlainTextEdit*)self)->updateRequest(*(QRect*)rect, dy);
}
void QPlainTextEdit_blockCountChanged(void* self, int newBlockCount) {
	((QPlainTextEdit*)self)->blockCountChanged(newBlockCount);
}
void QPlainTextEdit_modificationChanged(void* self, int arg1) {
	((QPlainTextEdit*)self)->modificationChanged(arg1);
}
void* QPlainTextEdit_firstVisibleBlock(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->firstVisibleBlock();
	return new QTextBlock(__r);
}
void* QPlainTextEdit_contentOffset(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->contentOffset();
	return new QPointF(__r);
}
void* QPlainTextEdit_blockBoundingRect(void* self, void* block) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->blockBoundingRect(*(QTextBlock*)block);
	return new QRectF(__r);
}
void* QPlainTextEdit_blockBoundingGeometry(void* self, void* block) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->blockBoundingGeometry(*(QTextBlock*)block);
	return new QRectF(__r);
}
void* QPlainTextEdit_getPaintContext(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->getPaintContext();
	return new QAbstractTextDocumentLayout::PaintContext(__r);
}
void QPlainTextEdit_zoomInF(void* self, float range_) {
	((QPlainTextEdit__Go*)self)->zoomInF(range_);
}
void QPlainTextEdit_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QPlainTextEdit__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QPlainTextEdit_setViewportMargins2(void* self, void* margins) {
	((QPlainTextEdit__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QPlainTextEdit_viewportMargins(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QPlainTextEdit_drawFrame(void* self, void* arg1) {
	((QPlainTextEdit__Go*)self)->drawFrame((QPainter*)arg1);
}
void QPlainTextEdit_setTabOrder(void* arg1, void* arg2) {
	QPlainTextEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QPlainTextEdit_mouseGrabber() {
	auto&& __r = QPlainTextEdit::mouseGrabber();
	return (void*)(__r);
}
void* QPlainTextEdit_keyboardGrabber() {
	auto&& __r = QPlainTextEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QPlainTextEdit_createWindowContainer(void* window) {
	auto&& __r = QPlainTextEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QPlainTextEdit_updateMicroFocus(void* self) {
	((QPlainTextEdit__Go*)self)->updateMicroFocus();
}
void QPlainTextEdit_create(void* self) {
	((QPlainTextEdit__Go*)self)->create();
}
void QPlainTextEdit_destroy(void* self) {
	((QPlainTextEdit__Go*)self)->destroy();
}
int QPlainTextEdit_focusNextChild(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->focusNextChild();
	return __r;
}
int QPlainTextEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QPlainTextEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPlainTextEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPlainTextEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPlainTextEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPlainTextEdit_sender(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QPlainTextEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QPlainTextEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->receivers(signal);
	return __r;
}
int QPlainTextEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QPlainTextEdit_devicePixelRatioFScale() {
	auto&& __r = QPlainTextEdit::devicePixelRatioFScale();
	return __r;
}
int QPlainTextEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QPlainTextEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QPlainTextEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QPlainTextEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPlainTextEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPlainTextEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPlainTextEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPlainTextEdit_find3(void* self, const char* exp, long long exp_len, int options) {
	QString exp_qstr = QString::fromUtf8(exp, exp_len);
	auto&& __r = ((QPlainTextEdit*)self)->find(exp_qstr, QFlags<QTextDocument::FindFlag>(options));
	return __r;
}
int QPlainTextEdit_find4(void* self, void* exp, int options) {
	auto&& __r = ((QPlainTextEdit*)self)->find(*(QRegularExpression*)exp, QFlags<QTextDocument::FindFlag>(options));
	return __r;
}
void QPlainTextEdit_moveCursor2(void* self, int operation, int mode) {
	((QPlainTextEdit*)self)->moveCursor((QTextCursor::MoveOperation)(operation), (QTextCursor::MoveMode)(mode));
}
void QPlainTextEdit_zoomIn2(void* self, int range_) {
	((QPlainTextEdit*)self)->zoomIn(range_);
}
void QPlainTextEdit_zoomOut2(void* self, int range_) {
	((QPlainTextEdit*)self)->zoomOut(range_);
}
void* QPlainTextEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QPlainTextEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QPlainTextEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QPlainTextEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QPlainTextEdit_updateMicroFocus2(void* self, int query) {
	((QPlainTextEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QPlainTextEdit_create2(void* self, unsigned long long arg1) {
	((QPlainTextEdit__Go*)self)->create(arg1);
}
void QPlainTextEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QPlainTextEdit__Go*)self)->create(arg1, initializeWindow);
}
void QPlainTextEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QPlainTextEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QPlainTextEdit_destroy2(void* self, int destroyWindow) {
	((QPlainTextEdit__Go*)self)->destroy(destroyWindow);
}
void QPlainTextEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QPlainTextEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QPlainTextEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPlainTextEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPlainTextEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPlainTextEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPlainTextEdit_loadResource_default(void* self, int type_, void* name) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::loadResource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
void* QPlainTextEdit_inputMethodQuery_default(void* self, int property) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::inputMethodQuery((Qt::InputMethodQuery)(property));
	return new QVariant(__r);
}
void* QPlainTextEdit_inputMethodQuery_default2(void* self, int query, void* argument) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::inputMethodQuery((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void* QPlainTextEdit_metaObject_default(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::metaObject();
	return (void*)(__r);
}
int QPlainTextEdit_event_default(void* self, void* e) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QPlainTextEdit_timerEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
void QPlainTextEdit_keyPressEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QPlainTextEdit_keyReleaseEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QPlainTextEdit_resizeEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QPlainTextEdit_paintEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QPlainTextEdit_mousePressEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QPlainTextEdit_mouseMoveEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QPlainTextEdit_mouseReleaseEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QPlainTextEdit_mouseDoubleClickEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)e);
}
int QPlainTextEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QPlainTextEdit_contextMenuEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)e);
}
void QPlainTextEdit_dragEnterEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)e);
}
void QPlainTextEdit_dragLeaveEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QPlainTextEdit_dragMoveEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QPlainTextEdit_dropEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)e);
}
void QPlainTextEdit_focusInEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QPlainTextEdit_focusOutEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QPlainTextEdit_showEvent_default(void* self, void* arg1) {
	((QPlainTextEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QPlainTextEdit_changeEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QPlainTextEdit_wheelEvent_default(void* self, void* e) {
	((QPlainTextEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void* QPlainTextEdit_createMimeDataFromSelection_default(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_createMimeDataFromSelection();
	return (void*)(__r);
}
int QPlainTextEdit_canInsertFromMimeData_default(void* self, void* source) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_canInsertFromMimeData((QMimeData*)source);
	return __r;
}
void QPlainTextEdit_insertFromMimeData_default(void* self, void* source) {
	((QPlainTextEdit__Go*)self)->cppgoBase_insertFromMimeData((QMimeData*)source);
}
void QPlainTextEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QPlainTextEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void QPlainTextEdit_scrollContentsBy_default(void* self, int dx, int dy) {
	((QPlainTextEdit__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QPlainTextEdit_doSetTextCursor_default(void* self, void* cursor) {
	((QPlainTextEdit__Go*)self)->cppgoBase_doSetTextCursor(*(QTextCursor*)cursor);
}
void* QPlainTextEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::minimumSizeHint();
	return new QSize(__r);
}
void* QPlainTextEdit_sizeHint_default(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::sizeHint();
	return new QSize(__r);
}
void QPlainTextEdit_setupViewport_default(void* self, void* viewport) {
	((QPlainTextEdit*)self)->QPlainTextEdit::setupViewport((QWidget*)viewport);
}
int QPlainTextEdit_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QPlainTextEdit_viewportEvent_default(void* self, void* arg1) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_viewportEvent((QEvent*)arg1);
	return __r;
}
void* QPlainTextEdit_viewportSizeHint_default(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QPlainTextEdit_initStyleOption_default(void* self, void* option) {
	((QPlainTextEdit__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QPlainTextEdit_devType_default(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::devType();
	return __r;
}
void QPlainTextEdit_setVisible_default(void* self, int visible) {
	((QPlainTextEdit*)self)->QPlainTextEdit::setVisible(visible);
}
int QPlainTextEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::heightForWidth(arg1);
	return __r;
}
int QPlainTextEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::hasHeightForWidth();
	return __r;
}
void* QPlainTextEdit_paintEngine_default(void* self) {
	auto&& __r = ((QPlainTextEdit*)self)->QPlainTextEdit::paintEngine();
	return (void*)(__r);
}
void QPlainTextEdit_enterEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QPlainTextEdit_leaveEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QPlainTextEdit_moveEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QPlainTextEdit_closeEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QPlainTextEdit_tabletEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QPlainTextEdit_actionEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QPlainTextEdit_hideEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QPlainTextEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QPlainTextEdit_initPainter_default(void* self, void* painter) {
	((QPlainTextEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QPlainTextEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QPlainTextEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QPlainTextEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QPlainTextEdit_childEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPlainTextEdit_customEvent_default(void* self, void* event) {
	((QPlainTextEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPlainTextEdit_connectNotify_default(void* self, void* signal) {
	((QPlainTextEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPlainTextEdit_disconnectNotify_default(void* self, void* signal) {
	((QPlainTextEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPlainTextEdit_connect_textChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::textChanged, (QPlainTextEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_undoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::undoAvailable, (QPlainTextEdit*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_redoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::redoAvailable, (QPlainTextEdit*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_copyAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::copyAvailable, (QPlainTextEdit*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::selectionChanged, (QPlainTextEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_cursorPositionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::cursorPositionChanged, (QPlainTextEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_updateRequest(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::updateRequest, (QPlainTextEdit*)self, [__id](const QRect& rect, int dy){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)&rect, dy);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_blockCountChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::blockCountChanged, (QPlainTextEdit*)self, [__id](int newBlockCount){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, newBlockCount);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_modificationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::modificationChanged, (QPlainTextEdit*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::windowTitleChanged, (QPlainTextEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::windowIconChanged, (QPlainTextEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::windowIconTextChanged, (QPlainTextEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::customContextMenuRequested, (QPlainTextEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::destroyed, (QPlainTextEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::destroyed, (QPlainTextEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextEdit*)self, &QPlainTextEdit::objectNameChanged, (QPlainTextEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
	__conns_QPlainTextEdit[__id] = __c;
}
void QPlainTextEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextEdit);
		auto __it = __conns_QPlainTextEdit.find(id);
		if (__it == __conns_QPlainTextEdit.end()) return;
		__c = __it->second;
		__conns_QPlainTextEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPlainTextEdit__Go_new(long long id) { return new QPlainTextEdit__Go(id); }
void QPlainTextEdit_delete(void* self) { delete (QPlainTextEdit*)self; }

}

#include <QtWidgets/qtableview.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtWidgets/qheaderview.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qregion.h>
#include <QtCore/qsize.h>
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
#include "gen_qtableview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTableView__Go : public QTableView {
public:
	long long __goId;
	unsigned char __goOn[91] = {};
	template<typename... A> QTableView__Go(long long id, A&&... a) : QTableView(std::forward<A>(a)...), __goId(id) {}
	using QTableView::rowMoved;
	using QTableView::columnMoved;
	using QTableView::rowResized;
	using QTableView::columnResized;
	using QTableView::rowCountChanged;
	using QTableView::columnCountChanged;
	using QTableView::state;
	using QTableView::scheduleDelayedItemsLayout;
	using QTableView::executeDelayedItemsLayout;
	using QTableView::setDirtyRegion;
	using QTableView::scrollDirtyRegion;
	using QTableView::dirtyRegionOffset;
	using QTableView::startAutoScroll;
	using QTableView::stopAutoScroll;
	using QTableView::doAutoScroll;
	using QTableView::dropIndicatorPosition;
	using QTableView::setViewportMargins;
	using QTableView::viewportMargins;
	using QTableView::drawFrame;
	using QTableView::updateMicroFocus;
	using QTableView::create;
	using QTableView::destroy;
	using QTableView::focusNextChild;
	using QTableView::focusPreviousChild;
	using QTableView::sender;
	using QTableView::senderSignalIndex;
	using QTableView::receivers;
	using QTableView::isSignalConnected;
	using QTableView::getDecodedMetricF;
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QTableView::scrollContentsBy(__a0, __a1); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QTableView::initViewItemOption(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTableView::paintEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTableView::timerEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTableView::dropEvent(__a0); }
	int cppgoBase_horizontalOffset() { return QTableView::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QTableView::verticalOffset(); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QTableView::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QTableView::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QTableView::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QTableView::selectedIndexes(); }
	void cppgoBase_updateGeometries() { QTableView::updateGeometries(); }
	QSize cppgoBase_viewportSizeHint() { return QTableView::viewportSizeHint(); }
	int cppgoBase_sizeHintForRow(int __a0) { return QTableView::sizeHintForRow(__a0); }
	int cppgoBase_sizeHintForColumn(int __a0) { return QTableView::sizeHintForColumn(__a0); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QTableView::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QTableView::horizontalScrollbarAction(__a0); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QTableView::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QTableView::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QTableView::currentChanged(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QTableView::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QTableView::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_updateEditorData() { QTableView::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QTableView::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QTableView::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QTableView::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QTableView::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QTableView::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QTableView::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QTableView::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QTableView::selectionCommand(__a0, __a1); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QTableView::startDrag(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTableView::focusNextPrevChild(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QTableView::event(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QTableView::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTableView::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTableView::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTableView::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTableView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTableView::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTableView::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTableView::dragLeaveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTableView::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTableView::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTableView::keyPressEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTableView::resizeEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTableView::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QTableView::eventFilter(__a0, __a1); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTableView::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTableView::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTableView::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QTableView::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTableView::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTableView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTableView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTableView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTableView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTableView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTableView::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTableView::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTableView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTableView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTableView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTableView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTableView::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTableView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTableView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTableView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTableView::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTableView::metaObject();
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QTableView::setModel(model);
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QTableView::setRootIndex(index);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QTableView::setSelectionModel(selectionModel);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QTableView::doItemsLayout();
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QTableView::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QTableView::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTableView::indexAt(p);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QTableView::scrollContentsBy(dx, dy);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QTableView::initViewItemOption(option);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QTableView::paintEvent(e);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTableView::timerEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTableView::dropEvent(event);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QTableView::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QTableView::verticalOffset();
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QTableView::moveCursor(cursorAction, modifiers);
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QTableView::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QTableView::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virtp_voidp) {
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
		return QTableView::selectedIndexes();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QTableView::updateGeometries();
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QTableView::viewportSizeHint();
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QTableView::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QTableView::sizeHintForColumn(column);
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QTableView::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QTableView::horizontalScrollbarAction(action);
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QTableView::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QTableView::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QTableView::currentChanged(current, previous);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[28] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QTableView::keyboardSearch(search);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QTableView::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QTableView::inputMethodQuery(query);
	}
	void reset() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QTableView::reset();
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QTableView::selectAll();
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTableView::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTableView::rowsAboutToBeRemoved(parent, start, end);
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QTableView::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QTableView::updateEditorGeometries();
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QTableView::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QTableView::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QTableView::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QTableView::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QTableView::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QTableView::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QTableView::selectionCommand(index, event);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QTableView::startDrag(supportedActions);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTableView::focusNextPrevChild(next);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTableView::event(event);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QTableView::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QTableView::mousePressEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QTableView::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QTableView::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QTableView::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTableView::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTableView::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTableView::dragLeaveEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTableView::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTableView::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QTableView::keyPressEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QTableView::resizeEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QTableView::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QTableView::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTableView::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTableView::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QTableView::setupViewport(viewport);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QTableView::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QTableView::contextMenuEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QTableView::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTableView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTableView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTableView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTableView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTableView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTableView::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTableView::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTableView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTableView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTableView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTableView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTableView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTableView::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QTableView::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTableView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTableView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTableView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTableView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTableView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTableView::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTableView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTableView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTableView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTableView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTableView;
static std::mutex __connsMu_QTableView;

extern "C" {

void* QTableView_base(void* self) { return (void*)static_cast<QAbstractItemView*>((QTableView*)self); }
void QTableView_chain(void* self, void** out) {
	QTableView* __p = (QTableView*)self;
	QAbstractItemView* __b0 = static_cast<QAbstractItemView*>(__p); out[0] = (void*)__b0;
	QAbstractScrollArea* __b1 = static_cast<QAbstractScrollArea*>(__b0); out[1] = (void*)__b1;
	QFrame* __b2 = static_cast<QFrame*>(__b1); out[2] = (void*)__b2;
	QWidget* __b3 = static_cast<QWidget*>(__b2); out[3] = (void*)__b3;
	QObject* __b4 = static_cast<QObject*>(__b3); out[4] = (void*)__b4;
}
void QTableView__set_override(void* self, int idx) { ((QTableView__Go*)self)->__goOn[idx] = 1; }
void* QTableView_fromQObject(void* o) { return (void*)qobject_cast<QTableView*>((QObject*)o); }
void* QTableView_new(long long __goid, void* parent) {
	return new QTableView__Go(__goid, (QWidget*)parent);
}
void* QTableView_new2(long long __goid) {
	return new QTableView__Go(__goid);
}
char* QTableView_tr(const char* s, long long* __out_len) {
	auto&& __r = QTableView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTableView_horizontalHeader(void* self) {
	auto&& __r = ((QTableView*)self)->horizontalHeader();
	return (void*)(__r);
}
void* QTableView_verticalHeader(void* self) {
	auto&& __r = ((QTableView*)self)->verticalHeader();
	return (void*)(__r);
}
void QTableView_setHorizontalHeader(void* self, void* header) {
	((QTableView*)self)->setHorizontalHeader((QHeaderView*)header);
}
void QTableView_setVerticalHeader(void* self, void* header) {
	((QTableView*)self)->setVerticalHeader((QHeaderView*)header);
}
int QTableView_rowViewportPosition(void* self, int row) {
	auto&& __r = ((QTableView*)self)->rowViewportPosition(row);
	return __r;
}
int QTableView_rowAt(void* self, int y) {
	auto&& __r = ((QTableView*)self)->rowAt(y);
	return __r;
}
void QTableView_setRowHeight(void* self, int row, int height) {
	((QTableView*)self)->setRowHeight(row, height);
}
int QTableView_rowHeight(void* self, int row) {
	auto&& __r = ((QTableView*)self)->rowHeight(row);
	return __r;
}
int QTableView_columnViewportPosition(void* self, int column) {
	auto&& __r = ((QTableView*)self)->columnViewportPosition(column);
	return __r;
}
int QTableView_columnAt(void* self, int x) {
	auto&& __r = ((QTableView*)self)->columnAt(x);
	return __r;
}
void QTableView_setColumnWidth(void* self, int column, int width) {
	((QTableView*)self)->setColumnWidth(column, width);
}
int QTableView_columnWidth(void* self, int column) {
	auto&& __r = ((QTableView*)self)->columnWidth(column);
	return __r;
}
int QTableView_isRowHidden(void* self, int row) {
	auto&& __r = ((QTableView*)self)->isRowHidden(row);
	return __r;
}
void QTableView_setRowHidden(void* self, int row, int hide) {
	((QTableView*)self)->setRowHidden(row, hide);
}
int QTableView_isColumnHidden(void* self, int column) {
	auto&& __r = ((QTableView*)self)->isColumnHidden(column);
	return __r;
}
void QTableView_setColumnHidden(void* self, int column, int hide) {
	((QTableView*)self)->setColumnHidden(column, hide);
}
void QTableView_setSortingEnabled(void* self, int enable) {
	((QTableView*)self)->setSortingEnabled(enable);
}
int QTableView_isSortingEnabled(void* self) {
	auto&& __r = ((QTableView*)self)->isSortingEnabled();
	return __r;
}
int QTableView_showGrid(void* self) {
	auto&& __r = ((QTableView*)self)->showGrid();
	return __r;
}
int QTableView_gridStyle(void* self) {
	auto&& __r = ((QTableView*)self)->gridStyle();
	return (int)(__r);
}
void QTableView_setGridStyle(void* self, int style) {
	((QTableView*)self)->setGridStyle((Qt::PenStyle)(style));
}
void QTableView_setWordWrap(void* self, int on) {
	((QTableView*)self)->setWordWrap(on);
}
int QTableView_wordWrap(void* self) {
	auto&& __r = ((QTableView*)self)->wordWrap();
	return __r;
}
void QTableView_setCornerButtonEnabled(void* self, int enable) {
	((QTableView*)self)->setCornerButtonEnabled(enable);
}
int QTableView_isCornerButtonEnabled(void* self) {
	auto&& __r = ((QTableView*)self)->isCornerButtonEnabled();
	return __r;
}
void QTableView_setSpan(void* self, int row, int column, int rowSpan, int columnSpan) {
	((QTableView*)self)->setSpan(row, column, rowSpan, columnSpan);
}
int QTableView_rowSpan(void* self, int row, int column) {
	auto&& __r = ((QTableView*)self)->rowSpan(row, column);
	return __r;
}
int QTableView_columnSpan(void* self, int row, int column) {
	auto&& __r = ((QTableView*)self)->columnSpan(row, column);
	return __r;
}
void QTableView_clearSpans(void* self) {
	((QTableView*)self)->clearSpans();
}
void QTableView_selectRow(void* self, int row) {
	((QTableView*)self)->selectRow(row);
}
void QTableView_selectColumn(void* self, int column) {
	((QTableView*)self)->selectColumn(column);
}
void QTableView_hideRow(void* self, int row) {
	((QTableView*)self)->hideRow(row);
}
void QTableView_hideColumn(void* self, int column) {
	((QTableView*)self)->hideColumn(column);
}
void QTableView_showRow(void* self, int row) {
	((QTableView*)self)->showRow(row);
}
void QTableView_showColumn(void* self, int column) {
	((QTableView*)self)->showColumn(column);
}
void QTableView_resizeRowToContents(void* self, int row) {
	((QTableView*)self)->resizeRowToContents(row);
}
void QTableView_resizeRowsToContents(void* self) {
	((QTableView*)self)->resizeRowsToContents();
}
void QTableView_resizeColumnToContents(void* self, int column) {
	((QTableView*)self)->resizeColumnToContents(column);
}
void QTableView_resizeColumnsToContents(void* self) {
	((QTableView*)self)->resizeColumnsToContents();
}
void QTableView_sortByColumn(void* self, int column, int order) {
	((QTableView*)self)->sortByColumn(column, (Qt::SortOrder)(order));
}
void QTableView_setShowGrid(void* self, int show) {
	((QTableView*)self)->setShowGrid(show);
}
void QTableView_rowMoved(void* self, int row, int oldIndex, int newIndex) {
	((QTableView__Go*)self)->rowMoved(row, oldIndex, newIndex);
}
void QTableView_columnMoved(void* self, int column, int oldIndex, int newIndex) {
	((QTableView__Go*)self)->columnMoved(column, oldIndex, newIndex);
}
void QTableView_rowResized(void* self, int row, int oldHeight, int newHeight) {
	((QTableView__Go*)self)->rowResized(row, oldHeight, newHeight);
}
void QTableView_columnResized(void* self, int column, int oldWidth, int newWidth) {
	((QTableView__Go*)self)->columnResized(column, oldWidth, newWidth);
}
void QTableView_rowCountChanged(void* self, int oldCount, int newCount) {
	((QTableView__Go*)self)->rowCountChanged(oldCount, newCount);
}
void QTableView_columnCountChanged(void* self, int oldCount, int newCount) {
	((QTableView__Go*)self)->columnCountChanged(oldCount, newCount);
}
int QTableView_state(void* self) {
	auto&& __r = ((QTableView__Go*)self)->state();
	return (int)(__r);
}
void QTableView_scheduleDelayedItemsLayout(void* self) {
	((QTableView__Go*)self)->scheduleDelayedItemsLayout();
}
void QTableView_executeDelayedItemsLayout(void* self) {
	((QTableView__Go*)self)->executeDelayedItemsLayout();
}
void QTableView_setDirtyRegion(void* self, void* region) {
	((QTableView__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QTableView_scrollDirtyRegion(void* self, int dx, int dy) {
	((QTableView__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QTableView_dirtyRegionOffset(void* self) {
	auto&& __r = ((QTableView__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QTableView_startAutoScroll(void* self) {
	((QTableView__Go*)self)->startAutoScroll();
}
void QTableView_stopAutoScroll(void* self) {
	((QTableView__Go*)self)->stopAutoScroll();
}
void QTableView_doAutoScroll(void* self) {
	((QTableView__Go*)self)->doAutoScroll();
}
int QTableView_dropIndicatorPosition(void* self) {
	auto&& __r = ((QTableView__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QTableView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QTableView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QTableView_setViewportMargins2(void* self, void* margins) {
	((QTableView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QTableView_viewportMargins(void* self) {
	auto&& __r = ((QTableView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QTableView_drawFrame(void* self, void* arg1) {
	((QTableView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QTableView_setTabOrder(void* arg1, void* arg2) {
	QTableView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTableView_mouseGrabber() {
	auto&& __r = QTableView::mouseGrabber();
	return (void*)(__r);
}
void* QTableView_keyboardGrabber() {
	auto&& __r = QTableView::keyboardGrabber();
	return (void*)(__r);
}
void* QTableView_find(unsigned long long arg1) {
	auto&& __r = QTableView::find(arg1);
	return (void*)(__r);
}
void* QTableView_createWindowContainer(void* window) {
	auto&& __r = QTableView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTableView_updateMicroFocus(void* self) {
	((QTableView__Go*)self)->updateMicroFocus();
}
void QTableView_create(void* self) {
	((QTableView__Go*)self)->create();
}
void QTableView_destroy(void* self) {
	((QTableView__Go*)self)->destroy();
}
int QTableView_focusNextChild(void* self) {
	auto&& __r = ((QTableView__Go*)self)->focusNextChild();
	return __r;
}
int QTableView_focusPreviousChild(void* self) {
	auto&& __r = ((QTableView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTableView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTableView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTableView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTableView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTableView_sender(void* self) {
	auto&& __r = ((QTableView__Go*)self)->sender();
	return (void*)(__r);
}
int QTableView_senderSignalIndex(void* self) {
	auto&& __r = ((QTableView__Go*)self)->senderSignalIndex();
	return __r;
}
int QTableView_receivers(void* self, const char* signal) {
	auto&& __r = ((QTableView__Go*)self)->receivers(signal);
	return __r;
}
int QTableView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTableView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTableView_devicePixelRatioFScale() {
	auto&& __r = QTableView::devicePixelRatioFScale();
	return __r;
}
int QTableView_encodeMetricF(int metric, double value) {
	auto&& __r = QTableView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTableView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTableView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTableView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTableView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTableView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTableView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTableView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTableView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTableView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTableView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTableView_updateMicroFocus2(void* self, int query) {
	((QTableView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTableView_create2(void* self, unsigned long long arg1) {
	((QTableView__Go*)self)->create(arg1);
}
void QTableView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTableView__Go*)self)->create(arg1, initializeWindow);
}
void QTableView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTableView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTableView_destroy2(void* self, int destroyWindow) {
	((QTableView__Go*)self)->destroy(destroyWindow);
}
void QTableView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTableView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTableView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTableView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTableView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTableView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTableView_metaObject_default(void* self) {
	auto&& __r = ((QTableView*)self)->QTableView::metaObject();
	return (void*)(__r);
}
void QTableView_setModel_default(void* self, void* model) {
	((QTableView*)self)->QTableView::setModel((QAbstractItemModel*)model);
}
void QTableView_setRootIndex_default(void* self, void* index) {
	((QTableView*)self)->QTableView::setRootIndex(*(QModelIndex*)index);
}
void QTableView_setSelectionModel_default(void* self, void* selectionModel) {
	((QTableView*)self)->QTableView::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QTableView_doItemsLayout_default(void* self) {
	((QTableView*)self)->QTableView::doItemsLayout();
}
void* QTableView_visualRect_default(void* self, void* index) {
	auto&& __r = ((QTableView*)self)->QTableView::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QTableView_scrollTo_default(void* self, void* index, int hint) {
	((QTableView*)self)->QTableView::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QTableView_indexAt_default(void* self, void* p) {
	auto&& __r = ((QTableView*)self)->QTableView::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QTableView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QTableView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QTableView_initViewItemOption_default(void* self, void* option) {
	((QTableView__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
void QTableView_paintEvent_default(void* self, void* e) {
	((QTableView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QTableView_timerEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTableView_dropEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QTableView_horizontalOffset_default(void* self) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QTableView_verticalOffset_default(void* self) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void* QTableView_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
void QTableView_setSelection_default(void* self, void* rect, int command) {
	((QTableView__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QTableView_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QTableView_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QTableView_updateGeometries_default(void* self) {
	((QTableView__Go*)self)->cppgoBase_updateGeometries();
}
void* QTableView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
int QTableView_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_sizeHintForRow(row);
	return __r;
}
int QTableView_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_sizeHintForColumn(column);
	return __r;
}
void QTableView_verticalScrollbarAction_default(void* self, int action) {
	((QTableView__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QTableView_horizontalScrollbarAction_default(void* self, int action) {
	((QTableView__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
int QTableView_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QTableView_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QTableView__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QTableView_currentChanged_default(void* self, void* current, void* previous) {
	((QTableView__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QTableView_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QTableView*)self)->QTableView::keyboardSearch(search_qstr);
}
void* QTableView_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QTableView*)self)->QTableView::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QTableView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QTableView*)self)->QTableView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QTableView_reset_default(void* self) {
	((QTableView*)self)->QTableView::reset();
}
void QTableView_selectAll_default(void* self) {
	((QTableView*)self)->QTableView::selectAll();
}
void QTableView_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QTableView__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QTableView_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QTableView__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QTableView_updateEditorData_default(void* self) {
	((QTableView__Go*)self)->cppgoBase_updateEditorData();
}
void QTableView_updateEditorGeometries_default(void* self) {
	((QTableView__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QTableView_verticalScrollbarValueChanged_default(void* self, int value) {
	((QTableView__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QTableView_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QTableView__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QTableView_closeEditor_default(void* self, void* editor, int hint) {
	((QTableView__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QTableView_commitData_default(void* self, void* editor) {
	((QTableView__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QTableView_editorDestroyed_default(void* self, void* editor) {
	((QTableView__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QTableView_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QTableView_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
void QTableView_startDrag_default(void* self, int supportedActions) {
	((QTableView__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
int QTableView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QTableView_event_default(void* self, void* event) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QTableView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QTableView_mousePressEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QTableView_mouseMoveEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QTableView_mouseReleaseEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QTableView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QTableView_dragEnterEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTableView_dragMoveEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QTableView_dragLeaveEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTableView_focusInEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QTableView_focusOutEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTableView_keyPressEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QTableView_resizeEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QTableView_inputMethodEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QTableView_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QTableView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTableView*)self)->QTableView::minimumSizeHint();
	return new QSize(__r);
}
void* QTableView_sizeHint_default(void* self) {
	auto&& __r = ((QTableView*)self)->QTableView::sizeHint();
	return new QSize(__r);
}
void QTableView_setupViewport_default(void* self, void* viewport) {
	((QTableView*)self)->QTableView::setupViewport((QWidget*)viewport);
}
void QTableView_wheelEvent_default(void* self, void* arg1) {
	((QTableView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QTableView_contextMenuEvent_default(void* self, void* arg1) {
	((QTableView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QTableView_changeEvent_default(void* self, void* arg1) {
	((QTableView__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QTableView_initStyleOption_default(void* self, void* option) {
	((QTableView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QTableView_devType_default(void* self) {
	auto&& __r = ((QTableView*)self)->QTableView::devType();
	return __r;
}
void QTableView_setVisible_default(void* self, int visible) {
	((QTableView*)self)->QTableView::setVisible(visible);
}
int QTableView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTableView*)self)->QTableView::heightForWidth(arg1);
	return __r;
}
int QTableView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTableView*)self)->QTableView::hasHeightForWidth();
	return __r;
}
void* QTableView_paintEngine_default(void* self) {
	auto&& __r = ((QTableView*)self)->QTableView::paintEngine();
	return (void*)(__r);
}
void QTableView_keyReleaseEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTableView_enterEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTableView_leaveEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTableView_moveEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTableView_closeEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTableView_tabletEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTableView_actionEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTableView_showEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QTableView_hideEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTableView_metric_default(void* self, int arg1) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTableView_initPainter_default(void* self, void* painter) {
	((QTableView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTableView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTableView_sharedPainter_default(void* self) {
	auto&& __r = ((QTableView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTableView_childEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTableView_customEvent_default(void* self, void* event) {
	((QTableView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTableView_connectNotify_default(void* self, void* signal) {
	((QTableView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTableView_disconnectNotify_default(void* self, void* signal) {
	((QTableView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTableView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::pressed, (QTableView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::clicked, (QTableView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::doubleClicked, (QTableView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::activated, (QTableView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::entered, (QTableView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::viewportEntered, (QTableView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::iconSizeChanged, (QTableView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::windowTitleChanged, (QTableView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::windowIconChanged, (QTableView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::windowIconTextChanged, (QTableView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::customContextMenuRequested, (QTableView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::destroyed, (QTableView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::destroyed, (QTableView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableView*)self, &QTableView::objectNameChanged, (QTableView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
	__conns_QTableView[__id] = __c;
}
void QTableView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTableView);
		auto __it = __conns_QTableView.find(id);
		if (__it == __conns_QTableView.end()) return;
		__c = __it->second;
		__conns_QTableView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTableView__Go_new(long long id) { return new QTableView__Go(id); }
void QTableView_delete(void* self) { delete (QTableView*)self; }

}

