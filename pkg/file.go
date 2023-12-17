package pkg

import (
	"path/filepath"
	"strings"
)

func GetDir(path string) string {
	return path[:strings.LastIndex(path, "/")]
}

func RenameWithSuffix(path, suffix string) string {
	var ext = filepath.Ext(path)
	return path[:len(path)-len(ext)] + suffix + ext
}
