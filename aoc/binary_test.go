package aoc

import "testing"

// =============================================================================
// HexToBin Tests
// =============================================================================

func TestHexToBin(t *testing.T) {
	// Note: HexToBin requires even-length hex strings (each byte = 2 hex chars)
	tests := []struct {
		input    string
		expected string
	}{
		{"00", "00000000"},
		{"01", "00000001"},
		{"0F", "00001111"},
		{"FF", "11111111"},
		{"A5", "10100101"},
		{"DEADBEEF", "11011110101011011011111011101111"},
	}

	for _, tc := range tests {
		result := HexToBin(tc.input)
		if result != tc.expected {
			t.Errorf("HexToBin(%s): expected %s, got %s", tc.input, tc.expected, result)
		}
	}
}

func TestHexToBinLowercase(t *testing.T) {
	// Should handle lowercase hex
	result := HexToBin("ff")
	if result != "11111111" {
		t.Errorf("HexToBin lowercase: expected 11111111, got %s", result)
	}
}

func TestHexToBinLong(t *testing.T) {
	// Test with longer hex strings
	result := HexToBin("0123456789ABCDEF")
	if len(result) != 64 { // 16 hex chars = 64 bits
		t.Errorf("HexToBin long: expected 64 bits, got %d", len(result))
	}
}

// =============================================================================
// BinToDec Tests
// =============================================================================

func TestBinToDec(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"100", 4},
		{"1111", 15},
		{"11111111", 255},
		{"100000000", 256},
	}

	for _, tc := range tests {
		result := BinToDec(tc.input)
		if result != tc.expected {
			t.Errorf("BinToDec(%s): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}

func TestBinToDecLarge(t *testing.T) {
	// Test with larger numbers
	// 2^32 = 4294967296
	result := BinToDec("100000000000000000000000000000000")
	if result != 4294967296 {
		t.Errorf("BinToDec large: expected 4294967296, got %d", result)
	}
}

func TestBinToDecWithLeadingZeros(t *testing.T) {
	result := BinToDec("00001111")
	if result != 15 {
		t.Errorf("BinToDec leading zeros: expected 15, got %d", result)
	}
}

func TestBinToDecEmpty(t *testing.T) {
	result := BinToDec("")
	if result != 0 {
		t.Errorf("BinToDec empty: expected 0, got %d", result)
	}
}

// =============================================================================
// Round Trip Tests
// =============================================================================

func TestHexBinRoundTrip(t *testing.T) {
	// Convert hex to bin and verify it makes sense
	hex := "FF"
	bin := HexToBin(hex)
	dec := BinToDec(bin)

	if dec != 255 {
		t.Errorf("Round trip FF: expected 255, got %d", dec)
	}
}

func TestHexBinRoundTripComplex(t *testing.T) {
	hex := "CAFE"
	bin := HexToBin(hex)
	dec := BinToDec(bin)

	// 0xCAFE = 51966
	if dec != 51966 {
		t.Errorf("Round trip CAFE: expected 51966, got %d", dec)
	}
}

// =============================================================================
// AoC Specific Patterns
// =============================================================================

func TestHexToBinAoCPattern(t *testing.T) {
	// Common AoC pattern: decode hex packet header
	// Example from AoC 2021 Day 16
	input := "D2FE28"
	bin := HexToBin(input)

	// First 3 bits are version
	version := BinToDec(bin[0:3])
	// Next 3 bits are type ID
	typeID := BinToDec(bin[3:6])

	if version != 6 {
		t.Errorf("AoC pattern version: expected 6, got %d", version)
	}
	if typeID != 4 {
		t.Errorf("AoC pattern typeID: expected 4, got %d", typeID)
	}
}

func TestBinaryBitManipulation(t *testing.T) {
	// Test extracting specific bits
	bin := HexToBin("F0") // 11110000

	// Extract high nibble
	high := BinToDec(bin[0:4])
	if high != 15 {
		t.Errorf("High nibble: expected 15, got %d", high)
	}

	// Extract low nibble
	low := BinToDec(bin[4:8])
	if low != 0 {
		t.Errorf("Low nibble: expected 0, got %d", low)
	}
}
