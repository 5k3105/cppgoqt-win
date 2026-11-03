package qt

type QPaintEngine__PaintEngineFeature int64

const (
	QPaintEngine__PrimitiveTransform QPaintEngine__PaintEngineFeature = 1
	QPaintEngine__PatternTransform QPaintEngine__PaintEngineFeature = 2
	QPaintEngine__PixmapTransform QPaintEngine__PaintEngineFeature = 4
	QPaintEngine__PatternBrush QPaintEngine__PaintEngineFeature = 8
	QPaintEngine__LinearGradientFill QPaintEngine__PaintEngineFeature = 16
	QPaintEngine__RadialGradientFill QPaintEngine__PaintEngineFeature = 32
	QPaintEngine__ConicalGradientFill QPaintEngine__PaintEngineFeature = 64
	QPaintEngine__AlphaBlend QPaintEngine__PaintEngineFeature = 128
	QPaintEngine__PorterDuff QPaintEngine__PaintEngineFeature = 256
	QPaintEngine__PainterPaths QPaintEngine__PaintEngineFeature = 512
	QPaintEngine__Antialiasing QPaintEngine__PaintEngineFeature = 1024
	QPaintEngine__BrushStroke QPaintEngine__PaintEngineFeature = 2048
	QPaintEngine__ConstantOpacity QPaintEngine__PaintEngineFeature = 4096
	QPaintEngine__MaskedBrush QPaintEngine__PaintEngineFeature = 8192
	QPaintEngine__PerspectiveTransform QPaintEngine__PaintEngineFeature = 16384
	QPaintEngine__BlendModes QPaintEngine__PaintEngineFeature = 32768
	QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = 65536
	QPaintEngine__RasterOpModes QPaintEngine__PaintEngineFeature = 131072
	QPaintEngine__PaintOutsidePaintEvent QPaintEngine__PaintEngineFeature = 536870912
	QPaintEngine__AllFeatures QPaintEngine__PaintEngineFeature = 4294967295
)

type QPaintEngine__DirtyFlag int64

const (
	QPaintEngine__DirtyPen QPaintEngine__DirtyFlag = 1
	QPaintEngine__DirtyBrush QPaintEngine__DirtyFlag = 2
	QPaintEngine__DirtyBrushOrigin QPaintEngine__DirtyFlag = 4
	QPaintEngine__DirtyFont QPaintEngine__DirtyFlag = 8
	QPaintEngine__DirtyBackground QPaintEngine__DirtyFlag = 16
	QPaintEngine__DirtyBackgroundMode QPaintEngine__DirtyFlag = 32
	QPaintEngine__DirtyTransform QPaintEngine__DirtyFlag = 64
	QPaintEngine__DirtyClipRegion QPaintEngine__DirtyFlag = 128
	QPaintEngine__DirtyClipPath QPaintEngine__DirtyFlag = 256
	QPaintEngine__DirtyHints QPaintEngine__DirtyFlag = 512
	QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = 1024
	QPaintEngine__DirtyClipEnabled QPaintEngine__DirtyFlag = 2048
	QPaintEngine__DirtyOpacity QPaintEngine__DirtyFlag = 4096
	QPaintEngine__AllDirty QPaintEngine__DirtyFlag = 65535
)

type QPaintEngine__PolygonDrawMode int64

const (
	QPaintEngine__OddEvenMode QPaintEngine__PolygonDrawMode = 0
	QPaintEngine__WindingMode QPaintEngine__PolygonDrawMode = 1
	QPaintEngine__ConvexMode QPaintEngine__PolygonDrawMode = 2
	QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = 3
)

type QPaintEngine__Type int64

const (
	QPaintEngine__X11 QPaintEngine__Type = 0
	QPaintEngine__Windows QPaintEngine__Type = 1
	QPaintEngine__QuickDraw QPaintEngine__Type = 2
	QPaintEngine__CoreGraphics QPaintEngine__Type = 3
	QPaintEngine__MacPrinter QPaintEngine__Type = 4
	QPaintEngine__QWindowSystem QPaintEngine__Type = 5
	QPaintEngine__OpenGL QPaintEngine__Type = 6
	QPaintEngine__Picture QPaintEngine__Type = 7
	QPaintEngine__SVG QPaintEngine__Type = 8
	QPaintEngine__Raster QPaintEngine__Type = 9
	QPaintEngine__Direct3D QPaintEngine__Type = 10
	QPaintEngine__Pdf QPaintEngine__Type = 11
	QPaintEngine__OpenVG QPaintEngine__Type = 12
	QPaintEngine__OpenGL2 QPaintEngine__Type = 13
	QPaintEngine__PaintBuffer QPaintEngine__Type = 14
	QPaintEngine__Blitter QPaintEngine__Type = 15
	QPaintEngine__Direct2D QPaintEngine__Type = 16
	QPaintEngine__User QPaintEngine__Type = 50
	QPaintEngine__MaxUser QPaintEngine__Type = 100
)

