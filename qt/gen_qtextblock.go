package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextBlock_new uintptr
	addr_QTextBlock_new2 uintptr
	addr_QTextBlock_OperatorAssign uintptr
	addr_QTextBlock_isValid uintptr
	addr_QTextBlock_OperatorEqual uintptr
	addr_QTextBlock_OperatorNotEqual uintptr
	addr_QTextBlock_OperatorLesser uintptr
	addr_QTextBlock_position uintptr
	addr_QTextBlock_length uintptr
	addr_QTextBlock_contains uintptr
	addr_QTextBlock_layout uintptr
	addr_QTextBlock_clearLayout uintptr
	addr_QTextBlock_blockFormat uintptr
	addr_QTextBlock_blockFormatIndex uintptr
	addr_QTextBlock_charFormat uintptr
	addr_QTextBlock_charFormatIndex uintptr
	addr_QTextBlock_textDirection uintptr
	addr_QTextBlock_text uintptr
	addr_QTextBlock_textFormats uintptr
	addr_QTextBlock_document uintptr
	addr_QTextBlock_textList uintptr
	addr_QTextBlock_userData uintptr
	addr_QTextBlock_setUserData uintptr
	addr_QTextBlock_userState uintptr
	addr_QTextBlock_setUserState uintptr
	addr_QTextBlock_revision uintptr
	addr_QTextBlock_setRevision uintptr
	addr_QTextBlock_isVisible uintptr
	addr_QTextBlock_setVisible uintptr
	addr_QTextBlock_blockNumber uintptr
	addr_QTextBlock_firstLineNumber uintptr
	addr_QTextBlock_setLineCount uintptr
	addr_QTextBlock_lineCount uintptr
	addr_QTextBlock_next uintptr
	addr_QTextBlock_previous uintptr
	addr_QTextBlock_fragmentIndex uintptr
	pg_QTextBlock_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextBlock_new, "QTextBlock_new")
	symaddr(&addr_QTextBlock_new2, "QTextBlock_new2")
	symaddr(&addr_QTextBlock_OperatorAssign, "QTextBlock_OperatorAssign")
	symaddr(&addr_QTextBlock_isValid, "QTextBlock_isValid")
	symaddr(&addr_QTextBlock_OperatorEqual, "QTextBlock_OperatorEqual")
	symaddr(&addr_QTextBlock_OperatorNotEqual, "QTextBlock_OperatorNotEqual")
	symaddr(&addr_QTextBlock_OperatorLesser, "QTextBlock_OperatorLesser")
	symaddr(&addr_QTextBlock_position, "QTextBlock_position")
	symaddr(&addr_QTextBlock_length, "QTextBlock_length")
	symaddr(&addr_QTextBlock_contains, "QTextBlock_contains")
	symaddr(&addr_QTextBlock_layout, "QTextBlock_layout")
	symaddr(&addr_QTextBlock_clearLayout, "QTextBlock_clearLayout")
	symaddr(&addr_QTextBlock_blockFormat, "QTextBlock_blockFormat")
	symaddr(&addr_QTextBlock_blockFormatIndex, "QTextBlock_blockFormatIndex")
	symaddr(&addr_QTextBlock_charFormat, "QTextBlock_charFormat")
	symaddr(&addr_QTextBlock_charFormatIndex, "QTextBlock_charFormatIndex")
	symaddr(&addr_QTextBlock_textDirection, "QTextBlock_textDirection")
	symaddr(&addr_QTextBlock_text, "QTextBlock_text")
	symaddr(&addr_QTextBlock_textFormats, "QTextBlock_textFormats")
	symaddr(&addr_QTextBlock_document, "QTextBlock_document")
	symaddr(&addr_QTextBlock_textList, "QTextBlock_textList")
	symaddr(&addr_QTextBlock_userData, "QTextBlock_userData")
	symaddr(&addr_QTextBlock_setUserData, "QTextBlock_setUserData")
	symaddr(&addr_QTextBlock_userState, "QTextBlock_userState")
	symaddr(&addr_QTextBlock_setUserState, "QTextBlock_setUserState")
	symaddr(&addr_QTextBlock_revision, "QTextBlock_revision")
	symaddr(&addr_QTextBlock_setRevision, "QTextBlock_setRevision")
	symaddr(&addr_QTextBlock_isVisible, "QTextBlock_isVisible")
	symaddr(&addr_QTextBlock_setVisible, "QTextBlock_setVisible")
	symaddr(&addr_QTextBlock_blockNumber, "QTextBlock_blockNumber")
	symaddr(&addr_QTextBlock_firstLineNumber, "QTextBlock_firstLineNumber")
	symaddr(&addr_QTextBlock_setLineCount, "QTextBlock_setLineCount")
	symaddr(&addr_QTextBlock_lineCount, "QTextBlock_lineCount")
	symaddr(&addr_QTextBlock_next, "QTextBlock_next")
	symaddr(&addr_QTextBlock_previous, "QTextBlock_previous")
	symaddr(&addr_QTextBlock_fragmentIndex, "QTextBlock_fragmentIndex")
	reg(&pg_QTextBlock_delete, "QTextBlock_delete")
}

type QTextBlock struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextBlock) UnsafeGoID() int64     { return this.vid }
func (this *QTextBlock) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextBlock(h uintptr) *QTextBlock {
	if h == 0 {
		return nil
	}
	return &QTextBlock{h: h}
}

func newownedQTextBlock(h uintptr) *QTextBlock {
	o := bindQTextBlock(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextBlock).Delete)
	}
	return o
}

func (this *QTextBlock) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextBlock) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextBlock(h unsafe.Pointer) *QTextBlock { return bindQTextBlock(uintptr(h)) }
func UnsafeNewQTextBlockOwned(h unsafe.Pointer) *QTextBlock { return newownedQTextBlock(uintptr(h)) }

func (this *QTextBlock) UnsafeRelease() { this.h = 0 }

func NewQTextBlock() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_new)
	__o := newownedQTextBlock(__r)
	return __o
}

func NewQTextBlock2(o *QTextBlock) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_new2, uintptr(o.UnsafePointer()))
	__o := newownedQTextBlock(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QTextBlock) OperatorAssign(o *QTextBlock) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_OperatorAssign, this.h, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return bindQTextBlock(__r)
}

func (this *QTextBlock) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBlock) OperatorEqual(o *QTextBlock) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_OperatorEqual, this.h, uintptr(o.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return __r != 0
}

func (this *QTextBlock) OperatorNotEqual(o *QTextBlock) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_OperatorNotEqual, this.h, uintptr(o.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return __r != 0
}

func (this *QTextBlock) OperatorLesser(o *QTextBlock) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_OperatorLesser, this.h, uintptr(o.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return __r != 0
}

func (this *QTextBlock) Position() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_position, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) Length() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_length, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) Contains(position int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_contains, this.h, uintptr(position))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBlock) Layout() *QTextLayout {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_layout, this.h)
	runtime.KeepAlive(this)
	return bindQTextLayout(__r)
}

func (this *QTextBlock) ClearLayout() {
	purego.SyscallN(addr_QTextBlock_clearLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBlock) BlockFormat() *QTextBlockFormat {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_blockFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlockFormat(__r)
}

func (this *QTextBlock) BlockFormatIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_blockFormatIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) CharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_charFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextBlock) CharFormatIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_charFormatIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) TextDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_textDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QTextBlock) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBlock_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextBlock) TextFormats() []*QTextLayout__FormatRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBlock_textFormats, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextLayout__FormatRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextLayout__FormatRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextBlock) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QTextBlock) TextList() *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_textList, this.h)
	runtime.KeepAlive(this)
	return bindQTextList(__r)
}

func (this *QTextBlock) UserData() *QTextBlockUserData {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_userData, this.h)
	runtime.KeepAlive(this)
	return bindQTextBlockUserData(__r)
}

func (this *QTextBlock) SetUserData(data *QTextBlockUserData) {
	purego.SyscallN(addr_QTextBlock_setUserData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QTextBlock) UserState() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_userState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) SetUserState(state int) {
	purego.SyscallN(addr_QTextBlock_setUserState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QTextBlock) Revision() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_revision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) SetRevision(rev int) {
	purego.SyscallN(addr_QTextBlock_setRevision, this.h, uintptr(rev))
	runtime.KeepAlive(this)
}

func (this *QTextBlock) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBlock) SetVisible(visible bool) {
	purego.SyscallN(addr_QTextBlock_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTextBlock) BlockNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_blockNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) FirstLineNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_firstLineNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) SetLineCount(count int) {
	purego.SyscallN(addr_QTextBlock_setLineCount, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QTextBlock) LineCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_lineCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) Next() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_next, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextBlock) Previous() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextBlock_previous, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextBlock) FragmentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlock_fragmentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlock) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextBlock_delete(this.h)
	}
	this.UnsafeRelease()
}
