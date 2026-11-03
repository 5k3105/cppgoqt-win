package qt

type QStringConverter__Encoding int64

const (
	QStringConverter__Utf8 QStringConverter__Encoding = 0
	QStringConverter__Utf16 QStringConverter__Encoding = 1
	QStringConverter__Utf16LE QStringConverter__Encoding = 2
	QStringConverter__Utf16BE QStringConverter__Encoding = 3
	QStringConverter__Utf32 QStringConverter__Encoding = 4
	QStringConverter__Utf32LE QStringConverter__Encoding = 5
	QStringConverter__Utf32BE QStringConverter__Encoding = 6
	QStringConverter__Latin1 QStringConverter__Encoding = 7
	QStringConverter__System QStringConverter__Encoding = 8
	QStringConverter__LastEncoding QStringConverter__Encoding = 8
)

type QStringConverter__FinalizeResultError int64

const (
	QStringConverter__FinalizeResultError__NoError QStringConverter__FinalizeResultError = 0
	QStringConverter__FinalizeResultError__InvalidCharacters QStringConverter__FinalizeResultError = 1
	QStringConverter__FinalizeResultError__NotEnoughSpace QStringConverter__FinalizeResultError = 2
)

