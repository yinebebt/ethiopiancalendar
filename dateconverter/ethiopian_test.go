package dateconverter

import (
	"testing"
	"time"
)

func TestEthiopian(t *testing.T) {
	got, err := Ethiopian(2022, 9, 28)
	if err != nil {
		t.Fatalf("Ethiopian: %v", err)
	}
	want := time.Date(2015, 1, 18, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("Ethiopian(2022, 9, 28) = %v, want %v", got, want)
	}
}

func TestGregorian(t *testing.T) {
	got, err := Gregorian(2015, 1, 18)
	if err != nil {
		t.Fatalf("Gregorian: %v", err)
	}
	want := time.Date(2022, 9, 28, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("Gregorian(2015, 1, 18) = %v, want %v", got, want)
	}
}

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name       string
		gy, gm, gd int
	}{
		{"meskerem", 2022, 9, 28},
		{"new_year", 2024, 9, 11},
		{"mid_year", 2025, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			et, err := Ethiopian(tt.gy, tt.gm, tt.gd)
			if err != nil {
				t.Fatalf("Ethiopian: %v", err)
			}
			back, err := Gregorian(et.Year(), int(et.Month()), et.Day())
			if err != nil {
				t.Fatalf("Gregorian: %v", err)
			}
			want := time.Date(tt.gy, time.Month(tt.gm), tt.gd, 0, 0, 0, 0, time.UTC)
			if !back.Equal(want) {
				t.Errorf("round-trip: got %v want %v (via ET %v)", back, want, et.Format("2006-01-02"))
			}
		})
	}
}

func TestEthiopian_invalid(t *testing.T) {
	if _, err := Ethiopian(0, 1, 1); err == nil {
		t.Fatal("expected error for invalid gregorian date")
	}
	if _, err := Ethiopian(1582, 10, 10); err == nil {
		t.Fatal("expected error for Gregorian reform gap")
	}
}

func TestGregorian_invalid(t *testing.T) {
	if _, err := Gregorian(2015, 13, 7); err == nil {
		t.Fatal("expected error for invalid pagumen day")
	}
	if _, err := Gregorian(0, 1, 1); err == nil {
		t.Fatal("expected error for invalid ethiopian date")
	}
}
