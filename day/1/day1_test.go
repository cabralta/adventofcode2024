package main

import (
	"fmt"
	"testing"
)

func TestGetDistance(t *testing.T) {

	var tests = []struct {
		name   string
		input1 int
		input2 int
		want   int
	}{
		{"Distance should be 30", 10, -20, 30},
		{"Distance should be 10", 10, 20, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := GetDistance(tt.input1, tt.input2)
			if ans != tt.want || err != nil {
				t.Errorf("got %s, want %s", fmt.Sprint(ans), fmt.Sprint(tt.want))
			}
		})
	}
}

func TestSumInts(t *testing.T) {

	var tests = []struct {
		name          string
		distancesList []int
		want          int
	}{
		{"Sum should be 11", []int{2, 1, 0, 1, 2, 5}, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SumInts(tt.distancesList)
			if ans != tt.want {
				t.Errorf("got %s, want %s", fmt.Sprint(ans), fmt.Sprint(tt.want))
			}
		})
	}

}

func TestCalcSimilarityScore(t *testing.T) {

	var tests = []struct {
		name  string
		List1 []int
		List2 []int
		want  int
	}{
		{"Sum should be 31", []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scores := CalcSimilarityScore(tt.List1, tt.List2)
			ans := SumInts(scores)
			if ans != tt.want {
				t.Errorf("got %s, want %s", fmt.Sprint(ans), fmt.Sprint(tt.want))
			}
		})
	}

}
