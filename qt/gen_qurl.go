package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUrl_new uintptr
	addr_QUrl_new2 uintptr
	addr_QUrl_new3 uintptr
	addr_QUrl_new4 uintptr
	addr_QUrl_OperatorAssign uintptr
	addr_QUrl_OperatorAssign2 uintptr
	addr_QUrl_swap uintptr
	addr_QUrl_setUrl uintptr
	addr_QUrl_url uintptr
	addr_QUrl_toString uintptr
	addr_QUrl_toDisplayString uintptr
	addr_QUrl_toEncoded uintptr
	addr_QUrl_fromEncoded uintptr
	addr_QUrl_fromUserInput uintptr
	addr_QUrl_isValid uintptr
	addr_QUrl_errorString uintptr
	addr_QUrl_isEmpty uintptr
	addr_QUrl_clear uintptr
	addr_QUrl_setScheme uintptr
	addr_QUrl_scheme uintptr
	addr_QUrl_setAuthority uintptr
	addr_QUrl_authority uintptr
	addr_QUrl_setUserInfo uintptr
	addr_QUrl_userInfo uintptr
	addr_QUrl_setUserName uintptr
	addr_QUrl_userName uintptr
	addr_QUrl_setPassword uintptr
	addr_QUrl_password uintptr
	addr_QUrl_setHost uintptr
	addr_QUrl_host uintptr
	addr_QUrl_setPort uintptr
	addr_QUrl_port uintptr
	addr_QUrl_setPath uintptr
	addr_QUrl_path uintptr
	addr_QUrl_fileName uintptr
	addr_QUrl_hasQuery uintptr
	addr_QUrl_setQuery uintptr
	addr_QUrl_setQuery2 uintptr
	addr_QUrl_query uintptr
	addr_QUrl_hasFragment uintptr
	addr_QUrl_fragment uintptr
	addr_QUrl_setFragment uintptr
	addr_QUrl_resolved uintptr
	addr_QUrl_isRelative uintptr
	addr_QUrl_isParentOf uintptr
	addr_QUrl_isLocalFile uintptr
	addr_QUrl_fromLocalFile uintptr
	addr_QUrl_toLocalFile uintptr
	addr_QUrl_detach uintptr
	addr_QUrl_isDetached uintptr
	addr_QUrl_fromPercentEncoding uintptr
	addr_QUrl_toPercentEncoding uintptr
	addr_QUrl_fromAce uintptr
	addr_QUrl_toAce uintptr
	addr_QUrl_idnWhitelist uintptr
	addr_QUrl_toStringList uintptr
	addr_QUrl_fromStringList uintptr
	addr_QUrl_setIdnWhitelist uintptr
	addr_QUrl_setUrl2 uintptr
	addr_QUrl_fromEncoded2 uintptr
	addr_QUrl_fromUserInput2 uintptr
	addr_QUrl_fromUserInput3 uintptr
	addr_QUrl_setAuthority2 uintptr
	addr_QUrl_authority2 uintptr
	addr_QUrl_setUserInfo2 uintptr
	addr_QUrl_userInfo2 uintptr
	addr_QUrl_setUserName2 uintptr
	addr_QUrl_userName2 uintptr
	addr_QUrl_setPassword2 uintptr
	addr_QUrl_password2 uintptr
	addr_QUrl_setHost2 uintptr
	addr_QUrl_host2 uintptr
	addr_QUrl_port2 uintptr
	addr_QUrl_setPath2 uintptr
	addr_QUrl_path2 uintptr
	addr_QUrl_fileName2 uintptr
	addr_QUrl_setQuery3 uintptr
	addr_QUrl_query2 uintptr
	addr_QUrl_fragment2 uintptr
	addr_QUrl_setFragment2 uintptr
	addr_QUrl_toPercentEncoding2 uintptr
	addr_QUrl_toPercentEncoding3 uintptr
	addr_QUrl_fromAce2 uintptr
	addr_QUrl_toAce2 uintptr
	addr_QUrl_fromStringList2 uintptr
	addr_QUrl_OperatorEqual uintptr
	addr_QUrl_OperatorNotEqual uintptr
	addr_QUrl_OperatorLesser uintptr
	addr_QUrl_OperatorGreater uintptr
	addr_QUrl_OperatorLesserOrEqual uintptr
	addr_QUrl_OperatorGreaterOrEqual uintptr
	pg_QUrl_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUrl_new, "QUrl_new")
	symaddr(&addr_QUrl_new2, "QUrl_new2")
	symaddr(&addr_QUrl_new3, "QUrl_new3")
	symaddr(&addr_QUrl_new4, "QUrl_new4")
	symaddr(&addr_QUrl_OperatorAssign, "QUrl_OperatorAssign")
	symaddr(&addr_QUrl_OperatorAssign2, "QUrl_OperatorAssign2")
	symaddr(&addr_QUrl_swap, "QUrl_swap")
	symaddr(&addr_QUrl_setUrl, "QUrl_setUrl")
	symaddr(&addr_QUrl_url, "QUrl_url")
	symaddr(&addr_QUrl_toString, "QUrl_toString")
	symaddr(&addr_QUrl_toDisplayString, "QUrl_toDisplayString")
	symaddr(&addr_QUrl_toEncoded, "QUrl_toEncoded")
	symaddr(&addr_QUrl_fromEncoded, "QUrl_fromEncoded")
	symaddr(&addr_QUrl_fromUserInput, "QUrl_fromUserInput")
	symaddr(&addr_QUrl_isValid, "QUrl_isValid")
	symaddr(&addr_QUrl_errorString, "QUrl_errorString")
	symaddr(&addr_QUrl_isEmpty, "QUrl_isEmpty")
	symaddr(&addr_QUrl_clear, "QUrl_clear")
	symaddr(&addr_QUrl_setScheme, "QUrl_setScheme")
	symaddr(&addr_QUrl_scheme, "QUrl_scheme")
	symaddr(&addr_QUrl_setAuthority, "QUrl_setAuthority")
	symaddr(&addr_QUrl_authority, "QUrl_authority")
	symaddr(&addr_QUrl_setUserInfo, "QUrl_setUserInfo")
	symaddr(&addr_QUrl_userInfo, "QUrl_userInfo")
	symaddr(&addr_QUrl_setUserName, "QUrl_setUserName")
	symaddr(&addr_QUrl_userName, "QUrl_userName")
	symaddr(&addr_QUrl_setPassword, "QUrl_setPassword")
	symaddr(&addr_QUrl_password, "QUrl_password")
	symaddr(&addr_QUrl_setHost, "QUrl_setHost")
	symaddr(&addr_QUrl_host, "QUrl_host")
	symaddr(&addr_QUrl_setPort, "QUrl_setPort")
	symaddr(&addr_QUrl_port, "QUrl_port")
	symaddr(&addr_QUrl_setPath, "QUrl_setPath")
	symaddr(&addr_QUrl_path, "QUrl_path")
	symaddr(&addr_QUrl_fileName, "QUrl_fileName")
	symaddr(&addr_QUrl_hasQuery, "QUrl_hasQuery")
	symaddr(&addr_QUrl_setQuery, "QUrl_setQuery")
	symaddr(&addr_QUrl_setQuery2, "QUrl_setQuery2")
	symaddr(&addr_QUrl_query, "QUrl_query")
	symaddr(&addr_QUrl_hasFragment, "QUrl_hasFragment")
	symaddr(&addr_QUrl_fragment, "QUrl_fragment")
	symaddr(&addr_QUrl_setFragment, "QUrl_setFragment")
	symaddr(&addr_QUrl_resolved, "QUrl_resolved")
	symaddr(&addr_QUrl_isRelative, "QUrl_isRelative")
	symaddr(&addr_QUrl_isParentOf, "QUrl_isParentOf")
	symaddr(&addr_QUrl_isLocalFile, "QUrl_isLocalFile")
	symaddr(&addr_QUrl_fromLocalFile, "QUrl_fromLocalFile")
	symaddr(&addr_QUrl_toLocalFile, "QUrl_toLocalFile")
	symaddr(&addr_QUrl_detach, "QUrl_detach")
	symaddr(&addr_QUrl_isDetached, "QUrl_isDetached")
	symaddr(&addr_QUrl_fromPercentEncoding, "QUrl_fromPercentEncoding")
	symaddr(&addr_QUrl_toPercentEncoding, "QUrl_toPercentEncoding")
	symaddr(&addr_QUrl_fromAce, "QUrl_fromAce")
	symaddr(&addr_QUrl_toAce, "QUrl_toAce")
	symaddr(&addr_QUrl_idnWhitelist, "QUrl_idnWhitelist")
	symaddr(&addr_QUrl_toStringList, "QUrl_toStringList")
	symaddr(&addr_QUrl_fromStringList, "QUrl_fromStringList")
	symaddr(&addr_QUrl_setIdnWhitelist, "QUrl_setIdnWhitelist")
	symaddr(&addr_QUrl_setUrl2, "QUrl_setUrl2")
	symaddr(&addr_QUrl_fromEncoded2, "QUrl_fromEncoded2")
	symaddr(&addr_QUrl_fromUserInput2, "QUrl_fromUserInput2")
	symaddr(&addr_QUrl_fromUserInput3, "QUrl_fromUserInput3")
	symaddr(&addr_QUrl_setAuthority2, "QUrl_setAuthority2")
	symaddr(&addr_QUrl_authority2, "QUrl_authority2")
	symaddr(&addr_QUrl_setUserInfo2, "QUrl_setUserInfo2")
	symaddr(&addr_QUrl_userInfo2, "QUrl_userInfo2")
	symaddr(&addr_QUrl_setUserName2, "QUrl_setUserName2")
	symaddr(&addr_QUrl_userName2, "QUrl_userName2")
	symaddr(&addr_QUrl_setPassword2, "QUrl_setPassword2")
	symaddr(&addr_QUrl_password2, "QUrl_password2")
	symaddr(&addr_QUrl_setHost2, "QUrl_setHost2")
	symaddr(&addr_QUrl_host2, "QUrl_host2")
	symaddr(&addr_QUrl_port2, "QUrl_port2")
	symaddr(&addr_QUrl_setPath2, "QUrl_setPath2")
	symaddr(&addr_QUrl_path2, "QUrl_path2")
	symaddr(&addr_QUrl_fileName2, "QUrl_fileName2")
	symaddr(&addr_QUrl_setQuery3, "QUrl_setQuery3")
	symaddr(&addr_QUrl_query2, "QUrl_query2")
	symaddr(&addr_QUrl_fragment2, "QUrl_fragment2")
	symaddr(&addr_QUrl_setFragment2, "QUrl_setFragment2")
	symaddr(&addr_QUrl_toPercentEncoding2, "QUrl_toPercentEncoding2")
	symaddr(&addr_QUrl_toPercentEncoding3, "QUrl_toPercentEncoding3")
	symaddr(&addr_QUrl_fromAce2, "QUrl_fromAce2")
	symaddr(&addr_QUrl_toAce2, "QUrl_toAce2")
	symaddr(&addr_QUrl_fromStringList2, "QUrl_fromStringList2")
	symaddr(&addr_QUrl_OperatorEqual, "QUrl_OperatorEqual")
	symaddr(&addr_QUrl_OperatorNotEqual, "QUrl_OperatorNotEqual")
	symaddr(&addr_QUrl_OperatorLesser, "QUrl_OperatorLesser")
	symaddr(&addr_QUrl_OperatorGreater, "QUrl_OperatorGreater")
	symaddr(&addr_QUrl_OperatorLesserOrEqual, "QUrl_OperatorLesserOrEqual")
	symaddr(&addr_QUrl_OperatorGreaterOrEqual, "QUrl_OperatorGreaterOrEqual")
	reg(&pg_QUrl_delete, "QUrl_delete")
}

type QUrl struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUrl) UnsafeGoID() int64     { return this.vid }
func (this *QUrl) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUrl(h uintptr) *QUrl {
	if h == 0 {
		return nil
	}
	return &QUrl{h: h}
}

func newownedQUrl(h uintptr) *QUrl {
	o := bindQUrl(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUrl).Delete)
	}
	return o
}

func (this *QUrl) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUrl) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUrl(h unsafe.Pointer) *QUrl { return bindQUrl(uintptr(h)) }
func UnsafeNewQUrlOwned(h unsafe.Pointer) *QUrl { return newownedQUrl(uintptr(h)) }

func (this *QUrl) UnsafeRelease() { this.h = 0 }

func NewQUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QUrl_new)
	__o := newownedQUrl(__r)
	return __o
}

func NewQUrl2(copy_ *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QUrl_new2, uintptr(copy_.UnsafePointer()))
	__o := newownedQUrl(__r)
	runtime.KeepAlive(copy_)
	return __o
}

func NewQUrl3(url string) *QUrl {
	var url_p uintptr
	if len(url) > 0 {
		url_p = uintptr(unsafe.Pointer(unsafe.StringData(url)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_new3, url_p, uintptr(len(url)))
	__o := newownedQUrl(__r)
	runtime.KeepAlive(url)
	return __o
}

func NewQUrl4(url string, mode QUrl__ParsingMode) *QUrl {
	var url_p uintptr
	if len(url) > 0 {
		url_p = uintptr(unsafe.Pointer(unsafe.StringData(url)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_new4, url_p, uintptr(len(url)), uintptr(mode))
	__o := newownedQUrl(__r)
	runtime.KeepAlive(url)
	return __o
}

func (this *QUrl) OperatorAssign(copy_ *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QUrl_OperatorAssign, this.h, uintptr(copy_.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(copy_)
	return bindQUrl(__r)
}

func (this *QUrl) OperatorAssignWithUrl(url string) *QUrl {
	var url_p uintptr
	if len(url) > 0 {
		url_p = uintptr(unsafe.Pointer(unsafe.StringData(url)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_OperatorAssign2, this.h, url_p, uintptr(len(url)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
	return bindQUrl(__r)
}

func (this *QUrl) Swap(other *QUrl) {
	purego.SyscallN(addr_QUrl_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QUrl) SetUrl(url string) {
	var url_p uintptr
	if len(url) > 0 {
		url_p = uintptr(unsafe.Pointer(unsafe.StringData(url)))
	}
	purego.SyscallN(addr_QUrl_setUrl, this.h, url_p, uintptr(len(url)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func (this *QUrl) Url() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_url, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) ToDisplayString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_toDisplayString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) ToEncoded() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QUrl_toEncoded, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func QUrl_FromEncoded(input *QByteArrayView) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QUrl_fromEncoded, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(input)
	return newownedQUrl(__r)
}

func QUrl_FromUserInput(userInput string) *QUrl {
	var userInput_p uintptr
	if len(userInput) > 0 {
		userInput_p = uintptr(unsafe.Pointer(unsafe.StringData(userInput)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_fromUserInput, userInput_p, uintptr(len(userInput)))
	runtime.KeepAlive(userInput)
	return newownedQUrl(__r)
}

func (this *QUrl) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrl) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrl) Clear() {
	purego.SyscallN(addr_QUrl_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QUrl) SetScheme(scheme string) {
	var scheme_p uintptr
	if len(scheme) > 0 {
		scheme_p = uintptr(unsafe.Pointer(unsafe.StringData(scheme)))
	}
	purego.SyscallN(addr_QUrl_setScheme, this.h, scheme_p, uintptr(len(scheme)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(scheme)
}

func (this *QUrl) Scheme() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_scheme, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetAuthority(authority string) {
	var authority_p uintptr
	if len(authority) > 0 {
		authority_p = uintptr(unsafe.Pointer(unsafe.StringData(authority)))
	}
	purego.SyscallN(addr_QUrl_setAuthority, this.h, authority_p, uintptr(len(authority)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(authority)
}

func (this *QUrl) Authority() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_authority, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetUserInfo(userInfo string) {
	var userInfo_p uintptr
	if len(userInfo) > 0 {
		userInfo_p = uintptr(unsafe.Pointer(unsafe.StringData(userInfo)))
	}
	purego.SyscallN(addr_QUrl_setUserInfo, this.h, userInfo_p, uintptr(len(userInfo)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(userInfo)
}

func (this *QUrl) UserInfo() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_userInfo, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetUserName(userName string) {
	var userName_p uintptr
	if len(userName) > 0 {
		userName_p = uintptr(unsafe.Pointer(unsafe.StringData(userName)))
	}
	purego.SyscallN(addr_QUrl_setUserName, this.h, userName_p, uintptr(len(userName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(userName)
}

func (this *QUrl) UserName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_userName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetPassword(password string) {
	var password_p uintptr
	if len(password) > 0 {
		password_p = uintptr(unsafe.Pointer(unsafe.StringData(password)))
	}
	purego.SyscallN(addr_QUrl_setPassword, this.h, password_p, uintptr(len(password)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(password)
}

func (this *QUrl) Password() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_password, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetHost(host string) {
	var host_p uintptr
	if len(host) > 0 {
		host_p = uintptr(unsafe.Pointer(unsafe.StringData(host)))
	}
	purego.SyscallN(addr_QUrl_setHost, this.h, host_p, uintptr(len(host)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(host)
}

func (this *QUrl) Host() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_host, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetPort(port int) {
	purego.SyscallN(addr_QUrl_setPort, this.h, uintptr(port))
	runtime.KeepAlive(this)
}

func (this *QUrl) Port() int {
	__r0, _, _ := purego.SyscallN(addr_QUrl_port, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUrl) SetPath(path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QUrl_setPath, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QUrl) Path() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_path, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) HasQuery() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_hasQuery, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrl) SetQuery(query string) {
	var query_p uintptr
	if len(query) > 0 {
		query_p = uintptr(unsafe.Pointer(unsafe.StringData(query)))
	}
	purego.SyscallN(addr_QUrl_setQuery, this.h, query_p, uintptr(len(query)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(query)
}

func (this *QUrl) SetQueryWithQuery(query *QUrlQuery) {
	purego.SyscallN(addr_QUrl_setQuery2, this.h, uintptr(query.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(query)
}

func (this *QUrl) Query() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_query, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) HasFragment() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_hasFragment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrl) Fragment() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fragment, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetFragment(fragment string) {
	var fragment_p uintptr
	if len(fragment) > 0 {
		fragment_p = uintptr(unsafe.Pointer(unsafe.StringData(fragment)))
	}
	purego.SyscallN(addr_QUrl_setFragment, this.h, fragment_p, uintptr(len(fragment)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fragment)
}

func (this *QUrl) Resolved(relative *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QUrl_resolved, this.h, uintptr(relative.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(relative)
	return newownedQUrl(__r)
}

func (this *QUrl) IsRelative() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_isRelative, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrl) IsParentOf(url *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_isParentOf, this.h, uintptr(url.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
	return __r != 0
}

func (this *QUrl) IsLocalFile() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_isLocalFile, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QUrl_FromLocalFile(localfile string) *QUrl {
	var localfile_p uintptr
	if len(localfile) > 0 {
		localfile_p = uintptr(unsafe.Pointer(unsafe.StringData(localfile)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_fromLocalFile, localfile_p, uintptr(len(localfile)))
	runtime.KeepAlive(localfile)
	return newownedQUrl(__r)
}

func (this *QUrl) ToLocalFile() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_toLocalFile, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) Detach() {
	purego.SyscallN(addr_QUrl_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QUrl) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QUrl_FromPercentEncoding(arg1 *QByteArray) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fromPercentEncoding, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func QUrl_ToPercentEncoding(arg1 string) *QByteArray {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_toPercentEncoding, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
	return newownedQByteArray(__r)
}

func QUrl_FromAce(domain *QByteArray) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fromAce, uintptr(domain.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(domain)
	return pgStr(__r, __n)
}

func QUrl_ToAce(domain string) *QByteArray {
	var domain_p uintptr
	if len(domain) > 0 {
		domain_p = uintptr(unsafe.Pointer(unsafe.StringData(domain)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_toAce, domain_p, uintptr(len(domain)))
	runtime.KeepAlive(domain)
	return newownedQByteArray(__r)
}

func QUrl_IdnWhitelist() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_idnWhitelist, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QUrl_ToStringList(uris []*QUrl) []string {
	uris_arr := make([]uintptr, len(uris))
	for __i, __e := range uris {
		uris_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var uris_p uintptr
	if len(uris_arr) > 0 {
		uris_p = uintptr(unsafe.Pointer(&uris_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_toStringList, uris_p, uintptr(len(uris)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(uris)
	runtime.KeepAlive(uris_arr)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QUrl_FromStringList(uris []string) []*QUrl {
	uris_arr := make([]uintptr, len(uris))
	for __i, __e := range uris {
		uris_arr[__i] = pgCString(__e)
	}
	var uris_p uintptr
	if len(uris_arr) > 0 {
		uris_p = uintptr(unsafe.Pointer(&uris_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fromStringList, uris_p, uintptr(len(uris)), uintptr(unsafe.Pointer(&__n)))
	for __i := range uris_arr {
		pgFree(uris_arr[__i])
	}
	runtime.KeepAlive(uris_arr)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QUrl_SetIdnWhitelist(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QUrl_setIdnWhitelist, arg1_p, uintptr(len(arg1)))
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func (this *QUrl) SetUrl2(url string, mode QUrl__ParsingMode) {
	var url_p uintptr
	if len(url) > 0 {
		url_p = uintptr(unsafe.Pointer(unsafe.StringData(url)))
	}
	purego.SyscallN(addr_QUrl_setUrl2, this.h, url_p, uintptr(len(url)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func QUrl_FromEncoded2(input *QByteArrayView, mode QUrl__ParsingMode) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QUrl_fromEncoded2, uintptr(input.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(input)
	return newownedQUrl(__r)
}

func QUrl_FromUserInput2(userInput string, workingDirectory string) *QUrl {
	var userInput_p uintptr
	if len(userInput) > 0 {
		userInput_p = uintptr(unsafe.Pointer(unsafe.StringData(userInput)))
	}
	var workingDirectory_p uintptr
	if len(workingDirectory) > 0 {
		workingDirectory_p = uintptr(unsafe.Pointer(unsafe.StringData(workingDirectory)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_fromUserInput2, userInput_p, uintptr(len(userInput)), workingDirectory_p, uintptr(len(workingDirectory)))
	runtime.KeepAlive(userInput)
	runtime.KeepAlive(workingDirectory)
	return newownedQUrl(__r)
}

func QUrl_FromUserInput3(userInput string, workingDirectory string, options QUrl__UserInputResolutionOption) *QUrl {
	var userInput_p uintptr
	if len(userInput) > 0 {
		userInput_p = uintptr(unsafe.Pointer(unsafe.StringData(userInput)))
	}
	var workingDirectory_p uintptr
	if len(workingDirectory) > 0 {
		workingDirectory_p = uintptr(unsafe.Pointer(unsafe.StringData(workingDirectory)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_fromUserInput3, userInput_p, uintptr(len(userInput)), workingDirectory_p, uintptr(len(workingDirectory)), uintptr(options))
	runtime.KeepAlive(userInput)
	runtime.KeepAlive(workingDirectory)
	return newownedQUrl(__r)
}

func (this *QUrl) SetAuthority2(authority string, mode QUrl__ParsingMode) {
	var authority_p uintptr
	if len(authority) > 0 {
		authority_p = uintptr(unsafe.Pointer(unsafe.StringData(authority)))
	}
	purego.SyscallN(addr_QUrl_setAuthority2, this.h, authority_p, uintptr(len(authority)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(authority)
}

func (this *QUrl) AuthorityWithOptions(options QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_authority2, this.h, uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetUserInfo2(userInfo string, mode QUrl__ParsingMode) {
	var userInfo_p uintptr
	if len(userInfo) > 0 {
		userInfo_p = uintptr(unsafe.Pointer(unsafe.StringData(userInfo)))
	}
	purego.SyscallN(addr_QUrl_setUserInfo2, this.h, userInfo_p, uintptr(len(userInfo)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(userInfo)
}

func (this *QUrl) UserInfoWithOptions(options QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_userInfo2, this.h, uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetUserName2(userName string, mode QUrl__ParsingMode) {
	var userName_p uintptr
	if len(userName) > 0 {
		userName_p = uintptr(unsafe.Pointer(unsafe.StringData(userName)))
	}
	purego.SyscallN(addr_QUrl_setUserName2, this.h, userName_p, uintptr(len(userName)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(userName)
}

func (this *QUrl) UserNameWithOptions(options QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_userName2, this.h, uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetPassword2(password string, mode QUrl__ParsingMode) {
	var password_p uintptr
	if len(password) > 0 {
		password_p = uintptr(unsafe.Pointer(unsafe.StringData(password)))
	}
	purego.SyscallN(addr_QUrl_setPassword2, this.h, password_p, uintptr(len(password)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(password)
}

func (this *QUrl) PasswordWith(arg1 QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_password2, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetHost2(host string, mode QUrl__ParsingMode) {
	var host_p uintptr
	if len(host) > 0 {
		host_p = uintptr(unsafe.Pointer(unsafe.StringData(host)))
	}
	purego.SyscallN(addr_QUrl_setHost2, this.h, host_p, uintptr(len(host)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(host)
}

func (this *QUrl) HostWith(arg1 QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_host2, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) PortWithDefaultPort(defaultPort int) int {
	__r0, _, _ := purego.SyscallN(addr_QUrl_port2, this.h, uintptr(defaultPort))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUrl) SetPath2(path string, mode QUrl__ParsingMode) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QUrl_setPath2, this.h, path_p, uintptr(len(path)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QUrl) PathWithOptions(options QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_path2, this.h, uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) FileNameWithOptions(options QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fileName2, this.h, uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetQuery2(query string, mode QUrl__ParsingMode) {
	var query_p uintptr
	if len(query) > 0 {
		query_p = uintptr(unsafe.Pointer(unsafe.StringData(query)))
	}
	purego.SyscallN(addr_QUrl_setQuery3, this.h, query_p, uintptr(len(query)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(query)
}

func (this *QUrl) QueryWith(arg1 QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_query2, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) FragmentWithOptions(options QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fragment2, this.h, uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrl) SetFragment2(fragment string, mode QUrl__ParsingMode) {
	var fragment_p uintptr
	if len(fragment) > 0 {
		fragment_p = uintptr(unsafe.Pointer(unsafe.StringData(fragment)))
	}
	purego.SyscallN(addr_QUrl_setFragment2, this.h, fragment_p, uintptr(len(fragment)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fragment)
}

func QUrl_ToPercentEncoding2(arg1 string, exclude *QByteArray) *QByteArray {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_toPercentEncoding2, arg1_p, uintptr(len(arg1)), uintptr(exclude.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(exclude)
	return newownedQByteArray(__r)
}

func QUrl_ToPercentEncoding3(arg1 string, exclude *QByteArray, include *QByteArray) *QByteArray {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_toPercentEncoding3, arg1_p, uintptr(len(arg1)), uintptr(exclude.UnsafePointer()), uintptr(include.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(exclude)
	runtime.KeepAlive(include)
	return newownedQByteArray(__r)
}

func QUrl_FromAce2(domain *QByteArray, options QUrl__AceProcessingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fromAce2, uintptr(domain.UnsafePointer()), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(domain)
	return pgStr(__r, __n)
}

func QUrl_ToAce2(domain string, options QUrl__AceProcessingOption) *QByteArray {
	var domain_p uintptr
	if len(domain) > 0 {
		domain_p = uintptr(unsafe.Pointer(unsafe.StringData(domain)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrl_toAce2, domain_p, uintptr(len(domain)), uintptr(options))
	runtime.KeepAlive(domain)
	return newownedQByteArray(__r)
}

func QUrl_FromStringList2(uris []string, mode QUrl__ParsingMode) []*QUrl {
	uris_arr := make([]uintptr, len(uris))
	for __i, __e := range uris {
		uris_arr[__i] = pgCString(__e)
	}
	var uris_p uintptr
	if len(uris_arr) > 0 {
		uris_p = uintptr(unsafe.Pointer(&uris_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrl_fromStringList2, uris_p, uintptr(len(uris)), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	for __i := range uris_arr {
		pgFree(uris_arr[__i])
	}
	runtime.KeepAlive(uris_arr)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QUrl) OperatorEqual(rhs *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrl) OperatorNotEqual(rhs *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrl) OperatorLesser(rhs *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrl) OperatorGreater(rhs *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrl) OperatorLesserOrEqual(rhs *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrl) OperatorGreaterOrEqual(rhs *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrl_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrl) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUrl_delete(this.h)
	}
	this.UnsafeRelease()
}
