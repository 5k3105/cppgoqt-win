package qt

type QCborStreamReader__Type int64

const (
	QCborStreamReader__UnsignedInteger QCborStreamReader__Type = 0
	QCborStreamReader__NegativeInteger QCborStreamReader__Type = 32
	QCborStreamReader__ByteString QCborStreamReader__Type = 64
	QCborStreamReader__ByteArray QCborStreamReader__Type = 64
	QCborStreamReader__TextString QCborStreamReader__Type = 96
	QCborStreamReader__String QCborStreamReader__Type = 96
	QCborStreamReader__Array QCborStreamReader__Type = 128
	QCborStreamReader__Map QCborStreamReader__Type = 160
	QCborStreamReader__Tag QCborStreamReader__Type = 192
	QCborStreamReader__SimpleType QCborStreamReader__Type = 224
	QCborStreamReader__HalfFloat QCborStreamReader__Type = 249
	QCborStreamReader__Float16 QCborStreamReader__Type = 249
	QCborStreamReader__Float QCborStreamReader__Type = 250
	QCborStreamReader__Double QCborStreamReader__Type = 251
	QCborStreamReader__Invalid QCborStreamReader__Type = 255
)

type QCborStreamReader__StringResultCode int64

const (
	QCborStreamReader__EndOfString QCborStreamReader__StringResultCode = 0
	QCborStreamReader__Ok QCborStreamReader__StringResultCode = 1
	QCborStreamReader__Error QCborStreamReader__StringResultCode = -1
)

