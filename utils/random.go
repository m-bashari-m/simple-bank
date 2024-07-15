package utils

import (
	"math/rand"
	"strings"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

var currencies = [...]string{"USD", "CAD", "EUR"}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < length; i++ {
		sb.WriteByte(alphabets[rand.Intn(k)])
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 999999)
}

func RandomCurrency() string {
	return currencies[rand.Intn(len(currencies))]
}
