package apprentice

import (
	"fmt"
	"path/filepath"
)

// Target represents the target for a compilation
type Target struct {
	GOOS      string
	GOARCH    string
	SourceDir string
}

// Name generates the binary name for the Target
func (t Target) Name() string {
	n := filepath.Base(ModulePath())
	if t.SourceDir != "" {
		n = t.SourceDir
	}
	name := fmt.Sprintf(
		"%s-%s-%s",
		n,
		t.GOOS,
		t.GOARCH,
	)

	if t.GOOS == "windows" {
		name += ".exe"
	}

	return name
}

func (t Target) String() string { return t.Name() }
