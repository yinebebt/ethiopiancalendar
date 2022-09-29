package ethioGrego

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Gregorian date object representation of provided Ethiopian date
func To_gregorian(year, month, date int) (time.Time, error) {
	var gregorian_date int
	var dateResult string

	if !isValid(year, month, date) {
		return time.Time{}, errors.New("not a valid date")
	}
	// Ethiopian new year in Gregorian calendar
	new_year_day := start_day_of_ethiopian(year)

	// September (Ethiopian) sees 7y difference
	gregorian_year := year + 7

	// Number of days in gregorian months starting with September (index 1)
	// Index 0 is reserved for leap years switches.
	gregorian_months := []int{0, 30, 31, 30, 31, 31, 28,
		31, 30, 31, 30, 31, 31, 30}

	//if next gregorian year is leap year, February has 29 days.
	next_year := gregorian_year + 1
	if (next_year%4 == 0 && next_year%100 != 0) || next_year%400 == 0 {
		gregorian_months[6] = 29
	}
	// calculate number of days up to that date
	until := ((month - 1) * 30) + date
	if until <= 37 && year <= 1575 { //mysterious rule
		until += 28
		gregorian_months[0] = 31
	} else {
		until += new_year_day - 1
	}

	// if ethiopian year is leap year, paguemain has six days
	if year-1%4 == 3 {
		until += 1
	}

	//calculate month and date incremently
	m := 0
	for i := range gregorian_months {
		if until <= gregorian_months[i] {
			m = i
			gregorian_date = until
			break
		} else {
			m = i
			until -= gregorian_months[i]
		}
	}
	//  if m > 4, we're already on next Gregorian year
	if m > 4 {
		gregorian_year += 1
	}
	//  Gregorian months ordered according to Ethiopian
	order := []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	gregorian_month := order[m]

	da := strconv.Itoa(gregorian_date)
	mon := strconv.Itoa(gregorian_month)
	if len(da) == 1 {
		da = "0" + da
	}
	if len(mon) == 1 {
		mon = "0" + mon
	}

	dateResult = "" + strconv.Itoa(gregorian_year) + "-" + mon + "-" + da
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		fmt.Print("unabe to parse dateResult.", err)
		return time.Time{}, errors.New("not a valid date")

	}
	return res, nil
}
