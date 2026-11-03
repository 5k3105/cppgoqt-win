package qt

type QCborError__Code int64

const (
	QCborError__UnknownError QCborError__Code = 1
	QCborError__AdvancePastEnd QCborError__Code = 3
	QCborError__InputOutputError QCborError__Code = 4
	QCborError__GarbageAtEnd QCborError__Code = 256
	QCborError__EndOfFile QCborError__Code = 257
	QCborError__UnexpectedBreak QCborError__Code = 258
	QCborError__UnknownType QCborError__Code = 259
	QCborError__IllegalType QCborError__Code = 260
	QCborError__IllegalNumber QCborError__Code = 261
	QCborError__IllegalSimpleType QCborError__Code = 262
	QCborError__InvalidUtf8String QCborError__Code = 516
	QCborError__DataTooLarge QCborError__Code = 1024
	QCborError__NestingTooDeep QCborError__Code = 1025
	QCborError__UnsupportedType QCborError__Code = 1026
	QCborError__NoError QCborError__Code = 0
)

