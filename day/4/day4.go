package main

import (
	"log/slog"
	"time"

	"adventofcode2024.com/day/utils"
)

func CalculateMatches(Input []string, Match string) int {

	SubTotal := 0

	for _, Str := range Input {
		Match := utils.RexString(Str, Match)
		SubTotal += len(Match)
	}

	return SubTotal
}

func CheckForPattern(Matrix [][]string, x int, y int, MatchTerm string) int {
	//Check East

	var EastString string
	var WestString string
	var SouthString string
	var NorthString string
	var NorthEastString string
	var NorthWestString string
	var SouthEastString string
	var SouthWestString string
	Matches := 0
	MatchLen := len(MatchTerm)
	MatchMax := MatchLen - 1
	//create strings to match with starting vector

	for i := 0; i < MatchLen; i++ {

		if (y + MatchMax) < len(Matrix[0]) {
			EastString += string(Matrix[x][y+i])
		}

		if (y - MatchMax) >= 0 {
			WestString += string(Matrix[x][y-i])
		}

		if (x + MatchMax) < len(Matrix) {
			SouthString += string(Matrix[x+i][y])

			if (y + MatchMax) < len(Matrix[0]) {
				SouthEastString += string(Matrix[x+i][y+i])
			}

			if (y - MatchMax) >= 0 {
				SouthWestString += string(Matrix[x+i][y-i])
			}
		}

		if (x - MatchMax) >= 0 {
			NorthString += string(Matrix[x-i][y])

			if (y + MatchMax) < len(Matrix[0]) {
				NorthEastString += string(Matrix[x-i][y+i])
			}

			if (y - MatchMax) >= 0 {
				NorthWestString += string(Matrix[x-i][y-i])
			}

		}

	}

	DirectionalStrings := []string{EastString, WestString, NorthString, SouthString, NorthEastString, NorthWestString, SouthEastString, SouthWestString}

	for _, Str := range DirectionalStrings {

		if Str == MatchTerm {
			Matches++
		}
	}

	return Matches
}

func main() {

	Total := 0
	//path := "testdata/test_input"
	path := "data/aoc_data"
	defer utils.TimeTrack(time.Now(), "Day 4")

	StringMatrix := utils.ReadFileIntoStringSlice(path)
	MatchStr := "XMAS"
	//MatchTerm := []byte(MatchStr)

	for x, line := range StringMatrix {
		for y, _ := range line {
			//fmt.Printf("%s ", string(char))
			Total += CheckForPattern(StringMatrix, x, y, MatchStr)

		}

		//fmt.Printf("\n")

	}
	slog.Info("Match Calculation Completed Successfully.", slog.Int("Total", Total))

}
