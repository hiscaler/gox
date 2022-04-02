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
		{"   ", false},
		{"   ", false},
		{"　　　", false},
		{`
  

`, false},
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

func TestIsBlank(t *testing.T) {
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
		b := IsBlank(testCase.String)
		if b != testCase.IsEmpty {
			t.Errorf("%d: %s except %v, actual %v", i, testCase.String, testCase.IsEmpty, b)
		}
	}
}

func TestContainsChinese(t *testing.T) {
	type testCast struct {
		String string
		Has    bool
	}
	testCasts := []testCast{
		{"", false},
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

func TestToNarrow(t *testing.T) {
	testCasts := []struct {
		tag      string
		string   string
		expected string
	}{
		{"t-letter", "aｂｃ", "abc"},
		{"t-number", "０１２３４５６７８９", "0123456789"},
		{"t-letter-number", "a０", "a0"},
		{"t-other", "～！＠#＄％＾＆＊（）－＋？", "~!@#$%^&*()-+?"},
	}
	for _, tc := range testCasts {
		value := ToNarrow(tc.string)
		assert.Equal(t, tc.expected, value, tc.tag)
	}
}

func TestToWiden(t *testing.T) {
	testCasts := []struct {
		tag      string
		string   string
		expected string
	}{
		{"t-letter", "abc", "ａｂｃ"},
		{"t-number", "0123456789", "０１２３４５６７８９"},
		{"t-letter-number", "a0", "ａ０"},
		{"t-other", "~!@#$%^&*()-+?", "～！＠＃＄％＾＆＊（）－＋？"},
	}
	for _, tc := range testCasts {
		value := ToWiden(tc.string)
		assert.Equal(t, tc.expected, value, tc.tag)
	}
}

func TestSplitMore(t *testing.T) {
	type testCast struct {
		Number int
		String string
		Seps   []string
		Values []string
	}
	testCasts := []testCast{
		{1, "abc", []string{}, []string{"abc"}},
		{2, "a b c", []string{}, []string{"a b c"}},
		{3, "a b c,d", []string{}, []string{"a b c,d"}},
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
		values := SplitMore(tc.String, tc.Seps...)
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
		{"5.0 out of 5 stars", []string{"5.0 out of", "stars"}, "5"},
		{"5.0 out of 5 stars", []string{"5.0 out of", "5", "stars"}, ""},
		{"a b a b c d e f g g f e d", []string{"a", "b", "c", "d", "f g"}, "e  g f e"},
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

func TestWordMatched(t *testing.T) {
	tests := []struct {
		tag           string
		string        string
		words         []string
		caseSensitive bool
		except        bool
	}{
		{"t1", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"Towels", "B"}, true, true},
		{"t2", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"towels", "B"}, false, true},
		{"t3.1", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"tow", "A", "B"}, false, true},
		{"t3.2", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"tow", "A", "B"}, true, false},
		{"t4", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"tow"}, false, false},
		{"t5.1", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"20"}, false, false},
		{"t5.2", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"200"}, false, true},
		{"t6", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"Blue Shop"}, true, true},
		{"t7.1", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"blue shop"}, false, true},
		{"t7.2", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"blue shop"}, true, false},
		{"t8", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"Scott "}, false, true},
		{"t9", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"Sheets "}, false, true},
		{"t10.1", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{`.`}, false, false},
		{"t10.2", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{`...................`}, false, false},
		{"t11", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"*"}, false, false},
		{"t12", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"***"}, false, false},
		{"t13.1", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{`.*`}, false, false},
		{"t13.2", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{`B.*x`}, false, false},
		{"t14", "Scott Blue Shop Towels in a Box - 200 Sheets.", []string{"."}, false, false},
		{"t14.1", "Scott Blue Shop Towels in a Box - 200 Sheets.", []string{".*"}, false, false},
		{"t14.2", "Scott Blue Shop Towels in a Box - 200 Sheets.", []string{"Sheets."}, false, true},
		{"t15", "Scott Blue Shop Towels in a Box - 200 Sheets?", []string{"?"}, false, false},
		{"t16", "Scott Blue Shop Towels in a Box - 200 Sheets", []string{"[Sheets]"}, false, false},
		{"t17", "Scott Blue Shop Towels in a Box a-a 200 Sheets", []string{"a-a"}, false, true},
		{"t18", "Scott Blue Shop Towels in a Box--200 Sheets", []string{"-"}, false, false},
		{"t19", "Scott Blue Shop Towels in a Box--200 Sheets", []string{"--"}, false, false},
		{"t20", "Scott Blue Shop Towels in a Box--200 Sheets", []string{"Box--200"}, false, true},
		{"t20.1", "Scott Blue Shop Towels in a Box~200 Sheets", []string{"Box"}, false, false},
		{"t21", "Scott Blue Shop Towels in a Box--200 Sheets 中文", []string{"中文"}, false, true},
		{"t22", "中文汉字", []string{"汉字"}, false, true},
		{"t23", "中a文b汉c字", []string{"汉c字"}, false, true},
		{"t24", "中a文b汉c字", []string{"汉C字"}, false, true},
		{"t25", "中a文b汉c字", []string{"汉C字"}, true, false},
	}
	for _, test := range tests {
		b := WordMatched(test.string, test.words, test.caseSensitive)
		assert.Equal(t, test.except, b, test.tag)
	}
}

func BenchmarkWordMatched(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordMatched("Scott Blue Shop Towels in a Box--200 Sheets", []string{"Throw Pillow Covers", "Throw Pillows", "Patio Furniture Pillows", "Pillow Covers", "Pillowcases", "Pillow Case", "Pillow Cover", "scot", "scottt", "blu", "Shop Towels"}, true)
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		tag           string
		string        string
		words         []string
		caseSensitive bool
		except        bool
	}{
		{"t1", "Hello world!", []string{"he", "He"}, false, true},
		{"t2", "Hello world!", []string{"he", "He"}, true, true},
		{"t3", "Hello world!", []string{"he"}, true, false},
		{"t4", "", []string{""}, true, true},
		{"t5", "", nil, true, true},
		{"t6", "", []string{}, true, true},
		{"t7", "Hello world!", []string{""}, true, true},
	}
	for _, test := range tests {
		b := StartsWith(test.string, test.words, test.caseSensitive)
		assert.Equal(t, test.except, b, test.tag)
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		tag           string
		string        string
		words         []string
		caseSensitive bool
		except        bool
	}{
		{"t1", "Hello world!", []string{"he", "He"}, false, false},
		{"t2", "Hello world!", []string{"he", "He"}, true, false},
		{"t3", "Hello world!", []string{"d!", "!"}, true, true},
		{"t4", "Hello world!", []string{"WORLD!"}, false, true},
		{"t5", "", []string{""}, true, true},
		{"t6", "", nil, true, true},
		{"t7", "", []string{}, true, true},
		{"t8", "Hello world!", []string{""}, true, true},
	}
	for _, test := range tests {
		b := EndsWith(test.string, test.words, test.caseSensitive)
		assert.Equal(t, test.except, b, test.tag)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		tag           string
		string        string
		words         []string
		caseSensitive bool
		except        bool
	}{
		{"t1", "Hello world!", []string{"ol", "LL"}, false, true},
		{"t2", "Hello world!", []string{"ol", "LL"}, true, false},
		{"t3", "Hello world!", []string{"notfound", "world"}, false, true},
		{"t4", "Hello world!", []string{"notfound", "world"}, true, true},
		{"t5", "", []string{""}, true, true},
		{"t6", "Hello world!", []string{""}, true, true},
	}
	for _, test := range tests {
		b := Contains(test.string, test.words, test.caseSensitive)
		assert.Equal(t, test.except, b, test.tag)
	}
}
func TestQuoteMeta(t *testing.T) {
	tests := []struct {
		tag      string
		string   string
		expected string
	}{
		{"t1", `.+\()[]$^*?`, `\.\+\\\(\)\[\]\$\^\*\?`},
		{"t1", `.+\()[]$^*?{}`, `\.\+\\\(\)\[\]\$\^\*\?{}`},
	}
	for _, test := range tests {
		b := QuoteMeta(test.string)
		assert.Equal(t, test.expected, b, test.tag)
	}
}
