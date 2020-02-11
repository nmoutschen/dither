package dither

import (
	"testing"
)

func BenchmarkReverseInterleave8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseInterleave(0xF0, 0x0F, 8)
	}
}

func BenchmarkReverseInterleave16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseInterleave(0xFF0F, 0xF0FF, 16)
	}
}

func BenchmarkReverseInterleave32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseInterleave(0xFF0FFF0F, 0xF0FFF0FF, 32)
	}
}

func TestReverseInterleave(t *testing.T) {
	testCases := []struct {
		A, B, BC, Expected uint
	}{
		{0b10, 0b00, 2, 0b0010},
		{0b0011, 0b0100, 4, 0b10100100},
	}

	for i, tc := range testCases {
		if res := reverseInterleave(tc.A, tc.B, tc.BC); res != tc.Expected {
			t.Errorf("reverseInterleave(%b, %b, %d) = %b; want %b for test case %d", tc.A, tc.B, tc.BC, res, tc.Expected, i)
		}
	}
}
