// Static-Qt plugin registration for the self-contained shim DLL: with a static Qt build
// there are no loadable plugin DLLs, so the platform + style plugins must be imported.
#include <QtPlugin>
Q_IMPORT_PLUGIN(QWindowsIntegrationPlugin)
Q_IMPORT_PLUGIN(QModernWindowsStylePlugin)
