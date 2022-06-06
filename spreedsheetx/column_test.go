package spreedsheetx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewColumn(t *testing.T) {
	column := NewColumn("A")
	assert.Equal(t, "A", column.Name())
	column.Next()
	assert.Equal(t, "B", column.Name())
	column.RightShift(2)
	assert.Equal(t, "D", column.Name())
	column.To("F")
	assert.Equal(t, "F", column.Name())
	column.Next()
	assert.Equal(t, "G", column.Name())
	_, err := column.To("ZZZ")
	assert.Equal(t, true, err != nil, "err")
}

func TestNewColumn2(t *testing.T) {
	column := NewColumn("A")
	column.To("ZZ")
	assert.Equal(t, "ZZ", column.Name())
	column.RightShift(2)
	assert.Equal(t, "AAB", column.Name())
}

func TestNewColumn3(t *testing.T) {
	column := NewColumn("ABC")
	assert.Equal(t, "ABC", column.Name())
	column.RightShift(2)
	assert.Equal(t, "ABE", column.Name())
	column.RightShift(53)
	assert.Equal(t, "ADF", column.Name())
	column.To("A")
	column.Next()
	assert.Equal(t, "B", column.Name())
	column.Next()
	assert.Equal(t, "C", column.Name())
	column.To("A")
	column.RightShift(26)
	assert.Equal(t, "AA", column.Name())
}

func TestNewColumn4(t *testing.T) {
	column := NewColumn("A")
	column.RightShift(1000)
	assert.Equal(t, "ALM", column.Name())
	column.To("A")
	column.RightShift(25)
	assert.Equal(t, "Z", column.Name())
	column.RightShift(1)
	assert.Equal(t, "AA", column.Name())
	column.To("A")
	column.RightShift(maxNumber - 1)
	assert.Equal(t, "XFD", column.Name())
}

func TestNewColumn5(t *testing.T) {
	column := NewColumn("A")
	column.Next()
	column.RightShift(26)
	assert.Equal(t, "AB", column.Name())
	column.Reset()
	assert.Equal(t, "A", column.Name())
	assert.Equal(t, "A", column.StartName())
	assert.Equal(t, "A", column.EndName())
}
