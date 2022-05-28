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
	column.To("ZZZ")
	assert.Equal(t, "ZZZ", column.Name())
	column.Next()
	assert.Equal(t, "AAAA", column.Name())
	column.To("ZZZ")
	assert.Equal(t, "ZZZ", column.Name())
	column.RightShift(2)
	assert.Equal(t, "AAAB", column.Name())
	assert.Equal(t, "A", column.StartName())
	assert.Equal(t, "A", column.Reset().Name())
	assert.Equal(t, "AAAB", column.EndName())
}

func TestNewColumn2(t *testing.T) {
	column := NewColumn("A")
	column.To("ZZ")
	assert.Equal(t, "ZZ", column.Name())
	column.RightShift(2)
	assert.Equal(t, "AAB", column.Name())
}

func TestNewColumn3(t *testing.T) {
	column := NewColumn("ZZZ")
	assert.Equal(t, "ZZZ", column.Name())
	column.RightShift(2)
	assert.Equal(t, "AAAB", column.Name())
	column.RightShift(53)
	assert.Equal(t, "AACC", column.Name())
	column.To("A").Next().Next()
	assert.Equal(t, "C", column.Name())
	column.To("A").RightShift(26)
	assert.Equal(t, "AA", column.Name())
}

// todo not pass
func TestNewColumn4(t *testing.T) {
	column := NewColumn("A")
	column.RightShift(1000)
	assert.Equal(t, "ALM", column.Name())
}
