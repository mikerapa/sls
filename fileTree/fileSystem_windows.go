// +build windows

package fileTree

import "syscall"

func isHiddenFile(filePath string) bool {
	prt, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return false
	}

	attributes, err := syscall.GetFileAttributes(prt)
	if err != nil {
		return false
	}

	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
}