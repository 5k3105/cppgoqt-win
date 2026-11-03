package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPdfWriter_base func(uintptr) uintptr
	addr_QPdfWriter_new uintptr
	addr_QPdfWriter_new2 uintptr
	addr_QPdfWriter_tr uintptr
	addr_QPdfWriter_setPdfVersion uintptr
	addr_QPdfWriter_pdfVersion uintptr
	addr_QPdfWriter_title uintptr
	addr_QPdfWriter_setTitle uintptr
	addr_QPdfWriter_creator uintptr
	addr_QPdfWriter_setCreator uintptr
	addr_QPdfWriter_documentId uintptr
	addr_QPdfWriter_setDocumentId uintptr
	addr_QPdfWriter_author uintptr
	addr_QPdfWriter_setAuthor uintptr
	addr_QPdfWriter_newPage uintptr
	addr_QPdfWriter_setResolution uintptr
	addr_QPdfWriter_resolution uintptr
	addr_QPdfWriter_setDocumentXmpMetadata uintptr
	addr_QPdfWriter_documentXmpMetadata uintptr
	addr_QPdfWriter_addFileAttachment uintptr
	addr_QPdfWriter_colorModel uintptr
	addr_QPdfWriter_setColorModel uintptr
	addr_QPdfWriter_outputIntent uintptr
	addr_QPdfWriter_setOutputIntent uintptr
	addr_QPdfWriter_connect uintptr
	addr_QPdfWriter_connect2 uintptr
	addr_QPdfWriter_setPageLayout uintptr
	addr_QPdfWriter_setPageSize uintptr
	addr_QPdfWriter_setPageOrientation uintptr
	addr_QPdfWriter_setPageMargins uintptr
	addr_QPdfWriter_pageLayout uintptr
	addr_QPdfWriter_setPageRanges uintptr
	addr_QPdfWriter_pageRanges uintptr
	addr_QPdfWriter_devType uintptr
	addr_QPdfWriter_paintingActive uintptr
	addr_QPdfWriter_width uintptr
	addr_QPdfWriter_height uintptr
	addr_QPdfWriter_widthMM uintptr
	addr_QPdfWriter_heightMM uintptr
	addr_QPdfWriter_logicalDpiX uintptr
	addr_QPdfWriter_logicalDpiY uintptr
	addr_QPdfWriter_physicalDpiX uintptr
	addr_QPdfWriter_physicalDpiY uintptr
	pg_QPdfWriter_devicePixelRatio func(uintptr) float64
	pg_QPdfWriter_devicePixelRatioF func(uintptr) float64
	addr_QPdfWriter_colorCount uintptr
	addr_QPdfWriter_depth uintptr
	pg_QPdfWriter_devicePixelRatioFScale func() float64
	pg_QPdfWriter_encodeMetricF func(int32, float64) int32
	addr_QPdfWriter_tr2 uintptr
	addr_QPdfWriter_tr3 uintptr
	addr_QPdfWriter_addFileAttachment2 uintptr
	addr_QPdfWriter_connect3 uintptr
	addr_QPdfWriter_connect4 uintptr
	addr_QPdfWriter_asQPagedPaintDevice uintptr
	pg_QPdfWriter_connect_destroyed func(uintptr, int64)
	pg_QPdfWriter_connect_destroyed2 func(uintptr, int64)
	pg_QPdfWriter_connect_objectNameChanged func(uintptr, int64)
	pg_QPdfWriter__disconnect func(int64)
	pg_QPdfWriter_delete func(uintptr)
)

func init() {
	reg(&pg_QPdfWriter_base, "QPdfWriter_base")
	symaddr(&addr_QPdfWriter_new, "QPdfWriter_new")
	symaddr(&addr_QPdfWriter_new2, "QPdfWriter_new2")
	symaddr(&addr_QPdfWriter_tr, "QPdfWriter_tr")
	symaddr(&addr_QPdfWriter_setPdfVersion, "QPdfWriter_setPdfVersion")
	symaddr(&addr_QPdfWriter_pdfVersion, "QPdfWriter_pdfVersion")
	symaddr(&addr_QPdfWriter_title, "QPdfWriter_title")
	symaddr(&addr_QPdfWriter_setTitle, "QPdfWriter_setTitle")
	symaddr(&addr_QPdfWriter_creator, "QPdfWriter_creator")
	symaddr(&addr_QPdfWriter_setCreator, "QPdfWriter_setCreator")
	symaddr(&addr_QPdfWriter_documentId, "QPdfWriter_documentId")
	symaddr(&addr_QPdfWriter_setDocumentId, "QPdfWriter_setDocumentId")
	symaddr(&addr_QPdfWriter_author, "QPdfWriter_author")
	symaddr(&addr_QPdfWriter_setAuthor, "QPdfWriter_setAuthor")
	symaddr(&addr_QPdfWriter_newPage, "QPdfWriter_newPage")
	symaddr(&addr_QPdfWriter_setResolution, "QPdfWriter_setResolution")
	symaddr(&addr_QPdfWriter_resolution, "QPdfWriter_resolution")
	symaddr(&addr_QPdfWriter_setDocumentXmpMetadata, "QPdfWriter_setDocumentXmpMetadata")
	symaddr(&addr_QPdfWriter_documentXmpMetadata, "QPdfWriter_documentXmpMetadata")
	symaddr(&addr_QPdfWriter_addFileAttachment, "QPdfWriter_addFileAttachment")
	symaddr(&addr_QPdfWriter_colorModel, "QPdfWriter_colorModel")
	symaddr(&addr_QPdfWriter_setColorModel, "QPdfWriter_setColorModel")
	symaddr(&addr_QPdfWriter_outputIntent, "QPdfWriter_outputIntent")
	symaddr(&addr_QPdfWriter_setOutputIntent, "QPdfWriter_setOutputIntent")
	symaddr(&addr_QPdfWriter_connect, "QPdfWriter_connect")
	symaddr(&addr_QPdfWriter_connect2, "QPdfWriter_connect2")
	symaddr(&addr_QPdfWriter_setPageLayout, "QPdfWriter_setPageLayout")
	symaddr(&addr_QPdfWriter_setPageSize, "QPdfWriter_setPageSize")
	symaddr(&addr_QPdfWriter_setPageOrientation, "QPdfWriter_setPageOrientation")
	symaddr(&addr_QPdfWriter_setPageMargins, "QPdfWriter_setPageMargins")
	symaddr(&addr_QPdfWriter_pageLayout, "QPdfWriter_pageLayout")
	symaddr(&addr_QPdfWriter_setPageRanges, "QPdfWriter_setPageRanges")
	symaddr(&addr_QPdfWriter_pageRanges, "QPdfWriter_pageRanges")
	symaddr(&addr_QPdfWriter_devType, "QPdfWriter_devType")
	symaddr(&addr_QPdfWriter_paintingActive, "QPdfWriter_paintingActive")
	symaddr(&addr_QPdfWriter_width, "QPdfWriter_width")
	symaddr(&addr_QPdfWriter_height, "QPdfWriter_height")
	symaddr(&addr_QPdfWriter_widthMM, "QPdfWriter_widthMM")
	symaddr(&addr_QPdfWriter_heightMM, "QPdfWriter_heightMM")
	symaddr(&addr_QPdfWriter_logicalDpiX, "QPdfWriter_logicalDpiX")
	symaddr(&addr_QPdfWriter_logicalDpiY, "QPdfWriter_logicalDpiY")
	symaddr(&addr_QPdfWriter_physicalDpiX, "QPdfWriter_physicalDpiX")
	symaddr(&addr_QPdfWriter_physicalDpiY, "QPdfWriter_physicalDpiY")
	reg(&pg_QPdfWriter_devicePixelRatio, "QPdfWriter_devicePixelRatio")
	reg(&pg_QPdfWriter_devicePixelRatioF, "QPdfWriter_devicePixelRatioF")
	symaddr(&addr_QPdfWriter_colorCount, "QPdfWriter_colorCount")
	symaddr(&addr_QPdfWriter_depth, "QPdfWriter_depth")
	reg(&pg_QPdfWriter_devicePixelRatioFScale, "QPdfWriter_devicePixelRatioFScale")
	reg(&pg_QPdfWriter_encodeMetricF, "QPdfWriter_encodeMetricF")
	symaddr(&addr_QPdfWriter_tr2, "QPdfWriter_tr2")
	symaddr(&addr_QPdfWriter_tr3, "QPdfWriter_tr3")
	symaddr(&addr_QPdfWriter_addFileAttachment2, "QPdfWriter_addFileAttachment2")
	symaddr(&addr_QPdfWriter_connect3, "QPdfWriter_connect3")
	symaddr(&addr_QPdfWriter_connect4, "QPdfWriter_connect4")
	symaddr(&addr_QPdfWriter_asQPagedPaintDevice, "QPdfWriter_asQPagedPaintDevice")
	reg(&pg_QPdfWriter_connect_destroyed, "QPdfWriter_connect_destroyed")
	reg(&pg_QPdfWriter_connect_destroyed2, "QPdfWriter_connect_destroyed2")
	reg(&pg_QPdfWriter_connect_objectNameChanged, "QPdfWriter_connect_objectNameChanged")
	reg(&pg_QPdfWriter__disconnect, "QPdfWriter__disconnect")
	reg(&pg_QPdfWriter_delete, "QPdfWriter_delete")
}

type QPdfWriter struct {
	h uintptr
	*QObject
}

func bindQPdfWriter(h uintptr) *QPdfWriter {
	if h == 0 {
		return nil
	}
	return &QPdfWriter{h: h, QObject: bindQObject(pg_QPdfWriter_base(h))}
}

func newownedQPdfWriter(h uintptr) *QPdfWriter { return bindQPdfWriter(h) }

func (this *QPdfWriter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPdfWriter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPdfWriter(h unsafe.Pointer) *QPdfWriter { return bindQPdfWriter(uintptr(h)) }
func UnsafeNewQPdfWriterOwned(h unsafe.Pointer) *QPdfWriter { return newownedQPdfWriter(uintptr(h)) }

func (this *QPdfWriter) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQPdfWriter(filename string) *QPdfWriter {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_new, filename_p, uintptr(len(filename)))
	__o := newownedQPdfWriter(__r)
	runtime.KeepAlive(filename)
	return __o
}

func NewQPdfWriter2(device *QIODevice) *QPdfWriter {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_new2, uintptr(device.UnsafePointer()))
	__o := newownedQPdfWriter(__r)
	runtime.KeepAlive(device)
	return __o
}

func QPdfWriter_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPdfWriter) SetPdfVersion(version QPagedPaintDevice__PdfVersion) {
	purego.SyscallN(addr_QPdfWriter_setPdfVersion, this.h, uintptr(version))
	runtime.KeepAlive(this)
}

func (this *QPdfWriter) PdfVersion() QPagedPaintDevice__PdfVersion {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_pdfVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPagedPaintDevice__PdfVersion(__r)
}

func (this *QPdfWriter) Title() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_title, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPdfWriter) SetTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QPdfWriter_setTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QPdfWriter) Creator() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_creator, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPdfWriter) SetCreator(creator string) {
	var creator_p uintptr
	if len(creator) > 0 {
		creator_p = uintptr(unsafe.Pointer(unsafe.StringData(creator)))
	}
	purego.SyscallN(addr_QPdfWriter_setCreator, this.h, creator_p, uintptr(len(creator)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(creator)
}

func (this *QPdfWriter) DocumentId() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_documentId, this.h)
	runtime.KeepAlive(this)
	return newownedQUuid(__r)
}

func (this *QPdfWriter) SetDocumentId(documentId *QUuid) {
	purego.SyscallN(addr_QPdfWriter_setDocumentId, this.h, uintptr(documentId.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(documentId)
}

func (this *QPdfWriter) Author() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_author, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPdfWriter) SetAuthor(author string) {
	var author_p uintptr
	if len(author) > 0 {
		author_p = uintptr(unsafe.Pointer(unsafe.StringData(author)))
	}
	purego.SyscallN(addr_QPdfWriter_setAuthor, this.h, author_p, uintptr(len(author)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(author)
}

func (this *QPdfWriter) NewPage() bool {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_newPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPdfWriter) SetResolution(resolution int) {
	purego.SyscallN(addr_QPdfWriter_setResolution, this.h, uintptr(resolution))
	runtime.KeepAlive(this)
}

func (this *QPdfWriter) Resolution() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_resolution, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) SetDocumentXmpMetadata(xmpMetadata *QByteArray) {
	purego.SyscallN(addr_QPdfWriter_setDocumentXmpMetadata, this.h, uintptr(xmpMetadata.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(xmpMetadata)
}

func (this *QPdfWriter) DocumentXmpMetadata() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_documentXmpMetadata, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QPdfWriter) AddFileAttachment(fileName string, data *QByteArray) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QPdfWriter_addFileAttachment, this.h, fileName_p, uintptr(len(fileName)), uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(data)
}

func (this *QPdfWriter) ColorModel() QPdfWriter__ColorModel {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_colorModel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPdfWriter__ColorModel(__r)
}

func (this *QPdfWriter) SetColorModel(model QPdfWriter__ColorModel) {
	purego.SyscallN(addr_QPdfWriter_setColorModel, this.h, uintptr(model))
	runtime.KeepAlive(this)
}

func (this *QPdfWriter) OutputIntent() *QPdfOutputIntent {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_outputIntent, this.h)
	runtime.KeepAlive(this)
	return newownedQPdfOutputIntent(__r)
}

func (this *QPdfWriter) SetOutputIntent(intent *QPdfOutputIntent) {
	purego.SyscallN(addr_QPdfWriter_setOutputIntent, this.h, uintptr(intent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(intent)
}

func QPdfWriter_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPdfWriter_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPdfWriter) SetPageLayout(pageLayout *QPageLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_setPageLayout, this.h, uintptr(pageLayout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageLayout)
	return __r != 0
}

func (this *QPdfWriter) SetPageSize(pageSize *QPageSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_setPageSize, this.h, uintptr(pageSize.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageSize)
	return __r != 0
}

func (this *QPdfWriter) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_setPageOrientation, this.h, uintptr(orientation))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPdfWriter) SetPageMargins(margins *QMarginsF, units QPageLayout__Unit) bool {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_setPageMargins, this.h, uintptr(margins.UnsafePointer()), uintptr(units))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return __r != 0
}

func (this *QPdfWriter) PageLayout() *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_pageLayout, this.h)
	runtime.KeepAlive(this)
	return newownedQPageLayout(__r)
}

func (this *QPdfWriter) SetPageRanges(ranges *QPageRanges) {
	purego.SyscallN(addr_QPdfWriter_setPageRanges, this.h, uintptr(ranges.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ranges)
}

func (this *QPdfWriter) PageRanges() *QPageRanges {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_pageRanges, this.h)
	runtime.KeepAlive(this)
	return newownedQPageRanges(__r)
}

func (this *QPdfWriter) DevType() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_devType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) PaintingActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_paintingActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPdfWriter) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) WidthMM() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_widthMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) HeightMM() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_heightMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) LogicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_logicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) LogicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_logicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) PhysicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_physicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) PhysicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_physicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) DevicePixelRatio() float64 {
	__r := pg_QPdfWriter_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPdfWriter) DevicePixelRatioF() float64 {
	__r := pg_QPdfWriter_devicePixelRatioF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPdfWriter) ColorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_colorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPdfWriter) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QPdfWriter_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QPdfWriter_DevicePixelRatioFScale() float64 {
	__r := pg_QPdfWriter_devicePixelRatioFScale()
	return __r
}

func QPdfWriter_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPdfWriter_encodeMetricF(int32(metric), value)
	return int(__r)
}

func QPdfWriter_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPdfWriter_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QPdfWriter) AddFileAttachment2(fileName string, data *QByteArray, mimeType string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var mimeType_p uintptr
	if len(mimeType) > 0 {
		mimeType_p = uintptr(unsafe.Pointer(unsafe.StringData(mimeType)))
	}
	purego.SyscallN(addr_QPdfWriter_addFileAttachment2, this.h, fileName_p, uintptr(len(fileName)), uintptr(data.UnsafePointer()), mimeType_p, uintptr(len(mimeType)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(data)
	runtime.KeepAlive(mimeType)
}

func QPdfWriter_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPdfWriter_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPdfWriter) AsQPagedPaintDevice() *QPagedPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPdfWriter_asQPagedPaintDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPagedPaintDevice(__r)
}

func (this *QPdfWriter) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPdfWriter_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPdfWriter) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPdfWriter_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPdfWriter) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPdfWriter_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPdfWriter) Disconnect(handle int64) {
	pg_QPdfWriter__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QPdfWriter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPdfWriter_delete(this.h)
	}
	this.UnsafeRelease()
}
