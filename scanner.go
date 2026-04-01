package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ScanForJavaFiles looks for .class or .java files to ensure the ZIP is valid
func ScanForJavaFiles(directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// We are looking for Java bytecode or source files
		if !info.IsDir() && (strings.HasSuffix(path, ".class") || strings.HasSuffix(path, ".java")) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no valid Java files found in the source")
	}

	return files, nil
}
