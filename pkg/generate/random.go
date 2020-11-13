package generate

import (
	"math/rand"
	"strings"
	"time"

	"github.com/icrowley/fake"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer, used for tests
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n, used for tests
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail returns a random email address
func RandomEmail() string {
	return fake.EmailAddress()
}