package qt

type QtMsgType int64

const (
	QtDebugMsg QtMsgType = 0
	QtWarningMsg QtMsgType = 1
	QtCriticalMsg QtMsgType = 2
	QtFatalMsg QtMsgType = 3
	QtInfoMsg QtMsgType = 4
	QtSystemMsg QtMsgType = 2
)

type QCborSimpleType int64

const (
	False QCborSimpleType = 20
	True QCborSimpleType = 21
	Null QCborSimpleType = 22
	Undefined QCborSimpleType = 23
)

type QCborTag int64
type QCborKnownTags int64

const (
	DateTimeString QCborKnownTags = 0
	UnixTime_t QCborKnownTags = 1
	PositiveBignum QCborKnownTags = 2
	NegativeBignum QCborKnownTags = 3
	Decimal QCborKnownTags = 4
	Bigfloat QCborKnownTags = 5
	COSE_Encrypt0 QCborKnownTags = 16
	COSE_Mac0 QCborKnownTags = 17
	COSE_Sign1 QCborKnownTags = 18
	ExpectedBase64url QCborKnownTags = 21
	ExpectedBase64 QCborKnownTags = 22
	ExpectedBase16 QCborKnownTags = 23
	EncodedCbor QCborKnownTags = 24
	Url QCborKnownTags = 32
	Base64url QCborKnownTags = 33
	Base64 QCborKnownTags = 34
	RegularExpression QCborKnownTags = 35
	MimeMessage QCborKnownTags = 36
	Uuid QCborKnownTags = 37
	COSE_Encrypt QCborKnownTags = 96
	COSE_Mac QCborKnownTags = 97
	COSE_Sign QCborKnownTags = 98
	Signature QCborKnownTags = 55799
)

type QCborNegativeInteger int64
