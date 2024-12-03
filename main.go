package main

import (
	"fmt"
	"github.com/crossbone-magister/timewlib"
	"os"
	"tag-percentage/logic"
	"tag-percentage/output"
)

const TargetTagConfig = "reports.tagpercentage.target"

func main() {
	parsed, err := timewlib.Parse(os.Stdin)
	if err == nil {
		var config timewlib.Configuration = parsed.Configuration
		var targetTag = config[TargetTagConfig]
		if targetTag != "" {
			intervals, err := timewlib.Process(parsed.Intervals)
			if err == nil {
				if len(intervals) > 0 {
					percentages, average := logic.CalculateTagPercentage(intervals, targetTag)
					for _, row := range output.FormatPercentages(percentages, targetTag, average) {
						fmt.Println(row)
					}
				} else {
					fmt.Println(timewlib.GenerateNoDataMessage(config))
				}
			} else {
				printErrorAndExit(err)
			}
		} else {
			fmt.Println("No target tag specified")
			os.Exit(1)
		}
	} else {
		printErrorAndExit(err)
	}

}

func printErrorAndExit(err error) {
	fmt.Printf("Error while reading timewarrior input: %s\n", err)
	os.Exit(1)
}
