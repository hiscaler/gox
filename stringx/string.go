package stringx

import (
	"regexp"
	"strings"
	"unicode"
)

// IsEmpty 判断字符串是否为空
func IsEmpty(s string) bool {
	if s == "" || strings.TrimSpace(s) == "" {
		return true
	}
	return false
}

// ToNumber 字符串转换为唯一数字
// https://stackoverflow.com/questions/5459436/how-can-i-generate-a-unique-int-from-a-unique-string
func ToNumber(s string) int {
	number := 0
	runes := []rune(s)
	for i, r := range runes {
		x := 0
		if i != 0 {
			x = int(runes[i-1])
		}
		number += ((x << 16) | (x >> 16)) ^ int(r)
	}
	return number
}

func ContainsChinese(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

// IsSafeCharacters Only include a-zA-Z0-9.-_
// Reference https://www.quora.com/What-are-valid-file-names
func IsSafeCharacters(str string) bool {
	if str == "" {
		return false
	}
	re, _ := regexp.Compile(`^[a-zA-Z0-9\.\-_][a-zA-Z0-9\.\-_]*$`)
	return re.MatchString(str)
}

// ToHalfWidth Full width string to half width
func ToHalfWidth(str string) string {
	replacePairs := make([]string, 0)
	chars := map[string]string{
		"０": "0",
		"１": "1",
		"２": "2",
		"３": "3",
		"４": "4",
		"５": "5",
		"６": "6",
		"７": "7",
		"８": "8",
		"９": "9",
		"ａ": "a",
		"ｂ": "b",
		"ｃ": "c",
		"ｄ": "d",
		"ｅ": "e",
		"ｆ": "f",
		"ｇ": "g",
		"ｈ": "h",
		"ｉ": "i",
		"ｊ": "j",
		"ｋ": "k",
		"ｌ": "l",
		"ｍ": "m",
		"ｎ": "n",
		"ｏ": "o",
		"ｐ": "p",
		"ｑ": "q",
		"ｒ": "r",
		"ｓ": "s",
		"ｔ": "t",
		"ｕ": "u",
		"ｖ": "v",
		"ｗ": "w",
		"ｘ": "x",
		"ｙ": "y",
		"ｚ": "z",
		"Ａ": "A",
		"Ｂ": "B",
		"Ｃ": "C",
		"Ｄ": "D",
		"Ｅ": "E",
		"Ｆ": "F",
		"Ｇ": "G",
		"Ｈ": "H",
		"Ｉ": "I",
		"Ｊ": "J",
		"Ｋ": "K",
		"Ｌ": "L",
		"Ｍ": "M",
		"Ｎ": "N",
		"Ｏ": "O",
		"Ｐ": "P",
		"Ｑ": "Q",
		"Ｒ": "R",
		"Ｓ": "S",
		"Ｔ": "T",
		"Ｕ": "U",
		"Ｖ": "V",
		"Ｗ": "W",
		"Ｘ": "X",
		"Ｙ": "Y",
		"Ｚ": "Z",
		"（": "(",
		"）": ")",
		"〔": "[",
		"〕": "]",
		"【": "[",
		"】": "]",
		"〖": "[",
		"〗": "]",
		"“": "\"",
		"”": "\"",
		"‘": "'",
		"’": "'",
		"｛": "{",
		"｝": "}",
		"《": "<",
		"》": ">",
		"：": ":",
		"。": ".",
		"，": ",",
		"、": ".",
		"；": ",",
		"？": "?",
		"！": "!",
		"…": "-",
		"‖": "|",
		"｜": "|",
		"〃": "\"",
		"＄": "$",
		"＠": "@",
		"＃": "#",
		"＾": "^",
		"＆": "&",
		"＊": "*",
		"％": "%",
		"＋": "+",
		"—": "-",
		"－": "-",
		"～": "-",
		"￣": "~",
	}
	for k, v := range chars {
		replacePairs = append(replacePairs, k, v)
	}

	return strings.NewReplacer(replacePairs...).Replace(str)
}
