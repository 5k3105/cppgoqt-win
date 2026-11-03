package qt

type QSurface__SurfaceClass int64

const (
	QSurface__Window QSurface__SurfaceClass = 0
	QSurface__Offscreen QSurface__SurfaceClass = 1
)

type QSurface__SurfaceType int64

const (
	QSurface__RasterSurface QSurface__SurfaceType = 0
	QSurface__OpenGLSurface QSurface__SurfaceType = 1
	QSurface__RasterGLSurface QSurface__SurfaceType = 2
	QSurface__OpenVGSurface QSurface__SurfaceType = 3
	QSurface__VulkanSurface QSurface__SurfaceType = 4
	QSurface__MetalSurface QSurface__SurfaceType = 5
	QSurface__Direct3DSurface QSurface__SurfaceType = 6
)

