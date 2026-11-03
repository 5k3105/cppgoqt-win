package qt

type QCborValue__EncodingOption int64

const (
	QCborValue__SortKeysInMaps QCborValue__EncodingOption = 1
	QCborValue__UseFloat QCborValue__EncodingOption = 2
	QCborValue__UseFloat16 QCborValue__EncodingOption = 6
	QCborValue__UseIntegers QCborValue__EncodingOption = 8
	QCborValue__NoTransformation QCborValue__EncodingOption = 0
)

type QCborValue__DiagnosticNotationOption int64

const (
	QCborValue__Compact QCborValue__DiagnosticNotationOption = 0
	QCborValue__LineWrapped QCborValue__DiagnosticNotationOption = 1
	QCborValue__ExtendedFormat QCborValue__DiagnosticNotationOption = 2
)

type QCborValue__Type int64

const (
	QCborValue__Integer QCborValue__Type = 0
	QCborValue__ByteArray QCborValue__Type = 64
	QCborValue__String QCborValue__Type = 96
	QCborValue__Array QCborValue__Type = 128
	QCborValue__Map QCborValue__Type = 160
	QCborValue__Tag QCborValue__Type = 192
	QCborValue__SimpleType QCborValue__Type = 256
	QCborValue__False QCborValue__Type = 276
	QCborValue__True QCborValue__Type = 277
	QCborValue__Null QCborValue__Type = 278
	QCborValue__Undefined QCborValue__Type = 279
	QCborValue__Double QCborValue__Type = 514
	QCborValue__DateTime QCborValue__Type = 65536
	QCborValue__Url QCborValue__Type = 65568
	QCborValue__RegularExpression QCborValue__Type = 65571
	QCborValue__Uuid QCborValue__Type = 65573
	QCborValue__Invalid QCborValue__Type = -1
)

