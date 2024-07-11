package util

import (
	"path"
	"strings"

	"github.com/lulzshadowwalker/art/internal/config"
)

func Template(filename string) string {
	return path.Join(config.GetProjectRoot(), "internal/template", filename)
}

func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}

func FirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	return s[:1]
}
