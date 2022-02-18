package stringx

import (
	"github.com/hiscaler/gox/slicex"
	"github.com/stretchr/testify/assert"
	"strings"
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

func TestSplit(t *testing.T) {
	type testCast struct {
		Number int
		String string
		Seps   []string
		Values []string
	}
	testCasts := []testCast{
		{1, "abc", []string{}, []string{"abc"}},
		{2, "a b c", []string{}, []string{"a", "b", "c"}},
		{3, "a b c,d", []string{}, []string{"a", "b", "c,d"}},
		{4, "a b c,d", []string{",", " "}, []string{"a", "b", "c", "d"}},
		{5, "a,b,c,d", []string{",", " "}, []string{"a", "b", "c", "d"}},
		{6, "a,b,c,d e", []string{",", " "}, []string{"a", "b", "c", "d", "e"}},
		{7, "a.,b,c,d e", []string{",", " "}, []string{"a.", "b", "c", "d", "e"}},
		{8, "a.,b,c,d e", []string{",", ".", " "}, []string{"a", "b", "c", "d", "e"}},
		{9, "a.,b,c,d e", []string{",", " "}, []string{"a.", "b", "c", "d", "e"}},
		{10, "a.,b,c,d e", []string{",", "", " "}, []string{"a", ".", "b", "c", "d", "e"}},
		{11, "hello, world!!!", []string{",", " ", "!"}, []string{"hello", "world"}},
		{12, "WaterWipes Original Baby Wipes, 99.9% Water, Unscented & Hypoallergenic for Sensitive Newborn Skin, 3 Packs (180 Count)", []string{",", " ", "!"}, []string{"WaterWipes", "Original", "Baby", "Wipes", "99.9%", "Water", "Unscented", "&", "Hypoallergenic", "for", "Sensitive", "Newborn", "Skin", "3", "Packs", "(180", "Count)"}},
	}
	for _, tc := range testCasts {
		values := SplitWord(tc.String, tc.Seps...)
		if !slicex.StringSliceEqual(values, tc.Values, false, false, true) {
			t.Errorf("%d except %#v, actual：%#v", tc.Number, tc.Values, values)
		}
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		Number int
		Value  interface{}
		Except string
	}{
		{1, false, "false"},
		{2, true, "true"},
		{3, 1, "1"},
		{4, 1.1, "1.1"},
		{5, "abc", "abc"},
		{6, [2]int{1, 2}, "[1,2]"},
		{7, []int{1, 2}, "[1,2]"},
		{8, []string{"a", "b"}, `["a","b"]`},
	}
	for _, testCase := range testCases {
		s := String(testCase.Value)
		if !strings.EqualFold(s, testCase.Except) {
			t.Errorf("%d except: %s, actual: %s", testCase.Number, testCase.Except, s)
		}
	}
}

func TestRemoveEmoji(t *testing.T) {
	testCases := []struct {
		Number       int
		BeforeString string
		AfterString  string
	}{
		{1, "👶hi", "hi"},
		{2, "1👰", "1"},
		{3, "1👉2🤟👉👰3🤟👉👶你好🤟", "123你好"},
		{4, "1👉2🤟👉👰3🤟👉👶你  　　好🤟", "123你  　　好"},
	}
	for _, testCase := range testCases {
		s := RemoveEmoji(testCase.BeforeString, true)
		if !strings.EqualFold(s, testCase.AfterString) {
			t.Errorf("%d except: %s, actual: %s", testCase.Number, testCase.AfterString, s)
		}
	}
}

func TestTrimSpecial(t *testing.T) {
	var testCases = []struct {
		string       string
		replacePairs []string
		expected     string
	}{
		{"  a", []string{}, "a"},
		{"  a", []string{"b", "c"}, "a"},
		{"  a", []string{"a", "c"}, ""},
		{" a       ", []string{}, "a"},
		{`
						
a

`, []string{}, "a"},
		{"  ab", []string{"b"}, "a"},
		{"  a b ", []string{"b"}, "a"},
		{"  a b b", []string{"b"}, "a"},
		{"  a b a", []string{"b"}, "a  a"},
	}
	for _, testCase := range testCases {
		actual := TrimSpecial(testCase.string, testCase.replacePairs...)
		if actual != testCase.expected {
			t.Errorf("TrimSpecial(%s, %s) = %s; expected %s", testCase.string, testCase.replacePairs, actual, testCase.expected)
		}
	}
}

func TestRemoveExtraSpace(t *testing.T) {
	var testCases = []struct {
		number   int
		string   string
		expected string
	}{
		{1, "  a", "a"},
		{2, "  a", "a"},
		{3, "  a", "a"},
		{4, " a       ", "a"},
		{5, `
						
a

`, "a"},
		{6, "  ab", "ab"},
		{7, "  a b ", "a b"},
		{8, "  a b    b", "a b b"},
		{9, "　　　hello,    world!", "hello, world!"},
		{10, `
　　　hello,    




					world!
`, "hello, world!"},
	}
	for _, testCase := range testCases {
		actual := RemoveExtraSpace(testCase.string)
		if actual != testCase.expected {
			t.Errorf("%d RemoveExtraSpace(%s) = '%s'; expected %s", testCase.number, testCase.string, actual, testCase.expected)
		}
	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		tag        string
		bytesValue []byte
		string     string
	}{
		{"t1", []byte{'a'}, "a"},
		{"t2", []byte("abc"), "abc"},
		{"t3", []byte("a b c "), "a b c "},
	}
	for _, test := range tests {
		b := ToBytes(test.string)
		assert.Equal(t, test.bytesValue, b, test.tag)
	}
}
