package zipx

import (
	"archive/zip"
	"context"
	"github.com/hiscaler/gox/filex"
	"golang.org/x/sync/errgroup"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type zipFile struct {
	header *zip.FileHeader
	data   *os.File
}

// Compress compresses files and saved, if compactDirectory is true, then will remove all directory path
func Compress(filename string, files []string, method uint16, compactDirectory bool) error {
	zFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer zFile.Close()
	zipWriter := zip.NewWriter(zFile)
	defer zipWriter.Close()

	zipFiles := make([]zipFile, len(files))
	errGrp, ctx := errgroup.WithContext(context.Background())
	for i, file := range files {
		f := file
		j := i
		errGrp.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				zf, e := addFile(f, method, compactDirectory)
				if e != nil {
					ctx.Done()
					return e
				}
				zipFiles[j] = zf
				return nil
			}
		})
	}
	err = errGrp.Wait()
	if err != nil {
		return err
	}

	for i := range zipFiles {
		if zipFiles[i].data == nil {
			continue
		}
		err = func(i int) error {
			defer zipFiles[i].data.Close()
			if err != nil {
				return err // For close all opened files
			}
			writer, e := zipWriter.CreateHeader(zipFiles[i].header)
			if e != nil {
				return e
			}

			_, e = io.Copy(writer, zipFiles[i].data)
			return e
		}(i)
	}
	return err
}

func addFile(filename string, method uint16, compactDirectory bool) (zipFile zipFile, err error) {
	pendingAddFile, err := os.Open(filename)
	if err != nil {
		return
	}

	zipFile.data = pendingAddFile
	info, err := pendingAddFile.Stat()
	if err != nil {
		return
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return
	}

	if compactDirectory {
		header.Name = filepath.Base(filename)
	} else {
		header.Name = filename
	}
	header.Method = method
	zipFile.header = header
	return
}

// UnCompress unzip source file to destination directory
func UnCompress(src, dst string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	defer r.Close()

	// Create destination directory if not exists
	if !filex.Exists(dst) {
		err = os.MkdirAll(dst, fs.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, file := range r.File {
		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		dir := filepath.Dir(path)
		if !filex.Exists(dir) {
			err = os.MkdirAll(dir, fs.ModePerm)
			if err != nil {
				return err
			}
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
	fw, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return err
}
