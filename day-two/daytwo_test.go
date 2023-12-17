package main

import (
	"testing"

	testingutils "github.com/prests/advent-of-code-2023/testing"
)

func TestCubeGameParser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  *CubeRoundResults
	}{
		{
			name:  "Simple game",
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: &CubeRoundResults{
				id: 1,
				roundResults: []*ColorResults{
					{
						red:  4,
						blue: 3,
					},
					{
						red:   1,
						green: 2,
						blue:  6,
					},
					{
						green: 2,
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseGameResults(test.input)
			testingutils.AssertNoError(t, err)
			testingutils.AssertDeepEqual(t, got, test.want)
		})
	}
}

func TestCalculateWins(t *testing.T) {
	t.Parallel()

	t.Run("calculate winnable games", func(t *testing.T) {
		t.Parallel()

		games := []*CubeRoundResults{
			{
				id: 1,
				roundResults: []*ColorResults{
					{
						red:  4,
						blue: 3,
					},
					{
						red:   1,
						green: 2,
						blue:  6,
					},
					{
						green: 2,
					},
				},
			},
			{
				id: 2,
				roundResults: []*ColorResults{
					{
						blue:  1,
						green: 2,
					},
					{
						green: 3,
						blue:  4,
						red:   1,
					},
					{
						green: 1,
						blue:  1,
					},
				},
			},
			{
				id: 3,
				roundResults: []*ColorResults{
					{
						green: 8,
						blue:  6,
						red:   20,
					},
					{
						blue:  5,
						red:   4,
						green: 13,
					},
					{
						green: 5,
						red:   1,
					},
				},
			},
		}

		got := CalculateWins(games)

		testingutils.AssertEqual(t, got, 3)
	})
}

func TestSmallestCube(t *testing.T) {
	t.Parallel()

	t.Run("Find smallest cubes for a game", func(t *testing.T) {
		t.Parallel()

		games := []*ColorResults{
			{
				red:  4,
				blue: 3,
			},
			{
				red:   1,
				green: 2,
				blue:  6,
			},
			{
				green: 2,
			},
		}

		got := FindSmallestCubes(games)

		testingutils.AssertDeepEqual(t, got, ColorResults{red: 4, blue: 6, green: 2})
	})

	t.Run("Calculate Powers", func(t *testing.T) {
		t.Parallel()

		gameColors := []ColorResults{
			{
				red:  4,
				blue: 3,
			},
			{
				red:   1,
				green: 2,
				blue:  6,
			},
			{
				green: 2,
			},
		}

		got := CalculatePowers(gameColors)

		testingutils.AssertEqual(t, got, 12)
	})
}
