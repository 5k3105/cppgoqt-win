package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QOpenGLExtraFunctions_base func(uintptr) uintptr
	addr_QOpenGLExtraFunctions_new uintptr
	addr_QOpenGLExtraFunctions_new2 uintptr
	addr_QOpenGLExtraFunctions_new3 uintptr
	addr_QOpenGLExtraFunctions_glReadBuffer uintptr
	addr_QOpenGLExtraFunctions_glCopyTexSubImage3D uintptr
	addr_QOpenGLExtraFunctions_glGenQueries uintptr
	addr_QOpenGLExtraFunctions_glIsQuery uintptr
	addr_QOpenGLExtraFunctions_glBeginQuery uintptr
	addr_QOpenGLExtraFunctions_glEndQuery uintptr
	addr_QOpenGLExtraFunctions_glGetQueryiv uintptr
	addr_QOpenGLExtraFunctions_glGetQueryObjectuiv uintptr
	addr_QOpenGLExtraFunctions_glUnmapBuffer uintptr
	addr_QOpenGLExtraFunctions_glBlitFramebuffer uintptr
	addr_QOpenGLExtraFunctions_glRenderbufferStorageMultisample uintptr
	addr_QOpenGLExtraFunctions_glFramebufferTextureLayer uintptr
	addr_QOpenGLExtraFunctions_glFlushMappedBufferRange uintptr
	addr_QOpenGLExtraFunctions_glBindVertexArray uintptr
	addr_QOpenGLExtraFunctions_glGenVertexArrays uintptr
	addr_QOpenGLExtraFunctions_glIsVertexArray uintptr
	addr_QOpenGLExtraFunctions_glGetIntegeri_v uintptr
	addr_QOpenGLExtraFunctions_glBeginTransformFeedback uintptr
	addr_QOpenGLExtraFunctions_glEndTransformFeedback uintptr
	addr_QOpenGLExtraFunctions_glBindBufferRange uintptr
	addr_QOpenGLExtraFunctions_glBindBufferBase uintptr
	addr_QOpenGLExtraFunctions_glGetTransformFeedbackVarying uintptr
	addr_QOpenGLExtraFunctions_glGetVertexAttribIiv uintptr
	addr_QOpenGLExtraFunctions_glGetVertexAttribIuiv uintptr
	addr_QOpenGLExtraFunctions_glVertexAttribI4i uintptr
	addr_QOpenGLExtraFunctions_glVertexAttribI4ui uintptr
	addr_QOpenGLExtraFunctions_glGetUniformuiv uintptr
	addr_QOpenGLExtraFunctions_glUniform1ui uintptr
	addr_QOpenGLExtraFunctions_glUniform2ui uintptr
	addr_QOpenGLExtraFunctions_glUniform3ui uintptr
	addr_QOpenGLExtraFunctions_glUniform4ui uintptr
	pg_QOpenGLExtraFunctions_glClearBufferfi func(uintptr, int32, int32, float32, int32)
	addr_QOpenGLExtraFunctions_glCopyBufferSubData uintptr
	addr_QOpenGLExtraFunctions_glGetActiveUniformBlockiv uintptr
	addr_QOpenGLExtraFunctions_glGetActiveUniformBlockName uintptr
	addr_QOpenGLExtraFunctions_glUniformBlockBinding uintptr
	addr_QOpenGLExtraFunctions_glDrawArraysInstanced uintptr
	addr_QOpenGLExtraFunctions_glGetInteger64v uintptr
	addr_QOpenGLExtraFunctions_glGetInteger64i_v uintptr
	addr_QOpenGLExtraFunctions_glGetBufferParameteri64v uintptr
	addr_QOpenGLExtraFunctions_glGenSamplers uintptr
	addr_QOpenGLExtraFunctions_glIsSampler uintptr
	addr_QOpenGLExtraFunctions_glBindSampler uintptr
	addr_QOpenGLExtraFunctions_glSamplerParameteri uintptr
	pg_QOpenGLExtraFunctions_glSamplerParameterf func(uintptr, int32, int32, float32)
	addr_QOpenGLExtraFunctions_glGetSamplerParameteriv uintptr
	addr_QOpenGLExtraFunctions_glGetSamplerParameterfv uintptr
	addr_QOpenGLExtraFunctions_glVertexAttribDivisor uintptr
	addr_QOpenGLExtraFunctions_glBindTransformFeedback uintptr
	addr_QOpenGLExtraFunctions_glGenTransformFeedbacks uintptr
	addr_QOpenGLExtraFunctions_glIsTransformFeedback uintptr
	addr_QOpenGLExtraFunctions_glPauseTransformFeedback uintptr
	addr_QOpenGLExtraFunctions_glResumeTransformFeedback uintptr
	addr_QOpenGLExtraFunctions_glProgramParameteri uintptr
	addr_QOpenGLExtraFunctions_glTexStorage2D uintptr
	addr_QOpenGLExtraFunctions_glTexStorage3D uintptr
	addr_QOpenGLExtraFunctions_glGetInternalformativ uintptr
	addr_QOpenGLExtraFunctions_glDispatchCompute uintptr
	addr_QOpenGLExtraFunctions_glDispatchComputeIndirect uintptr
	addr_QOpenGLExtraFunctions_glFramebufferParameteri uintptr
	addr_QOpenGLExtraFunctions_glGetFramebufferParameteriv uintptr
	addr_QOpenGLExtraFunctions_glGetProgramInterfaceiv uintptr
	addr_QOpenGLExtraFunctions_glGetProgramResourceName uintptr
	addr_QOpenGLExtraFunctions_glUseProgramStages uintptr
	addr_QOpenGLExtraFunctions_glActiveShaderProgram uintptr
	addr_QOpenGLExtraFunctions_glBindProgramPipeline uintptr
	addr_QOpenGLExtraFunctions_glGenProgramPipelines uintptr
	addr_QOpenGLExtraFunctions_glIsProgramPipeline uintptr
	addr_QOpenGLExtraFunctions_glGetProgramPipelineiv uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform1i uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform2i uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform3i uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform4i uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform1ui uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform2ui uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform3ui uintptr
	addr_QOpenGLExtraFunctions_glProgramUniform4ui uintptr
	pg_QOpenGLExtraFunctions_glProgramUniform1f func(uintptr, int32, int32, float32)
	pg_QOpenGLExtraFunctions_glProgramUniform2f func(uintptr, int32, int32, float32, float32)
	pg_QOpenGLExtraFunctions_glProgramUniform3f func(uintptr, int32, int32, float32, float32, float32)
	pg_QOpenGLExtraFunctions_glProgramUniform4f func(uintptr, int32, int32, float32, float32, float32, float32)
	addr_QOpenGLExtraFunctions_glValidateProgramPipeline uintptr
	addr_QOpenGLExtraFunctions_glGetProgramPipelineInfoLog uintptr
	addr_QOpenGLExtraFunctions_glBindImageTexture uintptr
	addr_QOpenGLExtraFunctions_glGetBooleani_v uintptr
	addr_QOpenGLExtraFunctions_glMemoryBarrier uintptr
	addr_QOpenGLExtraFunctions_glMemoryBarrierByRegion uintptr
	addr_QOpenGLExtraFunctions_glTexStorage2DMultisample uintptr
	addr_QOpenGLExtraFunctions_glGetMultisamplefv uintptr
	addr_QOpenGLExtraFunctions_glSampleMaski uintptr
	addr_QOpenGLExtraFunctions_glGetTexLevelParameteriv uintptr
	addr_QOpenGLExtraFunctions_glGetTexLevelParameterfv uintptr
	addr_QOpenGLExtraFunctions_glBindVertexBuffer uintptr
	addr_QOpenGLExtraFunctions_glVertexAttribFormat uintptr
	addr_QOpenGLExtraFunctions_glVertexAttribIFormat uintptr
	addr_QOpenGLExtraFunctions_glVertexAttribBinding uintptr
	addr_QOpenGLExtraFunctions_glVertexBindingDivisor uintptr
	addr_QOpenGLExtraFunctions_glBlendBarrier uintptr
	addr_QOpenGLExtraFunctions_glCopyImageSubData uintptr
	addr_QOpenGLExtraFunctions_glGetDebugMessageLog uintptr
	addr_QOpenGLExtraFunctions_glPopDebugGroup uintptr
	addr_QOpenGLExtraFunctions_glGetObjectLabel uintptr
	addr_QOpenGLExtraFunctions_glEnablei uintptr
	addr_QOpenGLExtraFunctions_glDisablei uintptr
	addr_QOpenGLExtraFunctions_glBlendEquationi uintptr
	addr_QOpenGLExtraFunctions_glBlendEquationSeparatei uintptr
	addr_QOpenGLExtraFunctions_glBlendFunci uintptr
	addr_QOpenGLExtraFunctions_glBlendFuncSeparatei uintptr
	addr_QOpenGLExtraFunctions_glColorMaski uintptr
	addr_QOpenGLExtraFunctions_glIsEnabledi uintptr
	addr_QOpenGLExtraFunctions_glFramebufferTexture uintptr
	pg_QOpenGLExtraFunctions_glPrimitiveBoundingBox func(uintptr, float32, float32, float32, float32, float32, float32, float32, float32)
	addr_QOpenGLExtraFunctions_glGetGraphicsResetStatus uintptr
	addr_QOpenGLExtraFunctions_glGetnUniformfv uintptr
	addr_QOpenGLExtraFunctions_glGetnUniformiv uintptr
	addr_QOpenGLExtraFunctions_glGetnUniformuiv uintptr
	pg_QOpenGLExtraFunctions_glMinSampleShading func(uintptr, float32)
	addr_QOpenGLExtraFunctions_glPatchParameteri uintptr
	addr_QOpenGLExtraFunctions_glGetTexParameterIiv uintptr
	addr_QOpenGLExtraFunctions_glGetTexParameterIuiv uintptr
	addr_QOpenGLExtraFunctions_glGetSamplerParameterIiv uintptr
	addr_QOpenGLExtraFunctions_glGetSamplerParameterIuiv uintptr
	addr_QOpenGLExtraFunctions_glTexBuffer uintptr
	addr_QOpenGLExtraFunctions_glTexBufferRange uintptr
	addr_QOpenGLExtraFunctions_glTexStorage3DMultisample uintptr
	pg_QOpenGLExtraFunctions_delete func(uintptr)
)

func init() {
	reg(&pg_QOpenGLExtraFunctions_base, "QOpenGLExtraFunctions_base")
	symaddr(&addr_QOpenGLExtraFunctions_new, "QOpenGLExtraFunctions_new")
	symaddr(&addr_QOpenGLExtraFunctions_new2, "QOpenGLExtraFunctions_new2")
	symaddr(&addr_QOpenGLExtraFunctions_new3, "QOpenGLExtraFunctions_new3")
	symaddr(&addr_QOpenGLExtraFunctions_glReadBuffer, "QOpenGLExtraFunctions_glReadBuffer")
	symaddr(&addr_QOpenGLExtraFunctions_glCopyTexSubImage3D, "QOpenGLExtraFunctions_glCopyTexSubImage3D")
	symaddr(&addr_QOpenGLExtraFunctions_glGenQueries, "QOpenGLExtraFunctions_glGenQueries")
	symaddr(&addr_QOpenGLExtraFunctions_glIsQuery, "QOpenGLExtraFunctions_glIsQuery")
	symaddr(&addr_QOpenGLExtraFunctions_glBeginQuery, "QOpenGLExtraFunctions_glBeginQuery")
	symaddr(&addr_QOpenGLExtraFunctions_glEndQuery, "QOpenGLExtraFunctions_glEndQuery")
	symaddr(&addr_QOpenGLExtraFunctions_glGetQueryiv, "QOpenGLExtraFunctions_glGetQueryiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetQueryObjectuiv, "QOpenGLExtraFunctions_glGetQueryObjectuiv")
	symaddr(&addr_QOpenGLExtraFunctions_glUnmapBuffer, "QOpenGLExtraFunctions_glUnmapBuffer")
	symaddr(&addr_QOpenGLExtraFunctions_glBlitFramebuffer, "QOpenGLExtraFunctions_glBlitFramebuffer")
	symaddr(&addr_QOpenGLExtraFunctions_glRenderbufferStorageMultisample, "QOpenGLExtraFunctions_glRenderbufferStorageMultisample")
	symaddr(&addr_QOpenGLExtraFunctions_glFramebufferTextureLayer, "QOpenGLExtraFunctions_glFramebufferTextureLayer")
	symaddr(&addr_QOpenGLExtraFunctions_glFlushMappedBufferRange, "QOpenGLExtraFunctions_glFlushMappedBufferRange")
	symaddr(&addr_QOpenGLExtraFunctions_glBindVertexArray, "QOpenGLExtraFunctions_glBindVertexArray")
	symaddr(&addr_QOpenGLExtraFunctions_glGenVertexArrays, "QOpenGLExtraFunctions_glGenVertexArrays")
	symaddr(&addr_QOpenGLExtraFunctions_glIsVertexArray, "QOpenGLExtraFunctions_glIsVertexArray")
	symaddr(&addr_QOpenGLExtraFunctions_glGetIntegeri_v, "QOpenGLExtraFunctions_glGetIntegeri_v")
	symaddr(&addr_QOpenGLExtraFunctions_glBeginTransformFeedback, "QOpenGLExtraFunctions_glBeginTransformFeedback")
	symaddr(&addr_QOpenGLExtraFunctions_glEndTransformFeedback, "QOpenGLExtraFunctions_glEndTransformFeedback")
	symaddr(&addr_QOpenGLExtraFunctions_glBindBufferRange, "QOpenGLExtraFunctions_glBindBufferRange")
	symaddr(&addr_QOpenGLExtraFunctions_glBindBufferBase, "QOpenGLExtraFunctions_glBindBufferBase")
	symaddr(&addr_QOpenGLExtraFunctions_glGetTransformFeedbackVarying, "QOpenGLExtraFunctions_glGetTransformFeedbackVarying")
	symaddr(&addr_QOpenGLExtraFunctions_glGetVertexAttribIiv, "QOpenGLExtraFunctions_glGetVertexAttribIiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetVertexAttribIuiv, "QOpenGLExtraFunctions_glGetVertexAttribIuiv")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexAttribI4i, "QOpenGLExtraFunctions_glVertexAttribI4i")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexAttribI4ui, "QOpenGLExtraFunctions_glVertexAttribI4ui")
	symaddr(&addr_QOpenGLExtraFunctions_glGetUniformuiv, "QOpenGLExtraFunctions_glGetUniformuiv")
	symaddr(&addr_QOpenGLExtraFunctions_glUniform1ui, "QOpenGLExtraFunctions_glUniform1ui")
	symaddr(&addr_QOpenGLExtraFunctions_glUniform2ui, "QOpenGLExtraFunctions_glUniform2ui")
	symaddr(&addr_QOpenGLExtraFunctions_glUniform3ui, "QOpenGLExtraFunctions_glUniform3ui")
	symaddr(&addr_QOpenGLExtraFunctions_glUniform4ui, "QOpenGLExtraFunctions_glUniform4ui")
	reg(&pg_QOpenGLExtraFunctions_glClearBufferfi, "QOpenGLExtraFunctions_glClearBufferfi")
	symaddr(&addr_QOpenGLExtraFunctions_glCopyBufferSubData, "QOpenGLExtraFunctions_glCopyBufferSubData")
	symaddr(&addr_QOpenGLExtraFunctions_glGetActiveUniformBlockiv, "QOpenGLExtraFunctions_glGetActiveUniformBlockiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetActiveUniformBlockName, "QOpenGLExtraFunctions_glGetActiveUniformBlockName")
	symaddr(&addr_QOpenGLExtraFunctions_glUniformBlockBinding, "QOpenGLExtraFunctions_glUniformBlockBinding")
	symaddr(&addr_QOpenGLExtraFunctions_glDrawArraysInstanced, "QOpenGLExtraFunctions_glDrawArraysInstanced")
	symaddr(&addr_QOpenGLExtraFunctions_glGetInteger64v, "QOpenGLExtraFunctions_glGetInteger64v")
	symaddr(&addr_QOpenGLExtraFunctions_glGetInteger64i_v, "QOpenGLExtraFunctions_glGetInteger64i_v")
	symaddr(&addr_QOpenGLExtraFunctions_glGetBufferParameteri64v, "QOpenGLExtraFunctions_glGetBufferParameteri64v")
	symaddr(&addr_QOpenGLExtraFunctions_glGenSamplers, "QOpenGLExtraFunctions_glGenSamplers")
	symaddr(&addr_QOpenGLExtraFunctions_glIsSampler, "QOpenGLExtraFunctions_glIsSampler")
	symaddr(&addr_QOpenGLExtraFunctions_glBindSampler, "QOpenGLExtraFunctions_glBindSampler")
	symaddr(&addr_QOpenGLExtraFunctions_glSamplerParameteri, "QOpenGLExtraFunctions_glSamplerParameteri")
	reg(&pg_QOpenGLExtraFunctions_glSamplerParameterf, "QOpenGLExtraFunctions_glSamplerParameterf")
	symaddr(&addr_QOpenGLExtraFunctions_glGetSamplerParameteriv, "QOpenGLExtraFunctions_glGetSamplerParameteriv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetSamplerParameterfv, "QOpenGLExtraFunctions_glGetSamplerParameterfv")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexAttribDivisor, "QOpenGLExtraFunctions_glVertexAttribDivisor")
	symaddr(&addr_QOpenGLExtraFunctions_glBindTransformFeedback, "QOpenGLExtraFunctions_glBindTransformFeedback")
	symaddr(&addr_QOpenGLExtraFunctions_glGenTransformFeedbacks, "QOpenGLExtraFunctions_glGenTransformFeedbacks")
	symaddr(&addr_QOpenGLExtraFunctions_glIsTransformFeedback, "QOpenGLExtraFunctions_glIsTransformFeedback")
	symaddr(&addr_QOpenGLExtraFunctions_glPauseTransformFeedback, "QOpenGLExtraFunctions_glPauseTransformFeedback")
	symaddr(&addr_QOpenGLExtraFunctions_glResumeTransformFeedback, "QOpenGLExtraFunctions_glResumeTransformFeedback")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramParameteri, "QOpenGLExtraFunctions_glProgramParameteri")
	symaddr(&addr_QOpenGLExtraFunctions_glTexStorage2D, "QOpenGLExtraFunctions_glTexStorage2D")
	symaddr(&addr_QOpenGLExtraFunctions_glTexStorage3D, "QOpenGLExtraFunctions_glTexStorage3D")
	symaddr(&addr_QOpenGLExtraFunctions_glGetInternalformativ, "QOpenGLExtraFunctions_glGetInternalformativ")
	symaddr(&addr_QOpenGLExtraFunctions_glDispatchCompute, "QOpenGLExtraFunctions_glDispatchCompute")
	symaddr(&addr_QOpenGLExtraFunctions_glDispatchComputeIndirect, "QOpenGLExtraFunctions_glDispatchComputeIndirect")
	symaddr(&addr_QOpenGLExtraFunctions_glFramebufferParameteri, "QOpenGLExtraFunctions_glFramebufferParameteri")
	symaddr(&addr_QOpenGLExtraFunctions_glGetFramebufferParameteriv, "QOpenGLExtraFunctions_glGetFramebufferParameteriv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetProgramInterfaceiv, "QOpenGLExtraFunctions_glGetProgramInterfaceiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetProgramResourceName, "QOpenGLExtraFunctions_glGetProgramResourceName")
	symaddr(&addr_QOpenGLExtraFunctions_glUseProgramStages, "QOpenGLExtraFunctions_glUseProgramStages")
	symaddr(&addr_QOpenGLExtraFunctions_glActiveShaderProgram, "QOpenGLExtraFunctions_glActiveShaderProgram")
	symaddr(&addr_QOpenGLExtraFunctions_glBindProgramPipeline, "QOpenGLExtraFunctions_glBindProgramPipeline")
	symaddr(&addr_QOpenGLExtraFunctions_glGenProgramPipelines, "QOpenGLExtraFunctions_glGenProgramPipelines")
	symaddr(&addr_QOpenGLExtraFunctions_glIsProgramPipeline, "QOpenGLExtraFunctions_glIsProgramPipeline")
	symaddr(&addr_QOpenGLExtraFunctions_glGetProgramPipelineiv, "QOpenGLExtraFunctions_glGetProgramPipelineiv")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform1i, "QOpenGLExtraFunctions_glProgramUniform1i")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform2i, "QOpenGLExtraFunctions_glProgramUniform2i")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform3i, "QOpenGLExtraFunctions_glProgramUniform3i")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform4i, "QOpenGLExtraFunctions_glProgramUniform4i")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform1ui, "QOpenGLExtraFunctions_glProgramUniform1ui")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform2ui, "QOpenGLExtraFunctions_glProgramUniform2ui")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform3ui, "QOpenGLExtraFunctions_glProgramUniform3ui")
	symaddr(&addr_QOpenGLExtraFunctions_glProgramUniform4ui, "QOpenGLExtraFunctions_glProgramUniform4ui")
	reg(&pg_QOpenGLExtraFunctions_glProgramUniform1f, "QOpenGLExtraFunctions_glProgramUniform1f")
	reg(&pg_QOpenGLExtraFunctions_glProgramUniform2f, "QOpenGLExtraFunctions_glProgramUniform2f")
	reg(&pg_QOpenGLExtraFunctions_glProgramUniform3f, "QOpenGLExtraFunctions_glProgramUniform3f")
	reg(&pg_QOpenGLExtraFunctions_glProgramUniform4f, "QOpenGLExtraFunctions_glProgramUniform4f")
	symaddr(&addr_QOpenGLExtraFunctions_glValidateProgramPipeline, "QOpenGLExtraFunctions_glValidateProgramPipeline")
	symaddr(&addr_QOpenGLExtraFunctions_glGetProgramPipelineInfoLog, "QOpenGLExtraFunctions_glGetProgramPipelineInfoLog")
	symaddr(&addr_QOpenGLExtraFunctions_glBindImageTexture, "QOpenGLExtraFunctions_glBindImageTexture")
	symaddr(&addr_QOpenGLExtraFunctions_glGetBooleani_v, "QOpenGLExtraFunctions_glGetBooleani_v")
	symaddr(&addr_QOpenGLExtraFunctions_glMemoryBarrier, "QOpenGLExtraFunctions_glMemoryBarrier")
	symaddr(&addr_QOpenGLExtraFunctions_glMemoryBarrierByRegion, "QOpenGLExtraFunctions_glMemoryBarrierByRegion")
	symaddr(&addr_QOpenGLExtraFunctions_glTexStorage2DMultisample, "QOpenGLExtraFunctions_glTexStorage2DMultisample")
	symaddr(&addr_QOpenGLExtraFunctions_glGetMultisamplefv, "QOpenGLExtraFunctions_glGetMultisamplefv")
	symaddr(&addr_QOpenGLExtraFunctions_glSampleMaski, "QOpenGLExtraFunctions_glSampleMaski")
	symaddr(&addr_QOpenGLExtraFunctions_glGetTexLevelParameteriv, "QOpenGLExtraFunctions_glGetTexLevelParameteriv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetTexLevelParameterfv, "QOpenGLExtraFunctions_glGetTexLevelParameterfv")
	symaddr(&addr_QOpenGLExtraFunctions_glBindVertexBuffer, "QOpenGLExtraFunctions_glBindVertexBuffer")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexAttribFormat, "QOpenGLExtraFunctions_glVertexAttribFormat")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexAttribIFormat, "QOpenGLExtraFunctions_glVertexAttribIFormat")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexAttribBinding, "QOpenGLExtraFunctions_glVertexAttribBinding")
	symaddr(&addr_QOpenGLExtraFunctions_glVertexBindingDivisor, "QOpenGLExtraFunctions_glVertexBindingDivisor")
	symaddr(&addr_QOpenGLExtraFunctions_glBlendBarrier, "QOpenGLExtraFunctions_glBlendBarrier")
	symaddr(&addr_QOpenGLExtraFunctions_glCopyImageSubData, "QOpenGLExtraFunctions_glCopyImageSubData")
	symaddr(&addr_QOpenGLExtraFunctions_glGetDebugMessageLog, "QOpenGLExtraFunctions_glGetDebugMessageLog")
	symaddr(&addr_QOpenGLExtraFunctions_glPopDebugGroup, "QOpenGLExtraFunctions_glPopDebugGroup")
	symaddr(&addr_QOpenGLExtraFunctions_glGetObjectLabel, "QOpenGLExtraFunctions_glGetObjectLabel")
	symaddr(&addr_QOpenGLExtraFunctions_glEnablei, "QOpenGLExtraFunctions_glEnablei")
	symaddr(&addr_QOpenGLExtraFunctions_glDisablei, "QOpenGLExtraFunctions_glDisablei")
	symaddr(&addr_QOpenGLExtraFunctions_glBlendEquationi, "QOpenGLExtraFunctions_glBlendEquationi")
	symaddr(&addr_QOpenGLExtraFunctions_glBlendEquationSeparatei, "QOpenGLExtraFunctions_glBlendEquationSeparatei")
	symaddr(&addr_QOpenGLExtraFunctions_glBlendFunci, "QOpenGLExtraFunctions_glBlendFunci")
	symaddr(&addr_QOpenGLExtraFunctions_glBlendFuncSeparatei, "QOpenGLExtraFunctions_glBlendFuncSeparatei")
	symaddr(&addr_QOpenGLExtraFunctions_glColorMaski, "QOpenGLExtraFunctions_glColorMaski")
	symaddr(&addr_QOpenGLExtraFunctions_glIsEnabledi, "QOpenGLExtraFunctions_glIsEnabledi")
	symaddr(&addr_QOpenGLExtraFunctions_glFramebufferTexture, "QOpenGLExtraFunctions_glFramebufferTexture")
	reg(&pg_QOpenGLExtraFunctions_glPrimitiveBoundingBox, "QOpenGLExtraFunctions_glPrimitiveBoundingBox")
	symaddr(&addr_QOpenGLExtraFunctions_glGetGraphicsResetStatus, "QOpenGLExtraFunctions_glGetGraphicsResetStatus")
	symaddr(&addr_QOpenGLExtraFunctions_glGetnUniformfv, "QOpenGLExtraFunctions_glGetnUniformfv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetnUniformiv, "QOpenGLExtraFunctions_glGetnUniformiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetnUniformuiv, "QOpenGLExtraFunctions_glGetnUniformuiv")
	reg(&pg_QOpenGLExtraFunctions_glMinSampleShading, "QOpenGLExtraFunctions_glMinSampleShading")
	symaddr(&addr_QOpenGLExtraFunctions_glPatchParameteri, "QOpenGLExtraFunctions_glPatchParameteri")
	symaddr(&addr_QOpenGLExtraFunctions_glGetTexParameterIiv, "QOpenGLExtraFunctions_glGetTexParameterIiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetTexParameterIuiv, "QOpenGLExtraFunctions_glGetTexParameterIuiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetSamplerParameterIiv, "QOpenGLExtraFunctions_glGetSamplerParameterIiv")
	symaddr(&addr_QOpenGLExtraFunctions_glGetSamplerParameterIuiv, "QOpenGLExtraFunctions_glGetSamplerParameterIuiv")
	symaddr(&addr_QOpenGLExtraFunctions_glTexBuffer, "QOpenGLExtraFunctions_glTexBuffer")
	symaddr(&addr_QOpenGLExtraFunctions_glTexBufferRange, "QOpenGLExtraFunctions_glTexBufferRange")
	symaddr(&addr_QOpenGLExtraFunctions_glTexStorage3DMultisample, "QOpenGLExtraFunctions_glTexStorage3DMultisample")
	reg(&pg_QOpenGLExtraFunctions_delete, "QOpenGLExtraFunctions_delete")
}

type QOpenGLExtraFunctions struct {
	h uintptr
	*QOpenGLFunctions
}

func bindQOpenGLExtraFunctions(h uintptr) *QOpenGLExtraFunctions {
	if h == 0 {
		return nil
	}
	return &QOpenGLExtraFunctions{h: h, QOpenGLFunctions: bindQOpenGLFunctions(pg_QOpenGLExtraFunctions_base(h))}
}

func newownedQOpenGLExtraFunctions(h uintptr) *QOpenGLExtraFunctions {
	o := bindQOpenGLExtraFunctions(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QOpenGLExtraFunctions).Delete)
	}
	return o
}

func (this *QOpenGLExtraFunctions) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOpenGLExtraFunctions) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOpenGLExtraFunctions(h unsafe.Pointer) *QOpenGLExtraFunctions { return bindQOpenGLExtraFunctions(uintptr(h)) }
func UnsafeNewQOpenGLExtraFunctionsOwned(h unsafe.Pointer) *QOpenGLExtraFunctions { return newownedQOpenGLExtraFunctions(uintptr(h)) }

func (this *QOpenGLExtraFunctions) UnsafeRelease() {
	this.h = 0
	if this.QOpenGLFunctions != nil {
		this.QOpenGLFunctions.UnsafeRelease()
	}
}

func NewQOpenGLExtraFunctions() *QOpenGLExtraFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_new)
	__o := newownedQOpenGLExtraFunctions(__r)
	return __o
}

func NewQOpenGLExtraFunctions2(context *QOpenGLContext) *QOpenGLExtraFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_new2, uintptr(context.UnsafePointer()))
	__o := newownedQOpenGLExtraFunctions(__r)
	runtime.KeepAlive(context)
	return __o
}

func NewQOpenGLExtraFunctions3(arg1 *QOpenGLExtraFunctions) *QOpenGLExtraFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQOpenGLExtraFunctions(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QOpenGLExtraFunctions) GlReadBuffer(mode uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glReadBuffer, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlCopyTexSubImage3D(target uint, level int, xoffset int, yoffset int, zoffset int, x int, y int, width int, height int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glCopyTexSubImage3D, this.h, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGenQueries(n int) uint {
	var ids_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGenQueries, this.h, uintptr(n), uintptr(unsafe.Pointer(&ids_out)))
	runtime.KeepAlive(this)
	return uint(ids_out)
}

func (this *QOpenGLExtraFunctions) GlIsQuery(id uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glIsQuery, this.h, uintptr(id))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlBeginQuery(target uint, id uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBeginQuery, this.h, uintptr(target), uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlEndQuery(target uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glEndQuery, this.h, uintptr(target))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetQueryiv(target uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetQueryiv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetQueryObjectuiv(id uint, pname uint) uint {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetQueryObjectuiv, this.h, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return uint(params_out)
}

func (this *QOpenGLExtraFunctions) GlUnmapBuffer(target uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glUnmapBuffer, this.h, uintptr(target))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlBlitFramebuffer(srcX0 int, srcY0 int, srcX1 int, srcY1 int, dstX0 int, dstY0 int, dstX1 int, dstY1 int, mask uint, filter uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBlitFramebuffer, this.h, uintptr(srcX0), uintptr(srcY0), uintptr(srcX1), uintptr(srcY1), uintptr(dstX0), uintptr(dstY0), uintptr(dstX1), uintptr(dstY1), uintptr(mask), uintptr(filter))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlRenderbufferStorageMultisample(target uint, samples int, internalformat uint, width int, height int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glRenderbufferStorageMultisample, this.h, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlFramebufferTextureLayer(target uint, attachment uint, texture uint, level int, layer int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glFramebufferTextureLayer, this.h, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(layer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlFlushMappedBufferRange(target uint, offset int64, length int64) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glFlushMappedBufferRange, this.h, uintptr(target), uintptr(offset), uintptr(length))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBindVertexArray(array uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindVertexArray, this.h, uintptr(array))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGenVertexArrays(n int) uint {
	var arrays_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGenVertexArrays, this.h, uintptr(n), uintptr(unsafe.Pointer(&arrays_out)))
	runtime.KeepAlive(this)
	return uint(arrays_out)
}

func (this *QOpenGLExtraFunctions) GlIsVertexArray(array uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glIsVertexArray, this.h, uintptr(array))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlGetIntegeri_v(target uint, index uint) int {
	var data_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetIntegeri_v, this.h, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(&data_out)))
	runtime.KeepAlive(this)
	return int(data_out)
}

func (this *QOpenGLExtraFunctions) GlBeginTransformFeedback(primitiveMode uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBeginTransformFeedback, this.h, uintptr(primitiveMode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlEndTransformFeedback() {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glEndTransformFeedback, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBindBufferRange(target uint, index uint, buffer uint, offset int64, size int64) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindBufferRange, this.h, uintptr(target), uintptr(index), uintptr(buffer), uintptr(offset), uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBindBufferBase(target uint, index uint, buffer uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindBufferBase, this.h, uintptr(target), uintptr(index), uintptr(buffer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetTransformFeedbackVarying(program uint, index uint, bufSize int) (int, int, uint, int8) {
	var length_out int32
	var size_out int32
	var type__out int32
	var name_out int8
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetTransformFeedbackVarying, this.h, uintptr(program), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&size_out)), uintptr(unsafe.Pointer(&type__out)), uintptr(unsafe.Pointer(&name_out)))
	runtime.KeepAlive(this)
	return int(length_out), int(size_out), uint(type__out), name_out
}

func (this *QOpenGLExtraFunctions) GlGetVertexAttribIiv(index uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetVertexAttribIiv, this.h, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetVertexAttribIuiv(index uint, pname uint) uint {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetVertexAttribIuiv, this.h, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return uint(params_out)
}

func (this *QOpenGLExtraFunctions) GlVertexAttribI4i(index uint, x int, y int, z int, w int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexAttribI4i, this.h, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlVertexAttribI4ui(index uint, x uint, y uint, z uint, w uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexAttribI4ui, this.h, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetUniformuiv(program uint, location int) uint {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetUniformuiv, this.h, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return uint(params_out)
}

func (this *QOpenGLExtraFunctions) GlUniform1ui(location int, v0 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glUniform1ui, this.h, uintptr(location), uintptr(v0))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlUniform2ui(location int, v0 uint, v1 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glUniform2ui, this.h, uintptr(location), uintptr(v0), uintptr(v1))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlUniform3ui(location int, v0 uint, v1 uint, v2 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glUniform3ui, this.h, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlUniform4ui(location int, v0 uint, v1 uint, v2 uint, v3 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glUniform4ui, this.h, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlClearBufferfi(buffer uint, drawbuffer int, depth float32, stencil int) {
	pg_QOpenGLExtraFunctions_glClearBufferfi(this.h, int32(buffer), int32(drawbuffer), depth, int32(stencil))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlCopyBufferSubData(readTarget uint, writeTarget uint, readOffset int64, writeOffset int64, size int64) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glCopyBufferSubData, this.h, uintptr(readTarget), uintptr(writeTarget), uintptr(readOffset), uintptr(writeOffset), uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetActiveUniformBlockiv(program uint, uniformBlockIndex uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetActiveUniformBlockiv, this.h, uintptr(program), uintptr(uniformBlockIndex), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetActiveUniformBlockName(program uint, uniformBlockIndex uint, bufSize int) (int, int8) {
	var length_out int32
	var uniformBlockName_out int8
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetActiveUniformBlockName, this.h, uintptr(program), uintptr(uniformBlockIndex), uintptr(bufSize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&uniformBlockName_out)))
	runtime.KeepAlive(this)
	return int(length_out), uniformBlockName_out
}

func (this *QOpenGLExtraFunctions) GlUniformBlockBinding(program uint, uniformBlockIndex uint, uniformBlockBinding uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glUniformBlockBinding, this.h, uintptr(program), uintptr(uniformBlockIndex), uintptr(uniformBlockBinding))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlDrawArraysInstanced(mode uint, first int, count int, instancecount int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glDrawArraysInstanced, this.h, uintptr(mode), uintptr(first), uintptr(count), uintptr(instancecount))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetInteger64v(pname uint) int64 {
	var data_out int64
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetInteger64v, this.h, uintptr(pname), uintptr(unsafe.Pointer(&data_out)))
	runtime.KeepAlive(this)
	return data_out
}

func (this *QOpenGLExtraFunctions) GlGetInteger64i_v(target uint, index uint) int64 {
	var data_out int64
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetInteger64i_v, this.h, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(&data_out)))
	runtime.KeepAlive(this)
	return data_out
}

func (this *QOpenGLExtraFunctions) GlGetBufferParameteri64v(target uint, pname uint) int64 {
	var params_out int64
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetBufferParameteri64v, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLExtraFunctions) GlGenSamplers(count int) uint {
	var samplers_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGenSamplers, this.h, uintptr(count), uintptr(unsafe.Pointer(&samplers_out)))
	runtime.KeepAlive(this)
	return uint(samplers_out)
}

func (this *QOpenGLExtraFunctions) GlIsSampler(sampler uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glIsSampler, this.h, uintptr(sampler))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlBindSampler(unit uint, sampler uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindSampler, this.h, uintptr(unit), uintptr(sampler))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlSamplerParameteri(sampler uint, pname uint, param int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glSamplerParameteri, this.h, uintptr(sampler), uintptr(pname), uintptr(param))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlSamplerParameterf(sampler uint, pname uint, param float32) {
	pg_QOpenGLExtraFunctions_glSamplerParameterf(this.h, int32(sampler), int32(pname), param)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetSamplerParameteriv(sampler uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetSamplerParameteriv, this.h, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetSamplerParameterfv(sampler uint, pname uint) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetSamplerParameterfv, this.h, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLExtraFunctions) GlVertexAttribDivisor(index uint, divisor uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexAttribDivisor, this.h, uintptr(index), uintptr(divisor))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBindTransformFeedback(target uint, id uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindTransformFeedback, this.h, uintptr(target), uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGenTransformFeedbacks(n int) uint {
	var ids_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGenTransformFeedbacks, this.h, uintptr(n), uintptr(unsafe.Pointer(&ids_out)))
	runtime.KeepAlive(this)
	return uint(ids_out)
}

func (this *QOpenGLExtraFunctions) GlIsTransformFeedback(id uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glIsTransformFeedback, this.h, uintptr(id))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlPauseTransformFeedback() {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glPauseTransformFeedback, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlResumeTransformFeedback() {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glResumeTransformFeedback, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramParameteri(program uint, pname uint, value int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramParameteri, this.h, uintptr(program), uintptr(pname), uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlTexStorage2D(target uint, levels int, internalformat uint, width int, height int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glTexStorage2D, this.h, uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlTexStorage3D(target uint, levels int, internalformat uint, width int, height int, depth int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glTexStorage3D, this.h, uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetInternalformativ(target uint, internalformat uint, pname uint, bufSize int) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetInternalformativ, this.h, uintptr(target), uintptr(internalformat), uintptr(pname), uintptr(bufSize), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlDispatchCompute(num_groups_x uint, num_groups_y uint, num_groups_z uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glDispatchCompute, this.h, uintptr(num_groups_x), uintptr(num_groups_y), uintptr(num_groups_z))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlDispatchComputeIndirect(indirect int64) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glDispatchComputeIndirect, this.h, uintptr(indirect))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlFramebufferParameteri(target uint, pname uint, param int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glFramebufferParameteri, this.h, uintptr(target), uintptr(pname), uintptr(param))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetFramebufferParameteriv(target uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetFramebufferParameteriv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetProgramInterfaceiv(program uint, programInterface uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetProgramInterfaceiv, this.h, uintptr(program), uintptr(programInterface), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetProgramResourceName(program uint, programInterface uint, index uint, bufSize int) (int, int8) {
	var length_out int32
	var name_out int8
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetProgramResourceName, this.h, uintptr(program), uintptr(programInterface), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&name_out)))
	runtime.KeepAlive(this)
	return int(length_out), name_out
}

func (this *QOpenGLExtraFunctions) GlUseProgramStages(pipeline uint, stages uint, program uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glUseProgramStages, this.h, uintptr(pipeline), uintptr(stages), uintptr(program))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlActiveShaderProgram(pipeline uint, program uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glActiveShaderProgram, this.h, uintptr(pipeline), uintptr(program))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBindProgramPipeline(pipeline uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindProgramPipeline, this.h, uintptr(pipeline))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGenProgramPipelines(n int) uint {
	var pipelines_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGenProgramPipelines, this.h, uintptr(n), uintptr(unsafe.Pointer(&pipelines_out)))
	runtime.KeepAlive(this)
	return uint(pipelines_out)
}

func (this *QOpenGLExtraFunctions) GlIsProgramPipeline(pipeline uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glIsProgramPipeline, this.h, uintptr(pipeline))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlGetProgramPipelineiv(pipeline uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetProgramPipelineiv, this.h, uintptr(pipeline), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform1i(program uint, location int, v0 int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform1i, this.h, uintptr(program), uintptr(location), uintptr(v0))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform2i(program uint, location int, v0 int, v1 int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform2i, this.h, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform3i(program uint, location int, v0 int, v1 int, v2 int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform3i, this.h, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform4i(program uint, location int, v0 int, v1 int, v2 int, v3 int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform4i, this.h, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform1ui(program uint, location int, v0 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform1ui, this.h, uintptr(program), uintptr(location), uintptr(v0))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform2ui(program uint, location int, v0 uint, v1 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform2ui, this.h, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform3ui(program uint, location int, v0 uint, v1 uint, v2 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform3ui, this.h, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform4ui(program uint, location int, v0 uint, v1 uint, v2 uint, v3 uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glProgramUniform4ui, this.h, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform1f(program uint, location int, v0 float32) {
	pg_QOpenGLExtraFunctions_glProgramUniform1f(this.h, int32(program), int32(location), v0)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform2f(program uint, location int, v0 float32, v1 float32) {
	pg_QOpenGLExtraFunctions_glProgramUniform2f(this.h, int32(program), int32(location), v0, v1)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform3f(program uint, location int, v0 float32, v1 float32, v2 float32) {
	pg_QOpenGLExtraFunctions_glProgramUniform3f(this.h, int32(program), int32(location), v0, v1, v2)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlProgramUniform4f(program uint, location int, v0 float32, v1 float32, v2 float32, v3 float32) {
	pg_QOpenGLExtraFunctions_glProgramUniform4f(this.h, int32(program), int32(location), v0, v1, v2, v3)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlValidateProgramPipeline(pipeline uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glValidateProgramPipeline, this.h, uintptr(pipeline))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetProgramPipelineInfoLog(pipeline uint, bufSize int) (int, int8) {
	var length_out int32
	var infoLog_out int8
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetProgramPipelineInfoLog, this.h, uintptr(pipeline), uintptr(bufSize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&infoLog_out)))
	runtime.KeepAlive(this)
	return int(length_out), infoLog_out
}

func (this *QOpenGLExtraFunctions) GlBindImageTexture(unit uint, texture uint, level int, layered byte, layer int, access uint, format uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindImageTexture, this.h, uintptr(unit), uintptr(texture), uintptr(level), uintptr(layered), uintptr(layer), uintptr(access), uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetBooleani_v(target uint, index uint) byte {
	var data_out uint8
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetBooleani_v, this.h, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(&data_out)))
	runtime.KeepAlive(this)
	return byte(data_out)
}

func (this *QOpenGLExtraFunctions) GlMemoryBarrier(barriers uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glMemoryBarrier, this.h, uintptr(barriers))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlMemoryBarrierByRegion(barriers uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glMemoryBarrierByRegion, this.h, uintptr(barriers))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlTexStorage2DMultisample(target uint, samples int, internalformat uint, width int, height int, fixedsamplelocations byte) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glTexStorage2DMultisample, this.h, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(fixedsamplelocations))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetMultisamplefv(pname uint, index uint) float32 {
	var val_out float32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetMultisamplefv, this.h, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(&val_out)))
	runtime.KeepAlive(this)
	return val_out
}

func (this *QOpenGLExtraFunctions) GlSampleMaski(maskNumber uint, mask uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glSampleMaski, this.h, uintptr(maskNumber), uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetTexLevelParameteriv(target uint, level int, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetTexLevelParameteriv, this.h, uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetTexLevelParameterfv(target uint, level int, pname uint) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetTexLevelParameterfv, this.h, uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLExtraFunctions) GlBindVertexBuffer(bindingindex uint, buffer uint, offset int64, stride int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBindVertexBuffer, this.h, uintptr(bindingindex), uintptr(buffer), uintptr(offset), uintptr(stride))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlVertexAttribFormat(attribindex uint, size int, type_ uint, normalized byte, relativeoffset uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexAttribFormat, this.h, uintptr(attribindex), uintptr(size), uintptr(type_), uintptr(normalized), uintptr(relativeoffset))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlVertexAttribIFormat(attribindex uint, size int, type_ uint, relativeoffset uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexAttribIFormat, this.h, uintptr(attribindex), uintptr(size), uintptr(type_), uintptr(relativeoffset))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlVertexAttribBinding(attribindex uint, bindingindex uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexAttribBinding, this.h, uintptr(attribindex), uintptr(bindingindex))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlVertexBindingDivisor(bindingindex uint, divisor uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glVertexBindingDivisor, this.h, uintptr(bindingindex), uintptr(divisor))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBlendBarrier() {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBlendBarrier, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlCopyImageSubData(srcName uint, srcTarget uint, srcLevel int, srcX int, srcY int, srcZ int, dstName uint, dstTarget uint, dstLevel int, dstX int, dstY int, dstZ int, srcWidth int, srcHeight int, srcDepth int) {
	__pack := []uintptr{uintptr(srcName), uintptr(srcTarget), uintptr(srcLevel), uintptr(srcX), uintptr(srcY), uintptr(srcZ), uintptr(dstName), uintptr(dstTarget), uintptr(dstLevel), uintptr(dstX), uintptr(dstY), uintptr(dstZ), uintptr(srcWidth), uintptr(srcHeight), uintptr(srcDepth)}
	purego.SyscallN(addr_QOpenGLExtraFunctions_glCopyImageSubData, this.h, uintptr(unsafe.Pointer(&__pack[0])))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetDebugMessageLog(count uint, bufSize int) (uint, uint, uint, uint, uint, int, int8) {
	var sources_out int32
	var types_out int32
	var ids_out int32
	var severities_out int32
	var lengths_out int32
	var messageLog_out int8
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glGetDebugMessageLog, this.h, uintptr(count), uintptr(bufSize), uintptr(unsafe.Pointer(&sources_out)), uintptr(unsafe.Pointer(&types_out)), uintptr(unsafe.Pointer(&ids_out)), uintptr(unsafe.Pointer(&severities_out)), uintptr(unsafe.Pointer(&lengths_out)), uintptr(unsafe.Pointer(&messageLog_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), uint(sources_out), uint(types_out), uint(ids_out), uint(severities_out), int(lengths_out), messageLog_out
}

func (this *QOpenGLExtraFunctions) GlPopDebugGroup() {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glPopDebugGroup, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetObjectLabel(identifier uint, name uint, bufSize int) (int, int8) {
	var length_out int32
	var label_out int8
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetObjectLabel, this.h, uintptr(identifier), uintptr(name), uintptr(bufSize), uintptr(unsafe.Pointer(&length_out)), uintptr(unsafe.Pointer(&label_out)))
	runtime.KeepAlive(this)
	return int(length_out), label_out
}

func (this *QOpenGLExtraFunctions) GlEnablei(target uint, index uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glEnablei, this.h, uintptr(target), uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlDisablei(target uint, index uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glDisablei, this.h, uintptr(target), uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBlendEquationi(buf uint, mode uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBlendEquationi, this.h, uintptr(buf), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBlendEquationSeparatei(buf uint, modeRGB uint, modeAlpha uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBlendEquationSeparatei, this.h, uintptr(buf), uintptr(modeRGB), uintptr(modeAlpha))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBlendFunci(buf uint, src uint, dst uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBlendFunci, this.h, uintptr(buf), uintptr(src), uintptr(dst))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlBlendFuncSeparatei(buf uint, srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glBlendFuncSeparatei, this.h, uintptr(buf), uintptr(srcRGB), uintptr(dstRGB), uintptr(srcAlpha), uintptr(dstAlpha))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlColorMaski(index uint, r byte, g byte, b byte, a byte) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glColorMaski, this.h, uintptr(index), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlIsEnabledi(target uint, index uint) byte {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glIsEnabledi, this.h, uintptr(target), uintptr(index))
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QOpenGLExtraFunctions) GlFramebufferTexture(target uint, attachment uint, texture uint, level int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glFramebufferTexture, this.h, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlPrimitiveBoundingBox(minX float32, minY float32, minZ float32, minW float32, maxX float32, maxY float32, maxZ float32, maxW float32) {
	pg_QOpenGLExtraFunctions_glPrimitiveBoundingBox(this.h, minX, minY, minZ, minW, maxX, maxY, maxZ, maxW)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetGraphicsResetStatus() uint {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLExtraFunctions_glGetGraphicsResetStatus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QOpenGLExtraFunctions) GlGetnUniformfv(program uint, location int, bufSize int) float32 {
	var params_out float32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetnUniformfv, this.h, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return params_out
}

func (this *QOpenGLExtraFunctions) GlGetnUniformiv(program uint, location int, bufSize int) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetnUniformiv, this.h, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetnUniformuiv(program uint, location int, bufSize int) uint {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetnUniformuiv, this.h, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return uint(params_out)
}

func (this *QOpenGLExtraFunctions) GlMinSampleShading(value float32) {
	pg_QOpenGLExtraFunctions_glMinSampleShading(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlPatchParameteri(pname uint, value int) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glPatchParameteri, this.h, uintptr(pname), uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlGetTexParameterIiv(target uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetTexParameterIiv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetTexParameterIuiv(target uint, pname uint) uint {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetTexParameterIuiv, this.h, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return uint(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetSamplerParameterIiv(sampler uint, pname uint) int {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetSamplerParameterIiv, this.h, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return int(params_out)
}

func (this *QOpenGLExtraFunctions) GlGetSamplerParameterIuiv(sampler uint, pname uint) uint {
	var params_out int32
	purego.SyscallN(addr_QOpenGLExtraFunctions_glGetSamplerParameterIuiv, this.h, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(&params_out)))
	runtime.KeepAlive(this)
	return uint(params_out)
}

func (this *QOpenGLExtraFunctions) GlTexBuffer(target uint, internalformat uint, buffer uint) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glTexBuffer, this.h, uintptr(target), uintptr(internalformat), uintptr(buffer))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlTexBufferRange(target uint, internalformat uint, buffer uint, offset int64, size int64) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glTexBufferRange, this.h, uintptr(target), uintptr(internalformat), uintptr(buffer), uintptr(offset), uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) GlTexStorage3DMultisample(target uint, samples int, internalformat uint, width int, height int, depth int, fixedsamplelocations byte) {
	purego.SyscallN(addr_QOpenGLExtraFunctions_glTexStorage3DMultisample, this.h, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(fixedsamplelocations))
	runtime.KeepAlive(this)
}

func (this *QOpenGLExtraFunctions) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOpenGLExtraFunctions_delete(this.h)
	}
	this.UnsafeRelease()
}
