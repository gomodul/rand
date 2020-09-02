package rand

import (
	r "math/rand"
	"time"
)

func Generate(n int, dict []rune) string {
	r.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = dict[r.Intn(len(dict))]
	}
	return string(b)
}
