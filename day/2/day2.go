package main

import (
	"flag"
	"log/slog"
	"runtime"
	"time"

	"adventofcode2024.com/day/utils"
)

const DampenMax = 1

func IsGradualAndDirectional(Value1 int, Value2 int, Inc bool) bool {

	Gradual := false
	Directional := false
	GradualAndDirectional := false

	Difference := Value1 - Value2
	Change := utils.Abs(Difference)

	if Change > 0 && Change < 4 {
		Gradual = true

		if Inc && Difference < 0 || !Inc && Difference > 0 {
			Directional = true
		}
	}

	if Gradual && Directional {
		GradualAndDirectional = true
	}

	return GradualAndDirectional
}

func IsReportSafe(Report []int) bool {

	Inc := false
	ReportCount := len(Report) - 1

	if Report[0] < Report[ReportCount] {
		//increasing
		Inc = true
	}

	for i := 1; i < len(Report); i++ {

		LevelSafe := IsGradualAndDirectional(Report[i-1], Report[i], Inc)
		if !LevelSafe {
			return false
		}
	}

	return true

}

func DampenReport(Report []int) bool {

	for i := 0; i < len(Report); i++ {
		//test if the report can be made safe
		if IsReportSafe(utils.RemoveItemFromSlice(Report, i)) {
			return true
		}

	}

	return false
}

func main() {

	var Path string
	var Debug bool
	var Dampen bool
	flag.StringVar(&Path, "Path", "data/input", "OS Path to an input file")
	flag.BoolVar(&Debug, "Debug", false, "Turn on debugging messages")
	flag.BoolVar(&Dampen, "DampenReports", false, "Turn on problem dampener")
	flag.Parse()
	//set logging level based on user input
	if Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	pc, _, _, _ := runtime.Caller(1)
	defer utils.TimeTrack(time.Now(), runtime.FuncForPC(pc).Name())

	Reports := utils.ReadFileIntoIntMatrix(Path)
	SafeCount := 0

	for i := 0; i < len(Reports); i++ {
		Safe := IsReportSafe(Reports[i])

		if Safe {
			slog.Debug("Report being processed is safe")
			SafeCount++
		} else {
			if Dampen {
				if DampenReport(Reports[i]) {
					slog.Debug("After problem dampening Report being processed is safe")
					SafeCount++
				} else {
					slog.Debug("After problem dampening Report being processed is unsafe")
				}
			} else {
				slog.Debug("Dampening not in use. Report being processed is unsafe")
			}
		}
	}

	slog.Info("Completed Successfully", slog.Int("SafeReports", SafeCount))

}
