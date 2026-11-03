package qt

type QXmlStreamReader__TokenType int64

const (
	QXmlStreamReader__NoToken QXmlStreamReader__TokenType = 0
	QXmlStreamReader__Invalid QXmlStreamReader__TokenType = 1
	QXmlStreamReader__StartDocument QXmlStreamReader__TokenType = 2
	QXmlStreamReader__EndDocument QXmlStreamReader__TokenType = 3
	QXmlStreamReader__StartElement QXmlStreamReader__TokenType = 4
	QXmlStreamReader__EndElement QXmlStreamReader__TokenType = 5
	QXmlStreamReader__Characters QXmlStreamReader__TokenType = 6
	QXmlStreamReader__Comment QXmlStreamReader__TokenType = 7
	QXmlStreamReader__DTD QXmlStreamReader__TokenType = 8
	QXmlStreamReader__EntityReference QXmlStreamReader__TokenType = 9
	QXmlStreamReader__ProcessingInstruction QXmlStreamReader__TokenType = 10
)

type QXmlStreamReader__ReadElementTextBehaviour int64

const (
	QXmlStreamReader__ErrorOnUnexpectedElement QXmlStreamReader__ReadElementTextBehaviour = 0
	QXmlStreamReader__IncludeChildElements QXmlStreamReader__ReadElementTextBehaviour = 1
	QXmlStreamReader__SkipChildElements QXmlStreamReader__ReadElementTextBehaviour = 2
)

type QXmlStreamReader__Error int64

const (
	QXmlStreamReader__NoError QXmlStreamReader__Error = 0
	QXmlStreamReader__UnexpectedElementError QXmlStreamReader__Error = 1
	QXmlStreamReader__CustomError QXmlStreamReader__Error = 2
	QXmlStreamReader__NotWellFormedError QXmlStreamReader__Error = 3
	QXmlStreamReader__PrematureEndOfDocumentError QXmlStreamReader__Error = 4
)

