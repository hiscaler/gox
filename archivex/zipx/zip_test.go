package zipx

import (
	"archive/zip"
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	err := Compress("./a.zip", []string{"./zip.go", "./testdata/a/a.text", "./testdata/b/b.txt"}, zip.Deflate, true)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCompressError(t *testing.T) {
	files := make([]string, 0)
	for i := 0; i <= 100; i++ {
		files = append(files, fmt.Sprintf("%d-not-exists.file", i))
	}
	err := Compress("./a.zip", files, zip.Deflate, true)
	if err == nil {
		t.Error("err is nil")
	} else {
		t.Logf("err = %s", err.Error())
	}
}
