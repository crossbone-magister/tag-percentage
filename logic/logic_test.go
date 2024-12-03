package logic

import (
	"github.com/crossbone-magister/timewlib"
	"testing"
	"time"
)

func TestCalculateTagPercentage(t *testing.T) {
	var interval1 = timewlib.NewInterval(10, 0, 11, 0)
	interval1.Tags = []string{"target"}
	var interval2 = timewlib.NewInterval(11, 0, 12, 0)
	percentagesPerDay, averagePercentage := CalculateTagPercentage([]timewlib.Interval{
		*interval1,
		*interval2,
	}, "target")
	var label = time.Now().Format("2006-01-02")
	if len(percentagesPerDay) <= 0 {
		t.Errorf("PercentagesPerDay is empty")
	}
	if _, ok := percentagesPerDay[label]; !ok {
		t.Errorf("PercentagesPerDay doesn't contain day %s", label)
	}
	if percentagesPerDay[label] != 50.0 {
		t.Errorf("PercentagesPerDay for day %s is not 0.5 but %f", label, percentagesPerDay[label])
	}
	if averagePercentage != 50.0 {
		t.Errorf("AveragePercentage is less than 0.5")
	}
}
