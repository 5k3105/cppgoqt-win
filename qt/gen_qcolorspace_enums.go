package qt

type QColorSpace__NamedColorSpace int64

const (
	QColorSpace__SRgb QColorSpace__NamedColorSpace = 1
	QColorSpace__SRgbLinear QColorSpace__NamedColorSpace = 2
	QColorSpace__AdobeRgb QColorSpace__NamedColorSpace = 3
	QColorSpace__DisplayP3 QColorSpace__NamedColorSpace = 4
	QColorSpace__ProPhotoRgb QColorSpace__NamedColorSpace = 5
	QColorSpace__Bt2020 QColorSpace__NamedColorSpace = 6
	QColorSpace__Bt2100Pq QColorSpace__NamedColorSpace = 7
	QColorSpace__Bt2100Hlg QColorSpace__NamedColorSpace = 8
)

type QColorSpace__Primaries int64

const (
	QColorSpace__Primaries__Custom QColorSpace__Primaries = 0
	QColorSpace__Primaries__SRgb QColorSpace__Primaries = 1
	QColorSpace__Primaries__AdobeRgb QColorSpace__Primaries = 2
	QColorSpace__Primaries__DciP3D65 QColorSpace__Primaries = 3
	QColorSpace__Primaries__ProPhotoRgb QColorSpace__Primaries = 4
	QColorSpace__Primaries__Bt2020 QColorSpace__Primaries = 5
)

type QColorSpace__TransferFunction int64

const (
	QColorSpace__TransferFunction__Custom QColorSpace__TransferFunction = 0
	QColorSpace__TransferFunction__Linear QColorSpace__TransferFunction = 1
	QColorSpace__TransferFunction__Gamma QColorSpace__TransferFunction = 2
	QColorSpace__TransferFunction__SRgb QColorSpace__TransferFunction = 3
	QColorSpace__TransferFunction__ProPhotoRgb QColorSpace__TransferFunction = 4
	QColorSpace__TransferFunction__Bt2020 QColorSpace__TransferFunction = 5
	QColorSpace__TransferFunction__St2084 QColorSpace__TransferFunction = 6
	QColorSpace__TransferFunction__Hlg QColorSpace__TransferFunction = 7
)

type QColorSpace__TransformModel int64

const (
	QColorSpace__TransformModel__ThreeComponentMatrix QColorSpace__TransformModel = 0
	QColorSpace__TransformModel__ElementListProcessing QColorSpace__TransformModel = 1
)

type QColorSpace__ColorModel int64

const (
	QColorSpace__ColorModel__Undefined QColorSpace__ColorModel = 0
	QColorSpace__ColorModel__Rgb QColorSpace__ColorModel = 1
	QColorSpace__ColorModel__Gray QColorSpace__ColorModel = 2
	QColorSpace__ColorModel__Cmyk QColorSpace__ColorModel = 3
)

