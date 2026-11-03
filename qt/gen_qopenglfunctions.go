package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QOpenGLFunctions_new uintptr
	addr_QOpenGLFunctions_new2 uintptr
	addr_QOpenGLFunctions_new3 uintptr
	addr_QOpenGLFunctions_openGLFeatures uintptr
	addr_QOpenGLFunctions_hasOpenGLFeature uintptr
	addr_QOpenGLFunctions_initializeOpenGLFunctions uintptr
	addr_QOpenGLFunctions_glBindTexture uintptr
	addr_QOpenGLFunctions_glBlendFunc uintptr
	addr_QOpenGLFunctions_glClear uintptr
	pg_QOpenGLFunctions_glClearColor func(uintptr, float32, float32, float32, float32)
	addr_QOpenGLFunctions_glClearStencil uintptr
	addr_QOpenGLFunctions_glColorMask uintptr
	addr_QOpenGLFunctions_glCopyTexImage2D uintptr
	addr_QOpenGLFunctions_glCopyTexSubImage2D uintptr
	addr_QOpenGLFunctions_glCullFace uintptr
	addr_QOpenGLFunctions_glDepthFunc uintptr
	addr_QOpenGLFunctions_glDepthMask uintptr
	addr_QOpenGLFunctions_glDisable uintptr
	addr_QOpenGLFunctions_glDrawArrays uintptr
	addr_QOpenGLFunctions_glEnable uintptr
	addr_QOpenGLFunctions_glFinish uintptr
	addr_QOpenGLFunctions_glFlush uintptr
	addr_QOpenGLFunctions_glFrontFace uintptr
	addr_QOpenGLFunctions_glGenTextures uintptr
	addr_QOpenGLFunctions_glGetBooleanv uintptr
	addr_QOpenGLFunctions_glGetError uintptr
	addr_QOpenGLFunctions_glGetFloatv uintptr
	addr_QOpenGLFunctions_glGetIntegerv uintptr
	addr_QOpenGLFunctions_glGetTexParameterfv uintptr
	addr_QOpenGLFunctions_glGetTexParameteriv uintptr
	addr_QOpenGLFunctions_glHint uintptr
	addr_QOpenGLFunctions_glIsEnabled uintptr
	addr_QOpenGLFunctions_glIsTexture uintptr
	pg_QOpenGLFunctions_glLineWidth func(uintptr, float32)
	addr_QOpenGLFunctions_glPixelStorei uintptr
	pg_QOpenGLFunctions_glPolygonOffset func(uintptr, float32, float32)
	addr_QOpenGLFunctions_glScissor uintptr
	addr_QOpenGLFunctions_glStencilFunc uintptr
	addr_QOpenGLFunctions_glStencilMask uintptr
	addr_QOpenGLFunctions_glStencilOp uintptr
	pg_QOpenGLFunctions_glTexParameterf func(uintptr, int32, int32, float32)
	addr_QOpenGLFunctions_glTexParameteri uintptr
	addr_QOpenGLFunctions_glViewport uintptr
	addr_QOpenGLFunctions_glActiveTexture uintptr
	addr_QOpenGLFunctions_glAttachShader uintptr
	addr_QOpenGLFunctions_glBindAttribLocation uintptr
	addr_QOpenGLFunctions_glBindBuffer uintptr
	addr_QOpenGLFunctions_glBindFramebuffer uintptr
	addr_QOpenGLFunctions_glBindRenderbuffer uintptr
	pg_QOpenGLFunctions_glBlendColor func(uintptr, float32, float32, float32, float32)
	addr_QOpenGLFunctions_glBlendEquation uintptr
	addr_QOpenGLFunctions_glBlendEquationSeparate uintptr
	addr_QOpenGLFunctions_glBlendFuncSeparate uintptr
	addr_QOpenGLFunctions_glCheckFramebufferStatus uintptr
	pg_QOpenGLFunctions_glClearDepthf func(uintptr, float32)
	addr_QOpenGLFunctions_glCompileShader uintptr
	addr_QOpenGLFunctions_glCreateProgram uintptr
	addr_QOpenGLFunctions_glCreateShader uintptr
	addr_QOpenGLFunctions_glDeleteProgram uintptr
	addr_QOpenGLFunctions_glDeleteShader uintptr
	pg_QOpenGLFunctions_glDepthRangef func(uintptr, float32, float32)
	addr_QOpenGLFunctions_glDetachShader uintptr
	addr_QOpenGLFunctions_glDisableVertexAttribArray uintptr
	addr_QOpenGLFunctions_glEnableVertexAttribArray uintptr
	addr_QOpenGLFunctions_glFramebufferRenderbuffer uintptr
	addr_QOpenGLFunctions_glFramebufferTexture2D uintptr
	addr_QOpenGLFunctions_glGenBuffers uintptr
	addr_QOpenGLFunctions_glGenerateMipmap uintptr
	addr_QOpenGLFunctions_glGenFramebuffers uintptr
	addr_QOpenGLFunctions_glGenRenderbuffers uintptr
	addr_QOpenGLFunctions_glGetActiveAttrib uintptr
	addr_QOpenGLFunctions_glGetActiveUniform uintptr
	addr_QOpenGLFunctions_glGetAttachedShaders uintptr
	addr_QOpenGLFunctions_glGetAttribLocation uintptr
	addr_QOpenGLFunctions_glGetBufferParameteriv uintptr
	addr_QOpenGLFunctions_glGetFramebufferAttachmentParameteriv uintptr
	addr_QOpenGLFunctions_glGetProgramiv uintptr
	addr_QOpenGLFunctions_glGetProgramInfoLog uintptr
	addr_QOpenGLFunctions_glGetRenderbufferParameteriv uintptr
	addr_QOpenGLFunctions_glGetShaderiv uintptr
	addr_QOpenGLFunctions_glGetShaderInfoLog uintptr
	addr_QOpenGLFunctions_glGetShaderPrecisionFormat uintptr
	addr_QOpenGLFunctions_glGetShaderSource uintptr
	addr_QOpenGLFunctions_glGetUniformfv uintptr
	addr_QOpenGLFunctions_glGetUniformiv uintptr
	addr_QOpenGLFunctions_glGetUniformLocation uintptr
	addr_QOpenGLFunctions_glGetVertexAttribfv uintptr
	addr_QOpenGLFunctions_glGetVertexAttribiv uintptr
	addr_QOpenGLFunctions_glIsBuffer uintptr
	addr_QOpenGLFunctions_glIsFramebuffer uintptr
	addr_QOpenGLFunctions_glIsProgram uintptr
	addr_QOpenGLFunctions_glIsRenderbuffer uintptr
	addr_QOpenGLFunctions_glIsShader uintptr
	addr_QOpenGLFunctions_glLinkProgram uintptr
	addr_QOpenGLFunctions_glReleaseShaderCompiler uintptr
	addr_QOpenGLFunctions_glRenderbufferStorage uintptr
	pg_QOpenGLFunctions_glSampleCoverage func(uintptr, float32, uint8)
	addr_QOpenGLFunctions_glStencilFuncSeparate uintptr
	addr_QOpenGLFunctions_glStencilMaskSeparate uintptr
	addr_QOpenGLFunctions_glStencilOpSeparate uintptr
	pg_QOpenGLFunctions_glUniform1f func(uintptr, int32, float32)
	addr_QOpenGLFunctions_glUniform1i uintptr
	pg_QOpenGLFunctions_glUniform2f func(uintptr, int32, float32, float32)
	addr_QOpenGLFunctions_glUniform2i uintptr
	pg_QOpenGLFunctions_glUniform3f func(uintptr, int32, float32, float32, float32)
	addr_QOpenGLFunctions_glUniform3i uintptr
	pg_QOpenGLFunctions_glUniform4f func(uintptr, int32, float32, float32, float32, float32)
	addr_QOpenGLFunctions_glUniform4i uintptr
	addr_QOpenGLFunctions_glUseProgram uintptr
	addr_QOpenGLFunctions_glValidateProgram uintptr
	pg_QOpenGLFunctions_glVertexAttrib1f func(uintptr, int32, float32)
	pg_QOpenGLFunctions_glVertexAttrib2f func(uintptr, int32, float32, float32)
	pg_QOpenGLFunctions_glVertexAttrib3f func(uintptr, int32, float32, float32, float32)
	pg_QOpenGLFunctions_glVertexAttrib4f func(uintptr, int32, float32, float32, float32, float32)
	pg_QOpenGLFunctions_delete func(uintptr)
)

func init() {
	symaddr(&addr_QOpenGLFunctions_new, "QOpenGLFunctions_new")
	symaddr(&addr_QOpenGLFunctions_new2, "QOpenGLFunctions_new2")
	symaddr(&addr_QOpenGLFunctions_new3, "QOpenGLFunctions_new3")
	symaddr(&addr_QOpenGLFunctions_openGLFeatures, "QOpenGLFunctions_openGLFeatures")
	symaddr(&addr_QOpenGLFunctions_hasOpenGLFeature, "QOpenGLFunctions_hasOpenGLFeature")
	symaddr(&addr_QOpenGLFunctions_initializeOpenGLFunctions, "QOpenGLFunctions_initializeOpenGLFunctions")
	symaddr(&addr_QOpenGLFunctions_glBindTexture, "QOpenGLFunctions_glBindTexture")
	symaddr(&addr_QOpenGLFunctions_glBlendFunc, "QOpenGLFunctions_glBlendFunc")
	symaddr(&addr_QOpenGLFunctions_glClear, "QOpenGLFunctions_glClear")
	reg(&pg_QOpenGLFunctions_glClearColor, "QOpenGLFunctions_glClearColor")
	symaddr(&addr_QOpenGLFunctions_glClearStencil, "QOpenGLFunctions_glClearStencil")
	symaddr(&addr_QOpenGLFunctions_glColorMask, "QOpenGLFunctions_glColorMask")
	symaddr(&addr_QOpenGLFunctions_glCopyTexImage2D, "QOpenGLFunctions_glCopyTexImage2D")
	symaddr(&addr_QOpenGLFunctions_glCopyTexSubImage2D, "QOpenGLFunctions_glCopyTexSubImage2D")
	symaddr(&addr_QOpenGLFunctions_glCullFace, "QOpenGLFunctions_glCullFace")
	symaddr(&addr_QOpenGLFunctions_glDepthFunc, "QOpenGLFunctions_glDepthFunc")
	symaddr(&addr_QOpenGLFunctions_glDepthMask, "QOpenGLFunctions_glDepthMask")
	symaddr(&addr_QOpenGLFunctions_glDisable, "QOpenGLFunctions_glDisable")
	symaddr(&addr_QOpenGLFunctions_glDrawArrays, "QOpenGLFunctions_glDrawArrays")
	symaddr(&addr_QOpenGLFunctions_glEnable, "QOpenGLFunctions_glEnable")
	symaddr(&addr_QOpenGLFunctions_glFinish, "QOpenGLFunctions_glFinish")
	symaddr(&addr_QOpenGLFunctions_glFlush, "QOpenGLFunctions_glFlush")
	symaddr(&addr_QOpenGLFunctions_glFrontFace, "QOpenGLFunctions_glFrontFace")
	symaddr(&addr_QOpenGLFunctions_glGenTextures, "QOpenGLFunctions_glGenTextures")
	symaddr(&addr_QOpenGLFunctions_glGetBooleanv, "QOpenGLFunctions_glGetBooleanv")
	symaddr(&addr_QOpenGLFunctions_glGetError, "QOpenGLFunctions_glGetError")
	symaddr(&addr_QOpenGLFunctions_glGetFloatv, "QOpenGLFunctions_glGetFloatv")
	symaddr(&addr_QOpenGLFunctions_glGetIntegerv, "QOpenGLFunctions_glGetIntegerv")
	symaddr(&addr_QOpenGLFunctions_glGetTexParameterfv, "QOpenGLFunctions_glGetTexParameterfv")
	symaddr(&addr_QOpenGLFunctions_glGetTexParameteriv, "QOpenGLFunctions_glGetTexParameteriv")
	symaddr(&addr_QOpenGLFunctions_glHint, "QOpenGLFunctions_glHint")
	symaddr(&addr_QOpenGLFunctions_glIsEnabled, "QOpenGLFunctions_glIsEnabled")
	symaddr(&addr_QOpenGLFunctions_glIsTexture, "QOpenGLFunctions_glIsTexture")
	reg(&pg_QOpenGLFunctions_glLineWidth, "QOpenGLFunctions_glLineWidth")
	symaddr(&addr_QOpenGLFunctions_glPixelStorei, "QOpenGLFunctions_glPixelStorei")
	reg(&pg_QOpenGLFunctions_glPolygonOffset, "QOpenGLFunctions_glPolygonOffset")
	symaddr(&addr_QOpenGLFunctions_glScissor, "QOpenGLFunctions_glScissor")
	symaddr(&addr_QOpenGLFunctions_glStencilFunc, "QOpenGLFunctions_glStencilFunc")
	symaddr(&addr_QOpenGLFunctions_glStencilMask, "QOpenGLFunctions_glStencilMask")
	symaddr(&addr_QOpenGLFunctions_glStencilOp, "QOpenGLFunctions_glStencilOp")
	reg(&pg_QOpenGLFunctions_glTexParameterf, "QOpenGLFunctions_glTexParameterf")
	symaddr(&addr_QOpenGLFunctions_glTexParameteri, "QOpenGLFunctions_glTexParameteri")
	symaddr(&addr_QOpenGLFunctions_glViewport, "QOpenGLFunctions_glViewport")
	symaddr(&addr_QOpenGLFunctions_glActiveTexture, "QOpenGLFunctions_glActiveTexture")
	symaddr(&addr_QOpenGLFunctions_glAttachShader, "QOpenGLFunctions_glAttachShader")
	symaddr(&addr_QOpenGLFunctions_glBindAttribLocation, "QOpenGLFunctions_glBindAttribLocation")
	symaddr(&addr_QOpenGLFunctions_glBindBuffer, "QOpenGLFunctions_glBindBuffer")
	symaddr(&addr_QOpenGLFunctions_glBindFramebuffer, "QOpenGLFunctions_glBindFramebuffer")
	symaddr(&addr_QOpenGLFunctions_glBindRenderbuffer, "QOpenGLFunctions_glBindRenderbuffer")
	reg(&pg_QOpenGLFunctions_glBlendColor, "QOpenGLFunctions_glBlendColor")
	symaddr(&addr_QOpenGLFunctions_glBlendEquation, "QOpenGLFunctions_glBlendEquation")
	symaddr(&addr_QOpenGLFunctions_glBlendEquationSeparate, "QOpenGLFunctions_glBlendEquationSeparate")
	symaddr(&addr_QOpenGLFunctions_glBlendFuncSeparate, "QOpenGLFunctions_glBlendFuncSeparate")
	symaddr(&addr_QOpenGLFunctions_glCheckFramebufferStatus, "QOpenGLFunctions_glCheckFramebufferStatus")
	reg(&pg_QOpenGLFunctions_glClearDepthf, "QOpenGLFunctions_glClearDepthf")
	symaddr(&addr_QOpenGLFunctions_glCompileShader, "QOpenGLFunctions_glCompileShader")
	symaddr(&addr_QOpenGLFunctions_glCreateProgram, "QOpenGLFunctions_glCreateProgram")
	symaddr(&addr_QOpenGLFunctions_glCreateShader, "QOpenGLFunctions_glCreateShader")
	symaddr(&addr_QOpenGLFunctions_glDeleteProgram, "QOpenGLFunctions_glDeleteProgram")
	symaddr(&addr_QOpenGLFunctions_glDeleteShader, "QOpenGLFunctions_glDeleteShader")
	reg(&pg_QOpenGLFunctions_glDepthRangef, "QOpenGLFunctions_glDepthRangef")
	symaddr(&addr_QOpenGLFunctions_glDetachShader, "QOpenGLFunctions_glDetachShader")
	symaddr(&addr_QOpenGLFunctions_glDisableVertexAttribArray, "QOpenGLFunctions_glDisableVertexAttribArray")
	symaddr(&addr_QOpenGLFunctions_glEnableVertexAttribArray, "QOpenGLFunctions_glEnableVertexAttribArray")
	symaddr(&addr_QOpenGLFunctions_glFramebufferRenderbuffer, "QOpenGLFunctions_glFramebufferRenderbuffer")
	symaddr(&addr_QOpenGLFunctions_glFramebufferTexture2D, "QOpenGLFunctions_glFramebufferTexture2D")
	symaddr(&addr_QOpenGLFunctions_glGenBuffers, "QOpenGLFunctions_glGenBuffers")
	symaddr(&addr_QOpenGLFunctions_glGenerateMipmap, "QOpenGLFunctions_glGenerateMipmap")
	symaddr(&addr_QOpenGLFunctions_glGenFramebuffers, "QOpenGLFunctions_glGenFramebuffers")
	symaddr(&addr_QOpenGLFunctions_glGenRenderbuffers, "QOpenGLFunctions_glGenRenderbuffers")
	symaddr(&addr_QOpenGLFunctions_glGetActiveAttrib, "QOpenGLFunctions_glGetActiveAttrib")
	symaddr(&addr_QOpenGLFunctions_glGetActiveUniform, "QOpenGLFunctions_glGetActiveUniform")
	symaddr(&addr_QOpenGLFunctions_glGetAttachedShaders, "QOpenGLFunctions_glGetAttachedShaders")
	symaddr(&addr_QOpenGLFunctions_glGetAttribLocation, "QOpenGLFunctions_glGetAttribLocation")
	symaddr(&addr_QOpenGLFunctions_glGetBufferParameteriv, "QOpenGLFunctions_glGetBufferParameteriv")
	symaddr(&addr_QOpenGLFunctions_glGetFramebufferAttachmentParameteriv, "QOpenGLFunctions_glGetFramebufferAttachmentParameteriv")
	symaddr(&addr_QOpenGLFunctions_glGetProgramiv, "QOpenGLFunctions_glGetProgramiv")
	symaddr(&addr_QOpenGLFunctions_glGetProgramInfoLog, "QOpenGLFunctions_glGetProgramInfoLog")
	symaddr(&addr_QOpenGLFunctions_glGetRenderbufferParameteriv, "QOpenGLFunctions_glGetRenderbufferParameteriv")
	symaddr(&addr_QOpenGLFunctions_glGetShaderiv, "QOpenGLFunctions_glGetShaderiv")
	symaddr(&addr_QOpenGLFunctions_glGetShaderInfoLog, "QOpenGLFunctions_glGetShaderInfoLog")
	symaddr(&addr_QOpenGLFunctions_glGetShaderPrecisionFormat, "QOpenGLFunctions_glGetShaderPrecisionFormat")
	symaddr(&addr_QOpenGLFunctions_glGetShaderSource, "QOpenGLFunctions_glGetShaderSource")
	symaddr(&addr_QOpenGLFunctions_glGetUniformfv, "QOpenGLFunctions_glGetUniformfv")
	symaddr(&addr_QOpenGLFunctions_glGetUniformiv, "QOpenGLFunctions_glGetUniformiv")
	symaddr(&addr_QOpenGLFunctions_glGetUniformLocation, "QOpenGLFunctions_glGetUniformLocation")
	symaddr(&addr_QOpenGLFunctions_glGetVertexAttribfv, "QOpenGLFunctions_glGetVertexAttribfv")
	symaddr(&addr_QOpenGLFunctions_glGetVertexAttribiv, "QOpenGLFunctions_glGetVertexAttribiv")
	symaddr(&addr_QOpenGLFunctions_glIsBuffer, "QOpenGLFunctions_glIsBuffer")
	symaddr(&addr_QOpenGLFunctions_glIsFramebuffer, "QOpenGLFunctions_glIsFramebuffer")
	symaddr(&addr_QOpenGLFunctions_glIsProgram, "QOpenGLFunctions_glIsProgram")
	symaddr(&addr_QOpenGLFunctions_glIsRenderbuffer, "QOpenGLFunctions_glIsRenderbuffer")
	symaddr(&addr_QOpenGLFunctions_glIsShader, "QOpenGLFunctions_glIsShader")
	symaddr(&addr_QOpenGLFunctions_glLinkProgram, "QOpenGLFunctions_glLinkProgram")
	symaddr(&addr_QOpenGLFunctions_glReleaseShaderCompiler, "QOpenGLFunctions_glReleaseShaderCompiler")
	symaddr(&addr_QOpenGLFunctions_glRenderbufferStorage, "QOpenGLFunctions_glRenderbufferStorage")
	reg(&pg_QOpenGLFunctions_glSampleCoverage, "QOpenGLFunctions_glSampleCoverage")
	symaddr(&addr_QOpenGLFunctions_glStencilFuncSeparate, "QOpenGLFunctions_glStencilFuncSeparate")
	symaddr(&addr_QOpenGLFunctions_glStencilMaskSeparate, "QOpenGLFunctions_glStencilMaskSeparate")
	symaddr(&addr_QOpenGLFunctions_glStencilOpSeparate, "QOpenGLFunctions_glStencilOpSeparate")
	reg(&pg_QOpenGLFunctions_glUniform1f, "QOpenGLFunctions_glUniform1f")
	symaddr(&addr_QOpenGLFunctions_glUniform1i, "QOpenGLFunctions_glUniform1i")
	reg(&pg_QOpenGLFunctions_glUniform2f, "QOpenGLFunctions_glUniform2f")
	symaddr(&addr_QOpenGLFunctions_glUniform2i, "QOpenGLFunctions_glUniform2i")
	reg(&pg_QOpenGLFunctions_glUniform3f, "QOpenGLFunctions_glUniform3f")
	symaddr(&addr_QOpenGLFunctions_glUniform3i, "QOpenGLFunctions_glUniform3i")
	reg(&pg_QOpenGLFunctions_glUniform4f, "QOpenGLFunctions_glUniform4f")
	symaddr(&addr_QOpenGLFunctions_glUniform4i, "QOpenGLFunctions_glUniform4i")
	symaddr(&addr_QOpenGLFunctions_glUseProgram, "QOpenGLFunctions_glUseProgram")
	symaddr(&addr_QOpenGLFunctions_glValidateProgram, "QOpenGLFunctions_glValidateProgram")
	reg(&pg_QOpenGLFunctions_glVertexAttrib1f, "QOpenGLFunctions_glVertexAttrib1f")
	reg(&pg_QOpenGLFunctions_glVertexAttrib2f, "QOpenGLFunctions_glVertexAttrib2f")
	reg(&pg_QOpenGLFunctions_glVertexAttrib3f, "QOpenGLFunctions_glVertexAttrib3f")
	reg(&pg_QOpenGLFunctions_glVertexAttrib4f, "QOpenGLFunctions_glVertexAttrib4f")
	reg(&pg_QOpenGLFunctions_delete, "QOpenGLFunctions_delete")
}

type QOpenGLFunctions struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QOpenGLFunctions) UnsafeGoID() int64     { return this.vid }
func (this *QOpenGLFunctions) UnsafeSetGoID(id int64) { this.vid = id }

func bindQOpenGLFunctions(h uintptr) *QOpenGLFunctions {
	if h == 0 {
		return nil
	}
	return &QOpenGLFunctions{h: h}
}

func newownedQOpenGLFunctions(h uintptr) *QOpenGLFunctions {
	o := bindQOpenGLFunctions(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QOpenGLFunctions).Delete)
	}
	return o
}

func (this *QOpenGLFunctions) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOpenGLFunctions) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOpenGLFunctions(h unsafe.Pointer) *QOpenGLFunctions { return bindQOpenGLFunctions(uintptr(h)) }
func UnsafeNewQOpenGLFunctionsOwned(h unsafe.Pointer) *QOpenGLFunctions { return newownedQOpenGLFunctions(uintptr(h)) }

func (this *QOpenGLFunctions) UnsafeRelease() { this.h = 0 }

func NewQOpenGLFunctions() *QOpenGLFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLFunctions_new)
	__o := newownedQOpenGLFunctions(__r)
	return __o
}

func NewQOpenGLFunctions2(context *QOpenGLContext) *QOpenGLFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLFunctions_new2, uintptr(context.UnsafePointer()))
	__o := newownedQOpenGLFunctions(__r)
	runtime.KeepAlive(context)
	return __o
}

func NewQOpenGLFunctions3(arg1 *QOpenGLFunctions) *QOpenGLFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLFunctions_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQOpenGLFunctions(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QOpenGLFunctions) OpenGLFeatures() QOpenGLFunctions__OpenGLFeature {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_openGLFeatures, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QOpenGLFunctions__OpenGLFeature(__r)
}

func (this *QOpenGLFunctions) HasOpenGLFeature(feature QOpenGLFunctions__OpenGLFeature) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_hasOpenGLFeature, this.h, uintptr(feature))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QOpenGLFunctions) InitializeOpenGLFunctions() {
	purego.SyscallN(addr_QOpenGLFunctions_initializeOpenGLFunctions, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBindTexture(target uint, texture uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBindTexture, this.h, uintptr(target), uintptr(texture))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBlendFunc(sfactor uint, dfactor uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBlendFunc, this.h, uintptr(sfactor), uintptr(dfactor))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlClear(mask uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glClear, this.h, uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlClearColor(red float32, green float32, blue float32, alpha float32) {
	pg_QOpenGLFunctions_glClearColor(this.h, red, green, blue, alpha)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlClearStencil(s int) {
	purego.SyscallN(addr_QOpenGLFunctions_glClearStencil, this.h, uintptr(s))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlColorMask(red byte, green byte, blue byte, alpha byte) {
	purego.SyscallN(addr_QOpenGLFunctions_glColorMask, this.h, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlCopyTexImage2D(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	purego.SyscallN(addr_QOpenGLFunctions_glCopyTexImage2D, this.h, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(border))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlCopyTexSubImage2D(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	purego.SyscallN(addr_QOpenGLFunctions_glCopyTexSubImage2D, this.h, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlCullFace(mode uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glCullFace, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDepthFunc(func_ uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glDepthFunc, this.h, uintptr(func_))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDepthMask(flag byte) {
	purego.SyscallN(addr_QOpenGLFunctions_glDepthMask, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDisable(cap_ uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glDisable, this.h, uintptr(cap_))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDrawArrays(mode uint, first int, count int) {
	purego.SyscallN(addr_QOpenGLFunctions_glDrawArrays, this.h, uintptr(mode), uintptr(first), uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlEnable(cap_ uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glEnable, this.h, uintptr(cap_))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlFinish() {
	purego.SyscallN(addr_QOpenGLFunctions_glFinish, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlFlush() {
	purego.SyscallN(addr_QOpenGLFunctions_glFlush, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlFrontFace(mode uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glFrontFace, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlGenTextures(n int) uint {
	var textures_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGenTextures, this.h, uintptr(n), uintptr(unsafe.Pointer(&textures_out)))
	runtime.KeepAlive(this)
	return uint(textures_out)
}

func (this *QOpenGLFunctions) GlGetBooleanv(pname uint) byte {
	var params_out uint8
	purego.SyscallN(addr_QOpenGLFunctions_glGetBooleanv, this.h, uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return byte(params_out)
}

func (this *QOpenGLFunctions) GlGetError() uint {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glGetError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QOpenGLFunctions) GlGetFloatv(pname uint) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLFunctions_glGetFloatv, this.h, uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLFunctions) GlGetIntegerv(pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetIntegerv, this.h, uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetTexParameterfv(target uint, pname uint) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLFunctions_glGetTexParameterfv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLFunctions) GlGetTexParameteriv(target uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetTexParameteriv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlHint(target uint, mode uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glHint, this.h, uintptr(target), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlIsEnabled(cap_ uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsEnabled, this.h, uintptr(cap_))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlIsTexture(texture uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsTexture, this.h, uintptr(texture))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlLineWidth(width float32) {
	pg_QOpenGLFunctions_glLineWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlPixelStorei(pname uint, param int) {
	purego.SyscallN(addr_QOpenGLFunctions_glPixelStorei, this.h, uintptr(pname), uintptr(param))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlPolygonOffset(factor float32, units float32) {
	pg_QOpenGLFunctions_glPolygonOffset(this.h, factor, units)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlScissor(x int, y int, width int, height int) {
	purego.SyscallN(addr_QOpenGLFunctions_glScissor, this.h, uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlStencilFunc(func_ uint, ref int, mask uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glStencilFunc, this.h, uintptr(func_), uintptr(ref), uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlStencilMask(mask uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glStencilMask, this.h, uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlStencilOp(fail uint, zfail uint, zpass uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glStencilOp, this.h, uintptr(fail), uintptr(zfail), uintptr(zpass))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlTexParameterf(target uint, pname uint, param float32) {
	pg_QOpenGLFunctions_glTexParameterf(this.h, int32(target), int32(pname), param)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlTexParameteri(target uint, pname uint, param int) {
	purego.SyscallN(addr_QOpenGLFunctions_glTexParameteri, this.h, uintptr(target), uintptr(pname), uintptr(param))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlViewport(x int, y int, width int, height int) {
	purego.SyscallN(addr_QOpenGLFunctions_glViewport, this.h, uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlActiveTexture(texture uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glActiveTexture, this.h, uintptr(texture))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlAttachShader(program uint, shader uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glAttachShader, this.h, uintptr(program), uintptr(shader))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBindAttribLocation(program uint, index uint, name string) {
	name_b := append([]byte(name), 0)
	purego.SyscallN(addr_QOpenGLFunctions_glBindAttribLocation, this.h, uintptr(program), uintptr(index), uintptr(unsafe.Pointer(&name_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
}

func (this *QOpenGLFunctions) GlBindBuffer(target uint, buffer uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBindBuffer, this.h, uintptr(target), uintptr(buffer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBindFramebuffer(target uint, framebuffer uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBindFramebuffer, this.h, uintptr(target), uintptr(framebuffer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBindRenderbuffer(target uint, renderbuffer uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBindRenderbuffer, this.h, uintptr(target), uintptr(renderbuffer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBlendColor(red float32, green float32, blue float32, alpha float32) {
	pg_QOpenGLFunctions_glBlendColor(this.h, red, green, blue, alpha)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBlendEquation(mode uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBlendEquation, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBlendEquationSeparate, this.h, uintptr(modeRGB), uintptr(modeAlpha))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlBlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glBlendFuncSeparate, this.h, uintptr(srcRGB), uintptr(dstRGB), uintptr(srcAlpha), uintptr(dstAlpha))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlCheckFramebufferStatus(target uint) uint {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glCheckFramebufferStatus, this.h, uintptr(target))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QOpenGLFunctions) GlClearDepthf(depth float32) {
	pg_QOpenGLFunctions_glClearDepthf(this.h, depth)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlCompileShader(shader uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glCompileShader, this.h, uintptr(shader))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlCreateProgram() uint {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glCreateProgram, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QOpenGLFunctions) GlCreateShader(type_ uint) uint {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glCreateShader, this.h, uintptr(type_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QOpenGLFunctions) GlDeleteProgram(program uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glDeleteProgram, this.h, uintptr(program))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDeleteShader(shader uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glDeleteShader, this.h, uintptr(shader))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDepthRangef(zNear float32, zFar float32) {
	pg_QOpenGLFunctions_glDepthRangef(this.h, zNear, zFar)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDetachShader(program uint, shader uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glDetachShader, this.h, uintptr(program), uintptr(shader))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlDisableVertexAttribArray(index uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glDisableVertexAttribArray, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlEnableVertexAttribArray(index uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glEnableVertexAttribArray, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlFramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glFramebufferRenderbuffer, this.h, uintptr(target), uintptr(attachment), uintptr(renderbuffertarget), uintptr(renderbuffer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlFramebufferTexture2D(target uint, attachment uint, textarget uint, texture uint, level int) {
	purego.SyscallN(addr_QOpenGLFunctions_glFramebufferTexture2D, this.h, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlGenBuffers(n int) uint {
	var buffers_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGenBuffers, this.h, uintptr(n), uintptr(unsafe.Pointer(&buffers_out)))
	runtime.KeepAlive(this)
	return uint(buffers_out)
}

func (this *QOpenGLFunctions) GlGenerateMipmap(target uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glGenerateMipmap, this.h, uintptr(target))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlGenFramebuffers(n int) uint {
	var framebuffers_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGenFramebuffers, this.h, uintptr(n), uintptr(unsafe.Pointer(&framebuffers_out)))
	runtime.KeepAlive(this)
	return uint(framebuffers_out)
}

func (this *QOpenGLFunctions) GlGenRenderbuffers(n int) uint {
	var renderbuffers_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGenRenderbuffers, this.h, uintptr(n), uintptr(unsafe.Pointer(&renderbuffers_out)))
	runtime.KeepAlive(this)
	return uint(renderbuffers_out)
}

func (this *QOpenGLFunctions) GlGetActiveAttrib(program uint, index uint, bufsize int, name []byte) (int, int, uint) {
	var length_out int32
	var size_out int32
	var type__out int32
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(&name[0]))
	}
	purego.SyscallN(addr_QOpenGLFunctions_glGetActiveAttrib, this.h, uintptr(program), uintptr(index), uintptr(bufsize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&size_out)), uintptr(unsafe.Pointer(&type__out)), name_p)
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return int(length_out), int(size_out), uint(type__out)
}

func (this *QOpenGLFunctions) GlGetActiveUniform(program uint, index uint, bufsize int, name []byte) (int, int, uint) {
	var length_out int32
	var size_out int32
	var type__out int32
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(&name[0]))
	}
	purego.SyscallN(addr_QOpenGLFunctions_glGetActiveUniform, this.h, uintptr(program), uintptr(index), uintptr(bufsize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&size_out)), uintptr(unsafe.Pointer(&type__out)), name_p)
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return int(length_out), int(size_out), uint(type__out)
}

func (this *QOpenGLFunctions) GlGetAttachedShaders(program uint, maxcount int) (int, uint) {
	var count_out int32
	var shaders_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetAttachedShaders, this.h, uintptr(program), uintptr(maxcount), uintptr(unsafe.Pointer(&count_out)), uintptr(unsafe.Pointer(&shaders_out)))
	runtime.KeepAlive(this)
	return int(count_out), uint(shaders_out)
}

func (this *QOpenGLFunctions) GlGetAttribLocation(program uint, name string) int {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glGetAttribLocation, this.h, uintptr(program), uintptr(unsafe.Pointer(&name_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	return int(__r)
}

func (this *QOpenGLFunctions) GlGetBufferParameteriv(target uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetBufferParameteriv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetFramebufferAttachmentParameteriv(target uint, attachment uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetFramebufferAttachmentParameteriv, this.h, uintptr(target), uintptr(attachment), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetProgramiv(program uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetProgramiv, this.h, uintptr(program), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetProgramInfoLog(program uint, bufsize int, infolog []byte) int {
	var length_out int32
	var infolog_p uintptr
	if len(infolog) > 0 {
		infolog_p = uintptr(unsafe.Pointer(&infolog[0]))
	}
	purego.SyscallN(addr_QOpenGLFunctions_glGetProgramInfoLog, this.h, uintptr(program), uintptr(bufsize), uintptr(unsafe.Pointer(&length_out)), infolog_p)
	runtime.KeepAlive(this)
	runtime.KeepAlive(infolog)
	return int(length_out)
}

func (this *QOpenGLFunctions) GlGetRenderbufferParameteriv(target uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetRenderbufferParameteriv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetShaderiv(shader uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetShaderiv, this.h, uintptr(shader), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetShaderInfoLog(shader uint, bufsize int, infolog []byte) int {
	var length_out int32
	var infolog_p uintptr
	if len(infolog) > 0 {
		infolog_p = uintptr(unsafe.Pointer(&infolog[0]))
	}
	purego.SyscallN(addr_QOpenGLFunctions_glGetShaderInfoLog, this.h, uintptr(shader), uintptr(bufsize), uintptr(unsafe.Pointer(&length_out)), infolog_p)
	runtime.KeepAlive(this)
	runtime.KeepAlive(infolog)
	return int(length_out)
}

func (this *QOpenGLFunctions) GlGetShaderPrecisionFormat(shadertype uint, precisiontype uint) (int, int) {
	var range__out int32
	var precision_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetShaderPrecisionFormat, this.h, uintptr(shadertype), uintptr(precisiontype), uintptr(unsafe.Pointer(&range__out)), uintptr(unsafe.Pointer(&precision_out)))
	runtime.KeepAlive(this)
	return int(range__out), int(precision_out)
}

func (this *QOpenGLFunctions) GlGetShaderSource(shader uint, bufsize int, source []byte) int {
	var length_out int32
	var source_p uintptr
	if len(source) > 0 {
		source_p = uintptr(unsafe.Pointer(&source[0]))
	}
	purego.SyscallN(addr_QOpenGLFunctions_glGetShaderSource, this.h, uintptr(shader), uintptr(bufsize), uintptr(unsafe.Pointer(&length_out)), source_p)
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
	return int(length_out)
}

func (this *QOpenGLFunctions) GlGetUniformfv(program uint, location int) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLFunctions_glGetUniformfv, this.h, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLFunctions) GlGetUniformiv(program uint, location int) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetUniformiv, this.h, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlGetUniformLocation(program uint, name string) int {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glGetUniformLocation, this.h, uintptr(program), uintptr(unsafe.Pointer(&name_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	return int(__r)
}

func (this *QOpenGLFunctions) GlGetVertexAttribfv(index uint, pname uint) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLFunctions_glGetVertexAttribfv, this.h, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLFunctions) GlGetVertexAttribiv(index uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLFunctions_glGetVertexAttribiv, this.h, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLFunctions) GlIsBuffer(buffer uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsBuffer, this.h, uintptr(buffer))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlIsFramebuffer(framebuffer uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsFramebuffer, this.h, uintptr(framebuffer))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlIsProgram(program uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsProgram, this.h, uintptr(program))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlIsRenderbuffer(renderbuffer uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsRenderbuffer, this.h, uintptr(renderbuffer))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlIsShader(shader uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLFunctions_glIsShader, this.h, uintptr(shader))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLFunctions) GlLinkProgram(program uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glLinkProgram, this.h, uintptr(program))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlReleaseShaderCompiler() {
	purego.SyscallN(addr_QOpenGLFunctions_glReleaseShaderCompiler, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlRenderbufferStorage(target uint, internalformat uint, width int, height int) {
	purego.SyscallN(addr_QOpenGLFunctions_glRenderbufferStorage, this.h, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlSampleCoverage(value float32, invert byte) {
	pg_QOpenGLFunctions_glSampleCoverage(this.h, value, uint8(invert))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlStencilFuncSeparate(face uint, func_ uint, ref int, mask uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glStencilFuncSeparate, this.h, uintptr(face), uintptr(func_), uintptr(ref), uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlStencilMaskSeparate(face uint, mask uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glStencilMaskSeparate, this.h, uintptr(face), uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlStencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glStencilOpSeparate, this.h, uintptr(face), uintptr(fail), uintptr(zfail), uintptr(zpass))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform1f(location int, x float32) {
	pg_QOpenGLFunctions_glUniform1f(this.h, int32(location), x)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform1i(location int, x int) {
	purego.SyscallN(addr_QOpenGLFunctions_glUniform1i, this.h, uintptr(location), uintptr(x))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform2f(location int, x float32, y float32) {
	pg_QOpenGLFunctions_glUniform2f(this.h, int32(location), x, y)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform2i(location int, x int, y int) {
	purego.SyscallN(addr_QOpenGLFunctions_glUniform2i, this.h, uintptr(location), uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform3f(location int, x float32, y float32, z float32) {
	pg_QOpenGLFunctions_glUniform3f(this.h, int32(location), x, y, z)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform3i(location int, x int, y int, z int) {
	purego.SyscallN(addr_QOpenGLFunctions_glUniform3i, this.h, uintptr(location), uintptr(x), uintptr(y), uintptr(z))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform4f(location int, x float32, y float32, z float32, w float32) {
	pg_QOpenGLFunctions_glUniform4f(this.h, int32(location), x, y, z, w)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUniform4i(location int, x int, y int, z int, w int) {
	purego.SyscallN(addr_QOpenGLFunctions_glUniform4i, this.h, uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlUseProgram(program uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glUseProgram, this.h, uintptr(program))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlValidateProgram(program uint) {
	purego.SyscallN(addr_QOpenGLFunctions_glValidateProgram, this.h, uintptr(program))
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlVertexAttrib1f(indx uint, x float32) {
	pg_QOpenGLFunctions_glVertexAttrib1f(this.h, int32(indx), x)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlVertexAttrib2f(indx uint, x float32, y float32) {
	pg_QOpenGLFunctions_glVertexAttrib2f(this.h, int32(indx), x, y)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlVertexAttrib3f(indx uint, x float32, y float32, z float32) {
	pg_QOpenGLFunctions_glVertexAttrib3f(this.h, int32(indx), x, y, z)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) GlVertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	pg_QOpenGLFunctions_glVertexAttrib4f(this.h, int32(indx), x, y, z, w)
	runtime.KeepAlive(this)
}

func (this *QOpenGLFunctions) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOpenGLFunctions_delete(this.h)
	}
	this.UnsafeRelease()
}
