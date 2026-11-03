#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPdfWriter_base(void* self);
void* QPdfWriter_fromQObject(void* o);
void* QPdfWriter_new(const char* filename, long long filename_len);
void* QPdfWriter_new2(void* device);
char* QPdfWriter_tr(const char* s, long long* __out_len);
void QPdfWriter_setPdfVersion(void* self, int version);
int QPdfWriter_pdfVersion(void* self);
char* QPdfWriter_title(void* self, long long* __out_len);
void QPdfWriter_setTitle(void* self, const char* title, long long title_len);
char* QPdfWriter_creator(void* self, long long* __out_len);
void QPdfWriter_setCreator(void* self, const char* creator, long long creator_len);
void* QPdfWriter_documentId(void* self);
void QPdfWriter_setDocumentId(void* self, void* documentId);
char* QPdfWriter_author(void* self, long long* __out_len);
void QPdfWriter_setAuthor(void* self, const char* author, long long author_len);
int QPdfWriter_newPage(void* self);
void QPdfWriter_setResolution(void* self, int resolution);
int QPdfWriter_resolution(void* self);
void QPdfWriter_setDocumentXmpMetadata(void* self, void* xmpMetadata);
void* QPdfWriter_documentXmpMetadata(void* self);
void QPdfWriter_addFileAttachment(void* self, const char* fileName, long long fileName_len, void* data);
int QPdfWriter_colorModel(void* self);
void QPdfWriter_setColorModel(void* self, int model);
void* QPdfWriter_outputIntent(void* self);
void QPdfWriter_setOutputIntent(void* self, void* intent);
void* QPdfWriter_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPdfWriter_connect2(void* sender, void* signal, void* receiver, void* method);
int QPdfWriter_setPageLayout(void* self, void* pageLayout);
int QPdfWriter_setPageSize(void* self, void* pageSize);
int QPdfWriter_setPageOrientation(void* self, int orientation);
int QPdfWriter_setPageMargins(void* self, void* margins, int units);
void* QPdfWriter_pageLayout(void* self);
void QPdfWriter_setPageRanges(void* self, void* ranges);
void* QPdfWriter_pageRanges(void* self);
int QPdfWriter_devType(void* self);
int QPdfWriter_paintingActive(void* self);
int QPdfWriter_width(void* self);
int QPdfWriter_height(void* self);
int QPdfWriter_widthMM(void* self);
int QPdfWriter_heightMM(void* self);
int QPdfWriter_logicalDpiX(void* self);
int QPdfWriter_logicalDpiY(void* self);
int QPdfWriter_physicalDpiX(void* self);
int QPdfWriter_physicalDpiY(void* self);
double QPdfWriter_devicePixelRatio(void* self);
double QPdfWriter_devicePixelRatioF(void* self);
int QPdfWriter_colorCount(void* self);
int QPdfWriter_depth(void* self);
double QPdfWriter_devicePixelRatioFScale();
int QPdfWriter_encodeMetricF(int metric, double value);
char* QPdfWriter_tr2(const char* s, const char* c, long long* __out_len);
char* QPdfWriter_tr3(const char* s, const char* c, int n, long long* __out_len);
void QPdfWriter_addFileAttachment2(void* self, const char* fileName, long long fileName_len, void* data, const char* mimeType, long long mimeType_len);
void* QPdfWriter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPdfWriter_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPdfWriter_asQPagedPaintDevice(void* self);
void QPdfWriter_connect_destroyed(void* self, long long id);
void QPdfWriter_connect_destroyed2(void* self, long long id);
void QPdfWriter_connect_objectNameChanged(void* self, long long id);
void QPdfWriter__disconnect(long long id);
void QPdfWriter_delete(void* self);

#ifdef __cplusplus
}
#endif
