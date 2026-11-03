#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QEventPoint_new();
void* QEventPoint_new2(int pointId, int state, void* scenePosition, void* globalPosition);
void* QEventPoint_new3(void* other);
void* QEventPoint_new4(int id);
void* QEventPoint_new5(int id, void* device);
void* QEventPoint_OperatorAssign(void* self, void* other);
int QEventPoint_OperatorEqual(void* self, void* other);
int QEventPoint_OperatorNotEqual(void* self, void* other);
void QEventPoint_swap(void* self, void* other);
void* QEventPoint_position(void* self);
void* QEventPoint_pressPosition(void* self);
void* QEventPoint_grabPosition(void* self);
void* QEventPoint_lastPosition(void* self);
void* QEventPoint_scenePosition(void* self);
void* QEventPoint_scenePressPosition(void* self);
void* QEventPoint_sceneGrabPosition(void* self);
void* QEventPoint_sceneLastPosition(void* self);
void* QEventPoint_globalPosition(void* self);
void* QEventPoint_globalPressPosition(void* self);
void* QEventPoint_globalGrabPosition(void* self);
void* QEventPoint_globalLastPosition(void* self);
void* QEventPoint_normalizedPosition(void* self);
void* QEventPoint_pos(void* self);
void* QEventPoint_startPos(void* self);
void* QEventPoint_scenePos(void* self);
void* QEventPoint_startScenePos(void* self);
void* QEventPoint_screenPos(void* self);
void* QEventPoint_startScreenPos(void* self);
void* QEventPoint_startNormalizedPos(void* self);
void* QEventPoint_normalizedPos(void* self);
void* QEventPoint_lastPos(void* self);
void* QEventPoint_lastScenePos(void* self);
void* QEventPoint_lastScreenPos(void* self);
void* QEventPoint_lastNormalizedPos(void* self);
void* QEventPoint_velocity(void* self);
int QEventPoint_state(void* self);
void* QEventPoint_device(void* self);
int QEventPoint_id(void* self);
void* QEventPoint_uniqueId(void* self);
unsigned long QEventPoint_timestamp(void* self);
unsigned long QEventPoint_lastTimestamp(void* self);
unsigned long QEventPoint_pressTimestamp(void* self);
double QEventPoint_timeHeld(void* self);
double QEventPoint_pressure(void* self);
double QEventPoint_rotation(void* self);
void* QEventPoint_ellipseDiameters(void* self);
int QEventPoint_isAccepted(void* self);
void QEventPoint_setAccepted(void* self);
void QEventPoint_setAccepted2(void* self, int accepted);
void QEventPoint_delete(void* self);

#ifdef __cplusplus
}
#endif
