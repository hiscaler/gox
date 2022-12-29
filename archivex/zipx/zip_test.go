package zipx

import (
	"archive/zip"
	"fmt"
	"github.com/hiscaler/gox/filex"
	"path/filepath"
	"testing"
)

var files []string

func init() {
	files = []string{
		"./zip.go",
		"./testdata/a/a.txt",
		"./testdata/b/b.txt",
		"./testdata/中国/你好.txt",
	}
}

func TestCompressCompactDirectory(t *testing.T) {
	err := Compress("./a.zip", files, zip.Deflate, true)
	if err != nil {
		t.Error(err)
	} else if !filex.Exists("./a.zip") {
		t.Error("zip file not exists")
	}
}

func TestCompressUnCompactDirectory(t *testing.T) {
	err := Compress("./a.zip", files, zip.Deflate, false)
	if err != nil {
		t.Error(err)
	} else if !filex.Exists("./a.zip") {
		t.Error("zip file not exists")
	}
}

func TestCompressError(t *testing.T) {
	notExistsFiles := make([]string, 0)
	for i := 0; i <= 100; i++ {
		notExistsFiles = append(notExistsFiles, fmt.Sprintf("%d-not-exists.file", i))
	}
	err := Compress("./a.zip", notExistsFiles, zip.Deflate, true)
	if err == nil {
		t.Error("err is nil")
	} else {
		t.Logf("err = %s", err.Error())
	}
}

func TestUnCompress(t *testing.T) {
	TestCompressUnCompactDirectory(t)
	err := UnCompress("./a.zip", "./a")
	if err != nil {
		t.Error(err.Error())
	} else {
		for _, file := range files {
			checkFile := filepath.Join("./a", file)
			if !filex.Exists(checkFile) {
				t.Errorf("%s is not exists", checkFile)
				break
			}
		}
	}
}
