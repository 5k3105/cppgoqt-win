package qt

type QIcon__Mode int64

const (
	QIcon__Normal QIcon__Mode = 0
	QIcon__Disabled QIcon__Mode = 1
	QIcon__Active QIcon__Mode = 2
	QIcon__Selected QIcon__Mode = 3
)

type QIcon__State int64

const (
	QIcon__On QIcon__State = 0
	QIcon__Off QIcon__State = 1
)

type QIcon__ThemeIcon int64

const (
	QIcon__ThemeIcon__AddressBookNew QIcon__ThemeIcon = 0
	QIcon__ThemeIcon__ApplicationExit QIcon__ThemeIcon = 1
	QIcon__ThemeIcon__AppointmentNew QIcon__ThemeIcon = 2
	QIcon__ThemeIcon__CallStart QIcon__ThemeIcon = 3
	QIcon__ThemeIcon__CallStop QIcon__ThemeIcon = 4
	QIcon__ThemeIcon__ContactNew QIcon__ThemeIcon = 5
	QIcon__ThemeIcon__DocumentNew QIcon__ThemeIcon = 6
	QIcon__ThemeIcon__DocumentOpen QIcon__ThemeIcon = 7
	QIcon__ThemeIcon__DocumentOpenRecent QIcon__ThemeIcon = 8
	QIcon__ThemeIcon__DocumentPageSetup QIcon__ThemeIcon = 9
	QIcon__ThemeIcon__DocumentPrint QIcon__ThemeIcon = 10
	QIcon__ThemeIcon__DocumentPrintPreview QIcon__ThemeIcon = 11
	QIcon__ThemeIcon__DocumentProperties QIcon__ThemeIcon = 12
	QIcon__ThemeIcon__DocumentRevert QIcon__ThemeIcon = 13
	QIcon__ThemeIcon__DocumentSave QIcon__ThemeIcon = 14
	QIcon__ThemeIcon__DocumentSaveAs QIcon__ThemeIcon = 15
	QIcon__ThemeIcon__DocumentSend QIcon__ThemeIcon = 16
	QIcon__ThemeIcon__EditClear QIcon__ThemeIcon = 17
	QIcon__ThemeIcon__EditCopy QIcon__ThemeIcon = 18
	QIcon__ThemeIcon__EditCut QIcon__ThemeIcon = 19
	QIcon__ThemeIcon__EditDelete QIcon__ThemeIcon = 20
	QIcon__ThemeIcon__EditFind QIcon__ThemeIcon = 21
	QIcon__ThemeIcon__EditPaste QIcon__ThemeIcon = 22
	QIcon__ThemeIcon__EditRedo QIcon__ThemeIcon = 23
	QIcon__ThemeIcon__EditSelectAll QIcon__ThemeIcon = 24
	QIcon__ThemeIcon__EditUndo QIcon__ThemeIcon = 25
	QIcon__ThemeIcon__FolderNew QIcon__ThemeIcon = 26
	QIcon__ThemeIcon__FormatIndentLess QIcon__ThemeIcon = 27
	QIcon__ThemeIcon__FormatIndentMore QIcon__ThemeIcon = 28
	QIcon__ThemeIcon__FormatJustifyCenter QIcon__ThemeIcon = 29
	QIcon__ThemeIcon__FormatJustifyFill QIcon__ThemeIcon = 30
	QIcon__ThemeIcon__FormatJustifyLeft QIcon__ThemeIcon = 31
	QIcon__ThemeIcon__FormatJustifyRight QIcon__ThemeIcon = 32
	QIcon__ThemeIcon__FormatTextDirectionLtr QIcon__ThemeIcon = 33
	QIcon__ThemeIcon__FormatTextDirectionRtl QIcon__ThemeIcon = 34
	QIcon__ThemeIcon__FormatTextBold QIcon__ThemeIcon = 35
	QIcon__ThemeIcon__FormatTextItalic QIcon__ThemeIcon = 36
	QIcon__ThemeIcon__FormatTextUnderline QIcon__ThemeIcon = 37
	QIcon__ThemeIcon__FormatTextStrikethrough QIcon__ThemeIcon = 38
	QIcon__ThemeIcon__GoDown QIcon__ThemeIcon = 39
	QIcon__ThemeIcon__GoHome QIcon__ThemeIcon = 40
	QIcon__ThemeIcon__GoNext QIcon__ThemeIcon = 41
	QIcon__ThemeIcon__GoPrevious QIcon__ThemeIcon = 42
	QIcon__ThemeIcon__GoUp QIcon__ThemeIcon = 43
	QIcon__ThemeIcon__HelpAbout QIcon__ThemeIcon = 44
	QIcon__ThemeIcon__HelpFaq QIcon__ThemeIcon = 45
	QIcon__ThemeIcon__InsertImage QIcon__ThemeIcon = 46
	QIcon__ThemeIcon__InsertLink QIcon__ThemeIcon = 47
	QIcon__ThemeIcon__InsertText QIcon__ThemeIcon = 48
	QIcon__ThemeIcon__ListAdd QIcon__ThemeIcon = 49
	QIcon__ThemeIcon__ListRemove QIcon__ThemeIcon = 50
	QIcon__ThemeIcon__MailForward QIcon__ThemeIcon = 51
	QIcon__ThemeIcon__MailMarkImportant QIcon__ThemeIcon = 52
	QIcon__ThemeIcon__MailMarkRead QIcon__ThemeIcon = 53
	QIcon__ThemeIcon__MailMarkUnread QIcon__ThemeIcon = 54
	QIcon__ThemeIcon__MailMessageNew QIcon__ThemeIcon = 55
	QIcon__ThemeIcon__MailReplyAll QIcon__ThemeIcon = 56
	QIcon__ThemeIcon__MailReplySender QIcon__ThemeIcon = 57
	QIcon__ThemeIcon__MailSend QIcon__ThemeIcon = 58
	QIcon__ThemeIcon__MediaEject QIcon__ThemeIcon = 59
	QIcon__ThemeIcon__MediaPlaybackPause QIcon__ThemeIcon = 60
	QIcon__ThemeIcon__MediaPlaybackStart QIcon__ThemeIcon = 61
	QIcon__ThemeIcon__MediaPlaybackStop QIcon__ThemeIcon = 62
	QIcon__ThemeIcon__MediaRecord QIcon__ThemeIcon = 63
	QIcon__ThemeIcon__MediaSeekBackward QIcon__ThemeIcon = 64
	QIcon__ThemeIcon__MediaSeekForward QIcon__ThemeIcon = 65
	QIcon__ThemeIcon__MediaSkipBackward QIcon__ThemeIcon = 66
	QIcon__ThemeIcon__MediaSkipForward QIcon__ThemeIcon = 67
	QIcon__ThemeIcon__ObjectRotateLeft QIcon__ThemeIcon = 68
	QIcon__ThemeIcon__ObjectRotateRight QIcon__ThemeIcon = 69
	QIcon__ThemeIcon__ProcessStop QIcon__ThemeIcon = 70
	QIcon__ThemeIcon__SystemLockScreen QIcon__ThemeIcon = 71
	QIcon__ThemeIcon__SystemLogOut QIcon__ThemeIcon = 72
	QIcon__ThemeIcon__SystemSearch QIcon__ThemeIcon = 73
	QIcon__ThemeIcon__SystemReboot QIcon__ThemeIcon = 74
	QIcon__ThemeIcon__SystemShutdown QIcon__ThemeIcon = 75
	QIcon__ThemeIcon__ToolsCheckSpelling QIcon__ThemeIcon = 76
	QIcon__ThemeIcon__ViewFullscreen QIcon__ThemeIcon = 77
	QIcon__ThemeIcon__ViewRefresh QIcon__ThemeIcon = 78
	QIcon__ThemeIcon__ViewRestore QIcon__ThemeIcon = 79
	QIcon__ThemeIcon__WindowClose QIcon__ThemeIcon = 80
	QIcon__ThemeIcon__WindowNew QIcon__ThemeIcon = 81
	QIcon__ThemeIcon__ZoomFitBest QIcon__ThemeIcon = 82
	QIcon__ThemeIcon__ZoomIn QIcon__ThemeIcon = 83
	QIcon__ThemeIcon__ZoomOut QIcon__ThemeIcon = 84
	QIcon__ThemeIcon__AudioCard QIcon__ThemeIcon = 85
	QIcon__ThemeIcon__AudioInputMicrophone QIcon__ThemeIcon = 86
	QIcon__ThemeIcon__Battery QIcon__ThemeIcon = 87
	QIcon__ThemeIcon__CameraPhoto QIcon__ThemeIcon = 88
	QIcon__ThemeIcon__CameraVideo QIcon__ThemeIcon = 89
	QIcon__ThemeIcon__CameraWeb QIcon__ThemeIcon = 90
	QIcon__ThemeIcon__Computer QIcon__ThemeIcon = 91
	QIcon__ThemeIcon__DriveHarddisk QIcon__ThemeIcon = 92
	QIcon__ThemeIcon__DriveOptical QIcon__ThemeIcon = 93
	QIcon__ThemeIcon__InputGaming QIcon__ThemeIcon = 94
	QIcon__ThemeIcon__InputKeyboard QIcon__ThemeIcon = 95
	QIcon__ThemeIcon__InputMouse QIcon__ThemeIcon = 96
	QIcon__ThemeIcon__InputTablet QIcon__ThemeIcon = 97
	QIcon__ThemeIcon__MediaFlash QIcon__ThemeIcon = 98
	QIcon__ThemeIcon__MediaOptical QIcon__ThemeIcon = 99
	QIcon__ThemeIcon__MediaTape QIcon__ThemeIcon = 100
	QIcon__ThemeIcon__MultimediaPlayer QIcon__ThemeIcon = 101
	QIcon__ThemeIcon__NetworkWired QIcon__ThemeIcon = 102
	QIcon__ThemeIcon__NetworkWireless QIcon__ThemeIcon = 103
	QIcon__ThemeIcon__Phone QIcon__ThemeIcon = 104
	QIcon__ThemeIcon__Printer QIcon__ThemeIcon = 105
	QIcon__ThemeIcon__Scanner QIcon__ThemeIcon = 106
	QIcon__ThemeIcon__VideoDisplay QIcon__ThemeIcon = 107
	QIcon__ThemeIcon__AppointmentMissed QIcon__ThemeIcon = 108
	QIcon__ThemeIcon__AppointmentSoon QIcon__ThemeIcon = 109
	QIcon__ThemeIcon__AudioVolumeHigh QIcon__ThemeIcon = 110
	QIcon__ThemeIcon__AudioVolumeLow QIcon__ThemeIcon = 111
	QIcon__ThemeIcon__AudioVolumeMedium QIcon__ThemeIcon = 112
	QIcon__ThemeIcon__AudioVolumeMuted QIcon__ThemeIcon = 113
	QIcon__ThemeIcon__BatteryCaution QIcon__ThemeIcon = 114
	QIcon__ThemeIcon__BatteryLow QIcon__ThemeIcon = 115
	QIcon__ThemeIcon__DialogError QIcon__ThemeIcon = 116
	QIcon__ThemeIcon__DialogInformation QIcon__ThemeIcon = 117
	QIcon__ThemeIcon__DialogPassword QIcon__ThemeIcon = 118
	QIcon__ThemeIcon__DialogQuestion QIcon__ThemeIcon = 119
	QIcon__ThemeIcon__DialogWarning QIcon__ThemeIcon = 120
	QIcon__ThemeIcon__FolderDragAccept QIcon__ThemeIcon = 121
	QIcon__ThemeIcon__FolderOpen QIcon__ThemeIcon = 122
	QIcon__ThemeIcon__FolderVisiting QIcon__ThemeIcon = 123
	QIcon__ThemeIcon__ImageLoading QIcon__ThemeIcon = 124
	QIcon__ThemeIcon__ImageMissing QIcon__ThemeIcon = 125
	QIcon__ThemeIcon__MailAttachment QIcon__ThemeIcon = 126
	QIcon__ThemeIcon__MailUnread QIcon__ThemeIcon = 127
	QIcon__ThemeIcon__MailRead QIcon__ThemeIcon = 128
	QIcon__ThemeIcon__MailReplied QIcon__ThemeIcon = 129
	QIcon__ThemeIcon__MediaPlaylistRepeat QIcon__ThemeIcon = 130
	QIcon__ThemeIcon__MediaPlaylistShuffle QIcon__ThemeIcon = 131
	QIcon__ThemeIcon__NetworkOffline QIcon__ThemeIcon = 132
	QIcon__ThemeIcon__PrinterPrinting QIcon__ThemeIcon = 133
	QIcon__ThemeIcon__SecurityHigh QIcon__ThemeIcon = 134
	QIcon__ThemeIcon__SecurityLow QIcon__ThemeIcon = 135
	QIcon__ThemeIcon__SoftwareUpdateAvailable QIcon__ThemeIcon = 136
	QIcon__ThemeIcon__SoftwareUpdateUrgent QIcon__ThemeIcon = 137
	QIcon__ThemeIcon__SyncError QIcon__ThemeIcon = 138
	QIcon__ThemeIcon__SyncSynchronizing QIcon__ThemeIcon = 139
	QIcon__ThemeIcon__UserAvailable QIcon__ThemeIcon = 140
	QIcon__ThemeIcon__UserOffline QIcon__ThemeIcon = 141
	QIcon__ThemeIcon__WeatherClear QIcon__ThemeIcon = 142
	QIcon__ThemeIcon__WeatherClearNight QIcon__ThemeIcon = 143
	QIcon__ThemeIcon__WeatherFewClouds QIcon__ThemeIcon = 144
	QIcon__ThemeIcon__WeatherFewCloudsNight QIcon__ThemeIcon = 145
	QIcon__ThemeIcon__WeatherFog QIcon__ThemeIcon = 146
	QIcon__ThemeIcon__WeatherShowers QIcon__ThemeIcon = 147
	QIcon__ThemeIcon__WeatherSnow QIcon__ThemeIcon = 148
	QIcon__ThemeIcon__WeatherStorm QIcon__ThemeIcon = 149
	QIcon__ThemeIcon__NThemeIcons QIcon__ThemeIcon = 150
)

