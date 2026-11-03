package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionViewItem_base func(uintptr) uintptr
	addr_QStyleOptionViewItem_new uintptr
	addr_QStyleOptionViewItem_new2 uintptr
	addr_QStyleOptionViewItem_OperatorAssign uintptr
	addr_QStyleOptionViewItem_displayAlignment_get uintptr
	addr_QStyleOptionViewItem_displayAlignment_set uintptr
	addr_QStyleOptionViewItem_decorationAlignment_get uintptr
	addr_QStyleOptionViewItem_decorationAlignment_set uintptr
	addr_QStyleOptionViewItem_textElideMode_get uintptr
	addr_QStyleOptionViewItem_textElideMode_set uintptr
	addr_QStyleOptionViewItem_decorationPosition_get uintptr
	addr_QStyleOptionViewItem_decorationPosition_set uintptr
	addr_QStyleOptionViewItem_decorationSize_get uintptr
	addr_QStyleOptionViewItem_decorationSize_set uintptr
	addr_QStyleOptionViewItem_font_get uintptr
	addr_QStyleOptionViewItem_font_set uintptr
	addr_QStyleOptionViewItem_showDecorationSelected_get uintptr
	addr_QStyleOptionViewItem_showDecorationSelected_set uintptr
	addr_QStyleOptionViewItem_features_get uintptr
	addr_QStyleOptionViewItem_features_set uintptr
	addr_QStyleOptionViewItem_locale_get uintptr
	addr_QStyleOptionViewItem_locale_set uintptr
	addr_QStyleOptionViewItem_widget_get uintptr
	addr_QStyleOptionViewItem_widget_set uintptr
	addr_QStyleOptionViewItem_index_get uintptr
	addr_QStyleOptionViewItem_index_set uintptr
	addr_QStyleOptionViewItem_checkState_get uintptr
	addr_QStyleOptionViewItem_checkState_set uintptr
	addr_QStyleOptionViewItem_icon_get uintptr
	addr_QStyleOptionViewItem_icon_set uintptr
	addr_QStyleOptionViewItem_text_get uintptr
	addr_QStyleOptionViewItem_text_set uintptr
	addr_QStyleOptionViewItem_viewItemPosition_get uintptr
	addr_QStyleOptionViewItem_viewItemPosition_set uintptr
	addr_QStyleOptionViewItem_backgroundBrush_get uintptr
	addr_QStyleOptionViewItem_backgroundBrush_set uintptr
	pg_QStyleOptionViewItem_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionViewItem_base, "QStyleOptionViewItem_base")
	symaddr(&addr_QStyleOptionViewItem_new, "QStyleOptionViewItem_new")
	symaddr(&addr_QStyleOptionViewItem_new2, "QStyleOptionViewItem_new2")
	symaddr(&addr_QStyleOptionViewItem_OperatorAssign, "QStyleOptionViewItem_OperatorAssign")
	symaddr(&addr_QStyleOptionViewItem_displayAlignment_get, "QStyleOptionViewItem_displayAlignment_get")
	symaddr(&addr_QStyleOptionViewItem_displayAlignment_set, "QStyleOptionViewItem_displayAlignment_set")
	symaddr(&addr_QStyleOptionViewItem_decorationAlignment_get, "QStyleOptionViewItem_decorationAlignment_get")
	symaddr(&addr_QStyleOptionViewItem_decorationAlignment_set, "QStyleOptionViewItem_decorationAlignment_set")
	symaddr(&addr_QStyleOptionViewItem_textElideMode_get, "QStyleOptionViewItem_textElideMode_get")
	symaddr(&addr_QStyleOptionViewItem_textElideMode_set, "QStyleOptionViewItem_textElideMode_set")
	symaddr(&addr_QStyleOptionViewItem_decorationPosition_get, "QStyleOptionViewItem_decorationPosition_get")
	symaddr(&addr_QStyleOptionViewItem_decorationPosition_set, "QStyleOptionViewItem_decorationPosition_set")
	symaddr(&addr_QStyleOptionViewItem_decorationSize_get, "QStyleOptionViewItem_decorationSize_get")
	symaddr(&addr_QStyleOptionViewItem_decorationSize_set, "QStyleOptionViewItem_decorationSize_set")
	symaddr(&addr_QStyleOptionViewItem_font_get, "QStyleOptionViewItem_font_get")
	symaddr(&addr_QStyleOptionViewItem_font_set, "QStyleOptionViewItem_font_set")
	symaddr(&addr_QStyleOptionViewItem_showDecorationSelected_get, "QStyleOptionViewItem_showDecorationSelected_get")
	symaddr(&addr_QStyleOptionViewItem_showDecorationSelected_set, "QStyleOptionViewItem_showDecorationSelected_set")
	symaddr(&addr_QStyleOptionViewItem_features_get, "QStyleOptionViewItem_features_get")
	symaddr(&addr_QStyleOptionViewItem_features_set, "QStyleOptionViewItem_features_set")
	symaddr(&addr_QStyleOptionViewItem_locale_get, "QStyleOptionViewItem_locale_get")
	symaddr(&addr_QStyleOptionViewItem_locale_set, "QStyleOptionViewItem_locale_set")
	symaddr(&addr_QStyleOptionViewItem_widget_get, "QStyleOptionViewItem_widget_get")
	symaddr(&addr_QStyleOptionViewItem_widget_set, "QStyleOptionViewItem_widget_set")
	symaddr(&addr_QStyleOptionViewItem_index_get, "QStyleOptionViewItem_index_get")
	symaddr(&addr_QStyleOptionViewItem_index_set, "QStyleOptionViewItem_index_set")
	symaddr(&addr_QStyleOptionViewItem_checkState_get, "QStyleOptionViewItem_checkState_get")
	symaddr(&addr_QStyleOptionViewItem_checkState_set, "QStyleOptionViewItem_checkState_set")
	symaddr(&addr_QStyleOptionViewItem_icon_get, "QStyleOptionViewItem_icon_get")
	symaddr(&addr_QStyleOptionViewItem_icon_set, "QStyleOptionViewItem_icon_set")
	symaddr(&addr_QStyleOptionViewItem_text_get, "QStyleOptionViewItem_text_get")
	symaddr(&addr_QStyleOptionViewItem_text_set, "QStyleOptionViewItem_text_set")
	symaddr(&addr_QStyleOptionViewItem_viewItemPosition_get, "QStyleOptionViewItem_viewItemPosition_get")
	symaddr(&addr_QStyleOptionViewItem_viewItemPosition_set, "QStyleOptionViewItem_viewItemPosition_set")
	symaddr(&addr_QStyleOptionViewItem_backgroundBrush_get, "QStyleOptionViewItem_backgroundBrush_get")
	symaddr(&addr_QStyleOptionViewItem_backgroundBrush_set, "QStyleOptionViewItem_backgroundBrush_set")
	reg(&pg_QStyleOptionViewItem_delete, "QStyleOptionViewItem_delete")
}

type QStyleOptionViewItem struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionViewItem(h uintptr) *QStyleOptionViewItem {
	if h == 0 {
		return nil
	}
	return &QStyleOptionViewItem{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionViewItem_base(h))}
}

func newownedQStyleOptionViewItem(h uintptr) *QStyleOptionViewItem {
	o := bindQStyleOptionViewItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionViewItem).Delete)
	}
	return o
}

func (this *QStyleOptionViewItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionViewItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionViewItem(h unsafe.Pointer) *QStyleOptionViewItem { return bindQStyleOptionViewItem(uintptr(h)) }
func UnsafeNewQStyleOptionViewItemOwned(h unsafe.Pointer) *QStyleOptionViewItem { return newownedQStyleOptionViewItem(uintptr(h)) }

func (this *QStyleOptionViewItem) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_new)
	__o := newownedQStyleOptionViewItem(__r)
	return __o
}

func NewQStyleOptionViewItem2(other *QStyleOptionViewItem) *QStyleOptionViewItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionViewItem(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionViewItem) OperatorAssign(arg1 *QStyleOptionViewItem) *QStyleOptionViewItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionViewItem(__r)
}

func (this *QStyleOptionViewItem) DisplayAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_displayAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionViewItem) SetDisplayAlignment(displayAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionViewItem_displayAlignment_set, this.h, uintptr(displayAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) DecorationAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_decorationAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionViewItem) SetDecorationAlignment(decorationAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionViewItem_decorationAlignment_set, this.h, uintptr(decorationAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) TextElideMode() TextElideMode {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_textElideMode_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextElideMode(__r)
}

func (this *QStyleOptionViewItem) SetTextElideMode(textElideMode TextElideMode) {
	purego.SyscallN(addr_QStyleOptionViewItem_textElideMode_set, this.h, uintptr(textElideMode))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) DecorationPosition() QStyleOptionViewItem__Position {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_decorationPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionViewItem__Position(__r)
}

func (this *QStyleOptionViewItem) SetDecorationPosition(decorationPosition QStyleOptionViewItem__Position) {
	purego.SyscallN(addr_QStyleOptionViewItem_decorationPosition_set, this.h, uintptr(decorationPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) DecorationSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_decorationSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionViewItem) SetDecorationSize(decorationSize *QSize) {
	purego.SyscallN(addr_QStyleOptionViewItem_decorationSize_set, this.h, uintptr(decorationSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(decorationSize)
}

func (this *QStyleOptionViewItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_font_get, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QStyleOptionViewItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QStyleOptionViewItem_font_set, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QStyleOptionViewItem) ShowDecorationSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_showDecorationSelected_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionViewItem) SetShowDecorationSelected(showDecorationSelected bool) {
	purego.SyscallN(addr_QStyleOptionViewItem_showDecorationSelected_set, this.h, uintptr(b2i(showDecorationSelected)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) Features() QStyleOptionViewItem__ViewItemFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionViewItem__ViewItemFeature(__r)
}

func (this *QStyleOptionViewItem) SetFeatures(features QStyleOptionViewItem__ViewItemFeature) {
	purego.SyscallN(addr_QStyleOptionViewItem_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_locale_get, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QStyleOptionViewItem) SetLocale(locale *QLocale) {
	purego.SyscallN(addr_QStyleOptionViewItem_locale_set, this.h, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
}

func (this *QStyleOptionViewItem) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_widget_get, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QStyleOptionViewItem) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QStyleOptionViewItem_widget_set, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStyleOptionViewItem) Index() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_index_get, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QStyleOptionViewItem) SetIndex(index *QModelIndex) {
	purego.SyscallN(addr_QStyleOptionViewItem_index_set, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QStyleOptionViewItem) CheckState() CheckState {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_checkState_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CheckState(__r)
}

func (this *QStyleOptionViewItem) SetCheckState(checkState CheckState) {
	purego.SyscallN(addr_QStyleOptionViewItem_checkState_set, this.h, uintptr(checkState))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionViewItem) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionViewItem_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionViewItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionViewItem) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionViewItem_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionViewItem) ViewItemPosition() QStyleOptionViewItem__ViewItemPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_viewItemPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionViewItem__ViewItemPosition(__r)
}

func (this *QStyleOptionViewItem) SetViewItemPosition(viewItemPosition QStyleOptionViewItem__ViewItemPosition) {
	purego.SyscallN(addr_QStyleOptionViewItem_viewItemPosition_set, this.h, uintptr(viewItemPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionViewItem) BackgroundBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionViewItem_backgroundBrush_get, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QStyleOptionViewItem) SetBackgroundBrush(backgroundBrush *QBrush) {
	purego.SyscallN(addr_QStyleOptionViewItem_backgroundBrush_set, this.h, uintptr(backgroundBrush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(backgroundBrush)
}

func (this *QStyleOptionViewItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionViewItem_delete(this.h)
	}
	this.UnsafeRelease()
}
