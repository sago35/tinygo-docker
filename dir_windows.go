package main

import (
	"path/filepath"
	"strings"
)

func modDir(dir string) (string, error) {
	wd := filepath.ToSlash(filepath.Join(`/`, strings.ToLower(dir[0:1]), dir[3:]))
	return wd, nil
}
