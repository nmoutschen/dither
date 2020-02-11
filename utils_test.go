package dither

import (
	"testing"
)

func TestXorshift(t *testing.T) {
	testCases := []struct {
		S, Expected uint
	}{
		{seed, 2148245535},
		{1, 270369},
		{0xFFFF, 3760001024},
	}

	for i, tc := range testCases {
		if res := xorshift(tc.S); res != tc.Expected {
			t.Errorf("xorshift(%d) = %d; want %d in test case %d", tc.S, res, tc.Expected, i)
		}
	}
}
