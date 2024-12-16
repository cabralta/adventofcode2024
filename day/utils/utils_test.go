package utils

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {

	var tests = []struct {
		name   string
		input1 int
		want   int
	}{
		{"Value should be positive", 10, 10},
		{"Distance should be 10", -10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Abs(tt.input1)
			if ans != tt.want {
				t.Errorf("got %s, want %s", fmt.Sprint(ans), fmt.Sprint(tt.want))
			}
		})
	}
}
