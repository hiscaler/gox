package randx

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"strings"
)

func generateValues(str string, len int, upper bool) string {
	if len <= 0 {
		return ""
	}
	buffer := bytes.NewBufferString(str)
	bigInt := big.NewInt(int64(buffer.Len()))
	buffer.Reset()
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		buffer.WriteByte(str[randomInt.Int64()])
	}
	s := buffer.String()
	if upper {
		s = strings.ToUpper(s)
	}
	return s
}

// Letter Generate letter rand string
func Letter(len int, upper bool) string {
	return generateValues("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", len, upper)
}

// Number Generate number rand string
func Number(len int) string {
	return generateValues("0123456789", len, false)
}

// Any Generate number and letter combined string
func Any(len int) string {
	return generateValues("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", len, false)
}
