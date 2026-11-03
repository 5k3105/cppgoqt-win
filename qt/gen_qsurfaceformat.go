package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSurfaceFormat_new uintptr
	addr_QSurfaceFormat_new2 uintptr
	addr_QSurfaceFormat_new3 uintptr
	addr_QSurfaceFormat_OperatorAssign uintptr
	addr_QSurfaceFormat_setDepthBufferSize uintptr
	addr_QSurfaceFormat_depthBufferSize uintptr
	addr_QSurfaceFormat_setStencilBufferSize uintptr
	addr_QSurfaceFormat_stencilBufferSize uintptr
	addr_QSurfaceFormat_setRedBufferSize uintptr
	addr_QSurfaceFormat_redBufferSize uintptr
	addr_QSurfaceFormat_setGreenBufferSize uintptr
	addr_QSurfaceFormat_greenBufferSize uintptr
	addr_QSurfaceFormat_setBlueBufferSize uintptr
	addr_QSurfaceFormat_blueBufferSize uintptr
	addr_QSurfaceFormat_setAlphaBufferSize uintptr
	addr_QSurfaceFormat_alphaBufferSize uintptr
	addr_QSurfaceFormat_setSamples uintptr
	addr_QSurfaceFormat_samples uintptr
	addr_QSurfaceFormat_setSwapBehavior uintptr
	addr_QSurfaceFormat_swapBehavior uintptr
	addr_QSurfaceFormat_hasAlpha uintptr
	addr_QSurfaceFormat_setProfile uintptr
	addr_QSurfaceFormat_profile uintptr
	addr_QSurfaceFormat_setRenderableType uintptr
	addr_QSurfaceFormat_renderableType uintptr
	addr_QSurfaceFormat_setMajorVersion uintptr
	addr_QSurfaceFormat_majorVersion uintptr
	addr_QSurfaceFormat_setMinorVersion uintptr
	addr_QSurfaceFormat_minorVersion uintptr
	addr_QSurfaceFormat_version uintptr
	addr_QSurfaceFormat_setVersion uintptr
	addr_QSurfaceFormat_stereo uintptr
	addr_QSurfaceFormat_setStereo uintptr
	addr_QSurfaceFormat_setOptions uintptr
	addr_QSurfaceFormat_setOption uintptr
	addr_QSurfaceFormat_testOption uintptr
	addr_QSurfaceFormat_options uintptr
	addr_QSurfaceFormat_swapInterval uintptr
	addr_QSurfaceFormat_setSwapInterval uintptr
	addr_QSurfaceFormat_colorSpace uintptr
	addr_QSurfaceFormat_setColorSpace uintptr
	addr_QSurfaceFormat_setColorSpace2 uintptr
	addr_QSurfaceFormat_setDefaultFormat uintptr
	addr_QSurfaceFormat_defaultFormat uintptr
	addr_QSurfaceFormat_setOption2 uintptr
	addr_QSurfaceFormat_OperatorEqual uintptr
	addr_QSurfaceFormat_OperatorNotEqual uintptr
	pg_QSurfaceFormat_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSurfaceFormat_new, "QSurfaceFormat_new")
	symaddr(&addr_QSurfaceFormat_new2, "QSurfaceFormat_new2")
	symaddr(&addr_QSurfaceFormat_new3, "QSurfaceFormat_new3")
	symaddr(&addr_QSurfaceFormat_OperatorAssign, "QSurfaceFormat_OperatorAssign")
	symaddr(&addr_QSurfaceFormat_setDepthBufferSize, "QSurfaceFormat_setDepthBufferSize")
	symaddr(&addr_QSurfaceFormat_depthBufferSize, "QSurfaceFormat_depthBufferSize")
	symaddr(&addr_QSurfaceFormat_setStencilBufferSize, "QSurfaceFormat_setStencilBufferSize")
	symaddr(&addr_QSurfaceFormat_stencilBufferSize, "QSurfaceFormat_stencilBufferSize")
	symaddr(&addr_QSurfaceFormat_setRedBufferSize, "QSurfaceFormat_setRedBufferSize")
	symaddr(&addr_QSurfaceFormat_redBufferSize, "QSurfaceFormat_redBufferSize")
	symaddr(&addr_QSurfaceFormat_setGreenBufferSize, "QSurfaceFormat_setGreenBufferSize")
	symaddr(&addr_QSurfaceFormat_greenBufferSize, "QSurfaceFormat_greenBufferSize")
	symaddr(&addr_QSurfaceFormat_setBlueBufferSize, "QSurfaceFormat_setBlueBufferSize")
	symaddr(&addr_QSurfaceFormat_blueBufferSize, "QSurfaceFormat_blueBufferSize")
	symaddr(&addr_QSurfaceFormat_setAlphaBufferSize, "QSurfaceFormat_setAlphaBufferSize")
	symaddr(&addr_QSurfaceFormat_alphaBufferSize, "QSurfaceFormat_alphaBufferSize")
	symaddr(&addr_QSurfaceFormat_setSamples, "QSurfaceFormat_setSamples")
	symaddr(&addr_QSurfaceFormat_samples, "QSurfaceFormat_samples")
	symaddr(&addr_QSurfaceFormat_setSwapBehavior, "QSurfaceFormat_setSwapBehavior")
	symaddr(&addr_QSurfaceFormat_swapBehavior, "QSurfaceFormat_swapBehavior")
	symaddr(&addr_QSurfaceFormat_hasAlpha, "QSurfaceFormat_hasAlpha")
	symaddr(&addr_QSurfaceFormat_setProfile, "QSurfaceFormat_setProfile")
	symaddr(&addr_QSurfaceFormat_profile, "QSurfaceFormat_profile")
	symaddr(&addr_QSurfaceFormat_setRenderableType, "QSurfaceFormat_setRenderableType")
	symaddr(&addr_QSurfaceFormat_renderableType, "QSurfaceFormat_renderableType")
	symaddr(&addr_QSurfaceFormat_setMajorVersion, "QSurfaceFormat_setMajorVersion")
	symaddr(&addr_QSurfaceFormat_majorVersion, "QSurfaceFormat_majorVersion")
	symaddr(&addr_QSurfaceFormat_setMinorVersion, "QSurfaceFormat_setMinorVersion")
	symaddr(&addr_QSurfaceFormat_minorVersion, "QSurfaceFormat_minorVersion")
	symaddr(&addr_QSurfaceFormat_version, "QSurfaceFormat_version")
	symaddr(&addr_QSurfaceFormat_setVersion, "QSurfaceFormat_setVersion")
	symaddr(&addr_QSurfaceFormat_stereo, "QSurfaceFormat_stereo")
	symaddr(&addr_QSurfaceFormat_setStereo, "QSurfaceFormat_setStereo")
	symaddr(&addr_QSurfaceFormat_setOptions, "QSurfaceFormat_setOptions")
	symaddr(&addr_QSurfaceFormat_setOption, "QSurfaceFormat_setOption")
	symaddr(&addr_QSurfaceFormat_testOption, "QSurfaceFormat_testOption")
	symaddr(&addr_QSurfaceFormat_options, "QSurfaceFormat_options")
	symaddr(&addr_QSurfaceFormat_swapInterval, "QSurfaceFormat_swapInterval")
	symaddr(&addr_QSurfaceFormat_setSwapInterval, "QSurfaceFormat_setSwapInterval")
	symaddr(&addr_QSurfaceFormat_colorSpace, "QSurfaceFormat_colorSpace")
	symaddr(&addr_QSurfaceFormat_setColorSpace, "QSurfaceFormat_setColorSpace")
	symaddr(&addr_QSurfaceFormat_setColorSpace2, "QSurfaceFormat_setColorSpace2")
	symaddr(&addr_QSurfaceFormat_setDefaultFormat, "QSurfaceFormat_setDefaultFormat")
	symaddr(&addr_QSurfaceFormat_defaultFormat, "QSurfaceFormat_defaultFormat")
	symaddr(&addr_QSurfaceFormat_setOption2, "QSurfaceFormat_setOption2")
	symaddr(&addr_QSurfaceFormat_OperatorEqual, "QSurfaceFormat_OperatorEqual")
	symaddr(&addr_QSurfaceFormat_OperatorNotEqual, "QSurfaceFormat_OperatorNotEqual")
	reg(&pg_QSurfaceFormat_delete, "QSurfaceFormat_delete")
}

type QSurfaceFormat struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSurfaceFormat) UnsafeGoID() int64     { return this.vid }
func (this *QSurfaceFormat) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSurfaceFormat(h uintptr) *QSurfaceFormat {
	if h == 0 {
		return nil
	}
	return &QSurfaceFormat{h: h}
}

func newownedQSurfaceFormat(h uintptr) *QSurfaceFormat {
	o := bindQSurfaceFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSurfaceFormat).Delete)
	}
	return o
}

func (this *QSurfaceFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSurfaceFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSurfaceFormat(h unsafe.Pointer) *QSurfaceFormat { return bindQSurfaceFormat(uintptr(h)) }
func UnsafeNewQSurfaceFormatOwned(h unsafe.Pointer) *QSurfaceFormat { return newownedQSurfaceFormat(uintptr(h)) }

func (this *QSurfaceFormat) UnsafeRelease() { this.h = 0 }

func NewQSurfaceFormat() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QSurfaceFormat_new)
	__o := newownedQSurfaceFormat(__r)
	return __o
}

func NewQSurfaceFormat2(options QSurfaceFormat__FormatOption) *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QSurfaceFormat_new2, uintptr(options))
	__o := newownedQSurfaceFormat(__r)
	return __o
}

func NewQSurfaceFormat3(other *QSurfaceFormat) *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QSurfaceFormat_new3, uintptr(other.UnsafePointer()))
	__o := newownedQSurfaceFormat(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QSurfaceFormat) OperatorAssign(other *QSurfaceFormat) *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QSurfaceFormat_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQSurfaceFormat(__r)
}

func (this *QSurfaceFormat) SetDepthBufferSize(size int) {
	purego.SyscallN(addr_QSurfaceFormat_setDepthBufferSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) DepthBufferSize() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_depthBufferSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetStencilBufferSize(size int) {
	purego.SyscallN(addr_QSurfaceFormat_setStencilBufferSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) StencilBufferSize() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_stencilBufferSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetRedBufferSize(size int) {
	purego.SyscallN(addr_QSurfaceFormat_setRedBufferSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) RedBufferSize() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_redBufferSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetGreenBufferSize(size int) {
	purego.SyscallN(addr_QSurfaceFormat_setGreenBufferSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) GreenBufferSize() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_greenBufferSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetBlueBufferSize(size int) {
	purego.SyscallN(addr_QSurfaceFormat_setBlueBufferSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) BlueBufferSize() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_blueBufferSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetAlphaBufferSize(size int) {
	purego.SyscallN(addr_QSurfaceFormat_setAlphaBufferSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) AlphaBufferSize() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_alphaBufferSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetSamples(numSamples int) {
	purego.SyscallN(addr_QSurfaceFormat_setSamples, this.h, uintptr(numSamples))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) Samples() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_samples, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetSwapBehavior(behavior QSurfaceFormat__SwapBehavior) {
	purego.SyscallN(addr_QSurfaceFormat_setSwapBehavior, this.h, uintptr(behavior))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) SwapBehavior() QSurfaceFormat__SwapBehavior {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_swapBehavior, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurfaceFormat__SwapBehavior(__r)
}

func (this *QSurfaceFormat) HasAlpha() bool {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_hasAlpha, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSurfaceFormat) SetProfile(profile QSurfaceFormat__OpenGLContextProfile) {
	purego.SyscallN(addr_QSurfaceFormat_setProfile, this.h, uintptr(profile))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) Profile() QSurfaceFormat__OpenGLContextProfile {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_profile, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurfaceFormat__OpenGLContextProfile(__r)
}

func (this *QSurfaceFormat) SetRenderableType(type_ QSurfaceFormat__RenderableType) {
	purego.SyscallN(addr_QSurfaceFormat_setRenderableType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) RenderableType() QSurfaceFormat__RenderableType {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_renderableType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurfaceFormat__RenderableType(__r)
}

func (this *QSurfaceFormat) SetMajorVersion(majorVersion int) {
	purego.SyscallN(addr_QSurfaceFormat_setMajorVersion, this.h, uintptr(majorVersion))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) MajorVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_majorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetMinorVersion(minorVersion int) {
	purego.SyscallN(addr_QSurfaceFormat_setMinorVersion, this.h, uintptr(minorVersion))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) MinorVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_minorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) Version() (int, int) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_version, this.h, uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), int(__b)
}

func (this *QSurfaceFormat) SetVersion(major int, minor int) {
	purego.SyscallN(addr_QSurfaceFormat_setVersion, this.h, uintptr(major), uintptr(minor))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) Stereo() bool {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_stereo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSurfaceFormat) SetStereo(enable bool) {
	purego.SyscallN(addr_QSurfaceFormat_setStereo, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) SetOptions(options QSurfaceFormat__FormatOption) {
	purego.SyscallN(addr_QSurfaceFormat_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) SetOption(option QSurfaceFormat__FormatOption) {
	purego.SyscallN(addr_QSurfaceFormat_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) TestOption(option QSurfaceFormat__FormatOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSurfaceFormat) Options() QSurfaceFormat__FormatOption {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurfaceFormat__FormatOption(__r)
}

func (this *QSurfaceFormat) SwapInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_swapInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSurfaceFormat) SetSwapInterval(interval int) {
	purego.SyscallN(addr_QSurfaceFormat_setSwapInterval, this.h, uintptr(interval))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) ColorSpace() *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QSurfaceFormat_colorSpace, this.h)
	runtime.KeepAlive(this)
	return bindQColorSpace(__r)
}

func (this *QSurfaceFormat) SetColorSpace(colorSpace *QColorSpace) {
	purego.SyscallN(addr_QSurfaceFormat_setColorSpace, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
}

func (this *QSurfaceFormat) SetColorSpaceWithColorSpace(colorSpace QSurfaceFormat__ColorSpace) {
	purego.SyscallN(addr_QSurfaceFormat_setColorSpace2, this.h, uintptr(colorSpace))
	runtime.KeepAlive(this)
}

func QSurfaceFormat_SetDefaultFormat(format *QSurfaceFormat) {
	purego.SyscallN(addr_QSurfaceFormat_setDefaultFormat, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(format)
}

func QSurfaceFormat_DefaultFormat() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QSurfaceFormat_defaultFormat)
	return newownedQSurfaceFormat(__r)
}

func (this *QSurfaceFormat) SetOption2(option QSurfaceFormat__FormatOption, on bool) {
	purego.SyscallN(addr_QSurfaceFormat_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QSurfaceFormat) OperatorEqual(rhs *QSurfaceFormat) bool {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSurfaceFormat) OperatorNotEqual(rhs *QSurfaceFormat) bool {
	__r0, _, _ := purego.SyscallN(addr_QSurfaceFormat_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSurfaceFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSurfaceFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
