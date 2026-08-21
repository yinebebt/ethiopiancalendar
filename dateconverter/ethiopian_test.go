package dateconverter

import (
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	pairs := []struct{ gy, gm, gd, ey, em, ed int }{
		{2022, 9, 28, 2015, 1, 18},
		{2022, 9, 11, 2015, 1, 1},
		{2025, 2, 2, 2017, 5, 25},
		{1582, 10, 4, 1575, 2, 7},
		{1582, 10, 15, 1575, 2, 8},
	}
	for _, tt := range pairs {
		g := time.Date(tt.gy, time.Month(tt.gm), tt.gd, 0, 0, 0, 0, time.UTC)
		e := time.Date(tt.ey, time.Month(tt.em), tt.ed, 0, 0, 0, 0, time.UTC)
		t.Run(g.Format("2006-01-02"), func(t *testing.T) {
			got, err := Ethiopian(tt.gy, tt.gm, tt.gd)
			if err != nil {
				t.Fatal(err)
			}
			if !got.Equal(e) {
				t.Errorf("Ethiopian: got %v, want %v", got, e)
			}

			got, err = Gregorian(tt.ey, tt.em, tt.ed)
			if err != nil {
				t.Fatal(err)
			}
			if !got.Equal(g) {
				t.Errorf("Gregorian: got %v, want %v", got, g)
			}
		})
	}
}

func TestGregorian_pagume(t *testing.T) {
	t.Parallel()

	got, err := Gregorian(2015, 13, 6)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(time.Date(2023, 9, 11, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("leap pagume 6: got %v", got)
	}

	got, err = Gregorian(2016, 13, 5)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(time.Date(2024, 9, 11, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("non-leap pagume 5: got %v", got)
	}
}

func TestEthiopian_leapDay(t *testing.T) {
	t.Parallel()
	got, err := Ethiopian(2024, 2, 29)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(time.Date(2016, 6, 21, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("got %v, want 2016-06-21", got)
	}
}

func TestInvalid(t *testing.T) {
	t.Parallel()
	if _, err := Ethiopian(0, 1, 1); err == nil {
		t.Error("year 0")
	}
	if _, err := Ethiopian(1582, 10, 10); err == nil {
		t.Error("reform gap")
	}
	if _, err := Gregorian(2015, 13, 7); err == nil {
		t.Error("pagume 7")
	}
	if _, err := Gregorian(2016, 13, 6); err == nil {
		t.Error("pagume 6 non-leap")
	}
}
