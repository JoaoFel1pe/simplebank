package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var name strings.Builder
	total_letters := len(alphabet)
	for i := 0; i < n; i++ {
		letter := alphabet[rand.Intn(total_letters)]
		name.WriteByte(letter)
	}
	return name.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "BRL"}
	total_currencies := len(currencies)
	return currencies[rand.Intn(total_currencies)]
}
