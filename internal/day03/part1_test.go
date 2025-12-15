package day03

import "testing"

func TestPart1Examples(t *testing.T) {
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
			expected: "357",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Part1(tt.input)
			if got != tt.expected {
				t.Fatalf("Part1 = %q, want %q", got, tt.expected)
			}
		})
	}
}
