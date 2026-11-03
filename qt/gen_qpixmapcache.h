#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QPixmapCache_cacheLimit();
void QPixmapCache_setCacheLimit(int arg1);
int QPixmapCache_find(const char* key, long long key_len, void* pixmap);
int QPixmapCache_find2(void* key, void* pixmap);
int QPixmapCache_insert(const char* key, long long key_len, void* pixmap);
void* QPixmapCache_insert2(void* pixmap);
int QPixmapCache_replace(void* key, void* pixmap);
void QPixmapCache_remove(const char* key, long long key_len);
void QPixmapCache_remove2(void* key);
void QPixmapCache_clear();
void QPixmapCache_delete(void* self);

#ifdef __cplusplus
}
#endif
