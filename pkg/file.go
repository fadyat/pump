package pkg

import (
	"os"
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

func HomeDirConfig(filename string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "pump", filename), nil
}
