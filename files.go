package apprentice

import "os"

// FileExists returns whether the given file or directory exists or not
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// FileSize returns the size of the given file in bytes
func FileSize(path string) int64 {
	f, err := os.Stat(path)
	if err != nil {
		return 0
	}

	return f.Size()
}
