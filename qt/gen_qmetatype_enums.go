package qt

type QMetaType__Type int64

const (
	QMetaType__Bool QMetaType__Type = 1
	QMetaType__Int QMetaType__Type = 2
	QMetaType__UInt QMetaType__Type = 3
	QMetaType__LongLong QMetaType__Type = 4
	QMetaType__ULongLong QMetaType__Type = 5
	QMetaType__Double QMetaType__Type = 6
	QMetaType__Long QMetaType__Type = 32
	QMetaType__Short QMetaType__Type = 33
	QMetaType__Char QMetaType__Type = 34
	QMetaType__Char16 QMetaType__Type = 56
	QMetaType__Char32 QMetaType__Type = 57
	QMetaType__ULong QMetaType__Type = 35
	QMetaType__UShort QMetaType__Type = 36
	QMetaType__UChar QMetaType__Type = 37
	QMetaType__Float QMetaType__Type = 38
	QMetaType__SChar QMetaType__Type = 40
	QMetaType__Nullptr QMetaType__Type = 51
	QMetaType__QCborSimpleType QMetaType__Type = 52
	QMetaType__Void QMetaType__Type = 43
	QMetaType__VoidStar QMetaType__Type = 31
	QMetaType__QChar QMetaType__Type = 7
	QMetaType__QString QMetaType__Type = 10
	QMetaType__QByteArray QMetaType__Type = 12
	QMetaType__QBitArray QMetaType__Type = 13
	QMetaType__QDate QMetaType__Type = 14
	QMetaType__QTime QMetaType__Type = 15
	QMetaType__QDateTime QMetaType__Type = 16
	QMetaType__QUrl QMetaType__Type = 17
	QMetaType__QLocale QMetaType__Type = 18
	QMetaType__QRect QMetaType__Type = 19
	QMetaType__QRectF QMetaType__Type = 20
	QMetaType__QSize QMetaType__Type = 21
	QMetaType__QSizeF QMetaType__Type = 22
	QMetaType__QLine QMetaType__Type = 23
	QMetaType__QLineF QMetaType__Type = 24
	QMetaType__QPoint QMetaType__Type = 25
	QMetaType__QPointF QMetaType__Type = 26
	QMetaType__QEasingCurve QMetaType__Type = 29
	QMetaType__QUuid QMetaType__Type = 30
	QMetaType__QVariant QMetaType__Type = 41
	QMetaType__QRegularExpression QMetaType__Type = 44
	QMetaType__QJsonValue QMetaType__Type = 45
	QMetaType__QJsonObject QMetaType__Type = 46
	QMetaType__QJsonArray QMetaType__Type = 47
	QMetaType__QJsonDocument QMetaType__Type = 48
	QMetaType__QCborValue QMetaType__Type = 53
	QMetaType__QCborArray QMetaType__Type = 54
	QMetaType__QCborMap QMetaType__Type = 55
	QMetaType__Float16 QMetaType__Type = 63
	QMetaType__QModelIndex QMetaType__Type = 42
	QMetaType__QPersistentModelIndex QMetaType__Type = 50
	QMetaType__QObjectStar QMetaType__Type = 39
	QMetaType__QVariantMap QMetaType__Type = 8
	QMetaType__QVariantList QMetaType__Type = 9
	QMetaType__QVariantHash QMetaType__Type = 28
	QMetaType__QVariantPair QMetaType__Type = 58
	QMetaType__QByteArrayList QMetaType__Type = 49
	QMetaType__QStringList QMetaType__Type = 11
	QMetaType__QFont QMetaType__Type = 4096
	QMetaType__QPixmap QMetaType__Type = 4097
	QMetaType__QBrush QMetaType__Type = 4098
	QMetaType__QColor QMetaType__Type = 4099
	QMetaType__QPalette QMetaType__Type = 4100
	QMetaType__QIcon QMetaType__Type = 4101
	QMetaType__QImage QMetaType__Type = 4102
	QMetaType__QPolygon QMetaType__Type = 4103
	QMetaType__QRegion QMetaType__Type = 4104
	QMetaType__QBitmap QMetaType__Type = 4105
	QMetaType__QCursor QMetaType__Type = 4106
	QMetaType__QKeySequence QMetaType__Type = 4107
	QMetaType__QPen QMetaType__Type = 4108
	QMetaType__QTextLength QMetaType__Type = 4109
	QMetaType__QTextFormat QMetaType__Type = 4110
	QMetaType__QTransform QMetaType__Type = 4112
	QMetaType__QMatrix4x4 QMetaType__Type = 4113
	QMetaType__QVector2D QMetaType__Type = 4114
	QMetaType__QVector3D QMetaType__Type = 4115
	QMetaType__QVector4D QMetaType__Type = 4116
	QMetaType__QQuaternion QMetaType__Type = 4117
	QMetaType__QPolygonF QMetaType__Type = 4118
	QMetaType__QColorSpace QMetaType__Type = 4119
	QMetaType__QSizePolicy QMetaType__Type = 8192
	QMetaType__FirstCoreType QMetaType__Type = 1
	QMetaType__LastCoreType QMetaType__Type = 63
	QMetaType__FirstGuiType QMetaType__Type = 4096
	QMetaType__LastGuiType QMetaType__Type = 4119
	QMetaType__FirstWidgetsType QMetaType__Type = 8192
	QMetaType__LastWidgetsType QMetaType__Type = 8192
	QMetaType__HighestInternalId QMetaType__Type = 8192
	QMetaType__QReal QMetaType__Type = 6
	QMetaType__UnknownType QMetaType__Type = 0
	QMetaType__User QMetaType__Type = 65536
)

type QMetaType__TypeFlag int64

const (
	QMetaType__NeedsConstruction QMetaType__TypeFlag = 1
	QMetaType__NeedsDestruction QMetaType__TypeFlag = 2
	QMetaType__RelocatableType QMetaType__TypeFlag = 4
	QMetaType__MovableType QMetaType__TypeFlag = 4
	QMetaType__PointerToQObject QMetaType__TypeFlag = 8
	QMetaType__IsEnumeration QMetaType__TypeFlag = 16
	QMetaType__SharedPointerToQObject QMetaType__TypeFlag = 32
	QMetaType__WeakPointerToQObject QMetaType__TypeFlag = 64
	QMetaType__TrackingPointerToQObject QMetaType__TypeFlag = 128
	QMetaType__IsUnsignedEnumeration QMetaType__TypeFlag = 256
	QMetaType__IsGadget QMetaType__TypeFlag = 512
	QMetaType__PointerToGadget QMetaType__TypeFlag = 1024
	QMetaType__IsPointer QMetaType__TypeFlag = 2048
	QMetaType__IsQmlList QMetaType__TypeFlag = 4096
	QMetaType__IsConst QMetaType__TypeFlag = 8192
	QMetaType__NeedsCopyConstruction QMetaType__TypeFlag = 16384
	QMetaType__NeedsMoveConstruction QMetaType__TypeFlag = 32768
)

