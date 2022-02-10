package filepathx

import (
	"github.com/hiscaler/gox/slicex"
	"os"
	"path/filepath"
	"testing"
)

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
					return path == filepath.Clean(root+"/test/2/2.txt")
				},
			},
			[]string{"2.txt"},
		},
		{
			3,
			root,
			WalkOption{
				CaseSensitive: false,
				Only:          []string{"2.txt"},
			},
			[]string{"2.txt"},
		},
		{
			4,
			root,
			WalkOption{
				CaseSensitive: false,
				Except:        []string{"2.txt"},
			},
			[]string{"filepath.go", "filepath_test.go", "1.1.txt"},
		},
		{
			5,
			root,
			WalkOption{
				CaseSensitive: false,
			},
			[]string{"filepath.go", "filepath_test.go", "1.1.txt", "2.txt"},
		},
		{
			6,
			root + "/test",
			WalkOption{},
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
