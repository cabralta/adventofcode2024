package utils

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Calculations
func Abs(Integer int) int {
	if Integer < 0 {
		return -Integer
	}

	return Integer
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

// Slice Functions

func StringsToInts(Input []string) []int {

	Ints := make([]int, 0)
	for _, StringToConvert := range Input {
		Int, _ := strconv.Atoi(StringToConvert)
		Ints = append(Ints, Int)
	}

	return Ints
}

func RemoveItemFromSlice(Slice []int, S int) []int {
	//We don't want to modify the existing slice but provide a new slice with the missing value
	NewSlice := make([]int, 0)
	NewSlice = append(NewSlice, Slice[:S]...)
	return append(NewSlice, Slice[S+1:]...)
}

func GetReverseStringSlice(Input []string) []string {
	ReverseSlice := make([]string, 0)

	for _, Str := range Input {
		ReverseSlice = append(ReverseSlice, ReverseString(Str))
	}

	return ReverseSlice
}

func GetVerticalStringSlice(Input []string) []string {

	VerticalSlice := make([]string, len(Input))

	for _, s := range Input {
		for j, AsciiRune := range s {
			str := string(AsciiRune)
			VerticalSlice[j] += str
		}
	}

	return VerticalSlice
}

func GetDiagonalStringSlice(Input []string) []string {

	DiagonalSlice := make([]string, ((2 * len(Input)) - 1))
	//MaxH := len(Input)-1
	MaxL := len(Input[0]) - 1

	//CalculateVectorDifference
	//MaxL + VectorDiff = DiagonalSlice[i]+=CurrentVector

	for i, line := range Input {

		for j, char := range line {
			VectorDiff := i - j
			DiagonalSliceIdx := MaxL - VectorDiff
			DiagonalSlice[DiagonalSliceIdx] += string(char)
		}
	}

	return DiagonalSlice
}

// String Functions
func RexString(Input string, Rex string) []string {

	rex, _ := regexp.Compile(Rex)
	return rex.FindAllString(Input, -1)

}

func ReverseString(Input string) string {

	ReverseString := make([]byte, len(Input))
	StrLen := len(Input) - 1

	for i, value := range Input {
		ReverseString[StrLen-i] = byte(value)
	}

	return string(ReverseString)

}

func ConvertStringToStringSlice(Input string) []string {
	StringSlice := make([]string, 0)

	for _, char := range Input {
		StringSlice = append(StringSlice, string(char))
	}
	return StringSlice
}

// File Operations
func ReadFileIntoStringSlice(Path string) [][]string {
	FileContents := make([][]string, 0)
	file, err := os.Open(Path)
	if err != nil {
		slog.Error("Error opening file.", Path, slog.Any("error", err))
		panic(err)
	}
	defer file.Close()
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		line := Scanner.Text()
		FileContents = append(FileContents, ConvertStringToStringSlice(line))

	}

	return FileContents
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

// General

func TimeTrack(Start time.Time, MethodName string) time.Duration {
	Elapsed := time.Since(Start)
	slog.Info("Execution duration", MethodName, Elapsed)

	return Elapsed
}
