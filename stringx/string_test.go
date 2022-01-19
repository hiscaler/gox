package stringx

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		String  string
		IsEmpty bool
	}{
		{"A", false},
		{"", true},
		{"   ", true},
		{"   ", true},
		{"　　　", true},
		{`
  

`, true},
		{`
  
a

`, false},
	}
	for i, testCase := range testCases {
		b := IsEmpty(testCase.String)
		if b != testCase.IsEmpty {
			t.Errorf("%d: %s except %v, actual %v", i, testCase.String, testCase.IsEmpty, b)
		}
	}
}

func TestIsSafeCharacters(t *testing.T) {
	type testCast struct {
		String string
		Safe   bool
	}
	testCasts := []testCast{
		{"", false},
		{" ", false},
		{"a", true},
		{"111", true},
		{"ａ", false},
		{"A_B", true},
		{"A_中B", false},
		{"a.b-c_", true},
		{"_.a.b-c_", true},
		{`\.a.b-c_`, false},
	}
	for _, tc := range testCasts {
		safe := IsSafeCharacters(tc.String)
		if safe != tc.Safe {
			t.Errorf("%s except %v, actual：%v", tc.String, tc.Safe, safe)
		}
	}
}

func TestContainsChinese(t *testing.T) {
	type testCast struct {
		String string
		Has    bool
	}
	testCasts := []testCast{
		{"a", false},
		{"A_B", false},
		{"A_中B", true},
	}
	for _, tc := range testCasts {
		has := ContainsChinese(tc.String)
		if has != tc.Has {
			t.Errorf("%s except %v, actual：%v", tc.String, tc.Has, has)
		}
	}
}

func TestToHalfWidth(t *testing.T) {
	testCasts := []struct {
		Before string
		After  string
	}{
		{"aｂｃ", "abc"},
		{"a０", "a0"},
		{"￣！＠#＄％＾＆＊（）－＋", "~!@#$%^&*()-+"},
		{"０１２３４５６７８９", "0123456789"},
		{"a０", "a0"},
	}
	for _, tc := range testCasts {
		after := ToHalfWidth(tc.Before)
		if after != tc.After {
			t.Errorf("%s except %v, actual：%v", tc.Before, tc.After, after)
		}
	}
}
