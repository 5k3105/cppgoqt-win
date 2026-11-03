package qt

type QTextOption__TabType int64

const (
	QTextOption__LeftTab QTextOption__TabType = 0
	QTextOption__RightTab QTextOption__TabType = 1
	QTextOption__CenterTab QTextOption__TabType = 2
	QTextOption__DelimiterTab QTextOption__TabType = 3
)

type QTextOption__WrapMode int64

const (
	QTextOption__NoWrap QTextOption__WrapMode = 0
	QTextOption__WordWrap QTextOption__WrapMode = 1
	QTextOption__ManualWrap QTextOption__WrapMode = 2
	QTextOption__WrapAnywhere QTextOption__WrapMode = 3
	QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = 4
)

type QTextOption__Flag int64

const (
	QTextOption__ShowTabsAndSpaces QTextOption__Flag = 1
	QTextOption__ShowLineAndParagraphSeparators QTextOption__Flag = 2
	QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = 4
	QTextOption__SuppressColors QTextOption__Flag = 8
	QTextOption__ShowDocumentTerminator QTextOption__Flag = 16
	QTextOption__ShowDefaultIgnorables QTextOption__Flag = 32
	QTextOption__DisableEmojiParsing QTextOption__Flag = 64
	QTextOption__IncludeTrailingSpaces QTextOption__Flag = 2147483648
)

