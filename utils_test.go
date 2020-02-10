package dither

import (
	"testing"
)

func BenchmarkBitInterleave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitInterleave(0xFFFFFFFF)
	}
}

func BenchmarkBitReverse16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitReverse(0x00F0, 16)
	}
}

func BenchmarkBitReverse32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitReverse(0x0000F0FF, 32)
	}
}

func BenchmarkBitReverse64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitReverse(0x00000000F0FFF0FF, 64)
	}
}

func TestBitInterleave(t *testing.T) {
	testCases := []struct {
		N, Expected uint
	}{
		{0b1111, 0b01010101},
		{0b1010, 0b01000100},
		{0xFFFF, 0x55555555},
		{0xFFFFFFFF, 0x5555555555555555},
		{0xFFFFFFFFFFFFFFFF, 0x5555555555555555},
	}

	for i, tc := range testCases {
		if res := bitInterleave(tc.N); res != tc.Expected {
			t.Errorf("bitInterleave(%x) = %x; want %x for test case %d", tc.N, res, tc.Expected, i)
		}
	}
}

func TestBitReverse(t *testing.T) {
	testCases := []struct {
		N, L, Expected uint
	}{
		{0b00001111, 8, 0b11110000},
		{0b01010000, 8, 0b00001010},
		{0b0110000, 7, 0b0000110},
		{0x00F0, 16, 0x0F00},
		{0x0000FF0F, 32, 0xF0FF0000},
		{0x00000000FF0FFF0F, 64, 0xF0FFF0FF00000000},
	}

	for i, tc := range testCases {
		if res := bitReverse(tc.N, tc.L); res != tc.Expected {
			t.Errorf("bitReverse(%x, %d) = %x; want %x for test case %d", tc.N, tc.L, res, tc.Expected, i)
		}
	}
}
