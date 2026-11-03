package qt

type QUrl__ParsingMode int64

const (
	QUrl__TolerantMode QUrl__ParsingMode = 0
	QUrl__StrictMode QUrl__ParsingMode = 1
	QUrl__DecodedMode QUrl__ParsingMode = 2
)

type QUrl__UrlFormattingOption int64

const (
	QUrl__None QUrl__UrlFormattingOption = 0
	QUrl__RemoveScheme QUrl__UrlFormattingOption = 1
	QUrl__RemovePassword QUrl__UrlFormattingOption = 2
	QUrl__RemoveUserInfo QUrl__UrlFormattingOption = 6
	QUrl__RemovePort QUrl__UrlFormattingOption = 8
	QUrl__RemoveAuthority QUrl__UrlFormattingOption = 30
	QUrl__RemovePath QUrl__UrlFormattingOption = 32
	QUrl__RemoveQuery QUrl__UrlFormattingOption = 64
	QUrl__RemoveFragment QUrl__UrlFormattingOption = 128
	QUrl__PreferLocalFile QUrl__UrlFormattingOption = 512
	QUrl__StripTrailingSlash QUrl__UrlFormattingOption = 1024
	QUrl__RemoveFilename QUrl__UrlFormattingOption = 2048
	QUrl__NormalizePathSegments QUrl__UrlFormattingOption = 4096
)

type QUrl__ComponentFormattingOption int64

const (
	QUrl__PrettyDecoded QUrl__ComponentFormattingOption = 0
	QUrl__EncodeSpaces QUrl__ComponentFormattingOption = 1048576
	QUrl__EncodeUnicode QUrl__ComponentFormattingOption = 2097152
	QUrl__EncodeDelimiters QUrl__ComponentFormattingOption = 12582912
	QUrl__EncodeReserved QUrl__ComponentFormattingOption = 16777216
	QUrl__DecodeReserved QUrl__ComponentFormattingOption = 33554432
	QUrl__FullyEncoded QUrl__ComponentFormattingOption = 32505856
	QUrl__FullyDecoded QUrl__ComponentFormattingOption = 133169152
)

type QUrl__UserInputResolutionOption int64

const (
	QUrl__DefaultResolution QUrl__UserInputResolutionOption = 0
	QUrl__AssumeLocalFile QUrl__UserInputResolutionOption = 1
)

type QUrl__AceProcessingOption int64

const (
	QUrl__IgnoreIDNWhitelist QUrl__AceProcessingOption = 1
	QUrl__AceTransitionalProcessing QUrl__AceProcessingOption = 2
)

