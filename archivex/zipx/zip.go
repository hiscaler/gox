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

func UnCompress(src, dst string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		if err = writeFile(file, path); err != nil {
			break
		}
	}
	return err
}

func writeFile(file *zip.File, path string) error {
	fr, err := file.Open()
	if err != nil {
		return err
	}
	defer fr.Close()

	fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return err
}
