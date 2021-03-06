package hash_test

import (
	"github.com/felipebool/bloom/pkg/hash"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddAndCheckValues_Bool(t *testing.T) {
	testCases := map[string]struct {
		set      []int
		get      []int
		expected []bool
	}{
		"simple case": {
			set:      []int{0, 1, 2, 3, 4},
			get:      []int{0, 1, 2, 3, 4},
			expected: []bool{true, true, true, true, true},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			boolHash := hash.NewBoolHash(100)
			boolHash.SetBits(tc.set)
			got := boolHash.GetValues(tc.get)

			require.Equal(t, len(tc.expected), len(got))
			require.ElementsMatch(t, got, tc.expected)
		})
	}
}
