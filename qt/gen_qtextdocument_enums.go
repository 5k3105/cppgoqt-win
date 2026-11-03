package qt

type QTextDocument__MetaInformation int64

const (
	QTextDocument__DocumentTitle QTextDocument__MetaInformation = 0
	QTextDocument__DocumentUrl QTextDocument__MetaInformation = 1
	QTextDocument__CssMedia QTextDocument__MetaInformation = 2
	QTextDocument__FrontMatter QTextDocument__MetaInformation = 3
)

type QTextDocument__MarkdownFeature int64

const (
	QTextDocument__MarkdownNoHTML QTextDocument__MarkdownFeature = 96
	QTextDocument__MarkdownDialectCommonMark QTextDocument__MarkdownFeature = 0
	QTextDocument__MarkdownDialectGitHub QTextDocument__MarkdownFeature = 1068812
)

type QTextDocument__FindFlag int64

const (
	QTextDocument__FindBackward QTextDocument__FindFlag = 1
	QTextDocument__FindCaseSensitively QTextDocument__FindFlag = 2
	QTextDocument__FindWholeWords QTextDocument__FindFlag = 4
)

type QTextDocument__ResourceType int64

const (
	QTextDocument__UnknownResource QTextDocument__ResourceType = 0
	QTextDocument__HtmlResource QTextDocument__ResourceType = 1
	QTextDocument__ImageResource QTextDocument__ResourceType = 2
	QTextDocument__StyleSheetResource QTextDocument__ResourceType = 3
	QTextDocument__MarkdownResource QTextDocument__ResourceType = 4
	QTextDocument__UserResource QTextDocument__ResourceType = 100
)

type QTextDocument__Stacks int64

const (
	QTextDocument__UndoStack QTextDocument__Stacks = 1
	QTextDocument__RedoStack QTextDocument__Stacks = 2
	QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = 3
)

