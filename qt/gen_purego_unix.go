//go:build !windows

package qt

import "github.com/ebitengine/purego"

// shimLibName is the default shim library (bare soname — the loader searches the usual paths;
// $CPPGO_LIB_QT overrides with an explicit path).
const shimLibName = "libshim_qt.so"

// openShimLib: RTLD_LAZY without RTLD_GLOBAL (= RTLD_LOCAL). Each package's shim library
// carries its OWN copy of the cppgo_* C->Go dispatch globals; under RTLD_GLOBAL those would
// interpose across packages (qt + webengine loaded together) and route every callback into
// whichever package init'd last — where the other package's objects aren't registered — so
// ALL signals/virtuals would silently no-op. RTLD_LOCAL keeps each lib's dispatch private.
func openShimLib(path string) (uintptr, error) {
	return purego.Dlopen(path, purego.RTLD_LAZY)
}

func shimSym(name string) (uintptr, error) {
	return purego.Dlsym(_lib, name)
}
