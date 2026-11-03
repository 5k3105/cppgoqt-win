#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCommandLineParser_new();
char* QCommandLineParser_tr(const char* sourceText, long long* __out_len);
void QCommandLineParser_setSingleDashWordOptionMode(void* self, int parsingMode);
void QCommandLineParser_setOptionsAfterPositionalArgumentsMode(void* self, int mode);
int QCommandLineParser_addOption(void* self, void* commandLineOption);
int QCommandLineParser_addOptions(void* self, void** options, long long options_len);
void* QCommandLineParser_addVersionOption(void* self);
void* QCommandLineParser_addHelpOption(void* self);
void QCommandLineParser_setApplicationDescription(void* self, const char* description, long long description_len);
char* QCommandLineParser_applicationDescription(void* self, long long* __out_len);
void QCommandLineParser_addPositionalArgument(void* self, const char* name, long long name_len, const char* description, long long description_len);
void QCommandLineParser_clearPositionalArguments(void* self);
void QCommandLineParser_process(void* self, char** arguments, long long arguments_len);
void QCommandLineParser_process2(void* self, void* app);
int QCommandLineParser_parse(void* self, char** arguments, long long arguments_len);
char* QCommandLineParser_errorText(void* self, long long* __out_len);
int QCommandLineParser_isSet(void* self, const char* name, long long name_len);
char* QCommandLineParser_value(void* self, const char* name, long long name_len, long long* __out_len);
char** QCommandLineParser_values(void* self, const char* name, long long name_len, long long* __out_len);
int QCommandLineParser_isSet2(void* self, void* option);
char* QCommandLineParser_value2(void* self, void* option, long long* __out_len);
char** QCommandLineParser_values2(void* self, void* option, long long* __out_len);
char** QCommandLineParser_positionalArguments(void* self, long long* __out_len);
char** QCommandLineParser_optionNames(void* self, long long* __out_len);
char** QCommandLineParser_unknownOptionNames(void* self, long long* __out_len);
void QCommandLineParser_showVersion(void* self);
void QCommandLineParser_showHelp(void* self);
char* QCommandLineParser_helpText(void* self, long long* __out_len);
void QCommandLineParser_showMessageAndExit(int type_, const char* message, long long message_len);
char* QCommandLineParser_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QCommandLineParser_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
void QCommandLineParser_addPositionalArgument2(void* self, const char* name, long long name_len, const char* description, long long description_len, const char* syntax, long long syntax_len);
void QCommandLineParser_showHelp2(void* self, int exitCode);
void QCommandLineParser_showMessageAndExit2(int type_, const char* message, long long message_len, int exitCode);
void QCommandLineParser_delete(void* self);

#ifdef __cplusplus
}
#endif
