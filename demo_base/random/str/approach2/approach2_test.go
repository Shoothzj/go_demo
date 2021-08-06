package approach2

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
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestApproach2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(randStr(10))
}

func BenchmarkApproach2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_ = randStr(10)
	}
}
