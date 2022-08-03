package filepathx

import (
	"github.com/hiscaler/gox/slicex"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestDirs(t *testing.T) {
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
				FilterFunc: func(path string) bool {
					return filepath.Base(path) == "2"
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
			[]string{"testdata", "1", "1.1", "1.1", "1.1.1"},
		},
		{
			5,
			root,
			WalkOption{
				CaseSensitive: false,
				Recursive:     true,
			},
			[]string{"testdata", "1", "1.1", "1.1", "2", "1.1.1"},
		},
		{
			6,
			root + "/testdata",
			WalkOption{
				Recursive: true,
			},
			[]string{"1", "1.1", "1.1", "2", "1.1.1"},
		},
		{
			7,
			root + "/testdata",
			WalkOption{
				Recursive: false,
			},
			[]string{"1", "2"},
		},
	}
	for _, testCase := range testCases {
		dirs := Dirs(testCase.Path, testCase.Option)
		for i, dir := range dirs {
			dirs[i] = filepath.Base(dir)
		}
		if !slicex.StringSliceEqual(dirs, testCase.Files, true, true, true) {
			t.Errorf("%d: except %v actual %v", testCase.Number, testCase.Files, dirs)
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
				FilterFunc: func(path string) bool {
					return filepath.Base(path) == "2.txt"
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
			[]string{"filepath.go", "filepath_test.go", "1.1.txt", "中文_ZH (1).txt", "中文_ZH (1).txt", "中文_ZH (9).txt", "0.txt"},
		},
		{
			5,
			root,
			WalkOption{
				CaseSensitive: false,
				Recursive:     true,
			},
			[]string{"filepath.go", "filepath_test.go", "1.1.txt", "2.txt", "中文_ZH (1).txt", "中文_ZH (1).txt", "中文_ZH (9).txt", "0.txt"},
		},
		{
			6,
			root + "/testdata",
			WalkOption{
				Recursive: true,
			},
			[]string{"1.1.txt", "2.txt", "中文_ZH (1).txt", "中文_ZH (1).txt", "中文_ZH (9).txt", "0.txt"},
		},
		{
			7,
			root + "/testdata",
			WalkOption{
				Recursive: false,
			},
			[]string{"0.txt"},
		},
		{
			8,
			root + "/testdata/1/1.1/1.1",
			WalkOption{
				Recursive: false,
			},
			[]string{"中文_ZH (1).txt"},
		},
		{
			9,
			"./testdata/1/1.1/1.1",
			WalkOption{
				Recursive: false,
			},
			[]string{"中文_ZH (1).txt"},
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

func TestGenerateDirNames(t *testing.T) {
	tests := []struct {
		tag           string
		string        string
		n             int
		level         int
		caseSensitive bool
		dirs          []string
	}{
		{"t1", "abc", 0, 1, true, []string{"abc"}},
		{"t2", "abc", 1, 1, true, []string{"a"}},
		{"t3", "abc", 1, 2, true, []string{"a", "b"}},
		{"t4", "abc", 1, 3, true, []string{"a", "b", "c"}},
		{"t5", "abc", 2, 1, true, []string{"ab"}},
		{"t6", "abc", 2, 2, true, []string{"ab", "c"}},
		{"t7", " a b c ", 2, 2, true, []string{"ab", "c"}},
		{"t7", " a b cdefghijklmn ", 2, 3, true, []string{"ab", "cd", "ef"}},
		{"t8", " a", 12, 3, true, []string{"a"}},
		{"t9", " a中文$b", 12, 3, true, []string{"ab"}},
	}
	for _, test := range tests {
		names := GenerateDirNames(test.string, test.n, test.level, test.caseSensitive)
		assert.Equal(t, test.dirs, names, test.tag)
	}
}

func TestExt(t *testing.T) {
	root, _ := os.Getwd()
	tests := []struct {
		tag  string
		path string
		b    []byte
		ext  string
	}{
		{"t1", "/a/b", nil, ""},
		{"t2", "https://golang.org/doc/gopher/fiveyears.jpg", nil, ".jpg"},
		{"t3", filepath.Join(root, "/testdata/2/2.txt"), nil, ".txt"},
		{"t4", filepath.Join(root, "/testdata/2/1.jpg"), nil, ".jpg"},
		{"t5", filepath.Join(root, "/testdata/2/1.pdf"), nil, ".pdf"},
		{"t6", filepath.Join(root, "/testdata/2/1111.pdf"), nil, ".pdf"},
		{"t7", filepath.Join(root, "/testdata/1.xlsx"), nil, ".xlsx"},
	}
	for _, test := range tests {
		ext := Ext(test.path, test.b)
		assert.Equal(t, test.ext, ext, test.tag)
	}
}
