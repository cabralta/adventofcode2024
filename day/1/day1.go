package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(Integer int) int {
	if Integer < 0 {
		return -Integer
	}

	return Integer
}

func GetDistance(LeftInput int, RightInput int) (int, error) {

	var Distance int = 0

	Distance = Abs(LeftInput - RightInput)
	//log.Printf("Distance: %v\n", Distance)
	return Distance, nil
}

// can optimize since its ints and shorten list as we go along
func CalcSimilarityScore(LeftList []int, RightList []int) []int {

	LeftSimilarityScores := make([]int, 0)

	for l := 0; l < len(LeftList); l++ {
		Score := 0
		for r := 0; r < len(RightList); r++ {
			//where values match increase score
			if LeftList[l] == RightList[r] {
				Score++
			}

		}
		LeftSimilarityScores = append(LeftSimilarityScores, LeftList[l]*Score)
	}

	return LeftSimilarityScores
}

func CalcDistances(LeftList []int, RightList []int) []int {

	if len(LeftList) == len(RightList) {
		Distances := make([]int, 0)
		for i := 0; i < len(LeftList); i++ {
			dist, err := GetDistance(LeftList[i], RightList[i])
			if err == nil {
				Distances = append(Distances, dist)
			} else {
				panic(err)
			}

		}

		return Distances

	} else {
		return nil
	}

}

func SumInts(IntList []int) int {
	Sum := 0
	for i := 0; i < len(IntList); i++ {
		Sum += IntList[i]
	}
	return Sum
}

func ReadInputFile(Path string) ([]int, []int) {

	l := make([]int, 0)
	r := make([]int, 0)

	file, _ := os.Open(Path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Fields(line)

		li, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}

		ri, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		l = append(l, li)
		r = append(r, ri)
	}

	return l, r
}

func main() {

	l, r := ReadInputFile("data/input")

	sort.Ints(l)
	sort.Ints(r)

	Distances := CalcDistances(l, r)
	SimilarityScores := CalcSimilarityScore(l, r)

	TotalSimilarityScore := SumInts(SimilarityScores)
	TotalDistance := SumInts(Distances)

	fmt.Printf("Total Distance: %v\nTotal Similarity: %v\n", TotalDistance, TotalSimilarityScore)
}
