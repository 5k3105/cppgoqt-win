package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QCompleter_base func(uintptr) uintptr
	addr_QCompleter_new uintptr
	addr_QCompleter_new2 uintptr
	addr_QCompleter_new3 uintptr
	addr_QCompleter_new4 uintptr
	addr_QCompleter_new5 uintptr
	addr_QCompleter_new6 uintptr
	addr_QCompleter_tr uintptr
	addr_QCompleter_setWidget uintptr
	addr_QCompleter_widget uintptr
	addr_QCompleter_setModel uintptr
	addr_QCompleter_model uintptr
	addr_QCompleter_setCompletionMode uintptr
	addr_QCompleter_completionMode uintptr
	addr_QCompleter_setFilterMode uintptr
	addr_QCompleter_filterMode uintptr
	addr_QCompleter_popup uintptr
	addr_QCompleter_setPopup uintptr
	addr_QCompleter_setCaseSensitivity uintptr
	addr_QCompleter_caseSensitivity uintptr
	addr_QCompleter_setModelSorting uintptr
	addr_QCompleter_modelSorting uintptr
	addr_QCompleter_setCompletionColumn uintptr
	addr_QCompleter_completionColumn uintptr
	addr_QCompleter_setCompletionRole uintptr
	addr_QCompleter_completionRole uintptr
	addr_QCompleter_wrapAround uintptr
	addr_QCompleter_maxVisibleItems uintptr
	addr_QCompleter_setMaxVisibleItems uintptr
	addr_QCompleter_completionCount uintptr
	addr_QCompleter_setCurrentRow uintptr
	addr_QCompleter_currentRow uintptr
	addr_QCompleter_currentIndex uintptr
	addr_QCompleter_currentCompletion uintptr
	addr_QCompleter_completionModel uintptr
	addr_QCompleter_completionPrefix uintptr
	addr_QCompleter_setCompletionPrefix uintptr
	addr_QCompleter_complete uintptr
	addr_QCompleter_setWrapAround uintptr
	addr_QCompleter_pathFromIndex uintptr
	addr_QCompleter_splitPath uintptr
	addr_QCompleter_activated uintptr
	addr_QCompleter_activated2 uintptr
	addr_QCompleter_highlighted uintptr
	addr_QCompleter_highlighted2 uintptr
	addr_QCompleter_connect uintptr
	addr_QCompleter_connect2 uintptr
	addr_QCompleter_sender uintptr
	addr_QCompleter_senderSignalIndex uintptr
	addr_QCompleter_receivers uintptr
	addr_QCompleter_isSignalConnected uintptr
	addr_QCompleter_tr2 uintptr
	addr_QCompleter_tr3 uintptr
	addr_QCompleter_complete2 uintptr
	addr_QCompleter_connect3 uintptr
	addr_QCompleter_connect4 uintptr
	addr_QCompleter_pathFromIndex_default uintptr
	addr_QCompleter_splitPath_default uintptr
	addr_QCompleter_metaObject_default uintptr
	addr_QCompleter_eventFilter_default uintptr
	addr_QCompleter_event_default uintptr
	addr_QCompleter_timerEvent_default uintptr
	addr_QCompleter_childEvent_default uintptr
	addr_QCompleter_customEvent_default uintptr
	addr_QCompleter_connectNotify_default uintptr
	addr_QCompleter_disconnectNotify_default uintptr
	pg_QCompleter_connect_activated func(uintptr, int64)
	pg_QCompleter_connect_activated2 func(uintptr, int64)
	pg_QCompleter_connect_highlighted func(uintptr, int64)
	pg_QCompleter_connect_highlighted2 func(uintptr, int64)
	pg_QCompleter_connect_destroyed func(uintptr, int64)
	pg_QCompleter_connect_destroyed2 func(uintptr, int64)
	pg_QCompleter_connect_objectNameChanged func(uintptr, int64)
	pg_QCompleter__disconnect func(int64)
	pg_QCompleter__set_override func(uintptr, int32)
	pg_QCompleter_delete func(uintptr)
)

func init() {
	reg(&pg_QCompleter_base, "QCompleter_base")
	symaddr(&addr_QCompleter_new, "QCompleter_new")
	symaddr(&addr_QCompleter_new2, "QCompleter_new2")
	symaddr(&addr_QCompleter_new3, "QCompleter_new3")
	symaddr(&addr_QCompleter_new4, "QCompleter_new4")
	symaddr(&addr_QCompleter_new5, "QCompleter_new5")
	symaddr(&addr_QCompleter_new6, "QCompleter_new6")
	symaddr(&addr_QCompleter_tr, "QCompleter_tr")
	symaddr(&addr_QCompleter_setWidget, "QCompleter_setWidget")
	symaddr(&addr_QCompleter_widget, "QCompleter_widget")
	symaddr(&addr_QCompleter_setModel, "QCompleter_setModel")
	symaddr(&addr_QCompleter_model, "QCompleter_model")
	symaddr(&addr_QCompleter_setCompletionMode, "QCompleter_setCompletionMode")
	symaddr(&addr_QCompleter_completionMode, "QCompleter_completionMode")
	symaddr(&addr_QCompleter_setFilterMode, "QCompleter_setFilterMode")
	symaddr(&addr_QCompleter_filterMode, "QCompleter_filterMode")
	symaddr(&addr_QCompleter_popup, "QCompleter_popup")
	symaddr(&addr_QCompleter_setPopup, "QCompleter_setPopup")
	symaddr(&addr_QCompleter_setCaseSensitivity, "QCompleter_setCaseSensitivity")
	symaddr(&addr_QCompleter_caseSensitivity, "QCompleter_caseSensitivity")
	symaddr(&addr_QCompleter_setModelSorting, "QCompleter_setModelSorting")
	symaddr(&addr_QCompleter_modelSorting, "QCompleter_modelSorting")
	symaddr(&addr_QCompleter_setCompletionColumn, "QCompleter_setCompletionColumn")
	symaddr(&addr_QCompleter_completionColumn, "QCompleter_completionColumn")
	symaddr(&addr_QCompleter_setCompletionRole, "QCompleter_setCompletionRole")
	symaddr(&addr_QCompleter_completionRole, "QCompleter_completionRole")
	symaddr(&addr_QCompleter_wrapAround, "QCompleter_wrapAround")
	symaddr(&addr_QCompleter_maxVisibleItems, "QCompleter_maxVisibleItems")
	symaddr(&addr_QCompleter_setMaxVisibleItems, "QCompleter_setMaxVisibleItems")
	symaddr(&addr_QCompleter_completionCount, "QCompleter_completionCount")
	symaddr(&addr_QCompleter_setCurrentRow, "QCompleter_setCurrentRow")
	symaddr(&addr_QCompleter_currentRow, "QCompleter_currentRow")
	symaddr(&addr_QCompleter_currentIndex, "QCompleter_currentIndex")
	symaddr(&addr_QCompleter_currentCompletion, "QCompleter_currentCompletion")
	symaddr(&addr_QCompleter_completionModel, "QCompleter_completionModel")
	symaddr(&addr_QCompleter_completionPrefix, "QCompleter_completionPrefix")
	symaddr(&addr_QCompleter_setCompletionPrefix, "QCompleter_setCompletionPrefix")
	symaddr(&addr_QCompleter_complete, "QCompleter_complete")
	symaddr(&addr_QCompleter_setWrapAround, "QCompleter_setWrapAround")
	symaddr(&addr_QCompleter_pathFromIndex, "QCompleter_pathFromIndex")
	symaddr(&addr_QCompleter_splitPath, "QCompleter_splitPath")
	symaddr(&addr_QCompleter_activated, "QCompleter_activated")
	symaddr(&addr_QCompleter_activated2, "QCompleter_activated2")
	symaddr(&addr_QCompleter_highlighted, "QCompleter_highlighted")
	symaddr(&addr_QCompleter_highlighted2, "QCompleter_highlighted2")
	symaddr(&addr_QCompleter_connect, "QCompleter_connect")
	symaddr(&addr_QCompleter_connect2, "QCompleter_connect2")
	symaddr(&addr_QCompleter_sender, "QCompleter_sender")
	symaddr(&addr_QCompleter_senderSignalIndex, "QCompleter_senderSignalIndex")
	symaddr(&addr_QCompleter_receivers, "QCompleter_receivers")
	symaddr(&addr_QCompleter_isSignalConnected, "QCompleter_isSignalConnected")
	symaddr(&addr_QCompleter_tr2, "QCompleter_tr2")
	symaddr(&addr_QCompleter_tr3, "QCompleter_tr3")
	symaddr(&addr_QCompleter_complete2, "QCompleter_complete2")
	symaddr(&addr_QCompleter_connect3, "QCompleter_connect3")
	symaddr(&addr_QCompleter_connect4, "QCompleter_connect4")
	symaddr(&addr_QCompleter_pathFromIndex_default, "QCompleter_pathFromIndex_default")
	symaddr(&addr_QCompleter_splitPath_default, "QCompleter_splitPath_default")
	symaddr(&addr_QCompleter_metaObject_default, "QCompleter_metaObject_default")
	symaddr(&addr_QCompleter_eventFilter_default, "QCompleter_eventFilter_default")
	symaddr(&addr_QCompleter_event_default, "QCompleter_event_default")
	symaddr(&addr_QCompleter_timerEvent_default, "QCompleter_timerEvent_default")
	symaddr(&addr_QCompleter_childEvent_default, "QCompleter_childEvent_default")
	symaddr(&addr_QCompleter_customEvent_default, "QCompleter_customEvent_default")
	symaddr(&addr_QCompleter_connectNotify_default, "QCompleter_connectNotify_default")
	symaddr(&addr_QCompleter_disconnectNotify_default, "QCompleter_disconnectNotify_default")
	reg(&pg_QCompleter_connect_activated, "QCompleter_connect_activated")
	reg(&pg_QCompleter_connect_activated2, "QCompleter_connect_activated2")
	reg(&pg_QCompleter_connect_highlighted, "QCompleter_connect_highlighted")
	reg(&pg_QCompleter_connect_highlighted2, "QCompleter_connect_highlighted2")
	reg(&pg_QCompleter_connect_destroyed, "QCompleter_connect_destroyed")
	reg(&pg_QCompleter_connect_destroyed2, "QCompleter_connect_destroyed2")
	reg(&pg_QCompleter_connect_objectNameChanged, "QCompleter_connect_objectNameChanged")
	reg(&pg_QCompleter__disconnect, "QCompleter__disconnect")
	reg(&pg_QCompleter__set_override, "QCompleter__set_override")
	reg(&pg_QCompleter_delete, "QCompleter_delete")
}

type QCompleter struct {
	h uintptr
	*QObject
}

func bindQCompleter(h uintptr) *QCompleter {
	if h == 0 {
		return nil
	}
	return &QCompleter{h: h, QObject: bindQObject(pg_QCompleter_base(h))}
}

func newownedQCompleter(h uintptr) *QCompleter { return bindQCompleter(h) }

func (this *QCompleter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCompleter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCompleter(h unsafe.Pointer) *QCompleter { return bindQCompleter(uintptr(h)) }
func UnsafeNewQCompleterOwned(h unsafe.Pointer) *QCompleter { return newownedQCompleter(uintptr(h)) }

func (this *QCompleter) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQCompleter() *QCompleter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCompleter_new, uintptr(__goid))
	__o := newownedQCompleter(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCompleter_virtIdx, QCompleter_enable)
	return __o
}

func NewQCompleter2(model *QAbstractItemModel) *QCompleter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCompleter_new2, uintptr(__goid), uintptr(model.UnsafePointer()))
	__o := newownedQCompleter(__r)
	runtime.KeepAlive(model)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCompleter_virtIdx, QCompleter_enable)
	return __o
}

func NewQCompleter3(completions []string) *QCompleter {
	__goid := nextVirtID()
	completions_arr := make([]uintptr, len(completions))
	for __i, __e := range completions {
		completions_arr[__i] = pgCString(__e)
	}
	var completions_p uintptr
	if len(completions_arr) > 0 {
		completions_p = uintptr(unsafe.Pointer(&completions_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCompleter_new3, uintptr(__goid), completions_p, uintptr(len(completions)))
	__o := newownedQCompleter(__r)
	for __i := range completions_arr {
		pgFree(completions_arr[__i])
	}
	runtime.KeepAlive(completions_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCompleter_virtIdx, QCompleter_enable)
	return __o
}

func NewQCompleter4(parent *QObject) *QCompleter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCompleter_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQCompleter(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCompleter_virtIdx, QCompleter_enable)
	return __o
}

func NewQCompleter5(model *QAbstractItemModel, parent *QObject) *QCompleter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCompleter_new5, uintptr(__goid), uintptr(model.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQCompleter(__r)
	runtime.KeepAlive(model)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCompleter_virtIdx, QCompleter_enable)
	return __o
}

func NewQCompleter6(completions []string, parent *QObject) *QCompleter {
	__goid := nextVirtID()
	completions_arr := make([]uintptr, len(completions))
	for __i, __e := range completions {
		completions_arr[__i] = pgCString(__e)
	}
	var completions_p uintptr
	if len(completions_arr) > 0 {
		completions_p = uintptr(unsafe.Pointer(&completions_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCompleter_new6, uintptr(__goid), completions_p, uintptr(len(completions)), uintptr(parent.UnsafePointer()))
	__o := newownedQCompleter(__r)
	for __i := range completions_arr {
		pgFree(completions_arr[__i])
	}
	runtime.KeepAlive(completions_arr)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCompleter_virtIdx, QCompleter_enable)
	return __o
}

func QCompleter_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QCompleter) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QCompleter_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QCompleter) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCompleter_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QCompleter) SetModel(c *QAbstractItemModel) {
	purego.SyscallN(addr_QCompleter_setModel, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QCompleter) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QCompleter_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QCompleter) SetCompletionMode(mode QCompleter__CompletionMode) {
	purego.SyscallN(addr_QCompleter_setCompletionMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QCompleter) CompletionMode() QCompleter__CompletionMode {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_completionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCompleter__CompletionMode(__r)
}

func (this *QCompleter) SetFilterMode(filterMode MatchFlag) {
	purego.SyscallN(addr_QCompleter_setFilterMode, this.h, uintptr(filterMode))
	runtime.KeepAlive(this)
}

func (this *QCompleter) FilterMode() MatchFlag {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_filterMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MatchFlag(__r)
}

func (this *QCompleter) Popup() *QAbstractItemView {
	__r, _, _ := purego.SyscallN(addr_QCompleter_popup, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemView(__r)
}

func (this *QCompleter) SetPopup(popup *QAbstractItemView) {
	purego.SyscallN(addr_QCompleter_setPopup, this.h, uintptr(popup.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(popup)
}

func (this *QCompleter) SetCaseSensitivity(caseSensitivity CaseSensitivity) {
	purego.SyscallN(addr_QCompleter_setCaseSensitivity, this.h, uintptr(caseSensitivity))
	runtime.KeepAlive(this)
}

func (this *QCompleter) CaseSensitivity() CaseSensitivity {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_caseSensitivity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CaseSensitivity(__r)
}

func (this *QCompleter) SetModelSorting(sorting QCompleter__ModelSorting) {
	purego.SyscallN(addr_QCompleter_setModelSorting, this.h, uintptr(sorting))
	runtime.KeepAlive(this)
}

func (this *QCompleter) ModelSorting() QCompleter__ModelSorting {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_modelSorting, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCompleter__ModelSorting(__r)
}

func (this *QCompleter) SetCompletionColumn(column int) {
	purego.SyscallN(addr_QCompleter_setCompletionColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QCompleter) CompletionColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_completionColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCompleter) SetCompletionRole(role int) {
	purego.SyscallN(addr_QCompleter_setCompletionRole, this.h, uintptr(role))
	runtime.KeepAlive(this)
}

func (this *QCompleter) CompletionRole() int {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_completionRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCompleter) WrapAround() bool {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_wrapAround, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCompleter) MaxVisibleItems() int {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_maxVisibleItems, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	purego.SyscallN(addr_QCompleter_setMaxVisibleItems, this.h, uintptr(maxItems))
	runtime.KeepAlive(this)
}

func (this *QCompleter) CompletionCount() int {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_completionCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCompleter) SetCurrentRow(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_setCurrentRow, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCompleter) CurrentRow() int {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_currentRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCompleter) CurrentIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QCompleter_currentIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QCompleter) CurrentCompletion() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_currentCompletion, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCompleter) CompletionModel() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QCompleter_completionModel, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QCompleter) CompletionPrefix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_completionPrefix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCompleter) SetCompletionPrefix(prefix string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QCompleter_setCompletionPrefix, this.h, prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
}

func (this *QCompleter) Complete() {
	purego.SyscallN(addr_QCompleter_complete, this.h)
	runtime.KeepAlive(this)
}

func (this *QCompleter) SetWrapAround(wrap bool) {
	purego.SyscallN(addr_QCompleter_setWrapAround, this.h, uintptr(b2i(wrap)))
	runtime.KeepAlive(this)
}

func (this *QCompleter) PathFromIndex(index *QModelIndex) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_pathFromIndex, this.h, uintptr(index.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return pgStr(__r, __n)
}

func (this *QCompleter) SplitPath(path string) []string {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_splitPath, this.h, path_p, uintptr(len(path)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCompleter) Activated(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QCompleter_activated, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QCompleter) ActivatedWithIndex(index *QModelIndex) {
	purego.SyscallN(addr_QCompleter_activated2, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QCompleter) Highlighted(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QCompleter_highlighted, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QCompleter) HighlightedWithIndex(index *QModelIndex) {
	purego.SyscallN(addr_QCompleter_highlighted2, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func QCompleter_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCompleter_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCompleter_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCompleter_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCompleter) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QCompleter_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QCompleter) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCompleter) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QCompleter_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QCompleter) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QCompleter_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCompleter_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QCompleter) CompleteWithRect(rect *QRect) {
	purego.SyscallN(addr_QCompleter_complete2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func QCompleter_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCompleter_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCompleter_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCompleter_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCompleter) PathFromIndexDefault(index *QModelIndex) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_pathFromIndex_default, this.h, uintptr(index.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return pgStr(__r, __n)
}

func (this *QCompleter) SplitPathDefault(path string) []string {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCompleter_splitPath_default, this.h, path_p, uintptr(len(path)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCompleter) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QCompleter_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QCompleter) EventFilterDefault(o *QObject, e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_eventFilter_default, this.h, uintptr(o.UnsafePointer()), uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QCompleter) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCompleter_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QCompleter) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QCompleter_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCompleter) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QCompleter_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCompleter) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCompleter_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCompleter) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCompleter_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCompleter) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCompleter_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCompleter) OnActivated(fn func(string)) int64 {
	id := registerCallback(func(text uintptr, text_len int64) { fn(pgStrNoFree(text, text_len)) })
	pg_QCompleter_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) OnActivatedWithIndex(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QCompleter_connect_activated2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) OnHighlighted(fn func(string)) int64 {
	id := registerCallback(func(text uintptr, text_len int64) { fn(pgStrNoFree(text, text_len)) })
	pg_QCompleter_connect_highlighted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) OnHighlightedWithIndex(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QCompleter_connect_highlighted2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCompleter_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QCompleter_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QCompleter_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCompleter) Disconnect(handle int64) {
	pg_QCompleter__disconnect(handle)
	unregisterCallback(handle)
}

func QCompleter_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -4079497074265104506:
		return 1
	case -7571154661167114460:
		return 2
	case 5217623514630858297:
		return 3
	case 1237236456312824991:
		return 4
	case -5409927581812374498:
		return 5
	case -1185315367526857833:
		return 6
	case 3769868229123868474:
		return 7
	case 7379029982170503008:
		return 8
	case -568369772982215962:
		return 9
	}
	return -1
}

func QCompleter_enable(h uintptr, idx int32) { pg_QCompleter__set_override(h, idx) }

func (this *QCompleter) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QCompleter) OnPathFromIndex(fn func(*QModelIndex) string) {
	registerVirt(this.UnsafeGoID(), -4079497074265104506, func(index uintptr) (uintptr, int64) {
		__s := fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))
		return pgCString(__s), int64(len(__s))
	}) // slot "pathFromIndex"
}

func (this *QCompleter) OnSplitPath(fn func(string) []string) {
	registerVirt(this.UnsafeGoID(), -7571154661167114460, func(path uintptr, path_len int64) (uintptr, int64) {
		__rs := fn(pgStrNoFree(path, path_len))
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "splitPath"
}

func (this *QCompleter) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(o uintptr, e uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(o)), UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "eventFilter"
}

func (this *QCompleter) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QCompleter) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QCompleter) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QCompleter) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QCompleter) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QCompleter) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QCompleter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCompleter_delete(this.h)
	}
	this.UnsafeRelease()
}
