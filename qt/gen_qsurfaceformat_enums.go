package qt

type QSurfaceFormat__FormatOption int64

const (
	QSurfaceFormat__StereoBuffers QSurfaceFormat__FormatOption = 1
	QSurfaceFormat__DebugContext QSurfaceFormat__FormatOption = 2
	QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = 4
	QSurfaceFormat__ResetNotification QSurfaceFormat__FormatOption = 8
	QSurfaceFormat__ProtectedContent QSurfaceFormat__FormatOption = 16
)

type QSurfaceFormat__SwapBehavior int64

const (
	QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = 0
	QSurfaceFormat__SingleBuffer QSurfaceFormat__SwapBehavior = 1
	QSurfaceFormat__DoubleBuffer QSurfaceFormat__SwapBehavior = 2
	QSurfaceFormat__TripleBuffer QSurfaceFormat__SwapBehavior = 3
)

type QSurfaceFormat__RenderableType int64

const (
	QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = 0
	QSurfaceFormat__OpenGL QSurfaceFormat__RenderableType = 1
	QSurfaceFormat__OpenGLES QSurfaceFormat__RenderableType = 2
	QSurfaceFormat__OpenVG QSurfaceFormat__RenderableType = 4
)

type QSurfaceFormat__OpenGLContextProfile int64

const (
	QSurfaceFormat__NoProfile QSurfaceFormat__OpenGLContextProfile = 0
	QSurfaceFormat__CoreProfile QSurfaceFormat__OpenGLContextProfile = 1
	QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = 2
)

type QSurfaceFormat__ColorSpace int64

const (
	QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = 0
	QSurfaceFormat__SRGBColorSpace QSurfaceFormat__ColorSpace = 1
)

