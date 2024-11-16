package pointerutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	s = "foo"
	i = 1
	b = true
)

func TestString(t *testing.T) {
	t.Run("returns a pointer from a literal", func(tt *testing.T) {
		assert.Equal(tt, "foo", *String("foo"))
	})

	t.Run("returns a pointer from a const", func(tt *testing.T) {
		assert.Equal(tt, "foo", *String(s))
	})

	t.Run("returns a pointer to the empty string", func(tt *testing.T) {
		assert.Equal(tt, "", *String(""))
	})
}

func TestEmptyString(t *testing.T) {
	t.Run("returns a pointer from a string", func(tt *testing.T) {
		assert.Equal(tt, "foo", *EmptyString("foo"))
	})

	t.Run("returns nil for the empty string", func(tt *testing.T) {
		assert.Nil(tt, EmptyString(""))
	})
}

func TestInt(t *testing.T) {
	t.Run("returns a pointer from a literal", func(tt *testing.T) {
		assert.Equal(tt, 1, *Int(1))
	})

	t.Run("returns a pointer from a const", func(tt *testing.T) {
		assert.Equal(tt, 1, *Int(i))
	})
}

func TestBool(t *testing.T) {
	t.Run("returns a pointer from a literal", func(tt *testing.T) {
		assert.Equal(tt, true, *Bool(true))
	})

	t.Run("returns a pointer from a const", func(tt *testing.T) {
		assert.Equal(tt, true, *Bool(b))
	})
}

func TestTime(t *testing.T) {
	t.Run("returns a pointer from a literal", func(tt *testing.T) {
		format := "2006-01-02"
		now := time.Now()
		assert.Equal(tt, now.Format(format), Time(time.Now()).Format(format))
	})
}

func TestEqual(t *testing.T) {
	t.Run("strings", func(tt *testing.T) {
		cases := []struct {
			name   string
			v1     *string
			v2     *string
			result bool
		}{
			{"both nil", nil, nil, true},
			{"first nil", nil, String("foo"), false},
			{"second nil", String("foo"), nil, false},
			{"both non-nil but different", String("foo"), String("bar"), false},
			{"both non-nil but same", String("foo"), String("foo"), true},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(tt *testing.T) {
				assert.Equal(tt, tc.result, Equal(tc.v1, tc.v2))
			})
		}
	})

	t.Run("ints", func(tt *testing.T) {
		cases := []struct {
			name   string
			v1     *int
			v2     *int
			result bool
		}{
			{"both nil", nil, nil, true},
			{"first nil", nil, Int(1), false},
			{"second nil", Int(1), nil, false},
			{"both non-nil but different", Int(1), Int(2), false},
			{"both non-nil but same", Int(1), Int(1), true},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(tt *testing.T) {
				assert.Equal(tt, tc.result, Equal(tc.v1, tc.v2))
			})
		}
	})
}

func TestEqualSlices(t *testing.T) {
	t.Run("strings", func(tt *testing.T) {
		cases := []struct {
			name   string
			v1     []string
			v2     []string
			result bool
		}{
			{"both nil", nil, nil, true},
			{"first nil", nil, []string{"foo"}, false},
			{"second nil", []string{"foo"}, nil, false},
			{"both non-nil but different", []string{"foo"}, []string{"bar"}, false},
			{"both non-nil but same", []string{"foo"}, []string{"foo"}, true},
			{"both non-nil but different lengths", []string{"foo"}, []string{"foo", "bar"}, false},
			{"both non-nil but different orders", []string{"bar", "foo"}, []string{"foo", "bar"}, false},
			{"both non-nil but same longer", []string{"foo", "bar", "baz"}, []string{"foo", "bar", "baz"}, true},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(tt *testing.T) {
				assert.Equal(tt, tc.result, EqualSlices(tc.v1, tc.v2))
			})
		}
	})

	t.Run("ints", func(tt *testing.T) {
		cases := []struct {
			name   string
			v1     []int
			v2     []int
			result bool
		}{
			{"both nil", nil, nil, true},
			{"first nil", nil, []int{1}, false},
			{"second nil", []int{1}, nil, false},
			{"both non-nil but different", []int{1}, []int{2}, false},
			{"both non-nil but same", []int{1}, []int{1}, true},
			{"both non-nil but different lengths", []int{1}, []int{1, 2}, false},
			{"both non-nil but different orders", []int{2, 1}, []int{1, 2}, false},
			{"both non-nil but same longer", []int{1, 2, 3}, []int{1, 2, 3}, true},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(tt *testing.T) {
				assert.Equal(tt, tc.result, EqualSlices(tc.v1, tc.v2))
			})
		}
	})
}
