package ethioGrego

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ToGregorian gives Gregorian date object representation of provided Ethiopian date
func ToGregorian(year, month, date int) (time.Time, error) {
	var gregorianDate int
	var dateResult string

	if !isValid(year, month, date) {
		return time.Time{}, errors.New("not a valid date")
	}
	// Ethiopian new year in Gregorian calendar
	newYearDay := startDayOfEthiopian(year)

	// September (Ethiopian) sees 7y difference
	gregorianYear := year + 7

	// Number of days in gregorian months starting with September (index 1)
	// Index 0 is reserved for leap years switches.
	gregorianMonths := []int{0, 30, 31, 30, 31, 31, 28,
		31, 30, 31, 30, 31, 31, 30}

	//if next gregorian year is leap year, February has 29 days.
	nextYear := gregorianYear + 1
	if (nextYear%4 == 0 && nextYear%100 != 0) || nextYear%400 == 0 {
		gregorianMonths[6] = 29
	}
	// calculate number of days up to that date
	until := ((month - 1) * 30) + date
	if until <= 37 && year <= 1575 { //mysterious rule
		until += 28
		gregorianMonths[0] = 31
	} else {
		until += newYearDay - 1
	}

	// if ethiopian year is leap year, paguemain has six days
	if year-1%4 == 3 {
		until += 1
	}

	//calculate month and date incrementally
	m := 0
	for i := range gregorianMonths {
		if until <= gregorianMonths[i] {
			m = i
			gregorianDate = until
			break
		} else {
			m = i
			until -= gregorianMonths[i]
		}
	}
	//  if m > 4, we're already on next Gregorian year
	if m > 4 {
		gregorianYear += 1
	}
	//  Gregorian months ordered according to Ethiopian
	order := []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	gregorianMonth := order[m]

	da := strconv.Itoa(gregorianDate)
	mon := strconv.Itoa(gregorianMonth)
	if len(da) == 1 {
		da = "0" + da
	}
	if len(mon) == 1 {
		mon = "0" + mon
	}

	dateResult = "" + strconv.Itoa(gregorianYear) + "-" + mon + "-" + da
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		fmt.Print("unable to parse dateResult.", err)
		return time.Time{}, errors.New("not a valid date")
	}
	return res, nil
}
