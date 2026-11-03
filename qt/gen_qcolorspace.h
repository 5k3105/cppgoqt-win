#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QColorSpace_new();
void* QColorSpace_new2(int namedColorSpace);
void* QColorSpace_new3(void* whitePoint, int transferFunction);
void* QColorSpace_new4(void* whitePoint, unsigned short* transferFunctionTable, long long transferFunctionTable_len);
void* QColorSpace_new5(int primaries, int transferFunction);
void* QColorSpace_new6(int primaries, float gamma);
void* QColorSpace_new7(int primaries, unsigned short* transferFunctionTable, long long transferFunctionTable_len);
void* QColorSpace_new8(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, int transferFunction);
void* QColorSpace_new9(void* primaryPoints, int transferFunction);
void* QColorSpace_new10(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, unsigned short* transferFunctionTable, long long transferFunctionTable_len);
void* QColorSpace_new11(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, unsigned short* redTransferFunctionTable, long long redTransferFunctionTable_len, unsigned short* greenTransferFunctionTable, long long greenTransferFunctionTable_len, unsigned short* blueTransferFunctionTable, long long blueTransferFunctionTable_len);
void* QColorSpace_new12(void* colorSpace);
void* QColorSpace_new13(void* whitePoint, int transferFunction, float gamma);
void* QColorSpace_new14(int primaries, int transferFunction, float gamma);
void* QColorSpace_new15(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, int transferFunction, float gamma);
void* QColorSpace_new16(void* primaryPoints, int transferFunction, float gamma);
void* QColorSpace_OperatorAssign(void* self, void* colorSpace);
void QColorSpace_swap(void* self, void* colorSpace);
int QColorSpace_primaries(void* self);
int QColorSpace_transferFunction(void* self);
float QColorSpace_gamma(void* self);
char* QColorSpace_description(void* self, long long* __out_len);
void QColorSpace_setDescription(void* self, const char* description, long long description_len);
void QColorSpace_setTransferFunction(void* self, int transferFunction);
void QColorSpace_setTransferFunction2(void* self, unsigned short* transferFunctionTable, long long transferFunctionTable_len);
void QColorSpace_setTransferFunctions(void* self, unsigned short* redTransferFunctionTable, long long redTransferFunctionTable_len, unsigned short* greenTransferFunctionTable, long long greenTransferFunctionTable_len, unsigned short* blueTransferFunctionTable, long long blueTransferFunctionTable_len);
void* QColorSpace_withTransferFunction(void* self, int transferFunction);
void* QColorSpace_withTransferFunction2(void* self, unsigned short* transferFunctionTable, long long transferFunctionTable_len);
void* QColorSpace_withTransferFunctions(void* self, unsigned short* redTransferFunctionTable, long long redTransferFunctionTable_len, unsigned short* greenTransferFunctionTable, long long greenTransferFunctionTable_len, unsigned short* blueTransferFunctionTable, long long blueTransferFunctionTable_len);
void QColorSpace_setPrimaries(void* self, int primariesId);
void QColorSpace_setPrimaries2(void* self, void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint);
void QColorSpace_setWhitePoint(void* self, void* whitePoint);
void* QColorSpace_whitePoint(void* self);
void QColorSpace_setPrimaryPoints(void* self, void* primaryPoints);
void* QColorSpace_primaryPoints(void* self);
int QColorSpace_transformModel(void* self);
int QColorSpace_colorModel(void* self);
void QColorSpace_detach(void* self);
int QColorSpace_isValid(void* self);
int QColorSpace_isValidTarget(void* self);
void* QColorSpace_fromIccProfile(void* iccProfile);
void* QColorSpace_iccProfile(void* self);
void* QColorSpace_transformationToColorSpace(void* self, void* colorspace);
void* QColorSpace_ToQVariant(void* self);
void QColorSpace_setTransferFunction3(void* self, int transferFunction, float gamma);
void* QColorSpace_withTransferFunction3(void* self, int transferFunction, float gamma);
int QColorSpace_OperatorEqual(void* self, void* colorSpace2);
int QColorSpace_OperatorNotEqual(void* self, void* colorSpace2);
void QColorSpace_delete(void* self);

#ifdef __cplusplus
}
#endif
