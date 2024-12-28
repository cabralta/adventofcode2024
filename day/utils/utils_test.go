package utils

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {

	var tests = []struct {
		name   string
		input1 int
		want   int
	}{
		{"Positive Value", 10, 10},
		{"Negative Value", -10, 10},
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

func TestStringToInt(t *testing.T) {

	var tests = []struct {
		name              string
		StringArrayToTest []string
		want              []int
	}{
		{"Conversion of Strings to Ints", []string{"0", "1", "2"}, []int{0, 1, 2}},
		{"Conversion expecting 0 for non-int values", []string{"error", "1.123", "-1.0"}, []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ans := StringsToInts(tt.StringArrayToTest)

			for i, _ := range ans {
				if ans[i] != tt.want[i] {
					t.Errorf("Slice values don't match, got %s, want %s", fmt.Sprint(ans), fmt.Sprint(tt.want))
				}
			}
		})
	}

}

func TestProductOfInts(t *testing.T) {
	var tests = []struct {
		name           string
		IntSliceToTest []int
		want           int
	}{
		{"Product of 1*2*3 is 6", []int{1, 2, 3}, 6},
		{"Product of 1*1*1 is 1", []int{1, 1, 1}, 1},
		{"Product of -1*2*3 is -6", []int{-1, 2, 3}, -6},
		{"Product of 0*1*1 is 0", []int{0, 1, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := GetProductOfInts(tt.IntSliceToTest)
			if ans != tt.want {
				t.Errorf("Product doesn't match expected result. Got %s, want %s", fmt.Sprint(ans), fmt.Sprint(tt.want))
			}
		})
	}
}

func TestReadFileIntoStringPanic(t *testing.T) {
	TestFilePath := "/tmp/file_that_should_not_exist"
	defer func() {
		recover()
	}()

	ReadFileIntoString(TestFilePath)

	t.Errorf("Invalid file should cause panic")
}

func TestReadFileIntoString(t *testing.T) {

	TestFilePath := "test_read_file_into_string_file"
	BytesToWrite := []byte("Test Line 1\nTest Line 2")
	Want := strings.Split(string(BytesToWrite[:]), "\n")

	err := os.WriteFile(TestFilePath, BytesToWrite, 0644)
	if err != nil {
		panic(err)
	}

	t.Run("Validate File Read Returns String Slice", func(t *testing.T) {
		ans := ReadFileIntoString(TestFilePath)

		if len(ans) == len(Want) {
			for i, _ := range ans {
				if ans[i] != Want[i] {
					t.Errorf("Strings don't match. Got %s, Want %s", fmt.Sprint(ans[i]), fmt.Sprint(Want[i]))
				}
			}
		} else {
			t.Errorf("Slice Lengths don't match. File was not parsed correctly. Got %s, Want %s", fmt.Sprint(len(ans)), fmt.Sprint(len(Want)))
		}
	})

	err = os.Remove(TestFilePath)
	if err != nil {
		panic(err)
	}

}

func TestReadFileIntoIntMatrixPanic(t *testing.T) {
	TestFilePath := "/tmp/file_that_should_not_exist"
	defer func() {
		recover()
	}()

	ReadFileIntoIntMatrix(TestFilePath)

	t.Errorf("Invalid file should cause panic")
}

func TestReadFileIntoIntMatrix(t *testing.T) {

	//Test can be improved to validate many different types of white space

	var tests = []struct {
		name         string
		TestFilePath string
		BytesToWrite []byte
		ExpectPanic  bool
		Want         [][]int
	}{
		{"Clean Conversion", "test_read_file_into_int_matrix_file", []byte("1 2\n3 4\n5 6"), false, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"Panic Conversion", "test_read_file_into_int_matrix_file", []byte("1 2\n3 4\n5 T"), true, [][]int{{1, 2}, {3, 4}, {5, 0}}},
	}

	for _, tt := range tests {

		err := os.WriteFile(tt.TestFilePath, tt.BytesToWrite, 0644)
		if err != nil {
			panic(err)
		}

		t.Run("Validate File Read Returns Slice of Slice with the same values", func(t *testing.T) {

			defer func() { recover() }()
			ans := ReadFileIntoIntMatrix(tt.TestFilePath)
			if tt.ExpectPanic {
				t.Errorf("Conversion should have caused panic")
			}

			if len(ans) != len(tt.Want) {
				t.Errorf("Slice lengths don't match. File was not parsed correctly. Got %s, Want %s", fmt.Sprint(len(ans)), fmt.Sprint(len(tt.Want)))
			} else {

				for i, ints := range ans {

					if len(ints) != len(tt.Want[i]) {

						t.Errorf("Slice item lengths don't match. File was not parsed correctly. Got %s, Want %s", fmt.Sprint(len(ans[i])), fmt.Sprint(len(tt.Want[i])))
					} else {

						for j, _ := range ints {

							if ans[i][j] != tt.Want[i][j] {
								t.Errorf("Values don't match in matrix. Got %s, Want %s", fmt.Sprint(ans[i][j]), fmt.Sprint(tt.Want[i][j]))
							}
						}
					}
				}
			}
		})

		err = os.Remove(tt.TestFilePath)
		if err != nil {
			panic(err)
		}
	}

}

func TestRemoveItemFromSlice(t *testing.T) {

	Ints := []int{0, 1, 2, 3, 4}
	IndexToRemove := 2
	Want := []int{0, 1, 3, 4}

	t.Run("Remove Item from Slice", func(t *testing.T) {

		ans := RemoveItemFromSlice(Ints, IndexToRemove)

		for i, _ := range ans {
			if ans[i] != Want[i] {
				t.Errorf("Slice values don't match, got %s, want %s", fmt.Sprint(ans), fmt.Sprint(Want))
			}
		}

	})

}

func TestTimeTrack(t *testing.T) {

	Elapsed := TimeTrack(time.Now(), "Test Method")

	if Elapsed == 0 {
		t.Errorf("Time Duration Not Recorded. Got %s", fmt.Sprint(Elapsed))
	}

}
