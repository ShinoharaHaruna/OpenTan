package utils

import "os"

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false // By default, when encountering other errors
}

// FilesExist 检查给定的文件路径是否都存在
func FilesExist(paths ...string) bool {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return false
		}
	}
	return true
}
