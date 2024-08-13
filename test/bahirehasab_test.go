package test

import (
	"fmt"
	"gitlab.com/yinebebt/ethiopiancalendar/internal/module/bahirehasab"
	"testing"
)

func TestBahireHasab(t *testing.T) {
	res, err := bahirehasab.BahireHasab(2016)
	fmt.Print(res, err)
	if res.Year.Year != 2016 {
		t.Errorf("Test faled, expected %v but got %v ", 2016, res.Year.Year)
	}

	if err != nil {
		t.Errorf("Test faled, expected %v but got %v ", nil, err)
	}
}
