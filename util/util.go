package util

import (
	"math/rand"
	"time"
)

// RandomString 随机用户名（10位）
func RandomString(n int) string {
	var letters = []byte("abcdefghijklnmopqrstuvwxyz1234567890ABCDEFGHIJKLNMOPQRSTUVWXYZ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
