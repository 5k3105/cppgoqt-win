#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractTextDocumentLayout_base(void* self);
void* QAbstractTextDocumentLayout_fromQObject(void* o);
char* QAbstractTextDocumentLayout_tr(const char* s, long long* __out_len);
void QAbstractTextDocumentLayout_draw(void* self, void* painter, void* context);
int QAbstractTextDocumentLayout_hitTest(void* self, void* point, int accuracy);
char* QAbstractTextDocumentLayout_anchorAt(void* self, void* pos, long long* __out_len);
char* QAbstractTextDocumentLayout_imageAt(void* self, void* pos, long long* __out_len);
void* QAbstractTextDocumentLayout_formatAt(void* self, void* pos);
void* QAbstractTextDocumentLayout_blockWithMarkerAt(void* self, void* pos);
int QAbstractTextDocumentLayout_pageCount(void* self);
void* QAbstractTextDocumentLayout_documentSize(void* self);
void* QAbstractTextDocumentLayout_frameBoundingRect(void* self, void* frame);
void* QAbstractTextDocumentLayout_blockBoundingRect(void* self, void* block);
void QAbstractTextDocumentLayout_setPaintDevice(void* self, void* device);
void* QAbstractTextDocumentLayout_paintDevice(void* self);
void* QAbstractTextDocumentLayout_document(void* self);
void QAbstractTextDocumentLayout_registerHandler(void* self, int objectType, void* component);
void QAbstractTextDocumentLayout_unregisterHandler(void* self, int objectType);
void* QAbstractTextDocumentLayout_handlerForObject(void* self, int objectType);
void QAbstractTextDocumentLayout_update(void* self);
void QAbstractTextDocumentLayout_updateBlock(void* self, void* block);
void QAbstractTextDocumentLayout_documentSizeChanged(void* self, void* newSize);
void QAbstractTextDocumentLayout_pageCountChanged(void* self, int newPages);
void* QAbstractTextDocumentLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractTextDocumentLayout_connect2(void* sender, void* signal, void* receiver, void* method);
char* QAbstractTextDocumentLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractTextDocumentLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void QAbstractTextDocumentLayout_unregisterHandler2(void* self, int objectType, void* component);
void QAbstractTextDocumentLayout_update2(void* self, void* arg1);
void* QAbstractTextDocumentLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractTextDocumentLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QAbstractTextDocumentLayout_connect_update(void* self, long long id);
void QAbstractTextDocumentLayout_connect_update2(void* self, long long id);
void QAbstractTextDocumentLayout_connect_updateBlock(void* self, long long id);
void QAbstractTextDocumentLayout_connect_documentSizeChanged(void* self, long long id);
void QAbstractTextDocumentLayout_connect_pageCountChanged(void* self, long long id);
void QAbstractTextDocumentLayout_connect_destroyed(void* self, long long id);
void QAbstractTextDocumentLayout_connect_destroyed2(void* self, long long id);
void QAbstractTextDocumentLayout_connect_objectNameChanged(void* self, long long id);
void QAbstractTextDocumentLayout__disconnect(long long id);
void QAbstractTextDocumentLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
