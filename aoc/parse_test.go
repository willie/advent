package aoc

import "testing"

func TestParseInts(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"1 2 3", []int{1, 2, 3}},
		{"move 5 from 3 to 7", []int{5, 3, 7}},
		{"x=10, y=-20", []int{10, -20}},
		{"no numbers here", []int{}},
		{"42", []int{42}},
		{"a1b2c3", []int{1, 2, 3}},
		{"-1 -2 -3", []int{-1, -2, -3}},
		{"range: 100-200", []int{100, -200}}, // Note: parses as 100 and -200
	}

	for _, tc := range tests {
		result := ParseInts(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("ParseInts(%q): expected %v, got %v", tc.input, tc.expected, result)
			continue
		}
		for i, v := range result {
			if v != tc.expected[i] {
				t.Errorf("ParseInts(%q): expected %v, got %v", tc.input, tc.expected, result)
				break
			}
		}
	}
}

func TestParseInt64s(t *testing.T) {
	result := ParseInt64s("big: 9999999999999")
	if len(result) != 1 || result[0] != 9999999999999 {
		t.Errorf("ParseInt64s: expected [9999999999999], got %v", result)
	}
}

func TestSplitInts(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []int
	}{
		{"1,2,3", ",", []int{1, 2, 3}},
		{"1|2|3", "|", []int{1, 2, 3}},
		{"10->20->30", "->", []int{10, 20, 30}},
		{"42", ",", []int{42}},
		{"1,,3", ",", []int{1, 3}}, // Empty parts skipped
	}

	for _, tc := range tests {
		result := SplitInts(tc.input, tc.sep)
		if len(result) != len(tc.expected) {
			t.Errorf("SplitInts(%q, %q): expected %v, got %v", tc.input, tc.sep, tc.expected, result)
			continue
		}
		for i, v := range result {
			if v != tc.expected[i] {
				t.Errorf("SplitInts(%q, %q): expected %v, got %v", tc.input, tc.sep, tc.expected, result)
				break
			}
		}
	}
}

func TestWords(t *testing.T) {
	result := Words("hello   world\tfoo")
	expected := []string{"hello", "world", "foo"}

	if len(result) != len(expected) {
		t.Errorf("Words: expected %v, got %v", expected, result)
	}
}

func TestMustInt64(t *testing.T) {
	if MustInt64("9999999999999") != 9999999999999 {
		t.Error("MustInt64: wrong result")
	}
}
