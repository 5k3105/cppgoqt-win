#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPluginMetaData__ElfNoteHeader_new(unsigned int payloadSize);
void* QPluginMetaData__ElfNoteHeader_new2(void* arg1);
unsigned int QPluginMetaData__ElfNoteHeader_n_namesz_get(void* self);
void QPluginMetaData__ElfNoteHeader_n_namesz_set(void* self, unsigned int n_namesz);
unsigned int QPluginMetaData__ElfNoteHeader_n_descsz_get(void* self);
void QPluginMetaData__ElfNoteHeader_n_descsz_set(void* self, unsigned int n_descsz);
unsigned int QPluginMetaData__ElfNoteHeader_n_type_get(void* self);
void QPluginMetaData__ElfNoteHeader_n_type_set(void* self, unsigned int n_type);
void* QPluginMetaData__ElfNoteHeader_header_get(void* self);
void QPluginMetaData__ElfNoteHeader_header_set(void* self, void* header);
void QPluginMetaData__ElfNoteHeader_delete(void* self);

#ifdef __cplusplus
}
#endif
