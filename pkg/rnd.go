package pkg

import "math/rand"

func RandString(l int) string {
	var (
		chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		b     = make([]rune, l)
	)

	for i := range b {
		// nolint:gosec // don't need cryptographically secure random
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}
