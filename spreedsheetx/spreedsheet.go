package spreedsheetx

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// ColumnName get cell name with name and offset
// A + 1 = B
// AA + 1 = AB
// ZZ + 1 = AAA
func ColumnName(name string, offset int) (string, error) {
	if name == "" {
		return "", errors.New("name is invalid")
	}
	if offset < 0 {
		return "", fmt.Errorf("offset must greater than 0")
	}
	name = strings.ToUpper(name)
	for _, v := range name {
		if !unicode.IsLetter(v) {
			return "", fmt.Errorf("%v is invalid cell name", name)
		}
	}
	a := 65
	z := 90
	n := len(name)
	numbers := make([]int, len(name)+offset/26+1)
	for i := range name {
		numbers[i] = int(name[n-i-1]) // reversed string ascii value
	}
	for i, number := range numbers {
		if i == 0 {
			number += offset
		}
		if number <= z {
			numbers[i] = number
		} else {
			numbers[i] = a
			numbers[i+1] += number - z
		}
	}

	n = len(numbers)
	sb := strings.Builder{}
	sb.Grow(n)
	for i := n - 1; i >= 0; i-- {
		number := numbers[i]
		if number == 0 {
			continue
		}
		if number < a {
			number += 64
		}
		sb.WriteRune(rune(number))
	}
	return sb.String(), nil
}
