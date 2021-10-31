// +build !windows

package fileTree

import (
	"path/filepath"
	"strings"
)


func isHiddenFile(filePath string) bool {
	return strings.HasPrefix(filepath.Base(filePath), ".")
}

