package randx

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"strings"
)

func generateValues(str string, len int) string {
	s := ""
	b := bytes.NewBufferString(str)
	bigInt := big.NewInt(int64(b.Len()))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		s += string(str[randomInt.Int64()])
	}
	return s
}

func Letter(len int, upper bool) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	s := generateValues(str, len)
	if upper {
		s = strings.ToUpper(s)
	} else {
		s = strings.ToLower(s)
	}
	return s
}

func Number(len int) string {
	return generateValues("0123456789", len)
}

func Any(len int, upper bool) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	s := generateValues(str, len)
	if upper {
		s = strings.ToUpper(s)
	} else {
		s = strings.ToLower(s)
	}
	return s
}
