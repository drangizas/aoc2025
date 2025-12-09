package day02

import "testing"

func TestPart2Examples(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "given example",
			input:    "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			expected: "4174379265",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Part2(tt.input)
			if got != tt.expected {
				t.Fatalf("Part2(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestContainsSequence(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			input:    "11",
			expected: true,
		},
		{
			input:    "111",
			expected: true,
		},
		{
			input:    "1010",
			expected: true,
		},
		{
			input:    "565656",
			expected: true,
		},
		{
			input:    "824824824",
			expected: true,
		},
		{
			input:    "2121212121",
			expected: true,
		},
		{
			input:    "123123123",
			expected: true,
		},
		{
			input:    "1212121212",
			expected: true,
		},
		{
			input:    "1111111",
			expected: true,
		},
		{
			input:    "12",
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := containsSequence(tt.input)
			if got != tt.expected {
				t.Fatalf("containsSequence(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
