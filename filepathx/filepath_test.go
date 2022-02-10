package filepathx

import (
	"fmt"
	"github.com/hiscaler/gox/slicex"
	"os"
	"path/filepath"
	"testing"
)

func TestDirs(t *testing.T) {
	root, _ := os.Getwd()
	fmt.Println(root)
	testCases := []struct {
		Number int
		Path   string
		Option WalkOption
		Files  []string
	}{
		{
			1,
			"/a/b",
			WalkOption{},
			[]string{},
		},
		{
			2,
			root,
			WalkOption{
				CaseSensitive: false,
				Filter: func(path string) bool {
					return path == "2"
				},
				Recursive: true,
			},
			[]string{"2"},
		},
		{
			3,
			root,
			WalkOption{
				CaseSensitive: false,
				Only:          []string{"2"},
				Recursive:     true,
			},
			[]string{"2"},
		},
		{
			4,
			root,
			WalkOption{
				CaseSensitive: false,
				Except:        []string{"2"},
				Recursive:     true,
			},
			[]string{"test", "1", "1.1"},
		},
		{
			5,
			root,
			WalkOption{
				CaseSensitive: false,
				Recursive:     true,
			},
			[]string{"test", "1", "1.1", "2"},
		},
		{
			6,
			root + "/test",
			WalkOption{
				Recursive: true,
			},
			[]string{"1", "1.1", "2"},
		},
	}
	for _, testCase := range testCases {
		files := Dirs(testCase.Path, testCase.Option)
		for i, file := range files {
			files[i] = filepath.Base(file)
		}
		if !slicex.StringSliceEqual(files, testCase.Files, true, true, true) {
			t.Errorf("%d: except %v actual %v", testCase.Number, testCase.Files, files)
		}
	}
}

func TestFiles(t *testing.T) {
	root, _ := os.Getwd()
	testCases := []struct {
		Number int
		Path   string
		Option WalkOption
		Files  []string
	}{
		{
			1,
			"/a/b",
			WalkOption{},
			[]string{},
		},
		{
			2,
			root,
			WalkOption{
				CaseSensitive: false,
				Filter: func(path string) bool {
					return path == "2.txt"
				},
				Recursive: true,
			},
			[]string{"2.txt"},
		},
		{
			3,
			root,
			WalkOption{
				CaseSensitive: false,
				Only:          []string{"2.txt"},
				Recursive:     true,
			},
			[]string{"2.txt"},
		},
		{
			4,
			root,
			WalkOption{
				CaseSensitive: false,
				Except:        []string{"2.txt"},
				Recursive:     true,
			},
			[]string{"filepath.go", "filepath_test.go", "1.1.txt"},
		},
		{
			5,
			root,
			WalkOption{
				CaseSensitive: false,
				Recursive:     true,
			},
			[]string{"filepath.go", "filepath_test.go", "1.1.txt", "2.txt"},
		},
		{
			6,
			root + "/test",
			WalkOption{
				Recursive: true,
			},
			[]string{"1.1.txt", "2.txt"},
		},
	}
	for _, testCase := range testCases {
		files := Files(testCase.Path, testCase.Option)
		for i, file := range files {
			files[i] = filepath.Base(file)
		}
		if !slicex.StringSliceEqual(files, testCase.Files, true, true, true) {
			t.Errorf("%d: except %v actual %v", testCase.Number, testCase.Files, files)
		}
	}
}
