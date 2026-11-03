package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRegularExpressionMatch_new uintptr
	addr_QRegularExpressionMatch_new2 uintptr
	addr_QRegularExpressionMatch_OperatorAssign uintptr
	addr_QRegularExpressionMatch_swap uintptr
	addr_QRegularExpressionMatch_regularExpression uintptr
	addr_QRegularExpressionMatch_matchType uintptr
	addr_QRegularExpressionMatch_matchOptions uintptr
	addr_QRegularExpressionMatch_hasMatch uintptr
	addr_QRegularExpressionMatch_hasPartialMatch uintptr
	addr_QRegularExpressionMatch_isValid uintptr
	addr_QRegularExpressionMatch_lastCapturedIndex uintptr
	addr_QRegularExpressionMatch_hasCaptured uintptr
	addr_QRegularExpressionMatch_hasCaptured2 uintptr
	addr_QRegularExpressionMatch_captured uintptr
	addr_QRegularExpressionMatch_capturedView uintptr
	addr_QRegularExpressionMatch_captured2 uintptr
	addr_QRegularExpressionMatch_capturedView2 uintptr
	addr_QRegularExpressionMatch_capturedTexts uintptr
	addr_QRegularExpressionMatch_capturedStart uintptr
	addr_QRegularExpressionMatch_capturedLength uintptr
	addr_QRegularExpressionMatch_capturedEnd uintptr
	addr_QRegularExpressionMatch_capturedStart2 uintptr
	addr_QRegularExpressionMatch_capturedLength2 uintptr
	addr_QRegularExpressionMatch_capturedEnd2 uintptr
	addr_QRegularExpressionMatch_captured3 uintptr
	addr_QRegularExpressionMatch_capturedView3 uintptr
	addr_QRegularExpressionMatch_capturedStart3 uintptr
	addr_QRegularExpressionMatch_capturedLength3 uintptr
	addr_QRegularExpressionMatch_capturedEnd3 uintptr
	pg_QRegularExpressionMatch_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRegularExpressionMatch_new, "QRegularExpressionMatch_new")
	symaddr(&addr_QRegularExpressionMatch_new2, "QRegularExpressionMatch_new2")
	symaddr(&addr_QRegularExpressionMatch_OperatorAssign, "QRegularExpressionMatch_OperatorAssign")
	symaddr(&addr_QRegularExpressionMatch_swap, "QRegularExpressionMatch_swap")
	symaddr(&addr_QRegularExpressionMatch_regularExpression, "QRegularExpressionMatch_regularExpression")
	symaddr(&addr_QRegularExpressionMatch_matchType, "QRegularExpressionMatch_matchType")
	symaddr(&addr_QRegularExpressionMatch_matchOptions, "QRegularExpressionMatch_matchOptions")
	symaddr(&addr_QRegularExpressionMatch_hasMatch, "QRegularExpressionMatch_hasMatch")
	symaddr(&addr_QRegularExpressionMatch_hasPartialMatch, "QRegularExpressionMatch_hasPartialMatch")
	symaddr(&addr_QRegularExpressionMatch_isValid, "QRegularExpressionMatch_isValid")
	symaddr(&addr_QRegularExpressionMatch_lastCapturedIndex, "QRegularExpressionMatch_lastCapturedIndex")
	symaddr(&addr_QRegularExpressionMatch_hasCaptured, "QRegularExpressionMatch_hasCaptured")
	symaddr(&addr_QRegularExpressionMatch_hasCaptured2, "QRegularExpressionMatch_hasCaptured2")
	symaddr(&addr_QRegularExpressionMatch_captured, "QRegularExpressionMatch_captured")
	symaddr(&addr_QRegularExpressionMatch_capturedView, "QRegularExpressionMatch_capturedView")
	symaddr(&addr_QRegularExpressionMatch_captured2, "QRegularExpressionMatch_captured2")
	symaddr(&addr_QRegularExpressionMatch_capturedView2, "QRegularExpressionMatch_capturedView2")
	symaddr(&addr_QRegularExpressionMatch_capturedTexts, "QRegularExpressionMatch_capturedTexts")
	symaddr(&addr_QRegularExpressionMatch_capturedStart, "QRegularExpressionMatch_capturedStart")
	symaddr(&addr_QRegularExpressionMatch_capturedLength, "QRegularExpressionMatch_capturedLength")
	symaddr(&addr_QRegularExpressionMatch_capturedEnd, "QRegularExpressionMatch_capturedEnd")
	symaddr(&addr_QRegularExpressionMatch_capturedStart2, "QRegularExpressionMatch_capturedStart2")
	symaddr(&addr_QRegularExpressionMatch_capturedLength2, "QRegularExpressionMatch_capturedLength2")
	symaddr(&addr_QRegularExpressionMatch_capturedEnd2, "QRegularExpressionMatch_capturedEnd2")
	symaddr(&addr_QRegularExpressionMatch_captured3, "QRegularExpressionMatch_captured3")
	symaddr(&addr_QRegularExpressionMatch_capturedView3, "QRegularExpressionMatch_capturedView3")
	symaddr(&addr_QRegularExpressionMatch_capturedStart3, "QRegularExpressionMatch_capturedStart3")
	symaddr(&addr_QRegularExpressionMatch_capturedLength3, "QRegularExpressionMatch_capturedLength3")
	symaddr(&addr_QRegularExpressionMatch_capturedEnd3, "QRegularExpressionMatch_capturedEnd3")
	reg(&pg_QRegularExpressionMatch_delete, "QRegularExpressionMatch_delete")
}

type QRegularExpressionMatch struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRegularExpressionMatch) UnsafeGoID() int64     { return this.vid }
func (this *QRegularExpressionMatch) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRegularExpressionMatch(h uintptr) *QRegularExpressionMatch {
	if h == 0 {
		return nil
	}
	return &QRegularExpressionMatch{h: h}
}

func newownedQRegularExpressionMatch(h uintptr) *QRegularExpressionMatch {
	o := bindQRegularExpressionMatch(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRegularExpressionMatch).Delete)
	}
	return o
}

func (this *QRegularExpressionMatch) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRegularExpressionMatch) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRegularExpressionMatch(h unsafe.Pointer) *QRegularExpressionMatch { return bindQRegularExpressionMatch(uintptr(h)) }
func UnsafeNewQRegularExpressionMatchOwned(h unsafe.Pointer) *QRegularExpressionMatch { return newownedQRegularExpressionMatch(uintptr(h)) }

func (this *QRegularExpressionMatch) UnsafeRelease() { this.h = 0 }

func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_new)
	__o := newownedQRegularExpressionMatch(__r)
	return __o
}

func NewQRegularExpressionMatch2(match *QRegularExpressionMatch) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_new2, uintptr(match.UnsafePointer()))
	__o := newownedQRegularExpressionMatch(__r)
	runtime.KeepAlive(match)
	return __o
}

func (this *QRegularExpressionMatch) OperatorAssign(match *QRegularExpressionMatch) *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_OperatorAssign, this.h, uintptr(match.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(match)
	return bindQRegularExpressionMatch(__r)
}

func (this *QRegularExpressionMatch) Swap(other *QRegularExpressionMatch) {
	purego.SyscallN(addr_QRegularExpressionMatch_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QRegularExpressionMatch) RegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_regularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QRegularExpressionMatch) MatchType() QRegularExpression__MatchType {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_matchType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRegularExpression__MatchType(__r)
}

func (this *QRegularExpressionMatch) MatchOptions() QRegularExpression__MatchOption {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_matchOptions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRegularExpression__MatchOption(__r)
}

func (this *QRegularExpressionMatch) HasMatch() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_hasMatch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpressionMatch) HasPartialMatch() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_hasPartialMatch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpressionMatch) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpressionMatch) LastCapturedIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_lastCapturedIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRegularExpressionMatch) HasCaptured(name string) bool {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_hasCaptured, this.h, name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r != 0
}

func (this *QRegularExpressionMatch) HasCapturedWithNth(nth int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_hasCaptured2, this.h, uintptr(nth))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpressionMatch) Captured() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_captured, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRegularExpressionMatch) CapturedView() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedView, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QRegularExpressionMatch) CapturedWithName(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_captured2, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QRegularExpressionMatch) CapturedViewWithName(name string) *QStringView {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedView2, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQStringView(__r)
}

func (this *QRegularExpressionMatch) CapturedTexts() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedTexts, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QRegularExpressionMatch) CapturedStart() int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedStart, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpressionMatch) CapturedLength() int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedLength, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpressionMatch) CapturedEnd() int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedEnd, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpressionMatch) CapturedStartWithName(name string) int64 {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedStart2, this.h, name_p, uintptr(len(name)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r
}

func (this *QRegularExpressionMatch) CapturedLengthWithName(name string) int64 {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedLength2, this.h, name_p, uintptr(len(name)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r
}

func (this *QRegularExpressionMatch) CapturedEndWithName(name string) int64 {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedEnd2, this.h, name_p, uintptr(len(name)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r
}

func (this *QRegularExpressionMatch) CapturedWithNth(nth int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_captured3, this.h, uintptr(nth), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRegularExpressionMatch) CapturedViewWithNth(nth int) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedView3, this.h, uintptr(nth))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QRegularExpressionMatch) CapturedStartWithNth(nth int) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedStart3, this.h, uintptr(nth))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpressionMatch) CapturedLengthWithNth(nth int) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedLength3, this.h, uintptr(nth))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpressionMatch) CapturedEndWithNth(nth int) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatch_capturedEnd3, this.h, uintptr(nth))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRegularExpressionMatch) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRegularExpressionMatch_delete(this.h)
	}
	this.UnsafeRelease()
}
