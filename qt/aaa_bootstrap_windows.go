//go:build windows

package qt

// Self-bootstrapping Windows shim: the purego loader reads $CPPGO_LIB_QT during the
// gen_*.go init() functions (symaddr -> _load -> LoadLibrary), so the DLL path must be set
// BEFORE those run. This file's init() runs first (the go tool orders a package's files
// lexically, and "aaa_" sorts before "gen_"). It extracts the embedded, static-Qt shim
// DLL to a per-user cache keyed by content hash and points $CPPGO_LIB_QT at it — so the
// app ships as a single self-contained .exe.

import (
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"os"
	"path/filepath"
)

//go:embed shim_qt.dll
var embeddedShim []byte

func init() {
	if os.Getenv("CPPGO_LIB_QT") != "" {
		return // explicit override wins
	}
	sum := sha256.Sum256(embeddedShim)
	name := "shim_qt-" + hex.EncodeToString(sum[:8]) + ".dll"

	dir := os.Getenv("LOCALAPPDATA")
	if dir == "" {
		dir = os.TempDir()
	}
	dir = filepath.Join(dir, "cppgoqt")
	_ = os.MkdirAll(dir, 0o755)
	path := filepath.Join(dir, name)

	// Extract once; reuse on later launches. The hash in the name means a new build's DLL
	// lands at a new path, so there is never a stale-DLL mismatch. A loaded DLL can't be
	// deleted on Windows, so we never delete — old hashes can be swept lazily if desired.
	if _, err := os.Stat(path); err != nil {
		tmp := path + ".tmp"
		if os.WriteFile(tmp, embeddedShim, 0o644) == nil {
			_ = os.Rename(tmp, path)
		}
	}
	_ = os.Setenv("CPPGO_LIB_QT", path)
}
