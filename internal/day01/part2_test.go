package day01

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
				"L68\n" +
				"L30\n" +
				"R48\n" +
				"L5\n" +
				"R60\n" +
				"L55\n" +
				"L1\n" +
				"L99\n" +
				"R14\n" +
				"L82\n",
			expected: "6",
		},
		{
			name: "back and forth on 0",
			input: "" +
				"L50\n" +
				"L5\n" +
				"R5\n",
			expected: "2",
		},
		{
			name:     "large rotation crosses zero multiple times",
			input:    "R1000\n",
			expected: "10",
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
