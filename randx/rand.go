package randx

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	randNumberChars = "0123456789"
	randLetterChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func generateValues(str string, n int, upper bool) string {
	if n <= 0 {
		return ""
	}

	sb := strings.Builder{}
	sb.Grow(n)
	bigInt := big.NewInt(int64(len(str)))
	for i := 0; i < n; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		sb.WriteByte(str[randomInt.Int64()])
	}
	s := sb.String()
	if upper {
		s = strings.ToUpper(s)
	}
	return s
}

// Letter Generate letter rand string
func Letter(n int, upper bool) string {
	return generateValues(randLetterChars, n, upper)
}

// Number Generate number rand string
func Number(n int) string {
	return generateValues(randNumberChars, n, false)
}

// Any Generate number and letter combined string
func Any(n int) string {
	return generateValues(randLetterChars+randNumberChars, n, false)
}
