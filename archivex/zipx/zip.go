package zipx

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Compress compresses files and saved, if compactDirectory is true, then will remove all directory path
func Compress(filename string, files []string, method uint16, compactDirectory bool) error {
	zipFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = addFile(zipWriter, file, method, compactDirectory); err != nil {
			return err
		}
	}
	return nil
}

func addFile(zipWriter *zip.Writer, filename string, method uint16, compactDirectory bool) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer fileToZip.Close()
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	if compactDirectory {
		header.Name = filepath.Base(filename)
	} else {
		header.Name = filename
	}
	header.Method = method
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
