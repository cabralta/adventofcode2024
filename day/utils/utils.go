package utils

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

func Abs(Integer int) int {
	if Integer < 0 {
		return -Integer
	}

	return Integer
}

func StringsToInts(Input []string) []int {

	Ints := make([]int, 0)
	for _, StringToConvert := range Input {
		Int, _ := strconv.Atoi(StringToConvert)
		Ints = append(Ints, Int)
	}

	return Ints
}


func GetProductOfInts(IntArr []int) int {

	Product := 0

	for Idx, Int := range IntArr {

		if Idx == 0 {
			Product = Int
		} else {
			Product = Product * Int
		}

	}

	return Product
}

func ReadFileIntoString(Path string) []string {
	FileContents := make([]string, 0)
	file, err := os.Open(Path)
	if err != nil {
		slog.Error("Error opening file.", Path, slog.Any("error", err))
		panic(err)
	}
	defer file.Close()
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		FileContents = append(FileContents, Scanner.Text())

	}

	return FileContents
}

func ReadFileIntoIntMatrix(Path string) [][]int {

	IntMatrix := make([][]int, 0)
	file, err := os.Open(Path)
	if err != nil {
		slog.Error("Error opening file.", Path, slog.Any("error", err))
		panic(err)
	}
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	for Scanner.Scan() {
		Row := Scanner.Text()
		Ints := make([]int, 0)
		Fields := strings.Fields(Row)
		for i := range Fields {

			IntValue, err := strconv.Atoi(Fields[i])
			if err != nil {
				slog.Error("Type Conversion Failed. Invalid Record Value.", slog.Any("error", err))
				panic(err)
			}
			Ints = append(Ints, IntValue)
		}
		IntMatrix = append(IntMatrix, Ints)

	}

	return IntMatrix
}

func RemoveItemFromSlice(Slice []int, S int) []int {
	//We don't want to modify the existing slice but provide a new slice with the missing value
	NewSlice := make([]int, 0)
	NewSlice = append(NewSlice, Slice[:S]...)
	return append(NewSlice, Slice[S+1:]...)
}

func TimeTrack(Start time.Time, MethodName string) {
	Elapsed := time.Since(Start)
	slog.Info("Execution duration", MethodName, Elapsed)
}
