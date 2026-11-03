package qt

type QUuid__Variant int64

const (
	QUuid__VarUnknown QUuid__Variant = -1
	QUuid__NCS QUuid__Variant = 0
	QUuid__DCE QUuid__Variant = 2
	QUuid__Microsoft QUuid__Variant = 6
	QUuid__Reserved QUuid__Variant = 7
)

type QUuid__Version int64

const (
	QUuid__VerUnknown QUuid__Version = -1
	QUuid__Time QUuid__Version = 1
	QUuid__EmbeddedPOSIX QUuid__Version = 2
	QUuid__Md5 QUuid__Version = 3
	QUuid__Name QUuid__Version = 3
	QUuid__Random QUuid__Version = 4
	QUuid__Sha1 QUuid__Version = 5
	QUuid__UnixEpoch QUuid__Version = 7
)

type QUuid__StringFormat int64

const (
	QUuid__WithBraces QUuid__StringFormat = 0
	QUuid__WithoutBraces QUuid__StringFormat = 1
	QUuid__Id128 QUuid__StringFormat = 3
)

