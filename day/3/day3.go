package main

import (
	"flag"
	"log/slog"
	"regexp"
	"runtime"
	"strings"
	"time"

	"adventofcode2024.com/day/utils"
)

func RexString(Input string, Rex string) []string {

	rex, _ := regexp.Compile(Rex)
	return rex.FindAllString(Input, -1)

}

func GetProductSum(ExtractedFunctions []string) int {

	EnableMul := true
	Sum := 0
	rex, _ := regexp.Compile("[0-9]?[0-9]?[0-9]")

	for _, Function := range ExtractedFunctions {

		if Function == "don't()" {
			EnableMul = false
		} else if Function == "do()" {
			EnableMul = true
		}

		if EnableMul && strings.HasPrefix(Function, "mul") {

			InputStrings := rex.FindAllString(Function, -1)
			Ints := utils.StringsToInts(InputStrings)
			Product := utils.GetProductOfInts(Ints)
			Sum = Sum + Product
		}

	}

	return Sum
}

func FixCorruptedMem(RawInput []string, HasCommands bool) int {

	var Exp string

	if HasCommands {
		Exp = "mul\\([0-9]?[0-9]?[0-9]?,[0-9]?[0-9]?[0-9]?\\)|do\\(\\)|don\\'t\\(\\)"
	} else {
		Exp = "mul\\([0-9]?[0-9]?[0-9]?,[0-9]?[0-9]?[0-9]?\\)"
	}

	Input := strings.Join(RawInput, " ")

	ExtractedString := RexString(Input, Exp)

	return GetProductSum(ExtractedString)
}

func main() {

	var Path string
	var Debug bool
	var HasCommands bool

	flag.StringVar(&Path, "Path", "testdata/input_test", "OS Path to an input file")
	flag.BoolVar(&Debug, "Debug", false, "Turn on debugging messages")
	flag.BoolVar(&HasCommands, "HasCommands", false, "Detect do() | don't() Commands")
	flag.Parse()
	//set logging level based on user input
	if Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	pc, _, _, _ := runtime.Caller(1)
	defer utils.TimeTrack(time.Now(), runtime.FuncForPC(pc).Name())

	StringsToParse := utils.ReadFileIntoString(Path)
	SumAll := FixCorruptedMem(StringsToParse, HasCommands)
	slog.Info("Completed Successfully", slog.Int("SumAllMatched", SumAll))

}
