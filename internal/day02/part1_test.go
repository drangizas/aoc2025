package day02

import "testing"

func TestPart1Examples(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "given example",
			input:    "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			expected: "1227775554",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Part1(tt.input)
			if got != tt.expected {
				t.Fatalf("Part1(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
