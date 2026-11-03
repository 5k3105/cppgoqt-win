package qt

type QPixelFormat__FieldWidth int64

const (
	QPixelFormat__ModelFieldWidth QPixelFormat__FieldWidth = 4
	QPixelFormat__FirstFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__SecondFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__ThirdFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__FourthFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__FifthFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__AlphaFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__AlphaUsageFieldWidth QPixelFormat__FieldWidth = 1
	QPixelFormat__AlphaPositionFieldWidth QPixelFormat__FieldWidth = 1
	QPixelFormat__PremulFieldWidth QPixelFormat__FieldWidth = 1
	QPixelFormat__TypeInterpretationFieldWidth QPixelFormat__FieldWidth = 4
	QPixelFormat__ByteOrderFieldWidth QPixelFormat__FieldWidth = 2
	QPixelFormat__SubEnumFieldWidth QPixelFormat__FieldWidth = 6
	QPixelFormat__UnusedFieldWidth QPixelFormat__FieldWidth = 9
	QPixelFormat__TotalFieldWidthByWidths QPixelFormat__FieldWidth = 64
)

type QPixelFormat__Field int64

const (
	QPixelFormat__ModelField QPixelFormat__Field = 0
	QPixelFormat__FirstField QPixelFormat__Field = 4
	QPixelFormat__SecondField QPixelFormat__Field = 10
	QPixelFormat__ThirdField QPixelFormat__Field = 16
	QPixelFormat__FourthField QPixelFormat__Field = 22
	QPixelFormat__FifthField QPixelFormat__Field = 28
	QPixelFormat__AlphaField QPixelFormat__Field = 34
	QPixelFormat__AlphaUsageField QPixelFormat__Field = 40
	QPixelFormat__AlphaPositionField QPixelFormat__Field = 41
	QPixelFormat__PremulField QPixelFormat__Field = 42
	QPixelFormat__TypeInterpretationField QPixelFormat__Field = 43
	QPixelFormat__ByteOrderField QPixelFormat__Field = 47
	QPixelFormat__SubEnumField QPixelFormat__Field = 49
	QPixelFormat__UnusedField QPixelFormat__Field = 55
	QPixelFormat__TotalFieldWidthByOffsets QPixelFormat__Field = 64
)

type QPixelFormat__ColorModel int64

const (
	QPixelFormat__RGB QPixelFormat__ColorModel = 0
	QPixelFormat__BGR QPixelFormat__ColorModel = 1
	QPixelFormat__Indexed QPixelFormat__ColorModel = 2
	QPixelFormat__Grayscale QPixelFormat__ColorModel = 3
	QPixelFormat__CMYK QPixelFormat__ColorModel = 4
	QPixelFormat__HSL QPixelFormat__ColorModel = 5
	QPixelFormat__HSV QPixelFormat__ColorModel = 6
	QPixelFormat__YUV QPixelFormat__ColorModel = 7
	QPixelFormat__Alpha QPixelFormat__ColorModel = 8
)

type QPixelFormat__AlphaUsage int64

const (
	QPixelFormat__UsesAlpha QPixelFormat__AlphaUsage = 0
	QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = 1
)

type QPixelFormat__AlphaPosition int64

const (
	QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = 0
	QPixelFormat__AtEnd QPixelFormat__AlphaPosition = 1
)

type QPixelFormat__AlphaPremultiplied int64

const (
	QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = 0
	QPixelFormat__Premultiplied QPixelFormat__AlphaPremultiplied = 1
)

type QPixelFormat__TypeInterpretation int64

const (
	QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = 0
	QPixelFormat__UnsignedShort QPixelFormat__TypeInterpretation = 1
	QPixelFormat__UnsignedByte QPixelFormat__TypeInterpretation = 2
	QPixelFormat__FloatingPoint QPixelFormat__TypeInterpretation = 3
)

type QPixelFormat__YUVLayout int64

const (
	QPixelFormat__YUV444 QPixelFormat__YUVLayout = 0
	QPixelFormat__YUV422 QPixelFormat__YUVLayout = 1
	QPixelFormat__YUV411 QPixelFormat__YUVLayout = 2
	QPixelFormat__YUV420P QPixelFormat__YUVLayout = 3
	QPixelFormat__YUV420SP QPixelFormat__YUVLayout = 4
	QPixelFormat__YV12 QPixelFormat__YUVLayout = 5
	QPixelFormat__UYVY QPixelFormat__YUVLayout = 6
	QPixelFormat__YUYV QPixelFormat__YUVLayout = 7
	QPixelFormat__NV12 QPixelFormat__YUVLayout = 8
	QPixelFormat__NV21 QPixelFormat__YUVLayout = 9
	QPixelFormat__IMC1 QPixelFormat__YUVLayout = 10
	QPixelFormat__IMC2 QPixelFormat__YUVLayout = 11
	QPixelFormat__IMC3 QPixelFormat__YUVLayout = 12
	QPixelFormat__IMC4 QPixelFormat__YUVLayout = 13
	QPixelFormat__Y8 QPixelFormat__YUVLayout = 14
	QPixelFormat__Y16 QPixelFormat__YUVLayout = 15
)

type QPixelFormat__ByteOrder int64

const (
	QPixelFormat__LittleEndian QPixelFormat__ByteOrder = 0
	QPixelFormat__BigEndian QPixelFormat__ByteOrder = 1
	QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = 2
)

