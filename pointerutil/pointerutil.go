package pointerutil

import (
	"time"
)

func String(s string) *string {
	return &s
}

func EmptyString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func Int(i int) *int {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

func Float64(f float64) *float64 {
	return &f
}

func Time(t time.Time) *time.Time {
	return &t
}

func Equal[T comparable](s1, s2 *T) bool {
	switch {
	case s1 == nil && s2 == nil:
		return true
	case s1 != nil && s2 != nil:
		return *s1 == *s2
	default:
		return false
	}
}

func EqualSlices[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return (s1 == nil && s2 == nil) || (s1 != nil && s2 != nil)
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
