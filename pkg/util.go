package pkg

import "time"

func Ptr[T any](v T) *T {
	return &v
}

func Now() *time.Time {
	f := time.Now().Format(time.RFC3339)
	t, _ := time.Parse(time.RFC3339, f)
	return &t
}
