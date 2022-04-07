package keyx

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	type testCase struct {
		Number int
		Values interface{}
		Key    string
	}

	b1 := []byte("")
	b2 := []byte("abc")
	testCases := []testCase{
		{1, []interface{}{1, 2, 3}, "123"},
		{2, []interface{}{0, -1, 2, 3}, "0-123"},
		{3, []interface{}{1.1, 2.12, 3.123}, "1.102.123.12"},
		{4, []interface{}{1.1, 2.12, 3.123}, "1.102.123.12"},
		{5, []interface{}{"a", "b", "c"}, "abc"},
		{6, []interface{}{"a", "b", "c", 1, 2, 3}, "abc123"},
		{7, []interface{}{true, true, false, false}, "truetruefalsefalse"},
		{8, []interface{}{[]int{1, 2, 3}}, "123"},
		{9, []interface{}{[...]int{1, 2, 3, 4}}, "1234"},
		{10, []interface{}{struct {
			Username string
			Age      int
		}{}}, "Age0Username"},
		{11, []interface{}{struct {
			Username string
			Age      int
		}{"John", 12}}, "Age12UsernameJohn"},
		{12, []interface{}{User{
			ID:   1,
			Name: "John",
		}}, "User:ID1NameJohn"},
		// byte
		{13, []interface{}{b1, b2}, "abc"},
		// rune
		{14, []interface{}{'a', 'b', 'c'}, "abc"},
		{15, []map[string]string{{"k1": "v1", "k2": "v2"}}, "k1v1k2v2"},
	}
	for _, tc := range testCases {
		key := Generate(tc.Values)
		if key != tc.Key {
			t.Errorf("%d: except：%s actual：%s", tc.Number, tc.Key, key)
		}
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate([]interface{}{struct {
			Username string
			Age      int
		}{"John", 12}})
	}
}
