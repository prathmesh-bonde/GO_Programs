package util

// generatng random data for testing

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// special func that gets called everytime pkg is used
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string using given alphabets
func RandomString(n int) string {
	var sb strings.Builder
	slen := len(alphabet)

	for i := 0; i < n; i++ {
		s := alphabet[rand.Intn(slen)]

		sb.WriteByte(s)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance generates a random amount of balance
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "INR"}
	clen := len(currency)

	return currency[rand.Intn(clen)]
}
