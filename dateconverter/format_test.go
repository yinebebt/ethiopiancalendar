package dateconverter

import (
	"testing"
	"time"
)

func TestFormatGregorian(t *testing.T) {
	d := time.Date(2012, time.July, 2, 0, 0, 0, 0, time.UTC)
	got := FormatGregorian(d)
	want := "Monday, July 2, 2012"
	if got != want {
		t.Errorf("FormatGregorian: got %q want %q", got, want)
	}
}

func TestFormatEthiopian(t *testing.T) {
	// 2015-01-18 E.C. == 2022-09-28 G.C. (Wednesday)
	got, err := FormatEthiopian(2015, 1, 18)
	if err != nil {
		t.Fatalf("FormatEthiopian: %v", err)
	}
	want := "መስከረም 18፣ 2015 (ረቡዕ)"
	if got != want {
		t.Errorf("FormatEthiopian: got %q want %q", got, want)
	}

	et, err := Ethiopian(2022, 9, 28)
	if err != nil {
		t.Fatalf("Ethiopian: %v", err)
	}
	got, err = FormatEthiopian(et.Year(), int(et.Month()), et.Day())
	if err != nil {
		t.Fatalf("FormatEthiopian from conversion: %v", err)
	}
	if got != want {
		t.Errorf("FormatEthiopian from conversion: got %q want %q", got, want)
	}
}
