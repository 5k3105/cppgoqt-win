#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleHints_base(void* self);
void* QStyleHints_fromQObject(void* o);
char* QStyleHints_tr(const char* s, long long* __out_len);
void QStyleHints_setMouseDoubleClickInterval(void* self, int mouseDoubleClickInterval);
int QStyleHints_mouseDoubleClickInterval(void* self);
int QStyleHints_mouseDoubleClickDistance(void* self);
int QStyleHints_touchDoubleTapDistance(void* self);
void QStyleHints_setMousePressAndHoldInterval(void* self, int mousePressAndHoldInterval);
int QStyleHints_mousePressAndHoldInterval(void* self);
void QStyleHints_setStartDragDistance(void* self, int startDragDistance);
int QStyleHints_startDragDistance(void* self);
void QStyleHints_setStartDragTime(void* self, int startDragTime);
int QStyleHints_startDragTime(void* self);
int QStyleHints_startDragVelocity(void* self);
void QStyleHints_setKeyboardInputInterval(void* self, int keyboardInputInterval);
int QStyleHints_keyboardInputInterval(void* self);
int QStyleHints_keyboardAutoRepeatRate(void* self);
double QStyleHints_keyboardAutoRepeatRateF(void* self);
void QStyleHints_setCursorFlashTime(void* self, int cursorFlashTime);
int QStyleHints_cursorFlashTime(void* self);
int QStyleHints_showIsFullScreen(void* self);
int QStyleHints_showIsMaximized(void* self);
int QStyleHints_showShortcutsInContextMenus(void* self);
void QStyleHints_setShowShortcutsInContextMenus(void* self, int showShortcutsInContextMenus);
int QStyleHints_contextMenuTrigger(void* self);
void QStyleHints_setContextMenuTrigger(void* self, int contextMenuTrigger);
int QStyleHints_menuSelectionWraps(void* self);
int QStyleHints_passwordMaskDelay(void* self);
void* QStyleHints_passwordMaskCharacter(void* self);
double QStyleHints_fontSmoothingGamma(void* self);
int QStyleHints_useRtlExtensions(void* self);
int QStyleHints_setFocusOnTouchRelease(void* self);
int QStyleHints_tabFocusBehavior(void* self);
void QStyleHints_setTabFocusBehavior(void* self, int tabFocusBehavior);
int QStyleHints_singleClickActivation(void* self);
int QStyleHints_useHoverEffects(void* self);
void QStyleHints_setUseHoverEffects(void* self, int useHoverEffects);
int QStyleHints_wheelScrollLines(void* self);
void QStyleHints_setWheelScrollLines(void* self, int scrollLines);
void QStyleHints_setMouseQuickSelectionThreshold(void* self, int threshold);
int QStyleHints_mouseQuickSelectionThreshold(void* self);
int QStyleHints_colorScheme(void* self);
void QStyleHints_setColorScheme(void* self, int scheme);
void QStyleHints_unsetColorScheme(void* self);
void* QStyleHints_accessibility(void* self);
void QStyleHints_cursorFlashTimeChanged(void* self, int cursorFlashTime);
void QStyleHints_keyboardInputIntervalChanged(void* self, int keyboardInputInterval);
void QStyleHints_mouseDoubleClickIntervalChanged(void* self, int mouseDoubleClickInterval);
void QStyleHints_mousePressAndHoldIntervalChanged(void* self, int mousePressAndHoldInterval);
void QStyleHints_startDragDistanceChanged(void* self, int startDragDistance);
void QStyleHints_startDragTimeChanged(void* self, int startDragTime);
void QStyleHints_tabFocusBehaviorChanged(void* self, int tabFocusBehavior);
void QStyleHints_useHoverEffectsChanged(void* self, int useHoverEffects);
void QStyleHints_showShortcutsInContextMenusChanged(void* self, int arg1);
void QStyleHints_contextMenuTriggerChanged(void* self, int contextMenuTrigger);
void QStyleHints_wheelScrollLinesChanged(void* self, int scrollLines);
void QStyleHints_mouseQuickSelectionThresholdChanged(void* self, int threshold);
void QStyleHints_colorSchemeChanged(void* self, int colorScheme);
void* QStyleHints_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QStyleHints_connect2(void* sender, void* signal, void* receiver, void* method);
char* QStyleHints_tr2(const char* s, const char* c, long long* __out_len);
char* QStyleHints_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QStyleHints_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QStyleHints_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QStyleHints_connect_cursorFlashTimeChanged(void* self, long long id);
void QStyleHints_connect_keyboardInputIntervalChanged(void* self, long long id);
void QStyleHints_connect_mouseDoubleClickIntervalChanged(void* self, long long id);
void QStyleHints_connect_mousePressAndHoldIntervalChanged(void* self, long long id);
void QStyleHints_connect_startDragDistanceChanged(void* self, long long id);
void QStyleHints_connect_startDragTimeChanged(void* self, long long id);
void QStyleHints_connect_tabFocusBehaviorChanged(void* self, long long id);
void QStyleHints_connect_useHoverEffectsChanged(void* self, long long id);
void QStyleHints_connect_showShortcutsInContextMenusChanged(void* self, long long id);
void QStyleHints_connect_contextMenuTriggerChanged(void* self, long long id);
void QStyleHints_connect_wheelScrollLinesChanged(void* self, long long id);
void QStyleHints_connect_mouseQuickSelectionThresholdChanged(void* self, long long id);
void QStyleHints_connect_colorSchemeChanged(void* self, long long id);
void QStyleHints_connect_destroyed(void* self, long long id);
void QStyleHints_connect_destroyed2(void* self, long long id);
void QStyleHints_connect_objectNameChanged(void* self, long long id);
void QStyleHints__disconnect(long long id);
void QStyleHints_delete(void* self);

#ifdef __cplusplus
}
#endif
