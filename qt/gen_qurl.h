#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUrl_new();
void* QUrl_new2(void* copy_);
void* QUrl_new3(const char* url, long long url_len);
void* QUrl_new4(const char* url, long long url_len, int mode);
void* QUrl_OperatorAssign(void* self, void* copy_);
void* QUrl_OperatorAssign2(void* self, const char* url, long long url_len);
void QUrl_swap(void* self, void* other);
void QUrl_setUrl(void* self, const char* url, long long url_len);
char* QUrl_url(void* self, long long* __out_len);
char* QUrl_toString(void* self, long long* __out_len);
char* QUrl_toDisplayString(void* self, long long* __out_len);
void* QUrl_toEncoded(void* self);
void* QUrl_fromEncoded(void* input);
void* QUrl_fromUserInput(const char* userInput, long long userInput_len);
int QUrl_isValid(void* self);
char* QUrl_errorString(void* self, long long* __out_len);
int QUrl_isEmpty(void* self);
void QUrl_clear(void* self);
void QUrl_setScheme(void* self, const char* scheme, long long scheme_len);
char* QUrl_scheme(void* self, long long* __out_len);
void QUrl_setAuthority(void* self, const char* authority, long long authority_len);
char* QUrl_authority(void* self, long long* __out_len);
void QUrl_setUserInfo(void* self, const char* userInfo, long long userInfo_len);
char* QUrl_userInfo(void* self, long long* __out_len);
void QUrl_setUserName(void* self, const char* userName, long long userName_len);
char* QUrl_userName(void* self, long long* __out_len);
void QUrl_setPassword(void* self, const char* password, long long password_len);
char* QUrl_password(void* self, long long* __out_len);
void QUrl_setHost(void* self, const char* host, long long host_len);
char* QUrl_host(void* self, long long* __out_len);
void QUrl_setPort(void* self, int port);
int QUrl_port(void* self);
void QUrl_setPath(void* self, const char* path, long long path_len);
char* QUrl_path(void* self, long long* __out_len);
char* QUrl_fileName(void* self, long long* __out_len);
int QUrl_hasQuery(void* self);
void QUrl_setQuery(void* self, const char* query, long long query_len);
void QUrl_setQuery2(void* self, void* query);
char* QUrl_query(void* self, long long* __out_len);
int QUrl_hasFragment(void* self);
char* QUrl_fragment(void* self, long long* __out_len);
void QUrl_setFragment(void* self, const char* fragment, long long fragment_len);
void* QUrl_resolved(void* self, void* relative);
int QUrl_isRelative(void* self);
int QUrl_isParentOf(void* self, void* url);
int QUrl_isLocalFile(void* self);
void* QUrl_fromLocalFile(const char* localfile, long long localfile_len);
char* QUrl_toLocalFile(void* self, long long* __out_len);
void QUrl_detach(void* self);
int QUrl_isDetached(void* self);
char* QUrl_fromPercentEncoding(void* arg1, long long* __out_len);
void* QUrl_toPercentEncoding(const char* arg1, long long arg1_len);
char* QUrl_fromAce(void* domain, long long* __out_len);
void* QUrl_toAce(const char* domain, long long domain_len);
char** QUrl_idnWhitelist(long long* __out_len);
char** QUrl_toStringList(void** uris, long long uris_len, long long* __out_len);
void** QUrl_fromStringList(char** uris, long long uris_len, long long* __out_len);
void QUrl_setIdnWhitelist(char** arg1, long long arg1_len);
void QUrl_setUrl2(void* self, const char* url, long long url_len, int mode);
void* QUrl_fromEncoded2(void* input, int mode);
void* QUrl_fromUserInput2(const char* userInput, long long userInput_len, const char* workingDirectory, long long workingDirectory_len);
void* QUrl_fromUserInput3(const char* userInput, long long userInput_len, const char* workingDirectory, long long workingDirectory_len, int options);
void QUrl_setAuthority2(void* self, const char* authority, long long authority_len, int mode);
char* QUrl_authority2(void* self, int options, long long* __out_len);
void QUrl_setUserInfo2(void* self, const char* userInfo, long long userInfo_len, int mode);
char* QUrl_userInfo2(void* self, int options, long long* __out_len);
void QUrl_setUserName2(void* self, const char* userName, long long userName_len, int mode);
char* QUrl_userName2(void* self, int options, long long* __out_len);
void QUrl_setPassword2(void* self, const char* password, long long password_len, int mode);
char* QUrl_password2(void* self, int arg1, long long* __out_len);
void QUrl_setHost2(void* self, const char* host, long long host_len, int mode);
char* QUrl_host2(void* self, int arg1, long long* __out_len);
int QUrl_port2(void* self, int defaultPort);
void QUrl_setPath2(void* self, const char* path, long long path_len, int mode);
char* QUrl_path2(void* self, int options, long long* __out_len);
char* QUrl_fileName2(void* self, int options, long long* __out_len);
void QUrl_setQuery3(void* self, const char* query, long long query_len, int mode);
char* QUrl_query2(void* self, int arg1, long long* __out_len);
char* QUrl_fragment2(void* self, int options, long long* __out_len);
void QUrl_setFragment2(void* self, const char* fragment, long long fragment_len, int mode);
void* QUrl_toPercentEncoding2(const char* arg1, long long arg1_len, void* exclude);
void* QUrl_toPercentEncoding3(const char* arg1, long long arg1_len, void* exclude, void* include);
char* QUrl_fromAce2(void* domain, int options, long long* __out_len);
void* QUrl_toAce2(const char* domain, long long domain_len, int options);
void** QUrl_fromStringList2(char** uris, long long uris_len, int mode, long long* __out_len);
int QUrl_OperatorEqual(void* self, void* rhs);
int QUrl_OperatorNotEqual(void* self, void* rhs);
int QUrl_OperatorLesser(void* self, void* rhs);
int QUrl_OperatorGreater(void* self, void* rhs);
int QUrl_OperatorLesserOrEqual(void* self, void* rhs);
int QUrl_OperatorGreaterOrEqual(void* self, void* rhs);
void QUrl_delete(void* self);

#ifdef __cplusplus
}
#endif
