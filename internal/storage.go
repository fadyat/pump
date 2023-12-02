package internal

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func readJson(path string, v interface{}) error {
	content, err := readFileBytes(path)

	switch err {
	case nil:
		return json.Unmarshal(content, v)
	case ErrFileNotFound:
		return nil
	default:
		return err
	}
}

func readFileBytes(path string) ([]byte, error) {
	b, err := os.ReadFile(path)

	switch {
	case err == nil:
		return b, nil
	case os.IsNotExist(err):
		return nil, ErrFileNotFound
	default:
		return nil, err
	}
}

func writeJson(path string, v interface{}) error {
	var data, err = json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	var dir = filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
