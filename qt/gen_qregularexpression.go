package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRegularExpression_new uintptr
	addr_QRegularExpression_new2 uintptr
	addr_QRegularExpression_new3 uintptr
	addr_QRegularExpression_new4 uintptr
	addr_QRegularExpression_patternOptions uintptr
	addr_QRegularExpression_setPatternOptions uintptr
	addr_QRegularExpression_OperatorAssign uintptr
	addr_QRegularExpression_swap uintptr
	addr_QRegularExpression_pattern uintptr
	addr_QRegularExpression_setPattern uintptr
	addr_QRegularExpression_isValid uintptr
	addr_QRegularExpression_patternErrorOffset uintptr
	addr_QRegularExpression_errorString uintptr
	addr_QRegularExpression_captureCount uintptr
	addr_QRegularExpression_namedCaptureGroups uintptr
	addr_QRegularExpression_match uintptr
	addr_QRegularExpression_match2 uintptr
	addr_QRegularExpression_matchView uintptr
	addr_QRegularExpression_globalMatch uintptr
	addr_QRegularExpression_globalMatch2 uintptr
	addr_QRegularExpression_globalMatchView uintptr
	addr_QRegularExpression_optimize uintptr
	addr_QRegularExpression_escape uintptr
	addr_QRegularExpression_wildcardToRegularExpression uintptr
	addr_QRegularExpression_anchoredPattern uintptr
	addr_QRegularExpression_escape2 uintptr
	addr_QRegularExpression_wildcardToRegularExpression2 uintptr
	addr_QRegularExpression_anchoredPattern2 uintptr
	addr_QRegularExpression_fromWildcard uintptr
	addr_QRegularExpression_match3 uintptr
	addr_QRegularExpression_match4 uintptr
	addr_QRegularExpression_match5 uintptr
	addr_QRegularExpression_match6 uintptr
	addr_QRegularExpression_match7 uintptr
	addr_QRegularExpression_match8 uintptr
	addr_QRegularExpression_matchView2 uintptr
	addr_QRegularExpression_matchView3 uintptr
	addr_QRegularExpression_matchView4 uintptr
	addr_QRegularExpression_globalMatch3 uintptr
	addr_QRegularExpression_globalMatch4 uintptr
	addr_QRegularExpression_globalMatch5 uintptr
	addr_QRegularExpression_globalMatch6 uintptr
	addr_QRegularExpression_globalMatch7 uintptr
	addr_QRegularExpression_globalMatch8 uintptr
	addr_QRegularExpression_globalMatchView2 uintptr
	addr_QRegularExpression_globalMatchView3 uintptr
	addr_QRegularExpression_globalMatchView4 uintptr
	addr_QRegularExpression_wildcardToRegularExpression3 uintptr
	addr_QRegularExpression_wildcardToRegularExpression4 uintptr
	addr_QRegularExpression_fromWildcard2 uintptr
	addr_QRegularExpression_fromWildcard3 uintptr
	addr_QRegularExpression_OperatorEqual uintptr
	addr_QRegularExpression_OperatorNotEqual uintptr
	pg_QRegularExpression_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRegularExpression_new, "QRegularExpression_new")
	symaddr(&addr_QRegularExpression_new2, "QRegularExpression_new2")
	symaddr(&addr_QRegularExpression_new3, "QRegularExpression_new3")
	symaddr(&addr_QRegularExpression_new4, "QRegularExpression_new4")
	symaddr(&addr_QRegularExpression_patternOptions, "QRegularExpression_patternOptions")
	symaddr(&addr_QRegularExpression_setPatternOptions, "QRegularExpression_setPatternOptions")
	symaddr(&addr_QRegularExpression_OperatorAssign, "QRegularExpression_OperatorAssign")
	symaddr(&addr_QRegularExpression_swap, "QRegularExpression_swap")
	symaddr(&addr_QRegularExpression_pattern, "QRegularExpression_pattern")
	symaddr(&addr_QRegularExpression_setPattern, "QRegularExpression_setPattern")
	symaddr(&addr_QRegularExpression_isValid, "QRegularExpression_isValid")
	symaddr(&addr_QRegularExpression_patternErrorOffset, "QRegularExpression_patternErrorOffset")
	symaddr(&addr_QRegularExpression_errorString, "QRegularExpression_errorString")
	symaddr(&addr_QRegularExpression_captureCount, "QRegularExpression_captureCount")
	symaddr(&addr_QRegularExpression_namedCaptureGroups, "QRegularExpression_namedCaptureGroups")
	symaddr(&addr_QRegularExpression_match, "QRegularExpression_match")
	symaddr(&addr_QRegularExpression_match2, "QRegularExpression_match2")
	symaddr(&addr_QRegularExpression_matchView, "QRegularExpression_matchView")
	symaddr(&addr_QRegularExpression_globalMatch, "QRegularExpression_globalMatch")
	symaddr(&addr_QRegularExpression_globalMatch2, "QRegularExpression_globalMatch2")
	symaddr(&addr_QRegularExpression_globalMatchView, "QRegularExpression_globalMatchView")
	symaddr(&addr_QRegularExpression_optimize, "QRegularExpression_optimize")
	symaddr(&addr_QRegularExpression_escape, "QRegularExpression_escape")
	symaddr(&addr_QRegularExpression_wildcardToRegularExpression, "QRegularExpression_wildcardToRegularExpression")
	symaddr(&addr_QRegularExpression_anchoredPattern, "QRegularExpression_anchoredPattern")
	symaddr(&addr_QRegularExpression_escape2, "QRegularExpression_escape2")
	symaddr(&addr_QRegularExpression_wildcardToRegularExpression2, "QRegularExpression_wildcardToRegularExpression2")
	symaddr(&addr_QRegularExpression_anchoredPattern2, "QRegularExpression_anchoredPattern2")
	symaddr(&addr_QRegularExpression_fromWildcard, "QRegularExpression_fromWildcard")
	symaddr(&addr_QRegularExpression_match3, "QRegularExpression_match3")
	symaddr(&addr_QRegularExpression_match4, "QRegularExpression_match4")
	symaddr(&addr_QRegularExpression_match5, "QRegularExpression_match5")
	symaddr(&addr_QRegularExpression_match6, "QRegularExpression_match6")
	symaddr(&addr_QRegularExpression_match7, "QRegularExpression_match7")
	symaddr(&addr_QRegularExpression_match8, "QRegularExpression_match8")
	symaddr(&addr_QRegularExpression_matchView2, "QRegularExpression_matchView2")
	symaddr(&addr_QRegularExpression_matchView3, "QRegularExpression_matchView3")
	symaddr(&addr_QRegularExpression_matchView4, "QRegularExpression_matchView4")
	symaddr(&addr_QRegularExpression_globalMatch3, "QRegularExpression_globalMatch3")
	symaddr(&addr_QRegularExpression_globalMatch4, "QRegularExpression_globalMatch4")
	symaddr(&addr_QRegularExpression_globalMatch5, "QRegularExpression_globalMatch5")
	symaddr(&addr_QRegularExpression_globalMatch6, "QRegularExpression_globalMatch6")
	symaddr(&addr_QRegularExpression_globalMatch7, "QRegularExpression_globalMatch7")
	symaddr(&addr_QRegularExpression_globalMatch8, "QRegularExpression_globalMatch8")
	symaddr(&addr_QRegularExpression_globalMatchView2, "QRegularExpression_globalMatchView2")
	symaddr(&addr_QRegularExpression_globalMatchView3, "QRegularExpression_globalMatchView3")
	symaddr(&addr_QRegularExpression_globalMatchView4, "QRegularExpression_globalMatchView4")
	symaddr(&addr_QRegularExpression_wildcardToRegularExpression3, "QRegularExpression_wildcardToRegularExpression3")
	symaddr(&addr_QRegularExpression_wildcardToRegularExpression4, "QRegularExpression_wildcardToRegularExpression4")
	symaddr(&addr_QRegularExpression_fromWildcard2, "QRegularExpression_fromWildcard2")
	symaddr(&addr_QRegularExpression_fromWildcard3, "QRegularExpression_fromWildcard3")
	symaddr(&addr_QRegularExpression_OperatorEqual, "QRegularExpression_OperatorEqual")
	symaddr(&addr_QRegularExpression_OperatorNotEqual, "QRegularExpression_OperatorNotEqual")
	reg(&pg_QRegularExpression_delete, "QRegularExpression_delete")
}

type QRegularExpression struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRegularExpression) UnsafeGoID() int64     { return this.vid }
func (this *QRegularExpression) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRegularExpression(h uintptr) *QRegularExpression {
	if h == 0 {
		return nil
	}
	return &QRegularExpression{h: h}
}

func newownedQRegularExpression(h uintptr) *QRegularExpression {
	o := bindQRegularExpression(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRegularExpression).Delete)
	}
	return o
}

func (this *QRegularExpression) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRegularExpression) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRegularExpression(h unsafe.Pointer) *QRegularExpression { return bindQRegularExpression(uintptr(h)) }
func UnsafeNewQRegularExpressionOwned(h unsafe.Pointer) *QRegularExpression { return newownedQRegularExpression(uintptr(h)) }

func (this *QRegularExpression) UnsafeRelease() { this.h = 0 }

func NewQRegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_new)
	__o := newownedQRegularExpression(__r)
	return __o
}

func NewQRegularExpression2(pattern string) *QRegularExpression {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_new2, pattern_p, uintptr(len(pattern)))
	__o := newownedQRegularExpression(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQRegularExpression3(re *QRegularExpression) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_new3, uintptr(re.UnsafePointer()))
	__o := newownedQRegularExpression(__r)
	runtime.KeepAlive(re)
	return __o
}

func NewQRegularExpression4(pattern string, options QRegularExpression__PatternOption) *QRegularExpression {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_new4, pattern_p, uintptr(len(pattern)), uintptr(options))
	__o := newownedQRegularExpression(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func (this *QRegularExpression) PatternOptions() QRegularExpression__PatternOption {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpression_patternOptions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRegularExpression__PatternOption(__r)
}

func (this *QRegularExpression) SetPatternOptions(options QRegularExpression__PatternOption) {
	purego.SyscallN(addr_QRegularExpression_setPatternOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QRegularExpression) OperatorAssign(re *QRegularExpression) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_OperatorAssign, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return bindQRegularExpression(__r)
}

func (this *QRegularExpression) Swap(other *QRegularExpression) {
	purego.SyscallN(addr_QRegularExpression_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QRegularExpression) Pattern() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_pattern, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRegularExpression) SetPattern(pattern string) {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	purego.SyscallN(addr_QRegularExpression_setPattern, this.h, pattern_p, uintptr(len(pattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QRegularExpression) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpression_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpression) PatternErrorOffset() int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpression_patternErrorOffset, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpression) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRegularExpression) CaptureCount() int {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpression_captureCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRegularExpression) NamedCaptureGroups() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_namedCaptureGroups, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QRegularExpression) Match(subject string) *QRegularExpressionMatch {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match, this.h, subject_p, uintptr(len(subject)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) MatchWithSubjectView(subjectView *QStringView) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match2, this.h, uintptr(subjectView.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) MatchView(subjectView *QStringView) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_matchView, this.h, uintptr(subjectView.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) GlobalMatch(subject string) *QRegularExpressionMatchIterator {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch, this.h, subject_p, uintptr(len(subject)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatchWithSubjectView(subjectView *QStringView) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch2, this.h, uintptr(subjectView.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatchView(subjectView *QStringView) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatchView, this.h, uintptr(subjectView.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) Optimize() {
	purego.SyscallN(addr_QRegularExpression_optimize, this.h)
	runtime.KeepAlive(this)
}

func QRegularExpression_Escape(str string) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_escape, str_p, uintptr(len(str)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func QRegularExpression_WildcardToRegularExpression(str string) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_wildcardToRegularExpression, str_p, uintptr(len(str)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func QRegularExpression_AnchoredPattern(expression string) string {
	var expression_p uintptr
	if len(expression) > 0 {
		expression_p = uintptr(unsafe.Pointer(unsafe.StringData(expression)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_anchoredPattern, expression_p, uintptr(len(expression)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(expression)
	return pgStr(__r, __n)
}

func QRegularExpression_EscapeWithStr(str *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_escape2, uintptr(str.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func QRegularExpression_WildcardToRegularExpressionWithStr(str *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_wildcardToRegularExpression2, uintptr(str.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func QRegularExpression_AnchoredPatternWithExpression(expression *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_anchoredPattern2, uintptr(expression.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(expression)
	return pgStr(__r, __n)
}

func QRegularExpression_FromWildcard(pattern *QStringView) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_fromWildcard, uintptr(pattern.UnsafePointer()))
	runtime.KeepAlive(pattern)
	return newownedQRegularExpression(__r)
}

func (this *QRegularExpression) Match2(subject string, offset int64) *QRegularExpressionMatch {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match3, this.h, subject_p, uintptr(len(subject)), uintptr(offset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) Match3(subject string, offset int64, matchType QRegularExpression__MatchType) *QRegularExpressionMatch {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match4, this.h, subject_p, uintptr(len(subject)), uintptr(offset), uintptr(matchType))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) Match4(subject string, offset int64, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match5, this.h, subject_p, uintptr(len(subject)), uintptr(offset), uintptr(matchType), uintptr(matchOptions))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) Match5(subjectView *QStringView, offset int64) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match6, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) Match6(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match7, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) Match7(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_match8, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType), uintptr(matchOptions))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) MatchView2(subjectView *QStringView, offset int64) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_matchView2, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) MatchView3(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_matchView3, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) MatchView4(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_matchView4, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType), uintptr(matchOptions))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpression) GlobalMatch2(subject string, offset int64) *QRegularExpressionMatchIterator {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch3, this.h, subject_p, uintptr(len(subject)), uintptr(offset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatch3(subject string, offset int64, matchType QRegularExpression__MatchType) *QRegularExpressionMatchIterator {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch4, this.h, subject_p, uintptr(len(subject)), uintptr(offset), uintptr(matchType))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatch4(subject string, offset int64, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatchIterator {
	var subject_p uintptr
	if len(subject) > 0 {
		subject_p = uintptr(unsafe.Pointer(unsafe.StringData(subject)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch5, this.h, subject_p, uintptr(len(subject)), uintptr(offset), uintptr(matchType), uintptr(matchOptions))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subject)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatch5(subjectView *QStringView, offset int64) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch6, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatch6(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch7, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatch7(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatch8, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType), uintptr(matchOptions))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatchView2(subjectView *QStringView, offset int64) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatchView2, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatchView3(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatchView3, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpression) GlobalMatchView4(subjectView *QStringView, offset int64, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_globalMatchView4, this.h, uintptr(subjectView.UnsafePointer()), uintptr(offset), uintptr(matchType), uintptr(matchOptions))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subjectView)
	return newownedQRegularExpressionMatchIterator(__r)
}

func QRegularExpression_WildcardToRegularExpression2(str string, options QRegularExpression__WildcardConversionOption) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_wildcardToRegularExpression3, str_p, uintptr(len(str)), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func QRegularExpression_WildcardToRegularExpression3(str *QStringView, options QRegularExpression__WildcardConversionOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_wildcardToRegularExpression4, uintptr(str.UnsafePointer()), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func QRegularExpression_FromWildcard2(pattern *QStringView, cs CaseSensitivity) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_fromWildcard2, uintptr(pattern.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(pattern)
	return newownedQRegularExpression(__r)
}

func QRegularExpression_FromWildcard3(pattern *QStringView, cs CaseSensitivity, options QRegularExpression__WildcardConversionOption) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpression_fromWildcard3, uintptr(pattern.UnsafePointer()), uintptr(cs), uintptr(options))
	runtime.KeepAlive(pattern)
	return newownedQRegularExpression(__r)
}

func (this *QRegularExpression) OperatorEqual(rhs *QRegularExpression) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpression_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRegularExpression) OperatorNotEqual(rhs *QRegularExpression) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpression_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRegularExpression) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRegularExpression_delete(this.h)
	}
	this.UnsafeRelease()
}
