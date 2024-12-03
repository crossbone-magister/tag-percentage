package logic

import (
	"fmt"
	"github.com/crossbone-magister/timewlib"
	"slices"
	"time"
)

func CalculateTagPercentage(intervals []timewlib.Interval, targetTag string) (map[string]float64, float64) {
	totalPerDay := make(map[string]time.Duration)
	totalTargetTagPerDay := make(map[string]time.Duration)
	for _, interval := range intervals {
		y, m, d := interval.StartDate()
		key := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		if _, ok := totalPerDay[key]; !ok {
			totalPerDay[key] = 0
		}
		totalPerDay[key] += interval.Duration()
		if slices.Contains(interval.Tags, targetTag) {
			totalTargetTagPerDay[key] += interval.Duration()
		}
	}
	var percentages = make(map[string]float64)
	var average = 0.0
	for day, total := range totalPerDay {
		percentage := (totalTargetTagPerDay[day].Seconds() / total.Seconds()) * 100
		percentages[day] = percentage
		average += percentage
	}
	average /= float64(len(totalPerDay))
	return percentages, average
}
