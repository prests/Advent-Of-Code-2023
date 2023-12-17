package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const MAX_RED_CUBES = 12
const MAX_GREEN_CUBES = 13
const MAX_BLUE_CUBES = 14

type ColorResults struct {
	red   int
	blue  int
	green int
}

type CubeRoundResults struct {
	id           int
	roundResults []*ColorResults
}

//go:embed day-two-input.txt
var efs embed.FS

func parseGameID(input string) (int, error) {
	idAsStr := strings.Split(input, " ")[1]

	return strconv.Atoi(idAsStr)
}

func parseRound(input string) (*ColorResults, error) {
	var roundColors ColorResults
	cubes := strings.Split(input, ",")

	for _, cubeStr := range cubes {
		color := strings.Split(strings.TrimSpace(cubeStr), " ")
		val, err := strconv.Atoi(color[0])
		if err != nil {
			return nil, err
		}

		switch color[1] {
		case "green":
			roundColors.green = val
		case "blue":
			roundColors.blue = val
		case "red":
			roundColors.red = val
		}
	}

	return &roundColors, nil
}

func parseRoundResults(input string) ([]*ColorResults, error) {
	var results []*ColorResults

	rounds := strings.Split(input, ";")
	for _, round := range rounds {
		roundResults, err := parseRound(round)
		if err != nil {
			return nil, err
		}

		results = append(results, roundResults)
	}

	return results, nil
}

func ParseGameResults(input string) (*CubeRoundResults, error) {
	gameData := strings.Split(input, ":")

	id, err := parseGameID(gameData[0])
	if err != nil {
		return nil, err
	}

	roundResults, err := parseRoundResults(gameData[1])
	if err != nil {
		return nil, err
	}

	return &CubeRoundResults{
		id,
		roundResults,
	}, nil
}

func isGameWinnable(roundResults []*ColorResults) bool {
	for _, colors := range roundResults {
		if colors.red > MAX_RED_CUBES || colors.blue > MAX_BLUE_CUBES || colors.green > MAX_GREEN_CUBES {
			return false
		}
	}

	return true
}

func CalculateWins(games []*CubeRoundResults) int {
	sum := 0

	for _, game := range games {
		if isGameWinnable(game.roundResults) {
			sum += game.id
		}
	}

	return sum
}

func FindSmallestCubes(roundResults []*ColorResults) ColorResults {
	var smallestCubes ColorResults

	for _, round := range roundResults {
		if round.blue > smallestCubes.blue {
			smallestCubes.blue = round.blue
		}

		if round.red > smallestCubes.red {
			smallestCubes.red = round.red
		}

		if round.green > smallestCubes.green {
			smallestCubes.green = round.green
		}
	}

	return smallestCubes
}

func CalculatePowers(gameColors []ColorResults) int {
	sum := 0

	for _, game := range gameColors {
		power := game.blue * game.green * game.red
		sum += power
	}

	return sum
}

func main() {
	file, err := efs.Open("day-two-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cubeResults []*CubeRoundResults
	for scanner.Scan() {
		input := scanner.Text()

		gameResults, err := ParseGameResults(input)
		if err != nil {
			log.Fatalf("Error parsing game: %v", err)
		}
		cubeResults = append(cubeResults, gameResults)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	sum := CalculateWins(cubeResults)
	fmt.Printf("sum: %d\n", sum)

	var smallestCubesPerGame []ColorResults
	for _, game := range cubeResults {
		smallestCubes := FindSmallestCubes(game.roundResults)
		smallestCubesPerGame = append(smallestCubesPerGame, smallestCubes)
	}

	powersSum := CalculatePowers(smallestCubesPerGame)
	fmt.Printf("game sum: %d\n", powersSum)
}
