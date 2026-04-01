package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// FastPackageZip takes a source folder and creates a Minecraft JAR/ZIP instantly
func FastPackageZip(sourceDir, outputPath string) error {
	// Create the target file
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// Walk through the folder and add files to the zip
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a header based on file info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the internal path correctly
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate // High speed compression
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})
}
