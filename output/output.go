package output

import (
	"fmt"
	"sort"
)

func FormatPercentages(percentages map[string]float64, targetTag string, average float64) []string {
	var output = make([]string, 0)
	var sortedDates = make([]string, 0, len(percentages))
	for key, _ := range percentages {
		sortedDates = append(sortedDates, key)
	}
	sort.Strings(sortedDates)
	for _, date := range sortedDates {
		output = append(output, fmt.Sprintf("%s - %s: %.2f%%", date, targetTag, percentages[date]))
	}
	output = append(output, fmt.Sprintf("Average percentage: %.2f%%", average))
	return output
}
