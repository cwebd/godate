package date

import (
	"testing"
	"time"
)

func TestDates(t *testing.T) {
	// Invalid dates.
	tests := []string{
		"",
		"HI",
		"20938258384738",
		"23.44",
	}
	for _, date := range tests {
		if d, err := Parse(date); err == nil {
			t.Errorf("Input %q was parsed to a date %v", date, d)
		}
	}

	// Full date, no ambiguity.
	correct := time.Date(1989, 6, 25, 0, 0, 0, 0, time.UTC)
	tests = []string{
		"25 Jun 1989",
		"Jun 25 1989",
		"Jun 25 '89",
		"June 25, 1989",
		"June 25, '89",
		"25 Jun 1989",
		"25 Jun '89",

		"6/25/89",
		"06/25/89",
		"6/25/1989",
		"06/25/1989",
		"25/6/89",
		"25/6/1989",
		"25/06/89",
		"25/06/1989",

		"6-25-89",
		"06-25-89",
		"6-25-1989",
		"06-25-1989",
		"25-6-89",
		"25-6-1989",
		"25-06-89",
		"25-06-1989",

		"19890625",
		"890625",

		"Sunday 25th June 1989",
	}

	for _, date := range tests {
		d, err := Parse(date)
		if err != nil {
			t.Errorf("Error parsing %q: %v", date, err)
		} else if d != correct {
			t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
		}
	}

	// Testing Other Date Ordinals

	// Test st date ordinal
	correct = time.Date(1989, 8, 1, 0, 0, 0, 0, time.UTC)
	date := "Tuesday 1st August 1989"
	d, err := Parse(date)
	if err != nil {
		t.Errorf("Error parsing %q: %v", date, err)
	} else if d != correct {
		t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
	}

	// Test nd date ordinal
	correct = time.Date(1989, 8, 2, 0, 0, 0, 0, time.UTC)
	date = "Wednesday 2nd August 1989"
	d, err = Parse(date)
	if err != nil {
		t.Errorf("Error parsing %q: %v", date, err)
	} else if d != correct {
		t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
	}

	// Test rd date ordinal
	correct = time.Date(1989, 8, 3, 0, 0, 0, 0, time.UTC)
	date = "Wednesday 3rd August 1989"
	d, err = Parse(date)
	if err != nil {
		t.Errorf("Error parsing %q: %v", date, err)
	} else if d != correct {
		t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
	}

	// Test th date ordinal
	correct = time.Date(1989, 8, 4, 0, 0, 0, 0, time.UTC)
	date = "Wednesday 4th August 1989"
	d, err = Parse(date)
	if err != nil {
		t.Errorf("Error parsing %q: %v", date, err)
	} else if d != correct {
		t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
	}

	// Just month and year
	correct = time.Date(1989, 6, 1, 0, 0, 0, 0, time.UTC)
	tests = []string{
		"Jun 1989",
		"Jun-1989",
		"Jun '89",
		"June 1989",
		"June '89",
		"6/89",
		"06/89",
		"6-89",
		"06-89",
		"198906",
		"8906",
		"0689",
	}

	for _, date := range tests {
		d, err := Parse(date)
		if err != nil {
			t.Errorf("Error prasing %s: %v", date, err)
		} else if d != correct {
			t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
		}
	}

	// Just month and day
	correct = time.Date(0, 6, 25, 0, 0, 0, 0, time.UTC)
	tests = []string{
		"Jun 25",
		"Jun-25",
		"June 25",
		"6/25",
		"06/25",
		"6-25",
	}

	for _, date := range tests {
		d, err := Parse(date)
		if err != nil {
			t.Errorf("Error prasing %q: %v", date, err)
		} else if d != correct {
			t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
		}
	}

	// Just month and year, ambiguous
	correct = time.Date(2009, 6, 1, 0, 0, 0, 0, time.UTC)
	tests = []string{
		"Jun 2009",
		"Jun-2009",
		"June 2009",
		"200906",
	}

	for _, date := range tests {
		d, err := Parse(date)
		if err != nil {
			t.Errorf("Error prasing %q: %v", date, err)
		} else if d != correct {
			t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
		}
	}

	// Full date with time at minute resolution.
	correct = time.Date(1989, 6, 25, 15, 45, 0, 0, time.UTC)
	tests = []string{
		"Jun 25 1989 15:45",
		"Jun 25 1989 3:45PM",
		"Jun 25 1989 03:45PM",
		"6/25/1989 15:45",
		"6/25/1989 03:45PM",
		"6/25/1989 3:45PM",
	}

	for _, date := range tests {
		d, err := Parse(date)
		if err != nil {
			t.Errorf("Error prasing %q: %v", date, err)
		} else if d != correct {
			t.Errorf("Dates for %q did not match:\n%v (actual)\n%v (expected)", date, d, correct)
		}
	}

	// TODO: Make more test cases for time of day, timezones.
}

func TestConvert(t *testing.T) {
	// TODO: Make test cases for time of day, timezones.
	tests := map[string]string{
		"Jan 02 2006":      "MMM dd yyyy",
		"January 02, 2006": "MMMM dd, yyyy",
		"01/02/2006":       "MM/dd/yyyy",
		"_2/1/06":          "d/M/yy",
		"01-2006":          "MM-yyyy",
		"Jan-2006":         "MMM-yyyy",
		"Mon 02 Jan '06":   "EEE dd MMM 'yy",
	}
	for f, correct := range tests {
		if uf := ConvertGoLayoutToUnicode(f); uf != correct {
			t.Errorf("Failed to convert %q to %q, got %q instead", f, correct, uf)
		}
	}
}
