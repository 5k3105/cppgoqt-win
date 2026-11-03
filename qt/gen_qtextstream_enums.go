package qt

type QTextStream__RealNumberNotation int64

const (
	QTextStream__SmartNotation QTextStream__RealNumberNotation = 0
	QTextStream__FixedNotation QTextStream__RealNumberNotation = 1
	QTextStream__ScientificNotation QTextStream__RealNumberNotation = 2
)

type QTextStream__FieldAlignment int64

const (
	QTextStream__AlignLeft QTextStream__FieldAlignment = 0
	QTextStream__AlignRight QTextStream__FieldAlignment = 1
	QTextStream__AlignCenter QTextStream__FieldAlignment = 2
	QTextStream__AlignAccountingStyle QTextStream__FieldAlignment = 3
)

type QTextStream__Status int64

const (
	QTextStream__Ok QTextStream__Status = 0
	QTextStream__ReadPastEnd QTextStream__Status = 1
	QTextStream__ReadCorruptData QTextStream__Status = 2
	QTextStream__WriteFailed QTextStream__Status = 3
)

type QTextStream__NumberFlag int64

const (
	QTextStream__ShowBase QTextStream__NumberFlag = 1
	QTextStream__ForcePoint QTextStream__NumberFlag = 2
	QTextStream__ForceSign QTextStream__NumberFlag = 4
	QTextStream__UppercaseBase QTextStream__NumberFlag = 8
	QTextStream__UppercaseDigits QTextStream__NumberFlag = 16
)

