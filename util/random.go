package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt генерирует случайное целое число от min до max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString генерирует случайную строку длины n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

//RandomOwner генерирует случайное имя владельца
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney генерирует случайную сумму денег
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency генерирует случайный код валюты
func RandomCurrency() string {
	currencies := []string{USD, EUR, TJS}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail генерирует случайное электронное письмо
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
