package zipx

import (
	"archive/zip"
	"context"
	"golang.org/x/sync/errgroup"
	"io"
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
	errGrp, errCtx := errgroup.WithContext(context.Background())
	for i, file := range files {
		f := file
		j := i
		errGrp.Go(func() error {
			select {
			case <-errCtx.Done():
				return nil
			default:
				zf, e := addFile(f, method, compactDirectory)
				if e != nil {
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

	for _, file := range zipFiles {
		if file.data == nil {
			continue
		}
		err = func() error {
			defer file.data.Close()
			if err != nil {
				return err // For closed all opened files
			}
			writer, e := zipWriter.CreateHeader(file.header)
			if e != nil {
				return e
			}

			_, e = io.Copy(writer, file.data)
			return e
		}()
	}
	return err
}

func addFile(filename string, method uint16, compactDirectory bool) (zipFile zipFile, err error) {
	pendingAddFile, err := os.Open(filename)
	if err != nil {
		return
	}

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
	zipFile.data = pendingAddFile
	return
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
