package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDebug_base func(uintptr) uintptr
	addr_QDebug_new uintptr
	addr_QDebug_new2 uintptr
	addr_QDebug_new3 uintptr
	addr_QDebug_new4 uintptr
	addr_QDebug_new5 uintptr
	addr_QDebug_OperatorAssign uintptr
	addr_QDebug_swap uintptr
	addr_QDebug_resetFormat uintptr
	addr_QDebug_space uintptr
	addr_QDebug_nospace uintptr
	addr_QDebug_maybeSpace uintptr
	addr_QDebug_verbosity uintptr
	addr_QDebug_verbosity2 uintptr
	addr_QDebug_setVerbosity uintptr
	addr_QDebug_autoInsertSpaces uintptr
	addr_QDebug_setAutoInsertSpaces uintptr
	addr_QDebug_quoteStrings uintptr
	addr_QDebug_setQuoteStrings uintptr
	addr_QDebug_quote uintptr
	addr_QDebug_noquote uintptr
	addr_QDebug_maybeQuote uintptr
	addr_QDebug_OperatorShiftLeft uintptr
	addr_QDebug_OperatorShiftLeft2 uintptr
	addr_QDebug_OperatorShiftLeft3 uintptr
	addr_QDebug_OperatorShiftLeft4 uintptr
	addr_QDebug_OperatorShiftLeft5 uintptr
	addr_QDebug_OperatorShiftLeft6 uintptr
	addr_QDebug_OperatorShiftLeft7 uintptr
	addr_QDebug_OperatorShiftLeft8 uintptr
	addr_QDebug_OperatorShiftLeft9 uintptr
	addr_QDebug_OperatorShiftLeft10 uintptr
	addr_QDebug_OperatorShiftLeft11 uintptr
	addr_QDebug_OperatorShiftLeft12 uintptr
	addr_QDebug_OperatorShiftLeft13 uintptr
	addr_QDebug_OperatorShiftLeft14 uintptr
	pg_QDebug_OperatorShiftLeft15 func(uintptr, float32) uintptr
	pg_QDebug_OperatorShiftLeft16 func(uintptr, float64) uintptr
	addr_QDebug_OperatorShiftLeft17 uintptr
	addr_QDebug_OperatorShiftLeft18 uintptr
	addr_QDebug_OperatorShiftLeft19 uintptr
	addr_QDebug_OperatorShiftLeft20 uintptr
	addr_QDebug_OperatorShiftLeft21 uintptr
	addr_QDebug_OperatorShiftLeft22 uintptr
	addr_QDebug_OperatorShiftLeft23 uintptr
	addr_QDebug_maybeQuote2 uintptr
	addr_QDebug_OperatorShiftLeft24 uintptr
	addr_QDebug_OperatorShiftLeft25 uintptr
	addr_QDebug_OperatorShiftLeft26 uintptr
	addr_QDebug_OperatorShiftLeft27 uintptr
	addr_QDebug_OperatorShiftLeft28 uintptr
	addr_QDebug_OperatorShiftLeft29 uintptr
	addr_QDebug_OperatorShiftLeft30 uintptr
	addr_QDebug_OperatorShiftLeft31 uintptr
	addr_QDebug_OperatorShiftLeft32 uintptr
	addr_QDebug_OperatorShiftLeft33 uintptr
	addr_QDebug_OperatorShiftLeft34 uintptr
	addr_QDebug_OperatorShiftLeft35 uintptr
	addr_QDebug_OperatorShiftLeft36 uintptr
	addr_QDebug_OperatorShiftLeft37 uintptr
	addr_QDebug_OperatorShiftLeft38 uintptr
	addr_QDebug_OperatorShiftLeft39 uintptr
	addr_QDebug_OperatorShiftLeft40 uintptr
	addr_QDebug_OperatorShiftLeft41 uintptr
	addr_QDebug_OperatorShiftLeft42 uintptr
	addr_QDebug_OperatorShiftLeft43 uintptr
	addr_QDebug_OperatorShiftLeft44 uintptr
	addr_QDebug_OperatorShiftLeft45 uintptr
	addr_QDebug_OperatorShiftLeft46 uintptr
	addr_QDebug_OperatorShiftLeft47 uintptr
	addr_QDebug_OperatorShiftLeft48 uintptr
	addr_QDebug_OperatorShiftLeft49 uintptr
	addr_QDebug_OperatorShiftLeft50 uintptr
	addr_QDebug_OperatorShiftLeft51 uintptr
	addr_QDebug_OperatorShiftLeft52 uintptr
	addr_QDebug_OperatorShiftLeft53 uintptr
	addr_QDebug_OperatorShiftLeft54 uintptr
	addr_QDebug_OperatorShiftLeft55 uintptr
	addr_QDebug_OperatorShiftLeft56 uintptr
	addr_QDebug_OperatorShiftLeft57 uintptr
	addr_QDebug_OperatorShiftLeft58 uintptr
	addr_QDebug_OperatorShiftLeft59 uintptr
	addr_QDebug_OperatorShiftLeft60 uintptr
	addr_QDebug_OperatorShiftLeft61 uintptr
	addr_QDebug_OperatorShiftLeft62 uintptr
	addr_QDebug_OperatorShiftLeft63 uintptr
	addr_QDebug_OperatorShiftLeft64 uintptr
	addr_QDebug_OperatorShiftLeft65 uintptr
	addr_QDebug_OperatorShiftLeft66 uintptr
	addr_QDebug_OperatorShiftLeft67 uintptr
	addr_QDebug_OperatorShiftLeft68 uintptr
	addr_QDebug_OperatorShiftLeft69 uintptr
	addr_QDebug_OperatorShiftLeft70 uintptr
	addr_QDebug_OperatorShiftLeft71 uintptr
	addr_QDebug_OperatorShiftLeft72 uintptr
	addr_QDebug_OperatorShiftLeft73 uintptr
	addr_QDebug_OperatorShiftLeft74 uintptr
	addr_QDebug_OperatorShiftLeft75 uintptr
	addr_QDebug_OperatorShiftLeft76 uintptr
	addr_QDebug_OperatorShiftLeft77 uintptr
	addr_QDebug_OperatorShiftLeft78 uintptr
	addr_QDebug_OperatorShiftLeft79 uintptr
	addr_QDebug_OperatorShiftLeft80 uintptr
	addr_QDebug_OperatorShiftLeft81 uintptr
	addr_QDebug_OperatorShiftLeft82 uintptr
	addr_QDebug_OperatorShiftLeft83 uintptr
	addr_QDebug_OperatorShiftLeft84 uintptr
	addr_QDebug_OperatorShiftLeft85 uintptr
	addr_QDebug_OperatorShiftLeft86 uintptr
	addr_QDebug_OperatorShiftLeft87 uintptr
	addr_QDebug_OperatorShiftLeft88 uintptr
	addr_QDebug_OperatorShiftLeft89 uintptr
	addr_QDebug_OperatorShiftLeft90 uintptr
	addr_QDebug_OperatorShiftLeft91 uintptr
	addr_QDebug_OperatorShiftLeft92 uintptr
	addr_QDebug_OperatorShiftLeft93 uintptr
	addr_QDebug_OperatorShiftLeft94 uintptr
	addr_QDebug_OperatorShiftLeft95 uintptr
	addr_QDebug_OperatorShiftLeft96 uintptr
	addr_QDebug_OperatorShiftLeft97 uintptr
	addr_QDebug_OperatorShiftLeft98 uintptr
	addr_QDebug_OperatorShiftLeft99 uintptr
	addr_QDebug_OperatorShiftLeft100 uintptr
	addr_QDebug_OperatorShiftLeft101 uintptr
	addr_QDebug_OperatorShiftLeft102 uintptr
	addr_QDebug_OperatorShiftLeft103 uintptr
	addr_QDebug_OperatorShiftLeft104 uintptr
	addr_QDebug_OperatorShiftLeft105 uintptr
	addr_QDebug_OperatorShiftLeft106 uintptr
	addr_QDebug_OperatorShiftLeft107 uintptr
	addr_QDebug_OperatorShiftLeft108 uintptr
	addr_QDebug_OperatorShiftLeft109 uintptr
	addr_QDebug_OperatorShiftLeft110 uintptr
	addr_QDebug_OperatorShiftLeft111 uintptr
	addr_QDebug_OperatorShiftLeft112 uintptr
	addr_QDebug_OperatorShiftLeft113 uintptr
	addr_QDebug_OperatorShiftLeft114 uintptr
	addr_QDebug_OperatorShiftLeft115 uintptr
	addr_QDebug_OperatorShiftLeft116 uintptr
	addr_QDebug_OperatorShiftLeft117 uintptr
	addr_QDebug_OperatorShiftLeft118 uintptr
	addr_QDebug_OperatorShiftLeft119 uintptr
	addr_QDebug_OperatorShiftLeft120 uintptr
	addr_QDebug_OperatorShiftLeft121 uintptr
	addr_QDebug_OperatorShiftLeft122 uintptr
	addr_QDebug_OperatorShiftLeft123 uintptr
	addr_QDebug_OperatorShiftLeft124 uintptr
	addr_QDebug_OperatorShiftLeft125 uintptr
	addr_QDebug_OperatorShiftLeft126 uintptr
	addr_QDebug_OperatorShiftLeft127 uintptr
	pg_QDebug_delete func(uintptr)
)

func init() {
	reg(&pg_QDebug_base, "QDebug_base")
	symaddr(&addr_QDebug_new, "QDebug_new")
	symaddr(&addr_QDebug_new2, "QDebug_new2")
	symaddr(&addr_QDebug_new3, "QDebug_new3")
	symaddr(&addr_QDebug_new4, "QDebug_new4")
	symaddr(&addr_QDebug_new5, "QDebug_new5")
	symaddr(&addr_QDebug_OperatorAssign, "QDebug_OperatorAssign")
	symaddr(&addr_QDebug_swap, "QDebug_swap")
	symaddr(&addr_QDebug_resetFormat, "QDebug_resetFormat")
	symaddr(&addr_QDebug_space, "QDebug_space")
	symaddr(&addr_QDebug_nospace, "QDebug_nospace")
	symaddr(&addr_QDebug_maybeSpace, "QDebug_maybeSpace")
	symaddr(&addr_QDebug_verbosity, "QDebug_verbosity")
	symaddr(&addr_QDebug_verbosity2, "QDebug_verbosity2")
	symaddr(&addr_QDebug_setVerbosity, "QDebug_setVerbosity")
	symaddr(&addr_QDebug_autoInsertSpaces, "QDebug_autoInsertSpaces")
	symaddr(&addr_QDebug_setAutoInsertSpaces, "QDebug_setAutoInsertSpaces")
	symaddr(&addr_QDebug_quoteStrings, "QDebug_quoteStrings")
	symaddr(&addr_QDebug_setQuoteStrings, "QDebug_setQuoteStrings")
	symaddr(&addr_QDebug_quote, "QDebug_quote")
	symaddr(&addr_QDebug_noquote, "QDebug_noquote")
	symaddr(&addr_QDebug_maybeQuote, "QDebug_maybeQuote")
	symaddr(&addr_QDebug_OperatorShiftLeft, "QDebug_OperatorShiftLeft")
	symaddr(&addr_QDebug_OperatorShiftLeft2, "QDebug_OperatorShiftLeft2")
	symaddr(&addr_QDebug_OperatorShiftLeft3, "QDebug_OperatorShiftLeft3")
	symaddr(&addr_QDebug_OperatorShiftLeft4, "QDebug_OperatorShiftLeft4")
	symaddr(&addr_QDebug_OperatorShiftLeft5, "QDebug_OperatorShiftLeft5")
	symaddr(&addr_QDebug_OperatorShiftLeft6, "QDebug_OperatorShiftLeft6")
	symaddr(&addr_QDebug_OperatorShiftLeft7, "QDebug_OperatorShiftLeft7")
	symaddr(&addr_QDebug_OperatorShiftLeft8, "QDebug_OperatorShiftLeft8")
	symaddr(&addr_QDebug_OperatorShiftLeft9, "QDebug_OperatorShiftLeft9")
	symaddr(&addr_QDebug_OperatorShiftLeft10, "QDebug_OperatorShiftLeft10")
	symaddr(&addr_QDebug_OperatorShiftLeft11, "QDebug_OperatorShiftLeft11")
	symaddr(&addr_QDebug_OperatorShiftLeft12, "QDebug_OperatorShiftLeft12")
	symaddr(&addr_QDebug_OperatorShiftLeft13, "QDebug_OperatorShiftLeft13")
	symaddr(&addr_QDebug_OperatorShiftLeft14, "QDebug_OperatorShiftLeft14")
	reg(&pg_QDebug_OperatorShiftLeft15, "QDebug_OperatorShiftLeft15")
	reg(&pg_QDebug_OperatorShiftLeft16, "QDebug_OperatorShiftLeft16")
	symaddr(&addr_QDebug_OperatorShiftLeft17, "QDebug_OperatorShiftLeft17")
	symaddr(&addr_QDebug_OperatorShiftLeft18, "QDebug_OperatorShiftLeft18")
	symaddr(&addr_QDebug_OperatorShiftLeft19, "QDebug_OperatorShiftLeft19")
	symaddr(&addr_QDebug_OperatorShiftLeft20, "QDebug_OperatorShiftLeft20")
	symaddr(&addr_QDebug_OperatorShiftLeft21, "QDebug_OperatorShiftLeft21")
	symaddr(&addr_QDebug_OperatorShiftLeft22, "QDebug_OperatorShiftLeft22")
	symaddr(&addr_QDebug_OperatorShiftLeft23, "QDebug_OperatorShiftLeft23")
	symaddr(&addr_QDebug_maybeQuote2, "QDebug_maybeQuote2")
	symaddr(&addr_QDebug_OperatorShiftLeft24, "QDebug_OperatorShiftLeft24")
	symaddr(&addr_QDebug_OperatorShiftLeft25, "QDebug_OperatorShiftLeft25")
	symaddr(&addr_QDebug_OperatorShiftLeft26, "QDebug_OperatorShiftLeft26")
	symaddr(&addr_QDebug_OperatorShiftLeft27, "QDebug_OperatorShiftLeft27")
	symaddr(&addr_QDebug_OperatorShiftLeft28, "QDebug_OperatorShiftLeft28")
	symaddr(&addr_QDebug_OperatorShiftLeft29, "QDebug_OperatorShiftLeft29")
	symaddr(&addr_QDebug_OperatorShiftLeft30, "QDebug_OperatorShiftLeft30")
	symaddr(&addr_QDebug_OperatorShiftLeft31, "QDebug_OperatorShiftLeft31")
	symaddr(&addr_QDebug_OperatorShiftLeft32, "QDebug_OperatorShiftLeft32")
	symaddr(&addr_QDebug_OperatorShiftLeft33, "QDebug_OperatorShiftLeft33")
	symaddr(&addr_QDebug_OperatorShiftLeft34, "QDebug_OperatorShiftLeft34")
	symaddr(&addr_QDebug_OperatorShiftLeft35, "QDebug_OperatorShiftLeft35")
	symaddr(&addr_QDebug_OperatorShiftLeft36, "QDebug_OperatorShiftLeft36")
	symaddr(&addr_QDebug_OperatorShiftLeft37, "QDebug_OperatorShiftLeft37")
	symaddr(&addr_QDebug_OperatorShiftLeft38, "QDebug_OperatorShiftLeft38")
	symaddr(&addr_QDebug_OperatorShiftLeft39, "QDebug_OperatorShiftLeft39")
	symaddr(&addr_QDebug_OperatorShiftLeft40, "QDebug_OperatorShiftLeft40")
	symaddr(&addr_QDebug_OperatorShiftLeft41, "QDebug_OperatorShiftLeft41")
	symaddr(&addr_QDebug_OperatorShiftLeft42, "QDebug_OperatorShiftLeft42")
	symaddr(&addr_QDebug_OperatorShiftLeft43, "QDebug_OperatorShiftLeft43")
	symaddr(&addr_QDebug_OperatorShiftLeft44, "QDebug_OperatorShiftLeft44")
	symaddr(&addr_QDebug_OperatorShiftLeft45, "QDebug_OperatorShiftLeft45")
	symaddr(&addr_QDebug_OperatorShiftLeft46, "QDebug_OperatorShiftLeft46")
	symaddr(&addr_QDebug_OperatorShiftLeft47, "QDebug_OperatorShiftLeft47")
	symaddr(&addr_QDebug_OperatorShiftLeft48, "QDebug_OperatorShiftLeft48")
	symaddr(&addr_QDebug_OperatorShiftLeft49, "QDebug_OperatorShiftLeft49")
	symaddr(&addr_QDebug_OperatorShiftLeft50, "QDebug_OperatorShiftLeft50")
	symaddr(&addr_QDebug_OperatorShiftLeft51, "QDebug_OperatorShiftLeft51")
	symaddr(&addr_QDebug_OperatorShiftLeft52, "QDebug_OperatorShiftLeft52")
	symaddr(&addr_QDebug_OperatorShiftLeft53, "QDebug_OperatorShiftLeft53")
	symaddr(&addr_QDebug_OperatorShiftLeft54, "QDebug_OperatorShiftLeft54")
	symaddr(&addr_QDebug_OperatorShiftLeft55, "QDebug_OperatorShiftLeft55")
	symaddr(&addr_QDebug_OperatorShiftLeft56, "QDebug_OperatorShiftLeft56")
	symaddr(&addr_QDebug_OperatorShiftLeft57, "QDebug_OperatorShiftLeft57")
	symaddr(&addr_QDebug_OperatorShiftLeft58, "QDebug_OperatorShiftLeft58")
	symaddr(&addr_QDebug_OperatorShiftLeft59, "QDebug_OperatorShiftLeft59")
	symaddr(&addr_QDebug_OperatorShiftLeft60, "QDebug_OperatorShiftLeft60")
	symaddr(&addr_QDebug_OperatorShiftLeft61, "QDebug_OperatorShiftLeft61")
	symaddr(&addr_QDebug_OperatorShiftLeft62, "QDebug_OperatorShiftLeft62")
	symaddr(&addr_QDebug_OperatorShiftLeft63, "QDebug_OperatorShiftLeft63")
	symaddr(&addr_QDebug_OperatorShiftLeft64, "QDebug_OperatorShiftLeft64")
	symaddr(&addr_QDebug_OperatorShiftLeft65, "QDebug_OperatorShiftLeft65")
	symaddr(&addr_QDebug_OperatorShiftLeft66, "QDebug_OperatorShiftLeft66")
	symaddr(&addr_QDebug_OperatorShiftLeft67, "QDebug_OperatorShiftLeft67")
	symaddr(&addr_QDebug_OperatorShiftLeft68, "QDebug_OperatorShiftLeft68")
	symaddr(&addr_QDebug_OperatorShiftLeft69, "QDebug_OperatorShiftLeft69")
	symaddr(&addr_QDebug_OperatorShiftLeft70, "QDebug_OperatorShiftLeft70")
	symaddr(&addr_QDebug_OperatorShiftLeft71, "QDebug_OperatorShiftLeft71")
	symaddr(&addr_QDebug_OperatorShiftLeft72, "QDebug_OperatorShiftLeft72")
	symaddr(&addr_QDebug_OperatorShiftLeft73, "QDebug_OperatorShiftLeft73")
	symaddr(&addr_QDebug_OperatorShiftLeft74, "QDebug_OperatorShiftLeft74")
	symaddr(&addr_QDebug_OperatorShiftLeft75, "QDebug_OperatorShiftLeft75")
	symaddr(&addr_QDebug_OperatorShiftLeft76, "QDebug_OperatorShiftLeft76")
	symaddr(&addr_QDebug_OperatorShiftLeft77, "QDebug_OperatorShiftLeft77")
	symaddr(&addr_QDebug_OperatorShiftLeft78, "QDebug_OperatorShiftLeft78")
	symaddr(&addr_QDebug_OperatorShiftLeft79, "QDebug_OperatorShiftLeft79")
	symaddr(&addr_QDebug_OperatorShiftLeft80, "QDebug_OperatorShiftLeft80")
	symaddr(&addr_QDebug_OperatorShiftLeft81, "QDebug_OperatorShiftLeft81")
	symaddr(&addr_QDebug_OperatorShiftLeft82, "QDebug_OperatorShiftLeft82")
	symaddr(&addr_QDebug_OperatorShiftLeft83, "QDebug_OperatorShiftLeft83")
	symaddr(&addr_QDebug_OperatorShiftLeft84, "QDebug_OperatorShiftLeft84")
	symaddr(&addr_QDebug_OperatorShiftLeft85, "QDebug_OperatorShiftLeft85")
	symaddr(&addr_QDebug_OperatorShiftLeft86, "QDebug_OperatorShiftLeft86")
	symaddr(&addr_QDebug_OperatorShiftLeft87, "QDebug_OperatorShiftLeft87")
	symaddr(&addr_QDebug_OperatorShiftLeft88, "QDebug_OperatorShiftLeft88")
	symaddr(&addr_QDebug_OperatorShiftLeft89, "QDebug_OperatorShiftLeft89")
	symaddr(&addr_QDebug_OperatorShiftLeft90, "QDebug_OperatorShiftLeft90")
	symaddr(&addr_QDebug_OperatorShiftLeft91, "QDebug_OperatorShiftLeft91")
	symaddr(&addr_QDebug_OperatorShiftLeft92, "QDebug_OperatorShiftLeft92")
	symaddr(&addr_QDebug_OperatorShiftLeft93, "QDebug_OperatorShiftLeft93")
	symaddr(&addr_QDebug_OperatorShiftLeft94, "QDebug_OperatorShiftLeft94")
	symaddr(&addr_QDebug_OperatorShiftLeft95, "QDebug_OperatorShiftLeft95")
	symaddr(&addr_QDebug_OperatorShiftLeft96, "QDebug_OperatorShiftLeft96")
	symaddr(&addr_QDebug_OperatorShiftLeft97, "QDebug_OperatorShiftLeft97")
	symaddr(&addr_QDebug_OperatorShiftLeft98, "QDebug_OperatorShiftLeft98")
	symaddr(&addr_QDebug_OperatorShiftLeft99, "QDebug_OperatorShiftLeft99")
	symaddr(&addr_QDebug_OperatorShiftLeft100, "QDebug_OperatorShiftLeft100")
	symaddr(&addr_QDebug_OperatorShiftLeft101, "QDebug_OperatorShiftLeft101")
	symaddr(&addr_QDebug_OperatorShiftLeft102, "QDebug_OperatorShiftLeft102")
	symaddr(&addr_QDebug_OperatorShiftLeft103, "QDebug_OperatorShiftLeft103")
	symaddr(&addr_QDebug_OperatorShiftLeft104, "QDebug_OperatorShiftLeft104")
	symaddr(&addr_QDebug_OperatorShiftLeft105, "QDebug_OperatorShiftLeft105")
	symaddr(&addr_QDebug_OperatorShiftLeft106, "QDebug_OperatorShiftLeft106")
	symaddr(&addr_QDebug_OperatorShiftLeft107, "QDebug_OperatorShiftLeft107")
	symaddr(&addr_QDebug_OperatorShiftLeft108, "QDebug_OperatorShiftLeft108")
	symaddr(&addr_QDebug_OperatorShiftLeft109, "QDebug_OperatorShiftLeft109")
	symaddr(&addr_QDebug_OperatorShiftLeft110, "QDebug_OperatorShiftLeft110")
	symaddr(&addr_QDebug_OperatorShiftLeft111, "QDebug_OperatorShiftLeft111")
	symaddr(&addr_QDebug_OperatorShiftLeft112, "QDebug_OperatorShiftLeft112")
	symaddr(&addr_QDebug_OperatorShiftLeft113, "QDebug_OperatorShiftLeft113")
	symaddr(&addr_QDebug_OperatorShiftLeft114, "QDebug_OperatorShiftLeft114")
	symaddr(&addr_QDebug_OperatorShiftLeft115, "QDebug_OperatorShiftLeft115")
	symaddr(&addr_QDebug_OperatorShiftLeft116, "QDebug_OperatorShiftLeft116")
	symaddr(&addr_QDebug_OperatorShiftLeft117, "QDebug_OperatorShiftLeft117")
	symaddr(&addr_QDebug_OperatorShiftLeft118, "QDebug_OperatorShiftLeft118")
	symaddr(&addr_QDebug_OperatorShiftLeft119, "QDebug_OperatorShiftLeft119")
	symaddr(&addr_QDebug_OperatorShiftLeft120, "QDebug_OperatorShiftLeft120")
	symaddr(&addr_QDebug_OperatorShiftLeft121, "QDebug_OperatorShiftLeft121")
	symaddr(&addr_QDebug_OperatorShiftLeft122, "QDebug_OperatorShiftLeft122")
	symaddr(&addr_QDebug_OperatorShiftLeft123, "QDebug_OperatorShiftLeft123")
	symaddr(&addr_QDebug_OperatorShiftLeft124, "QDebug_OperatorShiftLeft124")
	symaddr(&addr_QDebug_OperatorShiftLeft125, "QDebug_OperatorShiftLeft125")
	symaddr(&addr_QDebug_OperatorShiftLeft126, "QDebug_OperatorShiftLeft126")
	symaddr(&addr_QDebug_OperatorShiftLeft127, "QDebug_OperatorShiftLeft127")
	reg(&pg_QDebug_delete, "QDebug_delete")
}

type QDebug struct {
	h uintptr
	*QIODeviceBase
}

func bindQDebug(h uintptr) *QDebug {
	if h == 0 {
		return nil
	}
	return &QDebug{h: h, QIODeviceBase: bindQIODeviceBase(pg_QDebug_base(h))}
}

func newownedQDebug(h uintptr) *QDebug {
	o := bindQDebug(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDebug).Delete)
	}
	return o
}

func (this *QDebug) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDebug) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDebug(h unsafe.Pointer) *QDebug { return bindQDebug(uintptr(h)) }
func UnsafeNewQDebugOwned(h unsafe.Pointer) *QDebug { return newownedQDebug(uintptr(h)) }

func (this *QDebug) UnsafeRelease() {
	this.h = 0
	if this.QIODeviceBase != nil {
		this.QIODeviceBase.UnsafeRelease()
	}
}

func NewQDebug(device *QIODevice) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_new, uintptr(device.UnsafePointer()))
	__o := newownedQDebug(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQDebug2(string_ *QString) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_new2, uintptr(string_.UnsafePointer()))
	__o := newownedQDebug(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQDebug3(bytes *QByteArray) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_new3, uintptr(bytes.UnsafePointer()))
	__o := newownedQDebug(__r)
	runtime.KeepAlive(bytes)
	return __o
}

func NewQDebug4(t QtMsgType) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_new4, uintptr(t))
	__o := newownedQDebug(__r)
	return __o
}

func NewQDebug5(o *QDebug) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_new5, uintptr(o.UnsafePointer()))
	__o := newownedQDebug(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QDebug) OperatorAssign(other *QDebug) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQDebug(__r)
}

func (this *QDebug) Swap(other *QDebug) {
	purego.SyscallN(addr_QDebug_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QDebug) ResetFormat() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_resetFormat, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) Space() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_space, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) Nospace() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_nospace, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) MaybeSpace() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_maybeSpace, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) Verbosity(verbosityLevel int) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_verbosity, this.h, uintptr(verbosityLevel))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) Verbosity2() int {
	__r0, _, _ := purego.SyscallN(addr_QDebug_verbosity2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDebug) SetVerbosity(verbosityLevel int) {
	purego.SyscallN(addr_QDebug_setVerbosity, this.h, uintptr(verbosityLevel))
	runtime.KeepAlive(this)
}

func (this *QDebug) AutoInsertSpaces() bool {
	__r0, _, _ := purego.SyscallN(addr_QDebug_autoInsertSpaces, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDebug) SetAutoInsertSpaces(b bool) {
	purego.SyscallN(addr_QDebug_setAutoInsertSpaces, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QDebug) QuoteStrings() bool {
	__r0, _, _ := purego.SyscallN(addr_QDebug_quoteStrings, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDebug) SetQuoteStrings(b bool) {
	purego.SyscallN(addr_QDebug_setQuoteStrings, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QDebug) Quote() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_quote, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) Noquote() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_noquote, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) MaybeQuote() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_maybeQuote, this.h)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft(t *QChar) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithBool(t bool) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft2, this.h, uintptr(b2i(t)))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithChar(t int8) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft3, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithShort(t int16) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft4, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithUnsignedshort(t uint16) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft5, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithChar16_t(t uint16) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft6, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithChar32_t(t uint32) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft7, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithInt(t int) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft8, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithUnsignedint(t uint) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft9, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithLong(t int32) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft10, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithUnsignedlong(t uint32) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft11, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWith(t int64) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft12, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft2(t uint64) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft13, this.h, uintptr(t))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQfloat16(t *qfloat16) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft14, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithFloat(t float32) *QDebug {
	__r := pg_QDebug_OperatorShiftLeft15(this.h, t)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithDouble(t float64) *QDebug {
	__r := pg_QDebug_OperatorShiftLeft16(this.h, t)
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft3(t string) *QDebug {
	t_b := append([]byte(t), 0)
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft17, this.h, uintptr(unsafe.Pointer(&t_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t_b)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQString(t string) *QDebug {
	var t_p uintptr
	if len(t) > 0 {
		t_p = uintptr(unsafe.Pointer(unsafe.StringData(t)))
	}
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft18, this.h, t_p, uintptr(len(t)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQStringView(s *QStringView) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft19, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft4(t *QLatin1String) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft20, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQByteArray(t *QByteArray) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft21, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQByteArrayView(t *QByteArrayView) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft22, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQTextStreamManipulator(m *QTextStreamManipulator) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft23, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return bindQDebug(__r)
}

func (this *QDebug) MaybeQuoteWithChar(c int8) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_maybeQuote2, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQAnyStringView(s string) *QDebug {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft24, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQMetaType(m *QMetaType) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft25, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQObject(arg1 *QObject) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft26, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithCombination(combination *QKeyCombination) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft27, this.h, uintptr(combination.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(combination)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQVariantType(arg1 QVariant__Type) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft28, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQModelIndex(arg1 *QModelIndex) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft29, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPersistentModelIndex(arg1 *QPersistentModelIndex) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft30, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQBitArray(arg1 *QBitArray) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft31, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithModes(modes QIODeviceBase__OpenModeFlag) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft32, this.h, uintptr(modes))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQLocale(arg1 *QLocale) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft33, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithSt(st QCborSimpleType) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft34, this.h, uintptr(st))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithTg(tg QCborKnownTags) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft35, this.h, uintptr(tg))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft5(tg QCborTag) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft36, this.h, uintptr(tg))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQDateTime(arg1 *QDateTime) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft37, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQDate(arg1 *QDate) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft38, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQTime(arg1 *QTime) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft39, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithRe(re *QRegularExpression) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft40, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithPatternOptions(patternOptions QRegularExpression__PatternOption) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft41, this.h, uintptr(patternOptions))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithMatch(match *QRegularExpressionMatch) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft42, this.h, uintptr(match.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(match)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQUrl(arg1 *QUrl) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft43, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQUuid(arg1 *QUuid) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft44, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQCborValue(v *QCborValue) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft45, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQCborArray(a *QCborArray) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft46, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQCborMap(m *QCborMap) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft47, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithTz(tz *QTimeZone) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft48, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQFileInfo(arg1 *QFileInfo) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft49, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithFilters(filters QDir__Filter) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft50, this.h, uintptr(filters))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithDir(dir *QDir) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft51, this.h, uintptr(dir.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dir)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithItem(item *QEasingCurve) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft52, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQItemSelectionRange(arg1 *QItemSelectionRange) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft53, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQJsonDocument(arg1 *QJsonDocument) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft54, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQJsonValue(arg1 *QJsonValue) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft55, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQJsonArray(arg1 *QJsonArray) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft56, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQJsonObject(arg1 *QJsonObject) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft57, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithRevision(revision *QTypeRevision) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft58, this.h, uintptr(revision.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(revision)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithVersion(version *QVersionNumber) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft59, this.h, uintptr(version.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(version)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPoint(arg1 *QPoint) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft60, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPointF(p *QPointF) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft61, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQLine(p *QLine) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft62, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQLineF(p *QLineF) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft63, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQMargins(arg1 *QMargins) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft64, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQMarginsF(arg1 *QMarginsF) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft65, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithMime(mime *QMimeType) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft66, this.h, uintptr(mime.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mime)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithOv(ov *QOperatingSystemVersion) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft67, this.h, uintptr(ov.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ov)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPermission(arg1 *QPermission) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft68, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQSize(arg1 *QSize) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft69, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQSizeF(arg1 *QSizeF) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft70, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQRect(arg1 *QRect) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft71, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQRectF(arg1 *QRectF) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft72, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQStorageInfo(arg1 *QStorageInfo) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft73, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQColor(arg1 *QColor) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft74, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithFormat(format *QPixelFormat) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft75, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPolygon(arg1 *QPolygon) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft76, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPolygonF(arg1 *QPolygonF) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft77, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQRegion(arg1 *QRegion) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft78, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQTransform(arg1 *QTransform) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft79, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQImage(arg1 *QImage) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft80, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPixmap(arg1 *QPixmap) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft81, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQIcon(arg1 *QIcon) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft82, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithVector(vector *QVector2D) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft83, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft6(vector *QVector3D) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft84, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft7(vector *QVector4D) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft85, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQScreen(arg1 *QScreen) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft86, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQInputDevice(arg1 *QInputDevice) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft87, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPointingDevice(arg1 *QPointingDevice) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft88, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQEventPoint(arg1 *QEventPoint) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft89, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft8(arg1 *QEventPoint) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft90, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQKeySequence(arg1 *QKeySequence) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft91, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQEvent(arg1 *QEvent) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft92, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithTag(tag *QFont__Tag) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft93, this.h, uintptr(tag.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQFont(arg1 *QFont) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft94, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQBrush(arg1 *QBrush) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft95, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPen(arg1 *QPen) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft96, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQTextLength(arg1 *QTextLength) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft97, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQTextFormat(arg1 *QTextFormat) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft98, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPalette(arg1 *QPalette) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft99, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithIface(iface *QAccessibleInterface) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft100, this.h, uintptr(iface.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iface)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithEv(ev *QAccessibleEvent) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft101, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQAction(arg1 *QAction) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft102, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQSurfaceFormat(arg1 *QSurfaceFormat) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft103, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQCursor(arg1 *QCursor) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft104, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQWindow(arg1 *QWindow) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft105, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithColorSpace(colorSpace *QColorSpace) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft106, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithAxis(axis *QFontVariableAxis) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft107, this.h, uintptr(axis.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(axis)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQMatrix4x4(m *QMatrix4x4) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft108, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithCtx(ctx *QOpenGLContext) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft109, this.h, uintptr(ctx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ctx)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithCg(cg *QOpenGLContextGroup) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft110, this.h, uintptr(cg.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cg)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithPageSize(pageSize *QPageSize) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft111, this.h, uintptr(pageSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageSize)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithPageLayout(pageLayout *QPageLayout) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft112, this.h, uintptr(pageLayout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageLayout)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithPageRanges(pageRanges *QPageRanges) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft113, this.h, uintptr(pageRanges.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageRanges)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQPainterPath(arg1 *QPainterPath) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft114, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQQuaternion(q *QQuaternion) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft115, this.h, uintptr(q.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQSizePolicy(arg1 *QSizePolicy) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft116, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQWidget(arg1 *QWidget) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft117, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithOption(option *QStyleOption) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft118, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithDockWidget(dockWidget *QDockWidget) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft119, this.h, uintptr(dockWidget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockWidget)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQGesture(arg1 *QGesture) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft120, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQGestureEvent(arg1 *QGestureEvent) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft121, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft9(item *QGraphicsItem) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft122, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeft10(item *QGraphicsObject) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft123, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithChange(change QGraphicsItem__GraphicsItemChange) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft124, this.h, uintptr(change))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithFlag(flag QGraphicsItem__GraphicsItemFlag) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft125, this.h, uintptr(flag))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithFlags(flags QGraphicsItem__GraphicsItemFlag) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft126, this.h, uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QDebug) OperatorShiftLeftWithQGraphicsSceneEvent(arg1 *QGraphicsSceneEvent) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QDebug_OperatorShiftLeft127, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQDebug(__r)
}

func (this *QDebug) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDebug_delete(this.h)
	}
	this.UnsafeRelease()
}
