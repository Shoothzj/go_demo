package approach3

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(b)
}

func TestApproach3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(randStr(10))
}

func BenchmarkApproach3(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_ = randStr(10)
	}
}
