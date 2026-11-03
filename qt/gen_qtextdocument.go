package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextDocument_base func(uintptr) uintptr
	addr_QTextDocument_new uintptr
	addr_QTextDocument_new2 uintptr
	addr_QTextDocument_new3 uintptr
	addr_QTextDocument_new4 uintptr
	addr_QTextDocument_tr uintptr
	addr_QTextDocument_clone uintptr
	addr_QTextDocument_isEmpty uintptr
	addr_QTextDocument_clear uintptr
	addr_QTextDocument_setUndoRedoEnabled uintptr
	addr_QTextDocument_isUndoRedoEnabled uintptr
	addr_QTextDocument_isUndoAvailable uintptr
	addr_QTextDocument_isRedoAvailable uintptr
	addr_QTextDocument_availableUndoSteps uintptr
	addr_QTextDocument_availableRedoSteps uintptr
	addr_QTextDocument_revision uintptr
	addr_QTextDocument_setDocumentLayout uintptr
	addr_QTextDocument_documentLayout uintptr
	addr_QTextDocument_setMetaInformation uintptr
	addr_QTextDocument_metaInformation uintptr
	addr_QTextDocument_toHtml uintptr
	addr_QTextDocument_setHtml uintptr
	addr_QTextDocument_toMarkdown uintptr
	addr_QTextDocument_setMarkdown uintptr
	addr_QTextDocument_toRawText uintptr
	addr_QTextDocument_toPlainText uintptr
	addr_QTextDocument_setPlainText uintptr
	addr_QTextDocument_characterAt uintptr
	addr_QTextDocument_find uintptr
	addr_QTextDocument_find2 uintptr
	addr_QTextDocument_find3 uintptr
	addr_QTextDocument_find4 uintptr
	addr_QTextDocument_frameAt uintptr
	addr_QTextDocument_rootFrame uintptr
	addr_QTextDocument_object uintptr
	addr_QTextDocument_objectForFormat uintptr
	addr_QTextDocument_findBlock uintptr
	addr_QTextDocument_findBlockByNumber uintptr
	addr_QTextDocument_findBlockByLineNumber uintptr
	addr_QTextDocument_begin uintptr
	addr_QTextDocument_end uintptr
	addr_QTextDocument_firstBlock uintptr
	addr_QTextDocument_lastBlock uintptr
	addr_QTextDocument_setPageSize uintptr
	addr_QTextDocument_pageSize uintptr
	addr_QTextDocument_setDefaultFont uintptr
	addr_QTextDocument_defaultFont uintptr
	pg_QTextDocument_setSuperScriptBaseline func(uintptr, float64)
	pg_QTextDocument_superScriptBaseline func(uintptr) float64
	pg_QTextDocument_setSubScriptBaseline func(uintptr, float64)
	pg_QTextDocument_subScriptBaseline func(uintptr) float64
	pg_QTextDocument_setBaselineOffset func(uintptr, float64)
	pg_QTextDocument_baselineOffset func(uintptr) float64
	addr_QTextDocument_pageCount uintptr
	addr_QTextDocument_isModified uintptr
	addr_QTextDocument_print uintptr
	addr_QTextDocument_resource uintptr
	addr_QTextDocument_addResource uintptr
	addr_QTextDocument_allFormats uintptr
	addr_QTextDocument_markContentsDirty uintptr
	addr_QTextDocument_setUseDesignMetrics uintptr
	addr_QTextDocument_useDesignMetrics uintptr
	addr_QTextDocument_setLayoutEnabled uintptr
	addr_QTextDocument_isLayoutEnabled uintptr
	addr_QTextDocument_drawContents uintptr
	pg_QTextDocument_setTextWidth func(uintptr, float64)
	pg_QTextDocument_textWidth func(uintptr) float64
	pg_QTextDocument_idealWidth func(uintptr) float64
	pg_QTextDocument_indentWidth func(uintptr) float64
	pg_QTextDocument_setIndentWidth func(uintptr, float64)
	pg_QTextDocument_documentMargin func(uintptr) float64
	pg_QTextDocument_setDocumentMargin func(uintptr, float64)
	addr_QTextDocument_adjustSize uintptr
	addr_QTextDocument_size uintptr
	addr_QTextDocument_blockCount uintptr
	addr_QTextDocument_lineCount uintptr
	addr_QTextDocument_characterCount uintptr
	addr_QTextDocument_setDefaultStyleSheet uintptr
	addr_QTextDocument_defaultStyleSheet uintptr
	addr_QTextDocument_undo uintptr
	addr_QTextDocument_redo uintptr
	addr_QTextDocument_clearUndoRedoStacks uintptr
	addr_QTextDocument_maximumBlockCount uintptr
	addr_QTextDocument_setMaximumBlockCount uintptr
	addr_QTextDocument_defaultTextOption uintptr
	addr_QTextDocument_setDefaultTextOption uintptr
	addr_QTextDocument_baseUrl uintptr
	addr_QTextDocument_setBaseUrl uintptr
	addr_QTextDocument_defaultCursorMoveStyle uintptr
	addr_QTextDocument_setDefaultCursorMoveStyle uintptr
	addr_QTextDocument_contentsChange uintptr
	addr_QTextDocument_contentsChanged uintptr
	addr_QTextDocument_undoAvailable uintptr
	addr_QTextDocument_redoAvailable uintptr
	addr_QTextDocument_undoCommandAdded uintptr
	addr_QTextDocument_modificationChanged uintptr
	addr_QTextDocument_cursorPositionChanged uintptr
	addr_QTextDocument_blockCountChanged uintptr
	addr_QTextDocument_baseUrlChanged uintptr
	addr_QTextDocument_documentLayoutChanged uintptr
	addr_QTextDocument_undo2 uintptr
	addr_QTextDocument_redo2 uintptr
	addr_QTextDocument_appendUndoItem uintptr
	addr_QTextDocument_setModified uintptr
	addr_QTextDocument_connect uintptr
	addr_QTextDocument_connect2 uintptr
	addr_QTextDocument_sender uintptr
	addr_QTextDocument_senderSignalIndex uintptr
	addr_QTextDocument_receivers uintptr
	addr_QTextDocument_isSignalConnected uintptr
	addr_QTextDocument_tr2 uintptr
	addr_QTextDocument_tr3 uintptr
	addr_QTextDocument_clone2 uintptr
	addr_QTextDocument_toMarkdown2 uintptr
	addr_QTextDocument_setMarkdown2 uintptr
	addr_QTextDocument_find5 uintptr
	addr_QTextDocument_find6 uintptr
	addr_QTextDocument_find7 uintptr
	addr_QTextDocument_find8 uintptr
	addr_QTextDocument_find9 uintptr
	addr_QTextDocument_find10 uintptr
	addr_QTextDocument_drawContents2 uintptr
	addr_QTextDocument_clearUndoRedoStacks2 uintptr
	addr_QTextDocument_setModified2 uintptr
	addr_QTextDocument_connect3 uintptr
	addr_QTextDocument_connect4 uintptr
	addr_QTextDocument_clear_default uintptr
	addr_QTextDocument_metaObject_default uintptr
	addr_QTextDocument_createObject_default uintptr
	addr_QTextDocument_loadResource_default uintptr
	addr_QTextDocument_event_default uintptr
	addr_QTextDocument_eventFilter_default uintptr
	addr_QTextDocument_timerEvent_default uintptr
	addr_QTextDocument_childEvent_default uintptr
	addr_QTextDocument_customEvent_default uintptr
	addr_QTextDocument_connectNotify_default uintptr
	addr_QTextDocument_disconnectNotify_default uintptr
	pg_QTextDocument_connect_contentsChange func(uintptr, int64)
	pg_QTextDocument_connect_contentsChanged func(uintptr, int64)
	pg_QTextDocument_connect_undoAvailable func(uintptr, int64)
	pg_QTextDocument_connect_redoAvailable func(uintptr, int64)
	pg_QTextDocument_connect_undoCommandAdded func(uintptr, int64)
	pg_QTextDocument_connect_modificationChanged func(uintptr, int64)
	pg_QTextDocument_connect_cursorPositionChanged func(uintptr, int64)
	pg_QTextDocument_connect_blockCountChanged func(uintptr, int64)
	pg_QTextDocument_connect_baseUrlChanged func(uintptr, int64)
	pg_QTextDocument_connect_documentLayoutChanged func(uintptr, int64)
	pg_QTextDocument_connect_destroyed func(uintptr, int64)
	pg_QTextDocument_connect_destroyed2 func(uintptr, int64)
	pg_QTextDocument_connect_objectNameChanged func(uintptr, int64)
	pg_QTextDocument__disconnect func(int64)
	pg_QTextDocument__set_override func(uintptr, int32)
	pg_QTextDocument_delete func(uintptr)
)

func init() {
	reg(&pg_QTextDocument_base, "QTextDocument_base")
	symaddr(&addr_QTextDocument_new, "QTextDocument_new")
	symaddr(&addr_QTextDocument_new2, "QTextDocument_new2")
	symaddr(&addr_QTextDocument_new3, "QTextDocument_new3")
	symaddr(&addr_QTextDocument_new4, "QTextDocument_new4")
	symaddr(&addr_QTextDocument_tr, "QTextDocument_tr")
	symaddr(&addr_QTextDocument_clone, "QTextDocument_clone")
	symaddr(&addr_QTextDocument_isEmpty, "QTextDocument_isEmpty")
	symaddr(&addr_QTextDocument_clear, "QTextDocument_clear")
	symaddr(&addr_QTextDocument_setUndoRedoEnabled, "QTextDocument_setUndoRedoEnabled")
	symaddr(&addr_QTextDocument_isUndoRedoEnabled, "QTextDocument_isUndoRedoEnabled")
	symaddr(&addr_QTextDocument_isUndoAvailable, "QTextDocument_isUndoAvailable")
	symaddr(&addr_QTextDocument_isRedoAvailable, "QTextDocument_isRedoAvailable")
	symaddr(&addr_QTextDocument_availableUndoSteps, "QTextDocument_availableUndoSteps")
	symaddr(&addr_QTextDocument_availableRedoSteps, "QTextDocument_availableRedoSteps")
	symaddr(&addr_QTextDocument_revision, "QTextDocument_revision")
	symaddr(&addr_QTextDocument_setDocumentLayout, "QTextDocument_setDocumentLayout")
	symaddr(&addr_QTextDocument_documentLayout, "QTextDocument_documentLayout")
	symaddr(&addr_QTextDocument_setMetaInformation, "QTextDocument_setMetaInformation")
	symaddr(&addr_QTextDocument_metaInformation, "QTextDocument_metaInformation")
	symaddr(&addr_QTextDocument_toHtml, "QTextDocument_toHtml")
	symaddr(&addr_QTextDocument_setHtml, "QTextDocument_setHtml")
	symaddr(&addr_QTextDocument_toMarkdown, "QTextDocument_toMarkdown")
	symaddr(&addr_QTextDocument_setMarkdown, "QTextDocument_setMarkdown")
	symaddr(&addr_QTextDocument_toRawText, "QTextDocument_toRawText")
	symaddr(&addr_QTextDocument_toPlainText, "QTextDocument_toPlainText")
	symaddr(&addr_QTextDocument_setPlainText, "QTextDocument_setPlainText")
	symaddr(&addr_QTextDocument_characterAt, "QTextDocument_characterAt")
	symaddr(&addr_QTextDocument_find, "QTextDocument_find")
	symaddr(&addr_QTextDocument_find2, "QTextDocument_find2")
	symaddr(&addr_QTextDocument_find3, "QTextDocument_find3")
	symaddr(&addr_QTextDocument_find4, "QTextDocument_find4")
	symaddr(&addr_QTextDocument_frameAt, "QTextDocument_frameAt")
	symaddr(&addr_QTextDocument_rootFrame, "QTextDocument_rootFrame")
	symaddr(&addr_QTextDocument_object, "QTextDocument_object")
	symaddr(&addr_QTextDocument_objectForFormat, "QTextDocument_objectForFormat")
	symaddr(&addr_QTextDocument_findBlock, "QTextDocument_findBlock")
	symaddr(&addr_QTextDocument_findBlockByNumber, "QTextDocument_findBlockByNumber")
	symaddr(&addr_QTextDocument_findBlockByLineNumber, "QTextDocument_findBlockByLineNumber")
	symaddr(&addr_QTextDocument_begin, "QTextDocument_begin")
	symaddr(&addr_QTextDocument_end, "QTextDocument_end")
	symaddr(&addr_QTextDocument_firstBlock, "QTextDocument_firstBlock")
	symaddr(&addr_QTextDocument_lastBlock, "QTextDocument_lastBlock")
	symaddr(&addr_QTextDocument_setPageSize, "QTextDocument_setPageSize")
	symaddr(&addr_QTextDocument_pageSize, "QTextDocument_pageSize")
	symaddr(&addr_QTextDocument_setDefaultFont, "QTextDocument_setDefaultFont")
	symaddr(&addr_QTextDocument_defaultFont, "QTextDocument_defaultFont")
	reg(&pg_QTextDocument_setSuperScriptBaseline, "QTextDocument_setSuperScriptBaseline")
	reg(&pg_QTextDocument_superScriptBaseline, "QTextDocument_superScriptBaseline")
	reg(&pg_QTextDocument_setSubScriptBaseline, "QTextDocument_setSubScriptBaseline")
	reg(&pg_QTextDocument_subScriptBaseline, "QTextDocument_subScriptBaseline")
	reg(&pg_QTextDocument_setBaselineOffset, "QTextDocument_setBaselineOffset")
	reg(&pg_QTextDocument_baselineOffset, "QTextDocument_baselineOffset")
	symaddr(&addr_QTextDocument_pageCount, "QTextDocument_pageCount")
	symaddr(&addr_QTextDocument_isModified, "QTextDocument_isModified")
	symaddr(&addr_QTextDocument_print, "QTextDocument_print")
	symaddr(&addr_QTextDocument_resource, "QTextDocument_resource")
	symaddr(&addr_QTextDocument_addResource, "QTextDocument_addResource")
	symaddr(&addr_QTextDocument_allFormats, "QTextDocument_allFormats")
	symaddr(&addr_QTextDocument_markContentsDirty, "QTextDocument_markContentsDirty")
	symaddr(&addr_QTextDocument_setUseDesignMetrics, "QTextDocument_setUseDesignMetrics")
	symaddr(&addr_QTextDocument_useDesignMetrics, "QTextDocument_useDesignMetrics")
	symaddr(&addr_QTextDocument_setLayoutEnabled, "QTextDocument_setLayoutEnabled")
	symaddr(&addr_QTextDocument_isLayoutEnabled, "QTextDocument_isLayoutEnabled")
	symaddr(&addr_QTextDocument_drawContents, "QTextDocument_drawContents")
	reg(&pg_QTextDocument_setTextWidth, "QTextDocument_setTextWidth")
	reg(&pg_QTextDocument_textWidth, "QTextDocument_textWidth")
	reg(&pg_QTextDocument_idealWidth, "QTextDocument_idealWidth")
	reg(&pg_QTextDocument_indentWidth, "QTextDocument_indentWidth")
	reg(&pg_QTextDocument_setIndentWidth, "QTextDocument_setIndentWidth")
	reg(&pg_QTextDocument_documentMargin, "QTextDocument_documentMargin")
	reg(&pg_QTextDocument_setDocumentMargin, "QTextDocument_setDocumentMargin")
	symaddr(&addr_QTextDocument_adjustSize, "QTextDocument_adjustSize")
	symaddr(&addr_QTextDocument_size, "QTextDocument_size")
	symaddr(&addr_QTextDocument_blockCount, "QTextDocument_blockCount")
	symaddr(&addr_QTextDocument_lineCount, "QTextDocument_lineCount")
	symaddr(&addr_QTextDocument_characterCount, "QTextDocument_characterCount")
	symaddr(&addr_QTextDocument_setDefaultStyleSheet, "QTextDocument_setDefaultStyleSheet")
	symaddr(&addr_QTextDocument_defaultStyleSheet, "QTextDocument_defaultStyleSheet")
	symaddr(&addr_QTextDocument_undo, "QTextDocument_undo")
	symaddr(&addr_QTextDocument_redo, "QTextDocument_redo")
	symaddr(&addr_QTextDocument_clearUndoRedoStacks, "QTextDocument_clearUndoRedoStacks")
	symaddr(&addr_QTextDocument_maximumBlockCount, "QTextDocument_maximumBlockCount")
	symaddr(&addr_QTextDocument_setMaximumBlockCount, "QTextDocument_setMaximumBlockCount")
	symaddr(&addr_QTextDocument_defaultTextOption, "QTextDocument_defaultTextOption")
	symaddr(&addr_QTextDocument_setDefaultTextOption, "QTextDocument_setDefaultTextOption")
	symaddr(&addr_QTextDocument_baseUrl, "QTextDocument_baseUrl")
	symaddr(&addr_QTextDocument_setBaseUrl, "QTextDocument_setBaseUrl")
	symaddr(&addr_QTextDocument_defaultCursorMoveStyle, "QTextDocument_defaultCursorMoveStyle")
	symaddr(&addr_QTextDocument_setDefaultCursorMoveStyle, "QTextDocument_setDefaultCursorMoveStyle")
	symaddr(&addr_QTextDocument_contentsChange, "QTextDocument_contentsChange")
	symaddr(&addr_QTextDocument_contentsChanged, "QTextDocument_contentsChanged")
	symaddr(&addr_QTextDocument_undoAvailable, "QTextDocument_undoAvailable")
	symaddr(&addr_QTextDocument_redoAvailable, "QTextDocument_redoAvailable")
	symaddr(&addr_QTextDocument_undoCommandAdded, "QTextDocument_undoCommandAdded")
	symaddr(&addr_QTextDocument_modificationChanged, "QTextDocument_modificationChanged")
	symaddr(&addr_QTextDocument_cursorPositionChanged, "QTextDocument_cursorPositionChanged")
	symaddr(&addr_QTextDocument_blockCountChanged, "QTextDocument_blockCountChanged")
	symaddr(&addr_QTextDocument_baseUrlChanged, "QTextDocument_baseUrlChanged")
	symaddr(&addr_QTextDocument_documentLayoutChanged, "QTextDocument_documentLayoutChanged")
	symaddr(&addr_QTextDocument_undo2, "QTextDocument_undo2")
	symaddr(&addr_QTextDocument_redo2, "QTextDocument_redo2")
	symaddr(&addr_QTextDocument_appendUndoItem, "QTextDocument_appendUndoItem")
	symaddr(&addr_QTextDocument_setModified, "QTextDocument_setModified")
	symaddr(&addr_QTextDocument_connect, "QTextDocument_connect")
	symaddr(&addr_QTextDocument_connect2, "QTextDocument_connect2")
	symaddr(&addr_QTextDocument_sender, "QTextDocument_sender")
	symaddr(&addr_QTextDocument_senderSignalIndex, "QTextDocument_senderSignalIndex")
	symaddr(&addr_QTextDocument_receivers, "QTextDocument_receivers")
	symaddr(&addr_QTextDocument_isSignalConnected, "QTextDocument_isSignalConnected")
	symaddr(&addr_QTextDocument_tr2, "QTextDocument_tr2")
	symaddr(&addr_QTextDocument_tr3, "QTextDocument_tr3")
	symaddr(&addr_QTextDocument_clone2, "QTextDocument_clone2")
	symaddr(&addr_QTextDocument_toMarkdown2, "QTextDocument_toMarkdown2")
	symaddr(&addr_QTextDocument_setMarkdown2, "QTextDocument_setMarkdown2")
	symaddr(&addr_QTextDocument_find5, "QTextDocument_find5")
	symaddr(&addr_QTextDocument_find6, "QTextDocument_find6")
	symaddr(&addr_QTextDocument_find7, "QTextDocument_find7")
	symaddr(&addr_QTextDocument_find8, "QTextDocument_find8")
	symaddr(&addr_QTextDocument_find9, "QTextDocument_find9")
	symaddr(&addr_QTextDocument_find10, "QTextDocument_find10")
	symaddr(&addr_QTextDocument_drawContents2, "QTextDocument_drawContents2")
	symaddr(&addr_QTextDocument_clearUndoRedoStacks2, "QTextDocument_clearUndoRedoStacks2")
	symaddr(&addr_QTextDocument_setModified2, "QTextDocument_setModified2")
	symaddr(&addr_QTextDocument_connect3, "QTextDocument_connect3")
	symaddr(&addr_QTextDocument_connect4, "QTextDocument_connect4")
	symaddr(&addr_QTextDocument_clear_default, "QTextDocument_clear_default")
	symaddr(&addr_QTextDocument_metaObject_default, "QTextDocument_metaObject_default")
	symaddr(&addr_QTextDocument_createObject_default, "QTextDocument_createObject_default")
	symaddr(&addr_QTextDocument_loadResource_default, "QTextDocument_loadResource_default")
	symaddr(&addr_QTextDocument_event_default, "QTextDocument_event_default")
	symaddr(&addr_QTextDocument_eventFilter_default, "QTextDocument_eventFilter_default")
	symaddr(&addr_QTextDocument_timerEvent_default, "QTextDocument_timerEvent_default")
	symaddr(&addr_QTextDocument_childEvent_default, "QTextDocument_childEvent_default")
	symaddr(&addr_QTextDocument_customEvent_default, "QTextDocument_customEvent_default")
	symaddr(&addr_QTextDocument_connectNotify_default, "QTextDocument_connectNotify_default")
	symaddr(&addr_QTextDocument_disconnectNotify_default, "QTextDocument_disconnectNotify_default")
	reg(&pg_QTextDocument_connect_contentsChange, "QTextDocument_connect_contentsChange")
	reg(&pg_QTextDocument_connect_contentsChanged, "QTextDocument_connect_contentsChanged")
	reg(&pg_QTextDocument_connect_undoAvailable, "QTextDocument_connect_undoAvailable")
	reg(&pg_QTextDocument_connect_redoAvailable, "QTextDocument_connect_redoAvailable")
	reg(&pg_QTextDocument_connect_undoCommandAdded, "QTextDocument_connect_undoCommandAdded")
	reg(&pg_QTextDocument_connect_modificationChanged, "QTextDocument_connect_modificationChanged")
	reg(&pg_QTextDocument_connect_cursorPositionChanged, "QTextDocument_connect_cursorPositionChanged")
	reg(&pg_QTextDocument_connect_blockCountChanged, "QTextDocument_connect_blockCountChanged")
	reg(&pg_QTextDocument_connect_baseUrlChanged, "QTextDocument_connect_baseUrlChanged")
	reg(&pg_QTextDocument_connect_documentLayoutChanged, "QTextDocument_connect_documentLayoutChanged")
	reg(&pg_QTextDocument_connect_destroyed, "QTextDocument_connect_destroyed")
	reg(&pg_QTextDocument_connect_destroyed2, "QTextDocument_connect_destroyed2")
	reg(&pg_QTextDocument_connect_objectNameChanged, "QTextDocument_connect_objectNameChanged")
	reg(&pg_QTextDocument__disconnect, "QTextDocument__disconnect")
	reg(&pg_QTextDocument__set_override, "QTextDocument__set_override")
	reg(&pg_QTextDocument_delete, "QTextDocument_delete")
}

type QTextDocument struct {
	h uintptr
	*QObject
}

func bindQTextDocument(h uintptr) *QTextDocument {
	if h == 0 {
		return nil
	}
	return &QTextDocument{h: h, QObject: bindQObject(pg_QTextDocument_base(h))}
}

func newownedQTextDocument(h uintptr) *QTextDocument { return bindQTextDocument(h) }

func (this *QTextDocument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextDocument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextDocument(h unsafe.Pointer) *QTextDocument { return bindQTextDocument(uintptr(h)) }
func UnsafeNewQTextDocumentOwned(h unsafe.Pointer) *QTextDocument { return newownedQTextDocument(uintptr(h)) }

func (this *QTextDocument) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQTextDocument() *QTextDocument {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTextDocument_new, uintptr(__goid))
	__o := newownedQTextDocument(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextDocument_virtIdx, QTextDocument_enable)
	return __o
}

func NewQTextDocument2(text string) *QTextDocument {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQTextDocument(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextDocument_virtIdx, QTextDocument_enable)
	return __o
}

func NewQTextDocument3(parent *QObject) *QTextDocument {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTextDocument_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTextDocument(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextDocument_virtIdx, QTextDocument_enable)
	return __o
}

func NewQTextDocument4(text string, parent *QObject) *QTextDocument {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQTextDocument(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextDocument_virtIdx, QTextDocument_enable)
	return __o
}

func QTextDocument_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextDocument) Clone() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_clone, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QTextDocument) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) Clear() {
	purego.SyscallN(addr_QTextDocument_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	purego.SyscallN(addr_QTextDocument_setUndoRedoEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) IsUndoRedoEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isUndoRedoEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) IsUndoAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isUndoAvailable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) IsRedoAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isRedoAvailable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) AvailableUndoSteps() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_availableUndoSteps, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) AvailableRedoSteps() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_availableRedoSteps, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) Revision() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_revision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) SetDocumentLayout(layout *QAbstractTextDocumentLayout) {
	purego.SyscallN(addr_QTextDocument_setDocumentLayout, this.h, uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_documentLayout, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractTextDocumentLayout(__r)
}

func (this *QTextDocument) SetMetaInformation(info QTextDocument__MetaInformation, arg2 string) {
	var arg2_p uintptr
	if len(arg2) > 0 {
		arg2_p = uintptr(unsafe.Pointer(unsafe.StringData(arg2)))
	}
	purego.SyscallN(addr_QTextDocument_setMetaInformation, this.h, uintptr(info), arg2_p, uintptr(len(arg2)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg2)
}

func (this *QTextDocument) MetaInformation(info QTextDocument__MetaInformation) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_metaInformation, this.h, uintptr(info), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) ToHtml() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_toHtml, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) SetHtml(html string) {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	purego.SyscallN(addr_QTextDocument_setHtml, this.h, html_p, uintptr(len(html)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(html)
}

func (this *QTextDocument) ToMarkdown() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_toMarkdown, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) SetMarkdown(markdown string) {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	purego.SyscallN(addr_QTextDocument_setMarkdown, this.h, markdown_p, uintptr(len(markdown)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(markdown)
}

func (this *QTextDocument) ToRawText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_toRawText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) ToPlainText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_toPlainText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) SetPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextDocument_setPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextDocument) CharacterAt(pos int) *QChar {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_characterAt, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QTextDocument) Find(subString string) *QTextCursor {
	var subString_p uintptr
	if len(subString) > 0 {
		subString_p = uintptr(unsafe.Pointer(unsafe.StringData(subString)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find, this.h, subString_p, uintptr(len(subString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subString)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find2(subString string, cursor *QTextCursor) *QTextCursor {
	var subString_p uintptr
	if len(subString) > 0 {
		subString_p = uintptr(unsafe.Pointer(unsafe.StringData(subString)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find2, this.h, subString_p, uintptr(len(subString)), uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subString)
	runtime.KeepAlive(cursor)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) FindWithExpr(expr *QRegularExpression) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find3, this.h, uintptr(expr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(expr)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find3(expr *QRegularExpression, cursor *QTextCursor) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find4, this.h, uintptr(expr.UnsafePointer()), uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(expr)
	runtime.KeepAlive(cursor)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) FrameAt(pos int) *QTextFrame {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_frameAt, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQTextFrame(__r)
}

func (this *QTextDocument) RootFrame() *QTextFrame {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_rootFrame, this.h)
	runtime.KeepAlive(this)
	return bindQTextFrame(__r)
}

func (this *QTextDocument) Object(objectIndex int) *QTextObject {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_object, this.h, uintptr(objectIndex))
	runtime.KeepAlive(this)
	return bindQTextObject(__r)
}

func (this *QTextDocument) ObjectForFormat(arg1 *QTextFormat) *QTextObject {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_objectForFormat, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextObject(__r)
}

func (this *QTextDocument) FindBlock(pos int) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_findBlock, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) FindBlockByNumber(blockNumber int) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_findBlockByNumber, this.h, uintptr(blockNumber))
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_findBlockByLineNumber, this.h, uintptr(blockNumber))
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) Begin() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_begin, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) End() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_end, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) FirstBlock() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_firstBlock, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) LastBlock() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_lastBlock, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextDocument) SetPageSize(size *QSizeF) {
	purego.SyscallN(addr_QTextDocument_setPageSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QTextDocument) PageSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_pageSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QTextDocument) SetDefaultFont(font *QFont) {
	purego.SyscallN(addr_QTextDocument_setDefaultFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QTextDocument) DefaultFont() *QFont {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_defaultFont, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTextDocument) SetSuperScriptBaseline(baseline float64) {
	pg_QTextDocument_setSuperScriptBaseline(this.h, baseline)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) SuperScriptBaseline() float64 {
	__r := pg_QTextDocument_superScriptBaseline(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) SetSubScriptBaseline(baseline float64) {
	pg_QTextDocument_setSubScriptBaseline(this.h, baseline)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) SubScriptBaseline() float64 {
	__r := pg_QTextDocument_subScriptBaseline(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) SetBaselineOffset(baseline float64) {
	pg_QTextDocument_setBaselineOffset(this.h, baseline)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) BaselineOffset() float64 {
	__r := pg_QTextDocument_baselineOffset(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) PageCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_pageCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) IsModified() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isModified, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) Print(printer *QPagedPaintDevice) {
	purego.SyscallN(addr_QTextDocument_print, this.h, uintptr(printer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(printer)
}

func (this *QTextDocument) Resource(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_resource, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QTextDocument) AddResource(type_ int, name *QUrl, resource *QVariant) {
	purego.SyscallN(addr_QTextDocument_addResource, this.h, uintptr(type_), uintptr(name.UnsafePointer()), uintptr(resource.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(resource)
}

func (this *QTextDocument) AllFormats() []*QTextFormat {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_allFormats, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextFormat, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextFormat(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	purego.SyscallN(addr_QTextDocument_markContentsDirty, this.h, uintptr(from), uintptr(length))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	purego.SyscallN(addr_QTextDocument_setUseDesignMetrics, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) UseDesignMetrics() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_useDesignMetrics, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) SetLayoutEnabled(b bool) {
	purego.SyscallN(addr_QTextDocument_setLayoutEnabled, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) IsLayoutEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isLayoutEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocument) DrawContents(painter *QPainter) {
	purego.SyscallN(addr_QTextDocument_drawContents, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTextDocument) SetTextWidth(width float64) {
	pg_QTextDocument_setTextWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) TextWidth() float64 {
	__r := pg_QTextDocument_textWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) IdealWidth() float64 {
	__r := pg_QTextDocument_idealWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) IndentWidth() float64 {
	__r := pg_QTextDocument_indentWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) SetIndentWidth(width float64) {
	pg_QTextDocument_setIndentWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) DocumentMargin() float64 {
	__r := pg_QTextDocument_documentMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextDocument) SetDocumentMargin(margin float64) {
	pg_QTextDocument_setDocumentMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) AdjustSize() {
	purego.SyscallN(addr_QTextDocument_adjustSize, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) Size() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QTextDocument) BlockCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_blockCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) LineCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_lineCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) CharacterCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_characterCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) SetDefaultStyleSheet(sheet string) {
	var sheet_p uintptr
	if len(sheet) > 0 {
		sheet_p = uintptr(unsafe.Pointer(unsafe.StringData(sheet)))
	}
	purego.SyscallN(addr_QTextDocument_setDefaultStyleSheet, this.h, sheet_p, uintptr(len(sheet)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sheet)
}

func (this *QTextDocument) DefaultStyleSheet() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_defaultStyleSheet, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) Undo(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextDocument_undo, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextDocument) Redo(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextDocument_redo, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextDocument) ClearUndoRedoStacks() {
	purego.SyscallN(addr_QTextDocument_clearUndoRedoStacks, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) MaximumBlockCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_maximumBlockCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	purego.SyscallN(addr_QTextDocument_setMaximumBlockCount, this.h, uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) DefaultTextOption() *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_defaultTextOption, this.h)
	runtime.KeepAlive(this)
	return newownedQTextOption(__r)
}

func (this *QTextDocument) SetDefaultTextOption(option *QTextOption) {
	purego.SyscallN(addr_QTextDocument_setDefaultTextOption, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTextDocument) BaseUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_baseUrl, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QTextDocument) SetBaseUrl(url *QUrl) {
	purego.SyscallN(addr_QTextDocument_setBaseUrl, this.h, uintptr(url.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func (this *QTextDocument) DefaultCursorMoveStyle() CursorMoveStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_defaultCursorMoveStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CursorMoveStyle(__r)
}

func (this *QTextDocument) SetDefaultCursorMoveStyle(style CursorMoveStyle) {
	purego.SyscallN(addr_QTextDocument_setDefaultCursorMoveStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	purego.SyscallN(addr_QTextDocument_contentsChange, this.h, uintptr(from), uintptr(charsRemoved), uintptr(charsAdded))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) ContentsChanged() {
	purego.SyscallN(addr_QTextDocument_contentsChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) UndoAvailable(arg1 bool) {
	purego.SyscallN(addr_QTextDocument_undoAvailable, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) RedoAvailable(arg1 bool) {
	purego.SyscallN(addr_QTextDocument_redoAvailable, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) UndoCommandAdded() {
	purego.SyscallN(addr_QTextDocument_undoCommandAdded, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) ModificationChanged(m bool) {
	purego.SyscallN(addr_QTextDocument_modificationChanged, this.h, uintptr(b2i(m)))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) CursorPositionChanged(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextDocument_cursorPositionChanged, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextDocument) BlockCountChanged(newBlockCount int) {
	purego.SyscallN(addr_QTextDocument_blockCountChanged, this.h, uintptr(newBlockCount))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) BaseUrlChanged(url *QUrl) {
	purego.SyscallN(addr_QTextDocument_baseUrlChanged, this.h, uintptr(url.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func (this *QTextDocument) DocumentLayoutChanged() {
	purego.SyscallN(addr_QTextDocument_documentLayoutChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) Undo2() {
	purego.SyscallN(addr_QTextDocument_undo2, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) Redo2() {
	purego.SyscallN(addr_QTextDocument_redo2, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) AppendUndoItem(arg1 *QAbstractUndoItem) {
	purego.SyscallN(addr_QTextDocument_appendUndoItem, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextDocument) SetModified() {
	purego.SyscallN(addr_QTextDocument_setModified, this.h)
	runtime.KeepAlive(this)
}

func QTextDocument_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextDocument_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextDocument_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextDocument) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTextDocument) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextDocument) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTextDocument) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTextDocument_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextDocument_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTextDocument) CloneWithParent(parent *QObject) *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_clone2, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQTextDocument(__r)
}

func (this *QTextDocument) ToMarkdownWithFeatures(features QTextDocument__MarkdownFeature) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocument_toMarkdown2, this.h, uintptr(features), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocument) SetMarkdown2(markdown string, features QTextDocument__MarkdownFeature) {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	purego.SyscallN(addr_QTextDocument_setMarkdown2, this.h, markdown_p, uintptr(len(markdown)), uintptr(features))
	runtime.KeepAlive(this)
	runtime.KeepAlive(markdown)
}

func (this *QTextDocument) Find4(subString string, from int) *QTextCursor {
	var subString_p uintptr
	if len(subString) > 0 {
		subString_p = uintptr(unsafe.Pointer(unsafe.StringData(subString)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find5, this.h, subString_p, uintptr(len(subString)), uintptr(from))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subString)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find5(subString string, from int, options QTextDocument__FindFlag) *QTextCursor {
	var subString_p uintptr
	if len(subString) > 0 {
		subString_p = uintptr(unsafe.Pointer(unsafe.StringData(subString)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find6, this.h, subString_p, uintptr(len(subString)), uintptr(from), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subString)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find6(subString string, cursor *QTextCursor, options QTextDocument__FindFlag) *QTextCursor {
	var subString_p uintptr
	if len(subString) > 0 {
		subString_p = uintptr(unsafe.Pointer(unsafe.StringData(subString)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find7, this.h, subString_p, uintptr(len(subString)), uintptr(cursor.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subString)
	runtime.KeepAlive(cursor)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find7(expr *QRegularExpression, from int) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find8, this.h, uintptr(expr.UnsafePointer()), uintptr(from))
	runtime.KeepAlive(this)
	runtime.KeepAlive(expr)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find8(expr *QRegularExpression, from int, options QTextDocument__FindFlag) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find9, this.h, uintptr(expr.UnsafePointer()), uintptr(from), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(expr)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) Find9(expr *QRegularExpression, cursor *QTextCursor, options QTextDocument__FindFlag) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_find10, this.h, uintptr(expr.UnsafePointer()), uintptr(cursor.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(expr)
	runtime.KeepAlive(cursor)
	return newownedQTextCursor(__r)
}

func (this *QTextDocument) DrawContents2(painter *QPainter, rect *QRectF) {
	purego.SyscallN(addr_QTextDocument_drawContents2, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QTextDocument) ClearUndoRedoStacksWithHistoryToClear(historyToClear QTextDocument__Stacks) {
	purego.SyscallN(addr_QTextDocument_clearUndoRedoStacks2, this.h, uintptr(historyToClear))
	runtime.KeepAlive(this)
}

func (this *QTextDocument) SetModifiedWithBool(m bool) {
	purego.SyscallN(addr_QTextDocument_setModified2, this.h, uintptr(b2i(m)))
	runtime.KeepAlive(this)
}

func QTextDocument_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextDocument_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextDocument_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextDocument) ClearDefault() {
	purego.SyscallN(addr_QTextDocument_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextDocument) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTextDocument) CreateObjectDefault(f *QTextFormat) *QTextObject {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_createObject_default, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
	return bindQTextObject(__r)
}

func (this *QTextDocument) LoadResourceDefault(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextDocument_loadResource_default, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QTextDocument) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTextDocument) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocument_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTextDocument) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTextDocument_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextDocument) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTextDocument_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextDocument) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTextDocument_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextDocument) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTextDocument_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTextDocument) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTextDocument_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTextDocument) OnContentsChange(fn func(int, int, int)) int64 {
	id := registerCallback(func(from int32, charsRemoved int32, charsAdded int32) { fn(int(from), int(charsRemoved), int(charsAdded)) })
	pg_QTextDocument_connect_contentsChange(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnContentsChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextDocument_connect_contentsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnUndoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QTextDocument_connect_undoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnRedoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QTextDocument_connect_redoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnUndoCommandAdded(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextDocument_connect_undoCommandAdded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnModificationChanged(fn func(bool)) int64 {
	id := registerCallback(func(m int32) { fn(m != 0) })
	pg_QTextDocument_connect_modificationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnCursorPositionChanged(fn func(*QTextCursor)) int64 {
	id := registerCallback(func(cursor uintptr) { fn(UnsafeNewQTextCursor(unsafe.Pointer(cursor))) })
	pg_QTextDocument_connect_cursorPositionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnBlockCountChanged(fn func(int)) int64 {
	id := registerCallback(func(newBlockCount int32) { fn(int(newBlockCount)) })
	pg_QTextDocument_connect_blockCountChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnBaseUrlChanged(fn func(*QUrl)) int64 {
	id := registerCallback(func(url uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(url))) })
	pg_QTextDocument_connect_baseUrlChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnDocumentLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextDocument_connect_documentLayoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextDocument_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextDocument_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextDocument_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextDocument) Disconnect(handle int64) {
	pg_QTextDocument__disconnect(handle)
	unregisterCallback(handle)
}

func QTextDocument_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -778567940572639070:
		return 1
	case 7258893770181894642:
		return 2
	case -519616681208050567:
		return 3
	case 1237236456312824991:
		return 4
	case 5217623514630858297:
		return 5
	case -5409927581812374498:
		return 6
	case -1185315367526857833:
		return 7
	case 3769868229123868474:
		return 8
	case 7379029982170503008:
		return 9
	case -568369772982215962:
		return 10
	}
	return -1
}

func QTextDocument_enable(h uintptr, idx int32) { pg_QTextDocument__set_override(h, idx) }

func (this *QTextDocument) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTextDocument) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QTextDocument) OnCreateObject(fn func(*QTextFormat) *QTextObject) {
	registerVirt(this.UnsafeGoID(), 7258893770181894642, func(f uintptr) uintptr { return uintptr(fn(UnsafeNewQTextFormat(unsafe.Pointer(f))).UnsafePointer()) }) // slot "createObject"
}

func (this *QTextDocument) OnLoadResource(fn func(int, *QUrl) *QVariant) {
	registerVirt(this.UnsafeGoID(), -519616681208050567, func(type_ int32, name uintptr) uintptr { return uintptr(fn(int(type_), UnsafeNewQUrl(unsafe.Pointer(name))).UnsafePointer()) }) // slot "loadResource"
}

func (this *QTextDocument) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTextDocument) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTextDocument) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTextDocument) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTextDocument) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTextDocument) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTextDocument) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTextDocument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextDocument_delete(this.h)
	}
	this.UnsafeRelease()
}
