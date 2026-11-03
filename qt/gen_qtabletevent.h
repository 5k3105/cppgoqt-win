#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTabletEvent_base(void* self);
void QTabletEvent_chain(void* self, void** out);
void* QTabletEvent_new(int t, void* device, void* pos, void* globalPos, double pressure, float xTilt, float yTilt, float tangentialPressure, double rotation, float z, int keyState, int button, int buttons);
void* QTabletEvent_pos(void* self);
void* QTabletEvent_globalPos(void* self);
void* QTabletEvent_posF(void* self);
void* QTabletEvent_globalPosF(void* self);
int QTabletEvent_x(void* self);
int QTabletEvent_y(void* self);
int QTabletEvent_globalX(void* self);
int QTabletEvent_globalY(void* self);
double QTabletEvent_hiResGlobalX(void* self);
double QTabletEvent_hiResGlobalY(void* self);
long long QTabletEvent_uniqueId(void* self);
double QTabletEvent_pressure(void* self);
double QTabletEvent_rotation(void* self);
double QTabletEvent_z(void* self);
double QTabletEvent_tangentialPressure(void* self);
double QTabletEvent_xTilt(void* self);
double QTabletEvent_yTilt(void* self);
int QTabletEvent_registerEventType();
int QTabletEvent_registerEventType2(int hint);
void QTabletEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
