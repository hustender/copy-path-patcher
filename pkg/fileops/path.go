package fileops

import (
	"errors"
	"strings"
)

func IsValidPath(path string) (bool, error) {
	if len(path) == 0 {
		return false, errors.New("input string is empty")
	}
	if strings.Contains(path, "/") && strings.Contains(path, "\\") {
		return false, errors.New("mixed path separators are not allowed")
	}
	if !strings.ContainsAny(path, "/\\") {
		return false, errors.New("no valid path separator detected")
	}
	return true, nil
}
