package config

import (
	"os"
)

func GetProjectRoot() string {
	return "/Users/lulz/dev/art/"
	root := os.Getenv("PROJECT_ROOT")
	if root == "" {
		panic("env.PROJECT_ROOT not set")
	}

	return root
}
