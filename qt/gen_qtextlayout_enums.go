package qt

type QTextLayout__GlyphRunRetrievalFlag int64

const (
	QTextLayout__RetrieveGlyphIndexes QTextLayout__GlyphRunRetrievalFlag = 1
	QTextLayout__RetrieveGlyphPositions QTextLayout__GlyphRunRetrievalFlag = 2
	QTextLayout__RetrieveStringIndexes QTextLayout__GlyphRunRetrievalFlag = 4
	QTextLayout__RetrieveString QTextLayout__GlyphRunRetrievalFlag = 8
	QTextLayout__DefaultRetrievalFlags QTextLayout__GlyphRunRetrievalFlag = 3
	QTextLayout__RetrieveAll QTextLayout__GlyphRunRetrievalFlag = 65535
)

type QTextLayout__CursorMode int64

const (
	QTextLayout__SkipCharacters QTextLayout__CursorMode = 0
	QTextLayout__SkipWords QTextLayout__CursorMode = 1
)

