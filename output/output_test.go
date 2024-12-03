package output

import "testing"

func TestFormatPercentages(t *testing.T) {
	var percentages = map[string]float64{
		"02-12-2024": 36.74589,
	}

	formatted := FormatPercentages(percentages, "target", 47.231458)
	expectedLineFormat := "02-12-2024 - target: 36.75%"
	if formatted[0] != expectedLineFormat {
		t.Errorf("Incorrected formatting for single row. Expected '%s', got '%s'", expectedLineFormat, formatted[0])
	}
	expectedAverageFormat := "Average percentage: 47.23%"
	if formatted[1] != expectedAverageFormat {
		t.Errorf("Incorrected formatting for average row. Expected '%s', got '%s'", expectedLineFormat, formatted[0])
	}

}
