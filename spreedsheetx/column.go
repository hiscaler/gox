package spreedsheetx

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// https://support.microsoft.com/en-us/office/excel-specifications-and-limits-1672b34d-7043-467e-8e27-269d656771c3?ui=en-us&rs=en-us&ad=us
// Max index is 16384 XFD

var (
	rxColumnName = regexp.MustCompile("^[A-Za-z]{1,3}$")
)

const (
	minNumber = 1
	maxNumber = 16384
	a         = 64
)

type Column struct {
	startName string // 最开始操作的列
	endName   string // 最远到达的列
	current   string // 当前列
}

func isValidName(name string) bool {
	return rxColumnName.MatchString(name) && toNumber(name) <= maxNumber
}

func reverse(name string) []rune {
	d := []rune(name)
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	return d
}

func toNumber(name string) int {
	name = strings.ToUpper(name)
	n := len(name)
	switch n {
	case 0:
		return 0
	case 1:
		return int(rune(name[0])) - a
	default:
		number := 0
		for i, r := range reverse(name) {
			if i == 0 {
				number += int(r) - a
			} else {
				number += (int(r) - a) * int(math.Pow(26, float64(i)))
			}
		}
		return number
	}
}

func NewColumn(name string) *Column {
	name = strings.ToUpper(name)
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
func (c *Column) Next() (*Column, error) {
	return c.RightShift(1)
}

func (c *Column) Prev() (*Column, error) {
	return c.LeftShift(1)
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
	c.endName = c.startName
	return c
}

// To 跳转到指定的列
func (c *Column) To(name string) (*Column, error) {
	name = strings.ToUpper(name)
	if !isValidName(name) {
		return c, fmt.Errorf("invalid column name %s", name)
	}
	c.current = name
	c.setEndName(name)
	return c, nil
}

func (c *Column) RightShift(steps int) (*Column, error) {
	if steps <= 0 {
		return c, nil
	}
	return c.shift(steps)
}

func (c *Column) LeftShift(steps int) (*Column, error) {
	if steps <= 0 {
		return c, nil
	}
	return c.shift(-steps)
}

// RightShift 基于当前位置右移多少列
func (c *Column) shift(steps int) (*Column, error) {
	if steps == 0 {
		return c, nil
	}

	number := toNumber(c.current)
	number += steps
	if number > maxNumber {
		return c, errors.New("out of max columns")
	} else if number < minNumber {
		return c, errors.New("out of min columns")
	}

	sb := strings.Builder{}
	sb.Grow(3) // Max 3 letters
	times := 0
	for {
		times++
		quotient := number / 26
		remainder := number % 26
		if remainder == 0 {
			sb.WriteRune('Z')
		} else {
			sb.WriteRune(rune(a + remainder))
		}
		if quotient == 0 {
			break
		} else if quotient <= 26 {
			if quotient != 1 || (times >= 1 && remainder != 0) {
				sb.WriteRune(rune(a + quotient))
			}
			break
		}
		number = quotient
	}

	c.current = string(reverse(sb.String()))
	c.setEndName(c.current)
	return c, nil
}
