package pkg

import (
	"math/rand"
)

func TakeRand[T any](s []T) T {
	// nolint:gosec // don't need cryptographically secure random
	return s[rand.Intn(len(s))]
}
