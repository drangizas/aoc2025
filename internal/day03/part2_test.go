package day03

import "testing"

func TestPart2Examples(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "given example",
			input: "" +
				"987654321111111\n" +
				"811111111111119\n" +
				"234234234234278\n" +
				"818181911112111\n",
			expected: "3121910778619",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Part2(tt.input)
			if got != tt.expected {
				t.Fatalf("Part2 = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestFindMax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			input:    "987654321111111",
			expected: 987654321111,
		},
		{
			input:    "811111111111119",
			expected: 811111111119,
		},
		{
			input:    "234234234234278",
			expected: 434234234278,
		},
		{
			input:    "818181911112111",
			expected: 888911112111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := findMaxDigits(tt.input)
			if got != tt.expected {
				t.Fatalf("Part2 = %v, want %v", got, tt.expected)
			}
		})
	}
}
