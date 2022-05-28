package spreedsheetx

import (
	"fmt"
	"strings"
)

type Column struct {
	startName string // 最开始操作的列
	endName   string // 最远到达的列
	current   string // 当前列
}

func isValidName(name string) bool {
	return name != "" && rxEnglishLetter.MatchString(name)
}

func NewColumn(name string) *Column {
	if !isValidName(name) {
		panic("invalid column name")
	}

	return &Column{
		startName: name,
		endName:   name,
		current:   name,
	}
}

// ToFirst 到第一列，总是返回 A 列
func (c *Column) ToFirst() *Column {
	c.current = "A"
	return c
}

// Next 当前列的下一列
func (c *Column) Next() *Column {
	c.RightShift(1)
	return c
}

// StartName 返回最开始的列名
func (c Column) StartName() string {
	return c.startName
}

func (c *Column) setEndName(name string) *Column {
	if c.endName < name || len(c.endName) < len(name) {
		c.endName = name
	}
	return c
}

// EndName 返回最远到达的列名
func (c Column) EndName() string {
	return c.endName
}

// Name 当前列名
func (c Column) Name() string {
	return c.current
}

// NameWithRow 带行号的列名，比如：A1
func (c Column) NameWithRow(row int) string {
	return fmt.Sprintf("%s%d", c.current, row)
}

// Reset 重置到最开始的列（NewColumn 创建时的列）
func (c *Column) Reset() *Column {
	c.current = c.startName
	return c
}

// To 跳转到指定的列
func (c *Column) To(name string) *Column {
	if !isValidName(name) {
		panic("invalid column name")
	}
	c.current = name
	c.setEndName(name)
	return c
}

// RightShift 基于当前位置右移多少列
func (c *Column) RightShift(steps int) *Column {
	if steps <= 0 {
		return c
	}

	a := 65
	z := 90
	name := strings.ToUpper(c.current)
	n := len(c.current)
	numbers := make([]int, n+11)
	for i := range name {
		numbers[i] = int(name[n-i-1]) // reversed string ascii value
	}

	for i, number := range numbers {
		if i == 0 {
			number += steps/26*z + steps%26
		}
		if number <= z {
			numbers[i] = number
		} else {
			numbers[i] = number % z
			if i+1 >= len(numbers) {
				numbers = append(numbers, 0)
			}
			numbers[i+1] += number / z
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
	c.current = sb.String()
	c.setEndName(c.current)
	return c
}
