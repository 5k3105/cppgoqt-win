#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMessageLogger_new();
void* QMessageLogger_new2(const char* file, int line, const char* function);
void* QMessageLogger_new3(const char* file, int line, const char* function, const char* category);
void QMessageLogger_debug(void* self, const char* msg);
void QMessageLogger_noDebug(void* self, const char* arg1);
void QMessageLogger_info(void* self, const char* msg);
void QMessageLogger_warning(void* self, const char* msg);
void QMessageLogger_critical(void* self, const char* msg);
void QMessageLogger_fatal(void* self, const char* msg);
void QMessageLogger_debug2(void* self, void* cat, const char* msg);
void QMessageLogger_info2(void* self, void* cat, const char* msg);
void QMessageLogger_warning2(void* self, void* cat, const char* msg);
void QMessageLogger_critical2(void* self, void* cat, const char* msg);
void QMessageLogger_fatal2(void* self, void* cat, const char* msg);
void* QMessageLogger_debug3(void* self);
void* QMessageLogger_debug4(void* self, void* cat);
void* QMessageLogger_info3(void* self);
void* QMessageLogger_info4(void* self, void* cat);
void* QMessageLogger_warning3(void* self);
void* QMessageLogger_warning4(void* self, void* cat);
void* QMessageLogger_critical3(void* self);
void* QMessageLogger_critical4(void* self, void* cat);
void* QMessageLogger_fatal3(void* self);
void* QMessageLogger_fatal4(void* self, void* cat);
void* QMessageLogger_noDebug2(void* self);
void QMessageLogger_delete(void* self);

#ifdef __cplusplus
}
#endif
