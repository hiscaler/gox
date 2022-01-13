package filex

import (
	"os"
	"testing"
)

func TestIsDir(t *testing.T) {
	root, _ := os.Getwd()
	testCases := []struct {
		Path   string
		Except bool
	}{
		{"/a/b", false},
		{root, true},
		{root + "/file.go", false},
		{root + "/file", false},
	}
	for _, testCase := range testCases {
		v := IsDir(testCase.Path)
		if v != testCase.Except {
			t.Errorf("`%s` except %v actual %v", testCase.Path, testCase.Except, v)
		}
	}
}

func TestIsFile(t *testing.T) {
	root, _ := os.Getwd()
	testCases := []struct {
		Path   string
		Except bool
	}{
		{"/a/b", false},
		{root, false},
		{root + "/file.go", true},
		{root + "/file", false},
	}
	for _, testCase := range testCases {
		v := IsFile(testCase.Path)
		if v != testCase.Except {
			t.Errorf("`%s` except %v actual %v", testCase.Path, testCase.Except, v)
		}
	}
}

func TestExists(t *testing.T) {
	root, _ := os.Getwd()
	testCases := []struct {
		Path   string
		Except bool
	}{
		{"/a/b", false},
		{root, true},
		{root + "/file.go", true},
		{root + "/file", false},
	}
	for _, testCase := range testCases {
		v := Exists(testCase.Path)
		if v != testCase.Except {
			t.Errorf("`%s` except %v actual %v", testCase.Path, testCase.Except, v)
		}
	}
}
