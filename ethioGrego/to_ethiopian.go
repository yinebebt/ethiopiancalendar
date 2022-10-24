package ethioGrego

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Ethiopian date string representation of provided Gregorian date
func To_ethiopian(year, month, date int) (time.Time, error) {
	var tahissas int
	var ethiopian_date int
	var dateResult string

	if !isValid(year, month, date) {
		return time.Time{}, errors.New("not a valid date")
	}

	//  date between 5 and 14 of May 1582 are invalid
	if month == 10 && date >= 5 && date <= 14 && year == 1582 {
		fmt.Printf("Invalid Date between 5-14 May 1582.")
		return time.Time{}, errors.New("not a valid date")
	}

	// Number of days in gregorian months starting with January (index 1)
	//	Index 0 is reserved for leap years switches.
	gregorian_months := []int{0, 31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

	ethiopian_months := []int{0, 30, 30, 30, 30, 30, 30, 30,
		30, 30, 5, 30, 30, 30, 30}

	//  if gregorian leap year, February has 29 days.
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		gregorian_months[2] = 29
	}
	//  September sees 8y difference
	ethiopian_year := year - 8

	//  if ethiopian leap year pagumain has 6 days
	if ethiopian_year%4 == 3 {
		ethiopian_months[10] = 6
	} else {
		ethiopian_months[10] = 5
	}
	//  Ethiopian new year in Gregorian calendar
	new_year_day := start_day_of_ethiopian(year - 8)

	// calculate number of days up to that date
	until := 0
	for i := 1; i < month; i++ {
		until += gregorian_months[i]
	}
	until += date

	// # update tahissas (december) to match january 1st
	if ethiopian_year%4 == 0 {
		tahissas = 26
	} else {
		tahissas = 25
	}

	// take into account the 1582 change
	if year < 1582 {
		ethiopian_months[1] = 0
		ethiopian_months[2] = tahissas
	} else if until <= 277 && year == 1582 {
		ethiopian_months[1] = 0
		ethiopian_months[2] = tahissas
	} else {
		tahissas = new_year_day - 3
		ethiopian_months[1] = tahissas
	}
	// calculate month and date incremently
	m := 0
	for m = range ethiopian_months {
		if until <= ethiopian_months[m] {
			if m == 1 || ethiopian_months[m] == 0 {
				ethiopian_date = until + (30 - tahissas)
			} else {
				ethiopian_date = until
			}
			break
		} else {
			until -= ethiopian_months[m]
		}
	}
	//  if m > 4, we're already on next Ethiopian year
	if m > 10 {
		ethiopian_year += 1
	}
	// Ethiopian months ordered according to Gregorian
	order := []int{0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2, 3, 4}
	ethiopian_month := order[m]

	da := strconv.Itoa(ethiopian_date)
	mon := strconv.Itoa(ethiopian_month)
	if len(da) == 1 {
		da = "0" + da
	}
	if len(mon) == 1 {
		mon = "0" + mon
	}

	dateResult = strconv.Itoa(ethiopian_year) + "-" + mon + "-" + da
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		fmt.Print("unabe to parse dateResult.", err) //for debugging purpose
		return time.Time{}, errors.New("not a valid date")
	}
	return res, nil
}
