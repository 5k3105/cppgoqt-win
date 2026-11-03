#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPalette_new();
void* QPalette_new2(void* button);
void* QPalette_new3(int button);
void* QPalette_new4(void* button, void* window);
void* QPalette_new5(void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window);
void* QPalette_new6(void* windowText, void* window, void* light, void* dark, void* mid, void* text, void* base);
void* QPalette_new7(void* palette);
void* QPalette_OperatorAssign(void* self, void* palette);
void QPalette_swap(void* self, void* other);
void* QPalette_ToQVariant(void* self);
int QPalette_currentColorGroup(void* self);
void QPalette_setCurrentColorGroup(void* self, int cg);
void* QPalette_color(void* self, int cg, int cr);
void* QPalette_brush(void* self, int cg, int cr);
void QPalette_setColor(void* self, int cg, int cr, void* color);
void QPalette_setColor2(void* self, int cr, void* color);
void QPalette_setBrush(void* self, int cr, void* brush);
int QPalette_isBrushSet(void* self, int cg, int cr);
void QPalette_setBrush2(void* self, int cg, int cr, void* brush);
void QPalette_setColorGroup(void* self, int cr, void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window);
int QPalette_isEqual(void* self, int cr1, int cr2);
void* QPalette_color2(void* self, int cr);
void* QPalette_brush2(void* self, int cr);
void* QPalette_windowText(void* self);
void* QPalette_button(void* self);
void* QPalette_light(void* self);
void* QPalette_dark(void* self);
void* QPalette_mid(void* self);
void* QPalette_text(void* self);
void* QPalette_base(void* self);
void* QPalette_alternateBase(void* self);
void* QPalette_toolTipBase(void* self);
void* QPalette_toolTipText(void* self);
void* QPalette_window(void* self);
void* QPalette_midlight(void* self);
void* QPalette_brightText(void* self);
void* QPalette_buttonText(void* self);
void* QPalette_shadow(void* self);
void* QPalette_highlight(void* self);
void* QPalette_highlightedText(void* self);
void* QPalette_link(void* self);
void* QPalette_linkVisited(void* self);
void* QPalette_placeholderText(void* self);
void* QPalette_accent(void* self);
int QPalette_OperatorEqual(void* self, void* p);
int QPalette_OperatorNotEqual(void* self, void* p);
int QPalette_isCopyOf(void* self, void* p);
long long QPalette_cacheKey(void* self);
void* QPalette_resolve(void* self, void* other);
unsigned long long QPalette_resolveMask(void* self);
void QPalette_setResolveMask(void* self, unsigned long long mask);
void* QPalette_toQVariant(void* self);
void QPalette_delete(void* self);

#ifdef __cplusplus
}
#endif
