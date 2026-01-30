package aoc

import "testing"

// =============================================================================
// AtoI Tests
// =============================================================================

func TestAtoI(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0", 0},
		{"1", 1},
		{"42", 42},
		{"-5", -5},
		{"12345", 12345},
		{"-99999", -99999},
	}

	for _, tc := range tests {
		result := AtoI(tc.input)
		if result != tc.expected {
			t.Errorf("AtoI(%s): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}

func TestAtoILeadingZeros(t *testing.T) {
	result := AtoI("007")
	if result != 7 {
		t.Errorf("AtoI leading zeros: expected 7, got %d", result)
	}
}

func TestAtoIPositiveSign(t *testing.T) {
	// Note: strconv.Atoi handles positive signs
	result := AtoI("+42")
	if result != 42 {
		t.Errorf("AtoI positive sign: expected 42, got %d", result)
	}
}

// =============================================================================
// BtoI Tests
// =============================================================================

func TestBtoI(t *testing.T) {
	if BtoI(true) != 1 {
		t.Error("BtoI(true): expected 1")
	}
	if BtoI(false) != 0 {
		t.Error("BtoI(false): expected 0")
	}
}

func TestBtoIWithComparison(t *testing.T) {
	// Common AoC pattern: count matching conditions
	count := BtoI(5 > 3) + BtoI(2 < 1) + BtoI(10 == 10)
	if count != 2 {
		t.Errorf("BtoI comparison: expected 2, got %d", count)
	}
}

// =============================================================================
// ReplaceAll Tests
// =============================================================================

func TestReplaceAll(t *testing.T) {
	result := ReplaceAll("a,b;c:d", ",;:", " ")
	if result != "a b c d" {
		t.Errorf("ReplaceAll: expected 'a b c d', got '%s'", result)
	}
}

func TestReplaceAllEmpty(t *testing.T) {
	result := ReplaceAll("abc", "", "X")
	if result != "abc" {
		t.Errorf("ReplaceAll empty chars: expected 'abc', got '%s'", result)
	}
}

func TestReplaceAllNoMatch(t *testing.T) {
	result := ReplaceAll("hello world", "xyz", "!")
	if result != "hello world" {
		t.Errorf("ReplaceAll no match: expected 'hello world', got '%s'", result)
	}
}

func TestReplaceAllMultiple(t *testing.T) {
	// Replace all digits with X
	result := ReplaceAll("a1b2c3", "123456789", "X")
	if result != "aXbXcX" {
		t.Errorf("ReplaceAll digits: expected 'aXbXcX', got '%s'", result)
	}
}

func TestReplaceAllRemove(t *testing.T) {
	// Remove characters by replacing with empty string
	result := ReplaceAll("a,b,c,d", ",", "")
	if result != "abcd" {
		t.Errorf("ReplaceAll remove: expected 'abcd', got '%s'", result)
	}
}

func TestReplaceAllSpecialChars(t *testing.T) {
	result := ReplaceAll("hello[world]", "[]", "")
	if result != "helloworld" {
		t.Errorf("ReplaceAll special: expected 'helloworld', got '%s'", result)
	}
}

// =============================================================================
// AoC Parsing Patterns
// =============================================================================

func TestParsingPattern(t *testing.T) {
	// Common AoC pattern: clean up input line
	input := "move 5 from 3 to 7"
	cleaned := ReplaceAll(input, "movefrt", " ")
	// Now split on whitespace to get numbers

	// This cleans: "     5      3    7" (multiple spaces)
	// Would then need to handle splitting
	if len(cleaned) < len("move") {
		t.Error("Parsing pattern: string should be modified")
	}
}

func TestParsingNumbers(t *testing.T) {
	// Common pattern: extract numbers from string
	inputs := []string{"123", "-45", "0", "999"}
	sum := 0
	for _, s := range inputs {
		sum += AtoI(s)
	}
	if sum != 1077 {
		t.Errorf("Parsing numbers: expected 1077, got %d", sum)
	}
}

// =============================================================================
// Helpers Tests (without network)
// =============================================================================

func TestTestFunction(t *testing.T) {
	// The Test function prints output, we can't easily test it
	// but we can verify it doesn't panic
	Test("test", 42, 42)   // Should print PASS
	Test("test", 42, 99)   // Should print FAIL
	Test("test", int64(42), int64(42)) // int64 version
}

func TestRunFunction(t *testing.T) {
	// The Run function just prints, verify no panic
	Run("result", 42)
	Run("result", "hello")
}

func TestTestXFunction(t *testing.T) {
	// TestX with paired results and expected values
	TestX("multi", 1, 2, 3, 1, 2, 3) // All should pass
}

func TestRunXFunction(t *testing.T) {
	// RunX prints multiple results
	RunX("results", 1, 2, 3)
}
