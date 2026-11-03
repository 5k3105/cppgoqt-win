package qt

type QTextFormat__FormatType int64

const (
	QTextFormat__InvalidFormat QTextFormat__FormatType = -1
	QTextFormat__BlockFormat QTextFormat__FormatType = 1
	QTextFormat__CharFormat QTextFormat__FormatType = 2
	QTextFormat__ListFormat QTextFormat__FormatType = 3
	QTextFormat__FrameFormat QTextFormat__FormatType = 5
	QTextFormat__UserFormat QTextFormat__FormatType = 100
)

type QTextFormat__Property int64

const (
	QTextFormat__ObjectIndex QTextFormat__Property = 0
	QTextFormat__CssFloat QTextFormat__Property = 2048
	QTextFormat__LayoutDirection QTextFormat__Property = 2049
	QTextFormat__OutlinePen QTextFormat__Property = 2064
	QTextFormat__BackgroundBrush QTextFormat__Property = 2080
	QTextFormat__ForegroundBrush QTextFormat__Property = 2081
	QTextFormat__BackgroundImageUrl QTextFormat__Property = 2083
	QTextFormat__BlockAlignment QTextFormat__Property = 4112
	QTextFormat__BlockTopMargin QTextFormat__Property = 4144
	QTextFormat__BlockBottomMargin QTextFormat__Property = 4145
	QTextFormat__BlockLeftMargin QTextFormat__Property = 4146
	QTextFormat__BlockRightMargin QTextFormat__Property = 4147
	QTextFormat__TextIndent QTextFormat__Property = 4148
	QTextFormat__TabPositions QTextFormat__Property = 4149
	QTextFormat__BlockIndent QTextFormat__Property = 4160
	QTextFormat__LineHeight QTextFormat__Property = 4168
	QTextFormat__LineHeightType QTextFormat__Property = 4169
	QTextFormat__BlockNonBreakableLines QTextFormat__Property = 4176
	QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = 4192
	QTextFormat__HeadingLevel QTextFormat__Property = 4208
	QTextFormat__BlockQuoteLevel QTextFormat__Property = 4224
	QTextFormat__BlockCodeLanguage QTextFormat__Property = 4240
	QTextFormat__BlockCodeFence QTextFormat__Property = 4241
	QTextFormat__BlockMarker QTextFormat__Property = 4256
	QTextFormat__FirstFontProperty QTextFormat__Property = 8160
	QTextFormat__FontCapitalization QTextFormat__Property = 8160
	QTextFormat__FontLetterSpacing QTextFormat__Property = 8161
	QTextFormat__FontWordSpacing QTextFormat__Property = 8162
	QTextFormat__FontStyleHint QTextFormat__Property = 8163
	QTextFormat__FontStyleStrategy QTextFormat__Property = 8164
	QTextFormat__FontKerning QTextFormat__Property = 8165
	QTextFormat__FontHintingPreference QTextFormat__Property = 8166
	QTextFormat__FontFamilies QTextFormat__Property = 8167
	QTextFormat__FontStyleName QTextFormat__Property = 8168
	QTextFormat__FontLetterSpacingType QTextFormat__Property = 8169
	QTextFormat__FontStretch QTextFormat__Property = 8170
	QTextFormat__FontFamily QTextFormat__Property = 8192
	QTextFormat__FontPointSize QTextFormat__Property = 8193
	QTextFormat__FontSizeAdjustment QTextFormat__Property = 8194
	QTextFormat__FontSizeIncrement QTextFormat__Property = 8194
	QTextFormat__FontWeight QTextFormat__Property = 8195
	QTextFormat__FontItalic QTextFormat__Property = 8196
	QTextFormat__FontUnderline QTextFormat__Property = 8197
	QTextFormat__FontOverline QTextFormat__Property = 8198
	QTextFormat__FontStrikeOut QTextFormat__Property = 8199
	QTextFormat__FontFixedPitch QTextFormat__Property = 8200
	QTextFormat__FontPixelSize QTextFormat__Property = 8201
	QTextFormat__FontFeatures QTextFormat__Property = 8208
	QTextFormat__FontVariableAxes QTextFormat__Property = 8209
	QTextFormat__LastFontProperty QTextFormat__Property = 8209
	QTextFormat__TextUnderlineColor QTextFormat__Property = 8224
	QTextFormat__TextVerticalAlignment QTextFormat__Property = 8225
	QTextFormat__TextOutline QTextFormat__Property = 8226
	QTextFormat__TextUnderlineStyle QTextFormat__Property = 8227
	QTextFormat__TextToolTip QTextFormat__Property = 8228
	QTextFormat__TextSuperScriptBaseline QTextFormat__Property = 8229
	QTextFormat__TextSubScriptBaseline QTextFormat__Property = 8230
	QTextFormat__TextBaselineOffset QTextFormat__Property = 8231
	QTextFormat__IsAnchor QTextFormat__Property = 8240
	QTextFormat__AnchorHref QTextFormat__Property = 8241
	QTextFormat__AnchorName QTextFormat__Property = 8242
	QTextFormat__OldFontLetterSpacingType QTextFormat__Property = 8243
	QTextFormat__OldFontStretch QTextFormat__Property = 8244
	QTextFormat__OldTextUnderlineColor QTextFormat__Property = 8208
	QTextFormat__OldFontFamily QTextFormat__Property = 8192
	QTextFormat__ObjectType QTextFormat__Property = 12032
	QTextFormat__ListStyle QTextFormat__Property = 12288
	QTextFormat__ListIndent QTextFormat__Property = 12289
	QTextFormat__ListNumberPrefix QTextFormat__Property = 12290
	QTextFormat__ListNumberSuffix QTextFormat__Property = 12291
	QTextFormat__ListStart QTextFormat__Property = 12292
	QTextFormat__FrameBorder QTextFormat__Property = 16384
	QTextFormat__FrameMargin QTextFormat__Property = 16385
	QTextFormat__FramePadding QTextFormat__Property = 16386
	QTextFormat__FrameWidth QTextFormat__Property = 16387
	QTextFormat__FrameHeight QTextFormat__Property = 16388
	QTextFormat__FrameTopMargin QTextFormat__Property = 16389
	QTextFormat__FrameBottomMargin QTextFormat__Property = 16390
	QTextFormat__FrameLeftMargin QTextFormat__Property = 16391
	QTextFormat__FrameRightMargin QTextFormat__Property = 16392
	QTextFormat__FrameBorderBrush QTextFormat__Property = 16393
	QTextFormat__FrameBorderStyle QTextFormat__Property = 16400
	QTextFormat__TableColumns QTextFormat__Property = 16640
	QTextFormat__TableColumnWidthConstraints QTextFormat__Property = 16641
	QTextFormat__TableCellSpacing QTextFormat__Property = 16642
	QTextFormat__TableCellPadding QTextFormat__Property = 16643
	QTextFormat__TableHeaderRowCount QTextFormat__Property = 16644
	QTextFormat__TableBorderCollapse QTextFormat__Property = 16645
	QTextFormat__TableCellRowSpan QTextFormat__Property = 18448
	QTextFormat__TableCellColumnSpan QTextFormat__Property = 18449
	QTextFormat__TableCellTopPadding QTextFormat__Property = 18450
	QTextFormat__TableCellBottomPadding QTextFormat__Property = 18451
	QTextFormat__TableCellLeftPadding QTextFormat__Property = 18452
	QTextFormat__TableCellRightPadding QTextFormat__Property = 18453
	QTextFormat__TableCellTopBorder QTextFormat__Property = 18454
	QTextFormat__TableCellBottomBorder QTextFormat__Property = 18455
	QTextFormat__TableCellLeftBorder QTextFormat__Property = 18456
	QTextFormat__TableCellRightBorder QTextFormat__Property = 18457
	QTextFormat__TableCellTopBorderStyle QTextFormat__Property = 18458
	QTextFormat__TableCellBottomBorderStyle QTextFormat__Property = 18459
	QTextFormat__TableCellLeftBorderStyle QTextFormat__Property = 18460
	QTextFormat__TableCellRightBorderStyle QTextFormat__Property = 18461
	QTextFormat__TableCellTopBorderBrush QTextFormat__Property = 18462
	QTextFormat__TableCellBottomBorderBrush QTextFormat__Property = 18463
	QTextFormat__TableCellLeftBorderBrush QTextFormat__Property = 18464
	QTextFormat__TableCellRightBorderBrush QTextFormat__Property = 18465
	QTextFormat__ImageName QTextFormat__Property = 20480
	QTextFormat__ImageTitle QTextFormat__Property = 20481
	QTextFormat__ImageAltText QTextFormat__Property = 20482
	QTextFormat__ImageWidth QTextFormat__Property = 20496
	QTextFormat__ImageHeight QTextFormat__Property = 20497
	QTextFormat__ImageQuality QTextFormat__Property = 20500
	QTextFormat__ImageMaxWidth QTextFormat__Property = 20501
	QTextFormat__FullWidthSelection QTextFormat__Property = 24576
	QTextFormat__PageBreakPolicy QTextFormat__Property = 28672
	QTextFormat__UserProperty QTextFormat__Property = 1048576
)

type QTextFormat__ObjectTypes int64

const (
	QTextFormat__NoObject QTextFormat__ObjectTypes = 0
	QTextFormat__ImageObject QTextFormat__ObjectTypes = 1
	QTextFormat__TableObject QTextFormat__ObjectTypes = 2
	QTextFormat__TableCellObject QTextFormat__ObjectTypes = 3
	QTextFormat__UserObject QTextFormat__ObjectTypes = 4096
)

type QTextFormat__PageBreakFlag int64

const (
	QTextFormat__PageBreak_Auto QTextFormat__PageBreakFlag = 0
	QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = 1
	QTextFormat__PageBreak_AlwaysAfter QTextFormat__PageBreakFlag = 16
)

