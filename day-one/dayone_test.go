package main

import (
	"testing"

	testingutils "github.com/prests/advent-of-code-2023/testing"
)

func TestGetCalibrationValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "No digits in string",
			input: "nodigits",
			want:  0,
		},
		{
			name:  "One digit in string",
			input: "thereis1digithere",
			want:  11,
		},
		{
			name:  "Two digits in string",
			input: "Thereis2digitsin3rdtest",
			want:  23,
		},
		{
			name:  "Three digits in string",
			input: "0wow3wholedigits7",
			want:  7,
		},
		{
			name:  "Digit written out as a word",
			input: "thereisfive7",
			want:  57,
		},
		{
			name:  "Digits written out as word after digit",
			input: "there3threeone4four",
			want:  34,
		},
		{
			name:  "Weird combined spelling of a word",
			input: "eighthree",
			want:  83,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := GetCalibrationValue(test.input)

			testingutils.AssertEqual(t, got, test.want)
		})
	}
}

func TestGetSum(t *testing.T) {
	t.Parallel()

	t.Run("should sum up ints", func(t *testing.T) {
		t.Parallel()

		inputs := []int{80, 15, 5}

		got := GetSum(inputs)

		testingutils.AssertEqual(t, got, 100)
	})
}
