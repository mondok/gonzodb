package util

import "os"

// FileExists checks if a file exists
// on the local file system
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
