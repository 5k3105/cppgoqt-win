package qt

type QSettings__Status int64

const (
	QSettings__NoError QSettings__Status = 0
	QSettings__AccessError QSettings__Status = 1
	QSettings__FormatError QSettings__Status = 2
)

type QSettings__Format int64

const (
	QSettings__NativeFormat QSettings__Format = 0
	QSettings__IniFormat QSettings__Format = 1
	QSettings__Registry32Format QSettings__Format = 2
	QSettings__Registry64Format QSettings__Format = 3
	QSettings__InvalidFormat QSettings__Format = 16
	QSettings__CustomFormat1 QSettings__Format = 17
	QSettings__CustomFormat2 QSettings__Format = 18
	QSettings__CustomFormat3 QSettings__Format = 19
	QSettings__CustomFormat4 QSettings__Format = 20
	QSettings__CustomFormat5 QSettings__Format = 21
	QSettings__CustomFormat6 QSettings__Format = 22
	QSettings__CustomFormat7 QSettings__Format = 23
	QSettings__CustomFormat8 QSettings__Format = 24
	QSettings__CustomFormat9 QSettings__Format = 25
	QSettings__CustomFormat10 QSettings__Format = 26
	QSettings__CustomFormat11 QSettings__Format = 27
	QSettings__CustomFormat12 QSettings__Format = 28
	QSettings__CustomFormat13 QSettings__Format = 29
	QSettings__CustomFormat14 QSettings__Format = 30
	QSettings__CustomFormat15 QSettings__Format = 31
	QSettings__CustomFormat16 QSettings__Format = 32
)

type QSettings__Scope int64

const (
	QSettings__UserScope QSettings__Scope = 0
	QSettings__SystemScope QSettings__Scope = 1
)

