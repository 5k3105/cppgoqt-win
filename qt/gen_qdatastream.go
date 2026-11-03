package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDataStream_base func(uintptr) uintptr
	addr_QDataStream_new uintptr
	addr_QDataStream_new2 uintptr
	addr_QDataStream_new3 uintptr
	addr_QDataStream_new4 uintptr
	addr_QDataStream_device uintptr
	addr_QDataStream_setDevice uintptr
	addr_QDataStream_atEnd uintptr
	addr_QDataStream_status uintptr
	addr_QDataStream_setStatus uintptr
	addr_QDataStream_resetStatus uintptr
	addr_QDataStream_floatingPointPrecision uintptr
	addr_QDataStream_setFloatingPointPrecision uintptr
	addr_QDataStream_byteOrder uintptr
	addr_QDataStream_setByteOrder uintptr
	addr_QDataStream_version uintptr
	addr_QDataStream_setVersion uintptr
	addr_QDataStream_OperatorShiftRight uintptr
	addr_QDataStream_OperatorShiftRight2 uintptr
	addr_QDataStream_OperatorShiftRight3 uintptr
	addr_QDataStream_OperatorShiftRight4 uintptr
	addr_QDataStream_OperatorShiftRight5 uintptr
	addr_QDataStream_OperatorShiftRight6 uintptr
	addr_QDataStream_OperatorShiftRight7 uintptr
	addr_QDataStream_OperatorShiftRight8 uintptr
	addr_QDataStream_OperatorShiftRight9 uintptr
	addr_QDataStream_OperatorShiftRight10 uintptr
	addr_QDataStream_OperatorShiftRight11 uintptr
	addr_QDataStream_OperatorShiftRight12 uintptr
	addr_QDataStream_OperatorShiftRight13 uintptr
	addr_QDataStream_OperatorShiftRight14 uintptr
	addr_QDataStream_OperatorShiftLeft uintptr
	addr_QDataStream_OperatorShiftLeft2 uintptr
	addr_QDataStream_OperatorShiftLeft3 uintptr
	addr_QDataStream_OperatorShiftLeft4 uintptr
	addr_QDataStream_OperatorShiftLeft5 uintptr
	addr_QDataStream_OperatorShiftLeft6 uintptr
	addr_QDataStream_OperatorShiftLeft7 uintptr
	addr_QDataStream_OperatorShiftLeft8 uintptr
	addr_QDataStream_OperatorShiftLeft9 uintptr
	pg_QDataStream_OperatorShiftLeft10 func(uintptr, float32) uintptr
	pg_QDataStream_OperatorShiftLeft11 func(uintptr, float64) uintptr
	addr_QDataStream_OperatorShiftLeft12 uintptr
	addr_QDataStream_OperatorShiftLeft13 uintptr
	addr_QDataStream_OperatorShiftLeft14 uintptr
	addr_QDataStream_ToBool uintptr
	addr_QDataStream_readRawData uintptr
	addr_QDataStream_writeBytes uintptr
	addr_QDataStream_writeRawData uintptr
	addr_QDataStream_skipRawData uintptr
	addr_QDataStream_startTransaction uintptr
	addr_QDataStream_commitTransaction uintptr
	addr_QDataStream_rollbackTransaction uintptr
	addr_QDataStream_abortTransaction uintptr
	addr_QDataStream_isDeviceTransactionStarted uintptr
	addr_QDataStream_OperatorShiftLeft15 uintptr
	addr_QDataStream_OperatorShiftRight15 uintptr
	addr_QDataStream_OperatorShiftLeft16 uintptr
	addr_QDataStream_OperatorShiftRight16 uintptr
	addr_QDataStream_OperatorShiftLeft17 uintptr
	addr_QDataStream_OperatorShiftRight17 uintptr
	addr_QDataStream_OperatorShiftRight18 uintptr
	addr_QDataStream_OperatorShiftLeft18 uintptr
	addr_QDataStream_OperatorShiftLeft19 uintptr
	addr_QDataStream_OperatorShiftRight19 uintptr
	addr_QDataStream_OperatorShiftRight20 uintptr
	addr_QDataStream_OperatorShiftLeft20 uintptr
	addr_QDataStream_OperatorShiftRight21 uintptr
	addr_QDataStream_OperatorShiftLeft21 uintptr
	addr_QDataStream_OperatorShiftLeft22 uintptr
	addr_QDataStream_OperatorShiftRight22 uintptr
	addr_QDataStream_OperatorShiftLeft23 uintptr
	addr_QDataStream_OperatorShiftRight23 uintptr
	addr_QDataStream_OperatorShiftLeft24 uintptr
	addr_QDataStream_OperatorShiftRight24 uintptr
	addr_QDataStream_OperatorShiftLeft25 uintptr
	addr_QDataStream_OperatorShiftRight25 uintptr
	addr_QDataStream_OperatorShiftLeft26 uintptr
	addr_QDataStream_OperatorShiftRight26 uintptr
	addr_QDataStream_OperatorShiftLeft27 uintptr
	addr_QDataStream_OperatorShiftRight27 uintptr
	addr_QDataStream_OperatorShiftLeft28 uintptr
	addr_QDataStream_OperatorShiftRight28 uintptr
	addr_QDataStream_OperatorShiftLeft29 uintptr
	addr_QDataStream_OperatorShiftRight29 uintptr
	addr_QDataStream_OperatorShiftLeft30 uintptr
	addr_QDataStream_OperatorShiftRight30 uintptr
	addr_QDataStream_OperatorShiftLeft31 uintptr
	addr_QDataStream_OperatorShiftRight31 uintptr
	addr_QDataStream_OperatorShiftLeft32 uintptr
	addr_QDataStream_OperatorShiftRight32 uintptr
	addr_QDataStream_OperatorShiftLeft33 uintptr
	addr_QDataStream_OperatorShiftRight33 uintptr
	addr_QDataStream_OperatorShiftLeft34 uintptr
	addr_QDataStream_OperatorShiftRight34 uintptr
	addr_QDataStream_OperatorShiftLeft35 uintptr
	addr_QDataStream_OperatorShiftRight35 uintptr
	addr_QDataStream_OperatorShiftLeft36 uintptr
	addr_QDataStream_OperatorShiftRight36 uintptr
	addr_QDataStream_OperatorShiftLeft37 uintptr
	addr_QDataStream_OperatorShiftRight37 uintptr
	addr_QDataStream_OperatorShiftLeft38 uintptr
	addr_QDataStream_OperatorShiftRight38 uintptr
	addr_QDataStream_OperatorShiftLeft39 uintptr
	addr_QDataStream_OperatorShiftRight39 uintptr
	addr_QDataStream_OperatorShiftLeft40 uintptr
	addr_QDataStream_OperatorShiftRight40 uintptr
	addr_QDataStream_OperatorShiftLeft41 uintptr
	addr_QDataStream_OperatorShiftRight41 uintptr
	addr_QDataStream_OperatorShiftLeft42 uintptr
	addr_QDataStream_OperatorShiftRight42 uintptr
	addr_QDataStream_OperatorShiftLeft43 uintptr
	addr_QDataStream_OperatorShiftRight43 uintptr
	addr_QDataStream_OperatorShiftLeft44 uintptr
	addr_QDataStream_OperatorShiftRight44 uintptr
	addr_QDataStream_OperatorShiftLeft45 uintptr
	addr_QDataStream_OperatorShiftRight45 uintptr
	addr_QDataStream_OperatorShiftLeft46 uintptr
	addr_QDataStream_OperatorShiftRight46 uintptr
	addr_QDataStream_OperatorShiftLeft47 uintptr
	addr_QDataStream_OperatorShiftRight47 uintptr
	addr_QDataStream_OperatorShiftLeft48 uintptr
	addr_QDataStream_OperatorShiftRight48 uintptr
	addr_QDataStream_OperatorShiftLeft49 uintptr
	addr_QDataStream_OperatorShiftRight49 uintptr
	addr_QDataStream_OperatorShiftLeft50 uintptr
	addr_QDataStream_OperatorShiftRight50 uintptr
	addr_QDataStream_OperatorShiftLeft51 uintptr
	addr_QDataStream_OperatorShiftRight51 uintptr
	addr_QDataStream_OperatorShiftLeft52 uintptr
	addr_QDataStream_OperatorShiftRight52 uintptr
	addr_QDataStream_OperatorShiftLeft53 uintptr
	addr_QDataStream_OperatorShiftRight53 uintptr
	addr_QDataStream_OperatorShiftLeft54 uintptr
	addr_QDataStream_OperatorShiftRight54 uintptr
	addr_QDataStream_OperatorShiftLeft55 uintptr
	addr_QDataStream_OperatorShiftRight55 uintptr
	addr_QDataStream_OperatorShiftLeft56 uintptr
	addr_QDataStream_OperatorShiftRight56 uintptr
	addr_QDataStream_OperatorShiftLeft57 uintptr
	addr_QDataStream_OperatorShiftRight57 uintptr
	addr_QDataStream_OperatorShiftRight58 uintptr
	addr_QDataStream_OperatorShiftLeft58 uintptr
	addr_QDataStream_OperatorShiftLeft59 uintptr
	addr_QDataStream_OperatorShiftRight59 uintptr
	addr_QDataStream_OperatorShiftLeft60 uintptr
	addr_QDataStream_OperatorShiftRight60 uintptr
	addr_QDataStream_OperatorShiftLeft61 uintptr
	addr_QDataStream_OperatorShiftRight61 uintptr
	addr_QDataStream_OperatorShiftLeft62 uintptr
	addr_QDataStream_OperatorShiftRight62 uintptr
	addr_QDataStream_OperatorShiftLeft63 uintptr
	addr_QDataStream_OperatorShiftRight63 uintptr
	addr_QDataStream_OperatorShiftLeft64 uintptr
	addr_QDataStream_OperatorShiftRight64 uintptr
	addr_QDataStream_OperatorShiftLeft65 uintptr
	addr_QDataStream_OperatorShiftRight65 uintptr
	addr_QDataStream_OperatorShiftLeft66 uintptr
	addr_QDataStream_OperatorShiftRight66 uintptr
	addr_QDataStream_OperatorShiftLeft67 uintptr
	addr_QDataStream_OperatorShiftRight67 uintptr
	addr_QDataStream_OperatorShiftLeft68 uintptr
	addr_QDataStream_OperatorShiftRight68 uintptr
	addr_QDataStream_OperatorShiftLeft69 uintptr
	addr_QDataStream_OperatorShiftRight69 uintptr
	addr_QDataStream_OperatorShiftLeft70 uintptr
	addr_QDataStream_OperatorShiftRight70 uintptr
	addr_QDataStream_OperatorShiftLeft71 uintptr
	addr_QDataStream_OperatorShiftRight71 uintptr
	addr_QDataStream_OperatorShiftLeft72 uintptr
	addr_QDataStream_OperatorShiftRight72 uintptr
	addr_QDataStream_OperatorShiftLeft73 uintptr
	addr_QDataStream_OperatorShiftRight73 uintptr
	addr_QDataStream_OperatorShiftLeft74 uintptr
	addr_QDataStream_OperatorShiftRight74 uintptr
	addr_QDataStream_OperatorShiftLeft75 uintptr
	addr_QDataStream_OperatorShiftRight75 uintptr
	addr_QDataStream_OperatorShiftLeft76 uintptr
	addr_QDataStream_OperatorShiftRight76 uintptr
	addr_QDataStream_OperatorShiftLeft77 uintptr
	addr_QDataStream_OperatorShiftRight77 uintptr
	addr_QDataStream_OperatorShiftLeft78 uintptr
	addr_QDataStream_OperatorShiftRight78 uintptr
	addr_QDataStream_OperatorShiftLeft79 uintptr
	addr_QDataStream_OperatorShiftRight79 uintptr
	addr_QDataStream_OperatorShiftLeft80 uintptr
	addr_QDataStream_OperatorShiftRight80 uintptr
	addr_QDataStream_OperatorShiftLeft81 uintptr
	addr_QDataStream_OperatorShiftRight81 uintptr
	addr_QDataStream_OperatorShiftLeft82 uintptr
	addr_QDataStream_OperatorShiftRight82 uintptr
	addr_QDataStream_OperatorShiftRight83 uintptr
	addr_QDataStream_OperatorShiftLeft83 uintptr
	addr_QDataStream_OperatorShiftLeft84 uintptr
	addr_QDataStream_OperatorShiftRight84 uintptr
	addr_QDataStream_OperatorShiftLeft85 uintptr
	addr_QDataStream_OperatorShiftRight85 uintptr
	addr_QDataStream_OperatorShiftRight86 uintptr
	addr_QDataStream_OperatorShiftLeft86 uintptr
	addr_QDataStream_OperatorShiftLeft87 uintptr
	addr_QDataStream_OperatorShiftRight87 uintptr
	pg_QDataStream_delete func(uintptr)
)

func init() {
	reg(&pg_QDataStream_base, "QDataStream_base")
	symaddr(&addr_QDataStream_new, "QDataStream_new")
	symaddr(&addr_QDataStream_new2, "QDataStream_new2")
	symaddr(&addr_QDataStream_new3, "QDataStream_new3")
	symaddr(&addr_QDataStream_new4, "QDataStream_new4")
	symaddr(&addr_QDataStream_device, "QDataStream_device")
	symaddr(&addr_QDataStream_setDevice, "QDataStream_setDevice")
	symaddr(&addr_QDataStream_atEnd, "QDataStream_atEnd")
	symaddr(&addr_QDataStream_status, "QDataStream_status")
	symaddr(&addr_QDataStream_setStatus, "QDataStream_setStatus")
	symaddr(&addr_QDataStream_resetStatus, "QDataStream_resetStatus")
	symaddr(&addr_QDataStream_floatingPointPrecision, "QDataStream_floatingPointPrecision")
	symaddr(&addr_QDataStream_setFloatingPointPrecision, "QDataStream_setFloatingPointPrecision")
	symaddr(&addr_QDataStream_byteOrder, "QDataStream_byteOrder")
	symaddr(&addr_QDataStream_setByteOrder, "QDataStream_setByteOrder")
	symaddr(&addr_QDataStream_version, "QDataStream_version")
	symaddr(&addr_QDataStream_setVersion, "QDataStream_setVersion")
	symaddr(&addr_QDataStream_OperatorShiftRight, "QDataStream_OperatorShiftRight")
	symaddr(&addr_QDataStream_OperatorShiftRight2, "QDataStream_OperatorShiftRight2")
	symaddr(&addr_QDataStream_OperatorShiftRight3, "QDataStream_OperatorShiftRight3")
	symaddr(&addr_QDataStream_OperatorShiftRight4, "QDataStream_OperatorShiftRight4")
	symaddr(&addr_QDataStream_OperatorShiftRight5, "QDataStream_OperatorShiftRight5")
	symaddr(&addr_QDataStream_OperatorShiftRight6, "QDataStream_OperatorShiftRight6")
	symaddr(&addr_QDataStream_OperatorShiftRight7, "QDataStream_OperatorShiftRight7")
	symaddr(&addr_QDataStream_OperatorShiftRight8, "QDataStream_OperatorShiftRight8")
	symaddr(&addr_QDataStream_OperatorShiftRight9, "QDataStream_OperatorShiftRight9")
	symaddr(&addr_QDataStream_OperatorShiftRight10, "QDataStream_OperatorShiftRight10")
	symaddr(&addr_QDataStream_OperatorShiftRight11, "QDataStream_OperatorShiftRight11")
	symaddr(&addr_QDataStream_OperatorShiftRight12, "QDataStream_OperatorShiftRight12")
	symaddr(&addr_QDataStream_OperatorShiftRight13, "QDataStream_OperatorShiftRight13")
	symaddr(&addr_QDataStream_OperatorShiftRight14, "QDataStream_OperatorShiftRight14")
	symaddr(&addr_QDataStream_OperatorShiftLeft, "QDataStream_OperatorShiftLeft")
	symaddr(&addr_QDataStream_OperatorShiftLeft2, "QDataStream_OperatorShiftLeft2")
	symaddr(&addr_QDataStream_OperatorShiftLeft3, "QDataStream_OperatorShiftLeft3")
	symaddr(&addr_QDataStream_OperatorShiftLeft4, "QDataStream_OperatorShiftLeft4")
	symaddr(&addr_QDataStream_OperatorShiftLeft5, "QDataStream_OperatorShiftLeft5")
	symaddr(&addr_QDataStream_OperatorShiftLeft6, "QDataStream_OperatorShiftLeft6")
	symaddr(&addr_QDataStream_OperatorShiftLeft7, "QDataStream_OperatorShiftLeft7")
	symaddr(&addr_QDataStream_OperatorShiftLeft8, "QDataStream_OperatorShiftLeft8")
	symaddr(&addr_QDataStream_OperatorShiftLeft9, "QDataStream_OperatorShiftLeft9")
	reg(&pg_QDataStream_OperatorShiftLeft10, "QDataStream_OperatorShiftLeft10")
	reg(&pg_QDataStream_OperatorShiftLeft11, "QDataStream_OperatorShiftLeft11")
	symaddr(&addr_QDataStream_OperatorShiftLeft12, "QDataStream_OperatorShiftLeft12")
	symaddr(&addr_QDataStream_OperatorShiftLeft13, "QDataStream_OperatorShiftLeft13")
	symaddr(&addr_QDataStream_OperatorShiftLeft14, "QDataStream_OperatorShiftLeft14")
	symaddr(&addr_QDataStream_ToBool, "QDataStream_ToBool")
	symaddr(&addr_QDataStream_readRawData, "QDataStream_readRawData")
	symaddr(&addr_QDataStream_writeBytes, "QDataStream_writeBytes")
	symaddr(&addr_QDataStream_writeRawData, "QDataStream_writeRawData")
	symaddr(&addr_QDataStream_skipRawData, "QDataStream_skipRawData")
	symaddr(&addr_QDataStream_startTransaction, "QDataStream_startTransaction")
	symaddr(&addr_QDataStream_commitTransaction, "QDataStream_commitTransaction")
	symaddr(&addr_QDataStream_rollbackTransaction, "QDataStream_rollbackTransaction")
	symaddr(&addr_QDataStream_abortTransaction, "QDataStream_abortTransaction")
	symaddr(&addr_QDataStream_isDeviceTransactionStarted, "QDataStream_isDeviceTransactionStarted")
	symaddr(&addr_QDataStream_OperatorShiftLeft15, "QDataStream_OperatorShiftLeft15")
	symaddr(&addr_QDataStream_OperatorShiftRight15, "QDataStream_OperatorShiftRight15")
	symaddr(&addr_QDataStream_OperatorShiftLeft16, "QDataStream_OperatorShiftLeft16")
	symaddr(&addr_QDataStream_OperatorShiftRight16, "QDataStream_OperatorShiftRight16")
	symaddr(&addr_QDataStream_OperatorShiftLeft17, "QDataStream_OperatorShiftLeft17")
	symaddr(&addr_QDataStream_OperatorShiftRight17, "QDataStream_OperatorShiftRight17")
	symaddr(&addr_QDataStream_OperatorShiftRight18, "QDataStream_OperatorShiftRight18")
	symaddr(&addr_QDataStream_OperatorShiftLeft18, "QDataStream_OperatorShiftLeft18")
	symaddr(&addr_QDataStream_OperatorShiftLeft19, "QDataStream_OperatorShiftLeft19")
	symaddr(&addr_QDataStream_OperatorShiftRight19, "QDataStream_OperatorShiftRight19")
	symaddr(&addr_QDataStream_OperatorShiftRight20, "QDataStream_OperatorShiftRight20")
	symaddr(&addr_QDataStream_OperatorShiftLeft20, "QDataStream_OperatorShiftLeft20")
	symaddr(&addr_QDataStream_OperatorShiftRight21, "QDataStream_OperatorShiftRight21")
	symaddr(&addr_QDataStream_OperatorShiftLeft21, "QDataStream_OperatorShiftLeft21")
	symaddr(&addr_QDataStream_OperatorShiftLeft22, "QDataStream_OperatorShiftLeft22")
	symaddr(&addr_QDataStream_OperatorShiftRight22, "QDataStream_OperatorShiftRight22")
	symaddr(&addr_QDataStream_OperatorShiftLeft23, "QDataStream_OperatorShiftLeft23")
	symaddr(&addr_QDataStream_OperatorShiftRight23, "QDataStream_OperatorShiftRight23")
	symaddr(&addr_QDataStream_OperatorShiftLeft24, "QDataStream_OperatorShiftLeft24")
	symaddr(&addr_QDataStream_OperatorShiftRight24, "QDataStream_OperatorShiftRight24")
	symaddr(&addr_QDataStream_OperatorShiftLeft25, "QDataStream_OperatorShiftLeft25")
	symaddr(&addr_QDataStream_OperatorShiftRight25, "QDataStream_OperatorShiftRight25")
	symaddr(&addr_QDataStream_OperatorShiftLeft26, "QDataStream_OperatorShiftLeft26")
	symaddr(&addr_QDataStream_OperatorShiftRight26, "QDataStream_OperatorShiftRight26")
	symaddr(&addr_QDataStream_OperatorShiftLeft27, "QDataStream_OperatorShiftLeft27")
	symaddr(&addr_QDataStream_OperatorShiftRight27, "QDataStream_OperatorShiftRight27")
	symaddr(&addr_QDataStream_OperatorShiftLeft28, "QDataStream_OperatorShiftLeft28")
	symaddr(&addr_QDataStream_OperatorShiftRight28, "QDataStream_OperatorShiftRight28")
	symaddr(&addr_QDataStream_OperatorShiftLeft29, "QDataStream_OperatorShiftLeft29")
	symaddr(&addr_QDataStream_OperatorShiftRight29, "QDataStream_OperatorShiftRight29")
	symaddr(&addr_QDataStream_OperatorShiftLeft30, "QDataStream_OperatorShiftLeft30")
	symaddr(&addr_QDataStream_OperatorShiftRight30, "QDataStream_OperatorShiftRight30")
	symaddr(&addr_QDataStream_OperatorShiftLeft31, "QDataStream_OperatorShiftLeft31")
	symaddr(&addr_QDataStream_OperatorShiftRight31, "QDataStream_OperatorShiftRight31")
	symaddr(&addr_QDataStream_OperatorShiftLeft32, "QDataStream_OperatorShiftLeft32")
	symaddr(&addr_QDataStream_OperatorShiftRight32, "QDataStream_OperatorShiftRight32")
	symaddr(&addr_QDataStream_OperatorShiftLeft33, "QDataStream_OperatorShiftLeft33")
	symaddr(&addr_QDataStream_OperatorShiftRight33, "QDataStream_OperatorShiftRight33")
	symaddr(&addr_QDataStream_OperatorShiftLeft34, "QDataStream_OperatorShiftLeft34")
	symaddr(&addr_QDataStream_OperatorShiftRight34, "QDataStream_OperatorShiftRight34")
	symaddr(&addr_QDataStream_OperatorShiftLeft35, "QDataStream_OperatorShiftLeft35")
	symaddr(&addr_QDataStream_OperatorShiftRight35, "QDataStream_OperatorShiftRight35")
	symaddr(&addr_QDataStream_OperatorShiftLeft36, "QDataStream_OperatorShiftLeft36")
	symaddr(&addr_QDataStream_OperatorShiftRight36, "QDataStream_OperatorShiftRight36")
	symaddr(&addr_QDataStream_OperatorShiftLeft37, "QDataStream_OperatorShiftLeft37")
	symaddr(&addr_QDataStream_OperatorShiftRight37, "QDataStream_OperatorShiftRight37")
	symaddr(&addr_QDataStream_OperatorShiftLeft38, "QDataStream_OperatorShiftLeft38")
	symaddr(&addr_QDataStream_OperatorShiftRight38, "QDataStream_OperatorShiftRight38")
	symaddr(&addr_QDataStream_OperatorShiftLeft39, "QDataStream_OperatorShiftLeft39")
	symaddr(&addr_QDataStream_OperatorShiftRight39, "QDataStream_OperatorShiftRight39")
	symaddr(&addr_QDataStream_OperatorShiftLeft40, "QDataStream_OperatorShiftLeft40")
	symaddr(&addr_QDataStream_OperatorShiftRight40, "QDataStream_OperatorShiftRight40")
	symaddr(&addr_QDataStream_OperatorShiftLeft41, "QDataStream_OperatorShiftLeft41")
	symaddr(&addr_QDataStream_OperatorShiftRight41, "QDataStream_OperatorShiftRight41")
	symaddr(&addr_QDataStream_OperatorShiftLeft42, "QDataStream_OperatorShiftLeft42")
	symaddr(&addr_QDataStream_OperatorShiftRight42, "QDataStream_OperatorShiftRight42")
	symaddr(&addr_QDataStream_OperatorShiftLeft43, "QDataStream_OperatorShiftLeft43")
	symaddr(&addr_QDataStream_OperatorShiftRight43, "QDataStream_OperatorShiftRight43")
	symaddr(&addr_QDataStream_OperatorShiftLeft44, "QDataStream_OperatorShiftLeft44")
	symaddr(&addr_QDataStream_OperatorShiftRight44, "QDataStream_OperatorShiftRight44")
	symaddr(&addr_QDataStream_OperatorShiftLeft45, "QDataStream_OperatorShiftLeft45")
	symaddr(&addr_QDataStream_OperatorShiftRight45, "QDataStream_OperatorShiftRight45")
	symaddr(&addr_QDataStream_OperatorShiftLeft46, "QDataStream_OperatorShiftLeft46")
	symaddr(&addr_QDataStream_OperatorShiftRight46, "QDataStream_OperatorShiftRight46")
	symaddr(&addr_QDataStream_OperatorShiftLeft47, "QDataStream_OperatorShiftLeft47")
	symaddr(&addr_QDataStream_OperatorShiftRight47, "QDataStream_OperatorShiftRight47")
	symaddr(&addr_QDataStream_OperatorShiftLeft48, "QDataStream_OperatorShiftLeft48")
	symaddr(&addr_QDataStream_OperatorShiftRight48, "QDataStream_OperatorShiftRight48")
	symaddr(&addr_QDataStream_OperatorShiftLeft49, "QDataStream_OperatorShiftLeft49")
	symaddr(&addr_QDataStream_OperatorShiftRight49, "QDataStream_OperatorShiftRight49")
	symaddr(&addr_QDataStream_OperatorShiftLeft50, "QDataStream_OperatorShiftLeft50")
	symaddr(&addr_QDataStream_OperatorShiftRight50, "QDataStream_OperatorShiftRight50")
	symaddr(&addr_QDataStream_OperatorShiftLeft51, "QDataStream_OperatorShiftLeft51")
	symaddr(&addr_QDataStream_OperatorShiftRight51, "QDataStream_OperatorShiftRight51")
	symaddr(&addr_QDataStream_OperatorShiftLeft52, "QDataStream_OperatorShiftLeft52")
	symaddr(&addr_QDataStream_OperatorShiftRight52, "QDataStream_OperatorShiftRight52")
	symaddr(&addr_QDataStream_OperatorShiftLeft53, "QDataStream_OperatorShiftLeft53")
	symaddr(&addr_QDataStream_OperatorShiftRight53, "QDataStream_OperatorShiftRight53")
	symaddr(&addr_QDataStream_OperatorShiftLeft54, "QDataStream_OperatorShiftLeft54")
	symaddr(&addr_QDataStream_OperatorShiftRight54, "QDataStream_OperatorShiftRight54")
	symaddr(&addr_QDataStream_OperatorShiftLeft55, "QDataStream_OperatorShiftLeft55")
	symaddr(&addr_QDataStream_OperatorShiftRight55, "QDataStream_OperatorShiftRight55")
	symaddr(&addr_QDataStream_OperatorShiftLeft56, "QDataStream_OperatorShiftLeft56")
	symaddr(&addr_QDataStream_OperatorShiftRight56, "QDataStream_OperatorShiftRight56")
	symaddr(&addr_QDataStream_OperatorShiftLeft57, "QDataStream_OperatorShiftLeft57")
	symaddr(&addr_QDataStream_OperatorShiftRight57, "QDataStream_OperatorShiftRight57")
	symaddr(&addr_QDataStream_OperatorShiftRight58, "QDataStream_OperatorShiftRight58")
	symaddr(&addr_QDataStream_OperatorShiftLeft58, "QDataStream_OperatorShiftLeft58")
	symaddr(&addr_QDataStream_OperatorShiftLeft59, "QDataStream_OperatorShiftLeft59")
	symaddr(&addr_QDataStream_OperatorShiftRight59, "QDataStream_OperatorShiftRight59")
	symaddr(&addr_QDataStream_OperatorShiftLeft60, "QDataStream_OperatorShiftLeft60")
	symaddr(&addr_QDataStream_OperatorShiftRight60, "QDataStream_OperatorShiftRight60")
	symaddr(&addr_QDataStream_OperatorShiftLeft61, "QDataStream_OperatorShiftLeft61")
	symaddr(&addr_QDataStream_OperatorShiftRight61, "QDataStream_OperatorShiftRight61")
	symaddr(&addr_QDataStream_OperatorShiftLeft62, "QDataStream_OperatorShiftLeft62")
	symaddr(&addr_QDataStream_OperatorShiftRight62, "QDataStream_OperatorShiftRight62")
	symaddr(&addr_QDataStream_OperatorShiftLeft63, "QDataStream_OperatorShiftLeft63")
	symaddr(&addr_QDataStream_OperatorShiftRight63, "QDataStream_OperatorShiftRight63")
	symaddr(&addr_QDataStream_OperatorShiftLeft64, "QDataStream_OperatorShiftLeft64")
	symaddr(&addr_QDataStream_OperatorShiftRight64, "QDataStream_OperatorShiftRight64")
	symaddr(&addr_QDataStream_OperatorShiftLeft65, "QDataStream_OperatorShiftLeft65")
	symaddr(&addr_QDataStream_OperatorShiftRight65, "QDataStream_OperatorShiftRight65")
	symaddr(&addr_QDataStream_OperatorShiftLeft66, "QDataStream_OperatorShiftLeft66")
	symaddr(&addr_QDataStream_OperatorShiftRight66, "QDataStream_OperatorShiftRight66")
	symaddr(&addr_QDataStream_OperatorShiftLeft67, "QDataStream_OperatorShiftLeft67")
	symaddr(&addr_QDataStream_OperatorShiftRight67, "QDataStream_OperatorShiftRight67")
	symaddr(&addr_QDataStream_OperatorShiftLeft68, "QDataStream_OperatorShiftLeft68")
	symaddr(&addr_QDataStream_OperatorShiftRight68, "QDataStream_OperatorShiftRight68")
	symaddr(&addr_QDataStream_OperatorShiftLeft69, "QDataStream_OperatorShiftLeft69")
	symaddr(&addr_QDataStream_OperatorShiftRight69, "QDataStream_OperatorShiftRight69")
	symaddr(&addr_QDataStream_OperatorShiftLeft70, "QDataStream_OperatorShiftLeft70")
	symaddr(&addr_QDataStream_OperatorShiftRight70, "QDataStream_OperatorShiftRight70")
	symaddr(&addr_QDataStream_OperatorShiftLeft71, "QDataStream_OperatorShiftLeft71")
	symaddr(&addr_QDataStream_OperatorShiftRight71, "QDataStream_OperatorShiftRight71")
	symaddr(&addr_QDataStream_OperatorShiftLeft72, "QDataStream_OperatorShiftLeft72")
	symaddr(&addr_QDataStream_OperatorShiftRight72, "QDataStream_OperatorShiftRight72")
	symaddr(&addr_QDataStream_OperatorShiftLeft73, "QDataStream_OperatorShiftLeft73")
	symaddr(&addr_QDataStream_OperatorShiftRight73, "QDataStream_OperatorShiftRight73")
	symaddr(&addr_QDataStream_OperatorShiftLeft74, "QDataStream_OperatorShiftLeft74")
	symaddr(&addr_QDataStream_OperatorShiftRight74, "QDataStream_OperatorShiftRight74")
	symaddr(&addr_QDataStream_OperatorShiftLeft75, "QDataStream_OperatorShiftLeft75")
	symaddr(&addr_QDataStream_OperatorShiftRight75, "QDataStream_OperatorShiftRight75")
	symaddr(&addr_QDataStream_OperatorShiftLeft76, "QDataStream_OperatorShiftLeft76")
	symaddr(&addr_QDataStream_OperatorShiftRight76, "QDataStream_OperatorShiftRight76")
	symaddr(&addr_QDataStream_OperatorShiftLeft77, "QDataStream_OperatorShiftLeft77")
	symaddr(&addr_QDataStream_OperatorShiftRight77, "QDataStream_OperatorShiftRight77")
	symaddr(&addr_QDataStream_OperatorShiftLeft78, "QDataStream_OperatorShiftLeft78")
	symaddr(&addr_QDataStream_OperatorShiftRight78, "QDataStream_OperatorShiftRight78")
	symaddr(&addr_QDataStream_OperatorShiftLeft79, "QDataStream_OperatorShiftLeft79")
	symaddr(&addr_QDataStream_OperatorShiftRight79, "QDataStream_OperatorShiftRight79")
	symaddr(&addr_QDataStream_OperatorShiftLeft80, "QDataStream_OperatorShiftLeft80")
	symaddr(&addr_QDataStream_OperatorShiftRight80, "QDataStream_OperatorShiftRight80")
	symaddr(&addr_QDataStream_OperatorShiftLeft81, "QDataStream_OperatorShiftLeft81")
	symaddr(&addr_QDataStream_OperatorShiftRight81, "QDataStream_OperatorShiftRight81")
	symaddr(&addr_QDataStream_OperatorShiftLeft82, "QDataStream_OperatorShiftLeft82")
	symaddr(&addr_QDataStream_OperatorShiftRight82, "QDataStream_OperatorShiftRight82")
	symaddr(&addr_QDataStream_OperatorShiftRight83, "QDataStream_OperatorShiftRight83")
	symaddr(&addr_QDataStream_OperatorShiftLeft83, "QDataStream_OperatorShiftLeft83")
	symaddr(&addr_QDataStream_OperatorShiftLeft84, "QDataStream_OperatorShiftLeft84")
	symaddr(&addr_QDataStream_OperatorShiftRight84, "QDataStream_OperatorShiftRight84")
	symaddr(&addr_QDataStream_OperatorShiftLeft85, "QDataStream_OperatorShiftLeft85")
	symaddr(&addr_QDataStream_OperatorShiftRight85, "QDataStream_OperatorShiftRight85")
	symaddr(&addr_QDataStream_OperatorShiftRight86, "QDataStream_OperatorShiftRight86")
	symaddr(&addr_QDataStream_OperatorShiftLeft86, "QDataStream_OperatorShiftLeft86")
	symaddr(&addr_QDataStream_OperatorShiftLeft87, "QDataStream_OperatorShiftLeft87")
	symaddr(&addr_QDataStream_OperatorShiftRight87, "QDataStream_OperatorShiftRight87")
	reg(&pg_QDataStream_delete, "QDataStream_delete")
}

type QDataStream struct {
	h uintptr
	*QIODeviceBase
}

func bindQDataStream(h uintptr) *QDataStream {
	if h == 0 {
		return nil
	}
	return &QDataStream{h: h, QIODeviceBase: bindQIODeviceBase(pg_QDataStream_base(h))}
}

func newownedQDataStream(h uintptr) *QDataStream {
	o := bindQDataStream(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDataStream).Delete)
	}
	return o
}

func (this *QDataStream) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDataStream) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDataStream(h unsafe.Pointer) *QDataStream { return bindQDataStream(uintptr(h)) }
func UnsafeNewQDataStreamOwned(h unsafe.Pointer) *QDataStream { return newownedQDataStream(uintptr(h)) }

func (this *QDataStream) UnsafeRelease() {
	this.h = 0
	if this.QIODeviceBase != nil {
		this.QIODeviceBase.UnsafeRelease()
	}
}

func NewQDataStream() *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_new)
	__o := newownedQDataStream(__r)
	return __o
}

func NewQDataStream2(arg1 *QIODevice) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQDataStream(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQDataStream3(arg1 *QByteArray, flags QIODeviceBase__OpenModeFlag) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_new3, uintptr(arg1.UnsafePointer()), uintptr(flags))
	__o := newownedQDataStream(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQDataStream4(arg1 *QByteArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQDataStream(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QDataStream) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QDataStream_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QDataStream) SetDevice(arg1 *QIODevice) {
	purego.SyscallN(addr_QDataStream_setDevice, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDataStream) AtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_atEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDataStream) Status() QDataStream__Status {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_status, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDataStream__Status(__r)
}

func (this *QDataStream) SetStatus(status QDataStream__Status) {
	purego.SyscallN(addr_QDataStream_setStatus, this.h, uintptr(status))
	runtime.KeepAlive(this)
}

func (this *QDataStream) ResetStatus() {
	purego.SyscallN(addr_QDataStream_resetStatus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataStream) FloatingPointPrecision() QDataStream__FloatingPointPrecision {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_floatingPointPrecision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDataStream__FloatingPointPrecision(__r)
}

func (this *QDataStream) SetFloatingPointPrecision(precision QDataStream__FloatingPointPrecision) {
	purego.SyscallN(addr_QDataStream_setFloatingPointPrecision, this.h, uintptr(precision))
	runtime.KeepAlive(this)
}

func (this *QDataStream) ByteOrder() QDataStream__ByteOrder {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_byteOrder, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDataStream__ByteOrder(__r)
}

func (this *QDataStream) SetByteOrder(arg1 QDataStream__ByteOrder) {
	purego.SyscallN(addr_QDataStream_setByteOrder, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDataStream) Version() int {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_version, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDataStream) SetVersion(arg1 int) {
	purego.SyscallN(addr_QDataStream_setVersion, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDataStream) OperatorShiftRight(i int8) (*QDataStream, int8) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), i_io
}

func (this *QDataStream) OperatorShiftRightWith(i int8) (*QDataStream, int8) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight2, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), i_io
}

func (this *QDataStream) OperatorShiftRight2(i byte) (*QDataStream, byte) {
	i_io := uint8(i)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight3, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), byte(i_io)
}

func (this *QDataStream) OperatorShiftRight3(i int16) (*QDataStream, int16) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight4, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), i_io
}

func (this *QDataStream) OperatorShiftRight4(i uint16) (*QDataStream, uint16) {
	i_io := int16(i)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight5, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), uint16(i_io)
}

func (this *QDataStream) OperatorShiftRight5(i int) (*QDataStream, int) {
	i_io := int32(i)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight6, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), int(i_io)
}

func (this *QDataStream) OperatorShiftRight6(i uint) (*QDataStream, uint) {
	i_io := int32(i)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight7, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), uint(i_io)
}

func (this *QDataStream) OperatorShiftRight7(i int64) (*QDataStream, int64) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight8, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), i_io
}

func (this *QDataStream) OperatorShiftRight8(i uint64) (*QDataStream, uint64) {
	i_io := int64(i)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight9, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), uint64(i_io)
}

func (this *QDataStream) OperatorShiftRightWithBool(i bool) (*QDataStream, bool) {
	i_io := b2i(i)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight10, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), i_io != 0
}

func (this *QDataStream) OperatorShiftRightWithFloat(f float32) (*QDataStream, float32) {
	f_io := f
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight11, this.h, uintptr(unsafe.Pointer(&f_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), f_io
}

func (this *QDataStream) OperatorShiftRightWithDouble(f float64) (*QDataStream, float64) {
	f_io := f
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight12, this.h, uintptr(unsafe.Pointer(&f_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), f_io
}

func (this *QDataStream) OperatorShiftRightWithChar16_t(c uint16) (*QDataStream, uint16) {
	c_io := int16(c)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight13, this.h, uintptr(unsafe.Pointer(&c_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), uint16(c_io)
}

func (this *QDataStream) OperatorShiftRightWithChar32_t(c uint32) (*QDataStream, uint32) {
	c_io := int32(c)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight14, this.h, uintptr(unsafe.Pointer(&c_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), uint32(c_io)
}

func (this *QDataStream) OperatorShiftLeft(i int8) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWith(i int8) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft2(i byte) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft3, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft3(i int16) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft4, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft4(i uint16) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft5, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft5(i int) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft6, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft6(i uint) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft7, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft7(i int64) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft8, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft8(i uint64) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft9, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithFloat(f float32) *QDataStream {
	__r := pg_QDataStream_OperatorShiftLeft10(this.h, f)
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithDouble(f float64) *QDataStream {
	__r := pg_QDataStream_OperatorShiftLeft11(this.h, f)
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithStr(str string) *QDataStream {
	str_b := append([]byte(str), 0)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft12, this.h, uintptr(unsafe.Pointer(&str_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str_b)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithChar16_t(c uint16) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft13, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithChar32_t(c uint32) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft14, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_ToBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDataStream) ReadRawData(arg1 []byte, len_ int64) int64 {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QDataStream_readRawData, this.h, arg1_p, uintptr(len_))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func (this *QDataStream) WriteBytes(arg1 string, len_ int64) *QDataStream {
	arg1_b := append([]byte(arg1), 0)
	__r, _, _ := purego.SyscallN(addr_QDataStream_writeBytes, this.h, uintptr(unsafe.Pointer(&arg1_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1_b)
	return bindQDataStream(__r)
}

func (this *QDataStream) WriteRawData(arg1 string, len_ int64) int64 {
	arg1_b := append([]byte(arg1), 0)
	__r0, _, _ := purego.SyscallN(addr_QDataStream_writeRawData, this.h, uintptr(unsafe.Pointer(&arg1_b[0])), uintptr(len_))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1_b)
	return __r
}

func (this *QDataStream) SkipRawData(len_ int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_skipRawData, this.h, uintptr(len_))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDataStream) StartTransaction() {
	purego.SyscallN(addr_QDataStream_startTransaction, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataStream) CommitTransaction() bool {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_commitTransaction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDataStream) RollbackTransaction() {
	purego.SyscallN(addr_QDataStream_rollbackTransaction, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataStream) AbortTransaction() {
	purego.SyscallN(addr_QDataStream_abortTransaction, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataStream) IsDeviceTransactionStarted() bool {
	__r0, _, _ := purego.SyscallN(addr_QDataStream_isDeviceTransactionStarted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDataStream) OperatorShiftLeftWithQByteArray(arg1 *QByteArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft15, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQByteArray(arg1 *QByteArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight15, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQString(arg1 string) *QDataStream {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft16, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQString(arg1 *QString) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight16, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithChr(chr *QChar) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft17, this.h, uintptr(chr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(chr)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithChr(chr *QChar) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight17, this.h, uintptr(chr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(chr)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithCombination(combination *QKeyCombination) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight18, this.h, uintptr(combination.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(combination)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithCombination(combination *QKeyCombination) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft18, this.h, uintptr(combination.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(combination)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQfloat16(f *qfloat16) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft19, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQfloat16(f *qfloat16) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight19, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQVariant(p *QVariant) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight20, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQVariant(p *QVariant) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft20, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQVariantType(p QVariant__Type) (*QDataStream, QVariant__Type) {
	p_io := int32(p)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight21, this.h, uintptr(unsafe.Pointer(&p_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), QVariant__Type(p_io)
}

func (this *QDataStream) OperatorShiftLeftWithQVariantType(p QVariant__Type) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft21, this.h, uintptr(p))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQBitArray(arg1 *QBitArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft22, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQBitArray(arg1 *QBitArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight22, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQLocale(arg1 *QLocale) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft23, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQLocale(arg1 *QLocale) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight23, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithSt(st QCborSimpleType) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft24, this.h, uintptr(st))
	runtime.KeepAlive(this)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithSt(st QCborSimpleType) (*QDataStream, QCborSimpleType) {
	st_io := int32(st)
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight24, this.h, uintptr(unsafe.Pointer(&st_io)))
	runtime.KeepAlive(this)
	return bindQDataStream(__r), QCborSimpleType(st_io)
}

func (this *QDataStream) OperatorShiftLeftWithQDate(arg1 *QDate) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft25, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQDate(arg1 *QDate) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight25, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTime(arg1 *QTime) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft26, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTime(arg1 *QTime) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight26, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQDateTime(arg1 *QDateTime) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft27, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQDateTime(arg1 *QDateTime) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight27, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithRe(re *QRegularExpression) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft28, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithRe(re *QRegularExpression) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight28, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQUrl(arg1 *QUrl) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft29, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQUrl(arg1 *QUrl) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight29, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQUuid(arg1 *QUuid) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft30, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQUuid(arg1 *QUuid) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight30, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQCborValue(arg1 *QCborValue) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft31, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQCborValue(arg1 *QCborValue) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight31, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQCborArray(arg1 *QCborArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft32, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQCborArray(arg1 *QCborArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight32, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQCborMap(arg1 *QCborMap) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft33, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQCborMap(arg1 *QCborMap) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight33, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithTz(tz *QTimeZone) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft34, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithTz(tz *QTimeZone) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight34, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQEasingCurve(arg1 *QEasingCurve) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft35, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQEasingCurve(arg1 *QEasingCurve) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight35, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQJsonDocument(arg1 *QJsonDocument) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft36, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQJsonDocument(arg1 *QJsonDocument) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight36, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQJsonValue(arg1 *QJsonValue) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft37, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQJsonValue(arg1 *QJsonValue) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight37, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQJsonArray(arg1 *QJsonArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft38, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQJsonArray(arg1 *QJsonArray) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight38, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQJsonObject(arg1 *QJsonObject) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft39, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQJsonObject(arg1 *QJsonObject) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight39, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithRevision(revision *QTypeRevision) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft40, this.h, uintptr(revision.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(revision)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithRevision(revision *QTypeRevision) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight40, this.h, uintptr(revision.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(revision)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithVersion(version *QVersionNumber) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft41, this.h, uintptr(version.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(version)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithVersion(version *QVersionNumber) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight41, this.h, uintptr(version.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(version)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPoint(arg1 *QPoint) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft42, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPoint(arg1 *QPoint) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight42, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPointF(arg1 *QPointF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft43, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPointF(arg1 *QPointF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight43, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQLine(arg1 *QLine) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft44, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQLine(arg1 *QLine) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight44, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQLineF(arg1 *QLineF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft45, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQLineF(arg1 *QLineF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight45, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQMargins(arg1 *QMargins) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft46, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQMargins(arg1 *QMargins) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight46, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQMarginsF(arg1 *QMarginsF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft47, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQMarginsF(arg1 *QMarginsF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight47, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQSize(arg1 *QSize) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft48, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQSize(arg1 *QSize) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight48, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQSizeF(arg1 *QSizeF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft49, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQSizeF(arg1 *QSizeF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight49, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQRect(arg1 *QRect) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft50, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQRect(arg1 *QRect) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight50, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQRectF(arg1 *QRectF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft51, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQRectF(arg1 *QRectF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight51, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQColor(arg1 *QColor) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft52, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQColor(arg1 *QColor) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight52, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithPolygon(polygon *QPolygon) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft53, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithPolygon(polygon *QPolygon) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight53, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithArray(array *QPolygonF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft54, this.h, uintptr(array.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(array)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithArray(array *QPolygonF) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight54, this.h, uintptr(array.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(array)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQRegion(arg1 *QRegion) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft55, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQRegion(arg1 *QRegion) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight55, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTransform(arg1 *QTransform) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft56, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTransform(arg1 *QTransform) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight56, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQImage(arg1 *QImage) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft57, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQImage(arg1 *QImage) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight57, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPixmap(arg1 *QPixmap) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight58, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPixmap(arg1 *QPixmap) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft58, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQIcon(arg1 *QIcon) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft59, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQIcon(arg1 *QIcon) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight59, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQVector2D(arg1 *QVector2D) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft60, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQVector2D(arg1 *QVector2D) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight60, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQVector3D(arg1 *QVector3D) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft61, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQVector3D(arg1 *QVector3D) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight61, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQVector4D(arg1 *QVector4D) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft62, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQVector4D(arg1 *QVector4D) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight62, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithKs(ks *QKeySequence) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft63, this.h, uintptr(ks.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ks)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithKs(ks *QKeySequence) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight63, this.h, uintptr(ks.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ks)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQFontTag(arg1 *QFont__Tag) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft64, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQFontTag(arg1 *QFont__Tag) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight64, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQFont(arg1 *QFont) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft65, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQFont(arg1 *QFont) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight65, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQBrush(arg1 *QBrush) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft66, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQBrush(arg1 *QBrush) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight66, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPen(arg1 *QPen) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft67, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPen(arg1 *QPen) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight67, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextLength(arg1 *QTextLength) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft68, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextLength(arg1 *QTextLength) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight68, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextFormat(arg1 *QTextFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft69, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextFormat(arg1 *QTextFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight69, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextCharFormat(arg1 *QTextCharFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft70, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextCharFormat(arg1 *QTextCharFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight70, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextBlockFormat(arg1 *QTextBlockFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft71, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextBlockFormat(arg1 *QTextBlockFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight71, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextListFormat(arg1 *QTextListFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft72, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextListFormat(arg1 *QTextListFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight72, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextFrameFormat(arg1 *QTextFrameFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft73, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextFrameFormat(arg1 *QTextFrameFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight73, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQTextTableCellFormat(arg1 *QTextTableCellFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft74, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQTextTableCellFormat(arg1 *QTextTableCellFormat) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight74, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPalette(p *QPalette) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft75, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPalette(p *QPalette) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight75, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithCursor(cursor *QCursor) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft76, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithCursor(cursor *QCursor) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight76, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQColorSpace(arg1 *QColorSpace) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft77, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQColorSpace(arg1 *QColorSpace) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight77, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQMatrix4x4(arg1 *QMatrix4x4) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft78, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQMatrix4x4(arg1 *QMatrix4x4) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight78, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPageRanges(arg1 *QPageRanges) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft79, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPageRanges(arg1 *QPageRanges) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight79, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPainterPath(arg1 *QPainterPath) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft80, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPainterPath(arg1 *QPainterPath) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight80, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQPicture(p *QPicture) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft81, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQPicture(p *QPicture) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight81, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQQuaternion(arg1 *QQuaternion) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft82, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQQuaternion(arg1 *QQuaternion) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight82, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithItem(item *QStandardItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight83, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithItem(item *QStandardItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft83, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeftWithQSizePolicy(arg1 *QSizePolicy) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft84, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRightWithQSizePolicy(arg1 *QSizePolicy) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight84, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft9(item *QListWidgetItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft85, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRight9(item *QListWidgetItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight85, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRight10(item *QTableWidgetItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight86, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft10(item *QTableWidgetItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft86, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftLeft11(item *QTreeWidgetItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftLeft87, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) OperatorShiftRight11(item *QTreeWidgetItem) *QDataStream {
	__r, _, _ := purego.SyscallN(addr_QDataStream_OperatorShiftRight87, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQDataStream(__r)
}

func (this *QDataStream) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDataStream_delete(this.h)
	}
	this.UnsafeRelease()
}
